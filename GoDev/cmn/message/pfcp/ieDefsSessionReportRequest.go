package pfcp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/message/pfcp/utils"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)
// N4 消息

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

// 8.2.21	Report Type
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 39 (decimal)
	3 to 4	Length = n
	5	Spare	UPIR	ERIR	USAR	DLDR
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.21-1: Report Type*/
type IEReportType struct {
	IETypeLength
	TypeValue              uint8 // todo delete
	UPIR, ERIR, USAR, DLDR bool
}

func (i *IEReportType) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := utils.BoolToUint8(i.DLDR) +
		(utils.BoolToUint8(i.USAR) << 1) +
		(utils.BoolToUint8(i.ERIR) << 2) +
		(utils.BoolToUint8(i.UPIR) << 3)
	err = binary.Write(encBuf, binary.BigEndian, tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEReportType) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.DLDR, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}

	i.USAR, err = utils.GetBoolInByte(tmp, 2)
	if err != nil {
		return err
	}

	i.ERIR, err = utils.GetBoolInByte(tmp, 3)
	if err != nil {
		return err
	}
	i.UPIR, err = utils.GetBoolInByte(tmp, 4)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEReportType) Len() int {
	return int(i.Length)
}

func (i *IEReportType) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEReportType) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Octet 5 shall be encoded as follows:
-	Bit 1 – DLDR (Downlink Data Report): when set to 1, this indicates Downlink Data Report
-	Bit 2 – USAR (Usage Report): when set to 1, this indicates a Usage Report
-	Bit 3 – ERIR (Error Indication Report): when set to 1, this indicates an Error Indication Report.
-	Bit 4 – UPIR (User Plane Inactivity Report): when set to 1, this indicates a User Plane Inactivity Report.
-	Bit 5 to 8 – Spare, for future use and set to 0.
At least one bit shall be set to 1. Several bits may be set to 1.*/
const (
	DLDR = utils.Bit_0 //2 ^ 0
	USAR = utils.Bit_1 //2 ^ 1
	ERIR = utils.Bit_2 //2 ^ 2
	UPIR = utils.Bit_3 //2 ^ 3
)

func (i *IEReportType) Set(v uint8) {
	i.Type = IE_Report_Type
	i.Length = 1
	i.TypeValue = v
}

//7.5.8.2	Downlink Data Report IE within PFCP Session Report Request
/*Table 7.5.8.2-1: Downlink Data Report IE within PFCP Session Report Request
Octet 1 and 2		Downlink Data Report IE Type = 83 (decimal)
Octets 3 and 4		Length = n
Information elements	              P      IE Type

PDR ID	                              M      PDR ID
Downlink Data Service Information	  C      Downlink Data Service Information
*/
type IEDownlinkDataReport struct {
	IETypeLength
	PDRID             IEPDRID
	DLDataServiceInfo *IEDLDataServiceInfo

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEDownlinkDataReport) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//PDR ID	   M
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

	//  optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Downlink Data Service Information	  C
		case IE_Downlink_Data_Service_Information:
			// encode v
			vEnc, err = i.DLDataServiceInfo.Encode()
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

func (i *IEDownlinkDataReport) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEDownlinkDataReport) Len() int {
	return int(i.Length)
}

func (i *IEDownlinkDataReport) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEDownlinkDataReport) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//PDR ID	                          M
	case *IEPDRID:
		i.PDRID = *ie
		//Downlink Data Service Information	  C
	case *IEDLDataServiceInfo:
		i.DLDataServiceInfo = ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEDownlinkDataReport) Set() {
	i.Type = IE_Downlink_Data_Report
}

//8.2.27	Downlink Data Service Information
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 45 (decimal)
	3 to 4	Length = n
	5	Spare	                   QFII	PPI
	m	Spare	   Paging Policy Indication value
	p	Spare	   QFI
	q to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.27-1: Downlink Data Service Information*/
type IEDLDataServiceInfo struct {
	IETypeLength
	Flag      uint8 // todo delete
	QFII, PPI bool
	PPIValue  uint8 //Paging Policy Indication value
	QFI       uint8
}

