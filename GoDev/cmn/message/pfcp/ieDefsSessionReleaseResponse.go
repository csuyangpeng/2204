package pfcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/message/pfcp/utils"
	"net"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)
// N4 消息

// 7.5.7	PFCP Session Deletion Response
/*
Information elements	P

Cause	                M
Offending IE	        C
Load Control Information	O
Overload Control Information	O
Usage Report	        C
*/

type IEsSessionDelResponse struct {
	Cause               IECause
	OffendingIE         *IEOffendingIE                 `json:",omitempty"`
	LCI                 *IELCI                         `json:",omitempty"`
	OverloadControlInfo *IEOCI                         `json:",omitempty"`
	UsageReport         *IEUsageReportDeletionResponse `json:",omitempty"`

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEsSessionDelResponse) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Cause	                M
	// encode v
	vEnc, err := i.Cause.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{Type: uint16(IE_Cause), Length: uint16(len(vEnc))}
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
		//Offending IE	        C
		case IE_Offending_IE:
			// encode v
			vEnc, err = i.OffendingIE.Encode()
			if err != nil {
				return nil, err
			}
			//Load Control Information	O
		case IE_Load_Control_Information:
			// encode v
			vEnc, err = i.LCI.Encode()
			if err != nil {
				return nil, err
			}
			//Overload Control Information	O
		case IE_Overload_Control_Information:
			// encode v
			vEnc, err = i.OverloadControlInfo.Encode()
			if err != nil {
				return nil, err
			}
			//Usage Report	        C
		case IE_Usage_Report_Deletion_Response:
			// encode v
			vEnc, err = i.UsageReport.Encode()
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

func (i *IEsSessionDelResponse) Decode(data []byte) error {
	// parse IE
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEsSessionDelResponse) Len() int {
	return 0
}

func (i *IEsSessionDelResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEsSessionDelResponse) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Cause	                M
	case *IECause:
		i.Cause = *ie
		//Offending IE	        C
	case *IEOffendingIE:
		i.OffendingIE = ie
		//Load Control Information	O
	case *IELCI:
		i.LCI = ie
		//Overload Control Information	O
	case *IEOCI:
		i.OverloadControlInfo = ie
		//Usage Report	        C
	case *IEUsageReportDeletionResponse:
		i.UsageReport = ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

// 8.2.1 Cause
/*.		Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 19 (decimal)
3 to 4	Length = n
5	    Cause value
Figure 8.2.1-1: Cause*/
type IECause struct {
	IETypeLength
	CauseValue uint8 `json:",omitempty"`
}

func (i *IECause) Encode() (data []byte, err error) {
	// encode Tlv
	encBuf := bytes.NewBuffer(nil)
	/*err = binary.Write(encBuf, binary.BigEndian, i.Type)
	if err != nil {
		return
	}
	err = binary.Write(encBuf, binary.BigEndian, i.Length)
	if err != nil {
		return
	}*/
	// length check
	/*if 1 != i.Len() {
		return nil, errors.New("invalid length")
	}*/
	err = binary.Write(encBuf, binary.BigEndian, i.CauseValue)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IECause) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	tmpByte, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.CauseValue = tmpByte
	return nil
}

func (i *IECause) Len() int {
	return int(i.Length)
}

func (i *IECause) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECause) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IECause) Set() {
	i.Type = IE_Cause
	i.Length = 1
	i.CauseValue = Cause_Request_accepted
}

// Table 8.2.1-1: Cause values
const (
	Cause_Reserved = 0

	// Acceptance in a response
	Cause_Request_accepted = 1 // (success)
	//2-63	Spare.                           =    2

	// Rejection in a response
	Cause_Request_rejected                   = 64 //(reason not specified)
	Cause_Session_context_not_found          = 65
	Cause_Mandatory_IE_missing               = 66
	Cause_Conditional_IE_missing             = 67
	Cause_Invalid_length                     = 68
	Cause_Mandatory_IE_incorrect             = 69
	Cause_Invalid_Forwarding_Policy          = 70
	Cause_Invalid_FTEID_allocation_option    = 71
	Cause_No_established_PFCP_Association    = 72
	Cause_Rule_creation_modification_Failure = 73
	Cause_PFCP_entity_in_congestion          = 74
	Cause_No_resources_available             = 75
	Cause_Service_not_supported              = 76
	Cause_System_failure                     = 77
	//78 to 255	Spare for future use in a response message.
)

