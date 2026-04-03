package pfcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/message/pfcp/utils"
	"time"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)
// N4 消息

// IE 名称 --来源于消息
// IE type --来源于8.1IE列表
// IE 格式 --来源于8.2格式定义

//Remove PDR    C                                 Remove PDR
// IERemovePDR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
15	Remove PDR	Extendable / Table 7.5.4.6	Not Applicable

Table 7.5.4.6-1: Remove PDR IE within PFCP Session Modification Request
Octet 1 and 2		Remove PDR IE Type = 15 (decimal)
Octets 3 and 4		Length = n
Information elements	P	IE Type

PDR ID	                M	PDR ID
*/

//todo:Several IEs within the same IE type may be present to represent a list of PDRs to remove.
type IERemovePDR struct {
	IETypeLength
	PDRID IEPDRID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IERemovePDR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//PDR ID	                M	PDR ID
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

	return encBuf.Bytes(), nil
}

func (i *IERemovePDR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IERemovePDR) Len() int {
	return int(i.Length)
}

func (i *IERemovePDR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IERemovePDR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//PDR ID	                M	PDR ID
	case *IEPDRID:
		i.PDRID = *ie
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IERemovePDR) Set(v uint16) error {
	i.Type = IE_Remove_PDR
	i.PDRID.Set(v)

	return nil
}
func (i *IERemovePDR) Get() (v uint16, e error) {
	return i.PDRID.Get()
}

//Remove FAR	C                                 Remove FAR
// IERemoveFAR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
16	Remove FAR	Extendable / Table 7.5.4.7	Not Applicable
*/
/*Table 7.5.4.7-1: Remove FAR IE within PFCP Session Modification Request
Octet 1 and 2		Remove FAR IE Type = 16 (decimal)
Octets 3 and 4		Length = n
Information elements	P          IE Type

FAR ID	M                          FAR ID
*/

type IERemoveFAR struct {
	IETypeLength
	FARID IEFARID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IERemoveFAR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//FAR ID	M                          FAR ID
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

	return encBuf.Bytes(), nil
}

func (i *IERemoveFAR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IERemoveFAR) Len() int {
	return int(i.Length)
}

func (i *IERemoveFAR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IERemoveFAR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//FAR ID	M                          FAR ID
	case *IEFARID:
		i.FARID = *ie
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IERemoveFAR) Set(v uint32) error {
	i.Type = IE_Remove_FAR
	i.FARID.Set(v)
	return nil
}
func (i *IERemoveFAR) Get() (v uint32, e error) {
	return i.FARID.Get()
}

//Remove URR	C                                 Remove URR
// IERemoveURR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
17	Remove URR	Extendable / Table 7.5.4.8	Not Applicable
*/
/*Table 7.5.4.8-1: Remove URR IE within PFCP Session Modification Request
Octet 1 and 2		Remove URR IE Type = 17 (decimal)
Octets 3 and 4		Length = n
Information elements	P    IE Type

URR ID	                M    URR ID
*/
type IERemoveURR struct {
	IETypeLength
	URRID IEURRID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IERemoveURR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	// URR ID	                M    URR ID
	// encode v
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

	return encBuf.Bytes(), nil
}

func (i *IERemoveURR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IERemoveURR) Len() int {
	return int(i.Length)
}

func (i *IERemoveURR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IERemoveURR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//URR ID	                M    URR ID
	case *IEURRID:
		i.URRID = *ie
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IERemoveURR) Set(v uint32) error {
	i.Type = IE_Remove_URR
	i.URRID.Set(v)
	return nil
}
func (i *IERemoveURR) Get() (v uint32, e error) {
	return i.URRID.Get()
}

//Remove QER	C                                 Remove QER
// IERemoveQER
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
18	Remove QER	Extendable / Table 7.5.4.9	Not Applicable
*/
/*Table 7.5.4.9-1: Remove QER IE PFCP Session Modification Request
Octet 1 and 2		Remove QER IE Type = 18 (decimal)
Octets 3 and 4		Length = n
Information elements	P     IE Type

QER ID	               M      QER ID
*/
type IERemoveQER struct {
	IETypeLength
	QERID IEQERID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IERemoveQER) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	// QER ID	               M      QER ID
	// encode v
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

	return encBuf.Bytes(), nil
}

func (i *IERemoveQER) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IERemoveQER) Len() int {
	return int(i.Length)
}

func (i *IERemoveQER) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IERemoveQER) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//QER ID	               M      QER ID
	case *IEQERID:
		i.QERID = *ie
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IERemoveQER) Set(v uint32) error {
	i.Type = IE_Remove_QER
	i.QERID.Set(v)

	return nil
}
func (i *IERemoveQER) Get() (v uint32, e error) {
	return i.QERID.Get()
}

//Remove BAR	C                                 Remove BAR
// IERemoveBAR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
87	Remove BAR	Extendable / Table 7.5.4.12-1	Not Applicable
*/
/*Table 7.5.4.12-1: Remove BAR IE within PFCP Session Modification Request
Octet 1 and 2		Remove BAR IE Type = 87 (decimal)
Octets 3 and 4		Length = n
Information elements	P

BAR ID	                M
*/
type IERemoveBAR struct {
	IETypeLength
	BARID IEBARID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IERemoveBAR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	// BAR ID	                M
	// encode v
	vEnc, err := i.BARID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_BAR_ID), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IERemoveBAR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IERemoveBAR) Len() int {
	return int(i.Length)
}

func (i *IERemoveBAR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IERemoveBAR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//BAR ID	                M
	case *IEBARID:
		i.BARID = *ie
	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IERemoveBAR) Set(v uint8) error {
	i.Type = IE_Remove_BAR
	i.BARID.Set(v)
	return nil
}
func (i *IERemoveBAR) Get() (v uint8, e error) {
	return i.BARID.Get()
}

//Remove Traffic Endpoint	C                     Remove Traffic Endpoint
// IERemoveTrafficEndpoint
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
130	Remove Traffic Endpoint	Extendable / Table 7.5.4.14	Not Applicable
*/
/*Table 7.5.4.14-1: Remove Traffic Endpoint IE within Sx Session Modification Request
Octet 1 and 2		Remove Traffic Endpoint IE Type = 130 (decimal)
Octets 3 and 4		Length = n
Information elements	P

Traffic Endpoint ID	   M

*/
type IERemoveTrafficEndpoint struct {
	IETrafficEndpointID
}

