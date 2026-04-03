package nasie

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

type ARP struct {
	//mandatory
	//Defines the relative importance of a resource request. 1~15
	PriorityLevel byte

	//Defines whether a service data flow may get resources
	// that were already assigned to another service data flow with a lower priority level.
	PreemptCap PreemptionType

	//Defines whether a service data flow may lose the resources
	// assigned to it in order to admit a service data flow with higher priority level.
	PreemptVuln PreemptionType
}
type PreemptionType bool

func (p *PreemptionType) StoreWithString(val string) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	switch val {
	case "NOT_PREEMPT":
		*p = false
	case "MAY_PREEMPT":
		*p = true
	default:
		return fmt.Errorf("failed to store preemption(%s)", val)
	}
	return nil
}
