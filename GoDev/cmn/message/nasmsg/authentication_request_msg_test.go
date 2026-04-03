package nasmsg

import (
	"fmt"
	"testing"
)

func TestAuthenticationRequest_Encode(t *testing.T) {
	msg := AuthenticationRequestMsg{}

	msg.NgKSI.Ksi = 0
	msg.NgKSI.Tsc = false

	msg.Abba = []byte{0x00, 0x00}
	msg.Rand = [16]byte{0x47, 0x11, 0x47, 0x11, 0x47, 0x11, 0x47, 0x11, 0x47, 0x11, 0x47, 0x11, 0x47, 0x11, 0x47, 0x11}
	msg.Autn = [16]byte{0x11, 0x49, 0x52, 0x19, 0xc6, 0x70, 0xb9, 0xb9, 0xd8, 0xb9, 0x67, 0x6b, 0x45, 0xda, 0xd5, 0xfc}

	msg.IeFlags.Set(IeidAuthreqAutn)
	msg.IeFlags.Set(IeidAuthreqRand)

	buf, err := msg.Encode()
	if err != nil {
		fmt.Println(err)
	}
	out := fmt.Sprintf("%x", buf)
	expect := "7e0056000200002147114711471147114711471147114711201011495219c670b9b9d8b9676b45dad5fc"

	fmt.Printf("Encode Msg: %x\n", buf)
	fmt.Printf("Expect msg: %s\n", expect)

	if out != expect {
		t.Errorf("failed to encode.")
	}

	//0000   7e 00 56 00 02 00 00 21 47 11 47 11 47 11 47 11
	//0010   47 11 47 11 47 11 47 11 20 10 11 49 52 19 c6 70
	//0020   b9 b9 d8 b9 67 6b 45 da d5 fc
}
