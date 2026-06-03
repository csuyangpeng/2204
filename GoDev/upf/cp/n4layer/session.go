package n4layer

import (
	"container/list"
	"errors"
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/utils"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types/configure"

	"lite5gc/upf/context/n4context"
	"lite5gc/upf/context/pdrcontext"
	"lite5gc/upf/cp/pdr"
	"lite5gc/upf/metrics"
	"net"
	"net/rpc/jsonrpc"
	"strconv"
)

// 3GPP TS 29.244 V15.3.0 (2018-09)
// N4 消息
type N4Msg struct {
}

// 注册N4请求的响应
// N4 Session Release Request
// RPC 中方法名作为消息类型标识，第一个参数是消息的内容，第二个参数是需要返回的响应消息内容。
func (s *N4Msg) SessionReleaseRequest(req pfcp.SessionReleaseRequest,
	res *pfcp.SessionReleaseResponse) error {
	rlogger.FuncEntry(moduleTag, &req.PfcpHeader)

	res.PfcpHeader.Version = pfcp.Version
	res.PfcpHeader.MessageType = pfcp.PFCP_Session_Deletion_Response
	res.PfcpHeader.SEID = 0
	res.PfcpHeader.SequenceNumber = req.PfcpHeader.SequenceNumber
	// 读取当前上下文
	n4Cxt, err := n4context.GetN4Context(n4context.N4SessionIDKey(req.PfcpHeader.SEID))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, &req.PfcpHeader, "Failed to get N4 Context:%s", err)
		res.IE.Cause.CauseValue = pfcp.Cause_System_failure
		return nil
	}

	// N4 Session Release Response
	if req.PfcpHeader.MessageType != pfcp.PFCP_Session_Deletion_Request {
		//return errors.New("Session Release Request message type error: " +
		//	strconv.Itoa(int(req.PfcpHeader.MessageType)))
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Session Release Request message type error: %d", req.PfcpHeader.MessageType)
		res.IE.Cause.CauseValue = pfcp.Cause_Service_not_supported
		return nil
	}

	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "Session Release Request:")
	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "  .MessageType    %v", req.PfcpHeader.MessageType)
	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "  .SEID           %v", req.PfcpHeader.SEID)
	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "  .SequenceNumber %v", req.PfcpHeader.SequenceNumber)

	// todo meter release
	err = metrics.SessionMeterRelease(n4Cxt.MetricItems)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "release upf context Counter failed:%s", err)
		//return nil
	}
	err = metrics.SessionMeterRelease(n4Cxt.MetricItemsSnapshot)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "release upf context Counter Snapshot failed:%s", err)
		//return nil
	}

	//3.	The UPF identifies the N4 session context to be removed by the N4 Session ID and removes the whole session context.
	//The UPF responds with an N4 session release response message containing any information that the UPF has to provide to the SMF.
	// 释放会话,删除上下文
	err = n4context.DeleteN4Context(n4context.N4SessionIDKey(n4Cxt.SEID))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Failed to Delete N4 Context:%s", err)
		res.IE.Cause.CauseValue = pfcp.Cause_Session_context_not_found
	} else {
		res.IE.Cause.CauseValue = pfcp.Cause_Request_accepted
	}
	// delete PDR Table
	err = pdr.DeleteMatchPDRsTable(n4Cxt) // 1.1版本更新  DeleteMatchPDRsTable(n4Cxt) // 1.0 版本
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Failed to Delete MatchPDRs Table:%s", err)
		res.IE.Cause.CauseValue = pfcp.Cause_Session_context_not_found
	} else {
		res.IE.Cause.CauseValue = pfcp.Cause_Request_accepted
	}
	res.PfcpHeader.Version = pfcp.Version
	res.PfcpHeader.MPFlag = 0
	res.PfcpHeader.SFlag = 1

	res.PfcpHeader.MessageType = pfcp.PFCP_Session_Deletion_Response
	res.PfcpHeader.SEID = n4Cxt.SmfSEID.SEID
	res.PfcpHeader.SequenceNumber = req.PfcpHeader.SequenceNumber

	//res.IE.Cause.Set()
	//res.IE.Cause.CauseValue = pfcp.Cause_Request_accepted

	return nil
}

