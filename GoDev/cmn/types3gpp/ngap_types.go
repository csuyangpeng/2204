package types3gpp

import (
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

// Snssai struct definition
type Snssai struct {
	Sst    uint8
	Sd     uint32
	SdPrst bool
}

func (p *Snssai) String() string {
	return fmt.Sprintf("%d-%d", p.Sst, p.Sd)
}

func ConvertU32ToSd(sd []byte, val uint32, order types.ByteOrder) {
	if sd == nil || cap(sd) < SizeofSD {
		rlogger.Trace(types.ModCmn, rlogger.WARN, nil,
			"input sd is invalid %v", sd)
		return
	}

	var tmp [4]byte
	if order == types.LittleEndian {
		binary.LittleEndian.PutUint32(tmp[:], val)
		for i := 0; i < SizeofSD; i++ {
			sd[i] = tmp[i]
		}
	} else {
		binary.BigEndian.PutUint32(tmp[:], val)
		for i := 0; i < SizeofSD; i++ {
			sd[i] = tmp[i+1]
		}
	}

}

func ConvertSdToU32(sd []byte, order types.ByteOrder) uint32 {
	if sd == nil || cap(sd) < SizeofSD {
		rlogger.Trace(types.ModCmn, rlogger.WARN, nil,
			"input sd is invalid %v", sd)
		return 0
	}

	var tmp [4]byte
	for i := 0; i < SizeofSD; i++ {
		tmp[i+1] = sd[i]
	}

	if order == types.LittleEndian {
		return binary.LittleEndian.Uint32(tmp[:])
	} else {
		return binary.BigEndian.Uint32(tmp[:])
	}
}

// BPlmn struct definitaion
type BPlmn struct {
	Plmn             PlmnID
	SliceSupportList []Snssai
}

func (p *BPlmn) String() string {
	rtString := fmt.Sprintf("Plmn(%s), Snssai List{ ", &(p.Plmn))
	for i, v := range p.SliceSupportList {
		rtString += fmt.Sprintf("{#%d %s} ", i+1, &v)
	}

	return rtString
}

// AddSnssai add a Snssai into BPLMN
func (p *BPlmn) AddSnssai(snssai *Snssai) {
	p.SliceSupportList = append(p.SliceSupportList, *snssai)
}

// SupportedTA struct definition
type SupportedTA struct {
	Tac       TAC //[SizeofTAC]byte
	BPlmnList []BPlmn
}

func (p *SupportedTA) String() string {
	rtString := fmt.Sprintf("%v, BPlmn List{ ", p.Tac)
	for i, v := range p.BPlmnList {
		rtString += fmt.Sprintf("{ #%d %s }", i+1, &v)
	}

	return rtString
}

// AddBPlmn add a BPLMN into SupporttedTA
func (p *SupportedTA) AddBPlmn(bplmn *BPlmn) {
	p.BPlmnList = append(p.BPlmnList, *bplmn)
}

// GlobalGNBID struct definition
type GlobalGNBID struct {
	Plmn      PlmnID
	GNBLength uint16 //(22..32)
	GNBID     [SizeofGNBID]byte
}

func (p *GlobalGNBID) String() string {
	return fmt.Sprintf("plmn(%s), gNB Length(%d), gNB ID(%v)",
		&(p.Plmn), p.GNBLength, p.GNBID)
}

// GlobalRanNodeID struct definition
type GlobalRanNodeID struct {
	GNBID GlobalGNBID
}

func (p GlobalRanNodeID) String() string {
	return fmt.Sprintf("GlobalRanNodeId(GnbId(%s))", &(p.GNBID))
}

// NR Cell Id
type NRCellId [SizeofNrCellID]byte

// NR CGI
type NRCgi struct {
	PlmnID   PlmnID
	NrCellId NRCellId
}

func (p NRCgi) string() string {
	return fmt.Sprintf("NRCGI(%s,CellId(%s)",
		p.PlmnID, string(p.NrCellId[:]))
}

// timestamp
type TimeStampType [SizeofTimeStamp]byte

// User Location Information
type UserLocationInfo struct {
	NrCgi            NRCgi
	Tai              TAI
	IsTimeStampPrst  bool
	TimeStmp         TimeStampType
	IsPsCellInfoPrst bool
	PsCellInfo       NRCgi
}

func (p UserLocationInfo) String() string {
	return fmt.Sprintf("UserLocationInfo(%s,%s,TimeStamp(%v,%s) PsCellInformation(%v,%+v)",
		p.NrCgi, p.Tai, p.IsTimeStampPrst, string(p.TimeStmp[:]), p.IsPsCellInfoPrst, p.PsCellInfo)
}

//238.413 9.3.1.56
type TimeToWait uint8

const (
	Time1s  TimeToWait = 0
	Time2s  TimeToWait = 1
	Time5s  TimeToWait = 2
	Time10s TimeToWait = 3
	Time20s TimeToWait = 4
	Time60s TimeToWait = 5
)

type ARP struct {
	PriorityLevel uint16
	PreemptCap    PreemptionCapability
	PreemptVuln   PreemptionVulnerability
}

func (p ARP) String() string {
	return fmt.Sprintf("ARP{priorityLevel(%d), preemptCap(%s), preemptVuln(%s)}",
		p.PriorityLevel, p.PreemptCap, p.PreemptVuln)
}

type PreemptionCapability byte

const (
	NOT_PREEMPT PreemptionCapability = 0
	MAY_PREEMPT PreemptionCapability = 1
)

func (p *PreemptionCapability) StoreWithString(val string) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	switch val {
	case "NOT_PREEMPT":
		*p = NOT_PREEMPT
	case "MAY_PREEMPT":
		*p = MAY_PREEMPT
	default:
		return fmt.Errorf("failed to store preemption(%s)", val)
	}
	return nil
}
func (p PreemptionCapability) String() string {
	switch p {
	case NOT_PREEMPT:
		return "not_preempt"
	case MAY_PREEMPT:
		return "may_preempt"
	default:
		return "unknonw"
	}
}

