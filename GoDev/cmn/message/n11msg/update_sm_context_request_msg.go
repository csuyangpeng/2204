package n11msg

import (
	"github.com/willf/bitset"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	T "lite5gc/cmn/types3gpp"
)

//3GPP TS 29.502 V15.1.0(2018-09)
//6.1.6.2.4	Type: SMContextUpdateData
type UpdateSMContextRequestData struct {
	// Pei	C
	// This IE shall be present if it is available and has not been provided earlier to the SMF.
	// When present, this IE shall contain the permanent equipment identifier.
	Pei T.Pei

	// Gpsi	C
	// This IE shall be present if it is available and has not been provided earlier
	// to the SMF or if it has changed. When present, it shall contain the user's GPSI.
	Gpsi string

	// NfInstanceId	C
	// This IE shall be present upon inter-AMF change or mobility, or upon a N2 handover execution with AMF change.
	//When present, it shall contain the identifier of the serving NF (e.g. AMF).
	ServingNfId string

	// smContextStatusUri C
	// This IE shall be present if the servingNfId IE is present. It may be present otherwise.
	// When present, this IE shall include the callback URI to receive notification of SM context status.
	SmContextStatusUri string

	// Guami	C
	// This IE shall be present if the servingNfId of AMF is present.
	// When present, it shall contain the serving AMF's GUAMI.
	Guami T.Guami
	// servingNetwork		C
	// This IE shall be present if the servingNfId IE is present.
	// When present, it shall contain the serving core network operator PLMN ID.
	ServingNetwork T.PlmnID
	// backupAmfInfo	array(BackupAmfInfo)	C	1..N
	// This IE shall be included for the modification of the BackupAmfInfo
	// if the NF service consumer is an AMF and the AMF supports the AMF management without UDSF.
	// For deleting the backupAmfInfo, it shall contain the Null value.
	BackupAmfInfo []T.BackupAmfInfo
	// anType	AccessType	C
	// This IE shall be present upon a change of Access Network Type.
	// When present, this IE shall indicate the Access Network Type to which the PDU session is to be associated.
	AnType nasmsg.RegistrationResult
	// RatType	C
	// This IE shall be present and indicate the RAT Type used by the UE, if available, upon a change of RAT Type.
	// RAT限制定义了哪些无线接入技术（RAT）是UE不能用来接入网络的
	// 如果UE处于连接态，那么RAN对UE进行切换判决选择目标RAT时,要考虑每个网络的RAT限制，也就是不能选择一个受限的RAT作为目标切换RAT
	RatType T.RatType
	// presenceInLadn		C
	// This IE shall be present during a Service Request procedure (see subclause 5.2.2.3.2.2) ),
	// an Xn handover (see subclause 5.2.2.3.3) or a N2 handover execution (see subclause 5.2.2.3.4.3),
	// if the DNN of the PDU session corresponds to a LADN. When present,
	// it shall be set to "IN" or "OUT" to indicate that the UE is in or out of the LADN service area.
	PresenceInLadn PresenceState

	// UserLocation	C
	// This IE shall be present if it is available and if it needs to be reported to the SMF
	// (e.g. the user location has changed or the user plane of the PDU session is deactivated).
	// When present, this IE shall contain:
	//  -	the UE location information; and
	//  -	the timestamp, if available, indicating the UTC time when the UeLocation information was acquired.
	UeLocation UserLocation
	// ueTimeZone		C
	// This IE shall be present if it is available, the UE Time Zone has changed and needs to be reported to the SMF.
	// When present, this IE shall contain the UE Time Zone.
	UeTimeZone string
	// addUeLocation		O
	// Additional UE location.
	// This IE may be present, if anType indicates a non-3GPP access and a valid 3GPP access user location information is available.
	// When present, it shall contain:
	// -	the last known 3GPP access user location; and
	// -	the timestamp, if available, indicating the UTC time when the addUeLocation information was acquired.
	AddUeLocation UserLocation
	// upCnxState	UpCnxState	C
	// This IE shall be present to request the activation or the deactivation of the user plane connection of the PDU session.
	// When present, it shall be set as specified in subclause 5.2.2.3.2.
	UpCnxState UpCnxState
	// hoState	HoState	C
	// This IE shall be present to request the preparation, execution or cancellation of a handover of the PDU session.
	// When present, it shall be set as specified in subclause 5.2.2.3.4.
	HoState HoState
	// toBeSwitched		C
	// This IE shall be present during an Xn Handover (see subclause 5.2.2.3.3)
	// to request to switch the PDU session to a new downlink N3 tunnel endpoint.
	//
	//When present, it shall be set as follows:
	//- true: request to switch to the PDU session.
	//- false (default): no request to switch the PDU session.

	ToBeSwitched bool
	// failedToBeSwitched		C
	// This IE shall be present during an Xn Handover (see subclause 5.2.2.3.3)
	// if the PDU session failed to be setup in the target RAN.
	//
	//When present, it shall be to true to indicate that the PDU session failed to be setup in the target RAN.
	FailedToBeSwitched bool

	// n1SmMsg	RefToBinaryData	C
	// This IE shall be present if N1 SM Information has been received from the UE.
	// When present, this IE shall reference the N1 SM Message binary data (see subclause 6.1.6.4.2).
	N1SmMsg []byte
	// n2SmInfo	RefToBinaryData	C
	// This IE shall be present if N2 SM Information has been received from the AN.
	// When present, this IE shall reference the N2 SM Information binary data (see subclause 6.1.6.4.3).
	N2SmInfo string
	// n2SmInfoType		C
	// This IE shall be present if "n2SmInfo" attribute is present.
	//When present, this IE shall indicate the NGAP IE type for the NG AP SMF related IE container carried in "n2SmInfo" attribute.
	N2SmInfoType N2SmInfoType
	// targetServingNfId		C
	// This IE shall be present during a N2 handover preparation with AMF change.
	// When present, it shall contain the identifier of the target serving NF (e.g. AMF).
	TargetServingNfId int
	// dataForwarding	boolean	C
	// This IE shall be present and set as specified in subclause 5.2.2.3.9 during a 5GS to EPS handover.
	// When present, it shall be set as follows:
	// - true: indirect data forwarding is required;
	// - false (default): indirect data forwarding is not required.
	DataForwarding bool
	// epsBearerSetup	array(EpsBearerContainer)	C	0..N
	// This IE shall be present during a 5GS to EPS handover using the N26 interface.
	// When present, it shall include the EPS bearer context(s) successfully setup in EPS.
	EpsBearerSetup []string
	// revokeEbiList	array(EpsBearerId)	C	0..N
	// This IE shall be present to request the SMF to revoke some EBIs
	// (see subclause 4.11.1.4.1 of 3GPP TS 23.502 [3]). When present, it shall contain the EBIs to revoke.
	RevokeEbiList []int
	// release		C
	// This IE shall be used to indicate a network initiated PDU session release is requested.
	// This IE shall be present and set as specified in subclause 5.2.2.3.10 during P-CSCF restoration procedure.
	// When present, it shall be set as follows:
	// - true: PDU session release is required;
	// - false (default): PDU session release is not required.
	Release bool
	// 	Cause	O
	// When present, this IE shall indicate the cause for the requested modification,
	// e.g. the NF Service Consumer cause for requesting to deactivate the user plane connection of the PDU session.
	Cause Cause
	// NgApCause	C
	// This IE shall be present, if the information is available.
	// When present, this IE shall indicate the cause for the requested modification,
	// e.g. the NGAP cause for requesting to deactivate the user plane connection of the PDU session.
	NgApCause int
	// 5GMmCauseValue	5GMmCause	C
	// This IE shall be included if the AMF received a 5GMM cause code
	// from the UE during any network initiated PDU session modification or release procedure.
	// (e.g 5GMM Status message in response to a Downlink NAS Transport message carrying 5GSM payload).
	Mm5GCauseValue nas.Mm5gCause
	// sNssai		C
	// This IE shall be present, during an EPS to 5GS idle mode mobility or handover using the N26 interface,
	// if the S-NSSAI for the serving PLMN derived from the S-NSSAI of the home PLMN differs
	// from the S-NSSAI provided in the Create SM Context Request.
	// When present, it shall contain the S-NSSAI for the serving PLMN.
	SNssai nasie.SNssai
	// TraceData	C
	// This IE shall be included if trace is required to be activated,
	// modified or deactivated (see 3GPP TS 32.422 [22]).
	//For trace modification, it shall contain a complete replacement of trace data.
	//For trace deactivation, it shall contain the Null value.
	//TraceData T.TraceData

	// EpsInterWorkingIndication	O
	// This IE may be present if the indication has been provided during the PDU session creation,
	// and its value has changed after session creation or last update.
	// When present, this IE shall indicate whether the PDU session may possibly be moved
	// to EPS and whether N26 interface to be used during EPS interWorking procedures.
	EpsInterWorkingInd EpsInterworkingIndication

	//anTypeCanBeChanged	boolean	C	0..1
	// This IE shall be present and set to true to indicate
	// that the Access Network Type associated to the PDU session can be changed
	// (see subclause 5.2.2.3.2.4), during a Service Request procedure
	// (see subclause 4.2.3.2 of 3GPP TS 23.502 [3])), in response to paging or
	// NAS notification indicating non-3GPP access, when the PDU Session for
	// which the UE was paged or notified is in the List Of Allowed PDU Sessions provided by the UE,
	// and the AMF received N2 SM Information only or N1 SM Container and
	// N2 SM Information from the SMF in step 3a of subclause 4.2.3.3 of 3GPP TS 23.502 [3].
	//When present, it shall be set as follows:
	//- true: the access type of the PDU session can be changed.
	//- false: the access type of the PDU session cannot be changed (default).
	AnTypeCanBeChanged bool

	// Ie flags
	IeFlags bitset.BitSet
}

func (p UpdateSMContextRequestData) N11MsgDataIf() {}
