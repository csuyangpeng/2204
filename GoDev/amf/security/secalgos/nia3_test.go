package secalgos

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"testing"
	"unsafe"
)

func TestSet1(t *testing.T) {
	Key, _ := hex.DecodeString("00000000000000000000000000000000")
	Count := uint32(0x0)
	Bearer := uint32(0x0)
	Direction := uint32(0x0)
	Length := uint32(1) //bits,Plaintext length

	Plaintext := []uint32{0x00000000} // 32bit 补齐
	MacT := []uint32{0xc8a9595e}
	Mac := make([]byte, 4*uint(math.Ceil(float64(Length)/32))) //28
	fmt.Println("Ciphertext len ", len(Plaintext)*4)
	Eia3(Key, Count, Bearer, Direction, Length, Plaintext, Mac)
	fmt.Printf("-----MacT %08x\n", MacT)
	fmt.Printf("-----src Mac  %08x\n", Mac) //小端显示的
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	//var tt []byte = make([]byte, 4)
	//LittleEndian 指[]byte的字节序
	tt := binary.LittleEndian.Uint32(Mac)
	fmt.Printf("-----Mac  %08x\n", tt)
	if MacT[0] != tt {
		panic("MAC错误!")
	}

}

// 大端输入，大端输出
func TestSet1V1(t *testing.T) {
	Key, _ := hex.DecodeString("00000000000000000000000000000000")
	Count := uint32(0x0)
	Bearer := uint32(0x0)
	Direction := uint32(0x0)
	Length := uint32(1) //bits,Plaintext length

	PlaintextByte, _ := hex.DecodeString("00000001")
	fmt.Println("PlaintextByte:")
	fmt.Println(PlaintextByte)
	fmt.Printf("%08b\n", PlaintextByte)

	PlaintextInt := []uint32{0x00000001, 0x1} // 32bit 补齐
	fmt.Println("PlaintextInt:")
	fmt.Println(PlaintextInt)
	fmt.Printf("%032b\n", PlaintextInt)

	fmt.Println("PlaintextByte ZeroPadding:")
	PlaintextByte = ZeroPadding(PlaintextByte, 4)
	txtUint32, _ := BytesToUint32ArrayV1(PlaintextByte)
	fmt.Printf("%032b\n", txtUint32)

	Plaintext := []uint32{0x00000000} // 32bit 补齐，msg是大端序，最高位有效

	MacT := []uint32{0xc8a9595e}
	Mac := uint32(0)
	fmt.Println("Ciphertext len ", len(Plaintext)*4)
	Eia3V1(Key, Count, Bearer, Direction, Length, Plaintext, &Mac)
	fmt.Printf("-----MacT %08x\n", MacT)
	fmt.Printf("-----Mac  %08x\n", Mac)
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	//var tt []byte = make([]byte, 4)
	//binary.LittleEndian.PutUint32(tt, MacT[0])
	if Mac != MacT[0] {
		panic("MAC错误!")
	}

}

func TestSet2(t *testing.T) {
	Key, _ := hex.DecodeString("47054125561eb2dda94059da05097850")
	Count := uint32(0x561eb2dd)
	Bearer := uint32(0x14)
	Direction := uint32(0x0)
	Length := uint32(90) //bits,Plaintext length

	Plaintext := []uint32{0x00000000, 0x00000000, 0x00000000}
	MacT := []uint32{0x6719a088}
	Mac := make([]byte, unsafe.Sizeof(uint32(1)))
	fmt.Println("Ciphertext len ", len(Plaintext)*4)
	Eia3(Key, Count, Bearer, Direction, Length, Plaintext, Mac)
	fmt.Printf("-----MacT %08x\n", MacT)
	fmt.Printf("-----Mac  %08x\n", BytesToUint32Array(Mac))
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	var tt = make([]byte, 4)
	binary.LittleEndian.PutUint32(tt, MacT[0])
	if !bytes.Equal(Mac, tt) {
		panic("MAC错误!")
	}

}

