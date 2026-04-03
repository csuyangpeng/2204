package n11msg

import (
	"github.com/willf/bitset"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	T "lite5gc/cmn/types3gpp"
)

//3GPP TS 29.502V15.1.0(2018-09)
// 6.1.6.2.2	Type:SmContextCreateData
type SmContextCreateData struct {
	// Supi	C
	// This IE shall be present, except if the UE is emergency registered and UICCless.
	// When present, it shall contain the subscriber permanent identify.
	Supi T.Supi

	// unauthenticated Supi	C
	// This IE shall be present if the SUPI is present in the message but is not authenticated
	// and is for an emergency registered UE.
	// When present, it shall be set as follows:
	// - true: unauthenticated SUPI;
	// - false (default): authenticated SUPI.
	UnauthenticatedSupi bool

	// Pei	C
	// This IE shall be present if the UE is emergency registered and it is either UIClless or the SUPI is not authenticated.
	// For all other cases, this IE shall be present if it is available.
	// When present, it shall contain the permanent equipment identifier.
	Pei T.Pei

	// Gpsi	C
	// This IE shall be present if it is available. When present, it shall contain the user's GPSI.
	Gpsi string

	// PduSessionId C
	// This IE shall be present, except during an EPS to 5GS Idle mode mobility or handover using the N26 interface.
	// When present, it shall contain the PDU Session ID.
	PduSessionId nas.PduSessID

	// Dnn C
	// This IE shall be present, except during an EPS to 5GS Idle mode mobility or handover using the N26 interface.
	// When present, it shall contain the requested DNN.
	Dnn T.Apn

	// Snssai  C
	// This IE shall be present during the PDU session establishment procedure.
	// In this case, it shall contain the requested S-NSSAI for the serving PLMN.
	// This corresponds to an S-NSSAI from the allowed NSSAI.
	// This IE shall also be present during an EPS to 5GS idle mode mobility or
	// handover using the N26 interface. In this case, it shall contain the S-NSSAI configured in the AMF for EPS interworking.
	SNssai nasie.SNssai

	// HplmnSnssai 	C
	// This IE shall be present for a HR PDU session,
	// except during an EPS to 5GS idle mode mobility or handover using the N26 interface.
	// When present, it shall contain the requested S-NSSAI for the HPLMN.
	// This corresponds to an S-NSSAI from the subscribed S-NSSAI corresponding to the Snssai value included in the sNssai IE.
	HplmnSnssai nasie.SNssai

	//NfInstanceId	M
	// This IE shall contain the identifier of the serving NF (e.g. serving AMF).
	ServingNfId string

	// Guami	C
	// This IE shall contain the serving AMF's GUAMI.
	// It shall be included if the NF service consumer is an AMF.
	Guami T.Guami

	// serviceName		O
	// When present, this IE shall contain the name of the AMF service to which
	// SM context status notifications are to be sent (see see subclause 6.5.2.2 of 3GPP TS 29.500 [4]).
	// This IE may be included if the NF service consumer is an AMF.
	ServiceName T.ServiceName

	// servingNetwork		M
	// This IE shall contain the serving core network operator PLMN ID.
	ServingNetwork T.PlmnID

	// RequestType	C
	// This IE shall be present if the request relates to an existing PDU session
	// or an existing emergency PDU session, except during an EPS to 5GS idle mode
	// mobility or handover using the N26 interface. It may be present otherwise.
	// When present, it shall indicate whether the request refers to a new PDU session
	// or emergency PDU session, or to an existing PDU session or emergency PDU session.
	RequestType RequestType

	// n1SmMsg	RefToBinaryData	C
	// This IE shall be present and reference the N1 SM Message binary data (see subclause 6.1.6.4.2),
	// except during an EPS to 5GS Idle mode mobility or handover using N26.
	N1SmMsg []byte // byte

	// anType	M
	// This IE shall indicate the Access Network Type to which the PDU session is to be associated.
	AnType nasmsg.RegistrationResult

	// RatType	C
	// This IE shall be present and indicate the RAT Type used by the UE, if available.
	RatType T.RatType

	// PresenceState	C
	// This IE shall be present if the DNN corresponds to a LADN. When present,
	// it shall be set to "IN" or "OUT" to indicate that the UE is in or out of the LADN service area.
	PresenceInLadn PresenceState

	// UserLocation	C
	// This IE shall contain the UE location information, if it is available. See NOTE.
	UeLocation UserLocation

	// ueTimeZone	TimeZone	C
	// This IE shall contain the UE Time Zone, if it is available.
	UeTimeZone string

	// addUeLocation		O
	// Additional UE location.
	// This IE may be present, if anType indicates a non-3GPP access and valid 3GPP access user location information is available.
	// When present, it shall contain:
	// -	the last known 3GPP access user location; and
	// -	the timestamp, if available, indicating the UTC time when the addUeLocation information was acquired.
	AddUeLocation UserLocation

	// smContextStatusUri		M
	// This IE shall include the callback URI to receive notification of SM context status.
	SmContextStatusUri string

	// hSmfUri	C
	// This IE shall be present in HR roaming scenarios. When present,
	// it shall contain the URI of the Nsmf_PDUSession service of the selected H-SMF.
	// The URI shall be formatted as specified in subclause 6.1.1.
	HSmfUri string

	// oldPduSessionId		C
	// This IE shall be present if this information is received from the UE.
	// When present, it shall contain the old PDU Session ID received from the UE.
	// See subclauses 4.3.2.2.1 and 4.3.5.2 of 3GPP TS 23.502 [3].
	OldPduSessionId nas.PduSessID

	// pduSessionsActivateList	C    1..N
	// This IE shall be present, during an EPS to 5GS Idle mode mobility using the N26 interface,
	// if received in the Registration Request from the UE.
	//When present, it shall be set as received in the Registration Request.
	// It indicates all the PDU session(s) requested to be re-activated by the UE.
	PduSessionsActivateList []nas.PduSessID

	// ueEpsPdnConnection	C
	// This IE shall be present, during an EPS to 5GS Idle mode mobility or handover using the N26 interface.
	// When present, it shall contain an MME/SGSN UE EPS PDN connection including the EPS bearer context(s).
	UeEpsPdnConnection string

	// hoState	HoState	C
	// This IE shall be present during an EPS to 5GS handover using N26 interface,
	// to request the preparation of a handover of the PDU session.
	// When present, it shall be set as specified in subclause 5.2.2.2.3.
	HoState HoState

	// additionalHsmfUri		O     1..N
	// This IE may be present in HR roaming scenarios. When present,
	// it shall contain an array of URI of the Nsmf_PDUSession service of the additional H-SMFs discovered
	// by the AMF for the given DNN,hplmnSnssai and for this PDU session. If provided, the V-SMF shall use
	// these additional H-SMF(s) if the V-SMF is not able to receive any response from the H-SMF identified by hSmfUri.
	//The URI shall be formatted as specified in subclause 6.1.1.
	AdditionalHsmfUri []string

	// pcfId		O
	// When present, this IE shall contain the identifier of the PCF selected
	// by the AMF for the UE (for Access and Mobility Policy Control);
	// it shall be the V-PCF in LBO roaming and the H-PCF in HR roaming.
	PcfId string

	// nrfUri		O
	// This IE may be present to indicate the NRF to use for PCF selection
	// within the same network slice instance. When present, the SMF shall use the NRF URI to select the PCF.
	NrfUri string

	// supportedFeatures	C
	// This IE shall be present if at least one optional feature defined in subclause 6.1.8 is supported.
	SupportedFeatures string

	// selMode		C
	// This IE shall be present if it is available. When present,
	// it shall indicate whether the requested DNN corresponds to an explicitly subscribed DNN or
	// to the usage of a wildcard subscription.
	SelMode DnnSelectionMode

	// backupAmfInfo	C	1..N	This IE shall be included
	// if the NF service consumer is an AMF and the AMF supports the AMF management without UDSF for the following cases:
	// - First interaction with SMF.
	// - Modification of the BackupAmfInfo.
	BackupAmfInfo T.BackupAmfInfo

	// traceData		C
	// This IE shall be included if trace is required to be activated (see 3GPP TS 32.422 [22]).
	//TraceData T.TraceData

	// EpsInterworkingIndication	O    0..1
	// The AMF may provide the indication when a PGW-C+SMF is selected to serve the PDU Session.
	// When present, this IE shall indicate whether the PDU session may possibly be moved
	// to EPS and whether N26 interface to be used during EPS interworking procedures.
	// The AMF may derive the value of the indication from different sources,
	// like UE radio capabilities (e.g. "S1 mode supported"), UE subscription data
	// (e.g. "Core Network Type Restriction to EPC) and configurations.
	EpsInterworkingInd EpsInterworkingIndication

	//udmGroupId	NfGroupId	O	0..1
	//When present, it shall indicate the identity of the UDM group serving the UE.
	UdmGroupId string

	//routingIndicator	string	O	0..1
	//When present, it shall indicate the Routing Indicator of the UE.
	RoutingIndicator string

	//indirectForwardingFlag	boolean	C	0..1
	//The AMF shall include this indication during N26 based Handover procedure from EPS to 5GS
	//(see 3GPP TS 23.502 [3], subclause 4.11.1.2.2),
	//to inform the SMF of the applicability or non-applicability of indirect data forwarding.
	//When present, it shall be set as follows:
	//- True: indirect data forwarding is applicable
	//- False: indirect data forwarding is not applicable
	IndirectForwardingFlag bool

	// Ie flags
	IeFlags bitset.BitSet
}

