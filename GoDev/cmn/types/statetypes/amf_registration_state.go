package statetype

import "lite5gc/cmn/types"

// State for Registration Procedure definition
var (
	StateRegisterStart                = "register_start"
	StateRegisterAuthStart            = "register_auth_start"
	StateRegisterWfUdmGetAuth         = "register_auth_wait_for_udm_data"
	StateRegisterWfUdmGetAuthResp     = "register_auth_wait_for_udm_data_resp"
	StateRegisterWfAuthResp           = "register_wait_for_auth_resp"
	StateRegisterWfSecModCmp          = "register_wait_for_sec_mod_cmp"
	StateRegisterSecDone              = "register_wait_for_sec_done"
	StateRegisterWfUeCtxtTransferResp = "register_wait_for_UeContextTransferResp"
	StateRegisterWfIdentityResp       = "register_wait_for_identityResp"
	StateRegisterWfUeCmRegister       = "register_wait_for_ue_cm_registration"
	StateRegisterWfSdmGet             = "register_wait_for_sdm_get"
	StateRegisterWfInitCtxtSetupResp  = "register_wait_InitialCtxtSetupResp"
	StateRegisterWfUdmGetAmDataResp   = "register_wait_for_am_data_get"
	StateRegisterWfRegisterCmp        = "register_wait_for_registerCmp"
	StateRegisterEnd                  = "register_end"
)

// Event for Registration Procedure definition
var (
	EventRegisterRequest           = "regist_request"
	EventAuthRequest               = "auth_request"
	EventRegisterAuthResp          = "regist_auth_response"
	EventRegisterUdmGetAuth        = "regist_auth_get_udm_data"
	EventUdmGetAmData              = "regist_get_am_data"
	EventRegisterSecModCmp         = "regist_security_mod_complete"
	EventRegisterIdentityRequest   = "identity_request"
	EventRegisterInitCtxtSetupResp = "initial_ctxt_setup_response"
	EventRegisterComplete          = "register_complete"
)

//RegistCheckState return error which check the validation for the state string
func RegistCheckState(state string) error {
	switch state {
	case StateRegisterStart:
	case StateRegisterAuthStart:
	case StateRegisterWfUdmGetAuth:
	case StateRegisterWfUdmGetAuthResp:
	case StateRegisterWfUdmGetAmDataResp:
	case StateRegisterWfAuthResp:
	case StateRegisterWfSecModCmp:
	case StateRegisterSecDone:
	case StateRegisterWfUeCtxtTransferResp:
	case StateRegisterWfIdentityResp:
	case StateRegisterWfUeCmRegister:
	case StateRegisterWfSdmGet:
	case StateRegisterWfInitCtxtSetupResp:
	case StateRegisterWfRegisterCmp:
	case StateRegisterEnd:
	default:
		return types.ErrInvRegisterState
	}
	return nil
}
