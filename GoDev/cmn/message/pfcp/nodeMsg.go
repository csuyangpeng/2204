package pfcp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/message/pfcp/utils"
	"net"
	"time"
)

//7.4	PFCP Node Related Messages

//7.4.2.1	Heartbeat Request
type HeartbeatRequest struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IERecoveryTimeStamp
}

func (p *HeartbeatRequest) MarshalBinary() (data []byte, err error) {
	//	encode ie
	// 编码 必选 IE
	encBuf := bytes.NewBuffer(nil)

	//RecoveryTimeStamp	M
	//encode v
	vEnc, err := p.IE.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Recovery_Time_Stamp)
	tl.Length = uint16(len(vEnc))
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

func (p *HeartbeatRequest) UnmarshalBinary(data []byte) error {
	//	parse pfcp IE
	p.IE = &IERecoveryTimeStamp{}
	r := bytes.NewReader(data)
	// Tlv 解析
	var Tmp16 uint16
	err := binary.Read(r, binary.BigEndian, &Tmp16)
	if err != nil {
		return err
	}
	p.IE.Type = Tmp16

	err = binary.Read(r, binary.BigEndian, &Tmp16)
	if err != nil {
		return err
	}
	p.IE.Length = Tmp16

	//5 to 8	Recovery Time Stamp value
	var tmp32 uint32
	err = binary.Read(r, binary.BigEndian, &tmp32)
	if err != nil {
		return err
	}
	///0x83aa7e80 是1900年1月1日0点0分0秒 与 1970年01月01日0点0分0秒的差值
	p.IE.RecoveryTimeStamp = time.Unix(int64(tmp32)-0x83aa7e80, 0)
	return nil
}

func (p *HeartbeatRequest) Len() int {
	return int(p.PfcpHeader.Length)
}

//7.4.2.2	Heartbeat Response
type HeartbeatResponse struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IERecoveryTimeStamp
}

func (p *HeartbeatResponse) MarshalBinary() (data []byte, err error) {
	//	encode ie
	// 编码 必选 IE
	encBuf := bytes.NewBuffer(nil)

	//RecoveryTimeStamp	M
	//encode v
	ieEnc, err := p.IE.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Recovery_Time_Stamp)
	tl.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (p *HeartbeatResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
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
	p.IE = &IERecoveryTimeStamp{}
	p.IE.Type = t
	p.IE.Length = l
	//decode v
	err = p.IE.Decode(data[4 : 4+l])
	if err != nil {
		return err
	}

	return nil

}

func (p *HeartbeatResponse) Len() int {
	return int(p.PfcpHeader.Length)
}

/*Table 7.4.2.1-1: Information Elements in Heartbeat Request
Information elements	P
Recovery Time Stamp	M
*/
type IERecoveryTimeStamp struct {
	IETypeLength
	RecoveryTimeStamp time.Time
}

func (i *IERecoveryTimeStamp) Encode() (data []byte, err error) {
	// Tlv encode
	encBuf := bytes.NewBuffer(nil)
	/*err = binary.Write(encBuf, binary.BigEndian, i.Type)
	if err != nil {
		return
	}
	err = binary.Write(encBuf, binary.BigEndian, i.Length)
	if err != nil {
		return
	}*/
	//fmt.Printf("%#x\n", i.RecoveryTimeStamp.UTC().Unix())
	//fmt.Println(i.RecoveryTimeStamp.UTC())
	err = binary.Write(encBuf, binary.BigEndian, uint32(i.RecoveryTimeStamp.UTC().Unix()+0x83aa7e80))
	if err != nil {
		return
	}
	return encBuf.Bytes(), nil
}

func (i *IERecoveryTimeStamp) Decode(data []byte) error {
	//	parse pfcp IE todo Heartbeat
	r := bytes.NewReader(data)
	/*// Tlv parse
	var Tmp16 uint16
	err := binary.Read(r, binary.BigEndian, &Tmp16)
	if err != nil {
		return err
	}
	i.Type = Tmp16

	err = binary.Read(r, binary.BigEndian, &Tmp16)
	if err != nil {
		return err
	}
	i.Length = Tmp16*/
	//5 to 8	Recovery Time Stamp value
	var tmp32 uint32
	err := binary.Read(r, binary.BigEndian, &tmp32)
	if err != nil {
		return err
	}
	///0x83aa7e80 是1900年1月1日0点0分0秒 与 1970年01月01日0点0分0秒的差值
	i.RecoveryTimeStamp = time.Unix(int64(tmp32)-0x83aa7e80, 0)
	return nil
}

func (i *IERecoveryTimeStamp) Len() int {
	return int(i.Length)
}

func (i *IERecoveryTimeStamp) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IERecoveryTimeStamp) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

//7.4.3	PFCP PFD Management
//7.4.3.1	PFCP PFD Management Request
type PFCPPFDManagementRequest struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPPFDManagementRequest
}

func (p *PFCPPFDManagementRequest) MarshalBinary() (data []byte, err error) {
	// encode IE
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PFCPPFDManagementRequest) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPPFDManagementRequest{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPPFDManagementRequest) Len() int {
	return int(p.PfcpHeader.Length)
}

/*Table 7.4.3.1-1: Information Elements in PFCP PFD Management Request
Information elements	P
Application ID's PFDs	C
*/
type IEPFCPPFDManagementRequest struct {
	ApplicationIDsPFDs *IEApplicationIDsPFDs

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFCPPFDManagementRequest) Encode() (data []byte, err error) {
	// encode ie
	// Mandatory ie
	encBuf := bytes.NewBuffer(nil)

	//Application ID's PFDs	C
	vEnc, err := i.ApplicationIDsPFDs.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Application_IDs_PFDs)
	tl.Length = uint16(len(vEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}

	_, err = encBuf.Write(vEnc)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEPFCPPFDManagementRequest) Decode(data []byte) error {
	// parse IE
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEPFCPPFDManagementRequest) Len() int {
	return 0
}

func (i *IEPFCPPFDManagementRequest) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPPFDManagementRequest) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IEApplicationIDsPFDs:
		i.ApplicationIDsPFDs = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

/*Table 7.4.3.1-2: Application ID's PFDs
Octet 1 and 2		Application ID's PFDs IE Type = 58 (decimal)
Octets 3 and 4		Length = n
Information elements	P

Application ID	M
PFD context	C
*/
type IEApplicationIDsPFDs struct {
	IETypeLength
	ApplicationID *IEApplicationID
	PFD           *IEPFDContext

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

// grouped IE
func (i *IEApplicationIDsPFDs) Encode() (data []byte, err error) {
	// encode ie
	// Mandatory ie
	encBuf := bytes.NewBuffer(nil)

	// Application ID	M
	vEnc, err := i.ApplicationID.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Application_ID)
	tl.Length = uint16(len(vEnc))
	tlvEnv, err := tl.EncodeTlV(vEnc)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(tlvEnv)
	if err != nil {
		return
	}

	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//PFD context	C
		case IE_PFD_context:
			vEnc, err = i.PFD.Encode()
			if err != nil {
				return
			}
		default:

		}
		// TL 编码
		hd := IETypeLength{}
		hd.Type = uint16(id)
		hd.Length = uint16(len(vEnc))
		err = EncodeTl(encBuf, hd)
		if err != nil {
			return
		}
		_, err = encBuf.Write(vEnc)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEApplicationIDsPFDs) Decode(data []byte) error {
	// parse IE
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEApplicationIDsPFDs) Len() int {
	return int(i.Length)
}

func (i *IEApplicationIDsPFDs) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEApplicationIDsPFDs) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IEApplicationID:
		i.ApplicationID = Ie
	case *IEPFDContext:
		i.PFD = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

func (i *IEApplicationIDsPFDs) Set(v uint8) error {
	i.Type = IE_Application_IDs_PFDs
	return nil
}

/*Table 7.4.3.1-3: PFD context
Octet 1 and 2		PFD context IE Type = 59 (decimal)
Octets 3 and 4		Length = n
Information elements	P

PFD Contents 	M
*/

