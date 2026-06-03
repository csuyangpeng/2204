package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func Test_NgSetupResponseMsg_encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewNgSetupResponseMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	plmn := types3gpp.PlmnID{}
	val := []byte{0x46, 0x0f, 0x01}
	plmn.SetValue(val, types3gpp.LittleEndian)

	msg.AmfName = "foxconn"
	msg.RelativeAmfCapacity = 200

	guami := types3gpp.Guami{
		PlmnId: plmn,
	}
	guami.AmfId.SetAmfRegionID(8)
	setID := [2]uint8{1, 1}
	guami.AmfId.SetAmfSetID(setID[:])
	guami.AmfId.SetAmfPointer(1)
	msg.AddServedGuami(&guami)

	snssai := types3gpp.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	bplmn := types3gpp.BPlmn{Plmn: plmn}
	bplmn.AddSnssai(&snssai)
	bplmn.AddSnssai(&snssai)

	msg.AddPlmnSupport(&bplmn)

	fmt.Println(msg)
	fmt.Println("----------------------")
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func Test_NgSetupResponseMsg_decode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{32, 21, 0, 47, 0, 0, 4, 0, 80, 0, 10, 0, 100, 240, 16, 0, 1, 0, 136, 4, 64, 0, 86, 64, 1, 200, 0, 96, 0, 8, 0, 0, 100, 240, 16, 8, 1, 2, 0, 1, 0, 9, 3, 0, 102, 111, 120, 99, 111, 110, 110}
	fmt.Println(encMsg)

	setupResponseMsg := NewNgSetupResponseMsg()
	setupResponseMsg.SetOssCodecCtxt(ossCtxtPtr)
	decode := setupResponseMsg.Decode(encMsg)
	fmt.Println("result", decode)
}
