package main

import (
	"context"
	"flag"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
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
	gnbDataChanns := []router.DataChannel{}

	for i := 0; i < 2000; i++ {
		dchan := GnbRegist(appContext, i)
		gnbDataChanns = append(gnbDataChanns, dchan)
	}

	time.Sleep(time.Second * 5)

	for index, dchan := range gnbDataChanns {
		appContext.Wg.Add(1)
		go GnbRun(appContext, index, dchan)
	}
}

func GnbRegist(appContext *types.AppContext, i int) router.DataChannel {
	rlogger.FuncEntry(types.ModCmn, nil)

	defer func() {
		appContext.Ctx.Done()
	}()

	ctxt := appContext.Ctx

	routerCtrlChan, ok := ctxt.Value(types.RouterCtrlChanCK).(router.CtrlChannel)
	if !ok {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "failed to get router ctrl publish channel.")
		return nil
	}
	gnbChann := make(router.DataChannel, 500)

	//register
	msg := &router.ControlMsg{}
	msg.SrcAddr = router.RouteAddr{
		Type: router.GnbGR,
		Id:   uint32(i),
	}
	msg.Op = router.Register
	msg.PubChannel = gnbChann

	routerCtrlChan <- msg

	return gnbChann
}

func GnbRun(appContext *types.AppContext, i int, gnbChann router.DataChannel) {
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

	//send data ScGR_0
	//payload := Gnb2ScMsg{Name: "hello from gnb", Age: 18}
	//msgBuf := router.DataMsg{
	//	DestAddr: router.RouteAddr{
	//		Type: router.ScGR,
	//		Id:   uint32(rand.Intn(9)),
	//	},
	//	MsgData: payload,
	//}
	//SendMsg := &router.IpcMsg{
	//	MsgT: router.DP,
	//	MsgD: msgBuf,
	//}

	var rx_counter int
	var rx_old_counter int
	var tx_counter int
	var tx_old_coutner int
	//t2 := time.After(time.Second * 2)
	//t1 := time.NewTicker(time.Microsecond * 20)
	t1 := time.NewTicker(10 * time.Millisecond)
	//t1 := time.NewTicker(time.Millisecond)
	t2 := time.NewTicker(time.Second)
	for {
		select {
		case msgBuf := <-gnbChann:
			data := msgBuf.MsgD.(router.DataMsg).MsgData.(Gnb2ScMsg)
			//fmt.Printf("GNB goroutine (%d), receive msg from SC: %s  %d\n", i, data.Name, data.Age)
			rx_counter = rx_counter + data.Age
		case <-t1.C:
			payload := Gnb2ScMsg{Name: "hello from gnb", Age: 1, Id: i}
			msgBuf := router.DataMsg{
				DestAddr: router.RouteAddr{
					Type: router.ScGR,
					//Id:   router.UnknownId, //uint32(rand.Intn(20)),
					//Id: router.BroadcastId, //uint32(rand.Intn(20)),
					Id: uint32(rand.Intn(20)),
					//Id: uint32(0),
				},
				MsgData: payload,
			}
			SendMsg := &router.IpcMsg{
				MsgT: router.DP,
				MsgD: msgBuf,
			}

			//sendTime := time.Now()
			//go func() { routerChan <- SendMsg }()

			routerChan <- SendMsg

			//elapsed := time.Since(sendTime)
			//if elapsed > time.Microsecond*20000 {
			//	fmt.Println("cost time: ", elapsed)
			//}

			tx_counter++
		case <-t2.C:
			//fmt.Printf("GNB %d kpps rx %10d tx %10d \n",
			//	i, rx_counter-rx_old_counter, tx_counter-tx_old_coutner)
			rx_old_counter = rx_counter
			tx_old_coutner = tx_counter
		}
	}
	println(rx_old_counter, tx_old_coutner)
}

type Gnb2ScMsg struct {
	Name string
	Age  int
	Id   int
}

func (p Gnb2ScMsg) IpcMsgDataIf() {}

func ScStart(appContext *types.AppContext) {
	rlogger.FuncEntry(types.ModCmn, nil)

	for i := 0; i < 20; i++ {
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

	scChann := make(router.DataChannel, 500)

	//register
	msg := &router.ControlMsg{}
	msg.SrcAddr = router.RouteAddr{
		Type: router.ScGR,
		Id:   uint32(i),
	}
	msg.Op = router.Register
	msg.PubChannel = scChann

	routerCtrlChan <- msg

	//TODO need response from router?

	//send data ScGR_0
	//payload := Gnb2ScMsg{Name: "hello from sc", Age: 20}
	//msgBuf := router.DataMsg{
	//	DestAddr: router.RouteAddr{
	//		Type: router.GnbGR,
	//		Id:   uint32(rand.Intn(9)),
	//	},
	//	MsgData: payload,
	//}
	//SendMsg := &router.IpcMsg{
	//	MsgT: router.DP,
	//	MsgD: msgBuf,
	//}

	var rx_counter int
	var rx_old_counter int
	var tx_counter int
	var tx_old_coutner int
	//t2 := time.After(time.Second * 2)
	//t1 := time.NewTicker(time.Microsecond * 50)
	//t1 := time.NewTicker(time.Millisecond)
	t2 := time.NewTicker(time.Second)
	for {
		select {
		case msgBuf := <-scChann:
			data := msgBuf.MsgD.(router.DataMsg).MsgData.(Gnb2ScMsg)
			//fmt.Printf("SC goroutine (%d), receive msg from GNB: %s  %d\n", i, data.Name, data.Age)
			rx_counter = rx_counter + data.Age

			//case <-t1.C:
			payload := Gnb2ScMsg{Name: "hello from sc", Age: 1, Id: i}
			sendMsgBuf := router.DataMsg{
				DestAddr: router.RouteAddr{
					Type: router.GnbGR,
					Id:   uint32(data.Id),
					//Id: uint32(0),
				},
				MsgData: payload,
			}
			SendMsg := &router.IpcMsg{
				MsgT: router.DP,
				MsgD: sendMsgBuf,
			}
			//go func() { routerChan <- SendMsg }()
			routerChan <- SendMsg
			tx_counter++
		case <-t2.C:
			fmt.Printf("SC  %d kpps rx %10d tx %10d \n",
				i, (rx_counter-rx_old_counter)/1000, (tx_counter-tx_old_coutner)/1000)
			rx_old_counter = rx_counter
			tx_old_coutner = tx_counter
		}
	}

}

func testForTimer() {

	var rx_counter int
	var rx_old_counter int
	var tx_counter int
	var tx_old_coutner int
	//t2 := time.After(time.Second * 2)
	//t1 := time.NewTicker(time.Microsecond * 20)
	t1 := time.NewTicker(1 * time.Microsecond)
	//t1 := time.NewTimer(time.Microsecond * 5)
	//t1 := time.NewTicker(time.Millisecond)
	t2 := time.NewTicker(time.Second * 2)
	//t3 := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-t1.C:
			tx_counter++
			//t1.Reset(time.Microsecond * 5)
		case <-t2.C:
			fmt.Printf("Timer kpps rx %10d tx %10d \n",
				(rx_counter-rx_old_counter)/2000, (tx_counter-tx_old_coutner)/2000)
			rx_old_counter = rx_counter
			tx_old_coutner = tx_counter
			//case <-t3.C:
		}

	}
}
