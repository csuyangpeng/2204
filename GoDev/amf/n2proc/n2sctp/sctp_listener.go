package n2sctp

import (
	"context"
	"fmt"
	"lite5gc/cmn/redisclt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types/configure"
	"net"
	"strings"
	"sync"
	"time"

	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/sctp"
	"lite5gc/cmn/types"
)

// N2 Listener Structure
type N2Listener struct {
	ipAddr     sctp.SCTPAddr
	sctpOpts   SctpOptions
	ctx        context.Context
	cancel     context.CancelFunc
	conns      *sync.Map
	wg         *sync.WaitGroup
	mu         sync.Mutex
	bufferSize int // size of buffered channel
}

// New N2Listener
func NewN2Listener(ipaddr string, port int, options *SctpOptions, ctx context.Context) (*N2Listener, error) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	ipAddrs := []net.IPAddr{}

	for _, i := range strings.Split(ipaddr, ",") {
		if addrptr, err := net.ResolveIPAddr("ip", i); err == nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil, "Resolved address '%s' to %s", i, addrptr.String())
			ipAddrs = append(ipAddrs, *addrptr)
		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "Error resolving address '%s': %v", i, err)
			return nil, fmt.Errorf("%s, %s", types.ErrInvParm, err)
		}
	}

	sctpServer := &N2Listener{}
	sctpServer.ipAddr.IPAddrs = append(sctpServer.ipAddr.IPAddrs, ipAddrs...)
	sctpServer.ipAddr.Port = port

	sctpServer.sctpOpts = *options

	sctpServer.conns = &sync.Map{}
	sctpServer.wg = &sync.WaitGroup{}
	sctpServer.bufferSize = types.BufSize8192

	sctpServer.ctx, sctpServer.cancel = context.WithCancel(ctx)

	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil,
		"N2 Listener Sctp Address: %+v", sctpServer.ipAddr.String())

	// create the redis clt for communication with amf main process
	//gnblayer.RedisClt, err = redisclt.New(
	//	redisclt.Options{
	//		Addr:     "10.18.1.56:6379",
	//		Password: "",
	//		Prefix:   "cn_",
	//	})
	//if err != nil {
	//	panic(err)
	//}
	//err = gnblayer.RedisClt.CheckHealth()
	//if err != nil {
	//	fmt.Println("failed to connect to redis server")
	//	return nil, fmt.Errorf("failed to connect to redis server")
	//}
	err := redisclt.RedisCltInit()
	if err != nil {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.FATAL, nil,
			"failed to connect redis server(%s:%d) ",
			configure.SysConf.RedisAddr.Ip,
			configure.SysConf.RedisAddr.Port)
		return nil, fmt.Errorf("failed to connect to redis server")
	}
	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
		"connect to redis success(%s:%d)", configure.SysConf.RedisAddr.Ip, configure.SysConf.RedisAddr.Port)

	return sctpServer, nil
}

// NumofSctpConn return number of sctp connections
func (p *N2Listener) NumofSctpConn() int {
	var sz int
	p.conns.Range(func(k, v interface{}) bool {
		sz++
		return true
	})
	return sz
}

func (p *N2Listener) DelSctpConn(ngapId uint32) {
	p.conns.Delete(ngapId)
	return
}

func (p *N2Listener) destroy() {
	p.cancel()
	p.wg.Wait()
}

// Start function, activate the n2 sctp listener server
func (p *N2Listener) Start(appWg *sync.WaitGroup) error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)
	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "n2 sctp listener is running...")
	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, "sctp listen on ip addr (%s)", p.ipAddr.String())

	// sctp listen
	ln, err := sctp.ListenSCTP("sctp", &p.ipAddr)
	if err != nil {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to listen: %v", err)
		return err
	}
	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, "Sctp Listener on %s", ln.Addr().String())

	// no delay
	ln.SetNoDelay(true)

	initOpts := sctp.InitMsg{
		NumOstreams:    p.sctpOpts.NumOstreams,
		MaxInstreams:   p.sctpOpts.MaxInstreams,
		MaxAttempts:    p.sctpOpts.MaxAttempts,
		MaxInitTimeout: p.sctpOpts.MaxInitTimeout}
	ln.SetInitMsg(initOpts)
	hbInv := uint32(p.sctpOpts.HeatbeatInterval)
	pathMacRxt := uint32(p.sctpOpts.PathMaxRXT)
	ln.SetHeatbeatInterval(hbInv, pathMacRxt)

	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, "sctp listener main loop go routine start")
	defer func() {
		appWg.Done()
		if p := recover(); p != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "panics: %v", p)
		}

		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "sctp main loop go routine exit")
	}()

	//goroutine monitor for cancel singal
	appWg.Add(1)
	go func() {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, "sctp main loop monitor routine start")
		defer appWg.Done()
		select {
		case <-p.ctx.Done():
			p.destroy()
			ln.Close()
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "sctp main loop monitor routine exit")
		}
	}()

	//listener main loop
	var tempDelay time.Duration
	for {
		rawConn, err := ln.Accept()
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, "accept a new sctp connection")

		if err != nil {
			// handle the error for accept action
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay >= max {
					tempDelay = max
				}

				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, "accept error %v, retrying in %d", err, tempDelay)

				select {
				case <-time.After(tempDelay):
				case <-p.ctx.Done():
					rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "n2 listener mainloop go routine exit.")
					return fmt.Errorf("n2 listener mainloop go routine cancel triggered.")
				}
				continue
			}
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.FATAL, nil, "accept error (%s),not net error.", err)
			return err
		}
		tempDelay = 0

		// checking duplicate sctp connection
		var isSctpConnExist bool
		p.conns.Range(func(k, v interface{}) bool {
			c := v.(*N2SctpConn)
			if c.GetGnbIpAddr() == rawConn.RemoteAddr().String() {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "%v: %v is already exist \n", k, v)
				isSctpConnExist = true
				return true
			}
			return false
		})
		if isSctpConnExist {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "the sctp connection duplicated")
			continue
		}

		// check the max connections limitation

		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil,
			"Accepted Connection : %s", rawConn.RemoteAddr().String())

		sctpConn := rawConn.(*sctp.SCTPConn)

		// allocate the ngap id from id mgr
		gnbInstId, err := idmgr.GetInst().BorrowID(string(types.NGAP))

		n2SctpConn := NewN2SctpConn(gnbInstId, p, sctpConn, appWg)

		n2SctpConn.SetGnbIpAddr(strings.Split(sctpConn.RemoteAddr().String(), ":")[0])

		//insert into sctp connection map
		p.conns.Store(gnbInstId, n2SctpConn)

		go func() {
			n2SctpConn.Start(p.ctx)
		}()

		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
			"accepted client %s, gnb inst id (%d)",
			n2SctpConn.GetGnbIpAddr(),
			gnbInstId)

		p.conns.Range(func(k, v interface{}) bool {
			i := k.(uint32)
			c := v.(*N2SctpConn)
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil, "gnb inst id(%d) ip(%s)", i, c.GetGnbIpAddr())
			return true
		})
	}
}
