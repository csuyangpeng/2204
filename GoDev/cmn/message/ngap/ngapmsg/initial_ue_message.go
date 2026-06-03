package ngapmsg

import "C"
import (
	"encoding/binary"
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"unsafe"
)

//38413  9.2.5.1
// InitialUEMessage struct definition
type InitialUeMessage struct {
	RanUeNgapId           uint32
	NasPdu                []byte
	UserLocInfo           types3gpp.UserLocationInfo
	RrcEstablishmentCause types3gpp.RrcEstbCause

	//optional
	AmfSetId               uint16
	IsAmfSetIdPrst         bool
	UeContextRequest       uint16
	IsUeContextRequestPrst bool
	Stmsi5G                types3gpp.Stmsi5G
	IsStmsi5gPrst          bool
	AllowedNssai           []types3gpp.Snssai
	IsAllowedNssaiPrst     bool
	ctxt                   codec.NgapOssCtxt
}

// NewInitialUeMessage create a new Message
//func NewInitialUeMessage(ctxt codec.NgApOssCtxt) *InitialUeMessage {
func NewInitialUeMessage() *InitialUeMessage {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)
	return &InitialUeMessage{}
}

// SetOssCodecCtxt set OSS codec context
func (p *InitialUeMessage) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// AddAllowedNssai add NSSAI into List
func (p *InitialUeMessage) AddAllowedNssai(aNssaiPtr *types3gpp.Snssai) {
	p.AllowedNssai = append(p.AllowedNssai, *aNssaiPtr)
	p.IsAllowedNssaiPrst = true
}

// Encode the message into bytes
func (p *InitialUeMessage) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	initialUeMsg := codec.NewInitialUEMessageCodec()
	defer codec.DeleteInitialUEMessageCodec(initialUeMsg)

	// ran ue ngap id
	initialUeMsg.SetRanUeNgapId(uint(p.RanUeNgapId))
	initialUeMsg.SetNasPdu(string(p.NasPdu))

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
	initialUeMsg.SetUserLocationInfo(userLocationInfo)

	// rrc establishment cause
	initialUeMsg.SetRrcEstablishmentCause(byte(p.RrcEstablishmentCause))

	// AmfSetId
	if p.IsAmfSetIdPrst {
		amfSet := make([]byte, 2)
		binary.BigEndian.PutUint16(amfSet, p.AmfSetId)
		initialUeMsg.SetAmfSetId(&amfSet[0])
	}

	// UeContextRequest
	if p.IsUeContextRequestPrst {
		initialUeMsg.SetUeContextRequest(p.UeContextRequest)
	}

	// Stmsi5G
	if p.IsStmsi5gPrst {
		stmsi := codec.NewStmsi5G()
		defer codec.DeleteStmsi5G(stmsi)
		amfSet := make([]byte, 2)
		amfSet = p.Stmsi5G.GetAmfSetID()
		stmsi.SetSetId(&amfSet[0])
		stmsi.SetPointer(p.Stmsi5G.GetAmfPointer())
		tmsi := make([]byte, 4)
		binary.BigEndian.PutUint32(tmsi, p.Stmsi5G.GetTmsi())
		stmsi.SetTmsi(&tmsi[0])
		initialUeMsg.Set5gStmsi(stmsi)
	}

	// AllowedNssai
	for _, nssai := range p.AllowedNssai {
		snssai := codec.NewSNssai()
		defer codec.DeleteSNssai(snssai)
		snssai.SetSst(&nssai.Sst)
		if nssai.SdPrst == true {
			var sd [types3gpp.SizeofSD]byte
			types3gpp.ConvertU32ToSd(sd[:], nssai.Sd, types.BigEndian)
			snssai.SetSd(&sd[0])
			snssai.SetSdPresent(true)
		}
		initialUeMsg.AddAllowedNssai(snssai)
	}

	//encode
	msgBuffer := initialUeMsg.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

