package nodehandler

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/pfcpudp"
	pfcpv1 "lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	. "lite5gc/smf/n4layer/typedef"
	"time"
)

//6.2.6	PFCP Association Setup Procedure
func AssociationSetupProcedure(node *Node) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 使用本节的的server
	//todo 初始server值
	//pfcpudp.PfcpServer.PeerAddr.Port = configure.SmfConf.N4Conf.UPFPort
	node.Server = pfcpudp.PfcpServer
	// 构造setup消息
	msg := PFCPAssociationSetupRequestCreate(node)
	if msg == nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "pfcp message creation failed")
		return fmt.Errorf("pfcp message creation failed")
	}
	// 编码消息
	data, err := msg.Marshal()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "pfcp message encoding failed,err %s", err)
		return fmt.Errorf("pfcp msg marshal err %s", err)
	}
	// 发送消息到 udp server，使用UDP conn 的双工发送
	/*msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
		RemoteAdd: node.PeerAddr}
	// 异步，发送到chan
	ok := node.Server.SendMsgTo(msgCxt)
	if !ok {
		return fmt.Errorf("Channel full")
	}*/
	// 启动超时重传
	go TimeoutRetransmission(node, data)
	/*err = TimeoutRetransmission(node, data)
	if err != nil {
		//	超时失败
		return err
	}*/
	err = node.SendUdpMsg(data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "pfcp message sending failed,err %s", err)
		return err
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
		"SendN4UdpMsg to Peer : <%s>--><%s>: %#x\n", node.Server.UdpConn.LocalAddr(), node.PeerAddr, data)

	// 同步，等待响应
	//node.Wait.Add(1)
	//node.Wait.Wait()

	return nil
}

//The SMF initiates the N4 Association Setup procedure to request to setup
// an N4 association towards a UPF prior to establishing a first N4 session on this UPF.
// TS 29.244
//7.4.4.1	PFCP Association Setup Request
func PFCPAssociationSetupRequestCreate(node *Node) *pfcpv1.Message {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	node.SequenceNumber += 1
	msg := &pfcpv1.Message{}
	request := &pfcp.PFCPAssociationSetupRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Association_Setup_Request,
			Length:         0, // 编码后填充
			SequenceNumber: node.SequenceNumber},
	}

	request.IE = &pfcp.IEPFCPAssociationSetupRequest{}
	//Mandatory IE
	request.IE.NodeID = &pfcp.IENodeID{
		NodeIDType:  pfcp.NodeIDType_IPv4_address,
		NodeIDvalue: node.Server.LocalAddr.IP.To4(), // smf ip
	}
	//Mandatory IE
	request.IE.RecoveryTimeStamp = &pfcp.IERecoveryTimeStamp{
		// 1556588833 = nt64(0xe07227a1(wireshark) - 0x83aa7e80(1900到1970的秒差))
		RecoveryTimeStamp: time.Now(),
	}
	// 构造待编码消息
	msg.Header.Version = request.PfcpHeader.Version
	msg.Header.MessageType = request.PfcpHeader.MessageType
	msg.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	// 编码后填充
	msg.Header.Length = 0
	msg.Body = request
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}

// TS 29.244
//6.2.2	Heartbeat Procedure
func HeartbeatSetup(n *Node) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//1、在节点创建后启动3s定时器
	t1 := time.NewTimer(pfcpv1.T1)

	cxt, cancel := context.WithCancel(context.Background())
	n.NTimer.T1Cancel = cancel // stop timer
	//reSetcxt, reset := context.WithCancel(context.Background())
	reSetcxt := make(chan struct{}, 1)
	reset := func() {
		reSetcxt <- struct{}{}
	}
	// reset timer
	n.NTimer.T1Reset = reset

	n.NTimer.T1 = t1

	for {
		select {
		case <-reSetcxt:
			//3、收到节点内的任何消息（除节点释放消息），重置定时器
			t1.Reset(pfcpv1.T1)
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "node:%s,reset Heartbeat", n.NodeID)
			//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "node:%s,reset Heartbeat timer ", n.NodeID, t1)

		case <-cxt.Done(): // 关闭chan，不阻塞
			//4、释放节点后，关闭心跳定时器
			t1.Stop()
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "node:%s,stop Heartbeat", n.NodeID)
			return nil
		case <-t1.C:
			//	发送心跳消息,一个协程处理发送处理
			go sendHeartbeatMsg(n)
			t1.Reset(pfcpv1.T1)

		}
	}
	/*tr := time.AfterFunc(pfcpv1.T1, func() { // f 是非阻塞调用
		//2、3s到发送heartbeat消息
		// 创建消息
		msg := HeartbeatRequestCreate(n)
		// 编码发送
		data, err := msg.Marshal()
		if err != nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg marshal err %s", err)
		}
		// 发送消息到 udp server，使用UDP conn 的双工发送
		msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
			RemoteAdd: n.PeerAddr}
		// 异步，发送到chan
		ok := n.Server.SendMsgTo(msgCxt)
		if !ok {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Channel full")
		}
		// 启动响应超时重传
		err = TimeoutRetransmission(n, data)
		if err != nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg Timeout retransmission err: %s", err)
		}
		//设置流程的当前状态

	})
	*/
	return nil
}

