package pfcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/rlogger"
	"net"

	"lite5gc/cmn/message/pfcp/utils"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)
// N4 消息

// IE 名称 --来源于消息
// IE type --来源于8.1IE列表
// IE 格式 --来源于8.2格式定义

//Node ID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
60	Node ID	Extendable / Subclause 8.2.38	o-4 */ // o-4,如IPv4类型时length是5
/*8.2.38	Node ID
Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 60 (decimal)
3 to 4	Length = n
5	     Spare	        Node ID Type
6 to o	Node ID value
m to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.38-1: Node ID
*/
type IENodeID struct {
	IETypeLength
	NodeIDType  uint8  `json:",omitempty"`
	NodeIDvalue []byte `json:",omitempty"`
}

func (i *IENodeID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IENodeID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}
func (i *IENodeID) GetTl() (tl IETypeLength) {
	tl.Type = i.Type
	tl.Length = i.Length
	return tl
}
func (i *IENodeID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	//EncodeTl(encBuf, i.GetTl())
	//err = binary.Write(encBuf, binary.BigEndian, i.Type)
	//if err != nil {
	//	return
	//}
	//err = binary.Write(encBuf, binary.BigEndian, i.Length)
	//if err != nil {
	//	return
	//}
	//// length check
	//if (len(i.NodeIDvalue) + 1) != i.Len() {
	//	return nil, errors.New("invalid length")
	//}
	err = binary.Write(encBuf, binary.BigEndian, i.NodeIDType)
	if err != nil {
		return
	}
	err = binary.Write(encBuf, binary.BigEndian, i.NodeIDvalue)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IENodeID) Decode(data []byte) error {
	// parse V, todo
	r := bytes.NewReader(data)
	value := make([]byte, len(data))
	n, _ := r.Read(value)
	//Node ID Type 1-4
	if n < 1 {
		return errors.New("Invalid length")
	}
	v := utils.ByteOne(value[0])
	i.NodeIDType, _ = v.GetBits(1, 4)
	switch i.NodeIDType {
	case NodeIDType_IPv4_address:
		//	the Node ID value length shall be 4 Octet.
		if len(value[1:]) != 4 {
			return errors.New("Invalid length")
		}
		i.NodeIDvalue = value[1:]
	case NodeIDType_IPv6_address:
		//	the Node ID value length shall be 16 Octet.
		if len(value[1:]) != 16 {
			return errors.New("Invalid length")
		}
		i.NodeIDvalue = value[1:]
	case NodeIDType_FQDN_address:
		//string
		if len(value) > 1 {
			i.NodeIDvalue = value[1:]
		}
	}

	return nil
}

func (i *IENodeID) Len() int {

	return int(i.Length)
}

func (i *IENodeID) Set() {
	i.Type = IE_Node_ID
}

//IENodeIDType
/*Table 8.2.38-2: Node ID Type
Node ID Type Value (Decimal)	Node ID Type
0	IPv4 address
1	IPv6 address
2	FQDN
3 to 15	Spare, for future use.
*/
const (
	NodeIDType_IPv4_address = 0
	NodeIDType_IPv6_address = 1
	NodeIDType_FQDN_address = 2
)

func (i *IENodeID) SetValue(NodeIDType uint8, value []byte) {
	if value == nil {
		return
	}
	switch NodeIDType {
	case NodeIDType_IPv4_address:
		i.NodeIDvalue = value[:4]
	case NodeIDType_IPv6_address:
		i.NodeIDvalue = value[:16]
	case NodeIDType_FQDN_address:
		i.NodeIDvalue = value[:]
	default:
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "NodeIDType error:%d", NodeIDType)
	}
}

// CP F-SEID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
57	F-SEID	Extendable / Subclause 8.2 37	p+15-4 */
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 57 (decimal)
3 to 4	Length = n
5	Spare	Spare	Spare	Spare	Spare	Spare	V4	V6
6 to 13	SEID
m to (m+3)	IPv4 address
p to (p+15)	IPv6 address
k to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.37-1: F-SEID
*/
type IEFSEID struct {
	IETypeLength
	V4Flag   uint8 `json:",omitempty"` // V4:2,V6:1
	V6Flag   uint8
	SEID     uint64 `json:",omitempty"`
	IPv4Addr net.IP `json:",omitempty"` // 4 byte
	IPv6Addr net.IP `json:",omitempty"` // 16 byte
}

func (i *IEFSEID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEFSEID) checklen() error {
	// At least one of V4 and V6 shall be set to "1", and both may be set to "1".
	if 9 > i.Len() {
		return errors.New("invalid length")
	}
	if (i.V4Flag == 1) && (i.V6Flag == 0) {
		if 13 != i.Len() {
			return errors.New("invalid length")
		}
	}
	if (i.V4Flag == 0) && (i.V6Flag == 1) {
		//1+8+16
		if (25) != i.Len() {
			return errors.New("invalid length")
		}
	}
	if (i.V4Flag == 1) && (i.V6Flag == 1) {
		//1+8+4+16
		if (29) != i.Len() {
			return errors.New("invalid length")
		}
	}
	return nil
}
func (i *IEFSEID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEFSEID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	Spare	Spare	Spare	Spare	Spare	V4	V6
	tmpByte := i.V6Flag + (i.V4Flag << 1)
	err = binary.Write(encBuf, binary.BigEndian, tmpByte)
	if err != nil {
		return
	}
	//6 to 13	SEID
	err = binary.Write(encBuf, binary.BigEndian, i.SEID)
	if err != nil {
		return
	}
	//m to (m+3)	IPv4 address
	tmpBytes := []byte(i.IPv4Addr.To4())
	if i.V4Flag == 1 {
		err = binary.Write(encBuf, binary.BigEndian, tmpBytes)
		if err != nil {
			return
		}
	}
	//p to (p+15)	IPv6 address
	tmpBytes = []byte(i.IPv6Addr.To16())
	if i.V6Flag == 1 {
		err = binary.Write(encBuf, binary.BigEndian, tmpBytes)
		if err != nil {
			return
		}
	}
	return encBuf.Bytes(), nil

}

func (i *IEFSEID) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	var tmp uint8
	err := binary.Read(r, binary.BigEndian, &tmp)
	if err != nil {
		return err
	}
	i.V4Flag, _ = utils.GetBitInByte(tmp, 2)
	i.V6Flag, _ = utils.GetBitInByte(tmp, 1)

	err = binary.Read(r, binary.BigEndian, &i.SEID)
	if err != nil {
		return err
	}

	if i.V4Flag == 1 {
		var ip [4]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv4Addr = ip[:]
		return nil
	}
	if i.V6Flag == 1 {
		var ip [16]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv6Addr = ip[:]
		return nil
	}
	return nil
}

func (i *IEFSEID) Len() int {
	return int(i.Length)
}

const (
	IEFSEID_IPv4_address = 2
	IEFSEID_IPv6_address = 1
)

func (i *IEFSEID) Set() {
	i.Type = IE_F_SEID
}

// release of CN Tunnel Info
func (i *IEFSEID) Release() error {
	i = nil
	return nil
}

// Create PDR (Packet Detection Rule)
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
1	Create PDR	Extendable / Table 7.5.2.2-1	Not Applicable */
/*Octet 1 and 2		Create PDR IE Type = 1(decimal)
Octets 3 and 4		Length = n
Information elements	P

PDR ID  	M
Precedence	M
PDI	        M
Outer Header Removal 	C
FAR ID 	C
URR ID	C
QER ID 	C
Activate Predefined Rules 	C
*/
type IECreatePDR struct {
	IETypeLength
	PDRID              IEPDRID
	Precedence         IEPrecedence
	PDI                IEPDI
	OuterHeaderRemoval *IEOuterHeaderRemoval        `json:",omitempty"`
	FARID              *IEFARID                     `json:",omitempty"`
	URRIDs             []*IEURRID                   `json:",omitempty"`
	QERIDs             []*IEQERID                   `json:",omitempty"`
	ActPredefinedRules []*IEActivatePredefinedRules `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IECreatePDR) GetGrouped() bool {
	return i.Grouped
}

func (i *IECreatePDR) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	//PDR ID  	M
	case *IEPDRID:
		i.PDRID = *Ie
		//Precedence	M
	case *IEPrecedence:
		i.Precedence = *Ie
		//PDI	        M
	case *IEPDI:
		i.PDI = *Ie
		//Outer Header Removal 	C
	case *IEOuterHeaderRemoval:
		i.OuterHeaderRemoval = Ie
		//FAR ID 	C
	case *IEFARID:
		i.FARID = Ie
		//URR ID	C
	case *IEURRID:
		i.URRIDs = append(i.URRIDs, Ie)
		//QER ID 	C
	case *IEQERID:
		i.QERIDs = append(i.QERIDs, Ie)
		//Activate Predefined Rules 	C
	case *IEActivatePredefinedRules:
		i.ActPredefinedRules = append(i.ActPredefinedRules, Ie)

	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IECreatePDR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

// 组类型Ie解析
/*func GroupedIeParse(root PfcpIeInterface, data []byte) error {

	r := bytes.NewReader(data)
	// Tlv 解析
	// 当前 TL
	var tmp16 uint16
	err := binary.Read(r, binary.BigEndian, &tmp16)
	if err != nil {
		return err
	}
	t := tmp16

	err = binary.Read(r, binary.BigEndian, &tmp16)
	if err != nil {
		return err
	}
	l := tmp16
	if t > 32767 {
		//Figure 8.1.1-1: Information Element Format
		//	todo
		return fmt.Errorf("type error")
	}
	//Figure 8.1.1-2: 3GPP defined Information Element Format

	objModule := RegIeObject[int(t)]
	// 新建对象
	obj := CreateObject(objModule) // 对象的指针
	pfcpIe, ok := obj.(PfcpIeInterface)
	if !ok {
		//log
		return fmt.Errorf("type error")
	}
	ieData := data[4 : 4+l] // 当前V
	data = data[4+l:]       // 下一个TLV

	// 当前IE
	var grouped bool
	pfcpIeType, ok1 := obj.(PfcpGroupedIeInterface)
	if !ok1 {
		//log
		return fmt.Errorf("type error")
	}
	grouped = pfcpIeType.GetGrouped()
	if grouped {
		GroupedIeParse(pfcpIe, ieData)
	} else {
		err = pfcpIe.SetObject(t, l) // 当前IE的头设置
		if err != nil {
			return err
		}
		err = pfcpIe.Decode(ieData) // 当前IE的V解析
		if err != nil {
			return err
		}
		// 挂载当前对象
		err = root.SetObjectToParent(pfcpIe)
		if err != nil {
			return err
		}
	}

	//switch Ie := pfcpIe.(type) {
	//case *IEPDRID:
	//	i.PDRID = *Ie
	//	i.PDRID.Type = t
	//	i.PDRID.Length = l
	//
	//default:
	//
	//}
	return nil
}*/
func GroupedIeParse(root PfcpIeInterface, data []byte) error {
	for len(data) > 0 {
		r := bytes.NewReader(data)
		// Tlv 解析
		var tmp16 uint16
		err := binary.Read(r, binary.BigEndian, &tmp16)
		if err != nil {
			return err
		}
		t := tmp16

		err = binary.Read(r, binary.BigEndian, &tmp16)
		if err != nil {
			return err
		}
		l := tmp16
		if t > 32767 {
			//Figure 8.1.1-1: Information Element Format
			//	todo
			return fmt.Errorf("type error")
		}
		objModule, ok := RegIeObject[int(t)]
		if !ok {
			return fmt.Errorf("get register Ie error")
		}
		// 新建对象
		obj := CreateObject(objModule) // 对象的指针
		pfcpIe, ok := obj.(PfcpIeInterface)
		if !ok {
			//log
			return fmt.Errorf("type error.ie id :%d", t)
		}
		ieData := data[4 : 4+l] // 当前V
		data = data[4+l:]       // 下一个TLV

		err = pfcpIe.SetObject(t, l) // 当前IE的头设置
		if err != nil {
			return err
		}
		err = pfcpIe.Decode(ieData) // 当前IE的V解析
		if err != nil {
			return err
		}
		// 挂载当前对象
		err = root.SetObjectToParent(pfcpIe)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *IECreatePDR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)
	// Mandatory ie
	//PDR ID  	M
	// encode v
	vEnc, err := i.PDRID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Packet_Detection_Rule_ID), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	//Precedence	M
	// encode v
	vEnc, err = i.Precedence.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_Precedence), Length: uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	//PDI	        M
	// encode v
	vEnc, err = i.PDI.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_PDI), Length: uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Outer Header Removal 	C
		case IE_Outer_Header_Removal:
			vEnc, err = i.OuterHeaderRemoval.Encode()
			if err != nil {
				return nil, err
			}
			//FAR ID 	C
		case IE_FAR_ID:
			vEnc, err = i.FARID.Encode()
			if err != nil {
				return nil, err
			}
			//URR ID	C
		case IE_URR_ID:
			for _, v := range i.URRIDs[:len(i.URRIDs)-1] {
				//	encode v
				vEnc, err = v.Encode()
				if err != nil {
					return nil, err
				}
				// encode tl
				tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
				tlvEnc, err = tl.EncodeTlV(vEnc)
				if err != nil {
					return nil, err
				}
				_, err = encBuf.Write(tlvEnc)
				if err != nil {
					return nil, err
				}

			}
			// encode v
			vEnc, err = i.URRIDs[len(i.URRIDs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//QER ID 	C
		case IE_QER_ID:
			for _, v := range i.QERIDs[:len(i.QERIDs)-1] {
				//	encode v
				vEnc, err = v.Encode()
				if err != nil {
					return nil, err
				}
				// encode tl
				tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
				tlvEnc, err = tl.EncodeTlV(vEnc)
				if err != nil {
					return nil, err
				}
				_, err = encBuf.Write(tlvEnc)
				if err != nil {
					return nil, err
				}

			}
			// encode v
			vEnc, err = i.QERIDs[len(i.QERIDs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Activate Predefined Rules 	C
		case IE_Activate_Predefined_Rules:
			for _, v := range i.ActPredefinedRules[:len(i.ActPredefinedRules)-1] {
				//	encode v
				vEnc, err = v.Encode()
				if err != nil {
					return nil, err
				}
				// encode tl
				tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
				tlvEnc, err = tl.EncodeTlV(vEnc)
				if err != nil {
					return nil, err
				}
				_, err = encBuf.Write(tlvEnc)
				if err != nil {
					return nil, err
				}

			}
			// encode v
			vEnc, err = i.ActPredefinedRules[len(i.ActPredefinedRules)-1].Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")

		}
		// TL 编码
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IECreatePDR) Decode(data []byte) error {
	// parse v
	i.Grouped = true
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IECreatePDR) Len() int {
	return int(i.Length)
}

func (i *IECreatePDR) Set() {
	i.Type = IE_Create_PDR
}

//PDR ID
// 8.2.36	Packet Detection Rule ID (PDR ID)
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
56	Packet Detection Rule ID	Extendable / Subclause 8.2 36	2 */
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 56 (decimal)
3 to 4	Length = n
5 to 6	Rule ID
7to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.36-1: PDR ID
*/
type IEPDRID struct {
	IETypeLength
	RuleID uint16 `json:",omitempty"`
}

func (i *IEPDRID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEPDRID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPDRID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.RuleID)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEPDRID) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.RuleID)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEPDRID) Len() int {
	return int(i.Length)
}

func (i *IEPDRID) Set(v uint16) error {
	i.Type = IE_Packet_Detection_Rule_ID
	i.Length = 2
	i.RuleID = v
	return nil
}
func (i *IEPDRID) Get() (v uint16, e error) {
	return i.RuleID, nil
}

//Precedence
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
29	Precedence	Extendable / Subclause 8.2.11	4 */
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 29 (decimal)
3 to 4	Length = n
5 to 8	Precedence value
9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.11-1: Precedence
*/
type IEPrecedence struct {
	IETypeLength
	PrecedenceValue uint32 `json:",omitempty"`
}

func (i *IEPrecedence) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPrecedence) SetObjectToParent(child interface{}) error {

	return fmt.Errorf("implement me")
}

func (i *IEPrecedence) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	err = binary.Write(encBuf, binary.BigEndian, i.PrecedenceValue)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEPrecedence) Decode(data []byte) error {
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.PrecedenceValue)
	if err != nil {
		return err
	}

	return nil

}

func (i *IEPrecedence) Len() int {
	return int(i.Length)
}

func (i *IEPrecedence) Set(v uint32) {
	i.Type = IE_Precedence
	i.Length = 4
	i.PrecedenceValue = v
}

//PDI
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
2	PDI IE	Extendable / Table 7.5.2.2-2	Not Applicable */
/*Octet 1 and 2		PDI IE Type = 2 (decimal)
Octets 3 and 4		Length = n
Information elements	P

Source Interface	   M
Local F-TEID 	       O  F-TEID
Network Instance	   O
UE IP address 	       O
Traffic Endpoint ID	   C
SDF Filter	           O
Application ID	       O
Ethernet PDU Session Information	O
Ethernet Packet Filter	O
QFIs	O
Framed-Route	O
Framed-Routing	O
Framed-IPv6-Route	  O
*/
//NOTE 3:	SDF Filter IE(s) shall not be present if Ethernet Packet Filter IE(s) is present.
type IEPDI struct {
	IETypeLength
	SourceInterface   IESourceInterface
	LocalFTEID        *IEFTEID                  `json:",omitempty"`
	NetworkInstance   *IENetworkInstance        `json:",omitempty"`
	UEIPaddress       *IEUEIPaddress            `json:",omitempty"`
	TrafficEndpointID *IETrafficEndpointID      `json:",omitempty"`
	SDFFilters        []*IESDFFilter            `json:",omitempty"`
	ApplicationID     *IEApplicationID          `json:",omitempty"`
	EthPDUSessionInfo *IEEthernetPDUSessionInfo `json:",omitempty"`
	EthPacketFilters  []*IEEthernetPacketFilter `json:",omitempty"`
	QFIs              []*IEQFI                  `json:",omitempty"`
	FramedRoutes      []*IEFramedRoute          `json:",omitempty"`
	FramedRouting     *IEFramedRouting          `json:",omitempty"`
	FramedIPv6Routes  []*IEFramedIPv6Route      `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPDI) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPDI) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Source Interface	   M
	case *IESourceInterface:
		i.SourceInterface = *ie
		//Local F-TEID 	       O  F-TEID
	case *IEFTEID:
		i.LocalFTEID = ie
		//Network Instance	   O
	case *IENetworkInstance:
		i.NetworkInstance = ie
		//UE IP address 	       O
	case *IEUEIPaddress:
		i.UEIPaddress = ie
		//Traffic Endpoint ID	   C
	case *IETrafficEndpointID:
		i.TrafficEndpointID = ie
		//SDF Filter	           O
	case *IESDFFilter:
		i.SDFFilters = append(i.SDFFilters, ie)
		//Application ID	       O
	case *IEApplicationID:
		i.ApplicationID = ie
		//Ethernet PDU Session Information	O
	case *IEEthernetPDUSessionInfo:
		i.EthPDUSessionInfo = ie
		//Ethernet Packet Filter	O
	case *IEEthernetPacketFilter:
		i.EthPacketFilters = append(i.EthPacketFilters, ie)
		//QFIs	O
	case *IEQFI:
		i.QFIs = append(i.QFIs, ie)
		//Framed-Route	O
	case *IEFramedRoute:
		i.FramedRoutes = append(i.FramedRoutes, ie)
		//Framed-Routing	O
	case *IEFramedRouting:
		i.FramedRouting = ie
		//Framed-IPv6-Route	  O
	case *IEFramedIPv6Route:
		i.FramedIPv6Routes = append(i.FramedIPv6Routes)
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEPDI) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Source Interface	   M
	vEnc, err := i.SourceInterface.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Source_Interface), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Local F-TEID 	       O  F-TEID
		case IE_F_TEID:
			// encode v
			vEnc, err = i.LocalFTEID.Encode()
			if err != nil {
				return nil, err
			}
			//Network Instance	   O
		case IE_Network_Instance:
			// encode v
			vEnc, err = i.NetworkInstance.Encode()
			if err != nil {
				return nil, err
			}
			//UE IP address 	       O
		case IE_UE_IP_Address:
			// encode v
			vEnc, err = i.UEIPaddress.Encode()
			if err != nil {
				return nil, err
			}
			//Traffic Endpoint ID	   C
		case IE_Traffic_Endpoint_ID:
			vEnc, err = i.TrafficEndpointID.Encode()
			if err != nil {
				return nil, err
			}
			//SDF Filter	           O
		case IE_SDF_Filter:
			for _, v := range i.SDFFilters[:len(i.SDFFilters)-1] {
				// encode v
				vEnc, err = v.Encode()
				if err != nil {
					return nil, err
				}

				// TL 编码
				tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
				tlvEnc, err = tl.EncodeTlV(vEnc)
				if err != nil {
					return nil, err
				}
				_, err = encBuf.Write(tlvEnc)
				if err != nil {
					return nil, err
				}

			}
			// encode v
			vEnc, err = i.SDFFilters[len(i.SDFFilters)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Application ID	       O
		case IE_Application_ID:
			// encode v
			vEnc, err = i.ApplicationID.Encode()
			if err != nil {
				return nil, err
			}
			//Ethernet PDU Session Information	O
		case IE_Ethernet_PDU_Session_Information:
			// encode v
			vEnc, err = i.EthPDUSessionInfo.Encode()
			if err != nil {
				return nil, err
			}
			//Ethernet Packet Filter	O
		case IE_Ethernet_Packet_Filter:
			for _, v := range i.EthPacketFilters[:len(i.EthPacketFilters)-1] {
				// encode v
				vEnc, err = v.Encode()
				if err != nil {
					return nil, err
				}

				// encode tl
				tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
				tlvEnc, err = tl.EncodeTlV(vEnc)
				if err != nil {
					return nil, err
				}
				_, err = encBuf.Write(tlvEnc)
				if err != nil {
					return nil, err
				}
			}
			// encode v
			vEnc, err = i.EthPacketFilters[len(i.EthPacketFilters)-1].Encode()
			if err != nil {
				return nil, err
			}
			//QFIs	O
		case IE_QFI:
			for _, v := range i.QFIs[:len(i.QFIs)-1] {
				//	encode v
				vEnc, err = v.Encode()
				if err != nil {
					return nil, err
				}
				// encode tl
				tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
				tlvEnc, err = tl.EncodeTlV(vEnc)
				if err != nil {
					return nil, err
				}
				_, err = encBuf.Write(tlvEnc)
				if err != nil {
					return nil, err
				}

			}
			// encode v
			vEnc, err = i.QFIs[len(i.QFIs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Framed-Route	O
		case IE_Framed_Route:
			for _, v := range i.FramedRoutes[:len(i.FramedRoutes)-1] {
				//	encode v
				vEnc, err = v.Encode()
				if err != nil {
					return nil, err
				}
				// encode tl
				tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
				tlvEnc, err = tl.EncodeTlV(vEnc)
				if err != nil {
					return nil, err
				}
				_, err = encBuf.Write(tlvEnc)
				if err != nil {
					return nil, err
				}

			}
			// encode v
			vEnc, err = i.FramedRoutes[len(i.FramedRoutes)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Framed-Routing	O
		case IE_Framed_Routing:
			// encode v
			vEnc, err = i.FramedRouting.Encode()
			if err != nil {
				return nil, err
			}
			//Framed-IPv6-Route	  O
		case IE_Framed_IPv6_Route:
			for _, v := range i.FramedIPv6Routes[:len(i.FramedIPv6Routes)-1] {
				//	encode v
				vEnc, err = v.Encode()
				if err != nil {
					return nil, err
				}
				// encode tl
				tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
				tlvEnc, err = tl.EncodeTlV(vEnc)
				if err != nil {
					return nil, err
				}
				_, err = encBuf.Write(tlvEnc)
				if err != nil {
					return nil, err
				}

			}
			// encode v
			vEnc, err = i.FramedIPv6Routes[len(i.FramedIPv6Routes)-1].Encode()
			if err != nil {
				return nil, err
			}

		default:
			return nil, fmt.Errorf("Illegal IE")

		}
		// encode TL
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}
	return encBuf.Bytes(), nil
}

func (i *IEPDI) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEPDI) Len() int {
	return int(i.Length)
}

func (i *IEPDI) Set() {
	i.Type = IE_PDI
}

//Source Interface
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
20	Source Interface	Extendable / Subclause 8.2.2	1
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 20 (decimal)
3 to 4	Length = n
5 	           Spare	Interface value
6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.2-1: Source Interface
*/
type IESourceInterface struct {
	IETypeLength
	InterfaceValue uint8 `json:",omitempty"` // Table 8.2.2-1: Interface value
}

