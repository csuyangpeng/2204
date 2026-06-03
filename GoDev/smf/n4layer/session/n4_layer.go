package session

// 3GPP TS 23.501 V15.3.0 (2018-09)
// 3GPP TS 23.502 V15.3.0 (2018-09)
// 3GPP TS 29.244 V15.3.0 (2018-09)
import (
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
	"net/rpc/jsonrpc"
)

// PFCP_Session_Establishment_Request
func N4SessionEstablishmentRequest(pduSessCtxt *gctxt.PduSessContext) error {
	rlogger.FuncEntry(types.ModuleSmfN4, pduSessCtxt)
	//if smfUeCxt == nil {
	//	rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,  "Input parameter check failed")
	//	return errors.New("smfcontext is nil")
	//}
	//1 todo:the selected UPF
	upfIp, err := UpfSelect()

	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, pduSessCtxt, "Failure to select UPF:%s", err)
	}
	// 2 请求消息构造
	// 3GPP TS 23.502 V15.3.0 (2018-09) 4.3.2.2.1	Non-roaming and Roaming with Local Breakout
	// 10a.	The SMF sends an N4 Session Establishment Request to the UPF
	// provides Packet detection, enforcement ,reporting rules,
	// CN Tunnel Info,Inactivity Timer,Trace Requirements.
	reqN4, err := CreateN4SessionReqMsg(pduSessCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, pduSessCtxt, "Failure to create N4SessionReqMsg:%s", err)
	}

	// 3 发送请求消息，同步处理响应
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
	// The SMF also stores the relation between the N4 Session ID and PDU Session for a UE.

	return nil
}

// UpfSelect UPF select
func UpfSelect() (string, error) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//n4Port := configure.SmfConf.N4Conf.UPFPort
	//UpfIpN4Port = n4Layer.UPFIP.String() + ":" + strconv.Itoa(n4Port) //":8805"
	return UpfIpN4Port, nil
}

// upf 中处理
// PFCP_Session_Establishment_Response
func N4SessionEstablishmentResponse(smfUeCxt gctxt.UeContext, n4Cxt gctxt.N4SessionContext) error {
	return nil
}

