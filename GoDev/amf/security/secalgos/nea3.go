package secalgos

import (
	"encoding/binary"
	"fmt"
	"math"
)

// 大端输入
func Nea3Encrypt(key []byte, count []byte, bearer byte, direction byte, plaintext []byte, bLength int) (
	ciphertext []byte, err error) {
	// Entry validity check
	if len(key) != AuthKeyLengthBytes {
		err = fmt.Errorf("key length error, current value %d, expected value %d",
			len(key), AuthKeyLengthBytes)
		return nil, err
	}
	if len(count) != AuthCountLengthBytes {
		err = fmt.Errorf("count length error, current value %d, expected value %d",
			len(count), AuthCountLengthBytes)
		return nil, err
	}
	//  two-hex-digit value in the range 00 to 1F inclusive
	if bearer > AuthBearerMax {
		err = fmt.Errorf("bearer value error, current value %#x, "+
			"expected value in the range 00 to 1F inclusive",
			bearer)
		return nil, err
	}
	// The DIRECTION bit shall be 0 for uplink and 1 for downlink
	if direction > AuthDirectionMax {
		err = fmt.Errorf("direction value error, current value %#x, "+
			"expected value in the range 0 to 1 inclusive",
			direction)
		return nil, err
	}
	// plaintext and bLength consistency check
	if len(plaintext)*8 < bLength {
		err = fmt.Errorf("Plaintext length error, current value %d, expected value %d",
			bLength, len(plaintext)*8)
		return nil, err
	}

	Key := key
	Count := binary.BigEndian.Uint32(count)
	Bearer := uint32(bearer)
	Direction := uint32(direction)

	Length := uint32(bLength)

	textPadding := ZeroPadding(plaintext, 4)
	text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐

	CiphertextTmp := make([]uint32, uint(math.Ceil(float64(Length)/32)))
	// Eea3V1(ck []byte, count uint32, bearer uint32, direction uint32,
	//	length uint32, m []uint32, ciper []uint32)
	Eea3V1(Key, Count, Bearer, Direction, Length, text32, CiphertextTmp)
	ciphertext, err = Uint32ArrayToBytesV1(CiphertextTmp)
	if err != nil {
		return nil, err
	}
	ciphertext = ciphertext[:Length/8]
	return ciphertext, nil
}

func Nea3Decrypt(key []byte, count []byte, bearer byte, direction byte, ciphertext []byte, bLength int) (
	plaintext []byte, err error) {
	// Entry validity check
	if len(key) != AuthKeyLengthBytes {
		err = fmt.Errorf("key length error, current value %d, expected value %d",
			len(key), AuthKeyLengthBytes)
		return nil, err
	}
	if len(count) != AuthCountLengthBytes {
		err = fmt.Errorf("count length error, current value %d, expected value %d",
			len(count), AuthCountLengthBytes)
		return nil, err
	}
	//  two-hex-digit value in the range 00 to 1F inclusive
	if bearer > AuthBearerMax {
		err = fmt.Errorf("bearer value error, current value %#x, "+
			"expected value in the range 00 to 1F inclusive",
			bearer)
		return nil, err
	}
	// The DIRECTION bit shall be 0 for uplink and 1 for downlink
	if direction > AuthDirectionMax {
		err = fmt.Errorf("direction value error, current value %#x, "+
			"expected value in the range 0 to 1 inclusive",
			direction)
		return nil, err
	}
	// plaintext and bLength consistency check
	if len(ciphertext)*8 < bLength {
		err = fmt.Errorf("plaintext length error, current value %d, expected value %d",
			bLength, len(ciphertext)*8)
		return nil, err
	}

	Key := key
	Count := binary.BigEndian.Uint32(count)
	Bearer := uint32(bearer)
	Direction := uint32(direction)

	Length := uint32(bLength)

	textPadding := ZeroPadding(ciphertext, 4)
	text32, _ := BytesToUint32ArrayV1(textPadding) // 32bit 补齐

	plaintextTmp := make([]uint32, uint(math.Ceil(float64(Length)/32)))
	// Eea3V1(ck []byte, count uint32, bearer uint32, direction uint32,
	//	length uint32, m []uint32, ciper []uint32)
	Eea3V1(Key, Count, Bearer, Direction, Length, text32, plaintextTmp)
	plaintext, err = Uint32ArrayToBytesV1(plaintextTmp)
	if err != nil {
		return nil, err
	}
	plaintext = plaintext[:Length/8]
	return plaintext, nil
}
