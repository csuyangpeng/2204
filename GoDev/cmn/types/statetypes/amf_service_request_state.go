package statetype

import "lite5gc/cmn/types"

// state types definition
var (
	StateSrvReqStart    = "service_Start"
	StateSrvReqWfSerReq = "service_waitFor_ServiceRequest"

	StateSrvReqAuthStart   = "service_auth_start"
	StateSrvReqWfAuthResp  = "service_wait_for_auth_resp"
	StateSrvReqWfSecModCmp = "service_wait_for_sec_mod_cmp"
	StateSrvReqSecDone     = "service_wait_for_sec_done"

	StateSrvReqWfUpSmCtxtResp      = "service_waitFor_UpdateSmCtxtResp"
	StateSrvReqWfInitCtxtSetupResp = "service_waitFor_InitialCtxtSetupResp"
	StateSrvReqWfUpSmCtxtRespSec   = "service_waitFor_UpdateSmCtxtRespSec"
	StateSrvReqEnd                 = "service_waitFor_End"
)

// Service Request state model event definition
var (
	EventPagingReq              = "event_paging_request"
	EventSrvReqServiceReq       = "event_service_request"
	EventSrvReqServiceReqAuth   = "event_service_request_auth"
	EventSrvReqServiceReqPaging = "event_service_request_paging"

	EventSrvReqAuthResp  = "event_service_auth_response"
	EventSrvReqSecModCmp = "event_service_security_mod_complete"

	EventSrvReqUpdateSmCtxtResp  = "event_update_sm_ctxt_response"
	EventSrvReqInitCtxtSetupResp = "event_init_ctxt_setup_response"

	EventSrvReqInitCtxtSetupRespNoPsi = "event_init_ctxt_setup_response_when_noPsi"
	EventSrvReqUpdateSmCtxtRespSec    = "event_update_sm_ctxt_resp_second"
	EventSrvReqEnd                    = "event_service_request_end"
)

//RmCheckState return error, which check the validation for a state string
func SrvReqCheckState(state string) error {
	switch state {
	case StateSrvReqStart:
	case StateSrvReqWfSerReq:
	case StateSrvReqAuthStart:
	case StateSrvReqWfAuthResp:
	case StateSrvReqWfSecModCmp:
	case StateSrvReqSecDone:
	case StateSrvReqWfUpSmCtxtResp:
	case StateSrvReqWfInitCtxtSetupResp:
	case StateSrvReqWfUpSmCtxtRespSec:
	case StateSrvReqEnd:
	default:
		return types.ErrInvRmState
	}
	return nil
}
