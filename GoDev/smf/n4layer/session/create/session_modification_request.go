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

//N4 Session Modification Request
// PFCP_Session_Modification_Request
func SessionModificationRequest(msgData *gctxt.ScN4MsgData,
	n *pfcpv1.Node) (*pfcp.SessionModifyRequest, error) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)

	seid := msgData.SEID

	// TS 23.502 4.3.2.2	UE Requested PDU Session Establishment
	// 16a.	The SMF initiates an N4 Session Modification procedure with the UPF.
	// The SMF provides AN Tunnel Info to the UPF as well as the corresponding forwarding rules.

	// 4.4.1.3	N4 Session Modification procedure
	// The SMF sends an N4 session modification request message to the UPF that contains the update
	// 获取修改相关的参数
	// 获取SEID
	//seid := SEID // config
	//seid := sessionCtxt.SEID
	// 获取SMF的SEID
	n4Ctxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "Failed to get N4 Context:%s", err)
		return nil, err
	}

	var reqN4 pfcp.SessionModifyRequest
	// 获取UPF的SEID
	reqN4.PfcpHeader.Set()
	reqN4.PfcpHeader.MessageType = pfcp.PFCP_Session_Modification_Request
	reqN4.PfcpHeader.SEID = n4Ctxt.UpfSEID.SEID
	utils.Uint32AtomicAdd(&n.SequenceNumber, 1)
	reqN4.PfcpHeader.SequenceNumber = n.SequenceNumber

	// 2 请求消息构造
	// 3GPP TS 23.502 V15.3.0 (2018-09)
	// 16a.	The SMF initiates an N4 Session Modification procedure with the UPF.
	// The SMF provides AN Tunnel Info to the UPF as well as the corresponding forwarding rules.

	// TS 23.501 Table 5.8.2.11.6-1: Attributes within Forwarding Action Rule
	/*N4 Session ID
	Rule ID
	Action
	Network instance
	Destination interface
	Outer header creation
	Send end marker packet(s)
	Transport level marking
	Forwarding policy
	Request for Proxying in UPF
	Container for header enrichment
	Buffering Action Rule
	*/
	// cpFSEID  O
	/*ip := net.ParseIP(LocalIp)
	cpFSEID := NewFSEID(seid, pfcp.IEFSEID_IPv4_address, ip)
	//reqN4.IE.CPFSEID = cpFSEID
	reqN4.IE.CPFSEID.Set()*/

	reqN4.IE.UpdateFARs = make([]*pfcp.IEUpdateFAR, 0)
	reqN4.IE.CreatePDRs = make([]*pfcp.IECreatePDR, 0)
	reqN4.IE.CreateFARs = make([]*pfcp.IECreateFAR, 0)

	reqN4.IE.RemovePDRs = make([]*pfcp.IERemovePDR, 0)
	reqN4.IE.RemoveFARs = make([]*pfcp.IERemoveFAR, 0)

	for _, v := range msgData.PFCPParameters.PDRs {
		switch v.Tag {
		case gctxt.CreatePDRs:
			//1 映射PDR
			pdr, err := v.Fill()
			if err != nil {
				return nil, fmt.Errorf("failed to fill pdr")
			}
			far, err := v.FarI.Far.Fill()
			if err != nil {
				return nil, fmt.Errorf("failed to fill far")
			}

			qer, err := v.QerI[0].Qer.Fill()
			if err != nil {
				return nil, fmt.Errorf("failed to fill qer")
			}

			//2 构造IE
			reqN4.IE.IeFlags.Set(pfcp.IE_Create_PDR)

			createPdr := CreateIEPdr(pdr, msgData, n)
			//	add to pdr
			createPdr.IeFlags.Set(pfcp.IE_QER_ID)
			createPdr.QERIDs = append(createPdr.QERIDs, &pfcp.IEQERID{Value: qer.QERID.Value})
			reqN4.IE.CreatePDRs = append(reqN4.IE.CreatePDRs, createPdr)
			//Create FAR
			reqN4.IE.IeFlags.Set(pfcp.IE_Create_FAR)
			createFar := CreateIEFar(far, msgData, n)
			reqN4.IE.CreateFARs = append(reqN4.IE.CreateFARs, createFar)

			reqN4.IE.IeFlags.Set(pfcp.IE_Create_QER)
			createQer := CreateIEQer(qer, msgData, n)
			reqN4.IE.CreateQERs = append(reqN4.IE.CreateQERs, createQer)

		case gctxt.RemovePDRs:
			//1 映射QER
			if len(v.QerI) == 0 {
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "RemovePDRs:", v.RuleID, v.QerI)
				break
			}
			qer, err := v.QerI[0].Qer.Fill()
			if err != nil {
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "RemovePDRs:flow", err)
				return nil, fmt.Errorf("failed to fill qer")
			}
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "RemovePDRs:input flow", qer.QoSflowidentifier.Value)
			// print n4ctxt
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "RemovePDRs:%s", n4Ctxt)

			// 2 qfi to Qer id
			if qer.IeFlags.Test(pfcp.IE_QFI) {
				qfi := qer.QoSflowidentifier.Value
				for _, v := range n4Ctxt.QERs {
					rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "RemovePDRs:n4 flow", v.QoSflowidentifier.Value)
					if v.QoSflowidentifier.Value == qfi {
						qer.QERID.Value = v.QERID.Value
						rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "RemovePDRs:Qer ID", qer.QERID.Value)
						break
					}

				}

			}

			//3 构造IE
			reqN4.IE.IeFlags.Set(pfcp.IE_Remove_QER)
			removeQer := RemoveIEQer(qer, msgData, n)
			reqN4.IE.RemoveQERs = append(reqN4.IE.RemoveQERs, removeQer)
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "reqN4.IE.RemoveQERs", reqN4.IE.RemoveQERs)
			/*//1 映射PDR
			pdr, err := v.Fill()
			if err != nil {
				return nil, fmt.Errorf("failed to fill pdr")
			}
			// qfi to pdr id

			if pdr.PDI.IeFlags.Test(pfcp.IE_QFI) {
				for _, v := range pdr.PDI.QFIs {
					rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "RemovePDRs:QFI", v)
					for pdrid, qfi := range n4Ctxt.PDRQFIs {
						rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "RemovePDRs:pdrid", pdrid, qfi)
						for _, dv := range qfi {
							rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "RemovePDRs:dv", dv)
							if dv.Value == v.Value {
								pdr.PDRID.RuleID = pdrid
								rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "pdr-id:", pdrid)
								break
							}
						}
					}
				}

			}
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "RemovePDRs pdr-id:", pdr.PDRID.RuleID)
			// pdr id to far id
			for _, v := range n4Ctxt.PDRs {
				if v.PDRID.RuleID == pdr.PDRID.RuleID {
					pdr.IeFlags.Set(pfcp.IE_FAR_ID)
					pdr.FARID.Value = v.FARID.Value
					rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "pdr.FARID.Value:", pdr.FARID.Value)
				}
			}
			//2 构造IE
			// pdr
			reqN4.IE.IeFlags.Set(pfcp.IE_Remove_PDR)
			removePdr := RemoveIEPdr(pdr, pduSessCtxt, n)
			reqN4.IE.RemovePDRs = append(reqN4.IE.RemovePDRs, removePdr)
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "reqN4.IE.RemovePDRs", reqN4.IE.RemovePDRs)
			// far
			reqN4.IE.IeFlags.Set(pfcp.IE_Remove_FAR)
			removeFar := RemoveIEFar(pdr, pduSessCtxt, n)
			reqN4.IE.RemoveFARs = append(reqN4.IE.RemoveFARs, removeFar)
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "reqN4.IE.RemoveFARs", reqN4.IE.RemoveFARs)*/

		default:
			//1 映射PDR
			//更新下行far，获取gnb tunnel info

			far, err := v.FarI.Far.Fill()
			if err != nil {
				return nil, fmt.Errorf("failed to fill far")
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "N4SessionModificationRequest", err)
			}
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4SessionModificationRequest  ForwardingParameters:%+v", far.ForwardingParameters)
			//2 构造IE
			reqN4.IE.IeFlags.Set(pfcp.IE_Update_FAR)

			updateFar := UpdateIEFar(far)
			reqN4.IE.UpdateFARs = append(reqN4.IE.UpdateFARs, updateFar)
		}

	}
	// Create BAR	O
	var bar *pfcp.IECreateBAR

	if len(msgData.PFCPParameters.PDRs) > 0 &&
		msgData.PFCPParameters.PDRs[0].FarI.Far.IeFlags.Test(pfcp.IE_Create_BAR) {
		v, err := msgData.PFCPParameters.PDRs[0].FarI.Far.BAR.Fill()
		if err != nil {
			return nil, fmt.Errorf("failed to fill bar")
		}
		bar = v
		// UpdateBAR             *IEUpdateBAR type = 86
		{
			reqN4.IE.UpdateBAR = &pfcp.IEUpdateBARForSMR{}
			reqN4.IE.UpdateBAR.Set()
			reqN4.IE.UpdateBAR.BARID.Set(bar.BARID.Value)
			reqN4.IE.UpdateBAR.IeFlags.Set(pfcp.IE_Downlink_Data_Notification_Delay)
			reqN4.IE.UpdateBAR.DLDataNotificationDelay.Set(bar.DLDataNotificationDelay.Value)
			// SugBuffPacketsCount max value is 255
			reqN4.IE.UpdateBAR.IeFlags.Set(pfcp.IE_Suggested_Buffering_Packets_Count)
			reqN4.IE.UpdateBAR.SuggestedBufferingPacketsCount.Set(bar.SugBuffPacketsCount.CountValue)
		}
	}
	/*reqN4, err := CreateN4SessionModifcationReqMsg()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,  "Failure to create N4SessionReqMsg:%s", err)
	}*/

	return &reqN4, nil
}
func CreateIEPdr(pdr *pfcp.IECreatePDR, msgData *gctxt.ScN4MsgData,
	n *pfcpv1.Node) *pfcp.IECreatePDR {
	loccalTeid := GetFTEID(n) //GetTEID() // 每个PDR中使用同一个TEID
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "node teid:%d", loccalTeid)
	//smf pdu EstbPrcdCtxt
	//pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionModPrcdCtxt)
	//if !ok {
	//	rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, pduSessCtxt, "failed to get procedure context")
	//	return nil
	//}
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
		//	todo test
		pdr.PDI.UEIPaddress.IPv4Addr = msgData.PFCPParameters.PDRs[1].PDI.UEIPAddress
		rlogger.Trace(types.ModuleSmfN4, rlogger.DEBUG, nil, "ue ip: %s", pdr.PDI.UEIPaddress.IPv4Addr)

	}
	if pdr.PDI.IeFlags.Test(pfcp.IE_Network_Instance) {
		//pdr.PDI.NetworkInstance.NetworkInstance = string(types3gpp.EncodeLables([]byte(msgData.DNN.String())))
		pdr.PDI.NetworkInstance.NetworkInstance = string(types3gpp.EncodeLables([]byte(msgData.PFCPParameters.PDRs[0].FarI.Far.NetworkInstance)))
	}

	return pdr
}
func CreateIEFar(far *pfcp.IECreateFAR, msgData *gctxt.ScN4MsgData,
	n *pfcpv1.Node) *pfcp.IECreateFAR {
	//pCtxt, ok := pduSessCtxt.GetPrcdCtxt().(*prcdctxt.PduSessionModPrcdCtxt)
	//if !ok {
	//	rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, pduSessCtxt, "failed to get procedure context")
	//	return nil
	//}
	if far.IeFlags.Test(pfcp.IE_Forwarding_Parameters) {
		// DNN support
		if far.ForwardingParameters.IeFlags.Test(pfcp.IE_Network_Instance) {
			// todo test
			far.ForwardingParameters.NetworkInstance.NetworkInstance = string(
				types3gpp.EncodeLables([]byte(msgData.PFCPParameters.PDRs[0].FarI.Far.NetworkInstance))) //pCtxt.DNN.String()
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Far Dnn:%s", msgData.DNN.String())
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Far Dnn:%s", far.ForwardingParameters.NetworkInstance.NetworkInstance)
		}
	}

	//Duplicating Parameters 	C
	//BAR ID	O
	//far.BARID.Set(configure.SmfConf.N4Conf.BAR.BARID.Value)

	return far
}