func (i *IESourceInterface) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	err = encBuf.WriteByte(i.InterfaceValue)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IESourceInterface) Decode(data []byte) error {
	r := bytes.NewReader(data)
	var tmp uint8
	err := binary.Read(r, binary.BigEndian, &tmp)
	v := utils.ByteOne(tmp)
	i.InterfaceValue, err = v.GetBits(1, 4)
	if err != nil {
		return err
	}

	return nil
}

func (i *IESourceInterface) Len() int {
	return int(i.Length)
}

func (i *IESourceInterface) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IESourceInterface) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

//Table 8.2.2-1: Interface value
/*Interface value	Values (Decimal)
Access	        0
Core	        1
SGi-LAN/N6-LAN	2
CP-function     3
Spare	        4 to 15
*/
const (
	Access            = 0
	Core              = 1
	SGi_LAN_or_N6_LAN = 2
	CP_function       = 3
)

func (i *IESourceInterface) Set(v uint8) {
	i.Type = IE_Source_Interface
	i.InterfaceValue = v
}

//Local F-TEID :F-TEID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
21	F-TEID	Extendable / Subclause 8.2.3	q-4
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 21 (decimal)
3 to 4	Length = n
5 	Spare	            CHID	CH	V6	V4
6 to 9 	TEID
m to (m+3) 	IPv4 address
p to (p+15) 	IPv6 address
q 	CHOOSE ID
k to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.3-1: F-TEID
*/
type IEFTEID struct {
	IETypeLength
	CHIDFlag uint8 `json:",omitempty"`
	CHFlag   uint8 `json:",omitempty"`
	// If this bit is set to "1" and the CH bit is not set,
	// the IPv6 address field shall be present
	V6Flag uint8 `json:",omitempty"`
	// If this bit is set to "1" and the CH bit is not set,
	// the IPv4 address field shall be present
	V4Flag   uint8          `json:",omitempty"`
	TEID     types3gpp.Teid `json:",omitempty"`
	IPv4Addr net.IP         `json:",omitempty"` // 4 byte
	IPv6Addr net.IP         `json:",omitempty"` // 16 byte

	CHOOSEID uint8 `json:",omitempty"`
}

func (i *IEFTEID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := (i.V4Flag) + (i.V6Flag << 1) + (i.CHFlag << 2) + (i.CHIDFlag << 3)
	err = binary.Write(encBuf, binary.BigEndian, tmpByte)
	if err != nil {
		return
	}

	//6 to 9 	TEID
	if i.CHFlag == 0 {
		err = binary.Write(encBuf, binary.BigEndian, uint32(i.TEID))
		if err != nil {
			return
		}
	}

	//m to (m+3) 	IPv4 address
	if i.CHFlag == 0 && i.V4Flag == 1 {
		err = binary.Write(encBuf, binary.BigEndian, []byte(i.IPv4Addr.To4()))
		if err != nil {
			return
		}
	}
	//p to (p+15) 	IPv6 address
	if i.CHFlag == 0 && i.V6Flag == 1 {
		err = binary.Write(encBuf, binary.BigEndian, []byte(i.IPv6Addr.To16()))
		if err != nil {
			return
		}
	}
	//q 	CHOOSE ID
	if i.CHFlag == 1 && i.CHIDFlag == 1 {
		err = binary.Write(encBuf, binary.BigEndian, i.CHOOSEID)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEFTEID) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	var tmp uint8
	err := binary.Read(r, binary.BigEndian, &tmp)
	if err != nil {
		return err
	}
	i.CHIDFlag, _ = utils.GetBitInByte(tmp, 4)
	i.CHFlag, _ = utils.GetBitInByte(tmp, 3)
	i.V6Flag, _ = utils.GetBitInByte(tmp, 2)
	i.V4Flag, _ = utils.GetBitInByte(tmp, 1)

	var tmp32 uint32
	err = binary.Read(r, binary.BigEndian, &tmp32)
	if err != nil {
		return err
	}
	i.TEID = types3gpp.Teid(tmp32)

	if i.V4Flag == 1 && i.CHFlag == 0 {
		var ip [4]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv4Addr = ip[:]
		return nil
	}
	if i.V6Flag == 1 && i.CHFlag == 0 {
		var ip [16]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv6Addr = ip[:]
		return nil
	}
	if i.CHFlag == 1 && i.CHIDFlag == 1 {
		//	then the TEID, IPv4 address and IPv6 address fields shall not be present
		err = binary.Read(r, binary.BigEndian, &i.CHOOSEID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *IEFTEID) Len() int {
	return int(i.Length)
}

func (i *IEFTEID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEFTEID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEFTEID) Set() {
	i.Type = IE_F_TEID

}

//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
22	Network Instance	Variable Length / Subclause 8.2.4	Not Applicable
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 22 (decimal)
3 to 4	Length = n
5 to (n+4)	Network Instance
Figure 8.2.4-1: Network Instance
*/
type IENetworkInstance struct {
	IETypeLength
	NetworkInstance string //*types3gpp.Apn `json:",omitempty"`
}

func (i *IENetworkInstance) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	_, err = encBuf.Write([]byte(i.NetworkInstance))
	if err != nil {
		return
	}
	return encBuf.Bytes(), nil
}

func (i *IENetworkInstance) Decode(data []byte) error {
	// parse v
	i.NetworkInstance = string(data)
	return nil
}

func (i *IENetworkInstance) Len() int {
	return int(i.Length)
}

func (i *IENetworkInstance) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IENetworkInstance) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IENetworkInstance) Set() {
	i.Type = IE_Network_Instance
}

//UE IP address 	       O
//IEUEIPaddress
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
93	UE IP Address	Extendable / Subclause 8.2.62	p+15-1
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 93 (decimal)
3 to 4	Length = n
5	Spare	           IPv6D	S/D	V4	V6
m to (m+3)	IPv4 address
p to (p+15)	IPv6 address
r	IPv6 Prefix Delegation Bits
k to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.62-1: UE IP Address
*/
type IEUEIPaddress struct {
	IETypeLength
	//5	Spare	           IPv6D	S/D	V4	V6
	IPv6D uint8 `json:",omitempty"`
	//Bit 3 – S/D, In the PDI IE, if this bit is set to "0", this indicates a Source IP address;
	// if this bit is set to "1", this indicates a Destination IP address.
	SD     uint8 `json:",omitempty"`
	V4Flag uint8 `json:",omitempty"`
	V6Flag uint8 `json:",omitempty"`

	IPv4Addr net.IP `json:",omitempty"` // 4 byte
	IPv6Addr net.IP `json:",omitempty"` // 16 byte

	IPv6PrefixDelegationBits uint8 `json:",omitempty"`
}

func (i *IEUEIPaddress) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	           IPv6D	S/D	V4	V6
	tmpByte := i.V6Flag +
		(i.V4Flag << 1) +
		(i.SD << 2) +
		(i.IPv6D << 3)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	//m to (m+3)	IPv4 address
	if i.V4Flag == 1 {
		encBuf.Write([]byte(i.IPv4Addr.To4()))
	}
	//p to (p+15)	IPv6 address
	if i.V6Flag == 1 {
		encBuf.Write([]byte(i.IPv6Addr.To16()))
	}
	//r	IPv6 Prefix Delegation Bits
	if i.V6Flag == 1 && i.IPv6D == 1 {
		encBuf.WriteByte(i.IPv6PrefixDelegationBits)
	}

	return encBuf.Bytes(), nil
}

func (i *IEUEIPaddress) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.IPv6D, _ = utils.GetBitInByte(tmp, 4)
	i.SD, _ = utils.GetBitInByte(tmp, 3)
	i.V4Flag, _ = utils.GetBitInByte(tmp, 2)
	i.V6Flag, _ = utils.GetBitInByte(tmp, 1)

	if i.V4Flag == 1 {
		var ip [4]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv4Addr = ip[:]
		return nil
	}
	if i.V6Flag == 1 {
		var ip [16]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv6Addr = ip[:]
		return nil
	}

	if i.V6Flag == 1 && i.IPv6D == 1 {
		i.IPv6PrefixDelegationBits, _ = r.ReadByte()
	}
	return nil
}

func (i *IEUEIPaddress) Len() int {
	return int(i.Length)
}

func (i *IEUEIPaddress) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUEIPaddress) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEUEIPaddress) Set() {
	i.Type = IE_UE_IP_Address
}

func (i *IEUEIPaddress) V4Set(ipAddr net.IP, sd bool) {
	i.Type = IE_UE_IP_Address
	i.Length = 5
	i.IPv6D = 0
	//Bit 3 – S/D: This bit is only applicable to the UE IP Address IE in the PDI IE.
	// It shall be set to "0" and ignored by the receiver in IEs other than PDI IE.
	// In the PDI IE, if this bit is set to "0", this indicates a Source IP address;
	// if this bit is set to "1", this indicates a Destination IP address.
	i.SD = 0
	if sd == true {
		i.SD = 1
	}
	i.V4Flag = 1
	i.V6Flag = 0
	i.IPv6PrefixDelegationBits = 0
	i.IPv4Addr = ipAddr
}

//Traffic Endpoint ID	   C
// IETrafficEndpointID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
131	Traffic Endpoint ID	Extendable / Subclause 8.2.92	1
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 131 (decimal)
3 to 4	Length = n
5 	Traffic Endpoint ID value
6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.92-1: Traffic Endpoint ID
*/
type IETrafficEndpointID struct {
	IETypeLength
	//The Traffic Endpoint ID value shall be encoded as a binary integer value
	// within the range of 0 to 255.
	Value uint8 `json:",omitempty"`
}

func (i *IETrafficEndpointID) Encode() (data []byte, err error) {
	// encode v
	encV := bytes.NewBuffer(nil)

	err = binary.Write(encV, binary.BigEndian, i.Value)
	if err != nil {
		return
	}

	return encV.Bytes(), nil
}

func (i *IETrafficEndpointID) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.Value = tmp
	return nil
}

func (i *IETrafficEndpointID) Len() int {
	return int(i.Length)
}

func (i *IETrafficEndpointID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IETrafficEndpointID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IETrafficEndpointID) Set(value uint8) error {
	i.Type = IE_Traffic_Endpoint_ID
	i.Length = 1
	i.Value = value
	return nil
}

func (i *IETrafficEndpointID) Get() (uint8, error) {
	return i.Value, nil
}

// SDF Filter	           O
// IESDFFilter //SDF	Service Data Flow
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
23	SDF Filter	Extendable / Subclause 8.2.5	v+2-4
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 23 (decimal)
3 to 4	Length = n
5	Spare	       BID	FL	SPI	TTC	FD
6	Spare
m to (m+1)	Length of Flow Description
(m+2) to p	Flow Description
s to (s+1)	ToS Traffic Class
t to (t+3)	Security Parameter Index
v to (v+2)	Flow Label
w to (w+3)	SDF Filter ID
x to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.5-1: SDF Filter
*/
type IESDFFilter struct {
	IETypeLength
	BIDFlag                 bool   `json:",omitempty"`
	FLFlag                  bool   `json:",omitempty"`
	SPIFlag                 bool   `json:",omitempty"`
	TTCFlag                 bool   `json:",omitempty"`
	FDFlag                  bool   `json:",omitempty"`
	LengthofFlowDescription uint16 `json:",omitempty"`
	FlowDescription         []byte `json:",omitempty"`
	ToSTrafficClass         uint16 `json:",omitempty"`
	SecurityParameterIndex  uint32 `json:",omitempty"`
	FlowLabel               uint32 `json:",omitempty"` // v to (v+2)
	SDFFilterID             uint32 `json:",omitempty"`
}

func (i *IESDFFilter) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	       BID	FL	SPI	TTC	FD
	tmpByte := utils.BoolToUint8(i.FDFlag) +
		(utils.BoolToUint8(i.TTCFlag) << 1) +
		(utils.BoolToUint8(i.SPIFlag) << 2) +
		(utils.BoolToUint8(i.FLFlag) << 3) +
		(utils.BoolToUint8(i.BIDFlag) << 4)

	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}
	//6	Spare
	err = encBuf.WriteByte(0)
	if err != nil {
		return
	}

	//m to (m+1)	Length of Flow Description
	if i.FDFlag {
		err = binary.Write(encBuf, binary.BigEndian, i.LengthofFlowDescription)
		if err != nil {
			return
		}
		//(m+2) to p	Flow Description
		err = binary.Write(encBuf, binary.BigEndian, i.FlowDescription)
		if err != nil {
			return
		}
	}

	//s to (s+1)	ToS Traffic Class
	if i.TTCFlag {
		err = binary.Write(encBuf, binary.BigEndian, i.ToSTrafficClass)
		if err != nil {
			return
		}
	}
	//t to (t+3)	Security Parameter Index
	if i.SPIFlag {
		err = binary.Write(encBuf, binary.BigEndian, i.SecurityParameterIndex)
		if err != nil {
			return
		}
	}
	//v to (v+2)	Flow Label
	if i.FLFlag {

		tmpByte3 := make([]byte, 3)
		tmpByte3[0] = byte(i.FlowLabel >> 16)
		tmpByte3[1] = byte(i.FlowLabel >> 8)
		tmpByte3[2] = byte(i.FlowLabel)

		err = binary.Write(encBuf, binary.BigEndian, tmpByte3)
		if err != nil {
			return
		}
	}
	//w to (w+3)	SDF Filter ID
	if i.BIDFlag {
		err = binary.Write(encBuf, binary.BigEndian, i.SDFFilterID)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IESDFFilter) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.BIDFlag, _ = utils.GetBoolInByte(tmp, 5)
	i.FLFlag, _ = utils.GetBoolInByte(tmp, 4)
	i.SPIFlag, _ = utils.GetBoolInByte(tmp, 3)
	i.TTCFlag, _ = utils.GetBoolInByte(tmp, 2)
	i.FDFlag, _ = utils.GetBoolInByte(tmp, 1)
	// 6	Spare
	tmp, err = r.ReadByte()
	if err != nil {
		return err
	}

	if i.FDFlag {
		//m to (m+1)	Length of Flow Description
		err := binary.Read(r, binary.BigEndian, &i.LengthofFlowDescription)
		if err != nil {
			return err
		}

		tmp := make([]byte, i.LengthofFlowDescription)
		_, err = r.Read(tmp)
		if err != nil {
			return err
		}
		i.FlowDescription = tmp
	}
	if i.TTCFlag {
		//s to (s+1)	ToS Traffic Class
		err := binary.Read(r, binary.BigEndian, &i.LengthofFlowDescription)
		if err != nil {
			return err
		}
	}
	if i.SPIFlag {
		//t to (t+3)	Security Parameter Index
		err := binary.Read(r, binary.BigEndian, &i.SecurityParameterIndex)
		if err != nil {
			return err
		}
	}

	if i.FLFlag {
		//v to (v+2)	Flow Label
		var tmp [3]byte
		_, err := r.Read(tmp[:])
		if err != nil {
			return err
		}
		i.FlowLabel = uint32(tmp[0])<<16 + uint32(tmp[1])<<8 + uint32(tmp[2])
	}

	if i.BIDFlag {
		//w to (w+3)	SDF Filter ID
		err := binary.Read(r, binary.BigEndian, &i.SDFFilterID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *IESDFFilter) Len() int {
	return int(i.Length)
}

func (i *IESDFFilter) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IESDFFilter) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IESDFFilter) SetFlowDesc(FlowDescription []byte) error {
	i.Type = IE_SDF_Filter
	i.FDFlag = true
	if len(FlowDescription) > Byte_2_MAX_Value {
		return errors.New("Flow Description Over length ")
	}
	//FlowDescription = make([]byte, len(FlowDescription))
	//i.LengthofFlowDescription = uint16(copy(i.FlowDescription, FlowDescription))
	i.FlowDescription = FlowDescription
	i.LengthofFlowDescription = uint16(len(FlowDescription))
	i.Length = 4 + i.LengthofFlowDescription
	return nil
}
func (i *IESDFFilter) SetTTC(ToSTrafficClass uint16) error {
	i.Type = IE_SDF_Filter
	i.TTCFlag = true
	i.ToSTrafficClass = ToSTrafficClass
	i.Length = 4 //s to (s+1)	ToS Traffic Class
	return nil
}
func (i *IESDFFilter) SetSPI(SecurityParameterIndex uint32) error {
	i.Type = IE_SDF_Filter
	i.SPIFlag = true
	i.SecurityParameterIndex = SecurityParameterIndex
	i.Length = 6 //t to (t+3)	Security Parameter Index
	return nil
}
func (i *IESDFFilter) SetFL(FlowLabel uint32) error {
	i.Type = IE_SDF_Filter
	i.FLFlag = true
	i.FlowLabel = FlowLabel
	i.Length = 5 //v to (v+2)	Flow Label
	return nil
}
func (i *IESDFFilter) SetBFD(SDFFilterID uint32) error {
	i.Type = IE_SDF_Filter
	i.BIDFlag = true
	i.SDFFilterID = SDFFilterID
	i.Length = 6 //w to (w+3)	SDF Filter ID
	return nil
}
func (i *IESDFFilter) Get() interface{} {
	/*Bit 1 – FD (Flow Description): If this bit is set to "1", then the
	Length of Flow Description and the Flow Description fields shall be
	present, otherwise they shall not be present.
	-	Bit 2 – TTC (ToS Traffic Class): If this bit is set to "1", then
	the ToS Traffic Class field shall be present, otherwise the ToS Traffic
	Class field shall not be present.
	-	Bit 3 – SPI (Security Parameter Index): If this bit is set to "1",
	then the Security Parameter Index field shall be present, otherwise the
	Security Parameter Index field shall not be present.
	-	Bit 4 – FL (Flow Label): If this bit is set to "1", then the Flow
	Label field shall be present, otherwise the Flow Label field shall not be present.
	-	Bit 5 – BID (Bidirectional SDF Filter): If this bit is set to "1",
	then the SDF Filter ID shall be present, otherwise the SDF Filter ID
	shall not be present.
	-	Bit 6 to 8: Spare, for future use and set to 0.*/
	if i.FDFlag == true {
		return i.FlowDescription
	}
	if i.TTCFlag == true {
		return i.ToSTrafficClass
	}
	if i.SPIFlag == true {
		return i.SecurityParameterIndex
	}
	if i.FLFlag == true {
		return i.FlowLabel
	}
	if i.BIDFlag == true {
		return i.SDFFilterID
	}
	return nil
}

//Application ID	       O
// IEApplicationID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
24	Application ID	Variable Length / Subclause 8.2.6	Not Applicable
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 24 (decimal)
3 to 4	Length = n
5 to (n+4)	Application Identifier
Figure 8.2.6-1: Application ID
*/
type IEApplicationID struct {
	IETypeLength
	//The Application Identifier shall be encoded as an OctetString (see 3GPP TS 29.212 [8])
	ApplicationIdentifier []byte `json:",omitempty"`
}

func (i *IEApplicationID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.ApplicationIdentifier)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEApplicationID) Decode(data []byte) error {
	// parse v
	i.ApplicationIdentifier = data
	return nil
}

func (i *IEApplicationID) Len() int {
	return int(i.Length)
}

func (i *IEApplicationID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEApplicationID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEApplicationID) Set() {
	i.Type = IE_Application_ID
}

//Ethernet PDU Session Information	O
// IEEthernetPDUSessionInformation
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
142	Ethernet PDU Session Information	Extendable / Subclause 8.2.102	1
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 142 (decimal)
3 to 4	Length = n
5	Spare	                       ETHI
k to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.102-1: Ethernet PDU Session Information
*/
type IEEthernetPDUSessionInfo struct {
	IETypeLength
	ETHI uint8 `json:",omitempty"`
}

func (i *IEEthernetPDUSessionInfo) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	err = encBuf.WriteByte(i.ETHI)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEEthernetPDUSessionInfo) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.ETHI, err = utils.GetBitInByte(tmp, 1)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEthernetPDUSessionInfo) Len() int {
	return int(i.Length)
}

func (i *IEEthernetPDUSessionInfo) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEthernetPDUSessionInfo) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEEthernetPDUSessionInfo) Set() {
	i.Type = IE_Ethernet_PDU_Session_Information
	i.Length = 1
}