// N4 Session Establishment Request
func (s *N4Msg) SessionEstablishmentRequest(req pfcp.SessionEstablishmentRequest,
	res *pfcp.SessionEstablishmentResponse) error {
	rlogger.FuncEntry(moduleTag, &req.PfcpHeader)
	// 收到请求消息处理
	if req.PfcpHeader.MessageType != pfcp.PFCP_Session_Establishment_Request {
		rlogger.Trace(moduleTag, rlogger.ERROR, &req.PfcpHeader, "Message type error,Correct value 50(Session Establishment), Error value is %d", req.PfcpHeader.MessageType)
		return errors.New("Session Establishment Request message type error: " + strconv.Itoa(int(req.PfcpHeader.MessageType)))
	}

	rlogger.Trace(moduleTag, rlogger.INFO, &req.PfcpHeader, "Session Establishment Request:(MessageType %v,init SEID %v,SequenceNumber %v,smf SEID %v)",
		req.PfcpHeader.MessageType, req.PfcpHeader.SEID, req.PfcpHeader.SequenceNumber, req.IE.CPFSEID.SEID)

	// 添加UPF上下文
	// key生成UPF本地的SEID
	// 保存SMF的SEID
	res.IE.Cause.CauseValue = pfcp.Cause_Request_accepted
	ufpCxt := StoreUpfN4Context(req)
	if ufpCxt == nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, &req.PfcpHeader, "Failure of parameter validity check parameter.")
		res.IE.Cause.CauseValue = pfcp.Cause_System_failure
		ufpCxt = &n4context.N4SessionContext{}
	}
	// 发送原因为“成功”的N4会话建立响应;
	// N4 Session Establishment Response
	// 返回响应消息构造
	res.PfcpHeader.Version = pfcp.Version
	res.PfcpHeader.MPFlag = 0
	res.PfcpHeader.SFlag = 1

	res.PfcpHeader.MessageType = pfcp.PFCP_Session_Establishment_Response
	res.PfcpHeader.SEID = req.IE.CPFSEID.SEID
	res.PfcpHeader.SequenceNumber = req.PfcpHeader.SequenceNumber

	//res.IE.NodeID todo
	res.IE.NodeID.Set()
	res.IE.NodeID.SetValue(pfcp.NodeIDType_IPv4_address, []byte(net.ParseIP(UpfN4Layer.UpfIp).To4()))

	//res.IE.Cause
	res.IE.Cause.Set() //Cause_Request_accepted = 1 // (success)

	// 可选IE
	//res.IE.UPFSEID
	res.IE.UPFSEID = &pfcp.IEFSEID{}
	res.IE.UPFSEID.Set()
	res.IE.UPFSEID.SEID = ufpCxt.SEID

	res.IE.UPFSEID.V4Flag = 1
	res.IE.UPFSEID.IPv4Addr = net.ParseIP(UpfN4Layer.UpfIp).To4()

	res.IE.IeFlags.Set(pfcp.IE_F_SEID)

	// 释放req.PfcpHeader.SEID 的会话
	rlogger.Trace(moduleTag, rlogger.INFO, &req.PfcpHeader, "Session Establishment Response:(MessageType %v,SEID %v,SequenceNumber %v,UPF SEID %v)",
		res.PfcpHeader.MessageType, res.PfcpHeader.SEID, res.PfcpHeader.SequenceNumber, res.IE.UPFSEID.SEID)
	//rlogger.Trace(moduleTag, rlogger.INFO, req, "  .MessageType    %v", res.PfcpHeader.MessageType)
	//rlogger.Trace(moduleTag, rlogger.INFO, req, "  .SEID           %v", res.PfcpHeader.SEID)
	//rlogger.Trace(moduleTag, rlogger.INFO, req, "  .SequenceNumber %v", res.PfcpHeader.SequenceNumber)
	//rlogger.Trace(moduleTag, rlogger.INFO, req, "  .UPF SEID %v", res.IE.UPFSEID.SEID)

	// todo: test 应用规则
	/*tuple := &pdr.IpPacketHeaderFields{
		SrcIp:     net.ParseIP("192.0.2.10"),
		SrcPort:   5000,
		DstIp:     net.ParseIP("192.0.2.20"),
		DstPort:   7000,
		Protocol:  6,
		Direction: nasie.DownlinkOnly,
	}
	packetDetectionRule, _ := pdr.LookupPDRs(tuple)
	fmt.Printf("packetDetectionRule :%+v", packetDetectionRule)
	*/
	return nil
}

