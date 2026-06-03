package prcdctxt

import (
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"

	"github.com/willf/bitset"
)

type SessionModificationPrcdCtxt struct {
	BaseCtxt
	Psi         nas.PduSessID
	UpCnxState  types3gpp.PduSessStatus
	AmfSmCtxtId uint32
	IeFlags     bitset.BitSet
	counter     int
	N2SmInfo    map[uint32][]byte
}

func (p *SessionModificationPrcdCtxt) SetCounter(c int) {
	p.counter = c
}

func (p *SessionModificationPrcdCtxt) GetCounter() int {
	return p.counter
}

func (p *SessionModificationPrcdCtxt) GetIeFlag() *bitset.BitSet {
	return &p.IeFlags
}

// SetState set the procedure state
func (p *SessionModificationPrcdCtxt) SetNextState(state string) error {
	err := st.SessionModCheckState(state)
	if err != nil {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, nil, "state (%s) is not belong to Session Modification PrcdCtxt", state)
		return err
	}
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}

func NewSessionModPrcdCtxt() (sessModCtxt *SessionModificationPrcdCtxt) {
	sessModCtxt = &SessionModificationPrcdCtxt{}
	sessModCtxt.status = st.StateSessModStart
	sessModCtxt.N2SmInfo = make(map[uint32][]byte)
	return sessModCtxt
}
