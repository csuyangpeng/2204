package ngapmsg

import "C"
import (
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	T "lite5gc/cmn/types3gpp"

	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
	"net"
	"unsafe"
)

// PduSessResModifyReqTransfer struct definition
// 38.413 9.3.4.1
type PduSessResModifyReqTransfer struct {
	// 0..1
	UlNguUpTnlModifyList []*T.UlNguUpTnlModifyItem
	// 0..1
	QosFlowAddOrModReqList []*T.QosFlowAddOrModRequest
	// 0..1
	QosFlowToReleaseList []*T.QosFlowDesc
	//optional
	AddGtpTunnelList      []*T.GtpTunnel
	PduSessAMBR           T.Ambr
	NetworkInst           int
	CommonNetworkInstance string

	OptFlags bitset.BitSet
	ctxt     codec.NgapOssCtxt
}

const (
	PSRMRT_PduSessAMBR = iota
	PSRMRT_NetworkInst
	PSRMRT_AddGtpTunnelList
	PSRMRT_CommonNetworkInstance
)

func (p *PduSessResModifyReqTransfer) String() string {
	outStr := fmt.Sprintf("PduSessResModifyReqTransfer{"+
		"[%v SessAMBR(%s)],NetworkInst(%v,%d),CommonNetwork(%v,%s),",
		p.OptFlags.Test(PSRMRT_PduSessAMBR), p.PduSessAMBR,
		p.OptFlags.Test(PSRMRT_NetworkInst), p.NetworkInst,
		p.OptFlags.Test(PSRMRT_CommonNetworkInstance), p.CommonNetworkInstance)

	outStr += " UlNguUpTnlModifyList( "
	for _, v := range p.UlNguUpTnlModifyList {
		outStr += fmt.Sprintf("%s,", *v)
	}

	outStr += ") QosFlowAddOrModReqList( "
	for _, v := range p.QosFlowAddOrModReqList {
		outStr += fmt.Sprintf("%s,", *v)
	}

	outStr += ") QosFlowToReleaseList( "
	for _, v := range p.QosFlowToReleaseList {
		outStr += fmt.Sprintf("%s,", v)
	}
	outStr += fmt.Sprintf(") AddGtpTunnelList(%v),(", p.OptFlags.Test(PSRMRT_AddGtpTunnelList))
	for _, v := range p.AddGtpTunnelList {
		outStr += fmt.Sprintf("%+v,", v)
	}

	outStr += "}"
	return outStr
}

// NewPduSessResModifyReqTransfer create a new Message
func NewPduSessResModifyReqTransfer() *PduSessResModifyReqTransfer {

	return &PduSessResModifyReqTransfer{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResModifyReqTransfer) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *PduSessResModifyReqTransfer) AddQosFlowAddOrModReqList(qosFlow *T.QosFlowAddOrModRequest) {
	p.QosFlowAddOrModReqList = append(p.QosFlowAddOrModReqList, qosFlow)
}

func (p *PduSessResModifyReqTransfer) AddUlNguUpTnlModifyList(qosFlow *T.UlNguUpTnlModifyItem) {
	p.UlNguUpTnlModifyList = append(p.UlNguUpTnlModifyList, qosFlow)
}

func (p *PduSessResModifyReqTransfer) AddQosFlowToReleaseList(qosFlow *T.QosFlowDesc) {
	p.QosFlowToReleaseList = append(p.QosFlowToReleaseList, qosFlow)
}

func (p *PduSessResModifyReqTransfer) AddAddGtpTunnelList(qosFlow *T.GtpTunnel) {
	p.AddGtpTunnelList = append(p.AddGtpTunnelList, qosFlow)
}

