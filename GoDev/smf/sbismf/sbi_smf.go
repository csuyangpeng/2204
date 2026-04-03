package sbismf

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	namf_comm "lite5gc/openapi/Namf_Communication"
	nudm_sdm "lite5gc/openapi/Nudm_SubscriberDataManagement"
)

// global SbiSmf Layer pointer
var SbiSmfLayer *SbiSmf

type SbiSmf struct {
	MsgRouter *routeragent.MsgRouter

	nudmSdmApiClt  *nudm_sdm.APIClient
	namfCommApiClt *namf_comm.APIClient
}

func NewSbiSmf(appContext *types.AppContext) (sbi *SbiSmf, err error) {
	rlogger.FuncEntry(types.ModuleSmfSbi, nil)
	if appContext == nil {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "input para is nil")
		return nil, types.ErrInputParaNil
	}
	ctxt := appContext.Ctx
	routerChan, ok := ctxt.Value(types.RouterPublishChanCK).(router.DataChannel)
	if !ok {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "failed to get router publish channel.")
		return
	}
	routerCtrlChan, ok := ctxt.Value(types.RouterCtrlChanCK).(router.CtrlChannel)
	if !ok {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return
	}

	myEndpoint := router.RouteAddr{
		Type: router.SbiGR,
		Id:   1, // only one sbi per process
	}
	msgRouter := routeragent.NewMsgRouter(routerCtrlChan, routerChan, myEndpoint)

	sbi = &SbiSmf{
		MsgRouter: msgRouter,
	}

	return
}

func (p *SbiSmf) Initialize() {
	rlogger.FuncEntry(types.ModuleSmfSbi, nil)
	if p != nil {
		//register handler for msg router
		p.MsgRouter.RegisterHandler(router.SmfScGR, p.HandleSmfScMsg)

		//create api client
		p.CreateSMFClientToUDM()
		p.CreateSMFClientToAMF()
	} else {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "input para is nil")
	}
}

func Start(appContext *types.AppContext) {
	if appContext != nil {
		var err error
		SbiSmfLayer, err = NewSbiSmf(appContext)
		if err != nil {
			rlogger.Trace(types.ModuleSmfSbi, rlogger.FATAL, nil, "failed to create "+
				"SBI interface go routine, err:", err)
			panic(err)
		}

		SbiSmfLayer.Initialize()

		go SbiSmfLayer.Run(appContext)
	} else {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "input para is nil")
	}
}

func (p *SbiSmf) Run(appContext *types.AppContext) error {
	rlogger.FuncEntry(types.ModuleSmfSbi, nil)
	if appContext != nil {
		p.MsgRouter.Activate()

		appContext.Wg.Add(1)
		defer func() {
			appContext.Wg.Done()
		}()

		p.MsgRouter.LoopStart(appContext.Ctx)
	} else {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.ERROR, nil, "input para is nil")
		return types.ErrInputParaNil
	}
	return nil
}

func (p *SbiSmf) CreateSMFClientToAMF() {
	addr := fmt.Sprintf("%s://%s:%d", configure.SmfConf.Sbi.Amf.Scheme,
		configure.SmfConf.Sbi.Amf.Addr.Ip, configure.SmfConf.Sbi.Amf.Addr.Port)
	rlogger.Trace(types.ModuleSmfSbi, rlogger.INFO, nil, "SMFClientToAMF addr:%s", addr)

	commcfg := namf_comm.NewConfiguration()
	commcfg.SetBasePath(addr)
	p.namfCommApiClt = namf_comm.NewAPIClient(commcfg)
}

func (p *SbiSmf) CreateSMFClientToUDM() {
	addr := fmt.Sprintf("%s://%s:%d", configure.SmfConf.Sbi.Udm.Scheme,
		configure.SmfConf.Sbi.Udm.Addr.Ip, configure.SmfConf.Sbi.Udm.Addr.Port)
	rlogger.Trace(types.ModuleSmfSbi, rlogger.INFO, nil, "SMFClientToUDM addr:%s", addr)

	sdmcfg := nudm_sdm.NewConfiguration()
	sdmcfg.SetBasePath(addr)
	p.nudmSdmApiClt = nudm_sdm.NewAPIClient(sdmcfg)

}

func SendScMsg(msg *sbicmn.SbiMessage) {
	rlogger.FuncEntry(types.ModuleSmfSbi, nil)
	if msg != nil {
		sendMsgBuf := &router.DataMsg{
			DestAddr: router.RouteAddr{
				Type: router.SmfScGR,
				Id:   msg.ScInstId,
			},
			MsgData: msg,
		}
		SbiSmfLayer.MsgRouter.SendMessage(sendMsgBuf)
	} else {
		rlogger.Trace(types.ModuleSmfSbi, rlogger.INFO, nil, "input para is nil")
	}
}
