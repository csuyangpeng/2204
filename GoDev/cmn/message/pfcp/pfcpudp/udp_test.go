package pfcpudp_test

import (
	"fmt"
	"lite5gc/cmn/message/pfcp/pfcpudp"
	"net"
	"testing"
)

// udp server
// server
var udpAddr = net.UDPAddr{IP: net.ParseIP(pfcpudp.Pfcp_Default_IP), Port: pfcpudp.Pfcp_Default_Port}
var server = pfcpudp.NewUdpServer(&udpAddr)

func TestNewUdpServer(t *testing.T) {
	// 开启8805监听
	//StartN4Listen(server)
}

// client
// 1 重用DiaUDP 连接 ,ok
func TestSendUdpMsg(t *testing.T) {
	server.LocalAddr = &net.UDPAddr{IP: net.IPv4zero, Port: 8805}
	server.PeerAddr = &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8805}
	go func() {
		for i := 0; i < 10; i++ {
			recvRet := server.SendMsg([]byte("hello"))
			if recvRet != true {
				fmt.Printf("The buffer is full, discarding the message\n")
				//rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "The buffer is full, discarding the message")
			}
		}

	}()

	//SendUdpMsg(server)

}

// 2 相同五元组，DiaUDP 连接是否相同,是相同的，系统不允许
func TestDupSendUdpMsg(t *testing.T) {
	server.LocalAddr = &net.UDPAddr{IP: net.IPv4zero, Port: 8805}
	server.PeerAddr = &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8805}
	go func() {
		for i := 0; i < 10; i++ {
			recvRet := server.SendMsg([]byte("hello"))
			if recvRet != true {
				fmt.Printf("The buffer is full, discarding the message\n")
				//rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "The buffer is full, discarding the message")
			}
		}

	}()

	for {
		select {
		case UDPMsg := <-server.SendMsgListChan:
			conn, err := net.DialUDP("udp", server.LocalAddr, server.PeerAddr)
			if err != nil {
				fmt.Println(err)
				//rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "Failed to get UPD send connection!")
			}
			//dial udp 0.0.0.0:8805->127.0.0.1:8805:
			// bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.
			// 每个套接字地址（协议/网络地址/端口）通常只允许使用一次。
			fmt.Println(*conn)

			n, err := conn.Write(UDPMsg.Msgbuf)
			if err != nil {
				fmt.Printf("failed to write: %v", err)
			}
			//rlogger.Trace(types.ModCmn, rlogger.INFO, nil,  "<%s>--><%s>: %#x\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg[:n])
			//fmt.Printf("SendN4UdpMsg to Peer : <%s>--><%s>: %#x\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg.Msgbuf[:n])
			fmt.Printf("SendN4UdpMsg to Peer : <%s>--><%s>: %s\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg.Msgbuf[:n])
			conn.Close()
		}
	}

}

// 3 DiaUDP 连接接收请求消息
func TestDialSendUdpMsg(t *testing.T) {
	server.LocalAddr = &net.UDPAddr{IP: net.IPv4zero, Port: 8806}
	server.PeerAddr = &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 8805}
	go func() {
		for i := 0; i < 10; i++ {
			recvRet := server.SendMsg([]byte("hello"))
			if recvRet != true {
				fmt.Printf("The buffer is full, discarding the message\n")
				//rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "The buffer is full, discarding the message")
			}
		}

	}()

	conn, err := net.DialUDP("udp", server.LocalAddr, server.PeerAddr)
	if err != nil {
		fmt.Println(err)
		//rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "Failed to get UPD send connection!")
	}
	defer conn.Close()

	for {
		select {
		case UDPMsg := <-server.SendMsgListChan:
			// 发送请求
			n, err := conn.Write(UDPMsg.Msgbuf)
			if err != nil {
				fmt.Printf("failed to write: %v", err)
			}
			//rlogger.Trace(types.ModCmn, rlogger.INFO, nil,  "<%s>--><%s>: %#x\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg[:n])
			//fmt.Printf("SendN4UdpMsg to Peer : <%s>--><%s>: %#x\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg.Msgbuf[:n])
			fmt.Printf("SendN4UdpMsg to Peer : <%s>--><%s>: %s\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg.Msgbuf[:n])
			//	接收响应
			msg := make([]byte, 1024)
			n, err = conn.Read(msg) // dial conn ，不能用于接收消息
			fmt.Printf("ReceiveN4UdpMsg from Peer : <%s>--><%s>: %s\n", conn.RemoteAddr(), conn.LocalAddr(), msg[:n])
		}
	}

}

func TestDialReceive(t *testing.T) {
	server.LocalAddr = &net.UDPAddr{IP: net.IPv4zero, Port: 8805}
	server.PeerAddr = &net.UDPAddr{IP: net.ParseIP(""), Port: 0}

	conn, err := net.DialUDP("udp", server.LocalAddr, server.PeerAddr)
	if err != nil {
		fmt.Println(err)
		//rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "Failed to get UPD send connection!")
	}
	defer conn.Close()

	for {
		//	接收响应
		msg := make([]byte, 1024)
		n, _ := conn.Read(msg) // dial conn ，不能用于接收消息
		fmt.Printf("ReceiveN4UdpMsg from Peer : <%s>--><%s>: %s\n", conn.RemoteAddr(), conn.LocalAddr(), msg[:n])
	}

}