func TestSet3(t *testing.T) {
	Key, _ := hex.DecodeString("c9e6cec4607c72db000aefa88385ab0a")
	Count := uint32(0xa94059da)
	Bearer := uint32(0xa)
	Direction := uint32(0x1)
	Length := uint32(577) //bits,Plaintext length

	Plaintext := []uint32{0x983b41d4, 0x7d780c9e, 0x1ad11d7e, 0xb70391b1,
		0xde0b35da, 0x2dc62f83, 0xe7b78d63, 0x06ca0ea0,
		0x7e941b7b, 0xe91348f9, 0xfcb170e2, 0x217fecd9,
		0x7f9f68ad, 0xb16e5d7d, 0x21e569d2, 0x80ed775c,
		0xebde3f40, 0x93c53881, 0x00000000}
	MacT := []uint32{0xfae8ff0b}
	Mac := make([]byte, unsafe.Sizeof(uint32(1)))
	fmt.Println("Ciphertext len ", len(Plaintext)*4)
	Eia3(Key, Count, Bearer, Direction, Length, Plaintext, Mac)
	fmt.Printf("-----MacT %x\n", MacT)
	fmt.Printf("-----Mac  %x\n", BytesToUint32Array(Mac))
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	var tt = make([]byte, 4)
	binary.LittleEndian.PutUint32(tt, MacT[0])
	if !bytes.Equal(Mac, tt) {
		panic("MAC错误!")
	}

}

func TestSet4(t *testing.T) {
	Key, _ := hex.DecodeString("c8a48262d0c2e2bac4b96ef77e80ca59")
	Count := uint32(0x05097850)
	Bearer := uint32(0x10)
	Direction := uint32(0x1)
	Length := uint32(2079) //bits,Plaintext length

	Plaintext := []uint32{0xb546430b, 0xf87b4f1e, 0xe834704c, 0xd6951c36, 0xe26f108c, 0xf731788f, 0x48dc34f1, 0x678c0522,
		0x1c8fa7ff, 0x2f39f477, 0xe7e49ef6, 0x0a4ec2c3, 0xde24312a, 0x96aa26e1, 0xcfba5756, 0x3838b297,
		0xf47e8510, 0xc779fd66, 0x54b14338, 0x6fa639d3, 0x1edbd6c0, 0x6e47d159, 0xd94362f2, 0x6aeeedee,
		0x0e4f49d9, 0xbf841299, 0x5415bfad, 0x56ee82d1, 0xca7463ab, 0xf085b082, 0xb09904d6, 0xd990d43c,
		0xf2e062f4, 0x0839d932, 0x48b1eb92, 0xcdfed530, 0x0bc14828, 0x0430b6d0, 0xcaa094b6, 0xec8911ab,
		0x7dc36824, 0xb824dc0a, 0xf6682b09, 0x35fde7b4, 0x92a14dc2, 0xf4364803, 0x8da2cf79, 0x170d2d50,
		0x133fd494, 0x16cb6e33, 0xbea90b8b, 0xf4559b03, 0x732a01ea, 0x290e6d07, 0x4f79bb83, 0xc10e5800,
		0x15cc1a85, 0xb36b5501, 0x046e9c4b, 0xdcae5135, 0x690b8666, 0xbd54b7a7, 0x03ea7b6f, 0x220a5469,
		0xa568027e}
	MacT := []uint32{0x004ac4d6}
	Mac := make([]byte, unsafe.Sizeof(uint32(1)))
	fmt.Println("Ciphertext len ", len(Plaintext)*4)
	Eia3(Key, Count, Bearer, Direction, Length, Plaintext, Mac)
	fmt.Printf("-----MacT %08x\n", MacT)
	fmt.Printf("-----Mac  %08x\n", BytesToUint32Array(Mac))
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	var tt = make([]byte, 4)
	binary.LittleEndian.PutUint32(tt, MacT[0])
	if !bytes.Equal(Mac, tt) {
		panic("MAC错误!")
	}

}

