package create

import (
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"lite5gc/smf/smfcontext/gctxt"
)

// PFCP_Session_Establishment_Request
func SessionEstablishmentRequest(msgData *gctxt.ScN4MsgData, n *pfcpv1.Node) (*pfcp.SessionEstablishmentRequest, error) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)

	// 2 请求消息构造
	// 3GPP TS 23.502 V15.3.0 (2018-09) 4.3.2.2.1	Non-roaming and Roaming with Local Breakout
	// 10a.	The SMF sends an N4 Session Establishment Request to the UPF
	// provides Packet detection, enforcement ,reporting rules,
	// CN Tunnel Info,Inactivity Timer,Trace Requirements.
	reqN4, err := CreateN4SessionReqMsg(msgData, n)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failure to create N4SessionReqMsg:%s", err)
		return nil, err
	}

	/*// 3 发送请求消息，同步处理响应
	resN4, err := SendN4Msg(reqN4, upfIp)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, pduSessCtxt, "Failed to send N4 message:%s", err)
		return nil
	}
	if resN4 == nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, pduSessCtxt, "Failed to receive N4 response message:%s", err)
		return nil
	}

	// 4 响应成功，保存上下文, TS 23.502 4.4.1.2	N4 Session Establishment procedure
	n4Cxt := &gctxt.N4SessionContext{}
	if resN4.IE.Cause.CauseValue == pfcp.Cause_Request_accepted {

		StoreSmfN4Context(reqN4, resN4, n4Cxt)
		// 返回SEID给UE context
		pCtxt := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
		pCtxt.Seid = n4Cxt.SEID
	}
	// Todo: 保存supi + pdu session id 与SEID的对应关系
	// The SMF also stores the relation between the N4 Session ID and PDU Session for a UE.*/

	return reqN4, nil
}

