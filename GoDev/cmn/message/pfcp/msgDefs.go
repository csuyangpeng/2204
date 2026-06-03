package pfcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"io"
	"reflect"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)
const Port = 8805

type PFCPMSG uint8

// 7.3	Message Types
const (
	//PFCP Node related messages
	PFCP_Heartbeat_Request              = 1
	PFCP_Heartbeat_Response             = 2
	PFCP_PFD_Management_Request         = 3
	PFCP_PFD_Management_Response        = 4
	PFCP_Association_Setup_Request      = 5
	PFCP_Association_Setup_Response     = 6
	PFCP_Association_Update_Request     = 7
	PFCP_Association_Update_Response    = 8
	PFCP_Association_Release_Request    = 9
	PFCP_Association_Release_Response   = 10
	PFCP_Version_Not_Supported_Response = 11
	PFCP_Node_Report_Request            = 12
	PFCP_Node_Report_Response           = 13

	//PFCP Session related messages
	PFCP_Session_Establishment_Request  = 50 // 完成
	PFCP_Session_Establishment_Response = 51 // 完成
	PFCP_Session_Modification_Request   = 52
	PFCP_Session_Modification_Response  = 53
	PFCP_Session_Deletion_Request       = 54 // 完成
	PFCP_Session_Deletion_Response      = 55 // 完成
	PFCP_Session_Report_Request         = 56
	PFCP_Session_Report_Response        = 57
)

type PfcpMsgInterface interface {
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) error
	Len() int
}

type PfcpIeInterface interface {
	Encode() (data []byte, err error)
	Decode(data []byte) error
	Len() int

	SetObject(t uint16, l uint16) error
	SetObjectToParent(child interface{}) error
}

type PfcpGroupedIeInterface interface {
	GetGrouped() bool
}

// 7.5.8	PFCP Session Report Request
type SessionReportRequest struct {
	PfcpHeader PfcpHeaderforSession
	IE         IEsSessionReportRequest
}

func (s *SessionReportRequest) MarshalBinary() (data []byte, err error) {
	// IE encode
	data, err = s.IE.Encode()
	if err != nil {
		return nil, err
	}
	return
}

func (s *SessionReportRequest) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	err := s.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionReportRequest) Len() int {
	return int(s.PfcpHeader.Length)
}

// 7.5.9	PFCP Session Report Response
type SessionReportResponse struct {
	PfcpHeader PfcpHeaderforSession
	IE         IEsSessionReportResponse
}

func (s *SessionReportResponse) MarshalBinary() (data []byte, err error) {
	// IE encode
	data, err = s.IE.Encode()
	if err != nil {
		return nil, err
	}
	return
}

func (s *SessionReportResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	err := s.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionReportResponse) Len() int {
	return int(s.PfcpHeader.Length)
}

// 7.5.6	PFCP Session Deletion Request
type SessionReleaseRequest struct {
	// alias PfcpSessionDelRequest
	PfcpHeader PfcpHeaderforSession
	IE         IEsSessionDelRequest
}

func (s *SessionReleaseRequest) MarshalBinary() (data []byte, err error) {
	// IE encode
	data, err = s.IE.Encode()
	if err != nil {
		return nil, err
	}
	return
}

func (s *SessionReleaseRequest) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	err := s.IE.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionReleaseRequest) Len() int {
	return int(s.PfcpHeader.Length)
}

// 7.5.7	PFCP Session Deletion Response
/*
Information elements	P

Cause	                M
Offending IE	        C
Load Control Information	O
Overload Control Information	O
Usage Report	        C
*/
type SessionReleaseResponse struct {
	// alias PfcpSessionDelResponse
	PfcpHeader PfcpHeaderforSession
	IE         IEsSessionDelResponse
}

func (s *SessionReleaseResponse) MarshalBinary() (data []byte, err error) {
	// IE encode
	data, err = s.IE.Encode()
	if err != nil {
		return nil, err
	}
	return
}

func (s *SessionReleaseResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	err := s.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionReleaseResponse) Len() int {
	return int(s.PfcpHeader.Length)
}

// 7.5.2	PFCP Session Establishment Request
type SessionEstablishmentRequest struct {
	PfcpHeader PfcpHeaderforSession
	IE         IEsSessionEstablishmentRequest
}

func (s *SessionEstablishmentRequest) MarshalBinary() (data []byte, err error) {
	// IE encode
	data, err = s.IE.Encode()
	if err != nil {
		return nil, err
	}
	return
}

func (s *SessionEstablishmentRequest) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	err := s.IE.Decode(data)
	if err != nil {
		return err
	}

	/*r := bytes.NewReader(data)
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

	err = pfcpIe.SetObject(t, l) // 当前IE的头设置
	if err != nil {
		return err
	}
	pfcpIe.Decode(ieData) // 当前IE的V解析
	if err != nil {
		return err
	}

	// 挂载当前对象
	err = root.SetObjectToParent(pfcpIe)
	if err != nil {
		return err
	}*/

	return nil
}

