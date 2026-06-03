package derivevec

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego/logs"
	"lite5gc/cmn/types"
	"lite5gc/udm/arpf/milenage"
	"testing"
)

func TestDeriveHeAv(t *testing.T) {
	ueData := &types.UeSecContext{}
	ueData.SnName = "5G:mnc001.mcc460.3gppnetwork.org"
	ueData.Amf = [2]byte{0x80, 0x00}
	ueData.Sqn = [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	ueData.Key = [16]byte{0x46, 0x5b, 0x5c, 0xe8, 0xb1, 0x99, 0xb4, 0x9f, 0xaa, 0x5f, 0x0a, 0x2e, 0xe2, 0x38, 0xa6, 0xbc}
	ueData.Op = [16]byte{0xcd, 0xc2, 0x02, 0xd5, 0x12, 0x3e, 0x20, 0xf6, 0x2b, 0x6d, 0x67, 0x6a, 0xc7, 0x2c, 0xb3, 0x18}

	heav, err := DeriveHeAv(ueData)
	if err != nil {
		logs.Error("failed to derive He AV, %s", err)
	}

	fmt.Println(heav)
}

func TestDeriveHeAv_02(t *testing.T) {
	ueData := &types.UeSecContext{}
	//ueData.SnName = []byte{'5', 'G', ':', 0x64, 0xf0, 0x00}
	ueData.SnName = "5G:mnc000.mcc460.3gppnetwork.org"
	ueData.Amf = [2]byte{0x80, 0x00}
	ueData.Sqn = [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01}
	ueData.Key = [16]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef, 0x12, 0x34, 0x56, 0x78, 0x90, 0xab, 0xcd, 0xef}
	ueData.Op = [16]byte{0x63, 0xBF, 0xA5, 0x0E, 0xE6, 0x52, 0x33, 0x65, 0xFF, 0x14, 0xC1, 0xF4, 0x5F, 0x88, 0x73, 0x7D}

	heav, err := DeriveHeAv(ueData)
	if err != nil {
		logs.Error("failed to derive He AV, %s", err)
	}

	fmt.Println(heav)
}

func TestSqn(t *testing.T) {
	//sqn := [8]byte{0x00, 0x00, 0x12, 0x34, 0x56, 0x78, 0x90, 0xAB}
	sqn := [6]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0xAB}
	sqn8 := make([]byte, 8)
	for i := 0; i < 6; i++ {
		sqn8[i+2] = sqn[i]
	}

	sqnU64 := binary.BigEndian.Uint64(sqn8)

	s := make([]byte, 8)
	binary.BigEndian.PutUint64(s, sqnU64)
	msqn := make([]byte, 6)
	for i := 0; i < 6; i++ {
		msqn[i] = s[i+2]
	}

	fmt.Printf("Input SQN: %x, Output MSQN: %x", sqn, msqn)

}

