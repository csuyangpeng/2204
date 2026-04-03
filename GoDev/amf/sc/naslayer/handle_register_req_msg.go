/*
* Copyright(C),2020-2022
* Author:  xiaoyun
* Date:    11/13/20 3:29 AM
* Description:
	处理所有注册请求消息：
	1 解码注册请求的NAS消息
	2 创建UE上下文，填充信息
	3 不同的注册请求消息的分发处理
*/
package naslayer

import (
	"bytes"
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	statetype "lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/udm/sidf"
)

func (p *NasMgr) HandleRegistrationRequestMsg(ctx context.Context, n2connData *gctxt.N2ConnCtxt,
	plainNasMsg *bytes.Reader) nas.Mm5gCause {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	// 1 decode the nas message
	p.registrationRequest.Reset()
	err := p.registrationRequest.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to decode "+
			"registration request nas message, %s", err)
		return nas.InvalidMandInfo
	}
	rrMsg := &p.registrationRequest

	// 2 get the ue context with imsi/guti
	ueContext, rc := RetreiveUeContext(rrMsg, n2connData)
	if ueContext == nil {
		return nas.SystemFailure
	}

	switch rc {
	case nas.SuccessAccept:
	case nas.IdRequestNeed:
		isSucc := TriggerIdRequestPrcd(ctx, ueContext, rrMsg, n2connData.NeedInitCtxtSetupPrcd)
		if isSucc {
			return nas.SuccessAccept
		} else {
			return nas.IllegalUE
		}
	default:
		// illegal ue, return here
		return rc
	}

	// 3 save information in ue context
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil,
		"save in ue context ran_ue_ngap_id(%x)", n2connData.GnbConnID)
	ueContext.SetRanUeNgapId(n2connData.GnbConnID)

	// for test
	amf_id := ueContext.GetAmfUeNgapId()
	testCtxt, err := gctxt.GetUeContext(gctxt.AmfUeNgApId(amf_id))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to find the ue context with AmfUeNgApId (%x), AmfUeIdUeCtxtTable(%s)",
			amf_id,
			gctxt.DumpAmfUeIdUeCtxtTable())
	} else {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"amf ue ngap id (%x), ue_ngap_id(%x)", testCtxt.GetAmfUeNgapId(), testCtxt.GetRanUeNgapId())
	}

	// security Capability
	ueContext.SecurityCtxt.UeSecCapablity = rrMsg.UeSecCapablity
	// security header
	SecHdrTypeIncomingMsg, ok := ctx.Value(types.SecHeaderTypeInNasMsgCK).(nas.SecHeaderType)
	if ok {
		ueContext.SecHdrTypeIncomingMsg = SecHdrTypeIncomingMsg
	} else {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "No security header type, set to plain nas msg")
		ueContext.SecHdrTypeIncomingMsg = nas.PlainNasMsg
	}

	// 存进context，便于跳转状态机后再次使用
	ctx = context.WithValue(ctx, types.UeContextCK, ueContext)
	// 4 handle the registration message
	switch rrMsg.RegType {
	case nas.InitRegist:
		cause := p.handleInitRegisterRequest(ctx, n2connData, ueContext)
		if cause != nas.SuccessAccept {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueContext.GetImsiPtr(),
				"fail hand init reg req msg")
			return cause
		}
	case nas.MobRegistUpdating, nas.PeriodicRegUpdating, nas.EmergencyRegist:
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueContext.GetImsiPtr(),
			"unsupported reg type (%d)", rrMsg.RegType)
		return nas.MsgTypeNotCompatble
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ueContext.GetImsiPtr(),
			"unknown reg type (%d)", rrMsg.RegType)
		return nas.MsgTypeNotCompatble
	}

	return nas.SuccessAccept
}

