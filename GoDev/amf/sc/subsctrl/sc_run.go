package subsctrl

import (
	"context"
	"fmt"
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/redisclt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

// Start function: Initial and start sc module
func Start(appContext *types.AppContext) {
	rlogger.FuncEntry(types.ModuleAmfSc, nil)
	fmt.Println("sc is running")

	// less then 255
	scNum := 1 //configure.SysConf.SC.NumOfAmfScInst
	rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "AMF SC number: %d", scNum)

	// clear the sc inst id in redis
	redisclt.Agent.Del(types.AmfScInstId)

	for i := 0; i < scNum; i++ {
		id, err := idmgr.GetInst().BorrowID(string(types.SC))
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.FATAL, nil,
				"failed to borrow SC id, err:", err)
			return
		}
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil,
			"SC instance (%d) created", id)

		appContext.Ctx = context.WithValue(appContext.Ctx, types.ScIdCK, id)

		scCtrl, err := NewSubsCtrl(id, appContext)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.FATAL, nil,
				"failed to create subscriber control module, err:", err)
			panic(err)
		}

		err = scCtrl.SubsCtrlInit(id, appContext)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.FATAL, nil,
				"failed to initialize Subscriber Control Module, err:", err)
			panic(err)
		}

		appContext.Wg.Add(1)
		go scCtrl.Start(appContext)

		// save the sc instance id in redis
		redisclt.Agent.SetAdd(types.AmfScInstId, fmt.Sprintf("%d", id))
	}
}
