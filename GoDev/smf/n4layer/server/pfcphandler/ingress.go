package pfcphandler

import (
	"fmt"
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/types"
	"lite5gc/smf/n4layer/session/create"
	"lite5gc/smf/smfcontext/gctxt"
	"strings"

	//"lite5gc/smf/n4layer/node"
	node "lite5gc/smf/n4layer/typedef"
)

// 收消息
func ReceivePfcpMsgHandle(b []byte) {
	//解析头

	// 根据消息类型进入分发，返回响应
}
func HandlePFDManagementResponse(n *node.Node, res *pfcp.PFCPPFDManagementResponse) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 节点消息匹配检查
	if res.PfcpHeader.SequenceNumber != n.SequenceNumber {
		return fmt.Errorf("response msg mismatch")
	}

	//设置node的当前状态,可以发送请求
	n.NFsm.NodeFsm.Bfsm.Event(node.EventNodePFDAck, n)
	//node.State = node.NFsm.NodeFsm.Bfsm.Current()
	return nil
}

func HandleAssociationSetupResponse(n *node.Node, res *pfcp.PFCPAssociationSetupResponse) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 节点消息匹配检查
	if res.PfcpHeader.SequenceNumber != n.SequenceNumber {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
			"HandleAssociationSetupResponse SequenceNumber(%d),n.SequenceNumber(%d)",
			res.PfcpHeader.SequenceNumber, n.SequenceNumber)
		//return fmt.Errorf("response msg mismatch")
	}
	// 处理响应消息，保存UP信息
	if res.IE.Cause != nil && res.IE.Cause.CauseValue != pfcp.Cause_Request_accepted {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
			"HandleAssociationSetupResponse cause(%v)", res.IE.Cause.CauseValue)
		// 对端响应错误，删除节点
		//设置node的当前状态
		n.NFsm.NodeFsm.Bfsm.SetState(node.StateNodeRelease)
		n.NFsm.NodeFsm.Bfsm.Event(node.EventNodeReleaseAck, n)
		return fmt.Errorf("cause(%v)", res.IE.Cause.CauseValue)
	}
	for _, v := range res.IE.UserPlaneIPResourceInformation {
		if v.ASSONI {
			dnn := v.NetworkInstance
			if strings.Contains(strings.ToLower(dnn), "n3") { // 包含n3
				if v.V4 {
					n.N3Ip = v.IPv4address
				}
				if v.TEIDRI != 0 {
					//partition
					n.TEIDRI = v.TEIDRI
					n.TEIDRange = v.TEIDRange
					rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
						"HandleAssociationSetupResponse UserPlaneIPResourceInformation(TEIDRI:%v,TEIDRange:%v)", n.TEIDRI, n.TEIDRange)
					//	启动teid分配
					if n.TEIDRI > 32 {
						rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
							"invalid value,HandleAssociationSetupResponse UserPlaneIPResourceInformation(TEIDRI:%v)", n.TEIDRI)
					}
					maxTeid := uint32(0xFFFFFFFF) >> (n.TEIDRI)
					//SmfTEID type
					SmfTEID := n.NodeID
					idmgr.GetInst().RegisterIDMgr(SmfTEID, maxTeid)
					idmgr.GetInst().ReserveID(SmfTEID, 0)
					idmgr.GetInst().ReserveID(SmfTEID, maxTeid)
					rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
						"HandleAssociationSetupResponse UserPlaneIPResourceInformation(TEID Type:%v,maxTeid:%v)", SmfTEID, maxTeid)
				}
			}
		}

	}

	//设置node的当前状态,可以发送请求
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
		"HandleAssociationSetupResponse fsm(%s)", n.NFsm.NodeFsm.Bfsm.Current())
	n.NFsm.NodeFsm.Bfsm.SetState(node.StateNodeStart)
	n.NFsm.NodeFsm.Bfsm.Event(node.EventNodeSetupAck, n)
	//node.State = node.NFsm.NodeFsm.Bfsm.Current()
	return nil
}

func HandleAssociationUpdateRequest(n *node.Node, res *pfcp.PFCPAssociationUpdateRequest) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 重置心跳定时器
	n.NTimer.T1Reset()
	//设置node的当前状态
	n.NFsm.NodeFsm.Bfsm.Event(node.EventNodeUpdate, n)

	return nil
}

func HandleAssociationUpdateResponse(n *node.Node, res *pfcp.PFCPAssociationUpdateResponse) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 与请求匹配检查
	if n.Request == nil {
		return fmt.Errorf("Illegal message")
	}
	if n.Request.Header.MessageType != pfcp.PFCP_Association_Update_Request {
		return fmt.Errorf("Illegal message")
	}
	n.Request = nil
	// 节点消息匹配检查
	if res.PfcpHeader.SequenceNumber != n.SequenceNumber {
		return fmt.Errorf("response msg mismatch")
	}

	//设置node的当前状态,可以发送请求
	n.NFsm.NodeFsm.Bfsm.Event(node.EventNodeUpdateAck, n)
	//node.State = node.NFsm.NodeFsm.Bfsm.Current()
	return nil
}