func TestSet5(t *testing.T) {
	Key, _ := hex.DecodeString("6b8b08ee79e0b5982d6d128ea9f220cb")
	Count := uint32(0x561eb2dd)
	Bearer := uint32(0x1c)
	Direction := uint32(0x0)
	Length := uint32(5670) //bits,Plaintext length

	Plaintext := []uint32{0x5bad7247, 0x10ba1c56, 0xd5a315f8, 0xd40f6e09, 0x3780be8e, 0x8de07b69, 0x92432018, 0xe08ed96a,
		0x5734af8b, 0xad8a575d, 0x3a1f162f, 0x85045cc7, 0x70925571, 0xd9f5b94e, 0x454a77c1, 0x6e72936b,
		0xf016ae15, 0x7499f054, 0x3b5d52ca, 0xa6dbeab6, 0x97d2bb73, 0xe41b8075, 0xdce79b4b, 0x86044f66,
		0x1d4485a5, 0x43dd7860, 0x6e0419e8, 0x059859d3, 0xcb2b67ce, 0x0977603f, 0x81ff839e, 0x33185954,
		0x4cfbc8d0, 0x0fef1a4c, 0x8510fb54, 0x7d6b06c6, 0x11ef44f1, 0xbce107cf, 0xa45a06aa, 0xb360152b,
		0x28dc1ebe, 0x6f7fe09b, 0x0516f9a5, 0xb02a1bd8, 0x4bb0181e, 0x2e89e19b, 0xd8125930, 0xd178682f,
		0x3862dc51, 0xb636f04e, 0x720c47c3, 0xce51ad70, 0xd94b9b22, 0x55fbae90, 0x6549f499, 0xf8c6d399,
		0x47ed5e5d, 0xf8e2def1, 0x13253e7b, 0x08d0a76b, 0x6bfc68c8, 0x12f375c7, 0x9b8fe5fd, 0x85976aa6,
		0xd46b4a23, 0x39d8ae51, 0x47f680fb, 0xe70f978b, 0x38effd7b, 0x2f7866a2, 0x2554e193, 0xa94e98a6,
		0x8b74bd25, 0xbb2b3f5f, 0xb0a5fd59, 0x887f9ab6, 0x8159b717, 0x8d5b7b67, 0x7cb546bf, 0x41eadca2,
		0x16fc1085, 0x0128f8bd, 0xef5c8d89, 0xf96afa4f, 0xa8b54885, 0x565ed838, 0xa950fee5, 0xf1c3b0a4,
		0xf6fb71e5, 0x4dfd169e, 0x82cecc72, 0x66c850e6, 0x7c5ef0ba, 0x960f5214, 0x060e71eb, 0x172a75fc,
		0x1486835c, 0xbea65344, 0x65b055c9, 0x6a72e410, 0x52241823, 0x25d83041, 0x4b40214d, 0xaa8091d2,
		0xe0fb010a, 0xe15c6de9, 0x0850973b, 0xdf1e423b, 0xe148a237, 0xb87a0c9f, 0x34d4b476, 0x05b803d7,
		0x43a86a90, 0x399a4af3, 0x96d3a120, 0x0a62f3d9, 0x507962e8, 0xe5bee6d3, 0xda2bb3f7, 0x237664ac,
		0x7a292823, 0x900bc635, 0x03b29e80, 0xd63f6067, 0xbf8e1716, 0xac25beba, 0x350deb62, 0xa99fe031,
		0x85eb4f69, 0x937ecd38, 0x7941fda5, 0x44ba67db, 0x09117749, 0x38b01827, 0xbcc69c92, 0xb3f772a9,
		0xd2859ef0, 0x03398b1f, 0x6bbad7b5, 0x74f7989a, 0x1d10b2df, 0x798e0dbf, 0x30d65874, 0x64d24878,
		0xcd00c0ea, 0xee8a1a0c, 0xc753a279, 0x79e11b41, 0xdb1de3d5, 0x038afaf4, 0x9f5c682c, 0x3748d8a3,
		0xa9ec54e6, 0xa371275f, 0x1683510f, 0x8e4f9093, 0x8f9ab6e1, 0x34c2cfdf, 0x4841cba8, 0x8e0cff2b,
		0x0bcc8e6a, 0xdcb71109, 0xb5198fec, 0xf1bb7e5c, 0x531aca50, 0xa56a8a3b, 0x6de59862, 0xd41fa113,
		0xd9cd9578, 0x08f08571, 0xd9a4bb79, 0x2af271f6, 0xcc6dbb8d, 0xc7ec36e3, 0x6be1ed30, 0x8164c31c,
		0x7c0afc54, 0x1c000000}
	MacT := []uint32{0x0ca12792}
	Mac := make([]byte, unsafe.Sizeof(uint32(1)))
	fmt.Println("Ciphertext len ", len(Plaintext)*4)
	Eia3(Key, Count, Bearer, Direction, Length, Plaintext, Mac)
	fmt.Printf("-----MacT %08x\n", MacT)
	fmt.Printf("-----Mac  %08x\n", BytesToUint32Array(Mac))
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	var tt = make([]byte, 4)
	binary.LittleEndian.PutUint32(tt, MacT[0])
	if !bytes.Equal(Mac, tt) {
		panic("MAC错误!")
	}

}

