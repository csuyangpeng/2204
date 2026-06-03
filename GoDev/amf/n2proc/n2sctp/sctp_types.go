package n2sctp

// SctpOptions define the sctp options
type SctpOptions struct {
	//Sctp Initmsg Config
	NumOstreams    uint16
	MaxInstreams   uint16
	MaxAttempts    uint16
	MaxInitTimeout uint16
	// Sctp Heatbeat Config
	HeatbeatInterval uint16
	PathMaxRXT       uint16
}