func CreateIEQer(qer *pfcp.IECreateQER, msgData *gctxt.ScN4MsgData,
	n *pfcpv1.Node) *pfcp.IECreateQER {

	if qer.IeFlags.Test(pfcp.IE_MBR) {
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Qer MBR:%+v", qer.MaximumBitrate)
	}
	if qer.IeFlags.Test(pfcp.IE_GBR) {
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Qer GBR:%+v", qer.GuaranteedBitrate)
	}
	return qer
}

func UpdateIEFar(v *pfcp.IECreateFAR) *pfcp.IEUpdateFAR {
	/*FARID                 IEFARID
	ApplyAction           IEApplyAction
	UpdateForwardingPara  IEUpdateForwardingParameters
	UpdateDuplicatingPara IEUpdateDuplicatingParameters
	BARID                 IEBARID*/
	//updateFar.Set()
	//FAR ID	        M
	// FARID = 2
	updateFar := pfcp.IEUpdateFAR{}
	updateFar.Set()
	/*if v.FARID.Value != 2 {// 选取配置项FAR 2 构造updateFar
	//	fmt.Println("N4SessionModificationRequest sum FARID:", v.FARID.Value)
		continue
	}*/
	updateFar.FARID.Set(v.FARID.Value)

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4SessionModificationRequest FARID:%v", updateFar.FARID.Value)
	//Apply Action	C
	//updateFar.ApplyAction.Set(v.ApplyAction.Flag)
	updateFar.ApplyAction.Set(v.ApplyAction.Flag)
	updateFar.IeFlags.Set(pfcp.IE_Apply_Action)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4SessionModificationRequest ApplyAction:%v", updateFar.ApplyAction.Flag)
	//Forwarding Parameters	C
	if v.ForwardingParameters != nil && v.IeFlags.Test(pfcp.IE_Forwarding_Parameters) {
		fwPara := &pfcp.IEUpdateForwardingParameters{}
		updateFar.UpdateForwardingPara = fwPara
		updateFar.IeFlags.Set(pfcp.IE_Update_Forwarding_Parameters)
		fwPara.Set()
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4SessionModificationRequest  ForwardingParameters:%+v", updateFar.UpdateForwardingPara)
		// Destination Interface	C
		fwPara.DstInterface = &pfcp.IEDestinationInterface{}
		fwPara.DstInterface.Set(v.ForwardingParameters.DstInterface.Value)
		updateFar.UpdateForwardingPara.IeFlags.Set(pfcp.IE_Destination_Interface)
		// Network instance	    C
		// DNN support
		if v.ForwardingParameters.NetworkInstance != nil &&
			v.ForwardingParameters.IeFlags.Test(pfcp.IE_Network_Instance) {
			fwPara.NetworkInstance = &pfcp.IENetworkInstance{}
			fwPara.NetworkInstance.Set()
			updateFar.UpdateForwardingPara.IeFlags.Set(pfcp.IE_Network_Instance)
			fwPara.NetworkInstance.NetworkInstance = v.ForwardingParameters.NetworkInstance.NetworkInstance
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "Modify Far Dnn:%s", fwPara.NetworkInstance.NetworkInstance)
		}
		// Outer Header Creation 	C
		if v.ForwardingParameters.OuterHeaderCreation != nil &&
			v.ForwardingParameters.IeFlags.Test(pfcp.IE_Outer_Header_Creation) {
			outerHeaderCreation := &pfcp.IEOuterHeaderCreation{}
			fwPara.OuterHeaderCreation = outerHeaderCreation
			updateFar.UpdateForwardingPara.IeFlags.Set(pfcp.IE_Outer_Header_Creation)
			// gNB tunnel ID and IP
			outerHeaderCreation.SetGtpuIPv4(
				v.ForwardingParameters.OuterHeaderCreation.IPv4Addr,
				v.ForwardingParameters.OuterHeaderCreation.TEID)
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "N4SessionModificationRequest ForwardingParameters.OuterHeaderCreation:%+v", updateFar.UpdateForwardingPara.OuterHeaderCreation)
		} else {
			if v.FARID.Value == 2 {

				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "N4SessionModificationRequest ForwardingParameters OuterHeaderCreation is nil")
			}
		}
	} else {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "N4SessionModificationRequest ForwardingParameters is nil")
	}
	//	BAR ID	C                                BAR ID
	updateFar.BARID = &pfcp.IEBARID{}
	//updateFar.BARID.Set(configure.SmfConf.N4Conf.BAR.BARID.Value)
	if v.IeFlags.Test(pfcp.IE_BAR_ID) {
		updateFar.BARID.Set(v.BARID.Value)
	}
	return &updateFar
}

