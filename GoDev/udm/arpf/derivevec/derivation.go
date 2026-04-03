package derivevec

import (
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/udm/arpf/milenage"
)

func DeriveHeAv(ueSecData *types.UeSecContext) (*types.HeAvType, error) {
	rlogger.FuncEntry(types.ModuleUdm, ueSecData)

	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil, "UeSecData:%s", ueSecData)

	if ueSecData == nil {
		return nil, fmt.Errorf("invalid input para")
	}

	HeAv := &types.HeAvType{}

	//generate Rand
	//	randData, err := GenerateRandomBytes(int(types.RandSize))
	// for debug
	var err error
	randData := []byte{0x12, 0x87, 0x57, 0x60, 0xdc, 0x77, 0xaa, 0x4f, 0x37, 0xcb, 0x62, 0xca, 0xe0, 0xcc, 0xb9, 0xc2}
	if err != nil {
		return nil, fmt.Errorf("failed to generate rand")
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil,
		"generate rand(%x)", randData)

	// for test
	//randData = []byte{0xb4, 0x1f, 0xd2, 0xb4, 0x0f, 0x9a, 0xaf, 0x42, 0xd1, 0x54, 0xd1, 0xf9, 0x00, 0x3f, 0xfb, 0x1a}

	for i := 0; i < len(randData); i++ {
		HeAv.Rand[i] = randData[i]
	}

	var mdata *milenage.Milenage

	if ueSecData.IsOpc {
		mdata = milenage.New(
			ueSecData.Key[:],
			ueSecData.Opc[:],
			randData[:],
			ConvertUint64Sqn(ueSecData.Sqn),
			binary.BigEndian.Uint16(ueSecData.Amf[:]),
			true)
	} else {
		mdata = milenage.New(
			ueSecData.Key[:],
			ueSecData.Op[:],
			randData[:],
			ConvertUint64Sqn(ueSecData.Sqn),
			binary.BigEndian.Uint16(ueSecData.Amf[:]),
			false)
	}

	HeAv.MACA, err = mdata.F1()
	if err != nil {
		return nil, fmt.Errorf("failed to generate MAC.")
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil,
		"generate maca(%x)", HeAv.MACA)

	HeAv.XRes, HeAv.CK, HeAv.IK, HeAv.AK, err = mdata.F2345()
	if err != nil {
		return nil, fmt.Errorf("failed to generate XRes,CK,IK and AK.")
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil,
		"generate XRes(%x),CK(%x),IK(%x),AK(%x)",
		HeAv.XRes, HeAv.CK, HeAv.IK, HeAv.AK)

	for i := 0; i < 6; i++ {
		HeAv.SqnXorAk[i] = ueSecData.Sqn[i] ^ HeAv.AK[i]
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil,
		"generate sqn_xor_ak(%x)", HeAv.SqnXorAk)

	// AUTN = SQN XOR AK || AMF || MAC
	var autn []byte
	autn = append(autn, HeAv.SqnXorAk[:]...)
	autn = append(autn, ueSecData.Amf[:]...)
	autn = append(autn, HeAv.MACA[:]...)

	for i := 0; i < len(autn); i++ {
		HeAv.Autn[i] = autn[i]
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil, "generate autn(%x)", HeAv.Autn)

	// 33.501 A.4 XRES*
	fc := 0x6B
	snName := []byte(ueSecData.SnName)
	pn := [][]byte{snName, HeAv.Rand[:], HeAv.XRes[:]}
	ln := []uint16{uint16(len(ueSecData.SnName)), uint16(types.RandSize), uint16(types.XresSize)}

	// The input key KEY shall be equal to the concatenation CK || IK of CK and IK
	var key []byte
	key = append(key, HeAv.CK[:]...)
	key = append(key, HeAv.IK[:]...)
	err, xresS := KdfDerivation(key, byte(fc), pn, ln, 3)
	if err != nil {
		return nil, fmt.Errorf("failed to generate XRES*")
	}

	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil,
		"generaet original xres*(%x)", xresS)

	//The (X)RES* is identified with the 128 least significant bits of the output of the KDF.
	for i := 0; i < int(types.XressSize); i++ {
		HeAv.XResStar[i] = xresS[i+int(types.XressSize)]
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil,
		"genereate xres*(%x)", HeAv.XResStar)

	// 33.501 A.2 Kausf
	kausf := DeriveKausf(key, ueSecData.SnName, HeAv.SqnXorAk)
	for i := 0; i < len(kausf); i++ {
		HeAv.Kausf[i] = kausf[i]
	}
	rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil,
		"HeAv Infos:%s ", HeAv)
	return HeAv, nil
}

