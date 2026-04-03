package gtpv1u

import (
	"bytes"
	"encoding/binary"
	"lite5gc/cmn/rlogger"
)

// 3GPP TS 29.281 V15.5.0 (2018-12)
type N3MsgParser struct {
	Msgbuf           []byte
	GtpuMsg          *GTPPDU
	GPDUSessionUL    *GPDUSessionContUL // UL PDU SESSION INFORMATION
	EchoRequestMsg   *EchoRequest
	EchoResponseMsg  *EchoResponse
	ErrIndicationMsg *ErrIndication
}

// GTP-PDU协议消息
type GTPPDU struct {
	Gtpv1uHeader
	Gtpbody []byte
}

// G-PDU消息,消息类型：Msg_Type_G_PDU
type GPDU struct {
	Gtpv1uHeader
	TPDU []byte
}

// G-PDU消息,消息类型：Msg_Type_G_PDU,带扩展容器
// RAN --> UPF的上行消息
type GPDUSessionContUL struct {
	Gtpv1uHeader
	PDUSessionContainer PDUSessionContainerUL
	//TPDU []byte
	Gtpbody []byte
}

// G-PDU消息,消息类型：Msg_Type_G_PDU,带扩展容器
// UPF--> NG-RAN的下行消息
type GPDUSessionContDL struct {
	Gtpv1uHeader
	PDUSessionContainer PDUSessionContainerDL
	Gtpbody             []byte
}

// G-SIG消息
type GSigMsg struct {
}

// Path Management messages
type GPathMsg struct {
}

// Tunnel Management messages
type GTunMsg struct {
}

func init() {
}

// N3接口收UDP消息，解码处理

// 解码GTPv1-U消息
func (m *N3MsgParser) DecodeMsg() error {
	FuncEntry(moduleTag, nil)
	// the GTP-U Header
	// Always present fields
	r, payloadLen, err := m.mHeaderDecode()
	if err != nil {
		//fmt.Println(err)
		rlogger.Trace(moduleTag, rlogger.WARN, nil, " Always present fields error: %s", err)
		return err
	}
	//Optional fields,payloadLen 是去掉必选头之后的长度
	// 下一步可优化入参为m.oHeaderDecode(payloadLen)，保证r的健壮
	err = m.oHeaderDecode(r, payloadLen)
	if err != nil {
		//fmt.Println(err)
		rlogger.Trace(moduleTag, rlogger.WARN, nil, " Optional fields error: %s ", err)
		return err // 可选头异常放通
	}

	// GTP消息体
	//n, _ := r.Read(GtpMsg.Gtpbody)
	tPayload := r.Len()
	//fmt.Println("\nread paylaod length:", tPayload)
	//rlogger.Trace(moduleTag, rlogger.INFO, nil,  "\nread paylaod length:%d", tPayload)

	// read payload
	if tPayload != 0 && m.GtpuMsg.MessageType == Msg_Type_G_PDU {
		m.GtpuMsg.Gtpbody = make([]byte, r.Len())
		n, err := r.ReadAt(m.GtpuMsg.Gtpbody, int64(len(m.Msgbuf)-r.Len()))
		if err != nil && n == 0 { // no carry paylaod
			//fmt.Println("ReadAt Fail")
			rlogger.Trace(moduleTag, rlogger.WARN, nil, " ReadAt Fail %s", err)
			m.GtpuMsg.Gtpbody = []byte{}
			return ErrGTPMsgLen
		}

		//fmt.Printf("  read paylaod :  %#x\n", m.GtpuMsg.Gtpbody[:n])
		//rlogger.Trace(moduleTag, rlogger.INFO, nil,  "  read paylaod :  %#x\n", m.GtpuMsg.Gtpbody[:n])
	}

	return nil
}

