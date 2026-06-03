package main

import (
	"C"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
	"unsafe"

	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"

	"github.com/astaxie/beego/logs"
	"github.com/ishidawataru/sctp"
)

func main() {

	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	config := make(map[string]interface{})
	config["filename"] = "./sctpTestClt.log"
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	rlogger.Trace(MODULE_ID, types.INFO, nil, "logger initailzed success!")

	// var server = flag.Bool("server", false, "")
	var ip = flag.String("ip", "192.168.119.129", "")
	var port = flag.Int("sport", 0, "")
	var lport = flag.Int("lport", 0, "")

	flag.Parse()

	rlogger.Trace(MODULE_ID, types.DEBUG, nil, "ip = %s, ser port = %d, local port = %d", *ip, *port, *lport)
	ips := []net.IPAddr{}

	for _, i := range strings.Split(*ip, ",") {
		if a, err := net.ResolveIPAddr("ip", i); err == nil {
			rlogger.Trace(MODULE_ID, types.DEBUG, nil, "Resolved address '%s' to %s", i, a)
			ips = append(ips, *a)
		} else {
			rlogger.Trace(MODULE_ID, types.ERROR, nil, "Error resolving address '%s': %v", i, err)
		}
	}

	addr := &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    *port,
	}
	var laddr *sctp.SCTPAddr
	if *lport != 0 {
		laddr = &sctp.SCTPAddr{
			Port: *lport,
		}
	}
	conn, err := sctp.DialSCTP("sctp", laddr, addr)
	if err != nil {
		rlogger.Trace(MODULE_ID, types.FATAL, nil, "dest port(%d), failed to dial: %v", addr.Port, err)
		return
	}
	rlogger.Trace(MODULE_ID, types.DEBUG, nil, "Dail LocalAddr: %s; RemoteAddr: %s", conn.LocalAddr(), conn.RemoteAddr())

	ppid := 0
	for {
		info := &sctp.SndRcvInfo{
			Stream: uint16(ppid),
			PPID:   uint32(ppid),
		}
		ppid++
		conn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO)

		msg := createNgApSetupRequest(ossCtxtPtr)

		n, err := conn.Write(msg)
		if err != nil {
			rlogger.Trace(MODULE_ID, types.ERROR, nil, "failed to send message")
			continue
		}

		// n, err := conn.SCTPWrite([]byte(msg), info)
		// if err != nil {
		// 	rlogger.Trace(MODULE_ID, types.ERROR, nil, "failed to write: %v", err)
		// }
		rlogger.Trace(MODULE_ID, types.DEBUG, nil, "write: %d", n)
		buf := make([]byte, 254)
		n, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				rlogger.Trace(MODULE_ID, types.WARN, nil, "sctp connecton closed. err : %s", err)
				return
			} else {
				rlogger.Trace(MODULE_ID, types.WARN, nil, "Read Error. err : %s", err)
			}
		}

		// _, info, err = conn.SCTPRead(buf)
		// if err != nil {
		// 	rlogger.Trace(MODULE_ID, types.ERROR, nil, "failed to read: %v", err)
		// }
		log.Printf("read: info: %+v", info)
		rlogger.Trace(MODULE_ID, types.DEBUG, nil, "Read from 5GC: %s", string(buf))
		time.Sleep(time.Second * 3)
	}
}

func createNgApSetupRequest(ossCtxtPtr codec.NgApOssCtxt) []byte {

	ngSetupReqEncode := codec.NewNgSetupRequestCodec()
	ngSetupReqEncode.SetRanNodeName("RAN01")
	ngSetupReqEncode.SetPagingDrx(1)

	//for gloabl gnb id
	ggnbid := codec.NewGGnbId()
	plmnid := []byte("460")
	ggnbid.SetPlmnid(&plmnid[0])

	gnbid := codec.NewGnbId()
	gnbid.SetLen(22)
	val := []byte("ABC")
	gnbid.SetVal(&val[0])

	ggnbid.SetGnbId(gnbid)
	ngSetupReqEncode.SetGgnbId(ggnbid)

	//for support ta list
	snssai1 := codec.NewSNssai()
	sst1 := byte('A')
	snssai1.SetSst(&sst1)
	sd1 := []byte("111")
	snssai1.SetSd(&sd1[0])
	snssai1.SetSdPresent(true)

	snssai2 := codec.NewSNssai()
	sst2 := byte('B')
	snssai2.SetSst(&sst2)
	sd2 := []byte("222")
	snssai2.SetSd(&sd2[0])
	snssai2.SetSdPresent(true)

	ssList := codec.NewSNssaiVector()
	ssList.Add(snssai1)
	ssList.Add(snssai2)

	bplmn := codec.NewBPlmnItem()
	bplmnid := []byte("461")
	bplmn.SetPlmnid(&bplmnid[0])
	bplmn.SetSsList(ssList)

	bplmnList := codec.NewBPlmnItemVector()
	bplmnList.Add(bplmn)

	stai := codec.NewSupTAItem()
	tac := []byte("123")
	stai.SetTac(&tac[0])
	stai.SetBplmnList(bplmnList)

	ngSetupReqEncode.AddSupTAList(stai)

	ngSetupReqEncode.DumpMessage()
	msgBuf := ngSetupReqEncode.Encode(ossCtxtPtr)
	bufLen := msgBuf.GetLength()
	bufValue := msgBuf.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))

	return encodeBuffer
}
