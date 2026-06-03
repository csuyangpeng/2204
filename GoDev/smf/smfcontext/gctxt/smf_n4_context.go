package gctxt

import (
	"encoding/json"
	"fmt"
	"github.com/willf/bitset"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"net"
	"strconv"
)

// 3GPP TS 23.501 V15.3.0 (2018-09)
// 5.8.2.11.2	N4 Session Context
type N4SessionContext struct {
	PduCxt       *PduSessContext
	SEID         uint64
	UpfSEID      pfcp.IEFSEID
	PduSessionId nas.PduSessID
	IMSI         types3gpp.Imsi
	PDRs         []*pfcp.IECreatePDR
	//PARs                     []*pfcp.IECreateFAR
	URRs                     []*pfcp.IECreateURR
	QERs                     []*pfcp.IECreateQER
	FARs                     []*pfcp.IECreateFAR
	BAR                      *pfcp.IECreateBAR
	CreateTrafficEndpoints   []*pfcp.IECreateTrafficEndpoint
	PDNType                  *pfcp.IEPDNType
	UserPlaneInactivityTimer *pfcp.IEUserPlaneInactivityTimer
	UserID                   *pfcp.IEUserID
	TraceInformation         *pfcp.IETraceInformation
	//	CN tunnel info. // key is PDRID
	LocalFTEID map[uint16]*pfcp.IEFTEID
	//-	Network instance.
	NetworkInstance map[uint16]*pfcp.IENetworkInstance
	//-	QFIs.
	PDRQFIs map[uint16][]*pfcp.IEQFI
	//-	IP Packet Filter Set
	SDFFilters map[uint16][]*pfcp.IESDFFilter
	//Application Identifier
	ApplicationID map[uint16]*pfcp.IEApplicationID
	//Ethernet Packet Filter Set
	EthPacketFilters map[uint16][]*pfcp.IEEthernetPacketFilter
}

func (this *N4SessionContext) String() string {
	s := fmt.Sprintf("SEID:%d,", this.SEID)
	s += fmt.Sprintf("PduSessionId:%v,", this.PduSessionId)
	s += fmt.Sprintf("IMSI:%s,", this.IMSI)
	s += fmt.Sprintf("BAR:%s,", this.BAR)
	for i, v := range this.PDRs {
		tmp, _ := json.Marshal(v)
		s += fmt.Sprintf("\nPDR %d:%s,", i, tmp)
	}
	for i, v := range this.FARs {
		tmp, _ := json.Marshal(v)
		s += fmt.Sprintf("\nFAR %d:%s,", i, tmp)
	}
	for i, v := range this.URRs {
		tmp, _ := json.Marshal(v)
		s += fmt.Sprintf("\nURR %d:%s,", i, tmp)
	}
	for i, v := range this.QERs {
		tmp, _ := json.Marshal(v)
		s += fmt.Sprintf("\nQER %d:%s,", i, tmp)
	}
	return s
}

type N4Cause int8

const (
	Cause_Reserved N4Cause = 0

	// Acceptance in a response
	Cause_Request_accepted N4Cause = 1 // (success)
	//2-63	Spare.                           =    2

	// Rejection in a response
	Cause_Request_rejected                   N4Cause = 64 //(reason not specified)
	Cause_Session_context_not_found          N4Cause = 65
	Cause_Mandatory_IE_missing               N4Cause = 66
	Cause_Conditional_IE_missing             N4Cause = 67
	Cause_Invalid_length                     N4Cause = 68
	Cause_Mandatory_IE_incorrect             N4Cause = 69
	Cause_Invalid_Forwarding_Policy          N4Cause = 70
	Cause_Invalid_FTEID_allocation_option    N4Cause = 71
	Cause_No_established_PFCP_Association    N4Cause = 72
	Cause_Rule_creation_modification_Failure N4Cause = 73
	Cause_PFCP_entity_in_congestion          N4Cause = 74
	Cause_No_resources_available             N4Cause = 75
	Cause_Service_not_supported              N4Cause = 76
	Cause_System_failure                     N4Cause = 77
	//78 to 255	Spare for future use in a response message.
)

func (c *N4Cause) Error() string {
	return "PFCP Error:" + strconv.Itoa(int(*c))
}

type ParametersMap interface {
	Fill() (interface{}, error)
}

//3GPP TS 23.501 V15.3.0 (2018-09)
//5.8.2.11	Parameters for N4 session management
// smf N4参数指示
type SmfToPFCPParameters struct {
	//SEID                     uint64
	//SMFIP                    net.IP
	//TEID                     types3gpp.Teid

	PDRs                     []*PacketDetectionRule     `IeFlags:"IE_Create_PDR"`
	UserPlaneInactivityTimer IEUserPlaneInactivityTimer `IeFlags:"IE_User_Plane_Inactivity_Timer"`
	UserID                   IEUserID                   `IeFlags:"IE_User_ID"`
	// presence
	IeFlags bitset.BitSet
}

// Presence 表示本参数是否携带，true是携带.
type IEUserID struct {
	IMSI string
	//Presence bool
}
type IEUserPlaneInactivityTimer struct {
	TimerValue uint32
	//Presence   bool
}

type FTEID struct {
	TEID     types3gpp.Teid
	IPv4Addr net.IP
}

type SDFFilter struct {
	FlowDescription []byte
	IPFilterRule
}

