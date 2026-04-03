package nasie

import "fmt"

type Dynamic5Qi struct {
	//mandatory
	ResourceType QosResourceType

	// 1~127, 1 as the highest priority and 127 as the lowest priority.
	PriorityLevel byte

	//expressed in milliseconds. Minimum = 1.
	PacketDelayBudget uint

	//Examples:
	//Packer Error Rate 10-6 shall be encoded as "6".
	//Packer Error Rate 10-2 shall be encoded as "2".
	PacketErrRate uint

	//optional
	//This IE shall be present only for a GBR QoS flow or a Delay Critical GBR QoS flow.
	//expressed in milliseconds. Minimum = 1.
	AverWindow uint

	//This IE shall be present for a Delay Critical GBR QoS flow.
	//expressed in milliseconds. Minimum = 1.
	MaxDataBurstVol uint
}

type QosResourceType byte

const (
	NON_GBR          QosResourceType = 1 //Non-GBR QoS Flow.
	NON_CRITICAL_GBR QosResourceType = 2 //Non-delay critical GBR QoS flow.
	CRITICAL_GBR     QosResourceType = 2 //Delay critical GBR QoS flow.
)

func (p *QosResourceType) StoreWithString(val string) error {
	switch val {
	case "NON_GBR":
		*p = NON_GBR
	case "NON_CRITICAL_GBR":
		*p = NON_CRITICAL_GBR
	case "CRITICAL_GBR":
		*p = CRITICAL_GBR
	default:
		return fmt.Errorf("invalid qos resource type (%s)", val)
	}
	return nil
}