//Ethernet Packet Filter	O
// IEEthernetPacketFilter
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
132	Ethernet Packet Filter	Extendable / Table 7.5.2.2-3	Not Applicable
*/
/*Octet 1 and 2		Ethernet Packet Filter IE Type = 132 (decimal)
Octets 3 and 4		Length = n
Information elements	P

Ethernet Filter ID	C
Ethernet Filter Properties	C
MAC address	O
Ethertype	O
C-TAG	O
S-TAG	O
SDF Filter	O
*/
type IEEthernetPacketFilter struct {
	IETypeLength
	EthernetFilterID         *IEEthernetFilterID         `json:",omitempty"`
	EthernetFilterProperties *IEEthernetFilterProperties `json:",omitempty"`
	MACaddress               *IEMACaddress               `json:",omitempty"`
	Ethertype                *IEEthertype                `json:",omitempty"`
	CTAG                     *IECTAG                     `json:",omitempty"`
	STAG                     *IESTAG                     `json:",omitempty"`
	SDFFilter                *IESDFFilter                `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEEthernetPacketFilter) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)
	var vEnc []byte
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Ethernet Filter ID	C
		case IE_Ethernet_Filter_ID:
			// encode v
			vEnc, err = i.EthernetFilterID.Encode()
			if err != nil {
				return nil, err
			}
			//Ethernet Filter Properties	C
		case IE_Ethernet_Filter_Properties:
			vEnc, err = i.EthernetFilterProperties.Encode()
			if err != nil {
				return nil, err
			}
			//MAC address	O
		case IE_MAC_address:
			vEnc, err = i.MACaddress.Encode()
			if err != nil {
				return nil, err
			}
			//Ethertype	O
		case IE_Ethertype:
			vEnc, err = i.Ethertype.Encode()
			if err != nil {
				return nil, err
			}
			//C-TAG	O
		case IE_C_TAG:
			vEnc, err = i.CTAG.Encode()
			if err != nil {
				return nil, err
			}
			//S-TAG	O
		case IE_S_TAG:
			vEnc, err = i.STAG.Encode()
			if err != nil {
				return nil, err
			}

			//SDF Filter	O
		case IE_SDF_Filter:
			vEnc, err = i.SDFFilter.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")

		}
		// TL 编码
		tl := IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err := tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEEthernetPacketFilter) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEthernetPacketFilter) Len() int {
	return int(i.Length)
}

func (i *IEEthernetPacketFilter) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEthernetPacketFilter) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Ethernet Filter ID	C
	case *IEEthernetFilterID:
		i.EthernetFilterID = ie
		//Ethernet Filter Properties	C
	case *IEEthernetFilterProperties:
		i.EthernetFilterProperties = ie
		//MAC address	O
	case *IEMACaddress:
		i.MACaddress = ie
		//Ethertype	O
	case *IEEthertype:
		i.Ethertype = ie
		//C-TAG	O
	case *IECTAG:
		i.CTAG = ie
		//S-TAG	O
	case *IESTAG:
		i.STAG = ie
		//SDF Filter	O
	case *IESDFFilter:
		i.SDFFilter = ie
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEEthernetPacketFilter) Set() {
	i.Type = IE_Ethernet_Packet_Filter
}

//Ethernet Filter ID
//IEEthernetFilterID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
138	Ethernet Filter ID	Extendable / Subclause 8.2.98	4
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 138 (decimal)
3 to 4	Length = n
5 to 8	Ethernet Filter ID value
10 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.98-1: Ethernet Filter ID
*/
type IEEthernetFilterID struct {
	IETypeLength
	Value uint32 `json:",omitempty"`
}

func (i *IEEthernetFilterID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEEthernetFilterID) Decode(data []byte) error {
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.BigEndian, &i.Value)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEthernetFilterID) Len() int {
	return int(i.Length)
}

func (i *IEEthernetFilterID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEthernetFilterID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEEthernetFilterID) Set(value uint32) {
	i.Type = IE_Ethernet_Filter_ID
	i.Length = 4
	i.Value = value
}

//Ethernet Filter Properties
//IEEthernetFilterProperties
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
139	Ethernet Filter Properties	Extendable / Subclause 8.2.99	1
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 139 (decimal)
3 to 4	Length = n
5	Spare	                       BIDE
6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.99-1: Ethernet Filter Properties
*/
type IEEthernetFilterProperties struct {
	IETypeLength
	//Bidirectional Ethernet Filter
	BIDE uint8 `json:",omitempty"`
}

func (i *IEEthernetFilterProperties) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	err = encBuf.WriteByte(i.BIDE)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEEthernetFilterProperties) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.BIDE, err = utils.GetBitInByte(tmp, 1)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEthernetFilterProperties) Len() int {
	return int(i.Length)
}

func (i *IEEthernetFilterProperties) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEthernetFilterProperties) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEEthernetFilterProperties) Set(value uint8) {
	i.Type = IE_Ethernet_Filter_Properties
	i.Length = 1
	i.BIDE = value
}

//MAC address
//IEMACaddress
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
133	MAC address	Extendable / Subclause 8.2.93	s-1-4
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 133 (decimal)
3 to 4	Length = n
5	spare	          UDES USOU DEST SOUR
m to (m+5)	Source MAC address value
n to (n+5)	Destination MAC address value
o to (o+5)	Upper Source MAC address value
p to (p+5)	Upper Destination MAC address value
s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.93-1: MAC address
*/
const (
	IEMACaddress_SOURFlag = 1
	IEMACaddress_DESTFlag = 2
	IEMACaddress_USOUFlag = 4
	IEMACaddress_UDESFlag = 8
)

type IEMACaddress struct {
	IETypeLength
	// SOURFlag =1 ,DESTFlag=2,USOUFlag=4,UDESFlag=8
	SOURFlag bool
	DESTFlag bool
	USOUFlag bool
	UDESFlag bool
	//Flag             uint8   `json:",omitempty"`
	SrcMACValue      [6]byte `json:",omitempty"`
	DstMACValue      [6]byte `json:",omitempty"`
	UpperSrcMACValue [6]byte `json:",omitempty"`
	UpperDstMACValue [6]byte `json:",omitempty"`
}

func (i *IEMACaddress) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	spare	          UDES USOU DEST SOUR
	tmpByte := utils.BoolToUint8(i.SOURFlag) +
		(utils.BoolToUint8(i.DESTFlag) << 1) +
		(utils.BoolToUint8(i.USOUFlag) << 2) +
		(utils.BoolToUint8(i.UDESFlag) << 3)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	//m to (m+5)	Source MAC address value
	if i.SOURFlag {
		_, err = encBuf.Write(i.SrcMACValue[:])
		if err != nil {
			return nil, err
		}
	}
	//n to (n+5)	Destination MAC address value
	if i.DESTFlag {
		_, err = encBuf.Write(i.DstMACValue[:])
		if err != nil {
			return nil, err
		}
	}
	//o to (o+5)	Upper Source MAC address value
	if i.USOUFlag {
		_, err = encBuf.Write(i.UpperSrcMACValue[:])
		if err != nil {
			return nil, err
		}
	}
	//p to (p+5)	Upper Destination MAC address value
	if i.UDESFlag {
		_, err = encBuf.Write(i.UpperDstMACValue[:])
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEMACaddress) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.SOURFlag, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}
	i.DESTFlag, err = utils.GetBoolInByte(tmp, 2)
	if err != nil {
		return err
	}
	i.USOUFlag, err = utils.GetBoolInByte(tmp, 3)
	if err != nil {
		return err
	}
	i.UDESFlag, err = utils.GetBoolInByte(tmp, 4)
	if err != nil {
		return err
	}

	if i.SOURFlag {
		_, err := r.Read(i.SrcMACValue[:])
		if err != nil {
			return err
		}
	}

	if i.DESTFlag {
		_, err := r.Read(i.DstMACValue[:])
		if err != nil {
			return err
		}
	}
	if i.USOUFlag {
		_, err := r.Read(i.UpperSrcMACValue[:])
		if err != nil {
			return err
		}
	}

	if i.UDESFlag {
		_, err := r.Read(i.UpperDstMACValue[:])
		if err != nil {
			return err
		}
	}
	return nil

}

func (i *IEMACaddress) Len() int {
	return int(i.Length)
}

func (i *IEMACaddress) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEMACaddress) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Bit 1 – SOUR (Source): If this bit is set to "1", then the source MAC address
value is provided.
-	Bit 2 – DEST (Destination): If this bit is set to "1", then the destination
MAC address value is provided.
-	Bit 3 – USOU (Source): If this bit is set to "1", then the source MAC address
value contains the lower value and Upper Source MAC address value contains the
upper value of an MAC address range.
-	Bit 4 – UDES (Destination): If this bit is set to "1", then the destination
MAC address value contains the lower value and Upper Destination MAC address
value contains the upper value of an MAC address range.-	Bit 5 to 8: Spare,
for future use and set to 0.
*/
func (i *IEMACaddress) SetMAC(MACValue [6]byte, flag uint8) error {
	i.Type = IE_MAC_address
	switch flag {
	case IEMACaddress_SOURFlag:
		i.Length = 7
		i.SrcMACValue = MACValue
	case IEMACaddress_DESTFlag:
		i.Length = 7
		i.SrcMACValue = MACValue
		// TODO
	case IEMACaddress_USOUFlag:
		return errors.New("Unsupported flag")
	case IEMACaddress_UDESFlag:
		return errors.New("Unsupported flag")
	}
	return nil
}

func (i *IEMACaddress) GetMAC(flag uint8) ([6]byte, error) {
	switch flag {
	case IEMACaddress_SOURFlag:
		return i.SrcMACValue, nil
	case IEMACaddress_DESTFlag:
		return i.SrcMACValue, nil
		// TODO
	case IEMACaddress_USOUFlag:
	case IEMACaddress_UDESFlag:

	}
	return [6]byte{}, errors.New("Unsupported flag")
}

//Ethertype
//IEEthertype
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
136	Ethertype	Extendable / Subclause 8.2.96	2
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 136 (decimal)
3 to 4	Length = n
5 to 6	Ethertype
7 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.u-1: Ethernetype
*/
type IEEthertype struct {
	IETypeLength
	Ethertype uint16 `json:",omitempty"`
}

func (i *IEEthertype) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	err = binary.Write(encBuf, binary.BigEndian, i.Ethertype)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEEthertype) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.BigEndian, &i.Ethertype)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEthertype) Len() int {
	return int(i.Length)
}

func (i *IEEthertype) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEthertype) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEEthertype) Set(Ethertype uint16) error {
	i.Type = IE_Ethertype
	i.Length = 2
	i.Ethertype = Ethertype
	return nil
}
func (i *IEEthertype) Get() (uint16, error) {
	return i.Ethertype, nil
}

//C-TAG
//IECTAG
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
134	C-TAG	Extendable / Subclause 8.2.94	3
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 134 (decimal)
3 to 4	Length = n
5	Spare	                VID	DEI	PCP
6 	C-VID value	   DEI Flag	PCP value
7	C-VID Value
8 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.94-1: C-TAG (Customer-VLAN tag)
*/
const (
	IEFlag_PCP = 1
	IEFlag_DEI = 2
	IEFlag_VID = 4
)

type IECTAG struct {
	IETypeLength
	PCP bool
	DEI bool
	VID bool
	//Flag      uint8  `json:",omitempty"`
	PCPValue  uint8  `json:",omitempty"`
	DEIFlag   uint8  `json:",omitempty"`
	CVIDValue uint16 `json:",omitempty"`
}

func (i *IECTAG) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	                VID	DEI	PCP
	tmpByte := utils.BoolToUint8(i.PCP) +
		(utils.BoolToUint8(i.DEI) << 1) +
		(utils.BoolToUint8(i.VID) << 2)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	//6 	C-VID value	DEI Flag	PCP value
	tmpByte = i.PCPValue + (i.DEIFlag << 3) +
		(byte(i.CVIDValue>>8) << 4)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}
	//specified in IEEE 802.1Q [30] tag format
	//VID：即VLAN的标志符，占12个比特
	//7	C-VID Value ,为低位字节
	tmpByte = byte(i.CVIDValue)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IECTAG) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.PCP, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}
	i.DEI, err = utils.GetBoolInByte(tmp, 2)
	if err != nil {
		return err
	}
	i.VID, err = utils.GetBoolInByte(tmp, 3)
	if err != nil {
		return err
	}
	//6 	C-VID value	DEI Flag	PCP value
	tmp6, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp6)
	if i.PCP {

		i.PCPValue, err = v.GetBits(1, 3)
		if err != nil {
			return err
		}
	}
	if i.DEI {
		i.DEIFlag, err = v.GetBit(4)
		if err != nil {
			return err
		}
	}

	if i.VID {
		tmp7, err := v.GetBits(5, 8)
		if err != nil {
			return err
		}
		v1 := utils.ByteOne(tmp7)
		CVID1, _ := v1.RightShift(5)
		//7	C-VID Value
		CVID2, err := r.ReadByte()
		if err != nil {
			return err
		}
		//specified in IEEE 802.1Q [30] tag format
		//VID：即VLAN的标志符，占12个比特
		i.CVIDValue = uint16(CVID1)<<8 + uint16(CVID2)
	}

	return nil
}

func (i *IECTAG) Len() int {
	return int(i.Length)
}

func (i *IECTAG) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECTAG) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*-	Bit 1 – PCP: If this bit is set to "1", then PCP Value field shall used by
the receiver, otherwise the PCP Value field shall be ignored.
-	Bit 2 – DEI: If this bit is set to "1", then DEI flag field shall used by
the receiver, otherwise the DEI flag field shall be ignored.
-	Bit 3 – VID: If this bit is set to "1", then C-VID value field shall used
by the receiver, otherwise the VID Value field shall be ignored.
-	Bit 4 to 8 – spare and reserved for future use.
*/
func (i *IECTAG) Set(v uint16, flag uint8) (e error) {
	i.Type = IE_C_TAG
	i.Length = 3
	switch flag {
	case IEFlag_PCP:
		i.PCPValue = uint8(v)
	case IEFlag_DEI:
		i.DEIFlag = uint8(v)
	case IEFlag_VID:
		i.CVIDValue = v
	}
	return nil
}

func (i *IECTAG) Get() (v interface{}, e error) {
	/*switch i.Flag {
	case IEFlag_PCP:
		return i.PCPValue, nil
	case IEFlag_DEI:
		return i.DEIFlag, nil
	case IEFlag_VID:
		return i.SVIDValue, nil
	}*/
	return nil, nil
}

//S-TAG
//IESTAG
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
135	S-TAG	Extendable / Subclause 8.2.95	3
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 135 (decimal)
3 to 4	Length = n
5	Spare	VID	DEI	PCP
6	S-VID value	DEI Flag	PCP value
7	S-VID value
8 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.95-1: S-TAG (Service-VLAN tag)
*/
type IESTAG struct {
	IETypeLength
	PCP bool
	DEI bool
	VID bool
	//Flag      uint8  `json:",omitempty"`
	PCPValue  uint8  `json:",omitempty"`
	DEIFlag   uint8  `json:",omitempty"`
	SVIDValue uint16 `json:",omitempty"`
}

func (i *IESTAG) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	                VID	DEI	PCP
	tmpByte := utils.BoolToUint8(i.PCP) +
		(utils.BoolToUint8(i.DEI) << 1) +
		(utils.BoolToUint8(i.VID) << 2)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	//6	S-VID value	DEI Flag	PCP value
	tmpByte = i.PCPValue + (i.DEIFlag << 3) +
		(byte(i.SVIDValue>>8) << 4)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}
	//specified in IEEE 802.1Q [30] tag format
	//VID：即VLAN的标志符，占12个比特
	//7	S-VID value ,为低位字节
	tmpByte = byte(i.SVIDValue)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IESTAG) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.PCP, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}
	i.DEI, err = utils.GetBoolInByte(tmp, 2)
	if err != nil {
		return err
	}
	i.VID, err = utils.GetBoolInByte(tmp, 3)
	if err != nil {
		return err
	}
	//6 	C-VID value	DEI Flag	PCP value
	tmp6, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp6)
	if i.PCP {

		i.PCPValue, err = v.GetBits(1, 3)
		if err != nil {
			return err
		}
	}
	if i.DEI {
		i.DEIFlag, err = v.GetBit(4)
		if err != nil {
			return err
		}
	}

	if i.VID {
		tmp7, err := v.GetBits(5, 8)
		if err != nil {
			return err
		}
		v1 := utils.ByteOne(tmp7)
		CVID1, _ := v1.RightShift(5)
		//7	C-VID Value
		CVID2, err := r.ReadByte()
		if err != nil {
			return err
		}
		//specified in IEEE 802.1Q [30] tag format
		//VID：即VLAN的标志符，占12个比特
		i.SVIDValue = uint16(CVID1)<<8 + uint16(CVID2)
	}
	return nil
}

func (i *IESTAG) Len() int {
	return int(i.Length)
}

func (i *IESTAG) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IESTAG) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Bit 1 – PCP: If this bit is set to "1", then PCP Value field shall used by
the receiver, otherwise the PCP Value field shall be ignored.
-	Bit 2 – DEI: If this bit is set to "1", then DEI flag field shall used by
the receiver, otherwise the DEI flag field shall be ignored.
-	Bit 3 – VID: If this bit is set to "1", then VID value field shall used by
the receiver, otherwise the VID Value field shall be ignored.
-	Bit 4 to 8 – spare and reserved for future use.
*/
func (i *IESTAG) Set(v interface{}, flag uint8) (e error) {
	i.Type = IE_S_TAG
	i.Length = 3

	switch flag {
	case IEFlag_PCP:
		if value, ok := v.(uint8); ok {
			i.PCPValue = value
			return nil
		}
	case IEFlag_DEI:
		if value, ok := v.(uint8); ok {
			i.DEIFlag = value
			return nil
		}
	case IEFlag_VID:
		if value, ok := v.(uint16); ok {
			i.SVIDValue = value
			return nil
		}
	}
	return errors.New("Unsupported flag")
}

func (i *IESTAG) Get() (v interface{}, e error) {
	/*switch i.Flag {
	case IEFlag_PCP:
		return i.PCPValue, nil
	case IEFlag_DEI:
		return i.DEIFlag, nil
	case IEFlag_VID:
		return i.SVIDValue, nil
	}*/
	return nil, nil
}

//IEEthernetPacketFilter
//End--------------------------------------------------------------------------

//QFIs	O
// IEQFI
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
124	QFIs	Extendable / Subclause 8.2.89	m-4
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 124 (decimal)
3 to 4	Length = n
5	Spare	QFIs value
p to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.89-1: QFIs
*/
type IEQFI struct {
	IETypeLength
	Value uint8 `json:",omitempty"` //6 bits QoS Flow Identifier
}

func (i *IEQFI) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	err = encBuf.WriteByte(i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEQFI) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	i.Value, err = v.GetBits(1, 6)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEQFI) Len() int {
	return int(i.Length)
}

func (i *IEQFI) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEQFI) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEQFI) Get() (v uint8, e error) {
	return i.Value, nil
}

func (i *IEQFI) Set(v uint8) error {
	i.Type = IE_QFI
	i.Length = 1
	i.Value = v
	return nil
}

//Framed-Route	O
// IEFramedRoute
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
153	Framed-Route	Variable Length / Subclause 8.2.109	Not Applicable
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 153 (decimal)
3 to 4	Length = n
5 to (n+4)	Framed-Route
Figure 8.2.109-1: Framed-Route
*/
type IEFramedRoute struct {
	IETypeLength
	FramedRoute string `json:",omitempty"`
}

func (i *IEFramedRoute) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	_, err = encBuf.Write([]byte(i.FramedRoute))
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEFramedRoute) Decode(data []byte) error {
	//parse v
	i.FramedRoute = string(data)
	return nil
}

func (i *IEFramedRoute) Len() int {
	return int(i.Length)
}

func (i *IEFramedRoute) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEFramedRoute) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*func (i *IEFramedRoute) Set(v []byte) error {
	i.Type = IE_Framed_Route
	i.Length = uint16(len(v))
	i.FramedRoutes = make([]byte, len(v))
	copy(i.FramedRoutes, v)
	return nil
}
*/
func (i *IEFramedRoute) Set(v string) error {
	i.Type = IE_Framed_Route
	i.FramedRoute = v
	i.Length = uint16(len(v))
	return nil
}
func (i *IEFramedRoute) Get() (v string, e error) {
	return i.FramedRoute, nil
}

//Framed-Routing	O
// IEFramedRouting
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
154	Framed-Routing	Fixed Length / Subclause 8.2.110	4
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 154 (decimal)
3 to 4	Length = 4
5 to 8	Framed-Routing
Figure 8.2.110-1: Framed-Routing
*/

type IEFramedRouting struct {
	IETypeLength
	FramedRouting uint32 `json:",omitempty"`
}

func (i *IEFramedRouting) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.FramedRouting)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEFramedRouting) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.FramedRouting)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEFramedRouting) Len() int {
	return int(i.Length)
}

func (i *IEFramedRouting) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEFramedRouting) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEFramedRouting) Set(v uint32) error {
	i.Type = IE_Framed_Routing
	i.Length = 4
	i.FramedRouting = v
	return nil
}
func (i *IEFramedRouting) Get() (v uint32, e error) {
	return i.FramedRouting, nil
}

//Framed-IPv6-Route	  O
// IEFramedIPv6Route
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 155 (decimal)
	3 to 4	Length = n
	5 to (n+4)	Framed-IPv6-Route
Figure 8.2.z-1: Framed-IPv6-Route
The Framed-IPv6-Route field shall be encoded as an Octet String as the value
part of the Framed-IPv6-Route AVP specified in IETF RFC 3162 [38].
*/
type IEFramedIPv6Route struct {
	IETypeLength
	FramedIPv6Route string `json:",omitempty"`
}

func (i *IEFramedIPv6Route) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	_, err = encBuf.Write([]byte(i.FramedIPv6Route))
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEFramedIPv6Route) Decode(data []byte) error {
	//parse v
	i.FramedIPv6Route = string(data)
	return nil
}

func (i *IEFramedIPv6Route) Len() int {
	return int(i.Length)
}

func (i *IEFramedIPv6Route) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEFramedIPv6Route) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEFramedIPv6Route) Set(v string) error {
	i.Type = IE_Framed_IPv6_Route
	i.FramedIPv6Route = v
	i.Length = uint16(len(v))
	return nil
}
func (i *IEFramedIPv6Route) Get() (v string, e error) {
	return i.FramedIPv6Route, nil
}

//IEPDI
//End--------------------------------------------------------------------------

//Outer Header Removal
// IEOuterHeaderRemoval
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
95	Outer Header Removal	Extendable / Subclause 8.2.64	1
*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 95 (decimal)
3 to 4	Length = n
5	Outer Header Removal Description
6	GTP-U Extension Header Deletion
7 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.64-1: Outer Header Removal
*/
type IEOuterHeaderRemoval struct {
	IETypeLength
	Description                 uint8 `json:",omitempty"`
	GTPUExtensionHeaderDeletion uint8
}

func (i *IEOuterHeaderRemoval) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Outer Header Removal Description
	err = encBuf.WriteByte(i.Description)
	if err != nil {
		return
	}
	//todo  GTP-U Extension Header Deletion
	/*//6	GTP-U Extension Header Deletion ,5GS to EPS handover
	err = encBuf.WriteByte(i.GTPUExtensionHeaderDeletion)
	if err != nil {
		return
	}*/

	return encBuf.Bytes(), nil
}

func (i *IEOuterHeaderRemoval) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)

	if len(data) == 1 {
		tmp, err := r.ReadByte()
		if err != nil {
			return err
		}
		i.Description = tmp
	}
	if len(data) > 1 {
		tmp, err := r.ReadByte()
		if err != nil {
			return err
		}
		i.GTPUExtensionHeaderDeletion = tmp
	}
	return nil
}

func (i *IEOuterHeaderRemoval) Len() int {
	return int(i.Length)
}

func (i *IEOuterHeaderRemoval) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEOuterHeaderRemoval) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Table 8.2.64-1: Outer Header Removal Description
Outer Header to be removed from the incoming packet	Value (Decimal)
GTP-U/UDP/IPv4 (see NOTE 1)	0
GTP-U/UDP/IPv6 (see NOTE 1)	1
UDP/IPv4 (See NOTE 3)	2
UDP/IPv6 (See NOTE 3)	3
IPv4 (NOTE 6)	4
IPv6 (NOTE 6)	5
GTP-U/UDP/IP (NOTE 4)	6
VLAN S-TAG (See NOTE 5)	7
S-TAG and C-TAG (See NOTE 5)	8
For future use. Shall not be sent. If received, shall be interpreted as the value "1".	9 to 255
*/
type OuterHeaderRemoval uint8

const (
	IEOuterHRemoval_GTPU_UDP_IPv4 OuterHeaderRemoval = 0
	IEOuterHRemoval_GTPU_UDP_IPv6 OuterHeaderRemoval = 1
	IEOuterHRemoval_UDP_IPv4      OuterHeaderRemoval = 2
	IEOuterHRemoval_UDP_IPv6      OuterHeaderRemoval = 3

	IEOuterHRemoval_IPv4        OuterHeaderRemoval = 4
	IEOuterHRemoval_IPv6        OuterHeaderRemoval = 5
	IEOuterHRemoval_GTPU_UDP_IP OuterHeaderRemoval = 6
	IEOuterHRemoval_VLAN_STAG   OuterHeaderRemoval = 7
	IEOuterHRemoval_STAG_CTAG   OuterHeaderRemoval = 8
)

func (i *IEOuterHeaderRemoval) SetDescription(v uint8) error {
	if OuterHeaderRemoval(v) < IEOuterHRemoval_GTPU_UDP_IPv4 ||
		OuterHeaderRemoval(v) > IEOuterHRemoval_STAG_CTAG {
		return errors.New("Unsupported  type value")
	}

	i.Type = IE_Outer_Header_Removal
	i.Length = 2
	i.Description = v
	return nil
}
func (i *IEOuterHeaderRemoval) GetDescription() (v uint8, e error) {
	return i.Description, nil
}
func (i *IEOuterHeaderRemoval) SetGTPUExtensionHeaderDel(v uint8) error {
	if v > 1 {
		return errors.New("Unsupported  type value")
	}

	i.Type = IE_Outer_Header_Removal
	i.Length = 2
	i.GTPUExtensionHeaderDeletion = v
	return nil
}
func (i *IEOuterHeaderRemoval) GetGTPUExtensionHeaderDel() (uint8, error) {
	return i.GTPUExtensionHeaderDeletion, nil
}

//FAR ID
// IEFARID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
108	FAR ID	Extendable / Subclause 8.2.74	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 108 (decimal)
	3 to 4	Length = n
	5 to 8	FAR ID value
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.74-1: FAR ID
*/

type IEFARID struct {
	IETypeLength
	Value uint32 `json:",omitempty"`
}

func (i *IEFARID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEFARID) Decode(data []byte) error {
	//parse v
	//i.Value = binary.BigEndian.Uint32(data)

	r := bytes.NewReader(data)
	err := binary.Read(r, binary.BigEndian, &i.Value)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEFARID) Len() int {
	return int(i.Length)
}

func (i *IEFARID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEFARID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEFARID) Set(v uint32) error {
	i.Type = IE_FAR_ID
	i.Length = 4
	i.Value = v

	return nil
}
func (i *IEFARID) Get() (v uint32, e error) {
	return i.Value, nil
}

/*The bit 8 of octet 5 is used to indicate if the Rule ID is dynamically
allocated by the CP function or predefined in the UP function. If set to 0,
it indicates that the Rule is dynamically provisioned by the CP Function.
If set to 1, it indicates that the Rule is predefined in the UP Function. */
func (i *IEFARID) GetIndicatesRule() (r uint8, e error) {
	return utils.GetBitInUint32(i.Value, 32)
}

//QER ID
// IEQERID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
109	QER ID	Extendable / Subclause 8.2.75	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 109 (decimal)
	3 to 4	Length = n
	5 to 8	QER ID value
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.75-1: QER ID
*/
type IEQERID struct {
	IETypeLength
	Value uint32 `json:",omitempty"`
}

func (i *IEQERID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEQERID) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.BigEndian, &i.Value)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEQERID) Len() int {
	return int(i.Length)
}

func (i *IEQERID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l

	return nil
}

func (i *IEQERID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEQERID) Set(v uint32) error {
	i.Type = IE_QER_ID
	i.Length = 4
	i.Value = v
	return nil
}
func (i *IEQERID) Get() (v uint32, e error) {
	return i.Value, nil
}

//Activate Predefined Rules
// IEActivatePredefinedRules
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
106	Activate Predefined Rules 	Variable Length / Subclause 8.2.72	Not Applicable
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 106 (decimal)
	3 to 4	Length = n
	5 to (n+4)	Predefined Rules Name

Figure 8.2.72-1: Activate Predefined Rules
*/
type IEActivatePredefinedRules struct {
	IETypeLength
	RulesName string `json:",omitempty"`
}

func (i *IEActivatePredefinedRules) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	_, err = encBuf.Write([]byte(i.RulesName))
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEActivatePredefinedRules) Decode(data []byte) error {
	//	parse
	i.RulesName = string(data)
	return nil
}

func (i *IEActivatePredefinedRules) Len() int {
	return int(i.Length)
}

func (i *IEActivatePredefinedRules) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEActivatePredefinedRules) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEActivatePredefinedRules) Set(v string) error {
	i.Type = IE_Activate_Predefined_Rules
	i.Length = uint16(len(v))
	i.RulesName = v
	return nil
}
func (i *IEActivatePredefinedRules) Get() (v string, e error) {
	return i.RulesName, nil
}

// IECreatePDR
//End--------------------------------------------------------------------------

//Create FAR	M //Forwarding Action Rule
// IECreateFAR
/*Table 7.5.2.3-1: Create FAR IE within PFCP Session Establishment Request
Octet 1 and 2		Create FAR IE Type = 3 (decimal)
Octets 3 and 4		Length = n
Information elements	P

FAR ID	        M
Apply Action	M
Forwarding Parameters	C
Duplicating Parameters 	C
BAR ID	O

*/
type IECreateFAR struct {
	IETypeLength
	FARID                 IEFARID
	ApplyAction           IEApplyAction
	ForwardingParameters  *IEForwardingParameters  `json:",omitempty"`
	DuplicatingParameters *IEDuplicatingParameters `json:",omitempty"`
	BARID                 *IEBARID                 `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IECreateFAR) Encode() (data []byte, err error) {
	//	encode ie
	encBuf := bytes.NewBuffer(nil)

	//	Mandatory ie
	//	FAR ID	        M
	// encode v
	vEnc, err := i.FARID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_FAR_ID), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	//	Apply Action	M
	// encode v
	vEnc, err = i.ApplyAction.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_Apply_Action), Length: uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	//  optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//	Forwarding Parameters	C
		case IE_Forwarding_Parameters:
			// encode v
			vEnc, err = i.ForwardingParameters.Encode()
			if err != nil {
				return nil, err
			}
			//	Duplicating Parameters 	C
		case IE_Duplicating_Parameters:
			// encode v
			vEnc, err = i.DuplicatingParameters.Encode()
			if err != nil {
				return nil, err
			}
			//	BAR ID	O
		case IE_BAR_ID:
			// encode v
			vEnc, err = i.BARID.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")
		}
		// encode TL
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}
	return encBuf.Bytes(), nil
}

