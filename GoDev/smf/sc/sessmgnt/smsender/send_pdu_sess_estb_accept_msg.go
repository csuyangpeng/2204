package smsender

import (
	"context"
	"fmt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
	"lite5gc/openapi/models"
	"lite5gc/smf/sc/sessmgnt"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
	"net"
)

func SendPduSessEstbAcceptMsg(ctxt context.Context,
	prcdCtxt *prcdctxt.PduSessionEstbPrcdCtxt) error {
	td := []interface{}{ctxt, prcdCtxt}
	rlogger.FuncEntry(types.ModuleSmfSM, td)
	// SMF负责检查UE的会话请求是否和其订阅数据一致

	// Encode NAS Msg: pdu session establishment accept msg
	var pduSessEstbAcceptMsg nasmsg.PduSessionEstbAcceptMsg
	//set Accept's EPD
	pduSessEstbAcceptMsg.ExtendProtoDisc = nas.Epd5gsSessMgntMsg

	//set msgHeader
	pduSessEstbAcceptMsg.MsgHeader.PduSessionID = prcdCtxt.PduSessId
	pduSessEstbAcceptMsg.MsgHeader.PrcdTransactionID = prcdCtxt.Pti
	pduSessEstbAcceptMsg.MsgHeader.MessageType = nas.PduSessEstabishAccept

	// mandatory IEs
	pduSessEstbAcceptMsg.SessionType = prcdCtxt.PduSessType

	pduSessEstbAcceptMsg.SscMode = prcdCtxt.SscMode

	//set the Authorized QoS rules
	pduSessEstbAcceptMsg.AuthorizedQoSRules = configure.SmfConf.Rules

	pduSessEstbAcceptMsg.SessionAMBR = prcdCtxt.SessionAmbr

	//optional IEs

	//set the DNN
	pduSessEstbAcceptMsg.IeFlags.Set(nasmsg.Ieid_PduSessionEstbAcpt_DNN)
	pduSessEstbAcceptMsg.DNN = prcdCtxt.SelectDnn
	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, td, "Dnn(%s)", pduSessEstbAcceptMsg.DNN)

	//set the selected SSC mode
	if !prcdCtxt.SessionReqIeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_MaxNumberOfSPF) {
		prcdCtxt.MaxNumofSupPktFilter = 16
	}

	// deal "Maximum number of supported packet filters"
	// if "Maximum number of supported packet filters" exist,
	// SMF consider this number as the maximum number of packet filters.
	// Otherwise the SMF considers that the UE supports 16 packet filters for this PDU session.
	if !prcdCtxt.SessionReqIeFlags.Test(nasmsg.Ieid_PduSessionEstbReq_MaxNumberOfSPF) {
		prcdCtxt.MaxNumofSupPktFilter = 16
	}
	// The SMF shall ensure that the number of the packet filters used in the authorized QoS rules of the PDU Session
	// does not exceed the maximum number of packet filters supported by the UE for the PDU session.
	var sumOfnumber byte
	for i := 0; i < len(configure.SmfConf.Rules.QoSRules); i++ {
		sumOfnumber += configure.SmfConf.Rules.QoSRules[i].NumberOfPacketFilters
	}
	if uint16(sumOfnumber) > prcdCtxt.MaxNumofSupPktFilter {
		//读配置的时候，先将rules按照优先级排序，再从高到低取packet filter，超过max的就截掉，不要了
		//TODO
	}

	// dnn
	if prcdCtxt.DnnConfig != nil {
		//set Session AMBR
		pduSessEstbAcceptMsg.SessionAMBR = prcdCtxt.DnnConfig.SessionAmbr
		if prcdCtxt.UEIP != nil {
			pduSessEstbAcceptMsg.PDUaddress.PDUSessionType = types3gpp.Ipv4
			pduSessEstbAcceptMsg.PDUaddress.PDUAddressInfor = prcdCtxt.UEIP
			pduSessEstbAcceptMsg.IeFlags.Set(nasmsg.Ieid_PduSessionEstbAcpt_PDUaddress)
		}
	} else {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "no dnn config in procedure context")
	}

	//set 5GSM cause
	//pduSessEstbAcceptMsg.SMCause = nas.PDUSessionTypeIPv4OnlyAllowed

	//set PDU address
	//pduSessEstbAcceptMsg.PDUaddress.PDUSessionType = t3.Ipv4

	//set the S-NSSAI of the PDU session
	pduSessEstbAcceptMsg.SNSSAI = prcdCtxt.SNSSAI
	pduSessEstbAcceptMsg.IeFlags.Set(nasmsg.Ieid_PduSessionEstbAcpt_SNSSAI)

	//set Always-on PDU session indication
	//UE:  Always-on PDU session not requested or Always-on PDU session requested
	//SMF: Always-on PDU session not allowed   or Always-on PDU session required
	//if the UE did not include requested IE, the SMF shall not include the indication IE
	//pduSessEstbAcceptMsg.AlwaysOn = prcdCtxt.AlwaysOnPduSessReq

	//set Authorized QoS flow descriptions
	pduSessEstbAcceptMsg.IeFlags.Set(nasmsg.Ieid_PduSessionEstbAcpt_AuthorizedQoSFlowDescriptions)
	pduSessEstbAcceptMsg.AuthorizedQoSFlowDescriptions = configure.SmfConf.FlowDescr

	// Handle "Integrity protection maximum data rate" from request msg
	// TODO, for user-plane integrity protection

	// 5GSM capability
	// 1 RqoS:UE supports Reflective QoS or not.
	// if no ,set the RQ timer to "deactivated" or has a value of zero
	// 2 MH6-PDU:UE supports multi-homed IPv6 PDU session, false
	/*if prcdCtxt.SmCap.RqoS == true && smfConfMgr.GetRqoS() {
		var rqTimer nasie.GprsTimer
		rqTimer.Uint = nasie.Deactivated
		rqTimer.TimerValue = 0
		// RQ timer value, for UE's Reflective Qos
		pduSessEstbAcceptMsg.RQTimerValue = rqTimer
	}*/

	//N1 msg
	bytesPduSessEstbAcceptMsg, err := pduSessEstbAcceptMsg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "fail to encode pdu session accept msg")
		return fmt.Errorf("fail to encode pdu session accept msg")
	}

	//msg counter
	pm.PegCounter(statistics.PduSessEstablishAcceptCounter)

	//Encode N11 message: n1n2message transfer request
	var msgData n11msg.N1N2MessageTransferReqData
	N1Container := n11msg.N1MessageContainerIE{N1MsgClass: n11msg.SM_N1Info, N1MessageContent: bytesPduSessEstbAcceptMsg}
	// set n1 to msg data
	msgData.N1MessageContainer = &N1Container
	msgData.IeFlags.Set(n11msg.Ieid_n1MessageContainer)
	// set snssai to msg data
	msgData.SessionId = prcdCtxt.PduSessId
	msgData.IeFlags.Set(n11msg.Ieid_pdusessionId)

	// Encode N2 message
	n2Info := &n11msg.N2SmInformation{}
	n2Info.PduSessionId = prcdCtxt.PduSessId
	n2Info.IeFlags.Set(n11msg.Ieid_pdusessionId)
	n2Info.Snssai = prcdCtxt.SNSSAI
	n2Info.IeFlags.Set(n11msg.Ieid_sNssai)

	n2msg := ngapmsg.NewPduSessResSetupReqTransfer()

	sessMgnt, ok := ctxt.Value(types.SessMgntCK).(*sessmgnt.SessMGMT)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "failed to get session mangement")
		return fmt.Errorf("failed to get session mangement")
	}
	n2msg.SetOssCodecCtxt(sessMgnt.GetOssCtxt().GetOssCtxtPtr_m())

	dnnConfig := prcdCtxt.DnnConfig
	if dnnConfig != nil {
		n2msg.PduSessType = dnnConfig.DefaultPDUSessionType.Convert2NgApType()

		qosflow := &types3gpp.QosFlowSetupReqest{}
		qosflow.QosFlowInd = 1
		//qosflow.QosFlowInd = uint(configure.GetQosRlue1Qfi())
		qosflow.QosLevelParam.Arp.PriorityLevel = dnnConfig.QosProf.Arp.PriorityLevel
		qosflow.QosLevelParam.Arp.PreemptCap = dnnConfig.QosProf.Arp.PreemptCap
		qosflow.QosLevelParam.Arp.PreemptVuln = dnnConfig.QosProf.Arp.PreemptVuln
		qosflow.QosLevelParam.QosChats.IsDynamic = false
		// before:
		// qosflow.QosLevelParam.QosChats.NonDynamic5qi.FiveQI = dnnConfig.QosProf.NonDynamic5qi.FiveQI
		// after:
		qosflow.QosLevelParam.QosChats.NonDynamic5qi.FiveQI = uint(dnnConfig.QosProf.QI5)
		//qosflow.QosLevelParam.QosChats.NonDynamic5qi.OptFlags.Set(types3gpp.ND5QI_AverageWind)
		//qosflow.QosLevelParam.QosChats.NonDynamic5qi.AverageWindow = dnnConfig.QosProf.NonDynamic5qi.AverageWindow

		n2msg.QosFlowSetupReqList = []*types3gpp.QosFlowSetupReqest{qosflow}

		n2msg.PduSessAMBR.Uplink, err = types3gpp.StoreAmbrBitRate(prcdCtxt.SessionAmbr.UpLink.String())

		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "failed to set Sesson AMBR uplink")
		}
		n2msg.PduSessAMBR.Downlink, err = types3gpp.StoreAmbrBitRate(prcdCtxt.SessionAmbr.DownLink.String())
		if err != nil {
			rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "failed to set Sesson AMBR downlink")
		}
		rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, td, "n2msg.PduSessAMBR", n2msg.PduSessAMBR)

		n2msg.OptFlags.Set(ngapmsg.PSRSRT_PduSessAMBR)

	} else {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "no dnn config in procedure context")
	}

	c, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(prcdCtxt.Seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, td, "fail to get GetN4Context by seid, %s", err)
		return fmt.Errorf("fail to get GetN4Context by seid, %s", err)
	} else {
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, td, "c.PDRs[0].PDI.LocalFTEID", c.PDRs[0].PDI.LocalFTEID, c)
		if c.PDRs[0].PDI.LocalFTEID != nil {
			n2msg.GtpTunnel.SetTeid(uint32(c.PDRs[0].PDI.LocalFTEID.TEID))
			n2msg.GtpTunnel.SetIpAddr(c.PDRs[0].PDI.LocalFTEID.IPv4Addr)
		} else {
			n2msg.GtpTunnel.SetTeid(0)
			n2msg.GtpTunnel.SetIpAddr(net.IP{})
		}
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, []interface{}{ctxt, prcdCtxt, c}, "gtp tunnel info : %s", n2msg.GtpTunnel.String())
	}

	n2Info.N2InfoCont = &n11msg.N2InfoContent{}
	n2Info.N2InfoCont.NgapIeType = n11msg.NgapPduResSetupReq
	encodeMsg := n2msg.Encode()
	n2Info.N2InfoCont.NgapData = append(n2Info.N2InfoCont.NgapData, encodeMsg...)
	n2Info.IeFlags.Set(n11msg.Ieid_n2InfoContent)

	//set n2 info msg into msg data
	N2Container := &n11msg.N2InfoContainerIE{N2InforClass: n11msg.SM_N2Info, SmInfo: n2Info}
	N2Container.IeFlags.Set(n11msg.Ieid_n2SmInfo)
	msgData.N2InfoContainer = N2Container
	msgData.IeFlags.Set(n11msg.Ieid_n2InfoContainer)

	requestData := &models.N1N2MessageTransferRequest{}
	requestData.JsonData = sbicmn.Trans_N11ToModels_N1N2MsgTransferReqFormat(&msgData)
	requestData.BinaryDataN1Message = msgData.N1MessageContainer.N1MessageContent
	requestData.BinaryDataN2Information = msgData.N2InfoContainer.SmInfo.N2InfoCont.NgapData
	SendN1N2MsgTransfer(ctxt, requestData, prcdCtxt.IMSI.AddIMSIPrefix(), prcdCtxt.PduSessId, sessMgnt.GetScId())

	return nil
}
