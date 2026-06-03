package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
)

func (p *NasMgr) HandleUeDeRegistReqMsg(
	ctx context.Context,
	n2connData *gctxt.N2ConnCtxt,
	plainNasMsg *bytes.Reader) error {
	td := []interface{}{ctx, n2connData}
	rlogger.FuncEntry(types.ModuleAmfNas, td)

	p.deRegistrationRequest.Reset()
	//decode the message
	err := p.deRegistrationRequest.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to decode "+
			"de registration request nas message， %s", err)
		return fmt.Errorf("failed to decode nas message")
	}

	// get ue context
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "ueCtxt is not in context")
	}

	switch p.deRegistrationRequest.MobIDGuti.IdType {
	case nasie.Guti5g:
		guti := p.deRegistrationRequest.MobIDGuti.Guti5g
		ueCtxt, err = gctxt.GetUeContext(gctxt.GutiKey(guti.String()))
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to find the ue context with 5G Guti (%s), error(%s)", guti.String(), err)
			return fmt.Errorf("failed to find the ue context with 5G Guti (%s), error(%s)", guti.String(), err)
		}
	case nasie.Suci:
		fallthrough
	case nasie.STmsi5g:
		fallthrough
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td,
			"unsupported mobile id in UE DeRegistration request message")
		return fmt.Errorf("unsupported mobile id in UE DeRegistration request msg(%d)",
			p.deRegistrationRequest.MobIDGuti.IdType)
	}

	// With initial ue msg to tranport DeRegister
	// set n2 context info in ue context
	//ueCtxt.SetRanUeNgapId(n2connData.GnbConnID)
	//err = gctxt.UpdateUeContext(n2connData.AmfUeNgapID, ueCtxt)
	//if err != nil {
	//	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "failed to update the ue context with AmfUeNgApId (%d)",
	//		n2connData.AmfUeNgapID)
	//	return err
	//}
	// update the amf ue ngap id in ue context, should return old amf ue ngap id first
	//err = idmgr64.GetInst().ReturnID(string(types.AMFUeNgapId), ueCtxt.GetAmfUeNgapId())
	//if err != nil {
	//	rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to release AmfUeNgApId:(%d) "+
	//		",error(%s)", ueCtxt.GetAmfUeNgapId(), err)
	//}
	//ueCtxt.SetAmfUeNgapID(uint64(n2connData.AmfUeNgapID))

	// store the ue context in context
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	// check deregistration procedure
	ueDeRegistPrcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.DeRegistrationPrcdCtxt)
	if ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get DeRegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return types.ErrFailGetProcedureCtxt
	} else {
		ueDeRegistPrcdCtxt = prcdctxt.NewDeRegistration()
		ueCtxt.SetProcCtxt(ueDeRegistPrcdCtxt)
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "current states(%s)", ueDeRegistPrcdCtxt.BaseCtxt.GetCurrentState())
	}

	//store the information from income message
	switch p.deRegistrationRequest.MobIDGuti.IdType {
	case nasie.Suci:
		ueDeRegistPrcdCtxt.Suci = *p.deRegistrationRequest.MobIDGuti.Suci
	case nasie.Guti5g:
		ueDeRegistPrcdCtxt.Guti5g = *p.deRegistrationRequest.MobIDGuti.Guti5g
	case nasie.STmsi5g:
		fallthrough
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "unsupported mobile id in ue registration request.")
		return fmt.Errorf("unsupported mobile id in ue registration request msg(%d)", p.deRegistrationRequest.MobIDGuti.IdType)
	}

	ueDeRegistPrcdCtxt.NgKSI = p.deRegistrationRequest.NgKSI
	ueDeRegistPrcdCtxt.DeRegistrationType = p.deRegistrationRequest.DeRegistrationType

	//trigger the FSM
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, td, p.deRegistrationRequest.DeRegistrationType.SwithOff)
	if p.deRegistrationRequest.DeRegistrationType.SwithOff == true {
		err := statemgr.TriggerFsm(ctx,
			statemgr.DeRegister,
			ueCtxt.GetProcCtxt().GetCurrentState(),
			statetype.EventDeRegisterRequestBySwitchOff)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	} else {
		err := statemgr.TriggerFsm(ctx,
			statemgr.DeRegister,
			ueCtxt.GetProcCtxt().GetCurrentState(),
			statetype.EventDeRegisterRequest)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	}

	return nil
}
