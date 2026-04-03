package types3gpp

type NgapMessage interface {
	Encode() []byte
	Decode(msgbuf []byte) error
}

type MsgType int

const (
	InitiatingMessage   MsgType = 1
	SuccessfulOutcome   MsgType = 2
	UnsuccessfulOutcome MsgType = 3
	SctpShutdown        MsgType = 4
	TaTblUpdate         MsgType = 5
)

func (p MsgType) String() string {
	switch p {
	case InitiatingMessage:
		return "InitiatingMessage"
	case SuccessfulOutcome:
		return "SuccessfulOutcome"
	case UnsuccessfulOutcome:
		return "UnsuccessfulOutcome"
	case SctpShutdown:
		return "Sctp Shutdown"
	default:
		return "Unknown"
	}
}

type ProcedureCode int

const (
	AMFConfigurationUpdate                ProcedureCode = 0
	AMFStatusIndication                   ProcedureCode = 1
	CellTrafficTrace                      ProcedureCode = 2
	DeactivateTrace                       ProcedureCode = 3
	DownlinkNASTransport                  ProcedureCode = 4
	DownlinkNonUEAssociatedNRPPaTransport ProcedureCode = 5
	DownlinkRANConfigurationTransfer      ProcedureCode = 6
	DownlinkRANStatusTransfer             ProcedureCode = 7
	DownlinkUEAssociatedNRPPaTransport    ProcedureCode = 8
	ErrorIndication                       ProcedureCode = 9
	HandoverCancel                        ProcedureCode = 10
	HandoverNotification                  ProcedureCode = 11
	HandoverPreparation                   ProcedureCode = 12
	HandoverResourceAllocation            ProcedureCode = 13
	InitialContextSetup                   ProcedureCode = 14
	InitialUEMessage                      ProcedureCode = 15
	LocationReportingControl              ProcedureCode = 16
	LocationReportingFailureIndication    ProcedureCode = 17
	LocationReport                        ProcedureCode = 18
	NASNonDeliveryIndication              ProcedureCode = 19
	NGReset                               ProcedureCode = 20
	NGSetup                               ProcedureCode = 21
	OverloadStart                         ProcedureCode = 22
	OverloadStop                          ProcedureCode = 23
	Paging                                ProcedureCode = 24
	PathSwitchRequest                     ProcedureCode = 25
	PDUSessionResourceModify              ProcedureCode = 26
	PDUSessionResourceModifyIndication    ProcedureCode = 27
	PDUSessionResourceRelease             ProcedureCode = 28
	PDUSessionResourceSetup               ProcedureCode = 29
	PDUSessionResourceNotify              ProcedureCode = 30
	PrivateMessage                        ProcedureCode = 31
	PWSCancel                             ProcedureCode = 32
	PWSFailureIndication                  ProcedureCode = 33
	PWSRestartIndication                  ProcedureCode = 34
	RANConfigurationUpdate                ProcedureCode = 35
	RerouteNASRequest                     ProcedureCode = 36
	RRCInactiveTransitionReport           ProcedureCode = 37
	TraceFailureIndication                ProcedureCode = 38
	TraceStart                            ProcedureCode = 39
	UEContextModification                 ProcedureCode = 40
	UEContextRelease                      ProcedureCode = 41
	UEContextReleaseRequest               ProcedureCode = 42
	UERadioCapabilityCheck                ProcedureCode = 43
	UERadioCapabilityInfoIndication       ProcedureCode = 44
	UETNLABindingRelease                  ProcedureCode = 45
	UplinkNASTransport                    ProcedureCode = 46
	UplinkNonUEAssociatedNRPPaTransport   ProcedureCode = 47
	UplinkRANConfigurationTransfer        ProcedureCode = 48
	UplinkRANStatusTransfer               ProcedureCode = 49
	UplinkUEAssociatedNRPPaTransport      ProcedureCode = 50
	WriteReplaceWarning                   ProcedureCode = 51

	MaxProcedureCode ProcedureCode = 100
)

