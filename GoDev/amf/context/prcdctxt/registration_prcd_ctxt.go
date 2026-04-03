package prcdctxt

import (
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
	T "lite5gc/cmn/types3gpp"
)

// Registration is procedure context for Registration procedure
type RegistrationPrcdCtxt struct {
	BaseCtxt
	Guti                  T.Guti5G
	GutiReallocated       bool
	NgKSI                 nasie.NasKSI
	RegistType            nas.RegistrationType
	FiveGUpdate           nasie.FiveGUpdateType
	ForPending            bool
	NeedNgRanRcu          bool
	NeedInitCtxtSetupPrcd bool
	UeSecCapablity        T.SecurityCapability
	RequestNssai          nasie.Nssai
	ResStart              [types.RessSize]byte
	TimerId               int64
}

// NewRegistration return a registration procedure context
func NewRegistrationPcrd(imsiPtr *T.Imsi) (registCtxt *RegistrationPrcdCtxt) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, imsiPtr)
	registCtxt = &RegistrationPrcdCtxt{}
	registCtxt.imsi = *imsiPtr
	registCtxt.status = st.StateRegisterStart
	return registCtxt
}

func NewRegistrationPcrdNoIMSI() (registCtxt *RegistrationPrcdCtxt) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	registCtxt = &RegistrationPrcdCtxt{}
	registCtxt.status = st.StateRegisterStart
	return registCtxt
}

// SetState set the procedure state
func (p *RegistrationPrcdCtxt) SetNextState(state string) error {
	err := st.RegistCheckState(state)
	if err != nil {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "state (%s) is not belong to RegistrationPrcdCtxt", state)
		return err
	}
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}

func (p *RegistrationPrcdCtxt) FillCtxtFromMsg(rrmsg *nasmsg.RegistrationRequestMsg) {
	p.RegistType = rrmsg.RegType
	p.FiveGUpdate = rrmsg.FiveGUpdate
	p.RequestNssai = rrmsg.RequestNssai
	p.ForPending = rrmsg.ForPending
	p.UeSecCapablity = rrmsg.UeSecCapablity
}
