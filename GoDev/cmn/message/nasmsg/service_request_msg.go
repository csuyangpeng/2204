package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"

	"github.com/willf/bitset"
)

//24501-f10   8.2.16 (2018-09)
type ServiceRequestMsg struct {
	//mandatory
	//ngKSI 	NAS key set identifier	9.11.3.32	M	V
	NgKSI nasie.NasKSI
	//Service type	Service type	9.11.3.50	M	V
	ServiceType nas.ServiceTYPE
	//5G-S-TMSI	5GS mobile identity	9.11.3.4	M	LV
	MobileIdentity nasie.MobileIdentity

	//optional
	//40	Uplink data status	Uplink data status	9.11.3.57	O	TLV
	UplinkDataStatus nasie.SessionStatus
	//50	PDU session status	PDU session status	9.11.3.44	O	TLV
	PDUSessionStatus nasie.SessionStatus
	//25	Allowed PDU session status	Allowed PDU session status	9.11.3.13	O	TLV
	AllowedPDUSessionStatus nasie.SessionStatus

	NasMsgContainer []byte
	//Indicates whether an IE is assigned or it is an empty value
	// Ie flags
	IeFlags bitset.BitSet
}

func (p *ServiceRequestMsg) Reset() {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	p.NgKSI.Reset()
	p.ServiceType = 0
	p.MobileIdentity.Reset()
	p.UplinkDataStatus.Reset()
	p.PDUSessionStatus.Reset()
	p.AllowedPDUSessionStatus.Reset()
	p.IeFlags = bitset.BitSet{}
}

//type IeId uint
const (
	IeidServicereqUplinkdatastatus uint = iota
	IeidServicereqPdusessionstatus
	IeidServicereqAllowedpdusessioinstaus
	IeidServicereqNasmsgcontainer
)

func (p ServiceRequestMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	// mandatory IEs

	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.ServiceRequest
	encBuf = append(encBuf, header.Encode()...)

	//ngKSI && ServiceType
	ngKSISer := byte(p.ServiceType) << 4
	ngKSISer |= byte(utils.BoolToByte(p.NgKSI.Tsc)) << 3
	ngKSISer |= p.NgKSI.Ksi
	encBuf = append(encBuf, ngKSISer)

	//MobileIdentity, it's format is LV
	var mobileId []byte
	mobOctetValue, err := p.MobileIdentity.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "failed to encode mobile id, err(%s)", err)
		return nil, err
	}
	// L
	lenMobID := make([]byte, 2)
	binary.BigEndian.PutUint16(lenMobID, uint16(len(mobOctetValue)))
	mobileId = append(mobileId, lenMobID...)
	// V
	mobileId = append(mobileId, mobOctetValue...)
	encBuf = append(encBuf, mobileId...)

	//Optional IEs
	optIeOctet := p.EncodeOptIes()
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p ServiceRequestMsg) EncodeOptIes() []byte {
	var encBuf []byte
	// for other optional IEs:
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case IeidServicereqUplinkdatastatus:
			typeOctet := byte(nasie.IeiUplinkDataStatus)
			uplinkDataStatusByte := p.UplinkDataStatus.Encode()
			// T
			encBuf = append(encBuf, typeOctet)
			// L
			encBuf = append(encBuf, byte(len(uplinkDataStatusByte)))
			// V
			encBuf = append(encBuf, uplinkDataStatusByte...)
		case IeidServicereqPdusessionstatus:
			typeOctet := byte(nasie.IeiPduSessionStatus)
			pduSessionStatusByte := p.PDUSessionStatus.Encode()
			// T
			encBuf = append(encBuf, typeOctet)
			// L
			encBuf = append(encBuf, byte(len(pduSessionStatusByte)))
			// V
			encBuf = append(encBuf, pduSessionStatusByte...)
		case IeidServicereqAllowedpdusessioinstaus:
			typeOctet := byte(nasie.IeiAllowedPduSessStatus)
			allowedPDUSessioinStausByte := p.AllowedPDUSessionStatus.Encode()
			// T
			encBuf = append(encBuf, typeOctet)
			// L
			encBuf = append(encBuf, byte(len(allowedPDUSessioinStausByte)))
			// V
			encBuf = append(encBuf, allowedPDUSessioinStausByte...)
		}
	}
	return encBuf
}

func (p *ServiceRequestMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)

	//ngKSI && Service type
	ksiSerTypeOctet, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read ngKSI && Service type type")
	}
	serviceType, _ := utils.GetBitsValue(ksiSerTypeOctet, 5, 8)
	p.ServiceType = nas.ServiceTYPE(serviceType >> 4)
	p.NgKSI.Tsc, _ = utils.GetBitValue(ksiSerTypeOctet, 4)
	p.NgKSI.Ksi, _ = utils.GetBitsValue(ksiSerTypeOctet, 1, 3)

	//5G-S-TMSI
	err = p.MobileIdentity.Decode(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode mobile identity : %s", err)
	}

	// Optional IEs
	err = p.DecodeOptIes(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode optional ies: %s", err)
	}

	return nil
}

func (p *ServiceRequestMsg) DecodeOptIes(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	for {
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}
		switch nasie.Iei(ieType) {
		//Uplink data status
		case nasie.IeiUplinkDataStatus:
			err = p.UplinkDataStatus.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode upLink Data status failed")
				return nas.ErrDecodeNasIeFailed
			}
			p.IeFlags.Set(IeidServicereqUplinkdatastatus)
			//PDU session status
		case nasie.IeiPDUSessionStatus:
			err = p.PDUSessionStatus.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode PDU Session Status failed")
				return nas.ErrDecodeNasIeFailed
			}
			p.IeFlags.Set(IeidServicereqPdusessionstatus)
			//Allowed PDU session status
		case nasie.IeiAllowedPduSessStatus:
			err = p.AllowedPDUSessionStatus.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode Allowed PDU session Status failed")
				return nas.ErrDecodeNasIeFailed
			}
			p.IeFlags.Set(IeidServicereqAllowedpdusessioinstaus)
		case nasie.IeiNasMsgContainer:
			//length
			lengthBytes := make([]byte, 2)
			binary.Read(msgBuf, binary.BigEndian, lengthBytes)
			//value
			length := binary.BigEndian.Uint16(lengthBytes)
			leftBytes := make([]byte, length)
			binary.Read(msgBuf, binary.BigEndian, leftBytes)

			p.NasMsgContainer = leftBytes
			p.IeFlags.Set(IeidServicereqNasmsgcontainer)
		}
	}
}