type IPFilterRule struct {
	Protocol        string `json:"proto"`              // tcp  udp  ip
	SourceIP        string `json:"src_ip"`             // 对应DN IP//"192.0.2.10/24",
	DestinationIP   string `json:"dst_ip"`             // 对应终端IP//"192.0.2.10/24",
	SourcePort      string `json:"src_port,omitempty"` //"5000" or "6000-8000"
	DestinationPort string `json:"dst_port,omitempty"` //"5000" or "6000-8000"
}

type IEQFI struct {
	Value uint8
}

type OuterHeaderCreation struct {
	//gNB 侧teid
	Description uint16
	TEID        types3gpp.Teid
	IPv4Addr    net.IP
}

type BAR struct {
	BARID                   uint8
	DLDataNotificationDelay uint8
	SugBuffPacketsCount     uint8

	// presence
	IeFlags bitset.BitSet
}

func (this *BAR) Fill() (*pfcp.IECreateBAR, error) {
	bar := &pfcp.IECreateBAR{}
	//	Mandatory ie
	//BAR ID	M
	bar.BARID.Value = this.BARID
	//Downlink Data Notification Delay C
	bar.IeFlags.Set(pfcp.IE_Downlink_Data_Notification_Delay)
	bar.DLDataNotificationDelay.Value = this.BARID
	//Suggested Buffering Packets Count  C
	bar.IeFlags.Set(pfcp.IE_Suggested_Buffering_Packets_Count)
	bar.SugBuffPacketsCount.Set(this.SugBuffPacketsCount)
	return bar, nil
}

type IEHeaderEnrichment struct {
	HeaderType    uint8
	FieldNameLen  uint8
	FieldName     string
	FieldValueLen uint8
	FieldValue    string
}

type IEProxying struct {
	INS bool
	ARP bool
}

// 23.501 中对N4参考的要求
//5.8.2.11 Parameters for N4 session management

/*-	Packet Detection Rules (PDR) that contain information to classify traffic (PDU(s)) arriving at the UPF;
-	Forwarding Action Rules (FAR) that contain information on whether forwarding, dropping or buffering is
to be applied to a traffic identified by PDR(s);
-	Usage Reporting Rules (URR) contains information that defines how traffic identified by PDR(s) shall
be accounted as well as how a certain measurement shall be reported;
-	QoS Enforcement Rules (QER), that contain information related to QoS enforcement of traffic identified by PDR(s);
-	Trace Requirements.*/

// PDR
//5.8.2.11.3	Packet Detection Rule
type PacketDetectionRule struct {
	Tag PDROpr

	//N4 Session ID  M
	SEID uint64
	//Rule ID  M
	RuleID uint16
	//Precedence M
	Precedence uint32
	// M
	PDI PacketDetectionInfo
	//Outer header removal
	OuterHeaderRemoval uint8
	//Forwarding Action Rule ID
	//pfcp.IECreatePDR
	FarI FarInfo
	//List of Usage Reporting Rule(s)
	UrrI []UrrInfo

	//List of QoS Enforcement Rule(s)
	QerI []QerInfo

	// presence
	IeFlags bitset.BitSet
}

type PDROpr uint8

const (
	CreatePDRs PDROpr = 1
	RemovePDRs PDROpr = 2
	UpdatePDRs PDROpr = 3
)

func (this *PacketDetectionRule) Fill() (*pfcp.IECreatePDR, error) {
	pdr := &pfcp.IECreatePDR{}

	for id, e := this.IeFlags.NextSet(0); e; id, e = this.IeFlags.NextSet(id + 1) {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil, "fill PDR:%d", id)
		switch id {
		// Mandatory ie
		//PDR ID  	M
		case pfcp.IE_Packet_Detection_Rule_ID:
			pdr.PDRID.RuleID = this.RuleID
			//Precedence	M
		case pfcp.IE_Precedence:
			pdr.Precedence.PrecedenceValue = this.Precedence
			//PDI	        M
		case pfcp.IE_PDI:
			pdi, _ := this.PDI.Fill()
			pdr.PDI = *pdi
			// optional ie
			//Outer Header Removal 	C
		case pfcp.IE_Outer_Header_Removal:
			pdr.IeFlags.Set(id)
			pdr.OuterHeaderRemoval = &pfcp.IEOuterHeaderRemoval{}
			pdr.OuterHeaderRemoval.Description = this.OuterHeaderRemoval
			//FAR ID 	C
		case pfcp.IE_Create_FAR:
			pdr.IeFlags.Set(pfcp.IE_FAR_ID)
			pdr.FARID = &pfcp.IEFARID{}
			pdr.FARID.Value = this.FarI.Far.RuleID
			//URR ID	C
		case pfcp.IE_Create_URR:
			pdr.IeFlags.Set(pfcp.IE_URR_ID)
			for _, urr := range this.UrrI {
				urrid := &pfcp.IEURRID{}
				urrid.URRIDValue = urr.Urr.URRID
				pdr.URRIDs = append(pdr.URRIDs, urrid)
			}

			//QER ID 	C
		case pfcp.IE_Create_QER:
			pdr.IeFlags.Set(pfcp.IE_QER_ID)
			for _, QerInfo := range this.QerI {
				qerid := &pfcp.IEQERID{}
				qerid.Value = QerInfo.Qer.RuleID
				pdr.QERIDs = append(pdr.QERIDs, qerid)
			}
			//Activate Predefined Rules 	C
		case pfcp.IE_Activate_Predefined_Rules:
		case pfcp.IE_F_SEID:
		default:
			return nil, fmt.Errorf("Illegal IE(%d)", id)

		}
	}
	return pdr, nil
}

