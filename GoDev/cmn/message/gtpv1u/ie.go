package gtpv1u

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"lite5gc/cmn/rlogger"
)

// 3GPP TS 29.281 V15.5.0 (2018-12)
func (m *N3MsgParser) GtpuInformationElement(r *bytes.Reader, residualLen int) error {
	FuncEntry(moduleTag, nil)
	switch m.GtpuMsg.MessageType {
	case Msg_Type_Echo_Request:
		EchoRequestMsg := &EchoRequest{}
		EchoRequestMsg.Gtpv1uHeader = m.GtpuMsg.Gtpv1uHeader
		//返回解析结果buf
		m.EchoRequestMsg = EchoRequestMsg
		//解析对应消息，上行Echo Request消息解码
		err := EchoRequestMsg.IEHandler(r)
		if err != nil {
			return err
		}

	case Msg_Type_Echo_Response:
		EchoResponseMsg := &EchoResponse{}
		EchoResponseMsg.Gtpv1uHeader = m.GtpuMsg.Gtpv1uHeader
		//返回解析结果buf
		m.EchoResponseMsg = EchoResponseMsg
		//解析对应消息，上行Echo Response消息解码
		err := EchoResponseMsg.IEHandler(r)
		if err != nil {
			return err
		}

	case Msg_Type_Error_Indication:
		ErrIndicationMsg := &ErrIndication{} // 当扩展头存在时，N3MsgParser已经挂载了ErrIndication，只需追加IE信息
		ErrIndicationMsg.Gtpv1uHeader = m.GtpuMsg.Gtpv1uHeader
		ErrIndicationMsg.IEHandler(r)

		//返回解析结果buf
		if m.ErrIndicationMsg != nil {
			m.ErrIndicationMsg.IE = ErrIndicationMsg.IE
			return nil
		}
		m.ErrIndicationMsg = ErrIndicationMsg
		//fmt.Println(m.ErrIndicationMsg)

	case Msg_Type_Support_Ext_H_Notification:

	case Msg_Type_End_Marker:

	case Msg_Type_G_PDU:
		return nil
	default:
		return ErrUnknownSigMsg

	}
	return nil
}

//上行Echo Request消息解码
// 解析IE处理 ,当前支持一个Private Extension处理
func (e *EchoRequest) IEHandler(r *bytes.Reader) error {
	FuncEntry(moduleTag, nil)
	// 可选IE 默认不出现
	if r.Len() == 0 {
		return nil
	}

	err := IEPriExtDecode(r, &e.IE.IEPriExt)
	if err != nil {
		return err
	}
	//PriExt成功解析，即存在
	e.IE.isPresence = true
	return nil
}

// 下行发送Echo Request消息编码
// Private Extension
func (e *EchoRequest) EncodeIE(buf *bytes.Buffer) error {
	FuncEntry(moduleTag, nil)
	if e.MessageType != Msg_Type_Echo_Request {
		return ErrUnexpSigMsg
	}

	/*// 8.6 Private Extension
	type IEPriExt struct {
		Type     uint8
		Length   uint16
		ExtId    uint16
		ExtValue []byte //长度是Length-2
	}*/
	if e.IE.isPresence == false {
		return nil
	}
	binary.Write(buf, binary.BigEndian, e.IE.Type)
	binary.Write(buf, binary.BigEndian, e.IE.Length)
	binary.Write(buf, binary.BigEndian, e.IE.ExtId)
	n, err := buf.Write(e.IE.ExtValue)
	if err != nil {
		//fmt.Println("Private Extension write fail.")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Private Extension write fail:%s", err)

	}
	//fmt.Println("Private Extension value len :", n)

	//fmt.Printf("Private Extension buf %b ", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Private Extension value len :%d", n)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Private Extension buf %b ", buf.Bytes())

	return nil
}