//func CreateN4SessionReqMsg(pduSessCtxt *gctxt.PduSessContext) (*pfcp.SessionEstablishmentRequest, error) {
//	rlogger.FuncEntry(types.ModuleSmfN4, pduSessCtxt)
//
//	//todo 从上下文中取值
//	//getCxtFromUeContext(smfUeCxt, psi)
//	// 2 请求消息构造
//	var reqN4 pfcp.SessionEstablishmentRequest
//
//	//var resN4 pfcp.SessionEstablishmentResponse
//
//	// 请求消息构造
//	reqN4.PfcpHeader.Set()
//
//	//N4 Session ID
//	// 初始N4 Session时，消息头中SEID为0，LocalSEID是SMF本地分配的SEID。
//	reqN4.PfcpHeader.SEID = 0 //seid
//	reqN4.PfcpHeader.SequenceNumber = SequenceNumber
//	SequenceNumber += 1
//
//	// Packet detection
//	//TS 23.501 5.8.2.11.3	Packet Detection Rule
//	/*Attribute
//	N4 Session ID
//	Rule ID
//	Precedence
//	Packet 	Source interface
//	detection	UE IP address
//	information	Network instance (NOTE 1)
//	CN tunnel info
//	Packet Filter Set
//	Application ID
//	QoS Flow ID
//	Ethernet PDU Session Information
//	Outer header removal
//	Forwarding Action Rule ID
//	List of Usage Reporting Rule ID(s)
//	List of QoS Enforcement Rule ID(s)*/
//
//	// 3GPP TS 29.244 V15.3.0 (2018-09)
//	// Table 7.5.2.1-1: Information Elements in an PFCP Session Establishment Request
//
//	// 必选IE设置
//	reqN4.IE.NodeID.Set()
//	// CPFSEDI
//	seid, err := GetSEID()
//	if err != nil {
//		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, pduSessCtxt, "Get Session ID failed:%s", err)
//	}
//	//ip := net.ParseIP(LocalIp)
//	//cpFSEID := NewFSEID(seid, pfcp.IEFSEID_IPv4_address, ip)
//	cpFSEID := NewFSEID(seid, pfcp.IEFSEID_IPv4_address, configure.SmfConf.N4Conf.SMFIP)
//	reqN4.IE.CPFSEID = *cpFSEID
//	reqN4.IE.CPFSEID.Set()
//
//	// todo
//	loccalTeid := GetTEID() // 每个PDR中使用同一个TEID
//	// PDR Attribute:
//	reqN4.IE.CreatePDRs = make([]*pfcp.IECreatePDR, 0)
//
//	pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
//	if !ok {
//		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, pduSessCtxt, "failed to get procedure context")
//		return nil, fmt.Errorf("failed to get pdu sess estb prcd ctxt")
//	}
//
//	for _, v := range pduSessCtxt.PDR {
//		pdr := pfcp.IECreatePDR{}
//		reqN4.IE.CreatePDRs = append(reqN4.IE.CreatePDRs, &pdr)
//		//Rule ID M
//		pdr.Set()
//		pdr.PDRID.Set(v.PDRID.RuleID)
//		//Precedence M
//		pdr.Precedence.Set(v.Precedence.PrecedenceValue)
//
//		//Packet 	Source interface M
//		pdr.PDI.Set()
//		pdr.PDI.SourceInterface.Set(v.PDI.SourceInterface.InterfaceValue)
//		//detection	UE IP address   O
//		pdr.PDI.UEIPaddress = &pfcp.IEUEIPaddress{}
//		ueIP := pCtxt.UEIP
//		rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, pduSessCtxt, "ue ip: %s", ueIP.String())
//		pdr.PDI.UEIPaddress.V4Set(ueIP, false)
//
//		//information	Network instance  O
//		// -	UPF is connected to 5G-ANs in different IP domains.
//		pdr.PDI.NetworkInstance = &pfcp.IENetworkInstance{}
//		pdr.PDI.NetworkInstance.Set()
//		networkInstance := pCtxt.DNN.String()
//		pdr.PDI.NetworkInstance.NetworkInstance = networkInstance
//
//		//Traffic Endpoint ID  C
//		if v.PDI.TrafficEndpointID != nil {
//			pdr.PDI.TrafficEndpointID = &pfcp.IETrafficEndpointID{}
//			pdr.PDI.TrafficEndpointID.Set(v.PDI.TrafficEndpointID.Value)
//		}
//		//CN tunnel info  O
//		pdr.PDI.LocalFTEID = &pfcp.IEFTEID{}
//		pdr.PDI.LocalFTEID.Set()
//		pdr.PDI.LocalFTEID.TEID = loccalTeid          //GetTEID() // 每个PDR中使用同一个TEID
//		//pdr.PDI.LocalFTEID.IPv4Addr = n4Layer.N3UpfIP // todo UpfIP n4 to n3 N3UpfIP
//
//		//Packet Filter Set  O
//		if v.PDI.SDFFilters != nil {
//			for _, sdff := range v.PDI.SDFFilters {
//				sDFFilter := &pfcp.IESDFFilter{}
//				pdr.PDI.SDFFilters = append(pdr.PDI.SDFFilters, sDFFilter)
//				sDFFilter.SetFlowDesc(sdff.FlowDescription)
//			}
//		}
//		//Application ID  O
//		if v.PDI.ApplicationID != nil {
//			pdr.PDI.ApplicationID = &pfcp.IEApplicationID{}
//			pdr.PDI.ApplicationID.Set()
//			pdr.PDI.ApplicationID.ApplicationIdentifier = v.PDI.ApplicationID.ApplicationIdentifier
//		}
//		//QoS Flow ID  O
//		if v.PDI.QFIs != nil {
//			for i, v := range v.PDI.QFIs {
//				qfi := &pfcp.IEQFI{}
//				pdr.PDI.QFIs = append(pdr.PDI.QFIs, qfi)
//				if i < pfcp.IEFlag_Bit_7 {
//					qfi.Set(v.Value)
//				}
//			}
//		}
//		//Ethernet PDU Session Information
//		pdr.PDI.EthPDUSessionInfo = &pfcp.IEEthernetPDUSessionInfo{}
//		// ETHI (Ethernet Indication):  This bit shall be set to 1.
//		// This refers to all the DL traffic matching the Ethernet PDU session
//		pdr.PDI.EthPDUSessionInfo.Set()
//
//		//Outer header removal C
//		if v.OuterHeaderRemoval != nil {
//			pdr.OuterHeaderRemoval = &pfcp.IEOuterHeaderRemoval{}
//			pdr.OuterHeaderRemoval.SetDescription(v.OuterHeaderRemoval.Description) //(pfcp.IEOuterHRemoval_GTPU_UDP_IPv4)
//		}
//
//		//Forwarding Action Rule ID  O
//		if v.FARID != nil {
//			pdr.FARID = &pfcp.IEFARID{}
//			pdr.FARID.Set(v.FARID.Value)
//		}
//
//		//List of Usage Reporting Rule ID(s)
//		/*pdr.URRIDs = make([]*pfcp.IEURRID, 0)
//		urrid := pfcp.IEURRID{}
//		urrid.Set(URRID)
//		pdr.URRIDs = append(pdr.URRIDs, &urrid)
//
//		//List of QoS Enforcement Rule ID(s)
//		pdr.QERIDs = make([]*pfcp.IEQERID, 0)
//		qerid := pfcp.IEQERID{}
//		qerid.Set(QERID)
//		pdr.QERIDs = append(pdr.QERIDs, &qerid)*/
//
//	}
//	//len(smfUeCxt.GetPduSessCtxt(psi).PDR) < 2
//	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, pduSessCtxt, "Number of PDRs provided by SMF:%d", len(pduSessCtxt.PDR))
//
//	//Create FAR	M
//	reqN4.IE.CreateFARs = make([]*pfcp.IECreateFAR, 0)
//	for _, v := range pduSessCtxt.FAR {
//		far := pfcp.IECreateFAR{}
//		reqN4.IE.CreateFARs = append(reqN4.IE.CreateFARs, &far)
//		far.Set()
//		//FAR ID	        M
//		far.FARID.Set(v.FARID.Value)
//		//Apply Action	M
//		far.ApplyAction.Set(v.ApplyAction.Flag)
//		//Forwarding Parameters	C
//		if v.ForwardingParameters != nil {
//			fwPara := &pfcp.IEForwardingParameters{}
//			far.ForwardingParameters = fwPara
//			fwPara.Set()
//			// Destination Interface	M
//			fwPara.DstInterface.Set(v.ForwardingParameters.DstInterface.Value)
//			// Outer Header Creation 	C
//			if v.ForwardingParameters.OuterHeaderCreation != nil {
//				outerHeaderCreation := &pfcp.IEOuterHeaderCreation{}
//				fwPara.OuterHeaderCreation = outerHeaderCreation
//				outerHeaderCreation.SetGtpuIPv4(
//					v.ForwardingParameters.OuterHeaderCreation.IPv4Addr,
//					v.ForwardingParameters.OuterHeaderCreation.TEID)
//			}
//			// DNN support
//			fwPara.NetworkInstance = &pfcp.IENetworkInstance{}
//			fwPara.NetworkInstance.Set()
//			networkInstance := pCtxt.DNN.String()
//			fwPara.NetworkInstance.NetworkInstance = networkInstance
//			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, pduSessCtxt, "Far Dnn:%s", networkInstance)
//		}
//
//		//Duplicating Parameters 	C
//		//BAR ID	O
//		far.BARID = &pfcp.IEBARID{}
//		far.BARID.Set(configure.SmfConf.N4Conf.BAR.BARID.Value)
//
//	}
//
//	// Create BAR	O
//	reqN4.IE.CreateBAR = &pfcp.IECreateBAR{}
//	reqN4.IE.CreateBAR.Set()
//	reqN4.IE.CreateBAR.BARID.Set(configure.SmfConf.N4Conf.BAR.BARID.Value)
//	reqN4.IE.CreateBAR.SugBuffPacketsCount.Set(configure.SmfConf.N4Conf.BAR.SugBuffPacketsCount.CountValue)
//
//	// User Plane Inactivity Timer	O
//	reqN4.IE.UserPlaneInactivityTimer = &pfcp.IEUserPlaneInactivityTimer{}
//	reqN4.IE.UserPlaneInactivityTimer.Set(pduSessCtxt.UPInactiveTimer)
//	// User ID 	O
//	reqN4.IE.UserID = &pfcp.IEUserID{}
//	reqN4.IE.UserID.SetIMSI(pCtxt.IMSI)
//
//	return &reqN4, nil
//}
func SendN4Msg(reqN4 *pfcp.SessionEstablishmentRequest, ipPort string) (*pfcp.SessionEstablishmentResponse, error) {
	rlogger.FuncEntry(types.ModuleSmfN4, &reqN4.PfcpHeader)
	conn, err := jsonrpc.Dial("tcp", ipPort)
	if err != nil {
		fmt.Println("dailing error: ", err)
		return nil, err
	}
	defer conn.Close()

	var resN4 pfcp.SessionEstablishmentResponse

	rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, &reqN4.PfcpHeader, "SessionEstablishmentRequest: %+v", reqN4)
	rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, &reqN4.PfcpHeader, "ue ip:%s", reqN4.IE.CreatePDRs[0].PDI.UEIPaddress.IPv4Addr.String())

	SessionCall := conn.Go("N4Msg.SessionEstablishmentRequest", reqN4, &resN4, nil)
	//fmt.Printf("SessionReleaseResponse：%s\n", resN4.Str)
	//time.Sleep(1*time.Second)
	//fmt.Printf("SessionReleaseResponse：%s\n", resN4.Str)
	// 等待响应
	SessionCall = <-SessionCall.Done
	if SessionCall.Error != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, &reqN4.PfcpHeader, "Add: expected no error but got string %q",
			SessionCall.Error.Error())
		return nil, nil
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, &reqN4.PfcpHeader, "SessionEstablishmentResponse: %+v", resN4)
	rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, &reqN4.PfcpHeader, ".Version        %+v", resN4.PfcpHeader.Version)
	rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, &reqN4.PfcpHeader, ".MessageType    %v", resN4.PfcpHeader.MessageType)
	rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, &reqN4.PfcpHeader, ".SEID           %v", resN4.PfcpHeader.SEID)
	rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, &reqN4.PfcpHeader, ".SequenceNumber %v", resN4.PfcpHeader.SequenceNumber)
	return &resN4, nil
}

