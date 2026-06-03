package gctxt

import (
	"fmt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/cmn/message/udmdata"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	T "lite5gc/cmn/types3gpp"
)

//UeContext is struct for ue information
type UeContext struct {
	imsi T.Imsi

	registMngtState types.RMState
	connMngtState   types.CMState

	//procedure ctxt interface type
	procCtxt prcdctxt.Base

	// pdu session context
	pduSessCtxt map[nas.PduSessID]*AmfPduSessCtxt

	amfUeNgapID T.UeNgApID
	ranUeNgapID T.UeNgApID

	Guti5g          *T.Guti5G
	AllowSmsOverNAS bool //false - not Allow
	EquivalentPlmns T.PlmnList
	TaiList         T.TAIList
	AllowedNssai    []T.Snssai
	RejectNssai     []T.Snssai
	PduSessStatus   nasie.SessionStatus
	ServiceAreaList nasie.ServiceAreaList
	T3512           nasie.GprsTimer3
	T3502           nasie.GprsTimer2

	// subscription data
	AccMobSubsData  *udmdata.AccMobSubscribeData
	SmfSelSubsData  *udmdata.SmfSelSubscribeData
	UeCtxtInSmfData *udmdata.UeContextInSmfData

	// ue radio capability info
	UeRadioCapInfo string

	// security context
	SecurityCtxt

	// procedure timer container
	PrcdTimerCtxt map[TimerType]*TimerCtxt

	// the things that 5g core need to do after service request
	TodoThing                nas.MmMsgType
	IsReRegistrationRequired bool
}

// NewUeContext return a UeContext with imsi
func NewUeContext(imsi T.Imsi) *UeContext {
	return &UeContext{imsi: imsi,
		pduSessCtxt:     make(map[nas.PduSessID]*AmfPduSessCtxt, nas.MaxPSI),
		registMngtState: types.StateRmDeRegistered,
		connMngtState:   types.CmIdle,
		PrcdTimerCtxt:   make(map[TimerType]*TimerCtxt, 256)}
}

func NewUeContextByAmfUeNgapID(amfUeID T.UeNgApID) *UeContext {
	return &UeContext{amfUeNgapID: amfUeID,
		pduSessCtxt:     make(map[nas.PduSessID]*AmfPduSessCtxt, nas.MaxPSI),
		registMngtState: types.StateRmDeRegistered,
		connMngtState:   types.CmIdle,
		PrcdTimerCtxt:   make(map[TimerType]*TimerCtxt, 256)}
}

func (p *UeContext) SetPduSessCtxtNil() {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	p.pduSessCtxt = make(map[nas.PduSessID]*AmfPduSessCtxt)
}

func (p *UeContext) SetImsi(i *T.Imsi) {
	p.imsi = *i
}

func (p *UeContext) GetImsi() T.Imsi {
	return p.imsi
}

func (p *UeContext) GetImsiPtr() *T.Imsi {
	return &p.imsi
}
func (p *UeContext) GetSupi() *T.Supi {
	supi := &T.Supi{}
	supi.SetImsi(&p.imsi)
	err := supi.SetType(T.IMSIType)
	if err != nil {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.INFO, p, "wrong supi type")
		return nil
	}
	return supi
}

// GetRmState return the RM State
func (p *UeContext) GetRmState() types.RMState {
	return p.registMngtState
}

// SetRmState set the RM State
func (p *UeContext) SetRmState(state types.RMState) error {
	err := types.RmCheckState(state)
	if err != nil {
		return err
	}
	p.registMngtState = state
	return nil
}

// GetCmState return the CM State
func (p *UeContext) GetCmState() types.CMState {
	return p.connMngtState
}
func (p *UeContext) GetPDUSessionCtxts() map[nas.PduSessID]*AmfPduSessCtxt {
	return p.pduSessCtxt
}

// SetCmState set the CM State
func (p *UeContext) SetCmState(state types.CMState) error {
	err := types.CmCheckState(state)
	if err != nil {
		return err
	}
	p.connMngtState = state
	return nil
}

// GetProcCtxt return procedure context
func (p *UeContext) GetProcCtxt() prcdctxt.Base {
	return p.procCtxt
}

// SetProcCtxt set the procedure context
func (p *UeContext) SetProcCtxt(ctxt prcdctxt.Base) {
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.INFO, p, "set procedure context of UE to:%T", ctxt)
	p.procCtxt = ctxt
}

