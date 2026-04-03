package sbiamf

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	nsmf_pdu_session "lite5gc/openapi/Nsmf_PDUSession"
	nudm_sdm "lite5gc/openapi/Nudm_SubscriberDataManagement"
	nudm_ueau "lite5gc/openapi/Nudm_UEAuthentication"
	nudm_uecm "lite5gc/openapi/Nudm_UEContextManagement"
)

// global SbiSmf Layer pointer
var SbiAmfLayer *SbiAmf

type SbiAmf struct {
	MsgRouter            *routeragent.MsgRouter
	nudmSdmApiClt        *nudm_sdm.APIClient
	nudmUeCtxtApiClt     *nudm_uecm.APIClient
	nudmAuthApiClt       *nudm_ueau.APIClient
	nsmfPDUSessionApiClt *nsmf_pdu_session.APIClient
}

func NewSbiAmf(appContext *types.AppContext) (sbi *SbiAmf, err error) {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)

	ctxt := appContext.Ctx
	routerChan, ok := ctxt.Value(types.RouterPublishChanCK).(router.DataChannel)
	if !ok {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil, "failed to get router publish channel.")
		return
	}
	routerCtrlChan, ok := ctxt.Value(types.RouterCtrlChanCK).(router.CtrlChannel)
	if !ok {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return
	}

	myEndpoint := router.RouteAddr{
		Type: router.SbiGR,
		Id:   1, // only one sbi per process
	}
	msgRouter := routeragent.NewMsgRouter(routerCtrlChan, routerChan, myEndpoint)

	sbi = &SbiAmf{
		MsgRouter: msgRouter,
	}

	return
}
func (p *SbiAmf) Initialize() {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)

	//register handler for msg router
	p.MsgRouter.RegisterHandler(router.ScGR, p.HandleScMsg)

	//create api client
	p.CreateAMFClientToUDM()
	p.CreateAMFClientToSMF()
}

func Start(appContext *types.AppContext) {
	var err error
	SbiAmfLayer, err = NewSbiAmf(appContext)
	if err != nil {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.FATAL, nil, "failed to create "+
			"SBI interface go routine, err:", err)
		panic(err)
	}

	SbiAmfLayer.Initialize()

	go SbiAmfLayer.Run(appContext)
}

func (p *SbiAmf) Run(appContext *types.AppContext) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)

	p.MsgRouter.Activate()

	appContext.Wg.Add(1)
	defer func() {
		appContext.Wg.Done()
	}()

	p.MsgRouter.LoopStart(appContext.Ctx)
	return nil
}

func (p *SbiAmf) CreateAMFClientToSMF() {
	addr := fmt.Sprintf("%s://%s:%d", configure.AmfConf.Sbi.Smf.Scheme,
		configure.AmfConf.Sbi.Smf.Addr.Ip, configure.AmfConf.Sbi.Smf.Addr.Port)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "AMFClientToSMF addr: %s", addr)

	sessioncfg := nsmf_pdu_session.NewConfiguration()
	sessioncfg.SetBasePath(addr)
	p.nsmfPDUSessionApiClt = nsmf_pdu_session.NewAPIClient(sessioncfg)

}

func (p *SbiAmf) CreateAMFClientToUDM() {
	addr := fmt.Sprintf("%s://%s:%d", configure.AmfConf.Sbi.Udm.Scheme,
		configure.AmfConf.Sbi.Udm.Addr.Ip, configure.AmfConf.Sbi.Udm.Addr.Port)
	rlogger.Trace(types.ModuleAmfSbi, rlogger.INFO, nil, "AMFClientToUDM addr: %s", addr)

	sdmcfg := nudm_sdm.NewConfiguration()
	sdmcfg.SetBasePath(addr)
	p.nudmSdmApiClt = nudm_sdm.NewAPIClient(sdmcfg)

	ueCtxtCfg := nudm_uecm.NewConfiguration()
	ueCtxtCfg.SetBasePath(addr)
	p.nudmUeCtxtApiClt = nudm_uecm.NewAPIClient(ueCtxtCfg)

	authcfg := nudm_ueau.NewConfiguration()
	authcfg.SetBasePath(addr)
	p.nudmAuthApiClt = nudm_ueau.NewAPIClient(authcfg)
}

func (p *SbiAmf) SendMsg2SC(scInst uint32, msg *sbicmn.SbiMessage) {
	rlogger.FuncEntry(types.ModuleAmfSbi, p)
	sendMsgBuf := &router.DataMsg{
		DestAddr: router.RouteAddr{
			Type: router.ScGR,
			Id:   scInst,
		},
		MsgData: msg,
	}
	p.MsgRouter.SendMessage(sendMsgBuf)
	return
}

func SendScMsg(msg *sbicmn.SbiMessage) {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)
	sendMsgBuf := &router.DataMsg{
		DestAddr: router.RouteAddr{
			Type: router.ScGR,
			Id:   msg.ScInstId,
		},
		MsgData: msg,
	}
	SbiAmfLayer.MsgRouter.SendMessage(sendMsgBuf)
}