func (p *PduSessResModifyReqTransfer) Encode() []byte {
	rlogger.FuncEntry(types.ModCmn, nil)
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "Encode PduSessResModifyReqTransfer(%s)", p)

	transfer := codec.NewPduSesResModReqTransferCodec()
	defer codec.DeletePduSesResModReqTransferCodec(transfer)

	// pdu session ambr
	if p.OptFlags.Test(PSRSRT_PduSessAMBR) {
		transfer.SetSessAmbr(int64(p.PduSessAMBR.Uplink), int64(p.PduSessAMBR.Downlink))
	}

	// Network instance
	if p.OptFlags.Test(PSRSRT_NetworkInst) {
		transfer.SetNtwkInstance(uint64(p.NetworkInst))
	}

	//Common Network Instance
	if p.OptFlags.Test(PSRMRT_CommonNetworkInstance) {
		transfer.SetCmmnNtwrkInstance(p.CommonNetworkInstance)
	}

	// ul ngu up tnl modify
	for _, v := range p.UlNguUpTnlModifyList {
		ulNguUpTnlModItem := codec.NewUlNguUpTnlModifyItem()
		defer codec.DeleteUlNguUpTnlModifyItem(ulNguUpTnlModItem)

		upIpaddr := make([]byte, 4)
		binary.BigEndian.PutUint32(upIpaddr, v.UplinkUpTransLayerInfo.GetIpv4Long())
		upTeid := make([]byte, 4)
		binary.BigEndian.PutUint32(upTeid, v.UplinkUpTransLayerInfo.GetTeid())

		uplinkGtpTunnel := codec.NewGtpTunnel()
		defer codec.DeleteGtpTunnel(uplinkGtpTunnel)
		uplinkGtpTunnel.SetIpType(0) //todo
		uplinkGtpTunnel.SetTransportLayerAddr(&upIpaddr[0])
		uplinkGtpTunnel.SetGtpTeid(&upTeid[0])
		ulNguUpTnlModItem.SetUplinkGtpTunnel(uplinkGtpTunnel)

		downIpaddr := make([]byte, 4)
		binary.BigEndian.PutUint32(downIpaddr, v.DownlinkUpTransLayerInfo.GetIpv4Long())
		downTeid := make([]byte, 4)
		binary.BigEndian.PutUint32(downTeid, v.DownlinkUpTransLayerInfo.GetTeid())
		downlinkGtpTunnel := codec.NewGtpTunnel()
		defer codec.DeleteGtpTunnel(downlinkGtpTunnel)
		downlinkGtpTunnel.SetIpType(0) //todo
		downlinkGtpTunnel.SetTransportLayerAddr(&downIpaddr[0])
		downlinkGtpTunnel.SetGtpTeid(&downTeid[0])
		ulNguUpTnlModItem.SetDownlinkGtpTunnel(downlinkGtpTunnel)

		transfer.AddUlNguUpTnlModifyList(ulNguUpTnlModItem)
	}

	//additional ul ng tnl information
	for _, v := range p.AddGtpTunnelList {
		addGtpTunnelItem := codec.NewGtpTunnel()
		defer codec.DeleteGtpTunnel(addGtpTunnelItem)

		upIpaddr := make([]byte, 4)
		binary.BigEndian.PutUint32(upIpaddr, v.GetIpv4Long())
		upTeid := make([]byte, 4)
		binary.BigEndian.PutUint32(upTeid, v.GetTeid())

		additionalGtpTunnel := codec.NewGtpTunnel()
		defer codec.DeleteGtpTunnel(additionalGtpTunnel)
		additionalGtpTunnel.SetIpType(0) //todo
		additionalGtpTunnel.SetTransportLayerAddr(&upIpaddr[0])
		additionalGtpTunnel.SetGtpTeid(&upTeid[0])

		transfer.AddAdditionalGtpTunnelList(additionalGtpTunnel)
	}

	for _, v := range p.QosFlowAddOrModReqList {
		qosFlowItem := codec.NewQosFlowAddOrModReqItem()
		defer codec.DeleteQosFlowAddOrModReqItem(qosFlowItem)

		// qos flow indicator
		qosFlowItem.SetQosFlowInd(v.QosFlowInd)

		// qos flow e rab id
		if v.IsERabIdPrst {
			qosFlowItem.SetERABIdPresent(true)
			qosFlowItem.SetERABId(v.ERabId)
		}

		// qos flow level qos parameter
		if v.IsQosFlowParamPrst {
			qosFlowPara := codec.NewQosFlowLevelQosPara()
			defer codec.DeleteQosFlowLevelQosPara(qosFlowPara)

			// qos flow para - arp
			arp := codec.NewARP()
			defer codec.DeleteARP(arp)
			arp.SetEmptionCapability(byte(v.QosLevelParam.Arp.PreemptCap))
			arp.SetEmptionVulnerability(byte(v.QosLevelParam.Arp.PreemptVuln))
			arp.SetPriorityLevel(v.QosLevelParam.Arp.PriorityLevel)
			qosFlowPara.SetArp(arp)

			// qos flow para = qos character
			qosCc := codec.NewQosCharacter()
			defer codec.DeleteQosCharacter(qosCc)
			vQosCc := &(v.QosLevelParam.QosChats)
			if vQosCc.IsDynamic {
				qosCc.SetIsDynamic(true)
				dynamic5qi := codec.NewDynamic5QI()
				defer codec.DeleteDynamic5QI(dynamic5qi)
				dynamic5qi.SetPriorityLevelQos(int64(vQosCc.Dynamic5qi.PriorityLevelQos))
				dynamic5qi.SetPacketDelayBudget(int64(vQosCc.Dynamic5qi.PktDelayBuget))
				pktErrRat := codec.NewPktErrRate()
				defer codec.DeletePktErrRate(pktErrRat)
				pktErrRat.SetPERScalar(int64(vQosCc.Dynamic5qi.PktErrRate.PktErrRateScalar))
				pktErrRat.SetPERExponent(int64(vQosCc.Dynamic5qi.PktErrRate.PktErrRateExponent))
				dynamic5qi.SetPacketErrorRate(pktErrRat)
				if vQosCc.Dynamic5qi.OptFlags.Test(T.D5QI_FiveQI) {
					dynamic5qi.SetFiveQI(int64(vQosCc.Dynamic5qi.FiveQI))
					dynamic5qi.SetIs5QIPrst(true)
				}
				if vQosCc.Dynamic5qi.OptFlags.Test(T.D5QI_DelayCri) {
					dynamic5qi.SetDelayCritical(byte(vQosCc.Dynamic5qi.DelayCritical))
					dynamic5qi.SetIsDelayCriticalPrst(true)
				}
				if vQosCc.Dynamic5qi.OptFlags.Test(T.D5QI_AverageWind) {
					dynamic5qi.SetAveragingWindow(int64(vQosCc.Dynamic5qi.AverageWindow))
					dynamic5qi.SetIsAveWindowPrst(true)
				}
				if vQosCc.Dynamic5qi.OptFlags.Test((T.D5QI_MaxDataBrstVol)) {
					dynamic5qi.SetMaximumDataBurstVolume(int64(vQosCc.Dynamic5qi.MaxDataBurstVol))
					dynamic5qi.SetIsMaxDataBusrtVolPrst(true)
				}
				qosCc.SetDynamic5qi(dynamic5qi)
			} else {
				qosCc.SetIsDynamic(false)
				nonDynamic5qi := codec.NewNonDynamic5QI()
				defer codec.DeleteNonDynamic5QI(nonDynamic5qi)
				nonDynamic5qi.SetFiveQI(int64(vQosCc.NonDynamic5qi.FiveQI))
				if vQosCc.NonDynamic5qi.OptFlags.Test(T.ND5QI_AverageWind) {
					nonDynamic5qi.SetAveragingWindow(int64(vQosCc.NonDynamic5qi.AverageWindow))
					nonDynamic5qi.SetIsAveWindowPrst(true)
				}
				if vQosCc.NonDynamic5qi.OptFlags.Test(T.ND5QI_PriorityLevelQos) {
					nonDynamic5qi.SetIsPriorityLevelQosPrst(true)
					nonDynamic5qi.SetPriorityLevelQos(int64(vQosCc.NonDynamic5qi.PriorityLevelQos))
				}
				if vQosCc.NonDynamic5qi.OptFlags.Test(T.ND5QI_MaxDataBrstVol) {
					nonDynamic5qi.SetIsMaxDataBusrtVolPrst(true)
					nonDynamic5qi.SetMaximumDataBurstVolume(int64(vQosCc.NonDynamic5qi.MaxDataBurstVol))
				}
				qosCc.SetNonDynamic5qi(nonDynamic5qi)
			}
			qosFlowPara.SetQosCharacter(qosCc)

			// qos flow para - gbr qos info
			if v.QosLevelParam.OptFlags.Test(T.QFLQP_GbrInfo) {
				gbr := &(v.QosLevelParam.GbrQosInfo)
				qosFlowPara.SetGBRQosInfoPresent(true)

				gbrInfo := codec.NewGBRQosInformation()
				defer codec.DeleteGBRQosInformation(gbrInfo)
				gbrInfo.SetMaxFlowBitRateDL(int64(gbr.MaxFlowBitRateDL))
				gbrInfo.SetMaxFlowBitRateUL(int64(gbr.MaxFlowBitRateUL))
				gbrInfo.SetGuaFlowBitRateDL(int64(gbr.GuarantFlowBitRateDL))
				gbrInfo.SetGuaFlowBitRateUL(int64(gbr.GuarantFlowBitRateUL))
				if gbr.OptFlags.Test(T.GBRQI_MaxPktLossRateDL) {
					gbrInfo.SetMaxPacketLossRateDLPresent(true)
					gbrInfo.SetMaxPacketLossRateDL(uint(gbr.MaxPktLossRateDL))
				}
				if gbr.OptFlags.Test(T.GBRQI_MaxPktLossRateUL) {
					gbrInfo.SetMaxPacketLossRateULPresent(true)
					gbrInfo.SetMaxPacketLossRateUL(uint(gbr.MaxPktLossRateUL))
				}
				if gbr.OptFlags.Test(T.GBRQI_NotifyCtrl) {
					gbrInfo.SetNotiControlPresent(true)
					gbrInfo.SetNotiControl(byte(gbr.NotifyCtrl))
				}
				qosFlowPara.SetGBRQosInfo(gbrInfo)
			}

			//qos flow para - ref qos atr
			if v.QosLevelParam.OptFlags.Test(T.QFLQP_ReflecQosAttr) {
				qosFlowPara.SetRefQosAttrPresent(true)
				qosFlowPara.SetRefQosAttr(byte(v.QosLevelParam.RefQosAttr))
			}

			//qos flow para - add qos flow info
			if v.QosLevelParam.OptFlags.Test(T.QFLQP_AddQosFlowInfo) {
				qosFlowPara.SetAddQosFlowInfoPresent(true)
				qosFlowPara.SetAddQosFlowInfo(byte(v.QosLevelParam.AddQosFlowInfo))
			}
			qosFlowItem.SetQosFlowParaPresent(true)
			qosFlowItem.SetQosFlowLevQosPara(qosFlowPara)
		}

		transfer.AddQosFlowAddOrModReqList(qosFlowItem)
	}

	for _, v := range p.QosFlowToReleaseList {
		qosflow := codec.NewQosFlowCodecItem()
		defer codec.DeleteQosFlowCodecItem(qosflow)

		qosflow.SetQosFlowInd(uint(v.QowFlowInd))
		qosflow.SetCauseType(byte(v.Cause.Type))
		qosflow.SetCauseValue(byte(v.Cause.Value))

		transfer.AddQosFlowToReleaseList(qosflow)
	}

	msgBuffer := transfer.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResModifyReqTransfer) Decode(msgbuf []byte) error {

	transfer := codec.NewPduSesResModReqTransferCodec()
	defer codec.DeletePduSesResModReqTransferCodec(transfer)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if transfer.Decode(p.ctxt, msgBuffer) == true {
		// session ambr
		if transfer.IsSessAMBRPrst() {
			p.OptFlags.Set(PSRSRT_PduSessAMBR)
			p.PduSessAMBR.Uplink = uint64(transfer.GetSessAmbrUl())
			p.PduSessAMBR.Downlink = uint64(transfer.GetSessAmbrDl())
		}

		// network instance
		if transfer.IsNtwkInstancePrst() {
			p.OptFlags.Set(PSRSRT_NetworkInst)
			p.NetworkInst = int(transfer.GetNtwkInstance())
		}

		// Ul Ngu Up Tnl Modify List
		tnlModifyList := transfer.GetUlNguUpTnlModifyList()
		for i := 0; i < int(tnlModifyList.Size()); i++ {
			tnlModifyItem := tnlModifyList.Get(i)
			pTnlModifyInfo := &T.UlNguUpTnlModifyItem{}

			if tnlModifyItem.GetUplinkGtpTunnel().GetIpType() == 0 { //todo ipv4-0 ipv6-1
				ipaddr := utils.Conv2ByteSlice(tnlModifyItem.GetUplinkGtpTunnel().GetTransportLayerAddr(), 4)
				pTnlModifyInfo.UplinkUpTransLayerInfo.SetIpAddr(net.IPv4(ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]))
			} else {
				rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ipv6 unsupported") // todo set ipv6
			}
			teid := utils.Conv2ByteSlice(tnlModifyItem.GetUplinkGtpTunnel().GetGtpTeid(), 4)
			pTnlModifyInfo.UplinkUpTransLayerInfo.SetTeid(binary.BigEndian.Uint32(teid))

			if tnlModifyItem.GetDownlinkGtpTunnel().GetIpType() == 0 { //todo ipv4-0 ipv6-1
				ipaddr := utils.Conv2ByteSlice(tnlModifyItem.GetDownlinkGtpTunnel().GetTransportLayerAddr(), 4)
				pTnlModifyInfo.DownlinkUpTransLayerInfo.SetIpAddr(net.IPv4(ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]))
			} else {
				rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ipv6 unsupported") // todo set ipv6
			}
			teid = utils.Conv2ByteSlice(tnlModifyItem.GetDownlinkGtpTunnel().GetGtpTeid(), 4)
			pTnlModifyInfo.DownlinkUpTransLayerInfo.SetTeid(binary.BigEndian.Uint32(teid))

			p.AddUlNguUpTnlModifyList(pTnlModifyInfo)
		}

		//additional ul ng-u up tnl information
		if transfer.IsAddGtpTunnelPrst() {
			p.OptFlags.Set(PSRMRT_AddGtpTunnelList)
			additionalGtpTunnelList := transfer.GetAdditionalGtpTunnelList()
			for i := 0; i < int(additionalGtpTunnelList.Size()); i++ {
				additionalGtpTunnelItem := additionalGtpTunnelList.Get(i)
				additionalGtpTunnelInfo := &T.GtpTunnel{}

				if additionalGtpTunnelItem.GetIpType() == 0 { //todo ipv4-0 ipv6-1
					ipaddr := utils.Conv2ByteSlice(additionalGtpTunnelItem.GetTransportLayerAddr(), 4)
					additionalGtpTunnelInfo.SetIpAddr(net.IPv4(ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]))
				} else {
					rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ipv6 unsupported") // todo set ipv6
				}
				teid := utils.Conv2ByteSlice(additionalGtpTunnelItem.GetGtpTeid(), 4)
				additionalGtpTunnelInfo.SetTeid(binary.BigEndian.Uint32(teid))

				if additionalGtpTunnelItem.GetIpType() == 0 { //todo ipv4-0 ipv6-1
					ipaddr := utils.Conv2ByteSlice(additionalGtpTunnelItem.GetTransportLayerAddr(), 4)
					additionalGtpTunnelInfo.SetIpAddr(net.IPv4(ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]))
				} else {
					rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ipv6 unsupported") // todo set ipv6
				}
				teid = utils.Conv2ByteSlice(additionalGtpTunnelItem.GetGtpTeid(), 4)
				additionalGtpTunnelInfo.SetTeid(binary.BigEndian.Uint32(teid))

				p.AddAddGtpTunnelList(additionalGtpTunnelInfo)
			}
		}

		//Common Network Instance
		if transfer.IsCmmnNtwrkInstacePrst() {
			p.OptFlags.Set(PSRMRT_CommonNetworkInstance)
			p.CommonNetworkInstance = transfer.GetCmmnNtwrkInstance()
		}

		// qos flow add or modify request list
		qosFlowList := transfer.GetQosFlowAddOrModReqList()
		for i := 0; i < int(qosFlowList.Size()); i++ {
			qosflow := qosFlowList.Get(i)

			pQosFlow := &T.QosFlowAddOrModRequest{}

			//qfi
			pQosFlow.QosFlowInd = qosflow.GetQosFlowInd()
			if qosflow.GetERABIdPresent() {
				pQosFlow.ERabId = qosflow.GetERABId()
				pQosFlow.IsERabIdPrst = true
			}

			//arp
			qosParam := qosflow.GetQosFlowLevQosPara()
			pQosFlow.QosLevelParam.Arp.PriorityLevel = qosParam.GetArp().GetPriorityLevel()
			pQosFlow.QosLevelParam.Arp.PreemptCap =
				T.PreemptionCapability(qosParam.GetArp().GetEmptionCapability())
			pQosFlow.QosLevelParam.Arp.PreemptVuln =
				T.PreemptionVulnerability(qosParam.GetArp().GetEmptionVulnerability())

			//qos characteristic
			qosCc := qosParam.GetQosCharacter()
			pQosCc := &pQosFlow.QosLevelParam.QosChats
			if qosCc.GetIsDynamic() {
				dynamic5qi := qosCc.GetDynamic5qi()
				pQosCc.IsDynamic = true
				pQosCc.Dynamic5qi.PriorityLevelQos = uint(dynamic5qi.GetPriorityLevelQos())
				pQosCc.Dynamic5qi.PktDelayBuget = uint(dynamic5qi.GetPacketDelayBudget())
				pQosCc.Dynamic5qi.PktErrRate.PktErrRateScalar =
					uint(dynamic5qi.GetPacketErrorRate().GetPERScalar())
				pQosCc.Dynamic5qi.PktErrRate.PktErrRateExponent =
					uint(dynamic5qi.GetPacketErrorRate().GetPERExponent())
				if dynamic5qi.GetIs5QIPrst() {
					pQosCc.Dynamic5qi.OptFlags.Set(T.D5QI_FiveQI)
					pQosCc.Dynamic5qi.FiveQI = uint(dynamic5qi.GetFiveQI())
				}
				if dynamic5qi.GetIsDelayCriticalPrst() {
					pQosCc.Dynamic5qi.OptFlags.Set(T.D5QI_DelayCri)
					pQosCc.Dynamic5qi.DelayCritical = T.DelayCritical(dynamic5qi.GetDelayCritical())
				}
				if dynamic5qi.GetIsAveWindowPrst() {
					pQosCc.Dynamic5qi.OptFlags.Set(T.D5QI_AverageWind)
					pQosCc.Dynamic5qi.AverageWindow = uint(dynamic5qi.GetAveragingWindow())
				}
				if dynamic5qi.GetIsMaxDataBusrtVolPrst() {
					pQosCc.Dynamic5qi.OptFlags.Set(T.D5QI_MaxDataBrstVol)
					pQosCc.Dynamic5qi.MaxDataBurstVol = uint(dynamic5qi.GetMaximumDataBurstVolume())
				}
			} else {
				nonDynamic5qi := qosCc.GetNonDynamic5qi()
				pQosCc.IsDynamic = false
				pQosCc.NonDynamic5qi.FiveQI = uint(nonDynamic5qi.GetFiveQI())
				if nonDynamic5qi.GetIsPriorityLevelQosPrst() {
					pQosCc.NonDynamic5qi.OptFlags.Set(T.ND5QI_PriorityLevelQos)
					pQosCc.NonDynamic5qi.PriorityLevelQos = uint(nonDynamic5qi.GetPriorityLevelQos())
				}
				if nonDynamic5qi.GetIsAveWindowPrst() {
					pQosCc.NonDynamic5qi.OptFlags.Set(T.ND5QI_AverageWind)
					pQosCc.NonDynamic5qi.AverageWindow = uint(nonDynamic5qi.GetAveragingWindow())
				}
				if nonDynamic5qi.GetIsMaxDataBusrtVolPrst() {
					pQosCc.NonDynamic5qi.OptFlags.Set(T.ND5QI_MaxDataBrstVol)
					pQosCc.NonDynamic5qi.MaxDataBurstVol = uint(nonDynamic5qi.GetMaximumDataBurstVolume())
				}
			}

			// GBR info
			if qosParam.GetGBRQosInfoPresent() {
				gbrInfo := qosParam.GetGBRQosInfo()

				pQosFlow.QosLevelParam.OptFlags.Test(T.QFLQP_GbrInfo)
				pGbrInfo := &pQosFlow.QosLevelParam.GbrQosInfo

				pGbrInfo.MaxFlowBitRateUL = uint64(gbrInfo.GetMaxFlowBitRateUL())
				pGbrInfo.MaxFlowBitRateDL = uint64(gbrInfo.GetMaxFlowBitRateDL())
				pGbrInfo.GuarantFlowBitRateUL = uint64(gbrInfo.GetGuaFlowBitRateUL())
				pGbrInfo.GuarantFlowBitRateDL = uint64(gbrInfo.GetGuaFlowBitRateDL())

				if gbrInfo.GetNotiControlPresent() {
					pGbrInfo.OptFlags.Set(T.GBRQI_NotifyCtrl)
					pGbrInfo.NotifyCtrl = T.NotificationControl(gbrInfo.GetNotiControl())
				}
				if gbrInfo.GetMaxPacketLossRateDLPresent() {
					pGbrInfo.OptFlags.Set(T.GBRQI_MaxPktLossRateDL)
					pGbrInfo.MaxPktLossRateDL = gbrInfo.GetMaxPacketLossRateDL()
				}
				if gbrInfo.GetMaxPacketLossRateULPresent() {
					pGbrInfo.OptFlags.Set(T.GBRQI_MaxPktLossRateUL)
					pGbrInfo.MaxPktLossRateUL = gbrInfo.GetMaxPacketLossRateUL()
				}
			}

			// reflect qos atrribute
			if qosParam.GetRefQosAttrPresent() {
				pQosFlow.QosLevelParam.OptFlags.Test(T.QFLQP_ReflecQosAttr)
				pQosFlow.QosLevelParam.RefQosAttr = T.ReflectQosAtt(qosParam.GetRefQosAttr())
			}

			// additional qos info
			if qosParam.GetAddQosFlowInfoPresent() {
				pQosFlow.QosLevelParam.OptFlags.Test(T.QFLQP_AddQosFlowInfo)
				pQosFlow.QosLevelParam.AddQosFlowInfo = T.AddQosFlowInfo(qosParam.GetAddQosFlowInfo())
			}

			p.AddQosFlowAddOrModReqList(pQosFlow)
		}

		// qos flow to release list
		qosFlowRelList := transfer.GetQosFlowToReleaseList()
		for i := 0; i < int(qosFlowRelList.Size()); i++ {
			qosFlowRelItem := qosFlowRelList.Get(i)
			qosFlowRelInfo := &T.QosFlowDesc{}

			qosFlowRelInfo.QowFlowInd = uint8(qosFlowRelItem.GetQosFlowInd())
			qosFlowRelInfo.Cause.Type = T.CauseType(qosFlowRelItem.GetCauseType())
			qosFlowRelInfo.Cause.Value = T.CauseValue(qosFlowRelItem.GetCauseValue())
			p.QosFlowToReleaseList = append(p.QosFlowToReleaseList, qosFlowRelInfo)
		}

	} else {
		return fmt.Errorf("failed to decode msg Bufer")
	}

	return nil
}