// Decode the bytes into message
func (p *InitialUeMessage) Decode(msgbuf []byte) error {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	initialUeMsg := codec.NewInitialUEMessageCodec()
	defer codec.DeleteInitialUEMessageCodec(initialUeMsg)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if initialUeMsg.Decode(p.ctxt, msgBuffer) == true {
		p.RanUeNgapId = uint32(initialUeMsg.GetRanUeNgapId())
		p.NasPdu = []byte(initialUeMsg.GetNasPdu())

		p.UserLocInfo.NrCgi.PlmnID.SetValue(
			utils.Conv2ByteSlice(initialUeMsg.GetUserLocationInfo().GetNRCgi().GetPlmn(),
				types3gpp.SizeofPlmnID), types3gpp.BigEndian)

		copy(p.UserLocInfo.NrCgi.NrCellId[:],
			utils.Conv2ByteSlice(initialUeMsg.GetUserLocationInfo().GetNRCgi().GetNrCellId(),
				types3gpp.SizeofNrCellID))

		if initialUeMsg.GetUserLocationInfo().GetIsTimeStampPrst() {
			copy(p.UserLocInfo.TimeStmp[:],
				utils.Conv2ByteSlice(initialUeMsg.GetUserLocationInfo().GetTimeStamp(),
					types3gpp.SizeofTimeStamp))
			p.UserLocInfo.IsTimeStampPrst = true
		}
		if initialUeMsg.GetUserLocationInfo().GetPsCellInfoPrst() {
			p.UserLocInfo.PsCellInfo.PlmnID.SetValue(
				utils.Conv2ByteSlice(initialUeMsg.GetUserLocationInfo().GetPsCellInfo().GetPlmn(),
					types3gpp.SizeofPlmnID),
				types3gpp.BigEndian)
			copy(p.UserLocInfo.PsCellInfo.NrCellId[:],
				utils.Conv2ByteSlice(initialUeMsg.GetUserLocationInfo().GetPsCellInfo().GetNrCellId(),
					types3gpp.SizeofNrCellID))
			p.UserLocInfo.IsPsCellInfoPrst = true
		}

		p.RrcEstablishmentCause = types3gpp.RrcEstbCause(initialUeMsg.GetRrcEstablishmentCause())

		if initialUeMsg.Is5gStmsiPresent() {
			amfSetId := make([]byte, 2)
			copy(amfSetId,
				utils.Conv2ByteSlice(initialUeMsg.Get5gStmsi().GetSetId(), types3gpp.SizeofAmfSetID))
			p.Stmsi5G.SetAmfSetID(amfSetId)
			p.Stmsi5G.SetAmfPointer(initialUeMsg.Get5gStmsi().GetPointer())

			tmsi := make([]byte, 4)
			copy(tmsi,
				utils.Conv2ByteSlice(initialUeMsg.Get5gStmsi().GetTmsi(), 4))
			p.Stmsi5G.SetTmsi(binary.BigEndian.Uint32(tmsi))
			p.IsStmsi5gPrst = true
		}

		if initialUeMsg.IsAmfSetPresent() {
			amfSetId := make([]byte, 2)
			copy(amfSetId,
				utils.Conv2ByteSlice(initialUeMsg.GetAmfSetId(), types3gpp.SizeofAmfSetID))
			p.AmfSetId = binary.BigEndian.Uint16(amfSetId)

			p.IsAmfSetIdPrst = true
		}

		if initialUeMsg.IsUeContextRequestPresent() {
			p.IsUeContextRequestPrst = true
			p.UeContextRequest = initialUeMsg.GetUeContextRequest()
		}

		if initialUeMsg.IsAllowedNssaiPresent() {
			snssaiList := initialUeMsg.GetAllowedNssaiList()
			for i := 0; i < int(snssaiList.Size()); i++ {
				snssaiMsg := types3gpp.Snssai{}

				snnsai := snssaiList.Get(i)
				sst := snnsai.GetSst()
				snssaiMsg.Sst = *sst

				if snnsai.GetSdPresent() == true {
					var tmp [types3gpp.SizeofSD]byte
					copy(tmp[:],
						utils.Conv2ByteSlice(snnsai.GetSd(), types3gpp.SizeofSD))

					snssaiMsg.SdPrst = true
					snssaiMsg.Sd = types3gpp.ConvertSdToU32(tmp[:], types.BigEndian)
				}
				p.AddAllowedNssai(&snssaiMsg)
			}
		}

	} else {
		return fmt.Errorf("failed to decode msg buffer")
	}

	return nil
}