type IEPFDContext struct {
	IETypeLength
	PFDContexts *IEPFDContents

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFDContext) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	ieEnc, err := i.PFDContexts.Encode()
	if err != nil {
		return nil, err
	}
	// TL 编码
	tl := IETypeLength{}
	tl.Type = uint16(IE_PFD_contents)
	tl.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}

	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEPFDContext) Decode(data []byte) error {
	// parse v
	// parse IE
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo 必选IE检查
	return nil
}

func (i *IEPFDContext) Len() int {
	return int(i.Length)
}

func (i *IEPFDContext) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPFDContext) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IEPFDContents:
		i.PFDContexts = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

func (i *IEPFDContext) Set(v uint8) error {
	i.Type = IE_PFD_context
	return nil
}

//8.2.39	PFD Contents
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 61 (decimal)
	3 to 4	Length = n
	5	ADNP	AURL	AFD	DNP	CP	DN	URL	FD
	6	Spare
	m to (m+1)	Length of Flow Description
	(m+2) to p	Flow Description
	q to (q+1)	Length of URL
	(q+2) to r	URL
	s to (s+1)	Length of Domain Name
	(s+2) to t	Domain Name
	u to (u+1)	Length of Custom PFD Content
	(u+2) to v	Custom PFD Content
	w to (w+1)	Length of Domain Name Protocol
	(w+2) to x	Domain Name Protocol
	y to (y+1)	Length of Additional Flow Description
	(y+2) to z	Additional Flow Description
	a to (a+1)	Length of Additional URL
	(a+2) to b	Additional URL
	c to (c+1)	Length of Additional Domain Name and Domain Name Protocol
	(c+2) to d	Additional Domain Name and Domain Name Protocol
	e to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.39-1: PFD Contents*/
type IEPFDContents struct {
	IETypeLength
	FD      bool
	URLFlag bool
	DN      bool
	CP      bool
	DNP     bool
	AFD     bool
	AURL    bool
	ADNP    bool
	Spare   uint8

	LengthofFlowDes uint16
	FlowDescription string

	LengthofURL uint16
	URL         string

	//	s to (s+1)	Length of Domain Name
	//	(s+2) to t	Domain Name
	LengthofDomainName uint16
	DomainName         string
	//	u to (u+1)	Length of Custom PFD Content
	//	(u+2) to v	Custom PFD Content
	LengthofCustomPFDContent uint16
	CustomPFDContent         string
	//	w to (w+1)	Length of Domain Name Protocol
	//	(w+2) to x	Domain Name Protocol
	LengthofDomainNameProto uint16
	DomainNameProtocol      string
	//	y to (y+1)	Length of Additional Flow Description
	//	(y+2) to z	Additional Flow Description
	LengthofAddFlowDesc uint16
	AdditionalFlowDesc  string
	//	a to (a+1)	Length of Additional URL
	//	(a+2) to b	Additional URL
	LengthofAddURL uint16
	AdditionalURL  string
	//	c to (c+1)	Length of Additional Domain Name and Domain Name Protocol
	//	(c+2) to d	Additional Domain Name and Domain Name Protocol
	LengthofAddDNandDNProto uint16
	AddDNandDNProto         string
}

func (i *IEPFDContents) Encode() (data []byte, err error) {
	// encode v
	encV := bytes.NewBuffer(nil)

	// 5	ADNP	AURL	AFD	DNP	CP	DN	URL	FD
	var tmpByte byte
	tmpByte = utils.BoolToUint8(i.FD) +
		(utils.BoolToUint8(i.URLFlag) << 1) +
		(utils.BoolToUint8(i.DN) << 2) +
		(utils.BoolToUint8(i.CP) << 3) +
		(utils.BoolToUint8(i.DNP) << 4) +
		(utils.BoolToUint8(i.AFD) << 5) +
		(utils.BoolToUint8(i.AURL) << 6) +
		(utils.BoolToUint8(i.ADNP) << 7)

	err = encV.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}
	// 6	Spare
	err = encV.WriteByte(i.Spare)
	if err != nil {
		return nil, err
	}

	// m to (m+1)	Length of Flow Description
	// (m+2) to p	Flow Description
	if i.FD {

		i.LengthofFlowDes = uint16(len(i.FlowDescription))
		err = binary.Write(encV, binary.BigEndian, i.LengthofFlowDes)
		if err != nil {
			return nil, err
		}

		err = binary.Write(encV, binary.BigEndian, []byte(i.FlowDescription))
		if err != nil {
			return nil, err
		}
	}

	if i.URLFlag {
		i.LengthofURL = uint16(len(i.URL))
		err = binary.Write(encV, binary.BigEndian, i.LengthofURL)
		if err != nil {
			return nil, err
		}

		err = binary.Write(encV, binary.BigEndian, []byte(i.URL))
		if err != nil {
			return nil, err
		}
	}

	if i.DN {
		i.LengthofDomainName = uint16(len(i.DomainName))
		err = binary.Write(encV, binary.BigEndian, i.LengthofDomainName)
		if err != nil {
			return nil, err
		}

		err = binary.Write(encV, binary.BigEndian, []byte(i.DomainName))
		if err != nil {
			return nil, err
		}
	}

	if i.CP {
		i.LengthofCustomPFDContent = uint16(len(i.CustomPFDContent))
		err = binary.Write(encV, binary.BigEndian, i.LengthofCustomPFDContent)
		if err != nil {
			return nil, err
		}

		err = binary.Write(encV, binary.BigEndian, i.CustomPFDContent)
		if err != nil {
			return nil, err
		}
	}

	if i.DNP {
		i.LengthofDomainNameProto = uint16(len(i.DomainNameProtocol))
		err = binary.Write(encV, binary.BigEndian, i.LengthofDomainNameProto)
		if err != nil {
			return nil, err
		}

		err = binary.Write(encV, binary.BigEndian, i.DomainNameProtocol)
		if err != nil {
			return nil, err
		}
	}

	if i.AFD {
		i.LengthofAddFlowDesc = uint16(len(i.AdditionalFlowDesc))
		err = binary.Write(encV, binary.BigEndian, i.LengthofAddFlowDesc)
		if err != nil {
			return nil, err
		}

		err = binary.Write(encV, binary.BigEndian, i.AdditionalFlowDesc)
		if err != nil {
			return nil, err
		}
	}

	if i.AURL {
		i.LengthofAddURL = uint16(len(i.AdditionalURL))
		err = binary.Write(encV, binary.BigEndian, i.LengthofAddURL)
		if err != nil {
			return nil, err
		}

		err = binary.Write(encV, binary.BigEndian, i.AdditionalURL)
		if err != nil {
			return nil, err
		}
	}

	if i.ADNP {
		i.LengthofAddDNandDNProto = uint16(len(i.AddDNandDNProto))
		err = binary.Write(encV, binary.BigEndian, i.LengthofAddDNandDNProto)
		if err != nil {
			return nil, err
		}

		err = binary.Write(encV, binary.BigEndian, i.AddDNandDNProto)
		if err != nil {
			return nil, err
		}
	}

	return encV.Bytes(), nil
}