// 上行Echo Response 消息解码
// 解析IE处理 ,当前支持必选the Recovery information element
// 支持一个可选Private Extension处理
func (e *EchoResponse) IEHandler(r *bytes.Reader) error {
	FuncEntry(moduleTag, nil)
	// the Recovery information element decode
	// 必选IE，需处理错误
	err := IERecoveryDecode(r, &e.IE.Recovery)
	if err != nil {
		return err
	}
	// The Private Extension information element decode
	// 可选IE 默认不出现
	if r.Len() == 0 {
		return nil
	}
	err = IEPriExtDecode(r, &e.IE.PriExt)
	if err != nil {
		return err
	}
	//PriExt成功解析，即存在
	e.IE.IsPresence = true
	return nil
}

// IE 解析 Recovery
func IERecoveryDecode(r *bytes.Reader, recovery *IERecovery) error {
	FuncEntry(moduleTag, nil)
	/*Bits
	Octets		8	7	6	5	4	3	2	1
	1		Type = 14 (Decimal)
	2       Restart counter*/
	binary.Read(r, binary.BigEndian, &recovery.Type)
	if recovery.Type != IE_Type_Recovery {
		//fmt.Println("Unexpected Information Element ")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Unexpected Information Element ")
		return ErrUnexpIE
	}
	//The value of the restart counter shall be set to 0 by the sending entity and ignored by the receiving entity.
	binary.Read(r, binary.BigEndian, &recovery.ReCounter)
	//fmt.Println("\nRead IE fields :")
	//fmt.Println("  13 Octet Recovery.Type:", recovery.Type)
	//fmt.Println("  14 Octet Recovery.ReCounter:", recovery.ReCounter)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "\nRead IE fields :")
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  Recovery.Type:     %d", recovery.Type)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  Recovery.ReCounter:%d", recovery.ReCounter)
	return nil
}

// IE 解析 8.3 Tunnel Endpoint Identifier Data I
func IETEIDDecode(r *bytes.Reader, teid *IETunEpIdI) error {
	FuncEntry(moduleTag, nil)
	// Tunnel Endpoint Identifier Data I	Mandatory  8.3
	/*Bits
	Octets	8	7	6	5	4	3	2	1
	1 	Type = 16 (Decimal)
	2 - 5	Tunnel Endpoint Identifier Data I*/
	binary.Read(r, binary.BigEndian, &teid.Type)
	if teid.Type != IE_Type_Tun_E_Id_Data_I {
		//fmt.Println("Unexpected Information Element ")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Unexpected Information Element ")
		return ErrUnexpIE
	}

	binary.Read(r, binary.BigEndian, &teid.TeidI)
	//fmt.Println("\nRead IE fields Teid :")
	//fmt.Println("  13 Octet Recovery.Type:", teid.Type)
	//fmt.Println("  14 Octet Recovery.ReCounter:", teid.TeidI)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "\nRead IE fields Teid :")
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  Recovery.Type:%d", teid.Type)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  Recovery.ReCounter:%d", teid.TeidI)

	return nil
}

// IE 解析 8.4 GTP-U Peer Address
func IEPAddrDecode(r *bytes.Reader, paddr *IEPeerAddr) error {
	FuncEntry(moduleTag, nil)
	// GTP-U Peer Address	Mandatory	8.4
	/*Bits
	Octets	8	7	6	5	4	3	2	1
	1 	Type = 133 (Decimal)
	2 - 3 Length
	4 - 6	IPv4 or IPv6 Address*/
	// The Length field may have only two values (4 or 16)
	// that determine if the Value field contains IPv4 or IPv6 address.
	binary.Read(r, binary.BigEndian, &paddr.Type)
	if paddr.Type != IE_Type_Peer_Addr {
		//fmt.Println("Unexpected Information Element ")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Unexpected Information Element ")
		return ErrUnexpIE
	}

	binary.Read(r, binary.BigEndian, &paddr.Length)
	// IPv6
	if paddr.Length == IPv6len {
		paddr.IpAddr = make([]byte, IPv6len)
		binary.Read(r, binary.BigEndian, &paddr.IpAddr)
	} else {
		// IPv4
		paddr.IpAddr = make([]byte, IPv4len)
		binary.Read(r, binary.BigEndian, &paddr.IpAddr)
	}

	//fmt.Println("\nRead IE fields Peer Address :")
	//fmt.Println("  13 Octet Peer Address length:", paddr.Length)
	fmt.Println("  14 Octet Peer Address ReCounter:", paddr.IpAddr)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "\nRead IE fields Peer Address :")
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  13 Octet Peer Address length:%d", paddr.Length)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  14 Octet Peer Address ReCounter:%s", paddr.IpAddr)

	return nil
}

