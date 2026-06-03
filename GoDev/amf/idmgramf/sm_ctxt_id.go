package idmgramf

import (
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func BorrowSmCtxtId(scInstId uint32) (uint32, error) {

	rlogger.FuncEntry(types.ModuleAmfIdmgr, nil)

	smCtxtId, err := idmgr.GetInst().BorrowID(string(types.SmCtxtId))
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.ERROR, nil, "failed to get sm context id from id manager")
		return 0, err
	}
	// add amf sc instance id in sm ctxt id for message routing
	smCtxtId |= scInstId << types.SmcScidShift
	return smCtxtId, nil
}

func ReturnSmCtxtId(smCtxtId uint32) error {
	rlogger.FuncEntry(types.ModuleAmfIdmgr, nil)

	mask := ^types.SmcScidAnd
	id := smCtxtId & uint32(mask)
	err := idmgr.GetInst().ReturnID(string(types.SmCtxtId), id)
	if err != nil {
		rlogger.Trace(types.ModuleAmfIdmgr, rlogger.DEBUG, nil, "failed to return sm context id(%d)", id)
		return err
	}
	return nil
}

func RetrieveScId(smCtxtId uint32) uint32 {
	return (smCtxtId & types.SmcScidAnd) >> types.SmcScidShift
}
