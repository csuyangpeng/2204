package gtpv1u

import (
	"fmt"
	"net"
	"strconv"
)

// 3GPP TS 29.281 V15.5.0 (2018-12)
// 网络常量
const (
	GTPV1_U_UE_MTU = 1500 //IP包的最大长度1500字节
	//The GTP-U header is a variable length header whose minimum length is 8 bytes.
	GTPV1_U_HEADER_MIN_LEN             = 8    //GTP-U header 8字节
	GTPV1_U_T_PDU                      = 1358 //3GPP TS 23.060 V15.4.0 Annex C
	GTPV1_U_UDP_PORT                   = 2152 //协议默认端口
	GTPV1_U_HEADER_OPTIONAL_FIELDS_LEN = 4    //Optional fields 4字节

)

// 协议常量
const (
	Protocol_ZERO        = 0
	Protocol_version     = 1
	Protocol_Type        = 1
	Protocol_Present     = 1
	Protocol_Not_Present = 0
)

//协议结构
/*  3GPP TS 29.281 V15.5.0
Figure 5.1-1: Outline of the GTP-U Header
		Bits
Octets  8	7	6 |	5	4	3	2	1
1		Version	  | PT  (*)	E	S	PN
2		Message Type
3		Length (1st Octet)
4		Length (2nd Octet)
5		Tunnel Endpoint Identifier (1st Octet)
6		Tunnel Endpoint Identifier (2nd Octet)
7		Tunnel Endpoint Identifier (3rd Octet)
8		Tunnel Endpoint Identifier (4th Octet)
9		Sequence Number (1st Octet)1) 4)
10		Sequence Number (2nd Octet)1) 4)
11		N-PDU Number2) 4)
12		Next Extension Header Type3) 4)
*/

// the GTP-U Header
type Gtpv1uHeader struct {
	//Always present fields:
	// Version field: This field is used to determine the version of the GTP-U protocol.
	// The version number shall be set to '1'.
	Version uint8 // 3 bits
	// Protocol Type (PT): This bit is used as a protocol discriminator between
	// GTP (when PT is '1') and GTP' (when PT is '0').
	PT uint8 // 1 bit
	//(*) This bit is a spare bit. It shall be sent as '0'.
	// The receiver shall not evaluate this bit.
	Spare uint8 // 1 bit
	//Extension Header flag (E): This flag indicates the presence of a meaningful
	// value of the Next Extension Header field. When it is set to '0',
	// the Next Extension Header field either is not present or, if present,
	// shall not be interpreted. When it is set to '1', the Next Extension Header
	// field is present, and shall be interpreted, as described below in this section.
	EFlag uint8 // 1 bit
	//Sequence number flag (S): This flag indicates the presence of a meaningful
	// value of the Sequence Number field.
	SFlag uint8 // 1 bit
	//N-PDU Number flag (PN): This flag indicates the presence of a meaningful
	// value of the N-PDU Number field. When it is set to '0', the N-PDU Number
	// field either is not present, or, if present, shall not be interpreted.
	// When it is set to '1', the N-PDU Number field is present, and shall
	// be interpreted, as described below in this section.
	PNFlag uint8 // 1 bit

	//Message Type: This field indicates the type of GTP-U message.
	MessageType uint8 // 8 bit

	//Length: This field indicates the length in octets of the payload,
	// i.e. the rest of the packet following the mandatory part of the
	// GTP header (that is the first 8 octets). The Sequence Number,
	// the N-PDU Number or any Extension headers shall be considered
	// to be part of the payload, i.e. included in the length count.
	Length uint16 // 16 bit
	//Tunnel Endpoint Identifier (TEID): This field unambiguously
	// identifies a tunnel endpoint in the receiving GTP U protocol
	// entity. The receiving end side of a GTP tunnel locally assigns
	// the TEID value the transmitting side has to use.
	TEID uint32 // 32 bit

	//Optional fields:
	//Sequence Number: If Sequence Number field is used for G-PDUs
	//(T-PDUs+headers), an increasing sequence number for T-PDUs
	//is transmitted via GTP-U tunnels, when transmission order
	//must be preserved. For Supported Extension Headers Notification
	//and Error Indication messages, the Sequence Number shall be
	//ignored by the receiver, even though the S flag is set to '1'.
	SequenceNumber uint16
	//N-PDU Number: This field is used at the Inter SGSN Routeing
	//Area Update procedure and some inter-system handover procedures
	//(e.g. between 2G and 3G radio access networks).
	NPDUNumber uint8
	//Next Extension Header Type: This field defines the type of Extension
	// Header that follows this field in the GTP PDU.
	NextExtHeaderType uint8
}