func CreateN4SessionReqMsg(msgData *gctxt.ScN4MsgData, n *pfcpv1.Node) (*pfcp.SessionEstablishmentRequest, error) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	utils.Uint32AtomicAdd(&n.SequenceNumber, 1)
	// 2 请求消息构造

	// 请求消息构造
	//header fill
	request := &pfcp.SessionEstablishmentRequest{
		PfcpHeader: pfcp.PfcpHeaderforSession{
			Version:        pfcp.Version,
			SFlag:          pfcp.Flag,
			MessageType:    pfcp.PFCP_Session_Establishment_Request,
			Length:         0, // 编码后填充
			SEID:           0, // 初始N4 Session时，消息头中SEID为0，LocalSEID是SMF本地分配的SEID。
			SequenceNumber: n.SequenceNumber},
	}

	// Packet detection
	//TS 23.501 5.8.2.11.3	Packet Detection Rule
	/*Attribute
	N4 Session ID
	Rule ID
	Precedence
	Packet 	Source interface
	detection	UE IP address
	information	Network instance (NOTE 1)
	CN tunnel info
	Packet Filter Set
	Application ID
	QoS Flow ID
	Ethernet PDU Session Information
	Outer header removal
	Forwarding Action Rule ID
	List of Usage Reporting Rule ID(s)
	List of QoS Enforcement Rule ID(s)*/

	// 3GPP TS 29.244 V15.3.0 (2018-09)
	// Table 7.5.2.1-1: Information Elements in an PFCP Session Establishment Request

	// 必选IE设置
	request.IE = pfcp.IEsSessionEstablishmentRequest{}
	request.IE.NodeID = pfcp.IENodeID{
		NodeIDType:  pfcp.NodeIDType_IPv4_address,
		NodeIDvalue: n.Server.LocalAddr.IP.To4(), // smf ip
	}
	// FSEID
	seid := msgData.SEID

	request.IE.CPFSEID = pfcp.IEFSEID{
		V4Flag:   pfcp.Flag,
		SEID:     seid,
		IPv4Addr: n.Server.LocalAddr.IP.To4(), // smf ip
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Smf Seid:%d", seid)
	//loccalTeid := GetTEID() // 每个PDR中使用同一个TEID
	loccalTeid := GetFTEID(n) //GetTEID() // 每个PDR中使用同一个TEID
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "node teid:%d", loccalTeid)
	// PDR Attribute
	request.IE.CreatePDRs = make([]*pfcp.IECreatePDR, 0)
	request.IE.CreateFARs = make([]*pfcp.IECreateFAR, 0)
	request.IE.CreateQERs = make([]*pfcp.IECreateQER, 0)
	//smf pdu EstbPrcdCtxt
	//pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
	//if !ok {
	//	rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, pduSessCtxt, "failed to get procedure context")
	//	return nil, fmt.Errorf("failed to get pdu sess estb prcd ctxt")
	//}
	// pduSessCtxt.PFCPParameters
	for _, v := range msgData.PFCPParameters.PDRs {
		// Create PDR
		{
			pdr, err := v.Fill()
			if err != nil {
				return nil, fmt.Errorf("failed to fill pdr")
			}
			request.IE.CreatePDRs = append(request.IE.CreatePDRs, pdr)
			if pdr.PDI.IeFlags.Test(pfcp.IE_F_TEID) {
				//todo tunnel info
				pdr.PDI.LocalFTEID.TEID = types3gpp.Teid(uint32(loccalTeid) + (uint32(n.TEIDRange) << (32 - n.TEIDRI)))
				pdr.PDI.LocalFTEID.IPv4Addr = n.N3Ip.To16() //configure.SmfConf.N3Conf.UPFIP // upf n3 ip
				rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, nil,
					"node teid:%#x,TEIDRI:%v,TEIDRange:%v", loccalTeid, n.TEIDRI, n.TEIDRange)
				rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, nil,
					"teid:%#x(decimal:%d),n3 ip: %s", pdr.PDI.LocalFTEID.TEID, pdr.PDI.LocalFTEID.TEID, pdr.PDI.LocalFTEID.IPv4Addr)
			}
			if pdr.PDI.IeFlags.Test(pfcp.IE_UE_IP_Address) {
				//	todo
				pdr.PDI.UEIPaddress.IPv4Addr = v.PDI.UEIPAddress
				rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, nil, "ue ip: %s", pdr.PDI.UEIPaddress.IPv4Addr)

			}
			if pdr.PDI.IeFlags.Test(pfcp.IE_Network_Instance) {
				pdr.PDI.NetworkInstance.NetworkInstance = string(types3gpp.EncodeLables([]byte(msgData.DNN.String())))
			}
		}
		//Create FAR	M
		{
			far, err := v.FarI.Far.Fill()
			if err != nil {
				return nil, fmt.Errorf("failed to fill far")
			}
			request.IE.CreateFARs = append(request.IE.CreateFARs, far)
			if far.IeFlags.Test(pfcp.IE_Forwarding_Parameters) {
				// DNN support
				if far.ForwardingParameters.IeFlags.Test(pfcp.IE_Network_Instance) {
					far.ForwardingParameters.NetworkInstance.NetworkInstance = string(types3gpp.EncodeLables([]byte(v.FarI.Far.NetworkInstance))) //pCtxt.DNN.String()
					rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Far Dnn:%s", msgData.DNN.String())
					rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Far Dnn:%s", far.ForwardingParameters.NetworkInstance.NetworkInstance)
				}
			}

			//Duplicating Parameters 	C
			//BAR ID	O
			//far.BARID.Set(configure.SmfConf.N4Conf.BAR.BARID.Value)
		}
		// Create URR	C

		// Create QER	C
		{
			if len(v.QerI) > 0 && v.IeFlags.Test(pfcp.IE_Create_QER) {
				qer, err := v.QerI[0].Qer.Fill()
				if err != nil {
					return nil, fmt.Errorf("failed to fill qer")
				}
				request.IE.CreateQERs = append(request.IE.CreateQERs, qer)
				request.IE.IeFlags.Set(pfcp.IE_Create_QER)
				rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Qer:%+v", qer)
			}
		}
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Number of PDRs provided by SMF:%d", len(msgData.PFCPParameters.PDRs))

	// Create BAR	O
	if len(msgData.PFCPParameters.PDRs) > 0 {
		bar, err := msgData.PFCPParameters.PDRs[1].FarI.Far.BAR.Fill()
		if err != nil {
			return nil, fmt.Errorf("failed to fill bar")
		}
		request.IE.CreateBAR = bar
		request.IE.IeFlags.Set(pfcp.IE_Create_BAR)
	}

	// User Plane Inactivity Timer	O
	request.IE.UserPlaneInactivityTimer = &pfcp.IEUserPlaneInactivityTimer{}
	request.IE.UserPlaneInactivityTimer.Set(msgData.PFCPParameters.UserPlaneInactivityTimer.TimerValue)
	// User ID 	O
	request.IE.UserID = &pfcp.IEUserID{}
	request.IE.UserID.SetIMSI(msgData.IMSI)

	return request, nil
}

