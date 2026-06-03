package gtpv1u

import (
	"bytes"
	"encoding/binary"

	"lite5gc/cmn/rlogger"
)

const moduleTag rlogger.ModuleTag = "gtpu"

var FuncEntry = rlogger.FuncEntry

// 3GPP TS 29.281 V15.5.0 (2018-12)
func (m *N3MsgParser) GtpuExtensionHeader(r *bytes.Reader, residualLen int) error {
	FuncEntry(moduleTag, nil)
	switch m.GtpuMsg.MessageType {
	case Msg_Type_Echo_Request:
		if m.GtpuMsg.NextExtHeaderType != 0 {
			//fmt.Println("Echo Request NextExtHeaderType wrong")
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, ErrUnexpExtHeaderType)
			return ErrUnexpExtHeaderType
		}
	case Msg_Type_Echo_Response:

	case Msg_Type_Error_Indication:
		if m.GtpuMsg.NextExtHeaderType == ExtHT_NO_MORE_EXT_HEADER {
			return nil
		}
		if m.GtpuMsg.NextExtHeaderType != ExtHT_UDP_SOURCE_PORT {
			return ErrUnexpExtHeaderType
		}
		ErrIndExtHeadHandler(m, r)
	case Msg_Type_Support_Ext_H_Notification:

	case Msg_Type_End_Marker:

	case Msg_Type_G_PDU:
		m.GPDUHandler(r)
	default:
		return ErrUnknownSigMsg

	}
	return nil
}

func ErrIndExtHeadHandler(m *N3MsgParser, r *bytes.Reader) error {
	FuncEntry(moduleTag, nil)
	ErrIndicationMsg := &ErrIndication{}
	ErrIndicationMsg.Gtpv1uHeader = m.GtpuMsg.Gtpv1uHeader
	// 解码 UDPPort //可选
	ExtHUdpPortDecode(r, &ErrIndicationMsg.ExtH.UdpPort)
	m.ErrIndicationMsg = ErrIndicationMsg
	return nil
}

// 解码 Extension Header types : UDP Port
func ExtHUdpPortDecode(r *bytes.Reader, port *UDPPort) error {
	FuncEntry(moduleTag, nil)
	/*Bits
	Octets		8	7	6	5	4	3	2	1
	1		0x01 length
	2-3		UDP Port number
	4		Next Extension Header Type (note)*/

	binary.Read(r, binary.BigEndian, port.Length)
	binary.Read(r, binary.BigEndian, port.UDPPortNumber)
	binary.Read(r, binary.BigEndian, port.NextExtHeaderType)
	return nil

}

func (m *N3MsgParser) GPDUHandler(r *bytes.Reader) error {
	FuncEntry(moduleTag, nil)
	GtpMsg := m.GtpuMsg

	if GtpMsg.NextExtHeaderType == ExtHT_NO_MORE_EXT_HEADER {
		return nil
	}
	if GtpMsg.NextExtHeaderType == ExtHT_PDU_SESSION_CONTAINER {
		// PDU_SESSION_CONTAINER 处理,处理来自NG-RAN的PDU Session Information
		GPDUSMsg := &GPDUSessionContUL{}
		GPDUSMsg.Gtpv1uHeader = GtpMsg.Gtpv1uHeader
		// 解码 ExtHT_PDU_SESSION_CONTAINER
		/* Figure 5.2.1-1: Outline of the Extension Header Format
		    Octets 1    Extension Header Length
		   2 – m		Extension Header Content
		    m+1		    Next Extension Header Type
		*/
		// m为3
		// Extension Header Length
		//binary.Read(r, binary.BigEndian, &GPDUSMsg.PDUSessionContainer.Length)

		// UL PDU SESSION INFORMATION
		/*type ULPDUSession struct {
			    PDUType uint8  //4 bits
			    SpareA  uint8  //4 bits
			    SpareB  uint8  //2 bits
			    QFI     uint8  //6 bits QoS Flow Identifier
			    //Padding uint16 //16 bits//当前协议中不存在
		       }
		*/
		var topOctet uint8
		// Octets 1 Extension Header Length
		binary.Read(r, binary.BigEndian, &topOctet)
		GPDUSMsg.PDUSessionContainer.Length = topOctet
		// 1 Octet
		binary.Read(r, binary.BigEndian, &topOctet)
		//top front 4 bits
		GPDUSMsg.PDUSessionContainer.ULPDU.PDUType = topOctet >> ULPDUSession_PDUType_Len
		//top after 4 bits
		GPDUSMsg.PDUSessionContainer.ULPDU.SpareA = topOctet & 0xf
		//The second Octet
		topOctet = 0
		binary.Read(r, binary.BigEndian, &topOctet)
		//2bits
		GPDUSMsg.PDUSessionContainer.ULPDU.SpareB = topOctet >> (8 - ULPDUSession_SpareB_Len)
		// 6bits
		GPDUSMsg.PDUSessionContainer.ULPDU.QFI = topOctet & 0x3f
		//Padding 16bits
		//binary.Read(r, binary.BigEndian, &GPDUSMsg.PDUSessionContainer.ULPDU.Padding)

		// Next Extension Header Type
		binary.Read(r, binary.BigEndian, &GPDUSMsg.PDUSessionContainer.NextExtHeaderType)
		m.GPDUSessionUL = GPDUSMsg
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  13 Octet ExtHeaderLength:%d", GPDUSMsg.PDUSessionContainer.Length)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  14 Octet PDU Type:%d", GPDUSMsg.PDUSessionContainer.ULPDU.PDUType)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  15 Octet QFI     :%d", GPDUSMsg.PDUSessionContainer.ULPDU.QFI)
		rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "  16 Octet NextExtensionHeaderType :%d", GPDUSMsg.PDUSessionContainer.NextExtHeaderType)

		return nil
	}
	return nil

}

