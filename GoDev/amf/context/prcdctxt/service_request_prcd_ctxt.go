package prcdctxt

import (
	"github.com/willf/bitset"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	st "lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

const (
	SerReqUpDataSmCtxtRespFirst  = 1
	SerReqUpDataSmCtxtRespSecond = 2
)

// DeRegistration is procedure context for DeRegistration procedure
type ServiceRequestPrcdCtxt struct {
	BaseCtxt
	NgKSI                       nasie.NasKSI
	ServiceType                 nas.ServiceTYPE
	MobileIdentity              nasie.MobileIdentity
	ResStart                    [types.RessSize]byte
	UeSecCapablity              types3gpp.SecurityCapability
	Order                       int // 1: first, 2 second
	UplinkDataStatus            nasie.SessionStatus
	PDUSessionStatus            nasie.SessionStatus
	AllowedPDUSessionStatus     nasie.SessionStatus
	IeFlags                     bitset.BitSet
	counter                     int
	PduSessResSetupRespList     []types3gpp.PduSessResSetupRespItem
	PduSessResFailedToSetupList []types3gpp.PduSessResFailedToSetupItem
	GnbConnID                   uint32
	GnbInfo                     types3gpp.GnbInfo
	Tai                         types3gpp.TAI
	UpCnxState                  types3gpp.PduSessStatus
	N2SmInfo                    map[uint32][]byte
	PsiIdx                      []byte
	TimerId                     int64
}

func (p *ServiceRequestPrcdCtxt) SetCounter(c int) {
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, nil, "set counter from (%v) to (%v)", p.counter, c)
	p.counter = c
}

func (p *ServiceRequestPrcdCtxt) GetCounter() int {
	return p.counter
}

func (p *ServiceRequestPrcdCtxt) GetIeFlag() *bitset.BitSet {
	return &p.IeFlags
}

// SetState set the procedure state
func (p *ServiceRequestPrcdCtxt) SetNextState(state string) error {
	err := st.SrvReqCheckState(state)
	if err != nil {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "state (%s) is not belong to ServiceRequestPrcd", state)
		return err
	}
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}

func NewServiceRequestPrcdCtxt() (serviceRequestCtxt *ServiceRequestPrcdCtxt) {
	serviceRequestCtxt = &ServiceRequestPrcdCtxt{}
	serviceRequestCtxt.status = st.StateSrvReqStart
	serviceRequestCtxt.N2SmInfo = make(map[uint32][]byte)
	return serviceRequestCtxt
}