func (this *PacketDetectionRule) String() string {
	s := fmt.Sprintf("SEID:%d,", this.SEID)
	s += fmt.Sprintf("RuleID:%d,", this.RuleID)
	s += fmt.Sprintf("Precedence:%d,", this.Precedence)
	s += fmt.Sprintf("PDI:%s,", &this.PDI)
	s += fmt.Sprintf("OuterHeaderRemoval:%d,", this.OuterHeaderRemoval)
	s += fmt.Sprintf("FarI:%s,", &this.FarI)
	for i, v := range this.UrrI {
		s += fmt.Sprintf("UrrI %d:%s,", i, &v)
	}
	for i, v := range this.QerI {
		s += fmt.Sprintf("QerI %d:%s,", i, &v)
	}
	s += fmt.Sprintf("IeFlags:%s,", &this.IeFlags)
	return s
}

//Packet detection information
type PacketDetectionInfo struct {
	//Source interface
	SourceInterface uint8
	//UE IP address
	UEIPAddress net.IP
	//Network instance
	NetworkInstance string //types3gpp.Apn
	//CN tunnel info
	CNTunnel FTEID
	//Packet Filter Set
	PacketFilterSet SDFFilter
	//Application ID
	AppId string
	//QoS Flow ID
	QFI uint8
	//Ethernet PDU Session Information
	ETHI uint8 //pfcp.IEEthernetPDUSessionInfo

	// presence
	IeFlags bitset.BitSet
}

func (this *PacketDetectionInfo) Fill() (*pfcp.IEPDI, error) {
	pdi := &pfcp.IEPDI{}

	for id, e := this.IeFlags.NextSet(0); e; id, e = this.IeFlags.NextSet(id + 1) {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil, "fill PDI:%d", id)
		switch id {
		// Mandatory ie
		//Source Interface	   M
		case pfcp.IE_Source_Interface:
			pdi.SourceInterface.Type = pfcp.IE_Source_Interface
			pdi.SourceInterface.InterfaceValue = this.SourceInterface
			// optional ie
			//Local F-TEID 	       O  F-TEID
		case pfcp.IE_F_TEID:
			pdi.IeFlags.Set(id)
			pdi.LocalFTEID = &pfcp.IEFTEID{}
			pdi.LocalFTEID.V4Flag = pfcp.Flag
			pdi.LocalFTEID.TEID = this.CNTunnel.TEID         //todo 需传入
			pdi.LocalFTEID.IPv4Addr = this.CNTunnel.IPv4Addr //// upf n3 ip
			//pdr.PDI.LocalFTEID.V4Flag = pfcp.Flag
			//pdr.PDI.LocalFTEID.TEID = loccalTeid
			//pdr.PDI.LocalFTEID.IPv4Addr = configure.SmfConf.N3Conf.UPFIP // upf n3 ip
			//Network Instance	   O
		case pfcp.IE_Network_Instance:
			pdi.IeFlags.Set(id)
			pdi.NetworkInstance = &pfcp.IENetworkInstance{}
			pdi.NetworkInstance.NetworkInstance = this.NetworkInstance //pCtxt.DNN.String()
			//UE IP address 	       O
		case pfcp.IE_UE_IP_Address:
			pdi.IeFlags.Set(id)
			pdi.UEIPaddress = &pfcp.IEUEIPaddress{}
			pdi.UEIPaddress.SD = pfcp.Flag
			pdi.UEIPaddress.V4Flag = pfcp.Flag
			pdi.UEIPaddress.IPv4Addr = this.UEIPAddress //pCtxt.UEIP
			//Traffic Endpoint ID	   C
		case pfcp.IE_Traffic_Endpoint_ID:
			pdi.IeFlags.Set(id)
			pdi.TrafficEndpointID = &pfcp.IETrafficEndpointID{}
			pdi.TrafficEndpointID.Value = 1 // todo，没有使用
			//SDF Filter	           O
		case pfcp.IE_SDF_Filter:
			pdi.IeFlags.Set(id)
			pdi.SDFFilters = make([]*pfcp.IESDFFilter, 1)
			sDFFilter := &pfcp.IESDFFilter{}
			pdi.SDFFilters[0] = sDFFilter

			sDFFilter.Type = pfcp.IE_SDF_Filter
			sDFFilter.FDFlag = true
			sDFFilter.FlowDescription = this.PacketFilterSet.FlowDescription
			sDFFilter.LengthofFlowDescription = uint16(len(sDFFilter.FlowDescription))

			//Application ID	       O
		case pfcp.IE_Application_ID:
			pdi.IeFlags.Set(id)
			pdi.ApplicationID = &pfcp.IEApplicationID{}
			pdi.ApplicationID.ApplicationIdentifier = []byte(this.AppId)
			//Ethernet PDU Session Information	O
		case pfcp.IE_Ethernet_PDU_Session_Information:
			pdi.IeFlags.Set(id)
			pdi.EthPDUSessionInfo = &pfcp.IEEthernetPDUSessionInfo{}
			pdi.EthPDUSessionInfo.ETHI = this.ETHI
			//Ethernet Packet Filter	O
		case pfcp.IE_Ethernet_Packet_Filter:
			//QFIs	O
		case pfcp.IE_QFI:
			pdi.IeFlags.Set(id)
			pdi.QFIs = make([]*pfcp.IEQFI, 1)
			pdi.QFIs[0] = &pfcp.IEQFI{}
			pdi.QFIs[0].Type = pfcp.IE_QFI
			pdi.QFIs[0].Value = this.QFI
			//Framed-Route	O
		case pfcp.IE_Framed_Route:
			//Framed-Routing	O
		case pfcp.IE_Framed_Routing:
			//Framed-IPv6-Route	  O
		case pfcp.IE_Framed_IPv6_Route:
		default:
			return nil, fmt.Errorf("Illegal IE(%d)", id)

		}
	}
	return pdi, nil
}
func (this *PacketDetectionInfo) String() string {
	s := fmt.Sprintf("SourceInterface:%d,", this.SourceInterface)
	s += fmt.Sprintf("UEIPAddress:%s,", this.UEIPAddress)
	s += fmt.Sprintf("NetworkInstance:%s,", this.NetworkInstance)
	s += fmt.Sprintf("CNTunnelInfo:(TEID:%d,IP:%s),", this.CNTunnel.TEID, this.CNTunnel.IPv4Addr)
	s += fmt.Sprintf("PacketFilterSet:%s,", &this.PacketFilterSet)
	s += fmt.Sprintf("AppId:%s,", this.AppId)
	s += fmt.Sprintf("QFI:%d,", this.QFI)
	s += fmt.Sprintf("ETHI:%d,", this.ETHI)
	s += fmt.Sprintf("IeFlags:%s,", &this.IeFlags)
	return s
}

