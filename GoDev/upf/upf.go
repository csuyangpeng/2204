// UPF implements UPF function module in 5GC
package main

import (
	"context"
	"github.com/intel-go/nff-go/flow"
	"github.com/intel-go/nff-go/packet"
	flowT "github.com/intel-go/nff-go/types"
	"lite5gc/cmn/metric"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/utils"
	"lite5gc/upf/adapter"
	"lite5gc/upf/context/pdrcontext"
	"lite5gc/upf/cp/n4udp"
	"lite5gc/upf/cp/pdr"
	"lite5gc/upf/defs"
	"lite5gc/upf/metrics"
	"lite5gc/upf/up"
	"strings"

	"lite5gc/upf/context/ipport"

	"lite5gc/upf/service"

	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"flag"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	//"lite5gc/cmn/types/configure"
	//"lite5gc/upf/agent/cmdagent"
	//upfcfg "lite5gc/upf/agent/config"
	//_ "lite5gc/upf/utils/tracerlogger"
)

const moduleTag rlogger.ModuleTag = "upf"

//var
var signalsHdl = []os.Signal{os.Interrupt, syscall.SIGTERM}

// init()
var work_dir string = ""
var upfConfFile string = "config/cm_upf_conf.yaml"

func init() {
	//flag.StringVar(&upfConfFile, "c", "upf.log", "upf configuration file")
	flag.StringVar(&work_dir, "workdir", "", "Set upf workdir, where contain config and logs directors")

}

func main() {

	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	//regist signal handler
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, signalsHdl...)

	// upfCtxt
	upfCtxt := &types.AppContext{
		Name:     "upf",
		Wg:       &sync.WaitGroup{},
		ConfPath: upfConfFile, // GOPATH 下的相对路径
	}
	upfCtxt.Ctx, upfCtxt.Cancel = context.WithCancel(context.Background())

	//load upf configuration from config file
	adapter.LoadConfigUPF(upfConfFile)
	fmt.Printf("%+v\n", configure.CmUpfConf)

	//upfcfg.Start(upfCtxt, upfConfFile)

	//start rlogger
	err := rlogger.Initialize(
		utils.Convert2Bool(configure.UpfConf.Logger.Control),
		configure.UpfConf.Logger.Path,
		configure.UpfConf.Logger.Level)
	if err != nil {
		fmt.Println(err)
	}

	rlogger.Trace(moduleTag, rlogger.INFO, nil, "start rlogger.")

	//start agent listen
	// upfAgent.UpfListenAgent(upfCtxt)
	//oamagent.StartOamAgent(upfCtxt.Ctx)

	//id manager register ids
	//idmgr.GetInst().RegisterIDMgr(string(types.DPE), types.MAX_DPE_INST)
	//intialize ETCD TODO:需测试后打开
	//etcdAgent.Start(upfCtxt)

	// packet counter init
	metricMap, _ := metrics.UpfCounterInit()

	// 开启N4 udp server
	// 当前开启N4 tcp server
	// dnn configure load
	err = pdr.StoreDnnGwIpTable(configure.UpfConf.DnnNameGwIpMap)
	if err != nil {
		fmt.Println(err)
	}

	errN4 := n4udp.StartN4Server(upfCtxt) //n4layer.StartN4Server(upfCtxt)
	if errN4 != nil {
		panic("Failed to start n4 socket server ")
	}

	// Performance improvement ,收发包统计计数
	errN7 := StartPerformanceMonitor(upfCtxt, metricMap)
	if errN7 != nil {
		panic("Failed to start performance testing ")
	}
	// Performance improvement，tool pprof
	pprofOptimization(upfCtxt)

	//cmdagent.Start(upfCtxt)

	go SignalHandler(sigchan, upfCtxt)

	//main goroutine waiting group protect (delay 1 second to ensure wait group start work):
	time.Sleep(1 * time.Second)

	fmt.Println("UPF Process Stared.")
	//	todo nff-go
	err = service.StartBufferServer(upfCtxt)
	if err != nil {
		fmt.Println(err)
	}
	startNffGo()
	// 阻塞不退出
	upfCtxt.Wg.Wait()

	// Quit Procedure:
	//oamagent.StopOamAgent()
	fmt.Println("upf main exit complete")

}

func SignalHandler(signalchan chan os.Signal, appContext *types.AppContext) {
	rlogger.FuncEntry(moduleTag, nil)

	for {
		select {
		case sig := <-signalchan:
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "receive signal (%s).", sig)
			fmt.Printf("receive signal:%s\n", sig)
			appContext.Cancel()

			return
		}
	}
}