func RemoveIEPdr(pdr *pfcp.IECreatePDR, pduSessCtxt *gctxt.PduSessContext,
	n *pfcpv1.Node) *pfcp.IERemovePDR {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	RemPdr := &pfcp.IERemovePDR{}
	RemPdr.PDRID = pdr.PDRID
	if pdr.IeFlags.Test(pfcp.IE_FAR_ID) {
		remFar := &pfcp.IERemoveFAR{}
		remFar.FARID = *pdr.FARID
	}

	return RemPdr
}
func RemoveIEFar(pdr *pfcp.IECreatePDR, pduSessCtxt *gctxt.PduSessContext,
	n *pfcpv1.Node) *pfcp.IERemoveFAR {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	remFar := &pfcp.IERemoveFAR{}
	if pdr.IeFlags.Test(pfcp.IE_FAR_ID) {
		remFar.FARID = *pdr.FARID
	}

	return remFar
}

func RemoveIEQer(qer *pfcp.IECreateQER, msgData *gctxt.ScN4MsgData,
	n *pfcpv1.Node) *pfcp.IERemoveQER {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	RemQer := &pfcp.IERemoveQER{}
	RemQer.QERID = qer.QERID
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "RemoveIEQer:Qer ID", RemQer.QERID.Value)
	return RemQer
}

