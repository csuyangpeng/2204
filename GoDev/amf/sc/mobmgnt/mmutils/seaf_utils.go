package mmutils

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/security/seaf"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
)

func RetrieveAuthVector(ctxt context.Context,
	ueContext *gctxt.UeContext,
	suci *types3gpp.Suci,
	reRsyncInfo *types.AuthReRsyncData) error {

	if ueContext == nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, nil, "invalid input parameter ueContext")
		return types.ErrFailFindUeCtxt
	}

	rlogger.FuncEntry(types.ModuleAmfMM, ueContext.GetImsiPtr())

	// get sn name
	snName := GenerateSnName(configure.AmfConf.PlmnList.List[0])
	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueContext.GetImsiPtr(), "SN Name (%s)", snName)

	err, _, seav := seaf.RetreiveAuthVectorWithSuci(ctxt, suci, snName, reRsyncInfo)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueContext.GetImsiPtr(),
			"failed to retreive auth vection with suci. error(%s)", err)
		return fmt.Errorf("failed to retreive auth vection with suci, error(%s)", err)
	}

	//store seav
	ueContext.TempSecCtxt.AuthVector.Rand = seav.Rand
	ueContext.TempSecCtxt.AuthVector.Autn = seav.Autn
	ueContext.TempSecCtxt.AuthVector.HXresStar = seav.HXresStar
	ueContext.TempSecCtxt.Abba = [2]byte{0x00, 0x00}
	ueContext.TempSecCtxt.KnasIntKey = nil
	ueContext.TempSecCtxt.KnasEncKey = nil
	ueContext.TempSecCtxt.Kamf = nil
	ueContext.TempSecCtxt.Kgnb = nil

	return nil
}