func (i *IECreateFAR) Decode(data []byte) error {
	//parse v
	i.Grouped = true
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IECreateFAR) Len() int {
	return int(i.Length)
}

func (i *IECreateFAR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECreateFAR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//FAR ID	        M
	case *IEFARID:
		i.FARID = *ie
		//Apply Action	M
	case *IEApplyAction:
		i.ApplyAction = *ie
		//Forwarding Parameters	C
	case *IEForwardingParameters:
		i.ForwardingParameters = ie
		//Duplicating Parameters 	C
	case *IEDuplicatingParameters:
		i.DuplicatingParameters = ie
		//BAR ID	O
	case *IEBARID:
		i.BARID = ie

	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IECreateFAR) Set() error {
	i.Type = IE_Create_FAR

	return nil
}

//Apply Action
// IEApplyAction
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
44	Apply Action	Extendable / Subclause 8.2.26	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 44 (decimal)
	3 to 4	Length = n
	5	Spare	Spare	Spare	DUPL	NOCP	BUFF	FORW	DROP
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.26-1: Apply-Action
*/

type IEApplyAction struct {
	IETypeLength
	Flag uint8 `json:",omitempty"`
}

func (i *IEApplyAction) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Flag)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEApplyAction) Decode(data []byte) error {
	//parse v
	//todo 业务在使用，没有修改定义格式
	r := bytes.NewReader(data)
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.Flag = tmp

	return nil
}

func (i *IEApplyAction) Len() int {
	return int(i.Length)
}

func (i *IEApplyAction) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEApplyAction) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

const (
	DROP uint8 = IEFlag_Bit_1
	FORW uint8 = IEFlag_Bit_2
	BUFF uint8 = IEFlag_Bit_3
	NOCP uint8 = IEFlag_Bit_4
	DUPL uint8 = IEFlag_Bit_5
)

/*-	Bit 1 – DROP (Drop): when set to 1, this indicates a request to drop the
packets.
-	Bit 2 – FORW (Forward): when set to 1, this indicates a request to
forward the packets.
-	Bit 3 – BUFF (Buffer): when set to 1, this indicates a request to
buffer the packets.
-	Bit 4 – NOCP (Notify the CP function): when set to 1, this indicates
a request to notify the CP function about the arrival of a first downlink
packet being buffered.
-	Bit 5 – DUPL (Duplicate): when set to 1, this indicates a request
to duplicate the packets.
-	Bit 6 to 8 – Spare, for future use and set to 0.
One and only one of the DROP, FORW and BUFF flags shall be set to 1.
The NOCP flag may only be set if the BUFF flag is set.
The DUPL flag may be set with any of the DROP, FORW, BUFF and NOCP flags.
*/
func (i *IEApplyAction) Set(v uint8) error {
	i.Type = IE_Apply_Action
	i.Length = 1
	switch v { // 安全性检查
	case IEFlag_Bit_1, IEFlag_Bit_2, IEFlag_Bit_3:
		i.Flag = v
		//fallthrough
	case IEFlag_Bit_4 | IEFlag_Bit_3:
		i.Flag = v
	case IEFlag_Bit_5 | IEFlag_Bit_1,
		IEFlag_Bit_5 | IEFlag_Bit_2,
		IEFlag_Bit_5 | IEFlag_Bit_3,
		IEFlag_Bit_5 | IEFlag_Bit_4:
		i.Flag = v
	default:
		return errors.New("Unsupported  type value")
	}
	return nil
}

func (i *IEApplyAction) Get() (v uint8, e error) {
	return i.Flag, nil
}

//Forwarding Parameters
//IEForwardingParameters
/*Table 7.5.2.3-2: Forwarding Parameters IE in FAR
Octet 1 and 2		Forwarding Parameters IE Type = 4 (decimal)
Octets 3 and 4		Length = n
Information elements	P

Destination Interface	M
Network Instance	    O
Redirect Information	C
Outer Header Creation 	C
Transport Level Marking 	C
Forwarding Policy 	C
Header Enrichment	O
Linked Traffic Endpoint ID	C
Proxying	C


*/
type IEForwardingParameters struct {
	IETypeLength
	DstInterface            IEDestinationInterface
	NetworkInstance         *IENetworkInstance       `json:",omitempty"`
	RedirectInfo            *IERedirectInformation   `json:",omitempty"`
	OuterHeaderCreation     *IEOuterHeaderCreation   `json:",omitempty"`
	TransportLevelMarking   *IETransportLevelMarking `json:",omitempty"`
	ForwardingPolicy        *IEForwardingPolicy      `json:",omitempty"`
	HeaderEnrichment        *IEHeaderEnrichment      `json:",omitempty"`
	LinkedTrafficEndpointID *IETrafficEndpointID     `json:",omitempty"`
	Proxying                *IEProxying              `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEForwardingParameters) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Destination Interface	M
	vEnc, err := i.DstInterface.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Destination_Interface), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Network Instance	    O
		case IE_Network_Instance:
			// encode v
			vEnc, err = i.NetworkInstance.Encode()
			if err != nil {
				return nil, err
			}
			//Redirect Information	C
		case IE_Redirect_Information:
			// encode v
			vEnc, err = i.RedirectInfo.Encode()
			if err != nil {
				return nil, err
			}

			//Outer Header Creation 	C
		case IE_Outer_Header_Creation:
			// encode v
			vEnc, err = i.OuterHeaderCreation.Encode()
			if err != nil {
				return nil, err
			}
			//Transport Level Marking 	C
		case IE_Transport_Level_Marking:
			// encode v
			vEnc, err = i.TransportLevelMarking.Encode()
			if err != nil {
				return nil, err
			}
			//Forwarding Policy 	C
		case IE_Forwarding_Policy:
			// encode v
			vEnc, err = i.ForwardingPolicy.Encode()
			if err != nil {
				return nil, err
			}
			//Header Enrichment	O
		case IE_Header_Enrichment:
			// encode v
			vEnc, err = i.HeaderEnrichment.Encode()
			if err != nil {
				return nil, err
			}
			//Linked Traffic Endpoint ID	C
		case IE_Traffic_Endpoint_ID:
			// encode v
			vEnc, err = i.LinkedTrafficEndpointID.Encode()
			if err != nil {
				return nil, err
			}
			//Proxying	C
		case IE_Proxying:
			// encode v
			vEnc, err = i.Proxying.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")
		}
		// encode TL
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEForwardingParameters) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEForwardingParameters) Len() int {
	return int(i.Length)
}

func (i *IEForwardingParameters) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEForwardingParameters) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Destination Interface	M
	case *IEDestinationInterface:
		i.DstInterface = *ie
		//Network Instance	    O
	case *IENetworkInstance:
		i.NetworkInstance = ie
		//Redirect Information	C
	case *IERedirectInformation:
		i.RedirectInfo = ie
		//Outer Header Creation 	C
	case *IEOuterHeaderCreation:
		i.OuterHeaderCreation = ie
		//Transport Level Marking 	C
	case *IETransportLevelMarking:
		i.TransportLevelMarking = ie
		//Forwarding Policy 	C
	case *IEForwardingPolicy:
		i.ForwardingPolicy = ie
		//Header Enrichment	O
	case *IEHeaderEnrichment:
		i.HeaderEnrichment = ie
		//Linked Traffic Endpoint ID	C
	case *IETrafficEndpointID:
		i.LinkedTrafficEndpointID = ie
		//Proxying	C
	case *IEProxying:
		i.Proxying = ie
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEForwardingParameters) Set() error {
	i.Type = IE_Forwarding_Parameters

	return nil
}

//Destination Interface	M
// IEDestinationInterface
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
42	Destination Interface	Extendable / Subclause 8.2.24	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 42 (decimal)
	3 to 4	Length = n
	5 	Spare	           Interface value
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.24-1: Destination Interface
*/
type IEDestinationInterface struct {
	IETypeLength
	Value uint8 `json:",omitempty"`
}

func (i *IEDestinationInterface) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = encBuf.WriteByte(i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEDestinationInterface) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	i.Value, err = v.GetBits(1, 4)
	return nil
}

func (i *IEDestinationInterface) Len() int {
	return int(i.Length)
}

func (i *IEDestinationInterface) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEDestinationInterface) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Table 8.2.24-1: Interface value
Interface value	    Values (Decimal)

Access (see NOTE 1)	0
Core (see NOTE 1)	1
SGi-LAN/N6-LAN	    2
CP- Function	    3
LI Function (see NOTE 2)	4
Spare	5 to 15
*/
const (
	IEDestinationInterface_Access          uint8 = 0
	IEDestinationInterface_Core            uint8 = 1
	IEDestinationInterface_SGiLAN_or_N6LAN uint8 = 2
	IEDestinationInterface_CPFunction      uint8 = 3
	IEDestinationInterface_LIFunction      uint8 = 4
)

func (i *IEDestinationInterface) Set(v uint8) error {
	i.Type = IE_Destination_Interface
	i.Length = 1
	i.Value = v
	return nil
}
func (i *IEDestinationInterface) Get() (v uint8, e error) {
	return i.Value, nil
}

//Redirect Information	C
// IERedirectInformation
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
38	Redirect Information	Extendable /Subclause 8.2.20	8+a-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1-2	Type = 38 (decimal)
	3-4	Length = n
	5	Spare	           Redirect Address Type
	6-7	Redirect Server Address Length=a
	8-(8+a)	Redirect Server Address // todo 8-(7+a)
	(8+a+1) to (n+4) 	These octet(s) is/are present only if explicitly specified
Figure 8.2.20-1: Redirect Information
*/
type IERedirectInformation struct {
	IETypeLength
	AddrType      uint8  `json:",omitempty"`
	ServerAddrLen uint16 `json:",omitempty"`
	ServerAddress string `json:",omitempty"` //UTF8String format
}

func (i *IERedirectInformation) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	           Redirect Address Type
	err = encBuf.WriteByte(i.AddrType)
	if err != nil {
		return
	}

	//6-7	Redirect Server Address Length=a
	err = binary.Write(encBuf, binary.BigEndian, i.ServerAddrLen)
	if err != nil {
		return
	}
	//8-(8+a)	Redirect Server Address
	_, err = encBuf.Write([]byte(i.ServerAddress))
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IERedirectInformation) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)
	//5	Spare	           Redirect Address Type
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	i.AddrType, err = v.GetBits(1, 4)
	if err != nil {
		return err
	}
	//6-7	Redirect Server Address Length=a
	err = binary.Read(r, binary.BigEndian, &i.ServerAddrLen)
	if err != nil {
		return err
	}
	//8-(8+a)	Redirect Server Address
	tmpBytes := make([]byte, i.ServerAddrLen)
	_, err = r.Read(tmpBytes)
	if err != nil {
		return err
	}
	i.ServerAddress = string(tmpBytes)

	return nil
}

func (i *IERedirectInformation) Len() int {
	return int(i.Length)
}

func (i *IERedirectInformation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IERedirectInformation) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Table 8.2.20-1: Redirect Address Type
Redirect Address Type 	Value (Decimal)

IPv4 address	0
IPv6 address	1
URL	2
SIP URI	3
Spare, for future use.	4 to 15
*/
const (
	IERedirectInformation_IPv4Address = 0
	IERedirectInformation_IPv6Address = 1
	IERedirectInformation_URL         = 2
	IERedirectInformation_SIPURI      = 3
)

func (i *IERedirectInformation) Set(v string, AddrType uint8) error {
	i.Type = IE_Redirect_Information

	i.AddrType = AddrType
	i.ServerAddrLen = uint16(len(v))
	i.ServerAddress = v

	i.Length = 1 + 2 + i.ServerAddrLen

	return nil
}
func (i *IERedirectInformation) Get() (v string, AddrType uint8, e error) {
	return i.ServerAddress, i.AddrType, nil
}

//Outer Header Creation 	C
// IEOuterHeaderCreation
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
84	Outer Header Creation	Extendable / Subclause 8.2.56	r+1-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 84 (decimal)
	3 to 4	Length = n
	5 to 6	Outer Header Creation Description
	m to (m+3)	TEID
	p to (p+3)	IPv4 Address
	q to (q+15)	IPv6 Address
	r to (r+1)	Port Number
    t to (t+2)	C-TAG
    u to (u+2)	S-TAG
    s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.56-1: Outer Header Creation
*/
type IEOuterHeaderCreation struct {
	//gNB 侧teid
	IETypeLength
	Description uint16         `json:",omitempty"`
	TEID        types3gpp.Teid `json:",omitempty"`
	IPv4Addr    net.IP         `json:",omitempty"`
	IPv6Addr    net.IP         `json:",omitempty"`
	PortNumber  uint16         `json:",omitempty"` // 非标GTP头时使用

	CTAG uint32 `json:",omitempty"`
	STAG uint32 `json:",omitempty"`
}

func (i *IEOuterHeaderCreation) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5 to 6	Outer Header Creation Description
	err = binary.Write(encBuf, binary.BigEndian, i.Description)
	if err != nil {
		return nil, err
	}

	//m to (m+3)	TEID
	if i.Description == IEOuterHCreation_GTPU_UDP_IPv4 ||
		i.Description == IEOuterHCreation_GTPU_UDP_IPv6 ||
		i.Description == IEOuterHCreation_GTPU_UDP_IPv4+IEOuterHCreation_GTPU_UDP_IPv6 {
		tmp32 := uint32(i.TEID)
		err := binary.Write(encBuf, binary.BigEndian, tmp32)
		if err != nil {
			return nil, err
		}
	}

	//p to (p+3)	IPv4 Address
	if i.Description == IEOuterHCreation_GTPU_UDP_IPv4 ||
		i.Description == IEOuterHCreation_IPv4 ||
		i.Description == IEOuterHCreation_UDP_IPv4 {
		_, err := encBuf.Write([]byte(i.IPv4Addr.To4()))
		if err != nil {
			return nil, err
		}
	}
	//q to (q+15)	IPv6 Address
	if i.Description == IEOuterHCreation_GTPU_UDP_IPv6 ||
		i.Description == IEOuterHCreation_IPv6 ||
		i.Description == IEOuterHCreation_UDP_IPv6 {
		_, err := encBuf.Write([]byte(i.IPv6Addr.To16()))
		if err != nil {
			return nil, err
		}
	}
	//	r to (r+1)	Port Number
	if i.Description == IEOuterHCreation_UDP_IPv4 ||
		i.Description == IEOuterHCreation_UDP_IPv6 {
		err := binary.Write(encBuf, binary.BigEndian, i.PortNumber)
		if err != nil {
			return nil, err
		}
	}

	// t to (t+2)	C-TAG
	// u to (u+2)	S-TAG
	if (i.Description & IEOuterHCreation_CTAG) == IEOuterHCreation_CTAG {
		var tmpBytes3 [3]byte
		tmpBytes3[0] = byte(i.CTAG >> 16)
		tmpBytes3[1] = byte(i.CTAG >> 8)
		tmpBytes3[2] = byte(i.CTAG)
		_, err := encBuf.Write(tmpBytes3[:])
		if err != nil {
			return nil, err
		}
	}
	if (i.Description & IEOuterHCreation_STAG) == IEOuterHCreation_STAG {
		var tmpBytes3 [3]byte
		tmpBytes3[0] = byte(i.STAG >> 16)
		tmpBytes3[1] = byte(i.STAG >> 8)
		tmpBytes3[2] = byte(i.STAG)
		_, err := encBuf.Write(tmpBytes3[:])
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEOuterHeaderCreation) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	//5 to 6	Outer Header Creation Description
	err := binary.Read(r, binary.BigEndian, &i.Description)
	if err != nil {
		return err
	}
	if i.Description == IEOuterHCreation_GTPU_UDP_IPv4 ||
		i.Description == IEOuterHCreation_GTPU_UDP_IPv6 ||
		i.Description == IEOuterHCreation_GTPU_UDP_IPv4+IEOuterHCreation_GTPU_UDP_IPv6 {
		var tmp32 uint32
		err := binary.Read(r, binary.BigEndian, &tmp32)
		if err != nil {
			return err
		}
		i.TEID = types3gpp.Teid(tmp32)
	}

	if i.Description == IEOuterHCreation_GTPU_UDP_IPv4 ||
		i.Description == IEOuterHCreation_IPv4 ||
		i.Description == IEOuterHCreation_UDP_IPv4 {
		var tmpBytes4 [4]byte
		_, err := r.Read(tmpBytes4[:])
		if err != nil {
			return err
		}
		i.IPv4Addr = net.IP(tmpBytes4[:])
	}

	if i.Description == IEOuterHCreation_GTPU_UDP_IPv6 ||
		i.Description == IEOuterHCreation_IPv6 ||
		i.Description == IEOuterHCreation_UDP_IPv6 {
		var tmpBytes16 [16]byte
		_, err := r.Read(tmpBytes16[:])
		if err != nil {
			return err
		}
		i.IPv6Addr = net.IP(tmpBytes16[:])
	}
	//	r to (r+1)	Port Number
	if i.Description == IEOuterHCreation_UDP_IPv4 ||
		i.Description == IEOuterHCreation_UDP_IPv6 {
		err := binary.Read(r, binary.BigEndian, &i.PortNumber)
		if err != nil {
			return err
		}
	}

	// t to (t+2)	C-TAG
	// u to (u+2)	S-TAG
	if (i.Description & IEOuterHCreation_CTAG) == IEOuterHCreation_CTAG {
		var tmpBytes3 [3]byte
		_, err := r.Read(tmpBytes3[:])
		if err != nil {
			return err
		}
		i.CTAG = (uint32(tmpBytes3[0]) << 16) +
			(uint32(tmpBytes3[1]) << 8) +
			(uint32(tmpBytes3[2]))
	}
	if (i.Description & IEOuterHCreation_STAG) == IEOuterHCreation_STAG {
		var tmpBytes3 [3]byte
		_, err := r.Read(tmpBytes3[:])
		if err != nil {
			return err
		}
		i.STAG = (uint32(tmpBytes3[0]) << 16) +
			(uint32(tmpBytes3[1]) << 8) +
			(uint32(tmpBytes3[2]))
	}
	return nil
}

func (i *IEOuterHeaderCreation) Len() int {
	return int(i.Length)
}

func (i *IEOuterHeaderCreation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEOuterHeaderCreation) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

//Outer Header Creation Description
/*Table 8.2.56-1: Outer Header Creation Description
Octet / Bit	Outer Header to be created in the outgoing packet
5/1	GTP-U/UDP/IPv4 (NOTE 1), (NOTE 3)
5/2	GTP-U/UDP/IPv6 (NOTE 1), (NOTE 3)
5/3	UDP/IPv4 (NOTE 2, NOTE 5)
5/4	UDP/IPv6 (NOTE 2, NOTE 5)
5/5	IPv4 (NOTE 5)
5/6	IPv6 (NOTE 5)
5/7	C-TAG (see NOTE 4)
5/8	S-TAG (see NOTE 4)
*/
const (
	IEOuterHCreation_GTPU_UDP_IPv4 = 1 * 256 // 对应标准中的值256
	IEOuterHCreation_GTPU_UDP_IPv6 = 2 * 256
	IEOuterHCreation_UDP_IPv4      = 4 * 256
	IEOuterHCreation_UDP_IPv6      = 8 * 256

	IEOuterHCreation_IPv4 = 16 * 256
	IEOuterHCreation_IPv6 = 32 * 256
	IEOuterHCreation_CTAG = 64 * 256
	IEOuterHCreation_STAG = 128 * 256
)

//The TEID field shall be present if the Outer Header Creation
// Description requests the creation of a GTP-U header.
//IEOuterHCreation_GTPU_UDP_IPv4 = 1
func (i *IEOuterHeaderCreation) SetGtpuIPv4(v net.IP, teid types3gpp.Teid) error {
	i.Type = IE_Outer_Header_Creation
	i.Length = 2 + 4 + 4
	i.Description = IEOuterHCreation_GTPU_UDP_IPv4
	i.IPv4Addr = v
	i.TEID = teid

	return nil
}
func (i *IEOuterHeaderCreation) SetIPv4(v net.IP, port uint16) error {
	i.Type = IE_Outer_Header_Creation
	i.Length = 2 + 4 + 2 // Description + IPv4Addr + PortNumber
	i.Description = IEOuterHCreation_UDP_IPv4
	i.IPv4Addr = v
	i.PortNumber = port

	return nil
}
func (i *IEOuterHeaderCreation) GetGtpuIPv4() (v net.IP, teid types3gpp.Teid, e error) {
	return i.IPv4Addr, i.TEID, nil
}
func (i *IEOuterHeaderCreation) GetIPv4() (v net.IP, port uint16, e error) {
	return i.IPv4Addr, i.PortNumber, nil
}

// TODO ipv6

//Transport Level Marking 	C
// IETransportLevelMarking
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
30	Transport Level Marking	Extendable / Subclause 8.2.12	2
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 30 (decimal)
	3 to 4	Length = n
	5 to 6	ToS/Traffic Class
	7 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.12-1: Transport Level Marking
*/
//TODO the ToS/Traffic Class mask field. See subclause 5.3.15 of 3GPP TS 29.212 [8].
type IETransportLevelMarking struct {
	IETypeLength
	ToSTrafficClass uint16 `json:",omitempty"`
}

func (i *IETransportLevelMarking) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.ToSTrafficClass)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IETransportLevelMarking) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.ToSTrafficClass)
	if err != nil {
		return err
	}

	return nil
}

func (i *IETransportLevelMarking) Len() int {
	return int(i.Length)
}

func (i *IETransportLevelMarking) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IETransportLevelMarking) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IETransportLevelMarking) Set(v uint16) error {
	i.Type = IE_Transport_Level_Marking
	i.Length = 2
	i.ToSTrafficClass = v

	return nil
}
func (i *IETransportLevelMarking) Get() (v uint16, e error) {
	return i.ToSTrafficClass, nil
}

//Forwarding Policy 	C
//IEForwardingPolicy
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
41	Forwarding Policy	Extendable / Subclause 8.2.23	k-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 41 (decimal)
	3 to 4	Length = n
	5 	Forwarding Policy Identifier Length // a maximum length of 255 octets.
	j to k 	Forwarding Policy Identifier
	m to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.23-1: Forwarding Policy
*/
type IEForwardingPolicy struct {
	IETypeLength
	IdentifierLength uint8  `json:",omitempty"`
	Identifier       string `json:",omitempty"`
}

func (i *IEForwardingPolicy) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = encBuf.WriteByte(i.IdentifierLength)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write([]byte(i.Identifier))
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEForwardingPolicy) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.IdentifierLength)
	if err != nil {
		return err
	}
	tmpBytes := make([]byte, i.IdentifierLength)
	_, err = r.Read(tmpBytes)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEForwardingPolicy) Len() int {
	return int(i.Length)
}

func (i *IEForwardingPolicy) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEForwardingPolicy) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEForwardingPolicy) Set(v string) error {
	i.Type = IE_Forwarding_Policy

	i.IdentifierLength = uint8(len(v))
	i.Identifier = v

	i.Length = uint16(1 + i.IdentifierLength)

	return nil
}
func (i *IEForwardingPolicy) Get() (v string, e error) {
	return i.Identifier, nil
}

//Header Enrichment	O
// IEHeaderEnrichment
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
98	Header Enrichment	Extendable / Subclause 8.2.67	q-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 98 (decimal)
	3 to 4	Length = n
	5	Spare	      Header Type
	6	Length of Header Field Name
	7 to m	Header Field Name
	p	Length of Header Field Value
	(p+1) to q	Header Field Value
	s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.67-1: Header Enrichment
*/
type IEHeaderEnrichment struct {
	IETypeLength
	HeaderType    uint8  `json:",omitempty"`
	FieldNameLen  uint8  `json:",omitempty"`
	FieldName     string `json:",omitempty"`
	FieldValueLen uint8  `json:",omitempty"`
	FieldValue    string `json:",omitempty"`
}

func (i *IEHeaderEnrichment) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	      Header Type
	err = encBuf.WriteByte(i.HeaderType)
	if err != nil {
		return nil, err
	}
	//6	Length of Header Field Name
	err = encBuf.WriteByte(i.FieldNameLen)
	if err != nil {
		return nil, err
	}
	//7 to m	Header Field Name
	_, err = encBuf.Write([]byte(i.FieldName))
	if err != nil {
		return nil, err
	}
	//p	Length of Header Field Value
	err = encBuf.WriteByte(i.FieldValueLen)
	if err != nil {
		return nil, err
	}
	//(p+1) to q	Header Field Value
	_, err = encBuf.Write([]byte(i.FieldValue))
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEHeaderEnrichment) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	var tmp byte
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	i.HeaderType, err = v.GetBits(1, 5)
	if err != nil {
		return err
	}

	//6	Length of Header Field Name
	tmp1, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.FieldNameLen = tmp1

	//7 to m	Header Field Name

	tmpBytes := make([]byte, i.FieldNameLen)
	_, err = r.Read(tmpBytes)
	if err != nil {
		return err
	}
	i.FieldName = string(tmpBytes)
	//p	Length of Header Field Value
	tmp2, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.FieldValueLen = tmp2

	//(p+1) to q	Header Field Value
	tmpBytes2 := make([]byte, i.FieldValueLen)
	_, err = r.Read(tmpBytes2)
	if err != nil {
		return err
	}
	i.FieldValue = string(tmpBytes2)

	return nil
}

func (i *IEHeaderEnrichment) Len() int {
	return int(i.Length)
}

func (i *IEHeaderEnrichment) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEHeaderEnrichment) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Header Type indicates the type of the Header. It shall be encoded as defined in Table 8.2.67-1.
Table 8.2.67-1: Header Type
Header Type 	Value (Decimal)
HTTP	0
Spare, for future use.	1 to 31
*/
const IEHeaderEnrichment_HeaderType_HTTP = 0

func (i *IEHeaderEnrichment) Set(name string, value string) error {
	i.Type = IE_Header_Enrichment
	i.HeaderType = IEHeaderEnrichment_HeaderType_HTTP
	i.FieldNameLen = uint8(len(name))
	i.FieldName = name
	i.FieldValueLen = uint8(len(value))
	i.FieldValue = value

	return nil
}
func (i *IEHeaderEnrichment) Get() (name string, value string, e error) {
	return i.FieldName, i.FieldValue, nil
}

//Proxying	C
// IEProxying
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
137	Proxying	Extendable / Subclause 8.2.97	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 137 (decimal)
	3 to 4	Length = n
	5	Spare	                   INS	ARP
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.97-1: Proxying
*/
type IEProxying struct {
	IETypeLength
	//Flag uint8 `json:",omitempty"`
	INS bool
	ARP bool
}

func (i *IEProxying) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	                   INS	ARP
	tmpByte := utils.BoolToUint8(i.ARP) +
		(utils.BoolToUint8(i.INS) << 1)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEProxying) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)

	var tmp byte
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.INS, err = utils.GetBoolInByte(tmp, 2)
	if err != nil {
		return err
	}
	i.ARP, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEProxying) Len() int {
	return int(i.Length)
}

func (i *IEProxying) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEProxying) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*-	Bit 1 – ARP: If this bit is set to "1", then ARP proxying is performed in UPF.
-	Bit 2 – INS: If this bit is set to "1", then IPv6 Neighbour Solicitation
proxying is performed in UPF.
-	Bit 3 to 8 – spare and reserved for future use.
*/
// Flag range
const (
	IEProxying_ARP uint8 = 1
	IEProxying_INS uint8 = 2
)

/*func (i *IEProxying) Set(v uint8) error {
	i.Type = IE_Proxying
	i.Length = 1
	if v != 1 && v != 2 {
		return errors.New("Unsupported values,Valid values are 1 or 2")
	}
	i.Flag = v

	return nil
}
func (i *IEProxying) Get() (v uint8, e error) {

	return i.Flag, nil
}*/

// IEForwardingParameters
//End--------------------------------------------------------------------------

//Duplicating Parameters 	C
// IEDuplicatingParameters
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
5	Duplicating Parameters	Extendable / Table 7.5.2.3-3	Not Applicable
*/
// Table 7.5.2.3-3: Duplicating Parameters IE in FAR
/*Octet 1 and 2		Duplicating Parameters IE Type = 5 (decimal)
Octets 3 and 4		Length = n
Information elements	P

Destination Interface	M
Outer Header Creation 	C
Transport Level marking	C  IE Type :Transport Level Marking
Forwarding Policy 	C

*/

type IEDuplicatingParameters struct {
	IETypeLength
	DestinationInterface  IEDestinationInterface
	OuterHeaderCreation   *IEOuterHeaderCreation   `json:",omitempty"`
	TransportLevelmarking *IETransportLevelMarking `json:",omitempty"` // 协议上当前是小写的。
	ForwardingPolicy      *IEForwardingPolicy      `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEDuplicatingParameters) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Destination Interface	M
	vEnc, err := i.DestinationInterface.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Destination_Interface), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Outer Header Creation 	C
		case IE_Outer_Header_Creation:
			//	encode v
			vEnc, err = i.OuterHeaderCreation.Encode()
			if err != nil {
				return nil, err
			}
			//Transport Level marking	C  IE Type :Transport Level Marking
		case IE_Transport_Level_Marking:
			//	encode v
			vEnc, err = i.TransportLevelmarking.Encode()
			if err != nil {
				return nil, err
			}
			//Forwarding Policy 	C
		case IE_Forwarding_Policy:
			//	encode v
			vEnc, err = i.ForwardingPolicy.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")
		}
		// encode TL
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}
	return encBuf.Bytes(), nil
}

