package secalgos

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"github.com/jacobsa/crypto/cmac"
	"github.com/jacobsa/crypto/common"
	"math"
)

// The size of an AES-CMAC checksum, in bytes.
const Size = aes.BlockSize

const blockSize = Size

var subkeyZero []byte
var subkeyRb []byte

func init() {
	subkeyZero = bytes.Repeat([]byte{0x00}, blockSize)
	subkeyRb = append(bytes.Repeat([]byte{0x00}, blockSize-1), 0x87)
}

// 3GPP TS 33.501 V15.3.1 (2018-12)
// NIST SP 800-38B

// D.3	Integrity algorithms
// D.3.1.3	128-NIA2

func Nia2CMAC(key []byte, count []byte, bearer byte, direction byte, Msg []byte, bLength int) (
	mac []byte, err error) {
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
		err = fmt.Errorf("Bearer value error, current value %#x, expected value in the range 00 to 1f inclusive",
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
	if len(Msg)*8 < bLength {
		err = fmt.Errorf("Msg length error, current value %d, expected value %d",
			bLength, len(Msg)*8)
		return nil, err
	}
	// Mlen = BLENGTH + 64
	// Mlen := 128 //bits
	// M 构造
	// M := Count + Bearer + Direction + padding(26个0) + Message
	M := make([]byte, int(8+math.Ceil(float64(bLength)/8)))

	M = append(M[:0], count...)
	tmpByte := bearer<<3 | (direction&1)<<2 //(direction << 2)
	M = append(M, tmpByte)
	tmp3Byte := make([]byte, 3)
	M = append(M, tmp3Byte...)
	M = append(M, Msg...)

	ciph, _ := aes.NewCipher(key)
	k1, k2 := generateSubkeys(ciph)
	//fmt.Println("k1 :", hex.EncodeToString(k1))
	//fmt.Println("k2 :", hex.EncodeToString(k2))
	k1, k2 = k2, k1
	//Subkey Generation Steps:
	//	1. Let L = CIPHK(0^b).
	//L_test, _ := hex.DecodeString("6e4261385adfc1fcb7c85f0c469fb20c")

	//b := make([]byte, len(Key))
	//L, _ := ciph(Key, b)
	//fmt.Printf("L:%x\n",L)
	//L1, _ := ciph(L, b)
	//fmt.Printf("L:%x\n",L1)

	//	2. If MSB1(L) = 0, then K1 = L << 1;
	//	Else K1 = (L << 1) ⊕ Rb; see Sec. 5.3 for the definition of Rb.
	//	3. If MSB1(K1) = 0, then K2 = K1 << 1;
	//	Else K2 = (K1 << 1) ⊕ Rb.
	//	4. Return K1, K2.
	//
	h, _ := cmac.New(key)
	h.Write(M)
	b := make([]byte, 4)
	//fmt.Println(len(b[:0]), cap(b[:0])) // 0 4
	mac = h.Sum(b[:0])
	mac = mac[:4]
	return mac, nil
}

// MAC Verification
func MacVerify(dstMac []byte, srcMac []byte) bool {
	return bytes.Equal(dstMac, srcMac)

}

// Given the supplied cipher, whose block size must be 16 bytes, return two
// subkeys that can be used in MAC generation. See section 5.3 of NIST SP
// 800-38B. Note that the other NIST-approved block size of 8 bytes is not
// supported by this function.
func generateSubkeys(ciph cipher.Block) (k1 []byte, k2 []byte) {
	if ciph.BlockSize() != blockSize {
		panic("generateSubkeys requires a cipher with a block size of 16 bytes.")
	}

	// Step 1
	l := make([]byte, blockSize)
	ciph.Encrypt(l, subkeyZero)

	// Step 2: Derive the first subkey.
	if common.Msb(l) == 0 {
		// TODO(jacobsa): Accept a destination buffer in ShiftLeft and then hoist
		// the allocation in the else branch below.
		k1 = common.ShiftLeft(l)
	} else {
		k1 = make([]byte, blockSize)
		common.Xor(k1, common.ShiftLeft(l), subkeyRb)
	}

	// Step 3: Derive the second subkey.
	if common.Msb(k1) == 0 {
		k2 = common.ShiftLeft(k1)
	} else {
		k2 = make([]byte, blockSize)
		common.Xor(k2, common.ShiftLeft(k1), subkeyRb)
	}

	return
}
