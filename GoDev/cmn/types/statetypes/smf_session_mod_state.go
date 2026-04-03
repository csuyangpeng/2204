package statetype

import "lite5gc/cmn/types"

//Pdu Session state types definition
var (
	StatePduSessModStart                = "state_smf_pdu_sess_mod_start"
	StatePduSessModWf1stN4SessModResp   = "state_smf_pdu_sess_mod_wf_1st_n4_update_smctxt_resp"
	StatePduSessModWf1stUpdateSMCtxtReq = "state_smf_pdu_sess_mod_wf_1st_update_SM_context_request"
	StatePduSessModWf2ndN4SessModResp   = "state_smf_pdu_sess_mod_wf_2nd_n4_update_smctxt_resp"
	StatePduSessModWf2ndUpdateSMCtxtReq = "state_smf_pdu_sess_mod_wf_2nd_update_SM_context_request"
	StatePduSessModWf3rdN4SessModResp   = "state_smf_pdu_sess_mod_wf_3rd_n4_update_smctxt_resp"
	StatePduSessModWf3rdUpdateSMCtxtReq = "state_smf_pdu_sess_mod_wf_3rd_update_SM_context_request"
	StatePduSessModEnd                  = "state_smf_pdu_sess_mod_end"
)

//Pdu Session state event definition
var (
	EventPduSessMod1stUpdateSmCtxtResp  = "event_smf_pdu_session_mod_1st_update_n4_ctxt_response"
	EventPduSessMod1stUpdateN4SmCtxtReq = "event_smf_pdu_session_mod_1st_update_SM_ctxt_request"
	EventPduSessMod2ndUpdateSmCtxtResp  = "event_smf_pdu_session_mod_2nd_update_n4_ctxt_response"
	EventPduSessMod2ndUpdateN4SmCtxtReq = "event_smf_pdu_session_mod_2nd_update_SM_ctxt_request"
	EventPduSessMod3rdUpdateSmCtxtResp  = "event_smf_pdu_session_mod_3rd_update_n4_ctxt_response"
)

//Pdu Session CheckState return error, which check the validation for a state string
func PduSessModCheckState(state string) error {
	switch state {
	case StatePduSessModStart:
	case StatePduSessModWf1stN4SessModResp:
	case StatePduSessModWf1stUpdateSMCtxtReq:
	case StatePduSessModWf2ndN4SessModResp:
	case StatePduSessModWf2ndUpdateSMCtxtReq:
	case StatePduSessModWf3rdN4SessModResp:
	case StatePduSessModEnd:
	default:
		return types.ErrInvCmState
	}
	return nil
}