func (i *IERemoveTrafficEndpoint) Set(v uint8) error {
	i.IETrafficEndpointID.Set(v)
	i.Type = IE_Remove_Traffic_Endpoint

	return nil
}
func (i *IERemoveTrafficEndpoint) Get() (v uint8, e error) {
	return i.IETrafficEndpointID.Get()
}

//Create PDR	C                                 Create PDR
//Create FAR	C                                 Create FAR
//Create URR	C                                 Create URR
//Create QER	C                                 Create QER
//Create BAR	C                                 Create BAR
//Create Traffic Endpoint	C                     Create Traffic Endpoint

//Update PDR	C                                 Update PDR
// IEUpdatePDR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
9	Update PDR	Extendable / Table 7.5.4.2-1	Not Applicable
*/
/*Table 7.5.4.2-1: Update PDR IE within PFCP Session Modification Request
Octet 1 and 2		Update PDR IE Type = 9 (decimal)
Octets 3 and 4		Length = n
Information elements	P            IE Type

PDR ID	                M            PDR ID
Outer Header Removal 	C            Outer Header Removal
Precedence	C                        Precedence
PDI	C                                PDI
FAR ID 	C                            FAR ID
URR ID 	C                            URR ID
QER ID 	C                            QER ID
Activate Predefined Rules 	C        Activate Predefined Rules
Deactivate Predefined Rules 	C    Deactivate Predefined Rules

*/

