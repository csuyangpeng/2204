package routeragent

import (
	"context"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/types"
)

type MsgRouter struct {
	routerCtrlCh router.CtrlChannel
	routerDataCh router.DataChannel
	pubEndpoint  router.RouteAddr
	msgChann     router.DataChannel
	handlerMap   MsgHandlerMap
}

func (p MsgRouter) String() (strbuf string) {
	strbuf += fmt.Sprintf("Router Endpoint Info:\n")
	strbuf += fmt.Sprintln("routerCtrlCh:\n", p.routerCtrlCh)
	strbuf += fmt.Sprintln("routerDataCh:\n", p.routerDataCh)
	strbuf += fmt.Sprintln("pubEndpoint:\n", p.pubEndpoint)
	strbuf += fmt.Sprintln("msgChann:\n", p.msgChann)
	strbuf += fmt.Sprintln("handlerMap:\n", p.handlerMap)
	return strbuf
}

type MsgHandler func(ctx context.Context, msg *router.DataMsg) error
type MsgHandlerMap map[router.InstType]MsgHandler

func NewMsgRouter(rCtrlChan router.CtrlChannel,
	rDataChan router.DataChannel,
	endpoint router.RouteAddr) *MsgRouter {

	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	msgRouter := &MsgRouter{
		routerCtrlCh: rCtrlChan,
		routerDataCh: rDataChan,
		pubEndpoint:  endpoint,
		msgChann:     make(router.DataChannel, 2000),
		handlerMap:   make(MsgHandlerMap),
	}

	return msgRouter
}

func (p *MsgRouter) Activate() {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	//register the msg Router
	msg := &router.ControlMsg{}
	msg.SrcAddr = p.pubEndpoint
	msg.Op = router.Register
	msg.PubChannel = p.msgChann

	p.routerCtrlCh <- msg
}

func (p *MsgRouter) Deactivate() {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	//deregister the msg Router
	msg := &router.ControlMsg{}
	msg.SrcAddr = p.pubEndpoint
	msg.Op = router.Deregister
	//msg.PubChannel = p.msgChann

	p.routerCtrlCh <- msg
}

func (p *MsgRouter) LoopStart(ctx context.Context) {
	rlogger.FuncEntry(types.ModuleCmnRouter, ctx)
	rlogger.Trace(types.ModuleCmnRouter, rlogger.INFO, ctx, "MsgRouter LoopStart routine start")
	defer func() {
		//if p := recover(); p != nil {
		//	rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, nil,  "panics: %v", p)
		//}
		rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, ctx, "MsgRouter LoopStart routine exit")
	}()

	for {
		select {
		case <-ctx.Done(): // connection closed
			return

		case msgBuf := <-p.msgChann:
			dataMsg := msgBuf.MsgD.(*router.DataMsg)
			srcEndpiont := dataMsg.SrcAddr
			handler, ok := p.handlerMap[srcEndpiont.Type]
			if !ok {
				rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, ctx, "the message type %s have no handler.", srcEndpiont.Type)
			} else {
				_ = handler(ctx, dataMsg)
			}
		}
	}
}

func (p *MsgRouter) RegisterHandler(srcEpType router.InstType, handler MsgHandler) error {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)
	_, ok := p.handlerMap[srcEpType]
	if ok {
		rlogger.Trace(types.ModuleCmnRouter, rlogger.ERROR, nil, "the handler already registered for type %s", srcEpType)
		return fmt.Errorf("duplicated registed for type %s", srcEpType)
	}
	p.handlerMap[srcEpType] = handler
	return nil
}

func (p *MsgRouter) DeregisterHandler(srcEpType router.InstType) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	delete(p.handlerMap, srcEpType)
}

func (p *MsgRouter) SendMessage(sendMsg *router.DataMsg) {
	rlogger.FuncEntry(types.ModuleCmnRouter, nil)

	sendMsg.SrcAddr = p.pubEndpoint

	SendMsg := &router.IpcMsg{
		MsgT: router.DP,
		MsgD: sendMsg,
	}
	p.routerDataCh <- SendMsg
}