func (i *IEPFDContents) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	//5	ADNP	AURL	AFD	DNP	CP	DN	URL	FD
	tmpByte, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.FD, err = utils.GetBoolInByte(tmpByte, 1)
	if err != nil {
		return err
	}
	i.URLFlag, err = utils.GetBoolInByte(tmpByte, 2)
	if err != nil {
		return err
	}
	i.DN, err = utils.GetBoolInByte(tmpByte, 3)
	if err != nil {
		return err
	}
	i.CP, err = utils.GetBoolInByte(tmpByte, 4)
	if err != nil {
		return err
	}
	i.DNP, err = utils.GetBoolInByte(tmpByte, 5)
	if err != nil {
		return err
	}
	i.AFD, err = utils.GetBoolInByte(tmpByte, 6)
	if err != nil {
		return err
	}
	i.AURL, err = utils.GetBoolInByte(tmpByte, 7)
	if err != nil {
		return err
	}
	i.ADNP, err = utils.GetBoolInByte(tmpByte, 8)
	if err != nil {
		return err
	}

	i.Spare, err = r.ReadByte()
	if err != nil {
		return err
	}
	// m to (m+1)	Length of Flow Description
	// (m+2) to p	Flow Description
	if i.FD {
		err = binary.Read(r, binary.BigEndian, &i.LengthofFlowDes)
		if err != nil {
			return err
		}
		tmpBytes := make([]byte, i.LengthofFlowDes)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.FlowDescription = string(tmpBytes)
	}
	// q to (q+1)	Length of URL
	// (q+2) to r	URL
	if i.URLFlag {
		err = binary.Read(r, binary.BigEndian, &i.LengthofURL)
		if err != nil {
			return err
		}
		tmpBytes := make([]byte, i.LengthofURL)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.URL = string(tmpBytes)
	}
	// s to (s+1)	Length of Domain Name
	// (s+2) to t	Domain Name
	if i.DN {
		err = binary.Read(r, binary.BigEndian, &i.LengthofDomainName)
		if err != nil {
			return err
		}
		tmpBytes := make([]byte, i.LengthofDomainName)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.DomainName = string(tmpBytes)
	}

	// u to (u+1)	Length of Custom PFD Content
	// (u+2) to v	Custom PFD Content
	if i.CP {
		err = binary.Read(r, binary.BigEndian, &i.LengthofCustomPFDContent)
		if err != nil {
			return err
		}
		tmpBytes := make([]byte, i.LengthofCustomPFDContent)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.CustomPFDContent = string(tmpBytes)
	}

	// w to (w+1)	Length of Domain Name Protocol
	// (w+2) to x	Domain Name Protocol
	if i.DNP {
		err = binary.Read(r, binary.BigEndian, &i.LengthofDomainNameProto)
		if err != nil {
			return err
		}
		tmpBytes := make([]byte, i.LengthofDomainNameProto)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.DomainNameProtocol = string(tmpBytes)
	}
	// y to (y+1)	Length of Additional Flow Description
	// (y+2) to z	Additional Flow Description
	if i.AFD {
		err = binary.Read(r, binary.BigEndian, &i.LengthofAddFlowDesc)
		if err != nil {
			return err
		}
		tmpBytes := make([]byte, i.LengthofAddFlowDesc)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.AdditionalFlowDesc = string(tmpBytes)
	}

	// a to (a+1)	Length of Additional URL
	// (a+2) to b	Additional URL
	if i.AURL {
		err = binary.Read(r, binary.BigEndian, &i.LengthofAddURL)
		if err != nil {
			return err
		}
		tmpBytes := make([]byte, i.LengthofAddURL)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.AdditionalURL = string(tmpBytes)
	}

	//	c to (c+1)	Length of Additional Domain Name and Domain Name Protocol
	//	(c+2) to d	Additional Domain Name and Domain Name Protocol
	if i.ADNP {
		err = binary.Read(r, binary.BigEndian, &i.LengthofAddDNandDNProto)
		if err != nil {
			return err
		}
		tmpBytes := make([]byte, i.LengthofAddDNandDNProto)
		_, err = r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.AddDNandDNProto = string(tmpBytes)
	}

	return nil
}

func (i *IEPFDContents) Len() int {
	return int(i.Length)
}

func (i *IEPFDContents) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPFDContents) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFDContents) Set(v uint8) error {
	i.Type = IE_PFD_contents
	return nil
}

//7.4.3.2	PFCP PFD Management Response
type PFCPPFDManagementResponse struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPPFDManagementResponse
}

func (p *PFCPPFDManagementResponse) MarshalBinary() (data []byte, err error) {
	// encode IE
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (p *PFCPPFDManagementResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPPFDManagementResponse{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPPFDManagementResponse) Len() int {
	return int(p.PfcpHeader.Length)
}

/*Table 7.4.3.2-1: Information Elements in PFCP PFD Management Response
Information elements	P

Cause	M
Offending IE	C
*/
type IEPFCPPFDManagementResponse struct {
	Cause       *IECause
	OffendingIE *IEOffendingIE

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFCPPFDManagementResponse) Encode() (data []byte, err error) {
	// encode ie
	encBuf := bytes.NewBuffer(nil)

	// Mandatory ie
	//	Cause	M
	ieEnc, err := i.Cause.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Cause)
	tl.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}

	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}

	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		// Offending IE	C
		case IE_Offending_IE:
			ieEnc, err = i.OffendingIE.Encode()
			if err != nil {
				return nil, err
			}
		default:

		}
		// TL 编码
		hd := IETypeLength{}
		hd.Type = uint16(id)
		hd.Length = uint16(len(ieEnc))
		err = EncodeTl(encBuf, hd)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(ieEnc)
		if err != nil {
			return nil, err
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEPFCPPFDManagementResponse) Decode(data []byte) error {
	// parse IE
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo 必选IE检查
	return nil
}

func (i *IEPFCPPFDManagementResponse) Len() int {
	return 0
}

func (i *IEPFCPPFDManagementResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPPFDManagementResponse) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IECause:
		i.Cause = Ie
	case *IEOffendingIE:
		i.OffendingIE = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

// 7.4.4	PFCP Association messages
// 7.4.4.1	PFCP Association Setup Request
type PFCPAssociationSetupRequest struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPAssociationSetupRequest
}

func (p *PFCPAssociationSetupRequest) MarshalBinary() (data []byte, err error) {
	// IE 编码
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PFCPAssociationSetupRequest) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPAssociationSetupRequest{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPAssociationSetupRequest) Len() int {
	return int(p.PfcpHeader.Length)
}

//Table 7.4.4.1-1: Information Elements in a PFCP Association Setup Request
/*Information elements	P

Node ID	M
Recovery Time Stamp	M
UP Function Features	C
CP Function Features	C
User Plane IP Resource Information 	O
*/
type IEPFCPAssociationSetupRequest struct {
	NodeID                         *IENodeID
	RecoveryTimeStamp              *IERecoveryTimeStamp
	UPFunctionFeatures             *IEUPFunctionFeatures
	CPFunctionFeatures             *IECPFunctionFeatures
	UserPlaneIPResourceInformation []*IEUserPlaneIPResourceInformation

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFCPAssociationSetupRequest) GetGrouped() bool {
	return true
}

func (i *IEPFCPAssociationSetupRequest) Encode() (data []byte, err error) {
	//	encode ie
	// 编码 必选 IE
	encBuf := bytes.NewBuffer(nil)

	//Node ID	M
	ieEnc, err := i.NodeID.Encode()
	if err != nil {
		return nil, err
	}
	// TL 编码
	hd := IETypeLength{}
	hd.Type = uint16(IE_Node_ID)
	hd.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, hd)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}

	//Recovery Time Stamp	M
	ieEnc, err = i.RecoveryTimeStamp.Encode()
	if err != nil {
		return nil, err
	}
	// TL 编码
	hd.Type = uint16(IE_Recovery_Time_Stamp)
	hd.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, hd)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		case IE_UP_Function_Features: //UP Function Features	C
			ieEnc, err = i.UPFunctionFeatures.Encode()
			if err != nil {
				return nil, err
			}
			//_, err = encBuf.Write(ieEnc)
			//if err != nil {
			//	return nil, err
			//}
		case IE_CP_Function_Features: //CP Function Features	C
			ieEnc, err = i.CPFunctionFeatures.Encode()
			if err != nil {
				return nil, err
			}
			//_, err = encBuf.Write(ieEnc)
			//if err != nil {
			//	return nil, err
			//}
		case IE_User_Plane_IP_Resource_Information: //User Plane IP Resource Information 	O
			for _, v := range i.UserPlaneIPResourceInformation[:len(i.UserPlaneIPResourceInformation)-1] {
				//	encode v
				vEnc, err := v.Encode()
				if err != nil {
					return nil, err
				}
				// encode tl
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
			// encode v
			ieEnc, err = i.UserPlaneIPResourceInformation[len(i.UserPlaneIPResourceInformation)-1].Encode()
			if err != nil {
				return nil, err
			}
			//ieEnc, err = i.UserPlaneIPResourceInformation.Encode()
			//if err != nil {
			//	return nil, err
			//}
			//_, err = encBuf.Write(ieEnc)
			//if err != nil {
			//	return nil, err
			//}

		default:

		}
		// TL 编码
		hd := IETypeLength{}
		hd.Type = uint16(id)
		hd.Length = uint16(len(ieEnc))
		err = EncodeTl(encBuf, hd)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(ieEnc)
		if err != nil {
			return nil, err
		}
	}
	return encBuf.Bytes(), nil
}