/* Figure 5.2.1-1: Outline of the Extension Header Format
Octets 1    Extension Header Length
2 – m		Extension Header Content
m+1		    Next Extension Header Type
*/
// Figure 5.2.1-1: Outline of the Extension Header Format
// the GTP-U Extension Header
//The Extension Header Length field specifies the length of
//     the particular Extension header in 4 octets units.
//The Next Extension Header Type field specifies the type of any
//     Extension Header that may follow a particular Extension Header.
//     If no such Header follows, then the value of the Next Extension
//     Header Type shall be 0.
type Gtpv1uExtHeaderA struct {
	Length            uint8
	NextExtHeaderType uint8
}

type Gtpv1uExtHeader struct {
	Gtpv1uExtHeaderA
	Content []byte // 4*n -2 ,n is a positive integer.其中n是正整数
}

// 消息获得的扩展头域
type MsgExtHeader struct {
	ExtHeaderList []Gtpv1uExtHeader
}

/*
Bits 7 and 8 of the Next Extension Header Type define how the recipient
shall handle unknown Extension Types, see Figure 5.2.1-2.
Bits
8      7	Meaning
0       0	Comprehension of this extension header is not required.
An Intermediate Node shall forward it to any Receiver Endpoint
不需要理解此扩展标头。 中间节点应将其转发到任何接收器端点
0       1	Comprehension of this extension header is not required.
An Intermediate Node shall discard the Extension Header Content and
not forward it to any Receiver Endpoint. Other extension headers shall
be treated independently of this extension header.
不需要理解此扩展标头。 中间节点应丢弃扩展标头内容，而不是将其转发到任何接收器端点。
其他扩展标头应独立于此扩展标头进行处理。
1       0	Comprehension of this extension header is required by
the Endpoint Receiver but not by an Intermediate Node. An Intermediate
Node shall forward the whole field to the Endpoint Receiver.
端点接收器需要理解此扩展头，但中间节点不需要理解。 中间节点应将整个字段转发到端点接收器。
1        1	Comprehension of this header type is required by recipient
(either Endpoint Receiver or Intermediate Node)
接收方（端点接收者或中间节点）需要理解这种报头类型。
*/
// getNumberbit Gets the bits in the byte
// 7 6 5 4 3 2 1 0
func getBitInByte(srcByte uint8, n uint8) uint8 {
	return ((srcByte << (7 - n)) >> 7)
}

// uint8中的二进制保存到数组[]uint8
func ConvBinary(srcByte uint8) [8]uint8 {
	var dstBinary [8]uint8
	for i := 0; i < 8; i++ {
		dstBinary[i] = getBitInByte(srcByte, uint8(i))
	}
	return dstBinary
}

/*5.2.2	Extension Header types
Next Extension Header Field Value	Type of Extension Header
0000 0000	No more extension headers
0000 0001	Reserved - Control Plane only.
0000 0010	Reserved - Control Plane only.
0000 0011	Long PDCP PDU Number.
See NOTE 2.
0010 0000	Service Class Indicator
0100 0000	UDP Port. Provides the UDP Source Port of the triggering message.
1000 0001	RAN Container
1000 0010	Long PDCP PDU Number.
See NOTE 3.
1000 0011	Xw RAN Container
1000 0100	NR RAN Container
1000 0101	PDU Session Container. See NOTE 4.
1100 0000	PDCP PDU Number [4]-[5]. See NOTE 1.
1100 0001	Reserved - Control Plane only.
1100 0010	Reserved - Control Plane only.
NOTE 1:	As an exception to the comprehension rule specified above, for a G-PDU with a Next Extension Header Field set to the value "1100 0000", the SGW shall consider this corresponding extension header as 'comprehension not required'.
NOTE 2:	This value shall be used by a source eNB or gNB complying with this release of the specification.
NOTE 3:	This value shall not be used by a source eNB or gNB complying with this release of the specification. It may be received from a source eNB complying with an earlier release of the specification, i.e. not supporting the extension header value "0000 0011".
NOTE 4:	For a G-PDU with several Extension Headers, the PDU Session Container should be the first Extension Header.

Figure 5.2.1-3: Definition of Extension Header Type

*/
const (
	s_ExtHT_NO_MORE_EXT_HEADER      = "00000000" //No more extension headers
	s_ExtHT_LONG_PDCP_PDU_NUM       = "00000011" //0000 0011	Long PDCP PDU Number.
	s_ExtHT_SERVICE_CLASS_INDICATOR = "00100000" //0010 0000	Service Class Indicator
	s_ExtHT_UDP_SOURCE_PORT         = "01000000" //0100 0000	UDP Port. Provides the UDP Source Port of the triggering message.
	s_ExtHT_RAN_CONTAINER           = "10000001" //1000 0001	RAN Container
	s_ExtHT_Xw_RAN_CONTAINER        = "10000011" //1000 0011	Xw RAN Container
	s_ExtHT_NR_RAN_CONTAINER        = "10000100" //1000 0100	NR RAN Container
	s_ExtHT_PDU_SESSION_CONTAINER   = "10000101" //1000 0101	PDU Session Container. dec:133
	s_ExtHT_PDCP_PDU_NUM            = "11000000" //1100 0000	PDCP PDU Number [4]-[5].
)