//todo 增加统计数据
func StartPerformanceMonitor(Cxt *types.AppContext, m metric.Registry) error {
	// 15秒输出收发包信息
	go func(Cxt *types.AppContext) {
		rlogger.Trace(moduleTag, rlogger.INFO, nil, "Performancemonitor routine start")
		Cxt.Wg.Add(1)
		defer Cxt.Wg.Done()

		//定时打印
		t1 := time.NewTicker(time.Second * 15)
		for {
			select {
			case <-Cxt.Ctx.Done():
				rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Performancemonitor routine exit")
				return

			case <-t1.C:
				// 时间区间收包数

				// 重新计数
				defs.N3CountReceivePacket = 0
				//dpe.N3CountReceivePacket2 = 0
				atomic.StoreUint64(&defs.N3CountSendPacket, 0)
				//dpe.N3CountSendPacket = 0
				atomic.StoreUint64(&defs.N6CountReceivePacket, 0)
				//dpe.N6CountReceivePacket = 0
				atomic.StoreUint64(&defs.N6CountSendPacket, 0)
				//dpe.N6CountSendPacket = 0

			}
		}

	}(Cxt)
	return nil
}

func pprofOptimization(upfCtxt *types.AppContext) {
	go func(Cxt *types.AppContext) {
		log.Println(http.ListenAndServe("localhost:11201", nil))
	}(upfCtxt)
}

// log level define
// 正常数据流使用   debug
// 正常n4信令流使用 info
// 可接受异常使用   warn
// 不可接受业务异常 error