func HandleHeartbeatRequest(n *node.Node, req *pfcp.HeartbeatRequest) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	if n != nil {
		sn := n.Node
		n.NFsm.NodeHBFSM.Bfsm.Event(node.EventHeartbeatReqRecv, sn, req)
	}
	return nil
}

func HandleHeartbeatResponse(n *node.Node, res *pfcp.HeartbeatResponse) error {
	//rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 节点状态不变，重置心跳定时器
	if (n != nil) && (n.NTimer.T1RetransCancel != nil) {
		n.NTimer.T1RetransCancel()
	}
	// 节点消息匹配检查
	if res.PfcpHeader.SequenceNumber != n.SequenceNumber {
		return fmt.Errorf("response msg mismatch")
	}
	return nil
}

func HandleAssociationReleaseRequest(n *node.Node, req *pfcp.PFCPAssociationReleaseRequest) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 重置心跳定时器
	n.NTimer.T1Reset()
	//设置node的当前状态
	n.NFsm.NodeFsm.Bfsm.Event(node.EventNodeRelease, n)
	return nil
}

func HandleAssociationReleaseResponse(n *node.Node, req *pfcp.PFCPAssociationReleaseResponse) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 重置心跳定时器
	n.NTimer.T1Reset()
	//设置node的当前状态
	n.NFsm.NodeFsm.Bfsm.Event(node.EventNodeReleaseAck, n)
	return nil
}

func HandleNodeReportRequest(n *node.Node, req *pfcp.PFCPNodeReportRequest) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 重置心跳定时器
	n.NTimer.T1Reset()
	//设置node的当前状态
	n.NFsm.NodeFsm.Bfsm.Event(node.EventNodeReport, n)
	return nil
}

func PFCPVersionNotSupportedResponse(n *node.Node, req *pfcp.PFCPVersionNotSupportedResponse) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 重置心跳定时器
	n.NTimer.T1Reset()
	//设置node的当前状态
	n.NFsm.NodeFsm.Bfsm.Event(node.EventNodeVersionNotSupport, n)
	return nil
}

func HandleSessionEstablishmentResponse(n *node.Node, msg *pfcp.SessionEstablishmentResponse) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 节点状态不变，重置心跳定时器
	n.NTimer.T1Reset()

	// 匹配N4上下文
	seid := msg.PfcpHeader.SEID
	n4Cxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, &msg.PfcpHeader, "Failed to get N4 Context:%s", err)
		return err
	}
	// N4层处理响应消息,更新N4中的upf seid
	create.ResStoreSmfN4Context(msg, n4Cxt)
	scN4MsgData := &gctxt.ScN4MsgData{}
	scN4MsgData.SEID = n4Cxt.SEID
	scN4MsgData.Cause = gctxt.N4Cause(msg.IE.Cause.CauseValue)

	// 发送到N11口
	response := pfcpv1.ServiceMsg{}
	response.Msg.ID = pfcp.PFCPMSG(msg.PfcpHeader.MessageType)
	response.Msg.Cxt = scN4MsgData // 从上下文中获得
	response.RemoteAdd = n.PeerAddr
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "SessionEstablishmentResponse:")
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "SessionEstablishmentResponse:Msg.ID", msg.PfcpHeader.MessageType)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "SessionEstablishmentResponse:pduCxt.Cause", scN4MsgData.Cause.Error())
	// 发送给 SmfSc
	scid, ok := n.ScCxt.Value(types.SessSmfScCK).(uint32)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n.ScCxt, "no sc id found")
		return fmt.Errorf("no sc id found")
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "sc id found:%d", scid)
	routeragent.SendIpcMessage(n.ScCxt, router.SmfScGR, scid, response)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "SessionEstablishmentResponse:", response)
	return nil
}

func HandleSessionModificationResponse(n *node.Node, msg *pfcp.SessionModifyResponse) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 节点状态不变，重置心跳定时器
	n.NTimer.T1Reset()

	// 匹配N4上下文
	seid := msg.PfcpHeader.SEID
	n4Cxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, &msg.PfcpHeader, "Failed to get N4 Context:%s", err)
		return err
	}
	// N4层处理响应消息

	// 适配为 SmfSc 消息
	scN4MsgData := &gctxt.ScN4MsgData{}
	scN4MsgData.SEID = n4Cxt.SEID
	scN4MsgData.Cause = gctxt.N4Cause(msg.IE.Cause.CauseValue)

	// 发送到N11口
	response := pfcpv1.ServiceMsg{}
	response.Msg.ID = pfcp.PFCPMSG(msg.PfcpHeader.MessageType)
	response.Msg.Cxt = scN4MsgData // 从上下文中获得
	response.RemoteAdd = n.PeerAddr

	// 发送给 SmfSc
	scid, ok := n.ScCxt.Value(types.SessSmfScCK).(uint32)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n.ScCxt, "no sc id found")
		return fmt.Errorf("no sc id found")
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "sc id found:%d", scid)
	routeragent.SendIpcMessage(n.ScCxt, router.SmfScGR, scid, response)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "SessionModificationResponse:", response)
	return nil
}