type PreemptionVulnerability byte

const (
	NOT_PREEMPTABLE PreemptionVulnerability = 0
	PREEMPTABLE     PreemptionVulnerability = 1
)

func (p *PreemptionVulnerability) StoreWithString(val string) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	switch val {
	case "NOT_PREEMPTABLE":
		*p = NOT_PREEMPTABLE
	case "PREEMPTABLE":
		*p = PREEMPTABLE
	default:
		return fmt.Errorf("failed to store preemption(%s)", val)
	}
	return nil
}
func (p PreemptionVulnerability) String() string {
	switch p {
	case NOT_PREEMPTABLE:
		return "not_preemptable"
	case PREEMPTABLE:
		return "preemptable"
	default:
		return "unknonw"
	}
}

//SecurityIndication struct definition
type SecurityIndication struct {
	IntPrctInd             IntegrityProtectionInd
	ConfdPrctInd           ConfdProtectionInd
	MaxIntPrctDataRate     MaxIntProtectedDataRate
	IsMaxIntPrctRatePrst   bool
	MaxIntPrctDataRateDl   MaxIntProtectedDataRate
	IsMaxIntPrctRateDlPrst bool
}

func (p SecurityIndication) String() string {
	return fmt.Sprintf("SecuritIndication{IndPrctInd(%s),ConfdPrctInd(%s),"+
		"[%v MaxIntPrctDataRat(%s)]},[%v MaxIntPrctDataRateDl(%s)]",
		p.IntPrctInd, p.ConfdPrctInd, p.IsMaxIntPrctRatePrst, p.MaxIntPrctDataRate,
		p.IsMaxIntPrctRateDlPrst, p.MaxIntPrctDataRateDl)
}

//QOSFlowSetupReqest struct definition
type QosFlowSetupReqest struct {
	QosFlowInd    uint
	QosLevelParam QosFlowLevelQosParam
	IsERabIdPrst  bool
	ERabId        uint
}

func (p QosFlowSetupReqest) String() string {
	return fmt.Sprintf("QosFlowSetupRequest{QosFlowInd(%d),%s,[%v ERabId(%d)]}",
		p.QosFlowInd, p.QosLevelParam, p.IsERabIdPrst, p.ERabId)
}

//QosFlowAddOrModRequest struct definition
type QosFlowAddOrModRequest struct {
	QosFlowInd         uint
	IsQosFlowParamPrst bool
	QosLevelParam      QosFlowLevelQosParam
	IsERabIdPrst       bool
	ERabId             uint
}

