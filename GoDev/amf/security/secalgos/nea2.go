package secalgos

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
)

// 3GPP TS 33.501 V15.3.1 (2018-12)
// NIST SP 800-38A

// D.2	Ciphering algorithms
// D.2.1.3	128-NEA2
const (
	AuthKeyLengthBytes   = 16
	AuthKeyLengthBits    = 128
	AuthCountLengthBytes = 4
	AuthBearerMax        = 0x1f
	AuthDirectionMax     = 0x01
)

//Key, _ := hex.DecodeString("54f4e2e04c83786eec8fb5abe8e36566") // len 16 bytes
//Count, _ := hex.DecodeString("aca4f50f")                       //4 bytes
////Bearer, _ := hex.DecodeString("15")                            //5 bits
//Bearer := byte(0x0b) //5 bits
////Direction, _ := hex.DecodeString("1")                          //1 bit
//Direction := byte(0x0) //1 bit
////Length :=  bits
func Nea2Encrypt(key []byte, count []byte, bearer byte, direction byte, plaintext []byte, bLength int) (
	ciphertext []byte, err error) {
	// Entry validity check
	if len(key) != AuthKeyLengthBytes {
		err = fmt.Errorf("Key length error, current value %d, expected value %d",
			len(key), AuthKeyLengthBytes)
		return nil, err
	}
	if len(count) != AuthCountLengthBytes {
		err = fmt.Errorf("Count length error, current value %d, expected value %d",
			len(count), AuthCountLengthBytes)
		return nil, err
	}
	//  two-hex-digit value in the range 00 to 1F inclusive
	if bearer > AuthBearerMax {
		err = fmt.Errorf("Bearer value error, current value %#x, expected value in the range 00 to 1F inclusive",
			bearer)
		return nil, err
	}
	// The DIRECTION bit shall be 0 for uplink and 1 for downlink
	if direction > AuthDirectionMax {
		err = fmt.Errorf("Direction value error, current value %#x, expected value in the range 0 to 1 inclusive",
			direction)
		return nil, err
	}
	// plaintext and bLength consistency check
	if len(plaintext)*8 < bLength {
		err = fmt.Errorf("Plaintext length error, current value %d, expected value %d",
			bLength, len(plaintext)*8)
		return nil, err
	}

	// T1 构造
	//CounterT1 := Count + Bearer + Direction + padding(90个0)
	CounterT1 := make([]byte, 16)
	CounterT1 = append(CounterT1[:0], count...)
	tmpByte := bearer<<3 | (direction&0x1)<<2
	CounterT1 = append(CounterT1, tmpByte)
	tmp8Byte := make([]byte, 11)
	CounterT1 = append(CounterT1, tmp8Byte...)

	//fmt.Printf("CounterT1:%x\n", CounterT1)

	//Key128 := key
	iv := CounterT1

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Failure to New Cipher")
	}

	ctr := cipher.NewCTR(c, iv)
	encrypted := make([]byte, len(plaintext))

	ctr.XORKeyStream(encrypted, plaintext)
	// 清除多余bit
	encrypted[len(encrypted)-1] = (encrypted[len(encrypted)-1] >>
		uint(len(encrypted)*8-bLength)) <<
		uint(len(encrypted)*8-bLength)

	return encrypted, nil

}

//Key, _ := hex.DecodeString("54f4e2e04c83786eec8fb5abe8e36566") // len 16 bytes
//Count, _ := hex.DecodeString("aca4f50f")                       //4 bytes
////Bearer, _ := hex.DecodeString("15")                            //5 bits
//Bearer := byte(0x0b) //5 bits
////Direction, _ := hex.DecodeString("1")                          //1 bit
//Direction := byte(0x0) //1 bit
////Length := 3861 bits
// Entry validity check
func Nea2Decrypt(key []byte, count []byte, bearer byte, direction byte, ciphertext []byte, bLength int) (
	plaintext []byte, err error) {
	if len(key) != AuthKeyLengthBytes {
		err = fmt.Errorf("Key length error, current value %d, expected value %d",
			len(key), AuthKeyLengthBytes)
		return nil, err
	}
	if len(count) != AuthCountLengthBytes {
		err = fmt.Errorf("Count length error, current value %d, expected value %d",
			len(count), AuthCountLengthBytes)
		return nil, err
	}
	//  two-hex-digit value in the range 00 to 1F inclusive
	if bearer > AuthBearerMax {
		err = fmt.Errorf("Bearer value error, current value %#x, expected value in the range 00 to 1F inclusive",
			bearer)
		return nil, err
	}
	// The DIRECTION bit shall be 0 for uplink and 1 for downlink
	if direction > AuthDirectionMax {
		err = fmt.Errorf("Direction value error, current value %#x, expected value in the range 0 to 1 inclusive",
			direction)
		return nil, err
	}
	// plaintext and bLength consistency check
	if len(ciphertext)*8 < bLength {
		err = fmt.Errorf("Plaintext length error, current value %d, expected value %d",
			bLength, len(ciphertext)*8)
		return nil, err
	}
	// T1 构造
	//CounterT1 := Count + Bearer + Direction + padding(90个0)
	CounterT1 := make([]byte, 16)
	CounterT1 = append(CounterT1[:0], count...)
	tmpByte := bearer<<3 | (direction&0x1)<<2
	CounterT1 = append(CounterT1, tmpByte)
	tmp8Byte := make([]byte, 11)
	CounterT1 = append(CounterT1, tmp8Byte...)

	//Key128 := key
	iv := CounterT1

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Failure to New Cipher")
	}

	ctr := cipher.NewCTR(c, iv)
	decrypted := make([]byte, len(ciphertext))

	ctr.XORKeyStream(decrypted, ciphertext)
	// 清除多余bit
	decrypted[len(decrypted)-1] = (decrypted[len(decrypted)-1] >>
		uint(len(decrypted)*8-bLength)) <<
		uint(len(decrypted)*8-bLength)
	return decrypted, nil

}