func (i *IEDLDataServiceInfo) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	                   QFII	PPI
	tmpByte := utils.BoolToUint8(i.PPI) +
		(utils.BoolToUint8(i.QFII) << 1)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}
	//m	Spare	   Paging Policy Indication value
	if i.PPI {
		err = encBuf.WriteByte(i.PPIValue)
		if err != nil {
			return nil, err
		}
	}

	//p	Spare	   QFI
	if i.QFII {
		err = encBuf.WriteByte(i.QFI)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEDLDataServiceInfo) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	//5	Spare	                   QFII	PPI
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.PPI, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}
	i.QFII, err = utils.GetBoolInByte(tmp, 2)
	if err != nil {
		return err
	}
	//m	Spare	   Paging Policy Indication value
	if i.PPI {
		tmp, err := r.ReadByte()
		if err != nil {
			return err
		}
		v := utils.ByteOne(tmp)
		i.PPIValue, err = v.GetBits(1, 6)
		if err != nil {
			return err
		}
	}

	//p	Spare	   QFI
	if i.QFII {
		tmp, err := r.ReadByte()
		if err != nil {
			return err
		}
		v := utils.ByteOne(tmp)
		i.QFI, err = v.GetBits(1, 6)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *IEDLDataServiceInfo) Len() int {
	return int(i.Length)
}

func (i *IEDLDataServiceInfo) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEDLDataServiceInfo) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*The PPI flag in octet 5 indicates whether the Paging Policy Indication value in octet 'm' shall be present. If PPI is set to '1',
then the Paging Policy Indication value shall be present. If PPI is set to '0', then octet 'm' shall not be present.
The Paging Policy Indication value, in octet 'm', shall be encoded as the DSCP in TOS (IPv4) or TC (IPv6) information
received in the IP payload of the GTP-U packet from the PGW (see IETF RFC 2474 [13]).
The QFII flag in octet 5 indicates whether the QFI value in octet 'p' shall be present. If QFII is set to '1',
then the QFI value shall be present. If QFII is set to '0', then octet 'p' shall not be present.
The QFI value, in octet 'p', shall be encoded as the octet 5 of the QFI IE in subclause 8.2.89.*/
const (
	PPI_Flag  = utils.Bit_0 //2 ^ 0
	QFII_Flag = utils.Bit_1 //2 ^ 1
)

func (i *IEDLDataServiceInfo) Set() {
	i.Type = IE_Downlink_Data_Service_Information

}

// 7.5.8.4	Error Indication Report IE within PFCP Session Report Request
/*Information elements	              P      IE Type
Remote F-TEID	                      M      F-TEID
*/
type IEErrorIndicationReport struct {
	IETypeLength
	RemoteFTEID IEFTEID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEErrorIndicationReport) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Remote F-TEID	                      M      F-TEID
	// encode v
	vEnc, err := i.RemoteFTEID.Encode()
	if err != nil {
		return nil, err
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

	return encBuf.Bytes(), nil
}

func (i *IEErrorIndicationReport) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEErrorIndicationReport) Len() int {
	return int(i.Length)
}

func (i *IEErrorIndicationReport) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEErrorIndicationReport) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Remote F-TEID	                      M      F-TEID
	case *IEFTEID:
		i.RemoteFTEID = *ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEErrorIndicationReport) Set() {
	i.Type = IE_Error_Indication_Report

}

