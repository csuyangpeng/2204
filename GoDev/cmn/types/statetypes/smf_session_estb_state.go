package statetype

import "lite5gc/cmn/types"

//Pdu Session state types definition
var (
	StatePduSessEstbStart                 = "pdu_sess_estb_start"
	StatePduSessEstbWfUecmRegist          = "pdu_sess_estb_waitfor_uecm_registration"
	StatePduSessEstbWfSmGet               = "pdu_sess_estb_waitfor_sm_get"
	StatePduSessEstbWfN4SessEstbResp      = "pdu_sess_estb_waitfor_n4sess_estb_resp"
	StatePduSessEstbWfN1N2MsgTransferResp = "pdu_sess_estb_waitfor_n1n2_msg_transfer_resp"
	StatePduSessEstbWfUpdateSmCtxtReq     = "pdu_sess_estb_waitfor_update_smctxt_req"
	StatePduSessEstbWfN4SessModResp       = "pdu_sess_estb_waitfor_n4sess_modify_resp"
	StatePduSessEstbEnd                   = "pdu_sess_estb_end"
)

//Pdu Session state event definition
var (
	EventPduSessEstbReq                 = "pdu_session_estb_pdu_sess_estb_req"
	EventPduSessEstbSmGetAck           = "pdu_session_estb_sm_get_ack"
	EventPduSessEstbUecmRegistAck       = "pdu_session_estb_uecm_regist_ack"
	EventPduSessEstbN4SessEstbResp      = "pdu_session_estb_n4sess_estb_response"
	EventPduSessEstbN1N2MsgTransferResp = "pdu_session_estb_n1n2_msg_transfer_response"
	EventPduSessEstbUpdateSmCtxtReq     = "pdu_session_estb_update_smctxt_request"
	EventPduSessEstbN4SessModResp       = "pdu_session_estb_n4sess_modify_response"
)

//Pdu Session CheckState return error, which check the validation for a state string
func PduSessEstbCheckState(state string) error {
	switch state {
	case StatePduSessEstbStart:
	case StatePduSessEstbWfUecmRegist:
	case StatePduSessEstbWfSmGet:
	case StatePduSessEstbWfN4SessEstbResp:
	case StatePduSessEstbWfN1N2MsgTransferResp:
	case StatePduSessEstbWfUpdateSmCtxtReq:
	case StatePduSessEstbWfN4SessModResp:
	case StatePduSessEstbEnd:
	default:
		return types.ErrInvRmState
	}
	return nil
}
