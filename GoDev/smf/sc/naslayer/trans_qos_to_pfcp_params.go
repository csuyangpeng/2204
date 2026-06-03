package naslayer

import (
	"fmt"
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/smfcontext/gctxt"
	"strconv"
)

func TransCreateQosRulesToPFCPParams(PduSessModReq nasmsg.PduSessionModifyRequestMsg) (pfcps gctxt.SmfToPFCPParameters, err error) {
	rlogger.FuncEntry(types.ModuleSmfNas, nil)

	pfcps.IeFlags.Set(pfcp.IE_Create_PDR)
	pfcps.PDRs = make([]*gctxt.PacketDetectionRule, 2)
	pfcps.PDRs[0] = &(gctxt.PacketDetectionRule{}) //for uplink
	pfcps.PDRs[1] = &(gctxt.PacketDetectionRule{}) //for downlink

	pfcps.PDRs[0].Tag = gctxt.CreatePDRs
	pfcps.PDRs[1].Tag = gctxt.CreatePDRs
	for i := 0; i < len(PduSessModReq.RequestQosRules.QoSRules); i++ {

		qosRule := PduSessModReq.RequestQosRules.QoSRules[i]

		switch qosRule.PacketFilterLists.PFList[0].PktFilterDirection {
		case nasie.UplinkOnly:
			ruleID, err := idmgr.GetInst().BorrowID(string(types.PDRID))
			if err != nil {
				rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, nil, "borrow PDR Id Failed.", err)
				return pfcps, fmt.Errorf("borrow PDR Id Failed")
			}
			pfcps.PDRs[0].RuleID = uint16(ruleID)
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Packet_Detection_Rule_ID)

			pfcps.PDRs[0].PDI.SourceInterface = 0
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_PDI)
			pfcps.PDRs[0].PDI.IeFlags.Set(pfcp.IE_Source_Interface)

			pfcps.PDRs[0].OuterHeaderRemoval = 0 //default value is 0 in uplink
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Outer_Header_Removal)

			pfcps.PDRs[0].Precedence = uint32(qosRule.QoSRulePrecedence)
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Precedence)

			sdfs, _ := TransFilterToSDF(qosRule.PacketFilterLists.PFList[0].PacketFilterContents)
			pfcps.PDRs[0].PDI.PacketFilterSet.FlowDescription = []byte(sdfs)
			pfcps.PDRs[0].PDI.IeFlags.Set(pfcp.IE_SDF_Filter)

			pfcps.PDRs[0].PDI.QFI = uint8(qosRule.QoSFlowIdentifier)
			pfcps.PDRs[0].PDI.IeFlags.Set(pfcp.IE_QFI)

			//pfcps.PDRs[0].FarI.Far.RuleID = uint32(configure.SmfConf.N4Conf.PFCPParameters.)
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Create_FAR)
			pfcps.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Create_QER)
			pfcps.PDRs[0].QerI = make([]gctxt.QerInfo, 1)
			pfcps.PDRs[0].QerI[0] = gctxt.QerInfo{}
			pfcps.PDRs[0].QerI[0].Qer.IeFlags.Set(pfcp.IE_GBR)
			pfcps.PDRs[0].QerI[0].Qer.IeFlags.Set(pfcp.IE_MBR)

			qosDesc := PduSessModReq.RequestQosFlowDesc.Descr[0].ParameterList.ParmsList
			for i := 0; i < len(qosDesc); i++ {
				switch qosDesc[i].ParameterID {
				case nasie.GFBRUplink:
					pfcps.PDRs[0].QerI[0].Qer.GBR.ULGBR = uint64(qosDesc[i].GFBRUplinkContent.Tokbps())
				case nasie.GFBRDownlink:
					pfcps.PDRs[0].QerI[0].Qer.GBR.DLGBR = uint64(qosDesc[i].GFBRDownlinkContent.Tokbps())
				case nasie.MFBRUplink:
					pfcps.PDRs[0].QerI[0].Qer.MBR.ULMBR = uint64(qosDesc[i].MFBRUplinkContent.Tokbps())
				case nasie.MFBRDownlink:
					pfcps.PDRs[0].QerI[0].Qer.MBR.DLMBR = uint64(qosDesc[i].MFBRDownlinkContent.Tokbps())
				}
			}
		case nasie.DownlinkOnly:
			ruleID, err := idmgr.GetInst().BorrowID(string(types.PDRID))
			if err != nil {
				rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, nil, "borrow PDR Id Failed.", err)
				return pfcps, fmt.Errorf("borrow PDR Id Failed")
			}
			pfcps.PDRs[0].RuleID = uint16(ruleID)
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Packet_Detection_Rule_ID)

			pfcps.PDRs[0].PDI.SourceInterface = 1
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_PDI)
			pfcps.PDRs[0].PDI.IeFlags.Set(pfcp.IE_Source_Interface)

			pfcps.PDRs[0].OuterHeaderRemoval = 1 //default value is 0 in uplink
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Outer_Header_Removal)

			pfcps.PDRs[0].Precedence = uint32(qosRule.QoSRulePrecedence)
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Precedence)

			sdfs, _ := TransFilterToSDF(qosRule.PacketFilterLists.PFList[0].PacketFilterContents)
			pfcps.PDRs[0].PDI.PacketFilterSet.FlowDescription = []byte(sdfs)
			pfcps.PDRs[0].PDI.IeFlags.Set(pfcp.IE_SDF_Filter)

			pfcps.PDRs[0].PDI.QFI = uint8(qosRule.QoSFlowIdentifier)
			pfcps.PDRs[0].PDI.IeFlags.Set(pfcp.IE_QFI)

			//pfcps.PDRs[0].FarI.Far.RuleID = uint32(configure.GetFar2Id())
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Create_FAR)
			pfcps.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

			//pfcps.PDRs[0].FarI.Far.BAR.BARID = uint8(configure.GetBarId())
			pfcps.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_Create_BAR)

			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Create_QER)
			pfcps.PDRs[0].QerI = make([]gctxt.QerInfo, 1)
			pfcps.PDRs[0].QerI[0] = gctxt.QerInfo{}
			pfcps.PDRs[0].QerI[0].Qer.IeFlags.Set(pfcp.IE_GBR)
			pfcps.PDRs[0].QerI[0].Qer.IeFlags.Set(pfcp.IE_MBR)
			qosDesc := PduSessModReq.RequestQosFlowDesc.Descr[0].ParameterList.ParmsList
			for i := 0; i < len(qosDesc); i++ {
				switch qosDesc[i].ParameterID {
				case nasie.GFBRUplink:
					pfcps.PDRs[0].QerI[0].Qer.GBR.ULGBR = uint64(qosDesc[i].GFBRUplinkContent.Tokbps())
				case nasie.GFBRDownlink:
					pfcps.PDRs[0].QerI[0].Qer.GBR.DLGBR = uint64(qosDesc[i].GFBRDownlinkContent.Tokbps())
				case nasie.MFBRUplink:
					pfcps.PDRs[0].QerI[0].Qer.MBR.ULMBR = uint64(qosDesc[i].MFBRUplinkContent.Tokbps())
				case nasie.MFBRDownlink:
					pfcps.PDRs[0].QerI[0].Qer.MBR.DLMBR = uint64(qosDesc[i].MFBRDownlinkContent.Tokbps())
				}
			}
		case nasie.Bidirectional:
			ruleID, err := idmgr.GetInst().BorrowID(string(types.PDRID))
			if err != nil {
				rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, nil, "borrow PDR Id Failed.", err)
				return pfcps, fmt.Errorf("borrow PDR Id Failed")
			}
			pfcps.PDRs[0].RuleID = uint16(ruleID)
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Packet_Detection_Rule_ID)
			ruleID2, err := idmgr.GetInst().BorrowID(string(types.PDRID))
			if err != nil {
				rlogger.Trace(types.ModuleSmfNas, rlogger.DEBUG, nil, "borrow PDR Id Failed.", err)
				return pfcps, fmt.Errorf("borrow PDR Id Failed")
			}
			pfcps.PDRs[1].RuleID = uint16(ruleID2)
			pfcps.PDRs[1].IeFlags.Set(pfcp.IE_Packet_Detection_Rule_ID)

			pfcps.PDRs[0].PDI.SourceInterface = 0
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_PDI)
			pfcps.PDRs[0].PDI.IeFlags.Set(pfcp.IE_Source_Interface)
			pfcps.PDRs[1].PDI.SourceInterface = 1
			pfcps.PDRs[1].IeFlags.Set(pfcp.IE_PDI)
			pfcps.PDRs[1].PDI.IeFlags.Set(pfcp.IE_Source_Interface)

			pfcps.PDRs[0].Precedence = uint32(qosRule.QoSRulePrecedence)
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Precedence)
			pfcps.PDRs[1].Precedence = uint32(qosRule.QoSRulePrecedence)
			pfcps.PDRs[1].IeFlags.Set(pfcp.IE_Precedence)

			sdfs, _ := TransFilterToSDF(qosRule.PacketFilterLists.PFList[0].PacketFilterContents)
			pfcps.PDRs[0].PDI.PacketFilterSet.FlowDescription = []byte(sdfs)
			pfcps.PDRs[0].PDI.IeFlags.Set(pfcp.IE_SDF_Filter)
			sdfs, _ = TransFilterToSDF(qosRule.PacketFilterLists.PFList[0].PacketFilterContents)
			pfcps.PDRs[1].PDI.PacketFilterSet.FlowDescription = []byte(sdfs)
			pfcps.PDRs[1].PDI.IeFlags.Set(pfcp.IE_SDF_Filter)

			pfcps.PDRs[0].PDI.QFI = uint8(qosRule.QoSFlowIdentifier)
			pfcps.PDRs[0].PDI.IeFlags.Set(pfcp.IE_QFI)
			pfcps.PDRs[1].PDI.QFI = uint8(qosRule.QoSFlowIdentifier)
			pfcps.PDRs[1].PDI.IeFlags.Set(pfcp.IE_QFI)

			//pfcps.PDRs[0].FarI.Far.RuleID = uint32(configure.GetFar1Id())
			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Create_FAR)
			pfcps.PDRs[0].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)
			//pfcps.PDRs[1].FarI.Far.RuleID = uint32(configure.GetFar2Id())
			pfcps.PDRs[1].IeFlags.Set(pfcp.IE_Create_FAR)
			pfcps.PDRs[1].FarI.Far.IeFlags.Set(pfcp.IE_FAR_ID)

			//pfcps.PDRs[1].FarI.Far.BAR.BARID = uint8(configure.GetBarId())
			pfcps.PDRs[1].FarI.Far.IeFlags.Set(pfcp.IE_Create_BAR)

			pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Create_QER)
			pfcps.PDRs[0].QerI = make([]gctxt.QerInfo, 1)
			pfcps.PDRs[0].QerI[0] = gctxt.QerInfo{}
			pfcps.PDRs[0].QerI[0].Qer.IeFlags.Set(pfcp.IE_GBR)
			pfcps.PDRs[0].QerI[0].Qer.IeFlags.Set(pfcp.IE_MBR)
			qosDesc := PduSessModReq.RequestQosFlowDesc.Descr[0].ParameterList.ParmsList
			for i := 0; i < len(qosDesc); i++ {
				switch qosDesc[i].ParameterID {
				case nasie.GFBRUplink:
					pfcps.PDRs[0].QerI[0].Qer.GBR.ULGBR = uint64(qosDesc[i].GFBRUplinkContent.Tokbps())
				case nasie.GFBRDownlink:
					pfcps.PDRs[0].QerI[0].Qer.GBR.DLGBR = uint64(qosDesc[i].GFBRDownlinkContent.Tokbps())
				case nasie.MFBRUplink:
					pfcps.PDRs[0].QerI[0].Qer.MBR.ULMBR = uint64(qosDesc[i].MFBRUplinkContent.Tokbps())
				case nasie.MFBRDownlink:
					pfcps.PDRs[0].QerI[0].Qer.MBR.DLMBR = uint64(qosDesc[i].MFBRDownlinkContent.Tokbps())
				}
			}
			pfcps.PDRs[1].IeFlags.Set(pfcp.IE_Create_QER)
			pfcps.PDRs[1].QerI = make([]gctxt.QerInfo, 1)
			pfcps.PDRs[1].QerI[0] = gctxt.QerInfo{}
			pfcps.PDRs[1].QerI[0].Qer.IeFlags.Set(pfcp.IE_GBR)
			pfcps.PDRs[1].QerI[0].Qer.IeFlags.Set(pfcp.IE_MBR)
			qosDesc = PduSessModReq.RequestQosFlowDesc.Descr[0].ParameterList.ParmsList
			for i := 0; i < len(qosDesc); i++ {
				switch qosDesc[i].ParameterID {
				case nasie.GFBRUplink:
					pfcps.PDRs[1].QerI[0].Qer.GBR.ULGBR = uint64(qosDesc[i].GFBRUplinkContent.Tokbps())
				case nasie.GFBRDownlink:
					pfcps.PDRs[1].QerI[0].Qer.GBR.DLGBR = uint64(qosDesc[i].GFBRDownlinkContent.Tokbps())
				case nasie.MFBRUplink:
					pfcps.PDRs[1].QerI[0].Qer.MBR.ULMBR = uint64(qosDesc[i].MFBRUplinkContent.Tokbps())
				case nasie.MFBRDownlink:
					pfcps.PDRs[1].QerI[0].Qer.MBR.DLMBR = uint64(qosDesc[i].MFBRDownlinkContent.Tokbps())
				}
			}
		default:
			rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "unSupport packet filter direction")
			return pfcps, fmt.Errorf("unSupport packet filter direction")
		}
	}

	return pfcps, nil
}

