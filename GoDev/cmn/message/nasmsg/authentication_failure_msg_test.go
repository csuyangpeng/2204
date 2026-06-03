package nasmsg

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAuthenticationFailureMsg_Decode(t *testing.T) {
	//nasmsg := []byte{15300e0eb163bd17ddc653a4510b0e3497}
	//nasmsg := []byte{0x15, 0x30, 0x0e, 0x0e, 0xb1, 0x63, 0xbd, 0x17, 0xdd, 0xc6, 0x53, 0xa4, 0x51, 0x0b, 0x0e, 0x34, 0x97}
	//
	//msgbuf := bytes.NewReader(nasmsg)
	//
	//authFailure := AuthenticationFailureMsg{}
	//authFailure.Reset()
	//
	//err := authFailure.Decode(msgbuf)
	//if err != nil {
	//	t.Errorf("failed to decode msg")
	//}
	//
	//fmt.Println(authFailure.String())
	//2950986c08806db6794ce9349ed36e7c
	var data [10]byte

	data[0] = 130

	data[1] = 2

	fmt.Println(data)

	var str = byteString(data[:])

	fmt.Println(str)
}

func byteString(p []byte) string {
	var s string
	for i := 0; i < len(p); i++ {

		s += strconv.Itoa(int(p[i]))
	}
	return s
}