func (i *IEDuplicatingParameters) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEDuplicatingParameters) Len() int {
	return int(i.Length)
}

func (i *IEDuplicatingParameters) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEDuplicatingParameters) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Destination Interface	M
	case *IEDestinationInterface:
		i.DestinationInterface = *ie
		//Outer Header Creation 	C
	case *IEOuterHeaderCreation:
		i.OuterHeaderCreation = ie
		//Transport Level marking	C  IE Type :Transport Level Marking
	case *IETransportLevelMarking:
		i.TransportLevelmarking = ie
		//Forwarding Policy 	C
	case *IEForwardingPolicy:
		i.ForwardingPolicy = ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEDuplicatingParameters) Set(v uint8) error {
	i.Type = IE_Duplicating_Parameters

	return nil
}

//BAR ID	O
// IEBARID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
88	BAR ID	Extendable / Subclause 8.2.57	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 88 (decimal)
	3 to 4	Length = n
	5	BAR ID value
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.57-1: BAR ID
*/
type IEBARID struct {
	IETypeLength
	Value uint8 `json:",omitempty"`
}

func (i *IEBARID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = encBuf.WriteByte(i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEBARID) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.Value = tmp

	return nil
}

func (i *IEBARID) Len() int {
	return int(i.Length)
}

func (i *IEBARID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEBARID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEBARID) Set(v uint8) error {
	i.Type = IE_BAR_ID
	i.Length = 1
	i.Value = v

	return nil
}
func (i *IEBARID) Get() (v uint8, e error) {
	return i.Value, nil
}

// IECreateFAR
//End--------------------------------------------------------------------------

//Create URR	C  //Usage Reporting Rule
// IECreateURR
/*Table 7.5.2.4-1: Create URR IE within PFCP Session Establishment Request
 */
