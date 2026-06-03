package nassecurity

import (
	"fmt"
	"lite5gc/amf/context/gctxt"
	ausfmgr "lite5gc/ausf/manager"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/udm/arpf/derivevec"
)

func DeriveNasKeys(kamf []byte,
	encAlgo types3gpp.NREncryptAlgo,
	intAlgo types3gpp.NRIntPrctAlgo) ([]byte, []byte, error) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	if kamf == nil {
		return nil, nil, fmt.Errorf("invalid KAMF as input parameter")
	}

	// derive NR EncryptAlgo
	encKey, err := derivevec.DeriveNasKeys(kamf, types.NumNasEncAlg, byte(encAlgo))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to derive Enc Key, error(%s)", err)
	}
	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Derive Algo(%d), NasEncKey(%x)", encAlgo, encKey)

	// derive NR IntPrctAlgo, 33.501 5.11.1.2 A.8
	intKey, err := derivevec.DeriveNasKeys(kamf, types.NumNasIntAlg, byte(intAlgo))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to derive Int Key, error(%s)", err)
	}
	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Derive Algo(%d), NasIntKey(%x)", intAlgo, intKey)

	return encKey, intKey, nil
}

func DeriveGnbKey(kamf []byte, uplinkNasCount uint32, accessType byte) (error, []byte) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "derive kgnb with uplink nas counter(%d)", uplinkNasCount)

	//33.501 A.9 3gpp access - 0x00 0x01, non 3gpp access - 0x00 0x02
	DerKgnb, err := derivevec.DeriveKgnb(kamf, uplinkNasCount, accessType)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to derive kgnb, error(%s)", err)
		return fmt.Errorf("failed to derive kgnb, error(%s)", err), nil
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "derived GNB Key (%x)", DerKgnb)

	return nil, DerKgnb
}

func DeriveAmfKey(imsi types3gpp.Imsi, abba [2]byte) (error, []byte) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	ausfAvCtxt, err := ausfmgr.GetAusfAvContext((&imsi).String())
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to get ModuleAusf Av Context")
		return fmt.Errorf("failed to get ModuleAusf Av Context"), nil
	}

	supi := &types3gpp.Supi{}
	supi.SetImsi(&imsi)
	err = supi.SetType(types3gpp.IMSIType)
	if err != nil {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.INFO, nil, "wrong supi type")
		return fmt.Errorf("wrong supi type"), nil
	}

	//460010000000005
	//supi := []byte{4, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5}

	kamf, err := derivevec.DeriveKamf(ausfAvCtxt.KSeaf[:], []byte(supi.String()), abba[:])
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to derive kamf")
		return fmt.Errorf("failed to derive kamf"), nil
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "derived AmfKey (%x)", kamf)

	return nil, kamf
}

func DeriveSecurityKeys(p *gctxt.UeContext) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	//drive kamf
	err, kamf := DeriveAmfKey(p.GetImsi(), p.TempSecCtxt.Abba)
	if err != nil {
		return fmt.Errorf("failed to derive kamf, error(%s)", err)
	}

	//drive security key
	encKey, intKey, err := DeriveNasKeys(kamf, p.TempSecCtxt.CipheringAlg, p.TempSecCtxt.IntegrityAlg)
	if err != nil {
		return fmt.Errorf("failed to DeriveSecurityKyes, error(%s)", err)
	}

	//drive kgnb
	err, kgnb := DeriveGnbKey(kamf, p.TempSecCtxt.UplinkNasCount, gctxt.Access3GPP)
	if err != nil {
		return fmt.Errorf("failed to derive kgnb error(%s)", err)
	}

	//save the key in temp context
	p.TempSecCtxt.Kamf = kamf
	p.TempSecCtxt.KnasEncKey = encKey
	p.TempSecCtxt.KnasIntKey = intKey
	p.TempSecCtxt.Kgnb = kgnb

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, p, "derive keys:", p.SecurityCtxt)

	return nil
}

func UpdateDownlinkNasCounter(ueCtxt *gctxt.UeContext) {
	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)
	if ueCtxt == nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "invalid ue context")
		return
	}

	nasCount, overflow := IncreaseNasCounter(ueCtxt.DownlinkNasCount)
	if overflow {
		ueCtxt.ForceAuthNeed = true
	}
	ueCtxt.DownlinkNasCount = nasCount
	return
}
