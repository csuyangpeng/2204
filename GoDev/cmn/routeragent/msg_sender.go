package routeragent

import (
	"context"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/types"
)

func SendIpcMessage(ctxt context.Context,
	destType router.InstType,
	destInstId uint32,
	msg router.IpcMsgData) error {
	rlogger.FuncEntry(types.ModuleCmnRouter, ctxt)

	msgRouter, ok := ctxt.Value(types.MsgRouterCK).(*MsgRouter)
	if !ok {
		rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, ctxt, "no msg router found")
		return fmt.Errorf("no msg routerfound")
	}

	sendMsgBuf := &router.DataMsg{
		DestAddr: router.RouteAddr{
			Type: destType,
			Id:   destInstId,
		},
		MsgData: msg,
	}

	msgRouter.SendMessage(sendMsgBuf)

	return nil
}
