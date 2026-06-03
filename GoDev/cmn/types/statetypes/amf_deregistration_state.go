package statetype

import "lite5gc/cmn/types"

// State for DeRegistration Procedure definition
var (
	StateDeRegisterStart           = "DeRegister_start"
	StateDeRegisterWfRelSmCtxtResp = "DeRegister_wait_for_smf_resp"
	StateDeRegisterWfUeCtxtRelCmp  = "DeRegister_wait_for_ueCtxtRelCmp"
	StateDeRegisterEnd             = "DeRegister_end"
)

// Event for DeRegistration Procedure definition
var (
	EventDeRegisterRequest            = "deRegist_request"
	EventDeRegisterRelSmCtxtResp      = "deRegister_release_sm_ctxt_resp"
	EventDeRegisterUeCtxtRelCmp       = "deRegister_ueCtxtRelCmp"
	EventDeRegisterRequestBySwitchOff = "deRegist request by switch Off"
)

func DeRegistCheckState(state string) error {
	switch state {
	case StateDeRegisterStart:
	case StateDeRegisterWfRelSmCtxtResp:
	case StateDeRegisterWfUeCtxtRelCmp:
	case StateDeRegisterEnd:
	default:
		return types.ErrInvDeregisterState
	}
	return nil
}