func RetreiveUeContext(rrMsg *nasmsg.RegistrationRequestMsg, n2connData *gctxt.N2ConnCtxt) (*gctxt.UeContext, nas.Mm5gCause) {

	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	// create a temp ue context
	ueContext := gctxt.NewUeContextByAmfUeNgapID(types3gpp.UeNgApID(n2connData.AmfUeNgapID))
	// save the amf_ue_ngap_id index with ue context
	err := gctxt.AddIndexUeContext(gctxt.AmfUeNgApId(ueContext.GetAmfUeNgapId()), ueContext)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueContext.GetImsiPtr(),
			"failed to add amf ue ngap id for ue context table")
		return ueContext, nas.SystemFailure
	}

	idType := rrMsg.MobileIdentity.IdType
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "rrMsg.MobileIdentity.IdType :%s", idType)

	switch idType {
	case nasie.Guti5g:
		guti := rrMsg.MobileIdentity.Guti5g
		ueCtxtWithGuti, err := gctxt.GetUeContext(gctxt.GutiKey(guti.String()))
		if err != nil {
			// 该GUTI不是本核心网分配的，需启动identity request流程，获取SUCI
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil,
				"failed to find the ue context with 5g guti (%s), error(%s)",
				guti.String(),
				err)
			ueContext.Guti5g = guti
			return ueContext, nas.IdRequestNeed
		} else {
			// update the ue context stored in Globel Context table
			ueContext = ueCtxtWithGuti
			ueContext.SetAmfUeNgapID(uint64(n2connData.AmfUeNgapID))
			gctxt.UpdateUeContext(n2connData.AmfUeNgapID, ueContext)
		}
	case nasie.Suci:
		// decode the SUCI，if SUCI type is ciphered, decode the suci firstly
		imsi, rt := GetImsiFromMsg(rrMsg)
		if rt != nas.SuccessAccept {
			return ueContext, rt
		}

		switch rrMsg.RegType {
		case nas.InitRegist:
			temp, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
			if err != nil {
				// save the imsi in ue context
				ueContext.SetImsi(imsi)
				err = gctxt.AddIndexUeContext(gctxt.ImsiKey(imsi.GetValue()), ueContext)
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to add index(imsi:%s) "+
						"for ue context,error(%s)", imsi, err)
					return ueContext, nas.SystemFailure
				}
				rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, gctxt.DumpAmfUeIdUeCtxtTable())
			} else {
				// update the ueContext with stored UeContext
				ueContext = temp
				ueContext.SetAmfUeNgapID(uint64(n2connData.AmfUeNgapID))
				gctxt.UpdateUeContext(n2connData.AmfUeNgapID, ueContext)
				rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, gctxt.DumpAmfUeIdUeCtxtTable())
			}
		case nas.MobRegistUpdating, nas.PeriodicRegUpdating, nas.EmergencyRegist:
			fallthrough
		default:
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "unsupported register type")
			return ueContext, nas.MsgTypeNotCompatble
		}
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ueContext.GetImsiPtr(), "UE ID should be guti or suci")
		return ueContext, nas.IllegalUE
	}

	return ueContext, nas.SuccessAccept
}

func GetImsiFromMsg(rrMsg *nasmsg.RegistrationRequestMsg) (*types3gpp.Imsi, nas.Mm5gCause) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	var imsi *types3gpp.Imsi
	var err error

	switch rrMsg.MobileIdentity.Suci.GetProtectSchemeId() {
	case types3gpp.NullScheme:
		imsi, err = rrMsg.MobileIdentity.Suci.GetImsi()
		if err != nil {
			return nil, nas.IllegalUE
		}
	case types3gpp.ProfileA:
		output := rrMsg.MobileIdentity.Suci.GetSchemeOutputA()
		HomeNwPubKeyId := rrMsg.MobileIdentity.Suci.GetHomeNwPubKeyId()
		MsinHex, err := sidf.SuciDecryptA(output.EphPublicKey, output.Ciphertext, output.MacTag, HomeNwPubKeyId)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to de-conceal SUCI ", err)
			return nil, nas.IllegalUE
		}
		rrMsg.MobileIdentity.Suci.SetMsinHex2SchemeOutput(MsinHex)

		imsi, err = rrMsg.MobileIdentity.Suci.GetImsi()
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to get imsi from suci ", err)
			return nil, nas.IllegalUE
		}
	case types3gpp.ProfileB:
		fallthrough
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "unsupported protection scheme")
		return nil, nas.UeIdCannotDerivedbyNw
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "IMSI in SUCI:%s ", imsi.String())
	return imsi, nas.SuccessAccept
}

func TriggerIdRequestPrcd(ctx context.Context,
	ueContext *gctxt.UeContext,
	rrMsg *nasmsg.RegistrationRequestMsg,
	NeedInitCtxtSetupPrcd bool) bool {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	// 创建流程上下文，registration procedure context
	if ueContext.GetProcCtxt() != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, nil, "current procedure is already exist")
		// 24.501 5.5.1.2.8 e)
		// check the ies in registration message,
		// 1) if different,
		// the previously initiated the registration procedure for initial registration
		// shall be aborted and the new the registration procedure
		//for initial registration shall be executed
		// 2) else, ignore the second message
		// Currently, cancel the current procedure and process initial registration request
		ueContext.SetProcCtxt(nil)
	}

	//allocate the procedure context
	ueRegistPrcdCtxt := *prcdctxt.NewRegistrationPcrdNoIMSI()
	ueRegistPrcdCtxt.FillCtxtFromMsg(rrMsg)
	ueRegistPrcdCtxt.NeedInitCtxtSetupPrcd = NeedInitCtxtSetupPrcd
	ueContext.SetProcCtxt(&ueRegistPrcdCtxt)

	err := statemgr.TriggerFsm(ctx,
		statemgr.Register,
		statetype.StateRegisterStart,
		statetype.EventRegisterIdentityRequest)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
		return false
	}

	return true
}