// 8.2.22	Offending IE
/*.		Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 40 (decimal)
3 to 4	Length = 2
5 to 6	Type of the offending IE
Figure 8.2.22-1: Offending IE*/
type IEOffendingIE struct {
	IETypeLength
	TypeOffendingIE uint16 `json:",omitempty"`
}

func (i *IEOffendingIE) checklen() error {
	if 2 != i.Len() {
		return errors.New("invalid length")
	}
	return nil
}
func (i *IEOffendingIE) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.TypeOffendingIE)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil

}

func (i *IEOffendingIE) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.TypeOffendingIE)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEOffendingIE) Len() int {
	return int(i.Length)
}

func (i *IEOffendingIE) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEOffendingIE) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEOffendingIE) Set() {
	i.Type = IE_Offending_IE
	i.Length = 2
}

// Load Control Information, IE Type value(Decimal)51
// The Load Control Information grouped IE shall be encoded as shown in Figure 7.5.3.3-1
/*Octet 1 and 2		Load Control Information IE Type = 51 (decimal)
Octets 3 and 4		Length = n

Information elements	      P	Condition / Comment	  Appl.N4	IE Type
Load Control Sequence Number  M	See subclause 6.2.3.3.2 for the description and use of this parameter.	X	Sequence Number
Load Metric	                  M	See subclause 6.2.3.3.2 for the description and use of this parameter.	X	Metric
*/
type IELCI struct {
	// the Load Control Information (LCI) IE.
	IETypeLength
	SeqNum IESequenceNumber
	Metric IEMetric

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IELCI) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Load Control Sequence Number  M
	vEnc, err := i.SeqNum.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{uint16(IE_Sequence_Number), uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	//Load Metric	                  M
	vEnc, err = i.Metric.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{uint16(IE_Metric), uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IELCI) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IELCI) Len() int {
	return int(i.Length)
}

func (i *IELCI) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IELCI) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Load Control Sequence Number  M
	case *IESequenceNumber:
		i.SeqNum = *ie
		//Load Metric	                M
	case *IEMetric:
		i.Metric = *ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IELCI) Set() {
	i.Type = IE_Load_Control_Information
}

//Sequence Number
// IESequenceNumber
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
52	Sequence Number	Fixed Length / Subclause 8.2.33	4
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 52 (decimal)
	3 to 4	Length = n
	5 to 8	Sequence Number
Figure 8.2.33-1: Sequence Number
*/
type IESequenceNumber struct {
	IETypeLength
	Value uint32 `json:",omitempty"`
}

func (i *IESequenceNumber) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Value)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IESequenceNumber) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.BigEndian, &i.Value)
	if err != nil {
		return err
	}
	return nil
}

func (i *IESequenceNumber) Len() int {
	return int(i.Length)
}

func (i *IESequenceNumber) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IESequenceNumber) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IESequenceNumber) Set(v uint32) error {
	i.Type = IE_Sequence_Number
	i.Length = 4
	i.Value = v

	return nil
}
func (i *IESequenceNumber) Get() (v uint32, e error) {
	return i.Value, nil
}

//Metric
// IEMetric
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
53	Metric	Fixed Length / Subclause 8.2.34	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 53 (decimal)
	3 to 4	Length = n
	5	Metric
Figure 8.2.34-1: Metric
*/
//value in the range of 0 to 100 (inclusive)
type IEMetric struct {
	IETypeLength
	Metric uint8 `json:",omitempty"`
}

func (i *IEMetric) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	err = binary.Write(encBuf, binary.BigEndian, i.Metric)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEMetric) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.Metric = tmp
	return nil
}

func (i *IEMetric) Len() int {
	return int(i.Length)
}

func (i *IEMetric) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEMetric) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEMetric) Set(v uint8) error {
	i.Type = IE_Metric
	i.Length = 1
	i.Metric = v
	return nil
}
func (i *IEMetric) Get() (v uint8, e error) {
	return i.Metric, nil
}

