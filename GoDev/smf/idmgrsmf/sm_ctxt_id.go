package idmgrsmf

import (
	"lite5gc/cmn/idmgr64"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func ReturnSEID(seid uint64) error {
	rlogger.FuncEntry(types.ModuleSmf, nil)
	err := idmgr64.GetInst().ReturnID(string(types.SmfSEID), seid)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.DEBUG, nil,  "failed to return seid(%d)", seid)
		return err
	}
	return nil
}

func RetrieveScId(smCtxtId uint32) uint32 {
	return (smCtxtId & types.SmcScidAnd) >> types.SmcScidShift
}

// SMF assign the SEID uint64 类型，核心网内唯一，会话级别
func GetSEID() (uint64, error) {
	rlogger.FuncEntry(types.ModuleSmf, nil)
	seid, err := idmgr64.GetInst().BorrowID(string(types.SmfSEID))
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "SMF assign the SEID failed:%s", err)
	}
	return seid, err
}