//Forwarding Action Rule
type FarInfo struct {
	//FarID uint32
	Far ForwardActionRule
}

func (this *FarInfo) String() string {
	return fmt.Sprintf("FarID %d:%s", this.Far)
}

//Table 5.8.2.11.5-1: Attributes within Usage Reporting Rule
type IEMeasurementMethod struct {
	DURAT bool
	VOLUM bool
	EVENT bool
}
type IEReportingTriggers struct {
	//Flag5 uint8
	//Flag6 uint8
	PERIO bool
	VOLTH bool
	TIMTH bool
	QUHTI bool
	START bool
	STOPT bool
	DROTH bool
	LIUSA bool

	VOLQU bool
	TIMQU bool
	ENVCL bool
	MACAR bool
	EVETH bool
	EVEQU bool
}

func (this *IEReportingTriggers) Fill() (*pfcp.IEReportingTriggers, error) {
	rep := &pfcp.IEReportingTriggers{}
	//  PERIO bool
	rep.PERIO = this.PERIO
	//	VOLTH bool
	rep.VOLTH = this.VOLTH
	//	TIMTH bool
	rep.TIMTH = this.TIMTH
	//	QUHTI bool
	rep.QUHTI = this.QUHTI
	//	START bool
	rep.START = this.START
	//	STOPT bool
	rep.STOPT = this.STOPT
	//	DROTH bool
	rep.DROTH = this.DROTH
	//	LIUSA bool
	rep.LIUSA = this.LIUSA
	//
	//	VOLQU bool
	rep.VOLQU = this.VOLQU
	//	TIMQU bool
	rep.TIMQU = this.TIMQU
	//	ENVCL bool
	rep.ENVCL = this.ENVCL
	//	MACAR bool
	rep.MACAR = this.MACAR
	//	EVETH bool
	rep.EVETH = this.EVETH
	//	EVEQU bool
	rep.EVEQU = this.EVEQU

	return rep, nil
}

type IEVolumeThreshold struct {
	TOVOL          bool
	ULVOL          bool
	DLVOL          bool
	TotalVolume    uint64
	UplinkVolume   uint64
	DownlinkVolume uint64
}
type IEVolumeQuota struct {
	TOVOL          bool
	ULVOL          bool
	DLVOL          bool
	TotalVolume    uint64
	UplinkVolume   uint64
	DownlinkVolume uint64
}
type IEDroppedDLTrafficThreshold struct {
	DLPA                        bool
	DLBY                        bool
	DownlinkPackets             uint64
	NumberofBytesofDownlinkData uint64
}
type IEMeasurementInformation struct {
	MBQE bool
	INAM bool
	RADI bool
	ISTM bool
}
type URR struct {
	URRID uint32
	//Measurement Method
	MeasurementMethod IEMeasurementMethod
	//Reporting triggers
	ReportingTriggers IEReportingTriggers
	//Periodic measurement threshold
	MeasurementPeriod uint32
	//Volume measurement threshold
	VolumeThreshold IEVolumeThreshold
	VolumeQuota     IEVolumeQuota

	//Time measurement threshold
	TimeThreshold             uint32
	TimeQuota                 uint32
	QuotaHoldingTime          uint32
	DroppedDLTrafficThreshold IEDroppedDLTrafficThreshold
	MonitoringTime            uint32
	//Event measurement threshold
	EventThreshold uint32
	EventQuota     uint32
	//Event based reporting

	//Inactivity detection time
	InactivityDetectionTime uint32
	//Linked URR ID(s)
	LinkedURRID uint32
	//Measurement information
	MeasurementInformation IEMeasurementInformation

	EthernetInactivityTimer uint32

	// presence
	IeFlags bitset.BitSet
}