func StoreSmfN4Context(
	reqN4 *pfcp.SessionEstablishmentRequest,
	resN4 *pfcp.SessionEstablishmentResponse,
	n4Cxt *gctxt.N4SessionContext) {
	rlogger.FuncEntry(types.ModuleSmfN4, &reqN4.PfcpHeader)
	// 填充SMF N4上下文
	n4Cxt.SEID = resN4.PfcpHeader.SEID
	n4Cxt.UpfSEID = *(resN4.IE.UPFSEID)
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
	// 保存N4上下文
	gctxt.AddIndexN4Context(gctxt.N4SessionIDKey(n4Cxt.SEID), n4Cxt)
}

//N4 Session Modification Request
// PFCP_Session_Modification_Request
//func N4SessionModificationRequest(seid uint64, fars []pfcp.IECreateFAR) error {
//	rlogger.FuncEntry(types.ModuleSmfN4, nil)
//
//	// TS 23.502 4.3.2.2	UE Requested PDU Session Establishment
//	// 16a.	The SMF initiates an N4 Session Modification procedure with the UPF.
//	// The SMF provides AN Tunnel Info to the UPF as well as the corresponding forwarding rules.
//
//	// 4.4.1.3	N4 Session Modification procedure
//	// The SMF sends an N4 session modification request message to the UPF that contains the update
//	// 获取修改相关的参数
//	// 获取SEID
//	//seid := SEID // config
//	//seid := sessionCtxt.SEID
//	// 获取SMF的SEID
//	n4Ctxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(seid))
//	if err != nil {
//		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failure to get N4 Context:%s", err)
//		return err
//	}
//
//	var reqN4 pfcp.SessionModifyRequest
//	// 获取UPF的SEID
//	reqN4.PfcpHeader.Set()
//	reqN4.PfcpHeader.MessageType = pfcp.PFCP_Session_Modification_Request
//	reqN4.PfcpHeader.SEID = n4Ctxt.UpfSEID.SEID
//	reqN4.PfcpHeader.SequenceNumber = SequenceNumber
//	SequenceNumber += 1
//
//	// 2 请求消息构造
//	// 3GPP TS 23.502 V15.3.0 (2018-09)
//	// 16a.	The SMF initiates an N4 Session Modification procedure with the UPF.
//	// The SMF provides AN Tunnel Info to the UPF as well as the corresponding forwarding rules.
//
//	// TS 23.501 Table 5.8.2.11.6-1: Attributes within Forwarding Action Rule
//	/*N4 Session ID
//	Rule ID
//	Action
//	Network instance
//	Destination interface
//	Outer header creation
//	Send end marker packet(s)
//	Transport level marking
//	Forwarding policy
//	Request for Proxying in UPF
//	Container for header enrichment
//	Buffering Action Rule
//	*/
//	// cpFSEID  O
//	/*ip := net.ParseIP(LocalIp)
//	cpFSEID := NewFSEID(seid, pfcp.IEFSEID_IPv4_address, ip)
//	//reqN4.IE.CPFSEID = cpFSEID
//	reqN4.IE.CPFSEID.Set()*/
//
//	reqN4.IE.UpdateFARs = make([]*pfcp.IEUpdateFAR, 0)
//	//updateFar := pfcp.IEUpdateFAR{}
//	//reqN4.IE.UpdateFARs = append(reqN4.IE.UpdateFARs, &updateFar)
//	/*FARID                 IEFARID
//	ApplyAction           IEApplyAction
//	UpdateForwardingPara  IEUpdateForwardingParameters
//	UpdateDuplicatingPara IEUpdateDuplicatingParameters
//	BARID                 IEBARID*/
//	//updateFar.Set()
//	//FAR ID	        M
//	// FARID = 2
//	for _, v := range fars {
//		updateFar := pfcp.IEUpdateFAR{}
//		reqN4.IE.UpdateFARs = append(reqN4.IE.UpdateFARs, &updateFar)
//		updateFar.Set()
//		/*if v.FARID.Value != 2 {// 选取配置项FAR 2 构造updateFar
//		//	fmt.Println("N4SessionModificationRequest sum FARID:", v.FARID.Value)
//			continue
//		}*/
//		updateFar.FARID.Set(v.FARID.Value)
//		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4SessionModificationRequest FARID:%v", v.FARID.Value)
//		//Apply Action	C
//		//updateFar.ApplyAction.Set(v.ApplyAction.Flag)
//		updateFar.ApplyAction.Set(v.ApplyAction.Flag)
//		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4SessionModificationRequest ApplyAction:%v", v.ApplyAction.Flag)
//		//Forwarding Parameters	C
//		if v.ForwardingParameters != nil {
//			fwPara := &pfcp.IEUpdateForwardingParameters{}
//			updateFar.UpdateForwardingPara = fwPara
//			fwPara.Set()
//			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4SessionModificationRequest  ForwardingParameters:%+v", v.ForwardingParameters)
//			// Destination Interface	C
//			fwPara.DstInterface = &pfcp.IEDestinationInterface{}
//			fwPara.DstInterface.Set(v.ForwardingParameters.DstInterface.Value)
//			// Network instance	    C
//			// DNN support
//			if v.ForwardingParameters.NetworkInstance != nil {
//				fwPara.NetworkInstance = &pfcp.IENetworkInstance{}
//				fwPara.NetworkInstance.Set()
//				fwPara.NetworkInstance.NetworkInstance = v.ForwardingParameters.NetworkInstance.NetworkInstance
//				rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Modify Far Dnn:%s", fwPara.NetworkInstance.NetworkInstance)
//			}
//			// Outer Header Creation 	C
//			if v.ForwardingParameters.OuterHeaderCreation != nil {
//				outerHeaderCreation := &pfcp.IEOuterHeaderCreation{}
//				fwPara.OuterHeaderCreation = outerHeaderCreation
//				// gNB tunnel ID and IP
//				outerHeaderCreation.SetGtpuIPv4(
//					v.ForwardingParameters.OuterHeaderCreation.IPv4Addr,
//					v.ForwardingParameters.OuterHeaderCreation.TEID)
//				rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4SessionModificationRequest ForwardingParameters.OuterHeaderCreation:%+v", v.ForwardingParameters.OuterHeaderCreation)
//			} else {
//				if v.FARID.Value == 2 {
//
//					rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "N4SessionModificationRequest ForwardingParameters OuterHeaderCreation is nil")
//				}
//			}
//		} else {
//			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "N4SessionModificationRequest ForwardingParameters is nil")
//		}
//		//	BAR ID	C                                BAR ID
//		updateFar.BARID = &pfcp.IEBARID{}
//		updateFar.BARID.Set(configure.SmfConf.N4Conf.BAR.BARID.Value)
//
//	}
//
//	// UpdateBAR             *IEUpdateBAR type = 86
//	reqN4.IE.UpdateBAR = &pfcp.IEUpdateBARForSMR{}
//	reqN4.IE.UpdateBAR.Set()
//	reqN4.IE.UpdateBAR.BARID.Set(configure.SmfConf.N4Conf.BAR.BARID.Value)
//	reqN4.IE.UpdateBAR.DLDataNotificationDelay.Set(0)
//	// SugBuffPacketsCount max value is 255
//	reqN4.IE.UpdateBAR.SuggestedBufferingPacketsCount.Set(configure.SmfConf.N4Conf.BAR.SugBuffPacketsCount.CountValue)
//
//	/*reqN4, err := CreateN4SessionModifcationReqMsg()
//	if err != nil {
//		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,  "Failure to create N4SessionReqMsg:%s", err)
//	}*/
//	// 3 发送请求消息，同步处理响应
//	var resN4 pfcp.SessionModifyResponse
//	upfIpN4Port, _ := UpfSelect()
//	err = SendN4MsgI(reqN4, &resN4, upfIpN4Port)
//	if err != nil {
//		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failed to send N4 message:%s", err)
//		return nil
//	}
//
//	// 4 响应成功，更新上下文, TS 23.502 4.4.1.2	N4 Session Establishment procedure
//	// TS 29.244 7.5.5	PFCP Session Modification Response
//	if resN4.IE.Cause.CauseValue == pfcp.Cause_Request_accepted {
//		UpdateSmfN4Context(&reqN4, &resN4)
//	}
//	// Todo: 保存supi + pdu session id 与SEID的对应关系
//	// The SMF also stores the relation between the N4 Session ID and PDU Session for a UE.
//
//	return nil
//}