//var inport int = "enp2s0f1"
//var outport string = "enp2s0f1"
//var IpPorts []ipport.IpPort
//var N3Outport uint16
//	todo nff-go
func startNffGo() {
	inport := flag.Uint("inport", 0, "Port for receiving packets.")
	outport := flag.Uint("outport", 0, "Port for sending packets.")
	//numflows := flag.Uint("numflows", 2, "Number of output flows to use. First flow with number zero is used for dropped packets.")
	nostats := flag.Bool("nostats", false, "Disable statics HTTP server.")
	//NoHWTXChecksum := flag.Bool("nohwcsum", false, "Specify whether to use hardware offloading for checksums calculation (requires -csum).")
	dpdkLogLevel := flag.String("dpdk", "--log-level=0", "Passes an arbitrary argument to dpdk EAL.")
	sendCPUCoresPerPort := flag.Int("send-threads", 0, "Number of CPU cores to be occupied by Send routines.")
	tXQueuesNumberPerPort := flag.Int("tx-queues", 0, "Number of transmit queues to use on network card.")

	vector := flag.Bool("v", false, "vector version")
	//ipv4Addr := flag.String("ipv4", "10.18.1.67", "ipv4 address")
	ipv4Addr := &configure.UpfConf.N3.Ipv4

	flag.Parse()
	ipport.N3Outport = uint16(*outport)

	var statsServerAddres *net.TCPAddr = nil
	if !*nostats {
		// Set up address for stats web server
		statsServerAddres = &net.TCPAddr{
			Port: 8080,
			IP:   net.IP{10, 18, 1, 53},
		}
	}

	config := flow.Config{
		//CPUList:          "0-3",
		StatsHTTPAddress: statsServerAddres,
		DPDKArgs:         []string{*dpdkLogLevel},
		//NoPacketHeadChange: true,
		//HWTXChecksum: !(*NoHWTXChecksum),
		SendCPUCoresPerPort:   *sendCPUCoresPerPort,
		TXQueuesNumberPerPort: *tXQueuesNumberPerPort,
	}
	//arp
	//大端存
	//ipv4 := flowT.IPv4Address(22)<<24 | flowT.IPv4Address(0)<<16 | flowT.IPv4Address(18)<<8 | flowT.IPv4Address(172)
	// 返回大端序
	ipv4 := flowT.SliceToIPv4(net.ParseIP(*ipv4Addr).To4())
	// init port
	ipport.IpPorts = make([]ipport.IpPort, 2, 2)
	for i, port := range ipport.IpPorts {
		port.Index = uint16(i)
		port.Subnet.IPv4.Addr = ipv4
		port.Subnet.IPv4.Mask = 0x0000ffff
		//port.MacAddress = flow.GetPortMACAddress(uint16(i)) //dpdk 没有初始化，不能获取Mac地址，系统报错：Invalid port_id=0
		port.StaticARP = port.DstMacAddress != flowT.MACAddress{0, 0, 0, 0, 0, 0}
		myV4Checker := func(ipv4 flowT.IPv4Address) bool {
			// ip 检查
			return ipv4 == port.Subnet.IPv4.Addr
		}
		myV6Checker := func(ipv6 flowT.IPv6Address) bool {
			return ipv6 == port.Subnet.IPv6.Addr
		}
		port.NeighCache = packet.NewNeighbourTable(port.Index, port.MacAddress, myV4Checker, myV6Checker)
		ipport.IpPorts[i] = port
		fmt.Printf("%+v\n", port)
	}

	flow.CheckFatal(flow.SystemInit(&config))
	//n3 interface
	portChecker := func() error {
		if *inport < 2 {
			return nil
		}
		return fmt.Errorf("port support 0,1")
	}()
	flow.CheckFatal(portChecker)

	//inputFlow, err := flow.SetReceiver(uint16(*inport))
	inputFlow, err := flow.SetReceiver(ipport.IpPorts[*inport].Index)
	fmt.Println("***---NFF--GO---***", err)
	fmt.Println("***---NFF--GO---inport num=", *inport)
	fmt.Println("***---NFF--GO---inport index=", ipport.IpPorts[*inport].Index)
	//inputFlow, err := flow.SetReceiverOS("enp2s0f1")
	fmt.Println("***---NFF--GO---***")
	flow.CheckFatal(err)
	//flow.CheckFatal(flow.SetHandler(inputFlow, dumpPrintF2, nil))

	flow.CheckFatal(flow.SetIPForPort(uint16(*inport), ipv4))
	fmt.Println(flowT.IPv4ToBytes(ipv4))
	fmt.Println(flow.GetNameByPort(uint16(*inport)))
	fmt.Printf("%#x\n", flow.GetPortMACAddress(uint16(*inport)))
	ipport.IpPorts[*inport].MacAddress = flow.GetPortMACAddress(uint16(*inport))
	myV4Checker := func(ipv4 flowT.IPv4Address) bool {
		// ip 检查
		return ipv4 == ipport.IpPorts[*inport].Subnet.IPv4.Addr
	}
	myV6Checker := func(ipv6 flowT.IPv6Address) bool {
		return ipv6 == ipport.IpPorts[*inport].Subnet.IPv6.Addr
	}
	ipport.IpPorts[*inport].NeighCache = packet.NewNeighbourTable(ipport.IpPorts[*inport].Index, ipport.IpPorts[*inport].MacAddress, myV4Checker, myV6Checker)
	//// 接收arp 请求处理
	flow.CheckFatal(flow.DealARPICMP(inputFlow))

	n3cxt := up.HandlerContext{Msgcxt: &pdrcontext.DataFlowContext{}}
	n3cxt.Ipport = &ipport.IpPorts[*inport]
	fmt.Printf("port %d : %+v\n", *inport, ipport.IpPorts[*inport])
	fmt.Printf("cxt.Ipport %+v\n", n3cxt.Ipport)
	n6cxt := up.HandlerContext{Msgcxt: &pdrcontext.DataFlowContext{}}
	n6cxt.Ipport = &ipport.IpPorts[*inport]
	fmt.Printf("port %d : %+v\n", *inport, ipport.IpPorts[*inport])
	fmt.Printf("cxt.Ipport %+v\n", n6cxt.Ipport)
	//flow.CheckFatal(flow.SetHandler(inputFlow, dumpPrintPortCxt, cxt))
	// 接收arp响应处理
	flow.CheckFatal(flow.SetHandlerDrop(inputFlow, up.ReceiveArpHandler, n3cxt))

	// N3 消息保留
	// N3只接收udp 2152 消息
	secondFlow, err := flow.SetSeparator(inputFlow, up.NffN3Handler, nil)
	flow.CheckFatal(err)

	// N6 消息保留
	//flow.CheckFatal(flow.SetHandler(secondFlow, dumpPrintF2, nil))
	//dropFlow, err := flow.SetSeparator(secondFlow, dpe.NffN6Handler, n6cxt)
	//flow.CheckFatal(err)
	//flow.CheckFatal(flow.SetStopper(dropFlow))

	//gtp 编码
	if *vector {
		flow.CheckFatal(flow.SetVectorHandlerDrop(secondFlow, up.NffStartDPERawHandleVector, n6cxt))
	} else {
		flow.CheckFatal(flow.SetHandlerDrop(secondFlow, up.NffStartDPERawHandle, n6cxt))
	}
	//发送gtp消息到N3
	//flow.CheckFatal(flow.SetHandler(secondFlow, dumpPrint, nil))
	flow.CheckFatal(flow.SetSender(secondFlow, uint16(*outport)))

	// gtp 解码
	//flow.CheckFatal(flow.SetHandler(inputFlow, dpe.NffStartDPEHandle, cxt))
	if *vector {
		flow.CheckFatal(flow.SetVectorHandlerDrop(inputFlow, up.NffStartDPEHandleVectorV1, n3cxt))
	} else {
		flow.CheckFatal(flow.SetHandlerDrop(inputFlow, up.NffStartDPEHandleV1, n3cxt))
	}
	//flow.CheckFatal(flow.SetHandler(inputFlow, dumpPrint, nil))
	//flow.CheckFatal(flow.SetVectorHandler(inputFlow, dpe.MyHandler, nil))

	// 发送 body 到 N6
	//flow.CheckFatal(flow.SetHandlerDrop(inputFlow, dpe.NffN6SendHandler, nil))
	//flow.CheckFatal(flow.SetHandler(inputFlow, dumpPrintPortCxt, cxt))
	//flow.CheckFatal(flow.SetHandlerDrop(inputFlow, dpe.NffN6SendHandlerV1, n3cxt))
	flow.CheckFatal(flow.SetHandler(inputFlow, dumpPrint, nil))

	flow.CheckFatal(flow.SetSender(inputFlow, uint16(*outport)))
	fmt.Println("***---NFF--GO---*** start")
	//flow.CheckFatal(flow.SetSenderOS(inputFlow, outport))

	flow.CheckFatal(flow.SystemStart())

}