func (this *URR) Fill() (*pfcp.IECreateURR, error) {
	urr := &pfcp.IECreateURR{}

	for id, e := this.IeFlags.NextSet(0); e; id, e = this.IeFlags.NextSet(id + 1) {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil, "fill URR:%d", id)
		switch id {
		// Mandatory ie
		//URR ID	M
		case pfcp.IE_URR_ID:
			urr.URRID.Set(this.URRID)
			//Measurement Method	M
		case pfcp.IE_Measurement_Method:
			urr.MeasurementMethod.DURAT = this.MeasurementMethod.DURAT
			urr.MeasurementMethod.VOLUM = this.MeasurementMethod.VOLUM
			urr.MeasurementMethod.EVENT = this.MeasurementMethod.EVENT
			//Reporting Triggers	M
		case pfcp.IE_Reporting_Triggers:
			rep, _ := this.ReportingTriggers.Fill()
			urr.ReportingTriggers = *rep
			// optional ie
			//Measurement Period 	C
		case pfcp.IE_Measurement_Period:
			urr.IeFlags.Set(id)
			urr.MeasurementPeriod = &pfcp.IEMeasurementPeriod{}
			urr.MeasurementPeriod.Period = this.MeasurementPeriod
			//Volume Threshold	C
		case pfcp.IE_Volume_Threshold:
			urr.IeFlags.Set(id)
			urr.VolumeThreshold = &pfcp.IEVolumeThreshold{}
			urr.VolumeThreshold.TOVOL = this.VolumeThreshold.TOVOL
			urr.VolumeThreshold.ULVOL = this.VolumeThreshold.ULVOL
			urr.VolumeThreshold.DLVOL = this.VolumeThreshold.DLVOL
			urr.VolumeThreshold.TotalVolume = this.VolumeThreshold.TotalVolume
			urr.VolumeThreshold.UplinkVolume = this.VolumeThreshold.UplinkVolume
			urr.VolumeThreshold.DownlinkVolume = this.VolumeThreshold.DownlinkVolume
			//Volume Quota	C
		case pfcp.IE_Volume_Quota:
			urr.IeFlags.Set(id)
			urr.VolumeQuota = &pfcp.IEVolumeQuota{}
			urr.VolumeQuota.TOVOL = this.VolumeQuota.TOVOL
			urr.VolumeQuota.ULVOL = this.VolumeQuota.ULVOL
			urr.VolumeQuota.DLVOL = this.VolumeQuota.DLVOL
			urr.VolumeQuota.TotalVolume = this.VolumeQuota.TotalVolume
			urr.VolumeQuota.UplinkVolume = this.VolumeQuota.UplinkVolume
			urr.VolumeQuota.DownlinkVolume = this.VolumeQuota.DownlinkVolume
			//Event Threshold	C
		case pfcp.IE_Event_Threshold:
			urr.IeFlags.Set(id)
			urr.EventThreshold = &pfcp.IEEventThreshold{}
			urr.EventThreshold.EventThreshold = this.EventThreshold
			//Event Quota	C
		case pfcp.IE_Event_Quota:
			urr.IeFlags.Set(id)
			urr.EventQuota = &pfcp.IEEventQuota{}
			urr.EventQuota.SubsequentEventQuota = this.EventQuota
			//Time Threshold	C
		case pfcp.IE_Time_Threshold:
			urr.IeFlags.Set(id)
			urr.TimeThreshold = &pfcp.IETimeThreshold{}
			urr.TimeThreshold.TimeThreshold = this.TimeThreshold
			//Time Quota	C
		case pfcp.IE_Time_Quota:
			urr.IeFlags.Set(id)
			urr.TimeQuota = &pfcp.IETimeQuota{}
			urr.TimeQuota.Value = this.TimeQuota
			//Quota Holding Time	C
		case pfcp.IE_Quota_Holding_Time:
			urr.IeFlags.Set(id)
			urr.QuotaHoldingTime = &pfcp.IEQuotaHoldingTime{}
			urr.QuotaHoldingTime.Value = this.QuotaHoldingTime
			//Dropped DL Traffic Threshold	C
		case pfcp.IE_Dropped_DL_Traffic_Threshold:
			urr.IeFlags.Set(id)
			urr.DroppedDLTrafficThreshold = &pfcp.IEDroppedDLTrafficThreshold{}
			urr.DroppedDLTrafficThreshold.DLBY = this.DroppedDLTrafficThreshold.DLBY
			urr.DroppedDLTrafficThreshold.DLPA = this.DroppedDLTrafficThreshold.DLPA
			urr.DroppedDLTrafficThreshold.NumberofBytesofDownlinkData = this.DroppedDLTrafficThreshold.NumberofBytesofDownlinkData
			urr.DroppedDLTrafficThreshold.DownlinkPackets = this.DroppedDLTrafficThreshold.DownlinkPackets
			//Monitoring Time	O
		case pfcp.IE_Monitoring_Time:
			urr.IeFlags.Set(id)
			urr.MonitoringTime = &pfcp.IEMonitoringTime{}
			urr.MonitoringTime.MonitoringTime = this.MonitoringTime
			//Subsequent Volume Threshold	O
		case pfcp.IE_Subsequent_Volume_Threshold:
			//Subsequent Time Threshold	O
		case pfcp.IE_Subsequent_Time_Threshold:
			//Subsequent Volume Quota	O
		case pfcp.IE_Subsequent_Volume_Quota:
			//Subsequent Time Quota	O
		case pfcp.IE_Subsequent_Time_Quota:
			//Subsequent Event Threshold	O
		case pfcp.IE_Subsequent_Event_Threshold:
			//Subsequent Event Quota	O
		case pfcp.IE_Subsequent_Event_Quota:
			//Inactivity Detection Time	C
		case pfcp.IE_Inactivity_Detection_Time:
			urr.IeFlags.Set(id)
			urr.InactivityDetectionTime = &pfcp.IEInactivityDetectionTime{}
			urr.InactivityDetectionTime.InactDetectionTime = this.InactivityDetectionTime
			//Linked URR ID	C
		case pfcp.IE_LinkedURR_ID:
			urr.IeFlags.Set(id)
			urr.LinkedURRID = &pfcp.IELinkedURRID{}
			urr.LinkedURRID.Value = this.LinkedURRID
			//Measurement Information	C
		case pfcp.IE_Measurement_Information:
			urr.IeFlags.Set(id)
			urr.MeasurementInformation = &pfcp.IEMeasurementInformation{}
			urr.MeasurementInformation.MBQE = this.MeasurementInformation.MBQE
			urr.MeasurementInformation.INAM = this.MeasurementInformation.INAM
			urr.MeasurementInformation.RADI = this.MeasurementInformation.RADI
			urr.MeasurementInformation.ISTM = this.MeasurementInformation.ISTM
			//FAR ID for Quota Action	C
		case pfcp.IE_FAR_ID:
			//Ethernet Inactivity Timer	C
		case pfcp.IE_Ethernet_Inactivity_Timer:
			urr.IeFlags.Set(id)
			urr.EthernetInactivityTimer = &pfcp.IEEthernetInactivityTimer{}
			urr.EthernetInactivityTimer.TimerValue = this.EthernetInactivityTimer
			//Additional Monitoring Time	O
		case pfcp.IE_Additional_Monitoring_Time:

		default:
			return nil, fmt.Errorf("Illegal IE(%d)", id)
		}
	}

	return urr, nil
}

