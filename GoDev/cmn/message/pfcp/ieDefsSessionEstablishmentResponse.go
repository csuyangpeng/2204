package pfcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/message/pfcp/utils"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)
// N4 消息

// IE 名称 --来源于消息
// IE type --来源于8.1IE列表
// IE 格式 --来源于8.2格式定义

//Created PDR 	C
// IECreatedPDR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
8	Created PDR	Extendable / Table 7.5.3.2-1	Not Applicable
*/
/*Table 7.5.3.2-1: Created PDR IE within PFCP Session Establishment Response
Octet 1 and 2		Created PDR IE Type = 8 (decimal)
Octets 3 and 4		Length = n
Information elements	P

PDR ID	M
Local F-TEID 	C     F-TEID
*/
type IECreatedPDR struct {
	IETypeLength
	PDRID      IEPDRID
	LocalFTEID IEFTEID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet //可选
}

func (i *IECreatedPDR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//PDR ID	M
	vEnc, err := i.PDRID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{uint16(IE_Created_PDR), uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	// optional ie
	//Local F-TEID 	C     F-TEID
	if i.IeFlags.Test(IE_F_TEID) {
		//	encode v
		vEnc, err = i.LocalFTEID.Encode()
		if err != nil {
			return nil, err
		}
		// encode tl
		tl = IETypeLength{Type: uint16(IE_F_TEID), Length: uint16(len(vEnc))}
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

func (i *IECreatedPDR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IECreatedPDR) Len() int {
	return int(i.Length)
}

func (i *IECreatedPDR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECreatedPDR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//PDR ID	M
	case *IEPDRID:
		i.PDRID = *ie
		//Local F-TEID 	C     F-TEID
	case *IEFTEID:
		i.LocalFTEID = *ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IECreatedPDR) Set(v uint8) error {
	i.Type = IE_Created_PDR

	return nil
}

//Failed Rule ID	C
// IEFailedRuleID
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
114	Failed Rule ID	Extendable / Subclause 8.2.80	p-4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 114 (decimal)
	3 to 4	Length = n
	5	Spare	Rule ID Type
	6 to p	Rule ID value
	(p+1) to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.80-1: Failed Rule ID
*/
type IEFailedRuleID struct {
	IETypeLength
	RuleIDType uint8       `json:",omitempty"`
	Value      interface{} `json:",omitempty"`
}

func (i *IEFailedRuleID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	err = binary.Write(encBuf, binary.BigEndian, i.RuleIDType)
	if err != nil {
		return
	}
	switch i.RuleIDType {
	case RuleID_PDR:
		pdrId, ok := i.Value.(IEPDRID)
		if !ok {
			return nil, errors.New("Invalid type")
		}
		// 仅编码进对应的id值
		err = binary.Write(encBuf, binary.BigEndian, pdrId.RuleID)
		if err != nil {
			return
		}
	case RuleID_FAR:
		id, ok := i.Value.(IEFARID)
		if !ok {
			return nil, errors.New("Invalid type")
		}
		err = binary.Write(encBuf, binary.BigEndian, id.Value)
		if err != nil {
			return
		}
	case RuleID_QER:
		id, ok := i.Value.(IEQERID)
		if !ok {
			return nil, errors.New("Invalid type")
		}
		err = binary.Write(encBuf, binary.BigEndian, id.Value)
		if err != nil {
			return
		}
	case RuleID_URR:
		id, ok := i.Value.(IEURRID)
		if !ok {
			return nil, errors.New("Invalid type")
		}
		err = binary.Write(encBuf, binary.BigEndian, id.URRIDValue)
		if err != nil {
			return
		}
	case RuleID_BAR:
		id, ok := i.Value.(IEBARID)
		if !ok {
			return nil, errors.New("Invalid type")
		}
		err = binary.Write(encBuf, binary.BigEndian, id.Value)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEFailedRuleID) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)
	tmp5, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp5)
	// 5	Spare(6-8)	Rule ID Type(1-5)
	i.RuleIDType, err = v.GetBits(1, 5)
	if err != nil {
		return err
	}

	//6 to p	Rule ID value
	switch i.RuleIDType {
	case RuleID_PDR:
		pdrID := IEPDRID{}
		err := binary.Read(r, binary.BigEndian, &pdrID.RuleID)
		if err != nil {
			return err
		}
		i.Value = pdrID.RuleID
	case RuleID_FAR:
		farID := IEFARID{}
		err := binary.Read(r, binary.BigEndian, &farID.Value)
		if err != nil {
			return err
		}
		i.Value = farID.Value
	case RuleID_QER:
		qerID := IEQERID{}
		err := binary.Read(r, binary.BigEndian, &qerID.Value)
		if err != nil {
			return err
		}
		i.Value = qerID.Value
	case RuleID_URR:
		urrID := IEURRID{}
		err := binary.Read(r, binary.BigEndian, &urrID.URRIDValue)
		if err != nil {
			return err
		}
		i.Value = urrID.URRIDValue
	case RuleID_BAR:
		barID := IEBARID{}
		err := binary.Read(r, binary.BigEndian, &barID.Value)
		if err != nil {
			return err
		}
		i.Value = barID.Value

	default:
		return errors.New("Invalid type")
	}

	return nil
}

func (i *IEFailedRuleID) Len() int {
	return int(i.Length)
}

func (i *IEFailedRuleID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEFailedRuleID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Table 8.2.80-1: Rule ID Type
Rule ID Type	Value (Decimal)
PDR	0
FAR	1
QER	2
URR	3
BAR	4
For future use. Shall not be sent. If received, shall be interpreted as the value "1".	5 to 31
*/
const (
	RuleID_PDR = 0
	RuleID_FAR = 1
	RuleID_QER = 2
	RuleID_URR = 3
	RuleID_BAR = 4
)

func (i *IEFailedRuleID) Set(v interface{}, typeFlag uint8) error {
	i.Type = IE_Failed_Rule_ID
	i.RuleIDType = typeFlag
	i.Value = v
	return nil
}
func (i *IEFailedRuleID) Get() (v interface{}, e error) {
	return i.Value, nil
}

//Created Traffic Endpoint 	C
// IECreatedTrafficEndpoint
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
128	Created Traffic Endpoint	Extendable / Table 7.5.3.5	Not Applicable
*/
/*Table 7.5.3.5-1: Created Traffic Endpoint IE within Sx Session Establishment Response
Octet 1 and 2		Created Traffic Endpoint IE Type = 128 (decimal)
Octets 3 and 4		Length = n
Information elements	P   IE Type

Traffic Endpoint ID	    M   Traffic Endpoint ID
Local F-TEID 	        C     F-TEID
*/
type IECreatedTrafficEndpoint struct {
	IETypeLength
	TrafficEndpointID IETrafficEndpointID
	LocalFTEID        IEFTEID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IECreatedTrafficEndpoint) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Traffic Endpoint ID	    M
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

	//optional
	//Local F-TEID 	        C
	if i.IeFlags.Test(IE_F_TEID) {
		vEnc, err = i.LocalFTEID.Encode()
		if err != nil {
			return
		}
		//encode tl
		tl := IETypeLength{Type: uint16(IE_F_TEID), Length: uint16(len(vEnc))}
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

func (i *IECreatedTrafficEndpoint) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IECreatedTrafficEndpoint) Len() int {
	return int(i.Length)
}

func (i *IECreatedTrafficEndpoint) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECreatedTrafficEndpoint) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Traffic Endpoint ID	    M   Traffic Endpoint ID
	case *IETrafficEndpointID:
		i.TrafficEndpointID = *ie
		//Local F-TEID 	        C     F-TEID
	case *IEFTEID:
		i.LocalFTEID = *ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IECreatedTrafficEndpoint) Set(v uint8) error {
	i.Type = IE_Created_Traffic_Endpoint

	return nil
}

// IEsSessionEstablishmentResponse
//End--------------------------------------------------------------------------
// 317