func TestDeriveHeAv3(t *testing.T) {
	ueData := &types.UeSecContext{}
	ueData.SnName = "5G:mnc001.mcc460.3gppnetwork.org"
	ueData.Amf = [2]byte{0x80, 0x00}
	//ueData.Sqn = [6]byte{0x00, 0x00, 0x00, 0x00, 0xf6, 0x20}
	ueData.Sqn = [6]byte{0x00, 0x00, 0x00, 0x00, 0xfb, 0x21}
	ueData.Key = [16]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x35}
	ueData.Op = [16]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12}

	heav, err := DeriveHeAv(ueData)
	if err != nil {
		logs.Error("failed to derive He AV, %s", err)
	}

	fmt.Println(heav)

	randData := [16]byte{0xb4, 0x1f, 0xd2, 0xb4, 0x0f, 0x9a, 0xaf, 0x42, 0xd1, 0x54, 0xd1, 0xf9, 0x00, 0x3f, 0xfb, 0x1a}
	ak2, err := ComputeAK2(ueData, randData)
	fmt.Printf("ak2 %x\n", ak2)
	sqnms_xor_ak2 := [6]byte{0x0e, 0xb1, 0x63, 0xbd, 0x17, 0xdd}

	var SQNms [6]byte
	for i := 0; i < 6; i++ {
		SQNms[i] = sqnms_xor_ak2[i] ^ ak2[i]
	}
	fmt.Printf("sqnms %x\n", SQNms)

	sqn_xor_ak := [6]byte{0xd5, 0x94, 0x78, 0x92, 0x19, 0xbd}
	var SQN [6]byte
	for i := 0; i < 6; i++ {
		SQN[i] = sqn_xor_ak[i] ^ heav.AK[i]
	}
	fmt.Printf("sqnhe %x\n", SQN)

	//verify MAC_S
	//f1*K(SQNMS || RAND || AMF)
	// c653a4510b0e3497

	//randData := []byte{0xb4, 0x1f, 0xd2, 0xb4, 0x0f, 0x9a, 0xaf, 0x42, 0xd1, 0x54, 0xd1, 0xf9, 0x00, 0x3f, 0xfb, 0x1a}
	SQNmsSlice := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf6, 0x20}
	sqn1 := binary.BigEndian.Uint64(SQNmsSlice)
	amf1 := binary.BigEndian.Uint16(ueData.Amf[:])

	mdata := milenage.New(
		ueData.Key[:],
		ueData.Op[:],
		randData[:],
		sqn1,
		amf1, false)

	mac_s, err := mdata.F1Star()

	fmt.Printf("MAC_S %x Expect c653a4510b0e3497\n", mac_s)

}

func TestDeriveMac(t *testing.T) {
	ueData := &types.UeSecContext{}
	ueData.SnName = "5G:mnc001.mcc460.3gppnetwork.org"
	ueData.Amf = [2]byte{0x80, 0x00}
	//ueData.Sqn = [6]byte{0x00, 0x00, 0x00, 0x00, 0xf6, 0x20}
	ueData.Sqn = [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x21}
	ueData.Key = [16]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x35}
	ueData.Op = [16]byte{0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12, 0x34, 0x56, 0x78, 0x90, 0x12}

	heav, err := DeriveHeAv(ueData)
	if err != nil {
		logs.Error("failed to derive He AV, %s", err)
	}

	fmt.Println(heav)

	randData := [16]byte{0xb4, 0x1f, 0xd2, 0xb4, 0x0f, 0x9a, 0xaf, 0x42, 0xd1, 0x54, 0xd1, 0xf9, 0x00, 0x3f, 0xfb, 0x1a}
	ak2, err := ComputeAK2(ueData, randData)
	fmt.Printf("ak2 %x\n", ak2)
	sqnms_xor_ak2 := [6]byte{0x0e, 0xb1, 0x63, 0xbd, 0x17, 0xdd}

	var SQNms [6]byte
	for i := 0; i < 6; i++ {
		SQNms[i] = sqnms_xor_ak2[i] ^ ak2[i]
	}
	fmt.Printf("sqnms %x\n", SQNms)

	sqn_xor_ak := [6]byte{0xd5, 0x94, 0x78, 0x92, 0x19, 0xbd}
	var SQN [6]byte
	for i := 0; i < 6; i++ {
		SQN[i] = sqn_xor_ak[i] ^ heav.AK[i]
	}
	fmt.Printf("sqnhe %x\n", SQN)

	//verify MAC_S
	//f1*K(SQNMS || RAND || AMF)
	// c653a4510b0e3497

	//randData := []byte{0xb4, 0x1f, 0xd2, 0xb4, 0x0f, 0x9a, 0xaf, 0x42, 0xd1, 0x54, 0xd1, 0xf9, 0x00, 0x3f, 0xfb, 0x1a}
	SQNmsSlice := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf6, 0x20}
	sqn1 := binary.BigEndian.Uint64(SQNmsSlice)
	amf1 := binary.BigEndian.Uint16(ueData.Amf[:])

	mdata := milenage.New(
		ueData.Key[:],
		ueData.Op[:],
		randData[:],
		sqn1,
		amf1)

	mac_s, err := mdata.F1Star()

	fmt.Printf("MAC_S %x Expect c653a4510b0e3497\n", mac_s)

}

