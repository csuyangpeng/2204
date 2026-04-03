package types3gpp

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSuci_GetHomeNetworkId(t *testing.T) {
	//ie := []byte{0x00, 0x0d, 0x01, 0x64, 0x00, 0x00, 0x0f, 0xff, 0x00, 0x00, 0x32, 0x54, 0x06, 0x00, 0xf1}
	ie := []byte{0, 13, 1, 100, 0, 0, 0, 0, 0, 0, 50, 84, 6, 0, 241, 46, 2, 240, 240}

	suci := Suci{}
	suci.reset()

	msg := bytes.NewReader(ie)
	suci.Decode(msg)
	fmt.Println("suci = ", suci)
	imsi, err := suci.GetImsi()
	if err != nil {
		fmt.Println("erroor ", err)
	}

	fmt.Println("imsi = ", imsi.String())

}

// profile A
// MNC 2 digits
func TestSuci_Decode_A(t *testing.T) {
	expectImsi := "460000234560001"
	//ieSuci, _ := hex.DecodeString("0164f0000000010100f9de3c94c066c065968046ced38ccf2fa75e14c15fc2d2382ff83aaa522c781d0c06d0ddd02cc5dc5726d220")
	// 5GS mobile identity:
	ie, _ := hex.DecodeString("00350164f0000000010100f9de3c94c066c065968046ced38ccf2fa75e14c15fc2d2382ff83aaa522c781d0c06d0ddd02cc5dc5726d220")

	suci := Suci{}
	suci.reset()

	msg := bytes.NewReader(ie)
	suci.DecodeConcealed(msg)
	fmt.Println("suci = ", suci)

	//output := suci.GetSchemeOutputA()
	//MsinHex, err := sidf.SuciDecryptA(output.EphPublicKey, output.Ciphertext, output.MacTag,0)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println("MSINHex = ", MsinHex)
	/*imsi, err := suci.GetImsi()
	if err != nil {
		fmt.Println("erroor ", err)
	}*/
	// todo 支持9位数字与10位数字的号码，任一位数需根据imsiHex长度判断

	//suci.SetMsinHex2SchemeOutput(MsinHex)

	imsi, _ := suci.GetImsi()
	fmt.Println("MNC  = ", imsi.GetMncBytes())
	fmt.Println("IMSI = ", imsi.String())
	fmt.Println("MSIN = ", imsi.GetMsIn())
	fmt.Println("int MSIN = ", imsi.GetMsInValue())
	if expectImsi != imsi.String() {
		t.Error("Unexpected result")
	}
}

// msin is Odd digit
// MNC 3 digits
func TestSuciMsin9_Decode_A(t *testing.T) {
	expectImsi := "460000234560001"
	// 5GS mobile identity:
	ie, _ := hex.DecodeString("0035016400000000010160f357b44e9de67ee9852e577152cdb41a7c461bc545ae75c5e142959273316e50c79935dfb59a6334fa4a9417")

	suci := Suci{}
	suci.reset()

	msg := bytes.NewReader(ie)
	suci.DecodeConcealed(msg)
	fmt.Println("suci = ", suci)

	//output := suci.GetSchemeOutputA()
	//MsinHex, err := sidf.SuciDecryptA(output.EphPublicKey, output.Ciphertext, output.MacTag,0)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Printf("MSINHex = %#x\n", MsinHex)
	/*imsi, err := suci.GetImsi()
	if err != nil {
		fmt.Println("erroor ", err)
	}*/
	// todo 支持9位数字与10位数字的号码，任一位数需根据imsiHex长度判断

	//suci.SetMsinHex2SchemeOutput(MsinHex)

	imsi, _ := suci.GetImsi()
	fmt.Println("MNC  = ", imsi.GetMncBytes())
	fmt.Println("IMSI = ", imsi.String())
	fmt.Println("MSIN = ", imsi.GetMsIn())
	fmt.Println("int MSIN = ", imsi.GetMsInValue())
	if expectImsi != imsi.String() {
		t.Error("Unexpected result")
	}
}

// msin is 10 digits
// MNC 2 digits
func TestSuciMsin10(t *testing.T) {
	expectImsi := "460000234560001"
	// 5GS mobile identity:
	ie, _ := hex.DecodeString("000d0164f000000000002043650010")
	suci := Suci{}
	suci.reset()

	msg := bytes.NewReader(ie)
	suci.DecodeConcealed(msg)
	fmt.Println("suci = ", suci)

	imsi, _ := suci.GetImsi()
	fmt.Println("MNC  = ", imsi.GetMncBytes())
	fmt.Println("IMSI = ", imsi.String())
	fmt.Println("MSIN = ", imsi.GetMsIn())
	fmt.Println("int MSIN = ", imsi.GetMsInValue())
	if expectImsi != imsi.String() {
		t.Error("Unexpected result")
	}
}

// msin is 9 digits
// MNC 3 digits
func TestSuciMsin9(t *testing.T) {
	expectImsi := "460000234560001"
	// 5GS mobile identity:
	ie, _ := hex.DecodeString("000d016400000000000032540600f1")
	suci := Suci{}
	suci.reset()

	msg := bytes.NewReader(ie)
	suci.DecodeConcealed(msg)
	fmt.Println("suci = ", suci)

	imsi, _ := suci.GetImsi()
	fmt.Println("MNC  = ", imsi.GetMncBytes())
	fmt.Println("IMSI = ", imsi.String())
	fmt.Println("MSIN = ", imsi.GetMsIn())
	fmt.Println("int MSIN = ", imsi.GetMsInValue())
	if expectImsi != imsi.String() {
		t.Error("Unexpected result")
	}
}

// msin is 8 digits
// MNC 2 digits
func TestSuciMsin8(t *testing.T) {
	expectImsi := "4600002345601"
	// 5GS mobile identity:
	ie, _ := hex.DecodeString("000c0164f0000000000020436510")
	suci := Suci{}
	suci.reset()

	msg := bytes.NewReader(ie)
	suci.DecodeConcealed(msg)
	fmt.Println("suci = ", suci)

	imsi, _ := suci.GetImsi()
	fmt.Println("MNC  = ", imsi.GetMncBytes())
	fmt.Println("IMSI = ", imsi.String())
	fmt.Println("MSIN = ", imsi.GetMsIn())
	fmt.Println("int MSIN = ", imsi.GetMsInValue())
	if expectImsi != imsi.String() {
		t.Error("Unexpected result")
	}
}