/*Octet 1 and 2		Create URR IE Type = 6 (decimal)
Octets 3 and 4		Length = n
Information elements	P

URR ID	M
Measurement Method	M
Reporting Triggers	M
Measurement Period 	C
Volume Threshold	C
Volume Quota	C
Event Threshold	C
Event Quota	C
Time Threshold	C
Time Quota	C
Quota Holding Time	C
Dropped DL Traffic Threshold	C
Monitoring Time	O
Subsequent Volume Threshold	O
Subsequent Time Threshold	O
Subsequent Volume Quota	O

Subsequent Time Quota	O
Subsequent Event Threshold	O
Subsequent Event Quota	O

Inactivity Detection Time	C
Linked URR ID	C
Measurement Information	C
Time Quota Mechanism	C
Aggregated URRs	C
FAR ID for Quota Action	C
Ethernet Inactivity Timer	C
Additional Monitoring Time	O
*/
type IECreateURR struct {
	IETypeLength
	URRID                     IEURRID
	MeasurementMethod         IEMeasurementMethod
	ReportingTriggers         IEReportingTriggers
	MeasurementPeriod         *IEMeasurementPeriod         `json:",omitempty"`
	VolumeThreshold           *IEVolumeThreshold           `json:",omitempty"`
	VolumeQuota               *IEVolumeQuota               `json:",omitempty"`
	EventThreshold            *IEEventThreshold            `json:",omitempty"`
	EventQuota                *IEEventQuota                `json:",omitempty"`
	TimeThreshold             *IETimeThreshold             `json:",omitempty"`
	TimeQuota                 *IETimeQuota                 `json:",omitempty"`
	QuotaHoldingTime          *IEQuotaHoldingTime          `json:",omitempty"`
	DroppedDLTrafficThreshold *IEDroppedDLTrafficThreshold `json:",omitempty"`
	MonitoringTime            *IEMonitoringTime            `json:",omitempty"`
	//EventInformation          *IEEventInformation          `json:",omitempty"`
	SubsequentVolumeThreshold *IESubsequentVolumeThreshold `json:",omitempty"`
	SubsequentTimeThreshold   *IESubsequentTimeThreshold   `json:",omitempty"`
	SubsequentVolumeQuota     *IESubsequentVolumeQuota     `json:",omitempty"`
	SubsequentTimeQuota       *IESubsequentTimeQuota       `json:",omitempty"`

	SubsequentEventThreshold *IESubsequentEventThreshold `json:",omitempty"`
	SubsequentEventQuota     *IESubsequentEventQuota     `json:",omitempty"`

	InactivityDetectionTime *IEInactivityDetectionTime `json:",omitempty"`
	LinkedURRID             *IELinkedURRID             `json:",omitempty"`
	MeasurementInformation  *IEMeasurementInformation  `json:",omitempty"`
	//TimeQuotaMechanism        IETimeQuotaMechanism  // no n4
	//AggregatedURRs            IEAggregatedURRs
	FARIDforQuotaAction      *IEFARIDforQuotaAction      `json:",omitempty"`
	EthernetInactivityTimer  *IEEthernetInactivityTimer  `json:",omitempty"`
	AdditionalMonitoringTime *IEAdditionalMonitoringTime `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IECreateURR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//URR ID	M
	vEnc, err := i.URRID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_URR_ID), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	//Measurement Method	M
	vEnc, err = i.MeasurementMethod.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_Measurement_Method), Length: uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	//Reporting Triggers	M
	vEnc, err = i.ReportingTriggers.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_Reporting_Triggers), Length: uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Measurement Period 	C
		case IE_Measurement_Period:
			//	encode v
			vEnc, err = i.MeasurementPeriod.Encode()
			if err != nil {
				return nil, err
			}
			//Volume Threshold	C
		case IE_Volume_Threshold:
			//	encode v
			vEnc, err = i.VolumeThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Volume Quota	C
		case IE_Volume_Quota:
			//	encode v
			vEnc, err = i.VolumeQuota.Encode()
			if err != nil {
				return nil, err
			}
			//Event Threshold	C
		case IE_Event_Threshold:
			//	encode v
			vEnc, err = i.EventThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Event Quota	C
		case IE_Event_Quota:
			//	encode v
			vEnc, err = i.EventQuota.Encode()
			if err != nil {
				return nil, err
			}
			//Time Threshold	C
		case IE_Time_Threshold:
			//	encode v
			vEnc, err = i.TimeThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Time Quota	C
		case IE_Time_Quota:
			//	encode v
			vEnc, err = i.TimeQuota.Encode()
			if err != nil {
				return nil, err
			}
			//Quota Holding Time	C
		case IE_Quota_Holding_Time:
			//	encode v
			vEnc, err = i.QuotaHoldingTime.Encode()
			if err != nil {
				return nil, err
			}
			//Dropped DL Traffic Threshold	C
		case IE_Dropped_DL_Traffic_Threshold:
			//	encode v
			vEnc, err = i.DroppedDLTrafficThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Monitoring Time	O
		case IE_Monitoring_Time:
			//	encode v
			vEnc, err = i.MonitoringTime.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Volume Threshold	O
		case IE_Subsequent_Volume_Threshold:
			//	encode v
			vEnc, err = i.SubsequentVolumeThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Time Threshold	O
		case IE_Subsequent_Time_Threshold:
			//	encode v
			vEnc, err = i.SubsequentTimeThreshold.Encode()
			if err != nil {
				return nil, err
			}

			//Subsequent Volume Quota	O
		case IE_Subsequent_Volume_Quota:
			//	encode v
			vEnc, err = i.SubsequentVolumeQuota.Encode()
			if err != nil {
				return nil, err
			}
			//
			//Subsequent Time Quota	O
		case IE_Subsequent_Time_Quota:
			//	encode v
			vEnc, err = i.SubsequentTimeQuota.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Event Threshold	O
		case IE_Subsequent_Event_Threshold:
			//	encode v
			vEnc, err = i.SubsequentEventThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Event Quota	O
		case IE_Subsequent_Event_Quota:
			//	encode v
			vEnc, err = i.SubsequentEventQuota.Encode()
			if err != nil {
				return nil, err
			}
			//
			//Inactivity Detection Time	C
		case IE_Inactivity_Detection_Time:
			//	encode v
			vEnc, err = i.InactivityDetectionTime.Encode()
			if err != nil {
				return nil, err
			}
			//Linked URR ID	C
		case IE_LinkedURR_ID:
			//	encode v
			vEnc, err = i.LinkedURRID.Encode()
			if err != nil {
				return nil, err
			}
			//Measurement Information	C
		case IE_Measurement_Information:
			//	encode v
			vEnc, err = i.MeasurementInformation.Encode()
			if err != nil {
				return nil, err
			}

			//FAR ID for Quota Action	C
		case IE_FAR_ID:
			//	encode v
			vEnc, err = i.FARIDforQuotaAction.Encode()
			if err != nil {
				return nil, err
			}
			//Ethernet Inactivity Timer	C
		case IE_Ethernet_Inactivity_Timer:
			//	encode v
			vEnc, err = i.EthernetInactivityTimer.Encode()
			if err != nil {
				return nil, err
			}
			//Additional Monitoring Time	O
		case IE_Additional_Monitoring_Time:
			//	encode v
			vEnc, err = i.AdditionalMonitoringTime.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")
		}
		// encode TL
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IECreateURR) Decode(data []byte) error {
	//parse v
	i.Grouped = true
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IECreateURR) Len() int {
	return int(i.Length)
}

func (i *IECreateURR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECreateURR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//URR ID	M
	case *IEURRID:
		i.URRID = *ie
		//Measurement Method	M
	case *IEMeasurementMethod:
		i.MeasurementMethod = *ie
		//Reporting Triggers	M
	case *IEReportingTriggers:
		i.ReportingTriggers = *ie

		//Measurement Period 	C
	case *IEMeasurementPeriod:
		i.MeasurementPeriod = ie
		//Volume Threshold	C
	case *IEVolumeThreshold:
		i.VolumeThreshold = ie
		//Volume Quota	C
	case *IEVolumeQuota:
		i.VolumeQuota = ie
		//Event Threshold	C
	case *IEEventThreshold:
		i.EventThreshold = ie
		//Event Quota	C
	case *IEEventQuota:
		i.EventQuota = ie
		//Time Threshold	C
	case *IETimeThreshold:
		i.TimeThreshold = ie
		//Time Quota	C
	case *IETimeQuota:
		i.TimeQuota = ie
		//Quota Holding Time	C
	case *IEQuotaHoldingTime:
		i.QuotaHoldingTime = ie
		//Dropped DL Traffic Threshold	C
	case *IEDroppedDLTrafficThreshold:
		i.DroppedDLTrafficThreshold = ie
		//Monitoring Time	O
	case *IEMonitoringTime:
		i.MonitoringTime = ie
		//Subsequent Volume Threshold	O
	case *IESubsequentVolumeThreshold:
		i.SubsequentVolumeThreshold = ie
		//Subsequent Time Threshold	O
	case *IESubsequentTimeThreshold:
		i.SubsequentTimeThreshold = ie
		//Subsequent Volume Quota	O
	case *IESubsequentVolumeQuota:
		i.SubsequentVolumeQuota = ie
		//Subsequent Time Quota	O
	case *IESubsequentTimeQuota:
		i.SubsequentTimeQuota = ie
		//Subsequent Event Threshold	O
	case *IESubsequentEventThreshold:
		i.SubsequentEventThreshold = ie
		//Subsequent Event Quota	O
	case *IESubsequentEventQuota:
		i.SubsequentEventQuota = ie
		//Inactivity Detection Time	C
	case *IEInactivityDetectionTime:
		i.InactivityDetectionTime = ie
		//Linked URR ID	C
	case *IELinkedURRID:
		i.LinkedURRID = ie
		//Measurement Information	C
	case *IEMeasurementInformation:
		i.MeasurementInformation = ie
		//Time Quota Mechanism	C // no n4
		//Aggregated URRs	C
		//FAR ID for Quota Action	C
	case *IEFARIDforQuotaAction:
		i.FARIDforQuotaAction = ie
		//Ethernet Inactivity Timer	C
	case *IEEthernetInactivityTimer:
		i.EthernetInactivityTimer = ie
		//Additional Monitoring Time	O
	case *IEAdditionalMonitoringTime:
		i.AdditionalMonitoringTime = ie
	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IECreateURR) Set(v uint8) error {
	i.Type = IE_Create_URR

	return nil
}

//Measurement Method	M
// IEMeasurementMethod
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
62	Measurement Method	Extendable / Subclause 8.2.40	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 62 (decimal)
	3 to 4	Length = n
	5	Spare	Spare	Spare	Spare	Spare	EVENT	VOLUM	DURAT
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.40-1: Measurement Method
*/
/*Octet 5 shall be encoded as follows:
-	Bit 1 – DURAT (Duration): when set to 1, this indicates a request for
measuring the duration of the traffic.
-	Bit 2 – VOLUM (Volume): when set to 1, this indicates a request for
measuring the volume of the traffic.
-	Bit 3 – EVENT (Event): when set to 1, this indicates a request for
measuring the events.
-	Bit 4 to 8: Spare, for future use and set to 0.
At least one bit shall be set to 1. Several bits may be set to 1.
*/
const (
	IEMeasurementMethod_DURAT = 1
	IEMeasurementMethod_VOLUM = 2
	IEMeasurementMethod_EVENT = 4
)

type IEMeasurementMethod struct {
	IETypeLength
	Flag  uint8 `json:",omitempty"`
	DURAT bool
	VOLUM bool
	EVENT bool
}

func (i *IEMeasurementMethod) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	// 5	Spare	Spare	Spare	Spare	Spare	EVENT	VOLUM	DURAT
	tmpByte := utils.BoolToUint8(i.DURAT) +
		(utils.BoolToUint8(i.VOLUM) << 1) +
		(utils.BoolToUint8(i.EVENT) << 2)
	err = binary.Write(encBuf, binary.BigEndian, tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEMeasurementMethod) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	var tmp byte
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}

	i.DURAT, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}

	i.VOLUM, err = utils.GetBoolInByte(tmp, 2)
	if err != nil {
		return err
	}

	i.EVENT, err = utils.GetBoolInByte(tmp, 3)
	if err != nil {
		return err
	}

	return nil

}

func (i *IEMeasurementMethod) Len() int {
	return int(i.Length)
}

func (i *IEMeasurementMethod) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEMeasurementMethod) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEMeasurementMethod) Set(v uint8) error {
	i.Type = IE_Measurement_Method
	i.Length = 1
	i.Flag = v

	return nil
}
func (i *IEMeasurementMethod) Get() (v uint8, e error) {
	return i.Flag, nil
}

//Reporting Triggers	M
// IEReportingTriggers
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
37	Reporting Triggers	Extendable /Subclause 8.2.19	2
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 37 (decimal)
	3 to 4	Length = n
	5	LIUSA	DROTH	STOPT	START	QUHTI	TIMTH	VOLTH	PERIO
	6	Spare	Spare	EVEQU	EVETH	MACAR	ENVCL	TIMQU	VOLQU
	7 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.19-1: Reporting Triggers
*/
const (
	IEReportingTriggers_LIUSA = IEFlag_Bit_8
	IEReportingTriggers_DROTH = IEFlag_Bit_7
	IEReportingTriggers_STOPT = IEFlag_Bit_6
	IEReportingTriggers_START = IEFlag_Bit_5
	IEReportingTriggers_QUHTI = IEFlag_Bit_4
	IEReportingTriggers_TIMTH = IEFlag_Bit_3
	IEReportingTriggers_VOLTH = IEFlag_Bit_2
	IEReportingTriggers_PERIO = IEFlag_Bit_1

	IEReportingTriggers_EVEQU = IEFlag_Bit_6
	IEReportingTriggers_EVETH = IEFlag_Bit_5
	IEReportingTriggers_MACAR = IEFlag_Bit_4
	IEReportingTriggers_ENVCL = IEFlag_Bit_3
	IEReportingTriggers_TIMQU = IEFlag_Bit_2
	IEReportingTriggers_VOLQU = IEFlag_Bit_1
)

type IEReportingTriggers struct {
	IETypeLength
	Flag5 uint8 `json:",omitempty"`
	Flag6 uint8 `json:",omitempty"`
	PERIO bool
	VOLTH bool
	TIMTH bool
	QUHTI bool
	START bool
	STOPT bool
	DROTH bool
	LIUSA bool

	VOLQU bool
	TIMQU bool
	ENVCL bool
	MACAR bool
	EVETH bool
	EVEQU bool
}

func (i *IEReportingTriggers) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	LIUSA	DROTH	STOPT	START	QUHTI	TIMTH	VOLTH	PERIO
	tmpByte := utils.BoolToUint8(i.PERIO) +
		(utils.BoolToUint8(i.VOLTH) << 1) +
		(utils.BoolToUint8(i.TIMTH) << 2) +
		(utils.BoolToUint8(i.QUHTI) << 3) +
		(utils.BoolToUint8(i.START) << 4) +
		(utils.BoolToUint8(i.STOPT) << 5) +
		(utils.BoolToUint8(i.DROTH) << 6) +
		(utils.BoolToUint8(i.LIUSA) << 7)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}
	//6	Spare	Spare	EVEQU	EVETH	MACAR	ENVCL	TIMQU	VOLQU
	tmpByte = utils.BoolToUint8(i.VOLQU) +
		(utils.BoolToUint8(i.TIMQU) << 1) +
		(utils.BoolToUint8(i.ENVCL) << 2) +
		(utils.BoolToUint8(i.MACAR) << 3) +
		(utils.BoolToUint8(i.EVETH) << 4) +
		(utils.BoolToUint8(i.EVEQU) << 5)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEReportingTriggers) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	var tmp5 byte
	tmp5, err := r.ReadByte()
	if err != nil {
		return err
	}

	i.PERIO, err = utils.GetBoolInByte(tmp5, 1)
	if err != nil {
		return err
	}
	i.VOLTH, err = utils.GetBoolInByte(tmp5, 2)
	if err != nil {
		return err
	}
	i.TIMTH, err = utils.GetBoolInByte(tmp5, 3)
	if err != nil {
		return err
	}
	i.QUHTI, err = utils.GetBoolInByte(tmp5, 4)
	if err != nil {
		return err
	}
	i.START, err = utils.GetBoolInByte(tmp5, 5)
	if err != nil {
		return err
	}
	i.STOPT, err = utils.GetBoolInByte(tmp5, 6)
	if err != nil {
		return err
	}
	i.DROTH, err = utils.GetBoolInByte(tmp5, 7)
	if err != nil {
		return err
	}
	i.LIUSA, err = utils.GetBoolInByte(tmp5, 8)
	if err != nil {
		return err
	}

	// octet 6
	//6	Spare	Spare	EVEQU	EVETH	MACAR	ENVCL	TIMQU	VOLQU

	var tmp6 byte
	tmp6, err = r.ReadByte()
	if err != nil {
		return err
	}
	i.VOLQU, err = utils.GetBoolInByte(tmp6, 1)
	if err != nil {
		return err
	}
	i.TIMQU, err = utils.GetBoolInByte(tmp6, 2)
	if err != nil {
		return err
	}
	i.ENVCL, err = utils.GetBoolInByte(tmp6, 3)
	if err != nil {
		return err
	}
	i.MACAR, err = utils.GetBoolInByte(tmp6, 4)
	if err != nil {
		return err
	}
	i.EVETH, err = utils.GetBoolInByte(tmp6, 5)
	if err != nil {
		return err
	}
	i.EVEQU, err = utils.GetBoolInByte(tmp6, 6)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEReportingTriggers) Len() int {
	return int(i.Length)
}

func (i *IEReportingTriggers) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEReportingTriggers) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEReportingTriggers) Set5(v uint8) error {
	i.Type = IE_Reporting_Triggers
	i.Length = 2
	i.Flag5 = v

	return nil
}
func (i *IEReportingTriggers) Set6(v uint8) error {
	i.Type = IE_Reporting_Triggers
	i.Length = 2
	i.Flag6 = v

	return nil
}
func (i *IEReportingTriggers) Get5() (v uint8, e error) {
	return i.Flag5, nil
}
func (i *IEReportingTriggers) Get6() (v uint8, e error) {
	return i.Flag6, nil
}

//Measurement Period 	C
// IEMeasurementPeriod
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
64	Measurement Period	Extendable / Subclause 8.2.42	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 64 (decimal)
	3 to 4	Length = n
	5 to 8	Measurement Period
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.42-1: Measurement Period
*/

type IEMeasurementPeriod struct {
	IETypeLength
	Period uint32 `json:",omitempty"`
}

func (i *IEMeasurementPeriod) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Period)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEMeasurementPeriod) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.Period)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEMeasurementPeriod) Len() int {
	return int(i.Length)
}

func (i *IEMeasurementPeriod) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEMeasurementPeriod) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEMeasurementPeriod) Set(v uint32) error {
	i.Type = IE_Measurement_Period
	i.Length = 4
	i.Period = v

	return nil
}
func (i *IEMeasurementPeriod) Get() (v uint32, e error) {
	return i.Period, nil
}

//Volume Threshold	C
// IEVolumeThreshold
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
31	Volume Threshold	Extendable /Subclause 8.2.13	q+7-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 31 (decimal)
	3 to 4	Length = n
	5	Spare	               DLVOL	ULVOL	TOVOL
	m to (m+7)	Total Volume
	p to (p+7)	Uplink Volume
	q to (q+7)	Downlink Volume
	s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.13-1: Volume Threshold
*/
const (
	IEVolumeThreshold_DLVOL = IEFlag_Bit_3
	IEVolumeThreshold_ULVOL = IEFlag_Bit_2
	IEVolumeThreshold_TOVOL = IEFlag_Bit_1
)

type IEVolumeThreshold struct {
	IETypeLength
	//Flag           uint8  `json:",omitempty"`
	TOVOL          bool
	ULVOL          bool
	DLVOL          bool
	TotalVolume    uint64 `json:",omitempty"`
	UplinkVolume   uint64 `json:",omitempty"`
	DownlinkVolume uint64 `json:",omitempty"`
}

func (i *IEVolumeThreshold) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	               DLVOL	ULVOL	TOVOL
	tmpByte := utils.BoolToUint8(i.TOVOL) +
		(utils.BoolToUint8(i.ULVOL) << 1) +
		(utils.BoolToUint8(i.DLVOL) << 2)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	//m to (m+7)	Total Volume
	if i.TOVOL {
		err = binary.Write(encBuf, binary.BigEndian, i.TotalVolume)
		if err != nil {
			return
		}
	}
	//p to (p+7)	Uplink Volume
	if i.ULVOL {
		err = binary.Write(encBuf, binary.BigEndian, i.UplinkVolume)
		if err != nil {
			return
		}
	}
	//q to (q+7)	Downlink Volume
	if i.DLVOL {
		err = binary.Write(encBuf, binary.BigEndian, i.DownlinkVolume)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEVolumeThreshold) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	var tmp5 byte
	tmp5, err := r.ReadByte()
	if err != nil {
		return err
	}

	i.TOVOL, err = utils.GetBoolInByte(tmp5, 1)
	if err != nil {
		return err
	}
	i.ULVOL, err = utils.GetBoolInByte(tmp5, 2)
	if err != nil {
		return err
	}
	i.DLVOL, err = utils.GetBoolInByte(tmp5, 3)
	if err != nil {
		return err
	}

	if i.TOVOL {
		err := binary.Read(r, binary.BigEndian, &i.TotalVolume)
		if err != nil {
			return err
		}
	}

	if i.ULVOL {
		err := binary.Read(r, binary.BigEndian, &i.UplinkVolume)
		if err != nil {
			return err
		}
	}

	if i.DLVOL {
		err := binary.Read(r, binary.BigEndian, &i.DownlinkVolume)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *IEVolumeThreshold) Len() int {
	return int(i.Length)
}

func (i *IEVolumeThreshold) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEVolumeThreshold) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEVolumeThreshold) Set(v uint64, flag uint8) error {
	i.Type = IE_Volume_Threshold
	i.Length = 9
	switch flag {
	case IEVolumeThreshold_TOVOL:
		i.TotalVolume = v
	case IEVolumeThreshold_ULVOL:
		i.UplinkVolume = v
	case IEVolumeThreshold_DLVOL:
		i.DownlinkVolume = v
	default:
		return errors.New("Unsupported flag")
	}

	return nil
}

/*func (i *IEVolumeThreshold) Get() (v uint64, e error) {
	switch i.Flag {
	case IEVolumeThreshold_TOVOL:
		v = i.TotalVolume
	case IEVolumeThreshold_ULVOL:
		v = i.UplinkVolume
	case IEVolumeThreshold_DLVOL:
		v = i.DownlinkVolume
	default:
		return 0, errors.New("No valid value")
	}
	return v, nil
}*/

//Volume Quota	C
// IEVolumeQuota
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
73	Volume Quota	Extendable / Subclause 8.2.50	q+7-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 73 (decimal)
	3 to 4	Length = n
	5	Spare	DLVOL	ULVOL	TOVOL
	m to (m+7)	Total Volume
	p to (p+7)	Uplink Volume
	q to (q+7)	Downlink Volume
	S to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.50-1: Volume Quota
*/
const (
	IEVolumeQuota_DLVOL = IEFlag_Bit_3
	IEVolumeQuota_ULVOL = IEFlag_Bit_2
	IEVolumeQuota_TOVOL = IEFlag_Bit_1
)

type IEVolumeQuota struct {
	IETypeLength
	//Flag           uint8  `json:",omitempty"`
	TOVOL          bool
	ULVOL          bool
	DLVOL          bool
	TotalVolume    uint64 `json:",omitempty"`
	UplinkVolume   uint64 `json:",omitempty"`
	DownlinkVolume uint64 `json:",omitempty"`
}

func (i *IEVolumeQuota) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	DLVOL	ULVOL	TOVOL
	tmpByte := utils.BoolToUint8(i.TOVOL) +
		(utils.BoolToUint8(i.ULVOL) << 1) +
		(utils.BoolToUint8(i.DLVOL) << 2)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	//m to (m+7)	Total Volume
	if i.TOVOL {
		err = binary.Write(encBuf, binary.BigEndian, i.TotalVolume)
		if err != nil {
			return
		}
	}
	//p to (p+7)	Uplink Volume
	if i.ULVOL {
		err = binary.Write(encBuf, binary.BigEndian, i.UplinkVolume)
		if err != nil {
			return
		}
	}
	//q to (q+7)	Downlink Volume
	if i.DLVOL {
		err = binary.Write(encBuf, binary.BigEndian, i.DownlinkVolume)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEVolumeQuota) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	var tmp5 byte
	tmp5, err := r.ReadByte()
	if err != nil {
		return err
	}

	i.TOVOL, err = utils.GetBoolInByte(tmp5, 1)
	if err != nil {
		return err
	}
	i.ULVOL, err = utils.GetBoolInByte(tmp5, 2)
	if err != nil {
		return err
	}
	i.DLVOL, err = utils.GetBoolInByte(tmp5, 3)
	if err != nil {
		return err
	}

	if i.TOVOL {
		err := binary.Read(r, binary.BigEndian, &i.TotalVolume)
		if err != nil {
			return err
		}
	}

	if i.ULVOL {
		err := binary.Read(r, binary.BigEndian, &i.UplinkVolume)
		if err != nil {
			return err
		}
	}

	if i.DLVOL {
		err := binary.Read(r, binary.BigEndian, &i.DownlinkVolume)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *IEVolumeQuota) Len() int {
	return int(i.Length)
}

func (i *IEVolumeQuota) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEVolumeQuota) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*func (i *IEVolumeQuota) Set(v uint64, flag uint8) error {
	i.Type = IE_Volume_Quota
	i.Length = 9

	switch flag {
	case IEVolumeQuota_TOVOL:
		i.TotalVolume = v
	case IEVolumeQuota_ULVOL:
		i.UplinkVolume = v
	case IEVolumeQuota_DLVOL:
		i.DownlinkVolume = v
	default:
		return errors.New("Unsupported flag")
	}
	i.Flag = flag
	return nil
}*/
/*func (i *IEVolumeQuota) Get() (v uint64, e error) {
	switch i.Flag {
	case IEVolumeQuota_TOVOL:
		v = i.TotalVolume
	case IEVolumeQuota_ULVOL:
		v = i.UplinkVolume
	case IEVolumeQuota_DLVOL:
		v = i.DownlinkVolume
	default:
		return 0, errors.New("No valid value")
	}
	return v, nil
}*/

//Time Threshold	C
// IETimeThreshold
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
32	Time Threshold	Extendable /Subclause 8.2.14	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 32 (decimal)
	3 to 4	Length = n
	5 to 8	Time Threshold
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.14-1: Time Threshold
*/
type IETimeThreshold struct {
	IETypeLength
	TimeThreshold uint32 `json:",omitempty"`
}

func (i *IETimeThreshold) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.TimeThreshold)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IETimeThreshold) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.TimeThreshold)
	if err != nil {
		return err
	}
	return nil
}

func (i *IETimeThreshold) Len() int {
	return int(i.Length)
}

func (i *IETimeThreshold) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IETimeThreshold) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IETimeThreshold) Set(v uint32) error {
	i.Type = IE_Time_Threshold
	i.Length = 4
	i.TimeThreshold = v

	return nil
}
func (i *IETimeThreshold) Get() (v uint32, e error) {
	return i.TimeThreshold, nil
}

//Time Quota	C
// IETimeQuota
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
74	Time Quota	Extendable / Subclause 8.2.51	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 74 (decimal)
	3 to 4	Length = n
	5 to 8	Time Quota value
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.51-1: Time Quota
*/
type IETimeQuota struct {
	IETypeLength
	Value uint32 `json:",omitempty"`
}

func (i *IETimeQuota) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IETimeQuota) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.Value)
	if err != nil {
		return err
	}
	return nil
}

func (i *IETimeQuota) Len() int {
	return int(i.Length)
}

func (i *IETimeQuota) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IETimeQuota) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IETimeQuota) Set(v uint32) error {
	i.Type = IE_Time_Quota
	i.Length = 4
	i.Value = v

	return nil
}
func (i *IETimeQuota) Get() (v uint32, e error) {
	return i.Value, nil
}

//Quota Holding Time	C
// IEQuotaHoldingTime
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
71	Quota Holding Time	Extendable / Subclause 8.2.48	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 71 (decimal)
	3 to 4	Length = n
	5 to 8	Quota Holding Time value
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.48-1: Quota Holding Time
*/
type IEQuotaHoldingTime struct {
	IETypeLength
	Value uint32 `json:",omitempty"`
}

func (i *IEQuotaHoldingTime) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEQuotaHoldingTime) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.Value)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEQuotaHoldingTime) Len() int {
	return int(i.Length)
}

func (i *IEQuotaHoldingTime) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEQuotaHoldingTime) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEQuotaHoldingTime) Set(v uint32) error {
	i.Type = IE_Quota_Holding_Time
	i.Length = 4
	i.Value = v

	return nil
}
func (i *IEQuotaHoldingTime) Get() (v uint32, e error) {
	return i.Value, nil
}

//Dropped DL Traffic Threshold	C
// IEDroppedDLTrafficThreshold
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
72	Dropped DL Traffic Threshold	Extendable / Subclause 8.2.49	m+7-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 72 (decimal)
	3 to 4	Length = n
	5	Spare	                   DLBY	DLPA
	m to (m+7)	Downlink Packets
	o to (o+7)	Number of Bytes of Downlink Data
	s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.49-1: Dropped DL Traffic Threshold
*/
const (
	IEDroppedDLTrafficThreshold_DLBY = 2
	IEDroppedDLTrafficThreshold_DLPA = 1
)

type IEDroppedDLTrafficThreshold struct {
	IETypeLength
	//Flag                        uint8  `json:",omitempty"`
	DLPA                        bool
	DLBY                        bool
	DownlinkPackets             uint64 `json:",omitempty"`
	NumberofBytesofDownlinkData uint64 `json:",omitempty"`
}

func (i *IEDroppedDLTrafficThreshold) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := utils.BoolToUint8(i.DLPA) +
		(utils.BoolToUint8(i.DLBY) << 1)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	err = binary.Write(encBuf, binary.BigEndian, i.DownlinkPackets)
	if err != nil {
		return
	}

	err = binary.Write(encBuf, binary.BigEndian, i.NumberofBytesofDownlinkData)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEDroppedDLTrafficThreshold) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp5, err := r.ReadByte()
	i.DLPA, err = utils.GetBoolInByte(tmp5, 1)
	if err != nil {
		return err
	}
	i.DLBY, err = utils.GetBoolInByte(tmp5, 2)
	if err != nil {
		return err
	}

	if i.DLPA {
		err := binary.Read(r, binary.BigEndian, &i.DownlinkPackets)
		if err != nil {
			return err
		}
	}
	if i.DLBY {
		err := binary.Read(r, binary.BigEndian, &i.NumberofBytesofDownlinkData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *IEDroppedDLTrafficThreshold) Len() int {
	return int(i.Length)
}

func (i *IEDroppedDLTrafficThreshold) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEDroppedDLTrafficThreshold) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEDroppedDLTrafficThreshold) Set(v uint64, flag uint8) error {
	i.Type = IE_Dropped_DL_Traffic_Threshold
	i.Length = 9
	switch flag {
	case IEDroppedDLTrafficThreshold_DLPA:
		i.DownlinkPackets = v
	case IEDroppedDLTrafficThreshold_DLBY:
		i.NumberofBytesofDownlinkData = v
	default:
		return errors.New("Unsupported flag")
	}
	//i.Flag = flag
	return nil
}

/*func (i *IEDroppedDLTrafficThreshold) Get() (v uint64, e error) {
	switch i.Flag {
	case IEDroppedDLTrafficThreshold_DLPA:
		v = i.DownlinkPackets
	case IEDroppedDLTrafficThreshold_DLBY:
		v = i.NumberofBytesofDownlinkData
	default:
		return 0, errors.New("No valid value")
	}
	return v, nil
}
*/
//Monitoring Time	O
// IEMonitoringTime
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
33	Monitoring Time	Extendable /Subclause 8.2.15	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 33 (decimal)
	3 to 4	Length = n
	5 to 8	Monitoring Time
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.15-1: Monitoring Time
*/
type IEMonitoringTime struct {
	IETypeLength
	MonitoringTime uint32 `json:",omitempty"`
}

func (i *IEMonitoringTime) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.MonitoringTime)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEMonitoringTime) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.MonitoringTime)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEMonitoringTime) Len() int {
	return int(i.Length)
}

func (i *IEMonitoringTime) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEMonitoringTime) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEMonitoringTime) Set(v uint32) error {
	i.Type = IE_Monitoring_Time
	i.Length = 4
	i.MonitoringTime = v

	return nil
}
func (i *IEMonitoringTime) Get() (v uint32, e error) {
	return i.MonitoringTime, nil
}

// upgrade delete
//Event Information	O
// IEEventInformation
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
148	Event Information	Extendable / Table 7.5.2.4-4	Not Applicable
*/
/*Table 7.5.2.4-4: Event Information
Octet 1 and 2		Aggregated URRs = 148 (decimal)
Octets 3 and 4		Length = n
Information elements	P

Event ID	M
Event Threshold	M
*/
/*type IEEventInformation struct {
	IETypeLength
	EventID        IEEventID        `json:",omitempty"`
	EventThreshold IEEventThreshold `json:",omitempty"`
}

func (i *IEEventInformation) Set(v uint8) error {
	i.Type = IE_Event_Information

	return nil
}
*/

// upgrade delete
//Event ID	M
//IEEventID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
150	Event ID	Extendable / Subclause 8.2.106	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 150 (decimal)
	3 to 4	Length = n
	5 to 8	Event ID
	13 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.106-1: Event ID
*/
/*type IEEventID struct {
	IETypeLength
	EventID uint32 `json:",omitempty"`
}

func (i *IEEventID) Set(v uint32) error {
	i.Type = IE_Event_ID
	i.Length = 4
	i.EventID = v

	return nil
}
func (i *IEEventID) Get() (v uint32, e error) {
	return i.EventID, nil
}
*/
// 201903 version upgrade
//Event Threshold	M
// IEEventThreshold
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
149	Event ThresholdReporting	Extendable / Subclause 8.2.113Table 7.5.8.3-2A	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 149 (decimal)
	3 to 4	Length = n
	5 to 8	Event Threshold
	13 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.113-1: Event Threshold
*/
type IEEventThreshold struct {
	IETypeLength
	EventThreshold uint32 `json:",omitempty"`
}

func (i *IEEventThreshold) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.EventThreshold)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEEventThreshold) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.EventThreshold)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEventThreshold) Len() int {
	return int(i.Length)
}

func (i *IEEventThreshold) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEventThreshold) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEEventThreshold) Set(v uint32) error {
	i.Type = IE_Event_Threshold
	i.Length = 4
	i.EventThreshold = v

	return nil
}
func (i *IEEventThreshold) Get() (v uint32, e error) {
	return i.EventThreshold, nil
}

//Subsequent Volume Threshold	O
// IESubsequentVolumeThreshold
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
34	Subsequent Volume Threshold	Extendable /Subclause 8.2.16	q+7-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 34 (decimal)
	3 to 4	Length = n
	5	Spare	DLVOL	ULVOL	TOVOL
	m to (m+7)	Total Volume
	p to (p+7)	Uplink Volume
	q to (q+7)	Downlink Volume
	s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.16-1: Subsequent Volume Threshold
*/
const (
	IESubsequentVolumeThreshold_DLVOL = IEFlag_Bit_3
	IESubsequentVolumeThreshold_ULVOL = IEFlag_Bit_2
	IESubsequentVolumeThreshold_TOVOL = IEFlag_Bit_1
)

type IESubsequentVolumeThreshold struct {
	IETypeLength
	//Flag           uint8  `json:",omitempty"`
	TOVOL          bool
	ULVOL          bool
	DLVOL          bool
	TotalVolume    uint64 `json:",omitempty"`
	UplinkVolume   uint64 `json:",omitempty"`
	DownlinkVolume uint64 `json:",omitempty"`
}

func (i *IESubsequentVolumeThreshold) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := utils.BoolToUint8(i.TOVOL) +
		(utils.BoolToUint8(i.ULVOL) << 1) +
		(utils.BoolToUint8(i.DLVOL) << 2)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	if i.TOVOL {
		err = binary.Write(encBuf, binary.BigEndian, i.TotalVolume)
		if err != nil {
			return
		}
	}

	if i.ULVOL {
		err = binary.Write(encBuf, binary.BigEndian, i.UplinkVolume)
		if err != nil {
			return
		}
	}

	if i.DLVOL {
		err = binary.Write(encBuf, binary.BigEndian, i.DownlinkVolume)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IESubsequentVolumeThreshold) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	var tmp5 byte
	tmp5, err := r.ReadByte()
	if err != nil {
		return err
	}

	i.TOVOL, err = utils.GetBoolInByte(tmp5, 1)
	if err != nil {
		return err
	}
	i.ULVOL, err = utils.GetBoolInByte(tmp5, 2)
	if err != nil {
		return err
	}
	i.DLVOL, err = utils.GetBoolInByte(tmp5, 3)
	if err != nil {
		return err
	}

	if i.TOVOL {
		err := binary.Read(r, binary.BigEndian, &i.TotalVolume)
		if err != nil {
			return err
		}
	}

	if i.ULVOL {
		err := binary.Read(r, binary.BigEndian, &i.UplinkVolume)
		if err != nil {
			return err
		}
	}

	if i.DLVOL {
		err := binary.Read(r, binary.BigEndian, &i.DownlinkVolume)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *IESubsequentVolumeThreshold) Len() int {
	return int(i.Length)
}

func (i *IESubsequentVolumeThreshold) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IESubsequentVolumeThreshold) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IESubsequentVolumeThreshold) Set(v uint64, flag uint8) error {
	i.Type = IE_Subsequent_Volume_Threshold
	i.Length = 9

	switch flag {
	case IESubsequentVolumeThreshold_TOVOL:
		i.TotalVolume = v
	case IESubsequentVolumeThreshold_ULVOL:
		i.UplinkVolume = v
	case IESubsequentVolumeThreshold_DLVOL:
		i.DownlinkVolume = v
	default:
		return errors.New("Unsupported flag")
	}
	//i.Flag = flag
	return nil
}

/*func (i *IESubsequentVolumeThreshold) Get() (v uint64, e error) {
	switch i.Flag {
	case IESubsequentVolumeThreshold_TOVOL:
		v = i.TotalVolume
	case IESubsequentVolumeThreshold_ULVOL:
		v = i.UplinkVolume
	case IESubsequentVolumeThreshold_DLVOL:
		v = i.DownlinkVolume
	default:
		return 0, errors.New("No valid value")
	}
	return v, nil
}*/

//Subsequent Time Threshold	O
// IESubsequentTimeThreshold
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
35	Subsequent Time Threshold	Extendable /Subclause 8.2.17	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 35 (decimal)
	3 to 4	Length = n
	5 to 8	Subsequent Time Threshold
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.17-1: Subsequent Time Threshold
*/
type IESubsequentTimeThreshold struct {
	IETypeLength
	SubsequentTimeThreshold uint32 `json:",omitempty"`
}

func (i *IESubsequentTimeThreshold) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.SubsequentTimeThreshold)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IESubsequentTimeThreshold) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.SubsequentTimeThreshold)
	if err != nil {
		return err
	}
	return nil
}

func (i *IESubsequentTimeThreshold) Len() int {
	return int(i.Length)
}

func (i *IESubsequentTimeThreshold) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IESubsequentTimeThreshold) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IESubsequentTimeThreshold) Set(v uint32) error {
	i.Type = IE_Subsequent_Time_Threshold
	i.Length = 4
	i.SubsequentTimeThreshold = v

	return nil
}
func (i *IESubsequentTimeThreshold) Get() (v uint32, e error) {
	return i.SubsequentTimeThreshold, nil
}

//Subsequent Volume Quota	O
// IESubsequentVolumeQuota
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
121	Subsequent Volume Quota	Extendable / Subclause 8.2.86	q+7-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 121 (decimal)
	3 to 4	Length = n
	5	Spare	DLVOL	ULVOL	TOVOL
	m to (m+7)	Total Volume
	p to (p+7)	Uplink Volume
	q to (q+7)	Downlink Volume
	S to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.86-1: Subsequent Volume Quota
*/
//TODO
type IESubsequentVolumeQuota struct {
	IEVolumeQuota
}

/*func (i *IESubsequentVolumeQuota) Set(v uint64, flag uint8) error {
	i.IEVolumeQuota.Set(v, flag)
	i.Type = IE_Subsequent_Volume_Quota
	i.Length = 9
	return nil
}*/

//Subsequent Time Quota	O
// IESubsequentTimeQuota
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
122	Subsequent Time Quota	Extendable / Subclause 8.2.87	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 122 (decimal)
	3 to 4	Length = n
	5 to 8	Time Quota value
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.87-1: Subsequent Time Quota
*/
//TODO
type IESubsequentTimeQuota struct {
	IETimeQuota
}

func (i *IESubsequentTimeQuota) Set(v uint32) error {
	i.IETimeQuota.Set(v)
	i.Type = IE_Subsequent_Time_Quota
	return nil
}

//Inactivity Detection Time	C
// IEInactivityDetectionTime
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
36	Inactivity Detection Time	Extendable /Subclause 8.2.18	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 36 (decimal)
	3 to 4	Length = n
	5 to 8	Inactivity Detection Time
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.18-1: Inactivity Detection Time
*/
type IEInactivityDetectionTime struct {
	IETypeLength
	InactDetectionTime uint32 `json:",omitempty"`
}

func (i *IEInactivityDetectionTime) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.InactDetectionTime)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEInactivityDetectionTime) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.InactDetectionTime)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEInactivityDetectionTime) Len() int {
	return int(i.Length)
}

func (i *IEInactivityDetectionTime) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEInactivityDetectionTime) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEInactivityDetectionTime) Set(v uint32) error {
	i.Type = IE_Inactivity_Detection_Time
	i.Length = 4
	i.InactDetectionTime = v
	return nil
}
func (i *IEInactivityDetectionTime) Get() (v uint32, e error) {
	return i.InactDetectionTime, nil
}

//Linked URR ID	C
// IELinkedURRID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
82	Linked URR ID	Extendable / Subclause 8.2.55	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 82 (decimal)
	3 to 4	Length = n
	5 to 8	Linked URR ID value
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.55-1: Linked URR ID
*/
type IELinkedURRID struct {
	IETypeLength
	Value uint32 `json:",omitempty"`
}

func (i *IELinkedURRID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IELinkedURRID) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.Value)
	if err != nil {
		return err
	}
	return nil
}

func (i *IELinkedURRID) Len() int {
	return int(i.Length)
}

func (i *IELinkedURRID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IELinkedURRID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IELinkedURRID) Set(v uint32) error {
	i.Type = IE_LinkedURR_ID
	i.Length = 4
	i.Value = v

	return nil
}
func (i *IELinkedURRID) Get() (v uint32, e error) {
	return i.Value, nil
}

//Measurement Information	C
// IEMeasurementInformation
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
100	Measurement Information	Extendable / Subclause 8.2.68	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 100 (decimal)
	3 to 4	Length = n
	5	Spare	          ISTM RADI INAM MBQE
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.68-1: Measurement Information
*/
const (
	IEMeasurementInfo_ISTM = IEFlag_Bit_4
	IEMeasurementInfo_RADI = IEFlag_Bit_3
	IEMeasurementInfo_INAM = IEFlag_Bit_2
	IEMeasurementInfo_MBQE = IEFlag_Bit_1
)

type IEMeasurementInformation struct {
	IETypeLength
	//Flag uint8 `json:",omitempty"`
	MBQE bool
	INAM bool
	RADI bool
	ISTM bool
}

func (i *IEMeasurementInformation) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	          ISTM RADI INAM MBQE
	tmpByte := utils.BoolToUint8(i.MBQE) +
		(utils.BoolToUint8(i.INAM) << 1) +
		(utils.BoolToUint8(i.RADI) << 2) +
		(utils.BoolToUint8(i.ISTM) << 3)

	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEMeasurementInformation) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}

	i.MBQE, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}
	i.INAM, err = utils.GetBoolInByte(tmp, 2)
	if err != nil {
		return err
	}
	i.RADI, err = utils.GetBoolInByte(tmp, 3)
	if err != nil {
		return err
	}
	i.ISTM, err = utils.GetBoolInByte(tmp, 4)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEMeasurementInformation) Len() int {
	return int(i.Length)
}

func (i *IEMeasurementInformation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEMeasurementInformation) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*func (i *IEMeasurementInformation) Set(v uint8) error {
	i.Type = IE_Measurement_Information
	i.Length = 1
	i.Flag = v

	return nil
}
func (i *IEMeasurementInformation) Get() (v uint8, e error) {
	return i.Flag, nil
}*/

//Time Quota Mechanism	C
// IETimeQuotaMechanism
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
115	Time Quota Mechanism	Extendable / Subclause 8.2.81	5
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 115 (decimal)
	3 to 4	Length = n
	5	Spare	                    BTIT
	m to (m+3)	Base Time Interval
	w to (n+4)	These octet(s) is/are present only if explicitly specified

Figure 8.2.81-1: Time Quota Mechanism
*/
type IETimeQuotaMechanism struct {
	IETypeLength
	BTIT             uint8  `json:",omitempty"`
	BaseTimeInterval uint32 `json:",omitempty"`
}

/*Table 8.2.81-1: Base Time Interval Type
Base Time Interval Type 	Value (Decimal)
CTP	0
DTP	1
Spare, for future use.	2 to 3
*/
const (
	BTIT_CTP = 0
	CTP_DTP  = 1
)

func (i *IETimeQuotaMechanism) Set(v uint32, flag uint8) error {
	i.Type = IE_Time_Quota_Mechanism
	i.Length = 5
	if flag != 0 && flag != 1 {
		return errors.New("Unsupported flag")
	}
	i.BTIT = flag
	i.BaseTimeInterval = v
	return nil
}
func (i *IETimeQuotaMechanism) Get() (v uint32, flag uint8, e error) {
	return i.BaseTimeInterval, i.BTIT, nil
}

//Aggregated URRs	C //There is no IE in N4

//FAR ID for Quota Action	C  IE Type :FAR ID
// IEFARIDforQuotaAction
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
Table 7.5.2.4-1: Create URR IE within PFCP Session Establishment Request
*/
type IEFARIDforQuotaAction struct {
	IEFARID
}

//Ethernet Inactivity Timer	C
// IEEthernetInactivityTimer
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
146	Ethernet Inactivity Timer	Extendable / Subclause 8.2.105	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 146 (decimal)
	3 to 4	Length = n
	5 to 8	Ethernet Inactivity Timer
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.105-1: Ethernet Inactivity Timer
*/
type IEEthernetInactivityTimer struct {
	IETypeLength
	TimerValue uint32 `json:",omitempty"`
}

func (i *IEEthernetInactivityTimer) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.TimerValue)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEEthernetInactivityTimer) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.TimerValue)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEthernetInactivityTimer) Len() int {
	return int(i.Length)
}

func (i *IEEthernetInactivityTimer) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEthernetInactivityTimer) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEEthernetInactivityTimer) Set(v uint32) error {
	i.Type = IE_Ethernet_Inactivity_Timer
	i.Length = 4
	i.TimerValue = v

	return nil
}
func (i *IEEthernetInactivityTimer) Get() (v uint32, e error) {
	return i.TimerValue, nil
}

//Additional Monitoring Time	O
// IEAdditionalMonitoringTime
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
147	Additional Monitoring Time	Extendable / Table 7.5.2.4-3	Not Applicable
*/
/*Table 7.5.2.4-3: Additional Monitoring Time
Octet 1 and 2		Additional Monitoring Time = 147 (decimal)
Octets 3 and 4		Length = n
Information elements	P

Monitoring Time	M
Subsequent Volume Threshold	O
Subsequent Time Threshold	O
Subsequent Volume Quota	O
Subsequent Time Quota	O
Subsequent Event Threshold	O
Subsequent Event Quota	O
*/
type IEAdditionalMonitoringTime struct {
	IETypeLength
	MonitoringTime           IEMonitoringTime
	SubseqVolumeThreshold    *IESubsequentVolumeThreshold
	SubseqTimeThreshold      *IESubsequentTimeThreshold
	SubseqVolumeQuota        *IESubsequentVolumeQuota
	SubseqTimeQuota          *IESubsequentTimeQuota
	SubsequentEventThreshold *IESubsequentEventThreshold
	SubsequentEventQuota     *IESubsequentEventQuota

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEAdditionalMonitoringTime) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Monitoring Time	M
	//encode v
	vEnc, err := i.MonitoringTime.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Monitoring_Time), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Subsequent Volume Threshold	O
		case IE_Subsequent_Volume_Threshold:
			//encode v
			vEnc, err = i.SubseqVolumeThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Time Threshold	O
		case IE_Subsequent_Time_Threshold:
			//encode v
			vEnc, err = i.SubseqTimeThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Volume Quota	O
		case IE_Subsequent_Volume_Quota:
			//encode v
			vEnc, err = i.SubseqVolumeQuota.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Time Quota	O
		case IE_Subsequent_Time_Quota:
			vEnc, err = i.SubseqTimeQuota.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Event Threshold	O
		case IE_Subsequent_Event_Threshold:
			vEnc, err = i.SubsequentEventThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Event Quota	O
		case IE_Subsequent_Event_Quota:
			vEnc, err = i.SubsequentEventQuota.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")

		}
		// TL 编码
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEAdditionalMonitoringTime) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEAdditionalMonitoringTime) Len() int {
	return int(i.Length)
}

func (i *IEAdditionalMonitoringTime) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEAdditionalMonitoringTime) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Monitoring Time	M
	case *IEMonitoringTime:
		i.MonitoringTime = *ie
		//Subsequent Volume Threshold	O
	case *IESubsequentVolumeThreshold:
		i.SubseqVolumeThreshold = ie
		//Subsequent Time Threshold	O
	case *IESubsequentTimeThreshold:
		i.SubseqTimeThreshold = ie
		//Subsequent Volume Quota	O
	case *IESubsequentVolumeQuota:
		i.SubseqVolumeQuota = ie
		//Subsequent Time Quota	O
	case *IESubsequentTimeQuota:
		i.SubseqTimeQuota = ie
		//Subsequent Event Threshold	O
	case *IESubsequentEventThreshold:
		i.SubsequentEventThreshold = ie
		//Subsequent Event Quota	O
	case *IESubsequentEventQuota:
		i.SubsequentEventQuota = ie
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEAdditionalMonitoringTime) Set(v uint8) error {
	i.Type = IE_Additional_Monitoring_Time

	return nil
}

// IEAdditionalMonitoringTime
//End--------------------------------------------------------------------------

//Create QER	C
// IECreateQER
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
7	Create QER	Extendable / Table 7.5.2.5-1	Not Applicable
*/
/*Table 7.5.2.5-1: Create QER IE within PFCP Session Establishment Request
Octet 1 and 2		Create QER IE Type = 7 (decimal)
Octets 3 and 4		Length = n
Information elements	P

QER ID	M
QER Correlation ID	C
Gate Status	M
Maximum Bitrate	C     MBR
Guaranteed Bitrate	C GBR
QoS flow identifier	C QFI
Reflective QoS	C     RQI

Paging Policy Indicator	C
Averaging Window	O
*/

type IECreateQER struct {
	IETypeLength
	QERID                 IEQERID
	QERCorrelationID      IEQERCorrelationID      `json:",omitempty"`
	GateStatus            IEGateStatus            `json:",omitempty"`
	MaximumBitrate        IEMBR                   `json:",omitempty"`
	GuaranteedBitrate     IEGBR                   `json:",omitempty"`
	QoSflowidentifier     IEQFI                   `json:",omitempty"`
	ReflectiveQoS         IERQI                   `json:",omitempty"`
	PagingPolicyIndicator IEPagingPolicyIndicator `json:",omitempty"`
	AveragingWindow       IEAveragingWindow       `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IECreateQER) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//QER ID	M
	vEnc, err := i.QERID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_QER_ID), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	//Gate Status	M
	vEnc, err = i.GateStatus.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_Gate_Status), Length: uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//QER Correlation ID	C
		case IE_QER_Correlation_ID:
			vEnc, err = i.QERCorrelationID.Encode()
			if err != nil {
				return nil, err
			}
			//Maximum Bitrate	C     MBR
		case IE_MBR:
			vEnc, err = i.MaximumBitrate.Encode()
			if err != nil {
				return nil, err
			}
			//Guaranteed Bitrate	C GBR
		case IE_GBR:
			vEnc, err = i.GuaranteedBitrate.Encode()
			if err != nil {
				return nil, err
			}
			//QoS flow identifier	C QFIs
		case IE_QFI:
			vEnc, err = i.QoSflowidentifier.Encode()
			if err != nil {
				return nil, err
			}
			//Reflective QoS	C     RQI
		case IE_RQI:
			vEnc, err = i.ReflectiveQoS.Encode()
			if err != nil {
				return nil, err
			}
			//
			//Paging Policy Indicator	C
		case IE_Paging_Policy_Indicator:
			vEnc, err = i.PagingPolicyIndicator.Encode()
			if err != nil {
				return nil, err
			}
			//Averaging Window	O
		case IE_Averaging_Window:
			vEnc, err = i.AveragingWindow.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")
		}
		// TL 编码
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IECreateQER) Decode(data []byte) error {
	//parse v
	i.Grouped = true
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IECreateQER) Len() int {
	return int(i.Length)
}

