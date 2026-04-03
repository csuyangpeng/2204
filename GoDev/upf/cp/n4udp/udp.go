package n4udp

import (
	"fmt"
	"lite5gc/cmn/message/pfcp/pfcpudp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/upf/cp/n4layer"

	"net"
)

// udp server
// server
//var UdpAddr = net.UDPAddr{IP: net.ParseIP(pfcpudp.Pfcp_Default_IP), Port: pfcpudp.Pfcp_Default_Port}
//var Server = pfcpudp.NewUdpServer(&UdpAddr)
var Server = pfcpudp.PfcpServer

func StartPfcpUdpListen(s *pfcpudp.Server) error {
	rlogger.FuncEntry(moduleTag, nil)
	listener, err := net.ListenUDP(
		"udp", s.LocalAddr)
	if err != nil {
		panic(fmt.Errorf("failed to Start Pfcp Udp Listen %s", err))
	}

	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
	//rlogger.Trace(moduleTag, rlogger.INFO, nil, "Local: <%s> \n", listener.LocalAddr().String())
	defer listener.Close()

	// 处理消息
	s.UdpConn = listener
	for i := 0; i < pfcpudp.HandleMsg_GOROUTINE_NUMBER; i++ {
		//upfCxt.Wg.Add(1)
		fmt.Printf("Start Pfcp handle(%v) \n", i)
		go handleRequestMsg(s)
	}

	// performance improvement
	// 当前使用方式 2 ，静态申请
	for {
		var data [pfcpudp.UDP_BUFFER]byte

		n, remoteAddr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			//fmt.Printf("error during read: %s", err)
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "failed to read from udp")
			//if err == syscall.EINVAL {
			//	return err
			//}
			continue
		}
		//fmt.Printf("receive msg: <%#x>,from %s,to %s\n", data[:n], remoteAddr.String(), listener.LocalAddr().String())
		rlogger.Trace(moduleTag, rlogger.INFO, nil,
			"receive msg: <%#x> ,from %s,to %s", data[:n], remoteAddr.String(), listener.LocalAddr().String())
		//receive msg: <hello> ,from 127.0.0.1:56191,to [::]:8805
		msgCxt := pfcpudp.MsgCxt{Msgbuf: data[:n],
			RemoteAdd: remoteAddr,
		}
		recvRet := s.ReceiveMsg(msgCxt)
		if recvRet != true {
			//fmt.Printf("The buffer is full, discarding the message\n")
			rlogger.Trace(moduleTag, rlogger.ERROR, nil, "The buffer is full, discarding the message")
		}

	}
	return nil

}

// go,处理收的消息，并发送响应消息
func handleRequestMsg(s *pfcpudp.Server) error {
	rlogger.FuncEntry(moduleTag, nil)
	for {

		select {
		case msgCxt := <-s.RevMsgListChan:
			//fmt.Printf("receive msg:<%#x>,\nfrom %s,to %s \n", msgCxt.Msgbuf, s.UdpConn.LocalAddr().String(), msgCxt.RemoteAdd.String())
			rlogger.Trace(moduleTag, rlogger.INFO, nil,
				"receive msg:<%#x>,\n  from %s,to %s \n", msgCxt.Msgbuf, msgCxt.RemoteAdd.String(), s.UdpConn.LocalAddr().String())
			//pfcp decode，result msg object

			//todo pfcp decode
			msg := pfcpv1.Message{}
			err := msg.Unmarshal(msgCxt.Msgbuf)
			if err != nil {
				rlogger.Trace(moduleTag, rlogger.ERROR, nil,
					"receive msg:<%#x>,\n  unmarshal error:%s \n", msgCxt.Msgbuf, err)
				continue
			}
			rlogger.Trace(moduleTag, rlogger.INFO, nil, "receive msg:(%s)", msg.String())
			res := &pfcpv1.Message{}
			err = Dispatch(msg, res)
			if err != nil {
				rlogger.Trace(moduleTag, rlogger.ERROR, nil,
					"receive msg:(%s),\n  Dispatch error:%s \n",
					msg.String(), err)
				continue
			}
			// response message end processing
			if res.Header.MessageType == 0 {
				continue
			}
			rlogger.Trace(moduleTag, rlogger.INFO, nil, "send msg:(%s)", res.String())
			// todo pfcp encode
			data, err := res.Marshal()
			if err != nil {
				rlogger.Trace(moduleTag, rlogger.ERROR, nil, "failed to marshale,error:%s", err)
				continue
			}
			s.UdpConn.WriteToUDP(data, msgCxt.RemoteAdd)
			//fmt.Printf("send msg:<%#x>,from %s,to %s \n", data, s.UdpConn.LocalAddr().String(), msgCxt.RemoteAdd.String())
			rlogger.Trace(moduleTag, rlogger.INFO, nil,
				"send msg:<%#x>,from %s,to %s \n", data, s.UdpConn.LocalAddr().String(), msgCxt.RemoteAdd.String())
			rlogger.Trace(moduleTag, rlogger.INFO, nil, "success to send msg")
		}
	}

}

// client

func SendUdpMsg(s *pfcpudp.Server) error {
	rlogger.FuncEntry(moduleTag, nil)
	//对端地址

	//srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	//dstAddr := &net.UDPAddr{IP: ip, Port: port}

	conn, err := net.DialUDP("udp", s.LocalAddr, s.PeerAddr)
	if err != nil {
		fmt.Println(err)
		//rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Failed to get UPD send connection!")
		return err
	}
	defer conn.Close()
	fmt.Println(conn)
	for {
		select {
		case UDPMsg := <-s.SendMsgListChan:
			n, err := conn.Write(UDPMsg.Msgbuf)
			if err != nil {
				fmt.Printf("failed to write: %v", err)
			}
			//rlogger.Trace(moduleTag, rlogger.INFO, nil,  "<%s>--><%s>: %#x\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg[:n])
			fmt.Printf("SendN4UdpMsg to Peer : <%s>--><%s>: %s\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg.Msgbuf[:n])
		}
	}

	return nil
}

func StartN4Server(Cxt *types.AppContext) error {
	rlogger.FuncEntry(moduleTag, nil)
	n4layer.UpfN4Layer.UpfIp = configure.UpfConf.N4.Local.Ipv4
	n4layer.UpfN4Layer.N3Ip = configure.UpfConf.N3.Ipv4
	// 开启N4监听
	//Cxt.Wg.Add(1)
	//go StartN4Listen(Cxt)
	Server.LocalAddr.IP = net.ParseIP(configure.UpfConf.N4.Local.Ipv4)
	Server.LocalAddr.Port = configure.UpfConf.N4.Local.Port
	go StartPfcpUdpListen(Server)

	return nil
}
