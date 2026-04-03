package pfcp

import (
	"bytes"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/message/pfcp/utils"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)
// N4 消息

// 7.5.9	PFCP Session Report Response
/*Table 7.5.9.1-1: Information Elements in a PFCP Session Report Response
Information elements	              P      IE Type

Cause	                              M      Cause
Offending IE	                      C      Offending IE
Update BAR	                          C      Update BAR
PFCPSRRsp-Flags	                      C      PFCPSRRsp-Flags*/

// 8.2.32	PFCPSRRsp-Flags
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 50 (decimal)
	3 to 4	Length = n
	5	Spare	Spare	Spare	Spare	Spare	Spare	Spare 	DROBU
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.32-1: PFCPSRRsp-Flags*/
type IEPFCPSRRspFlags struct {
	IETypeLength
	DROBU bool
}

func (i *IEPFCPSRRspFlags) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	tmpByte := utils.BoolToUint8(i.DROBU)
	err = encBuf.WriteByte(tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEPFCPSRRspFlags) Decode(data []byte) error {
	//	parse v
	r := bytes.NewReader(data)

	tmp, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.DROBU, err = utils.GetBoolInByte(tmp, 1)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEPFCPSRRspFlags) Len() int {
	return int(i.Length)
}

func (i *IEPFCPSRRspFlags) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPFCPSRRspFlags) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPSRRspFlags) Set(v bool) error {
	i.Type = IE_PFCPSRRsp_Flags
	i.Length = 1
	i.DROBU = v
	return nil
}

/*-	Bit 1 – DROBU (Drop Buffered Packets): if this bit is set to 1, it indicates that the UP function shall
drop all the packets currently buffered for the PFCP session, if any, prior to further applying the action
specified in the Apply Action value of the FARs.*/
const DROBU = utils.Bit_0

//Update BAR	C                                 Update BAR
// IEUpdateBAR
//Network Instance
/*IE Type value(Decimal)  Information elements	Comment/Reference	Number of Fixed Octets
12	Update BAR (PFCP Session Report Response)	Extendable / Table 7.5.9.2-1 	Not Applicable
*/
/*Table 7.5.9.2-1: Update BAR IE in PFCP Session Report Response
Octet 1 and 2		Update BAR IE Type = 12 (decimal)
Octets 3 and 4		Length = n
Information elements	P                  IE Type

BAR ID	M                                  BAR ID
Downlink Data Notification Delay	C      Downlink Data Notification Delay
DL Buffering Duration	C                  DL Buffering Duration
DL Buffering Suggested Packet Count	O      DL Buffering Suggested Packet Count
Suggested Buffering Packets Count	C      Suggested Buffering Packets Count
*/
type IEUpdateBAR struct {
	IETypeLength
	BARID                           IEBARID
	DLDataNotificationDelay         IEDownlinkDataNotificationDelay
	DLBufferingDuration             IEDLBufferingDuration
	DLBufferingSuggestedPacketCount IEDLBufferingSuggestedPacketCount
	SuggestedBufferingPacketsCount  IESuggestedBufferingPacketsCount

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUpdateBAR) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//BAR ID	M                                  BAR ID
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

	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//Downlink Data Notification Delay	C
		case IE_Downlink_Data_Notification_Delay:
			//	encode v
			vEnc, err = i.DLDataNotificationDelay.Encode()
			if err != nil {
				return nil, err
			}
			//DL Buffering Duration	C
		case IE_DL_Buffering_Duration:
			//	encode v
			vEnc, err = i.DLBufferingDuration.Encode()
			if err != nil {
				return nil, err
			}
			//DL Buffering Suggested Packet Count	O
		case IE_DL_Buffering_Suggested_Packet_Count:
			//	encode v
			vEnc, err = i.DLBufferingSuggestedPacketCount.Encode()
			if err != nil {
				return nil, err
			}
			//Suggested Buffering Packets Count	C
		case IE_Suggested_Buffering_Packets_Count:
			//	encode v
			vEnc, err = i.SuggestedBufferingPacketsCount.Encode()
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

func (i *IEUpdateBAR) Decode(data []byte) error {
	//parse v
	err := GroupedIeParse(i, data)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEUpdateBAR) Len() int {
	return int(i.Length)
}

func (i *IEUpdateBAR) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUpdateBAR) SetObjectToParent(child interface{}) error {
	switch ie := child.(type) {
	//BAR ID	M                                  BAR ID
	case *IEBARID:
		i.BARID = *ie
		//Downlink Data Notification Delay	C      Downlink Data Notification Delay
	case *IEDownlinkDataNotificationDelay:
		i.DLDataNotificationDelay = *ie
		//DL Buffering Duration	C                  DL Buffering Duration
	case *IEDLBufferingDuration:
		i.DLBufferingDuration = *ie
		//DL Buffering Suggested Packet Count	O      DL Buffering Suggested Packet Count
	case *IEDLBufferingSuggestedPacketCount:
		i.DLBufferingSuggestedPacketCount = *ie
		//Suggested Buffering Packets Count	C      Suggested Buffering Packets Count
	case *IESuggestedBufferingPacketsCount:
		i.SuggestedBufferingPacketsCount = *ie

	default:
		return fmt.Errorf("Illegal IE")
	}

	return nil
}

func (i *IEUpdateBAR) Set(v uint8) error {
	i.Type = IE_Update_BAR

	return nil
}

// 75
