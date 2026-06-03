package utils

import (
	"fmt"
	"github.com/willf/bitset"
	"testing"
)

// e.g. UT Template
func TestGetBitInByte(t *testing.T) {
	//11001111 207
	ExpectedResults := [8]byte{1, 1, 1, 1, 0, 0, 1, 1} //前底位，后高位

	bit1, _ := GetBitInByte(uint8(207), 1)
	bit2, _ := GetBitInByte(uint8(207), 2)
	bit3, _ := GetBitInByte(uint8(207), 3)
	bit4, _ := GetBitInByte(uint8(207), 4)
	bit5, _ := GetBitInByte(uint8(207), 5)
	bit6, _ := GetBitInByte(uint8(207), 6)
	bit7, _ := GetBitInByte(uint8(207), 7)
	bit8, _ := GetBitInByte(uint8(207), 8)

	ActualResult := [8]byte{bit1, bit2, bit3, bit4, bit5, bit6, bit7, bit8}
	if ActualResult != ExpectedResults {
		t.Errorf("fault.")
	}
	fmt.Println("ExpectedResults", ExpectedResults)
	fmt.Println("ActualResult   ", ActualResult)

}

func TestGetBitInUint32(t *testing.T) {
	//11111110 01110001 65137
	ExpectedResults := [8]byte{1, 0, 0, 0, 1, 1, 1, 0} //前底位，后高位
	bit1, _ := GetBitInUint32(uint32(65137), 1)        //1
	bit2, _ := GetBitInUint32(uint32(65137), 2)        //0
	bit3, _ := GetBitInUint32(uint32(65137), 8)        //0
	bit4, _ := GetBitInUint32(uint32(65137), 9)        //0
	bit5, _ := GetBitInUint32(uint32(65137), 12)       //1
	bit6, _ := GetBitInUint32(uint32(65137), 13)       //1
	bit7, _ := GetBitInUint32(uint32(65137), 15)       //1
	bit8, _ := GetBitInUint32(uint32(65137), 32)       //0

	ActualResult := [8]byte{bit1, bit2, bit3, bit4, bit5, bit6, bit7, bit8}
	if ActualResult != ExpectedResults {
		t.Errorf("fault.")
	}
	fmt.Println("ExpectedResults", ExpectedResults)
	fmt.Println("ActualResult   ", ActualResult)

}

func TestBitsConstant(t *testing.T) {
	fmt.Println(BIT[0], "~", BIT[7])
	fmt.Printf("BIT:%-8b,\nBIT:%-8d\n", BIT, BIT)

	b := bitset.New(8)
	SET_BIT(b, 5)
	fmt.Printf("BIT:%-8b,\nBIT:%-8d\n", b, b)

	b1 := &bitset.BitSet{}
	SET_BIT(b1, 6)
	fmt.Printf("BIT:%-8b,\nBIT:%-8d\n", b1, b1)
	//	out
	/*1 ~ 128
	BIT:[1        10       100      1000     10000    100000   1000000  10000000],
	BIT:[1        2        4        8        16       32       64       128     ]
	BIT:&{1000     [100000  ]},
	BIT:&{8        [32      ]}
	BIT:&{111      [1000000 ]},
	BIT:&{7        [64      ]}*/

	d := uint8(0)
	SetByte(&d, 7)
	SetByte(&d, 0)
	SetByte(&d, 56)
	fmt.Println("SetByte d:")
	fmt.Printf("BIT:%-8b,\nBIT:%-8d\n", d, d)
	//	out
	/*d:
	BIT:10000001,
	BIT:129 */

	// 按位清空
	ReSetByte(&d, 7)
	//ReSetByte(&d,0)
	fmt.Println("ReSetByte d:")
	fmt.Printf("BIT:%08b,\nBIT:%08d\n", d, d)

	isOk := IsSetByte(d, 0)
	fmt.Println("IsSetByte d:", isOk)
	fmt.Printf("BIT:%08b,\nBIT:%08d\n", d, d)
	/*ReSetByte d:
	BIT:00000001,
	BIT:00000001
	IsSetByte d: true
	BIT:00000001,
	BIT:00000001*/
}
