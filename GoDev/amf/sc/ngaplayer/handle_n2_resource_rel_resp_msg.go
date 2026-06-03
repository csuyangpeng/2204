package ngaplayer

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func (p *LayerMgr) handleN2ResRelRespMsg() error {
	rlogger.FuncEntry(types.ModuleAmfNgap,nil)
	rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, nil, "received handleN2ResRelRespMsg, don't do anything."+
		"there should be have another UL NAS msg carrying session release complete")
	return nil
}
