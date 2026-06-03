package statetype

import "lite5gc/cmn/types"

// state types definition
var (
	StateSessModStart              = "state_amf_pdu_sess_mod_Start"
	StateSessModWf1stUpSmCtxtResp  = "state_amf_pdu_sess_mod_wf_1st_UpdateSmCtxt_Resp"
	StateSessModWfN2SessResp       = "state_amf_pdu_sess_mod_wf_N2_session_response"
	StateSessModWf2ndUpSmCtxtResp  = "state_amf_pdu_sess_mod_wf_2nd_UpdateSmCtxt_Resp"
	StateSessModWfN2NASUplinkTrans = "state_amf_pdu_sess_mod_wf_N2_NAS_uplink_transfer"
	StateSessModWf3rdUpSmCtxtResp  = "state_amf_pdu_sess_mod_wf_3rd_UpdateSmCtxt_Resp"
	StateSessModEnd                = "state_amf_pdu_sess_mod_End"
)

//state model event definition
var (
	EventSessModWf1stUpSmCtxtReq   = "event_amf_pdu_sess_mod_1st_UpdateSmCtxt_Resp"
	EventSessModWfN2SessReq        = "event_amf_pdu_sess_mod_N2_session_response"
	EventSessModWf2ndUpSmCtxtReq   = "event_amf_pdu_sess_mod_2nd_UpdateSmCtxt_Resp"
	EventSessModWfN2NASUplinkTrans = "event_amf_pdu_sess_mod_N2_NAS_uplink_transfer"
	EventSessModWf3rdUpSmCtxtReq   = "event_amf_pdu_sess_mod_3rd_UpdateSmCtxt_Resp"
	EventSessModEnd                = "event_amf_pdu_sess_mod_end"
)

func SessionModCheckState(state string) error {
	switch state {
	case StateSessModStart:
	case StateSessModWf1stUpSmCtxtResp:
	case StateSessModWfN2SessResp:
	case StateSessModWf2ndUpSmCtxtResp:
	case StateSessModWfN2NASUplinkTrans:
	case StateSessModWf3rdUpSmCtxtResp:
	case StateSessModEnd:
	default:
		return types.ErrInvDeregisterState
	}
	return nil
}
