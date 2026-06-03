package secalgos

import (
	"bytes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCTR_AES_NEA2_Set2(t *testing.T) {

	Key, _ := hex.DecodeString("2bd6459f82c440e0952c49104805ff48") // len 16 bytes
	Count, _ := hex.DecodeString("c675a64b")                       //4 bytes
	//Bearer, _ := hex.DecodeString("15")                            //5 bits
	Bearer := byte(0x0c) //5 bits
	//Direction, _ := hex.DecodeString("1")                          //1 bit
	Direction := byte(0x1) //1 bit
	Length := 798          // bits
	ptext := "7ec61272743bf1614726446a6c38ced166f6ca76eb5430044286346cef130f92" +
		"922b03450d3a9975e5bd2ea0eb55ad8e1b199e3ec4316020e9a1b285e7627953" +
		"59b7bdfd39bef4b2484583d5afe082aee638bf5fd5a606193901a08f4ab41aab" +
		"9b134880"
	ctext := "5961605353c64bdca15b195e288553a910632506d6200aa790c4c806c99904cf" +
		"2445cc50bb1cf168a49673734e081b57e324ce5259c0e78d4cd97b870976503c" +
		"0943f2cb5ae8f052c7b7d392239587b8956086bcab18836042e2e6ce42432a17" +
		"105c53d0"
	//T1 := "c675a64b640000000000000000000000"
	Plaintext, _ := hex.DecodeString(ptext)  //798 bits
	Ciphertext, _ := hex.DecodeString(ctext) //798 bits

	resultEncrypted, err := Nea2Encrypt(Key, Count, Bearer, Direction, Plaintext, Length)
	if err != nil {
		t.Errorf("Nea2Encrypt:%s", err)
	}

	fmt.Println("test Plaintext :", len(Plaintext))
	fmt.Println("test Ciphertext:", len(Ciphertext))
	fmt.Printf("resultEncrypted:%x\n", resultEncrypted)

	if out := resultEncrypted[0:len(Ciphertext)]; !bytes.Equal(out, Ciphertext) {
		t.Errorf("CTR\ninpt %x\nhave %x\nwant %x", Plaintext, out, Ciphertext)
	}

	// 解密
	resultDecrypted, err := Nea2Decrypt(Key, Count, Bearer, Direction, Ciphertext, Length)
	if err != nil {
		t.Errorf("Nea2Decrypt:%s", err)
	}
	fmt.Printf("resultDecrypted:%x\n", resultDecrypted)
	if out := resultDecrypted[0:len(Ciphertext)]; !bytes.Equal(out, Plaintext) {
		t.Errorf("CTR\ninpt %x\nhave %x\nwant %x", Ciphertext, out, Plaintext)
	}

}

func TestCTR_AES_NEA2_Msg(t *testing.T) {

	Key, _ := hex.DecodeString("2bd6459f82c440e0952c49104805ff48") // len 16 bytes
	Count, _ := hex.DecodeString("c675a64b")                       //4 bytes
	//Bearer, _ := hex.DecodeString("15")                            //5 bits
	Bearer := byte(0x0c) //5 bits
	//Direction, _ := hex.DecodeString("1")                          //1 bit
	Direction := byte(0x1) //1 bit
	Length := 0            // bits
	Plaintext := []byte("hello msg send")
	Length = 8 * len(Plaintext)

	resultEncrypted, err := Nea2Encrypt(Key, Count, Bearer, Direction, Plaintext, Length)
	if err != nil {
		t.Errorf("Nea2Encrypt:%s", err)
	}

	fmt.Println("test Plaintext :", len(Plaintext))
	fmt.Printf("resultEncrypted:%x\n", resultEncrypted)

	// 解密
	resultDecrypted, err := Nea2Decrypt(Key, Count, Bearer, Direction, resultEncrypted, Length)
	if err != nil {
		t.Errorf("Nea2Decrypt:%s", err)
	}
	fmt.Printf("resultDecrypted:%x\n", resultDecrypted)
	if out := resultDecrypted[0:len(Plaintext)]; !bytes.Equal(out, Plaintext) {
		t.Errorf("CTR\ninpt %x\nhave %x\nwant %x", Plaintext, out, Plaintext)
	}
	fmt.Printf("CTR\ninpt :%s\nhave :%s\n", Plaintext, resultDecrypted)
}