func (p ProcedureCode) String() string {
	var rt string
	switch p {
	case AMFConfigurationUpdate:
		rt = "AMFConfigurationUpdate"
	case AMFStatusIndication:
		rt = "AMFStatusIndication"
	case CellTrafficTrace:
		rt = "CellTrafficTrace"
	case DeactivateTrace:
		rt = "DeactivateTrace"
	case DownlinkNASTransport:
		rt = "DownlinkNASTransport"
	case DownlinkNonUEAssociatedNRPPaTransport:
		rt = "DownlinkNonUEAssociatedNRPPaTransport"
	case DownlinkRANConfigurationTransfer:
		rt = "DownlinkRANConfigurationTransfer"
	case DownlinkRANStatusTransfer:
		rt = "DownlinkRANStatusTransfer"
	case DownlinkUEAssociatedNRPPaTransport:
		rt = "DownlinkUEAssociatedNRPPaTransport"
	case ErrorIndication:
		rt = "ErrorIndication"
	case HandoverCancel:
		rt = "HandoverCancel"
	case HandoverNotification:
		rt = "HandoverNotification"
	case HandoverPreparation:
		rt = "HandoverPreparation"
	case HandoverResourceAllocation:
		rt = "HandoverResourceAllocation"
	case InitialContextSetup:
		rt = "InitialContextSetup"
	case InitialUEMessage:
		rt = "InitialUEMessage"
	case LocationReportingControl:
		rt = "LocationReportingControl"
	case LocationReportingFailureIndication:
		rt = "LocationReportingFailureIndication"
	case LocationReport:
		rt = "LocationReport"
	case NASNonDeliveryIndication:
		rt = "NASNonDeliveryIndication"
	case NGReset:
		rt = "NGReset"
	case NGSetup:
		rt = "NGSetup"
	case OverloadStart:
		rt = "OverloadStart"
	case OverloadStop:
		rt = "OverloadStop"
	case Paging:
		rt = "Paging"
	case PathSwitchRequest:
		rt = "PathSwitchRequest"
	case PDUSessionResourceModify:
		rt = "PDUSessionResourceModify"
	case PDUSessionResourceModifyIndication:
		rt = "PDUSessionResourceModifyIndication"
	case PDUSessionResourceRelease:
		rt = "PDUSessionResourceRelease"
	case PDUSessionResourceSetup:
		rt = "PDUSessionResourceSetup"
	case PDUSessionResourceNotify:
		rt = "PDUSessionResourceNotify"
	case PrivateMessage:
		rt = "PrivateMessage"
	case PWSCancel:
		rt = "PWSCancel"
	case PWSFailureIndication:
		rt = "PWSFailureIndication"
	case PWSRestartIndication:
		rt = "PWSRestartIndication"
	case RANConfigurationUpdate:
		rt = "RANConfigurationUpdate"
	case RerouteNASRequest:
		rt = "RerouteNASRequest"
	case RRCInactiveTransitionReport:
		rt = "RRCInactiveTransitionReport"
	case TraceFailureIndication:
		rt = "TraceFailureIndication"
	case TraceStart:
		rt = "TraceStart"
	case UEContextModification:
		rt = "UEContextModification"
	case UEContextRelease:
		rt = "UEContextRelease"
	case UEContextReleaseRequest:
		rt = "UEContextReleaseRequest"
	case UERadioCapabilityCheck:
		rt = "UERadioCapabilityCheck"
	case UERadioCapabilityInfoIndication:
		rt = "UERadioCapabilityInfoIndication"
	case UETNLABindingRelease:
		rt = "UETNLABindingRelease"
	case UplinkNASTransport:
		rt = "UplinkNASTransport"
	case UplinkNonUEAssociatedNRPPaTransport:
		rt = "UplinkNonUEAssociatedNRPPaTransport"
	case UplinkRANConfigurationTransfer:
		rt = "UplinkRANConfigurationTransfer"
	case UplinkRANStatusTransfer:
		rt = "UplinkRANStatusTransfer"
	case UplinkUEAssociatedNRPPaTransport:
		rt = "UplinkUEAssociatedNRPPaTransport"
	case WriteReplaceWarning:
		rt = "WriteReplaceWarning"
	default:
		rt = "Unknown"
	}
	return rt
}
