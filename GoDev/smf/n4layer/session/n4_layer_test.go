package session_test

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/smf/n4layer/server"
	"lite5gc/smf/sc/smfcontext/prcdctxt"
	"net"
	"sync"
	"testing"
	"time"

	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/logger"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"

	n4layer "lite5gc/smf/n4layer/session"
	"lite5gc/smf/sc/smfcontext"
)

func TestCreateN4SessionReqMsgPaging(t *testing.T) {
	fmt.Println("SMF Process Stared.")
	smfCtxt := &types.AppContext{
		Name:     "smf",
		Wg:       &sync.WaitGroup{},
		ConfPath: types.DefConfFileSmf, //DEF_CONF_FILE,
	}
	smfCtxt.Ctx, smfCtxt.Cancel = context.WithCancel(context.Background())
	configure.LoadConfigCmn("ini", types.DefConfFileAmf)

	//load config
	var confPath = types.DefConfFileSmf
	//F:\work\go\src\5gc\oam\cm\configure\files\smf.config
	//          "src/5gc/oam/cm/configure/files/smf.config"
	err := configure.LoadConfigSmf("ini", confPath)
	if err != nil {
		panic("load config file failed")
	}
	////tag
	//start logger
	//logger.Init(configure.CmnConf.Logger.LogPath, configure.CmnConf.Logger.LogLevel)
	logger.Init()
	if configure.CmnConf.Logger.LogFile {
		logger.SetLoggerToFile(configure.CmnConf.Logger.LogPath, configure.CmnConf.Logger.LogLevel)
	}
	if configure.CmnConf.Logger.LogConsole {
		logger.SetLoggerToConsole(configure.CmnConf.Logger.LogLevel)
	}
	//id manager register ids
	idmgr.GetInst().RegisterIDMgr(string(types.DPE), types.MaxDpeInst)
	// n4 layer
	n4 := server.NewN4Layer()
	n4.Initialize(smfCtxt)
	n4.Start(smfCtxt)

	// 当前开启N4 tcp server
	errN4 := server.StartN4Server(smfCtxt)
	if errN4 != nil {
		panic("Failed to start n4 socket server ")
	}

	smfUeCxt := &smfcontext.UeContext{}
	psi := nas.PduSessID(1)
	smfUeCxt.IMSI.StoreWithNAI("1234565678@cn", 8)

	smfUeCxt.PduSessCtxts = make(map[nas.PduSessID]*smfcontext.PduSessContext, 0)

	//pduSessCtxt := smfcontext.PduSessContext{}
	pduSessCtxt := smfcontext.NewPduSessContext(psi)
	smfUeCxt.PduSessCtxts[psi] = pduSessCtxt

	pduSessCtxt.PDR = configure.SmfConf.N4Conf.PDR
	pduSessCtxt.FAR = configure.SmfConf.N4Conf.FAR

	pduSessCtxt.UEIP = net.ParseIP("192.1.2.3")
	// UPInactiveTimer
	pduSessCtxt.UPInactiveTimer = configure.SmfConf.N4Conf.UPInactiveTimer

	//pduSessCtxt.PDR[0].PDI.UEIPaddress = &pfcp.IEUEIPaddress{}
	//create a pdu session establishment prcd ctxt
	prcdCtxt := prcdctxt.NewPduSessEstbPrcdCtxt(psi)
	pduSessCtxt.SetPrcdCtxt(prcdCtxt)

	//store the information in procedure ctxt
	//prcdCtxt.IMSI = ueCtx
	prcdCtxt.PduSessId = pduSessCtxt.PduSessionId
	//prcdCtxt.Pti = p.PduSessEstbRequest.Pti
	//prcdCtxt.PduSessType = p.PduSessEstbRequest.SessionType
	//prcdCtxt.SscMode = p.PduSessEstbRequest.SscMode
	//prcdCtxt.SmCap = p.PduSessEstbRequest.SmCapability
	//prcdCtxt.MaxNumofSupPktFilter = p.PduSessEstbRequest.MaxNumberOfSPF
	//prcdCtxt.AlwaysOnPduSessReq = p.PduSessEstbRequest.AlwaysOn
	//prcdCtxt.SessionReqIeFlags = p.PduSessEstbRequest.IeFlags

	err = n4layer.N4SessionEstablishmentRequest(pduSessCtxt)
	if err != nil {
		t.Errorf("Failed to create message：%s", err)
	}
	time.Sleep(time.Second * 1)

	for _, far := range pduSessCtxt.FAR {
		far.ForwardingParameters.OuterHeaderCreation = &pfcp.IEOuterHeaderCreation{TEID: 0, IPv4Addr: []byte{192, 0, 2, 20}}
	}
	// 去激活N3
	//pduSessCtxt.FAR[0].ApplyAction.Set(12)
	//pduSessCtxt.FAR[1].ApplyAction.Set(12)
	err = n4layer.N4SessionModificationRequest(pduSessCtxt.SEID, pduSessCtxt.FAR)
	if err != nil {
		t.Errorf("Failed to modify message：%s", err)
	}
	/*seid := smfUeCxt.PduSessCtxts[psi].SEID
	err = n4layer.N4SessionReleaseRequest(seid)
	if err != nil {
		t.Errorf("Failed to release message：%s", err)
	}*/
	// 重复发送
	err = n4layer.N4SessionModificationRequest(pduSessCtxt.SEID, pduSessCtxt.FAR)
	if err != nil {
		t.Errorf("Failed to modify message：%s", err)
	}
	err = n4layer.N4SessionModificationRequest(pduSessCtxt.SEID, pduSessCtxt.FAR)
	if err != nil {
		t.Errorf("Failed to modify message：%s", err)
	}
	err = n4layer.N4SessionModificationRequest(pduSessCtxt.SEID, pduSessCtxt.FAR)
	if err != nil {
		t.Errorf("Failed to modify message：%s", err)
	}
	smfCtxt.Wg.Wait()
}
func TestCreateN4SessionReqMsg(t *testing.T) {
	fmt.Println("SMF Process Stared.")
	smfCtxt := &types.AppContext{
		Name:     "smf",
		Wg:       &sync.WaitGroup{},
		ConfPath: types.DefConfFileSmf, //DEF_CONF_FILE,
	}
	smfCtxt.Ctx, smfCtxt.Cancel = context.WithCancel(context.Background())
	configure.LoadConfigCmn("ini", types.DefConfFileAmf)

	//load config
	var confPath = types.DefConfFileSmf
	//F:\work\go\src\5gc\oam\cm\configure\files\smf.config
	//          "src/5gc/oam/cm/configure/files/smf.config"
	err := configure.LoadConfigSmf("ini", confPath)
	if err != nil {
		panic("load config file failed")
	}
	////tag
	//start logger
	logger.Init(configure.CmnConf.Logger.LogPath, configure.CmnConf.Logger.LogLevel)

	//id manager register ids
	idmgr.GetInst().RegisterIDMgr(string(types.DPE), types.MaxDpeInst)
	// n4 layer
	n4 := server.NewN4Layer()
	n4.Initialize(smfCtxt)
	n4.Start(smfCtxt)

	// 当前开启N4 tcp server
	errN4 := server.StartN4Server(smfCtxt)
	if errN4 != nil {
		panic("Failed to start n4 socket server ")
	}

	smfUeCxt := &smfcontext.UeContext{}
	psi := nas.PduSessID(1)
	smfUeCxt.IMSI.StoreWithNAI("1234565678@cn", 8)

	smfUeCxt.PduSessCtxts = make(map[nas.PduSessID]*smfcontext.PduSessContext, 0)

	//pduSessCtxt := smfcontext.PduSessContext{}
	pduSessCtxt := smfcontext.NewPduSessContext(psi)
	smfUeCxt.PduSessCtxts[psi] = pduSessCtxt

	pduSessCtxt.PDR = configure.SmfConf.N4Conf.PDR
	pduSessCtxt.FAR = configure.SmfConf.N4Conf.FAR

	pduSessCtxt.UEIP = net.ParseIP("192.0.2.20")
	// UPInactiveTimer
	pduSessCtxt.UPInactiveTimer = configure.SmfConf.N4Conf.UPInactiveTimer

	//pduSessCtxt.PDR[0].PDI.UEIPaddress = &pfcp.IEUEIPaddress{}
	//create a pdu session establishment prcd ctxt
	prcdCtxt := prcdctxt.NewPduSessEstbPrcdCtxt(psi)
	pduSessCtxt.SetPrcdCtxt(prcdCtxt)

	//store the information in procedure ctxt
	//prcdCtxt.IMSI = ueCtx
	prcdCtxt.PduSessId = pduSessCtxt.PduSessionId
	//prcdCtxt.Pti = p.PduSessEstbRequest.Pti
	//prcdCtxt.PduSessType = p.PduSessEstbRequest.SessionType
	//prcdCtxt.SscMode = p.PduSessEstbRequest.SscMode
	//prcdCtxt.SmCap = p.PduSessEstbRequest.SmCapability
	//prcdCtxt.MaxNumofSupPktFilter = p.PduSessEstbRequest.MaxNumberOfSPF
	//prcdCtxt.AlwaysOnPduSessReq = p.PduSessEstbRequest.AlwaysOn
	//prcdCtxt.SessionReqIeFlags = p.PduSessEstbRequest.IeFlags

	err = n4layer.N4SessionEstablishmentRequest(pduSessCtxt)
	if err != nil {
		t.Errorf("Failed to create message：%s", err)
	}
	time.Sleep(time.Second * 1)

	for _, far := range pduSessCtxt.FAR {
		far.ForwardingParameters.OuterHeaderCreation = &pfcp.IEOuterHeaderCreation{TEID: 0, IPv4Addr: []byte{192, 0, 2, 20}}
	}
	err = n4layer.N4SessionModificationRequest(pduSessCtxt.SEID, pduSessCtxt.FAR)
	if err != nil {
		t.Errorf("Failed to modify message：%s", err)
	}
	/*seid := smfUeCxt.PduSessCtxts[psi].SEID
	err = n4layer.N4SessionReleaseRequest(seid)
	if err != nil {
		t.Errorf("Failed to release message：%s", err)
	}*/
	smfCtxt.Wg.Wait()
}
func TestN4SessionModificationRequest(t *testing.T) {
	fmt.Println("SMF Process Stared.")
	smfCtxt := &types.AppContext{
		Name:     "smf",
		Wg:       &sync.WaitGroup{},
		ConfPath: types.DefConfFileSmf, //DEF_CONF_FILE,
	}
	smfCtxt.Ctx, smfCtxt.Cancel = context.WithCancel(context.Background())
	configure.LoadConfigCmn("ini", types.DefConfFileAmf)

	//load config
	var confPath = types.DefConfFileSmf
	//F:\work\go\src\5gc\oam\cm\configure\files\smf.config
	//          "src/5gc/oam/cm/configure/files/smf.config"
	err := configure.LoadConfigSmf("ini", confPath)
	if err != nil {
		panic("load config file failed")
	}
	////tag
	//start logger
	//logger.Init(configure.CmnConf.Logger.LogPath, configure.CmnConf.Logger.LogLevel)
	logger.Init()
	if configure.CmnConf.Logger.LogFile {
		logger.SetLoggerToFile(configure.CmnConf.Logger.LogPath, configure.CmnConf.Logger.LogLevel)
	}
	if configure.CmnConf.Logger.LogConsole {
		logger.SetLoggerToConsole(configure.CmnConf.Logger.LogLevel)
	}
	//id manager register ids
	idmgr.GetInst().RegisterIDMgr(string(types.DPE), types.MaxDpeInst)

	n4 := server.NewN4Layer()
	n4.Initialize(smfCtxt)
	n4.Start(smfCtxt)
	smfUeCxt := &smfcontext.UeContext{}
	psi := nas.PduSessID(1)
	smfUeCxt.IMSI.StoreWithNAI("1234565678@cn", 8)

	smfUeCxt.PduSessCtxts = make(map[nas.PduSessID]*smfcontext.PduSessContext, 0)

	pduSessCtxt := smfcontext.PduSessContext{}
	smfUeCxt.PduSessCtxts[psi] = &pduSessCtxt

	pduSessCtxt.PDR = configure.SmfConf.N4Conf.PDR
	pduSessCtxt.FAR = configure.SmfConf.N4Conf.FAR

	pduSessCtxt.UEIP = net.ParseIP("192.0.2.20")
	// UPInactiveTimer
	pduSessCtxt.UPInactiveTimer = configure.SmfConf.N4Conf.UPInactiveTimer

	err = n4layer.N4SessionEstablishmentRequest(&pduSessCtxt)
	if err != nil {
		t.Errorf("Failed to create message：%s", err)
	}
	time.Sleep(time.Second * 1)

	for _, far := range pduSessCtxt.FAR {
		far.ForwardingParameters.OuterHeaderCreation = &pfcp.IEOuterHeaderCreation{TEID: 0, IPv4Addr: []byte{192, 0, 2, 20}}
	}
	err = n4layer.N4SessionModificationRequest(pduSessCtxt.SEID, pduSessCtxt.FAR)
	if err != nil {
		t.Errorf("Failed to modify message：%s", err)
	}
	time.Sleep(time.Second * 5)
	// 去激活N3
	pduSessCtxt.FAR[0].ApplyAction.Set(12)
	pduSessCtxt.FAR[1].ApplyAction.Set(12)

	err = n4layer.N4SessionModificationRequest(pduSessCtxt.SEID, pduSessCtxt.FAR)
	if err != nil {
		t.Errorf("Failed to create message：%s", err)
	}
}
