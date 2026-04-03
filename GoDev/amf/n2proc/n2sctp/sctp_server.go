// Package ngsctp implement the sctp layer on AMF,
// terminate sctp layer communication
package n2sctp

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
)

var SctpSer *N2Listener
var err error

// Start start a ng sctp goroutine,
// which will listen and process all sctp related msg and event
func Start(appContext *types.AppContext) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	appWg := appContext.Wg

	//read from config for ip and port
	n2configIP := configure.AmfConf.N2.Ip
	n2configPort := configure.AmfConf.N2.Port

	// read from config for parameter
	sctpPara := &SctpOptions{}
	sctpPara.NumOstreams = uint16(configure.AmfConf.Sctp.Initmsg.NumOstreams)
	sctpPara.MaxInstreams = uint16(configure.AmfConf.Sctp.Initmsg.MaxInstreams)
	sctpPara.MaxAttempts = uint16(configure.AmfConf.Sctp.Initmsg.MaxAttempts)
	sctpPara.MaxInitTimeout = uint16(configure.AmfConf.Sctp.Initmsg.MaxInitTimeout)
	sctpPara.HeatbeatInterval = uint16(configure.AmfConf.Sctp.Heatbeat.Interval)
	sctpPara.PathMaxRXT = uint16(configure.AmfConf.Sctp.Heatbeat.PathMaxRXT)

	SctpSer, err = NewN2Listener(n2configIP, n2configPort, sctpPara, appContext.Ctx)
	if err != nil {
		fmt.Println("Failed to initialize ng sctp server, err:", err)
		panic("ng sctp server initialized failed")
	}

	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, "ng sctp server initialized success!")

	appWg.Add(1)
	go SctpSer.Start(appWg)
}
