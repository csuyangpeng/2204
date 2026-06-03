package statetype

import "lite5gc/cmn/types"

var (
	StatePduSessSerReqStart                = "smf_pdu_sess_serreq_start"
	StatePduSessSerReqWfN4ModResp          = "smf_pdu_sess_serreq_wf_n4_modify_resp"
	StatePduSessSerReqWfUpdateSmCtxtReqSec = "smf_pdu_sess_serreq_wf_update_smctxt_req_sec"
	StatePduSessSerReqWfN4ModRespSec       = "smf_pdu_sess_serreq_wf_n4_modify_resp_sec"
	StatePduSessSerReqEnd                  = "smf_pdu_sess_serreq_end"
)

var (
	EventPduSessSerReqStart              = "smf_event_pdu_sess_serreq_start"
	EventPduSessSerReqN4ModifyResp       = "smf_event_pdu_sess_serreq_n4_modify_resp"
	EventPduSessSerReqUpdateSmCtxtReqSec = "smf_event_pdu_sess_serreq_wf_update_smctxt_req_sec"
	EventPduSessSerReqN4ModifyRespSec    = "smf_event_pdu_sess_serreq_n4_modify_resp_sec"
)

func PduSessSerReqCheckState(state string) error {
	switch state {
	case StatePduSessSerReqStart:
	case StatePduSessSerReqWfN4ModResp:
	case StatePduSessSerReqWfUpdateSmCtxtReqSec:
	case StatePduSessSerReqWfN4ModRespSec:
	case StatePduSessSerReqEnd:
	default:
		return types.ErrInvRmState
	}
	return nil
}
