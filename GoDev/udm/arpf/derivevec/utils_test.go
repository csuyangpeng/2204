package derivevec

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"testing"
)

func TestGenerateRandomBytes(t *testing.T) {
	rand, err := GenerateRandomBytes(16)
	if err != nil {
		logs.Error("failed to GenerateRandomBytes")
	}
	fmt.Printf("%x", rand)
}

func TestKdfDerivation(t *testing.T) {
	key := [16]byte{0x46, 0x5b, 0x5c, 0xe8, 0xb1, 0x99, 0xb4, 0x9f, 0xaa, 0x5f, 0x0a, 0x2e, 0xe2, 0x38, 0xa6, 0xbc}
	fc := 0x6A
	pn := [][]byte{{0x31, 0x32, 0x33, 0x34, 0x35},
		{0xaa, 0x68, 0x9c, 0x64, 0x83, 0x70}}
	ln := []uint16{5, 6}
	err, output := KdfDerivation(key[:], byte(fc), pn, ln, 2)
	if err != nil {
		logs.Error("failed to KdfDerivation")
	}

	fmt.Printf("%x", output)
}