//7.4.2.1	Heartbeat Request
func HeartbeatRequestCreate(node *Node) *pfcpv1.Message {
	//rlogger.FuncEntry(types.ModuleSmfN4, nil)
	node.SequenceNumber += 1
	msg := &pfcpv1.Message{}
	request := &pfcp.HeartbeatRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Heartbeat_Request,
			Length:         0, // 编码后填充
			SequenceNumber: node.SequenceNumber},
	}
	//Mandatory IE
	request.IE = &pfcp.IERecoveryTimeStamp{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Recovery_Time_Stamp,
			Length: 4},
		RecoveryTimeStamp: time.Now(),
	}

	// 构造待编码消息
	// Encoding message filling
	msg.HeaderSet(request.PfcpHeader)
	msg.BodySet(request)

	//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}

func sendHeartbeatMsg(n *Node) { // f 是非阻塞调用
	//rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//2、3s到发送heartbeat消息
	// 创建消息
	msg := HeartbeatRequestCreate(n)
	// 编码发送
	data, err := msg.Marshal()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg marshal err %s", err)
	}
	// 发送消息到 udp server，使用UDP conn 的双工发送
	/*msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
		RemoteAdd: n.PeerAddr}
	// 异步，发送到chan
	ok := n.Server.SendMsgTo(msgCxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Channel full")
	}*/
	err = n.SendUdpMsg(data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg send err %s", err)
		return
	}
	// 启动响应超时重传
	err = HeartbeatTimeoutRetransmission(n, data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg Timeout retransmission err: %s", err)
	}
	//设置流程的当前状态

}
func HeartbeatTimeoutRetransmission(n *Node, data []byte) error {
	//rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//1、发送请求后，启动超时定时器
	cxt, cancel := context.WithCancel(context.Background())
	n.NTimer.T1RetransCancel = cancel
	t2 := time.NewTimer(pfcpv1.T2)
	var RetrCount = 0

	for {
		select {
		case <-cxt.Done():
			//2、收到响应消息，关闭定时器
			t2.Stop()
			//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "turn off retransmission,Heartbeat")
			return nil
		case <-t2.C:
			if RetrCount == pfcpv1.MaxRetransT2 {
				//3、超时后，重发请求，发送3次，无响应，上报失败响应
				t2.Stop()
				// 对端没有响应，关闭心跳发送，删除节点
				//设置node的当前状态
				n.NFsm.NodeFsm.Bfsm.SetState(StateNodeRelease)
				n.NFsm.NodeFsm.Bfsm.Event(EventNodeReleaseAck, n)
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "no response from peer,Heartbeat", n.PeerAddr.String())
				return fmt.Errorf("timeout")
			}
			t2.Reset(pfcpv1.T2)
			{
				// 发送消息到 udp server，使用UDP conn 的双工发送
				/*msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
					RemoteAdd: n.PeerAddr}
				// 异步，发送到chan
				ok := n.Server.SendMsgTo(msgCxt)
				if !ok {
					rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Channel full")
				}*/
				err := n.SendUdpMsg(data)
				if err != nil {
					rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "message sending failed from Heartbeat:%s,peer(%s)", err, n.PeerAddr)
					return err
				}
				RetrCount += 1
				rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
					"SendN4UdpMsg to Peer : <%s>--><%s>: %#x\n", n.Server.UdpConn.LocalAddr(), n.PeerAddr, data)
			}
		}
	}

	return nil
}