func (p *SmContextCreateData) SetSupi(supi *T.Supi) {
	p.IeFlags.Set(Ieid_supi)
	p.Supi = *supi
}

func (p *SmContextCreateData) SetUnauthenticatedSupi(value bool) {
	p.IeFlags.Set(Ieid_unauthenticatedSupi)
	p.UnauthenticatedSupi = value
}

func (p *SmContextCreateData) SetPei(pei *T.Pei) {
	p.IeFlags.Set(Ieid_pei)
	p.Pei = *pei
}

func (p *SmContextCreateData) SetGpsi(value string) {
	p.IeFlags.Set(Ieid_gpsi)
	p.Gpsi = value
}

func (p *SmContextCreateData) SetPduSessionId(value nas.PduSessID) {
	p.IeFlags.Set(Ieid_pdusessionId)
	p.PduSessionId = value
}

func (p *SmContextCreateData) SetDnn(dnn *T.Apn) {
	p.IeFlags.Set(Ieid_dnn)
	p.Dnn = *dnn
}

func (p *SmContextCreateData) SetSnssai(snssai *nasie.SNssai) {
	p.IeFlags.Set(Ieid_sNssai)
	p.SNssai = *snssai
}

func (p *SmContextCreateData) SetHplmnSnssai(hsnssai *nasie.SNssai) {
	p.IeFlags.Set(Ieid_hplmnSnssai)
	p.HplmnSnssai = *hsnssai
}