func SendN4MsgI(reqN4 interface{}, resN4 interface{}, ipPort string) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	conn, err := jsonrpc.Dial("tcp", ipPort)
	if err != nil {
		fmt.Println("dailing error: ", err)
		return err
	}
	defer conn.Close()

	//var resN4 pfcp.SessionEstablishmentResponse

	//fmt.Printf("SessionReleaseRequest: %+v\n", reqN4)
	var serviceMethod string
	switch reqN4.(type) {
	case pfcp.SessionModifyRequest:
		serviceMethod = "N4Msg.SessionModifyRequest"
		//fmt.Println("N4Msg.SessionModifyRequest")
	case pfcp.SessionReleaseRequest:
		serviceMethod = "N4Msg.SessionReleaseRequest"
		//fmt.Println("N4Msg.SessionReleaseRequest")
	default:
		serviceMethod = "N4Msg.SessionModifyRequest"
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "serviceMethod error")
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
	//todo udp send
	//server.SendUdpMsg(server.Server)
	return nil
}

func UpdateSmfN4Context(
	reqN4 *pfcp.SessionModifyRequest,
	resN4 *pfcp.SessionModifyResponse) {
	rlogger.FuncEntry(types.ModuleSmfN4, &reqN4.PfcpHeader)
	// 读取当前上下文
	// *gctxt.N4SessionContext
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &reqN4.PfcpHeader, "reqN4.PfcpHeader.SEID :%v", resN4.PfcpHeader.SEID)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &reqN4.PfcpHeader, "gctxt.LengthOfN4ContextTbl:%v", gctxt.LengthOfN4ContextTbl(gctxt.N4SessionIDCxtType))

	n4Cxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(resN4.PfcpHeader.SEID))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, &reqN4.PfcpHeader, "Failed to get N4 Context:%s", err)
		return
	}

	// 填充SMF N4上下文
	n4Cxt.SEID = resN4.PfcpHeader.SEID
	//n4Cxt.PDRs = reqN4.IE.UpdatePDR
	//n4Cxt.FARs = reqN4.IE.UpdateFAR
	// 可选参数

	// 更新上下文
	//n4Cxt.LocalFTEID = make(map[uint16]*pfcp.IEFTEID)
	// FARs
	for _, src := range reqN4.IE.UpdateFARs {
		// 更新上下文中的FARs
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &reqN4.PfcpHeader, "Src Update FAR: %+v", src)
		for _, dst := range n4Cxt.FARs {
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &reqN4.PfcpHeader, "Dst Update FAR: %+v", dst)
			if dst.FARID.Value == src.FARID.Value {
				//dst.FARID = src.FARID
				rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &reqN4.PfcpHeader, "Update FAR ApplyAction : %+v", src.ApplyAction.Flag)
				dst.ApplyAction = src.ApplyAction
				UpdateForwardParameters(src, dst)
				if src.BARID != nil {
					dst.BARID = src.BARID
				}
			}
		}
		//-	Network instance.
	}
	// paging support
	n4Cxt.BAR.BARID = reqN4.IE.UpdateBAR.BARID
	n4Cxt.BAR.DLDataNotificationDelay.Value = reqN4.IE.UpdateBAR.DLDataNotificationDelay.Value
	n4Cxt.BAR.SugBuffPacketsCount = reqN4.IE.UpdateBAR.SuggestedBufferingPacketsCount

	// 更新N4上下文
	gctxt.UpdateN4Context(gctxt.N4SessionIDKey(n4Cxt.SEID), n4Cxt)
}