// 消息级IE列表解析
func (i *IEPFCPAssociationSetupRequest) Decode(data []byte) error {
	// parse IE
	i.Grouped = true
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo 必选IE检查
	return nil
}

// 应用于消息的IE组
func nodeGroupedIeParse(root PfcpIeInterface, data []byte) error {
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
			return fmt.Errorf("type error")
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
	// 必选

	// 可选
	return nil
}

func (i *IEPFCPAssociationSetupRequest) Len() int {
	return 0
}

func (i *IEPFCPAssociationSetupRequest) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPAssociationSetupRequest) SetObjectToParent(child interface{}) error {

	switch Ie := child.(type) {
	case *IENodeID:
		i.NodeID = Ie
	case *IERecoveryTimeStamp:
		i.RecoveryTimeStamp = Ie
	case *IEUPFunctionFeatures:
		i.UPFunctionFeatures = Ie
	case *IECPFunctionFeatures:
		i.CPFunctionFeatures = Ie
	case *IEUserPlaneIPResourceInformation:
		i.UserPlaneIPResourceInformation = append(i.UserPlaneIPResourceInformation, Ie)
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

// UP Function Features
/*
8.2.25	UP Function Features
		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 43 (decimal)
	3 to 4	Length = n
	5 to 6	Supported-Features
	7 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.25-1: UP Function Features
*/
type IEUPFunctionFeatures struct {
	IETypeLength
	SupportedFeatures uint16
}

func (i *IEUPFunctionFeatures) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	/*err = binary.Write(encBuf, binary.BigEndian, i.Type)
	if err != nil {
		return
	}
	err = binary.Write(encBuf, binary.BigEndian, i.Length)
	if err != nil {
		return
	}*/
	err = binary.Write(encBuf, binary.BigEndian, i.SupportedFeatures)
	if err != nil {
		return
	}
	return encBuf.Bytes(), nil
}

func (i *IEUPFunctionFeatures) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.SupportedFeatures)
	if err != nil {
		return err
	}

	return nil
}

func (i *IEUPFunctionFeatures) Len() int {
	return int(i.Length)
}

func (i *IEUPFunctionFeatures) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUPFunctionFeatures) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEUPFunctionFeatures) Set(v uint8) error {
	i.Type = IE_UP_Function_Features
	return nil
}

const (
	BitMask16 = 1 << 15
	BitMask15 = 1 << 14
	BitMask14 = 1 << 13
	BitMask13 = 1 << 12
	BitMask12 = 1 << 11
	BitMask11 = 1 << 10
	BitMask10 = 1 << 9
	BitMask9  = 1 << 8

	BitMask8 = 1 << 7
	BitMask7 = 1 << 6
	BitMask6 = 1 << 5
	BitMask5 = 1 << 4
	BitMask4 = 1 << 3
	BitMask3 = 1 << 2
	BitMask2 = 1 << 1
	BitMask1 = 1
)

// Table 8.2.25-1: UP Function Features
const (
	//Feature Octet / Bit	Feature
	//5/1	BUCP
	BUCP = BitMask1
	//5/2	DDND
	DDND = BitMask2
	//5/3	DLBD
	DLBD = BitMask3
	//5/4	TRST
	TRST = BitMask4
	//5/5	FTUP
	FTUP = BitMask5
	//5/6	PFDM
	PFDM = BitMask6
	//5/7	HEEU
	HEEU = BitMask7
	//5/8	TREU
	TREU = BitMask8
	//6/1	EMPU
	EMPU = BitMask9
	//6/2	PDIU
	PDIU = BitMask10
	//6/3	UDBC
	UDBC = BitMask11
	//6/4	QUOAC
	QUOAC = BitMask12
	//6/5	TRACE
	TRACE = BitMask13
	//6/6	FRRT
	FRRT = BitMask14
	//6/7	PFDE
	PFDE = BitMask15
)

//CP Function Features
/*
8.2.58	CP Function Features
		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 89 (decimal)
	3 to 4	Length = n
	5 	Supported-Features
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.58-1: CP Function Features
*/
type IECPFunctionFeatures struct {
	IETypeLength
	SupportedFeatures uint8
}

func (i *IECPFunctionFeatures) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	/*err = binary.Write(encBuf, binary.BigEndian, i.Type)
	if err != nil {
		return
	}
	err = binary.Write(encBuf, binary.BigEndian, i.Length)
	if err != nil {
		return
	}*/
	err = binary.Write(encBuf, binary.BigEndian, i.SupportedFeatures)
	if err != nil {
		return
	}
	return encBuf.Bytes(), nil
}

func (i *IECPFunctionFeatures) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.SupportedFeatures)
	if err != nil {
		return err
	}

	return nil
}

func (i *IECPFunctionFeatures) Len() int {
	return int(i.Length)
}

func (i *IECPFunctionFeatures) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IECPFunctionFeatures) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IECPFunctionFeatures) Set(v uint8) error {
	i.Type = IE_CP_Function_Features
	return nil
}

//Table 8.2.58-1: CP Function Features
const (
	//	Feature Octet / Bit	Feature
	//5/1	LOAD
	LOAD = BitMask1
	//5/2	OVRL
	OVRL = BitMask2
)

//8.2.82	User Plane IP Resource Information
/*
		Bits
	Octets	8	  7	      6	    5 4 3   2	1
	1 to 2	Type = 116 (decimal)
	3 to 4	Length = n
	5 	Spare	ASSOSI	ASSONI	TEIDRI	V6	V4
	6  	TEID Range
	m to (m+3) 	IPv4 address
	p to (p+15) 	IPv6 address
	k to l	Network Instance
	r 	Spare	Source Interface
	s to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.82-1: User Plane IP Resource Information
*/
type IEUserPlaneIPResourceInformation struct {
	IETypeLength
	V4     bool
	V6     bool
	TEIDRI uint8 // 0x00011100
	ASSONI bool
	ASSOSI bool

	TEIDRange       uint8
	IPv4address     net.IP
	IPv6address     net.IP
	NetworkInstance string // dnn
	SourceInterface uint8  // 0x00001111
}

func (i *IEUserPlaneIPResourceInformation) Encode() (data []byte, err error) {
	encV := bytes.NewBuffer(nil)

	//5 	Spare	ASSOSI	ASSONI	TEIDRI	V6	V4
	var tmpByte byte
	tmpByte = utils.BoolToUint8(i.V4) +
		(utils.BoolToUint8(i.V6) << 1) +
		(i.TEIDRI << 2) +
		(utils.BoolToUint8(i.ASSONI) << 5) +
		(utils.BoolToUint8(i.ASSOSI) << 6)
	err = binary.Write(encV, binary.BigEndian, tmpByte)
	if err != nil {
		return
	}

	//6  	TEID Range
	if i.TEIDRI != 0 {
		err = binary.Write(encV, binary.BigEndian, i.TEIDRange)
		if err != nil {
			return
		}
	}
	//m to (m+3) 	IPv4 address
	if i.V4 {
		err = binary.Write(encV, binary.BigEndian, []byte(i.IPv4address.To4()))
		if err != nil {
			return
		}

	}
	if i.V6 {
		err = binary.Write(encV, binary.BigEndian, []byte(i.IPv6address.To16()))
		if err != nil {
			return
		}
	}

	//k to l	Network Instance
	if i.ASSONI {
		encV.Write([]byte(i.NetworkInstance))
	}

	if i.ASSOSI {
		encV.WriteByte(i.SourceInterface)
	}

	/*// encode tl
	encBuf := bytes.NewBuffer(nil)
	err = binary.Write(encBuf, binary.BigEndian, i.Type)
	if err != nil {
		return
	}
	i.Length = uint16(encV.Len())
	err = binary.Write(encBuf, binary.BigEndian, i.Length)
	if err != nil {
		return
	}*/

	return encV.Bytes(), nil
}

