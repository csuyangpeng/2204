package prcdctxt

import (
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

// DeRegistration is procedure context for DeRegistration procedure
type DeRegistrationPrcdCtxt struct {
	BaseCtxt
	Suci               types3gpp.Suci
	Guti5g             types3gpp.Guti5G
	DeRegistrationType nasie.DeRegistrationTypeIE
	NgKSI              nasie.NasKSI
	pduSessCounter     int
}

func (p *DeRegistrationPrcdCtxt) SetCounter(c int) {
	p.pduSessCounter = c
}

func (p *DeRegistrationPrcdCtxt) GetCounter() int {
	return p.pduSessCounter
}

func (p *DeRegistrationPrcdCtxt) SetNextState(state string) error {
	err := st.DeRegistCheckState(state)
	if err != nil {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "state (%s) is not belong to DeRegistCheckState", state)
		return err
	}
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}

func NewDeRegistration() (deregisterCtxt *DeRegistrationPrcdCtxt) {
	deregisterCtxt = &DeRegistrationPrcdCtxt{}
	deregisterCtxt.status = st.StateDeRegisterStart
	return deregisterCtxt
}
