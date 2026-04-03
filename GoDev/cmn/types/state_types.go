package types

type RMState string

const (
	StateRmDeRegistered RMState = "De-registered"
	StateRmRegistered   RMState = "Registered"
)

//RmCheckState return error, which check the validation for a state string
func RmCheckState(state RMState) error {
	switch state {
	case StateRmDeRegistered:
	case StateRmRegistered:
	default:
		return ErrInvRmState
	}
	return nil
}

type CMState string

const (
	CmIdle      CMState = "Idle"
	CmConnected CMState = "Connected"
)

func CmCheckState(state CMState) error {
	switch state {
	case CmIdle:
	case CmConnected:
	default:
		return ErrInvRmState
	}
	return nil
}

//Pdu Session state types definition
var (
	StatePduSessInactive            = "PduSessionInactive"
	StatePduSessActive              = "PduSessionActive"
	StatePduSessInactivePending     = "PduSessionInactivePending"
	StatePduSessModificationPending = "PduSessionModificationPending"
)

//Pdu Session state event definition
var (
	EventPduSessEstbAccept = "pdu_session_estb_accept"
	EventPduSessEstbReject = "pdu_session_estb_reject"
)

//Pdu Session CheckState return error, which check the validation for a state string
func PduSessCheckState(state string) error {
	switch state {
	case StatePduSessInactive:
	case StatePduSessActive:
	case StatePduSessInactivePending:
	case StatePduSessModificationPending:
	default:
		return ErrInvRmState
	}
	return nil
}