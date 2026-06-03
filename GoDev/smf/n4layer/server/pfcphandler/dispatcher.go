package pfcphandler

import (
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/n4layer/node/nodehandler"
	. "lite5gc/smf/n4layer/typedef"
)

//const MODULE_ID = logger.PACKAGE_N4LAYER_MODULE_ID
// 处理上行消息
func NodeDispatch(peerIp string, msg pfcpv1.Message, res *pfcpv1.Message) error {
	//rlogger.FuncEntry(types.ModuleSmfN4, nil)
	n, err := GetNode(peerIp)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"The corresponding processing node does not exist(%s),err:%s", peerIp, err)
	}
	//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "MessageType %v,SEID %v,SequenceNumber %v",
	//	msg.Header.MessageType, msg.Header.SEID, msg.Header.SequenceNumber)
	switch msg.Header.MessageType {
	case pfcp.PFCP_Heartbeat_Request:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		//	receive request：smf <-- upf
		rMsg, ok := msg.Body.(*pfcp.HeartbeatRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		err := HandleHeartbeatRequest(n, rMsg)
		if err != nil {
			return err
		}
		return nil
	case pfcp.PFCP_Heartbeat_Response:
		//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
		//	"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.HeartbeatResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		err := HandleHeartbeatResponse(n, rMsg)
		if err != nil {
			return err
		}

		return nil
	case pfcp.PFCP_PFD_Management_Response:
		rMsg, ok := msg.Body.(*pfcp.PFCPPFDManagementResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		err := HandlePFDManagementResponse(n, rMsg)
		if err != nil {
			return err
		}

		return nil
	case pfcp.PFCP_Association_Setup_Response:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.PFCPAssociationSetupResponse)
		if !ok {
			return fmt.Errorf("type error")
			rlogger.FuncEntry(types.ModuleSmfN4, nil)
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		err := HandleAssociationSetupResponse(n, rMsg)
		if err != nil {
			rlogger.FuncEntry(types.ModuleSmfN4, nil)
			return err
		}
		rlogger.FuncEntry(types.ModuleSmfN4, nil)
		return nil
	case pfcp.PFCP_Association_Update_Request:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		//	receive request：smf <-- upf
		rMsg, ok := msg.Body.(*pfcp.PFCPAssociationUpdateRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		// 处理请求
		err := HandleAssociationUpdateRequest(n, rMsg)
		if err != nil {
			return err
		}
		// 发送响应消息
		msg := nodehandler.AssociationUpdateResponseCreate(n, rMsg)
		// 编码消息
		data, err := msg.Marshal()
		if err != nil {
			return fmt.Errorf("Pfcp msg marshal err %s", err)
		}
		// 发送消息到 udp server，使用UDP conn 的双工发送
		//msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
		//	RemoteAdd: n.PeerAddr}
		// 异步，发送到chan
		err = n.SendUdpMsg(data)
		if err != nil {
			return err
		}
		return nil
	case pfcp.PFCP_Association_Update_Response:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.PFCPAssociationUpdateResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		err := HandleAssociationUpdateResponse(n, rMsg)
		if err != nil {
			return err
		}

		return nil
	case pfcp.PFCP_Association_Release_Request:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.PFCPAssociationReleaseRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		// 处理请求
		err := HandleAssociationReleaseRequest(n, rMsg)
		if err != nil {
			return err
		}
		// 发送响应消息
		msg := nodehandler.AssociationReleaseResponseCreate(n, rMsg)
		// 编码消息
		data, err := msg.Marshal()
		if err != nil {
			return fmt.Errorf("Pfcp msg marshal err %s", err)
		}
		// 发送消息到 udp server，使用UDP conn 的双工发送
		//msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
		//	RemoteAdd: n.PeerAddr}
		// 异步，发送到chan
		err = n.SendUdpMsg(data)
		if err != nil {
			return err
		}
		//设置node的当前状态,使用下行消息的状态机
		n.NFsm.NodeFsm.Bfsm.Event(EventNodeReleaseAck, n)
		return nil
	case pfcp.PFCP_Association_Release_Response:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.PFCPAssociationReleaseResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		err := HandleAssociationReleaseResponse(n, rMsg)
		if err != nil {
			return err
		}
		return nil
		//4.4.3.4	N4 Report Procedure
		// receive request:smf <-- upf
	case pfcp.PFCP_Node_Report_Request:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.PFCPNodeReportRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		// 处理请求
		err := HandleNodeReportRequest(n, rMsg)
		if err != nil {
			return err
		}
		// 发送响应消息
		msg := nodehandler.NodeReportResponseCreate(n, rMsg)
		// 编码消息
		data, err := msg.Marshal()
		if err != nil {
			return fmt.Errorf("Pfcp msg marshal err %s", err)
		}
		// 发送消息到 udp server
		//msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
		//	RemoteAdd: n.PeerAddr}
		// 异步，发送到chan
		err = n.SendUdpMsg(data)
		if err != nil {
			return err
		}
		//设置node的当前状态,使用下行消息的状态机
		n.NFsm.NodeFsm.Bfsm.Event(EventNodeReportAck, n)
		return nil
	case pfcp.PFCP_Version_Not_Supported_Response:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.PFCPVersionNotSupportedResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		err := PFCPVersionNotSupportedResponse(n, rMsg)
		if err != nil {
			return err
		}
		return nil
	case pfcp.PFCP_Session_Establishment_Response:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.SessionEstablishmentResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		//rmsg.IE.Cause.CauseValue = pfcp.Cause_Request_accepted
		err := HandleSessionEstablishmentResponse(n, rMsg)
		if err != nil {
			return err
		}
		return nil
	case pfcp.PFCP_Session_Modification_Response:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.SessionModifyResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		err := HandleSessionModificationResponse(n, rMsg)
		if err != nil {
			return err
		}
		return nil
	case pfcp.PFCP_Session_Deletion_Response:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.SessionReleaseResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		err := HandleSessionDeletionResponse(n, rMsg)
		if err != nil {
			return err
		}
		return nil
	case pfcp.PFCP_Session_Report_Request:
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
			"NodeDispatch MessageType(%d)", msg.Header.MessageType)
		rMsg, ok := msg.Body.(*pfcp.SessionReportRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		msg.HeaderFillToMsg(&rMsg.PfcpHeader)
		// 处理请求
		err := HandleSessionReportRequest(n, rMsg)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("Unknown PFCP message type: %d", msg.Header.MessageType)

	}
	return nil
}