// IE 解析 8.6 Private Extension
func IEPriExtDecode(r *bytes.Reader, priext *IEPriExt) error {
	FuncEntry(moduleTag, nil)
	/*// 8.6 Private Extension
	type IEPriExt struct {
		Type     uint8
		Length   uint16
		ExtId    uint16
		ExtValue []byte //长度是Length-2
	}*/
	err := binary.Read(r, binary.BigEndian, &priext.Type)
	if err != nil {
		return err
	}
	//fmt.Println("  15 Octet Private Extension Type:", priext.Type)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Octet Private Extension Type:%d", priext.Type)

	if priext.Type != IE_Type_Private_Ext {
		//fmt.Println("Unexpected Information Element ")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Unexpected Information Element ")
		return ErrUnexpIE
	}
	binary.Read(r, binary.BigEndian, &priext.Length)
	binary.Read(r, binary.BigEndian, &priext.ExtId)
	//负数时转换为有符号整数，防止溢出
	ExtValueLen := int32(priext.Length) - 2 // ExtId 占用2 octets
	if ExtValueLen < 0 {
		return ErrInvalidIELen
	}
	priext.ExtValue = make([]byte, ExtValueLen)
	n, err := r.Read(priext.ExtValue)
	if err != nil && n == 0 {
		//fmt.Println("ExtValue Read Fail")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "ExtValue Read Fail")
		priext.ExtValue = []byte{}
		return ErrInvalidIELen
	}

	return nil
}

// ErrIndication Error Indication IE Decode
/*Tunnel Endpoint Identifier Data I	Mandatory  8.3
GTP-U Peer Address	Mandatory	8.4
Private Extension	Optional	8.6*/
func (e *ErrIndication) IEHandler(r *bytes.Reader) error {
	FuncEntry(moduleTag, nil)
	// Tunnel Endpoint Identifier Data I	Mandatory  8.3
	// 必选IE，需处理错误
	err := IETEIDDecode(r, &e.IE.Teid)
	if err != nil {
		return err
	}
	// GTP-U Peer Address	Mandatory	8.4
	err = IEPAddrDecode(r, &e.IE.PeerAddr)
	if err != nil {
		return err
	}

	// The Private Extension information element decode
	// 可选IE 默认不出现
	if r.Len() == 0 {
		return nil
	}
	err = IEPriExtDecode(r, &e.IE.PriExt)
	if err != nil {
		return err
	}
	//成功解析，即存在
	e.IE.IsPresence = true
	return nil
}

// 下行发送Echo Response 消息编码
// 当前支持必选the Recovery information element
// 支持一个可选Private Extension处理
func (e *EchoResponse) EncodeIE(buf *bytes.Buffer) error {
	FuncEntry(moduleTag, nil)
	if e.MessageType != Msg_Type_Echo_Response {
		return ErrUnexpSigMsg
	}
	err := IERecoveryEncode(buf, &e.IE.Recovery)
	if err != nil {
		return err
	}
	if e.IE.IsPresence == false {
		return nil
	}
	err = IEPriExtEncode(buf, &e.IE.PriExt)
	if err != nil {
		return err
	}
	return nil
}