func UpdateSmfN4ContextV1(
	reqN4 *pfcp.SessionModifyRequest, seid uint64) {
	rlogger.FuncEntry(types.ModuleSmfN4, &reqN4.PfcpHeader)
	// 读取当前上下文
	// *gctxt.N4SessionContext
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &reqN4.PfcpHeader, "gctxt.LengthOfN4ContextTbl:%v",
		gctxt.LengthOfN4ContextTbl(gctxt.N4SessionIDCxtType))

	n4Cxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, &reqN4.PfcpHeader, "Failed to get N4 Context:%s", err)
		return
	}

	// 填充SMF N4上下文
	//n4Cxt.PduCxt = Cxt
	n4Cxt.SEID = seid
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
	if reqN4.IE.UpdateBAR != nil {
		n4Cxt.BAR.BARID = reqN4.IE.UpdateBAR.BARID
		n4Cxt.BAR.DLDataNotificationDelay.Value = reqN4.IE.UpdateBAR.DLDataNotificationDelay.Value
		n4Cxt.BAR.SugBuffPacketsCount = reqN4.IE.UpdateBAR.SuggestedBufferingPacketsCount
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "UpdateSmfN4ContextV1 SMF Seid %+v", n4Cxt.SEID)
	// 更新N4上下文
	gctxt.UpdateN4Context(gctxt.N4SessionIDKey(n4Cxt.SEID), n4Cxt)
}
func UpdateSmfN4ContextV2(
	reqN4 *pfcp.SessionModifyRequest,
	resN4 *pfcp.SessionModifyResponse) {
	rlogger.FuncEntry(types.ModuleSmfN4, &reqN4.PfcpHeader)
	// 读取当前上下文
	// *gctxt.N4SessionContext
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &reqN4.PfcpHeader, "reqN4.PfcpHeader.SEID :%v", resN4.PfcpHeader.SEID)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &reqN4.PfcpHeader, "gctxt.LengthOfN4ContextTbl:%v",
		gctxt.LengthOfN4ContextTbl(gctxt.N4SessionIDCxtType))

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
