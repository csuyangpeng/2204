/** Copyright(C),2020-2022
* Author: zmj
* Date: 12/22/20 2:56 PM
* Description:
 */
package ngapmsg

import "C"
import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"unsafe"
)

// UeContextReleaseRequest struct definition
type UeRadioCapInfoIndMsg struct {
	RanUeNgapId uint32
	AmfUeNgapId uint64

	// 23502 4.2.6
	// If List of PDU Session ID(s) with active N3 user plane is included in step 1b (n2 ue release request),
	// the step 5 to 7 (amf -> smf -> upf :update SM context & n4 modification)
	// are performed before step 2(n2 ue release command)
	UeRadioCapbility             string
	UeRadioCapbility4PagingNr    string
	HasPagingNr                  bool
	UeRadioCapbility4PagingEutra string
	HasPagingEutra               bool

	ctxt codec.NgapOssCtxt
}

// NewUeRadioCapInfoIndMsg create a new Message
func NewUeRadioCapInfoIndMsg() *UeRadioCapInfoIndMsg {
	return &UeRadioCapInfoIndMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *UeRadioCapInfoIndMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *UeRadioCapInfoIndMsg) String() string {
	return fmt.Sprintf("UeRadioCapInfoInd:RanNgapId(%d), AmfNgapId(%d), UeRadioCapbility(%x), "+
		"HasPagingNr(%v), UeRadioCap4PagingNr(%x)"+
		"HasPagingEutra(%v), UeRadioCap4PagingEutra(%x)",
		p.RanUeNgapId,
		p.AmfUeNgapId,
		p.UeRadioCapbility,
		p.HasPagingNr,
		p.UeRadioCapbility4PagingNr,
		p.HasPagingEutra,
		p.UeRadioCapbility4PagingEutra)
}

func (p *UeRadioCapInfoIndMsg) Encode() []byte {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	req := codec.NewUeRadioCapaInfoIndicationCodec()
	defer codec.DeleteUeRadioCapaInfoIndicationCodec(req)

	req.SetRanUeNgapId(uint(p.RanUeNgapId))
	req.SetAmfUeNgapId(p.AmfUeNgapId)

	paging := codec.NewUeRadioCapForPaging()
	defer codec.DeleteUeRadioCapForPaging(paging)

	req.SetUERadioCapability(p.UeRadioCapbility)
	if p.HasPagingNr {
		paging.SetUeRadioCapForPagingNrPrst(true)
		paging.SetUeRadioCapForPagingNr(p.UeRadioCapbility4PagingNr)
		req.SetUeRadioCapForPaging(paging)
	}

	if p.HasPagingEutra {
		paging.SetUeRadioCapForPagingEutraPrst(true)
		paging.SetUeRadioCapForPagingEutra(p.UeRadioCapbility4PagingEutra)
	}

	msgBuffer := req.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *UeRadioCapInfoIndMsg) Decode(msgBuf []byte) error {
	rlogger.FuncEntry(types.ModuleCmnNgap, nil)

	req := codec.NewUeRadioCapaInfoIndicationCodec()
	defer codec.DeleteUeRadioCapaInfoIndicationCodec(req)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgBuf)))
	msgBuffer.SetValue(&msgBuf[0])

	if req.Decode(p.ctxt, msgBuffer) == true {
		p.RanUeNgapId = uint32(req.GetRanUeNgapId())
		p.AmfUeNgapId = req.GetAmfUeNgapId()

		if req.IsUeRadioCapForPagingPrst() {
			paging := req.GetUeRadioCapForPaging()

			if paging.GetUeRadioCapForPagingNrPrst() {
				p.HasPagingNr = true
				p.UeRadioCapbility4PagingNr = paging.GetUeRadioCapForPagingNr()
			}

			if paging.GetUeRadioCapForPagingEutraPrst() {
				p.HasPagingEutra = true
				p.UeRadioCapbility4PagingEutra = paging.GetUeRadioCapForPagingEutra()
			}
		}
	} else {
		return fmt.Errorf("failed to decode msg Bufer")
	}
	return nil
}
