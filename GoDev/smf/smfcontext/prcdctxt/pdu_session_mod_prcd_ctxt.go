package prcdctxt

import (
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"net"
)

type PduSessionModPrcdCtxt struct {
	BaseCtxt
	PduSessId    nas.PduSessID
	Pti          nas.PrcdTransID
	Cause        nas.Sm5gCause
	IsCauseExist bool

	IMSI   types3gpp.Imsi
	DNN    types3gpp.Apn
	SNSSAI nasie.SNssai
	UEIP   net.IP
	Seid   uint64

	SbiMessage *sbicmn.SbiHandlerMessage
}

func NewPduSessionModPrcdCtxt(psi nas.PduSessID) *PduSessionModPrcdCtxt {
	prcdCtxt := &PduSessionModPrcdCtxt{}
	prcdCtxt.PduSessId = psi
	prcdCtxt.status = statetype.StatePduSessModStart
	return prcdCtxt
}

// SetState set the procedure state
func (p *PduSessionModPrcdCtxt) SetNextState(state string) error {
	err := statetype.PduSessModCheckState(state)
	if err != nil {
		return err
	}
	rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, nil, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}