// DL PDU SESSION INFORMATION encode , PDU SESSION CONTAINER ExtHeader
func (g *GPDUSessionContDL) EncodeExtHeader(buf *bytes.Buffer) error {
	FuncEntry(moduleTag, nil)
	if g.NextExtHeaderType != ExtHT_PDU_SESSION_CONTAINER {
		return ErrExtHeaderType
	}
	//PDU Session Container Extension Header Length
	//It is 4 octets long, and therefore the Length field has value 1.
	binary.Write(buf, binary.BigEndian, g.PDUSessionContainer.Length)
	/*type DLPDUSession struct {
		PDUType uint8 //4 bits
		SpareA  uint8 // 4 bits
		PPP     uint8 //1 bits Paging Policy Presence
		RQI     uint8 //1 bits Reflective QoS Indication
		QFI     uint8 //6 bits QoS Flow Identifier
		PPI     uint8 //3 bits Paging Policy Indicator
		SpareB  uint8 //5 bits
		Padding uint8 //3*8 bits//当PPP=1时，PPI、SpareB、Padding存在，否则都不存在
	}*/

	//PDU Session Container
	var topLine uint8
	topLine = (g.PDUSessionContainer.PDUType << 4) | //top 4 bits
		(g.PDUSessionContainer.SpareA & 0xf) //4 bits
	//top octet
	binary.Write(buf, binary.BigEndian, topLine)

	//The second octet，PPP，RQI,QFI
	var currentLine uint8
	currentLine = (g.PDUSessionContainer.PPP << 7) | // 1 bit PPP 第8位
		(g.PDUSessionContainer.RQI << 6 & 0x40) | //1 bit RQI 第7位
		(g.PDUSessionContainer.QFI & 0x3f) //6 bits QFI 后6位
	binary.Write(buf, binary.BigEndian, currentLine)

	//The second octet，PPI,SpareB
	currentLine = 0
	if g.PDUSessionContainer.Length == 2 { //当PPP=1时，PPI、SpareB、Padding存在，否则都不存在
		currentLine = (g.PDUSessionContainer.PPI << 5) | // 3 bits PPI 前3位
			(g.PDUSessionContainer.SpareB & 0x1f) // 1 bit SpareB 后5位
		binary.Write(buf, binary.BigEndian, currentLine)
		// Padding 3 byte
		currentLine = 0
		binary.Write(buf, binary.BigEndian, currentLine)
		binary.Write(buf, binary.BigEndian, currentLine)
		binary.Write(buf, binary.BigEndian, currentLine)
	}

	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ExtHeader buf %b ", buf.Bytes())
	// PDU Session Container NextExtHeaderType is 0
	binary.Write(buf, binary.BigEndian, g.PDUSessionContainer.NextExtHeaderType)
	return nil
}

// gNB simulator 上行扩展头编码
// UL PDU SESSION INFORMATION encode , PDU SESSION CONTAINER ExtHeader
func (g *GPDUSessionContUL) EncodeExtHeader(buf *bytes.Buffer) error {
	FuncEntry(moduleTag, nil)
	if g.NextExtHeaderType != ExtHT_PDU_SESSION_CONTAINER {
		return ErrExtHeaderType
	}
	//PDU Session Container Extension Header Length
	//It is 4 octets long, and therefore the Length field has value 1.
	binary.Write(buf, binary.BigEndian, g.PDUSessionContainer.Length)
	/*type ULPDUSession struct {
	PDUType uint8  //4 bits
	SpareA  uint8  //4 bits
	SpareB  uint8  //2 bits
	QFI     uint8  //6 bits QoS Flow Identifier
	//Padding uint16 //16 bits // 当前协议不使用
	}*/

	//PDU Session Container
	var topLine uint8
	topLine = (g.PDUSessionContainer.ULPDU.PDUType << 4) | //top 4 bits
		(g.PDUSessionContainer.ULPDU.SpareA & 0xf) //4 bits
	//top octet
	binary.Write(buf, binary.BigEndian, topLine)

	//The second octet，SpareB,QFI
	var currentLine uint8
	currentLine = (g.PDUSessionContainer.ULPDU.SpareB << 6) | // 2 bit
		(g.PDUSessionContainer.ULPDU.QFI & 0x3f) //6 bits QFI 后6位
	binary.Write(buf, binary.BigEndian, currentLine)

	//The second octet，PPI,SpareB
	currentLine = 0
	if g.PDUSessionContainer.Length != 1 { //当PPP=1时，PPI、SpareB、Padding存在，否则都不存在
		// Padding 1 byte
		//currentLine = 0
		//binary.Write(buf, binary.BigEndian, currentLine)
		rlogger.Trace(moduleTag, rlogger.WARN, nil, "Failed to ExtHeader PDUSessionContainer.Length")
	}

	rlogger.Trace(moduleTag, rlogger.DEBUG, nil, "ExtHeader buf %b ", buf.Bytes())
	// PDU Session Container NextExtHeaderType is 0
	binary.Write(buf, binary.BigEndian, g.PDUSessionContainer.NextExtHeaderType)
	return nil
}
