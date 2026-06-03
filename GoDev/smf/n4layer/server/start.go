package server

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
	"net"
)

var n4Layer *N4Layer

type N4Layer struct {
	UPFIP     net.IP
	N3UpfIP   net.IP
	MsgRouter *routeragent.MsgRouter
	// current node
	node *pfcpv1.Node

	ScId uint32
}

func NewN4Layer() *N4Layer {
	return &N4Layer{}
}

func (p *N4Layer) Initialize(appContext *types.AppContext) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//p.UPFIP = configure.SmfConf.N4Conf.UPFIP
	//p.N3UpfIP = configure.SmfConf.N3Conf.UPFIP

	errN4 := StartN4Server(appContext)
	if errN4 != nil {
		panic("Failed to start n4 socket server ")
	}

	return nil
}

func (p *N4Layer) Start(appContext *types.AppContext) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	n4Layer = p
	ctxt := appContext.Ctx
	// 默认处理节点
	//add router agent
	routerChan, ok := ctxt.Value(types.RouterPublishChanCK).(router.DataChannel)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to get router publish channel.")
		return fmt.Errorf("failed to get router publish channel.")
	}
	routerCtrlChan, ok := ctxt.Value(types.RouterCtrlChanCK).(router.CtrlChannel)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return fmt.Errorf("failed to get router ctrl publish channel.")
	}

	myEndpoint := router.RouteAddr{
		Type: router.PfcpNodeGR,
		Id:   0, // 默认处理节点
	}
	msgRouter := routeragent.NewMsgRouter(routerCtrlChan, routerChan, myEndpoint)

	//register handler for msg router
	msgRouter.RegisterHandler(router.SmfScGR, p.HandleDefaultNodeSessionMsg)

	msgRouter.Activate()
	p.MsgRouter = msgRouter
	ctxt = context.WithValue(ctxt, types.MsgRouterCK, p.MsgRouter)
	ctxt = context.WithValue(ctxt, types.SessSmfScCK, p.ScId)

	go func() {
		appContext.Wg.Add(1)
		defer func() {
			appContext.Wg.Done()
		}()
		goid := utils.Goid()
		ctxt = context.WithValue(ctxt, types.GoroutineID, goid)
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "main node goroutine id %d,start", goid)
		msgRouter.LoopStart(ctxt)
	}()
	// todo zoujun
	ScId, ok := ctxt.Value(types.SessSmfScCK).(uint32)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "---no sc id found 0")
	} else {
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "---sc id found 0:%d", ScId)
	}
	return nil
}

func (p *N4Layer) HandleDefaultNodeSessionMsg(ctx context.Context, msg *router.DataMsg) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)

	//new a context for message handler
	type key = struct{}
	ctxt_new := context.WithValue(ctx, key{}, "start handle message")

	msgData := msg.MsgData.(pfcpv1.ServiceMsg)
	rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, ctxt_new, "smfsc pfcp message type(%d)", msgData.Msg.ID)
	HandleRequestMsg(ctxt_new, msgData)
	return nil
}

// 对应处理节点
func (p *N4Layer) startRouterAgent(ctxt context.Context, id uint32) (context.Context, error) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//add router agent
	routerChan, ok := ctxt.Value(types.RouterPublishChanCK).(router.DataChannel)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to get router publish channel.")
		return nil, fmt.Errorf("failed to get router publish channel.")
	}
	routerCtrlChan, ok := ctxt.Value(types.RouterCtrlChanCK).(router.CtrlChannel)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return nil, fmt.Errorf("failed to get router ctrl publish channel.")
	}

	myEndpoint := router.RouteAddr{
		Type: router.PfcpNodeGR,
		Id:   id, // 当前处理节点
	}
	msgRouter := routeragent.NewMsgRouter(routerCtrlChan, routerChan, myEndpoint)

	//register handler for msg router
	msgRouter.RegisterHandler(router.SmfScGR, p.HandleNodeSessionMsg)
	msgRouter.Activate()
	p.MsgRouter = msgRouter
	ctxt = context.WithValue(ctxt, types.MsgRouterCK, p.MsgRouter)
	go func() {
		/*appContext.Wg.Add(1)
		defer func() {
			appContext.Wg.Done()
		}()*/
		goid := utils.Goid()
		ctxt = context.WithValue(ctxt, types.GoroutineID, goid)
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "node goroutine id %d,start", goid)
		// 动态释放node router
		msgRouter.LoopStart(ctxt)
	}()

	return ctxt, nil
}

func (p *N4Layer) HandleNodeSessionMsg(cxt context.Context, msg *router.DataMsg) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)

	//new a context for message handler
	type key = struct{}
	ctxt_new := context.WithValue(cxt, key{}, "start handle message")

	msgData := msg.MsgData.(pfcpv1.ServiceMsg)
	rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, ctxt_new, "smfsc pfcp message type(%d)", msgData.Msg.ID)
	HandleRequestMsg(ctxt_new, msgData)
	return nil
}
