package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func Test_InitContextSetupFailMsg_encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	failSetupContext := NewInitContextSetupFailMsg()
	failSetupContext.SetOssCodecCtxt(ossCtxtPtr)

	failSetupContext.AmfNgapId = 5
	failSetupContext.RanNgapId = 6
	failSetupContext.Cause.Type = 1
	failSetupContext.Cause.Value = 2
	failsetupItem := &types3gpp.PduSessResFailedToSetupItem{}
	failsetupItem.PduSessionId = 5
	failsetupItem.UnsuccessTransfer = string([]byte{0, 32})
	failSetupContext.AddPduSessResFailToSetupItem(failsetupItem)

	fmt.Println(failSetupContext)
	encodeMsg := failSetupContext.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func Test_InitContextSetupFailMsg_decode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{64, 14, 0, 31, 0, 0, 4, 0, 10, 64, 2, 0, 5, 0, 85, 64, 2, 0, 6, 0, 132, 64, 6, 0, 0, 5, 2, 0, 32, 0, 15, 64, 2, 0, 128}
	fmt.Println(encMsg)

	msg := NewInitContextSetupFailMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
}
