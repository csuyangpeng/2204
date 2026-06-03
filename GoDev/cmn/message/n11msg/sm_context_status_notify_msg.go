package n11msg

//3GPP TS 29.502V15.1.0(2018-09)
//6.1.6.2.8	Type: SMContextStatusNotification
type SmContextStatusNotifyMsg struct {
	// StatusInfo	M
	// This IE shall contain status information about the SM context.
	statusInfo StatusInfo
}

type StatusInfo struct {
	resourceStatus ResourceStatus
	cause          Cause
}

type ResourceStatus byte

const (
	RELEASED ResourceStatus = 1
)

type Cause byte

const (
	//Release due to Handover
	REL_DUE_TO_HO Cause = iota

	//Mobility due to EPS fallback for IMS voice is on-going.
	EPS_FALLBACK

	//Release due to user plane Security requirements that cannot be fulfilled.
	REL_DUE_TO_UP_SEC

	//Release due to the DNN based congestion control.
	DNN_CONGESTION

	//Release due to the S-NSSAI based congestion control.
	SNSSAI_CONGESTION

	//Release due to PDU session reactivation.
	REL_DUE_TO_REACTIVATION

	//The 5G AN did not respond to the request initiated by the network.
	G5_AN_NOT_RESPONDING

	//Release due to the associated S-NSSAI becomes no longer available.
	REL_DUE_TO_SLICE_NOT_AVAILABLE
)
