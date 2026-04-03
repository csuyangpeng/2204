package nas

type NasIeId uint

//for IeFlags in PduSession Release NAS Msg
const (
	Ieid_SMCause uint = iota
	Ieid_ExtendProtocolConfigOpt
	Ieid_BackOffTimer
	FiveGSMCongestionReAttemptIn
)
