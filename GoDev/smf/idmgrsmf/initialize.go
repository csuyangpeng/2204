package idmgrsmf

import (
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/idmgr64"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func SmfIdMgrInit() error {
	rlogger.FuncEntry(types.ModuleSmf, nil)

	err := smfRegisterIdMgr()
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to register id mgr")
		return types.ErrFailRegisterIdMgr
	}

	err = smfReserveIdMgr()
	if err != nil {
		rlogger.Trace(types.ModuleCmnIdMgr, rlogger.ERROR, nil, "fail to reserve id")
		return types.ErrFailReserveIdMgr
	}

	return nil
}

func smfRegisterIdMgr() error {
	rlogger.FuncEntry(types.ModuleSmf, nil)

	err := idmgr.GetInst().RegisterIDMgr(string(types.SmfTEID), types.MaxNumSmfTeid)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to register TEID")
		return types.ErrFailRegisterIdMgr
	}

	err = idmgr.GetInst().RegisterIDMgr(string(types.SmfSc), types.MaxSmfScInst)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to register SC id mgr")
		return types.ErrFailRegisterIdMgr
	}

	err = idmgr64.GetInst().RegisterIDMgr(string(types.SmfSEID), types.MaxNumSmfSeid)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to register SEID")
		return types.ErrFailRegisterIdMgr
	}

	err = idmgr.GetInst().RegisterIDMgr(string(types.PDRID), types.MaxNumPdrid)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to register PDR ID")
		return types.ErrFailRegisterIdMgr
	}

	return nil
}

func smfReserveIdMgr() error {
	rlogger.FuncEntry(types.ModuleSmf, nil)

	//29281_CR0037R2_(Rel-10)_C4-110320(e.g. all 0's, or all 1's)
	err := idmgr.GetInst().ReserveID(string(types.SmfTEID), 0)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to reserve TEID id")
		return types.ErrFailReserveIdMgr
	}

	err = idmgr.GetInst().ReserveID(string(types.SmfTEID), 0xffffffff)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to reserve TEID id")
		return types.ErrFailReserveIdMgr
	}

	//smf.config 已經配置了兩個pdr
	err = idmgr.GetInst().ReserveID(string(types.PDRID), 1)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to reserve PDR id")
		return types.ErrFailReserveIdMgr
	}

	err = idmgr.GetInst().ReserveID(string(types.PDRID), 2)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to reserve PDR id")
		return types.ErrFailReserveIdMgr
	}

	return nil
}
