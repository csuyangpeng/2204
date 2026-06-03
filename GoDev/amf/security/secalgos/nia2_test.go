package secalgos

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"
)

// ok
func TestNIA2_CMAC_Set2(t *testing.T) {
	Key, _ := hex.DecodeString("d3c5d592327fb11c4035c6680af8c6d1") // len 16 bytes

	Count, _ := hex.DecodeString("398a59b4") //4 bytes
	//Bearer, _ := hex.DecodeString("15")                            //5 bits
	Bearer := byte(0x1a) //5 bits
	//Direction, _ := hex.DecodeString("1")                          //1 bit
	Direction := byte(0x1)      //1 bit
	BLength := 64               // bits
	Mtext := "484583d5afe082ae" // MESSAGE

	Plaintext, _ := hex.DecodeString(Mtext) //64 bits
	tmpByte := Bearer<<3 | (Direction&1)<<2
	fmt.Printf("tmpByte:%T,%08b\n", tmpByte, tmpByte)
	fmt.Println("test Plaintext :", len(Plaintext))
	fmt.Printf("Plaintext:%08b\n", Plaintext)
	// Mlen = BLENGTH + 64

	Mac, err := Nia2CMAC(Key, Count, Bearer, Direction, Plaintext, BLength)
	if err != nil {
		t.Errorf("Nia2CMAC:%s", err)
	}
	mac_test, _ := hex.DecodeString("b93787e6493ff113ad73d3e01e826d73")
	fmt.Println("mac :", hex.EncodeToString(Mac[:4]))
	if !bytes.Equal(Mac, mac_test[:4]) {
		t.Errorf("%s: mac = %s", Mac, mac_test)
		return
	}
	// 01110000              00001000
	tmpByte = 0x0e<<3 | 0x03<<2 //(0x04 & direction << 2)

	fmt.Printf("%08b\n", 0x0e<<3)
	fmt.Printf("%08b\n", 0x03<<7)
	fmt.Printf("%08b\n", tmpByte)

	ccdd := 0x0e << 3
	ffd := Bearer << 3
	fmt.Printf("%T\n", ccdd)
	fmt.Printf("%T\n", ffd)
}

func TestMacVerify(t *testing.T) {
	dstMac, err := hex.DecodeString("b93787e6")
	if err != nil {
		t.Errorf("decode string:%s", err)
	}
	srcMac, err := hex.DecodeString("b93787e6")
	if err != nil {
		t.Errorf("decode string:%s", err)
	}
	fmt.Println(dstMac, srcMac)
	if !MacVerify(dstMac, srcMac) {
		t.Errorf("%s: mac = %s", dstMac, srcMac)
	}
}

// Entry validity check
func TestNIA2ValidityCheck(t *testing.T) {
	Key, _ := hex.DecodeString("d3c5d592327fb11c4035c6680af8c6d1") // len 16 bytes

	Count, _ := hex.DecodeString("398a59b4") //4 bytes // hex.DecodeString 奇数hex串丢弃最后的一个字符
	//Bearer, _ := hex.DecodeString("15")                            //5 bits
	Bearer := byte(0x1a) //5 bits
	//Direction, _ := hex.DecodeString("1")                          //1 bit
	Direction := byte(0x1)      //1 bit
	BLength := 64               // bits
	Mtext := "484583d5afe082ae" // MESSAGE

	Plaintext, _ := hex.DecodeString(Mtext) //64 bits
	tmpByte := Bearer<<3 | (Direction&1)<<2
	fmt.Printf("tmpByte:%T,%08b\n", tmpByte, tmpByte)
	fmt.Println("test Plaintext :", len(Plaintext))
	fmt.Printf("Plaintext:%08b\n", Plaintext)
	// Mlen = BLENGTH + 64

	Mac, err := Nia2CMAC(Key, Count, Bearer, Direction, Plaintext, BLength)
	if err != nil {
		t.Errorf("Nia2CMAC:%s", err)
	}
	mac_test, _ := hex.DecodeString("b93787e6493ff113ad73d3e01e826d73")
	fmt.Println("mac :", hex.EncodeToString(Mac[:4]))
	if !bytes.Equal(Mac, mac_test[:4]) {
		t.Errorf("%s: mac = %s", Mac, mac_test)
		return
	}
}

func BenchmarkNIA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Key, _ := hex.DecodeString("d3c5d592327fb11c4035c6680af8c6d1") // len 16 bytes

		Count, _ := hex.DecodeString("398a59b4") //4 bytes // hex.DecodeString 奇数hex串丢弃最后的一个字符
		//Bearer, _ := hex.DecodeString("15")                            //5 bits
		Bearer := byte(0x1a) //5 bits
		//Direction, _ := hex.DecodeString("1")                          //1 bit
		Direction := byte(0x1)      //1 bit
		BLength := 64               // bits
		Mtext := "484583d5afe082ae" // MESSAGE

		Plaintext, _ := hex.DecodeString(Mtext) //64 bits
		// Mlen = BLENGTH + 64

		Mac, err := Nia2CMAC(Key, Count, Bearer, Direction, Plaintext, BLength)
		if err != nil {
			b.Errorf("Nia2CMAC:%s", err)
		}
		mac_test, _ := hex.DecodeString("b93787e6493ff113ad73d3e01e826d73")
		if !bytes.Equal(Mac, mac_test[:4]) {
			b.Errorf("%s: mac = %s", Mac, mac_test)
			return
		}
		//	BenchmarkNIA2-4   	 1000000	      1368 ns/op
	}
}
