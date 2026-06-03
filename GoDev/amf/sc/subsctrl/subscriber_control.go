//Package subsctrl is controller of SC goroutie
package subsctrl

import (
	"context"
	"fmt"
	"lite5gc/amf/sc/naslayer"
	"lite5gc/amf/sc/ngaplayer"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/amf/sc/sbilayer"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/redisclt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
)

//SubscriberCtrl is the NasMgr for the whole SC Module
type SubscriberCtrl struct {
	NgapLayer  *ngaplayer.LayerMgr
	NgapSender *ngapsender.NgapSender
	NasLayer   *naslayer.NasMgr
	StateMgr   *statemgr.StateMgr
	TimerMgr   *timermgr.TimerMgr
	MsgRouter  *routeragent.MsgRouter
}

func (p SubscriberCtrl) String() (strbuf string) {
	strbuf += fmt.Sprintln("SubscriberCtrl Info:")
	strbuf += fmt.Sprintln("ScNgLayer:\n", *p.NgapLayer)
	strbuf += fmt.Sprintln("NasLayer:\n", *p.NasLayer)
	strbuf += fmt.Sprintln("StateMgr:\n", *p.StateMgr)
	strbuf += fmt.Sprintln("TimerMgr:\n", *p.TimerMgr)
	strbuf += fmt.Sprintln("MsgRouter:\n", *p.MsgRouter)
	return strbuf
}

//NewSubsCtrl return a SubscriberCtrl pointer
func NewSubsCtrl(scID uint32, appContext *types.AppContext) (sc *SubscriberCtrl, err error) {
	rlogger.FuncEntry(types.ModuleAmfSc, nil)

	ctxt := appContext.Ctx

	routerChan, ok := ctxt.Value(types.RouterPublishChanCK).(router.DataChannel)
	if !ok {
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "failed to get router publish channel")
		return
	}
	routerCtrlChan, ok := ctxt.Value(types.RouterCtrlChanCK).(router.CtrlChannel)
	if !ok {
		rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "failed to get router ctrl publish channel")
		return
	}

	myEndpoint := router.RouteAddr{
		Type: router.ScGR,
		Id:   scID,
	}
	msgRouter := routeragent.NewMsgRouter(routerCtrlChan, routerChan, myEndpoint)

	sc = &SubscriberCtrl{
		NgapLayer: ngaplayer.NewLayerMgr(scID),
		NasLayer:  naslayer.NewLayerMgr(scID),
		StateMgr:  statemgr.NewStateMgr(),
		MsgRouter: msgRouter,
	}
	return
}

//SubsCtrlInit - the initialization for SubscriberCtrl
func (p *SubscriberCtrl) SubsCtrlInit(scID uint32, appContext *types.AppContext) (err error) {
	rlogger.FuncEntry(types.ModuleAmfSc, nil)

	// initialize timer manager
	p.TimerMgr = timermgr.NewTimerMgr(appContext.Ctx, 2, 500)

	//initial scng layer
	p.NgapLayer.Init()
	p.NgapSender = ngapsender.NewNgapSender(scID, p.NgapLayer.GetOssCtxt())

	// clear the redis temp info
	// remove the redis list for amf_sc_<sc id>
	redisclt.Agent.Del(fmt.Sprintf("%s%s", types.AmfProc, fmt.Sprintf("%d", scID)))

	//register handler for msg router
	//err = p.MsgRouter.RegisterHandler(router.GnbGR, p.ScNgLayer.HandleNgapMsg)
	//if err != nil {
	//	rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "fail to process ngap msg")
	//	return fmt.Errorf("fail to process ngap msg")
	//}
	err = p.MsgRouter.RegisterHandler(router.SbiGR, sbilayer.HandleSBIMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "fail to process sbi msg")
		return fmt.Errorf("fail to process sbi msg")
	}
	return
}

//Start - go routine for the Sc Component
func (p *SubscriberCtrl) Start(appContext *types.AppContext) {
	rlogger.FuncEntry(types.ModuleAmfSc, nil)

	ctxt := appContext.Ctx
	ctxt = context.WithValue(ctxt, types.StateMgrCK, p.StateMgr)
	ctxt = context.WithValue(ctxt, types.NasLayerCK, p.NasLayer)
	ctxt = context.WithValue(ctxt, types.NgapLayerCK, p.NgapLayer)
	ctxt = context.WithValue(ctxt, types.NgapSenderCK, p.NgapSender)
	ctxt = context.WithValue(ctxt, types.ScTimerMgrCK, p.TimerMgr)
	ctxt = context.WithValue(ctxt, types.MsgRouterCK, p.MsgRouter)
	p.MsgRouter.Activate()

	defer func() {
		p.stop()
		appContext.Wg.Done()
	}()

	appContext.Wg.Add(1)
	go func() {
		defer appContext.Wg.Done()
		p.NgapLayer.LoopProcN2Msg(ctxt)
	}()

	p.MsgRouter.LoopStart(ctxt)

	rlogger.Trace(types.ModuleAmfSc, rlogger.DEBUG, nil, "return sc")
	return
}

func (p *SubscriberCtrl) stop() {
	rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil,
		"amf subscriber ctrl (%d) stopped", p.NasLayer.ScInstId)

	p.TimerMgr.Destroy()
}