// Overload Control Information,  IE Type = 54 (decimal)
/*Table 7.5.3.4-1: Overload Control Information IE within PFCP Session Establishment Response
Octet 1 and 2		Overload Control Information IE Type = 54 (decimal)
Octets 3 and 4		Length = n

Information elements	            P	Condition / Comment	Appl.N4	IE Type
Overload Control Sequence Number	M	See subclause 6.2.4.3.2 for the description and use of this parameter.	X	X	X	X	Sequence Number
Overload Reduction Metric	        M	See subclause 6.2.4.3.2 for the description and use of this parameter.	X	X	X	X	Metric
Period of Validity	                M	See subclause 6.2.4.3.2 for the description and use of this parameter.	X	X	X	X	Timer
Overload Control Information Flags	C	This IE shall be included if any of flag in this IE is set.  	X	X	X	X	OCI Flags
*/
type IEOCI struct {
	// the Overload Control Information (OCI) IE
	IETypeLength
	SeqNum           IESequenceNumber
	Metric           IEMetric
	PeriodofValidity IETimer
	OCIFlags         IEOCIFlags

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEOCI) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//Overload Control Sequence Number	M
	vEnc, err := i.SeqNum.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl := IETypeLength{uint16(IE_Sequence_Number), uint16(len(vEnc))}
	tlvEnc, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	//Overload Reduction Metric	        M
	vEnc, err = i.Metric.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{uint16(IE_Metric), uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}
	//Period of Validity	                M
	vEnc, err = i.PeriodofValidity.Encode()
	if err != nil {
		return nil, err
	}
	//encode tl
	tl = IETypeLength{uint16(IE_Timer), uint16(len(vEnc))}
	tlvEnc, err = tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnc)
	if err != nil {
		return nil, err
	}

	//optional
	//Overload Control Information Flags	C
	if i.IeFlags.Test(IE_OCI_Flags) {
		vEnc, err = i.OCIFlags.Encode()
		if err != nil {
			return nil, err
		}
		// encode TL
		tl = IETypeLength{Type: uint16(IE_OCI_Flags), Length: uint16(len(vEnc))}
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

func (i *IEOCI) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEOCI) Len() int {
	return int(i.Length)
}

func (i *IEOCI) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEOCI) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//Overload Control Sequence Number	M
	case *IESequenceNumber:
		i.SeqNum = *ie
		//Overload Reduction Metric	        M
	case *IEMetric:
		i.Metric = *ie
		//Period of Validity	                M
	case *IETimer:
		i.PeriodofValidity = *ie
		//Overload Control Information Flags	C
	case *IEOCIFlags:
		i.OCIFlags = *ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEOCI) Set() error {
	i.Type = IE_Overload_Control_Information
	return nil
}

//Timer
// IETimer
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
55	Timer	Extendable / Subclause 8.2 35	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 55 (decimal)
	3 to 4	Length = n
	5	Timer unit	   Timer value
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.35-1: Timer
*/
type IETimer struct {
	IETypeLength
	Unit  uint8 `json:",omitempty"`
	Value uint8 `json:",omitempty"`
}

func (i *IETimer) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := (i.Unit << 5) + (i.Value)
	err = binary.Write(encBuf, binary.BigEndian, tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IETimer) Decode(data []byte) error {
	r := bytes.NewReader(data)
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmp)
	//Timer value
	//Bits 5 to 1 represent the binary coded timer value.
	i.Value, err = v.GetBits(1, 5)
	if err != nil {
		return err
	}
	//Timer unit
	//Bits 6 to 8 defines the timer value unit for the timer
	i.Unit, err = v.GetBits(6, 8)
	v1 := utils.ByteOne(i.Unit)
	i.Unit, _ = v1.RightShift(6)
	if err != nil {
		return err
	}

	return nil
}

func (i *IETimer) Len() int {
	return int(i.Length)
}

func (i *IETimer) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IETimer) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

//Timer value
//Bits 5 to 1 represent the binary coded timer value.
/*Timer unit
Bits 6 to 8 defines the timer value unit for the timer as follows:
8 7 6
0 0 0  value is incremented in multiples of 2 seconds
0 0 1  value is incremented in multiples of 1 minute
0 1 0  value is incremented in multiples of 10 minutes
0 1 1  value is incremented in multiples of 1 hour
1 0 0  value is incremented in multiples of 10 hours
1 1 1  value indicates that the timer is infinite
*/
const (
	IETimer_seconds2  = 0
	IETimer_minute1   = 1
	IETimer_minutes10 = 2
	IETimer_hour1     = 3
	IETimer_hour10    = 4
)