func (i *IEUserPlaneIPResourceInformation) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	//5 	Spare	ASSOSI	ASSONI	TEIDRI	V6	V4
	tmpByte, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.V4, err = utils.GetBoolInByte(tmpByte, 1)
	if err != nil {
		return err
	}
	i.V6, err = utils.GetBoolInByte(tmpByte, 2)
	if err != nil {
		return err
	}

	v := utils.ByteOne(tmpByte)
	tmpTEIDRI, err := v.GetBits(3, 5)
	v1 := utils.ByteOne(tmpTEIDRI)
	i.TEIDRI, _ = v1.RightShift(3)
	if err != nil {
		return err
	}

	i.ASSONI, err = utils.GetBoolInByte(tmpByte, 6)
	if err != nil {
		return err
	}

	i.ASSOSI, err = utils.GetBoolInByte(tmpByte, 7)
	if err != nil {
		return err
	}

	//6  	TEID Range
	if i.TEIDRI != 0 {
		err = binary.Read(r, binary.BigEndian, &i.TEIDRange)
		if err != nil {
			return err
		}
	}

	if i.V4 {
		var ip [4]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv4address = ip[:]
	}
	if i.V6 {
		var ip [16]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv6address = ip[:]
	}

	//k to l	Network Instance
	if i.ASSONI && i.ASSOSI {
		tmpBytes := make([]byte, r.Len())
		_, err := r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.NetworkInstance = string(tmpBytes[:len(tmpBytes)-1])
		i.SourceInterface = tmpBytes[len(tmpBytes)-1] & 0x0f
	}

	if i.ASSONI && (!i.ASSOSI) {
		tmpBytes := make([]byte, r.Len())
		_, err := r.Read(tmpBytes)
		if err != nil {
			return err
		}
		i.NetworkInstance = string(tmpBytes)
	}

	if (!i.ASSONI) && i.ASSOSI {

		tmpbyte1, err := r.ReadByte()
		if err != nil {
			return err
		}
		i.SourceInterface = tmpbyte1 & 0x0f
	}

	return nil
}

func (i *IEUserPlaneIPResourceInformation) Len() int {
	return int(i.Length)
}

func (i *IEUserPlaneIPResourceInformation) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUserPlaneIPResourceInformation) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEUserPlaneIPResourceInformation) Set(v uint8) error {

	i.Type = IE_User_Plane_IP_Resource_Information
	return nil
}

/*func (i *IEUserPlaneIPResourceInformation) Get() (v uint8, e error) {
	panic("implement me")
}*/

//7.4.4.2	PFCP Association Setup Response
/*
Table 7.4.4.2-1: Information Elements in a PFCP Association Setup Response
Information elements	P
Node ID	M
Cause	M
Recovery Time Stamp	M
UP Function Features	C
CP Function Features	C
User Plane IP Resource Information 	O
*/
type PFCPAssociationSetupResponse struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPAssociationSetupResponse
}

func (p *PFCPAssociationSetupResponse) MarshalBinary() (data []byte, err error) {
	// IE 编码
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PFCPAssociationSetupResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPAssociationSetupResponse{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPAssociationSetupResponse) Len() int {
	return int(p.PfcpHeader.Length)
}

type IEPFCPAssociationSetupResponse struct {
	NodeID                         *IENodeID
	Cause                          *IECause
	RecoveryTimeStamp              *IERecoveryTimeStamp
	UPFunctionFeatures             *IEUPFunctionFeatures
	CPFunctionFeatures             *IECPFunctionFeatures
	UserPlaneIPResourceInformation []*IEUserPlaneIPResourceInformation

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFCPAssociationSetupResponse) Encode() (data []byte, err error) {
	//	encode ie
	// 编码 必选 IE ,TLV 值已经填充，
	encBuf := bytes.NewBuffer(nil)

	//Node ID	M
	ieEnc, err := i.NodeID.Encode()
	if err != nil {
		return
	}
	// TL 编码
	hd := IETypeLength{}
	hd.Type = uint16(IE_Node_ID)
	hd.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, hd)
	if err != nil {
		return
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return
	}
	//Cause	M
	ieEnc, err = i.Cause.Encode()
	if err != nil {
		return
	}
	// TL 编码
	hd.Type = uint16(IE_Cause)
	hd.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, hd)
	if err != nil {
		return
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return
	}

	//Recovery Time Stamp	M
	ieEnc, err = i.RecoveryTimeStamp.Encode()
	if err != nil {
		return
	}
	// TL 编码
	hd.Type = uint16(IE_Recovery_Time_Stamp)
	hd.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, hd)
	if err != nil {
		return
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return
	}
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		case IE_UP_Function_Features: //UP Function Features	C
			ieEnc, err = i.UPFunctionFeatures.Encode()
			if err != nil {
				return
			}
			/*_, err = encBuf.Write(ieEnc)
			if err != nil {
				return
			}*/
		case IE_CP_Function_Features: //CP Function Features	C
			ieEnc, err = i.CPFunctionFeatures.Encode()
			if err != nil {
				return
			}
			/*_, err = encBuf.Write(ieEnc)
			if err != nil {
				return
			}*/
		case IE_User_Plane_IP_Resource_Information: //User Plane IP Resource Information 	O
			for _, v := range i.UserPlaneIPResourceInformation[:len(i.UserPlaneIPResourceInformation)-1] {
				//	encode v
				vEnc, err := v.Encode()
				if err != nil {
					return nil, err
				}
				// encode tl
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
			// encode v
			ieEnc, err = i.UserPlaneIPResourceInformation[len(i.UserPlaneIPResourceInformation)-1].Encode()
			if err != nil {
				return nil, err
			}
			//ieEnc, err = i.UserPlaneIPResourceInformation.Encode()
			//if err != nil {
			//	return
			//}
			/*_, err = encBuf.Write(ieEnc)
			if err != nil {
				return
			}*/

		default:

		}
		// TL 编码
		hd := IETypeLength{}
		hd.Type = uint16(id)
		hd.Length = uint16(len(ieEnc))
		err = EncodeTl(encBuf, hd)
		if err != nil {
			return
		}
		_, err = encBuf.Write(ieEnc)
		if err != nil {
			return
		}
	}
	return encBuf.Bytes(), nil
}

func (i *IEPFCPAssociationSetupResponse) Decode(data []byte) error {
	// parse v
	i.Grouped = true
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo 必选IE检查
	return nil
}

func (i *IEPFCPAssociationSetupResponse) Len() int {
	return 0
}

func (i *IEPFCPAssociationSetupResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPAssociationSetupResponse) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IENodeID:
		i.NodeID = Ie
	case *IECause:
		i.Cause = Ie
	case *IERecoveryTimeStamp:
		i.RecoveryTimeStamp = Ie
	case *IEUPFunctionFeatures:
		i.UPFunctionFeatures = Ie
	case *IECPFunctionFeatures:
		i.CPFunctionFeatures = Ie
	case *IEUserPlaneIPResourceInformation:
		i.UserPlaneIPResourceInformation = append(i.UserPlaneIPResourceInformation, Ie)
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

//7.4.4.3	PFCP Association Update Request
type PFCPAssociationUpdateRequest struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPAssociationUpdateRequest
}

func (p *PFCPAssociationUpdateRequest) MarshalBinary() (data []byte, err error) {
	// IE 编码
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PFCPAssociationUpdateRequest) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPAssociationUpdateRequest{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPAssociationUpdateRequest) Len() int {
	return 0
}