//AssociationSetupRequest
func TimeoutRetransmission(n *Node, data []byte) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//1、发送请求后，启动超时定时器
	cxt, cancel := context.WithCancel(context.Background())
	n.NTimer.T2Cancel = cancel
	t2 := time.NewTimer(pfcpv1.T2)
	var RetrCount = 0

	for {
		select {
		case <-cxt.Done():
			//2、收到响应消息，关闭定时器
			t2.Stop()
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "turn off retransmission")
			return nil
		case <-t2.C:
			if RetrCount == pfcpv1.MaxRetransT2 {
				//3、超时后，重发请求，发送3次，无响应，上报失败响应
				t2.Stop()
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "no response from peer", n.PeerAddr.String())
				//设置node的当前状态
				n.NFsm.NodeFsm.Bfsm.SetState(StateNodeRelease)
				n.NFsm.NodeFsm.Bfsm.Event(EventNodeReleaseAck, n)
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "no response from peer,Heartbeat", n.PeerAddr.String())
				return fmt.Errorf("timeout")
			}
			t2.Reset(pfcpv1.T2)
			{
				/*// 发送消息到 udp server，使用UDP conn 的双工发送
				msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
					RemoteAdd: n.PeerAddr}
				// 异步，发送到chan
				ok := n.Server.SendMsgTo(msgCxt)
				if !ok {
					rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Channel full")
				}*/
				err := n.SendUdpMsg(data)
				if err != nil {
					rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "message sending failed:%s,peer(%s)", err, n.PeerAddr)
					return err
				}
				RetrCount += 1
				rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
					"SendN4UdpMsg to Peer : <%s>--><%s>: %#x\n", n.Server.UdpConn.LocalAddr(), n.PeerAddr, data)
			}
		}
	}

	return nil
}

//7.4.2.2	Heartbeat Response
func HeartbeatResponseCreate(req *pfcp.HeartbeatRequest) *pfcpv1.Message {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	msg := &pfcpv1.Message{}
	response := &pfcp.HeartbeatResponse{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Heartbeat_Response,
			Length:         0, // 编码后填充
			SequenceNumber: req.PfcpHeader.SequenceNumber},
	}
	//Mandatory IE
	response.IE = &pfcp.IERecoveryTimeStamp{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Recovery_Time_Stamp,
			Length: 4},
		RecoveryTimeStamp: req.IE.RecoveryTimeStamp,
	}

	// 构造待编码消息
	// Encoding message filling
	msg.HeaderSet(response.PfcpHeader)
	msg.BodySet(response)

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}

//3GPP TS 23.502 V15.4.0 (2018-12)
//4.4.3.2	N4 Association Update Procedure
// smf side
// send request:smf --> upf
//Message Type: PFCP Association Update Request (7)
func SendAssociationUpdateRequest(n *Node) { // f 是非阻塞调用
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//检查流程的当前状态
	if StateNodeActive != n.NFsm.NodeFsm.Bfsm.Current() {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
			"Current status(%s) message is illegal", n.NFsm.NodeFsm.Bfsm.Current())
		return
	}
	// 创建消息
	msg := AssociationUpdateRequestCreate(n)
	// 编码发送
	data, err := msg.Marshal()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "pfcp msg marshal err %s", err)
	}
	/*// 发送消息到 udp server，使用UDP conn 的双工发送
	msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
		RemoteAdd: n.PeerAddr}
	// 异步，发送到chan
	ok := n.Server.SendMsgTo(msgCxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Channel full")
	}*/
	err = n.SendUdpMsg(data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "pfcp msg send err %s", err)
		return
	}
	// 用于响应的匹配
	n.Request = msg
	// 启动响应超时重传
	err = TimeoutRetransmission(n, data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "pfcp msg Timeout retransmission err: %s", err)
	}
}

//7.4.4.3	PFCP Association Update Request
func AssociationUpdateRequestCreate(n *Node) *pfcpv1.Message {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	n.SequenceNumber += 1
	msg := &pfcpv1.Message{}
	request := &pfcp.PFCPAssociationUpdateRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Association_Update_Request,
			Length:         0, // 编码后填充
			SequenceNumber: n.SequenceNumber},
	}
	//Mandatory IE
	request.IE = &pfcp.IEPFCPAssociationUpdateRequest{}
	request.IE.NodeID = &pfcp.IENodeID{
		NodeIDType:  0,
		NodeIDvalue: n.Server.LocalAddr.IP.To4(), // smf ip
	}

	request.IE.CPFunctionFeatures = &pfcp.IECPFunctionFeatures{
		SupportedFeatures: pfcp.LOAD,
	}

	request.IE.IeFlags.Set(pfcp.IE_CP_Function_Features)

	// 构造待编码消息
	// Encoding message filling
	msg.HeaderSet(request.PfcpHeader)
	msg.BodySet(request)

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}

