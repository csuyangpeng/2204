package idmgramf

import (
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func AmfIdMgrInit() error {
	rlogger.FuncEntry(types.ModuleAmfIdmgr, nil)

	err := amfRegisterIdMgr()
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "fail to register id mgr")
		return types.ErrFailRegisterIdMgr
	}

	err = amfReserveIdMgr()
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "fail to reserve id")
		return types.ErrFailReserveIdMgr
	}
	return nil
}

func amfRegisterIdMgr() error {
	rlogger.FuncEntry(types.ModuleAmfIdmgr, nil)

	err := idmgr.GetInst().RegisterIDMgr(string(types.SC), types.MaxScInst)
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "fail to register SC id mgr")
		return types.ErrFailRegisterIdMgr
	}

	err = idmgr.GetInst().RegisterIDMgr(string(types.NGAP), types.MaxNgapInst)
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "fail to register NGAP id mgr")
		return types.ErrFailRegisterIdMgr
	}

	err = idmgr.GetInst().RegisterIDMgr(string(types.AMFUeNgapId), types.MaxNumAmfNgapId)
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "fail to register AMFUeNgapId id mgr")
		return types.ErrFailRegisterIdMgr
	}

	err = idmgr.GetInst().RegisterIDMgr(string(types.TMSI), types.MaxNumTmsi)
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "fail to register TMSI id mgr")
		return types.ErrFailRegisterIdMgr
	}
	return nil
}

func amfReserveIdMgr() error {
	rlogger.FuncEntry(types.ModuleAmfIdmgr, nil)

	err := idmgr.GetInst().ReserveID(string(types.SC), 0)
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "fail to reserve SC id 0")
		return types.ErrFailReserveIdMgr
	}

	err = idmgr.GetInst().ReserveID(string(types.TMSI), 1)
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "fail to reserve TMSI id")
		return types.ErrFailReserveIdMgr
	}

	// 23003 f60 2.4
	// The network shall not allocate a TMSI with all 32 bits equal to 1
	// (this is because the TMSI must be stored in the SIM,
	// and the SIM uses 4 octets with all bits equal to 1 to indicate that no valid TMSI is available).
	// skip 1111 1111 1111 1111 1111 1111 1111 1111
	err = idmgr.GetInst().ReserveID(string(types.TMSI), 0xffffffff)
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "fail to reserve TMSI id")
		return types.ErrFailReserveIdMgr
	}
	return nil
}