/*
Table 7.4.4.3-1: Information Elements in a PFCP Association Update Request
Information elements	P
Node ID	M
UP Function Features	O
CP Function Features	O
PFCP Association Release Request	C
Graceful Release Period	C
User Plane IP Resource Information 	O
*/
type IEPFCPAssociationUpdateRequest struct {
	NodeID                         *IENodeID
	UPFunctionFeatures             *IEUPFunctionFeatures
	CPFunctionFeatures             *IECPFunctionFeatures
	PFCPAssociationReleaseRequest  *IEPFCPAssociationReleaseRequest
	GracefulReleasePeriod          *IEGracefulReleasePeriod
	UserPlaneIPResourceInformation []*IEUserPlaneIPResourceInformation

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFCPAssociationUpdateRequest) Encode() (data []byte, err error) {
	//	encode ie
	// 编码 必选 IE
	encBuf := bytes.NewBuffer(nil)

	//Node ID	M
	ieEnc, err := i.NodeID.Encode() // Encode v
	if err != nil {
		return nil, err
	}
	// TL 编码
	hd := IETypeLength{}
	hd.Type = uint16(IE_Node_ID)
	hd.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, hd)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}
	// optional
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		case IE_Offending_IE: //UP Function Features	C
			ieEnc, err = i.UPFunctionFeatures.Encode() // Encode v
			if err != nil {
				return nil, err
			}

		case IE_F_SEID: //CP Function Features	C
			ieEnc, err = i.CPFunctionFeatures.Encode() // Encode v
			if err != nil {
				return nil, err
			}

		case IE_IE_PFCP_Association_Release_Request:
			ieEnc, err = i.PFCPAssociationReleaseRequest.Encode() // Encode v
			if err != nil {
				return nil, err
			}

		case IE_Graceful_Release_Period:
			ieEnc, err = i.GracefulReleasePeriod.Encode() // Encode v
			if err != nil {
				return nil, err
			}

		case IE_User_Plane_IP_Resource_Information: //User Plane IP Resource Information 	O
			for _, v := range i.UserPlaneIPResourceInformation[:len(i.UserPlaneIPResourceInformation)-1] {
				//	encode v
				vEnc, err := v.Encode()
				if err != nil {
					return nil, err
				}
				// encode tl
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
			// encode v
			ieEnc, err = i.UserPlaneIPResourceInformation[len(i.UserPlaneIPResourceInformation)-1].Encode()
			if err != nil {
				return nil, err
			}
			/*ieEnc, err = i.UserPlaneIPResourceInformation.Encode() // Encode v
			if err != nil {
				return nil, err
			}*/

		default:

		}
		// TL 编码
		hd := IETypeLength{}
		hd.Type = uint16(id)
		hd.Length = uint16(len(ieEnc))
		err = EncodeTl(encBuf, hd)
		if err != nil {
			return nil, err
		}
		_, err = encBuf.Write(ieEnc)
		if err != nil {
			return nil, err
		}
	}
	return encBuf.Bytes(), nil
}

// 消息级IE列表解析
func (i *IEPFCPAssociationUpdateRequest) Decode(data []byte) error {
	// parse v
	i.Grouped = true
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo 必选IE检查
	return nil
}

func (i *IEPFCPAssociationUpdateRequest) Len() int {
	return 0
}

func (i *IEPFCPAssociationUpdateRequest) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPAssociationUpdateRequest) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IENodeID:
		i.NodeID = Ie
	case *IEUPFunctionFeatures:
		i.UPFunctionFeatures = Ie
	case *IECPFunctionFeatures:
		i.CPFunctionFeatures = Ie
	case *IEPFCPAssociationReleaseRequest:
		i.PFCPAssociationReleaseRequest = Ie
	case *IEGracefulReleasePeriod:
		i.GracefulReleasePeriod = Ie
	case *IEUserPlaneIPResourceInformation:
		i.UserPlaneIPResourceInformation = append(i.UserPlaneIPResourceInformation, Ie)
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

//8.2.78	Graceful Release Period
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 112 (decimal)
	3 to 4	Length = n
	5	Timer unit	Timer value
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.78-1: Graceful Release Period*/
type IEGracefulReleasePeriod struct {
	IETypeLength
	TimerUnit  uint8 // 0x1110 0000
	TimerValue uint8 // 0x0001 1111
}

func (i *IEGracefulReleasePeriod) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	var tmpByte uint8
	tmpByte = (i.TimerUnit << 5) + i.TimerValue
	err = binary.Write(encBuf, binary.BigEndian, tmpByte)
	if err != nil {
		return
	}

	return encBuf.Bytes(), nil
}

func (i *IEGracefulReleasePeriod) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)
	var tmpByte byte
	err := binary.Read(r, binary.BigEndian, &tmpByte)
	if err != nil {
		return err
	}
	v := utils.ByteOne(tmpByte)
	i.TimerUnit, err = v.GetBits(6, 8)
	v1 := utils.ByteOne(i.TimerUnit)
	i.TimerUnit, _ = v1.RightShift(6)
	if err != nil {
		return err
	}
	i.TimerValue, err = v.GetBits(1, 5)
	if err != nil {
		return err
	}
	return nil
}

func (i *IEGracefulReleasePeriod) Len() int {
	return int(i.Length)
}

func (i *IEGracefulReleasePeriod) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEGracefulReleasePeriod) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEGracefulReleasePeriod) Set(v uint8) error {
	i.Type = IE_Graceful_Release_Period
	return nil
}

//Table 8.2.78.1: Graceful Release Period information element
/*Timer unit
Bits 6 to 8 defines the timer value unit for the timer as follows:
Bits
8 7 6
0 0 0  value is incremented in multiples of 2 seconds
0 0 1  value is incremented in multiples of 1 minute
0 1 0  value is incremented in multiples of 10 minutes
0 1 1  value is incremented in multiples of 1 hour
1 0 0  value is incremented in multiples of 10 hours
1 1 1  value indicates that the timer is infinite*/
const (
	TimerUnit_2_Seconds  = time.Second * 2
	TimerUnit_1_Minute   = time.Minute
	TimerUnit_10_Minutes = time.Minute * 10
	TimerUnit_1_Hour     = time.Hour
	TimerUnit_10_Hours   = time.Hour * 10
)

// IE
//8.2.77	PFCP Association Release Request
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 111 (decimal)
	3 to 4	Length = n
	5	Spare	SARR
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.77-1: PFCP Association Release Request*/
type IEPFCPAssociationReleaseRequest struct {
	IETypeLength
	SARR bool
}

func (i *IEPFCPAssociationReleaseRequest) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)
	/*err = binary.Write(encBuf, binary.BigEndian, i.Type)
	if err != nil {
		return
	}
	err = binary.Write(encBuf, binary.BigEndian, i.Length)
	if err != nil {
		return
	}*/
	err = binary.Write(encBuf, binary.BigEndian, i.SARR)
	if err != nil {
		return
	}
	return encBuf.Bytes(), nil
}

func (i *IEPFCPAssociationReleaseRequest) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)
	var tmpByte byte
	err := binary.Read(r, binary.BigEndian, &tmpByte)
	if err != nil {
		return err
	}
	i.SARR, _ = utils.GetBoolInByte(tmpByte, 1)

	return nil
}

func (i *IEPFCPAssociationReleaseRequest) Len() int {
	return int(i.Length)
}

func (i *IEPFCPAssociationReleaseRequest) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEPFCPAssociationReleaseRequest) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPAssociationReleaseRequest) Set(v uint8) error {
	i.Type = IE_IE_PFCP_Association_Release_Request
	return nil
}

//7.4.4.4	PFCP Association Update Response
type PFCPAssociationUpdateResponse struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPAssociationUpdateResponse
}

