package configure

import (
	"lite5gc/cmn/message/pfcp"
	"lite5gc/smf/smfcontext/gctxt"
)

type SystemConfig struct {
	RedisAddr IpAddr
	DBConf    DBConf
}

type ScConfig struct {
	NumOfAmfScInst int
	NumOfSmfScInst int
}

func (p *SystemConfig) ConvertSystemConfig(cmData *CmSysConfig) {
	p.RedisAddr = cmData.RedisSerAddr
	p.DBConf = cmData.BDConfig
}

func GetPFCPParams() gctxt.SmfToPFCPParameters {
	p := gctxt.SmfToPFCPParameters{}
	// PDR
	p.PDRs = make([]*gctxt.PacketDetectionRule, 2)
	p.PDRs[0] = &(gctxt.PacketDetectionRule{})
	p.PDRs[1] = &(gctxt.PacketDetectionRule{})
	p.IeFlags.Set(pfcp.IE_Create_PDR)

	// 上行
	p.PDRs[0].RuleID = 1
	p.PDRs[0].IeFlags.Set(pfcp.IE_Packet_Detection_Rule_ID)

	p.PDRs[0].PDI.SourceInterface = pfcp.Access //uplink
	p.PDRs[0].PDI.IeFlags.Set(pfcp.IE_Source_Interface)
	p.PDRs[0].IeFlags.Set(pfcp.IE_PDI)

	precedence := CmSmfConf.Rules[0].Precedence
	p.PDRs[0].Precedence = uint32(precedence)
	p.PDRs[0].IeFlags.Set(pfcp.IE_Precedence)

	sdf := CmSmfConf.Rules[0].PacketFilterLists[0].Descriptions[0]
	p.PDRs[0].PDI.PacketFilterSet.FlowDescription = []byte(sdf)
	p.PDRs[0].PDI.IeFlags.Set(pfcp.IE_SDF_Filter)

	qfi := CmSmfConf.Rules[0].QFI
	p.PDRs[0].PDI.QFI = qfi
	p.PDRs[0].PDI.IeFlags.Set(pfcp.IE_QFI)

	p.PDRs[0].OuterHeaderRemoval = 0
	p.PDRs[0].IeFlags.Set(pfcp.IE_Outer_Header_Removal)

	// QER
	p.PDRs[0].IeFlags.Set(pfcp.IE_Create_QER)
	p.PDRs[0].QerI = make([]gctxt.QerInfo, 1)
	p.PDRs[0].QerI[0].Qer.RuleID = 1
	p.PDRs[0].QerI[0].Qer.IeFlags.Set(pfcp.IE_QER_ID)

	p.PDRs[0].QerI[0].Qer.QFI = qfi
	p.PDRs[0].QerI[0].Qer.IeFlags.Set(pfcp.IE_QFI)

	// FAR
	p.PDRs[0].IeFlags.Set(pfcp.IE_Create_FAR)
	p.PDRs[0].FarI.Far.RuleID = 1
	p.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	p.PDRs[0].FarI.Far.Action = gctxt.FORW
	p.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Apply_Action)

	p.PDRs[0].FarI.Far.NetworkInstance = SmfConf.DnnInfo[0].Name
	p.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Network_Instance)

	p.PDRs[0].FarI.Far.DstInterface = gctxt.DestinationInterface_Core
	p.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Destination_Interface)

	// 下行
	p.PDRs[1].RuleID = 2
	p.PDRs[1].IeFlags.Set(pfcp.IE_Packet_Detection_Rule_ID)

	p.PDRs[1].PDI.SourceInterface = pfcp.Core //uplink
	p.PDRs[1].PDI.IeFlags.Set(pfcp.IE_Source_Interface)
	p.PDRs[1].IeFlags.Set(pfcp.IE_PDI)

	p.PDRs[1].Precedence = uint32(precedence)
	p.PDRs[1].IeFlags.Set(pfcp.IE_Precedence)

	p.PDRs[1].PDI.PacketFilterSet.FlowDescription = []byte(sdf)
	p.PDRs[1].PDI.IeFlags.Set(pfcp.IE_SDF_Filter)

	p.PDRs[1].PDI.QFI = qfi
	p.PDRs[1].PDI.IeFlags.Set(pfcp.IE_QFI)

	// FAR
	p.PDRs[1].IeFlags.Set(pfcp.IE_Create_FAR)
	p.PDRs[1].FarI.Far.RuleID = 2
	p.PDRs[1].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

	p.PDRs[1].FarI.Far.Action = gctxt.FORW
	p.PDRs[1].FarI.Far.IeFlags.Set(pfcp.IE_Apply_Action)

	p.PDRs[1].FarI.Far.NetworkInstance = SmfConf.DnnInfo[0].Name
	p.PDRs[1].FarI.Far.IeFlags.Set(pfcp.IE_Network_Instance)

	p.PDRs[1].FarI.Far.DstInterface = gctxt.DestinationInterface_Access
	p.PDRs[1].FarI.Far.IeFlags.Set(pfcp.IE_Destination_Interface)

	// bar
	p.PDRs[1].FarI.Far.BAR.BARID = 1
	p.PDRs[1].FarI.Far.IeFlags.Set(pfcp.IE_Create_BAR)

	p.PDRs[1].FarI.Far.BAR.DLDataNotificationDelay = CmSmfConf.BAR.DLDataNotificationDelay
	p.PDRs[1].FarI.Far.BAR.SugBuffPacketsCount = CmSmfConf.BAR.SugBuffPacketsCount

	return p
}
