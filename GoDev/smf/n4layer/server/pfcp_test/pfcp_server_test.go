package test_test

import (
	"context"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/router"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/utils"
	"lite5gc/smf/n4layer/server"
	"lite5gc/smf/n4layer/server/pfcphandler"
	"lite5gc/smf/n4layer/session/sessionhandler"
	"lite5gc/smf/n4layer/typedef"
	"net"
	"runtime"
	"strings"
	"time"

	//"lite5gc/smf/sc/sessmgnt"
	//"lite5gc/smf/sc/subsctrl"
	"fmt"
	"sync"
	"testing"
)

//	todo：引用ngap 的模块是smsender.SendN1N2MsgTransfer(n4Cxt, n.ScCxt)
var smf_conf_file string
var amf_conf_file string

func TestN4SessionProcedure(t *testing.T) {
	//	创建会话流程
	smf_conf_file = types.DefConfFileSmf
	amf_conf_file = types.DefConfFileAmf
	// 创建输入参数
	fmt.Println("smf process stared.")
	smfCtxt := &types.AppContext{
		Name:     "smf",
		Wg:       &sync.WaitGroup{},
		ConfPath: types.DefConfFileSmf, //DEF_CONF_FILE,
	}
	smfCtxt.Ctx, smfCtxt.Cancel = context.WithCancel(context.Background())
	//initial router manager
	routerCtrlChan, routerPubChann := router.Start(smfCtxt)
	smfCtxt.Ctx = context.WithValue(smfCtxt.Ctx, types.RouterCtrlChanCK, routerCtrlChan)
	smfCtxt.Ctx = context.WithValue(smfCtxt.Ctx, types.RouterPublishChanCK, routerPubChann)
	//load amf configuration
	err := configure.LoadConfigCmn("ini", amf_conf_file)
	if err != nil {
		panic("load common config file failed")
	}
	//load smf configuration
	err = configure.LoadConfigSmf("ini", smf_conf_file)
	if err != nil {
		panic("load smf config file failed")
	}
	//start logger
	logger.Init()
	logger.SetLoggerToConsole(configure.CmnConf.Logger.LogLevel)
	//go SmfWrapperStart(smfCtxt)
	fmt.Println("smf is running...", smfCtxt)
	fmt.Println(configure.SmfConf.N4Conf.PFCPParameters)
	// 启动 n4 server
	SmfWrapperStart(smfCtxt)
	// node中需要使用server.conn，等待server创建完成后，添加到节点中
	time.Sleep(time.Second) // 流程太快

	ScId := uint32(1)
	cxt := context.WithValue(smfCtxt.Ctx, types.SessSmfScCK, ScId)
	ScId, ok := cxt.Value(types.SessSmfScCK).(uint32)
	if !ok {
		fmt.Printf("---no sc id found 0")
	} else {
		fmt.Printf("---sc id found 0:%d", ScId)
	}

	//1、从sc中收到初始Establish session消息

	{
		//构造消息:PFCP_Session_Establishment_Request
		var msg pfcpv1.ServiceMsg
		msg.RemoteAdd = &net.UDPAddr{}
		msg.RemoteAdd.IP = configure.SmfConf.N4Conf.UPFIP
		msg.RemoteAdd.Port = configure.SmfConf.N4Conf.UPFPort
		server.HandleRequestMsg(cxt, msg)
		time.Sleep(time.Second * 30)
		//	session 模块
		//n := &typedef.Node{}
		//msg := pfcpv1.ServiceMsg{}
		//sessionhandler.SmfSessionHandle(n.Node, msg)
	}

	//2、向sc通道中发送更新消息

	//3、向sc通道中发送删除消息

}

func TestN4SessionProcedureExample(t *testing.T) {
	//	创建会话流程
	smf_conf_file = types.DefConfFileSmf
	// 创建输入参数
	fmt.Println("smf process stared.")
	smfCtxt := &types.AppContext{
		Name:     "smf",
		Wg:       &sync.WaitGroup{},
		ConfPath: types.DefConfFileSmf, //DEF_CONF_FILE,
	}
	//load smf configuration
	err := configure.LoadConfigSmf("ini", smf_conf_file)
	if err != nil {
		panic("load smf config file failed")
	}
	//go SmfWrapperStart(smfCtxt)
	fmt.Println("smf is running...", smfCtxt)
	fmt.Println(configure.SmfConf.N4Conf.PFCPParameters)
	//1、向sc通道中发送创建消息
	{
		//n := &typedef.Node{}
		//nodehandler.AssociationSetupProcedure(n)
		//	node 没有引用amf/ngap
	}
	{
		//	session 模块
		n := &typedef.Node{}
		msg := pfcpv1.ServiceMsg{}
		sessionhandler.SmfSessionHandle(n.Node, msg)
	}

	{
		//	server 模块
		var peerIp string
		var msg pfcpv1.Message
		var res *pfcpv1.Message
		pfcphandler.NodeDispatch(peerIp, msg, res)

		//	todo：引用ngap 的模块是smsender.SendN1N2MsgTransfer(n4Cxt, n.ScCxt)
	}
	//2、向sc通道中发送更新消息

	//3、向sc通道中发送删除消息

}

// 创建节点流程：创建、更新、释放

// 心跳检测流程：创建、心跳、心跳超时、释放

// 消息重传流程：节点创建、重置对端IP，使不可达、重传消息、释放节点

//const MODULE_ID = logger.PACKAGE_SMFWRAPPER_MODULE_ID

func SmfWrapperStart(appContext *types.AppContext) {

	// n4 layer
	n4 := server.NewN4Layer()
	err := n4.Initialize(appContext)
	if err != nil {
		return
	}

	/*// less then 255
	scNum := configure.SmfConf.SysConf.NumofScInst
	for i := 0; i < scNum; i++ {
		id, err := idmgr.GetInst().BorrowID(string(types.SmfSc))
		if err != nil {
			return
		}

		scCtrl, err := subsctrl.NewSubsCtrl(uint32(id), appContext)
		if err != nil {
			panic(err)
		}

		err = scCtrl.SubsCtrlInit(appContext)
		if err != nil {
			panic(err)
		}

		go scCtrl.Start(appContext)
		// todo zoujun
		SessMgmt, ok := appContext.Ctx.Value(types.SessMgntCK).(*sessmgnt.SessMgmt)
		if !ok {
			fmt.Println("---no sc id found set")
		} else {
			fmt.Printf("---sc id found:%d set", SessMgmt.GetScId())
		}
		time.Sleep(time.Second) // todo sc id cxt store
		n4.ScId = id
	}*/
	n4.ScId = 1
	err = n4.Start(appContext)
	if err != nil {
		panic(err)
	}
}

func TestGetgoid(t *testing.T) {
	goid := utils.Goid()
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println(string(buf[:n]))
	fmt.Println(strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine ")))
	fmt.Println(goid)
}