func TestNewCipher_CTR_AES_NEA2_Set2(t *testing.T) {
	Key, _ := hex.DecodeString("2bd6459f82c440e0952c49104805ff48") // len 16 bytes
	Count, _ := hex.DecodeString("c675a64b")                       //4 bytes
	//Bearer, _ := hex.DecodeString("15")                            //5 bits
	Bearer := byte(0x0c) //5 bits
	//Direction, _ := hex.DecodeString("1")                          //1 bit
	Direction := byte(0x1) //1 bit
	Length := 798          // bits
	ptext := "7ec61272743bf1614726446a6c38ced166f6ca76eb5430044286346cef130f92" +
		"922b03450d3a9975e5bd2ea0eb55ad8e1b199e3ec4316020e9a1b285e7627953" +
		"59b7bdfd39bef4b2484583d5afe082aee638bf5fd5a606193901a08f4ab41aab" +
		"9b134880"
	ctext := "5961605353c64bdca15b195e288553a910632506d6200aa790c4c806c99904cf" +
		"2445cc50bb1cf168a49673734e081b57e324ce5259c0e78d4cd97b870976503c" +
		"0943f2cb5ae8f052c7b7d392239587b8956086bcab18836042e2e6ce42432a17" +
		"105c53d0"
	//T1 := "c675a64b640000000000000000000000"
	Plaintext, _ := hex.DecodeString(ptext)  //798 bits
	Ciphertext, _ := hex.DecodeString(ctext) //798 bits

	nea2, err := NewNea2(Key, Count, Bearer, Direction)
	if err != nil {
		t.Errorf("New Nea2:%s", err)
	}

	fmt.Println("test Plaintext :", len(Plaintext))
	fmt.Println("test Ciphertext:", len(Ciphertext))

	resultEncrypted, err := nea2.Encrypt(Plaintext, Length)
	if err != nil {
		t.Errorf("Nea2 Encrypt:%s", err)
	}
	fmt.Printf("resultEncrypted:%x\n", resultEncrypted)

	if out := resultEncrypted[0:len(Ciphertext)]; !bytes.Equal(out, Ciphertext) {
		t.Errorf("CTR\ninpt %x\nhave %x\nwant %x", Plaintext, out, Ciphertext)
	}

	// 解密
	resultDecrypted, err := nea2.Decrypt(Ciphertext, Length)
	if err != nil {
		t.Errorf("Nea2 Decrypt:%s", err)
	}
	fmt.Printf("resultDecrypted:%x\n", resultDecrypted)
	if out := resultDecrypted[0:len(Ciphertext)]; !bytes.Equal(out, Plaintext) {
		t.Errorf("CTR\ninpt %x\nhave %x\nwant %x", Ciphertext, out, Plaintext)
	}

}