func (p *SmContextCreateData) SetGuami(guami T.Guami) {
	p.IeFlags.Set(Ieid_guami)
	p.Guami = guami
}

func (p *SmContextCreateData) SetServiceName(serName string) {
	p.IeFlags.Set(Ieid_serviceName)
	p.ServiceName = T.SetFromString(serName)
}

func (p *SmContextCreateData) SetRequestType(rtype RequestType) {
	p.IeFlags.Set(Ieid_requestType)
	p.RequestType = rtype
}

func (p *SmContextCreateData) SetN1SmMsg(msg []byte) {
	p.IeFlags.Set(Ieid_n1SmMsg)
	p.N1SmMsg = msg
}

func (p *SmContextCreateData) SetRatType(ratType T.RatType) {
	p.IeFlags.Set(Ieid_ratType)
	p.RatType = ratType
}

func (p *SmContextCreateData) SetPresenceInLadn(pInLadn PresenceState) {
	p.IeFlags.Set(Ieid_presenceInLadn)
	p.PresenceInLadn = pInLadn
}

func (p *SmContextCreateData) SetUeLocation(ul *UserLocation) {
	p.IeFlags.Set(Ieid_ueLocation)
	p.UeLocation = *ul
}

func (p *SmContextCreateData) SetUeTimeZone(utz string) {
	p.IeFlags.Set(Ieid_ueTimeZone)
	p.UeTimeZone = utz
}