func (p *PFCPAssociationUpdateResponse) MarshalBinary() (data []byte, err error) {
	// encode IE
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PFCPAssociationUpdateResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPAssociationUpdateResponse{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPAssociationUpdateResponse) Len() int {
	return 0
}

/*
Table 7.4.4.4-1: Information Elements in a PFCP Association Update Response
Information elements	P
Node ID	M
Cause	M
UP Function Features	O
CP Function Features	O
*/
type IEPFCPAssociationUpdateResponse struct {
	NodeID             *IENodeID
	Cause              *IECause
	UPFunctionFeatures *IEUPFunctionFeatures
	CPFunctionFeatures *IECPFunctionFeatures

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFCPAssociationUpdateResponse) Encode() (data []byte, err error) {
	// encode ie
	// Mandatory ie
	encBuf := bytes.NewBuffer(nil)

	// Node ID	M
	ieEnc, err := i.NodeID.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Node_ID)
	tl.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}

	//Cause	M
	ieEnc, err = i.Cause.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl = IETypeLength{}
	tl.Type = uint16(IE_Cause)
	tl.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}

	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//UP Function Features	O
		case IE_UP_Function_Features:
			ieEnc, err = i.UPFunctionFeatures.Encode()
			if err != nil {
				return
			}
			/*_, err = encBuf.Write(ieEnc)
			if err != nil {
				return
			}*/
			//CP Function Features	O
		case IE_CP_Function_Features:
			ieEnc, err = i.CPFunctionFeatures.Encode()
			if err != nil {
				return
			}
			/*_, err = encBuf.Write(ieEnc)
			if err != nil {
				return
			}*/
		default:

		}
		// TL 编码
		hd := IETypeLength{}
		hd.Type = uint16(id)
		hd.Length = uint16(len(ieEnc))
		err = EncodeTl(encBuf, hd)
		if err != nil {
			return
		}
		_, err = encBuf.Write(ieEnc)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEPFCPAssociationUpdateResponse) Decode(data []byte) error {
	// parse IE
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEPFCPAssociationUpdateResponse) Len() int {
	return 0
}

func (i *IEPFCPAssociationUpdateResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPAssociationUpdateResponse) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IENodeID:
		i.NodeID = Ie
	case *IECause:
		i.Cause = Ie
	case *IEUPFunctionFeatures:
		i.UPFunctionFeatures = Ie
	case *IECPFunctionFeatures:
		i.CPFunctionFeatures = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

//7.4.4.5	PFCP Association Release Request
type PFCPAssociationReleaseRequest struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEGroupedPFCPAssociationReleaseRequest
}

func (p *PFCPAssociationReleaseRequest) MarshalBinary() (data []byte, err error) {
	// encode IE
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PFCPAssociationReleaseRequest) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEGroupedPFCPAssociationReleaseRequest{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPAssociationReleaseRequest) Len() int {
	return int(p.PfcpHeader.Length)
}

/*Table 7.4.4.5-1: Information Elements in a PFCP Association Release Request
Information elements	P
Node ID	M
*/
type IEGroupedPFCPAssociationReleaseRequest struct {
	NodeID *IENodeID

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEGroupedPFCPAssociationReleaseRequest) Encode() (data []byte, err error) {
	// encode ie
	// Mandatory ie
	encBuf := bytes.NewBuffer(nil)

	//Node ID	M
	ieEnc, err := i.NodeID.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Node_ID)
	tl.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (i *IEGroupedPFCPAssociationReleaseRequest) Decode(data []byte) error {
	// parse IE
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEGroupedPFCPAssociationReleaseRequest) Len() int {
	return 0
}

func (i *IEGroupedPFCPAssociationReleaseRequest) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEGroupedPFCPAssociationReleaseRequest) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IENodeID:
		i.NodeID = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

//7.4.4.6	PFCP Association Release Response
type PFCPAssociationReleaseResponse struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPAssociationReleaseResponse
}

func (p *PFCPAssociationReleaseResponse) MarshalBinary() (data []byte, err error) {
	// encode IE
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PFCPAssociationReleaseResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPAssociationReleaseResponse{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPAssociationReleaseResponse) Len() int {
	return int(p.PfcpHeader.Length)
}

/*Table 7.4.4.6-1: Information Elements in a PFCP Association Release Response
Information elements	P
Node ID	M
Cause	M
*/
type IEPFCPAssociationReleaseResponse struct {
	NodeID *IENodeID
	Cause  *IECause

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFCPAssociationReleaseResponse) Encode() (data []byte, err error) {
	// encode ie
	// Mandatory ie
	encBuf := bytes.NewBuffer(nil)

	//Node ID	M
	vEnc, err := i.NodeID.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Node_ID)
	tl.Length = uint16(len(vEnc))
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
		return nil, err
	}
	// encode Tl
	tl = IETypeLength{}
	tl.Type = uint16(IE_Cause)
	tl.Length = uint16(len(vEnc))
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

func (i *IEPFCPAssociationReleaseResponse) Decode(data []byte) error {
	// parse IE
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo check mandatory ie
	return nil
}

func (i *IEPFCPAssociationReleaseResponse) Len() int {
	return 0
}

func (i *IEPFCPAssociationReleaseResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPAssociationReleaseResponse) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IENodeID:
		i.NodeID = Ie
	case *IECause:
		i.Cause = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

// 7.4.4.7	PFCP Version Not Supported Response
type PFCPVersionNotSupportedResponse struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPVersionNotSupportedResponse
}

func (p *PFCPVersionNotSupportedResponse) MarshalBinary() (data []byte, err error) {
	// encode IE
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PFCPVersionNotSupportedResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPVersionNotSupportedResponse{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPVersionNotSupportedResponse) Len() int {
	return int(p.PfcpHeader.Length)
}

type IEPFCPVersionNotSupportedResponse struct {
}

func (i *IEPFCPVersionNotSupportedResponse) Encode() (data []byte, err error) {
	return data, nil
}

func (i *IEPFCPVersionNotSupportedResponse) Decode(data []byte) error {
	return nil
}

func (i *IEPFCPVersionNotSupportedResponse) Len() int {
	return 0
}

func (i *IEPFCPVersionNotSupportedResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPVersionNotSupportedResponse) SetObjectToParent(child interface{}) error {
	return nil
}

//7.4.5	PFCP Node Report Procedure
//7.4.5.1	PFCP Node Report Request
type PFCPNodeReportRequest struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPNodeReportRequest
}

func (p *PFCPNodeReportRequest) MarshalBinary() (data []byte, err error) {
	// encode IE
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (p *PFCPNodeReportRequest) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPNodeReportRequest{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPNodeReportRequest) Len() int {
	return int(p.PfcpHeader.Length)
}

/*Table 7.4.5.1.1-1: Information Elements in PFCP Node Report Request
Information elements	P

Node ID	M
Node Report Type	M
User Plane Path Failure Report	C
*/
type IEPFCPNodeReportRequest struct {
	NodeID                     *IENodeID
	NodeReportType             *IENodeReportType
	UserPlanePathFailureReport *IEUserPlanePathFailureReport

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFCPNodeReportRequest) Encode() (data []byte, err error) {
	// encode ie
	// Mandatory ie
	encBuf := bytes.NewBuffer(nil)

	// Node ID	M
	vEnc, err := i.NodeID.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Node_ID)
	tl.Length = uint16(len(vEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(vEnc)
	if err != nil {
		return nil, err
	}
	// Node Report Type	M
	vEnc, err = i.NodeReportType.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl = IETypeLength{}
	tl.Type = uint16(IE_Node_Report_Type)
	tl.Length = uint16(len(vEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(vEnc)
	if err != nil {
		return nil, err
	}

	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		//UP Function Features	O
		case IE_User_Plane_Path_Failure_Report:
			vEnc, err = i.UserPlanePathFailureReport.Encode()
			if err != nil {
				return
			}
		default:
		}
		// TL 编码
		hd := IETypeLength{}
		hd.Type = uint16(id)
		hd.Length = uint16(len(vEnc))
		err = EncodeTl(encBuf, hd)
		if err != nil {
			return
		}
		_, err = encBuf.Write(vEnc)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEPFCPNodeReportRequest) Decode(data []byte) error {
	// parse IE
	i.Grouped = true
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo 必选IE检查
	return nil
}

func (i *IEPFCPNodeReportRequest) Len() int {
	return 0
}

func (i *IEPFCPNodeReportRequest) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPNodeReportRequest) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	//Node ID	M
	case *IENodeID:
		i.NodeID = Ie
		//Node Report Type	M
	case *IENodeReportType:
		i.NodeReportType = Ie
		//User Plane Path Failure Report	C
	case *IEUserPlanePathFailureReport:
		i.UserPlanePathFailureReport = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

// 7.4.5.1.2	User Plane Path Failure Report IE within PFCP Node Report Request
/*Table 7.4.5.1.2-1: User Plane Path Failure Report IE within PFCP Node Report Request
Octet 1 and 2			User Plane Path Failure Report IE Type = 102 (decimal)
Octets 3 and 4			Length = n

Information elements	P

Remote GTP-U Peer  	M
*/
type IEUserPlanePathFailureReport struct {
	IETypeLength
	RemoteGTPUPeer *IERemoteGTPUPeer

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEUserPlanePathFailureReport) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	ieEnc, err := i.RemoteGTPUPeer.Encode()
	if err != nil {
		return nil, err
	}
	// TL 编码
	tl := IETypeLength{}
	tl.Type = uint16(IE_Remote_GTPU_Peer)
	tl.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}

	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

// grouped IE parse
func (i *IEUserPlanePathFailureReport) Decode(data []byte) error {
	// parse v
	// parse IE
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo 必选IE检查
	return nil
}

func (i *IEUserPlanePathFailureReport) Len() int {
	return int(i.Length)
}

func (i *IEUserPlanePathFailureReport) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IEUserPlanePathFailureReport) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IERemoteGTPUPeer:
		i.RemoteGTPUPeer = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