func TestSet5LittleEndian(t *testing.T) {
	Key, _ := hex.DecodeString("6b8b08ee79e0b5982d6d128ea9f220cb")
	Count := uint32(0x561eb2dd)
	Bearer := uint32(0x1c)
	Direction := uint32(0x0)
	Length := uint32(5670) //bits,Plaintext length

	fmt.Println("dec:", Count)
	bcount := make([]byte, 4)
	binary.LittleEndian.PutUint32(bcount, Count)
	fmt.Println("Hex:", bcount)
	//strconv.ParseUint()
	strcount := strconv.FormatUint(uint64(Count), 2)
	fmt.Println("Bin:", strcount)

	Plaintext := []uint32{0x5bad7247, 0x10ba1c56, 0xd5a315f8, 0xd40f6e09, 0x3780be8e, 0x8de07b69, 0x92432018, 0xe08ed96a,
		0x5734af8b, 0xad8a575d, 0x3a1f162f, 0x85045cc7, 0x70925571, 0xd9f5b94e, 0x454a77c1, 0x6e72936b,
		0xf016ae15, 0x7499f054, 0x3b5d52ca, 0xa6dbeab6, 0x97d2bb73, 0xe41b8075, 0xdce79b4b, 0x86044f66,
		0x1d4485a5, 0x43dd7860, 0x6e0419e8, 0x059859d3, 0xcb2b67ce, 0x0977603f, 0x81ff839e, 0x33185954,
		0x4cfbc8d0, 0x0fef1a4c, 0x8510fb54, 0x7d6b06c6, 0x11ef44f1, 0xbce107cf, 0xa45a06aa, 0xb360152b,
		0x28dc1ebe, 0x6f7fe09b, 0x0516f9a5, 0xb02a1bd8, 0x4bb0181e, 0x2e89e19b, 0xd8125930, 0xd178682f,
		0x3862dc51, 0xb636f04e, 0x720c47c3, 0xce51ad70, 0xd94b9b22, 0x55fbae90, 0x6549f499, 0xf8c6d399,
		0x47ed5e5d, 0xf8e2def1, 0x13253e7b, 0x08d0a76b, 0x6bfc68c8, 0x12f375c7, 0x9b8fe5fd, 0x85976aa6,
		0xd46b4a23, 0x39d8ae51, 0x47f680fb, 0xe70f978b, 0x38effd7b, 0x2f7866a2, 0x2554e193, 0xa94e98a6,
		0x8b74bd25, 0xbb2b3f5f, 0xb0a5fd59, 0x887f9ab6, 0x8159b717, 0x8d5b7b67, 0x7cb546bf, 0x41eadca2,
		0x16fc1085, 0x0128f8bd, 0xef5c8d89, 0xf96afa4f, 0xa8b54885, 0x565ed838, 0xa950fee5, 0xf1c3b0a4,
		0xf6fb71e5, 0x4dfd169e, 0x82cecc72, 0x66c850e6, 0x7c5ef0ba, 0x960f5214, 0x060e71eb, 0x172a75fc,
		0x1486835c, 0xbea65344, 0x65b055c9, 0x6a72e410, 0x52241823, 0x25d83041, 0x4b40214d, 0xaa8091d2,
		0xe0fb010a, 0xe15c6de9, 0x0850973b, 0xdf1e423b, 0xe148a237, 0xb87a0c9f, 0x34d4b476, 0x05b803d7,
		0x43a86a90, 0x399a4af3, 0x96d3a120, 0x0a62f3d9, 0x507962e8, 0xe5bee6d3, 0xda2bb3f7, 0x237664ac,
		0x7a292823, 0x900bc635, 0x03b29e80, 0xd63f6067, 0xbf8e1716, 0xac25beba, 0x350deb62, 0xa99fe031,
		0x85eb4f69, 0x937ecd38, 0x7941fda5, 0x44ba67db, 0x09117749, 0x38b01827, 0xbcc69c92, 0xb3f772a9,
		0xd2859ef0, 0x03398b1f, 0x6bbad7b5, 0x74f7989a, 0x1d10b2df, 0x798e0dbf, 0x30d65874, 0x64d24878,
		0xcd00c0ea, 0xee8a1a0c, 0xc753a279, 0x79e11b41, 0xdb1de3d5, 0x038afaf4, 0x9f5c682c, 0x3748d8a3,
		0xa9ec54e6, 0xa371275f, 0x1683510f, 0x8e4f9093, 0x8f9ab6e1, 0x34c2cfdf, 0x4841cba8, 0x8e0cff2b,
		0x0bcc8e6a, 0xdcb71109, 0xb5198fec, 0xf1bb7e5c, 0x531aca50, 0xa56a8a3b, 0x6de59862, 0xd41fa113,
		0xd9cd9578, 0x08f08571, 0xd9a4bb79, 0x2af271f6, 0xcc6dbb8d, 0xc7ec36e3, 0x6be1ed30, 0x8164c31c,
		0x7c0afc54, 0x1c000000}
	MacT := []uint32{0x0ca12792}
	Mac := make([]byte, unsafe.Sizeof(uint32(1)))
	fmt.Println("Ciphertext len ", len(Plaintext)*4)
	Eia3(Key, Count, Bearer, Direction, Length, Plaintext, Mac)
	fmt.Printf("-----MacT %08x\n", MacT)
	fmt.Printf("-----Mac  %08x\n", BytesToUint32Array(Mac))
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	var tt = make([]byte, 4)
	binary.LittleEndian.PutUint32(tt, MacT[0])
	if !bytes.Equal(Mac, tt) {
		panic("MAC错误!")
	}

}