func StoreUpfN4Context(req pfcp.SessionEstablishmentRequest) *n4context.N4SessionContext {
	rlogger.FuncEntry(moduleTag, &req.PfcpHeader)
	// 上下文中维护SMF SEID与 UFP SEID的对应关系
	seid, err := n4context.GetSEID()
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, &req.PfcpHeader, "Get Session ID failed:%s", err)
		return nil
	}
	// N4 session context
	n4Cxt := &n4context.N4SessionContext{}
	n4Cxt.SEID = seid // key upf 本地SEID
	n4Cxt.SmfSEID = req.IE.CPFSEID
	n4Cxt.PDRs = req.IE.CreatePDRs
	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "Store upf n4 context:upf SEID %v,smf SEID %v,PDRs length:%d",
		n4Cxt.SEID, n4Cxt.SmfSEID, len(n4Cxt.PDRs))
	//rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "  .SEID:       %v", n4Cxt.SEID)
	//rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "  .Smf SEID:   %v", n4Cxt.SmfSEID)
	//rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "  .PDRs length:%d", len(n4Cxt.PDRs))

	// 可选参数
	n4Cxt.URRs = req.IE.CreateURRs
	n4Cxt.QERs = req.IE.CreateQERs
	n4Cxt.FARs = req.IE.CreateFARs
	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "Store upf n4 context:CreateFARs length:%d", len(n4Cxt.FARs))
	// paging support
	n4Cxt.BAR = req.IE.CreateBAR

	n4Cxt.CreateTrafficEndpoints = req.IE.CreateTrafficEndpoints
	n4Cxt.PDNType = req.IE.PDNType
	n4Cxt.UserPlaneInactivityTimer = req.IE.UserPlaneInactivityTimer
	n4Cxt.UserID = req.IE.UserID
	n4Cxt.TraceInformation = req.IE.TraceInformation
	// 申请空间
	n4Cxt.LocalFTEID = make(map[uint16]*pfcp.IEFTEID)
	n4Cxt.NetworkInstance = make(map[uint16]*pfcp.IENetworkInstance)
	n4Cxt.PDRQFIs = make(map[uint16][]*pfcp.IEQFI)
	n4Cxt.SDFFilters = make(map[uint16][]*pfcp.IESDFFilter)
	n4Cxt.ApplicationID = make(map[uint16]*pfcp.IEApplicationID)
	n4Cxt.EthPacketFilters = make(map[uint16][]*pfcp.IEEthernetPacketFilter)
	for i, v := range req.IE.CreatePDRs {
		//	CN tunnel info.
		n4Cxt.LocalFTEID[v.PDRID.RuleID] = v.PDI.LocalFTEID // PDR对应的TEID
		//-	Network instance.
		n4Cxt.NetworkInstance[v.PDRID.RuleID] = v.PDI.NetworkInstance
		//-	QFIs.
		n4Cxt.PDRQFIs[v.PDRID.RuleID] = v.PDI.QFIs
		//-	IP Packet Filter Set
		// todo: 设置为json格式,放通
		/*if v.PDRID.RuleID == 1 && len(v.PDI.SDFFilters) > 0 {
			v.PDI.SDFFilters[0].FlowDescription = []byte(`{"action":"permit","dir":"in","proto":"ip","src_ip":"192.168.240.0/0","dst_ip":"192.168.103.200/0","src_port":"0","dst_port":"0"}`)
		} else if len(v.PDI.SDFFilters) > 0 {
			v.PDI.SDFFilters[0].FlowDescription = []byte(`{"action":"permit","dir":"out","proto":"ip","src_ip":"172.16.3.200/0","dst_ip":"192.168.240.0/0","src_port":"0","dst_port":"0-0"}`)

		}*/
		n4Cxt.SDFFilters[v.PDRID.RuleID] = v.PDI.SDFFilters
		//Application Identifier
		n4Cxt.ApplicationID[v.PDRID.RuleID] = v.PDI.ApplicationID
		//Ethernet Packet Filter Set
		n4Cxt.EthPacketFilters[v.PDRID.RuleID] = v.PDI.EthPacketFilters
		if v.PDI.LocalFTEID != nil {
			rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "Store upf n4 context:PDRs[%d].LocalFTEID.TEID:%v", i, v.PDI.LocalFTEID.TEID)
		}
		if v.PDI.QFIs != nil {
			if v.PDI.QFIs[0] != nil {
				rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "Store upf n4 context:PDRs[%d].PDI.QFIs[0]:%v", i, v.PDI.QFIs[0].Value)
			}
		}

	}
	// 23.501 5.8.3.2	Buffering at UPF
	n4Cxt.Buffer = list.New()

	// todo counter
	n4Cxt.MetricItems, err = metrics.SessionCounterInit()
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Add upf context Counter failed:%s", err)
		return nil
	}
	n4Cxt.MetricItemsSnapshot, err = metrics.SessionCounterInit()
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Add upf context Counter Snapshot failed:%s", err)
		return nil
	}
	// 保存当前上下文
	err = n4context.AddIndexN4Context(n4context.N4SessionIDKey(seid), n4Cxt)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Add upf context failed:%s", err)
		return nil
	}
	// 存储并应用请求中收到的规则
	err = pdr.ConfigPDRsTable(n4Cxt) //todo:更新1.1 ConfigPDRsTable(n4Cxt) //Configuration rule table
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Config PDRs Table failed:%s", err)
	}
	return n4Cxt
}