func (i *IEUserPlanePathFailureReport) Set(v uint8) error {
	i.Type = IE_User_Plane_Path_Failure_Report
	return nil
}

//8.2.70	Remote GTP-U Peer
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 103 (decimal)
	3 to 4	Length = n
	5	Spare	V4	V6
	m to (m+3)	IPv4 address
	p to (p+15)	IPv6 address
	k to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.70-1: Remote GTP-U Peer*/
type IERemoteGTPUPeer struct {
	IETypeLength
	V6          bool
	V4          bool
	IPv4address net.IP
	IPv6address net.IP
}

func (i *IERemoteGTPUPeer) Encode() (data []byte, err error) {
	encV := bytes.NewBuffer(nil)

	// 5	Spare	V4	V6
	var tmpByte byte
	tmpByte = utils.BoolToUint8(i.V6) +
		(utils.BoolToUint8(i.V4) << 1)

	err = encV.WriteByte(tmpByte)
	if err != nil {
		return nil, err
	}

	// m to (m+3)	IPv4 address
	if i.V4 {
		err = binary.Write(encV, binary.BigEndian, []byte(i.IPv4address.To4()))
		if err != nil {
			return nil, err
		}
	}
	// p to (p+15)	IPv6 address
	if i.V6 {
		err = binary.Write(encV, binary.BigEndian, []byte(i.IPv6address.To16()))
		if err != nil {
			return
		}
	}

	return encV.Bytes(), nil
}

func (i *IERemoteGTPUPeer) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	//5	Spare	V4	V6
	tmpByte, err := r.ReadByte()
	if err != nil {
		return err
	}
	i.V4, err = utils.GetBoolInByte(tmpByte, 2)
	if err != nil {
		return err
	}
	i.V6, err = utils.GetBoolInByte(tmpByte, 1)
	if err != nil {
		return err
	}

	if i.V4 {
		var ip [4]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv4address = ip[:]
	}

	if i.V6 {
		var ip [16]byte
		_, err := r.Read(ip[:])
		if err != nil {
			return err
		}
		i.IPv6address = ip[:]
	}

	return nil
}

func (i *IERemoteGTPUPeer) Len() int {
	return int(i.Length)
}

func (i *IERemoteGTPUPeer) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IERemoteGTPUPeer) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IERemoteGTPUPeer) Set(v uint8) error {
	i.Type = IE_Remote_GTPU_Peer
	return nil
}

//8.2.69	Node Report Type
/*		Bits
	Octets	8	7	6	5	4	3	2	1
	1 to 2	Type = 101 (decimal)
	3 to 4	Length = n
	5	Spare	UPFR
	6 to (n+4)	These octet(s) is/are present only if explicitly specified
Figure 8.2.69-1: Node Report Type*/
type IENodeReportType struct {
	IETypeLength
	UPFR bool
}

func (i *IENodeReportType) Encode() (data []byte, err error) {
	// encode v
	encBuf := bytes.NewBuffer(nil)

	err = binary.Write(encBuf, binary.BigEndian, i.UPFR)
	if err != nil {
		return nil, err
	}
	return encBuf.Bytes(), nil
}

func (i *IENodeReportType) Decode(data []byte) error {
	// parse v
	r := bytes.NewReader(data)

	err := binary.Read(r, binary.BigEndian, &i.UPFR)
	if err != nil {
		return err
	}

	return nil
}

func (i *IENodeReportType) Len() int {
	return int(i.Length)
}

func (i *IENodeReportType) SetObject(t uint16, l uint16) error {
	i.Type = t
	i.Length = l
	return nil
}

func (i *IENodeReportType) SetObjectToParent(child interface{}) error {
	return fmt.Errorf("implement me")
}

func (i *IENodeReportType) Set(v uint8) error {
	i.Type = IE_Node_Report_Type
	return nil
}

//7.4.5.2	PFCP Node Report Response
type PFCPNodeReportResponse struct {
	PfcpHeader PfcpHeaderforNode
	IE         *IEPFCPNodeReportResponse
}

func (p *PFCPNodeReportResponse) MarshalBinary() (data []byte, err error) {
	// encode IE
	data, err = p.IE.Encode()
	if err != nil {
		return nil, err
	}
	return data, err
}

func (p *PFCPNodeReportResponse) UnmarshalBinary(data []byte) error {
	//	Grouped IE parse
	p.IE = &IEPFCPNodeReportResponse{}
	err := p.IE.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (p *PFCPNodeReportResponse) Len() int {
	return 0
}

/*Table 7.4.5.2.1-1: Information Elements in PFCP Node Report Response
Information elements	P

Node ID	M
Cause	M
Offending IE	C
*/
type IEPFCPNodeReportResponse struct {
	NodeID      *IENodeID
	Cause       *IECause
	OffendingIE *IEOffendingIE

	// add
	Grouped bool
	// Mandatory
	M []uint16
	// presence
	IeFlags bitset.BitSet
}

func (i *IEPFCPNodeReportResponse) Encode() (data []byte, err error) {
	// encode ie
	// Mandatory ie
	encBuf := bytes.NewBuffer(nil)

	//Node ID	M
	ieEnc, err := i.NodeID.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl := IETypeLength{}
	tl.Type = uint16(IE_Node_ID)
	tl.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return
	}

	//Cause	M
	ieEnc, err = i.Cause.Encode()
	if err != nil {
		return nil, err
	}
	// encode Tl
	tl = IETypeLength{}
	tl.Type = uint16(IE_Cause)
	tl.Length = uint16(len(ieEnc))
	err = EncodeTl(encBuf, tl)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(ieEnc)
	if err != nil {
		return
	}

	// optional ie
	for id, e := i.IeFlags.NextSet(0); e; id, e = i.IeFlags.NextSet(id + 1) {
		switch id {
		// Offending IE	C
		case IE_Offending_IE:
			ieEnc, err = i.OffendingIE.Encode()
			if err != nil {
				return
			}
		default:

		}
		// TL 编码
		hd := IETypeLength{}
		hd.Type = uint16(id)
		hd.Length = uint16(len(ieEnc))
		err = EncodeTl(encBuf, hd)
		if err != nil {
			return
		}
		_, err = encBuf.Write(ieEnc)
		if err != nil {
			return
		}
	}

	return encBuf.Bytes(), nil
}

func (i *IEPFCPNodeReportResponse) Decode(data []byte) error {
	// parse IE
	i.Grouped = true
	err := nodeGroupedIeParse(i, data)
	if err != nil {
		return err
	}
	// todo 必选IE检查
	return nil
}

func (i *IEPFCPNodeReportResponse) Len() int {
	return 0
}

func (i *IEPFCPNodeReportResponse) SetObject(t uint16, l uint16) error {
	return fmt.Errorf("implement me")
}

func (i *IEPFCPNodeReportResponse) SetObjectToParent(child interface{}) error {
	switch Ie := child.(type) {
	case *IENodeID:
		i.NodeID = Ie
	case *IECause:
		i.Cause = Ie
	case *IEOffendingIE:
		i.OffendingIE = Ie
	default:
		return errors.New("Illegal IE")
	}

	return nil
}

// 3080