func (s *SessionEstablishmentRequest) Len() int {
	return int(s.PfcpHeader.Length)
}

// 7.5.3	PFCP Session Establishment Response
type SessionEstablishmentResponse struct {
	PfcpHeader PfcpHeaderforSession
	IE         IEsSessionEstablishmentResponse
}

func (s *SessionEstablishmentResponse) MarshalBinary() (data []byte, err error) {
	// IE encode
	data, err = s.IE.Encode()
	if err != nil {
		return nil, err
	}
	return
}

/*func encodeGroupedIe(ies PfcpIeInterface) (data []byte, err error) {
	// grouped IE encode

	//encode TLV
	data, err = ies.Encode()
	return
}*/

func (s *SessionEstablishmentResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	err := s.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionEstablishmentResponse) Len() int {
	return int(s.PfcpHeader.Length)
}

//7.5.4	PFCP Session Modification Request
type SessionModifyRequest struct {
	PfcpHeader PfcpHeaderforSession
	IE         IEsSessionModifyRequest
}

func (s *SessionModifyRequest) MarshalBinary() (data []byte, err error) {
	// IE encode
	data, err = s.IE.Encode()
	if err != nil {
		return nil, err
	}
	return
}

func (s *SessionModifyRequest) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	err := s.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionModifyRequest) Len() int {
	return int(s.PfcpHeader.Length)
}

//7.5.5	PFCP Session Modification Response
type SessionModifyResponse struct {
	PfcpHeader PfcpHeaderforSession
	IE         IEsSessionModifyResponse
}

func (s *SessionModifyResponse) MarshalBinary() (data []byte, err error) {
	// IE encode
	data, err = s.IE.Encode()
	if err != nil {
		return nil, err
	}
	return
}

func (s *SessionModifyResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	err := s.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *SessionModifyResponse) Len() int {
	return int(s.PfcpHeader.Length)
}