var flag1 = 0

func dumpPrint(currentPacket *packet.Packet, context flow.UserContext) {
	if flag1 < 9 /*dump first three packets */ {
		fmt.Println("----------------------------------------------------------")
		fmt.Printf("GTP F1\n")

		fmt.Printf("%v", currentPacket.Ether)
		currentPacket.ParseL3()
		ipv4 := currentPacket.GetIPv4()
		if ipv4 != nil {
			fmt.Printf("%v", ipv4)
			tcp, udp, _ := currentPacket.ParseAllKnownL4ForIPv4()
			if tcp != nil {
				fmt.Printf("%v", tcp)
			} else if udp != nil {
				fmt.Printf("%v", udp)
				gtp := currentPacket.GTPIPv4FastParsing()
				fmt.Printf("%v", gtp)
			} else {
				println("ERROR")
			}
		} else {
			println("ERROR")
		}
		fmt.Println("----------------------------------------------------------")
		flag1++
	}
}

func dumpPrintF2(currentPacket *packet.Packet, context flow.UserContext) {
	//if flag1 < 9 /*dump first three packets */ {
	fmt.Printf("drop F2")
	fmt.Printf("%v", currentPacket.Ether)
	currentPacket.ParseL3()
	ipv4 := currentPacket.GetIPv4()
	if ipv4 != nil {
		fmt.Printf("%v", ipv4)
		tcp, udp, _ := currentPacket.ParseAllKnownL4ForIPv4()
		if tcp != nil {
			fmt.Printf("%v", tcp)
		} else if udp != nil {
			fmt.Printf("%v", udp)
			gtp := currentPacket.GTPIPv4FastParsing()
			fmt.Printf("%v", gtp)
		} else {
			println("ERROR")
		}
	} else {
		println("ERROR")
	}
	fmt.Println("----------------------------------------------------------")
	//flag1++
	//}
}

func dumpPrintPortCxt(currentPacket *packet.Packet, context flow.UserContext) {
	//if flag1 < 9 /*dump first three packets */ {
	fmt.Printf("Port context\n")
	fmt.Printf("%v\n", currentPacket.Ether)
	port := context.(up.HandlerContext)
	fmt.Printf("%+v\n", port)
	fmt.Printf("%+v\n", port.Ipport)

	currentPacket.ParseL3()
	ipv4 := currentPacket.GetIPv4()
	if ipv4 != nil {
		fmt.Printf("%v", ipv4)
		tcp, udp, _ := currentPacket.ParseAllKnownL4ForIPv4()
		if tcp != nil {
			fmt.Printf("%v", tcp)
		} else if udp != nil {
			fmt.Printf("%v", udp)
			gtp := currentPacket.GTPIPv4FastParsing()
			fmt.Printf("%v", gtp)
		} else {
			println("ERROR")
		}
	} else {
		println("ERROR")
	}
	fmt.Println("----------------------------------------------------------")
	//flag1++
	//}
}

func DnnListToMap(dnnlist string) map[string]string {
	DnnNameGwIpMap := make(map[string]string)
	gwIpMap := dnnlist
	if len(gwIpMap) == 0 {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Failed to load N6 dnn gateway ip from config file, set to default.")
		DnnNameGwIpMap[defs.Dnn] = defs.DnnIP
	} else {
		// dnn_name_gw_ip_map = cmnet:172.16.1.200,cmnet2:172.16.1.201
		kvStr := strings.Split(gwIpMap, ",")
		var DnnGwkv defs.DnnGwIpMapKV
		for _, vStr := range kvStr {
			// cmnet:172.16.1.200
			kv := strings.Split(vStr, ":")
			if len(kv) == 2 {
				DnnGwkv.DnnName = kv[0]
				DnnGwkv.GwIp = kv[1]
				DnnNameGwIpMap[DnnGwkv.DnnName] = DnnGwkv.GwIp
			}
		}
	}
	return DnnNameGwIpMap
}