//7.5.8.3	Usage Report IE within PFCP Session Report Request
//Usage Report	C                            Usage Report
// IEUsageReportWithinRepReq
/*Table 7.5.8.3-1: Usage Report IE within PFCP Session Report Request
Octet 1 and 2		Usage Report IE Type = 80 (decimal)
Octets 3 and 4		Length = n
Information elements	P       IE Type

URR ID	M
UR-SEQN	M
Usage Report Trigger	M
Start Time	C
End Time	C
Volume Measurement	C
Duration Measurement	C
Application Detection Information	C
UE IP address	C
Network Instance	C
Time of First Packet	C
Time of Last Packet	C
Usage Information 	C
Query URR Reference	C
Event Time Stamp 	C
Ethernet Traffic Information	C
*/
type IEUsageReportWithinRepReq struct {
	IETypeLength
	URRID                      IEURRID
	URSEQN                     IEURSEQN
	UsageReportTrigger         IEUsageReportTrigger
	StartTime                  *IEStartTime
	EndTime                    *IEEndTime
	VolumeMeasurement          *IEVolumeMeasurement
	DurationMeasurement        *IEDurationMeasurement
	AppDetectionInformation    *IEAppDetectionInformation
	UEIPaddress                *IEUEIPaddress
	NetworkInstance            *IENetworkInstance
	TimeofFirstPacket          *IETimeofFirstPacket
	TimeofLastPacket           *IETimeofLastPacket
	UsageInformation           *IEUsageInformation
	QueryURRReference          *IEQueryURRReference
	EventTimeStamp             *IEEventTimeStamp
	EthernetTrafficInformation *IEEthernetTrafficInformation

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUsageReportWithinRepReq) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//URR ID	M                       URR ID
	//encode v
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

	//UR-SEQN	M                       UR-SEQN
	//encode v
	vEnc, err = i.URSEQN.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_UR_SEQN), Length: uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	//Usage Report Trigger	M       Usage Report Trigger
	//encode v
	vEnc, err = i.UsageReportTrigger.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{Type: uint16(IE_Usage_Report_Trigger), Length: uint16(len(vEnc))}
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
		//Start Time	C                   Start Time
		case IE_Start_Time:
			// encode v
			vEnc, err = i.StartTime.Encode()
			if err != nil {
				return nil, err
			}
			//End Time	C                   End Time
		case IE_End_Time:
			// encode v
			vEnc, err = i.EndTime.Encode()
			if err != nil {
				return nil, err
			}

			//Volume Measurement	    C       Volume Measurement
		case IE_Volume_Measurement:
			// encode v
			vEnc, err = i.VolumeMeasurement.Encode()
			if err != nil {
				return nil, err
			}
			//Duration Measurement	C       Duration Measurement
		case IE_Duration_Measurement:
			// encode v
			vEnc, err = i.DurationMeasurement.Encode()
			if err != nil {
				return nil, err
			}
			//Application Detection Information	C
		case IE_Application_Detection_Information:
			// encode v
			vEnc, err = i.AppDetectionInformation.Encode()
			if err != nil {
				return nil, err
			}
			//UE IP address	C
		case IE_UE_IP_Address:
			// encode v
			vEnc, err = i.UEIPaddress.Encode()
			if err != nil {
				return nil, err
			}
			//Network Instance	C
		case IE_Network_Instance:
			// encode v
			vEnc, err = i.NetworkInstance.Encode()
			if err != nil {
				return nil, err
			}
			//Time of First Packet	C
		case IE_Time_of_First_Packet:
			// encode v
			vEnc, err = i.TimeofFirstPacket.Encode()
			if err != nil {
				return nil, err
			}
			//Time of Last Packet	C
		case IE_Time_of_Last_Packet:
			// encode v
			vEnc, err = i.TimeofLastPacket.Encode()
			if err != nil {
				return nil, err
			}

			//Usage Information 	C
		case IE_Usage_Information:
			// encode v
			vEnc, err = i.UsageInformation.Encode()
			if err != nil {
				return nil, err
			}

			//Query URR Reference	C           Query URR Reference
		case IE_Query_URR_Reference:
			// encode v
			vEnc, err = i.QueryURRReference.Encode()
			if err != nil {
				return nil, err
			}
			//Event Time Stamp 	C
		case IE_Event_Time_Stamp:
			vEnc, err = i.EventTimeStamp.Encode()
			if err != nil {
				return nil, err
			}
			//Ethernet Traffic Information	C
		case IE_Ethernet_Traffic_Information:
			// encode v
			vEnc, err = i.EthernetTrafficInformation.Encode()
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

func (i *IEUsageReportWithinRepReq) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUsageReportWithinRepReq) Len() int {
	return int(i.Length)
}

func (i *IEUsageReportWithinRepReq) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUsageReportWithinRepReq) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//URR ID	M
	case *IEURRID:
		i.URRID = *ie
		//UR-SEQN	M
	case *IEURSEQN:
		i.URSEQN = *ie
		//Usage Report Trigger	M
	case *IEUsageReportTrigger:
		i.UsageReportTrigger = *ie
		//Start Time	C
	case *IEStartTime:
		i.StartTime = ie
		//End Time	C
	case *IEEndTime:
		i.EndTime = ie
		//Volume Measurement	C
	case *IEVolumeMeasurement:
		i.VolumeMeasurement = ie
		//Duration Measurement	C
	case *IEDurationMeasurement:
		i.DurationMeasurement = ie
		//Application Detection Information	C
	case *IEAppDetectionInformation:
		i.AppDetectionInformation = ie
		//UE IP address	C
	case *IEUEIPaddress:
		i.UEIPaddress = ie
		//Network Instance	C
	case *IENetworkInstance:
		i.NetworkInstance = ie
		//Time of First Packet	C
	case *IETimeofFirstPacket:
		i.TimeofFirstPacket = ie
		//Time of Last Packet	C
	case *IETimeofLastPacket:
		i.TimeofLastPacket = ie
		//Usage Information 	C
	case *IEUsageInformation:
		i.UsageInformation = ie
		//Query URR Reference	C
	case *IEQueryURRReference:
		i.QueryURRReference = ie
		//Event Time Stamp 	C
	case *IEEventTimeStamp:
		i.EventTimeStamp = ie
		//Ethernet Traffic Information	C
	case *IEEthernetTrafficInformation:
		i.EthernetTrafficInformation = ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEUsageReportWithinRepReq) Set(v uint8) error {
	i.Type = IE_Usage_Report_Request

	return nil
}