//Table 7.5.2.1-1: Information Elements in an PFCP Session Establishment Request
/*Information elements	P

Node ID	    M
CP F-SEID   M F-SEID
Create PDR	M
Create FAR	M
Create URR	C
Create QER	C
Create BAR	O
Create Traffic Endpoint	C
PDN Type	C
User Plane Inactivity Timer	O
User ID 	O
Trace Information O
*/
type IEsSessionEstablishmentRequest struct {
	NodeID                   IENodeID
	CPFSEID                  IEFSEID
	CreatePDRs               []*IECreatePDR // 可以有多个
	CreateFARs               []*IECreateFAR
	CreateURRs               []*IECreateURR              `json:",omitempty"`
	CreateQERs               []*IECreateQER              `json:",omitempty"`
	CreateBAR                *IECreateBAR                `json:",omitempty"`
	CreateTrafficEndpoints   []*IECreateTrafficEndpoint  `json:",omitempty"`
	PDNType                  *IEPDNType                  `json:",omitempty"`
	UserPlaneInactivityTimer *IEUserPlaneInactivityTimer `json:",omitempty"`
	UserID                   *IEUserID                   `json:",omitempty"`
	TraceInformation         *IETraceInformation         `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEsSessionEstablishmentRequest) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Node ID	    M
	// encode v
	vEnc, err := i.NodeID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Node_ID), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	//CP F-SEID   M F-SEID
	// encode v
	vEnc, err = i.CPFSEID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_F_SEID), Length: uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	//Create PDR	M
	for _, v := range i.CreatePDRs {
		// encode v
		vEnc, err = v.Encode()
		if err != nil {
			return nil, err
		}
		//encode tl
		tl = IETypeLength{Type: uint16(IE_Create_PDR), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}
	//Create FAR	M
	for _, v := range i.CreateFARs {
		// encode v
		vEnc, err = v.Encode()
		if err != nil {
			return nil, err
		}
		//encode tl
		tl = IETypeLength{Type: uint16(IE_Create_FAR), Length: uint16(len(vEnc))}
		tlvEnc, err = tl.EncodeTlV(vEnc)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(tlvEnc)
		if err != nil {
			return nil, err
		}
	}

	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Create URR	C
		case IE_Create_URR:
			for _, v := range i.CreateURRs[:len(i.CreateURRs)-1] {
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
			vEnc, err = i.CreateURRs[len(i.CreateURRs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Create QER	C
		case IE_Create_QER:
			for _, v := range i.CreateQERs[:len(i.CreateQERs)-1] {
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
			vEnc, err = i.CreateQERs[len(i.CreateQERs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Create BAR	O
		case IE_Create_BAR:
			vEnc, err = i.CreateBAR.Encode()
			if err != nil {
				return nil, err
			}
			//Create Traffic Endpoint	C
		case IE_Create_Traffic_Endpoint:
			for _, v := range i.CreateTrafficEndpoints[:len(i.CreateTrafficEndpoints)-1] {
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
			vEnc, err = i.CreateTrafficEndpoints[len(i.CreateTrafficEndpoints)-1].Encode()
			if err != nil {
				return nil, err
			}
			//PDN Type	C
		case IE_PDN_Type:
			vEnc, err = i.PDNType.Encode()
			if err != nil {
				return nil, err
			}
			//User Plane Inactivity Timer	O
		case IE_User_Plane_Inactivity_Timer:
			vEnc, err = i.UserPlaneInactivityTimer.Encode()
			if err != nil {
				return nil, err
			}
			//User ID 	O
		case IE_User_ID:
			vEnc, err = i.UserID.Encode()
			if err != nil {
				return nil, err
			}
			//Trace Information O
		case IE_Trace_Information:
			vEnc, err = i.TraceInformation.Encode()
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

func (i *IEsSessionEstablishmentRequest) Decode(data []byte) error {
	// parse IE
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEsSessionEstablishmentRequest) Len() int {
	return 0
}

func (i *IEsSessionEstablishmentRequest) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEsSessionEstablishmentRequest) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	//Node ID	    M
	case *IENodeID:
		i.NodeID = *Ie
		//CP F-SEID   M F-SEID
	case *IEFSEID:
		i.CPFSEID = *Ie
		//Create PDR	M
	case *IECreatePDR:
		i.CreatePDRs = append(i.CreatePDRs, Ie)
		//Create FAR	M
	case *IECreateFAR:
		i.CreateFARs = append(i.CreateFARs, Ie)
		//Create URR	C
	case *IECreateURR:
		i.CreateURRs = append(i.CreateURRs, Ie)
		//Create QER	C
	case *IECreateQER:
		i.CreateQERs = append(i.CreateQERs, Ie)
		//Create BAR	O
	case *IECreateBAR:
		i.CreateBAR = Ie
		//Create Traffic Endpoint	C
	case *IECreateTrafficEndpoint:
		i.CreateTrafficEndpoints = append(i.CreateTrafficEndpoints, Ie)
		//PDN Type	C
	case *IEPDNType:
		i.PDNType = Ie
		//User Plane Inactivity Timer	O
	case *IEUserPlaneInactivityTimer:
		i.UserPlaneInactivityTimer = Ie
		//User ID 	O
	case *IEUserID:
		i.UserID = Ie
		//Trace Information O
	case *IETraceInformation:
		i.TraceInformation = Ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

//Table 7.5.3.1-1: Information Elements in a PFCP Session Establishment Response
/*Information elements	P

Node ID	M
Cause	M
Offending IE	C
UP F-SEID	    C  F-SEID
Created PDR 	C
Load Control Information	    O
Overload Control Information	O
Failed Rule ID	C
Created Traffic Endpoint 	C
*/
type IEsSessionEstablishmentResponse struct {
	NodeID                 IENodeID
	Cause                  IECause
	OffendingIE            *IEOffendingIE            `json:",omitempty"`
	UPFSEID                *IEFSEID                  `json:",omitempty"`
	CreatedPDR             *IECreatedPDR             `json:",omitempty"`
	LoadControlInfo        *IELCI                    `json:",omitempty"`
	OverloadControlInfo    *IEOCI                    `json:",omitempty"`
	FailedRuleID           *IEFailedRuleID           `json:",omitempty"`
	CreatedTrafficEndpoint *IECreatedTrafficEndpoint `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEsSessionEstablishmentResponse) Encode() (data []byte, err error) {
	//	encode ie
	encBuf := bytes.NewBuffer(data[:0])

	// Mandatory ie
	//Node ID	M
	// encode v
	vEnc, err := i.NodeID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Node_ID), Length: uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	//Cause	M
	vEnc, err = i.Cause.Encode()
	if err != nil {
		return
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_Cause), Length: uint16(len(vEnc))}
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
		case IE_Offending_IE: //Offending IE	C
			vEnc, err = i.OffendingIE.Encode()
			if err != nil {
				return
			}
		case IE_F_SEID: //UP F-SEID	    C  F-SEID
			vEnc, err = i.UPFSEID.Encode()
			if err != nil {
				return
			}
		case IE_Created_PDR: //Created PDR 	C
			vEnc, err = i.CreatedPDR.Encode()
			if err != nil {
				return
			}

		case IE_Load_Control_Information: //Load Control Information	    O
			vEnc, err = i.LoadControlInfo.Encode()
			if err != nil {
				return
			}

		case IE_Overload_Control_Information: //Overload Control Information	O
			vEnc, err = i.OverloadControlInfo.Encode()
			if err != nil {
				return
			}

		case IE_Failed_Rule_ID: //Failed Rule ID	C
			vEnc, err = i.FailedRuleID.Encode()
			if err != nil {
				return
			}

		case IE_Created_Traffic_Endpoint: //Created Traffic Endpoint 	C
			vEnc, err = i.CreatedTrafficEndpoint.Encode()
			if err != nil {
				return
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

func (i *IEsSessionEstablishmentResponse) Decode(data []byte) error {
	// parse IE
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEsSessionEstablishmentResponse) Len() int {
	return 0
}

func (i *IEsSessionEstablishmentResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEsSessionEstablishmentResponse) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	//Node ID	M
	case *IENodeID:
		i.NodeID = *Ie
		//Cause	M
	case *IECause:
		i.Cause = *Ie
		//Offending IE	C
	case *IEOffendingIE:
		i.OffendingIE = Ie
		//UP F-SEID	    C  F-SEID
	case *IEFSEID:
		i.UPFSEID = Ie
		//Created PDR 	C
	case *IECreatedPDR:
		i.CreatedPDR = Ie
		//Load Control Information	    O
	case *IELCI:
		i.LoadControlInfo = Ie
		//Overload Control Information	O
	case *IEOCI:
		i.OverloadControlInfo = Ie
		//Failed Rule ID	C
	case *IEFailedRuleID:
		i.FailedRuleID = Ie
		//Created Traffic Endpoint 	C
	case *IECreatedTrafficEndpoint:
		i.CreatedTrafficEndpoint = Ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

//Table 7.5.4.1-1: Information Elements in a PFCP Session Modification Request
/*Information elements	P                     IE Type

CP F-SEID	C                                 F-SEID
Remove PDR	C                                 Remove PDR
Remove FAR	C                                 Remove FAR
Remove URR	C                                 Remove URR
Remove QER	C                                 Remove QER
Remove BAR	C                                 Remove BAR
Remove Traffic Endpoint	C                     Remove Traffic Endpoint
Create PDR	C                                 Create PDR
Create FAR	C                                 Create FAR
Create URR	C                                 Create URR
Create QER	C                                 Create QER
Create BAR	C                                 Create BAR
Create Traffic Endpoint	C                     Create Traffic Endpoint
Update PDR	C                                 Update PDR
Update FAR	C                                 Update FAR
Update URR	C                                 Update URR

Update QER	C                                 Update QER
Update BAR	C                                 Update BAR
Update Traffic Endpoint	C                     Update Traffic Endpoint
PFCPSMReq-Flags	C                             PFCPSMReq-Flags
Query URR	C                                 Query URR

User Plane Inactivity Timer	C                 User Plane Inactivity Timer
Query URR Reference	O                         Query URR Reference
Trace Information   O                         Trace Information

*/
type IEsSessionModifyRequest struct {
	CPFSEID               *IEFSEID                 `json:",omitempty"`
	RemovePDRs            []*IERemovePDR           `json:",omitempty"`
	RemoveFARs            []*IERemoveFAR           `json:",omitempty"`
	RemoveURRs            []*IERemoveURR           `json:",omitempty"`
	RemoveQERs            []*IERemoveQER           `json:",omitempty"`
	RemoveBAR             *IERemoveBAR             `json:",omitempty"`
	RemoveTrafficEndpoint *IERemoveTrafficEndpoint `json:",omitempty"`
	CreatePDRs            []*IECreatePDR           `json:",omitempty"`
	CreateFARs            []*IECreateFAR           `json:",omitempty"`
	CreateURRs            []*IECreateURR           `json:",omitempty"`
	CreateQERs            []*IECreateQER           `json:",omitempty"`
	CreateBAR             *IECreateBAR             `json:",omitempty"`
	CreateTrafficEndpoint *IECreateTrafficEndpoint `json:",omitempty"`
	UpdatePDRs            []*IEUpdatePDR           `json:",omitempty"`
	UpdateFARs            []*IEUpdateFAR           `json:",omitempty"`
	UpdateURRs            []*IEUpdateURR           `json:",omitempty"`

	UpdateQERs            []*IEUpdateQER           `json:",omitempty"`
	UpdateBAR             *IEUpdateBARForSMR       `json:",omitempty"`
	UpdateTrafficEndpoint *IEUpdateTrafficEndpoint `json:",omitempty"`
	PFCPSMReqFlags        *IEPFCPSMReqFlags        `json:",omitempty"`
	QueryURRs             []*IEQueryURR            `json:",omitempty"`

	UserPlaneInactivityTimer *IEUserPlaneInactivityTimer `json:",omitempty"`
	QueryURRReference        *IEQueryURRReference        `json:",omitempty"`
	TraceInformation         *IETraceInformation         `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEsSessionModifyRequest) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	var tlvEnc []byte
	var vEnc []byte
	var tl IETypeLength
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//CP F-SEID	C                                 F-SEID
		case IE_F_SEID:
			vEnc, err = i.CPFSEID.Encode()
			if err != nil {
				return nil, err
			}
			//Remove PDR	C                                 Remove PDR
		case IE_Remove_PDR:
			for _, v := range i.RemovePDRs[:len(i.RemovePDRs)-1] {
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
			vEnc, err = i.RemovePDRs[len(i.RemovePDRs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Remove FAR	C                                 Remove FAR
		case IE_Remove_FAR:
			for _, v := range i.RemoveFARs[:len(i.RemoveFARs)-1] {
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
			vEnc, err = i.RemoveFARs[len(i.RemoveFARs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Remove URR	C                                 Remove URR
		case IE_Remove_URR:
			for _, v := range i.RemoveURRs[:len(i.RemoveURRs)-1] {
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
			vEnc, err = i.RemoveURRs[len(i.RemoveURRs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Remove QER	C                                 Remove QER
		case IE_Remove_QER:
			for _, v := range i.RemoveQERs[:len(i.RemoveQERs)-1] {
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
			vEnc, err = i.RemoveQERs[len(i.RemoveQERs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Remove BAR	C                                 Remove BAR
		case IE_Remove_BAR:
			// encode v
			vEnc, err = i.RemoveBAR.Encode()
			if err != nil {
				return nil, err
			}
			//Remove Traffic Endpoint	C                     Remove Traffic Endpoint
		case IE_Remove_Traffic_Endpoint:
			// encode v
			vEnc, err = i.RemoveTrafficEndpoint.Encode()
			if err != nil {
				return nil, err
			}
			//Create PDR	C                                 Create PDR
		case IE_Create_PDR:
			for _, v := range i.CreatePDRs[:len(i.CreatePDRs)-1] {
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
			vEnc, err = i.CreatePDRs[len(i.CreatePDRs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Create FAR	C                                 Create FAR
		case IE_Create_FAR:
			for _, v := range i.CreateFARs[:len(i.CreateFARs)-1] {
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
			vEnc, err = i.CreateFARs[len(i.CreateFARs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Create URR	C                                 Create URR
		case IE_Create_URR:
			for _, v := range i.CreateURRs[:len(i.CreateURRs)-1] {
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
			vEnc, err = i.CreateURRs[len(i.CreateURRs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Create QER	C                                 Create QER
		case IE_Create_QER:
			for _, v := range i.CreateQERs[:len(i.CreateQERs)-1] {
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
			vEnc, err = i.CreateQERs[len(i.CreateQERs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Create BAR	C                                 Create BAR
		case IE_Create_BAR:
			// encode v
			vEnc, err = i.CreateBAR.Encode()
			if err != nil {
				return nil, err
			}
			//Create Traffic Endpoint	C                     Create Traffic Endpoint
		case IE_Create_Traffic_Endpoint:
			// encode v
			vEnc, err = i.CreateTrafficEndpoint.Encode()
			if err != nil {
				return nil, err
			}
			//Update PDR	C                                 Update PDR
		case IE_Update_PDR:
			for _, v := range i.UpdatePDRs[:len(i.UpdatePDRs)-1] {
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
			vEnc, err = i.UpdatePDRs[len(i.UpdatePDRs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Update FAR	C                                 Update FAR
		case IE_Update_FAR:
			for _, v := range i.UpdateFARs[:len(i.UpdateFARs)-1] {
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
			vEnc, err = i.UpdateFARs[len(i.UpdateFARs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Update URR	C                                 Update URR
		case IE_Update_URR:
			for _, v := range i.UpdateURRs[:len(i.UpdateURRs)-1] {
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
			vEnc, err = i.UpdateURRs[len(i.UpdateURRs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//
			//Update QER	C                                 Update QER
		case IE_Update_QER:
			for _, v := range i.UpdateQERs[:len(i.UpdateQERs)-1] {
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
			vEnc, err = i.UpdateQERs[len(i.UpdateQERs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//Update BAR	C                                 Update BAR
		case IE_Update_BAR:
			// encode v
			vEnc, err = i.UpdateBAR.Encode()
			if err != nil {
				return nil, err
			}
			//Update Traffic Endpoint	C                     Update Traffic Endpoint
		case IE_Update_Traffic_Endpoint:
			// encode v
			vEnc, err = i.UpdateTrafficEndpoint.Encode()
			if err != nil {
				return nil, err
			}
			//PFCPSMReq-Flags	C                             PFCPSMReq-Flags
		case IE_PFCPSMReq_Flags:
			// encode v
			vEnc, err = i.PFCPSMReqFlags.Encode()
			if err != nil {
				return nil, err
			}
			//Query URR	C                                 Query URR
		case IE_Query_URR:
			for _, v := range i.QueryURRs[:len(i.QueryURRs)-1] {
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
			vEnc, err = i.QueryURRs[len(i.QueryURRs)-1].Encode()
			if err != nil {
				return nil, err
			}
			//
			//User Plane Inactivity Timer	C                 User Plane Inactivity Timer
		case IE_User_Plane_Inactivity_Timer:
			// encode v
			vEnc, err = i.UserPlaneInactivityTimer.Encode()
			if err != nil {
				return nil, err
			}
			//Query URR Reference	O                         Query URR Reference
		case IE_Query_URR_Reference:
			// encode v
			vEnc, err = i.QueryURRReference.Encode()
			if err != nil {
				return nil, err
			}
			//Trace Information   O                         Trace Information
		case IE_Trace_Information:
			// encode v
			vEnc, err = i.TraceInformation.Encode()
			if err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("Illegal IE")
		}
		// TL 编码
		tl := IETypeLength{Type: uint16(id), Length: uint16(len(vEnc))}
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

func (i *IEsSessionModifyRequest) Decode(data []byte) error {
	// parse IE
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEsSessionModifyRequest) Len() int {
	return 0
}

func (i *IEsSessionModifyRequest) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEsSessionModifyRequest) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	//CP F-SEID	C                                 F-SEID
	case *IEFSEID:
		i.CPFSEID = Ie
		//Remove PDR	C                                 Remove PDR
	case *IERemovePDR:
		i.RemovePDRs = append(i.RemovePDRs, Ie)
		//Remove FAR	C                                 Remove FAR
	case *IERemoveFAR:
		i.RemoveFARs = append(i.RemoveFARs, Ie)
		//Remove URR	C                                 Remove URR
	case *IERemoveURR:
		i.RemoveURRs = append(i.RemoveURRs, Ie)
		//Remove QER	C                                 Remove QER
	case *IERemoveQER:
		i.RemoveQERs = append(i.RemoveQERs, Ie)
		//Remove BAR	C                                 Remove BAR
	case *IERemoveBAR:
		i.RemoveBAR = Ie
		//Remove Traffic Endpoint	C                     Remove Traffic Endpoint
	case *IERemoveTrafficEndpoint:
		i.RemoveTrafficEndpoint = Ie
		//Create PDR	C                                 Create PDR
	case *IECreatePDR:
		i.CreatePDRs = append(i.CreatePDRs, Ie)
		//Create FAR	C                                 Create FAR
	case *IECreateFAR:
		i.CreateFARs = append(i.CreateFARs, Ie)
		//Create URR	C                                 Create URR
	case *IECreateURR:
		i.CreateURRs = append(i.CreateURRs, Ie)
		//Create QER	C                                 Create QER
	case *IECreateQER:
		i.CreateQERs = append(i.CreateQERs, Ie)
		//Create BAR	C                                 Create BAR
	case *IECreateBAR:
		i.CreateBAR = Ie
		//Create Traffic Endpoint	C                     Create Traffic Endpoint
	case *IECreateTrafficEndpoint:
		i.CreateTrafficEndpoint = Ie
		//Update PDR	C                                 Update PDR
	case *IEUpdatePDR:
		i.UpdatePDRs = append(i.UpdatePDRs, Ie)
		//Update FAR	C                                 Update FAR
	case *IEUpdateFAR:
		i.UpdateFARs = append(i.UpdateFARs, Ie)
		//Update URR	C                                 Update URR
	case *IEUpdateURR:
		i.UpdateURRs = append(i.UpdateURRs, Ie)
		//
		//Update QER	C                                 Update QER
	case *IEUpdateQER:
		i.UpdateQERs = append(i.UpdateQERs, Ie)
		//Update BAR	C                                 Update BAR
	case *IEUpdateBARForSMR:
		i.UpdateBAR = Ie
		//Update Traffic Endpoint	C                     Update Traffic Endpoint
	case *IEUpdateTrafficEndpoint:
		i.UpdateTrafficEndpoint = Ie
		//PFCPSMReq-Flags	C                             PFCPSMReq-Flags
	case *IEPFCPSMReqFlags:
		i.PFCPSMReqFlags = Ie
		//Query URR	C                                 Query URR
	case *IEQueryURR:
		i.QueryURRs = append(i.QueryURRs, Ie)
		//
		//User Plane Inactivity Timer	C                 User Plane Inactivity Timer
	case *IEUserPlaneInactivityTimer:
		i.UserPlaneInactivityTimer = Ie
		//Query URR Reference	O                         Query URR Reference
	case *IEQueryURRReference:
		i.QueryURRReference = Ie
		//Trace Information   O                         Trace Information
	case *IETraceInformation:
		i.TraceInformation = Ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

//7.5.5	PFCP Session Modification Response
/*Table 7.5.5.1-1: Information Elements in a PFCP Session Modification Response
Information elements	              P      IE Type

Cause	        M                            Cause
Offending IE	C                            Offending IE
Created PDR	    C                            Created PDR
Load Control Information	          O      Load Control Information
Overload Control Information	      O      Overload Control Information
Usage Report	C                            Usage Report
Failed Rule ID	C                            Failed Rule ID
Additional Usage Reports Information	C    Additional Usage Reports Information
Created/Updated Traffic Endpoint	    C    Created Traffic Endpoint
*/
type IEsSessionModifyResponse struct {
	Cause                      IECause
	OffendingIE                *IEOffendingIE                       `json:",omitempty"`
	CreatedPDR                 *IECreatedPDR                        `json:",omitempty"`
	LoadControlInformation     *IELCI                               `json:",omitempty"`
	OverloadControlInformation *IEOCI                               `json:",omitempty"`
	UsageReport                *IEUsageReportModResp                `json:",omitempty"`
	FailedRuleID               *IEFailedRuleID                      `json:",omitempty"`
	AddUsageReportsInfo        *IEAdditionalUsageReportsInformation `json:",omitempty"`
	CreatedTrafficEndpoint     *IECreatedTrafficEndpoint            `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEsSessionModifyResponse) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Cause	        M                            Cause
	// encode v
	vEnc, err := i.Cause.Encode()
	if err != nil {
		return nil, err
	}
	// encode tl
	tl := IETypeLength{uint16(IE_Cause), uint16(len(vEnc))}
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
		//Offending IE	C                            Offending IE
		case IE_Offending_IE:
			vEnc, err = i.OffendingIE.Encode()
			if err != nil {
				return nil, err
			}
			//Created PDR	    C                            Created PDR
		case IE_Created_PDR:
			vEnc, err = i.CreatedPDR.Encode()
			if err != nil {
				return nil, err
			}
			//Load Control Information	          O      Load Control Information
		case IE_Load_Control_Information:
			vEnc, err = i.LoadControlInformation.Encode()
			if err != nil {
				return nil, err
			}
			//Overload Control Information	      O      Overload Control Information
		case IE_Overload_Control_Information:
			vEnc, err = i.OverloadControlInformation.Encode()
			if err != nil {
				return nil, err
			}

			//Usage Report	C                            Usage Report
		case IE_Usage_Report_Modification_Response:
			// encode v
			vEnc, err = i.UsageReport.Encode()
			if err != nil {
				return nil, err
			}

			//Failed Rule ID	C                            Failed Rule ID
		case IE_Failed_Rule_ID:
			// encode v
			vEnc, err = i.FailedRuleID.Encode()
			if err != nil {
				return nil, err
			}
			//Additional Usage Reports Information	C    Additional Usage Reports Information
		case IE_Additional_Usage_Reports_Information:
			// encode v
			vEnc, err = i.AddUsageReportsInfo.Encode()
			if err != nil {
				return nil, err
			}
			//Created/Updated Traffic Endpoint	    C    Created Traffic Endpoint
		case IE_Created_Traffic_Endpoint:
			// encode v
			vEnc, err = i.CreatedTrafficEndpoint.Encode()
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

func (i *IEsSessionModifyResponse) Decode(data []byte) error {
	// parse IE
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEsSessionModifyResponse) Len() int {
	return 0
}

func (i *IEsSessionModifyResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEsSessionModifyResponse) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	//Cause	        M                            Cause
	case *IECause:
		i.Cause = *Ie
		//Offending IE	C
	case *IEOffendingIE:
		i.OffendingIE = Ie
		//Created PDR	    C                            Created PDR
	case *IECreatedPDR:
		i.CreatedPDR = Ie
		//Load Control Information	          O      Load Control Information
	case *IELCI:
		i.LoadControlInformation = Ie
		//Overload Control Information	      O      Overload Control Information
	case *IEOCI:
		i.OverloadControlInformation = Ie
		//Usage Report	C                            Usage Report
	case *IEUsageReportModResp:
		i.UsageReport = Ie
		//Failed Rule ID	C                            Failed Rule ID
	case *IEFailedRuleID:
		i.FailedRuleID = Ie
		//Additional Usage Reports Information	C    Additional Usage Reports Information
	case *IEAdditionalUsageReportsInformation:
		i.AddUsageReportsInfo = Ie
		//Created/Updated Traffic Endpoint	    C    Created Traffic Endpoint
	case *IECreatedTrafficEndpoint:
		i.CreatedTrafficEndpoint = Ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

// 7.5.8	PFCP Session Report Request
/*Table 7.5.8-1: Information Elements in a PFCP Session Report Request
Information elements	              P      IE Type

Report Type	                          M      Report Type
Downlink Data Report	              C      Downlink Data Report
Usage Report	                      C      Usage Report
Error Indication Report	              C      Error Indication Report
Load Control Information	          O      Load Control Information
Overload Control Information	      O      Overload Control Information
Additional Usage Reports Information  C	    Additional Usage Reports Information
*/
type IEsSessionReportRequest struct {
	ReportType                        IEReportType
	DownlinkDataReport                *IEDownlinkDataReport
	UsageReport                       *IEUsageReportWithinRepReq
	ErrorIndicationReport             *IEErrorIndicationReport
	LoadControlInformation            *IELCI
	OverloadControlInformation        *IEOCI
	AdditionalUsageReportsInformation IEAdditionalUsageReportsInformation

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEsSessionReportRequest) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Report Type	                          M      Report Type
	// encode v
	vEnc, err := i.ReportType.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Report_Type), Length: uint16(len(vEnc))}
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
		//Downlink Data Report	              C      Downlink Data Report
		case IE_Downlink_Data_Report:
			//	encode v
			vEnc, err = i.DownlinkDataReport.Encode()
			if err != nil {
				return nil, err
			}
			//Usage Report	                      C      Usage Report
		case IE_Usage_Report_Request:
			//	encode v
			vEnc, err = i.UsageReport.Encode()
			if err != nil {
				return nil, err
			}
			//Error Indication Report	              C      Error Indication Report
		case IE_Error_Indication_Report:
			//	encode v
			vEnc, err = i.ErrorIndicationReport.Encode()
			if err != nil {
				return nil, err
			}
			//Load Control Information	          O      Load Control Information
		case IE_Load_Control_Information:
			//	encode v
			vEnc, err = i.LoadControlInformation.Encode()
			if err != nil {
				return nil, err
			}
			//Overload Control Information	      O      Overload Control Information
		case IE_Overload_Control_Information:
			//	encode v
			vEnc, err = i.OverloadControlInformation.Encode()
			if err != nil {
				return nil, err
			}
			//Additional Usage Reports Information  C	    Additional Usage Reports Information
		case IE_Additional_Usage_Reports_Information:
			//	encode v
			vEnc, err = i.AdditionalUsageReportsInformation.Encode()
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