func TransFilterToSDF(component nasie.PacketFilterContent) (s string, e error) {
	rlogger.FuncEntry(types.ModuleSmfNas, nil)
	rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "input PacketFilterContent", component)
	s = "permit out"
	var protocolS, toS, localPortS, RemotePortS string
	for i := 0; i < len(component.PacketFilterContentID); i++ {
		switch component.PacketFilterContentID[i] {
		case nasie.IPv4RemoteAddressType:
			toS = " to " + component.IPV4RemoteAddr.IP.String() + "/32"
			rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "toS ", toS)
		case nasie.IPv4LocalAddressType:
		case nasie.IPv6RemoteAddressPrefixLengthType:
		case nasie.IPv6LocalAddressPrefixLengthType:
		case nasie.ProtocolIdentifierNextHeaderType:
			protocolS = " " + strconv.Itoa(int(component.V4ProtocolIdOrV6NextHeader))
			rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "protocolS ", protocolS)
		case nasie.SingleLocalPortType, nasie.SingleRemotePortType:
		case nasie.LocalPortRangeType:
			localPortS = " port " + strconv.FormatInt(int64(component.LocalPortRangeLowLimit), 10) +
				"-" + strconv.FormatInt(int64(component.LocalPortRangeHighLimit), 10)
			rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "localPortS ", localPortS)
		case nasie.RemotePortRangeType:
			RemotePortS = " port " + strconv.FormatInt(int64(component.RemotePortRangeLowLimit), 10) +
				"-" + strconv.FormatInt(int64(component.RemotePortRangeHighLimit), 10)
			rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "RemotePortS ", RemotePortS)
		case nasie.SecurityParameterIndexType:
			rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "pass")
		case nasie.TypeOfServiceTrafficClassType:
			rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "pass")
		case nasie.DestinationMACAddressType:
		case nasie.SourceMACAddressType:
		case nasie.FlowLabelType:
			rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "pass")
		case nasie.Q8021CTAGVIDType:
		case nasie.Q8021STAGVIDType:
		case nasie.Q8021CTAGPCPDEIType:
		case nasie.Q8021STAGPCPDEIType:
		case nasie.EthertypeType:
		default:
		}
	}
	if protocolS != "" {
		s += protocolS
	} else {
		s += " ip"
	}
	if localPortS != "" {
		s += " from 0.0.0.0/0 " + localPortS
	} else {
		s += " from any"
	}
	if RemotePortS != "" {
		if toS != "" {
			s += toS + RemotePortS
		} else {
			s += " 0.0.0.0/0 " + RemotePortS
		}
	} else {
		s += " to assigned"
	}
	rlogger.Trace(types.ModuleSmfNas, rlogger.INFO, nil, "SDF string = ", s)
	return s, nil
}

func TransDeleteQosRulesToPFCPParams(PduSessModReq nasmsg.PduSessionModifyRequestMsg) (pfcps gctxt.SmfToPFCPParameters) {
	rlogger.FuncEntry(types.ModuleSmfNas, nil)
	pfcps.IeFlags.Set(pfcp.IE_Create_PDR)
	pfcps.PDRs = make([]*gctxt.PacketDetectionRule, 1)
	pfcps.PDRs[0] = &(gctxt.PacketDetectionRule{})
	pfcps.PDRs[0].Tag = gctxt.RemovePDRs
	pfcps.PDRs[0].IeFlags.Set(pfcp.IE_Create_QER)
	pfcps.PDRs[0].QerI = make([]gctxt.QerInfo, 1)
	pfcps.PDRs[0].QerI[0].Qer.QFI = uint8(PduSessModReq.RequestQosFlowDesc.Descr[0].QFI)
	pfcps.PDRs[0].QerI[0].Qer.IeFlags.Set(pfcp.IE_QFI)
	return pfcps
}