//4.4.3.3	N4 Association Release Procedure
// smf side
// send request:smf --> upf
//Message Type: PFCP Association Release Request (9)
func SendAssociationReleaseRequest(n *Node) { // f 是非阻塞调用
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//检查流程的当前状态
	if StateNodeActive != n.NFsm.NodeFsm.Bfsm.Current() {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
			"Current status(%s) message is illegal", n.NFsm.NodeFsm.Bfsm.Current())
		return
	}
	// 创建消息
	msg := AssociationReleaseRequestCreate(n)
	// 编码发送
	data, err := msg.Marshal()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg marshal err %s", err)
	}
	/*// 发送消息到 udp server，使用UDP conn 的双工发送
	msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
		RemoteAdd: n.PeerAddr}
	// 异步，发送到chan
	ok := n.Server.SendMsgTo(msgCxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Channel full")
	}*/
	err = n.SendUdpMsg(data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg send err %s", err)
		return
	}
	// 设置状态机
	n.NFsm.NodeFsm.Bfsm.Event(EventNodeRelease, n)
	// 用于响应的匹配
	n.Request = msg
	// 启动响应超时重传
	err = TimeoutRetransmission(n, data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg Timeout retransmission err: %s", err)
	}
}

//7.4.4.4	PFCP Association Update Response
func AssociationUpdateResponseCreate(n *Node, req *pfcp.PFCPAssociationUpdateRequest) *pfcpv1.Message {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	msg := &pfcpv1.Message{}
	response := &pfcp.PFCPAssociationUpdateResponse{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Association_Update_Response,
			Length:         0, // 编码后填充
			SequenceNumber: req.PfcpHeader.SequenceNumber},
	}
	//Mandatory IE
	response.IE = &pfcp.IEPFCPAssociationUpdateResponse{}
	response.IE.NodeID = &pfcp.IENodeID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Node_ID,
			Length: 5,
		},
		NodeIDType:  0,
		NodeIDvalue: n.Server.LocalAddr.IP.To4(),
	}
	response.IE.Cause = &pfcp.IECause{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Cause,
			Length: 1,
		},
		CauseValue: 1,
	}

	// 构造待编码消息
	// Encoding message filling
	msg.HeaderSet(response.PfcpHeader)
	msg.BodySet(response)

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}

//7.4.4.5	PFCP Association Release Request
func AssociationReleaseRequestCreate(n *Node) *pfcpv1.Message {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	n.SequenceNumber += 1
	msg := &pfcpv1.Message{}
	response := &pfcp.PFCPAssociationReleaseRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Association_Release_Request,
			Length:         0, // 编码后填充
			SequenceNumber: n.SequenceNumber},
	}
	//Mandatory IE
	response.IE = &pfcp.IEGroupedPFCPAssociationReleaseRequest{}
	response.IE.NodeID = &pfcp.IENodeID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Node_ID,
			Length: 5,
		},
		NodeIDType:  0,
		NodeIDvalue: n.Server.LocalAddr.IP.To4(),
	}

	// 构造待编码消息
	// Encoding message filling
	msg.HeaderSet(response.PfcpHeader)
	msg.BodySet(response)

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}

//7.4.4.6	PFCP Association Release Response
func AssociationReleaseResponseCreate(n *Node, req *pfcp.PFCPAssociationReleaseRequest) *pfcpv1.Message {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	msg := &pfcpv1.Message{}
	response := &pfcp.PFCPAssociationReleaseResponse{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Association_Release_Response,
			Length:         0, // 编码后填充
			SequenceNumber: req.PfcpHeader.SequenceNumber},
	}
	//Mandatory IE
	response.IE = &pfcp.IEPFCPAssociationReleaseResponse{}
	response.IE.NodeID = &pfcp.IENodeID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Node_ID,
			Length: 5,
		},
		NodeIDType:  0,
		NodeIDvalue: n.Server.LocalAddr.IP.To4(),
	}
	response.IE.Cause = &pfcp.IECause{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Cause,
			Length: 1,
		},
		CauseValue: 1,
	}

	// 构造待编码消息
	// Encoding message filling
	msg.HeaderSet(response.PfcpHeader)
	msg.BodySet(response)

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}