func HandleSessionDeletionResponse(n *node.Node, msg *pfcp.SessionReleaseResponse) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 节点状态不变，重置心跳定时器
	n.NTimer.T1Reset()
	// 匹配N4上下文
	seid := msg.PfcpHeader.SEID
	n4Cxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, &msg.PfcpHeader, "Failed to get N4 Context:%s", err)
		return err
	}
	// N4层处理响应消息
	scN4MsgData := &gctxt.ScN4MsgData{}
	scN4MsgData.SEID = n4Cxt.SEID
	scN4MsgData.Cause = gctxt.N4Cause(msg.IE.Cause.CauseValue)

	// 发送到N11口
	response := pfcpv1.ServiceMsg{}
	response.Msg.ID = pfcp.PFCPMSG(msg.PfcpHeader.MessageType)
	response.Msg.Cxt = scN4MsgData // 从上下文中获得
	response.RemoteAdd = n.PeerAddr

	// 发送给 SmfSc
	scid, ok := n.ScCxt.Value(types.SessSmfScCK).(uint32)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n.ScCxt, "no sc id found")
		return fmt.Errorf("no sc id found")
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "sc id found:%d", scid)
	routeragent.SendIpcMessage(n.ScCxt, router.SmfScGR, scid, response)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "HandleSessionDeletionResponse:", response)
	// 3 在upf成功响应后释放本地 N4 context
	err = gctxt.DeleteN4Context(gctxt.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failure to delete N4 Context:%s", err)
		return fmt.Errorf("Failure to delete N4 Context:%s", err)
	}
	return nil
}

func HandleSessionReportRequest(n *node.Node, msg *pfcp.SessionReportRequest) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 节点状态不变，重置心跳定时器
	n.NTimer.T1Reset()
	// N4层处理响应消息
	res := &pfcp.SessionReportResponse{}
	err := create.SessionReportRequest(*msg, res)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n.ScCxt, err)
		return err
	}
	// smf received Data Notification Msg from Upf,
	// then send Namf_Communication_N1N2MessageTransfer to AMF
	n4Cxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(msg.PfcpHeader.SEID))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n.ScCxt, "no n4 cxt found")
		return err
	}
	/*err = smsender.SendN1N2MsgTransfer(n4Cxt, n.ScCxt)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, &msg.PfcpHeader, "failed to send n1n2 transfer msg")
		return fmt.Errorf("failed to send n1n2 transfer msg")
	}*/
	// 适配为 SmfSc 消息
	//if n4Cxt.PduCxt == nil {
	//	rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n.ScCxt, "session ctxt is nil")
	//	return fmt.Errorf("session ctxt is nil")
	//}

	scN4MsgData := &gctxt.ScN4MsgData{}
	scN4MsgData.SEID = n4Cxt.SEID
	scN4MsgData.Cause = gctxt.N4Cause(pfcp.Cause_Request_accepted)

	// 发送到N11口
	serviceMsg := pfcpv1.ServiceMsg{}
	serviceMsg.Msg.ID = pfcp.PFCPMSG(msg.PfcpHeader.MessageType)
	serviceMsg.Msg.Cxt = scN4MsgData // 从上下文中获得
	serviceMsg.RemoteAdd = n.PeerAddr

	// 发送给 SmfSc
	scid, ok := n.ScCxt.Value(types.SessSmfScCK).(uint32)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n.ScCxt, "no sc id found")
		return fmt.Errorf("no sc id found")
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "sc id found:%d", scid)
	routeragent.SendIpcMessage(n.ScCxt, router.SmfScGR, scid, serviceMsg)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, n.ScCxt, "SessionReportRequest:", serviceMsg)

	// 发送响应到upf
	resMsg := &pfcpv1.Message{}

	// 构造待编码消息
	// Encoding message filling
	resMsg.HeaderSet(res.PfcpHeader)
	resMsg.BodySet(res)
	// 编码消息
	data, err := resMsg.Marshal()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg marshal err %s", err)
		return err
	}

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, resMsg.String())
	err = n.SendUdpMsg(data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to send udp")
	}
	return nil
}
