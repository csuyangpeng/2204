package types

import (
	"context"
	"sync"
)

// AppContext define the contexts transfer during modules
type AppContext struct {
	Name     string
	Ctx      context.Context
	Cancel   context.CancelFunc
	Wg       *sync.WaitGroup
	ConfPath string
}

// CtxtKey define the keys in Context
type CtxtKey int

const (
	NgSctpSerCK CtxtKey = iota
	ImsiCK
	ScIdCK
	StateMgrCK
	UeContextCK
	SmHeaderCK
	SessMgntCK
	NgapLayerCK
	NgapSenderCK
	NasLayerCK
	ScTimerMgrCK
	RouterCtrlChanCK
	RouterPublishChanCK
	MsgRouterCK
	TimerTypeCK
	SmfStateMgrCK
	SmfNasLayerCK
	SmfScTimerMgrCK
	SmfPduSessCtxtCK
	SmfN11MsgDataCK
	SmfSbiHandlerMsgCK
	SmfUpdateSmCtxtReqMsgCK
	PduSessIdCK
	SbiMsgCK
	ReleaseSMContextRequestDataCK
	SessrelAcptCK
	UplinkMsgCK
	SessSmfScCK
	PduSessionEstbPrcdCtxtCK
	PduSessionAnRelSerReqCtxtCK
	GoroutineID
	SecHeaderTypeInNasMsgCK
	PduSessionModPrcdCtxtCK
)
