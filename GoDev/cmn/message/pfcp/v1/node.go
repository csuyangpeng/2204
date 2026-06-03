package pfcpv1

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/pfcpudp"
	"lite5gc/smf/smfcontext/gctxt"
	"net"
	"sync"
	"time"
)

/*var (
	nodePool       map[int]*Node
	indexNodeCount = 0
)*/

type NodeState int

const (
	NodeStart       NodeState = 0
	NodeActive      NodeState = 1
	NodeUpdate      NodeState = 2
	NodeDeactivated NodeState = 3
)

type NodeRelatedState int

const (
	NodeMsgStart   NodeRelatedState = 0
	NodeMsgSetup   NodeRelatedState = 1
	NodeMsgUpdate  NodeRelatedState = 2
	NodeMsgRelease NodeRelatedState = 3
)

func ReceiveNode(seq int) {}

type Node struct {
	//A Node ID may be set to an FQDN or an IP address (see subclause 8.2.38).
	//使用对端IP标识一个节点
	NodeID string
	State  NodeState // 用于业务消息发送前的检查，在active 事件中设置
	// 初始对端IP
	PeerAddr *net.UDPAddr
	// udp server
	Server *pfcpudp.Server

	// session established f-SEID，会话创建后F-SEID中的IP
	SeidIp net.IP
	// teid range
	N3Ip      net.IP
	TEIDRI    uint8 // partition,分区的位数
	TEIDRange uint8 // partition value，,分区位的值，smf 分配Teid的位数是32-TEIDRI

	//心跳响应
	HeartbeatResponse bool

	//upf 主动关闭
	UpfShutdown bool

	// Sequence Number,节点的当前事务编号,初始值为0
	SequenceNumber uint32

	// 记录节点管理消息的中间状态
	NodeAssociationState NodeRelatedState

	// 业务消息接口
	ScCxt           context.Context    // 与N11 sc 交互的上下文
	ServiceMsg      *ServiceMsgChan    // 用做响应上传到smf sc
	recvMsgBuffList chan ServiceMsg    // 节点内部缓存业务消息
	recvMsgList     chan ServiceMsg    // 节点内正常session消息转发处理，用于多并发
	Cxt             context.Context    // 节点开启的goroutine使用，关闭后退出
	CxtCancel       context.CancelFunc // 用于释放节点资源，关闭后退出

	Request  *Message
	Response *Message
	Count    uint32

	Mu   sync.Mutex
	Wait sync.WaitGroup
}

func CreateNode(cxt context.Context, name *net.UDPAddr) (node *Node, err error) {
	cxt, cancel := context.WithCancel(cxt)
	return &Node{NodeID: name.IP.String(),
		PeerAddr:        name,
		State:           NodeStart,
		recvMsgList:     make(chan ServiceMsg, CHANCAP),
		recvMsgBuffList: make(chan ServiceMsg, CHANCAP),
		Cxt:             cxt,
		CxtCancel:       cancel}, nil
}

func RemoveNode(index int) {}

// 定时器定义
const (
	// 心跳定时器
	T1 = 60 * time.Second

	//	消息重传定时器
	T2 = 1000 * time.Millisecond
)

const (
	MaxRetransT2 = 3
)

type NodeTimer struct {
	//Heartbeat timer
	T1       *time.Timer
	T1Reset  func() // Heartbeat timer reset
	T1Cancel context.CancelFunc
	//Heartbeat Retrans timer
	//T1Retrans       *time.Timer
	T1RetransCancel context.CancelFunc

	//AssociationSetup Request
	T2       *time.Timer
	T2Cancel context.CancelFunc

	MaxRetransNumT2 int
}

func (n *Node) ServiceMsgReveive(msg ServiceMsg) bool {
	select {
	case n.recvMsgList <- msg:
		return true
	default:
		return false //队列已满，消息丢弃
	}
}
func (n *Node) ServiceMsgReveiveBuff(msg ServiceMsg) bool {
	select {
	case n.recvMsgBuffList <- msg:
		return true
	default:
		return false //队列已满，消息丢弃
	}
}

func (n *Node) ServiceMsgHandler() chan ServiceMsg {

	return n.recvMsgList

}
func (n *Node) ServiceMsgBuffHandler() chan ServiceMsg {

	return n.recvMsgBuffList

}
func (n *Node) SendUdpMsg(msg []byte) error {
	conn := n.Server.UdpConn
	_, err := conn.WriteToUDP(msg, n.PeerAddr)
	if err != nil {
		return fmt.Errorf("failed to write: %v", err)
	}
	//rlogger.Trace(types.ModCmn, rlogger.INFO, nil,  "<%s>--><%s>: %#x\n", conn.LocalAddr(), conn.RemoteAddr(), UDPMsg[:n])
	//fmt.Printf("SendN4UdpMsg to Peer : <%s>--><%s>: %#x\n", conn.LocalAddr(), n.PeerAddr, msg[:num])
	return nil
}