func (i *IECreateQER) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECreateQER) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//QER ID	M
	case *IEQERID:
		i.QERID = *ie
		//QER Correlation ID	C
	case *IEQERCorrelationID:
		i.QERCorrelationID = *ie
		//Gate Status	M
	case *IEGateStatus:
		i.GateStatus = *ie
		//Maximum Bitrate	C     MBR
	case *IEMBR:
		i.MaximumBitrate = *ie
		//Guaranteed Bitrate	C GBR
	case *IEGBR:
		i.GuaranteedBitrate = *ie
		//QoS flow identifier	C QFI
	case *IEQFI:
		i.QoSflowidentifier = *ie
		//Reflective QoS	C     RQI
	case *IERQI:
		i.ReflectiveQoS = *ie
		//Paging Policy Indicator	C
	case *IEPagingPolicyIndicator:
		i.PagingPolicyIndicator = *ie
		//Averaging Window	O
	case *IEAveragingWindow:
		i.AveragingWindow = *ie
	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IECreateQER) Set(v uint8) error {
	i.Type = IE_Create_QER

	return nil
}

//QER Correlation ID	C
// IEQERCorrelationID
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
28	QER Correlation ID	Extendable / Subclause 8.2.10	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 28 (decimal)
	3 to 4	Length = n
	5 to 8	QER Correlation ID value
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.10-1: QER Correlation ID

*/
type IEQERCorrelationID struct {
	IETypeLength
	Value uint32 `json:",omitempty"`
}

func (i *IEQERCorrelationID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEQERCorrelationID) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.Value)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEQERCorrelationID) Len() int {
	return int(i.Length)
}

func (i *IEQERCorrelationID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEQERCorrelationID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEQERCorrelationID) Set(v uint32) error {
	i.Type = IE_QER_Correlation_ID
	i.Length = 4
	i.Value = v

	return nil
}
func (i *IEQERCorrelationID) Get() (v uint32, e error) {
	return i.Value, nil
}

//Gate Status	M
// IEGateStatus
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
25	Gate Status	Extendable / Subclause 8.2.7	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 25 (decimal)
	3 to 4	Length = n
	5	Spare	            UL Gate	DL Gate
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.7-1: Gate Status
*/
type IEGateStatus struct {
	IETypeLength
	ULGate uint8 `json:",omitempty"`
	DLGate uint8 `json:",omitempty"`
}

func (i *IEGateStatus) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	            UL Gate（3-4）	DL Gate（1-2）
	tmpByte := i.DLGate + (i.ULGate << 2)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEGateStatus) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	i.DLGate, err = v.GetBits(1, 2)
	if err != nil {
		return err
	}
	i.ULGate, err = v.GetBits(3, 4)
	v1 := utils.ByteOne(i.ULGate)
	i.ULGate, _ = v1.RightShift(3)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEGateStatus) Len() int {
	return int(i.Length)
}

func (i *IEGateStatus) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEGateStatus) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Table 8.2.7-1: UL Gate
UL Gate 	Value (Decimal)
OPEN	0
CLOSED	1
For future use. Shall not be sent. If received, shall be interpreted as the value "1".	2, 3

Table 8.2.7-2: DL Gate
DL Gate 	Value (Decimal)
OPEN	0
CLOSED	1
For future use. Shall not be sent. If received, shall be interpreted as the value "1".	2, 3
*/
const (
	IEGateStatus_OPEN   = 0
	IEGateStatus_CLOSED = 1
)

func (i *IEGateStatus) SetUL(v uint8) error {
	i.Type = IE_Gate_Status
	i.Length = 1
	i.ULGate = v
	return nil
}
func (i *IEGateStatus) SetDL(v uint8) error {
	i.Type = IE_Gate_Status
	i.Length = 1
	i.DLGate = v
	return nil
}

//Maximum Bitrate	C     MBR
// IEMBR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
26	MBR	Extendable / Subclause 8.2.8	10
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 26 (decimal)
	3 to 4	Length = n
	5 to 9	UL MBR
	10 to 14	DL MBR
	15 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.8-1: MBR
 as kilobits per second (1 kbps = 1000 bps) in binary value
*/
type IEMBR struct {
	IETypeLength
	ULMBR uint64 `json:",omitempty"`
	DLMBR uint64 `json:",omitempty"`
}

func (i *IEMBR) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5 to 9	UL MBR
	var tmpByte5 [5]byte
	tmpByte5[0] = byte(i.ULMBR >> 32)
	tmpByte5[1] = byte(i.ULMBR >> 24)
	tmpByte5[2] = byte(i.ULMBR >> 16)
	tmpByte5[3] = byte(i.ULMBR >> 8)
	tmpByte5[4] = byte(i.ULMBR)
	_, err = encBuf.Write(tmpByte5[:])
	if err != nil {
		return
	}

	//10 to 14	DL MBR
	tmpByte5[0] = byte(i.DLMBR >> 32)
	tmpByte5[1] = byte(i.DLMBR >> 24)
	tmpByte5[2] = byte(i.DLMBR >> 16)
	tmpByte5[3] = byte(i.DLMBR >> 8)
	tmpByte5[4] = byte(i.DLMBR)
	_, err = encBuf.Write(tmpByte5[:])
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEMBR) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	//5 to 9	UL MBR
	var tmpByte5 [5]byte
	_, err := r.Read(tmpByte5[:])
	if err != nil {
		return err
	}
	i.ULMBR = (uint64(tmpByte5[0]) << 32) + (uint64(tmpByte5[1]) << 24) +
		(uint64(tmpByte5[2]) << 16) + (uint64(tmpByte5[3]) << 8) + (uint64(tmpByte5[4]))

	//10 to 14	DL MBR
	_, err = r.Read(tmpByte5[:])
	if err != nil {
		return err
	}
	i.DLMBR = (uint64(tmpByte5[0]) << 32) + (uint64(tmpByte5[1]) << 24) +
		(uint64(tmpByte5[2]) << 16) + (uint64(tmpByte5[3]) << 8) + (uint64(tmpByte5[4]))

	return nil
}

func (i *IEMBR) Len() int {
	return int(i.Length)
}

func (i *IEMBR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEMBR) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEMBR) Set(ulv uint64, dlv uint64) error {
	i.Type = IE_MBR
	i.Length = 10
	i.ULMBR = ulv
	i.DLMBR = dlv

	return nil
}
func (i *IEMBR) Get() (ulv uint64, dlv uint64, e error) {
	return i.ULMBR, i.DLMBR, nil
}

//
//Guaranteed Bitrate	C GBR
// IEGBR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
27	GBR	Extendable / Subclause 8.2.9	10
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 27 (decimal)
	3 to 4	Length = n
	5 to 9	UL GBR
	10 to 14	DL GBR
	15 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.9-1: GBR
 as kilobits per second (1 kbps = 1000 bps) in binary value
*/
type IEGBR struct {
	IETypeLength
	ULGBR uint64 `json:",omitempty"`
	DLGBR uint64 `json:",omitempty"`
}

func (i *IEGBR) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5 to 9	UL GBR
	var tmpByte5 [5]byte
	tmpByte5[0] = byte(i.ULGBR >> 32)
	tmpByte5[1] = byte(i.ULGBR >> 24)
	tmpByte5[2] = byte(i.ULGBR >> 16)
	tmpByte5[3] = byte(i.ULGBR >> 8)
	tmpByte5[4] = byte(i.ULGBR)
	_, err = encBuf.Write(tmpByte5[:])
	if err != nil {
		return
	}

	//10 to 14	DL GBR
	tmpByte5[0] = byte(i.DLGBR >> 32)
	tmpByte5[1] = byte(i.DLGBR >> 24)
	tmpByte5[2] = byte(i.DLGBR >> 16)
	tmpByte5[3] = byte(i.DLGBR >> 8)
	tmpByte5[4] = byte(i.DLGBR)
	_, err = encBuf.Write(tmpByte5[:])
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEGBR) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	//5 to 9	UL GBR
	var tmpByte5 [5]byte
	_, err := r.Read(tmpByte5[:])
	if err != nil {
		return err
	}
	i.ULGBR = (uint64(tmpByte5[0]) << 32) + (uint64(tmpByte5[1]) << 24) +
		(uint64(tmpByte5[2]) << 16) + (uint64(tmpByte5[3]) << 8) + (uint64(tmpByte5[4]))

	//10 to 14	DL GBR
	_, err = r.Read(tmpByte5[:])
	if err != nil {
		return err
	}
	i.DLGBR = (uint64(tmpByte5[0]) << 32) + (uint64(tmpByte5[1]) << 24) +
		(uint64(tmpByte5[2]) << 16) + (uint64(tmpByte5[3]) << 8) + (uint64(tmpByte5[4]))

	return nil
}

func (i *IEGBR) Len() int {
	return int(i.Length)
}

func (i *IEGBR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEGBR) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEGBR) Set(ulv uint64, dlv uint64) error {
	i.Type = IE_GBR
	i.Length = 10
	i.ULGBR = ulv
	i.DLGBR = dlv

	return nil
}
func (i *IEGBR) Get() (ulv uint64, dlv uint64, e error) {
	return i.ULGBR, i.DLGBR, nil
}

//Reflective QoS	C     RQI
// IERQI
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
123	RQI	Extendable / Subclause 8.2.88	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 123 (decimal)
	3 to 4	Length = n
	5	Spare	                        RQI
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.88-1: RQI
*/
type IERQI struct {
	IETypeLength
	RQI bool `json:",omitempty"`
}

func (i *IERQI) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = encBuf.WriteByte(utils.BoolToUint8(i.RQI))
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IERQI) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.RQI, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}

	return nil
}

func (i *IERQI) Len() int {
	return int(i.Length)
}

func (i *IERQI) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IERQI) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IERQI) Set(v bool) error {
	i.Type = IE_RQI
	i.Length = 1
	i.RQI = v

	return nil
}
func (i *IERQI) Get() (v bool, e error) {
	return i.RQI, nil
}

// IEQERCorrelationID
//End--------------------------------------------------------------------------

//Create BAR	O
// IECreateBAR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
85	Create BAR	Extendable / Table 7.5.2.6-1	Not Applicable
*/
/*Table 7.5.2.6-1: Create BAR IE within PFCP Session Establishment Request
Octet 1 and 2		Create BAR IE Type = 85 (decimal)
Octets 3 and 4		Length = n

Information elements	P

BAR ID	M
Downlink Data Notification Delay C
Suggested Buffering Packets Count  C
*/

