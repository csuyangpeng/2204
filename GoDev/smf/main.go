package main

import (
	"context"
	"flag"
	"fmt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/utils"
	"lite5gc/oam/cm"
	"lite5gc/smf/idmgrsmf"
	"lite5gc/smf/n4layer/server"
	"lite5gc/smf/sbismf"
	"lite5gc/smf/sbismf/sbiInit"
	"lite5gc/smf/sc/subsctrl"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var signalsHdl = []os.Signal{os.Interrupt, syscall.SIGTERM}
var smf_conf_file string

func init() {
	flag.StringVar(&smf_conf_file, "config", types.DefConfFileSmf, "smf configuration file")
}

func main() {
	flag.Parse()

	// config
	err := cm.LoadSmfConfig(smf_conf_file)
	if err != nil {
		rlogger.Trace(types.ModuleAmf, rlogger.ERROR, nil, "fail to load config")
		return
	}

	// logger initialize with configuration
	rlogger.Initialize(utils.Convert2Bool(configure.SmfConf.Logger.Control),
		configure.SmfConf.Logger.Path,
		configure.SmfConf.Logger.Level)

	// id mgr
	err = idmgrsmf.SmfIdMgrInit()
	if err != nil {
		rlogger.Trace(types.ModuleSmf,rlogger.ERROR,nil,"fail to init id mgr")
		return
	}

	//register signal handler
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, signalsHdl...)

	smfCtxt := &types.AppContext{
		Name: "smf",
		Wg:   &sync.WaitGroup{},
	}
	smfCtxt.Ctx, smfCtxt.Cancel = context.WithCancel(context.Background())

	//initial router manager
	routerCtrlChan, routerPubChann := router.Start(smfCtxt)
	smfCtxt.Ctx = context.WithValue(smfCtxt.Ctx, types.RouterCtrlChanCK, routerCtrlChan)
	smfCtxt.Ctx = context.WithValue(smfCtxt.Ctx, types.RouterPublishChanCK, routerPubChann)

	//init statistics
	statistics.Init()

	// subscriber control
	scNum := 1 //configure.SysConf.SC.NumOfSmfScInst
	rlogger.Trace(types.ModuleSmf, rlogger.INFO, nil, "SC config reader success, scNum(%d)", scNum)
	for i := 0; i < scNum; i++ {
		id, err := idmgr.GetInst().BorrowID(string(types.SmfSc))
		if err != nil {
			rlogger.Trace(types.ModuleSmf, rlogger.FATAL, nil, "Failed to Borrow Id for Smf SC, err:", err)
			return
		}
		rlogger.Trace(types.ModuleSmf, rlogger.WARN, nil, "Smf Sc Instance (%d) created.", id)

		scCtrl, err := subsctrl.NewSubsCtrl(id, smfCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleSmf, rlogger.FATAL, nil, "failed to create "+
				"Subscriber Control Module, err:", err)
			panic(err)
		}

		err = scCtrl.SubsCtrlInit(smfCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleSmf, rlogger.FATAL, nil, "failed to initialize "+
				"Subscriber Control Module, err:", err)
			panic(err)
		}

		go scCtrl.Start(smfCtxt)
	}

	// n4 layer
	n4 := server.NewN4Layer()
	err = n4.Initialize(smfCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "smf n4 layer initialize failed.")
		return
	}
	err = n4.Start(smfCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleSmf, rlogger.FATAL, nil, "failed to n4 start "+
			"err:", err)
		panic(err)
	}

	sbiInit.Initialize()
	sbismf.Start(smfCtxt)

	go SignalHandler(sigChan, smfCtxt)

	//main goroutine waiting group protect (delay 1 second to ensure wait group start work):

	time.Sleep(1 * time.Second)

	smfCtxt.Wg.Wait()

	fmt.Println("smf main exit complete")
}

func SignalHandler(signalchan chan os.Signal, appContext *types.AppContext) {
	rlogger.FuncEntry(types.ModuleSmf, nil)
	for {
		select {
		case sig := <-signalchan:
			rlogger.Trace(types.ModuleSmf, rlogger.ERROR, nil, "receive signal (%s).", sig)
			fmt.Printf("receive signal:%s\n", sig)
			appContext.Cancel()

			return
		}
	}
}