const (
	ExtHT_NO_MORE_EXT_HEADER      = 0   //No more extension headers
	ExtHT_LONG_PDCP_PDU_NUM       = 3   //0000 0011	Long PDCP PDU Number.
	ExtHT_SERVICE_CLASS_INDICATOR = 32  //0010 0000	Service Class Indicator
	ExtHT_UDP_SOURCE_PORT         = 64  //0100 0000	UDP Port. Provides the UDP Source Port of the triggering message.
	ExtHT_RAN_CONTAINER           = 129 //1000 0001	RAN Container
	ExtHT_Xw_RAN_CONTAINER        = 131 //1000 0011	Xw RAN Container
	ExtHT_NR_RAN_CONTAINER        = 132 //1000 0100	NR RAN Container
	ExtHT_PDU_SESSION_CONTAINER   = 133 //1000 0101	PDU Session Container. dec:133
	ExtHT_PDCP_PDU_NUM            = 192 //1100 0000	PDCP PDU Number [4]-[5].
)

//Parse2toInt("00000000") 转换 2进制字符串到10 进制int
func Parse2toInt(s string) (uint8, error) {
	i64, err := strconv.ParseInt(s, 2, 0)
	if err != nil {
		fmt.Println("Parse binary to Int fail.")
	}
	return uint8(i64), err
}

// 5.2.2.1	UDP Port
/*This extension header may be transmitted in Error Indication messages to
provide the UDP Source Port of the G-PDU that triggered the Error Indication.
It is 4 octets long, and therefore the Length field has value 1.
		Bits
Octets		8	7	6	5	4	3	2	1
1		0x01
2-3		UDP Port number
4		Next Extension Header Type (note)
*/

type UDPPort struct {
	Gtpv1uExtHeaderA // Length ,NextExtHeaderType
	UDPPortNumber    uint16
}

// 5.2.2.2	PDCP PDU Number
/*This extension header is transmitted, for example in UTRAN, at SRNS relocation
time, to provide the PDCP sequence number of not yet acknowledged N-PDUs. It
is 4 octets long, and therefore the Length field has value 1.
Bits
Octets		8	7	6	5	4	3	2	1
1		0x01
2		PDCP PDU number
3		PDCP PDU number.
4		Next Extension Header Type (Note 2)
*/

type PDCPPDU struct {
	Gtpv1uExtHeaderA // Length ,NextExtHeaderType
	PDCPPDUNumber    uint16
}

// 5.2.2.3	Service Class Indicator
/*This extension header identifies the service class indicator (SCI) associated
with the T-PDU carried by the downlink G-PDU. This information may be used by
the A/Gb mode GERAN access for improved radio utilisation.

Bits
Octets		8	7	6	5	4	3	2	1
1		0x01
2		Service Class Indicator
3		Spare
4		Next Extension Header Type (note)
*/
type SCI struct {
	Gtpv1uExtHeaderA // Length ,NextExtHeaderType
	SCIndicator      uint8
	spare            uint8 // the spare shall be set to zero
}

// 5.2.2.4	RAN Container
/*This extension header may be transmitted in a G-PDU over the X2 user plane
interface between the eNBs. The RAN Container has a variable length and its
content is specified in 3GPP TS 36.425 [25]. A G-PDU message with this
extension header may be sent without a T-PDU.

Bits
Octets		8	7	6	5	4	3	2	1
1		0xn
2-(4n -1)		RAN Container
4n		Next Extension Header Type (NOTE)
*/
type RANContainer struct {
	Gtpv1uExtHeaderA        // Length ,NextExtHeaderType
	RANContainer     []byte // The RAN Container has a variable length
}