// 分段加密
func TestNewCipher_CTR_AES_NEA2_Set2_Stream(t *testing.T) {
	Key, _ := hex.DecodeString("2bd6459f82c440e0952c49104805ff48") // len 16 bytes
	Count, _ := hex.DecodeString("c675a64b")                       //4 bytes
	//Bearer, _ := hex.DecodeString("15")                            //5 bits
	Bearer := byte(0x0c) //5 bits
	//Direction, _ := hex.DecodeString("1")                          //1 bit
	Direction := byte(0x1) //1 bit
	Length := 798          // bits
	ptext := "7ec61272743bf1614726446a6c38ced166f6ca76eb5430044286346cef130f92" +
		"922b03450d3a9975e5bd2ea0eb55ad8e1b199e3ec4316020e9a1b285e7627953" +
		"59b7bdfd39bef4b2484583d5afe082aee638bf5fd5a606193901a08f4ab41aab" +
		"9b134880"
	ctext := "5961605353c64bdca15b195e288553a910632506d6200aa790c4c806c99904cf" +
		"2445cc50bb1cf168a49673734e081b57e324ce5259c0e78d4cd97b870976503c" +
		"0943f2cb5ae8f052c7b7d392239587b8956086bcab18836042e2e6ce42432a17" +
		"105c53d0"
	//T1 := "c675a64b640000000000000000000000"
	Plaintext, _ := hex.DecodeString(ptext)  //798 bits
	Ciphertext, _ := hex.DecodeString(ctext) //798 bits

	nea2, err := NewNea2(Key, Count, Bearer, Direction)
	if err != nil {
		t.Errorf("New Nea2:%s", err)
	}

	fmt.Println("test Plaintext :", len(Plaintext))
	fmt.Println("test Ciphertext:", len(Ciphertext))
	bLength := Length
	//resultEncrypted, err := nea2.Encrypt(Plaintext, Length)
	if bLength > len(Plaintext)*8 {
		t.Errorf("invalid bLength")
	}

	keyStream := cipher.NewCTR(nea2.c, nea2.iv)
	//n.keyStream = keyStream
	encrypted := make([]byte, len(Plaintext))
	// 分段加密
	fmt.Println("text length:", len(Plaintext))
	keyStream.XORKeyStream(encrypted[:len(Plaintext)/4], Plaintext[:len(Plaintext)/4])
	keyStream.XORKeyStream(encrypted[len(Plaintext)/4:len(Plaintext)/2], Plaintext[len(Plaintext)/4:len(Plaintext)/2])
	keyStream.XORKeyStream(encrypted[len(Plaintext)/2:], Plaintext[len(Plaintext)/2:])
	// 清除多余bit
	encrypted[len(encrypted)-1] = (encrypted[len(encrypted)-1] >>
		uint(len(encrypted)*8-bLength)) <<
		uint(len(encrypted)*8-bLength)

	resultEncrypted := encrypted
	fmt.Printf("resultEncrypted:%x\n", resultEncrypted)

	if out := resultEncrypted[0:len(Ciphertext)]; !bytes.Equal(out, Ciphertext) {
		t.Errorf("CTR\ninpt %x\nhave %x\nwant %x", Plaintext, out, Ciphertext)
	}

	// 解密
	resultDecrypted, err := nea2.Decrypt(Ciphertext, Length)
	if err != nil {
		t.Errorf("Nea2 Decrypt:%s", err)
	}
	fmt.Printf("resultDecrypted:%x\n", resultDecrypted)
	if out := resultDecrypted[0:len(Ciphertext)]; !bytes.Equal(out, Plaintext) {
		t.Errorf("CTR\ninpt %x\nhave %x\nwant %x", Ciphertext, out, Plaintext)
	}

}

func BenchmarkNEA2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Key, _ := hex.DecodeString("2bd6459f82c440e0952c49104805ff48") // len 16 bytes
		Count, _ := hex.DecodeString("c675a64b")                       //4 bytes
		//Bearer, _ := hex.DecodeString("15")                            //5 bits
		Bearer := byte(0x0c) //5 bits
		//Direction, _ := hex.DecodeString("1")                          //1 bit
		Direction := byte(0x1) //1 bit
		Length := 0            // bits
		Plaintext := []byte("hello msg send")
		Length = 8 * len(Plaintext)
		// 加密
		resultEncrypted, err := Nea2Encrypt(Key, Count, Bearer, Direction, Plaintext, Length)
		if err != nil {
			b.Errorf("Nea2Encrypt:%s", err)
		}

		// 解密
		resultDecrypted, err := Nea2Decrypt(Key, Count, Bearer, Direction, resultEncrypted, Length)
		if err != nil {
			b.Errorf("Nea2Decrypt:%s", err)
		}
		if out := resultDecrypted[0:len(Plaintext)]; !bytes.Equal(out, Plaintext) {
			b.Errorf("CTR\ninpt %x\nhave %x\nwant %x", Plaintext, out, Plaintext)
		}
		//	BenchmarkNEA2-4   	  500000	      2447 ns/op
	}
}
