package main

import (
	"context"
	"flag"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/utils"
	"lite5gc/oam/cm"
	"lite5gc/udm/dbmgr"
	"lite5gc/udm/sbi/addsbirouters"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var signalsHdl = []os.Signal{os.Interrupt, syscall.SIGTERM}

var udm_conf_file string
var sys_conf_file string

func init() {
	rlogger.SetLogOutput(os.Stdout)
	rlogger.SetLogLevel(rlogger.TRACE)

	flag.StringVar(&udm_conf_file, "config", types.DefConfFileUdm, "udm configuration file")
	flag.StringVar(&sys_conf_file, "sysconf", types.DefConfFileSys, "system configuration file")
}

func main() {

	// load the configuration
	cm.LoadSysConfig(sys_conf_file)
	cm.LoadUdmConfig(udm_conf_file)

	//fmt.Println(configure.UdmConf)
	//fmt.Println(configure.SysConf)

	// logger initialize with configuration
	rlogger.Initialize(utils.Convert2Bool(configure.UdmConf.Logger.Control),
		configure.UdmConf.Logger.Path,
		configure.UdmConf.Logger.Level)

	// create the context
	udmCtxt := &types.AppContext{
		Name: "udm",
		Wg:   &sync.WaitGroup{},
	}
	udmCtxt.Ctx, udmCtxt.Cancel = context.WithCancel(context.Background())

	//init statistics
	// statistics.Init()

	//start agent listen
	//oamagent.StartOamAgent(udmCtxt.Ctx)

	// database initialize db manager
	dbmgr.DbInit()

	// sbi
	addsbirouters.Start(udmCtxt)

	//interrupt Signal Handler
	SignalHandler(udmCtxt)

	// Quit Procedure:
	//oamagent.StopOamAgent()
	fmt.Println("udm main exit complete")
}

func SignalHandler(appContext *types.AppContext) {

	//regist signal handler
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, signalsHdl...)

	for {
		select {
		case sig := <-sigchan:
			fmt.Printf("receive signal:%s, UDM exits the main program \n", sig)
			rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "receive signal (%s).", sig)
			rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil, "udm main exit complete")
			appContext.Cancel()
			return
		}
	}
}