func (i *IETimer) Set(v uint8, unit uint8) error {
	i.Type = IE_Timer
	i.Length = 1
	i.Unit = unit
	i.Value = v
	return nil
}
func (i *IETimer) Get() (v uint8, unit uint8, e error) {
	return i.Value, i.Unit, nil
}

/*type IETimer struct {
	PeriodValidity time.Duration //持续时间以Int64纳秒计数表示两个瞬间之间经过的时间。该代表将可代表的最长期限限制在约290年。
}*/

//OCI Flags
// IEOCIFlags
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
110	OCI Flags	Extendable / Subclause 8.2.76	1
*/
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 110 (decimal)
	3 to 4	Length = n
	5	Spare	                      AOCI
	s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.76-1: OCI Flags
*/
//AOCI: Associate OCI with Node ID
type IEOCIFlags struct {
	IETypeLength
	AOCI bool `json:",omitempty"`
}

func (i *IEOCIFlags) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	err = binary.Write(encBuf, binary.BigEndian, i.AOCI)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEOCIFlags) Decode(data []byte) error {
	r := bytes.NewReader(data)
	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	//Bit 1 – AOCI: Associate OCI with Node ID
	i.AOCI, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEOCIFlags) Len() int {
	return int(i.Length)
}

func (i *IEOCIFlags) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEOCIFlags) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEOCIFlags) Set(v bool) error {
	i.Type = IE_OCI_Flags
	i.Length = 1
	i.AOCI = v
	return nil
}
func (i *IEOCIFlags) Get() (v bool, e error) {
	return i.AOCI, nil
}

// Usage Report
//7.5.7.2	Usage Report IE within PFCP Session Deletion Response
/*Table 7.5.7.2-1: Usage Report IE within PFCP Session Deletion Response
Octet 1 and 2		Usage Report IE Type = 79 (decimal)
Octets 3 and 4		Length = n
*/
/*Information elements	P

URR ID	M
UR-SEQN	M
Usage Report Trigger	M
Start Time	C
End Time	C
Volume Measurement	C
Duration Measurement	C
Time of First Packet	C
Time of Last Packet	C
Usage Information 	C
Ethernet Traffic Information	C
*/
//IE_Usage_Report_Deletion_Response
type IEUsageReportDeletionResponse struct {
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
	EthernetTrafficInformation *IEEthernetTrafficInformation

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUsageReportDeletionResponse) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//URR ID	M
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
	//UR-SEQN	M
	// encode v
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
	//Usage Report Trigger	M
	// encode v
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
		//Start Time	C
		case IE_Start_Time:
			//	encode v
			vEnc, err = i.StartTime.Encode()
			if err != nil {
				return nil, err
			}
			//End Time	C
		case IE_End_Time:
			//	encode v
			vEnc, err = i.EndTime.Encode()
			if err != nil {
				return nil, err
			}
			//Volume Measurement	C
		case IE_Volume_Measurement:
			//	encode v
			vEnc, err = i.VolumeMeasurement.Encode()
			if err != nil {
				return nil, err
			}
			//Duration Measurement	C
		case IE_Duration_Measurement:
			//	encode v
			vEnc, err = i.DurationMeasurement.Encode()
			if err != nil {
				return nil, err
			}
			//Time of First Packet	C
		case IE_Time_of_First_Packet:
			//	encode v
			vEnc, err = i.TimeofFirstPacket.Encode()
			if err != nil {
				return nil, err
			}
			//Time of Last Packet	C
		case IE_Time_of_Last_Packet:
			//	encode v
			vEnc, err = i.TimeofLastPacket.Encode()
			if err != nil {
				return nil, err
			}
			//Usage Information 	C
		case IE_Usage_Information:
			//	encode v
			vEnc, err = i.UsageInformation.Encode()
			if err != nil {
				return nil, err
			}
			//Ethernet Traffic Information	C
		case IE_Ethernet_Traffic_Information:
			//	encode v
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

func (i *IEUsageReportDeletionResponse) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUsageReportDeletionResponse) Len() int {
	return int(i.Length)
}

