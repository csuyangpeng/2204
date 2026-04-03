package naslayer

import (
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

type IntegrationStatusType uint8

const (
	IntegrityFailed    = 0
	IntegrityValidated = 1
)

func ValidateProtectedMsg(ueCtxt *gctxt.UeContext, nasSeqNum uint8, mac []byte, nasMsg []byte) IntegrationStatusType {
	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)

	var intPrctResult IntegrationStatusType = IntegrityFailed

	integrityAlg := ueCtxt.IntegrityAlg
	uplinkNasCount := ueCtxt.UplinkNasCount
	knasIntKey := ueCtxt.KnasIntKey

	NasCount := uplinkNasCount
	overflow := nassecurity.StoreUplinkNasCounter(nasSeqNum, &NasCount)

	bearId := gctxt.Access3GPP
	direction := types.NasUplink

	var plainNasMsg []byte
	plainNasMsg = append(plainNasMsg, nasSeqNum)
	plainNasMsg = append(plainNasMsg, nasMsg...)

	if nassecurity.VerifyIntegrity(integrityAlg,
		knasIntKey,
		NasCount,
		bearId,
		direction,
		plainNasMsg,
		mac) {
		intPrctResult = IntegrityValidated
	}

	// if pass the verfication, store the uplink nas counter
	if intPrctResult == IntegrityValidated {
		ueCtxt.UplinkNasCount = NasCount
		if overflow {
			ueCtxt.SecurityCtxt.ForceAuthNeed = true
		}
	}

	return intPrctResult
}

func DecipherMessage(ueCtxt *gctxt.UeContext, cipheredMsg []byte) ([]byte, error) {
	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)

	cipheringAlg := ueCtxt.CipheringAlg
	uplinkNasCount := ueCtxt.UplinkNasCount
	knasEncKey := ueCtxt.KnasEncKey
	direction := types.NasUplink
	bearId := gctxt.Access3GPP

	plainNasMsg, err := nassecurity.DecipherNasMsg(cipheringAlg,
		knasEncKey,
		uplinkNasCount,
		bearId,
		direction,
		cipheredMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ueCtxt, "failed to decipher nas message, error(%s)", err)
		return nil, fmt.Errorf("failed to decipher nas message, error(%s)", err)
	}

	return plainNasMsg, nil
}
