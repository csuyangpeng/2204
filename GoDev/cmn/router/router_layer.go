package router

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func Start(appContext *types.AppContext) (CtrlChannel, DataChannel) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	routerMgr := NewRouterMgr(appContext)

	routerMgr.run(appContext)

	// channels for input for router goroutine
	return routerMgr.ctrlChan, routerMgr.dataChan
}
