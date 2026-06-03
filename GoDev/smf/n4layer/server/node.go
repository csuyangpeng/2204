package server

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/pfcpudp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
	"lite5gc/smf/n4layer/node/nodehandler"
	"lite5gc/smf/n4layer/session/sessionhandler"
	. "lite5gc/smf/n4layer/typedef"
	"lite5gc/smf/smfcontext/gctxt"
	"net"
	"sync"
)

// node N11消息接口
func init() {
	N11Msg = pfcpv1.NewServiceMsgChan()

	//go N11Msg.HandleRequestMsg(HandleRequestMsg)

}

var lock sync.Mutex

// smf sc发起业务处理,node注册协程处理
func HandleRequestMsg(cxt context.Context, msg pfcpv1.ServiceMsg) (*pfcpv1.ServiceMsg, error) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "msg.RemoteAdd.IP", msg.RemoteAdd.IP)
	//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "msg.Msg", msg.Msg.Cxt.PFCPParameters)
	// 判断节点是否存在，
	peerAddr := msg.RemoteAdd.IP.String()
	lock.Lock()
	v, err := GetNode(peerAddr)
	lock.Unlock()
	// 存在，且是active状态，进入转发处理
	if err == nil {
		v.Mu.Lock()
		state := v.State
		v.Mu.Unlock()
		if state == pfcpv1.NodeActive {
			// todo 启动多线程处理
			ok := v.ServiceMsgReveive(msg)
			if !ok {
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
					"node goroutine id %d,channel full, message dropped", cxt.Value(types.GoroutineID))
			}
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
				"node goroutine id %v,ServiceMsg(%v)", cxt.Value(types.GoroutineID), msg.Msg.ID)
			//res, err := sessionhandler.SmfSessionHandle(v.Node, msg)
			return nil, nil
		} else {
			// 	Deactivated 状态，缓存消息，已启动创建流程
			//缓存消息
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
				"node goroutine id %v,node Deactivated", cxt.Value(types.GoroutineID))
			ok := v.ServiceMsgReveiveBuff(msg)
			if !ok {
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
					"node goroutine id %d,channel full, message dropped", cxt.Value(types.GoroutineID))
				return nil, fmt.Errorf("channel full")
			}
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
				"node goroutine id %v,buffer ServiceMsg(%v)", cxt.Value(types.GoroutineID), msg.Msg.ID)
			//rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, msg.Msg)
		}
		return nil, nil
	}
	// 不存在，创建node，lock到创建完成，发送消息到节点内部缓存，节点本身启动协程负责启动setup流程，阻塞处理

	if err != nil { // 不存在，创建node，启动一个协程负责创建节点
		msg.RemoteAdd.Port = pfcp.Port
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "create new node ", err)
		node, err := CreateNode(cxt, msg.RemoteAdd)
		if err != nil {
			return nil, err
		}
		// 动态释放node router
		cxt = node.Cxt
		// 启动接收通道router
		id := utils.IpToUint32(node.PeerAddr.IP)
		// todo zoujun
		scid, ok := cxt.Value(types.SessSmfScCK).(uint32)
		if !ok {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "---no sc id found 1")
		} else {
			rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, nil, "---sc id found 1:%d", scid)
		}
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "main node goroutine id %v", cxt.Value(types.GoroutineID))
		cxt, err := n4Layer.startRouterAgent(cxt, id)
		if err != nil {
			return nil, fmt.Errorf("start RouterAgent fault")
		}
		// todo zoujun
		scid, ok = cxt.Value(types.SessSmfScCK).(uint32)
		if !ok {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "---no sc id found 2")
		} else {
			rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, nil, "---sc id found 2:%d", scid)
		}
		node.ScCxt = cxt
		// 添加到 node pool
		/*lock.Lock()
		err = AddNode(peerAddr, node)
		if err != nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "AddNode(%s), err:%s ", peerAddr, err)
			return nil, err
		}
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "AddNode success(%s)", peerAddr)
		lock.Unlock()*/
		//缓存消息
		tag := node.ServiceMsgReveiveBuff(msg)
		if tag == false {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
				"main node goroutine id %d,channel full, message dropped", cxt.Value(types.GoroutineID))
			return nil, fmt.Errorf("channel full")
		}
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"main node goroutine id %v,buffer ServiceMsg(%v)", cxt.Value(types.GoroutineID), msg.Msg.ID)
		return nil, nil
	}
	return nil, nil
}

/*func CreateNode(name *net.UDPAddr) (n *Node, err error) {
	nodeData, err := pfcpv1.CreateNode(name)
	if err != nil {
		return nil, err
	}
	nfsm, err := NewNodeFSM()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failed to create nodeProcFSM, err:%s", err)
	}
	nfsmHB, err := NewNodeProcFSM()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failed to create nodeProcFSM, err:%s", err)
	}
	n = &Node{Node: nodeData,
		NFsm: &NodeFSMs{NodeFsm: nfsm, NodeHBFSM: nfsmHB}}
	return n, nil
}*/
func CreateNode(cxt context.Context, name *net.UDPAddr) (n *Node, err error) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	nodeData, err := pfcpv1.CreateNode(cxt, name)
	if err != nil {
		return nil, err
	}
	nfsm, err := NewNodeFSM()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failed to create nodeProcFSM, err:%s", err)
	}
	nfsmHB, err := NewNodeProcFSM()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failed to create nodeProcFSM, err:%s", err)
	}
	n = &Node{Node: nodeData,
		NFsm: &NodeFSMs{NodeFsm: nfsm, NodeHBFSM: nfsmHB}}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Start creating node:%s", n.NodeID)

	lock.Lock()
	err = AddNode(name.IP.String(), n)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "add node(%s), err:%s ", name.IP, err)
		return nil, err
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "add node success(%s)", name.IP)
	lock.Unlock()
	// 启动node接收业务消息处理
	for i := 0; i < pfcpudp.HandleMsg_GOROUTINE_NUMBER; i++ {
		// 收消息,2个处理协程
		go handleReceiveScMsg(n)
	}

	// 启动节点创建流程
	go handleNodeCreateHandle(n)

	return n, nil
}

