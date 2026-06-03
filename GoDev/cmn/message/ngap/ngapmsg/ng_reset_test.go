package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"testing"
)

func Test_NgResetMsg_encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewNgResetMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.CauseType = 1
	msg.CauseValue = 2

	msg.ResetTypeChoice.IsNgInterfacePrst = true
	msg.ResetTypeChoice.NgInterface = 0

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func Test_NgResetMsg_decode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{0, 20, 0, 14, 0, 0, 2, 0, 15, 64, 2, 0, 128, 0, 88, 0, 1, 0}
	fmt.Println(encMsg)

	resetMsg := NewNgResetMsg()
	resetMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := resetMsg.Decode(encMsg)
	fmt.Println("result", decode)
}
