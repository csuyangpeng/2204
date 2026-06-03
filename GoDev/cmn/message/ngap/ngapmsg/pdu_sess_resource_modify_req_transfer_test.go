package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	T "lite5gc/cmn/types3gpp"
	"net"
	"testing"
)

func Test_PduSessResModifyReqTransfer_encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewPduSessResModifyReqTransfer()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	qosflow := &T.QosFlowAddOrModRequest{}

	qosflow.QosFlowInd = 5

	qosflow.IsERabIdPrst = true
	qosflow.ERabId = 9

	//qosflow.IsQosFlowParamPrst = true
	//qosflow.QosLevelParam.Arp.PriorityLevel = 2
	//qosflow.QosLevelParam.Arp.PreemptCap = T.NOT_PREEMPT
	//qosflow.QosLevelParam.Arp.PreemptVuln = T.PREEMPTABLE
	//
	//qosflow.QosLevelParam.QosChats.IsDynamic = false
	//qosflow.QosLevelParam.QosChats.NonDynamic5qi.FiveQI = 1
	//qosflow.QosLevelParam.QosChats.NonDynamic5qi.OptFlags.Set(T.ND5QI_PriorityLevelQos)
	//qosflow.QosLevelParam.QosChats.NonDynamic5qi.PriorityLevelQos = 200
	//qosflow.QosLevelParam.QosChats.NonDynamic5qi.OptFlags.Set(T.ND5QI_AverageWind)
	//qosflow.QosLevelParam.QosChats.NonDynamic5qi.AverageWindow = 500
	//qosflow.QosLevelParam.QosChats.NonDynamic5qi.OptFlags.Set(T.ND5QI_MaxDataBrstVol)
	//qosflow.QosLevelParam.QosChats.NonDynamic5qi.MaxDataBurstVol = 1000
	//
	//qosflow.QosLevelParam.OptFlags.Set(T.QFLQP_GbrInfo)
	//qosflow.QosLevelParam.GbrQosInfo.MaxFlowBitRateUL = 100000
	//qosflow.QosLevelParam.GbrQosInfo.MaxFlowBitRateDL = 110000
	//qosflow.QosLevelParam.GbrQosInfo.GuarantFlowBitRateUL = 200000
	//qosflow.QosLevelParam.GbrQosInfo.GuarantFlowBitRateDL = 220000
	//qosflow.QosLevelParam.GbrQosInfo.OptFlags.Set(T.GBRQI_NotifyCtrl)
	//qosflow.QosLevelParam.GbrQosInfo.NotifyCtrl = T.NotificationEnabled
	//qosflow.QosLevelParam.GbrQosInfo.OptFlags.Set(T.GBRQI_MaxPktLossRateUL)
	//qosflow.QosLevelParam.GbrQosInfo.MaxPktLossRateUL = 500
	//qosflow.QosLevelParam.GbrQosInfo.OptFlags.Set(T.GBRQI_MaxPktLossRateDL)
	//qosflow.QosLevelParam.GbrQosInfo.MaxPktLossRateDL = 600
	//
	//qosflow.QosLevelParam.OptFlags.Set(T.QFLQP_ReflecQosAttr)
	//qosflow.QosLevelParam.RefQosAttr = T.SubjectTo
	//
	//qosflow.QosLevelParam.OptFlags.Set(T.QFLQP_AddQosFlowInfo)
	//qosflow.QosLevelParam.AddQosFlowInfo = T.MoreLikely

	msg.AddQosFlowAddOrModReqList(qosflow)

	//qosflow1 := &T.QosFlowAddOrModRequest{}
	//
	//qosflow1.QosFlowInd = 6
	//
	//qosflow1.IsERabIdPrst = true
	//qosflow1.ERabId = 10
	//
	//qosflow1.QosLevelParam.Arp.PriorityLevel = 3
	//qosflow1.QosLevelParam.Arp.PreemptCap = T.NOT_PREEMPT
	//qosflow1.QosLevelParam.Arp.PreemptVuln = T.PREEMPTABLE
	//
	//qosflow1.QosLevelParam.QosChats.IsDynamic = true
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.PriorityLevelQos = 100
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.PktDelayBuget = 5000
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.PktErrRate.PktErrRateScalar = 5
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.PktErrRate.PktErrRateExponent = 10
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.OptFlags.Set(T.D5QI_FiveQI)
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.FiveQI = 8
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.OptFlags.Set(T.D5QI_DelayCri)
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.DelayCritical = T.DelayCri
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.OptFlags.Set(T.D5QI_AverageWind)
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.AverageWindow = 500
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.OptFlags.Set(T.D5QI_MaxDataBrstVol)
	//qosflow1.QosLevelParam.QosChats.Dynamic5qi.MaxDataBurstVol = 1000
	//
	//msg.AddQosFlowAddOrModReqList(qosflow1)

	tnlModInfo := &T.UlNguUpTnlModifyItem{}
	tnlModInfo.UplinkUpTransLayerInfo.SetIpAddr(net.ParseIP("10.180.9.242"))
	tnlModInfo.UplinkUpTransLayerInfo.SetTeid(9524)
	tnlModInfo.DownlinkUpTransLayerInfo.SetIpAddr(net.ParseIP("10.180.9.245"))
	tnlModInfo.DownlinkUpTransLayerInfo.SetTeid(9524)
	msg.AddUlNguUpTnlModifyList(tnlModInfo)
	//
	//relQosFlowList := &T.QosFlowDesc{}
	//relQosFlowList.QowFlowInd = 5
	//relQosFlowList.Cause.Type = T.CT_RadioNetwork
	//relQosFlowList.Cause.Value = T.Radiok_CauseRadioNetwork_unspecified
	//msg.AddQosFlowToReleaseList(relQosFlowList)

	msg.OptFlags.Set(PSRSRT_PduSessAMBR)
	msg.PduSessAMBR.Uplink = 20000
	msg.PduSessAMBR.Downlink = 500000

	msg.OptFlags.Set(PSRSRT_NetworkInst)
	msg.NetworkInst = 12

	fmt.Println(msg)
	fmt.Println("----------------------")
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

//func Test_PduSessResModifyReqTransfer_decode(t *testing.T) {
//
//	ossCtxt := codec.NewOssCtxt()
//	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
//	defer codec.DeleteOssCtxt(ossCtxt)
//
//	encMsg := []byte{0,0,7,0,136,0,72,5,6,2,243,28,2,19,136,11,1,10,0,8,0,1,244,0,3,232,9,192,4,127,255,255,253,82,10,225,192,1,128,2,0,200,0,1,244,0,3,232,5,192,4,127,255,255,253,113,0,1,173,176,32,1,134,160,32,3,91,96,32,3,13,64,0,2,88,0,1,244,18,0,129,0,2,0,11,0,138,0,2,1,0,0,134,0,1,16,0,127,0,1,0,0,139,0,22,7,240,10,180,9,242,55,37,0,0,48,244,13,1,0,0,0,0,55,37,0,0,0,130,0,7,8,7,161,32,16,78,32}
//	//fmt.Println(encMsg)
//
//	tansfer := NewPduSessResSetupReqTransfer()
//	tansfer.SetOssCodecCtxt(ossCtxtPtr)
//	decode := tansfer.Decode(encMsg)
//	fmt.Println("result", decode)
//	fmt.Println("decoded msg:", tansfer)
//}
