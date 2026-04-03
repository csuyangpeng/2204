package ngapmsg

import (
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"

	"fmt"
)

func TestUeContextReleaseCmdMsg_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewUeContextReleaseCmdMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.UeNgapIdType = types3gpp.UeNgApIdPairType
	msg.AmfNgapId = 100
	msg.RanNgapId = 200
	msg.Cause.Type = types3gpp.CT_RadioNetwork
	msg.Cause.Value = types3gpp.Radiok_handover_cancelled

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
func TestUeContextReleaseCmdMsg_Decode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{0, 41, 0, 17, 0, 0, 2, 0, 15, 64, 2, 1, 64, 0, 114, 0, 4, 0, 100, 0, 200}

	cmnNgAp := codec.NewNgapCodec()
	cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))

	cmnNgAp.DecodeHeader(ossCtxtPtr)

	fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	msg := NewUeContextReleaseCmdMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(msg)
}
func TestUeContextReleaseCmdMsg02_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewUeContextReleaseCmdMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.UeNgapIdType = types3gpp.AmfUeNgApIdType
	msg.AmfNgapId = 100
	//msg.IsRppPrst = true
	//msg.RanPagingPriority = 5
	msg.Cause.Type = types3gpp.CT_RadioNetwork
	msg.Cause.Value = types3gpp.Radiok_handover_cancelled

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestUeContextReleaseCmdMsg02_Decode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{0, 41, 0, 15, 0, 0, 2, 0, 15, 64, 2, 1, 64, 0, 114, 0, 2, 64, 100}

	cmnNgAp := codec.NewNgapCodec()
	cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))

	cmnNgAp.DecodeHeader(ossCtxtPtr)

	fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	msg := NewUeContextReleaseCmdMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(msg)
}
