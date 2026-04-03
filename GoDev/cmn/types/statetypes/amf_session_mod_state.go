package statetype

import "lite5gc/cmn/types"

// state types definition
var (
	StateSessRelStart             = "amf_pdu_sess_rel_Start"
	StateSessRelWfUpSmCtxtResp    = "amf_pdu_sess_rel_UpdateSmCtxt_Resp"
	StateSessRelWfUpSmCtxtRespSec = "amf_pdu_sess_rel_N2_Resource_Release_Resp_Sec"
	StateSessRelWfUpSmCtxtReqSec  = "amf_pdu_sess_rel_UpdateSmCtxt_Resp_Sec"
	StateSessRelEnd               = "amf_pdu_sess_rel_End"
)

//state model event definition
var (
	EventSessRelReq            = "amf_event_pdu_session_rel_release_request"
	EventSessRelUpSmCtxtResp   = "amf_event_pdu_session_rel_UpdateSmCtxt_Resp"
	EventSessRelUpSmCtxtReqSec = "amf_event_pdu_session_rel_UpdateSmCtxt_Resp_Sec"
	EventSessRelEnd            = "amf_event_pdu_session_rel_End"
)

func SessionRelCheckState(state string) error {
	switch state {
	case StateSessRelStart:
	case StateSessRelWfUpSmCtxtResp:
	case StateSessRelWfUpSmCtxtRespSec:
	case StateSessRelWfUpSmCtxtReqSec:
	case StateSessRelEnd:
	default:
		return types.ErrInvDeregisterState
	}
	return nil
}