func TestSet1V11(t *testing.T) {
	Key, _ := hex.DecodeString("00000000000000000000000000000000")
	Count := uint32(0x0)
	Bearer := uint32(0x0)
	Direction := uint32(0x0)
	Length := uint32(1) //bits,Plaintext length

	PlaintextByte, _ := hex.DecodeString("00000000")

	MacT := []uint32{0xc8a9595e}
	Mac := uint32(0)

	fmt.Println("PlaintextByte:")
	fmt.Println(PlaintextByte)
	fmt.Printf("%08b\n", PlaintextByte)

	fmt.Println("PlaintextByte ZeroPadding:")
	PlaintextByte = ZeroPadding(PlaintextByte, 4)
	txtUint32, _ := BytesToUint32ArrayV1(PlaintextByte)
	fmt.Printf("%032b\n", txtUint32)

	fmt.Println("txtUint32 len ", len(txtUint32)*4)
	Eia3V1(Key, Count, Bearer, Direction, Length, txtUint32, &Mac)
	fmt.Printf("-----MacT %08x\n", MacT) // 大端打印
	fmt.Printf("-----Mac  %08x\n", Mac)
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	//var tt []byte = make([]byte, 4)
	//binary.LittleEndian.PutUint32(tt, MacT[0])
	if Mac != MacT[0] {
		panic("MAC错误!")
	}

}

