package statetype

import "lite5gc/cmn/types"

var (
	StatePduSessRelStart                = "smf_pdu_sess_rel_start"
	StatePduSessRelWfUpdateSmCtxtReqSec = "smf_pdu_sess_rel_waitfor_update_smctxt_request_second"
	StatePduSessRelEnd                  = "smf_pdu_sess_rel_end"
)

var (
	EventPduSessRelReq                = "smf_event_pdu_session_rel_req"
	EventPduSessRelUpdateSmCtxtReqSec = "smf_event_pdu_session_rel_update_smctxt_request_second"
)

func PduSessRelCheckState(state string) error {
	switch state {
	case StatePduSessRelStart:
	case StatePduSessRelWfUpdateSmCtxtReqSec:
	case StatePduSessRelEnd:
	default:
		return types.ErrInvRmState
	}
	return nil
}