func DeriveKausf(key []byte, snName string, sqnxorak [types.SqnxorakSize]byte) []byte {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	// 33.501 A.2 Kausf
	fc := 0x6A
	pn := [][]byte{[]byte(snName), sqnxorak[:]}
	ln := []uint16{uint16(len(snName)), uint16(types.SqnxorakSize)}

	// The input key KEY shall be equal to the concatenation CK || IK of CK and IK
	err, kausf := KdfDerivation(key[:], byte(fc), pn, ln, 2)
	if err != nil {
		return nil
	}

	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil,
		"generate kausf(%x) from sn_name(%s),sqn_xor_ak(%x), key-CK||IK(%x)",
		kausf,
		snName,
		sqnxorak,
		key)

	return kausf
}

func ComputeAK2(ueSecData *types.UeSecContext, rand [types.RandSize]byte) ([types.AkSize]byte, error) {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	var ak [types.AkSize]byte
	if ueSecData == nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "invalid input parameter, UeSecData")
		return ak, fmt.Errorf("invalid input parameter, UeSecData")
	}

	var mdata *milenage.Milenage
	if ueSecData.IsOpc {
		mdata = milenage.New(
			ueSecData.Key[:],
			ueSecData.Opc[:],
			rand[:],
			ConvertUint64Sqn(ueSecData.Sqn),
			binary.BigEndian.Uint16(ueSecData.Amf[:]),
			true)
	} else {
		mdata = milenage.New(
			ueSecData.Key[:],
			ueSecData.Op[:],
			rand[:],
			ConvertUint64Sqn(ueSecData.Sqn),
			binary.BigEndian.Uint16(ueSecData.Amf[:]),
			false)
	}

	ak, err := mdata.F5Star()
	if err != nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil,
			"failed to compute AK, error(%s)", err)
	}

	return ak, nil
}

func DeriveAusfAv(heav *types.HeAvType, snName string, ausfAv *types.AusfAvContext) error {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	if heav == nil || ausfAv == nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "invalid input parameter")
		return fmt.Errorf("inavlid input parameter")
	}

	ausfAv.Rand = heav.Rand
	ausfAv.Autn = heav.Autn

	// 33.501 A.5 HXRES* from XRES*
	err, hxress := DeriveHResS(heav.XResStar[:], heav.Rand[:])
	ausfAv.HXresStar = hxress
	rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil,
		"generate hxres*(%x) from xres*(%x) and rand(%x)",
		hxress, heav.XResStar, heav.Rand)

	// 33.501 A.6 Kseaf
	fc := 0x6C
	snNameBytes := []byte(snName)
	pn := [][]byte{snNameBytes}
	ln := []uint16{uint16(len(snName))}

	err, kaesf := KdfDerivation(heav.Kausf[:], byte(fc), pn, ln, 1)
	if err != nil {
		return fmt.Errorf("failed to generate Kaesf")
	}
	for i := 0; i < len(kaesf); i++ {
		ausfAv.KSeaf[i] = kaesf[i]
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil,
		"generate kseaf(%x) from kausf(%x) and sn_name(%s)", ausfAv.KSeaf, heav.Kausf, snName)

	return nil
}

