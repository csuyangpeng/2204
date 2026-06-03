package ngapmsg

import "C"
import (
	"fmt"
	"github.com/willf/bitset"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"unsafe"
)

// NgResetAckMsg struct definition
type NgResetAckMsg struct {
	UeAssLogicalNgConnList []*types3gpp.UeAssLogicalNgConnItem

	OptFlags bitset.BitSet
	ctxt     codec.NgapOssCtxt
}

const (
	ICSR_UeAssLogicalNgConn = iota
)

// NewNgResetAckMsg create a new Message
func NewNgResetAckMsg() *NgResetAckMsg {
	return &NgResetAckMsg{}
}

// SetOssCodecCtxt set OSS codec context
func (p *NgResetAckMsg) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

//AddUeAssLogicalNgConn add UeAssLogicalNgConn into UeAssLogicalNgConnList
func (p *NgResetAckMsg) AddUeAssLogicalNgConnItem(
	assLogicalNgConnItem *types3gpp.UeAssLogicalNgConnItem) {
	p.UeAssLogicalNgConnList = append(p.UeAssLogicalNgConnList, assLogicalNgConnItem)
}

func (p *NgResetAckMsg) String() string {
	rtStr := fmt.Sprintf("UeAssLogicalNgConnList:{Prst(%v) \n", p.OptFlags.Test(ICSR_UeAssLogicalNgConn))
	for _, v := range p.UeAssLogicalNgConnList {
		rtStr += fmt.Sprintf("%s,", v)
	}
	return rtStr + "}}"
}

func (p *NgResetAckMsg) Encode() []byte {

	resetAckCodec := codec.NewNgResetAckCodec()
	defer codec.DeleteNgResetAckCodec(resetAckCodec)

	//Ue Ass Logical Ng Conn
	if p.OptFlags.Test(ICSR_UeAssLogicalNgConn) {
		for _, v := range p.UeAssLogicalNgConnList {
			ueAssLogicalNgConnItem := codec.NewUeAssLogicalNgConn()
			defer codec.DeleteUeAssLogicalNgConn(ueAssLogicalNgConnItem)

			if v.IsAmfUeNgapIdPrst {
				ueAssLogicalNgConnItem.SetAmfUeNgapIdPrst(true)
				ueAssLogicalNgConnItem.SetAmfUeNgapId(uint64(v.AmfUeNgapId))
			}

			if v.IsRanUeNgapIdPrst {
				ueAssLogicalNgConnItem.SetRanUeNgapIdPrst(true)
				ueAssLogicalNgConnItem.SetRanUeNgapId(uint(v.RanUeNgapId))
			}

			resetAckCodec.AddUeAssLogicalNgConnList(ueAssLogicalNgConnItem)

		}
	}

	msgBuf := resetAckCodec.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuf)
	bufLen := msgBuf.GetLength()
	bufValue := msgBuf.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *NgResetAckMsg) Decode(msgbuf []byte) error {

	resetAckCodec := codec.NewNgResetAckCodec()
	defer codec.DeleteNgResetAckCodec(resetAckCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])
	if resetAckCodec.Decode(p.ctxt, msgBuffer) == true {
		//Ue Ass Logical Ng Conn
		if resetAckCodec.IsUeAssLogicalNgConnPrst() {
			ueAssNgConList := resetAckCodec.GetUeAssLogicalNgConnList()
			for i := 0; i < int(ueAssNgConList.Size()); i++ {
				item := &types3gpp.UeAssLogicalNgConnItem{}
				item.IsAmfUeNgapIdPrst = ueAssNgConList.Get(i).GetAmfUeNgapIdPrst()
				if item.IsAmfUeNgapIdPrst {
					item.AmfUeNgapId = uint32(ueAssNgConList.Get(i).GetAmfUeNgapId())
				}
				item.IsRanUeNgapIdPrst = ueAssNgConList.Get(i).GetRanUeNgapIdPrst()
				if item.IsRanUeNgapIdPrst {
					item.RanUeNgapId = uint32(ueAssNgConList.Get(i).GetRanUeNgapId())
				}
				p.UeAssLogicalNgConnList = append(p.UeAssLogicalNgConnList, item)
			}
			p.OptFlags.Set(ICSR_UeAssLogicalNgConn)
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
