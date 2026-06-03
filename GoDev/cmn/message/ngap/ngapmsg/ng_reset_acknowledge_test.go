package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func Test_NgResetAck_encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewNgResetAckMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.OptFlags.Set(ICSR_UeAssLogicalNgConn)
	item := &types3gpp.UeAssLogicalNgConnItem{}
	item.IsAmfUeNgapIdPrst = true
	item.AmfUeNgapId = 1
	item.IsRanUeNgapIdPrst = true
	item.RanUeNgapId = 2

	msg.AddUeAssLogicalNgConnItem(item)

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func Test_NgResetAckMsg_decode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{32, 20, 0, 12, 0, 0, 1, 0, 111, 64, 5, 1, 96, 1, 0, 2}
	fmt.Println(encMsg)

	resetMsg := NewNgResetAckMsg()
	resetMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := resetMsg.Decode(encMsg)
	fmt.Println("result", decode)
}