func (i *IEUsageReportDeletionResponse) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUsageReportDeletionResponse) SetObjectToParent(child interface{}) error {
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
		//Time of First Packet	C
	case *IETimeofFirstPacket:
		i.TimeofFirstPacket = ie
		//Time of Last Packet	C
	case *IETimeofLastPacket:
		i.TimeofLastPacket = ie
		//Usage Information 	C
	case *IEUsageInformation:
		i.UsageInformation = ie
		//Ethernet Traffic Information	C
	case *IEEthernetTrafficInformation:
		i.EthernetTrafficInformation = ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
  81	           URR ID	 Extendable / Subclause 8.2.54	4*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 81 (decimal)
3 to 4	Length = n
5 to 8	URR ID value
9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.54-1: URR ID
*/
type IEURRID struct {
	IETypeLength
	URRIDValue uint32 `json:",omitempty"`
}

func (i *IEURRID) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.URRIDValue)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEURRID) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.BigEndian, &i.URRIDValue)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEURRID) Len() int {
	return int(i.Length)
}

func (i *IEURRID) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEURRID) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEURRID) Set(value uint32) {
	i.Type = IE_URR_ID
	i.Length = 4
	i.URRIDValue = value
}

func (i *IEURRID) Get() (v uint32, e error) {
	return i.URRIDValue, nil
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
  104	UR-SEQN	                Fixed Length / Subclause 8.2.71	4*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 104 (decimal)
3 to 4	Length = n
5 to 8	UR-SEQN
Figure 8.2.71-1: UR-SEQN*/
//The UR-SEQN (Usage Report Sequence Number) IE
type IEURSEQN struct {
	IETypeLength
	URSEQN uint32 `json:",omitempty"`
}

func (i *IEURSEQN) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.URSEQN)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEURSEQN) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)
	err := binary.Read(r, binary.BigEndian, &i.URSEQN)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEURSEQN) Len() int {
	return int(i.Length)
}

func (i *IEURSEQN) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l

	return nil
}

func (i *IEURSEQN) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEURSEQN) Set(value uint32) {
	i.Type = IE_UR_SEQN
	i.Length = 4
	i.URSEQN = value
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
  63	Usage Report Trigger	Extendable / Subclause 8.2.41	1*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 63 (decimal)
3 to 4	Length = n
5	IMMER	DROTH	STOPT	START	QUHTI	TIMTH	VOLTH	PERIO
6	EVETH	MACAR	ENVCL	MONIT	TERMR	LIUSA	TIMQU	VOLQU
7	Spare	                                                EVEQU
8 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.41-1: Usage Report Trigger*/
//The Usage Report Trigger IE
type IEUsageReportTrigger struct {
	IETypeLength
	//Value1 uint8
	IMMER, DROTH, STOPT, START, QUHTI, TIMTH, VOLTH, PERIO bool
	//Value2 uint8
	EVETH, MACAR, ENVCL, MONIT, TERMR, LIUSA, TIMQU, VOLQU bool
	EVEQU                                                  bool
}

func (i *IEUsageReportTrigger) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := utils.BoolToUint8(i.PERIO) +
		(utils.BoolToUint8(i.VOLTH) << 1) +
		(utils.BoolToUint8(i.TIMTH) << 2) +
		(utils.BoolToUint8(i.QUHTI) << 3) +
		(utils.BoolToUint8(i.START) << 4) +
		(utils.BoolToUint8(i.STOPT) << 5) +
		(utils.BoolToUint8(i.DROTH) << 6) +
		(utils.BoolToUint8(i.IMMER) << 7)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	tmpByte = utils.BoolToUint8(i.VOLQU) +
		(utils.BoolToUint8(i.TIMQU) << 1) +
		(utils.BoolToUint8(i.LIUSA) << 2) +
		(utils.BoolToUint8(i.TERMR) << 3) +
		(utils.BoolToUint8(i.MONIT) << 4) +
		(utils.BoolToUint8(i.ENVCL) << 5) +
		(utils.BoolToUint8(i.MACAR) << 6) +
		(utils.BoolToUint8(i.EVETH) << 7)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	tmpByte = utils.BoolToUint8(i.EVEQU)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEUsageReportTrigger) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)
	//Octet 5	IMMER	DROTH	STOPT	START	QUHTI	TIMTH	VOLTH	PERIO
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
	i.IMMER, err = utils.GetBoolInByte(tmp5, 8)
	if err != nil {
		return err
	}
	//	Octet 6 EVETH	MACAR	ENVCL	MONIT	TERMR	LIUSA	TIMQU	VOLQU
	tmp6, err := r.ReadByte()
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
	i.LIUSA, err = utils.GetBoolInByte(tmp6, 3)
	if err != nil {
		return err
	}
	i.TERMR, err = utils.GetBoolInByte(tmp6, 4)
	if err != nil {
		return err
	}
	i.MONIT, err = utils.GetBoolInByte(tmp6, 5)
	if err != nil {
		return err
	}
	i.ENVCL, err = utils.GetBoolInByte(tmp6, 6)
	if err != nil {
		return err
	}
	i.MACAR, err = utils.GetBoolInByte(tmp6, 7)
	if err != nil {
		return err
	}
	i.EVETH, err = utils.GetBoolInByte(tmp6, 8)
	if err != nil {
		return err
	}
	// Octet 7 Spare	                                                EVEQU
	tmp7, err := r.ReadByte()
	if err != nil {
		return err
	}

	i.EVEQU, err = utils.GetBoolInByte(tmp7, 1)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEUsageReportTrigger) Len() int {
	return int(i.Length)
}

