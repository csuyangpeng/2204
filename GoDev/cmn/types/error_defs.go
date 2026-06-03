package types

import "errors"

// Common error definition
var (
	ErrInvParm     = errors.New("invalid parameter")
	ErrWorngIpcMsg = errors.New("incorrect ipc message")
)

// Error related to message process
var (
	ErrFailFindUeCtxt = errors.New("fail to find UE Context")

	ErrFailDelUeCtxt     = errors.New("fail to delete UE Context")
	ErrFailTriggerFsm    = errors.New("fail to trigger fsm")
	ErrFailFindGnbInstId = errors.New("fail to find gnb instance id")
	ErrFailEncodeNasMsg  = errors.New("fail to encode nas msg")
	ErrFailSendNgapMsg   = errors.New("fail to send ngap msg")
	ErrFailSendNasMsg    = errors.New("fail to send nas msg")

	ErrFailRegisterIdMgr = errors.New("fail to register id mgr")
	ErrFailReserveIdMgr  = errors.New("fail to reserve id mgr")
	ErrFailGetIdMgr  = errors.New("fail to get id mgr")

	ErrFailGetProcedureCtxt = errors.New("fail to get procedure context")
	ErrFailFindSessionCtxt  = errors.New("fail to find session Context")
	ErrFailDelSessionCtxt  = errors.New("fail to del session Context")

	ErrFailFindNgapLayer = errors.New("failed to find sc ngap egress")
	ErrFailFindTimerMgr  = errors.New("failed to find timer mgr")
	ErrFailFindStateMgr  = errors.New("failed to find state mgr")
	ErrFailSetState  = errors.New("failed to set state")
	ErrFailFindNasLayer  = errors.New("failed to find nas layer")

	ErrFailGetImsi        = errors.New("failed to find imsi")
	ErrInvRmState         = errors.New("invalid register management state")
	ErrInvCmState         = errors.New("invalid session management state")
	ErrInvRegisterState   = errors.New("invalid register procedure state")
	ErrInvDeregisterState = errors.New("invalid deregister procedure state")

	ErrInputParaNil = errors.New("input parameter is nil")

	ErrFailParseDNN = errors.New("fail to parse dnn")
)