func HandlePfcpHeartbeatRequest(req pfcp.HeartbeatRequest, res *pfcp.HeartbeatResponse) error {
	res.PfcpHeader = req.PfcpHeader
	res.PfcpHeader.MessageType = pfcp.PFCP_Heartbeat_Response
	res.IE = req.IE
	return nil
}

func HandlePfcpHeartbeatResponse(msg *pfcp.HeartbeatResponse) error {
	return nil
}
func HandleAssociationSetupRequest(msg *pfcp.PFCPAssociationSetupRequest) error {
	return nil
}

func HandleAssociationSetupResponse(msg *pfcp.PFCPAssociationSetupResponse) error {
	return nil
}

func HandleAssociationUpdateRequest(msg *pfcp.PFCPAssociationUpdateRequest) error {
	return nil
}

func HandleAssociationUpdateResponse(msg *pfcp.PFCPAssociationUpdateResponse) error {
	return nil
}

func HandleAssociationReleaseRequest(msg *pfcp.PFCPAssociationReleaseRequest) error {
	return nil
}
func HandleAssociationReleaseResponse(msg *pfcp.PFCPAssociationReleaseResponse) error {
	return nil
}

func HandleNodeReportRequest(msg *pfcp.PFCPNodeReportRequest) error {
	return nil
}

func HandleNodeReportResponse(msg *pfcp.PFCPNodeReportResponse) error {
	return nil
}

//HandlePfcpPFDManagementRequest
func HandlePfcpPFDManagementRequest(msg *pfcp.PFCPPFDManagementRequest) error {
	return nil
}

//HandlePfcpPFDManagementResponse
func HandlePfcpPFDManagementResponse(msg *pfcp.PFCPPFDManagementResponse) error {
	return nil
}

const (
	CHANCAP = 1000 // udp receive channel buffer //1000
)

type ServiceMsg struct {
	Msg       SmfToNode
	RemoteAdd *net.UDPAddr
}

type SmfToNode struct {
	ID  pfcp.PFCPMSG
	Cxt *gctxt.ScN4MsgData
}

func (p ServiceMsg) IpcMsgDataIf() {}

type ServiceMsgChan struct {
	ReqMsgList chan ServiceMsg // request msg list
	ResMsgList chan ServiceMsg //  response msg list
}

func NewServiceMsgChan() *ServiceMsgChan {
	return &ServiceMsgChan{
		ReqMsgList: make(chan ServiceMsg, CHANCAP),
		ResMsgList: make(chan ServiceMsg, CHANCAP)}
}

// ReceiveMsg, Send the request message to the channel
func (s *ServiceMsgChan) ReceiveReqMsg(msg ServiceMsg) bool {
	select {
	case s.ReqMsgList <- msg:
		return true
	default:
		return false //队列已满，消息丢弃

	}
	//fmt.Printf("chan: %s",<-s.MsgListChan)
}

// SendMsg, Send the response message to the channel
func (s *ServiceMsgChan) ReceiveResMsg(msg ServiceMsg) bool {
	select {
	case s.ResMsgList <- msg:
		return true
	default:
		return false //队列已满，消息丢弃
	}
	//fmt.Printf("chan: %s",<-s.MsgListChan)
}

// 处理请求消息
func (s *ServiceMsgChan) HandleRequestMsg(f func(m ServiceMsg) (*ServiceMsg, error)) {
	for {

		select {
		case msgCxt := <-s.ReqMsgList:
			fmt.Printf("send msg:<%s> ,from %s,to %s \n", msgCxt.Msg, msgCxt.RemoteAdd.String())
			response, err := f(msgCxt)
			if err != nil {
				fmt.Println(err)
				//	发送失败响应
				s.ReceiveResMsg(*response)
			}
		}
	}
}

// 业务层处理响应消息
func (s *ServiceMsgChan) HandleResponseMsg(f func()) {
	for {

		select {
		case msgCxt := <-s.ResMsgList:
			fmt.Printf("send msg:<%s> ,from %s,to %s \n", msgCxt.Msg, msgCxt.RemoteAdd.String())
			f()
		}
	}
}
