package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func TestUeContextReleaseReqMsg_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewUeContextReleaseReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.AmfUeNgapId = 100
	msg.RanUeNgapId = 200

	msg.PduSessIdList = append(msg.PduSessIdList, 5)
	msg.PduSessIdList = append(msg.PduSessIdList, 6)

	msg.Cause.Type = types3gpp.CT_RadioNetwork
	msg.Cause.Value = types3gpp.Radiok_unknown_targetID
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestUeContextReleaseReqMsg_Decode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{0, 42, 64, 30, 0, 0, 4, 0, 133, 0, 5, 1, 0, 6, 0, 5, 0, 15, 64, 2, 3, 0, 0, 10, 0, 2, 0, 100, 0, 85, 0, 2, 0, 200}

	cmnNgAp := codec.NewNgapCodec()
	cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))

	cmnNgAp.DecodeHeader(ossCtxtPtr)

	fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	msg := NewUeContextReleaseReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Printf("decode Msg: %s", msg)
}