func StoreSmfN4Context(
	reqN4 *pfcp.SessionEstablishmentRequest,
	//resN4 *pfcp.SessionEstablishmentResponse,
	n4Cxt *gctxt.N4SessionContext) {
	rlogger.FuncEntry(types.ModuleSmfN4, &reqN4.PfcpHeader)
	// 填充SMF N4上下文
	//SMF SEID
	n4Cxt.SEID = reqN4.IE.CPFSEID.SEID
	//n4Cxt.UpfSEID = *(resN4.IE.UPFSEID) //todo 在响应中保存
	n4Cxt.PDRs = reqN4.IE.CreatePDRs
	n4Cxt.FARs = reqN4.IE.CreateFARs
	// 可选参数
	n4Cxt.URRs = reqN4.IE.CreateURRs
	n4Cxt.QERs = reqN4.IE.CreateQERs
	n4Cxt.FARs = reqN4.IE.CreateFARs
	n4Cxt.BAR = reqN4.IE.CreateBAR

	n4Cxt.CreateTrafficEndpoints = reqN4.IE.CreateTrafficEndpoints
	n4Cxt.PDNType = reqN4.IE.PDNType
	n4Cxt.UserPlaneInactivityTimer = reqN4.IE.UserPlaneInactivityTimer
	n4Cxt.UserID = reqN4.IE.UserID
	n4Cxt.TraceInformation = reqN4.IE.TraceInformation

	// 申请空间
	n4Cxt.LocalFTEID = make(map[uint16]*pfcp.IEFTEID)
	n4Cxt.NetworkInstance = make(map[uint16]*pfcp.IENetworkInstance)
	n4Cxt.PDRQFIs = make(map[uint16][]*pfcp.IEQFI)
	n4Cxt.SDFFilters = make(map[uint16][]*pfcp.IESDFFilter)
	n4Cxt.ApplicationID = make(map[uint16]*pfcp.IEApplicationID)
	n4Cxt.EthPacketFilters = make(map[uint16][]*pfcp.IEEthernetPacketFilter)

	for _, v := range reqN4.IE.CreatePDRs {
		//	CN tunnel info.
		n4Cxt.LocalFTEID[v.PDRID.RuleID] = v.PDI.LocalFTEID // PDR对应的TEID
		//-	Network instance.
		n4Cxt.NetworkInstance[v.PDRID.RuleID] = v.PDI.NetworkInstance
		//-	QFIs.
		n4Cxt.PDRQFIs[v.PDRID.RuleID] = v.PDI.QFIs
		//-	IP Packet Filter Set
		n4Cxt.SDFFilters[v.PDRID.RuleID] = v.PDI.SDFFilters
		//Application Identifier
		n4Cxt.ApplicationID[v.PDRID.RuleID] = v.PDI.ApplicationID
		//Ethernet Packet Filter Set
		n4Cxt.EthPacketFilters[v.PDRID.RuleID] = v.PDI.EthPacketFilters
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "StoreSmfN4Context SMF Seid %+v", n4Cxt.SEID)

	// 保存N4上下文
	gctxt.AddIndexN4Context(gctxt.N4SessionIDKey(n4Cxt.SEID), n4Cxt)
}

func ResStoreSmfN4Context(
	//reqN4 *pfcp.SessionEstablishmentRequest,
	resN4 *pfcp.SessionEstablishmentResponse,
	n4Cxt *gctxt.N4SessionContext) {
	rlogger.FuncEntry(types.ModuleSmfN4, &resN4.PfcpHeader)
	// 填充SMF N4上下文
	//upf seid
	if resN4.IE.UPFSEID != nil {
		n4Cxt.UpfSEID = *(resN4.IE.UPFSEID) //todo 在响应中保存
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "ResStoreSmfN4Context,Smf Seid(%v),UPF Seid (%+v)", n4Cxt.SEID, n4Cxt.UpfSEID)

	// 更新N4上下文
	gctxt.UpdateN4Context(gctxt.N4SessionIDKey(n4Cxt.SEID), n4Cxt)
}
