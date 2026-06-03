package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/mobmgnt/procedure"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

func (p *NasMgr) HandleServiceRequestMsg(ctx context.Context, n2connData *gctxt.N2ConnCtxt,
	plainNasMsg *bytes.Reader) error {
	td := []interface{}{ctx, n2connData}
	rlogger.FuncEntry(types.ModuleAmfNas, td)

	p.serviceRequest.Reset()

	//decode the message
	err := p.serviceRequest.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to decode "+
			"service request nas message， %s", err)
		return fmt.Errorf("failed to decode "+
			"service request nas message， %s", err)
	}
	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "p.serviceRequest", p.serviceRequest)

	// get ue context through 5G S-TMSI
	stmsiKey := p.serviceRequest.MobileIdentity.Stmsi5g.GetKey()
	ueCtxt, err := gctxt.GetUeContext(gctxt.StmsiKey(stmsiKey))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "failed to find "+
			"the ue context for Stmsi5g :%s", p.serviceRequest.MobileIdentity.Stmsi5g.String())
		return fmt.Errorf("failed to find "+
			"the ue context for Stmsi5g :%s", p.serviceRequest.MobileIdentity.Stmsi5g.String())
	}

	if p.serviceRequest.IeFlags.Test(nasmsg.IeidServicereqNasmsgcontainer) {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "ueCtxt.IsSecCtxtEstb:", ueCtxt.IsSecCtxtEstb)
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "ueCtxt.CipheringAlg:", ueCtxt.CipheringAlg)

		var msgBuf *bytes.Reader

		//Decipher Message
		if ueCtxt.IsSecCtxtEstb && ueCtxt.CipheringAlg != types3gpp.NEA0 {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "decipher msg:", p.serviceRequest.NasMsgContainer)
			plainMsg, err := DecipherMessage(ueCtxt, p.serviceRequest.NasMsgContainer)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "decipher nas message failed.")
				return fmt.Errorf("decipher nas message failed.")
			}
			p.serviceRequest.NasMsgContainer = plainMsg
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "plain msg:", plainMsg)
		}
		msgBuf = bytes.NewReader(p.serviceRequest.NasMsgContainer)

		// plain msg
		// 1. decode extract the EPD
		epd, err := nas.GetEpd(msgBuf)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to decode nas epd(%s)", epd)
			return fmt.Errorf("failed to decode nas epd")
		}
		//2. decode Security header type and Spare half octet
		secHeaderType, err := nas.GetSecHeaderType(msgBuf)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to decode nas security header type(%)", secHeaderType)
			return fmt.Errorf("failed to decode nas security header type")
		}
		//3.Decode Message Type identity
		msgType, err := msgBuf.ReadByte()
		if err != nil {
			return fmt.Errorf("failed to decode nas msg type(%s), error(%s)", msgType, err)
		}
		//decode the message
		err = p.serviceRequest.Decode(msgBuf)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to decode "+
				"service request nas message， %s", err)
			return fmt.Errorf("failed to decode "+
				"service request nas message， %s", err)
		}
		rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "p.serviceRequest", p)

	}

	SecHdrTypeIncomingMsg, ok := ctx.Value(types.SecHeaderTypeInNasMsgCK).(nas.SecHeaderType)
	if ok {
		ueCtxt.SecHdrTypeIncomingMsg = SecHdrTypeIncomingMsg
	} else {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "No security header type, set to plain nas msg.")
		ueCtxt.SecHdrTypeIncomingMsg = nas.PlainNasMsg
	}

	//set ue context
	ueCtxt.SetRanUeNgapId(n2connData.GnbConnID)

	err = gctxt.UpdateUeContext(gctxt.AmfUeNgApId(n2connData.AmfUeNgapID), ueCtxt)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "failed to update the ue context with AmfUeNgApId (%d)",
			n2connData.AmfUeNgapID)
		return err
	}

	ueCtxt.SetAmfUeNgapID(uint64(n2connData.AmfUeNgapID))

	if len(ueCtxt.GetPDUSessionCtxts()) == 0 {
		//when follow on pending = 1 (in initial UE msg)
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "there is no estb session")
	}

	//allocate service request prcd ctxt
	var srvPrcdCtxt *prcdctxt.ServiceRequestPrcdCtxt

	if ueCtxt.GetProcCtxt() != nil {
		var ok bool
		srvPrcdCtxt, ok = ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
		if ok {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
			return types.ErrFailGetProcedureCtxt
		} else {
			//allocate a new procedure context
			srvPrcdCtxt = prcdctxt.NewServiceRequestPrcdCtxt()
		}
	} else {
		//allocate a new procedure context
		srvPrcdCtxt = prcdctxt.NewServiceRequestPrcdCtxt()
	}

	srvPrcdCtxt.IeFlags = p.serviceRequest.IeFlags
	srvPrcdCtxt.ServiceType = p.serviceRequest.ServiceType
	srvPrcdCtxt.NgKSI = p.serviceRequest.NgKSI
	srvPrcdCtxt.UeSecCapablity = ueCtxt.UeSecCapablity
	srvPrcdCtxt.MobileIdentity = p.serviceRequest.MobileIdentity
	srvPrcdCtxt.AllowedPDUSessionStatus = p.serviceRequest.AllowedPDUSessionStatus
	srvPrcdCtxt.PDUSessionStatus = p.serviceRequest.PDUSessionStatus
	srvPrcdCtxt.UplinkDataStatus = p.serviceRequest.UplinkDataStatus
	srvPrcdCtxt.GnbConnID = n2connData.GnbConnID
	srvPrcdCtxt.GnbInfo = n2connData.GnbInfo
	//srvPrcdCtxt.Tai = n2connData.Tai

	ueCtxt.SetProcCtxt(srvPrcdCtxt)

	// check authentication
	authNeeded := procedure.CheckAuthentication(ueCtxt.SecurityCtxt, srvPrcdCtxt.NgKSI, srvPrcdCtxt.UeSecCapablity)

	// store the ue context in psm context
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "AuthNeed (%v)", authNeeded)
	if authNeeded {
		//trigger the FSM
		err := statemgr.TriggerFsm(ctx,
			statemgr.ServiceRequest,
			statetype.StateSrvReqAuthStart,
			statetype.EventSrvReqServiceReqAuth)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	} else {
		// trigger FSM
		err := statemgr.TriggerFsm(ctx,
			statemgr.ServiceRequest,
			srvPrcdCtxt.GetCurrentState(),
			statetype.EventSrvReqServiceReq)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	}
	return nil
}
