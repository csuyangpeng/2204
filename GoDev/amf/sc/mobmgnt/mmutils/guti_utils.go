package mmutils

import (
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
)

func Allocate5GGuti() (types3gpp.Guti5G, error) {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	var guti types3gpp.Guti5G
	// plmn
	plmn := configure.AmfConf.PlmnList.List[0]
	plmn.ByteOrder = types3gpp.LittleEndian
	guti.SetPlmn(&plmn)

	// amf id
	amfId := configure.AmfConf.Service.AmfIdentifier
	guti.SetAmfID(&amfId)

	// generate tmsi
	tmsi, err := idmgr.GetInst().BorrowID(string(types.TMSI))
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil, "borrow TMSI ID failed.", err)
		return guti, fmt.Errorf("borrow TMSI ID failed, err: %s", err)
	}
	// add the restart counter int the first byte
	rstCounter := utils.GetRestartCounter()
	rstCounter = rstCounter << 28
	tmsi = (tmsi & types.AmfTmsiFilter) | rstCounter
	guti.SetTmsi(tmsi)

	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil, " the Allocate 5G Guti: %s", guti.String())
	return guti, nil
}

func Check5GGuti(guti types3gpp.Guti5G) (bool, error) {
	rlogger.FuncEntry(types.ModuleAmfMM, &guti)

	idList, err := idmgr.GetInst().GetIDList(types.ModuleName(types.TMSI))
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, nil, "fail to get id Mgr")
		return false, types.ErrFailGetIdMgr
	}
	var isTmsiAssigned bool
	for _, v := range idList {
		tmsiInGuti := guti.GetTmsi() & types.AmfTmsiFilter
		if tmsiInGuti == v {
			rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, nil, "tmsiInGuti:%d", tmsiInGuti)
			isTmsiAssigned = true
		}
	}
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, nil, "isTmsiAssigned:%x", isTmsiAssigned)

	gutiPlmn := guti.GetPlmn().GetValue(types3gpp.BigEndian)

	configPlmn := configure.AmfConf.PlmnList.List[0].GetValue(types3gpp.BigEndian)
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, nil, "gutiPlmn, configPlmn:%x", gutiPlmn, configPlmn)

	gutiAMFId := *guti.GetAmfID()
	configAMFId := configure.AmfConf.Service.AmfIdentifier
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, nil, "gutiAMFId:%x, configAMFId:%x", gutiPlmn, configPlmn)

	if gutiPlmn == configPlmn && gutiAMFId == configAMFId && isTmsiAssigned {
		return true, nil
	}
	return false, fmt.Errorf("the guti not assigned by AMF")
}

func AllocatingGUTIResources(ueCtxt *gctxt.UeContext) nas.Mm5gCause {
	rlogger.FuncEntry(types.ModuleAmfMM, ueCtxt.GetImsiPtr())

	guti5gValue, err := Allocate5GGuti()
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "when allocate 5G Guti , err: ", err)
		return nas.SystemFailure
	}
	ueCtxt.Guti5g = &guti5gValue

	err = gctxt.AddIndexUeContext(gctxt.GutiKey(guti5gValue.String()), ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to add Guti5g for ue context table", err)
		return nas.SystemFailure
	}

	tmsiKey := ueCtxt.Guti5g.GetStmsiKey()
	err = gctxt.AddIndexUeContext(gctxt.StmsiKey(tmsiKey), ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to add Stmsi for ue context table, err (%s),tmsi Key (%s) ", err, tmsiKey)
		return nas.SystemFailure
	}

	return nas.SuccessAccept
}

func RecyclingGUTIResources(ueCtxt *gctxt.UeContext) nas.Mm5gCause {
	rlogger.FuncEntry(types.ModuleAmfMM, ueCtxt.GetImsiPtr())

	//return old resource
	err := idmgr.GetInst().ReturnID(string(types.TMSI), ueCtxt.Guti5g.GetTmsi())
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to return TMSI id")
		return nas.SystemFailure
	}

	err = gctxt.DeleteUeContext(gctxt.GutiKey(ueCtxt.Guti5g.String()), ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to release ueCtxt by index(Guti5g:%d), error(%s)", ueCtxt.Guti5g.String(), err)
		return nas.SystemFailure
	}

	err = gctxt.DeleteUeContext(gctxt.StmsiKey(ueCtxt.Guti5g.GetStmsiKey()), ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to release ueCtxt by index(Stmsi:%d), error(%s)", ueCtxt.Guti5g.GetStmsiKey(), err)
		return nas.SystemFailure
	}

	return nas.SuccessAccept
}
