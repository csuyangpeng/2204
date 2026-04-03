package smutility

import (
	"context"
	"fmt"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/message/udmdata"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
	"lite5gc/smf/sc/sessmgnt"
	"lite5gc/smf/smfcontext/gctxt"
	"net"
)

type UpdateSmCtxtRespData struct {
	Imsi           types3gpp.Imsi
	Snssai         string
	Dnn            string
	Seid           uint64
	UpCnxState     n11msg.UpCnxState
	DirectRespFlag bool
}

func GenPduSessResSetupReqTransferMsg(ctxt context.Context, msgData *UpdateSmCtxtRespData) (error, []byte) {
	rlogger.FuncEntry(types.ModuleSmfSM, msgData)

	ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(msgData.Imsi.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, nil, "failed  to find ue context by imsi %s", msgData.Imsi.String())
		return types.ErrFailFindUeCtxt, nil
	}

	// get the session subscriber data from UDM
	var sessionSubMsg *udmdata.SessMgntSubscripitonData
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, msgData.Snssai)

	for _, v := range ueCtxt.SessMgntSubsDataMap {
		if msgData.Snssai == v.SingleNssai.String() {
			sessionSubMsg = &v
		}
	}
	if sessionSubMsg == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "fail to get sessionSubMsg")
		return fmt.Errorf("fail to get sessionSubMsg"), nil
	}
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, ueCtxt.SessMgntSubsDataMap)

	sessMgnt, ok := ctxt.Value(types.SessMgntCK).(*sessmgnt.SessMGMT)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed  to get session mangement")
		return fmt.Errorf("failed  to get session mangement"), nil
	}

	n2msg := ngapmsg.NewPduSessResSetupReqTransfer()
	n2msg.SetOssCodecCtxt(sessMgnt.GetOssCtxt().GetOssCtxtPtr_m())

	//subsData := sessionSubMsg.DnnConfigs[udmdata.DnnName(pduSessCtxt.DNN.String())]
	dnn, ok := sessionSubMsg.DnnConfigs[msgData.Dnn] //basic.com.mcc460.mnc001.gprs
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed  to get DnnConfigs by dnn name", msgData.Dnn)
		return fmt.Errorf("failed to get DnnConfigs by dnn name"), nil
	} else {
		n2msg.PduSessType = dnn.DefaultPDUSessionType.Convert2NgApType()

		qosflow := &types3gpp.QosFlowSetupReqest{}
		qosflow.QosFlowInd = uint(configure.SmfConf.Rules.QoSRules[0].QoSFlowIdentifier)
		qosflow.QosLevelParam.Arp.PriorityLevel = dnn.QosProf.Arp.PriorityLevel
		qosflow.QosLevelParam.Arp.PreemptCap = dnn.QosProf.Arp.PreemptCap
		qosflow.QosLevelParam.Arp.PreemptVuln = dnn.QosProf.Arp.PreemptVuln
		qosflow.QosLevelParam.QosChats.IsDynamic = false
		//qosflow.QosLevelParam.QosChats.NonDynamic5qi.FiveQI = dnn.QosProf.NonDynamic5qi.FiveQI
		qosflow.QosLevelParam.QosChats.NonDynamic5qi.FiveQI = uint(dnn.QosProf.QI5)

		//qosflow.QosLevelParam.QosChats.NonDynamic5qi.OptFlags.Set(types3gpp.ND5QI_AverageWind)
		//qosflow.QosLevelParam.QosChats.NonDynamic5qi.AverageWindow = dnn.QosProf.NonDynamic5qi.AverageWindow

		n2msg.QosFlowSetupReqList = []*types3gpp.QosFlowSetupReqest{qosflow}
	}

	c, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(msgData.Seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "failed to get GetN4Context by seid, %s", err)
		return fmt.Errorf("failed to GtpTunnel info(mandatory IE)"), nil
	} else {
		if c.PDRs[0].PDI.LocalFTEID != nil {
			n2msg.GtpTunnel.SetTeid(uint32(c.PDRs[0].PDI.LocalFTEID.TEID))
			n2msg.GtpTunnel.SetIpAddr(c.PDRs[0].PDI.LocalFTEID.IPv4Addr)
		} else {
			n2msg.GtpTunnel.SetTeid(0)
			n2msg.GtpTunnel.SetIpAddr(net.IP{})
		}
	}
	// N2 msg
	rlogger.Trace(types.ModuleSmfSM, rlogger.DEBUG, []interface{}{ueCtxt, msgData, c}, "gtp tunnel info : %s", n2msg.GtpTunnel.String())
	encodeMsg := n2msg.Encode()

	return nil, encodeMsg
}