func (p *UeContext) SetAmfUeNgapID(id uint64) {
	p.amfUeNgapID = T.UeNgApID(id)
}

func (p *UeContext) GetAmfUeNgapId() uint64 {
	return uint64(p.amfUeNgapID)
}

func (p *UeContext) SetRanUeNgapId(id uint32) {
	p.ranUeNgapID = T.UeNgApID(id)
}

func (p *UeContext) GetRanUeNgapId() uint32 {
	return uint32(p.ranUeNgapID)
}

func (p *UeContext) GetPduSessCtxt(psi nas.PduSessID) *AmfPduSessCtxt {
	sctxt, ok := p.pduSessCtxt[psi]
	if !ok {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "pdu session id(%v) is not exist", psi)
		return nil
	}
	return sctxt
}

func (p *UeContext) SetPduSessCtxt(psi nas.PduSessID, pt *AmfPduSessCtxt) {
	p.pduSessCtxt[psi] = pt
}

func (p *UeContext) AddPduSessCtxt(psi nas.PduSessID, pctxt *AmfPduSessCtxt) error {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	if _, ok := p.pduSessCtxt[psi]; ok {
		return fmt.Errorf("pdu session id(%d) is duplicated", int(psi))
	}
	p.pduSessCtxt[psi] = pctxt
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "psi(%d) is add to ueCxt", psi)
	return nil
}

func (p *UeContext) DelPduSessCtxt(psi nas.PduSessID) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	delete(p.pduSessCtxt, psi)
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "delete pdu session context,psi(%d)", psi)
}

func (p *UeContext) IsPduSessExist(psi nas.PduSessID) bool {
	_, ok := p.pduSessCtxt[psi]
	if !ok {
		return false
	}
	return true
}

func (p *UeContext) GetPsiList(status T.PduSessStatus) []byte {
	var rtPsi []byte
	for k, v := range p.pduSessCtxt {
		if v.Status == status {
			rtPsi = append(rtPsi, byte(k))
		}
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "PSI(%d),Status(%s)", k, v.Status)
	}
	return rtPsi
}

func (p *UeContext) GetUePsiStatus() nasie.SessionStatus {
	p.PduSessStatus.PsiA = 0
	p.PduSessStatus.PsiB = 0
	for _, v := range p.pduSessCtxt {
		if v.Status == T.SessActived {
			switch v.Psi {
			//psiA.bit1: no use
			case 1:
				p.PduSessStatus.PsiA |= 0x02 //psiA.bit2 == psi 1
			case 2:
				p.PduSessStatus.PsiA |= 0x04
			case 3:
				p.PduSessStatus.PsiA |= 0x08
			case 4:
				p.PduSessStatus.PsiA |= 0x10
			case 5:
				p.PduSessStatus.PsiA |= 0x20
			case 6:
				p.PduSessStatus.PsiA |= 0x40
			case 7:
				p.PduSessStatus.PsiA |= 0x80
			case 8:
				p.PduSessStatus.PsiB |= 0x01
			case 9:
				p.PduSessStatus.PsiB |= 0x02
			case 10:
				p.PduSessStatus.PsiB |= 0x04
			case 11:
				p.PduSessStatus.PsiB |= 0x08
			case 12:
				p.PduSessStatus.PsiB |= 0x10
			case 13:
				p.PduSessStatus.PsiB |= 0x20
			case 14:
				p.PduSessStatus.PsiB |= 0x40
			case 15:
				p.PduSessStatus.PsiB |= 0x80 //psiB.bit8 == psi 15
			}
		}
	}
	rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, p, "p.PduSessStatus(%v)", p.PduSessStatus)
	return p.PduSessStatus
}

func (p *UeContext) UpdateSecCtxt() {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)

	p.Kamf = p.TempSecCtxt.Kamf
	p.Abba = p.TempSecCtxt.Abba
	p.Kgnb = p.TempSecCtxt.Kgnb
	p.UplinkNasCount = p.TempSecCtxt.UplinkNasCount
	p.DownlinkNasCount = p.TempSecCtxt.DownlinkNasCount
	p.CipheringAlg = p.TempSecCtxt.CipheringAlg
	p.IntegrityAlg = p.TempSecCtxt.IntegrityAlg
	p.AuthVector = p.TempSecCtxt.AuthVector
	p.KnasEncKey = p.TempSecCtxt.KnasEncKey
	p.KnasIntKey = p.TempSecCtxt.KnasIntKey

	p.TempSecCtxt.Reset()
}