func (p *SmContextCreateData) SetAddUeLocation(ul *UserLocation) {
	p.IeFlags.Set(Ieid_addUeLocation)
	p.AddUeLocation = *ul
}

func (p *SmContextCreateData) SetHSmfUri(hsmfuri string) {
	p.IeFlags.Set(Ieid_hSmfUri)
	p.HSmfUri = hsmfuri
}

func (p *SmContextCreateData) SetOldPduSessionId(psid nas.PduSessID) {
	p.IeFlags.Set(Ieid_oldPduSessionId)
	p.OldPduSessionId = psid
}

func (p *SmContextCreateData) SetPduSessionsActivateList(psList []nas.PduSessID) {
	p.IeFlags.Set(Ieid_pduSessionsActivateList)
	p.PduSessionsActivateList = append(p.PduSessionsActivateList, psList...)
}

func (p *SmContextCreateData) SetUeEpsPdnConnection(value string) {
	p.IeFlags.Set(Ieid_ueEpsPdnConnection)
	p.UeEpsPdnConnection = value
}

func (p *SmContextCreateData) SetHoState(value HoState) {
	p.IeFlags.Set(Ieid_hoState)
	p.HoState = value
}

func (p *SmContextCreateData) SetAdditionalHsmfUri(aHsmfUri []string) {
	p.IeFlags.Set(Ieid_additionalHsmfUri)
	p.AdditionalHsmfUri = append(p.AdditionalHsmfUri, aHsmfUri...)
}

func (p *SmContextCreateData) SetPcfId(value string) {
	p.IeFlags.Set(Ieid_pcfId)
	p.PcfId = value
}

func (p *SmContextCreateData) SetNrfUri(value string) {
	p.IeFlags.Set(Ieid_nrfUri)
	p.NrfUri = value
}

func (p *SmContextCreateData) SetSupportedFeatures(value string) {
	p.IeFlags.Set(Ieid_supportedFeatures)
	p.SupportedFeatures = value
}

func (p *SmContextCreateData) SetSelMode(value DnnSelectionMode) {
	p.IeFlags.Set(Ieid_selMode)
	p.SelMode = value
}

func (p *SmContextCreateData) SetBackupAmfInfo(value *T.BackupAmfInfo) {
	p.IeFlags.Set(Ieid_backupAmfInfo)
	p.BackupAmfInfo = *value
}

//func (p *SmContextCreateData) SetTraceData(value *T.TraceData) {
//	p.IeFlags.Set(Ieid_traceData)
//	p.TraceData = *value
//}

func (p *SmContextCreateData) SetEpsInterworkingInd(value EpsInterworkingIndication) {
	p.IeFlags.Set(Ieid_epsInterworkingInd)
	p.EpsInterworkingInd = value
}
func (p *SmContextCreateData) SetUdmGroupId(s string) {
	p.IeFlags.Set(Ieid_udmGroupId)
	p.UdmGroupId = s
}

func (p *SmContextCreateData) SetRoutingIndicator(s string) {
	p.IeFlags.Set(Ieid_RoutingIndicator)
	p.RoutingIndicator = s
}

func (p *SmContextCreateData) SetIndirectForwardingFlag(s bool) {
	p.IeFlags.Set(Ieid_IndirectForwardingFlag)
	p.IndirectForwardingFlag = s
}
func (p SmContextCreateData) N11MsgDataIf() {}