func IERecoveryEncode(buf *bytes.Buffer, e *IERecovery) error {
	FuncEntry(moduleTag, nil)
	/*Bits
	Octets		8	7	6	5	4	3	2	1
	1		Type = 14 (Decimal)
	2       Restart counter*/
	binary.Write(buf, binary.BigEndian, e.Type)
	binary.Write(buf, binary.BigEndian, e.ReCounter)

	//fmt.Printf("Recovery buf %#x \n", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Recovery buf %#x \n", buf.Bytes())
	return nil
}

func IEPriExtEncode(buf *bytes.Buffer, e *IEPriExt) error {
	FuncEntry(moduleTag, nil)
	/* 8.6 Private Extension
	type IEPriExt struct {
		Type     uint8
		Length   uint16
		ExtId    uint16
		ExtValue []byte //长度是Length-2
	}*/
	binary.Write(buf, binary.BigEndian, e.Type)
	binary.Write(buf, binary.BigEndian, e.Length)
	binary.Write(buf, binary.BigEndian, e.ExtId)
	n, err := buf.Write(e.ExtValue) // 不定长
	if err != nil {
		//fmt.Println("Private Extension write fail.")
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Private Extension write fail.")
	}
	//fmt.Println("Private Extension value len :", n)
	//fmt.Printf("Private Extension buf %#x \n", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Private Extension value len :%d", n)
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Private Extension buf %#x \n", buf.Bytes())

	return nil
}

func IETEIDEncode(buf *bytes.Buffer, i *IETunEpIdI) error {
	FuncEntry(moduleTag, nil)
	// Tunnel Endpoint Identifier Data I	Mandatory  8.3
	/*Bits
	Octets	8	7	6	5	4	3	2	1
	1 	Type = 16 (Decimal)

	2 - 5	Tunnel Endpoint Identifier Data I*/
	binary.Write(buf, binary.BigEndian, i.Type)
	binary.Write(buf, binary.BigEndian, i.TeidI)

	//fmt.Printf("Tunnel Endpoint Identifier Data I buf: %#x \n", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "Tunnel Endpoint Identifier Data I buf: %#x \n", buf.Bytes())
	return nil
}

func IEPAddrEncode(buf *bytes.Buffer, i *IEPeerAddr) error {
	FuncEntry(moduleTag, nil)
	// GTP-U Peer Address	Mandatory	8.4
	/*Bits
	Octets	8	7	6	5	4	3	2	1
	1 	Type = 133 (Decimal)
	2 - 3 Length
	4 - 6	IPv4 or IPv6 Address*/
	// The Length field may have only two values (4 or 16)
	// that determine if the Value field contains IPv4 or IPv6 address.
	binary.Write(buf, binary.BigEndian, i.Type)
	binary.Write(buf, binary.BigEndian, i.Length)
	//IPv6
	if i.Length == IPv6len {

		binary.Write(buf, binary.BigEndian, []uint8(i.IpAddr))
	} else {
		//IPv4
		binary.Write(buf, binary.BigEndian, []uint8(i.IpAddr)) // 定长，uint8比byte效率高
	}

	//fmt.Printf("GTP-U Peer Address buf: %#x \n", buf.Bytes())
	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "GTP-U Peer Address buf: %#x \n", buf.Bytes())
	return nil
}

// ErrIndication Error Indication IE encode
/*Tunnel Endpoint Identifier Data I	Mandatory  8.3
GTP-U Peer Address	Mandatory	8.4
Private Extension	Optional	8.6*/
func (e *ErrIndication) EncodeIE(buf *bytes.Buffer) error {
	FuncEntry(moduleTag, nil)
	if e.MessageType != Msg_Type_Error_Indication {
		return ErrUnexpSigMsg
	}

	err := IETEIDEncode(buf, &e.IE.Teid)
	if err != nil {
		return err
	}

	err = IEPAddrEncode(buf, &e.IE.PeerAddr)
	if err != nil {
		return err
	}

	if e.IE.IsPresence == false {
		return nil
	}
	err = IEPriExtEncode(buf, &e.IE.PriExt)
	if err != nil {
		return err
	}
	return nil
}
