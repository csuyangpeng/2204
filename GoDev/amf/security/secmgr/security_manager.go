package secmgr

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
)

func SelectNasAlgo(ueCtxt *gctxt.UeContext,
	ueSecCap *types3gpp.SecurityCapability) (types3gpp.SelNasSecAlgo, error) {
	rlogger.FuncEntry(types.ModuleAmfSec, ueCtxt.GetImsiPtr())

	// 1 read the security capability from configuration
	configSecCap := configure.AmfConf.NAS.SecCap
	rlogger.Trace(types.ModuleAmfSec, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
		"amf config nas algos(%s)", configSecCap)

	// 2 get the matched security algorithm
	intAlgo := configSecCap.MatchNrIntAlgo(ueSecCap.GetNrIntPrctAlgo())
	encAlgo := configSecCap.MatchNrEncAlgo(ueSecCap.GetNrEncAlgo())
	rlogger.Trace(types.ModuleAmfSec, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
		"selected integrity algorithm(%s)", intAlgo)
	rlogger.Trace(types.ModuleAmfSec, rlogger.DEBUG, ueCtxt.GetImsiPtr(),
		"selected encipher algorithm(%s)", encAlgo)

	var selNasSecAlgo types3gpp.SelNasSecAlgo
	selNasSecAlgo.SetNrIntPrctAlgo(intAlgo)
	selNasSecAlgo.SetNrEncAlgo(encAlgo)

	// 3 save the selected algorithm in ue context
	ueCtxt.SecurityCtxt.TempSecCtxt.IntegrityAlg = intAlgo
	ueCtxt.SecurityCtxt.TempSecCtxt.CipheringAlg = encAlgo

	return selNasSecAlgo, nil
}
