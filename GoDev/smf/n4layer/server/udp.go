package server

import (
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/pfcpudp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/n4layer/server/pfcphandler"
	"net"
	"syscall"
)

// udp Server

func StartPfcpUdpListen(s *pfcpudp.Server) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//fmt.Println("---conn---StartPfcpUdpListen", s.LocalAddr)
	listener, err := net.ListenUDP(
		"udp", s.LocalAddr)
	if err != nil {
		fmt.Println("---conn---StartPfcpUdpListen err", err)
		panic(err)
	}

	//fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Local: <%s> \n", listener.LocalAddr().String())
	defer listener.Close()

	// 处理消息
	s.UdpConn = listener
	//fmt.Println("---conn---", s.UdpConn)
	for i := 0; i < pfcpudp.HandleMsg_GOROUTINE_NUMBER; i++ {
		//upfCxt.Wg.Add(1)
		// 收消息,2个处理协程
		go handleReceiveMsg(s)
	}
	// 发消息，
	//go StartPfcpUdpSendServer(s)
	// performance improvement
	// 当前使用方式 2 ，静态申请
	for {
		var data [pfcpudp.UDP_BUFFER]byte

		n, remoteAddr, err := listener.ReadFromUDP(data[:])
		if err != nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "error during read: %s", err)
			if err == syscall.EINVAL {
				return err
			}
			continue
		}
		//fmt.Printf("receive msg: <%#x> ,from %s,to %s\n", data[:n], remoteAddr.String(), listener.LocalAddr().String())
		//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
		//	"receive msg: <%#x> ,from %s,to %s\n", data[:n], remoteAddr.String(), listener.LocalAddr().String())
		//receive msg: <hello> ,from 127.0.0.1:56191,to [::]:8805
		msgCxt := pfcpudp.MsgCxt{Msgbuf: data[:n],
			RemoteAdd: remoteAddr}
		recvRet := s.ReceiveMsg(msgCxt)
		if recvRet != true {
			//fmt.Printf("The buffer is full, discarding the message\n")
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "The buffer is full, discarding the message")
		}

	}
	return nil

}

// client

func StartPfcpUdpSendServer(s *pfcpudp.Server) error {
	//对端地址

	//srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	//dstAddr := &net.UDPAddr{IP: ip, Port: port}

	/*conn, err := net.DialUDP("udp", s.LocalAddr, s.PeerAddr)
	if err != nil {
		fmt.Println(err)
		//rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failed to get UPD send connection!")
		return err
	}
	defer conn.Close()*/
	//fmt.Println(conn)
	for {
		select {
		case UDPMsg := <-s.SendMsgListChan:
			conn := s.UdpConn
			n, err := conn.WriteToUDP(UDPMsg.Msgbuf, UDPMsg.RemoteAdd)
			if err != nil {
				fmt.Printf("failed to write: %v", err)
			}
			//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,  "<%s>--><%s>: %#x\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg[:n])
			//fmt.Printf("SendN4UdpMsg to Peer : <%s>--><%s>: %#x\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg.Msgbuf[:n])
			fmt.Printf("SendN4UdpMsg to Peer : <%s>--><%s>: %s\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg.Msgbuf[:n])
		}
	}

	return nil
}

func SendMsgHandler(msg interface{}) error {
	pfcpMsg := pfcpv1.Message{}
	switch request := msg.(type) {
	//PDR ID  	M
	case *pfcp.SessionEstablishmentRequest:

		pfcpMsg.Header.Version = request.PfcpHeader.Version
		pfcpMsg.Header.SFlag = request.PfcpHeader.SFlag
		pfcpMsg.Header.MessageType = request.PfcpHeader.MessageType
		pfcpMsg.Header.Length = request.PfcpHeader.Length
		pfcpMsg.Header.SEID = request.PfcpHeader.SEID
		pfcpMsg.Header.SequenceNumber = request.PfcpHeader.SequenceNumber

		pfcpMsg.Body = request
	default:
		return fmt.Errorf("Illegal IE")
	}
	data, err := pfcpMsg.Marshal()
	if err != nil {
		return err
	}
	pfcpudp.PfcpServer.SendMsg(data)
	return nil
}

func handleReceiveMsg(s *pfcpudp.Server) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	for {

		select {
		case msgCxt := <-s.RevMsgListChan:
			rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, nil,
				"receive msg:<%#x> ,from %s,to %s", msgCxt.Msgbuf,
				msgCxt.RemoteAdd.String(), s.UdpConn.LocalAddr().String())
			//分发消息到对应的节点处理
			peerIp := msgCxt.RemoteAdd.IP.String()
			/*	node, ok := node.NodePool[peerIp]
				if !ok {
					rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
						"The corresponding processing node does not exist(%s)", peerIp)
				}*/
			//todo pfcp decode
			msg := pfcpv1.Message{}
			err := msg.Unmarshal(msgCxt.Msgbuf)
			if err != nil {
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
					"Invalid message(%s),data:%#x", err, msgCxt.Msgbuf)
			}
			//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "receive msg:(%s)", msg.String())
			res := &pfcpv1.Message{}
			err = pfcphandler.NodeDispatch(peerIp, msg, res)
			if err != nil {
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,

					"Invalid message(%s)", err)
			}
			/*// todo pfcp encode
			data, err := res.Marshal()
			if err == nil {
				s.UdpConn.WriteToUDP(data, msgCxt.RemoteAdd)
			}*/
		}
	}

}
