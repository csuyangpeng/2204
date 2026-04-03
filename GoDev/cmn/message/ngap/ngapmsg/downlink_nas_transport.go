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

// DownlinkNasTransportMessage struct definition
type DownlinkNasTransportMessage struct {
	RanUeNgapId             uint32
	AmfUeNgapId             uint64
	NasPdu                  []byte
	OldAmfName              string
	IsOldAmfNamePrst        bool
	RanPagingPriority       uint16
	IsRanPagingPriorityPrst bool
	IndexToRfsp             uint64
	IsIndexToRfspPrst       bool
	UeAmbr                  types3gpp.Ambr
	IsUeAmbrPrst            bool
	AllowedNssai            []types3gpp.Snssai
	IsAllowedNssaiPrst      bool
	ctxt                    codec.NgapOssCtxt
}

func (p *DownlinkNasTransportMessage) DumpMsg() {}

// NewDownlinkNasTransportMessage create a new Message
func NewDownlinkNasTransportMessage() *DownlinkNasTransportMessage {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	return &DownlinkNasTransportMessage{}
}

// SetOssCodecCtxt set OSS codec context
func (p *DownlinkNasTransportMessage) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// AddAllowedNssai add NSSAI into List
func (p *DownlinkNasTransportMessage) AddAllowedNssai(aNssaiPtr *types3gpp.Snssai) {
	p.AllowedNssai = append(p.AllowedNssai, *aNssaiPtr)
	p.IsAllowedNssaiPrst = true
}

// Encode the message into bytes
func (p *DownlinkNasTransportMessage) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	dlNasTransportMsg := codec.NewDownlinkNASTransportCodec()
	defer codec.DeleteDownlinkNASTransportCodec(dlNasTransportMsg)

	// ran ue ngap id
	dlNasTransportMsg.SetRanUeNgapId(uint(p.RanUeNgapId))
	// amf ue ngap id
	dlNasTransportMsg.SetAmfUeNgapId(uint64(p.AmfUeNgapId))
	// nas message
	dlNasTransportMsg.SetNasPdu(string(p.NasPdu))

	// old amf name
	if p.IsOldAmfNamePrst {
		dlNasTransportMsg.SetOldAmfName(p.OldAmfName)
	}

	// ranPagingPriority
	if p.IsRanPagingPriorityPrst {
		dlNasTransportMsg.SetRanPagingPriority(p.RanPagingPriority)
	}

	// index to rfsp
	if p.IsIndexToRfspPrst {
		dlNasTransportMsg.SetIndexToRfsp(int64(p.IndexToRfsp))
	}

	// ue ambr
	if p.IsUeAmbrPrst {
		ueAmbr := codec.NewUeAmbr()
		defer codec.DeleteUeAmbr(ueAmbr)

		ueAmbr.SetDownlink(int64(p.UeAmbr.Downlink))
		ueAmbr.SetUplink(int64(p.UeAmbr.Uplink))
		dlNasTransportMsg.SetUeAmbr(ueAmbr)
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
		dlNasTransportMsg.AddAllowedNssai(snssai)
	}

	//encode
	msgBuffer := dlNasTransportMsg.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

// Decode the bytes into message
func (p *DownlinkNasTransportMessage) Decode(msgbuf []byte) error {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	dlNasTransportMsg := codec.NewDownlinkNASTransportCodec()
	defer codec.DeleteDownlinkNASTransportCodec(dlNasTransportMsg)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if dlNasTransportMsg.Decode(p.ctxt, msgBuffer) == true {
		p.RanUeNgapId = uint32(dlNasTransportMsg.GetRanUeNgapId())
		p.AmfUeNgapId = uint64(dlNasTransportMsg.GetAmfUeNgapId())
		p.NasPdu = []byte(dlNasTransportMsg.GetNasPdu())

		if dlNasTransportMsg.IsOldAmfNamePresent() {
			p.OldAmfName = dlNasTransportMsg.GetOldAmfName()
			p.IsOldAmfNamePrst = true
		}

		if dlNasTransportMsg.IsRanPagingPriorityPresent() {
			p.RanPagingPriority = dlNasTransportMsg.GetRanPagingPriority()
			p.IsRanPagingPriorityPrst = true
		}

		if dlNasTransportMsg.IsIndexToRfspPresent() {
			p.IndexToRfsp = uint64(dlNasTransportMsg.GetIndexToRfsp())
			p.IsIndexToRfspPrst = true
		}

		if dlNasTransportMsg.IsUeAmbrPresent() {
			p.UeAmbr.Uplink = uint64(dlNasTransportMsg.GetUeAmbr().GetUplink())
			p.UeAmbr.Downlink = uint64(dlNasTransportMsg.GetUeAmbr().GetDownlink())
			p.IsUeAmbrPrst = true
		}

		if dlNasTransportMsg.IsAllowedNssaiPresent() {
			snssaiList := dlNasTransportMsg.GetAllowedNssaiList()
			for i := 0; i < int(snssaiList.Size()); i++ {
				snssaiMsg := types3gpp.Snssai{}

				snnsai := snssaiList.Get(i)
				sst := snnsai.GetSst()
				snssaiMsg.Sst = *sst

				if snnsai.GetSdPresent() == true {
					var tmp [types3gpp.SizeofSD]byte
					copy(tmp[:],
						utils.Conv2ByteSlice(snnsai.GetSd(), types3gpp.SizeofSD))

					snssaiMsg.SdPrst =true
					snssaiMsg.Sd = types3gpp.ConvertSdToU32(tmp[:],types.BigEndian)
				}
				p.AddAllowedNssai(&snssaiMsg)
			}
		}

	} else {
		return fmt.Errorf("failed to decode msg buffer")
	}

	return nil
}