func TestDeriveKausf2(t *testing.T) {

	sqnxorak, _ := hex.DecodeString("a2a5a793dd80")
	fmt.Println(sqnxorak)
}
func TestDeriveKausf(t *testing.T) {

	sqnxorak, _ := hex.DecodeString("a2a5a793dd80")

	rand, _ := hex.DecodeString("0100000000000000f71f99024d4f52e1") // len 16 bytes
	key, _ := hex.DecodeString("12345678901234567890123456789035")
	op, _ := hex.DecodeString("12345678901234567890123456789012")
	amf, _ := hex.DecodeString("8000")
	sqn := [6]byte{0, 0, 0, 0, 0, 0}
	mdata := milenage.New(
		key[:],
		op[:],
		rand[:],
		ConvertUint64Sqn(sqn),
		binary.BigEndian.Uint16(amf[:]))

	MACA, err := mdata.F1()
	if err != nil {
		fmt.Printf("failed to generate MAC.error(%s)\n", err)
	}
	fmt.Printf("MACA(%x)\n", MACA)

	XRes, CK, IK, AK, err := mdata.F2345()
	if err != nil {
		fmt.Printf("failed to generate XRes,CK,IK and AK.\n")
	}
	fmt.Printf("XRes(%x),CK(%x),IK(%x),AK %x\n", XRes, CK, IK, AK)

	var rSqn [6]byte
	for i := 0; i < 6; i++ {
		rSqn[i] = sqnxorak[i] ^ AK[i]
	}

	fmt.Printf("get sqn (%x)", rSqn)

	snName := "5G:mnc001.mcc460.3gppnetwork.org"
	fc := 0x6B
	sn := []byte(snName)
	pn := [][]byte{sn, rand[:], XRes[:]}
	ln := []uint16{uint16(len(sn)), uint16(types.RandSize), uint16(types.XresSize)}

	// The input key KEY shall be equal to the concatenation CK || IK of CK and IK
	var key_ckik []byte
	key_ckik = append(key_ckik, CK[:]...)
	key_ckik = append(key_ckik, IK[:]...)
	err, xresS := KdfDerivation(key_ckik, byte(fc), pn, ln, 3)

	fmt.Printf("XRes* 256 (%x)\n", xresS)
	sqnAK := [6]byte{}
	for i := 0; i < 6; i++ {
		sqnAK[i] = sqnxorak[i]
	}
	fmt.Printf("derive Kausf, ckik(%x),snname(%x),sqn_xor_ak(%x)", key_ckik, snName, sqnAK)
	kausf := DeriveKausf(key_ckik, snName, sqnAK)
	fmt.Printf("Kausf(%x)\n", kausf)

	fc = 0x6C
	snNameBytes := []byte(snName)
	pn = [][]byte{snNameBytes}
	ln = []uint16{uint16(len(snName))}

	err, kseaf := KdfDerivation(kausf, byte(fc), pn, ln, 1)
	fmt.Printf("Kseaf(%x)\n", kseaf)

	var supiStr string = "460010000000005"
	supi := []byte(supiStr)
	//supi := []byte{4,6,0,0,1,0,0,0,0,0,0,0,0,0,5}
	abba := []byte{0, 0}
	kamf, _ := DeriveKamf(kseaf, supi, abba)

	knint, _ := DeriveNasKeys(kamf, 0x02, 0x02)
	fmt.Printf("Knint(%x)", knint)

	//kgnb
	kgnb, err := DeriveKgnb(kamf, 0, 0x01)
	fmt.Printf("Kgnb(%x)", kgnb)

}

func TestDeriveKgnb_keys(t *testing.T) {
	kgnb, _ := hex.DecodeString("bc091e7a2bfac95cc4b7b06c56b4893b4a14670d90033eab1d4852b291a3d431")
	knint, _ := DeriveNasKeys(kgnb, 0x04, 0x02)
	fmt.Printf("Knint(%x)", knint)
}

//func TestDeriveAusfAv(t *testing.T) {
//
////	knint, _ := hex.DecodeString("9ba9057f9b6b763f671eeb53e1c5206e")
////	nasMsg, _ := hex.DecodeString("007e005d020102f0f0e1")
////	_, mac := nassecurity.GenerateIntegrity(types3gpp.NIA2, knint, 0, 1, nasMsg)
////	fmt.Printf("NIA2 MAC: %x", mac)
////}