func UpdateForwardParameters(src *pfcp.IEUpdateFAR, dst *pfcp.IECreateFAR) {
	if src.UpdateForwardingPara == nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters is nil ", src.FARID.Value)
		return
	}
	if dst.ForwardingParameters == nil {
		dst.ForwardingParameters = &pfcp.IEForwardingParameters{}
	}
	if src.UpdateForwardingPara.DstInterface != nil {
		dst.ForwardingParameters.DstInterface.Value = src.UpdateForwardingPara.DstInterface.Value
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters,DstInterface:%v ",
			src.FARID.Value,
			src.UpdateForwardingPara.DstInterface.Value)
	}

	if src.UpdateForwardingPara.OuterHeaderCreation != nil {
		if dst.ForwardingParameters != nil {
			dst.ForwardingParameters.OuterHeaderCreation = src.UpdateForwardingPara.OuterHeaderCreation
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters,OuterHeaderCreation:%v ",
				src.FARID.Value, src.UpdateForwardingPara.OuterHeaderCreation)
		} else {
			//dst.ForwardingParameters.TransportLevelMarking = src.UpdateForwardingPara.TransportLevelMarking
			// AN Tunnel Information
			forwardingParameters := &pfcp.IEForwardingParameters{OuterHeaderCreation: &pfcp.IEOuterHeaderCreation{}}
			dst.ForwardingParameters = forwardingParameters

			dst.ForwardingParameters.OuterHeaderCreation = src.UpdateForwardingPara.OuterHeaderCreation
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters,create OuterHeaderCreation:%v ",
				src.FARID.Value, src.UpdateForwardingPara.OuterHeaderCreation)
		}

	}

	if src.UpdateForwardingPara.NetworkInstance != nil {
		dst.ForwardingParameters.NetworkInstance = src.UpdateForwardingPara.NetworkInstance
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "FarID(%d) Update Forwarding Parameters,NetworkInstance:%v ",
			src.FARID.Value, src.UpdateForwardingPara.NetworkInstance)
	}

}

