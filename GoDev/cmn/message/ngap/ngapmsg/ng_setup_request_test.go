package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func TestCreateNgSetupRequestMsg_1(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewNgSetupRequestMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	plmn := types3gpp.PlmnID{}
	val := []byte{0x46, 0x0f, 0x01}
	plmn.SetValue(val, types3gpp.LittleEndian)

	gnbID := types3gpp.GlobalGNBID{
		Plmn:      plmn,
		GNBLength: 24,
		GNBID:     [types3gpp.SizeofGNBID]byte{0x12, 0x34, 0x56, 0x78},
	}
	msg.GRanNodeID.GNBID = gnbID

	msg.RanNodeName = "shenzhen01"
	msg.RanNodeNamePrst = true
	msg.DefPagingDRX = 4

	snssai := types3gpp.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	bplmn := types3gpp.BPlmn{Plmn: plmn}
	bplmn.AddSnssai(&snssai)
	bplmn.AddSnssai(&snssai)

	supportedTA := types3gpp.SupportedTA{
		Tac: types3gpp.TAC{0x44, 0x55, 0x66}}
	supportedTA.AddBPlmn(&bplmn)

	msg.AddSupportTA(&supportedTA)
	fmt.Println(msg)

	encodeMsg := msg.Encode()

	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestDecode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{0, 21, 0, 60, 0, 0, 4, 0, 102, 0, 15, 0, 0, 68, 85, 102, 0, 100, 240, 16, 0, 1, 0, 136, 4, 64, 0, 27, 0, 8, 0, 100, 240, 16, 16, 18, 52, 86, 0, 82, 64, 12, 4, 128, 115, 104, 101, 110, 122, 104, 101, 110, 48, 49, 0, 21, 64, 6, 192, 4, 127, 255, 255, 251}
	fmt.Println(encMsg)

	cmnNgAp := codec.NewNgapCodec()
	cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))

	cmnNgAp.DecodeHeader(ossCtxtPtr)

	fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	setupRequestMsg := NewNgSetupRequestMsg()
	setupRequestMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := setupRequestMsg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Println(setupRequestMsg)
}