func (i *IEsSessionReportRequest) Decode(data []byte) error {
	// parse IE
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEsSessionReportRequest) Len() int {
	return 0
}

func (i *IEsSessionReportRequest) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEsSessionReportRequest) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	//Report Type	                          M      Report Type
	case *IEReportType:
		i.ReportType = *Ie
		//Downlink Data Report	              C      Downlink Data Report
	case *IEDownlinkDataReport:
		i.DownlinkDataReport = Ie
		//Usage Report	                      C      Usage Report
	case *IEUsageReportWithinRepReq:
		i.UsageReport = Ie
		//Error Indication Report	              C      Error Indication Report
	case *IEErrorIndicationReport:
		i.ErrorIndicationReport = Ie
		//Load Control Information	          O      Load Control Information
	case *IELCI:
		i.LoadControlInformation = Ie
		//Overload Control Information	      O      Overload Control Information
	case *IEOCI:
		i.OverloadControlInformation = Ie
		//Additional Usage Reports Information  C	    Additional Usage Reports Information
	case *IEAdditionalUsageReportsInformation:
		i.AdditionalUsageReportsInformation = *Ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

// 7.5.9	PFCP Session Report Response
/*Table 7.5.9.1-1: Information Elements in a PFCP Session Report Response
Information elements	              P      IE Type

Cause	                              M      Cause
Offending IE	                      C      Offending IE
Update BAR	                          C      Update BAR
PFCPSRRsp-Flags	                      C      PFCPSRRsp-Flags
*/
type IEsSessionReportResponse struct {
	Cause          IECause
	OffendingIE    IEOffendingIE
	UpdateBAR      IEUpdateBAR
	PFCPSRRspFlags IEPFCPSRRspFlags

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEsSessionReportResponse) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Cause	          M
	// encode v
	vEnc, err := i.Cause.Encode()
	if err != nil {
		return nil, err
	}
	// encode tl
	tl := IETypeLength{uint16(IE_Cause), uint16(len(vEnc))}
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
		//Offending IE	                      C      Offending IE
		case IE_Offending_IE:
			//	encode v
			vEnc, err = i.OffendingIE.Encode()
			if err != nil {
				return nil, err
			}
			//Update BAR	                          C      Update BAR
		case IE_Update_BAR:
			//	encode v
			vEnc, err = i.UpdateBAR.Encode()
			if err != nil {
				return nil, err
			}
			//PFCPSRRsp-Flags	                      C      PFCPSRRsp-Flags
		case IE_PFCPSRRsp_Flags:
			//	encode v
			vEnc, err = i.PFCPSRRspFlags.Encode()
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