// 5.2.2.5	Xw RAN Container
/*This extension header may be transmitted in a G-PDU over the Xw user plane
interface between the eNB and the WLAN Termination (WT). The Xw RAN Container
has a variable length and its content is specified in 3GPP TS 36.465 [27].
	A G-PDU message with this extension header may be sent without a T-PDU.

Bits
Octets		8	7	6	5	4	3	2	1
1		0xn
2-(4n -1)		Xw RAN Container
4n		Next Extension Header Type (NOTE)
*/

type XwRANContainer struct {
	Gtpv1uExtHeaderA        // Length ,NextExtHeaderType
	XwRANContainer   []byte // The Xw RAN Container has a variable length
}

// 5.2.2.6	NR RAN Container
/*This extension header may be transmitted in a G-PDU over the X2-U, Xn-U and
F1-U user plane interfaces, within NG-RAN and, for EN-DC, within E-UTRAN.
	The NR RAN Container has a variable length and its content is specified in
3GPP TS 38.425 [30]. A G-PDU message with this extension header may be sent
without a T-PDU.

Bits
Octets		8	7	6	5	4	3	2	1
1		0xn
2-(4n -1)		NR RAN Container
4n		Next Extension Header Type (NOTE)
*/
type NRRANContainer struct {
	Gtpv1uExtHeaderA        // Length ,NextExtHeaderType
	NRRANContainer   []byte // The NR RAN Container has a variable length
}

// 5.2.2.7	PDU Session Container
/*This extension header may be transmitted in a G-PDU over the N3 and N9 user
plane interfaces, between NG-RAN and UPF, or between two UPFs. The PDU Session
Container has a variable length and its content is specified in 3GPP TS 38.415 [31].

Bits
Octets		8	7	6	5	4	3	2	1
1		0xn
2-(4n -1)		PDU Session Container
4n		Next Extension Header Type (NOTE)
*/
type PDUSessionContainerDL struct {
	// The PDU Session Container has a variable length
	Gtpv1uExtHeaderA
	DLPDUSession
}

type PDUSessionContainerUL struct {
	// The PDU Session Container has a variable length
	Gtpv1uExtHeaderA
	ULPDU ULPDUSession
}

// PDUSessionContainer
// Its content is specified in 3GPP TS 38.415
// 5.4 Elementary procedures

// 5.5.2.1	DL PDU SESSION INFORMATION (PDU Type 0)
/*Bits	                        Number of Octets
7	6	5	4	3	2	1	0
PDU Type (=0) |Spare	        1
PPP|RQI|QoS Flow Identifier 	1
PPI	      Spare	               0 or 1
Padding 	                   0-3

Figure 5.5.2.1-1: DL PDU SESSION INFORMATION (PDU Type 0) Format
*/
type DLPDUSession struct {
	PDUType uint8 //4 bits
	SpareA  uint8 // 4 bits
	PPP     uint8 //1 bits Paging Policy Presence
	RQI     uint8 //1 bits Reflective QoS Indication
	QFI     uint8 //6 bits QoS Flow Identifier
	PPI     uint8 //3 bits Paging Policy Indicator
	SpareB  uint8 //5 bits
	Padding uint8 //8 bits //当PPP存在时，length为2，Padding应填充3byte，当PPP不存在时，length为1，不用填充
}

// 5.5.2.2	UL PDU SESSION INFORMATION (PDU Type 1)
/*Bits	                      Number of Octets
7	6	5	4	3	2	1	0
PDU Type (=1)	Spare	      1
Spare  QoS Flow Identifier 	  1
Padding 	                 0-3

Figure 5.5.2.2-1: UL PDU SESSION INFORMATION (PDU Type 1) Format
*/
type ULPDUSession struct {
	PDUType uint8 //4 bits
	SpareA  uint8 //4 bits
	SpareB  uint8 //2 bits
	QFI     uint8 //6 bits QoS Flow Identifier
	//Padding uint16 //16 bits // 当前协议不使用
}

//bits
const (
	ULPDUSession_PDUType_Len = 4
	ULPDUSession_SpareA_Len  = 4
	ULPDUSession_SpareB_Len  = 2
	ULPDUSession_QFI_Len     = 6
	ULPDUSession_Padding_Len = 16
)

