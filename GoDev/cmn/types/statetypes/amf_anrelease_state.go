package statetype

import "lite5gc/cmn/types"

// state types definition
var (
	StateAnRelStart         = "anRelease_Start"
	StateAnRelWfRelCmp      = "anRelease_UeCtxtRelCmp"
	StateAnRelWfUpSmCtxtAck = "anRelease_UpdateSmCtxtAck"
	StateAnRelWfEnd         = "anRelease_End"
)

//An Release state model event definition
var (
	EventAnRelUeCtxtRelReq    = "ue_ctxt_release_request"
	EventAnRelUeCtxtRelCmp    = "ue_ctxt_release_complete"
	EventAnRelUpdateSmCtxtAck = "update_sm_ctxt_ack"
)

//RmCheckState return error, which check the validation for a state string
func AnRelCheckState(state string) error {
	switch state {
	case StateAnRelStart:
	case StateAnRelWfRelCmp:
	case StateAnRelWfUpSmCtxtAck:
	case StateAnRelWfEnd:
	default:
		return types.ErrInvRmState
	}
	return nil
}
