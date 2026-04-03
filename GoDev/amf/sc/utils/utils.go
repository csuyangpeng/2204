package utils

import (
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func Convert2NgapSNssai(snssai *nasie.SNssai) *types3gpp.Snssai {
	sns := &types3gpp.Snssai{}
	sns.Sst = snssai.Sst
	if snssai.Ind == nasie.SstSd {
		sns.SdPrst = true
		sns.Sd = types3gpp.ConvertSdToU32(snssai.Sd[:], types.BigEndian)
	}
	return sns
}

func GetNgapInstId(ueCtxt *gctxt.UeContext) (uint32, error) {
	rlogger.FuncEntry(types.ModuleAmfNgap, ueCtxt)

	n2ctxt, err := gctxt.GetN2ConnContext(gctxt.AmfUeNgApId(ueCtxt.GetAmfUeNgapId()))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, ueCtxt, "failed to find ne connection context")
		return 0, err
	}

	return n2ctxt.GnbInfo.GnbInstId, nil
}

func EncodeSecPrctNasMsg(ueCtxt *gctxt.UeContext, secHeaderType nas.SecHeaderType,
	plainNasMsg []byte) (nasMsg []byte, err error) {

	rlogger.FuncEntry(types.ModuleAmfNgap, ueCtxt)

	header := nas.MmNasSecMessageHeader{}
	header.ExtendProtoDisc = nas.Epd5gsMobMgntMsg
	header.SecurityHeaderType = secHeaderType

	cipheringAlg := ueCtxt.CipheringAlg
	integrityAlg := ueCtxt.IntegrityAlg
	knasEncKey := ueCtxt.KnasEncKey
	knasIntKey := ueCtxt.KnasIntKey
	downlinkNasCount := ueCtxt.DownlinkNasCount
	direction := types.NasDownlink
	bearid := gctxt.Access3GPP

	var cipherNasMsg = plainNasMsg

	if secHeaderType == nas.IntegrityPrtctCipher || secHeaderType == nas.IntegrityPrtctCipherNewSecCtxt {
		// cipher the message
		cipherNasMsg, err = nassecurity.CipherNasMsg(cipheringAlg,
			knasEncKey,
			downlinkNasCount,
			bearid,
			direction,
			plainNasMsg)
		if err != nil {
			return nil, fmt.Errorf("cipher nas message failed. error(%s)", err)
		}
	}

	//sqn
	header.Sqn = nassecurity.GetUplinkNasCounter(downlinkNasCount)

	// generate mac for nas message
	//if ueCtxt.SecurityCtxt.KnasIntKey == nil {
	//	err := nassecurity.DeriveSecurityKeys(ueCtxt)
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to EncodeSecPrctNasMsg, error(%s)", err)
	//	}
	//}

	var intNasMsg []byte
	intNasMsg = append(intNasMsg, header.Sqn)
	intNasMsg = append(intNasMsg, cipherNasMsg...)

	err, mac := nassecurity.GenerateIntegrityMac(integrityAlg,
		knasIntKey,
		downlinkNasCount,
		bearid,
		direction,
		intNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to generate int mac for nas message. error(%s)", err)
		return nil, fmt.Errorf("failed to EncodeSecPrctNasMsg. error(%s)", err)
	}

	header.Mac = mac

	secNasMsg := header.Encode()

	secNasMsg = append(secNasMsg, cipherNasMsg...)

	return secNasMsg, nil
}