func (p QosFlowAddOrModRequest) String() string {
	return fmt.Sprintf("QosFlowAddOrModRequest{QosFlowInd(%d),[%v,QosParam(%s)],[%v ERabId(%d)]}",
		p.QosFlowInd, p.IsQosFlowParamPrst, p.QosLevelParam, p.IsERabIdPrst, p.ERabId)
}

//QosFlowLevelQosParam struct definition
type QosFlowLevelQosParam struct {
	QosChats       QosCharacters
	Arp            ARP
	GbrQosInfo     GBRQosInfo
	RefQosAttr     ReflectQosAtt
	AddQosFlowInfo AddQosFlowInfo
	OptFlags       bitset.BitSet
}

func (p QosFlowLevelQosParam) String() string {
	return fmt.Sprintf("QosFlowLevelQosParam{%s,%s,[%v %s],[%v RefQosAttrr(%s)],"+
		"[%v AddQosFlowInfo(%s)]}",
		p.QosChats, p.Arp,
		p.OptFlags.Test(QFLQP_GbrInfo), p.GbrQosInfo,
		p.OptFlags.Test(QFLQP_ReflecQosAttr), p.RefQosAttr,
		p.OptFlags.Test(QFLQP_AddQosFlowInfo), p.AddQosFlowInfo)
}

const (
	QFLQP_GbrInfo uint = iota
	QFLQP_ReflecQosAttr
	QFLQP_AddQosFlowInfo
)

// GBRQosInfo
type GBRQosInfo struct {
	MaxFlowBitRateDL     uint64
	MaxFlowBitRateUL     uint64
	GuarantFlowBitRateDL uint64
	GuarantFlowBitRateUL uint64
	NotifyCtrl           NotificationControl
	MaxPktLossRateDL     uint
	MaxPktLossRateUL     uint
	OptFlags             bitset.BitSet
}

const (
	GBRQI_NotifyCtrl uint = iota
	GBRQI_MaxPktLossRateDL
	GBRQI_MaxPktLossRateUL
)

func (p GBRQosInfo) String() string {
	return fmt.Sprintf("GBRQosInfo{MaxFlowBitRateDL(%d),UL(%d),"+
		"GuarantFlowBitRateDL(%d),UL(%d),[%v NotifyCtrl(%s)],[%v MaxPktLossRateDL(%d)],"+
		"[%v MaxPktLossRateUL(%d)]}",
		p.MaxFlowBitRateDL, p.MaxFlowBitRateUL, p.GuarantFlowBitRateDL, p.GuarantFlowBitRateUL,
		p.OptFlags.Test(GBRQI_NotifyCtrl), p.NotifyCtrl,
		p.OptFlags.Test(GBRQI_MaxPktLossRateDL), p.MaxPktLossRateDL,
		p.OptFlags.Test(GBRQI_MaxPktLossRateUL), p.MaxPktLossRateUL)
}

// QosCharacters
type QosCharacters struct {
	IsDynamic     bool
	NonDynamic5qi NonDynamic5QI
	Dynamic5qi    Dynamic5QI
}

func (p QosCharacters) String() string {
	if p.IsDynamic {
		return fmt.Sprintf("QosCharacters(%s)", p.Dynamic5qi)
	} else {
		return fmt.Sprintf("QosCharacters(%s)", p.NonDynamic5qi)
	}
}

//NonDynamic5QI struct definition
type NonDynamic5QI struct {
	FiveQI           uint
	PriorityLevelQos uint
	AverageWindow    uint
	MaxDataBurstVol  uint
	OptFlags         bitset.BitSet
}

const (
	ND5QI_PriorityLevelQos uint = iota
	ND5QI_AverageWind
	ND5QI_MaxDataBrstVol
)

func (p NonDynamic5QI) String() string {
	return fmt.Sprintf("NonDynamic5QI{5QI(%d),"+
		"[%v PriorityLevelQos(%d)],[%v AverageWindow(%d)],[%v MaxDataBurstVol(%d)]}",
		p.FiveQI,
		p.OptFlags.Test(ND5QI_PriorityLevelQos), p.PriorityLevelQos,
		p.OptFlags.Test(ND5QI_AverageWind), p.AverageWindow,
		p.OptFlags.Test(ND5QI_MaxDataBrstVol), p.MaxDataBurstVol)
}

