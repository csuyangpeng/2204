package nasmsg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

//24501 Table 8.2.17.1.1
type ServiceAcceptMsg struct {

	//optional
	//50	PDU session status	PDU session status	9.11.3.44	O	TLV
	PDUSessionStatus nasie.SessionStatus
	//26	PDU session reactivation result	PDU session reactivation result	9.11.3.42	O	TLV
	PDUSessReactiveResult nasie.SessionStatus
	//7E	PDU session reactivation result error cause	9.11.3.43	O	TLV-E
	PDUSessReactiveResultErr nasie.SessionReactErrCause
	//78	EAP message	EAP message	9.11.2.2	O	TLV-E
	EAPMsg []byte

	//Indicates whether an IE is assigned or it is an empty value
	// Ie flags
	IeFlags bitset.BitSet
}

//type IeId uint
const (
	IeidServiceacptPdusessionstatus uint = iota
	IeidServiceacptPdusessreactiveresult
	IeidServiceacptPdusessreactiveresulterr
	IeidServiceacptEapmsg
)

func NewServiceAcceptMsg() *ServiceAcceptMsg {
	return &ServiceAcceptMsg{}
}

func (p ServiceAcceptMsg) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	// mandatory IEs

	//header
	header := nas.MmNasMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = nas.PlainNasMsg
	header.MessageType = nas.ServiceAccept
	encBuf = append(encBuf, header.Encode()...)

	//Optional IEs
	optIeOctet := p.EncodeOptIes()
	encBuf = append(encBuf, optIeOctet...)

	return encBuf, nil
}

func (p ServiceAcceptMsg) EncodeOptIes() []byte {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	var encBuf []byte
	// for other optional IEs:
	for i, e := p.IeFlags.NextSet(0); e; i, e = p.IeFlags.NextSet(i + 1) {
		switch i {
		case IeidServiceacptPdusessionstatus:
			typeOctet := byte(nasie.IeiPDUSessionStatus)
			uplinkDataStatusByte := p.PDUSessionStatus.Encode()
			// T
			encBuf = append(encBuf, typeOctet)
			// L
			encBuf = append(encBuf, byte(len(uplinkDataStatusByte)))
			// V
			encBuf = append(encBuf, uplinkDataStatusByte...)
		case IeidServiceacptPdusessreactiveresult:
			typeOctet := byte(nasie.IeiPDUSessionReactivationResult)
			pduSessReactiveResultByte := p.PDUSessReactiveResult.Encode()
			// T
			encBuf = append(encBuf, typeOctet)
			// L
			encBuf = append(encBuf, byte(len(pduSessReactiveResultByte)))
			// V
			encBuf = append(encBuf, pduSessReactiveResultByte...)
		case IeidServiceacptPdusessreactiveresulterr:
			typeOctet := byte(nasie.IeiPDUSessionReactivationResultErrorCause)
			pduSessReactiveResultErrByte := p.PDUSessReactiveResultErr.Encode()
			// T
			encBuf = append(encBuf, typeOctet)
			// L
			errLen := make([]byte, 2)
			binary.BigEndian.PutUint16(errLen, uint16(len(pduSessReactiveResultErrByte)))
			encBuf = append(encBuf, errLen...)
			// V
			encBuf = append(encBuf, pduSessReactiveResultErrByte...)
		case IeidServiceacptEapmsg:
			// T
			encBuf = append(encBuf, byte(nasie.IeiEAPMessage))
			// L
			eapLen := make([]byte, 2)
			binary.BigEndian.PutUint16(eapLen, uint16(len(p.EAPMsg)))
			encBuf = append(encBuf, eapLen...)
			// V
			encBuf = append(encBuf, p.EAPMsg...)
		}
	}
	return encBuf
}

func (p *ServiceAcceptMsg) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "bytes need to decode:", msgBuf)

	// Optional IEs
	err := p.DecodeOptIes(msgBuf)
	if err != nil {
		return fmt.Errorf("failed to decode optional ies")
	}
	return nil
}

func (p *ServiceAcceptMsg) DecodeOptIes(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleCmnMsg, nil)
	for {
		ieType, err := msgBuf.ReadByte()
		if err != nil {
			rlogger.Trace(types.ModuleCmnMsg, rlogger.DEBUG, nil, "no more bytes")
			return nil
		}
		switch nasie.Iei(ieType) {
		//PDU session status
		case nasie.IeiPDUSessionStatus:
			err = p.PDUSessionStatus.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode PDU Session Status failed")
				return nas.ErrDecodeNasIeFailed
			}
		case nasie.IeiPDUSessionReactivationResult:
			err = p.PDUSessReactiveResult.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode PDU Session Reactive Result failed")
				return nas.ErrDecodeNasIeFailed
			}
		case nasie.IeiPDUSessionReactivationResultErrorCause:
			err = p.PDUSessReactiveResultErr.Decode(msgBuf)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode PDU Session Reactive Result err failed")
				return nas.ErrDecodeNasIeFailed
			}
		case nasie.IeiEAPMessage:
			//L
			lenBytes := make([]byte, 2)
			err := binary.Read(msgBuf, binary.BigEndian, &lenBytes)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode EAP Message len failed")
				return nas.ErrDecodeNasIeFailed
			}
			eapLen := binary.BigEndian.Uint16(lenBytes)
			//V
			p.EAPMsg = make([]byte, eapLen)
			err = binary.Read(msgBuf, binary.BigEndian, &p.EAPMsg)
			if err != nil {
				rlogger.Trace(types.ModuleCmnMsg, rlogger.ERROR, nil, "Decode EAP Message failed")
				return nas.ErrDecodeNasIeFailed
			}
		}
	}
}