// oHeaderDecode Optional fields decode
func (m *N3MsgParser) oHeaderDecode(r *bytes.Reader, payloadLen int) error {
	FuncEntry(moduleTag, nil)
	//r := bytes.NewReader(m.Msgbuf[payloadLen:]) // 优化可以使用
	//安全检查
	if r == nil {
		return ErrNil
	}
	if //(payloadLen > GTPV1_U_T_PDU) || //test 放通大包
	payloadLen < 0 {
		return ErrInvalidIELen
	}

	if m.GtpuMsg.EFlag == Protocol_Present ||
		m.GtpuMsg.SFlag == Protocol_Present ||
		m.GtpuMsg.PNFlag == Protocol_Present {
		binary.Read(r, binary.BigEndian, &m.GtpuMsg.SequenceNumber)
		binary.Read(r, binary.BigEndian, &m.GtpuMsg.NPDUNumber)
		binary.Read(r, binary.BigEndian, &m.GtpuMsg.NextExtHeaderType)
		// Optional fields S,PN,E的长度，默认值是4bytes
		oFieldLen := payloadLen - r.Len()
		//fmt.Println("\nRead Optional fields :", oFieldLen)
		//fmt.Println("  9-10 Octet SequenceNumber:", m.GtpuMsg.SequenceNumber)
		//fmt.Println("  11 Octet NPDUNumber:", m.GtpuMsg.NPDUNumber)
		//fmt.Println("  12 Octet NextExtHeaderType:", m.GtpuMsg.NextExtHeaderType)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "\nRead Optional fields :%d", oFieldLen)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  9-10 Octet SequenceNumber:%d", m.GtpuMsg.SequenceNumber)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  11 Octet NPDUNumber:%d", m.GtpuMsg.NPDUNumber)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  12 Octet NextExtHeaderType:%d", m.GtpuMsg.NextExtHeaderType)

		// 扩展头域解析,根据GtpMsg.MessageType解析head与IE
		// 以业务消息类型控制解析的头域,r.Len() is Residual byte
		// 无扩展头域，直接返回成功nil
		err := m.GtpuExtensionHeader(r, r.Len()) //GtpuExtensionHeader(r *bytes.Reader, residualLen int) error
		if err != nil {
			//fmt.Println(err)
			rlogger.Trace(moduleTag, rlogger.WARN, nil, " Error: %s", err)
			return err
		}

	}
	// IE 解析
	err := m.GtpuInformationElement(r, r.Len()) //GtpuInformationElement(r *bytes.Reader, residualLen int) error
	if err != nil {
		//fmt.Println(err)
		rlogger.Trace(moduleTag, rlogger.WARN, nil, " Error: %s", err)
		return err
	}

	return nil
}

// mHeaderDecode AlwaysPresentFields
func (m *N3MsgParser) mHeaderDecode() (*bytes.Reader, int, error) {
	FuncEntry(moduleTag, nil)
	//安全检查
	if len(m.Msgbuf) < GTPV1_U_HEADER_MIN_LEN {
		return nil, Protocol_ZERO, ErrInvalidIELen
	}
	GtpMsg := GTPPDU{Gtpv1uHeader: Gtpv1uHeader{}}
	//make([]byte, len(m.Msgbuf)),

	r := bytes.NewReader(m.Msgbuf)
	var topOctet uint8
	// 1 Octet
	binary.Read(r, binary.BigEndian, &topOctet)
	GtpMsg.Version = topOctet >> 5
	//top 3 bits
	GtpMsg.PT = topOctet >> 4 & 1
	// 1 bit
	GtpMsg.Spare = topOctet >> 3 & 1
	// 1 bit
	GtpMsg.EFlag = topOctet >> 2 & 1
	// 1 bit
	GtpMsg.SFlag = topOctet >> 1 & 1
	// 1 bit
	GtpMsg.PNFlag = topOctet & 1
	// 1 bit
	// Message Type 2 Octet
	binary.Read(r, binary.BigEndian, &GtpMsg.MessageType)
	//Length 3-4 Octet
	binary.Read(r, binary.BigEndian, &GtpMsg.Length)
	//Tunnel Endpoint Identifier 5-8 Octet
	binary.Read(r, binary.BigEndian, &GtpMsg.TEID)
	// 长度在1500范围内
	mFieldLen := int(r.Size()) - r.Len()
	payloadLen := r.Len()
	//fmt.Println("Read always present fields length:", mFieldLen)
	//fmt.Println("Read Optional fields + payload length:", payloadLen)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Read always present fields length:%d", mFieldLen)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Read Optional fields + payload length:%d", payloadLen)

	//fmt.Println("1 Octet:")
	//fmt.Println(GtpMsg.Version, GtpMsg.PT, GtpMsg.Spare, GtpMsg.EFlag, GtpMsg.SFlag, GtpMsg.PNFlag)

	//fmt.Println("1 Octet GtpMsg.Version:", GtpMsg.Version)
	//fmt.Println("1 Octet GtpMsg.PT:", GtpMsg.PT)
	//fmt.Println("1 Octet GtpMsg.Spare:", GtpMsg.Spare)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "1 Octet GtpMsg.Version:%d", GtpMsg.Version)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "1 Octet GtpMsg.PT:%d", GtpMsg.PT)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "1 Octet GtpMsg.Spare:%d", GtpMsg.Spare)

	//fmt.Println("1 Octet GtpMsg.EFlag:", GtpMsg.EFlag)
	//fmt.Println("1 Octet GtpMsg.SFlag:", GtpMsg.SFlag)
	//fmt.Println("1 Octet GtpMsg.PNFlag:", GtpMsg.PNFlag)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "1 Octet GtpMsg.EFlag:%d", GtpMsg.EFlag)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "1 Octet GtpMsg.SFlag:%d", GtpMsg.SFlag)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "1 Octet GtpMsg.PNFlag:%d", GtpMsg.PNFlag)

	//fmt.Println("2 Octet MessageType:", GtpMsg.MessageType)
	//fmt.Println("3-4 Octet Length:   ", GtpMsg.Length)
	//fmt.Println("5-8 Octet TEID:     ", GtpMsg.TEID)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "2 Octet MessageType:%d", GtpMsg.MessageType)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "3-4 Octet Length:   %d", GtpMsg.Length)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "5-8 Octet TEID:     %d", GtpMsg.TEID)
	// 有效性检查 GtpMsg.Version =1 ，GtpMsg.PT =1
	if GtpMsg.Version != Protocol_version || GtpMsg.PT != Protocol_Type {
		rlogger.Trace(moduleTag, rlogger.WARN, nil, " Gtp protocol version %d:%s", GtpMsg.Version, ErrVersion)
		return nil, Protocol_ZERO, ErrVersion
	}
	m.GtpuMsg = &GtpMsg
	return r, payloadLen, nil
}