func (i *IEUsageReportTrigger) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUsageReportTrigger) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEUsageReportTrigger) Set() {
	i.Type = IE_Usage_Report_Trigger
	i.Length = 3

}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
75	Start Time	Extendable / Subclause 8.2.52	4*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 75 (decimal)
3 to 4	Length = n
5 to 8	Start Time
9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.52-1: Start Time*/
// The Start Time IE
type IEStartTime struct {
	IETypeLength
	StartTime uint32 `json:",omitempty"` // seconds relative to 00:00:00 on 1 January 1900.
}

func (i *IEStartTime) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.StartTime)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEStartTime) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.StartTime)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEStartTime) Len() int {
	return int(i.Length)
}

func (i *IEStartTime) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEStartTime) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEStartTime) Set() {
	i.Type = IE_Start_Time
	i.Length = 4
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
76	End Time	Extendable / Subclause 8.2.53	4*/
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 76 (decimal)
3 to 4	Length = n
5 to 8	End Time
9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.53-1: End Time
*/
// The End Time IE
type IEEndTime struct {
	IETypeLength
	EndTime uint32 `json:",omitempty"`
}

func (i *IEEndTime) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.EndTime)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEEndTime) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.EndTime)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEndTime) Len() int {
	return int(i.Length)
}

func (i *IEEndTime) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEndTime) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEEndTime) Set() {
	i.Type = IE_End_Time
	i.Length = 4
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
66	Volume Measurement	Extendable / Subclause 8.2.44	q+7-4 */
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 66 (decimal)
3 to 4	Length = n
5	Spare	               DLVOL ULVOL TOVOL
m to (m+7)	Total Volume
p to (p+7)	Uplink Volume
q to (q+7)	Downlink Volume
s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.44-1: Volume Measurement
*/
// The Volume Measurement IE
type IEVolumeMeasurement struct {
	IETypeLength
	DLVOL, ULVOL, TOVOL bool
	TotalVolume         uint16 `json:",omitempty"`
	UplinkVolume        uint16 `json:",omitempty"`
	DownlinkVolume      uint16 `json:",omitempty"`
}

func (i *IEVolumeMeasurement) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Spare	               DLVOL ULVOL TOVOL
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

func (i *IEVolumeMeasurement) Decode(data []byte) error {
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

func (i *IEVolumeMeasurement) Len() int {
	return int(i.Length)
}

func (i *IEVolumeMeasurement) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEVolumeMeasurement) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEVolumeMeasurement) Set() {
	i.Type = IE_Volume_Measurement
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
67	Duration Measurement	Extendable / Subclause 8.2.45	4 */
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 67 (decimal)
3 to 4	Length = n
5 to 8	Duration value
9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.45-1: Duration Measurement
*/
// The Duration Measurement IE
type IEDurationMeasurement struct {
	IETypeLength
	Durationvalue uint32 `json:",omitempty"` // time in seconds
}

func (i *IEDurationMeasurement) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.Durationvalue)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEDurationMeasurement) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.Durationvalue)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEDurationMeasurement) Len() int {
	return int(i.Length)
}

