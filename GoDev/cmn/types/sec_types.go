package types

import (
	"fmt"
)

const (
	KSize         uint8  = 16 //in Byte. Should be 128 bits (16) or 256 bits (32)
	RandSize      uint8  = 16
	OpSize        uint8  = 16
	OpcSize       uint8  = 16
	AmfSize       uint8  = 2
	AutnSize      uint8  = 16
	AutsSize      uint8  = 14
	SqnSize       uint8  = 6
	MccMncSize    uint8  = 3
	MSINSize      uint8  = 5
	EccPubKeySize uint8  = 65
	SnNameSize    uint16 = 1024
	MacaSize      uint8  = 8
	SqnxorakSize  uint8  = 6
	AkSize        uint8  = 6
	XresSize      uint8  = 8
	CkSize        uint8  = 16
	IkSize        uint8  = 16
	XressSize     uint8  = 16
	RessSize      uint8  = 16
	KausfSize     uint8  = 32
	KseafSize     uint8  = 32
	KamfSize      uint8  = 32
)

type SupiType struct {
	MccMnc [MccMncSize]byte
	Msin   [MSINSize]byte
}

type SuciType struct {
	Mccmnc    [MccMncSize]byte
	Msin      [MSINSize]byte
	EccPubKey [EccPubKeySize]byte
}

type SnName [SnNameSize]byte

type HeAvType struct {
	MACA     [MacaSize]byte
	SqnXorAk [SqnxorakSize]byte
	AK       [AkSize]byte
	XRes     [XresSize]byte
	CK       [CkSize]byte
	IK       [IkSize]byte
	Rand     [RandSize]byte
	Autn     [AutnSize]byte
	XResStar [XressSize]byte
	Kausf    [KausfSize]byte
}

func (p *HeAvType) String() string {
	var dump string
	dump = fmt.Sprintf("\nMAC_A  (%x)\n", p.MACA)
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
	Sqn        [SqnSize]byte
	Key        [KSize]byte
	Op         [OpSize]byte
	IsOpc      bool
	Opc        [OpcSize]byte
	Amf        [AmfSize]byte
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
	dump += fmt.Sprintf("SnName (%s)\n", p.SnName)

	return dump
}

type SeAvType struct {
	Rand      [RandSize]byte
	Autn      [AutnSize]byte
	HXresStar [XressSize]byte
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
	Rand       [RandSize]byte
	Autn       [AutnSize]byte
	HXresStar  [XressSize]byte
	KSeaf      [KseafSize]byte
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
	Rand      [RandSize]byte
	Autn      [AutnSize]byte
	HXresStar [XressSize]byte
	Kamf      [KamfSize]byte
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
	for i, _ := range p.Rand {
		p.Rand[i] = 0
	}
	for i, _ := range p.Autn {
		p.Autn[i] = 0
	}
	for i, _ := range p.HXresStar {
		p.HXresStar[i] = 0
	}
	for i, _ := range p.Kamf {
		p.Kamf[i] = 0
	}
}

type AuthReRsyncData struct {
	Rand [RandSize]byte
	Auts [AutsSize]byte
}

//33.501 A8
//N-NAS-enc-alg	0x01
//N-NAS-int-alg	0x02
//N-RRC-enc-alg	0x03
//N-RRC-int-alg	0x04
//N-UP-enc-alg	0x05
//N-UP-int-alg	0x06
const (
	NumNasEncAlg byte = 0x01
	NumNasIntAlg byte = 0x02
)

// 33.504 6.4.3.1
// The DIRECTION bit shall be set to 0 for uplink and 1 for downlink.
const (
	NasUplink   byte = 0x00
	NasDownlink byte = 0x01
)