// N4 Session Modify Request
func (s *N4Msg) SessionModifyRequest(req pfcp.SessionModifyRequest,
	res *pfcp.SessionModifyResponse) error {
	rlogger.FuncEntry(moduleTag, &req.PfcpHeader)
	//N4 Session Establishment Response
	if req.PfcpHeader.MessageType != pfcp.PFCP_Session_Modification_Request {
		return errors.New("Session Modify Request message type error: " +
			strconv.Itoa(int(req.PfcpHeader.MessageType)))
	}
	//fmt.Printf("SessionModifyRequest: %+v\n", req)
	rlogger.Trace(moduleTag, rlogger.INFO, &req.PfcpHeader, "SessionModifyRequest: %+v", req.PfcpHeader)
	// 读取当前上下文
	n4Cxt, err := n4context.GetN4Context(n4context.N4SessionIDKey(req.PfcpHeader.SEID))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Failed to get N4 Context:%s", err)
		return err
	}
	// 填充UPF N4上下文,pfcpHeader中携带的是local SEID
	n4Cxt.SEID = req.PfcpHeader.SEID
	// 可选参数

	// 更新上下文
	// FARs
	for _, src := range req.IE.UpdateFARs {
		if src == nil {
			continue
		}
		have := false
		// 更新上下文中的FARs
		rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "src Update FAR Len:%d", len(req.IE.UpdateFARs))
		rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "dst Cxt FAR len:%d", len(n4Cxt.FARs))
		for _, dst := range n4Cxt.FARs {
			if dst == nil {
				continue
			}
			if dst.FARID.Value == src.FARID.Value {
				have = true
				//dst.FARID = src.FARID
				rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "Src FAR:%+v", src)
				rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "Dst FAR:%+v", dst)
				dst.ApplyAction = src.ApplyAction
				UpdateForwardParameters(src, dst)
				dst.BARID = src.BARID
				//rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "Src FAR:%+v", src)
				//rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "Dst FAR:%+v", dst)

			}
		}
		// far id 不存在,返回 73
		if !have {
			res.PfcpHeader.Version = pfcp.Version
			res.PfcpHeader.MPFlag = 0
			res.PfcpHeader.SFlag = 1

			res.PfcpHeader.MessageType = pfcp.PFCP_Session_Modification_Response
			res.PfcpHeader.SEID = n4Cxt.SmfSEID.SEID
			res.PfcpHeader.SequenceNumber = req.PfcpHeader.SequenceNumber
			//res.IE.Cause.Set()
			res.IE.Cause.CauseValue = pfcp.Cause_Rule_creation_modification_Failure
			rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "SessionModifyResponse SMF SEID : %d,cause : %v", n4Cxt.SmfSEID.SEID,
				res.IE.Cause.CauseValue)
			return nil
		}
		//-	Network instance.
	} // end FARs
	// BAR
	if req.IE.UpdateBAR != nil {
		if n4Cxt.BAR == nil {
			n4Cxt.BAR = &pfcp.IECreateBAR{}
			n4Cxt.BAR.Set()
		}
		n4Cxt.BAR.BARID = req.IE.UpdateBAR.BARID
		n4Cxt.BAR.DLDataNotificationDelay = req.IE.UpdateBAR.DLDataNotificationDelay
		// 使用SugBuffPacketsCount作为down link data 的初始值
		n4Cxt.BAR.SugBuffPacketsCount = req.IE.UpdateBAR.SuggestedBufferingPacketsCount
	}

	// 更新N4上下文
	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "upf context update SEID : %d", n4Cxt.SEID)
	err = n4context.UpdateN4Context(n4context.N4SessionIDKey(n4Cxt.SEID), n4Cxt)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Update config PDRs Table failed:%s", err)
	}
	// 存储并应用请求中收到的规则
	err = pdr.ConfigPDRsTableUpdate(n4Cxt)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Cxt, "Update config PDRs Table failed:%s", err)
	}
	//test todo
	/*buf := []byte{1, 1, 1, 1, 1, 1, 1, 1}
	n4Cxt.Buffer.PushBack(buf)
	buf1:= []byte{2, 1, 1, 1, 1, 1, 1, 1}
	n4Cxt.Buffer.PushBack(buf1)
	buf2:= []byte{3, 1, 1, 1, 1, 1, 1, 1}
	n4Cxt.Buffer.PushBack(buf2)
	*/
	// todo 激活n3通道时，发送缓存数据到n3
	// 1、发送buffer引用到sendingList
	// 2、守候的处理函数：接收任务，从sendingList中循环获取一个buffer引用，阻塞的发送到buffer chan
	// 3、守候处理函数：从阻塞的buffer chan取消息，并完成发送流程
	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "PagingSend SessionModifyRequest n4Cxt.Buffer.Len : %d", n4Cxt.Buffer.Len())
	if n4Cxt.Buffer.Len() > 0 {
		SendingList.Rw.Lock()
		SendingList.SendList.PushBack(n4Cxt.Buffer)
		SendingList.Rw.Unlock()
		rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "PagingSend SessionModifyRequest SendingList.SendList len : %d", SendingList.SendList.Len())
		// 清空buffer
		n4Cxt.Buffer = list.New()
		// 通知有新任务,非阻塞通知
		select {
		case SendingList.State <- struct{}{}:
			rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "PagingSend SessionModifyRequest SendingList receive task ")

		default:
			rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "PagingSend SessionModifyRequest SendingList receive repeat task ")
			//队列已满，丢弃重复的通知
		}
		rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "PagingSend SendingList length: %d", SendingList.SendList.Len())
		//SendBufferMsg(n4Cxt)
	}

	res.PfcpHeader.Version = pfcp.Version
	res.PfcpHeader.MPFlag = 0
	res.PfcpHeader.SFlag = 1

	res.PfcpHeader.MessageType = pfcp.PFCP_Session_Modification_Response
	res.PfcpHeader.SEID = n4Cxt.SmfSEID.SEID
	res.PfcpHeader.SequenceNumber = req.PfcpHeader.SequenceNumber
	//fmt.Println("smf context update SEID : ", n4Cxt.SmfSEID.SEID)
	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "SessionModifyResponse SMF SEID : %d", n4Cxt.SmfSEID.SEID)
	//res.IE.Cause.Set()
	res.IE.Cause.Set() //Cause_Request_accepted = 1 // (success)
	// todo: test 应用规则
	/*tuple := &pdr.IpPacketHeaderFields{
		SrcIp:     net.ParseIP("192.0.2.10"),
		SrcPort:   5000,
		DstIp:     net.ParseIP("192.0.2.20"),
		DstPort:   7000,
		Protocol:  17,
		Direction: nasie.UplinkOnly,
	}
	packetDetectionRule, _ := pdr.LookupULPDRs(uint32(n4Cxt.PDRs[0].PDI.LocalFTEID.TEID),tuple)
	fmt.Printf("packetDetectionRule :%+v", packetDetectionRule)
	*/
	return nil
}