//Dynamic5QI struct definition
type Dynamic5QI struct {
	PriorityLevelQos uint
	PktDelayBuget    uint
	PktErrRate       PacketErrorRate
	FiveQI           uint
	DelayCritical    DelayCritical
	AverageWindow    uint
	MaxDataBurstVol  uint
	OptFlags         bitset.BitSet
}

const (
	D5QI_FiveQI uint = iota
	D5QI_DelayCri
	D5QI_AverageWind
	D5QI_MaxDataBrstVol
)

func (p Dynamic5QI) String() string {
	return fmt.Sprintf("Dynamic5QI{PriorityLevelQos(%d),PktDelayBuget(%d),PktErrRate(%s),"+
		"[%v 5QI(%d)],[%v DelayCritical(%s)],[%v AverageWindow(%d)],[%v MaxDataBurstVol(%d)]}",
		p.PriorityLevelQos, p.PktDelayBuget, p.PktErrRate,
		p.OptFlags.Test(D5QI_FiveQI), p.FiveQI,
		p.OptFlags.Test(D5QI_DelayCri), p.DelayCritical,
		p.OptFlags.Test(D5QI_AverageWind), p.AverageWindow,
		p.OptFlags.Test(D5QI_MaxDataBrstVol), p.MaxDataBurstVol)
}

type PacketErrorRate struct {
	PktErrRateScalar   uint
	PktErrRateExponent uint
}

func (p PacketErrorRate) String() string {
	return fmt.Sprintf("scalar(%d),exponent(%d)", p.PktErrRateScalar, p.PktErrRateExponent)
}

//PDUSessionResourceSetupListSUReq struct definition
type PduSessResSetupReqItem struct {
	PduSessionId            uint8
	IsNasPduPrst            bool
	NasPdu                  string
	Snssai                  *Snssai
	PduSessResSetupReqTrans string
}

func (p *PduSessResSetupReqItem) String() string {
	return fmt.Sprintf("PudSessResSetupReqItem{PduSessId(%d),"+
		"[%v NasPdu(%v)],SNssai(%s),Transfer(%v)}",
		p.PduSessionId, p.IsNasPduPrst, p.NasPdu, p.Snssai, []byte(p.PduSessResSetupReqTrans))
}

//PDUSessionResourceModifyRequestItem struct definition
type PduSessResModifyReqItem struct {
	PduSessionId          uint8
	IsNasPduPrst          bool
	NasPdu                string
	PduSessResModReqTrans string
	IsSnssaiPrst          bool
	Snssai                Snssai
}

func (p *PduSessResModifyReqItem) String() string {
	return fmt.Sprintf("PudSessResModifyReqItem{PduSessId(%d),"+
		"[%v NasPdu(%v)],Transfer(%v),[%v Snssai(%+v)]}",
		p.PduSessionId, p.IsNasPduPrst, p.NasPdu, []byte(p.PduSessResModReqTrans), p.IsSnssaiPrst, p.Snssai)
}

type PerformRst uint8

const (
	Performed    PerformRst = 0 // same with ngap spec 0
	NotPerformed PerformRst = 1 // same with ngap spec 1
)

func (p PerformRst) String() string {
	switch p {
	case Performed:
		return "performed"
	case NotPerformed:
		return "not_performed"
	default:
		return "unknown"
	}
}

//SecurityResult struct definition
type SecurityResult struct {
	IntPrctRst   PerformRst
	ConfdPrctRst PerformRst
}

func (p SecurityResult) String() string {
	return fmt.Sprintf("SecurityResult{IndPrctRst(%s),ConfdPrctRst(%s)",
		p.IntPrctRst, p.ConfdPrctRst)
}

type QosFlowDesc struct {
	QowFlowInd uint8
	Cause      NgapCause
}

func (p *QosFlowDesc) String() string {
	return fmt.Sprintf("qfi(%d),cause(%s)", p.QowFlowInd, p.Cause)
}

type PduSessResSetupRespItem struct {
	PduSessionId                uint8
	PduSessResSetupRespTransfer string
}

