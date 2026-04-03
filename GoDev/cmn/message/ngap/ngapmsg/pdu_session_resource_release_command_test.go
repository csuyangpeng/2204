package ngapmsg

import (
	"fmt"
	"testing"

	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
)

func TestPduSessResRelCmmdMsg_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewPduSessResRelRespMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.RanUeNGAPId = 5
	msg.AmfUeNGAPId = 6
	relItem := &types3gpp.PduSessResRelCmmdItem{}
	relItem.PduSessionId = 5
	relItem.PduSessResRelCmmdTransfer = string([]byte{0, 32})
	//msg.AddPduSessResRelCmmdItem(relItem)
	//
	//msg.RanPagePriority = 10
	//msg.IsRanPagePriorityPrst = true
	//msg.IsNasPduPrst = true
	//msg.NasPdu = []byte("123456789")

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestPduSessResRelCmmdMsg_Decode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{0, 28, 0, 44, 0, 0, 5, 0, 79, 0, 6, 0, 0, 5, 2, 0, 32, 0, 83, 64, 1, 9, 0, 38, 64, 10, 9, 49, 50, 51, 52, 53, 54, 55, 56, 57, 0, 10, 0, 2, 0, 6, 0, 85, 0, 2, 0, 5}

	cmnNgAp := codec.NewNgapCodec()
	cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))

	cmnNgAp.DecodeHeader(ossCtxtPtr)

	fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	msg := NewPduSessResRelCmmdMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(msg)
}