// 5.5.3.1	PDU Type
/*Description: The PDU Type indicates the structure of the PDU session UP
frame. The field takes the value of the PDU Type it identifies; i.e. "0"
for PDU Type 0. The PDU type is in bit 4 to bit 7 in the first octet of the frame.
Value range: {0= DL PDU SESSION INFORMATION, 1=UL PDU SESSION INFORMATION,
2-15=reserved for future PDU type extensions}.
Field length: 4 bits.
*/
const (
	PDU_Type_DL_PDU_Session_Information = 0
	PDU_Type_UL_PDU_Session_Information = 1
)

//5.5.3.2	Spare
/*Description: The spare field is set to "0" by the sender and should not be
interpreted by the receiver. This field is reserved for later versions.
Value range: (0–2^n-1).
Field Length: n bits.
*/
const (
	SPARE_ZERO = 0
)

// 5.5.3.3	QoS Flow Identifier (QFI)
/*Description: When present this parameter indicates the QoS Flow
Identifier of the QoS flow to which the transferred packet belongs.
Value range: {0..2^6-1}.
Field length: 6 bits.
23.501 QOS
*/
const (
	QFI_Start = 0
	QFI_End   = 63
)

// 5.5.3.4	Reflective QoS Indicator (RQI)
/*Description: This parameter indicates activation of the reflective QoS
towards the UE for the transferred packet as described in clause 5.4.1.1.
It is used only in the downlink direction. If RQA (Reflective QoS Activation)
has not been configured for the involved QoS flow, the RQI shall be ignored
by the NG-RAN node.
Value range: {0= Reflective QoS activation not triggered, 1= Reflective QoS
activation triggered}.
Field length: 1 bit.
*/
const (
	RQI_Not_Triggered = 0
	RQI_Act_Triggered = 1
)

// 5.5.3.5	Padding
/*Description: The padding is included at the end of the frame to ensure that
the PDU Session user plane protocol PDU length (including padding and the
future extension) is (n*4– 2) octets, where n is a positive integer. If
there is any future extension, the padding should be added after the future
extensions.
Field Length: 0–3 octets.
*/
const (
	Padding_ZERO = 0
)

// 5.5.3.6	Paging Policy Presence (PPP)
/*Description: When present this parameter indicates the presence of the
Paging Policy Indicator (PPI).
Value range: {0= Paging Policy Indicator not present, 1= Paging Policy
Indicator present}.
Field length: 1 bit.
*/
const (
	PPP_Not_Present = 0
	PPP_Present     = 1
)

// 5.5.3.7	Paging Policy Indicator (PPI)
/*Description: The Paging Policy Indicator is used for paging policy
differentiation (see details in 3GPP TS 23.501 [5]). This field applies
to PDU sessions of IP type.
Value range: {0..2^3-1}.
Field length: 3 bits.
*/
const (
	PPI_A = 0
	PPI_B = 7
)

// 6	GTP-U Message Formats
/*Table 6.1-1: Messages in GTP-U
Message Type value (Decimal)	GTP-C	GTP-U	GTP'
1	Echo Request		          X 	X	  x
2	Echo Response		          X	    X	  x
26	Error Indication			        X
31	Supported Extension
      Headers Notification	      X	    X
254	End Marker			                X
255	G-PDU			                    X
*/
const (
	Msg_Type_Echo_Request               = 1
	Msg_Type_Echo_Response              = 2
	Msg_Type_Error_Indication           = 26
	Msg_Type_Support_Ext_H_Notification = 31
	Msg_Type_End_Marker                 = 254
	Msg_Type_G_PDU                      = 255
)

//IE
/*Table 8.1-1: Information Elements
IE Type Value	Format	Information Element	Reference
0-13	TV	Reserved in 3GPP TS 29.060 [6]
14	TV	Recovery	8.2
15	TV	Reserved in 3GPP TS 29.060 [6]
16	TV	Tunnel Endpoint Identifier Data I	8.3
17-132	TV/TLV	Reserved in 3GPP TS 29.060 [6]
133	TLV	GSN Address. See NOTE 1.	8.4
134-140	TLV	Reserved in 3GPP TS 29.060 [6]
141	TLV	Extension Header Type List	8.5
142-254	TLV	Reserved in 3GPP TS 29.060 [6]
255	TLV	Private Extension	8.6
NOTE 1:	This IE is named as " GTP-U Peer Address" in the rest of this specification.
The most significant bit in the Type field is set to 0 when the TV format is
used and set to 1 for the TLV format.
*/
const (
	IE_Type_Recovery        = 14
	IE_Type_Tun_E_Id_Data_I = 16
	IE_Type_Peer_Addr       = 133
	IE_Type_Ext_H_Type_List = 141
	IE_Type_Private_Ext     = 255
)