func (p *PduSessResSetupRespItem) String() string {
	return fmt.Sprintf("psi(%d), transfer(%v)",
		p.PduSessionId, []byte(p.PduSessResSetupRespTransfer))
}

type PduSessResRelRespItem struct {
	PduSessionId              uint8
	PduSessResRelRespTransfer string
}

func (p *PduSessResRelRespItem) String() string {
	return fmt.Sprintf("psi(%d), transfer(%v)",
		p.PduSessionId, []byte(p.PduSessResRelRespTransfer))
}

type PduSessResModifyRespItem struct {
	PduSessionId              uint8
	PduSessResRelMdfyTransfer string
}

func (p *PduSessResModifyRespItem) String() string {
	return fmt.Sprintf("psi(%d), transfer(%v)",
		p.PduSessionId, []byte(p.PduSessResRelMdfyTransfer))
}

type PduSessResFailModifyItem struct {
	PduSessionId                  uint8
	PduSessResRelFailMdfyTransfer string
}

func (p *PduSessResFailModifyItem) String() string {
	return fmt.Sprintf("psi(%d), transfer(%v)",
		p.PduSessionId, []byte(p.PduSessResRelFailMdfyTransfer))
}

type PduSessResRelCmmdItem struct {
	PduSessionId              uint8
	PduSessResRelCmmdTransfer string
}

func (p *PduSessResRelCmmdItem) String() string {
	return fmt.Sprintf("psi(%d), transfer(%v)",
		p.PduSessionId, []byte(p.PduSessResRelCmmdTransfer))
}

type PduSessResFailedToSetupItem struct {
	PduSessionId      uint8
	UnsuccessTransfer string
}

func (p *PduSessResFailedToSetupItem) String() string {
	return fmt.Sprintf("psi(%d), transfer(%v)",
		p.PduSessionId, []byte(p.UnsuccessTransfer))
}

type UeNgapIdType uint8

const (
	UeNgApIdPairType UeNgapIdType = 1
	AmfUeNgApIdType  UeNgapIdType = 2
)

func (p UeNgapIdType) String() string {
	switch p {
	case UeNgApIdPairType:
		return "UeNgapIdPair"
	case AmfUeNgApIdType:
		return "AmfUeNgapId"
	default:
		return "invalidType"
	}
}

type AmfPagingTarget byte

const (
	PT_RanNode AmfPagingTarget = 1
	PT_TAI     AmfPagingTarget = 2
)

func (p AmfPagingTarget) String() string {
	var rt string
	switch p {
	case PT_RanNode:
		rt = "RanNode"
	case PT_TAI:
		rt = "TAI"
	default:
		rt = "invalidTarget"
	}

	return "Target(" + rt + ")"
}

type RecommendRanNode4Page struct {
	PageTarget AmfPagingTarget
	RanNodeId  GlobalRanNodeID
	Tai        TAI
}

func (p RecommendRanNode4Page) String() string {
	return fmt.Sprintf("RecmdNode(%s,%s,%s)",
		p.PageTarget, p.RanNodeId, p.Tai)
}

const (
	MaxnoofRecommendedCells    byte = 16
	maxnoofRedommendedRANNodes byte = 16
)

type RecommendCell4Page struct {
	Cgi        NRCgi
	TimeStayed int
	IsTSPrst   bool
}

func (p RecommendCell4Page) String() string {
	return fmt.Sprintf("RecmdCell(%s,%d)", p.Cgi, p.TimeStayed)
}

type RecommendPagingInfo struct {
	CellInfo    []RecommendCell4Page
	RanNodeInfo []RecommendRanNode4Page
}

func (p RecommendPagingInfo) String() string {
	var rtStr string
	rtStr = "CellInfo("
	for _, v := range p.CellInfo {
		rtStr += fmt.Sprintf("%s", v)
	}
	rtStr += "),RanNodeInfo("
	for _, v := range p.RanNodeInfo {
		rtStr += fmt.Sprintf("%s", v)
	}
	rtStr += ")"

	return rtStr
}

//enum for Paging DRX
type PagingDRX uint8