func UpdateForwardParameters(src *pfcp.IEUpdateFAR, dst *pfcp.IECreateFAR) {
	if src.UpdateForwardingPara == nil {
		rlogger.Trace(moduleTag, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters is nil ", src.FARID.Value)
		return
	}
	if dst.ForwardingParameters == nil {
		dst.ForwardingParameters = &pfcp.IEForwardingParameters{}
	}
	if src.UpdateForwardingPara.DstInterface != nil {
		dst.ForwardingParameters.DstInterface.Value = src.UpdateForwardingPara.DstInterface.Value
		rlogger.Trace(moduleTag, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters,DstInterface:%v ",
			src.FARID.Value,
			src.UpdateForwardingPara.DstInterface.Value)
	}

	if src.UpdateForwardingPara.OuterHeaderCreation != nil {
		if dst.ForwardingParameters != nil {
			dst.ForwardingParameters.OuterHeaderCreation = src.UpdateForwardingPara.OuterHeaderCreation
			rlogger.Trace(moduleTag, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters,OuterHeaderCreation:%v ",
				src.FARID.Value, src.UpdateForwardingPara.OuterHeaderCreation)
		} else {
			//dst.ForwardingParameters.TransportLevelMarking = src.UpdateForwardingPara.TransportLevelMarking
			// AN Tunnel Information
			forwardingParameters := &pfcp.IEForwardingParameters{OuterHeaderCreation: &pfcp.IEOuterHeaderCreation{}}
			dst.ForwardingParameters = forwardingParameters

			dst.ForwardingParameters.OuterHeaderCreation = src.UpdateForwardingPara.OuterHeaderCreation
			rlogger.Trace(moduleTag, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters,create OuterHeaderCreation:%v ",
				src.FARID.Value, src.UpdateForwardingPara.OuterHeaderCreation)
		}

	}

	if src.UpdateForwardingPara.NetworkInstance != nil {
		dst.ForwardingParameters.NetworkInstance = src.UpdateForwardingPara.NetworkInstance
		rlogger.Trace(moduleTag, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters,NetworkInstance:%v ",
			src.FARID.Value, src.UpdateForwardingPara.NetworkInstance)
	}

}

func SendBufferMsg(n4Cxt *n4context.N4SessionContext) bool {
	rlogger.FuncEntry(moduleTag, n4Cxt)
	l := n4Cxt.Buffer
	rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "SendBufferMsg len:%d ", l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		// fmt.Println(e.Value)
		rlogger.Trace(moduleTag, rlogger.INFO, n4Cxt, "%d SendBufferMsg:%v ", i, e.Value)
		i++
		select {
		case UpfN4Layer.BufferMsg <- e.Value.([]byte):
			return true
		default:
			return false //队列已满，消息丢弃
		}
	}
	n4Cxt.Buffer = list.New()
	return true
}
func SendBufferMsgObstructive(l *list.List) bool {
	rlogger.FuncEntry(moduleTag, nil)
	rlogger.Trace(moduleTag, rlogger.INFO, nil, "SendBufferMsg list len:%d ", l.Len())
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		// fmt.Println(e.Value)
		rlogger.Trace(moduleTag, rlogger.INFO, nil, "%d SendBufferMsg:%v ", i, e.Value)
		i++
		select {
		case UpfN4Layer.BufferMsg <- e.Value.([]byte):
			rlogger.Trace(moduleTag, rlogger.INFO, nil, "SendBufferMsg value:%v ", e.Value.([]byte))

		}
	}

	return true
}