//List of Usage Reporting Rule(s)
type UrrInfo struct {
	//UrrID uint32
	Urr URR
}

func (this *UrrInfo) String() string {
	return fmt.Sprintf("%+v", *this)
}

//List of QoS Enforcement Rule(s)
type QerInfo struct {
	//QerID uint32
	Qer QoSEnforcementRule
}

func (q *QerInfo) String() string {
	return fmt.Sprintf("%+v", *q)
}

// QER
//5.8.2.11.4	QoS Enforcement Rule
type QoSEnforcementRule struct {
	//N4 Session ID
	//SEID uint64
	//Rule ID
	RuleID uint32
	//QoS Enforcement Rule correlation ID， only used for interworking with EPC.
	QERcorrelationID uint32
	//Gate status UL/DL,Values are: open, close
	ULGateStatus GateStatusT
	DLGateStatus GateStatusT
	//nasie.ParametersIE
	//Maximum bitrate
	MBR IEMBR
	//Guaranteed bitrate
	//GBRQoSFlow
	GBR IEGBR
	//Averaging window
	AverageWindow uint32
	// QFI
	QFI uint8
	//Down-link flow level marking ,For UPF, this is for controlling the setting of the RQI in the encapsulation header.
	//
	DLFlowLevelMarking IERQI //ReflectiveQoS
	//Packet rate，only used for interworking with EPC.
	PacketRate uint64

	// presence
	IeFlags bitset.BitSet
}

func (this *QoSEnforcementRule) Fill() (*pfcp.IECreateQER, error) {
	qer := &pfcp.IECreateQER{}

	for id, e := this.IeFlags.NextSet(0); e; id, e = this.IeFlags.NextSet(id + 1) {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil, "fill QoSEnforcementRule:%d", id)
		switch id {
		// Mandatory ie
		//QER ID	M
		case pfcp.IE_QER_ID:
			qer.QERID.Value = this.RuleID
			//Gate Status	M
		case pfcp.IE_Gate_Status:
			qer.GateStatus.ULGate = uint8(this.ULGateStatus)
			qer.GateStatus.DLGate = uint8(this.DLGateStatus)
			// optional ie
			//QER Correlation ID	C
		case pfcp.IE_QER_Correlation_ID:
			qer.IeFlags.Set(id)
			qer.QERCorrelationID.Value = this.QERcorrelationID

			//Maximum Bitrate	C     MBR
		case pfcp.IE_MBR:
			qer.IeFlags.Set(id)
			qer.MaximumBitrate.ULMBR = this.MBR.ULMBR
			qer.MaximumBitrate.DLMBR = this.MBR.DLMBR
			//Guaranteed Bitrate	C GBR
		case pfcp.IE_GBR:
			qer.IeFlags.Set(id)
			qer.GuaranteedBitrate.ULGBR = this.GBR.ULGBR
			qer.GuaranteedBitrate.DLGBR = this.GBR.DLGBR
			//QoS flow identifier	C QFIs
		case pfcp.IE_QFI:
			qer.IeFlags.Set(id)
			qer.QoSflowidentifier.Value = this.QFI
			//Reflective QoS	C     RQI
		case pfcp.IE_RQI:
			qer.IeFlags.Set(id)
			qer.ReflectiveQoS.RQI = this.DLFlowLevelMarking.RQI
			//Paging Policy Indicator	C
		case pfcp.IE_Paging_Policy_Indicator:
			//Averaging Window	O
		case pfcp.IE_Averaging_Window:
			qer.IeFlags.Set(id)
			qer.AveragingWindow.AveragingWindow = this.AverageWindow
		default:
			return nil, fmt.Errorf("Illegal IE(%d)", id)
		}
	}
	return qer, nil
}

type IEMBR struct {
	ULMBR uint64
	DLMBR uint64
}
type IEGBR struct {
	ULGBR uint64
	DLGBR uint64
}
type IERQI struct {
	RQI bool
}

//Gate status UL/DL,Values are: open, close
type GateStatusT byte