const (
	PagingDrx_v32  PagingDRX = 0
	PagingDrx_v64  PagingDRX = 1
	PagingDrx_v128 PagingDRX = 2
	PagingDrx_v256 PagingDRX = 3
)

func (p PagingDRX) String() string {
	var str string
	switch p {
	case PagingDrx_v32:
		str = "v32"
	case PagingDrx_v64:
		str = "v64"
	case PagingDrx_v128:
		str = "v128"
	case PagingDrx_v256:
		str = "v256"
	default:
		str = "unknown"
	}
	return fmt.Sprintf("PagingDRX(%s)", str)
}

//enum for Paging DRX
type PagingOrigin uint8

const (
	Non_3Gpp PagingOrigin = 0
)

func (p PagingOrigin) String() string {
	var str string
	switch p {
	case Non_3Gpp:
		str = "non_3gpp"
	default:
		str = "unknown"
	}
	return fmt.Sprintf("PagingOrigin(%s)", str)
}

type PagingPriority uint8

const (
	PagingPri_lvl1 PagingPriority = 0
	PagingPri_lvl2 PagingPriority = 1
	PagingPri_lvl3 PagingPriority = 2
	PagingPri_lvl4 PagingPriority = 3
	PagingPri_lvl5 PagingPriority = 4
	PagingPri_lvl6 PagingPriority = 5
	PagingPri_lvl7 PagingPriority = 6
	PagingPri_lvl8 PagingPriority = 7
)

func (p PagingPriority) String() string {
	var str string
	switch p {
	case PagingPri_lvl1:
		str = "priolevel1"
	case PagingPri_lvl2:
		str = "priolevel2"
	case PagingPri_lvl3:
		str = "priolevel3"
	case PagingPri_lvl4:
		str = "priolevel4"
	case PagingPri_lvl5:
		str = "priolevel5"
	case PagingPri_lvl6:
		str = "priolevel6"
	case PagingPri_lvl7:
		str = "priolevel7"
	case PagingPri_lvl8:
		str = "priolevel8"
	default:
		str = "unknown"
	}
	return fmt.Sprintf("PagingPriority(%s)", str)
}

// RRC Estb Cause
type RrcEstbCause byte

const (
	Rrc_estb_cause_emergency          RrcEstbCause = 0
	Rrc_estb_cause_highPriorityAccess RrcEstbCause = 1
	Rrc_estb_cause_mt_Access          RrcEstbCause = 2
	Rrc_estb_cause_mo_Signalling      RrcEstbCause = 3
	Rrc_estb_cause_mo_Data            RrcEstbCause = 4
	Rrc_estb_cause_mo_VoiceCall       RrcEstbCause = 5
	Rrc_estb_cause_mo_VideoCall       RrcEstbCause = 6
	Rrc_estb_cause_mo_SMS             RrcEstbCause = 7
	Rrc_estb_cause_mps_PriorityAccess RrcEstbCause = 8
	Rrc_estb_cause_mcs_PriorityAccess RrcEstbCause = 9
	Rrc_estb_cause_notAvailable       RrcEstbCause = 10
)

func (p RrcEstbCause) String() string {
	var str string
	switch p {
	case Rrc_estb_cause_emergency:
		str = "emergency"
	case Rrc_estb_cause_highPriorityAccess:
		str = "highPriorityAccess"
	case Rrc_estb_cause_mt_Access:
		str = "mt_Access"
	case Rrc_estb_cause_mo_Signalling:
		str = "mo_Signalling"
	case Rrc_estb_cause_mo_Data:
		str = "mo_Data"
	case Rrc_estb_cause_mo_VoiceCall:
		str = "mo_VoiceCall"
	case Rrc_estb_cause_mo_VideoCall:
		str = "mo_VideoCall"
	case Rrc_estb_cause_mo_SMS:
		str = "mo_SMS"
	case Rrc_estb_cause_mps_PriorityAccess:
		str = "mps_PriorityAccess"
	case Rrc_estb_cause_mcs_PriorityAccess:
		str = "mcs_PriorityAccess"
	case Rrc_estb_cause_notAvailable:
		str = "notAvailable"
	default:
		str = "unknown"
	}
	return fmt.Sprintf("RRC Establishment cause(%s)", str)
}

type QosFlowMapInd byte

