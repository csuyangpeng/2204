package types3gpp

import (
	"testing"
)

func TestImsi_StoreImsiString(t *testing.T) {
	//imsiString := "460010123456789"
	//
	//imsi := &Imsi{}
	//imsi.StoreImsiString(imsiString,2)
	//
	//if imsiString != imsi.String() {
	//	t.Errorf("store Imsi string error.")
	//}
	//fmt.Println(string(1))
}

func TestImsi_StoreBcdArray(t *testing.T) {
	bcdArray := []byte{15, 0x46, 0x00, 0x10, 0x12, 0x34, 0x56, 0x78, 0x90}
	imsi := &Imsi{}
	err := imsi.StoreBcdArray(bcdArray, 2)
	if err != nil {
		t.Errorf("failed to store bcd array. err：%s", err)
	}

	expImsiString := "460010123456789"
	if expImsiString != imsi.String() {
		t.Errorf("expect %s, but return %s", expImsiString, imsi.String())
	}
}

func TestImsi_StoreBcdArray_1(t *testing.T) {
	bcdArray := []byte{10, 0x46, 0x00, 0x10, 0x12, 0x34}
	imsi := &Imsi{}
	err := imsi.StoreBcdArray(bcdArray, 2)
	if err != nil {
		t.Errorf("failed to store bcd array. err：%s", err)
	}

	expImsiString := "4600101234"
	if expImsiString != imsi.String() {
		t.Errorf("expect %s, but return %s", expImsiString, imsi.String())
	}
}

func TestImsi_StoreBcdArray_2(t *testing.T) {
	bcdArray := []byte{6, 0x46, 0x00, 0x10}
	imsi := &Imsi{}
	err := imsi.StoreBcdArray(bcdArray, 3)
	if err != nil {
		t.Errorf("failed to store bcd array. err：%s", err)
	}

	expImsiString := "460010"
	if expImsiString != imsi.String() {
		t.Errorf("expect %s, but return %s", expImsiString, imsi.String())
	}
}

func TestImsi_GetMcc(t *testing.T) {
	bcdArray := []byte{15, 0x46, 0x00, 0x10, 0x12, 0x34, 0x56, 0x78, 0x90}
	imsi := &Imsi{}
	err := imsi.StoreBcdArray(bcdArray, 2)
	if err != nil {
		t.Errorf("failed to store bcd array. err：%s", err)
	}

	expMcc := 460
	if uint16(expMcc) != imsi.GetMcc() {
		t.Errorf("expect %d, but return %d", expMcc, imsi.GetMcc())
	}
}

func TestImsi_GetMnc_1(t *testing.T) {
	bcdArray := []byte{15, 0x46, 0x00, 0x10, 0x12, 0x34, 0x56, 0x78, 0x90}
	imsi := &Imsi{}
	err := imsi.StoreBcdArray(bcdArray, 2)
	if err != nil {
		t.Errorf("failed to store bcd array. err：%s", err)
	}

	expMnc := 1
	if uint16(expMnc) != imsi.GetMnc() {
		t.Errorf("expect %d, but return %d", expMnc, imsi.GetMnc())
	}
}

func TestImsi_GetMnc_2(t *testing.T) {
	bcdArray := []byte{15, 0x46, 0x00, 0x10, 0x12, 0x34, 0x56, 0x78, 0x90}
	imsi := &Imsi{}
	err := imsi.StoreBcdArray(bcdArray, 3)
	if err != nil {
		t.Errorf("failed to store bcd array. err：%s", err)
	}

	expMnc := 10
	if uint16(expMnc) != imsi.GetMnc() {
		t.Errorf("expect %d, but return %d", expMnc, imsi.GetMnc())
	}
}

func TestImsi_GetNAIString(t *testing.T) {
	bcdArray := []byte{15, 0x46, 0x00, 0x10, 0x12, 0x34, 0x56, 0x78, 0x90}
	imsi := &Imsi{}
	err := imsi.StoreBcdArray(bcdArray, 2)
	if err != nil {
		t.Errorf("failed to store bcd array. err：%s", err)
	}

	expNAIStr := "460010123456789@nai.5gc.mnc001.mcc460.3gppnetwork.org"

	naiStr := imsi.GetNAIString()
	if expNAIStr != naiStr {
		t.Errorf("expect %s, \nbut return %s", expNAIStr, naiStr)
	}
}

func TestImsi_StoreWithNAI(t *testing.T) {
	naiStr := "460010123456789@nai.5gc.mnc01.mcc460.3gppnetwork.org"
	imsi := &Imsi{}

	err := imsi.StoreWithNAI(naiStr, 2)
	if err != nil {
		t.Errorf("failed to store with NAI string. err：%s", err)
	}

	expImsiStr := "460010123456789"
	if expImsiStr != imsi.String() {
		t.Errorf("expect %s, but return %s", expImsiStr, imsi.String())
	}
}
