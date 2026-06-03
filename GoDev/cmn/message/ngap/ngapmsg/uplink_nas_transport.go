package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"unsafe"
)

// UplinkNasTransportMessage struct definition
type UplinkNasTransportMessage struct {
	RanUeNgapId uint32
	AmfUeNgapId uint64

	NasPdu      []byte
	UserLocInfo types3gpp.UserLocationInfo
	ctxt        codec.NgapOssCtxt
}

// NewUplinkNasTransportMessage create a new Message
func NewUplinkNasTransportMessage() *UplinkNasTransportMessage {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	return &UplinkNasTransportMessage{}
}

// SetOssCodecCtxt set OSS codec context
func (p *UplinkNasTransportMessage) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// Encode the message into bytes
func (p *UplinkNasTransportMessage) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	ulNasTransportMsg := codec.NewUplinkNASTransportCodec()
	defer codec.DeleteUplinkNASTransportCodec(ulNasTransportMsg)

	// ran ue ngap id
	ulNasTransportMsg.SetRanUeNgapId(uint(p.RanUeNgapId))
	// amf ue ngap id
	ulNasTransportMsg.SetAmfUeNgapId(uint64(p.AmfUeNgapId))
	// nas message
	ulNasTransportMsg.SetNasPdu(string(p.NasPdu))

	// user location info
	userLocationInfo := codec.NewUserLocationInfo()
	defer codec.DeleteUserLocationInfo(userLocationInfo)

	nrcgi := codec.NewNrCgi()
	defer codec.DeleteNrCgi(nrcgi)
	plmn := p.UserLocInfo.NrCgi.PlmnID.GetValue(types3gpp.BigEndian)
	nrcgi.SetPlmn(&plmn[0])
	nrCellId := p.UserLocInfo.NrCgi.NrCellId
	nrcgi.SetNrCellId(&nrCellId[0])
	userLocationInfo.SetNRCgi(nrcgi)

	if p.UserLocInfo.IsPsCellInfoPrst {
		userLocationInfo.SetPsCellInfoPrst(true)
		psCellInfo := codec.NewNrCgi()
		defer codec.DeleteNrCgi(psCellInfo)
		plmn1 := p.UserLocInfo.PsCellInfo.PlmnID.GetValue(types3gpp.BigEndian)
		psCellInfo.SetPlmn(&plmn1[0])
		nrCellId1 := p.UserLocInfo.PsCellInfo.NrCellId
		psCellInfo.SetNrCellId(&nrCellId1[0])
		userLocationInfo.SetPsCellInfo(psCellInfo)
	}

	plmn = p.UserLocInfo.Tai.Plmn.GetValue(types3gpp.BigEndian)
	tai := codec.NewTaiType()
	defer codec.DeleteTaiType(tai)
	tai.SetPlmn(&plmn[0])
	tac := p.UserLocInfo.Tai.Tac
	tai.SetTac(&tac[0])
	userLocationInfo.SetTai(tai)
	if p.UserLocInfo.IsTimeStampPrst {
		userLocationInfo.SetIsTimeStampPrst(true)
		userLocationInfo.SetTimeStamp(&(p.UserLocInfo.TimeStmp[0]))
	}
	ulNasTransportMsg.SetUserLocationInfo(userLocationInfo)

	//encode
	msgBuffer := ulNasTransportMsg.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

// Decode the bytes into message
func (p *UplinkNasTransportMessage) Decode(msgbuf []byte) error {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	ulNasTransportMsg := codec.NewUplinkNASTransportCodec()
	defer codec.DeleteUplinkNASTransportCodec(ulNasTransportMsg)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if ulNasTransportMsg.Decode(p.ctxt, msgBuffer) == true {
		p.RanUeNgapId = uint32(ulNasTransportMsg.GetRanUeNgapId())
		p.AmfUeNgapId = ulNasTransportMsg.GetAmfUeNgapId()
		p.NasPdu = []byte(ulNasTransportMsg.GetNasPdu())

		p.UserLocInfo.NrCgi.PlmnID.SetValue(
			utils.Conv2ByteSlice(ulNasTransportMsg.GetUserLocationInfo().GetNRCgi().GetPlmn(),
				types3gpp.SizeofPlmnID),
			types3gpp.BigEndian)
		copy(p.UserLocInfo.NrCgi.NrCellId[:],
			utils.Conv2ByteSlice(ulNasTransportMsg.GetUserLocationInfo().GetNRCgi().GetNrCellId(),
				types3gpp.SizeofNrCellID))
		if ulNasTransportMsg.GetUserLocationInfo().GetIsTimeStampPrst() {
			copy(p.UserLocInfo.TimeStmp[:],
				utils.Conv2ByteSlice(ulNasTransportMsg.GetUserLocationInfo().GetTimeStamp(),
					types3gpp.SizeofTimeStamp))
			p.UserLocInfo.IsTimeStampPrst = true
		}
		if ulNasTransportMsg.GetUserLocationInfo().GetPsCellInfoPrst() {
			p.UserLocInfo.PsCellInfo.PlmnID.SetValue(
				utils.Conv2ByteSlice(ulNasTransportMsg.GetUserLocationInfo().GetPsCellInfo().GetPlmn(),
					types3gpp.SizeofPlmnID),
				types3gpp.BigEndian)
			copy(p.UserLocInfo.PsCellInfo.NrCellId[:],
				utils.Conv2ByteSlice(ulNasTransportMsg.GetUserLocationInfo().GetPsCellInfo().GetNrCellId(),
					types3gpp.SizeofNrCellID))
			p.UserLocInfo.IsPsCellInfoPrst = true
		}

	} else {
		return fmt.Errorf("failed to decode msg buffer")
	}

	return nil
}