// 作为client时，发起N4请求
// N4 Data Notification --> PFCP Session Report Request
func N4SessionReportRequest(n4Ctxt *n4context.N4SessionContext, dfCxt *pdrcontext.DataFlowContext) error {
	rlogger.FuncEntry(moduleTag, n4Ctxt)

	/*// get N4 context
	n4Ctxt, err := upfcontext.GetN4Context(upfcontext.N4SessionIDKey(dfCxt.SEID))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil,  "Failure to get N4 Context:%s", err)
		// 本地无上下文，返回成功
		return nil
	}*/

	// 2.	The UPF sends an N4 session Report request message to the SMF.
	//*pfcp.SessionReleaseRequest
	reqN4, err := CreateN4ReportRequestMsg(n4Ctxt, dfCxt)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Ctxt, "Failed to Create N4 session release request message :%s", err)
		return err
	}
	// 3 发送请求消息，同步处理响应
	var resN4 pfcp.SessionReportResponse
	upfIpN4Port := n4Ctxt.SmfSEID.IPv4Addr.String() + ":" + strconv.Itoa(configure.UpfConf.N4.Smf.Port) //":8806" //127.0.0.1
	err = SendN4MsgI(*reqN4, &resN4, upfIpN4Port)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Ctxt, "Failed to send N4 message:%s", err)
		return err
	}
	rlogger.Trace(moduleTag, rlogger.INFO, n4Ctxt, "Send N4 report request message to :%s", upfIpN4Port)
	// 4 响应成功,SMF响应失败消息
	if resN4.IE.Cause.CauseValue != pfcp.Cause_Request_accepted {
		return fmt.Errorf("Failed to Report N4 context,cause=%d", resN4.IE.Cause.CauseValue)
	}
	/*//1.	SMF receives the trigger to remove the N4 session context for the PDU Session.
	// delete N4 context
	// 在upf成功响应后释放本地 N4 context
	err = smfcontext.DeleteN4Context(smfcontext.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil,  "Failure to delete N4 Context:%s", err)
		return err
	}*/

	// 返回成功
	return nil
}

