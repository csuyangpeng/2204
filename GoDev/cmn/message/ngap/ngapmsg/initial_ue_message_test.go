package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func TestInitialUeMessage_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewInitialUeMessage()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.RanUeNgapId = 1234
	var nasData string = "123456789987654321"
	msg.NasPdu = []byte(nasData)

	plmn := types3gpp.PlmnID{}
	val := []byte{0x46, 0x0f, 0x01}
	plmn.SetValue(val, types3gpp.LittleEndian)

	msg.UserLocInfo.NrCgi.PlmnID = plmn
	msg.UserLocInfo.NrCgi.NrCellId = [types3gpp.SizeofNrCellID]byte{0x12, 0x34, 0x56, 0x78, 0x90}
	msg.UserLocInfo.Tai.Plmn = plmn
	msg.UserLocInfo.Tai.Tac = types3gpp.TAC{0x44, 0x55, 0x66}
	msg.UserLocInfo.IsTimeStampPrst = false

	msg.IsStmsi5gPrst = true
	msg.Stmsi5G.SetTmsi(12345)
	msg.Stmsi5G.SetAmfSetID([]byte{0x12, 0x34})
	msg.Stmsi5G.SetAmfPointer(15)

	fmt.Println(msg)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestInitialUeMessage_Decode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()

	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{0x00, 0x0f, 0x40, 0x51, 0x00, 0x00, 0x06, 0x00, 0x55, 0x00, 0x05, 0xc0, 0x2c, 0x00, 0x00, 0x00, 0x00, 0x26, 0x00, 0x19, 0x18, 0x7e, 0x01, 0xc3, 0xdc, 0xd6, 0x06, 0x03, 0x7e, 0x00, 0x45, 0x11, 0x00, 0x0b, 0xf2, 0x64, 0xf0, 0x00, 0x05, 0x00, 0x83, 0x00, 0x00, 0x00, 0x00, 0x00, 0x79, 0x00, 0x0f, 0x40, 0x64, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x64, 0xf0, 0x00, 0x00, 0x00, 0x01, 0x00, 0x5a, 0x40, 0x01, 0x18, 0x00, 0x1a, 0x00, 0x07, 0x00, 0x20, 0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x70, 0x40, 0x01, 0x00}
	fmt.Println(encMsg)

	msg := NewInitialUeMessage()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	err := msg.Decode(encMsg)

	fmt.Println("result", err)
	fmt.Println(msg.Stmsi5G.AmfSetID[0], msg.Stmsi5G.AmfSetID[1])
	fmt.Println(msg.Stmsi5G.String())
	fmt.Println(msg)
}