// node 节点处理协程，一个节点可以启动多个
func handleReceiveScMsg(n *Node) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	for {
		select {
		case <-n.Cxt.Done(): // 节点删除
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "node:%s,Release node goroutine", n.NodeID)
			return
		case msgCxt := <-n.ServiceMsgHandler():
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
				"service msg send, node:%s,len:%d", n.NodeID, len(n.ServiceMsgHandler()))
			_, err := sessionhandler.SmfSessionHandle(n.Node, msgCxt)
			if err != nil {
				// send response to smfsc
				ReleaseHandleReceiveScMsg(n, msgCxt)
			}
		}
	}
}
func handleReceiveScMsgBuff(n *Node) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	for {
		select {
		case msgCxt := <-n.ServiceMsgBuffHandler():
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
				"buff send, node:%s,len:%d", n.NodeID, len(n.ServiceMsgBuffHandler()))
			_, err := sessionhandler.SmfSessionHandle(n.Node, msgCxt)
			if err != nil {
				// send response to smfsc
				ReleaseHandleReceiveScMsg(n, msgCxt)
			}
		default:
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Buff send complete")
			return
		}
	}
}

func handleNodeCreateHandle(n *Node) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//节点内消息处理
	//	节点发送请求消息
	n.Mu.Lock()
	if n.State != pfcpv1.NodeActive {
		n.Mu.Unlock()
		//启动setup流程
		err := nodehandler.AssociationSetupProcedure(n)
		if err != nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
				"failed,node:%s,AssociationSetupRequest:%s", n.NodeID, err)
			return err
		}
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "node:%s,start Heartbeat", n.NodeID)
		//	todo 启动心跳检测
		go nodehandler.HeartbeatSetup(n)
	} else {
		n.Mu.Unlock()
	}

	// 收到AssociationSetup Response 后启动

	return nil
}

func ReleaseNodeRouter(cxt context.Context) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)

	msgRouter, ok := cxt.Value(types.MsgRouterCK).(*routeragent.MsgRouter)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, cxt, "no msg router found")
	}
	msgRouter.Deactivate()
}

// 释放缓存中的消息，发送失败响应
func ReleaseHandleReceiveScMsgBuff(n *Node) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	for {
		select {
		case msgCxt := <-n.ServiceMsgBuffHandler():
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
				"buff send, node:%s,len:%d", n.NodeID, len(n.ServiceMsgBuffHandler()))
			// 发送到N11 smfsc
			response := pfcpv1.ServiceMsg{}
			switch msgCxt.Msg.ID {
			case pfcp.PFCP_Session_Establishment_Request:
				response.Msg.ID = pfcp.PFCP_Session_Establishment_Response
			case pfcp.PFCP_Session_Modification_Request:
				response.Msg.ID = pfcp.PFCP_Session_Modification_Response
			case pfcp.PFCP_Session_Deletion_Request:
				response.Msg.ID = pfcp.PFCP_Session_Deletion_Response
			case pfcp.PFCP_Session_Report_Response:
				response.Msg.ID = pfcp.PFCP_Session_Report_Request
			}

			response.Msg.Cxt = msgCxt.Msg.Cxt
			response.Msg.Cxt.Cause = gctxt.Cause_No_established_PFCP_Association

			response.RemoteAdd = n.PeerAddr
			// 发送给 SmfSc
			scid, ok := n.ScCxt.Value(types.SessSmfScCK).(uint32)
			if !ok {
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n.ScCxt, "no sc id found")
				break
			}
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "sc id found:%d", scid)
			routeragent.SendIpcMessage(n.ScCxt, router.SmfScGR, scid, response)
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "SessionEstablishmentResponse:", response)
		default:
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Buff back send complete")
			return
		}
	}
}

// 异常情况发送响应给smfsc
func ReleaseHandleReceiveScMsg(n *Node, msgCxt pfcpv1.ServiceMsg) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	{
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"buff send, node:%s,msg:%+v", n.NodeID, msgCxt)
		// 发送到N11 smfsc
		response := pfcpv1.ServiceMsg{}
		switch msgCxt.Msg.ID {
		case pfcp.PFCP_Session_Establishment_Request:
			response.Msg.ID = pfcp.PFCP_Session_Establishment_Response
		case pfcp.PFCP_Session_Modification_Request:
			response.Msg.ID = pfcp.PFCP_Session_Modification_Response
		case pfcp.PFCP_Session_Deletion_Request:
			response.Msg.ID = pfcp.PFCP_Session_Deletion_Response
		case pfcp.PFCP_Session_Report_Response:
			response.Msg.ID = pfcp.PFCP_Session_Report_Request
		}

		if msgCxt.Msg.Cxt == nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
				"failed to PduCxt is nil,msg.Msg.Cxt(%p)", msgCxt.Msg.Cxt)
			return
		}
		response.Msg.Cxt = msgCxt.Msg.Cxt

		response.Msg.Cxt.Cause = gctxt.Cause_System_failure

		response.RemoteAdd = n.PeerAddr
		// 发送给 SmfSc
		scid, ok := n.ScCxt.Value(types.SessSmfScCK).(uint32)
		if !ok {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n.ScCxt, "no sc id found")
			return
		}
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "sc id found:%d", scid)
		routeragent.SendIpcMessage(n.ScCxt, router.SmfScGR, scid, response)
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "SessionEstablishmentResponse:", response)
	}
}
