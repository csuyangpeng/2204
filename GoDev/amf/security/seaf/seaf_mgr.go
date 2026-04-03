package seaf

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	ausfmgr "lite5gc/ausf/manager"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/udm/arpf/derivevec"
)

type SeafMgr struct {
	ABBApara []byte
}

func RetreiveAuthVectorWithSuci(ctxt context.Context,
	suci *types3gpp.Suci,
	snName string,
	reRsyncInfo *types.AuthReRsyncData) (error, *types3gpp.Supi, *types.SeAvType) {

	rlogger.FuncEntry(types.ModuleAmfSec, suci)

	err, supi, seav := ausfmgr.HandleUeAuthenticationReqMsg(ctxt, suci, snName, reRsyncInfo)

	if err != nil {
		return fmt.Errorf("failed to handle ue auth req msg"), nil, nil
	}

	return err, supi, seav
}

func RetreiveAuthVectorWithSupi(ctxt context.Context,
	supi *types3gpp.Supi,
	snName string,
	reRsyncInfo *types.AuthReRsyncData) (error, *types3gpp.Supi, *types.SeAvType) {
	rlogger.FuncEntry(types.ModuleAmfSec, supi)

	suci := &types3gpp.Suci{}
	suci.SetFromImsi(supi.GetImsi())

	err, supi, seav := ausfmgr.HandleUeAuthenticationReqMsg(ctxt, suci, snName, reRsyncInfo)

	if err != nil {
		return fmt.Errorf("failed to handle ue auth req msg"), nil, nil
	}

	return err, supi, seav
}

func HandleAuthenticationResponse(ueCtxt *gctxt.UeContext, ResStart []byte) error {
	rlogger.FuncEntry(types.ModuleAmfSec, ueCtxt)

	if ueCtxt == nil || ResStart == nil {
		return fmt.Errorf("invalid UeCtxt or ResStart as input parameter")
	}

	HXResStar := ueCtxt.TempSecCtxt.AuthVector.HXresStar

	// Setup 1: Calculate HRES* from RES* in authentication response message
	err, HResStar := derivevec.DeriveHResS(ResStart, ueCtxt.TempSecCtxt.AuthVector.Rand[:])
	if err != nil {
		rlogger.Trace(types.ModuleAmfSec, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to derive HRes*, err(%s)", err)
		return fmt.Errorf("failed to derive HRes*, err(%s)", err)
	}
	rlogger.Trace(types.ModuleAmfSec, rlogger.INFO, ueCtxt.GetImsiPtr(),
		"generate HRes*(%x) from Res*(%x) and Rand(%x)",
		HResStar, ResStart, ueCtxt.TempSecCtxt.AuthVector.Rand)

	// Setup 2: Compare  HRES* with HXRES*
	if HResStar != HXResStar {
		rlogger.Trace(types.ModuleAmfSec, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"HRES*(%x), mismatch with HXRES*(%x)", HResStar, HXResStar)
		return fmt.Errorf("Res* mismatch with XRES*")
	}
	rlogger.Trace(types.ModuleAmfSec, rlogger.DEBUG, ueCtxt.GetImsiPtr(), "Res* validation passed.")

	// Setup 3: send RES* to AUSF for further verification
	// TODO send RES* to AUSF

	return nil
}

func GetKgnb(ueCtxt *gctxt.UeContext) ([]byte, error) {
	rlogger.FuncEntry(types.ModuleAmfSec, ueCtxt)

	if ueCtxt == nil {
		return nil, fmt.Errorf("invalid UeCtxt or ResStart as input parameter")
	}

	var Kgnb []byte
	if ueCtxt.SecurityCtxt.Kgnb == nil {
		rlogger.Trace(types.ModuleAmfSec, rlogger.ERROR, ueCtxt.GetImsiPtr(), "invalid kgnb key")
		return nil, fmt.Errorf("invalid kgnb in ue security context")
	}

	Kgnb = ueCtxt.SecurityCtxt.Kgnb

	return Kgnb, nil
}
