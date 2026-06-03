package main

import (
	"context"
	"flag"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

const MODULE_ID = rlogger.PACKAGE_ROUTERAGENT_TEST_MAIN_MODULE_ID

var cpuprofile = flag.String("p", "", "write cpu profile to file")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	//fm, err := os.OpenFile("./mem.out", os.O_RDWR|os.O_CREATE, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer func() {
	//	pprof.WriteHeapProfile(fm)
	//	fm.Close()
	//}()

	amfCtxt := &types.AppContext{
		Name:     "amf",
		Wg:       &sync.WaitGroup{},
		ConfPath: types.DEF_CONF_FILE,
	}
	amfCtxt.Ctx, amfCtxt.Cancel = context.WithCancel(context.Background())

	//load config
	var confPath = types.DEF_CONF_FILE
	err := configure.LoadConfig("ini", confPath)
	if err != nil {
		panic("load config file failed")
	}

	//start logger
	rlogger.Init(configure.CmnConf.rlogger.LogPath, configure.CmnConf.rlogger.LogLevel)

	// start router
	routerCtrlChan, routerPubChann := router.Start(amfCtxt)
	amfCtxt.Ctx = context.WithValue(amfCtxt.Ctx, types.RouterCtrlChanCK, routerCtrlChan)
	amfCtxt.Ctx = context.WithValue(amfCtxt.Ctx, types.RouterPublishChanCK, routerPubChann)

	// start sc goroutine
	ScStart(amfCtxt)
	//time.Sleep(time.Second * 5)
	// start gnb goroutine
	GnbStart(amfCtxt)

	//go testForTimer()

	fmt.Println("amf is running...")

	//main goroutine waiting
	time.Sleep(time.Second * 10)
	//routeManager.DisplayRouteTable()

	//amfCtxt.Wg.Wait()
	go func() {
		t2 := time.NewTicker(time.Second)
		for {
			select {
			case <-t2.C:
				router.ShowRouteTable()
			}
		}
	}()

	http.HandleFunc("/", handler)
	http.ListenAndServe(":11200", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	p := pprof.Lookup("goroutine")
	p.WriteTo(w, 1)
}

func GnbStart(appContext *types.AppContext) {
	rlogger.FuncEntry(types.ModCmn, nil)

	for i := 0; i < 1; i++ {
		appContext.Wg.Add(1)
		go GnbRun(appContext, i)
	}
}

func GnbRun(appContext *types.AppContext, i int) {
	rlogger.FuncEntry(types.ModCmn, nil)

	defer func() {
		appContext.Ctx.Done()
	}()

	ctxt := appContext.Ctx

	routerChan, ok := ctxt.Value(types.RouterPublishChanCK).(router.DataChannel)
	if !ok {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to get router publish channel.")
		return
	}
	routerCtrlChan, ok := ctxt.Value(types.RouterCtrlChanCK).(router.CtrlChannel)
	if !ok {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return
	}

	myEndpoint := router.RouteAddr{
		Type: router.GnbGR,
		Id:   uint32(i),
	}
	msgRouter := routeragent.NewMsgRouter(routerCtrlChan, routerChan, myEndpoint)

	msgRouter.RegisterHandler(router.ScGR, handleMsgSc2Gnb)

	ctxt = context.WithValue(ctxt, types.MsgRouterCK, msgRouter)

	msgRouter.Activate()

	go msgRouter.LoopStart(ctxt)

	go gnbTriggerMsg(ctxt)
}

func handleMsgSc2Gnb(ctx context.Context, msg *router.DataMsg) error {
	rlogger.FuncEntry(types.ModCmn, nil)

	data := msg.MsgData.(Gnb2ScMsg)
	fmt.Printf("GNB goroutine, receive msg from SC: %s  %d\n", data.Name, data.Age)
	//payload := Gnb2ScMsg{Name: "hello from sc", Age: 1, Id: 1}
	//
	//sendMsgBuf := router.DataMsg{
	//	DestAddr: router.RouteAddr{
	//		Type: router.ScGR,
	//		Id:   uint32(msg.SrcAddr.Id),
	//	},
	//	MsgData: payload,
	//}
	//
	//msgRouter, ok := ctx.Value(types.MsgRouterCK).(*routeragent.MsgRouter))
	//if !ok {
	//	rlogger.Trace(types.ModCmn, rlogger.ERROR, nil,  "failed to get router ctrl publish channel.")
	//	return nil
	//}
	//msgRouter.SendMessage(sendMsgBuf)
	return nil
}
func gnbTriggerMsg(ctx context.Context) {
	rlogger.FuncEntry(types.ModCmn, nil)
	msgRouter, ok := ctx.Value(types.MsgRouterCK).(*routeragent.MsgRouter)
	if !ok {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return
	}

	t1 := time.NewTicker(10 * time.Millisecond)
	t2 := time.NewTicker(time.Second)
	for {
		select {
		case <-t1.C:
			payload := Gnb2ScMsg{Name: "hello from gnb", Age: 1, Id: 1}

			sendMsgBuf := &router.DataMsg{
				DestAddr: router.RouteAddr{
					Type: router.ScGR,
					Id:   router.UnknownId,
				},
				MsgData: payload,
			}
			msgRouter.SendMessage(sendMsgBuf)
			//tx_counter ++
		case <-t2.C:
			//fmt.Printf("GNB %d kpps rx %10d tx %10d \n",
			//	i, rx_counter-rx_old_counter, tx_counter-tx_old_coutner)
			//rx_old_counter = rx_counter
			//tx_old_coutner = tx_counter
		}
	}
}

type Gnb2ScMsg struct {
	Name string
	Age  int
	Id   int
}

func (p Gnb2ScMsg) IpcMsgDataIf() {}

func ScStart(appContext *types.AppContext) {
	rlogger.FuncEntry(types.ModCmn, nil)

	for i := 0; i < 1; i++ {
		appContext.Wg.Add(1)
		go ScRun(appContext, i)
	}
}

func ScRun(appContext *types.AppContext, i int) {
	rlogger.FuncEntry(types.ModCmn, nil)

	defer func() {
		appContext.Ctx.Done()
	}()

	ctxt := appContext.Ctx

	routerChan, ok := ctxt.Value(types.RouterPublishChanCK).(router.DataChannel)
	if !ok {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to get router publish channel.")
		return
	}
	routerCtrlChan, ok := ctxt.Value(types.RouterCtrlChanCK).(router.CtrlChannel)
	if !ok {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return
	}

	myEndpoint := router.RouteAddr{
		Type: router.ScGR,
		Id:   uint32(i),
	}
	msgRouter := routeragent.NewMsgRouter(routerCtrlChan, routerChan, myEndpoint)

	msgRouter.RegisterHandler(router.GnbGR, handleMsgGnb2Sc)

	ctxt = context.WithValue(ctxt, types.MsgRouterCK, msgRouter)

	msgRouter.Activate()

	msgRouter.LoopStart(ctxt)
}

func handleMsgGnb2Sc(ctx context.Context, msg *router.DataMsg) error {
	rlogger.FuncEntry(types.ModCmn, nil)

	//data := msg.MsgData.(Gnb2ScMsg)

	payload := Gnb2ScMsg{Name: "hello from sc", Age: 1, Id: 1}

	sendMsgBuf := &router.DataMsg{
		DestAddr: router.RouteAddr{
			Type: router.GnbGR,
			Id:   uint32(msg.SrcAddr.Id),
		},
		MsgData: payload,
	}

	msgRouter, ok := ctx.Value(types.MsgRouterCK).(*routeragent.MsgRouter)
	if !ok {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return nil
	}
	msgRouter.SendMessage(sendMsgBuf)
	return nil
}