// 编码GTPv1-U消息
// var sendMsgbuf [1024]byte
func (g *GTPPDU) EncodeMsg() ([]byte, error) {
	FuncEntry(moduleTag, nil)
	buf := new(bytes.Buffer)
	n := 0
	// TODO: 增加有效性检查
	EncodeMHeader(&g.Gtpv1uHeader, buf)
	n = buf.Len()

	//fmt.Println("head buf.len", n)
	//fmt.Println("head buf ", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head+option buf.len %d", n)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head+option buf %#x", buf.Bytes())

	EncodeOHeader(&g.Gtpv1uHeader, buf)
	// write payload
	bodyLen, err := buf.Write(g.Gtpbody)
	if err != nil {
		//fmt.Println("Gtpbody write fail.")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, " Gtpbody write fail:%s", err)
	}
	//fmt.Println("bodyLen:", bodyLen)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "bodyLen: %d", bodyLen)

	out := buf.Bytes()

	return out, nil
}

//EncodeOHeader Optional fields encode
func EncodeOHeader(g *Gtpv1uHeader, buf *bytes.Buffer) {
	FuncEntry(moduleTag, nil)
	if g.EFlag == Protocol_Present ||
		g.SFlag == Protocol_Present ||
		g.PNFlag == Protocol_Present {
		binary.Write(buf, binary.BigEndian, g.SequenceNumber)
		binary.Write(buf, binary.BigEndian, g.NPDUNumber)
		binary.Write(buf, binary.BigEndian, g.NextExtHeaderType)
		n := buf.Len()
		//fmt.Println("head+option buf.len", n)
		//fmt.Println("head+option buf", buf.Bytes())
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head+option buf.len %d", n)
		//rlogger.Trace(moduleTag, rlogger.INFO, nil,  "head+option buf %#x", buf.Bytes())
	}
}

// EncodeMHeader always present fields encode
func EncodeMHeader(g *Gtpv1uHeader, buf *bytes.Buffer) {
	FuncEntry(moduleTag, nil)
	var topLine uint8
	topLine = g.Version<<5 | //top 3 bits
		g.PT<<4 | //1 bits
		g.Spare<<3 | //1 bits
		g.EFlag<<2 | //1 bits
		g.SFlag<<1 | //1 bits
		g.PNFlag
	//1 bits
	binary.Write(buf, binary.BigEndian, topLine)

	//fmt.Printf("head topline buf %b ", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head topline buf %b", buf.Bytes())

	binary.Write(buf, binary.BigEndian, g.MessageType)
	binary.Write(buf, binary.BigEndian, g.Length)
	binary.Write(buf, binary.BigEndian, g.TEID)
}

