package gctxt

import (
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

type SecurityStatusType uint8

const (
	NoSecurity                     SecurityStatusType = 0
	IntegrityPrtc                  SecurityStatusType = 1
	IntegrityPrtctCipher           SecurityStatusType = 2
	IntegrityPrtctNewSecCtxt       SecurityStatusType = 3
	IntegrityPrtctCipherNewSecCtxt SecurityStatusType = 4
	MaxSecHeaderTypeValue          SecurityStatusType = 5
)

const (
	AccessZero    byte = 0x00
	Access3GPP    byte = 0x01
	AccessNon3Gpp byte = 0x02
)

type SecurityCtxt struct {
	ForceAuthNeed         bool
	IsSecCtxtEstb         bool
	SecHdrTypeIncomingMsg nas.SecHeaderType
	SyncFailureProceeding bool

	Kamf []byte
	Abba [2]byte
	Kgnb []byte

	UeSecCapablity types3gpp.SecurityCapability

	CipheringAlg types3gpp.NREncryptAlgo
	IntegrityAlg types3gpp.NRIntPrctAlgo

	NgKsi nasie.NasKSI

	AuthVector types.AmfAuthVector
	KnasEncKey []byte
	KnasIntKey []byte

	UplinkNasCount   uint32
	DownlinkNasCount uint32

	// 33.501 6.4.2.2
	//The value of the unique NAS connection identifier shall be set to "0x01" for 3GPP access and set to "0x02" for non-3GPP access.
	NasConnId byte

	TempSecCtxt TempSecurityCtxt
}

type TempSecurityCtxt struct {
	Kamf             []byte
	Abba             [2]byte
	Kgnb             []byte
	UplinkNasCount   uint32
	DownlinkNasCount uint32
	CipheringAlg     types3gpp.NREncryptAlgo
	IntegrityAlg     types3gpp.NRIntPrctAlgo
	AuthVector       types.AmfAuthVector
	KnasEncKey       []byte
	KnasIntKey       []byte
}

func (p *TempSecurityCtxt) Reset() {
	p.Kamf = nil
	for i := range p.Abba {
		p.Abba[i] = 0
	}
	p.Kgnb = nil
	p.UplinkNasCount = 0
	p.DownlinkNasCount = 0
	p.CipheringAlg = types3gpp.NEA0
	p.IntegrityAlg = types3gpp.NIA0
	p.AuthVector.Reset()
	p.KnasEncKey = nil
	p.KnasIntKey = nil
}
