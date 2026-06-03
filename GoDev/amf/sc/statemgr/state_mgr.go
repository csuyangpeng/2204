//Package statemgr is the manager for all FSMs in SC go routine
package statemgr

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statemgr/procedure/anrelease"
	deregister "lite5gc/amf/sc/statemgr/procedure/deregistration"
	register "lite5gc/amf/sc/statemgr/procedure/registration"
	"lite5gc/amf/sc/statemgr/procedure/servicerequest"
	"lite5gc/amf/sc/statemgr/procedure/sessionmodification"
	"lite5gc/amf/sc/statemgr/procedure/sessionrelease"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

// StateMgr control all the State Machine
type StateMgr struct {
	RegisterFsm       *register.RegisterProcFSM
	DeRegisterFsm     *deregister.DeRegisterProcFSM
	AnReleaseFsm      *anrelease.AnRelProcFSM
	ServiceRequestFsm *servicerequest.SerReqProcFSM
	SessRelRequestFsm *sessionrelease.SessRelProcFSM
	SessModRequestFsm *sessionmodification.SessModProcFSM
}

type FsmType uint8

const (
	Register FsmType = iota
	DeRegister
	AnRelease
	ServiceRequest
	SessionRelease
	SessionModify
)

func (p StateMgr) String() (strbuf string) {
	strbuf += fmt.Sprintf("StateManager Info:\n")
	strbuf += fmt.Sprintln("RegisterProcFSM: ", *p.RegisterFsm)
	strbuf += fmt.Sprintln("DeRegisterProcFSM: ", *p.DeRegisterFsm)
	strbuf += fmt.Sprintln("AnRelProcFSM: ", *p.AnReleaseFsm)
	strbuf += fmt.Sprintln("SerReqProcFSM: ", *p.ServiceRequestFsm)
	strbuf += fmt.Sprintln("SessRelProcFSM: ", *p.SessRelRequestFsm)
	strbuf += fmt.Sprintln("SessModProcFSM: ", *p.SessModRequestFsm)
	return strbuf
}

// NewStateMgr return a State Machine StateMgr
func NewStateMgr() *StateMgr {
	rlogger.FuncEntry(types.ModuleAmfState, nil)
	registerFsm, err := register.NewRegisterProcFSM()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "failed to create Regist fsm: ", err)
		return nil
	}
	deRegisterFsm, err := deregister.NewDeRegisterProcFSM()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "failed to create deRegist fsm: ", err)
		return nil
	}
	anReleaseFsm, err := anrelease.NewAnRelProcFSM()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "failed to create an Release fsm: ", err)
		return nil
	}
	serviceRequestFsm, err := servicerequest.NewSerReqProcFSM()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "failed to create service Request fsm: ", err)
		return nil
	}
	sessRelRequestFsm, err := sessionrelease.NewSessRelProcFSM()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "failed to create session release fsm: ", err)
		return nil
	}
	sessModRequestFsm, err := sessionmodification.NewSessModProcFSM()
	if err != nil {
		rlogger.Trace(types.ModuleAmfState, rlogger.FATAL, nil, "failed to create session release fsm: ", err)
		return nil
	}
	return &StateMgr{
		RegisterFsm:       registerFsm,
		DeRegisterFsm:     deRegisterFsm,
		AnReleaseFsm:      anReleaseFsm,
		ServiceRequestFsm: serviceRequestFsm,
		SessRelRequestFsm: sessRelRequestFsm,
		SessModRequestFsm: sessModRequestFsm,
	}
}

//t 状态机类型，state 设置状态， event 触发事件
func TriggerFsm(ctx context.Context, t FsmType, state string, event string) (err error) {
	stateMgr, ok := ctx.Value(types.StateMgrCK).(*StateMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "no FSM allocated")
		return types.ErrFailFindStateMgr
	}
	//get the ueCtxt
	ueCtxt := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	imsi := ueCtxt.GetImsi()
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(), "imsi(%s), state(%s),event(%s)",
		imsi.String(), state, event)
	switch t {
	case Register:
		stateMgr.RegisterFsm.Bfsm.SetState(state)
		err = stateMgr.RegisterFsm.Bfsm.Event(event, ctx)
	case DeRegister:
		stateMgr.DeRegisterFsm.Bfsm.SetState(state)
		err = stateMgr.DeRegisterFsm.Bfsm.Event(event, ctx)
	case AnRelease:
		stateMgr.AnReleaseFsm.Bfsm.SetState(state)
		err = stateMgr.AnReleaseFsm.Bfsm.Event(event, ctx)
	case ServiceRequest:
		stateMgr.ServiceRequestFsm.Bfsm.SetState(state)
		err = stateMgr.ServiceRequestFsm.Bfsm.Event(event, ctx)
	case SessionRelease:
		stateMgr.SessRelRequestFsm.Bfsm.SetState(state)
		err = stateMgr.SessRelRequestFsm.Bfsm.Event(event, ctx)
	case SessionModify:
		stateMgr.SessModRequestFsm.Bfsm.SetState(state)
		err = stateMgr.SessModRequestFsm.Bfsm.Event(event, ctx)
	default:
		return fmt.Errorf("unsupport FSM type")
	}
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger FSM, err", err)
		return err
	}
	return nil
}