// 8.2 Recovery
/*Bits
Octets		8	7	6	5	4	3	2	1
1		Type = 14 (Decimal)
2       Restart counter
*/
type IERecovery struct {
	Type      uint8
	ReCounter uint8
}

// 8.3	Tunnel Endpoint Identifier Data I
/*Bits
Octets		8	7	6	5	4	3	2	1
1		Type = 16 (Decimal)
2-5     Tunnel Endpoint Identifier Data I
*/
type IETunEpIdI struct {
	Type  uint8
	TeidI uint32
}

// 8.4	GTP-U Peer Address
/*Bits
Octets		8	7	6	5	4	3	2	1
1		Type = 133 (Decimal)
2-3     Length (4 or 16) net.IPv4len IPv6len
4-n    IPv4 or IPv6 Address
*/
// The Length field contains the length of the information element excluding
// the Type and Length field.
// IP address lengths (bytes).
const (
	IPv4len = net.IPv4len
	IPv6len = net.IPv6len
)

type IEPeerAddr struct {
	Type   uint8
	Length uint16
	IpAddr net.IP //[]byte
}

// 8.5	Extension Header Type List
/*Bits
Octets		8	7	6	5	4	3	2	1
1		Type = 141 (Decimal)
2       Length (value=n)
3-n+2   Extension types list
*/
type IEExtHTList struct {
	Type         uint8
	Length       uint8
	ExtTypesList []byte
}

// 8.6 Private Extension
type IEPriExt struct {
	Type     uint8
	Length   uint16
	ExtId    uint16
	ExtValue []byte //长度是Length-2
}

// 7.2	Path Management Messages
/*7.2.1	Echo Request
A GTP-U peer may send an Echo Request on a path to the other GTP-U peer to find
out if it is alive (see section Path Failure).
Table 7.2.1-1: Information Elements in an Echo Request
Information element |	Presence requirement	| Reference
Private Extension	|    Optional	            | 8.6
*/
type EchoRequestIE struct {
	IEPriExt
	isPresence bool
}

type EchoRequest struct {
	Gtpv1uHeader
	IE EchoRequestIE
}

// 7.2.2	Echo Response
/*The message shall be sent as a response to a received Echo Request.
Table 7.2.2-1: Information Elements in an Echo Response
Information element 	Presence requirement	Reference
Recovery	            Mandatory	            8.2
Private Extension	    Optional	            8.6
*/
type EchoResponseIE struct {
	Recovery   IERecovery
	PriExt     IEPriExt
	IsPresence bool
}

type EchoResponse struct {
	Gtpv1uHeader
	IE EchoResponseIE
}

// 7.2.3	Supported Extension Headers Notification
/*This message indicates a list of supported Extension Headers that the GTP
entity on the identified IP address can support.
Table 7.2.3-1: Information Elements in Supported Extension Headers Notification
Information element	       Presence requirement Reference
Extension Header Type List	Mandatory       	8.5
*/
type SupportExtHNotification struct {
	IEExtHTList
}

// 7.3	Tunnel Management Messages
//7.3.1	Error Indication
/*Table 7.3.1-1: Information Elements in an Error Indication
Information element	         Presence requirement	Reference
Tunnel Endpoint Identifier Data I	Mandatory	8.3
GTP-U Peer Address	               Mandatory	8.4
Private Extension	               Optional	8.6
*/
type ErrIndicationIE struct {
	Teid       IETunEpIdI
	PeerAddr   IEPeerAddr
	PriExt     IEPriExt
	IsPresence bool //Private Extension presence
}
type ErrIndicationExt struct {
	UdpPort UDPPort
}
type ErrIndication struct {
	Gtpv1uHeader
	ExtH ErrIndicationExt
	IE   ErrIndicationIE
}

// 7.3.2	End Marker
/*Table 7.3.2-1: Information Elements in End Marker message
Information element	Presence requirement	Reference
Private Extension	Optional	8.6
*/
type EndMarker struct {
	IEPriExt
	IsPresence bool
}