func (i *IEsSessionReportResponse) Decode(data []byte) error {
	// parse IE
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEsSessionReportResponse) Len() int {
	return 0
}

func (i *IEsSessionReportResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEsSessionReportResponse) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	//Cause	                              M      Cause
	case *IECause:
		i.Cause = *Ie
		//Offending IE	                      C      Offending IE
	case *IEOffendingIE:
		i.OffendingIE = *Ie
		//Update BAR	                          C      Update BAR
	case *IEUpdateBAR:
		i.UpdateBAR = *Ie
		//PFCPSRRsp-Flags	                      C      PFCPSRRsp-Flags
	case *IEPFCPSRRspFlags:
		i.PFCPSRRspFlags = *Ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func CreateObject(i interface{}) interface{} {

	t := reflect.ValueOf(i).Type()
	v := reflect.New(t).Interface() // 返回的是对象的接口类型

	return v
}

func EncodeTl(encBuf io.Writer, tl IETypeLength) error {
	// encode Tl
	//encBuf := bytes.NewBuffer(nil)
	err := binary.Write(encBuf, binary.BigEndian, tl.Type)
	if err != nil {
		return err
	}
	err = binary.Write(encBuf, binary.BigEndian, tl.Length)
	if err != nil {
		return err
	}

	return nil
}

// 1146