// 定义cipher
type nea2 struct {
	key       []byte
	count     []byte
	bearer    byte
	direction byte
	c         cipher.Block
	iv        []byte
	//keyStream cipher.Stream // 一个stream只能使用1次
}

func (n *nea2) Encrypt(plaintext []byte, bLength int) (
	ciphertext []byte, err error) {
	//panic("implement me")
	if bLength > len(plaintext)*8 {
		return nil, errors.New("invalid bLength")
	}

	keyStream := cipher.NewCTR(n.c, n.iv)
	//n.keyStream = keyStream

	encrypted := make([]byte, len(plaintext))

	keyStream.XORKeyStream(encrypted, plaintext)
	// 清除多余bit
	encrypted[len(encrypted)-1] = (encrypted[len(encrypted)-1] >>
		uint(len(encrypted)*8-bLength)) <<
		uint(len(encrypted)*8-bLength)

	return encrypted, nil
}

func (n *nea2) Decrypt(ciphertext []byte, bLength int) (
	plaintext []byte, err error) {
	//panic("implement me")
	if bLength > len(ciphertext)*8 {
		return nil, errors.New("invalid bLength")
	}
	decrypted := make([]byte, len(ciphertext))

	keyStream := cipher.NewCTR(n.c, n.iv)

	keyStream.XORKeyStream(decrypted, ciphertext)
	// 清除多余bit
	decrypted[len(decrypted)-1] = (decrypted[len(decrypted)-1] >>
		uint(len(decrypted)*8-bLength)) <<
		uint(len(decrypted)*8-bLength)

	return decrypted, nil
}

// 新建cipher
func NewCipherBlock(key []byte) (cipher.Block, error) {
	if len(key) != AuthKeyLengthBytes {
		err := fmt.Errorf("Key length error, current value %d, expected value %d",
			len(key), AuthKeyLengthBytes)
		return nil, err
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Failure to New Cipher")
	}
	return c, nil
}

func NewNea2(key []byte, count []byte, bearer byte, direction byte) (*nea2, error) {
	nea := nea2{
		key:       key,
		count:     count,
		bearer:    bearer,
		direction: direction,
	}
	if len(key) != AuthKeyLengthBytes {
		err := fmt.Errorf("Key length error, current value %d, expected value %d",
			len(key), AuthKeyLengthBytes)
		return nil, err
	}
	if len(count) != AuthCountLengthBytes {
		err := fmt.Errorf("Count length error, current value %d, expected value %d",
			len(count), AuthCountLengthBytes)
		return nil, err
	}
	//  two-hex-digit value in the range 00 to 1F inclusive
	if bearer > AuthBearerMax {
		err := fmt.Errorf("Bearer value error, current value %0x, expected value in the range 00 to 1F inclusive",
			bearer)
		return nil, err
	}
	// The DIRECTION bit shall be 0 for uplink and 1 for downlink
	if direction > AuthDirectionMax {
		err := fmt.Errorf("Direction value error, current value %0x, expected value in the range 0 to 1 inclusive",
			direction)
		return nil, err
	}

	// T1 构造
	//CounterT1 := Count + Bearer + Direction + padding(90个0)
	CounterT1 := make([]byte, 16)
	CounterT1 = append(CounterT1[:0], count...)
	tmpByte := bearer<<3 | (direction&0x1)<<2
	CounterT1 = append(CounterT1, tmpByte)
	tmp8Byte := make([]byte, 11)
	CounterT1 = append(CounterT1, tmp8Byte...)

	fmt.Printf("CounterT1:%x\n", CounterT1)

	//Key128 := key
	iv := CounterT1
	nea.iv = iv
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("Failure to New Cipher")
	}
	nea.c = c
	/*keyStream := cipher.NewCTR(c, iv)
	nea.keyStream = keyStream*/
	return &nea, nil
}
