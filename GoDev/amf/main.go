// Main entry for amf process
package main

import (
	"context"
	"flag"
	"fmt"
	"lite5gc/amf/idmgramf"
	"lite5gc/amf/sbiamf"
	"lite5gc/amf/sbiamf/sbiinit"
	"lite5gc/amf/sc/statistics"
	"lite5gc/amf/sc/subsctrl"
	"lite5gc/cmn/redisclt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/utils"
	"lite5gc/oam/cm"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var amf_conf_file string
var sys_conf_file string

func init() {
	//rlogger.SetLogOutput("logs/amf.log")
	//rlogger.SetLogOutput(os.Stdout)
	//rlogger.SetLogLevel(rlogger.TRACE)

	flag.StringVar(&amf_conf_file, "config", types.DefConfFileAmf, "amf configuration file")
	flag.StringVar(&sys_conf_file, "sysconf", types.DefConfFileSys, "system configuration file")
}

func main() {
	flag.Parse()

	// load the configuration
	cm.LoadSysConfig(sys_conf_file)
	err := cm.LoadAmfConfig(amf_conf_file)
	if err != nil {
		rlogger.Trace(types.ModuleAmf, rlogger.ERROR, nil, "fail to load config")
		return
	}

	// logger initialize with configuration
	rlogger.Initialize(utils.Convert2Bool(configure.AmfConf.Logger.Control),
		configure.AmfConf.Logger.Path,
		configure.AmfConf.Logger.Level)

	// initialize the id manager
	err = idmgramf.AmfIdMgrInit()
	if err != nil {
		rlogger.Trace(types.ModuleAmf, rlogger.ERROR, nil, "fail to init id mgr")
		return
	}

	// create the context
	amfCtxt := &types.AppContext{
		Name: "amf",
		Wg:   &sync.WaitGroup{},
	}
	amfCtxt.Ctx, amfCtxt.Cancel = context.WithCancel(context.Background())

	//initial router manager
	routerCtrlChan, routerPubChan := router.Start(amfCtxt)
	amfCtxt.Ctx = context.WithValue(amfCtxt.Ctx, types.RouterCtrlChanCK, routerCtrlChan)
	amfCtxt.Ctx = context.WithValue(amfCtxt.Ctx, types.RouterPublishChanCK, routerPubChan)

	// initialize redis client
	err = redisclt.RedisCltInit()
	if err != nil {
		rlogger.Trace(types.ModuleAmf, rlogger.ERROR, nil,
			"fail to connect to redis server %s",
			fmt.Sprintf("%s:%d", configure.SysConf.RedisAddr.Ip, configure.SysConf.RedisAddr.Port))
		return
	}

	// initialize statistics
	statistics.Init()

	// start subscriber control module
	subsctrl.Start(amfCtxt)

	// start sbi
	sbiinit.Initialize()
	sbiamf.Start(amfCtxt)

	rlogger.Trace(types.ModuleAmf, rlogger.INFO, nil, "amf is running...")
	fmt.Println("amf start running")

	// handle the signals
	go listenSignal(amfCtxt)

	amfCtxt.Wg.Wait()

	fmt.Println("amf main exit complete")
}

func listenSignal(appContext *types.AppContext) {
	sigs := make(chan os.Signal, 5)
	signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1,
		syscall.SIGUSR2, syscall.SIGTSTP)
	select {
	case sig := <-sigs:
		fmt.Println("exit process, sigs: ", sig)
		appContext.Cancel()
		return
	}
}
