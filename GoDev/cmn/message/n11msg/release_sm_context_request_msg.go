package n11msg

import (
	"lite5gc/cmn/nas"
	"lite5gc/cmn/types3gpp"
)

//29502  6.1.6.2.6
//Type: SMContextReleaseData
type ReleaseSMContextRequestData struct {
	// C	0..1
	// This IE shall be present, if the information is available.
	// When present, this IE shall indicate the NF Service Consumer cause for the requested SM context release.
	Cause Cause

	// C	0..1
	// This IE shall be present, if the information is available.
	// When present, this IE shall indicate the NGAP cause for the requested SM context release.
	NgApCause types3gpp.NgapCause

	//C	0..1
	// This IE shall be included if the PDU session is released by the AMF due to any 5GMM failure.
	// When present, this IE shall contain the 5GMM cause code value received from the UE.
	Mm5GCauseValue nas.Mm5gCause

	//C	0..1	This IE shall be present, if available.
	//When present, it shall contain the UE location information. See NOTE.
	UeLocation UserLocation

	//C	0..1	This IE shall be present, if available.
	//When present, it shall contain the UE location information.
	UeTimeZone string

	//O	0..1	Additional UE location.
	//This IE may be present, if anType previously reported is a non-3GPP access and a valid 3GPP access user location information is available.
	//When present, it shall contain:
	//-	the last known 3GPP access user location; and
	//-	the timestamp, if available, indicating the UTC time when the addUeLocation information was acquired.
	//See NOTE.
	AddUeLocation UserLocation

	//C	0..1	This IE shall be present and set to "true" during a 5GS to EPS Idle mode mobility or handover, for a Home Routed PDU session. When present, it shall be set as follows:
	//- true: release the SM context and PDU session in the V-SMF only;
	//- false (default): release the SM context and PDU session in V-SMF and H-SMF.
	VsmfReleaseOnly bool
}

func (p ReleaseSMContextRequestData) N11MsgDataIf() {}
