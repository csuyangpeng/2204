package ngapmsg

import (
	"fmt"
	"testing"

	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
)

func TestPduSessResRelRespMsg_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewPduSessResRelRespMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.AmfUeNGAPId = 5
	msg.RanUeNGAPId = 6

	plmn := types3gpp.PlmnID{}
	val := []byte{0x46, 0x0f, 0x01}
	plmn.SetValue(val, types3gpp.LittleEndian)

	msg.UserLocInfo.NrCgi.PlmnID = plmn
	msg.UserLocInfo.NrCgi.NrCellId = [types3gpp.SizeofNrCellID]byte{0x12, 0x34, 0x56, 0x78, 0x90}
	msg.UserLocInfo.Tai.Plmn = plmn
	msg.UserLocInfo.Tai.Tac = types3gpp.TAC{0x44, 0x55, 0x66}
	msg.UserLocInfo.IsTimeStampPrst = false

	item := &types3gpp.PduSessResRelRespItem{}
	item.PduSessionId = 5

	item.PduSessResRelRespTransfer = string([]byte{0})
	msg.AddPduSessResRelRespItem(item)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestPduSessResRelRespMsg_Decode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{}

	cmnNgAp := codec.NewNgapCodec()
	cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))

	cmnNgAp.DecodeHeader(ossCtxtPtr)

	fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	msg := NewPduSessResRelRespMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(msg)
}
