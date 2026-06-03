package main

import (
	"context"
	"lite5gc/amf/n2proc/n2sctp"
	"lite5gc/cmn/rlogger"
	"sync"

	"lite5gc/cmn/types"
)

func main() {

	rlogger.Initialize(true,
		"./sctp_ser.log",
		"debug")

	rlogger.Trace("sctp_ser", rlogger.DEBUG, nil, "logger initailzed success!")

	amfN2Ctxt := &types.AppContext{
		Name: "amf",
		Wg:   &sync.WaitGroup{},
	}
	amfN2Ctxt.Ctx, amfN2Ctxt.Cancel = context.WithCancel(context.Background())

	// star the main loop
	n2sctp.Start(amfN2Ctxt)

	amfN2Ctxt.Wg.Wait()

}