// IEsSessionReportRequest
func CreateN4ReportRequestMsg(n4Ctxt *n4context.N4SessionContext,
	dfCxt *pdrcontext.DataFlowContext) (*pfcp.SessionReportRequest, error) {
	rlogger.FuncEntry(moduleTag, n4Ctxt)
	var reqN4 pfcp.SessionReportRequest
	// 获取UPF的SEID
	reqN4.PfcpHeader.Set()
	reqN4.PfcpHeader.MessageType = pfcp.PFCP_Session_Report_Request
	reqN4.PfcpHeader.SEID = n4Ctxt.SmfSEID.SEID
	reqN4.PfcpHeader.SequenceNumber = SequenceNumber
	SequenceNumber += 1

	// 2 请求消息构造
	reqN4.IE = pfcp.IEsSessionReportRequest{}
	// todo
	reqN4.IE.ReportType.Set(pfcp.DLDR)

	reqN4.IE.DownlinkDataReport = &pfcp.IEDownlinkDataReport{}
	reqN4.IE.DownlinkDataReport.Set()
	reqN4.IE.DownlinkDataReport.PDRID.Set(dfCxt.RuleID)

	reqN4.IE.DownlinkDataReport.DLDataServiceInfo.Set()

	reqN4.IE.DownlinkDataReport.DLDataServiceInfo.QFI = dfCxt.DP.QFI
	reqN4.IE.DownlinkDataReport.DLDataServiceInfo.Flag &= pfcp.QFII_Flag

	if utils.IsSetByte(dfCxt.DP.PPP, 7) { // todo 定义7
		reqN4.IE.DownlinkDataReport.DLDataServiceInfo.PPIValue = dfCxt.DP.PPI
		reqN4.IE.DownlinkDataReport.DLDataServiceInfo.Flag &= pfcp.PPI_Flag
	}

	return &reqN4, nil
}

