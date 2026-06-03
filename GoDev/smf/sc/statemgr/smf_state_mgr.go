package statemgr

import (
	"context"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/sc/statemgr/procedure/sessestbprcd"
	"lite5gc/smf/sc/statemgr/procedure/sessmodprcd"
	"lite5gc/smf/sc/statemgr/procedure/sessrelprcd"
	"lite5gc/smf/sc/statemgr/procedure/sessserreqprcd"
)

// SmfStateMgr control all the State Machine
type SmfStateMgr struct {
	PduSessESTBFsm *sessestbprcd.PduSessEstbPrcdFSM
	PduSessRelFsm  *sessrelprcd.PduSessRelPrcdFSM
	PduSessModFsm  *sessmodprcd.PduSessModPrcdFSM
	SerReqFsm      *sessserreqprcd.PduSessSerReqPrcdFSM
}

type FsmType uint8

const (
	SessionESTB FsmType = iota
	SessionRel
	SessionMod
	SerReq
)

// NewStateMgr return a State Machine SmfStateMgr
func NewStateMgr() *SmfStateMgr {
	rlogger.FuncEntry(types.ModuleSmfState, nil)

	pduSessRelFsm, err := sessrelprcd.NewPduSessRelPrcdFSM()
	if err != nil {
		rlogger.Trace(types.ModuleSmfState, rlogger.FATAL, nil, "failed to create fsm: ", err)
		return nil
	}
	pduSessEstbFsm, err := sessestbprcd.NewPduSessEstbPrcdFSM()
	if err != nil {
		rlogger.Trace(types.ModuleSmfState, rlogger.FATAL, nil, "failed to create fsm: ", err)
		return nil
	}
	pduSessModFsm, err := sessmodprcd.NewPduSessModPrcdFSM()
	if err != nil {
		rlogger.Trace(types.ModuleSmfState, rlogger.FATAL, nil, "failed to create fsm: ", err)
		return nil
	}
	serReqFsm, err := sessserreqprcd.NewPduSessSerReqPrcdFSM()
	if err != nil {
		rlogger.Trace(types.ModuleSmfState, rlogger.FATAL, nil, "failed to create fsm: ", err)
		return nil
	}
	return &SmfStateMgr{
		PduSessRelFsm:  pduSessRelFsm,
		PduSessESTBFsm: pduSessEstbFsm,
		PduSessModFsm:  pduSessModFsm,
		SerReqFsm:      serReqFsm,
	}
}

//t 状态机类型，state 设置状态， event 触发事件
func TriggerSmfFsm(ctx context.Context, t FsmType, state string, event string) (err error) {
	stateMgr, ok := ctx.Value(types.SmfStateMgrCK).(*SmfStateMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "no FSM allocated")
		return types.ErrFailFindStateMgr
	}
	switch t {
	case SessionESTB:
		stateMgr.PduSessESTBFsm.Bfsm.SetState(state)
		err = stateMgr.PduSessESTBFsm.Bfsm.Event(event, ctx)
	case SessionRel:
		stateMgr.PduSessRelFsm.Bfsm.SetState(state)
		err = stateMgr.PduSessRelFsm.Bfsm.Event(event, ctx)
	case SessionMod:
		stateMgr.PduSessModFsm.Bfsm.SetState(state)
		err = stateMgr.PduSessModFsm.Bfsm.Event(event, ctx)
	case SerReq:
		stateMgr.SerReqFsm.Bfsm.SetState(state)
		err = stateMgr.SerReqFsm.Bfsm.Event(event, ctx)
	default:
		return fmt.Errorf("unsupport FSM type")
	}
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger FSM, err", err)
		return err
	}
	return nil
}
