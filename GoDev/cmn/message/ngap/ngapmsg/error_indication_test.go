package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"testing"
)

func Test_ErrIndicationMsg_encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewErrIndicationMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.IsAmfUeNgapIdPrst = true
	msg.AmfUeNgapId = 1
	msg.IsRanUeNgapIdPrst = true
	msg.RanUeNgapId = 2
	msg.IsCausePrst = true
	msg.CauseType = 1
	msg.CauseValue = 2

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
	fmt.Println([]byte(string([]byte{1, 1})))
}

func Test_ErrIndicationMsg_decode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{0, 9, 64, 21, 0, 0, 3, 0, 10, 64, 2, 0, 1, 0, 85, 64, 2, 0, 2, 0, 15, 64, 2, 0, 128}
	fmt.Println(encMsg)

	errIndMsg := NewErrIndicationMsg()
	errIndMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := errIndMsg.Decode(encMsg)
	fmt.Println("result", decode)
}
