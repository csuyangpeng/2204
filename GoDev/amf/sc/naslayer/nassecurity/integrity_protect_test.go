package nassecurity

import (
	"encoding/hex"
	"fmt"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func TestGenerateIntegrity(t *testing.T) {
	knint, _ := hex.DecodeString("9ba9057f9b6b763f671eeb53e1c5206e")
	//knint, _ := hex.DecodeString("e29e2a782c6d9354e16fbb651129b6b2")  //supi string; algo id 0x20
	//knint, _ := hex.DecodeString("4d2030914130b84e8909e0c32062e9c6")
	nasMsg, _ := hex.DecodeString("007e005d020102f0f0e1")
	_, mac := GenerateIntegrityMac(types3gpp.NIA2, knint, 0, 0, nasMsg)
	fmt.Printf("NIA2 MAC should be (56a7c495), return(%x)", mac)
}
