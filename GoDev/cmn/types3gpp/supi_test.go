package types3gpp

import (
	"lite5gc/cmn/mcctbl"
	"testing"
)

func TestSupi_SetImsi(t *testing.T) {
	bcdArray := []byte{15, 0x46, 0x00, 0x10, 0x12, 0x34, 0x56, 0x78, 0x90}
	imsi := &Imsi{}
	err := imsi.StoreBcdArray(bcdArray, 2)
	if err != nil {
		t.Errorf("failed to store bcd array. err：%s", err)
	}

	supi := &Supi{}
	supi.SetImsi(imsi)
	expStr := "type: imsi, value: 460010123456789"
	rtStr := supi.String()

	if expStr != rtStr {
		t.Errorf("exp: %s, but return :%s", expStr, rtStr)
	}
}

func TestSupi_SetNAI(t *testing.T) {
	naiStr := "460120123456789@nai.5gc.mnc001.mcc460.3gppnetwork.org"

	mccTbl := &mcctbl.MccTable{}
	_ = mccTbl.AddItem("460", 2)

	mccStr := naiStr[0:3]
	mncLen, _ := mccTbl.GetMncLen(mccStr)

	supi := &Supi{}
	err := supi.SetNAI(naiStr, mncLen)
	if err != nil {
		t.Errorf("failed to set NAI.")
	}

	expStr := "type: nai, value: 460120123456789@nai.5gc.mnc012.mcc460.3gppnetwork.org"
	rtStr := supi.String()

	if expStr != rtStr {
		t.Errorf("exp: %s, but return :%s", expStr, rtStr)
	}
}

func TestSupi_GetImsi(t *testing.T) {
	bcdArray := []byte{15, 0x46, 0x00, 0x10, 0x12, 0x34, 0x56, 0x78, 0x90}
	imsi := &Imsi{}
	err := imsi.StoreBcdArray(bcdArray, 2)
	if err != nil {
		t.Errorf("failed to store bcd array. err：%s", err)
	}

	supi := &Supi{}
	supi.SetImsi(imsi)
	rtImsi := supi.GetImsi()

	if *imsi != *rtImsi {
		t.Errorf("exp: %v, but return :%v", imsi, rtImsi)
	}
}

func TestSupi_GetNAI(t *testing.T) {
	naiStr := "460120123456789@nai.5gc.mnc001.mcc460.3gppnetwork.org"

	mccTbl := &mcctbl.MccTable{}
	_ = mccTbl.AddItem("460", 2)

	mccStr := naiStr[0:3]
	mncLen, _ := mccTbl.GetMncLen(mccStr)

	supi := &Supi{}
	err := supi.SetNAI(naiStr, mncLen)
	if err != nil {
		t.Errorf("failed to set NAI.")
	}

	expStr := "460120123456789@nai.5gc.mnc012.mcc460.3gppnetwork.org"
	rtStr := supi.GetNAI()

	if expStr != rtStr {
		t.Errorf("exp: %s, but return :%s", expStr, rtStr)
	}
}
