package prcdctxt

import (
	"github.com/willf/bitset"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/message/udmdata"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"net"
)

type PduSessionEstbPrcdCtxt struct {
	BaseCtxt
	PduSessId        nas.PduSessID
	IsExistPduSess   bool
	SelectDnn        types3gpp.Apn
	SessMgntSubsData *udmdata.SessMgntSubscripitonData
	DnnConfig        *nasie.DNNConfiguration
	UEIP             net.IP
	Seid             uint64

	// NAS msg of session request
	Pti                  nas.PrcdTransID
	PduSessType          types3gpp.PduSessType
	SscMode              nas.SSCMode
	SmCap                nas.SMCapability
	MaxNumofSupPktFilter uint16
	AlwaysOnPduSessReq   bool
	//Indicates whether an IE is assigned or it is an empty value
	SessionReqIeFlags bitset.BitSet

	//information from create sm context request
	SmfSmCtxtId uint32
	AmfSmCtxtId uint32
	DNN         types3gpp.Apn
	SNSSAI      nasie.SNssai
	IMSI        types3gpp.Imsi
	Supi        types3gpp.Supi
	SelMode     n11msg.DnnSelectionMode
	IsAuthed    bool
	ServingNfId string
	IeFlags     bitset.BitSet

	//info for n1n2 message transfer resp
	Cause       n11msg.N1N2MessageTransferCause
	SupFeatures string

	SessionAmbr nasie.SessionAmbr

	PduState n11msg.UpCnxState

	SbiMessage *sbicmn.SbiHandlerMessage
}

// NewRegistration return a registration procedure context
func NewPduSessEstbPrcdCtxt(psi nas.PduSessID) *PduSessionEstbPrcdCtxt {
	prcdCtxt := &PduSessionEstbPrcdCtxt{}
	prcdCtxt.PduSessId = psi
	prcdCtxt.status = statetype.StatePduSessEstbStart
	prcdCtxt.IsExistPduSess = false
	return prcdCtxt
}

// SetState set the procedure state
func (p *PduSessionEstbPrcdCtxt) SetNextState(state string) error {
	err := statetype.PduSessEstbCheckState(state)
	if err != nil {
		return err
	}
	rlogger.Trace(types.ModuleSmfCtxt, rlogger.DEBUG, nil, "change State from (%s) to (%s)", p.status, state)
	p.status = state
	return nil
}
