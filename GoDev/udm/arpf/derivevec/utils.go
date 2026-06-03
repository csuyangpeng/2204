package derivevec

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func KdfDerivation(key []byte, fc byte, pn [][]byte, ln []uint16, n uint8) (error, []byte) {
	var s []byte

	// S = FC || P0 || L0 || P1 || L1 || P2 || L2 || P3 || L3 ||... || Pn || Ln
	s = append(s, fc)
	pNum := len(pn)
	if pNum != len(ln) {
		return fmt.Errorf("invalid input pn and ln, number mismatch."), nil
	}

	for i := 0; i < pNum; i++ {
		pnVal := pn[i]
		lnVal := ln[i]
		if len(pnVal) != int(lnVal) {
			return fmt.Errorf("length mismatch for p%d:%d and l%d:%d", i, len(pnVal), i, lnVal), nil
		}

		s = append(s, pnVal[:]...)

		lnBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(lnBytes, lnVal)
		s = append(s, lnBytes[:]...)
	}

	//fmt.Println("construct S:", s)

	output := HmacSha256(s, key)

	return nil, output
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// HmacSha256 computes the HMAC-SHA256 of the message with the provided secret
func HmacSha256(message, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return h.Sum(nil)
}

func Sha256(message []byte) []byte {
	h := sha256.New()
	h.Write(message)
	return h.Sum(nil)
}

func ConvertUint64Sqn(sqn [6]byte) uint64 {

	sqn8 := make([]byte, 8)
	sqn8[0] = 0x00
	sqn8[1] = 0x00
	for i := 0; i < 6; i++ {
		sqn8[i+2] = sqn[i]
	}

	return binary.BigEndian.Uint64(sqn8)
}
