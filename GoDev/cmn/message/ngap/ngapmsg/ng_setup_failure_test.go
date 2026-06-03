package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"testing"
)

func Test_NgSetupFailure_encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewNgSetupFailureMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.RelativeTimeToWait = 3
	msg.CauseType = 3
	msg.CauseValue = 1

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func Test_NgSetupFailureMsg_decode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{64, 21, 0, 8, 0, 0, 1, 0, 15, 64, 1, 68}
	fmt.Println(encMsg)

	setupFailureMsg := NewNgSetupFailureMsg()
	setupFailureMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := setupFailureMsg.Decode(encMsg)
	fmt.Println("result", decode)
}
