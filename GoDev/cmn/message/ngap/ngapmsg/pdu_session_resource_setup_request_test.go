package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func TestPduSessResSetupReqMsg_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewPduSessResSetupReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.AmfUeNgApId = 1
	msg.RanPagePriority = 10
	msg.RanUeNgApId = 2
	msg.IsRanPagePriorityPrst = true
	msg.IsNasPduPrst = true
	msg.NasPdu = "123456789"
	msg.IsUeAmbrPrst = true
	msg.UeAmbr.Uplink = 5000000
	msg.UeAmbr.Downlink = 6000000

	item := &types3gpp.PduSessResSetupReqItem{}
	item.PduSessionId = 5
	item.IsNasPduPrst = true
	item.NasPdu = "987654321"
	item.PduSessResSetupReqTrans = string([]byte{88, 32, 7, 161, 32, 16, 78, 32, 1, 240, 10, 180, 9, 242, 55, 37, 0, 0, 0, 8, 20, 24, 11, 204, 112, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 9, 192, 4, 127, 255, 255, 253, 82, 10, 225, 192, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 5, 192, 4, 127, 255, 255, 253, 113, 0, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18})
	item.Snssai = &types3gpp.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddPduSessResSetupReqItem(item)

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestPduSessResSetupReqMsg_Decode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{0, 29, 0, 128, 149, 0, 0, 5, 0, 74, 0, 111, 0, 64, 5, 9, 57, 56, 55, 54, 53, 52, 51, 50, 49, 66, 32, 17, 34, 51, 92, 88, 32, 7, 161, 32, 16, 78, 32, 1, 240, 10, 180, 9, 242, 55, 37, 0, 0, 0, 8, 20, 24, 11, 204, 112, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 9, 192, 4, 127, 255, 255, 253, 82, 10, 225, 192, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 5, 192, 4, 127, 255, 255, 253, 113, 0, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18, 0, 83, 64, 1, 9, 0, 38, 0, 10, 9, 49, 50, 51, 52, 53, 54, 55, 56, 57, 0, 10, 0, 2, 0, 1, 0, 85, 0, 2, 0, 2}

	cmnNgAp := codec.NewNgapCodec()
	cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))

	cmnNgAp.DecodeHeader(ossCtxtPtr)

	fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	msg := NewPduSessResSetupReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(msg)
}
