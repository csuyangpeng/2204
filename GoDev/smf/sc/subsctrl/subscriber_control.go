package subsctrl

import (
	"context"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
	"lite5gc/smf/sc/n4layer"
	"lite5gc/smf/sc/naslayer"
	"lite5gc/smf/sc/sbilayer"
	"lite5gc/smf/sc/sessmgnt"
	"lite5gc/smf/sc/statemgr"
)

//SubscriberCtrl is the NasMgr for the whole SC Module
type SubscriberCtrl struct {
	NasLayer  *naslayer.NasMgr
	StateMgr  *statemgr.SmfStateMgr
	SessMgmt  *sessmgnt.SessMGMT
	TimerMgr  *timermgr.TimerMgr
	MsgRouter *routeragent.MsgRouter
}

//NewSubsCtrl return a SubscriberCtrl pointer
func NewSubsCtrl(scID uint32, appContext *types.AppContext) (sc *SubscriberCtrl, err error) {
	rlogger.FuncEntry(types.ModuleSmf, nil)

	ctxt := appContext.Ctx
	routerChan, ok := ctxt.Value(types.RouterPublishChanCK).(router.DataChannel)
	if !ok {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "failed to get router publish channel.")
		return
	}
	routerCtrlChan, ok := ctxt.Value(types.RouterCtrlChanCK).(router.CtrlChannel)
	if !ok {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return
	}

	myEndpoint := router.RouteAddr{
		Type: router.SmfScGR,
		Id:   scID,
	}
	msgRouter := routeragent.NewMsgRouter(routerCtrlChan, routerChan, myEndpoint)

	sc = &SubscriberCtrl{
		StateMgr:  statemgr.NewStateMgr(),
		SessMgmt:  sessmgnt.NewSessMgmt(scID),
		NasLayer:  &naslayer.NasMgr{},
		MsgRouter: msgRouter,
	}

	return
}

//SubsCtrlInit - the initialization for SubscriberCtrl
func (p *SubscriberCtrl) SubsCtrlInit(appContext *types.AppContext) (err error) {
	// initialize timer manager
	p.TimerMgr = timermgr.NewTimerMgr(appContext.Ctx, 2, 500)

	//initial nas layer manager
	p.SessMgmt.Initialize()

	//register handler for msg router
	err = p.MsgRouter.RegisterHandler(router.PfcpNodeGR, n4layer.HandleNodeSessionMsg)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "fail to process pfcp msg")
		return fmt.Errorf("fail to process pfcp msg")
	}

	err = p.MsgRouter.RegisterHandler(router.SbiGR, sbilayer.HandleSBIMsg)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.INFO, nil, "fail to process sbi msg")
		return fmt.Errorf("fail to process sbi msg")
	}
	return
}

//Start - go routine for the Sc Component
func (p *SubscriberCtrl) Start(appContext *types.AppContext) (err error) {
	rlogger.FuncEntry(types.ModuleSmf, nil)

	ctxt := appContext.Ctx
	ctxt = context.WithValue(ctxt, types.SmfStateMgrCK, p.StateMgr)
	ctxt = context.WithValue(ctxt, types.SmfNasLayerCK, p.NasLayer)
	ctxt = context.WithValue(ctxt, types.SmfScTimerMgrCK, p.TimerMgr)
	ctxt = context.WithValue(ctxt, types.SessMgntCK, p.SessMgmt)
	ctxt = context.WithValue(ctxt, types.MsgRouterCK, p.MsgRouter)
	appContext.Ctx = ctxt

	p.MsgRouter.Activate()

	appContext.Wg.Add(1)
	defer func() {
		p.stop()
		appContext.Wg.Done()
	}()
	p.MsgRouter.LoopStart(ctxt)

	return
}

func (p *SubscriberCtrl) stop() {
	rlogger.Trace(types.ModuleSmf, rlogger.INFO, nil, "SMF SubscriberCtrl %p to stop", p)
	p.TimerMgr.Destroy()
	rlogger.Trace(types.ModuleSmf, rlogger.INFO, nil, "SMF SubscriberCtrl %p stopped", p)
}
