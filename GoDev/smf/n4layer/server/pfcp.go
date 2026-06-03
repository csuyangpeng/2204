package server

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/pfcp/pfcpudp"
	"lite5gc/cmn/types/configure"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"strconv"

	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

var SMFCtxt context.Context

// receive response
// json rpc server 目前实行是TCP， 请求消息的UDP目标端口号应为8805.它是PFCP的注册端口号。

func StartN4Server(Cxt *types.AppContext) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 开启N4监听
	//go StartN4Listen(Cxt)
	n4Ip := configure.SmfConf.N4Conf.SMFIP
	n4Port := configure.SmfConf.N4Conf.SMFPort // 8805
	pfcpudp.PfcpServer.LocalAddr.IP = n4Ip
	pfcpudp.PfcpServer.LocalAddr.Port = n4Port
	//fmt.Println("-----------===")
	go StartPfcpUdpListen(pfcpudp.PfcpServer)
	return nil
}

func StartN4Listen(Cxt *types.AppContext) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)

	// 注册rpc服务模块
	//rpc.Register(new(N4Msg))
	n4Ip := configure.SmfConf.N4Conf.SMFIP.String()
	n4Port := configure.SmfConf.N4Conf.SMFPort // 8805
	ipAddr := n4Ip + ":" + strconv.Itoa(n4Port)
	lis, err := net.Listen("tcp4", ipAddr) //"10.180.8.51:8805"
	if err != nil {
		fmt.Printf("fatal error: %s", err)
		return

	}

	fmt.Fprintf(os.Stdout, "N4 server start connection %s\n", ipAddr)
	SMFCtxt = Cxt.Ctx

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4Listen server routine start")
	Cxt.Wg.Add(1)
	defer func() {
		Cxt.Wg.Done()
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "N4Listen server routine exit")
	}()
	fdClosed := false
	go func() {
		select {
		case <-Cxt.Ctx.Done():
			fdClosed = true
			lis.Close()
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "N4Listen server monitor routine exit")
			return
		}
	}()

	for {
		conn, err := lis.Accept() // 接收客户端连接请求
		if err != nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "N4Listen server Accept error occur: ", err)
			if fdClosed {
				break
			}
			continue
		}

		go func(conn net.Conn) { // 并发处理客户端请求
			//fmt.Fprintf(os.Stdout, "%s", "new client in coming\n")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

// json rpc client

// send request