func DeriveHResS(xress []byte, rand []byte) (err error, hxress [types.XressSize]byte) {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	if xress == nil || rand == nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "invalid input parameter")
		return fmt.Errorf("inavlid input parameter"), hxress
	}

	// 33.501 A.5 HXRES* from XRES*
	var s []byte
	s = append(s, rand...)
	s = append(s, xress...)
	out := Sha256(s)
	for i := 0; i < int(types.XressSize); i++ {
		hxress[i] = out[i+int(types.XressSize)]
	}
	return nil, hxress
}

func DeriveKamf(kseaf []byte, supi []byte, abba []byte) ([]byte, error) {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	if kseaf == nil || abba == nil || supi == nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "invalid input parameter Kausf or ABBA")
		return nil, fmt.Errorf("invalid input Kausf or ABBA")
	} else {
		rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil, "DeriveKamf, kseaf(%x), ABBA(%x),Supi(%x)", kseaf, abba, supi)
	}

	// 33.501 A.7 Kamf derivation
	fc := 0x6D

	pn := [][]byte{supi, abba}
	ln := []uint16{uint16(len(supi)), uint16(len(abba))}

	err, kamf := KdfDerivation(kseaf, byte(fc), pn, ln, 2)
	if err != nil {
		return nil, fmt.Errorf("failed to generate Kamf")
	}

	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil, "Kamf (%x)", kamf)
	return kamf, nil
}

func DeriveKgnb(kamf []byte, uplinkNasCount uint32, accessType byte) ([]byte, error) {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	if kamf == nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "invalid input parameter Kamf")
		return nil, fmt.Errorf("invalid input Kamf")
	}

	// 33.501 A.9 Kgnb derivation
	fc := 0x6E

	nasCountBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(nasCountBytes, uplinkNasCount)
	accessTypeBytes := make([]byte, 1)
	accessTypeBytes[0] = accessType

	pn := [][]byte{nasCountBytes, accessTypeBytes}
	ln := []uint16{uint16(len(nasCountBytes)), uint16(len(accessTypeBytes))}

	err, kgnb := KdfDerivation(kamf, byte(fc), pn, ln, 2)
	if err != nil {
		return nil, fmt.Errorf("failed to generate Kgnb")
	}

	return kgnb, nil
}

func DeriveNasKeys(kamf []byte, algoType byte, algoId byte) ([]byte, error) {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	if kamf == nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "invalid input parameter Kamf")
		return nil, fmt.Errorf("invalid input Kamf")
	} else {
		rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil, "derive kna, kamf(%x),algo type(%x),algo id(%x)", kamf, algoType, algoId)
	}

	// 33.501 A.8 Algorithm key
	fc := 0x69

	algoTypes := make([]byte, 1)
	algoTypes[0] = algoType

	algoIds := make([]byte, 1)
	algoIds[0] = ConvertAlgoIdentify(algoId)

	pn := [][]byte{algoTypes, algoIds}
	ln := []uint16{uint16(len(algoTypes)), uint16(len(algoIds))}

	err, key := KdfDerivation(kamf, byte(fc), pn, ln, 2)
	if err != nil {
		return nil, fmt.Errorf("failed to generate algoKey")
	}

	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil, "Nas Key 256 (%x)", key)
	// get the 128 least significant bits from key
	outKey := make([]byte, 16)
	for i := 0; i < 16; i++ {
		outKey[i] = key[i+16]
	}

	return outKey, nil
}

func ConvertAlgoIdentify(algoId byte) byte {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	// 33.501 A.8 Algorithm key  5.11.1.1  5.11.1.2
	//"0000"         NEA0			Null ciphering algorithm;
	//"0001"         128-NEA1		128-bit SNOW 3G based algorithm;
	//"0010"         128-NEA2		128-bit AES based algorithm; and
	//"0011"         128-NEA3		128-bit ZUC based algorithm.

	//"0000"         NIA0			Null Integrity Protection algorithm;
	//"0001"         128-NIA1		128-bit SNOW 3G based algorithm;
	//"0010"         128-NIA2		128-bit AES based algorithm; and
	//"0011"         128-NIA3		128-bit ZUC based algorithm.  5.11.1.1

	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil, "algorithm identifier (%x)", algoId)
	return algoId
}