func TestSet2V1(t *testing.T) {
	Key, _ := hex.DecodeString("47054125561eb2dda94059da05097850")
	Count := uint32(0x561eb2dd)
	Bearer := uint32(0x14)
	Direction := uint32(0x0)
	Length := uint32(90) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("000000000000000000000000")
	MacT := []uint32{0x6719a088}
	//Mac := make([]byte, unsafe.Sizeof(uint32(1)))
	Mac := uint32(0)
	fmt.Println("Plaintext     len ", len(Plaintext))

	PlaintextByte := ZeroPadding(Plaintext, 4)
	txtUint32, _ := BytesToUint32ArrayV1(PlaintextByte)
	fmt.Println("PlaintextByte len ", len(PlaintextByte))
	fmt.Println("txtUint32     len ", len(txtUint32))

	Eia3V1(Key, Count, Bearer, Direction, Length, txtUint32, &Mac)
	fmt.Printf("-----MacT %08x\n", MacT[0])
	fmt.Printf("-----Mac  %08x\n", Mac)
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	if Mac != MacT[0] {
		panic("MAC错误!")
	}

}
func TestSet3V1(t *testing.T) {
	Key, _ := hex.DecodeString("c9e6cec4607c72db000aefa88385ab0a")
	Count := uint32(0xa94059da)
	Bearer := uint32(0xa)
	Direction := uint32(0x1)
	Length := uint32(577) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("983b41d47d780c9e1ad11d7eb70391b1" +
		"de0b35da2dc62f83e7b78d6306ca0ea0" +
		"7e941b7be91348f9fcb170e2217fecd9" +
		"7f9f68adb16e5d7d21e569d280ed775c" +
		"ebde3f4093c5388100000000")
	MacT := []uint32{0xfae8ff0b}
	Mac := uint32(0)
	fmt.Println("Plaintext     len ", len(Plaintext)*8)

	PlaintextByte := ZeroPadding(Plaintext, 4)
	txtUint32, _ := BytesToUint32ArrayV1(PlaintextByte)
	fmt.Println("PlaintextByte len ", len(PlaintextByte)*8)
	fmt.Println("txtUint32     len ", len(txtUint32)*32)

	Eia3V1(Key, Count, Bearer, Direction, Length, txtUint32, &Mac)
	fmt.Printf("-----MacT %08x\n", MacT[0])
	fmt.Printf("-----Mac  %08x\n", Mac)
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	if Mac != MacT[0] {
		panic("MAC错误!")
	}

}

func TestSet4V1(t *testing.T) {
	Key, _ := hex.DecodeString("c8a48262d0c2e2bac4b96ef77e80ca59")
	Count := uint32(0x05097850)
	Bearer := uint32(0x10)
	Direction := uint32(0x1)
	Length := uint32(2079) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("b546430bf87b4f1ee834704cd6951c36e26f108cf731788f48dc34f1678c05221c8fa7ff2f39f477" +
		"e7e49ef60a4ec2c3de24312a96aa26e1cfba57563838b297f47e8510c779fd6654b143386fa639d31edbd6c06e47d159d94362f26aeeedee0e4f49d9" +
		"bf8412995415bfad56ee82d1ca7463abf085b082b09904d6d990d43cf2e062f40839d93248b1eb92cdfed5300bc148280430b6d0caa094b6ec8911ab" +
		"7dc36824b824dc0af6682b0935fde7b492a14dc2f43648038da2cf79170d2d50133fd49416cb6e33bea90b8bf4559b03732a01ea290e6d074f79bb83" +
		"c10e580015cc1a85b36b5501046e9c4bdcae5135690b8666bd54b7a703ea7b6f220a5469a568027e")
	MacT := []uint32{0x004ac4d6}
	Mac := uint32(0)
	fmt.Println("Plaintext     len ", len(Plaintext)*8)

	PlaintextByte := ZeroPadding(Plaintext, 4)
	txtUint32, _ := BytesToUint32ArrayV1(PlaintextByte)
	fmt.Println("PlaintextByte len ", len(PlaintextByte)*8)
	fmt.Println("txtUint32     len ", len(txtUint32)*32)
	if len(txtUint32) == 0 {
		panic("txtUint32 入参为空")
	}
	Eia3V1(Key, Count, Bearer, Direction, Length, txtUint32, &Mac)
	fmt.Printf("-----MacT %08x\n", MacT[0])
	fmt.Printf("-----Mac  %08x\n", Mac)
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	if Mac != MacT[0] {
		panic("MAC错误!")
	}

}