// TS 23.502 4.4.1.4	N4 Session Release procedure
/*
1.	SMF receives the trigger to remove the N4 session context for the PDU Session.
2.	The SMF sends an N4 session release request message to the UPF.
3.	The UPF identifies the N4 session context to be removed by the N4 Session ID and removes the whole session context.
The UPF responds with an N4 session release response message containing any information that the UPF has to provide to the SMF.
4.	The SMF interacts with the network entity which triggered this procedure (e.g. AMF or PCF).
*/
func N4SessionReleaseRequest(seid uint64) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	/*if smfUeCxt == nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,  "Input parameter check failed")
		return errors.New("smfcontext is nil") // todo err 返回一个error对象中可以增加错误码，对应具体的错误
	}
	// get SEID
	pduSessContext := smfUeCxt.GetPduSessCtxt(psi)
	if pduSessContext == nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,  "Failed to get pdu session Context:%s")
		return errors.New("pduSessContext is nil")
	}

	seid := pduSessContext.SEID*/
	// get N4 context
	n4Ctxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n4Ctxt, "Failure to get N4 Context:%s", err)
		// 本地无上下文，返回成功
		return nil
	}

	// 2.	The SMF sends an N4 session release request message to the UPF.
	//*pfcp.SessionReleaseRequest
	reqN4, err := CreateN4ReleaseRequestMsg(n4Ctxt)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n4Ctxt, "Failed to Create N4 session release request message :%s", err)
		return err
	}
	// 3 发送请求消息，同步处理响应
	var resN4 pfcp.SessionReleaseResponse
	upfIpN4Port, _ := UpfSelect()
	err = SendN4MsgI(*reqN4, &resN4, upfIpN4Port)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n4Ctxt, "Failed to send N4 message:%s", err)
		return err
	}

	// 4 响应成功,UPF释放上下文成功
	if resN4.IE.Cause.CauseValue != pfcp.Cause_Request_accepted {
		return fmt.Errorf("Failed to release N4 context,cause=%d", resN4.IE.Cause.CauseValue)
	}
	//1.	SMF receives the trigger to remove the N4 session context for the PDU Session.
	// delete N4 context
	// 在upf成功响应后释放本地 N4 context
	err = gctxt.DeleteN4Context(gctxt.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n4Ctxt, "Failure to delete N4 Context:%s", err)
		return err
	}

	// 返回成功
	return nil
}

// IEsSessionDelRequest
func CreateN4ReleaseRequestMsg(n4Ctxt *gctxt.N4SessionContext) (*pfcp.SessionReleaseRequest, error) {

	var reqN4 pfcp.SessionReleaseRequest
	// 获取UPF的SEID
	reqN4.PfcpHeader.Set()
	reqN4.PfcpHeader.MessageType = pfcp.PFCP_Session_Deletion_Request
	reqN4.PfcpHeader.SEID = n4Ctxt.UpfSEID.SEID
	reqN4.PfcpHeader.SequenceNumber = SequenceNumber
	SequenceNumber += 1

	// 2 请求消息构造
	reqN4.IE = pfcp.IEsSessionDelRequest{}

	return &reqN4, nil
}
