package pfcpudp

import (
	"net"
)

//const (
//	PFCP_PORT        = 8805
//	PFCP_MAX_UDP_LEN = 2048
//)

//type PfcpServer struct {
//	Addr string
//	Conn *net.UDPConn
//}
//
//func NewPfcpServer(addr string) (PfcpServer, error) {}
//
//func (p *PfcpServer) Listen() error {}
//
//func (p *PfcpServer) ReadFrom(msg *pfcp.Message) (*net.UDPAddr, error) {}
//
//func (p *PfcpServer) WriteTo(msg pfcp.Message, addr *net.UDPAddr) error {}
//
//func (p *PfcpServer) Close() error {}
//
//// Send a PFCP message and close UDP connection
//func SendPfcpMessage(msg pfcp.Message, srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) error {}
//
//// Receive a PFCP message and close UDP connection
//func ReceivePfcpMessage(msg *pfcp.Message, srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) error {}

const (
	CHANCAP    = 1000 // udp receive channel buffer //1000
	UDP_BUFFER = 2048 //2048 socket receive buffer

	HandleMsg_GOROUTINE_NUMBER = 2

	Pfcp_Default_IP   = "172.16.1.10"
	Pfcp_Default_Port = 8805
)

type MsgCxt struct {
	Msgbuf    []byte
	RemoteAdd *net.UDPAddr
}

// N4 udp server
type Server struct {
	LocalAddr *net.UDPAddr //N4 ip
	PeerAddr  *net.UDPAddr //N4 ip // 保存接收消息时的IP

	RevMsgListChan  chan MsgCxt // receive UPF msg list
	SendMsgListChan chan MsgCxt // send msg list，不使用，直接发送UDP

	UdpConn *net.UDPConn // send response
}

func NewUdpServer(udpAddr *net.UDPAddr) *Server {
	return &Server{LocalAddr: udpAddr,
		RevMsgListChan:  make(chan MsgCxt, CHANCAP),
		SendMsgListChan: make(chan MsgCxt, CHANCAP)}
}

// ReceiveMsg Send the received message to the channel
func (s *Server) ReceiveMsg(msgCxt MsgCxt) bool {
	select {
	case s.RevMsgListChan <- msgCxt:
		return true
	default:
		return false //队列已满，消息丢弃

	}
	//fmt.Printf("chan: %s",<-s.MsgListChan)
}

// SendMsg Send the received message to the channel
func (s *Server) SendMsg(msg []byte) bool {

	msgCxt := MsgCxt{Msgbuf: msg}
	select {
	case s.SendMsgListChan <- msgCxt:
		return true
	default:
		return false //队列已满，消息丢弃

	}
	//fmt.Printf("chan: %s",<-s.MsgListChan)
}

func (s *Server) SendMsgTo(msg MsgCxt) bool {

	select {
	case s.SendMsgListChan <- msg:
		return true
	default:
		return false //队列已满，消息丢弃

	}
	//fmt.Printf("chan: %s",<-s.MsgListChan)
}

// Server
var udpAddr = net.UDPAddr{IP: net.ParseIP(Pfcp_Default_IP), Port: Pfcp_Default_Port}
var PfcpServer = NewUdpServer(&udpAddr)
