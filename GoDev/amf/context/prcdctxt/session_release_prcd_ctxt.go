package prcdctxt

import (
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"

	"github.com/willf/bitset"
)

type SessionReleasePrcdCtxt struct {
	BaseCtxt
	Psi        nas.PduSessID
	UpCnxState types3gpp.PduSessStatus
	IeFlags    bitset.BitSet
	counter    int
	N2SmInfo   map[uint32][]byte
}

func (p *SessionReleasePrcdCtxt) SetCounter(c int) {
	p.counter = c
}

func (p *SessionReleasePrcdCtxt) GetCounter() int {
	return p.counter
}

func (p *SessionReleasePrcdCtxt) GetIeFlag() *bitset.BitSet {
	return &p.IeFlags
}

// SetState set the procedure state
func (p *SessionReleasePrcdCtxt) SetNextState(state string) error {
	err := st.SessionRelCheckState(state)
	if err != nil {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, nil, "state (%s) is not belong to SessionReleasePrcdCtxt", state)
		return err
	}
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}

func NewSessionRelPrcdCtxt() (sessRelCtxt *SessionReleasePrcdCtxt) {
	sessRelCtxt = &SessionReleasePrcdCtxt{}
	sessRelCtxt.status = st.StateSessRelStart
	sessRelCtxt.N2SmInfo = make(map[uint32][]byte)
	return sessRelCtxt
}
