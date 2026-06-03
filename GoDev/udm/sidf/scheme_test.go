package sidf

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSUCIDeconcealing(t *testing.T) {

}

func TestSuciDecryptA(t *testing.T) {

	UEephPubKey, _ := hex.DecodeString("b2e92f836055a255837debf850b528997ce0201cb82adfe4be1f587d07d8457d")
	ciphertext, _ := hex.DecodeString("cb02352410")
	MacTag, _ := hex.DecodeString("cddd9e730ef3fa87")

	expectSupi, _ := hex.DecodeString("00012080f6")
	fmt.Printf("UEephPubKey:%#x\n", UEephPubKey)
	fmt.Printf("ciphertext:%x\n", ciphertext)
	fmt.Printf("MacTag:%x\n", MacTag)

	supi, err := SuciDecryptA(UEephPubKey, ciphertext, MacTag, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("supi:%x", supi)
	if !bytes.Equal(supi, expectSupi) {
		t.Error("Unexpected results")
	}
}

func TestHNwkeyID(t *testing.T) {

	UEephPubKey, _ := hex.DecodeString("b2e92f836055a255837debf850b528997ce0201cb82adfe4be1f587d07d8457d")
	ciphertext, _ := hex.DecodeString("cb02352410")
	MacTag, _ := hex.DecodeString("cddd9e730ef3fa87")

	//expectSupi, _ := hex.DecodeString("00012080f6")
	fmt.Printf("UEephPubKey:%#x\n", UEephPubKey)
	fmt.Printf("ciphertext:%x\n", ciphertext)
	fmt.Printf("MacTag:%x\n", MacTag)

	_, err := SuciDecryptA(UEephPubKey, ciphertext, MacTag, 1)
	if err.Error() != "ecies: MAC Tag verification error" {
		t.Error(err)
	}
	/*fmt.Printf("supi:%x", supi)
	if !bytes.Equal(supi, expectSupi) {
		t.Error("Unexpected results")
	}*/
}
