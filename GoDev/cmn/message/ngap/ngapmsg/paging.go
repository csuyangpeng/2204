package ngapmsg

import "C"
import (
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"unsafe"
)

type PagingMsg struct {
	UePagingId       types3gpp.Stmsi5G
	TaiListForPaging []*types3gpp.TAI
	// opt ie
	PagingDrx              types3gpp.PagingDRX
	PagingPriority         types3gpp.PagingPriority
	UeRadioCap4PagingNR    string
	UeRadioCap4PagingEutra string
	PagingOrigin           types3gpp.PagingOrigin
	OptFlags               bitset.BitSet

	ctxt codec.NgapOssCtxt
}

const (
	ICSR_PagingDrx = iota
	ICSR_PagingPriority
	ICSR_UeRaidoCapPagingNR
	ICSR_UeRaidoCapPagingEutra
	ICSR_PagingOrigin
	ICSR_UeRaidoCapForPaging
)

//NewPaging create a new Message
func NewPagingMsg() *PagingMsg {
	return &PagingMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PagingMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// add tai list for paging
func (p *PagingMsg) AddTaiPagingList(item *types3gpp.TAI) {
	p.TaiListForPaging = append(p.TaiListForPaging, item)
}

func (p *PagingMsg) String() string {
	rtStr := fmt.Sprintf("PagingMsg{"+
		"PagingDrx(%v %s),PagingPriority[%v %s],PagingOrigin[%v (%s)],UePagingId(%d),"+
		"ForPaging[%v,UeRadioCap4PagingNR(%v,%s),UeRadioCap4PagingEutra(%v,%s)]}",
		p.OptFlags.Test(ICSR_PagingDrx),
		p.PagingDrx,
		p.OptFlags.Test(ICSR_PagingPriority),
		p.PagingPriority,
		p.OptFlags.Test(ICSR_PagingOrigin),
		p.PagingOrigin,
		p.UePagingId,
		p.OptFlags.Test(ICSR_UeRaidoCapForPaging),
		p.OptFlags.Test(ICSR_UeRaidoCapPagingNR),
		p.UeRadioCap4PagingNR,
		p.OptFlags.Test(ICSR_UeRaidoCapPagingEutra),
		p.UeRadioCap4PagingEutra)

	rtStr += "Tai List For Paging{"
	for _, v := range p.TaiListForPaging {
		rtStr += fmt.Sprintf("%s,", v)
	}
	return rtStr + "}}"
}

func (p *PagingMsg) Encode() []byte {

	pagingCodec := codec.NewPagingCodec()
	defer codec.DeletePagingCodec(pagingCodec)

	// encode ue paging identity
	uePagingId := codec.NewUePagingIdentity()
	defer codec.DeleteUePagingIdentity(uePagingId)

	amfSet := make([]byte, 2)
	amfSet = p.UePagingId.GetAmfSetID()
	uePagingId.SetSetId(&amfSet[0])
	uePagingId.SetPointer(p.UePagingId.GetAmfPointer())
	tmsi := make([]byte, 4)
	binary.BigEndian.PutUint32(tmsi, p.UePagingId.GetTmsi())
	uePagingId.SetTmsi(&tmsi[0])
	pagingCodec.SetUePagingIdentity(uePagingId)

	// encode Tai List For Paging
	for _, v := range p.TaiListForPaging {
		tai := codec.NewTaiType()
		defer codec.DeleteTaiType(tai)

		taiPlmnId := v.Plmn.GetValue(types3gpp.BigEndian)
		tai.SetPlmn(&taiPlmnId[0])
		tac := v.Tac.GetByteSlice()
		tai.SetTac(&tac[0])

		pagingCodec.AddTaiPagingList(tai)
	}

	// optional IEs
	if p.OptFlags.Test(ICSR_PagingOrigin) {
		pagingCodec.SetPagingOrigin(uint8(p.PagingOrigin))
	}

	if p.OptFlags.Test(ICSR_PagingDrx) {
		pagingCodec.SetPagingDrx(uint8(p.PagingDrx))
	}

	if p.OptFlags.Test(ICSR_PagingPriority) {
		pagingCodec.SetPagingPriority(uint8(p.PagingPriority))
	}

	ueRadioCap := codec.NewUeRadioCapForPaging()
	defer codec.DeleteUeRadioCapForPaging(ueRadioCap)

	if p.OptFlags.Test(ICSR_UeRaidoCapForPaging) {
		if p.OptFlags.Test(ICSR_UeRaidoCapPagingNR) {
			ueRadioCap.SetUeRadioCapForPagingNrPrst(true)
			ueRadioCap.SetUeRadioCapForPagingNr(p.UeRadioCap4PagingNR)
		}
		if p.OptFlags.Test(ICSR_UeRaidoCapPagingEutra) {
			ueRadioCap.SetUeRadioCapForPagingEutraPrst(true)
			ueRadioCap.SetUeRadioCapForPagingEutra(p.UeRadioCap4PagingEutra)
		}
		pagingCodec.SetUeRadioCapForPaging(ueRadioCap)
	}

	msgBuffer := pagingCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PagingMsg) Decode(msgbuf []byte) error {

	pagingCodec := codec.NewPagingCodec()
	defer codec.DeletePagingCodec(pagingCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if pagingCodec.Decode(p.ctxt, msgBuffer) == true {

		// ue paging identity
		amfSetId := make([]byte, 2)
		copy(amfSetId,
			utils.Conv2ByteSlice(pagingCodec.GetUePagingIdentity().GetSetId(), types3gpp.SizeofAmfSetID))
		p.UePagingId.SetAmfSetID(amfSetId)
		p.UePagingId.SetAmfPointer(pagingCodec.GetUePagingIdentity().GetPointer())

		tmsi := make([]byte, 4)
		copy(tmsi, utils.Conv2ByteSlice(pagingCodec.GetUePagingIdentity().GetTmsi(), 4))
		p.UePagingId.SetTmsi(binary.BigEndian.Uint32(tmsi))

		// tailist for paging
		tailistforPaging := pagingCodec.GetTaiPagingList()
		for i := 0; i < int(tailistforPaging.Size()); i++ {
			item := &types3gpp.TAI{}

			taitype := tailistforPaging.Get(i)

			item.Plmn.SetValue(
				utils.Conv2ByteSlice(taitype.GetPlmn(), types3gpp.SizeofPlmnID),
				types3gpp.BigEndian)

			tac := utils.Conv2ByteSlice(taitype.GetTac(), types3gpp.SizeofTAC)
			for i, v := range tac {
				item.Tac[i] = v
			}

			p.TaiListForPaging = append(p.TaiListForPaging, item)
		}

		// paging drx
		if pagingCodec.IsPagingDrxPrst() {
			p.PagingDrx = types3gpp.PagingDRX(pagingCodec.GetPagingDrx())
			p.OptFlags.Set(ICSR_PagingDrx)
		}
		// paging priority
		if pagingCodec.IsPagingPriorityPrst() {
			p.PagingPriority = types3gpp.PagingPriority(pagingCodec.GetPagingPriority())
			p.OptFlags.Set(ICSR_PagingPriority)
		}
		// paging origin
		if pagingCodec.IsPagingOriginPrst() {
			p.PagingOrigin = types3gpp.PagingOrigin(pagingCodec.GetPagingOrigin())
			p.OptFlags.Set(ICSR_PagingOrigin)
		}
		//UeRadioCapbilityforPaging
		if pagingCodec.IsUeRadioCapForPagingPrst() {
			if pagingCodec.GetUeRadioCapForPaging().GetUeRadioCapForPagingNrPrst() {
				p.UeRadioCap4PagingNR = pagingCodec.GetUeRadioCapForPaging().GetUeRadioCapForPagingNr()
				p.OptFlags.Set(ICSR_UeRaidoCapPagingNR)
			}
			if pagingCodec.GetUeRadioCapForPaging().GetUeRadioCapForPagingEutraPrst() {
				p.UeRadioCap4PagingEutra = pagingCodec.GetUeRadioCapForPaging().GetUeRadioCapForPagingEutra()
				p.OptFlags.Set(ICSR_UeRaidoCapPagingEutra)
			}
			p.OptFlags.Set(ICSR_UeRaidoCapForPaging)
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}
	return nil
}
