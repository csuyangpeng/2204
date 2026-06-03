package sectypes

import (
	"crypto/sha256"
	"fmt"
)

const (
	AKA_EAP uint8 = iota
	AKA_5G
)
const (
	K_SIZE         uint8  = 16 //in Byte. Should be 128 bits (16) or 256 bits (32)
	Rand_SIZE      uint8  = 16
	OP_SIZE        uint8  = 16
	OPC_SIZE       uint8  = 16
	AMF_SIZE       uint8  = 2
	Autn_SIZE      uint8  = 16
	Auts_SIZE      uint8  = 14
	SQN_SIZE       uint8  = 6
	MCCMNC_SIZE    uint8  = 3
	MSIN_SIZE      uint8  = 5
	ECCPubKey_SIZE uint8  = 65
	SnName_SIZE    uint16 = 1024
	MACA_SIZE      uint8  = 8
	SqnXorAk_SIZE  uint8  = 6
	AK_SIZE        uint8  = 6
	XRES_SIZE      uint8  = 8
	CK_SIZE        uint8  = 16
	IK_SIZE        uint8  = 16
	/* HMAC-SHA256 MAC Size */
	HMAC_SHA256_SIZE       = sha256.Size
	XResS_SIZE       uint8 = 16
	ResS_SIZE        uint8 = 16
	Kausf_SIZE       uint8 = 32
	Kseaf_SIZE       uint8 = 32
	Kamf_SIZE        uint8 = 32
	KnasEnc_SIZE     uint8 = 16
	KnasInt_SIZE     uint8 = 16
)

//33.501 A8
//N-NAS-enc-alg	0x01
//N-NAS-int-alg	0x02
//N-RRC-enc-alg	0x03
//N-RRC-int-alg	0x04
//N-UP-enc-alg	0x05
//N-UP-int-alg	0x06
const (
	N_NAS_ENC_ALG byte = 0x01
	N_NAS_INT_ALG byte = 0x02
)

// 33.504 6.4.3.1
// The DIRECTION bit shall be set to 0 for uplink and 1 for downlink.
const (
	NAS_UPLINK   byte = 0x00
	NAS_DOWNLINK byte = 0x01
)

type SupiType struct {
	MccMnc [MCCMNC_SIZE]byte
	Msin   [MSIN_SIZE]byte
}

type SuciType struct {
	Mccmnc    [MCCMNC_SIZE]byte
	Msin      [MSIN_SIZE]byte
	EccPubKey [ECCPubKey_SIZE]byte
}

type SnName [SnName_SIZE]byte

type HeAvType struct {
	MACA     [MACA_SIZE]byte
	SqnXorAk [SqnXorAk_SIZE]byte
	AK       [AK_SIZE]byte
	XRes     [XRES_SIZE]byte
	CK       [CK_SIZE]byte
	IK       [IK_SIZE]byte
	Rand     [Rand_SIZE]byte
	Autn     [Autn_SIZE]byte
	XResStar [XResS_SIZE]byte
	Kausf    [Kausf_SIZE]byte
}

func (p *HeAvType) String() string {
	var dump string
	dump = fmt.Sprintf("MAC_A  (%x)\n", p.MACA)
	dump += fmt.Sprintf("AK     (%x)\n", p.AK)
	dump += fmt.Sprintf("XRes   (%x)\n", p.XRes)
	dump += fmt.Sprintf("CK     (%x)\n", p.CK)
	dump += fmt.Sprintf("IK     (%x)\n", p.IK)
	dump += fmt.Sprintf("RAND   (%x)\n", p.Rand)
	dump += fmt.Sprintf("AUTN   (%x)\n", p.Autn)
	dump += fmt.Sprintf("XRes*  (%x)\n", p.XResStar)
	dump += fmt.Sprintf("Kausf* (%x)\n", p.Kausf)

	return dump
}

type UeSecContext struct {
	Imsi       string
	AuthMethod uint8
	Sqn        [SQN_SIZE]byte
	Key        [K_SIZE]byte
	Op         [OP_SIZE]byte
	IsOpc      bool
	Opc        [OPC_SIZE]byte
	Amf        [AMF_SIZE]byte
	SnName     string
}

func (p *UeSecContext) String() string {
	var dump string
	dump = fmt.Sprintf("Sqn    (%x)\n", p.Sqn)
	dump += fmt.Sprintf("Key   (%x)\n", p.Key)
	dump += fmt.Sprintf("Op    (%x)\n", p.Op)
	dump += fmt.Sprintf("IsOpc  (%v)\n", p.IsOpc)
	dump += fmt.Sprintf("Opc   (%x)\n", p.Opc)
	dump += fmt.Sprintf("Amf   (%d)\n", p.Amf)
	dump += fmt.Sprintf("SnName (%x)\n", p.SnName)

	return dump
}

type SeAvType struct {
	Rand      [Rand_SIZE]byte
	Autn      [Autn_SIZE]byte
	HXresStar [XResS_SIZE]byte
}

func (p *SeAvType) String() string {
	var dump string
	dump += fmt.Sprintf("RAND   (%x)\n", p.Rand)
	dump += fmt.Sprintf("AUTN   (%x)\n", p.Autn)
	dump += fmt.Sprintf("HXres*  (%x)\n", p.HXresStar)
	return dump
}

type AusfAvContext struct {
	Imsi       string
	AuthMethod uint8
	Rand       [Rand_SIZE]byte
	Autn       [Autn_SIZE]byte
	HXresStar  [XResS_SIZE]byte
	KSeaf      [Kseaf_SIZE]byte
}

func (p *AusfAvContext) String() string {
	var dump string
	dump += fmt.Sprintf("imsi (%s)\n", p.Imsi)
	dump += fmt.Sprintf("AuthMethod (%d)\n", p.AuthMethod)
	dump += fmt.Sprintf("RAND   (%x)\n", p.Rand)
	dump += fmt.Sprintf("AUTN   (%x)\n", p.Autn)
	dump += fmt.Sprintf("HXres*  (%x)\n", p.HXresStar)
	dump += fmt.Sprintf("KSeaf  (%x)\n", p.KSeaf)
	return dump
}

type AmfAuthVector struct {
	Rand      [Rand_SIZE]byte
	Autn      [Autn_SIZE]byte
	HXresStar [XResS_SIZE]byte
	Kamf      [Kamf_SIZE]byte
}

func (p *AmfAuthVector) String() string {
	var dump string
	dump += fmt.Sprintf("RAND   (%x)\n", p.Rand)
	dump += fmt.Sprintf("AUTN   (%x)\n", p.Autn)
	dump += fmt.Sprintf("HXres*  (%x)\n", p.HXresStar)
	dump += fmt.Sprintf("Kamf  (%x)\n", p.Kamf)
	return dump
}
func (p *AmfAuthVector) Reset() {
	for i := range p.Rand {
		p.Rand[i] = 0
	}
	for i := range p.Autn {
		p.Autn[i] = 0
	}
	for i := range p.HXresStar {
		p.HXresStar[i] = 0
	}
	for i := range p.Kamf {
		p.Kamf[i] = 0
	}
}

type AuthReRsyncData struct {
	Rand [Rand_SIZE]byte
	Auts [Auts_SIZE]byte
}