const (
	Open  GateStatusT = 1
	Close GateStatusT = 0
)

//Maximum bitrate
/*-	APN-AMBR (for a QER that is referenced by all relevant Packet Detection Rules of all PDN Connections to an APN) (NOTE 1).
-	Session-AMBR (for a QER that is referenced by all relevant Packet Detection Rules of the PDU Session)
-	QoS Flow MBR (for a QER that is referenced by all Packet Detection Rules of a QoS Flow)
-	SDF MBR (for a QER that is referenced by the uplink/downlink Packet Detection Rule of a SDF)
-	Bearer MBR (for a QER that is referenced by all relevant Packet Detection Rules of a bearer) (NOTE 1).
NOTE 1: only used for interworking with EPC
*/
type MaximumBitrateT byte

const (
	APN_AMBR     MaximumBitrateT = 0
	ession_AMBR  MaximumBitrateT = 1
	QoS_Flow_MBR MaximumBitrateT = 2
	SDF_MBR      MaximumBitrateT = 3
	Bearer_MBR   MaximumBitrateT = 4
)

//Guaranteed bitrate
/*This field contains:
-	QoS Flow GBR (for a QER that is referenced by all Packet Detection Rules of a QoS Flow)
-	Bearer GBR (for a QER that is referenced by all relevant Packet Detection Rules of a bearer) (NOTE 1).
NOTE 1: only used for interworking with EPC*/
type GuaranteedBitrateT byte

const (
	QoSFlow_GBR GuaranteedBitrateT = 0
	Bearer_GBR  GuaranteedBitrateT = 1
)

// URR
//5.8.2.11.5	Usage Reporting Rule todo

// FAR
//5.8.2.11.6	Forwarding Action Rule
type ForwardActionRule struct {
	//N4 Session ID
	//SEID uint64
	//Rule ID
	RuleID uint32
	//Action
	Action ActionT //pfcp.IEApplyAction

	//ForwardingParameters start
	//Network instance//(NOTE 2)
	NetworkInstance string //*types3gpp.Apn
	//Destination interface//(NOTE 3)
	DstInterface DestinationInterfaceT //pfcp.IEDestinationInterface
	//Outer header creation//(NOTE 3)
	OuterHeaderCreation OuterHeaderCreation

	//Send end marker packet(s)//(NOTE 2)
	SendEndMarker byte
	//Transport level marking//(NOTE 3)
	//pfcp.IETransportLevelMarking
	ToSTrafficClass uint16
	//Forwarding policy//(NOTE 3)
	//pfcp.IEForwardingPolicy
	ForwardingPolicy string
	//Request for Proxying in UPF , ARP proxying
	//pfcp.IEProxying
	Proxying IEProxying
	//Container for header enrichment//(NOTE 2) ,Only relevant for the uplink direction.
	HeaderEnrichment IEHeaderEnrichment
	//Buffering Action Rule//(NOTE 5)
	BAR BAR
	//ForwardingParameters end

	// presence
	IeFlags bitset.BitSet
}

func (this *ForwardActionRule) Fill() (*pfcp.IECreateFAR, error) {
	far := &pfcp.IECreateFAR{}
	far.ForwardingParameters = &pfcp.IEForwardingParameters{}
	for id, e := this.IeFlags.NextSet(0); e; id, e = this.IeFlags.NextSet(id + 1) {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil, "fill ForwardActionRule:%d", id)
		switch id {
		//	Mandatory ie
		//	FAR ID	        M
		case pfcp.IE_FAR_ID:
			far.FARID.Value = this.RuleID
			//	Apply Action	M
		case pfcp.IE_Apply_Action:
			far.ApplyAction.Flag = uint8(this.Action)
			if this.Action == FORW {
				far.IeFlags.Set(pfcp.IE_Forwarding_Parameters)
			}
			//  optional ie
			//	Forwarding Parameters	C
			//	Mandatory ie
			//Destination Interface	M
		case pfcp.IE_Destination_Interface:
			//far.ForwardingParameters.IeFlags.Set(id)
			far.ForwardingParameters.DstInterface.Value = uint8(this.DstInterface)
			//  optional ie
			//Network Instance	    O
		case pfcp.IE_Network_Instance:
			far.ForwardingParameters.IeFlags.Set(id)
			far.ForwardingParameters.NetworkInstance = &pfcp.IENetworkInstance{} //dnn
			far.ForwardingParameters.NetworkInstance.NetworkInstance = this.NetworkInstance
			//Redirect Information	C
		case pfcp.IE_Redirect_Information:
			//Outer Header Creation 	C
		case pfcp.IE_Outer_Header_Creation:
			far.ForwardingParameters.IeFlags.Set(id)
			far.ForwardingParameters.OuterHeaderCreation = &pfcp.IEOuterHeaderCreation{}
			far.ForwardingParameters.OuterHeaderCreation.Description = this.OuterHeaderCreation.Description
			far.ForwardingParameters.OuterHeaderCreation.TEID = this.OuterHeaderCreation.TEID
			far.ForwardingParameters.OuterHeaderCreation.IPv4Addr = this.OuterHeaderCreation.IPv4Addr

			//Transport Level Marking 	C
		case pfcp.IE_Transport_Level_Marking:
			far.ForwardingParameters.IeFlags.Set(id)
			far.ForwardingParameters.TransportLevelMarking = &pfcp.IETransportLevelMarking{}
			far.ForwardingParameters.TransportLevelMarking.ToSTrafficClass = this.ToSTrafficClass
			//Forwarding Policy 	C
		case pfcp.IE_Forwarding_Policy:
			far.ForwardingParameters.IeFlags.Set(id)
			far.ForwardingParameters.ForwardingPolicy = &pfcp.IEForwardingPolicy{}
			far.ForwardingParameters.ForwardingPolicy.Identifier = this.ForwardingPolicy
			far.ForwardingParameters.ForwardingPolicy.IdentifierLength = uint8(len(this.ForwardingPolicy))
			//Header Enrichment	O
		case pfcp.IE_Header_Enrichment:
			far.ForwardingParameters.IeFlags.Set(id)
			far.ForwardingParameters.HeaderEnrichment = &pfcp.IEHeaderEnrichment{}
			far.ForwardingParameters.HeaderEnrichment.HeaderType = this.HeaderEnrichment.HeaderType
			far.ForwardingParameters.HeaderEnrichment.FieldNameLen = this.HeaderEnrichment.FieldNameLen
			far.ForwardingParameters.HeaderEnrichment.FieldName = this.HeaderEnrichment.FieldName
			far.ForwardingParameters.HeaderEnrichment.FieldValueLen = this.HeaderEnrichment.FieldValueLen
			far.ForwardingParameters.HeaderEnrichment.FieldValue = this.HeaderEnrichment.FieldValue
			//Linked Traffic Endpoint ID	C
		case pfcp.IE_Traffic_Endpoint_ID:
			//Proxying	C
		case pfcp.IE_Proxying:
			far.ForwardingParameters.IeFlags.Set(id)
			far.ForwardingParameters.Proxying = &pfcp.IEProxying{}
			far.ForwardingParameters.Proxying.ARP = this.Proxying.ARP
			far.ForwardingParameters.Proxying.INS = this.Proxying.INS

			//	Duplicating Parameters 	C
		case pfcp.IE_Duplicating_Parameters:
			//	BAR ID	O
		case pfcp.IE_Create_BAR:
			far.IeFlags.Set(pfcp.IE_BAR_ID)
			far.BARID = &pfcp.IEBARID{}
			far.BARID.Value = this.BAR.BARID

		default:
			return nil, fmt.Errorf("Illegal IE %d", id)
		}
	}
	rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil, "fill ForwardActionRule input:%+v", this)
	rlogger.Trace(types.ModuleSmfCtxt, rlogger.INFO, nil, "fill ForwardActionRule output:%+v", far)
	return far, nil
}

