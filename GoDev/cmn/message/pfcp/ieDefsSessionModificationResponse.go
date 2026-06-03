package pfcp

import (
	"bytes"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/message/pfcp/utils"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)
// N4 消息

// IE 名称 --来源于消息
// IE type --来源于8.1IE列表
// IE 格式 --来源于8.2格式定义

//Usage Report	C                            Usage Report
// IEUsageReportModResp
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
78	Usage Report (in Session Modification Response)	Extendable / Table 7.5.5.2-1	Not Applicable
*/
/*Table 7.5.5.2-1: Usage Report IE within PFCP Session Modification Response
Octet 1 and 2		Usage Report IE Type = 78 (decimal)
Octets 3 and 4		Length = n
Information elements	P       IE Type

URR ID	M                       URR ID
UR-SEQN	M                       UR-SEQN
Usage Report Trigger	M       Usage Report Trigger
Start Time	C                   Start Time
End Time	C                   End Time
Volume Measurement	    C       Volume Measurement
Duration Measurement	C       Duration Measurement
Time of First Packet	C       Time of First Packet
Time of Last Packet	C           Time of Last Packet
Usage Information 	C           Usage Information
Query URR Reference	C           Query URR Reference
Ethernet Traffic Information C  Ethernet Traffic Information
*/
type IEUsageReportModResp struct {
	IETypeLength
	URRID                      IEURRID
	URSEQN                     IEURSEQN
	UsageReportTrigger         IEUsageReportTrigger
	StartTime                  *IEStartTime
	EndTime                    *IEEndTime
	VolumeMeasurement          *IEVolumeMeasurement
	DurationMeasurement        *IEDurationMeasurement
	TimeofFirstPacket          *IETimeofFirstPacket
	TimeofLastPacket           *IETimeofLastPacket
	UsageInformation           *IEUsageInformation
	QueryURRReference          *IEQueryURRReference
	EthernetTrafficInformation *IEEthernetTrafficInformation

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUsageReportModResp) Encode() (data []byte, err error) {
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
			//Time of First Packet	C       Time of First Packet
		case IE_Time_of_First_Packet:
			// encode v
			vEnc, err = i.TimeofFirstPacket.Encode()
			if err != nil {
				return nil, err
			}
			//Time of Last Packet	C           Time of Last Packet
		case IE_Time_of_Last_Packet:
			// encode v
			vEnc, err = i.TimeofLastPacket.Encode()
			if err != nil {
				return nil, err
			}
			//Usage Information 	C           Usage Information
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
			//Ethernet Traffic Information C  Ethernet Traffic Information
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

func (i *IEUsageReportModResp) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUsageReportModResp) Len() int {
	return int(i.Length)
}

func (i *IEUsageReportModResp) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUsageReportModResp) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//URR ID	M                       URR ID
	case *IEURRID:
		i.URRID = *ie
		//UR-SEQN	M                       UR-SEQN
	case *IEURSEQN:
		i.URSEQN = *ie
		//Usage Report Trigger	M       Usage Report Trigger
	case *IEUsageReportTrigger:
		i.UsageReportTrigger = *ie
		//Start Time	C                   Start Time
	case *IEStartTime:
		i.StartTime = ie
		//End Time	C                   End Time
	case *IEEndTime:
		i.EndTime = ie
		//Volume Measurement	    C       Volume Measurement
	case *IEVolumeMeasurement:
		i.VolumeMeasurement = ie
		//Duration Measurement	C       Duration Measurement
	case *IEDurationMeasurement:
		i.DurationMeasurement = ie
		//Time of First Packet	C       Time of First Packet
	case *IETimeofFirstPacket:
		i.TimeofFirstPacket = ie
		//Time of Last Packet	C           Time of Last Packet
	case *IETimeofLastPacket:
		i.TimeofLastPacket = ie
		//Usage Information 	C           Usage Information
	case *IEUsageInformation:
		i.UsageInformation = ie
		//Query URR Reference	C           Query URR Reference
	case *IEQueryURRReference:
		i.QueryURRReference = ie
		//Ethernet Traffic Information C  Ethernet Traffic Information
	case *IEEthernetTrafficInformation:
		i.EthernetTrafficInformation = ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEUsageReportModResp) Set(v uint8) error {
	i.Type = IE_Usage_Report_Modification_Response

	return nil
}

//Additional Usage Reports Information	C    Additional Usage Reports Information
// IEAdditionalUsageReportsInformation
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
126	Additional Usage Reports Information	Extendable / Subclause 8.2.91	2
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 126 (decimal)
	3 to 4	Length = n
	5	  AURI 	Number of Additional Usage Reports value
	6	Number of Additional Usage Reports value
	7 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.91-1: Additional Usage Reports Information
*/
type IEAdditionalUsageReportsInformation struct {
	IETypeLength
	AURI   bool   `json:",omitempty"`
	Number uint16 `json:",omitempty"`
}

func (i *IEAdditionalUsageReportsInformation) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := (utils.BoolToUint8(i.AURI) << 7) + byte(i.Number>>8)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}
	tmpByte = byte(i.Number)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}
	return encBuf.Bytes(), nil
}

func (i *IEAdditionalUsageReportsInformation) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.AURI, err = utils.GetBoolInByte(tmp, 8)
	if err != nil {
		return err
	}
	tmpByte2, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.Number = uint16((tmp<<1>>1)<<8) + uint16(tmpByte2)

	return nil
}

func (i *IEAdditionalUsageReportsInformation) Len() int {
	return int(i.Length)
}

func (i *IEAdditionalUsageReportsInformation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEAdditionalUsageReportsInformation) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEAdditionalUsageReportsInformation) Set(v uint16, AURI bool) error {
	i.Type = IE_Additional_Usage_Reports_Information
	i.Length = 2
	i.AURI = AURI
	i.Number = v
	return nil
}
func (i *IEAdditionalUsageReportsInformation) Get() (v uint16, AURI bool, e error) {
	return i.Number, i.AURI, nil
}

// 86