// 触发DL PDU SESSION INFORMATION 时调用
// DL PDU SESSION INFORMATION
func (g *GPDUSessionContDL) EncodeMsg() ([]byte, error) {
	FuncEntry(moduleTag, nil)
	buf := new(bytes.Buffer)
	//n := 0
	// TODO: 增加标准头填充
	EncodeMHeader(&g.Gtpv1uHeader, buf)
	//n = buf.Len()

	//fmt.Println("head buf.len", n)
	//fmt.Println("head buf ", buf.Bytes())
	//rlogger.Trace(moduleTag, rlogger.INFO, nil,  "head buf.len %d", n)
	//rlogger.Trace(moduleTag, rlogger.INFO, nil,  "head buf %#x", buf.Bytes())
	EncodeOHeader(&g.Gtpv1uHeader, buf)
	//DL PDU SESSION INFORMATION
	//PDU Session Container
	g.EncodeExtHeader(buf)
	//
	// write payload
	bodyLen, err := buf.Write(g.Gtpbody)
	if err != nil {
		//fmt.Println("Gtpbody write fail.")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, " Gtpbody write fail:%s", err)
	}
	//fmt.Println("bodyLen:", bodyLen)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "bodyLen: %d", bodyLen)
	//

	out := buf.Bytes()

	return out, nil
}

func (e *EchoRequest) EncodeMsg() ([]byte, error) {
	FuncEntry(moduleTag, nil)
	buf := new(bytes.Buffer)
	n := 0
	// TODO: 增加标准头填充
	EncodeMHeader(&e.Gtpv1uHeader, buf)
	n = buf.Len()

	//fmt.Println("head buf.len", n)
	//fmt.Println("head buf ", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head buf.len %d", n)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head buf %#x", buf.Bytes())
	EncodeOHeader(&e.Gtpv1uHeader, buf)
	//Echo Request
	// Private Extension
	e.EncodeIE(buf)

	out := buf.Bytes()

	return out, nil
}

func (e *EchoResponse) EncodeMsg() ([]byte, error) {
	FuncEntry(moduleTag, nil)
	buf := new(bytes.Buffer)
	n := 0
	// TODO: 外部增加标准头填充
	EncodeMHeader(&e.Gtpv1uHeader, buf)
	n = buf.Len()

	//fmt.Println("head buf.len", n)
	//fmt.Println("head buf ", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head buf.len %d", n)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head buf %#x", buf.Bytes())

	EncodeOHeader(&e.Gtpv1uHeader, buf)
	//Echo Response
	// Private Extension
	e.EncodeIE(buf)

	out := buf.Bytes()

	return out, nil
}

// ErrIndication Error Indication encode
func (e *ErrIndication) EncodeMsg() ([]byte, error) {
	FuncEntry(moduleTag, nil)
	buf := new(bytes.Buffer)
	n := 0
	// TODO: 外部增加标准头填充
	EncodeMHeader(&e.Gtpv1uHeader, buf)
	n = buf.Len()
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head buf.len %d", n)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head buf %#x", buf.Bytes())

	//fmt.Println("head buf.len", n)
	//fmt.Println("head buf ", buf.Bytes())
	EncodeOHeader(&e.Gtpv1uHeader, buf)

	e.EncodeIE(buf)

	out := buf.Bytes()

	return out, nil
}

// gNB simulator 上行消息编码
func (g *GPDUSessionContUL) EncodeMsg() ([]byte, error) {
	FuncEntry(moduleTag, nil)
	buf := new(bytes.Buffer)
	n := 0
	// TODO: 增加标准头填充
	EncodeMHeader(&g.Gtpv1uHeader, buf)
	n = buf.Len()

	//fmt.Println("head buf.len", n)
	//fmt.Println("head buf ", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head buf.len %d", n)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "head buf %#x", buf.Bytes())
	EncodeOHeader(&g.Gtpv1uHeader, buf)
	//DL PDU SESSION INFORMATION
	//PDU Session Container
	g.EncodeExtHeader(buf)
	//
	// write payload
	bodyLen, err := buf.Write(g.Gtpbody)
	if err != nil {
		//fmt.Println("Gtpbody write fail.")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, " Gtpbody write fail:%s", err)
	}
	//fmt.Println("bodyLen:", bodyLen)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "bodyLen: %d", bodyLen)
	//

	out := buf.Bytes()

	return out, nil
}