func (i *IEDurationMeasurement) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEDurationMeasurement) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEDurationMeasurement) Set() {
	i.Type = IE_Duration_Measurement
	i.Length = 4
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
69	Time of First Packet	Extendable / Subclause 8.2.46	4 */
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 69 (decimal)
3 to 4	Length = n
5 to 8	Time of First Packet
9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.46-1: Time of First Packet
*/
// The Time of First Packet IE
type IETimeofFirstPacket struct {
	IETypeLength
	TimeFirstPacket uint32 `json:",omitempty"` //the time in seconds
}

func (i *IETimeofFirstPacket) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.TimeFirstPacket)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IETimeofFirstPacket) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.TimeFirstPacket)
	if err != nil {
		return err
	}
	return nil
}

func (i *IETimeofFirstPacket) Len() int {
	return int(i.Length)
}

func (i *IETimeofFirstPacket) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IETimeofFirstPacket) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IETimeofFirstPacket) Set() {
	i.Type = IE_Time_of_First_Packet
	i.Length = 4
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
70	Time of Last Packet	Extendable / Subclause 8.2.47	4 */
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 70 (decimal)
3 to 4	Length = n
5 to 8	Time of Last Packet
9 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.47-1: Time of Last Packet
*/
// The Time of Last Packet IE
type IETimeofLastPacket struct {
	IETypeLength
	TimeLastPacket uint32 `json:",omitempty"`
}

func (i *IETimeofLastPacket) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.TimeLastPacket)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IETimeofLastPacket) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.TimeLastPacket)
	if err != nil {
		return err
	}
	return nil
}

func (i *IETimeofLastPacket) Len() int {
	return int(i.Length)
}

func (i *IETimeofLastPacket) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IETimeofLastPacket) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IETimeofLastPacket) Set() {
	i.Type = IE_Time_of_Last_Packet
	i.Length = 4
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
90	Usage Information	Extendable / Subclause 8.2.59	1 */
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 90 (decimal)
3 to 4	Length = n
5	Spare	           UBE	UAE	AFT	BEF
6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.59-1: Usage Information
*/
// The Usage Information IE
type IEUsageInformation struct {
	IETypeLength
	UBE, UAE, AFT, BEF bool
}

func (i *IEUsageInformation) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := utils.BoolToUint8(i.BEF) +
		(utils.BoolToUint8(i.AFT) << 1) +
		(utils.BoolToUint8(i.UAE) << 2) +
		(utils.BoolToUint8(i.UBE) << 3)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEUsageInformation) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	tmp5, err := r.ReadByte()
	i.BEF, err = utils.GetBoolInByte(tmp5, 1)
	if err != nil {
		return err
	}
	i.AFT, err = utils.GetBoolInByte(tmp5, 2)
	if err != nil {
		return err
	}
	i.UAE, err = utils.GetBoolInByte(tmp5, 3)
	if err != nil {
		return err
	}
	i.UBE, err = utils.GetBoolInByte(tmp5, 4)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEUsageInformation) Len() int {
	return int(i.Length)
}

func (i *IEUsageInformation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUsageInformation) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEUsageInformation) Set() {
	i.Type = IE_Usage_Information
	i.Length = 1
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
143	Ethernet Traffic Information	Extendable / Table 7.5.8.3-3	Not Applicable */

/*Table 7.5.8.3-3: Ethernet Traffic Information IE within Usage Report IE
Octet 1 and 2		Ethernet Traffic Information IE Type = 143 (decimal)
Octets 3 and 4		Length = n

Information elements	P   IE Type

MAC Addresses Detected	C   MAC Addresses Detected
MAC Addresses Removed	C   MAC Addresses Removed*/
type IEEthernetTrafficInformation struct {
	IETypeLength
	MACAddrDetected IEMACAddrDetected
	MACAddrRemoved  IEMACAddrRemoved

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEEthernetTrafficInformation) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	var vEnc, tlvEnc []byte
	var tl IETypeLength
	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//MAC Addresses Detected	C   MAC Addresses Detected
		case IE_MAC_Addresses_Detected:
			vEnc, err = i.MACAddrDetected.Encode()
			if err != nil {
				return nil, err
			}

			//MAC Addresses Removed	C   MAC Addresses Removed
		case IE_MAC_Addresses_Removed:
			vEnc, err = i.MACAddrRemoved.Encode()
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

func (i *IEEthernetTrafficInformation) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEEthernetTrafficInformation) Len() int {
	return int(i.Length)
}

