package gctxt

import (
	"lite5gc/cmn/message/udmdata"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/smf/smfcontext/prcdctxt"
)

//UeContext is struct for ue information
type UeContext struct {
	IMSI types3gpp.Imsi
	Pei  types3gpp.Pei
	Gpsi string

	// pdu session context
	PduSessCtxts   map[nas.PduSessID]*PduSessContext
	PegIdleAlready bool
	PegConnAlready bool

	// sess magnment subscription data
	SessMgntSubsDataMap []udmdata.SessMgntSubscripitonData

	//serving amf
	Guami types3gpp.Guami

	//pti list
	PtiListInUse []nas.PrcdTransID //维护网络侧发出去的SM流程中的PTI
}

// NewUeContext return a UeContext with imsi
func NewUeContext(imsi *types3gpp.Imsi) *UeContext {
	return &UeContext{IMSI: *imsi,
		PduSessCtxts:        make(map[nas.PduSessID]*PduSessContext),
		SessMgntSubsDataMap: make([]udmdata.SessMgntSubscripitonData, nas.MaxPSI),
		PtiListInUse:        make([]nas.PrcdTransID, nas.MaxPTI)}
}

func (p *UeContext) GetPduSessCtxt(psi nas.PduSessID) *PduSessContext {
	rlogger.FuncEntry(types.ModuleSmfCtxt, nil)

	if val, ok := p.PduSessCtxts[psi]; ok {
		return val
	} else {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, p,
			"failed to find the pdu session context with psi(%d)", psi)
		return nil
	}

}

// GetProcCtxt return procedure context
func (p *UeContext) GetProcCtxt(psi nas.PduSessID) prcdctxt.Base {
	rlogger.FuncEntry(types.ModuleSmfCtxt, p)

	if val, ok := p.PduSessCtxts[psi]; ok {
		return val.prcCtxt
	} else {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, p,
			"failed to find the pdu session context with psi(%d)", psi)
		return nil
	}
}

func (p *UeContext) DeletePduSessCtxt(psi nas.PduSessID) error {
	rlogger.FuncEntry(types.ModuleSmfCtxt, p.GetImsiPtr())
	if _, ok := p.PduSessCtxts[psi]; ok {
		delete(p.PduSessCtxts, psi)
	} else {
		rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, p.GetImsiPtr(),
			"failed to find the pdu session context with psi(%d)", psi)
		return types.ErrFailFindSessionCtxt
	}
	return nil
}

func (p *UeContext) GetImsiPtr() *types3gpp.Imsi {
	return &p.IMSI
}