func SendN4MsgI(reqN4 interface{}, resN4 interface{}, ipPort string) error {
	rlogger.FuncEntry(moduleTag, nil)
	conn, err := jsonrpc.Dial("tcp", ipPort)
	if err != nil {
		fmt.Println("dailing error: ", err)
		return err
	}
	defer conn.Close()

	var serviceMethod string
	switch reqN4.(type) {
	case pfcp.SessionReportRequest:
		serviceMethod = "N4Msg.SessionReportRequest"
		//fmt.Println("N4Msg.SessionModifyRequest")
	default:
		serviceMethod = "N4Msg.SessionReportRequest"
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "serviceMethod error")
		return nil
	}

	SessionCall := conn.Go(serviceMethod, reqN4, resN4, nil)
	//fmt.Printf("SessionReleaseResponse：%s\n", resN4.Str)
	//time.Sleep(1*time.Second)
	//fmt.Printf("SessionReleaseResponse：%s\n", resN4.Str)
	// 等待响应
	SessionCall = <-SessionCall.Done
	if SessionCall.Error != nil {
		fmt.Printf("Add: expected no error but got string %q",
			SessionCall.Error.Error())
		return nil
	}
	//fmt.Printf("SessionModifyResponse: %+v\n", resN4)

	return nil
}

func N4SessionDeactivationBuffering(buf []byte, dfCxt *pdrcontext.DataFlowContext) error {
	rlogger.FuncEntry(moduleTag, nil)

	// get N4 context
	n4Ctxt, err := n4context.GetN4Context(n4context.N4SessionIDKey(dfCxt.SEID))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Ctxt, "Failure to get N4 Context:%s", err)
		// 本地无上下文，返回成功
		return nil
	}
	if n4Ctxt.BAR == nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Ctxt, "Discard messages when the buffer capacity set nil")
		return nil
	}
	if n4Ctxt.Buffer.Len() < int(n4Ctxt.BAR.SugBuffPacketsCount.CountValue) {
		n4Ctxt.Buffer.PushBack(buf)
		rlogger.Trace(moduleTag, rlogger.INFO, n4Ctxt, "Current buffer length:%d", n4Ctxt.Buffer.Len())
	} else {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Ctxt,
			"Discard messages when the buffer capacity exceeds %d", n4Ctxt.BAR.SugBuffPacketsCount.CountValue)
		return nil
	}
	// 如果是第一个下行数据到达，则通知SMF
	if n4Ctxt.Buffer.Len() == 1 {
		// 发送DL数据到达通知给SMF
		//N4SessionReportRequest(n4Ctxt, dfCxt)
		SessionReportRequest(n4Ctxt, dfCxt)
	}

	return nil
}
