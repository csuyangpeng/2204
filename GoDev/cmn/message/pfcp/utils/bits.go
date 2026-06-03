package utils

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/willf/bitset"
)

// Bit constants
const (
	Bit_0 uint8 = 1
	Bit_1 uint8 = 1 << 1
	Bit_2 uint8 = 1 << 2
	Bit_3 uint8 = 1 << 3
	Bit_4 uint8 = 1 << 4
	Bit_5 uint8 = 1 << 5
	Bit_6 uint8 = 1 << 6
	Bit_7 uint8 = 1 << 7
)

var BIT = [8]uint8{Bit_0, Bit_1, Bit_2, Bit_3, Bit_4, Bit_5, Bit_6, Bit_7}

// Set bit n to 1
func SetByte(b *uint8, n uint) {
	if n < 8 {
		//*b = *b | (1 << n)
		*b = *b | BIT[n]
	}
}

// set bit n to 0
func ReSetByte(b *uint8, n uint) {
	if n < 8 {
		*b = *b &^ (BIT[n])
	}
}

// If this bit is set to "1"
func IsSetByte(b uint8, n uint) bool {
	if n < 8 {
		b = b & (BIT[n])
	}
	return b == BIT[n]
}

// bit set
// Set bit n to 1
func SET_BIT(b *bitset.BitSet, n uint) {
	//btSet := bitset.BitSet{}
	b.Set(n)
}

// getNumberbit Gets the bit value in the byte
func GetBitInByte(srcByte uint8, n uint8) (uint8, error) {
	if (n >= 1) && (n <= 8) {
		return ((srcByte << (8 - n)) >> 7), nil
	}
	return 0, errors.New("The n is between 1 and 8")
}

func GetBoolInByte(srcByte uint8, n uint8) (bool, error) {
	if (n >= 1) && (n <= 8) {
		return ((srcByte << (8 - n)) >> 7) == 1, nil
	}
	return false, errors.New("The n is between 1 and 8")
}

// getBitInUint32 Gets the bit value in the uint32
// 小端字节序
func GetBitInUint32(srcByte uint32, n uint8) (uint8, error) {
	if (n >= 1) && (n <= 32) {
		return uint8((srcByte << (32 - n)) >> 31), nil
	}
	return 0, errors.New("The n is between 0 and 7")
}

func SkipIe(msgBuf *bytes.Reader) {
	length, _ := msgBuf.ReadByte()
	leftBytes := make([]byte, length)
	binary.Read(msgBuf, binary.BigEndian, leftBytes)
}

func SkipIeOneByte(msgBuf *bytes.Reader) {
	msgBuf.ReadByte()
}

func SkipBigIe(msgBuf *bytes.Reader) {

	lengthBytes := make([]byte, 2)
	binary.Read(msgBuf, binary.BigEndian, lengthBytes)
	length := binary.BigEndian.Uint16(lengthBytes)

	leftBytes := make([]byte, length)
	binary.Read(msgBuf, binary.BigEndian, leftBytes)
}

// byte decode
type ByteOne byte

//Right shift ,n from 1 to 8,n 当前位额索引
func (b *ByteOne) RightShift(n uint8) (byte, error) {
	if (n >= 1) && (n <= 8) {
		return byte(*b) >> (n - 1), nil
	}
	return 0, fmt.Errorf("invalid input parameter, index(%d)", n)

}

//Left shift
func (b *ByteOne) LeftShift(n uint8) (byte, error) {
	if (n >= 1) && (n <= 8) {
		return byte(*b) >> (n - 1), nil
	}
	return 0, fmt.Errorf("invalid input parameter, index(%d)", n)

}

func (b *ByteOne) GetBit(index uint8) (uint8, error) {
	if index > 8 || index < 1 {
		return 0, fmt.Errorf("invalid input parameter, index(%d)", index)
	}

	var val byte = 1
	val = val << (index - 1)
	if (byte(*b))&val != 0 {
		return 1, nil
	} else {
		return 0, nil
	}
}

// fromIdx: 1 to 8
func (b *ByteOne) GetBits(fromIdx uint8, toIdx uint8) (byte, error) {
	if fromIdx > 8 || fromIdx < 1 ||
		toIdx > 8 || toIdx < 1 ||
		fromIdx > toIdx {
		return 0, fmt.Errorf("invalid input parameter, fromIdx(%d), toIdx(%d)", fromIdx, toIdx)
	}

	var val byte = 0
	for i := fromIdx; i <= toIdx; i++ {
		var tmp byte = 1
		val |= tmp << (i - 1)
	}

	return byte(*b) & val, nil
}

// bool to uint8
func BoolToUint8(v bool) uint8 {
	if v {
		return 1
	} else {
		return 0
	}
}