func (i *IEEthernetTrafficInformation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEEthernetTrafficInformation) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//MAC Addresses Detected	C
	case *IEMACAddrDetected:
		i.MACAddrDetected = *ie
		//MAC Addresses Removed	    C
	case *IEMACAddrRemoved:
		i.MACAddrRemoved = *ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEEthernetTrafficInformation) Set() {
	i.Type = IE_Ethernet_Traffic_Information
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
144	MAC Addresses Detected	Extendable / Subclause 8.2.103	7 */
// 8.2.103	MAC Addresses Detected
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 144 (decimal)
3 to 4	Length = n
5	Number of MAC addresses (k)
6 to 11	MAC address value 1
o to (o+5)	MAC address value 2 // 6 Octets
p to (p+5)	…
q to (q+5)	MAC address value k
s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.103-1: MAC addresses Detected
*/
type IEMACAddrDetected struct {
	IETypeLength
	NumberMACAddr uint8              `json:",omitempty"`
	MACAddrValues []net.HardwareAddr `json:",omitempty"` // 6 byte
}

func (i *IEMACAddrDetected) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Number of MAC addresses (k)
	i.NumberMACAddr = uint8(len(i.MACAddrValues))
	err = binary.Write(encBuf, binary.BigEndian, i.NumberMACAddr)
	if err != nil {
		return nil, err
	}

	//6 to 11	MAC address value 1
	//MAC address values shall be encoded as 12-digit hexadecimal numbers.
	for _, v := range i.MACAddrValues {
		if len(v) >= 6 {
			_, err = encBuf.Write(v[:6])
			if err != nil {
				return nil, err
			}
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEMACAddrDetected) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.NumberMACAddr)
	if err != nil {
		return err
	}

	i.MACAddrValues = make([]net.HardwareAddr, i.NumberMACAddr)
	for n, _ := range i.MACAddrValues {
		Bytes6 := make(net.HardwareAddr, 6)
		err := binary.Read(r, binary.BigEndian, Bytes6)
		if err != nil {
			return err
		}
		i.MACAddrValues[n] = Bytes6
	}

	return nil
}

func (i *IEMACAddrDetected) Len() int {
	return int(i.Length)
}

func (i *IEMACAddrDetected) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEMACAddrDetected) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEMACAddrDetected) Set() {
	i.Type = IE_MAC_Addresses_Detected
	i.Length = 7 //当前协议支持1个MAC
}

/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
1145	MAC Addresses Removed	Extendable / Subclause 8.2.104	7 */
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = 145 (decimal)
3 to 4	Length = n
5	Number of MAC addresses (k)
6 to 11	MAC address value 1
o to (o+5)	MAC address value 2
p to (p+5)	…
q to (q+5)	MAC address value k
s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.104-1: MAC addresses Removed
*/
// 8.2.104	MAC Addresses Removed
type IEMACAddrRemoved struct {
	IETypeLength
	NumberMACAddr uint8              `json:",omitempty"`
	MACAddrValues []net.HardwareAddr `json:",omitempty"` // 6 byte
}

func (i *IEMACAddrRemoved) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	//5	Number of MAC addresses (k)
	i.NumberMACAddr = uint8(len(i.MACAddrValues))
	err = binary.Write(encBuf, binary.BigEndian, i.NumberMACAddr)
	if err != nil {
		return nil, err
	}

	//6 to 11	MAC address value 1
	//MAC address values shall be encoded as 12-digit hexadecimal numbers.
	for _, v := range i.MACAddrValues {
		if len(v) >= 6 {
			_, err = encBuf.Write(v[:6])
			if err != nil {
				return nil, err
			}
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEMACAddrRemoved) Decode(data []byte) error {
	//parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.NumberMACAddr)
	if err != nil {
		return err
	}

	i.MACAddrValues = make([]net.HardwareAddr, i.NumberMACAddr)
	for n, _ := range i.MACAddrValues {
		Bytes6 := make(net.HardwareAddr, 6)
		err := binary.Read(r, binary.BigEndian, Bytes6)
		if err != nil {
			return err
		}
		i.MACAddrValues[n] = Bytes6
	}

	return nil
}

func (i *IEMACAddrRemoved) Len() int {
	return int(i.Length)
}

func (i *IEMACAddrRemoved) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEMACAddrRemoved) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEMACAddrRemoved) Set() {
	i.Type = IE_MAC_Addresses_Removed
	i.Length = 7 //当前协议支持1个MAC
}

// 1010