type IECreateBAR struct {
	IETypeLength
	BARID                   IEBARID
	DLDataNotificationDelay IEDownlinkDataNotificationDelay
	SugBuffPacketsCount     IESuggestedBufferingPacketsCount `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (this *IECreateBAR) String() string {
	s := fmt.Sprintf("BARID:%v,", this.BARID.Value)
	return s
}

func (i *IECreateBAR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//BAR ID	M
	vEnc, err := i.BARID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{uint16(IE_BAR_ID), uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Downlink Data Notification Delay C
		case IE_Downlink_Data_Notification_Delay:
			//	encode v
			vEnc, err = i.DLDataNotificationDelay.Encode()
			if err != nil {
				return nil, err
			}
			//Suggested Buffering Packets Count  C
		case IE_Suggested_Buffering_Packets_Count:
			//	encode v
			vEnc, err = i.SugBuffPacketsCount.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")
		}
		// encode TL
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IECreateBAR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IECreateBAR) Len() int {
	return int(i.Length)
}

func (i *IECreateBAR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECreateBAR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//BAR ID	M
	case *IEBARID:
		i.BARID = *ie
		//Downlink Data Notification Delay C
	case *IEDownlinkDataNotificationDelay:
		i.DLDataNotificationDelay = *ie
		//Suggested Buffering Packets Count  C
	case *IESuggestedBufferingPacketsCount:
		i.SugBuffPacketsCount = *ie
	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IECreateBAR) Set() error {
	i.Type = IE_Create_BAR
	return nil
}

//Suggested Buffering Packets Count C
// IESuggestedBufferingPacketsCount
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
140	Suggested Buffering Packets Count	Extendable / Subclause 8.2.100
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 140 (decimal)
	3 to 4	Length = n
	5 	Packet count value
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.100-1: Suggested Buffering Packets Count
*/
type IESuggestedBufferingPacketsCount struct {
	IETypeLength
	CountValue uint8 `json:",omitempty"`
}

func (i *IESuggestedBufferingPacketsCount) Encode() (data []byte, err error) {
	// encode v
	return []byte{i.CountValue}, nil
}

func (i *IESuggestedBufferingPacketsCount) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.CountValue = tmp

	return nil
}

func (i *IESuggestedBufferingPacketsCount) Len() int {
	return int(i.Length)
}

func (i *IESuggestedBufferingPacketsCount) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IESuggestedBufferingPacketsCount) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IESuggestedBufferingPacketsCount) Set(v uint8) error {
	i.Type = IE_Suggested_Buffering_Packets_Count
	i.Length = 1
	i.CountValue = v
	return nil
}
func (i *IESuggestedBufferingPacketsCount) Get() (v uint8, e error) {
	return i.CountValue, nil
}

// IECreateBAR
//End--------------------------------------------------------------------------

//Create Traffic Endpoint	C
// IECreateTrafficEndpoint
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
127	Create Traffic Endpoint	Extendable / Table 7.5.2.7	Not Applicable
*/
/*Table 7.5.2.7-1: Create Traffic Endpoint IE within PFCP Session Establishment Request
Octet 1 and 2	Create Traffic Endpoint IE Type = 127(decimal)
Octets 3 and 4	Length = n
Information elements	P

Traffic Endpoint ID	M
Local F-TEID 	    O      F-TEID
Network Instance	O
UE IP address 	O
Ethernet PDU Session Information	O
Framed-Route	O
Framed-Routing	O
Framed-IPv6-Route	O

*/
type IECreateTrafficEndpoint struct {
	IETypeLength
	TrafficEndpointID IETrafficEndpointID
	LocalFTEID        *IEFTEID                  `json:",omitempty"`
	NetworkInstance   *IENetworkInstance        `json:",omitempty"`
	UEIPaddress       *IEUEIPaddress            `json:",omitempty"`
	EthPDUSessionInfo *IEEthernetPDUSessionInfo `json:",omitempty"`
	FramedRoute       *IEFramedRoute            `json:",omitempty"`
	FramedRouting     *IEFramedRouting          `json:",omitempty"`
	FramedIPv6Route   *IEFramedIPv6Route        `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IECreateTrafficEndpoint) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Traffic Endpoint ID	M
	vEnc, err := i.TrafficEndpointID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Traffic_Endpoint_ID), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Local F-TEID 	    O      F-TEID
		case IE_F_TEID:
			//	encode v
			vEnc, err = i.LocalFTEID.Encode()
			if err != nil {
				return nil, err
			}
			//Network Instance	O
		case IE_Network_Instance:
			//	encode v
			vEnc, err = i.NetworkInstance.Encode()
			if err != nil {
				return nil, err
			}
			//UE IP address 	O
		case IE_UE_IP_Address:
			//	encode v
			vEnc, err = i.UEIPaddress.Encode()
			if err != nil {
				return nil, err
			}
			//Ethernet PDU Session Information	O
		case IE_Ethernet_PDU_Session_Information:
			//	encode v
			vEnc, err = i.EthPDUSessionInfo.Encode()
			if err != nil {
				return nil, err
			}
			//Framed-Route	O
		case IE_Framed_Route:
			//	encode v
			vEnc, err = i.FramedRoute.Encode()
			if err != nil {
				return nil, err
			}
			//Framed-Routing	O
		case IE_Framed_Routing:
			//	encode v
			vEnc, err = i.FramedRouting.Encode()
			if err != nil {
				return nil, err
			}
			//Framed-IPv6-Route	O
		case IE_Framed_IPv6_Route:
			//	encode v
			vEnc, err = i.FramedIPv6Route.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")
		}
		// encode TL
		tl = IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IECreateTrafficEndpoint) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IECreateTrafficEndpoint) Len() int {
	return int(i.Length)
}

func (i *IECreateTrafficEndpoint) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECreateTrafficEndpoint) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Traffic Endpoint ID	M
	case *IETrafficEndpointID:
		i.TrafficEndpointID = *ie
		//Local F-TEID 	    O      F-TEID
	case *IEFTEID:
		i.LocalFTEID = ie
		//Network Instance	O
	case *IENetworkInstance:
		i.NetworkInstance = ie
		//UE IP address 	O
	case *IEUEIPaddress:
		i.UEIPaddress = ie
		//Ethernet PDU Session Information	O
	case *IEEthernetPDUSessionInfo:
		i.EthPDUSessionInfo = ie
		//Framed-Route	O
	case *IEFramedRoute:
		i.FramedRoute = ie
		//Framed-Routing	O
	case *IEFramedRouting:
		i.FramedRouting = ie
		//Framed-IPv6-Route	O
	case *IEFramedIPv6Route:
		i.FramedIPv6Route = ie
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IECreateTrafficEndpoint) Set(v uint8) error {
	i.Type = IE_Create_Traffic_Endpoint
	return nil
}

// IECreateTrafficEndpoint
//End--------------------------------------------------------------------------

//PDN Type	C
// IEPDNType
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
113	PDN Type	Extendable / Subclause 8.2.79	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 113 (decimal)
	3 to 4	Length = n
	5	Spare	               PDN Type
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.79-1: PDN Type

*/
type IEPDNType struct {
	IETypeLength
	PDNType uint8 `json:",omitempty"`
}

func (i *IEPDNType) Encode() (data []byte, err error) {
	// encode v
	return []byte{i.PDNType}, nil
}

func (i *IEPDNType) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	i.PDNType, err = v.GetBits(1, 3)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEPDNType) Len() int {
	return int(i.Length)
}

func (i *IEPDNType) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPDNType) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEPDNType) Set(v uint8) error {
	i.Type = IE_PDN_Type
	i.Length = 1

	return nil
}
func (i *IEPDNType) Get() (v uint8, e error) {
	return i.PDNType, nil
}

/*Table 8.2.79-1: PDN Type
PDN Type 	Value (Decimal)
IPv4	1
IPv6	2
IPv4v6	3
Non-IP	4
Ethernet	5
For future use. Shall not be sent. 	0, 6, 7*/
const (
	PDNType_IPv4     uint8 = 1
	PDNType_IPv6     uint8 = 2
	PDNType_IPv4v6   uint8 = 3
	PDNType_NorIP    uint8 = 4
	PDNType_Ethernet uint8 = 5
)

//User Plane Inactivity Timer	O
// IEUserPlaneInactivityTimer
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
117	User Plane Inactivity Timer	Extendable /Subclause 8.2.83	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 117 (decimal)
	3 to 4	Length = n
	5 to 8	User Plane Inactivity Timer
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.83-1: User Plane Inactivity Timer
*/
type IEUserPlaneInactivityTimer struct {
	IETypeLength
	TimerValue uint32 `json:",omitempty"`
}

func (i *IEUserPlaneInactivityTimer) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.TimerValue)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEUserPlaneInactivityTimer) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.TimerValue)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUserPlaneInactivityTimer) Len() int {
	return int(i.Length)
}

func (i *IEUserPlaneInactivityTimer) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUserPlaneInactivityTimer) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEUserPlaneInactivityTimer) Set(v uint32) error {
	i.Type = IE_User_Plane_Inactivity_Timer
	i.Length = 4
	i.TimerValue = v
	return nil
}
func (i *IEUserPlaneInactivityTimer) Get() (v uint32, e error) {
	return i.TimerValue, nil
}

//User ID 	O
// IEUserID
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
141	User ID	Extendable / Subclause 8.2.101	h-1-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 141 (decimal)
	3 to 4	Length = n
	5 	Spare	           NAIF	MSISDNF	IMEIF	IMSIF
	6  	Length of IMSI
	7 to a	IMSI
	b	Length of IMEI
	(b+1) to c	IMEI
	d	Length of MSISDN
	(d+1) to e	MSISDN
	f	Length of NAI
	(f+1) to g	NAI
	h to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.101-1: User ID
*/
const (
	IEUserID_NAIF    = IEFlag_Bit_4
	IEUserID_MSISDNF = IEFlag_Bit_3
	IEUserID_IMEIF   = IEFlag_Bit_2
	IEUserID_IMSIF   = IEFlag_Bit_1
)

type IEUserID struct {
	IETypeLength
	//Flag         uint8          `json:",omitempty"`
	IMSIF        bool
	IMEIF        bool
	MSISDNF      bool
	NAIF         bool
	LengthofIMSI uint8  `json:",omitempty"`
	IMSI         string `json:",omitempty"`

	LengthofIMEI uint8  `json:",omitempty"`
	IMEI         string `json:",omitempty"` // todo

	LengthofMSISDN uint8
	MSISDN         string `json:",omitempty"` // todo

	LengthofNAI uint8
	NAI         string `json:",omitempty"` // todo
}

func (i *IEUserID) Encode() (data []byte, err error) {
	//encode v
	encBuf := bytes.NewBuffer(nil)

	//5 	Spare	           NAIF	MSISDNF	IMEIF	IMSIF
	tmpByte := utils.BoolToUint8(i.IMSIF) +
		(utils.BoolToUint8(i.IMEIF) << 1) +
		(utils.BoolToUint8(i.MSISDNF) << 2) +
		(utils.BoolToUint8(i.NAIF) << 3)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	//6 Length of IMSI
	//7 to a	IMSI
	if i.IMSIF {
		err = encBuf.WriteByte(i.LengthofIMSI)
		if err != nil {
			return nil, err
		}

		_, err = encBuf.Write([]byte(i.IMSI))
		if err != nil {
			return nil, err
		}
	}
	//b	Length of IMEI
	//(b+1) to c	IMEI
	if i.IMEIF {
		err = encBuf.WriteByte(i.LengthofIMEI)
		if err != nil {
			return nil, err
		}

		_, err = encBuf.Write([]byte(i.IMEI))
		if err != nil {
			return nil, err
		}
	}
	//d	Length of MSISDN
	//(d+1) to e	MSISDN
	if i.MSISDNF {
		err = encBuf.WriteByte(i.LengthofMSISDN)
		if err != nil {
			return nil, err
		}

		_, err = encBuf.Write([]byte(i.MSISDN))
		if err != nil {
			return nil, err
		}
	}
	//f	Length of NAI
	//(f+1) to g	NAI
	if i.NAIF {
		err = encBuf.WriteByte(i.LengthofNAI)
		if err != nil {
			return nil, err
		}

		_, err = encBuf.Write([]byte(i.NAI))
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEUserID) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.IMSIF, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}
	i.IMEIF, err = utils.GetBoolInByte(tmp, 2)
	if err != nil {
		return err
	}

	i.MSISDNF, err = utils.GetBoolInByte(tmp, 3)
	if err != nil {
		return err
	}
	i.NAIF, err = utils.GetBoolInByte(tmp, 4)
	if err != nil {
		return err
	}

	if i.IMSIF {
		tmp, err := r.ReadByte()
		if err != nil {
			return err
		}
		i.LengthofIMSI = tmp

		tmpBytes := make([]byte, i.LengthofIMSI)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.IMSI = string(tmpBytes)
	}

	if i.IMEIF {
		tmp, err := r.ReadByte()
		if err != nil {
			return err
		}
		i.LengthofIMEI = tmp

		tmpBytes := make([]byte, i.LengthofIMEI)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.IMEI = string(tmpBytes)
	}

	if i.MSISDNF {
		tmp, err := r.ReadByte()
		if err != nil {
			return err
		}
		i.LengthofMSISDN = tmp

		tmpBytes := make([]byte, i.LengthofMSISDN)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.MSISDN = string(tmpBytes)
	}

	if i.NAIF {
		tmp, err := r.ReadByte()
		if err != nil {
			return err
		}
		i.LengthofNAI = tmp

		tmpBytes := make([]byte, i.LengthofNAI)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.NAI = string(tmpBytes)
	}

	return nil
}

func (i *IEUserID) Len() int {
	return int(i.Length)
}

func (i *IEUserID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUserID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEUserID) SetIMSI(v types3gpp.Imsi) error {
	i.Type = IE_User_ID
	//i.Length = uint16(2 + i.LengthofIMSI)
	i.IMSI = v.String()
	i.IMSIF = true
	return nil
}
func (i *IEUserID) SetINAI(v string) error {
	i.Type = IE_User_ID
	//i.Length = uint16(2 + XX)
	i.NAI = v
	i.NAIF = true
	return nil
}

func (i *IEUserID) GetIMSI() (v string, e error) {
	return i.IMSI, nil
}
func (i *IEUserID) GetNAI() (v string, e error) {
	return i.NAI, nil
}

//Trace Information O
// IETraceInformation
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
152	Trace Information	Extendable / Subclause 8.2.108	q-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 152 (decimal)
	3 to 4	Length = n
	5	MCC digit 2	MCC digit 1
	6	MNC digit 3	MCC digit 3
	7	MNC digit 2	MNC digit 1
	8 to10	Trace ID
	11	Length of Triggering Events
	12 to m	Triggering Events
	m+1	Session Trace Depth
	m+2	Length of List of Interfaces
	(m+3) to p	List of Interfaces
	p+1	Length of IP Address of Trace Collection Entity
	(p+2) to q	IP Address of Trace Collection Entity
	(q+1) to (n-4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.108-1: Trace Information
*/
type IETraceInformation struct {
	IETypeLength
	MCC     [3]byte `json:",omitempty"`
	MNC     [3]byte `json:",omitempty"`
	TraceID uint32  `json:",omitempty"`

	LenTriggeringEvents uint8  `json:",omitempty"`
	TriggeringEvents    []byte `json:",omitempty"`

	SessionTraceDepth uint8 `json:",omitempty"`

	LenListofInterfaces uint8  `json:",omitempty"`
	ListofInterfaces    []byte `json:",omitempty"`

	LenIPAddress uint8  `json:",omitempty"`
	IPAddress    net.IP `json:",omitempty"`
}

func (i *IETraceInformation) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//  5	MCC digit 2	MCC digit 1
	//MCC1 1-4,MCC2 5-8
	tmpByte := (i.MCC[0]) + (i.MCC[1] << 4)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	//	6	MNC digit 3	MCC digit 3
	tmpByte = (i.MCC[2]) + (i.MNC[2] << 4)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}
	//	7	MNC digit 2	MNC digit 1
	tmpByte = (i.MNC[0]) + (i.MNC[1] << 4)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}
	//	8 to10	Trace ID
	var tmpBytes [3]byte
	tmpBytes[0] = byte(i.TraceID >> 16)
	tmpBytes[1] = byte(i.TraceID >> 8)
	tmpBytes[2] = byte(i.TraceID)
	//	11	Length of Triggering Events
	//	12 to m	Triggering Events
	tmpByte = i.LenTriggeringEvents
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(i.TriggeringEvents)
	if err != nil {
		return nil, err
	}
	//	m+1	Session Trace Depth
	tmpByte = i.SessionTraceDepth
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}
	//	m+2	Length of List of Interfaces
	//	(m+3) to p	List of Interfaces
	tmpByte = i.LenListofInterfaces
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(i.ListofInterfaces)
	if err != nil {
		return nil, err
	}
	//	p+1	Length of IP Address of Trace Collection Entity
	//	(p+2) to q	IP Address of Trace Collection Entity
	tmpByte = i.LenIPAddress
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write([]byte(i.IPAddress))
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IETraceInformation) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	//MCC[1]
	i.MCC[0], err = v.GetBits(1, 4)
	if err != nil {
		return err
	}
	//MCC[2]
	i.MCC[1], err = v.GetBits(5, 8)
	v1 := utils.ByteOne(i.MCC[1])
	i.MCC[1], _ = v1.RightShift(5)
	if err != nil {
		return err
	}
	// 6	MNC digit 3	MCC digit 3
	tmp6, err := r.ReadByte()
	if err != nil {
		return err
	}
	v = utils.ByteOne(tmp6)
	//MCC[3]
	i.MCC[2], err = v.GetBits(1, 4)
	if err != nil {
		return err
	}

	//MNC[3]
	i.MNC[2], err = v.GetBits(5, 8)
	v2 := utils.ByteOne(i.MNC[2])
	i.MNC[2], _ = v2.RightShift(5)
	if err != nil {
		return err
	}

	//7	MNC digit 2	MNC digit 1
	tmp7, err := r.ReadByte()
	if err != nil {
		return err
	}
	v = utils.ByteOne(tmp7)
	//MNC[2]
	i.MNC[1], err = v.GetBits(5, 8)
	v3 := utils.ByteOne(i.MNC[1])
	i.MNC[1], _ = v3.RightShift(5)
	if err != nil {
		return err
	}
	//MNC[1]
	i.MNC[0], err = v.GetBits(1, 4)
	if err != nil {
		return err
	}

	//8 to10	Trace ID
	var tmpBytes8 [3]byte
	_, err = r.Read(tmpBytes8[:])
	if err != nil {
		return err
	}
	i.TraceID = (uint32(tmpBytes8[0]) << 16) +
		(uint32(tmpBytes8[1]) << 8) +
		(uint32(tmpBytes8[2]))

	//11	Length of Triggering Events
	tmp11, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.LenTriggeringEvents = tmp11

	//12 to m	Triggering Events
	tmpBytes12 := make([]byte, i.LenTriggeringEvents)
	_, err = r.Read(tmpBytes12)
	if err != nil {
		return err
	}
	i.TriggeringEvents = tmpBytes12

	//m+1	Session Trace Depth
	tmp13, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.SessionTraceDepth = tmp13

	//m+2	Length of List of Interfaces
	tmp14, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.LenListofInterfaces = tmp14

	//(m+3) to p	List of Interfaces
	tmpBytes15 := make([]byte, i.LenListofInterfaces)
	_, err = r.Read(tmpBytes15)
	if err != nil {
		return err
	}
	i.ListofInterfaces = tmpBytes15

	//gth of IP Address of Trace Collection Entity
	tmp16, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.LenIPAddress = tmp16
	//(p+2) to q	IP Address of Trace Collection Entity
	tmpBytes17 := make([]byte, i.LenIPAddress)
	_, err = r.Read(tmpBytes17)
	if err != nil {
		return err
	}
	i.IPAddress = tmpBytes17

	return nil
}

func (i *IETraceInformation) Len() int {
	return int(i.Length)
}

func (i *IETraceInformation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IETraceInformation) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

//todo
func (i *IETraceInformation) SetTraceID(v uint32) error {
	i.Type = IE_Trace_Information
	i.TraceID = v
	return nil
}
func (i *IETraceInformation) GetTraceID() (v uint32, e error) {
	return i.TraceID, nil
}

// IEsSessionEstablishmentRequest
//End--------------------------------------------------------------------------

// 201903 version upgrade

//8.2.112	Event Quota
// IEEventQuota
/*The Event Quota IE type shall be encoded as shown in Figure 8.2.112-1.
It contains the event quota to be monitored by the UP function.
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 148 (decimal)
	3 to 4	Length = n
	5 to 8	Subsequent Event Quota
	13 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.112-1: Event Quota
*/
type IEEventQuota struct {
	IETypeLength
	SubsequentEventQuota uint32
}

func (i *IEEventQuota) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.SubsequentEventQuota)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEEventQuota) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.SubsequentEventQuota)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEventQuota) Len() int {
	return int(i.Length)
}

func (i *IEEventQuota) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEventQuota) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEEventQuota) Set(v uint32) error {
	i.Type = IE_Event_Quota
	i.Length = 4
	i.SubsequentEventQuota = v
	return nil
}

func (i *IEEventQuota) Get() (v uint32, e error) {
	return i.SubsequentEventQuota, nil
}

// IESubsequentEventQuota
//150	Subsequent Event QuotaID	Extendable / Subclause 8.2.106	4
/*
		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 150 (decimal)
	3 to 4	Length = n
	5 to 8	Subsequent Event Quota
	13 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.106-1: Subsequent Event Quota
*/

type IESubsequentEventQuota struct {
	IETypeLength
	SubsequentEventQuota uint32
}

func (i *IESubsequentEventQuota) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.SubsequentEventQuota)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IESubsequentEventQuota) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.SubsequentEventQuota)
	if err != nil {
		return err
	}
	return nil
}

func (i *IESubsequentEventQuota) Len() int {
	return int(i.Length)
}

func (i *IESubsequentEventQuota) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IESubsequentEventQuota) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IESubsequentEventQuota) Set(v uint32) error {
	i.Type = IE_Subsequent_Event_Quota
	i.Length = 4
	i.SubsequentEventQuota = v
	return nil
}

func (i *IESubsequentEventQuota) Get() (v uint32, e error) {
	return i.SubsequentEventQuota, nil
}

// IESubsequentEventThreshold
//151	Subsequent Event Threshold	Extendable / Subclause 8.2.107	4
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 151 (decimal)
	3 to 4	Length = n
	5 to 8	Subsequent Event Threshold
	13 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.107-1: Subsequent Event Threshold
*/
type IESubsequentEventThreshold struct {
	IETypeLength
	SubsequentEventThreshold uint32
}

func (i *IESubsequentEventThreshold) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.SubsequentEventThreshold)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IESubsequentEventThreshold) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.SubsequentEventThreshold)
	if err != nil {
		return err
	}
	return nil
}

func (i *IESubsequentEventThreshold) Len() int {
	return int(i.Length)
}

func (i *IESubsequentEventThreshold) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IESubsequentEventThreshold) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IESubsequentEventThreshold) Set(v uint32) error {
	i.Type = IE_Subsequent_Event_Threshold
	i.Length = 4
	i.SubsequentEventThreshold = v
	return nil
}

func (i *IESubsequentEventThreshold) Get() (v uint32, e error) {
	return i.SubsequentEventThreshold, nil
}

// IEEventTimeStamp
//156	Event Time Stamp 	Extendable / Subclause 8.2.114	4
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 156 (decimal)
	3 to 4	Length = n
	5 to 8	Event Time Stamp
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.114-1: Event Time Stamp
*/
type IEEventTimeStamp struct {
	IETypeLength
	EventTimeStamp uint32
}

func (i *IEEventTimeStamp) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.EventTimeStamp)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEEventTimeStamp) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.EventTimeStamp)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEventTimeStamp) Len() int {
	return int(i.Length)
}

func (i *IEEventTimeStamp) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEventTimeStamp) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEEventTimeStamp) Set(v uint32) error {
	i.Type = IE_Event_Time_Stamp
	i.Length = 4
	i.EventTimeStamp = v
	return nil
}

func (i *IEEventTimeStamp) Get() (uint32, error) {
	return i.EventTimeStamp, nil
}

// IEAveragingWindow
//157	Averaging Window	Extendable /Subclause 8.2.115	4
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 157 (decimal)
	3 to 4	Length = n
	5 to 8	Averaging Window
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.115-1: Averaging Window
*/
type IEAveragingWindow struct {
	IETypeLength
	AveragingWindow uint32
}

func (i *IEAveragingWindow) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.AveragingWindow)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEAveragingWindow) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.AveragingWindow)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEAveragingWindow) Len() int {
	return int(i.Length)
}

func (i *IEAveragingWindow) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEAveragingWindow) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEAveragingWindow) Set(v uint32) error {
	i.Type = IE_Averaging_Window
	i.Length = 4
	i.AveragingWindow = v
	return nil
}

func (i *IEAveragingWindow) Get() (uint32, error) {
	return i.AveragingWindow, nil
}

// IEPagingPolicyIndicator
//158	Paging Policy Indicator	Extendable / Subclause 8.2.116	1
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 158 (decimal)
	3 to 4	Length = n
	5	Spare	               PPI value
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.116-1: Paging Policy Indicator (PPI)
The PPI value shall be encoded as a value between 0 and 7, as specified in clause 5.5.3.7 of 3GPP TS 38.415 [34].
*/
type IEPagingPolicyIndicator struct {
	IETypeLength
	PPI uint8
}

func (i *IEPagingPolicyIndicator) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = encBuf.WriteByte(i.PPI)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEPagingPolicyIndicator) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	i.PPI, err = v.GetBits(1, 3)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEPagingPolicyIndicator) Len() int {
	return int(i.Length)
}

func (i *IEPagingPolicyIndicator) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPagingPolicyIndicator) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEPagingPolicyIndicator) Set(v uint8) error {
	i.Type = IE_Paging_Policy_Indicator
	i.Length = 1
	i.PPI = v
	return nil
}

func (i *IEPagingPolicyIndicator) Get() (uint8, error) {
	return i.PPI, nil
}

// 201903 version upgrade
//End--------------------------------------------------------------------------
//8874
