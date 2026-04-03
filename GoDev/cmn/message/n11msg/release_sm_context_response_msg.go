package n11msg

//29502  6.1.6.2.7
//Type: SMContextRetrieveData
type ReleaseSMContextResponseData struct {
	// C	0..1
	// This IE shall be present if it is available.
	// When present, it shall contain the target MME capabilities.
	targetMmeCap MmeCapabilities
}

type MmeCapabilities struct {
	// C	0..1
	// This IE shall be present if non-IP PDN type is supported.
	// It may be present otherwise. When present, this IE shall be set as follows:
	//- true: non-IP PDN type is supported;
	//- false (default): non-IP PDN type is not supported.
	nonIpSupported bool
}

func (p ReleaseSMContextResponseData) N11MsgDataIf() {}