func (far *ForwardActionRule) String() string {
	s := fmt.Sprintf("RuleID:%d,", far.RuleID)
	s += fmt.Sprintf("Action:%s,", far.Action)
	s += fmt.Sprintf("NetworkInstance:%s,", far.NetworkInstance)
	s += fmt.Sprintf("DstInterface:%s,", far.DstInterface)
	tmp, _ := json.Marshal(far.OuterHeaderCreation)
	s += fmt.Sprintf("OuterHeaderCreation:%s,", tmp)
	s += fmt.Sprintf("SendEndMarker:%d,", far.SendEndMarker)
	s += fmt.Sprintf("ToSTrafficClass:%d,", far.ToSTrafficClass)
	s += fmt.Sprintf("ForwardingPolicy:%s,", far.ForwardingPolicy)
	s += fmt.Sprintf("Proxying:%d,", far.Proxying)
	tmp, _ = json.Marshal(far.HeaderEnrichment)
	s += fmt.Sprintf("HeaderEnrichment:%s,", tmp)
	tmp, _ = json.Marshal(far.BAR)
	s += fmt.Sprintf("BAR:%s,", tmp)
	return s
}

//Action
//forwarded, duplicated, dropped or buffered.
type ActionT byte

func (a ActionT) String() string {
	switch a {
	case FORW:
		return fmt.Sprintf("%d:forward", a)
	case DUPL:
		return fmt.Sprintf("%d:duplicat", a)
	case DROP:
		return fmt.Sprintf("%d:dropped", a)
	case BUFF:
		return fmt.Sprintf("%d:buffered", a)
	case NOCP:
		return fmt.Sprintf("%d:Notify the CP function", a)
	default:
		return fmt.Sprintf("%d:null", a)

	}

}

const (
	//forwarded
	FORW ActionT = pfcp.IEFlag_Bit_2
	//duplicated
	DUPL ActionT = pfcp.IEFlag_Bit_5
	//dropped
	DROP ActionT = pfcp.IEFlag_Bit_1
	//buffered
	BUFF ActionT = pfcp.IEFlag_Bit_3
	// Notify the CP function
	NOCP ActionT = pfcp.IEFlag_Bit_4
)

const NOCPBUFF ActionT = pfcp.IEFlag_Bit_4 + pfcp.IEFlag_Bit_3

//Destination interface//(NOTE 3)
//Contains the values "access side", "core side", "SMF" or "N6-LAN".
type DestinationInterfaceT byte

func (d DestinationInterfaceT) String() string {
	return fmt.Sprintf("%d", d)
}

const (
	DestinationInterface_Access          DestinationInterfaceT = 0
	DestinationInterface_Core            DestinationInterfaceT = 1
	DestinationInterface_SGiLAN_or_N6LAN DestinationInterfaceT = 2
	DestinationInterface_CPFunction      DestinationInterfaceT = 3
	DestinationInterface__LIFunction     DestinationInterfaceT = 4
)

//5.8.2.11.7	Usage Report generated by UPF todo

//5.8.3	Explicit Buffer Management