//Table 7.5.8.3-2: Application Detection Information IE within Usage Report IE
/*Octet 1 and 2		Application Detection Information IE Type = 68 (decimal)
Octets 3 and 4		Length = n
Information elements	P

Application ID	M
Application Instance ID	C
Flow Information	C
*/
type IEAppDetectionInformation struct {
	IETypeLength
	ApplicationID         IEApplicationID
	ApplicationInstanceID *IEApplicationInstanceID
	FlowInformation       *IEFlowInformation

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEAppDetectionInformation) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Application ID	M
	// encode v
	vEnc, err := i.ApplicationID.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Application_ID), Length: uint16(len(vEnc))}
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
		//Application Instance ID	C
		case IE_Application_Instance_ID:
			//	encode v
			vEnc, err = i.ApplicationInstanceID.Encode()
			if err != nil {
				return nil, err
			}
			//Flow Information	C
		case IE_Flow_Information:
			//	encode v
			vEnc, err = i.FlowInformation.Encode()
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

func (i *IEAppDetectionInformation) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEAppDetectionInformation) Len() int {
	return int(i.Length)
}

func (i *IEAppDetectionInformation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEAppDetectionInformation) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Application ID	M
	case *IEApplicationID:
		i.ApplicationID = *ie
		//Application Instance ID	C
	case *IEApplicationInstanceID:
		i.ApplicationInstanceID = ie
		//Flow Information	C
	case *IEFlowInformation:
		i.FlowInformation = ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

// IEApplicationInstanceID
// 8.2.60	Application Instance ID
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 91 (decimal)
	3 to 4	Length = n
	5 to (n+4)	Application Instance Identifier
Figure 8.2.60-1: Application Instance ID*/
type IEApplicationInstanceID struct {
	IETypeLength
	AppInstanceId string
}

func (i *IEApplicationInstanceID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	_, err = encBuf.Write([]byte(i.AppInstanceId))
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEApplicationInstanceID) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmpBytes := make([]byte, i.Length)
	_, err := r.Read(tmpBytes)
	if err != nil {
		return err
	}
	i.AppInstanceId = string(tmpBytes)

	return nil
}

func (i *IEApplicationInstanceID) Len() int {
	return int(i.Length)
}

func (i *IEApplicationInstanceID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEApplicationInstanceID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

// 8.2.61	Flow Information
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 92 (decimal)
	3 to 4	Length = n
	5	Spare	                Flow Direction
	6 to 7	Length of Flow Description
	8 to m	Flow Description
	p to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.61-1: Flow Information
*/
type IEFlowInformation struct {
	IETypeLength
	FlowDirection           uint8
	LengthofFlowDescription uint16
	FlowDescription         string
}

func (i *IEFlowInformation) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	                Flow Direction
	err = encBuf.WriteByte(i.FlowDirection)
	if err != nil {
		return nil, err
	}
	//6 to 7	Length of Flow Description
	i.LengthofFlowDescription = uint16(len(i.FlowDescription))
	err = binary.Write(encBuf, binary.BigEndian, i.LengthofFlowDescription)
	if err != nil {
		return
	}
	//8 to m	Flow Description
	_, err = encBuf.Write([]byte(i.FlowDescription))
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEFlowInformation) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	//5	Spare	                Flow Direction
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	i.FlowDirection, err = v.GetBits(1, 3)
	if err != nil {
		return err
	}
	//6 to 7	Length of Flow Description
	err = binary.Read(r, binary.BigEndian, &i.LengthofFlowDescription)
	if err != nil {
		return err
	}
	//8 to m	Flow Description
	tmpBytes := make([]byte, i.LengthofFlowDescription)
	err = binary.Read(r, binary.BigEndian, &tmpBytes)
	if err != nil {
		return err
	}
	i.FlowDescription = string(tmpBytes)
	return nil
}

func (i *IEFlowInformation) Len() int {
	return int(i.Length)
}

func (i *IEFlowInformation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEFlowInformation) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

/*Table 8.2.61-1: Flow Direction
Flow Direction 	Value (Decimal)
Unspecified	0
Downlink (traffic to the UE)	1
Uplink (traffic from the UE)	2
Bidirectional	3
For future use. Shall not be sent. If received, shall be interpreted as the value "0".	4 to 7
*/
const (
	Unspecified   = 0
	Downlink      = 1
	Uplink        = 2
	Bidirectional = 3
)

// 121
