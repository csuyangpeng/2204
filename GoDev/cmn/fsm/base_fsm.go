package fsm

import (
	"fmt"
)

// StateModel define a basic state model structure
type StateModel struct {
	Event  string
	Src    string
	Dest   string
	CbFunc Callback
}

// BaseFsm is a wrapper for fsm, a third party fsm impltementation
type BaseFsm struct {
	Bfsm *FSM
}

func (p BaseFsm) String() (strbuf string) {
	strbuf += fmt.Sprintln("BaseFsm Info: ")
	strbuf += fmt.Sprintln(p.Bfsm)

	return strbuf
}

// NewBaseFsm return a BaseFsm pointer with intial state
func NewBaseFsm(initState string) *BaseFsm {
	fsm := NewFSM(
		initState,
		Events{},
		Callbacks{},
	)
	return &BaseFsm{Bfsm: fsm}
}

// RegisterEvent register a event trigger module into the FSM
func (p *BaseFsm) RegisterEvent(event string, srcState []string, dstState string, cbFunc Callback) error {
	if p == nil {
		return fmt.Errorf("invalid BaseFsm")
	}

	eventDesc := EventDesc{
		Name: event,
		Src:  srcState,
		Dst:  dstState,
	}
	events := []EventDesc{eventDesc}
	p.Bfsm.AddEventDesc(events)

	callbacks := map[string]Callback{event: cbFunc}
	p.Bfsm.AddCallBack(callbacks)
	return nil
}

// CreateFsm return a native fsm
func CreateFsm(initState string) *FSM {
	fsm := NewFSM(
		initState,
		Events{},
		Callbacks{},
	)
	return fsm
}