type IEUpdatePDR struct {
	IETypeLength
	PDRID                     *IEPDRID
	OuterHeaderRemoval        *IEOuterHeaderRemoval
	Precedence                *IEPrecedence
	PDI                       *IEPDI
	FARID                     *IEFARID
	URRID                     *IEURRID
	QERID                     *IEQERID
	ActivatePredefinedRules   []*IEActivatePredefinedRules
	DeactivatePredefinedRules []*IEDeactivatePredefinedRules

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUpdatePDR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//PDR ID	                M            PDR ID
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
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Outer Header Removal 	C            Outer Header Removal
		case IE_Outer_Header_Removal:
			//	encode v
			vEnc, err = i.OuterHeaderRemoval.Encode()
			if err != nil {
				return nil, err
			}
			//Precedence	C                        Precedence
		case IE_Precedence:
			//	encode v
			vEnc, err = i.Precedence.Encode()
			if err != nil {
				return nil, err
			}
			//PDI	C                                PDI
		case IE_PDI:
			//	encode v
			vEnc, err = i.PDI.Encode()
			if err != nil {
				return nil, err
			}
			//FAR ID 	C                            FAR ID
		case IE_FAR_ID:
			//	encode v
			vEnc, err = i.FARID.Encode()
			if err != nil {
				return nil, err
			}
			//URR ID 	C                            URR ID
		case IE_URR_ID:
			//	encode v
			vEnc, err = i.URRID.Encode()
			if err != nil {
				return nil, err
			}
			//QER ID 	C                            QER ID
		case IE_QER_ID:
			//	encode v
			vEnc, err = i.QERID.Encode()
			if err != nil {
				return nil, err
			}
			//Activate Predefined Rules 	C        Activate Predefined Rules
		case IE_Activate_Predefined_Rules:
			for _, v := range i.ActivatePredefinedRules[:len(i.ActivatePredefinedRules)-1] {
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
			vEnc, err = i.ActivatePredefinedRules[len(i.ActivatePredefinedRules)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Deactivate Predefined Rules 	C    Deactivate Predefined Rules
		case IE_Deactivate_Predefined_Rules:
			for _, v := range i.DeactivatePredefinedRules[:len(i.DeactivatePredefinedRules)-1] {
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
			vEnc, err = i.DeactivatePredefinedRules[len(i.DeactivatePredefinedRules)-1].Encode()
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

func (i *IEUpdatePDR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUpdatePDR) Len() int {
	return int(i.Length)
}

func (i *IEUpdatePDR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUpdatePDR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//PDR ID	                M            PDR ID
	case *IEPDRID:
		i.PDRID = ie
		//Outer Header Removal 	C            Outer Header Removal
	case *IEOuterHeaderRemoval:
		i.OuterHeaderRemoval = ie
		//Precedence	C                        Precedence
	case *IEPrecedence:
		i.Precedence = ie
		//PDI	C                                PDI
	case *IEPDI:
		i.PDI = ie
		//FAR ID 	C                            FAR ID
	case *IEFARID:
		i.FARID = ie
		//URR ID 	C                            URR ID
	case *IEURRID:
		i.URRID = ie
		//QER ID 	C                            QER ID
	case *IEQERID:
		i.QERID = ie
		//Activate Predefined Rules 	C        Activate Predefined Rules
	case *IEActivatePredefinedRules:
		i.ActivatePredefinedRules = append(i.ActivatePredefinedRules, ie)
		//Deactivate Predefined Rules 	C    Deactivate Predefined Rules
	case *IEDeactivatePredefinedRules:
		i.DeactivatePredefinedRules = append(i.DeactivatePredefinedRules, ie)

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEUpdatePDR) Set() error {
	i.Type = IE_Update_PDR

	return nil
}

//Deactivate Predefined Rules
// IEDeactivatePredefinedRules
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
107	Deactivate Predefined Rules 	Variable Length / Subclause 8.2.73	Not Applicable
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 107 (decimal)
	3 to 4	Length = n
	5 to (n+4)	Predefined Rules Name

Figure 8.2.73-1: Deactivate Predefined Rules
*/
type IEDeactivatePredefinedRules struct {
	IETypeLength
	RulesName string
}

func (i *IEDeactivatePredefinedRules) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	_, err = encBuf.Write([]byte(i.RulesName))
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEDeactivatePredefinedRules) Decode(data []byte) error {
	//	parse v
	//	5 to (n+4)	Predefined Rules Name
	i.RulesName = string(data)

	return nil
}

func (i *IEDeactivatePredefinedRules) Len() int {
	return int(i.Length)
}

func (i *IEDeactivatePredefinedRules) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEDeactivatePredefinedRules) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEDeactivatePredefinedRules) Set(v string) error {
	i.Type = IE_Deactivate_Predefined_Rules
	i.Length = uint16(len(v))
	i.RulesName = v
	return nil
}
func (i *IEDeactivatePredefinedRules) Get() (v string, e error) {
	return i.RulesName, nil
}

// IEUpdatePDR
//End--------------------------------------------------------------------------

//Update FAR	C                                 Update FAR
// IEUpdateFAR
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
10	Update FAR	Extendable / Table 7.5.4.3-1	Not Applicable
*/
/*Table 7.5.4.3-1: Update FAR IE within PFCP Session Modification Request
Octet 1 and 2		Update FAR IE Type = 10 (decimal)
Octets 3 and 4		Length = n
Information elements	P                IE Type

FAR ID	M                                FAR ID
Apply Action	C                        Apply Action
Update Forwarding parameters	C        Update Forwarding Parameters
Update Duplicating Parameters 	C        Update Duplicating Parameters
BAR ID	C                                BAR ID

*/
type IEUpdateFAR struct {
	IETypeLength
	FARID                 IEFARID
	ApplyAction           IEApplyAction
	UpdateForwardingPara  *IEUpdateForwardingParameters
	UpdateDuplicatingPara *IEUpdateDuplicatingParameters
	BARID                 *IEBARID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUpdateFAR) Encode() (data []byte, err error) {
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

	//  optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Apply Action	C                        Apply Action
		case IE_Apply_Action:
			// encode v
			vEnc, err = i.ApplyAction.Encode()
			if err != nil {
				return nil, err
			}
			//Update Forwarding parameters	C        Update Forwarding Parameters
		case IE_Update_Forwarding_Parameters:
			// encode v
			vEnc, err = i.UpdateForwardingPara.Encode()
			if err != nil {
				return nil, err
			}
			//Update Duplicating Parameters 	C        Update Duplicating Parameters
		case IE_Update_Duplicating_Parameters:
			// encode v
			vEnc, err = i.UpdateDuplicatingPara.Encode()
			if err != nil {
				return nil, err
			}
			//BAR ID	C                                BAR ID
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

func (i *IEUpdateFAR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUpdateFAR) Len() int {
	return int(i.Length)
}

func (i *IEUpdateFAR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUpdateFAR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//FAR ID	M                                FAR ID
	case *IEFARID:
		i.FARID = *ie
		//Apply Action	C                        Apply Action
	case *IEApplyAction:
		i.ApplyAction = *ie
		//Update Forwarding parameters	C        Update Forwarding Parameters
	case *IEUpdateForwardingParameters:
		i.UpdateForwardingPara = ie
		//Update Duplicating Parameters 	C        Update Duplicating Parameters
	case *IEUpdateDuplicatingParameters:
		i.UpdateDuplicatingPara = ie
		//BAR ID	C                                BAR ID
	case *IEBARID:
		i.BARID = ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEUpdateFAR) Set() error {
	i.Type = IE_Update_FAR
	return nil
}

//Update Forwarding Parameters
// IEUpdateForwardingParameters
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
11	Update Forwarding Parameters	Extendable / Table 7.5.4.3-2	Not Applicable
*/
/*Table 7.5.4.3-2: Update Forwarding Parameters IE in FAR
Octet 1 and 2		Update Forwarding Parameters IE Type = 11 (decimal)
Octets 3 and 4		Length = n
Information elements	P        IE Type

Destination Interface	C        Destination Interface
Network instance	    C        Network Instance
Redirect Information	C        Redirect Information
Outer Header Creation 	C        Outer Header Creation
Transport Level Marking	C        Transport Level Marking
Forwarding Policy 	C            Forwarding Policy
Header Enrichment	C            Header Enrichment
PFCPSMReq-Flags	    C            PFCPSMReq-Flags
Linked Traffic Endpoint ID	C    Traffic Endpoint ID

*/
type IEUpdateForwardingParameters struct {
	IETypeLength
	DstInterface          *IEDestinationInterface  `json:",omitempty"`
	NetworkInstance       *IENetworkInstance       `json:",omitempty"`
	RedirectInfo          *IERedirectInformation   `json:",omitempty"`
	OuterHeaderCreation   *IEOuterHeaderCreation   `json:",omitempty"`
	TransportLevelMarking *IETransportLevelMarking `json:",omitempty"`
	ForwardingPolicy      *IEForwardingPolicy      `json:",omitempty"`
	HeaderEnrichment      *IEHeaderEnrichment      `json:",omitempty"`
	PFCPSMReqFlags        *IEPFCPSMReqFlags        `json:",omitempty"`
	TrafficEndpointID     *IETrafficEndpointID     `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUpdateForwardingParameters) Encode() (data []byte, err error) {
	//	encode ie
	encBuf := bytes.NewBuffer(nil)

	//	Mandatory ie
	var vEnc, tlvEnc []byte
	var tl IETypeLength

	//  optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Destination Interface	C        Destination Interface
		case IE_Destination_Interface:
			//	encode v
			vEnc, err = i.DstInterface.Encode()
			if err != nil {
				return nil, err
			}
			//Network instance	    C        Network Instance
		case IE_Network_Instance:
			//	encode v
			vEnc, err = i.NetworkInstance.Encode()
			if err != nil {
				return nil, err
			}
			//Redirect Information	C        Redirect Information
		case IE_Redirect_Information:
			//	encode v
			vEnc, err = i.RedirectInfo.Encode()
			if err != nil {
				return nil, err
			}
			//Outer Header Creation 	C        Outer Header Creation
		case IE_Outer_Header_Creation:
			//	encode v
			vEnc, err = i.OuterHeaderCreation.Encode()
			if err != nil {
				return nil, err
			}
			//Transport Level Marking	C        Transport Level Marking
		case IE_Transport_Level_Marking:
			//	encode v
			vEnc, err = i.TransportLevelMarking.Encode()
			if err != nil {
				return nil, err
			}
			//Forwarding Policy 	C            Forwarding Policy
		case IE_Forwarding_Policy:
			//	encode v
			vEnc, err = i.ForwardingPolicy.Encode()
			if err != nil {
				return nil, err
			}
			//Header Enrichment	C            Header Enrichment
		case IE_Header_Enrichment:
			//	encode v
			vEnc, err = i.HeaderEnrichment.Encode()
			if err != nil {
				return nil, err
			}
			//PFCPSMReq-Flags	    C            PFCPSMReq-Flags
		case IE_PFCPSMReq_Flags:
			//	encode v
			vEnc, err = i.PFCPSMReqFlags.Encode()
			if err != nil {
				return nil, err
			}
			//Linked Traffic Endpoint ID	C    Traffic Endpoint ID
		case IE_Traffic_Endpoint_ID:
			//	encode v
			vEnc, err = i.TrafficEndpointID.Encode()
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

func (i *IEUpdateForwardingParameters) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUpdateForwardingParameters) Len() int {
	return int(i.Length)
}

func (i *IEUpdateForwardingParameters) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUpdateForwardingParameters) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Destination Interface	C        Destination Interface
	case *IEDestinationInterface:
		i.DstInterface = ie
		//Network instance	    C        Network Instance
	case *IENetworkInstance:
		i.NetworkInstance = ie
		//Redirect Information	C        Redirect Information
	case *IERedirectInformation:
		i.RedirectInfo = ie
		//Outer Header Creation 	C        Outer Header Creation
	case *IEOuterHeaderCreation:
		i.OuterHeaderCreation = ie
		//Transport Level Marking	C        Transport Level Marking
	case *IETransportLevelMarking:
		i.TransportLevelMarking = ie
		//Forwarding Policy 	C            Forwarding Policy
	case *IEForwardingPolicy:
		i.ForwardingPolicy = ie
		//Header Enrichment	C            Header Enrichment
	case *IEHeaderEnrichment:
		i.HeaderEnrichment = ie
		//PFCPSMReq-Flags	    C            PFCPSMReq-Flags
	case *IEPFCPSMReqFlags:
		i.PFCPSMReqFlags = ie
		//Linked Traffic Endpoint ID	C    Traffic Endpoint ID
	case *IETrafficEndpointID:
		i.TrafficEndpointID = ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEUpdateForwardingParameters) Set() error {
	i.Type = IE_Update_Forwarding_Parameters

	return nil
}

//PFCPSMReq-Flags	    C
// IEPFCPSMReqFlags
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
49	PFCPSMReq-Flags	Extendable / Subclause 8.2.31	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 49 (decimal)
	3 to 4	Length = n
	5	Spare	Spare	Spare	Spare	Spare	QAURR	SNDEM	DROBU
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.31-1: PFCPSMReq-Flags
*/
const (
	IEPFCPSMReqFlags_QAURR = IEFlag_Bit_3
	IEPFCPSMReqFlags_SNDEM = IEFlag_Bit_2
	IEPFCPSMReqFlags_DROBU = IEFlag_Bit_1
)

type IEPFCPSMReqFlags struct {
	IETypeLength
	DROBU bool
	SNDEM bool
	QAURR bool
}

func (i *IEPFCPSMReqFlags) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := utils.BoolToUint8(i.DROBU) +
		(utils.BoolToUint8(i.SNDEM) << 1) +
		(utils.BoolToUint8(i.QAURR) << 2)

	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEPFCPSMReqFlags) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEPFCPSMReqFlags) Len() int {
	return int(i.Length)
}

func (i *IEPFCPSMReqFlags) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPFCPSMReqFlags) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

//func (i *IEPFCPSMReqFlags) Set(v uint8) error {
//	i.Type = IE_PFCPSMReq_Flags
//	i.Length = 1
//	i.Flag = v
//
//	return nil
//}
//func (i *IEPFCPSMReqFlags) Get() (v uint8, e error) {
//	return i.Flag, nil
//}

// IEUpdateForwardingParameters
//End--------------------------------------------------------------------------

//Update Duplicating Parameters
// IEUpdateDuplicatingParameters
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
105	Update Duplicating Parameters	Extendable / Table 7.5.4.3-3	Not Applicable
*/
/*Table 7.5.4.3-3: Update Duplicating Parameters IE in FAR
Octet 1 and 2		Update Duplicating Parameters IE Type = 105 (decimal)
Octets 3 and 4		Length = n
Information elements	P      IE Type

Destination Interface	C      Destination Interface
Outer Header Creation 	C      Outer Header Creation
Transport Level Marking C      Transport Level Marking
Forwarding Policy 	C          Forwarding Policy

*/
type IEUpdateDuplicatingParameters struct {
	IETypeLength
	DstInterface          *IEDestinationInterface
	OuterHeaderCreation   *IEOuterHeaderCreation
	TransportLevelMarking *IETransportLevelMarking
	ForwardingPolicy      *IEForwardingPolicy

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUpdateDuplicatingParameters) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie

	var vEnc, tlvEnc []byte
	var tl IETypeLength
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Destination Interface	C      Destination Interface
		case IE_Destination_Interface:
			// encode v
			vEnc, err = i.DstInterface.Encode()
			if err != nil {
				return nil, err
			}
			//Outer Header Creation 	C      Outer Header Creation
		case IE_Outer_Header_Creation:
			// encode v
			vEnc, err = i.OuterHeaderCreation.Encode()
			if err != nil {
				return nil, err
			}
			//Transport Level Marking C      Transport Level Marking
		case IE_Transport_Level_Marking:
			// encode v
			vEnc, err = i.TransportLevelMarking.Encode()
			if err != nil {
				return nil, err
			}
			//Forwarding Policy 	C          Forwarding Policy
		case IE_Forwarding_Policy:
			// encode v
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

func (i *IEUpdateDuplicatingParameters) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUpdateDuplicatingParameters) Len() int {
	return int(i.Length)
}

func (i *IEUpdateDuplicatingParameters) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUpdateDuplicatingParameters) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Destination Interface	C      Destination Interface
	case *IEDestinationInterface:
		i.DstInterface = ie
		//Outer Header Creation 	C      Outer Header Creation
	case *IEOuterHeaderCreation:
		i.OuterHeaderCreation = ie
		//Transport Level Marking C      Transport Level Marking
	case *IETransportLevelMarking:
		i.TransportLevelMarking = ie
		//Forwarding Policy 	C          Forwarding Policy
	case *IEForwardingPolicy:
		i.ForwardingPolicy = ie

	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IEUpdateDuplicatingParameters) Set() error {
	i.Type = IE_Update_Duplicating_Parameters

	return nil
}

// IEUpdateFAR
//End--------------------------------------------------------------------------

//Update URR	C                                 Update URR
// IEUpdateURR
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
13	Update URR	Extendable / Table 7.5.4.4	Not Applicable
*/
/*Table 7.5.4.4-1: Update URR IE within PFCP Session Modification Request
Octet 1 and 2		Update URR IE Type = 13 (decimal)
Octets 3 and 4		Length = n
Information elements	P            IE Type

URR ID	M                            URR ID
Measurement Method	C                Measurement Method
Reporting Triggers	C                Reporting Triggers
Measurement Period 	C                Measurement Period
Volume Threshold	C                Volume Threshold
Volume Quota	C                    Volume Quota
Time Threshold	C                    Time Threshold
Time Quota	C                        Time Quota

Event Threshold	C                    Event Threshold
Event Quota	C                        Event Quota

Quota Holding Time	C                Quota Holding Time
Dropped DL Traffic Threshold C       Dropped DL Traffic Threshold
Monitoring Time	C                    Monitoring Time

Subsequent Volume Threshold	C        Subsequent Volume Threshold
Subsequent Time Threshold	C        Subsequent Time Threshold

Subsequent Volume Quota	C            Subsequent Volume Quota
Subsequent Time Quota	C            Subsequent Time Quota

Subsequent Event Threshold	O        Subsequent Event Threshold
Subsequent EventQuota	O            Subsequent EventQuota

Inactivity Detection Time	C        Inactivity Detection Time
Linked URR ID 	C                    Linked URR ID
Measurement Information	C            Measurement Information

FAR ID for Quota Action	C            FAR ID
Ethernet Inactivity Timer	C        Ethernet Inactivity Timer
Additional Monitoring Time	O        Additional Monitoring Time
*/

type IEUpdateURR struct {
	IETypeLength
	URRID             IEURRID
	MeasurementMethod *IEMeasurementMethod
	ReportingTriggers *IEReportingTriggers
	MeasurementPeriod *IEMeasurementPeriod
	VolumeThreshold   *IEVolumeThreshold
	VolumeQuota       *IEVolumeQuota
	TimeThreshold     *IETimeThreshold
	TimeQuota         *IETimeQuota

	EventThreshold *IEEventThreshold
	EventQuota     *IEEventQuota

	QuotaHoldingTime          *IEQuotaHoldingTime
	DroppedDLTrafficThreshold *IEDroppedDLTrafficThreshold
	MonitoringTime            *IEMonitoringTime

	SubsequentVolumeThreshold *IESubsequentVolumeThreshold
	SubsequentTimeThreshold   *IESubsequentTimeThreshold

	SubsequentVolumeQuota *IESubsequentVolumeQuota
	SubsequentTimeQuota   *IESubsequentTimeQuota

	SubsequentEventThreshold *IESubsequentEventThreshold
	SubsequentEventQuota     *IESubsequentEventQuota

	InactivityDetectionTime *IEInactivityDetectionTime
	LinkedURRID             *IELinkedURRID
	MeasurementInformation  *IEMeasurementInformation

	FARID                    *IEFARID
	EthernetInactivityTimer  *IEEthernetInactivityTimer
	AdditionalMonitoringTime *IEAdditionalMonitoringTime

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUpdateURR) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//URR ID	M                            URR ID
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

	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Measurement Method	C                Measurement Method
		case IE_Measurement_Method:
			// encode v
			vEnc, err = i.MeasurementMethod.Encode()
			if err != nil {
				return nil, err
			}
			//Reporting Triggers	C                Reporting Triggers
		case IE_Reporting_Triggers:
			// encode v
			vEnc, err = i.ReportingTriggers.Encode()
			if err != nil {
				return nil, err
			}

			//Measurement Period 	C                Measurement Period
		case IE_Measurement_Period:
			// encode v
			vEnc, err = i.MeasurementPeriod.Encode()
			if err != nil {
				return nil, err
			}

			//Volume Threshold	C                Volume Threshold
		case IE_Volume_Threshold:
			// encode v
			vEnc, err = i.VolumeThreshold.Encode()
			if err != nil {
				return nil, err
			}

			//Volume Quota	C                    Volume Quota
		case IE_Volume_Quota:
			// encode v
			vEnc, err = i.VolumeQuota.Encode()
			if err != nil {
				return nil, err
			}

			//Time Threshold	C                    Time Threshold
		case IE_Time_Threshold:
			// encode v
			vEnc, err = i.TimeThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Time Quota	C                        Time Quota
		case IE_Time_Quota:
			// encode v
			vEnc, err = i.TimeQuota.Encode()
			if err != nil {
				return nil, err
			}

			//
			//Event Threshold	C                    Event Threshold
		case IE_Event_Threshold:
			// encode v
			vEnc, err = i.EventThreshold.Encode()
			if err != nil {
				return nil, err
			}

			//Event Quota	C                        Event Quota
		case IE_Event_Quota:
			// encode v
			vEnc, err = i.EventQuota.Encode()
			if err != nil {
				return nil, err
			}
			//
			//Quota Holding Time	C                Quota Holding Time
		case IE_Quota_Holding_Time:
			// encode v
			vEnc, err = i.QuotaHoldingTime.Encode()
			if err != nil {
				return nil, err
			}
			//Dropped DL Traffic Threshold C       Dropped DL Traffic Threshold
		case IE_Dropped_DL_Traffic_Threshold:
			// encode v
			vEnc, err = i.DroppedDLTrafficThreshold.Encode()
			if err != nil {
				return nil, err
			}

			//Monitoring Time	C                    Monitoring Time
		case IE_Monitoring_Time:
			// encode v
			vEnc, err = i.MonitoringTime.Encode()
			if err != nil {
				return nil, err
			}

			//
			//Subsequent Volume Threshold	C        Subsequent Volume Threshold
		case IE_Subsequent_Volume_Threshold:
			// encode v
			vEnc, err = i.SubsequentVolumeThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent Time Threshold	C        Subsequent Time Threshold
		case IE_Subsequent_Time_Threshold:
			// encode v
			vEnc, err = i.SubsequentTimeThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//
			//Subsequent Volume Quota	C            Subsequent Volume Quota
		case IE_Subsequent_Volume_Quota:
			// encode v
			vEnc, err = i.SubsequentVolumeQuota.Encode()
			if err != nil {
				return nil, err
			}

			//Subsequent Time Quota	C            Subsequent Time Quota
		case IE_Subsequent_Time_Quota:
			// encode v
			vEnc, err = i.SubsequentTimeQuota.Encode()
			if err != nil {
				return nil, err
			}
			//
			//Subsequent Event Threshold	O        Subsequent Event Threshold
		case IE_Subsequent_Event_Threshold:
			// encode v
			vEnc, err = i.SubsequentEventThreshold.Encode()
			if err != nil {
				return nil, err
			}
			//Subsequent EventQuota	O            Subsequent EventQuota
		case IE_Subsequent_Event_Quota:
			// encode v
			vEnc, err = i.SubsequentEventQuota.Encode()
			if err != nil {
				return nil, err
			}
			//
			//Inactivity Detection Time	C        Inactivity Detection Time
		case IE_Inactivity_Detection_Time:
			// encode v
			vEnc, err = i.InactivityDetectionTime.Encode()
			if err != nil {
				return nil, err
			}
			//Linked URR ID 	C                    Linked URR ID
		case IE_LinkedURR_ID:
			// encode v
			vEnc, err = i.LinkedURRID.Encode()
			if err != nil {
				return nil, err
			}

			//Measurement Information	C            Measurement Information
		case IE_Measurement_Information:
			// encode v
			vEnc, err = i.MeasurementInformation.Encode()
			if err != nil {
				return nil, err
			}
			//
			//FAR ID for Quota Action	C            FAR ID
		case IE_FAR_ID:
			// encode v
			vEnc, err = i.FARID.Encode()
			if err != nil {
				return nil, err
			}
			//Ethernet Inactivity Timer	C        Ethernet Inactivity Timer
		case IE_Ethernet_Inactivity_Timer:
			// encode v
			vEnc, err = i.EthernetInactivityTimer.Encode()
			if err != nil {
				return nil, err
			}
			//Additional Monitoring Time	O        Additional Monitoring Time
		case IE_Additional_Monitoring_Time:
			// encode v
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

func (i *IEUpdateURR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUpdateURR) Len() int {
	return int(i.Length)
}

func (i *IEUpdateURR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUpdateURR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//URR ID	M                            URR ID
	case *IEURRID:
		i.URRID = *ie
		//Measurement Method	C                Measurement Method
	case *IEMeasurementMethod:
		i.MeasurementMethod = ie
		//Reporting Triggers	C                Reporting Triggers
	case *IEReportingTriggers:
		i.ReportingTriggers = ie
		//Measurement Period 	C                Measurement Period
	case *IEMeasurementPeriod:
		i.MeasurementPeriod = ie
		//Volume Threshold	C                Volume Threshold
	case *IEVolumeThreshold:
		i.VolumeThreshold = ie
		//Volume Quota	C                    Volume Quota
	case *IEVolumeQuota:
		i.VolumeQuota = ie
		//Time Threshold	C                    Time Threshold
	case *IETimeThreshold:
		i.TimeThreshold = ie
		//Time Quota	C                        Time Quota
	case *IETimeQuota:
		i.TimeQuota = ie
		//
		//Event Threshold	C                    Event Threshold
	case *IEEventThreshold:
		i.EventThreshold = ie
		//Event Quota	C                        Event Quota
	case *IEEventQuota:
		i.EventQuota = ie
		//
		//Quota Holding Time	C                Quota Holding Time
	case *IEQuotaHoldingTime:
		i.QuotaHoldingTime = ie
		//Dropped DL Traffic Threshold C       Dropped DL Traffic Threshold
	case *IEDroppedDLTrafficThreshold:
		i.DroppedDLTrafficThreshold = ie
		//Monitoring Time	C                    Monitoring Time
	case *IEMonitoringTime:
		i.MonitoringTime = ie
		//
		//Subsequent Volume Threshold	C        Subsequent Volume Threshold
	case *IESubsequentVolumeThreshold:
		i.SubsequentVolumeThreshold = ie
		//Subsequent Time Threshold	C        Subsequent Time Threshold
	case *IESubsequentTimeThreshold:
		i.SubsequentTimeThreshold = ie
		//
		//Subsequent Volume Quota	C            Subsequent Volume Quota
	case *IESubsequentVolumeQuota:
		i.SubsequentVolumeQuota = ie
		//Subsequent Time Quota	C            Subsequent Time Quota
	case *IESubsequentTimeQuota:
		i.SubsequentTimeQuota = ie
		//
		//Subsequent Event Threshold	O        Subsequent Event Threshold
	case *IESubsequentEventThreshold:
		i.SubsequentEventThreshold = ie
		//Subsequent EventQuota	O            Subsequent EventQuota
	case *IESubsequentEventQuota:
		i.SubsequentEventQuota = ie
		//
		//Inactivity Detection Time	C        Inactivity Detection Time
	case *IEInactivityDetectionTime:
		i.InactivityDetectionTime = ie
		//Linked URR ID 	C                    Linked URR ID
	case *IELinkedURRID:
		i.LinkedURRID = ie
		//Measurement Information	C            Measurement Information
	case *IEMeasurementInformation:
		i.MeasurementInformation = ie
		//
		//FAR ID for Quota Action	C            FAR ID
	case *IEFARID:
		i.FARID = ie
		//Ethernet Inactivity Timer	C        Ethernet Inactivity Timer
	case *IEEthernetInactivityTimer:
		i.EthernetInactivityTimer = ie
		//Additional Monitoring Time	O        Additional Monitoring Time
	case *IEAdditionalMonitoringTime:
		i.AdditionalMonitoringTime = ie

	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IEUpdateURR) Set(v uint8) error {
	i.Type = IE_Update_URR

	return nil
}

//Update QER	C                                 Update QER
// IEUpdateQER
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
14	Update QER	Extendable / Table 7.5.4.5	Not Applicable
*/
/*Table 7.5.4.5-1: Update QER IE within PFCP Session Modification Request
Octet 1 and 2		Update QER IE Type = 14 (decimal)
Octets 3 and 4		Length = n
Information elements	P   IE Type

QER ID	M                   QER ID
QER Correlation ID	C       QER Correlation ID
Gate Status	C               Gate Status
Maximum Bitrate	C           MBR
Guaranteed Bitrate	C       GBR
QoS flow identifier	C       QFI
Reflective QoS	C           RQI

Paging Policy Indicator	C   Paging Policy Indicato
Averaging Window	O       Averaging Window
*/
type IEUpdateQER struct {
	IETypeLength
	QERID                 IEQERID
	QERCorrelationID      *IEQERCorrelationID
	GateStatus            *IEGateStatus
	MaximumBitrate        *IEMBR
	GuaranteedBitrate     *IEGBR
	QoSflowidentifier     *IEQFI
	ReflectiveQoS         *IERQI
	PagingPolicyIndicator *IEPagingPolicyIndicator
	AveragingWindow       *IEAveragingWindow

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUpdateQER) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//QER ID	M                   QER ID
	//encode v
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
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//QER Correlation ID	C       QER Correlation ID
		case IE_QER_Correlation_ID:
			vEnc, err = i.QERCorrelationID.Encode()
			if err != nil {
				return nil, err
			}
			//Gate Status	C               Gate Status
		case IE_Gate_Status:
			vEnc, err = i.GateStatus.Encode()
			if err != nil {
				return nil, err
			}
			//Maximum Bitrate	C           MBR
		case IE_MBR:
			vEnc, err = i.MaximumBitrate.Encode()
			if err != nil {
				return nil, err
			}
			//Guaranteed Bitrate	C       GBR
		case IE_GBR:
			vEnc, err = i.GuaranteedBitrate.Encode()
			if err != nil {
				return nil, err
			}
			//QoS flow identifier	C       QFIs
		case IE_QFI:
			vEnc, err = i.QoSflowidentifier.Encode()
			if err != nil {
				return nil, err
			}
			//Reflective QoS	C           RQI
		case IE_RQI:
			vEnc, err = i.ReflectiveQoS.Encode()
			if err != nil {
				return nil, err
			}
			//
			//Paging Policy Indicator	C   Paging Policy Indicato
		case IE_Paging_Policy_Indicator:
			vEnc, err = i.PagingPolicyIndicator.Encode()
			if err != nil {
				return nil, err
			}

			//Averaging Window	O       Averaging Window
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

func (i *IEUpdateQER) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUpdateQER) Len() int {
	return int(i.Length)
}

func (i *IEUpdateQER) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUpdateQER) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//QER ID	M                   QER ID
	case *IEQERID:
		i.QERID = *ie
		//QER Correlation ID	C       QER Correlation ID
	case *IEQERCorrelationID:
		i.QERCorrelationID = ie
		//Gate Status	C               Gate Status
	case *IEGateStatus:
		i.GateStatus = ie
		//Maximum Bitrate	C           MBR
	case *IEMBR:
		i.MaximumBitrate = ie
		//Guaranteed Bitrate	C       GBR
	case *IEGBR:
		i.GuaranteedBitrate = ie
		//QoS flow identifier	C       QFI
	case *IEQFI:
		i.QoSflowidentifier = ie
		//Reflective QoS	C           RQI
	case *IERQI:
		i.ReflectiveQoS = ie
		//
		//Paging Policy Indicator	C   Paging Policy Indicato
	case *IEPagingPolicyIndicator:
		i.PagingPolicyIndicator = ie
		//Averaging Window	O       Averaging Window
	case *IEAveragingWindow:
		i.AveragingWindow = ie

	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IEUpdateQER) Set(v uint8) error {
	i.Type = IE_Update_QER

	return nil
}

//Update BAR	C                                 Update BAR
// IEUpdateBAR
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
86	Update BAR (Session Modification Request)	Extendable / Table 7.5.4.11-1	Not Applicable
*/
/*Table 7.5.4.11-1: Update BAR IE within PFCP Session Modification Request
Octet 1 and 2                              Update BAR IE Type = 86 (decimal)
Octets 3 and 4                             Length = n
Information elements	P                  IE Type

BAR ID	M                                  BAR ID
Downlink Data Notification Delay	C      Downlink Data Notification Delay
Suggested Buffering Packets Count	C      Suggested Buffering Packets Count
*/
type IEUpdateBARForSMR struct {
	IETypeLength
	BARID                          IEBARID
	DLDataNotificationDelay        IEDownlinkDataNotificationDelay
	SuggestedBufferingPacketsCount IESuggestedBufferingPacketsCount

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUpdateBARForSMR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//BAR ID	M                                  BAR ID
	vEnc, err := i.BARID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_BAR_ID), Length: uint16(len(vEnc))}
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
		//Downlink Data Notification Delay	C
		case IE_Downlink_Data_Notification_Delay:
			vEnc, err = i.DLDataNotificationDelay.Encode()
			if err != nil {
				return nil, err
			}
			//Suggested Buffering Packets Count	C
		case IE_Suggested_Buffering_Packets_Count:
			vEnc, err = i.SuggestedBufferingPacketsCount.Encode()
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

func (i *IEUpdateBARForSMR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUpdateBARForSMR) Len() int {
	return int(i.Length)
}

func (i *IEUpdateBARForSMR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUpdateBARForSMR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//BAR ID	M                                  BAR ID
	case *IEBARID:
		i.BARID = *ie
		//Downlink Data Notification Delay	C      Downlink Data Notification Delay
	case *IEDownlinkDataNotificationDelay:
		i.DLDataNotificationDelay = *ie
		//Suggested Buffering Packets Count	C      Suggested Buffering Packets Count
	case *IESuggestedBufferingPacketsCount:
		i.SuggestedBufferingPacketsCount = *ie

	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IEUpdateBARForSMR) Set() error {
	i.Type = IE_Update_BAR_Modification_Request

	return nil
}

//Downlink Data Notification Delay	C
// IEDownlinkDataNotificationDelay
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
46	Downlink Data Notification Delay	Extendable / Subclause 8.2.28	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 46 (decimal)
	3 to 4	Length = n
	5	Delay Value in integer multiples of 50 millisecs, or zero
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.28-1: Downlink Data Notification Delay
*/
type IEDownlinkDataNotificationDelay struct {
	IETypeLength
	Value uint8 `json:",omitempty"`
}

func (i *IEDownlinkDataNotificationDelay) Encode() (data []byte, err error) {
	// encode v
	return []byte{i.Value}, nil
}

func (i *IEDownlinkDataNotificationDelay) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.Value = tmp

	return nil
}

func (i *IEDownlinkDataNotificationDelay) Len() int {
	return int(i.Length)
}

func (i *IEDownlinkDataNotificationDelay) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEDownlinkDataNotificationDelay) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

const (
	IEDownlinkDataNotificationDelay_Unit = 50 * time.Millisecond
)

func (i *IEDownlinkDataNotificationDelay) Set(v uint8) error {
	i.Type = IE_Downlink_Data_Notification_Delay
	i.Length = 1
	i.Value = v

	return nil
}
func (i *IEDownlinkDataNotificationDelay) Get() (v uint8, e error) {
	return i.Value, nil
}

//DL Buffering Duration	C
// IEDLBufferingDuration
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
47	DL Buffering Duration	Extendable / Subclause 8.2.29	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 47 (decimal)
	3 to 4	Length = n
	5	Timer unit	Timer value
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.29-1: DL Buffering Duration
*/
type IEDLBufferingDuration struct {
	IETimer
}

func (i *IEDLBufferingDuration) Set(v uint8, unit uint8) error {
	i.IETimer.Set(v, unit)
	i.Type = IE_DL_Buffering_Duration

	return nil
}
func (i *IEDLBufferingDuration) Get() (v uint8, unit uint8, e error) {
	return i.IETimer.Get()
}

//DL Buffering Suggested Packet Count	O
// IEDLBufferingSuggestedPacketCount
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
48	DL Buffering Suggested Packet Count	Variable / Subclause 8.2.30	Not Applicable
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 48 (decimal)
	3 to 4	Length = n
	5 to n+4	Packet Count Value
Figure 8.2.30-1: DL Buffering Suggested Packet Count
*/
type IEDLBufferingSuggestedPacketCount struct {
	IETypeLength
	CountValue uint16 `json:",omitempty"`

	//The length shall be set to 1 or 2 octets.
	flag uint8
}

func (i *IEDLBufferingSuggestedPacketCount) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//The length shall be set to 1 or 2 octets.
	if i.flag == 2 {
		err = binary.Write(encBuf, binary.BigEndian, i.CountValue)
		if err != nil {
			return nil, err
		}
	}
	if i.flag == 1 {
		tmp := byte(i.CountValue)
		err = binary.Write(encBuf, binary.BigEndian, tmp)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEDLBufferingSuggestedPacketCount) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	if i.Length == 2 {
		err := binary.Read(r, binary.BigEndian, &i.CountValue)
		if err != nil {
			return err
		}
	}
	if i.Length == 1 {
		tmp, err := r.ReadByte()
		if err != nil {
			return err
		}
		i.CountValue = uint16(tmp)
	}
	return nil
}

func (i *IEDLBufferingSuggestedPacketCount) Len() int {
	return int(i.Length)
}

func (i *IEDLBufferingSuggestedPacketCount) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEDLBufferingSuggestedPacketCount) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEDLBufferingSuggestedPacketCount) Set(v uint16) error {
	i.Type = IE_DL_Buffering_Suggested_Packet_Count
	i.Length = 2
	i.CountValue = v
	return nil
}
func (i *IEDLBufferingSuggestedPacketCount) Get() (v uint16, e error) {
	return i.CountValue, nil
}

// IEUpdateBAR
//End--------------------------------------------------------------------------

//Update Traffic Endpoint	C                     Update Traffic Endpoint
// IEUpdateTrafficEndpoint
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
129	Update Traffic Endpoint	Extendable / Table 7.5.4.13	Not Applicable
*/
/*Table 7.5.4.13-1: Update Traffic Endpoint IE within Sx Session Modification Request
Octet 1 and 2	Update Traffic Endpoint Type = 129 (decimal)
Octets 3 and 4	Length = n
Information elements	P    IE Type

Traffic Endpoint ID	   M     Traffic Endpoint ID
Local F-TEID 	       C     F-TEID
Network Instance	   O     Network Instance
UE IP address 	C            UE IP address
Framed-Route	C            Framed-Route
Framed-Routing	C            Framed-Routing
Framed-IPv6-Route	C        Framed-IPv6-Route

*/
type IEUpdateTrafficEndpoint struct {
	IETypeLength
	TrafficEndpointID IETrafficEndpointID
	LocalFTEID        *IEFTEID
	NetworkInstance   *IENetworkInstance
	UEIPaddress       *IEUEIPaddress
	FramedRoute       *IEFramedRoute
	FramedRouting     *IEFramedRouting
	FramedIPv6Route   *IEFramedIPv6Route

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUpdateTrafficEndpoint) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Traffic Endpoint ID	   M     Traffic Endpoint ID
	vEnc, err := i.TrafficEndpointID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{uint16(IE_Traffic_Endpoint_ID), uint16(len(vEnc))}
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
		//Local F-TEID 	       C     F-TEID
		case IE_F_TEID:
			//	encode v
			vEnc, err = i.LocalFTEID.Encode()
			if err != nil {
				return nil, err
			}
			//Network Instance	   O     Network Instance
		case IE_Network_Instance:
			//	encode v
			vEnc, err = i.NetworkInstance.Encode()
			if err != nil {
				return nil, err
			}
			//UE IP address 	C            UE IP address
		case IE_UE_IP_Address:
			//	encode v
			vEnc, err = i.UEIPaddress.Encode()
			if err != nil {
				return nil, err
			}
			//Framed-Route	C            Framed-Route
		case IE_Framed_Route:
			//	encode v
			vEnc, err = i.FramedRoute.Encode()
			if err != nil {
				return nil, err
			}
			//Framed-Routing	C            Framed-Routing
		case IE_Framed_Routing:
			//	encode v
			vEnc, err = i.FramedRouting.Encode()
			if err != nil {
				return nil, err
			}
			//Framed-IPv6-Route	C        Framed-IPv6-Route
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

func (i *IEUpdateTrafficEndpoint) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUpdateTrafficEndpoint) Len() int {
	return int(i.Length)
}

func (i *IEUpdateTrafficEndpoint) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUpdateTrafficEndpoint) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Traffic Endpoint ID	   M     Traffic Endpoint ID
	case *IETrafficEndpointID:
		i.TrafficEndpointID = *ie
		//Local F-TEID 	       C     F-TEID
	case *IEFTEID:
		i.LocalFTEID = ie
		//Network Instance	   O     Network Instance
	case *IENetworkInstance:
		i.NetworkInstance = ie
		//UE IP address 	C            UE IP address
	case *IEUEIPaddress:
		i.UEIPaddress = ie
		//Framed-Route	C            Framed-Route
	case *IEFramedRoute:
		i.FramedRoute = ie
		//Framed-Routing	C            Framed-Routing
	case *IEFramedRouting:
		i.FramedRouting = ie
		//Framed-IPv6-Route	C        Framed-IPv6-Route
	case *IEFramedIPv6Route:
		i.FramedIPv6Route = ie

	default:
		return fmt.Errorf("Illegal IE")
	}
	return nil
}

func (i *IEUpdateTrafficEndpoint) Set(v uint8) error {
	i.Type = IE_Update_Traffic_Endpoint

	return nil
}

//PFCPSMReq-Flags	C                             PFCPSMReq-Flags

//Query URR	C                                 Query URR
// IEQueryURR
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
77	Query URR	Extendable / Table 7.5.4.10-1	Not Applicable
*/
/*Table 7.5.4.10-1: Query URR IE within PFCP Session Modification Request
Octet 1 and 2		Query URR IE Type = 77 (decimal)
Octets 3 and 4		Length = n
Information elements	P   IE Type

URR ID	                M   URR ID
*/
type IEQueryURR struct {
	IEURRID
}

func (i *IEQueryURR) Set(v uint32) error {
	i.IEURRID.Set(v)
	i.Type = IE_Query_URR

	return nil
}
func (i *IEQueryURR) Get() (v uint32, e error) {
	return i.IEURRID.Get()
}

//User Plane Inactivity Timer	C                 User Plane Inactivity Timer
//Query URR Reference	O                         Query URR Reference
// IEQueryURRReference
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
125	Query URR Reference	Extendable / Subclause 8.2.90	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 125 (decimal)
	3 to 4	Length = n
	5 to 8	Query URR Reference value
	9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.90-1: Query URR Reference
*/
type IEQueryURRReference struct {
	IEURRID
}

func (i *IEQueryURRReference) Set(v uint32) error {
	i.IEURRID.Set(v)
	i.Type = IE_Query_URR_Reference

	return nil
}
func (i *IEQueryURRReference) Get() (v uint32, e error) {
	return i.IEURRID.Get()
}

//Trace Information   O                         Trace Information

// end 1488 20200113