func TestSet5V1(t *testing.T) {
	Key, _ := hex.DecodeString("6b8b08ee79e0b5982d6d128ea9f220cb")
	Count := uint32(0x561eb2dd)
	Bearer := uint32(0x1c)
	Direction := uint32(0x0)
	Length := uint32(5670) //bits,Plaintext length

	Plaintext, _ := hex.DecodeString("5bad724710ba1c56d5a315f8d40f6e093780be8e8de07b6992432018e08ed96a" +
		"5734af8bad8a575d3a1f162f85045cc770925571d9f5b94e454a77c16e72936b" +
		"f016ae157499f0543b5d52caa6dbeab697d2bb73e41b8075dce79b4b86044f66" +
		"1d4485a543dd78606e0419e8059859d3cb2b67ce0977603f81ff839e33185954" +
		"4cfbc8d00fef1a4c8510fb547d6b06c611ef44f1bce107cfa45a06aab360152b" +
		"28dc1ebe6f7fe09b0516f9a5b02a1bd84bb0181e2e89e19bd8125930d178682f" +
		"3862dc51b636f04e720c47c3ce51ad70d94b9b2255fbae906549f499f8c6d399" +
		"47ed5e5df8e2def113253e7b08d0a76b6bfc68c812f375c79b8fe5fd85976aa6" +
		"d46b4a2339d8ae5147f680fbe70f978b38effd7b2f7866a22554e193a94e98a6" +
		"8b74bd25bb2b3f5fb0a5fd59887f9ab68159b7178d5b7b677cb546bf41eadca2" +
		"16fc10850128f8bdef5c8d89f96afa4fa8b54885565ed838a950fee5f1c3b0a4" +
		"f6fb71e54dfd169e82cecc7266c850e67c5ef0ba960f5214060e71eb172a75fc" +
		"1486835cbea6534465b055c96a72e4105224182325d830414b40214daa8091d2" +
		"e0fb010ae15c6de90850973bdf1e423be148a237b87a0c9f34d4b47605b803d7" +
		"43a86a90399a4af396d3a1200a62f3d9507962e8e5bee6d3da2bb3f7237664ac" +
		"7a292823900bc63503b29e80d63f6067bf8e1716ac25beba350deb62a99fe031" +
		"85eb4f69937ecd387941fda544ba67db0911774938b01827bcc69c92b3f772a9" +
		"d2859ef003398b1f6bbad7b574f7989a1d10b2df798e0dbf30d6587464d24878" +
		"cd00c0eaee8a1a0cc753a27979e11b41db1de3d5038afaf49f5c682c3748d8a3" +
		"a9ec54e6a371275f1683510f8e4f90938f9ab6e134c2cfdf4841cba88e0cff2b" +
		"0bcc8e6adcb71109b5198fecf1bb7e5c531aca50a56a8a3b6de59862d41fa113" +
		"d9cd957808f08571d9a4bb792af271f6cc6dbb8dc7ec36e36be1ed308164c31c" +
		"7c0afc541c000000")

	MacT := []uint32{0x0ca12792}
	Mac := uint32(0)
	fmt.Println("Plaintext     len ", len(Plaintext)*8)

	PlaintextByte := ZeroPadding(Plaintext, 4)
	txtUint32, _ := BytesToUint32ArrayV1(PlaintextByte)
	fmt.Println("PlaintextByte len ", len(PlaintextByte)*8)
	fmt.Println("txtUint32     len ", len(txtUint32)*32)
	if len(txtUint32) == 0 {
		panic("txtUint32 入参为空")
	}
	Eia3V1(Key, Count, Bearer, Direction, Length, txtUint32, &Mac)
	fmt.Printf("-----MacT %08x\n", MacT[0])
	fmt.Printf("-----Mac  %08x\n", Mac)
	//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
	if Mac != MacT[0] {
		panic("MAC错误!")
	}
}

func BenchmarkEIA3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Key, _ := hex.DecodeString("47054125561eb2dda94059da05097850")
		Count := uint32(0x561eb2dd)
		Bearer := uint32(0x14)
		Direction := uint32(0x0)
		Length := uint32(90) //bits,Plaintext length

		Plaintext, _ := hex.DecodeString("000000000000000000000000")
		MacT := []uint32{0x6719a088}
		//Mac := make([]byte, unsafe.Sizeof(uint32(1)))
		Mac := uint32(0)

		PlaintextByte := ZeroPadding(Plaintext, 4)
		txtUint32, _ := BytesToUint32ArrayV1(PlaintextByte)

		Eia3V1(Key, Count, Bearer, Direction, Length, txtUint32, &Mac)
		//tt := C.GoBytes((unsafe.Pointer)(&testM[0]), C.int(unsafe.Sizeof(testM)))
		if Mac != MacT[0] {
			panic("MAC错误!")
		}
		//	BenchmarkEIA3-4   	 1000000	      1192 ns/op
	}
}
