package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func TestPduSessResModCmdMsg_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewPduSessResModifyReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.RanUeNgApId = 5
	msg.AmfUeNgApId = 6
	relItem := &types3gpp.PduSessResModifyReqItem{}
	relItem.PduSessionId = 5
	relItem.PduSessResModReqTrans = string([]byte{0, 0, 4, 0, 130, 0, 7, 8, 7, 161, 32, 16, 78, 32, 0, 140, 0, 21, 32, 15, 128, 10, 180, 9, 242, 0, 0, 37, 52, 1, 240, 10, 180, 9, 245, 0, 0, 37, 52, 0, 135, 0, 3, 0, 130, 164, 0, 137, 0, 4, 0, 10, 0, 0})
	msg.AddPduSessResModifyReqItem(relItem)

	msg.RanPagePriority = 10
	msg.IsRanPagePriorityPrst = true

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
