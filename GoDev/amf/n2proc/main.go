package main

import (
	"context"
	"flag"
	"fmt"
	"lite5gc/amf/n2proc/n2sctp"
	"lite5gc/cmn/rlogger"
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
	flag.StringVar(&amf_conf_file, "config", types.DefConfFileAmf, "amf configuration file")
	flag.StringVar(&sys_conf_file, "sysconf", types.DefConfFileSys, "system configuration file")
}

func main() {
	flag.Parse()

	//debug for pprof
	//go func() {
	//	ip := "0.0.0.0:6060"
	//	if err := http.ListenAndServe(ip, nil); err != nil {
	//		fmt.Printf("start pprof failed on %s\n", ip)
	//		os.Exit(1)
	//	}
	//}()

	// load the configuration
	cm.LoadSysConfig(sys_conf_file)
	err := cm.LoadAmfConfig(amf_conf_file)
	if err != nil {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "fail to load config")
		return
	}

	// logger initialize with configuration
	rlogger.Initialize(utils.Convert2Bool(configure.AmfConf.Logger.Control),
		configure.AmfConf.Logger.Path,
		configure.AmfConf.Logger.Level)

	// create the context
	amfN2Ctxt := &types.AppContext{
		Name: "amf",
		Wg:   &sync.WaitGroup{},
	}
	amfN2Ctxt.Ctx, amfN2Ctxt.Cancel = context.WithCancel(context.Background())

	// star the main loop
	n2sctp.Start(amfN2Ctxt)
	fmt.Println("amf n2 process is running...")

	// handle the signals
	go listenSignal(amfN2Ctxt)

	amfN2Ctxt.Wg.Wait()

	fmt.Println("amf n2 process exit")
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
