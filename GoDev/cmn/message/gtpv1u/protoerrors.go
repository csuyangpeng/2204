package gtpv1u

// 3GPP TS 29.281 V15.5.0 (2018-12)
// 9.1	Protocol Errors-- 3GPP TS 29.060 [6], subclauses 11.1.
type gtpuProtoErr struct {
	s string
}

func (g *gtpuProtoErr) Error() string {
	return "Protocol Error:" + string(g.s)
}

var (
	ErrVersion       = &gtpuProtoErr{"Different GTP Versions"}
	ErrGTPMsgLen     = &gtpuProtoErr{"GTP Message Length Errors"}
	ErrUnknownSigMsg = &gtpuProtoErr{"Unknown GTP Signalling Message"}
	ErrUnexpSigMsg   = &gtpuProtoErr{"Unexpected GTP Signalling Message"}
	//ErrMissMIE       = &gtpuProtoErr{"Missing Mandatorily Present Information Element"}
	ErrInvalidIELen = &gtpuProtoErr{"Invalid IE Length"}
	//ErrInvalidMIE    = &gtpuProtoErr{"Invalid Mandatory Information Element"}
	//ErrInvalidOIE    = &gtpuProtoErr{"Invalid Optional Information Element"}
	//ErrUnknownIE     = &gtpuProtoErr{"Unknown Information Element"}
	ErrOutSeqIE = &gtpuProtoErr{"Out of Sequence Information Elements"}
	ErrUnexpIE  = &gtpuProtoErr{"Unexpected Information Element"}
	//ErrRepIE         = &gtpuProtoErr{"Repeated Information Elements"}
	//ErrIncorrectOIE  = &gtpuProtoErr{"Incorrect Optional Information Elements"}

	ErrExtHeaderType      = &gtpuProtoErr{"Extension Header Type Error"}
	ErrUnexpExtHeaderType = &gtpuProtoErr{"Unexpected Extension header Type"}
	ErrNil                = &gtpuProtoErr{"Null pointer Error"}
)
