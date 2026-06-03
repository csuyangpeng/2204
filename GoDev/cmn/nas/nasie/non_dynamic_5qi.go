package nasie

type NonDynamic5Qi struct {
	//optional
	// 1~127, 1 as the highest priority and 127 as the lowest priority.
	//When present, it contains the 5QI Priority Level value that
	// overrides the standardized or pre-configured value.
	PriorityLevel byte

	//expressed in milliseconds. Minimum = 1.
	//This IE may be present for a GBR QoS flow or a Delay Critical GBR QoS flow.
	// When present, it contains the Averaging Window that
	// overrides the standardized or pre-configured value.
	AverWindow uint

	//expressed in milliseconds. Minimum = 1.
	//This IE may be present for a Delay Critical GBR QoS flow.
	// When present, it contains the Maximum Data Burst Volume value that
	// overrides the standardized or pre-configured value
	MaxDataBurstVol uint
}