//4.4.3.4	N4 Report Procedure
// receive request:smf <-- upf
//7.4.5	PFCP Node Report Procedure
//7.4.5.2	PFCP Node Report Response
func NodeReportResponseCreate(n *Node, req *pfcp.PFCPNodeReportRequest) *pfcpv1.Message {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	msg := &pfcpv1.Message{}
	response := &pfcp.PFCPNodeReportResponse{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Node_Report_Response,
			Length:         0, // 编码后填充
			SequenceNumber: req.PfcpHeader.SequenceNumber},
	}
	//Mandatory IE
	response.IE = &pfcp.IEPFCPNodeReportResponse{}
	response.IE.NodeID = &pfcp.IENodeID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Node_ID,
			Length: 5,
		},
		NodeIDType:  0,
		NodeIDvalue: n.Server.LocalAddr.IP.To4(),
	}
	response.IE.Cause = &pfcp.IECause{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Cause,
			Length: 1,
		},
		CauseValue: 1,
	}

	// 构造待编码消息
	// Encoding message filling
	msg.HeaderSet(response.PfcpHeader)
	msg.BodySet(response)

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}

//4.4.3.5	N4 PFD management Procedure
// send request:smf --> upf
//7.4.3	PFCP PFD Management
////7.4.3.1	PFCP PFD Management Request
func SendPFDManagementRequest(n *Node) { // f 是非阻塞调用
	rlogger.FuncEntry(types.ModuleSmfN4, nil)

	//检查流程的当前状态
	if StateNodeActive != n.NFsm.NodeFsm.Bfsm.Current() {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
			"Current status(%s) message is illegal", n.NFsm.NodeFsm.Bfsm.Current())
		return
	}
	// 创建消息
	msg := PFDManagementRequestCreate(n)
	// 编码发送
	data, err := msg.Marshal()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg marshal err %s", err)
	}
	// 发送消息到 udp server，使用UDP conn 的双工发送
	/*msgCxt := pfcpudp.MsgCxt{Msgbuf: data,
		RemoteAdd: n.PeerAddr}
	// 异步，发送到chan
	ok := n.Server.SendMsgTo(msgCxt)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Channel full")
	}*/
	err = n.SendUdpMsg(data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg send err %s", err)
		return
	}
	// 设置状态机
	n.NFsm.NodeFsm.Bfsm.Event(EventNodePFD, n)
	// 用于响应的匹配
	n.Request = msg
	// 启动响应超时重传
	err = TimeoutRetransmission(n, data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Pfcp msg Timeout retransmission err: %s", err)
	}
}

//7.4.3.1	PFCP PFD Management Request
func PFDManagementRequestCreate(n *Node) *pfcpv1.Message {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)

	n.SequenceNumber += 1
	msg := &pfcpv1.Message{}
	request := &pfcp.PFCPPFDManagementRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_PFD_Management_Request,
			Length:         0, // 编码后填充
			SequenceNumber: n.SequenceNumber},
	}
	//Mandatory IE
	request.IE = &pfcp.IEPFCPPFDManagementRequest{}
	request.IE.ApplicationIDsPFDs = &pfcp.IEApplicationIDsPFDs{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Application_IDs_PFDs,
			Length: 5,
		},
	}

	request.IE.ApplicationIDsPFDs.ApplicationID = &pfcp.IEApplicationID{} // ipv4
	request.IE.ApplicationIDsPFDs.ApplicationID.ApplicationIdentifier = []byte("app1")

	request.IE.ApplicationIDsPFDs.IeFlags.Set(pfcp.IE_PFD_context)
	request.IE.ApplicationIDsPFDs.PFD = &pfcp.IEPFDContext{}
	request.IE.ApplicationIDsPFDs.PFD.PFDContexts = &pfcp.IEPFDContents{}
	request.IE.ApplicationIDsPFDs.PFD.PFDContexts.FD = true
	request.IE.ApplicationIDsPFDs.PFD.PFDContexts.LengthofFlowDes = 5
	request.IE.ApplicationIDsPFDs.PFD.PFDContexts.FlowDescription = "flow1"
	// 构造待编码消息
	// Encoding message filling
	msg.HeaderSet(request.PfcpHeader)
	msg.BodySet(request)

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}

//7.4.4.7	PFCP Version Not Supported Response
func VersionNotSupportedResponseCreate(n *Node, seq uint32) *pfcpv1.Message {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)

	msg := &pfcpv1.Message{}
	response := &pfcp.PFCPVersionNotSupportedResponse{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Version_Not_Supported_Response,
			Length:         0, // 编码后填充
			SequenceNumber: seq},
	}
	//Mandatory IE
	response.IE = &pfcp.IEPFCPVersionNotSupportedResponse{}

	// 构造待编码消息
	// Encoding message filling
	msg.HeaderSet(response.PfcpHeader)
	msg.BodySet(response)

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.String())

	return msg
}