const (
	QosFlowMapInd_UL QosFlowMapInd = 0
	QosFlowMapInd_DL QosFlowMapInd = 1
)

func (p QosFlowMapInd) String() string {
	var str string
	switch p {
	case QosFlowMapInd_UL:
		str = "UL"
	case QosFlowMapInd_DL:
		str = "DL"
	default:
		str = "unknown"
	}
	return fmt.Sprintf("QosFlowMappingInd(%s)", str)
}

type AssocQosFlowItem struct {
	QosFlowIdentifier int
	IsMappingIndPrst  bool
	MappingIndication QosFlowMapInd
}

func (p AssocQosFlowItem) String() string {
	return fmt.Sprintf("AssocQosFlowItem(QosFlwoId(%d),[%v,%s]",
		p.QosFlowIdentifier, p.IsMappingIndPrst, p.MappingIndication)
}

type NgInterfaceReset byte

const (
	ResetAll NgInterfaceReset = 0
)

// Reset Type
type ResetType struct {
	IsNgInterfacePrst bool
	NgInterface       NgInterfaceReset
}

func (p ResetType) String() string {
	return fmt.Sprintf("ResetType(%v,%v)",
		p.IsNgInterfacePrst, p.NgInterface)
}

type UeAssLogicalNgConnItem struct {
	AmfUeNgapId       uint32
	IsAmfUeNgapIdPrst bool
	RanUeNgapId       uint32
	IsRanUeNgapIdPrst bool
}

func (p UeAssLogicalNgConnItem) String() string {
	return fmt.Sprintf("AmfUeNgapId(%v,%v),RanUeNgapId(%v,%v)",
		p.IsAmfUeNgapIdPrst, p.AmfUeNgapId, p.IsRanUeNgapIdPrst, p.RanUeNgapId)
}

type UlNguUpTnlModifyItem struct {
	UplinkUpTransLayerInfo   GtpTunnel
	DownlinkUpTransLayerInfo GtpTunnel
}

func (p UlNguUpTnlModifyItem) String() string {
	return fmt.Sprintf("UplinkUpTransLayerInfo(%s),DownlinkUpTransLayerInfo(%s)",
		p.UplinkUpTransLayerInfo, p.DownlinkUpTransLayerInfo)
}

type AddQosFlowPerTnlInfo struct {
	AddUpTransLayerInfo GtpTunnel
	AssocQosFlowList    []AssocQosFlowItem
}

// AddAssocQosFlowList add a AssocQosFlowItem into AddQosFlowPerTnlInfo
func (p *AddQosFlowPerTnlInfo) AddAssocQosFlowItem(assocQosFlowItem *AssocQosFlowItem) {
	p.AssocQosFlowList = append(p.AssocQosFlowList, *assocQosFlowItem)
}

func (p AddQosFlowPerTnlInfo) String() string {
	return fmt.Sprintf("AddUpTransLayerInfo(%s)",
		p.AddUpTransLayerInfo)
}

type VolumeTimeReport struct {
	StartTimeStamp     TAC
	EndTimeStamp       TAC
	UsageCountUplink   uint64
	UsageCountDownlink uint64
}

type RATType byte

const (
	Nr    RATType = 0
	Eutra RATType = 1
)

func (p RATType) String() string {
	var str string
	switch p {
	case Nr:
		str = "NR"
	case Eutra:
		str = "EUTRA"
	default:
		str = "unknonw"
	}
	return str
}

type PduSessionUsageReport struct {
	RatType              RATType
	VolumeTimeReportList []VolumeTimeReport
}

type QosFlowUsageReport struct {
	QosFlowId              uint
	RatType                RATType
	QosFlowsTimeReportList []VolumeTimeReport
}

//AddQosFlowUsageReport Add QosFlowUsageReport into QosFlowUsageReportList
func (p *QosFlowUsageReport) AddQosFlowTimeReport(timeReport *VolumeTimeReport) {
	p.QosFlowsTimeReportList = append(p.QosFlowsTimeReportList, *timeReport)
}

type SecRATUsageInformation struct {
	PduSessUsageReport           PduSessionUsageReport
	IsPduSessUsageReportPrst     bool
	QosFlowUsageReportList       []QosFlowUsageReport
	IsQosFlowUsageReportListPrst bool
}
