package sbilayer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sbiamf/n11layer"
	"lite5gc/amf/sc/mobmgnt/mmutils"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/amf/sc/statistics"
	"lite5gc/ausf/manager"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types/sectypes"
	statetype "lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
	"lite5gc/udm/arpf/derivevec"
	"strings"
)

func HandleSBIMsg(ctx context.Context, msg *router.DataMsg) error {
	rlogger.FuncEntry(types.ModuleAmfSc, nil)

	message, ok := msg.MsgData.(*sbicmn.SbiMessage)
	if !ok {
		rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "invalid IPC message - SBI2ScMsg")
		return fmt.Errorf("invalid IPC message - SBI2ScMsg")
	}

	//dispatch message here
	switch message.MsgType {
	case sbicmn.GetAuthDataMsgResponse:
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "received GetAuthDataMsgResponse")
		authData := message.MsgData.(*sbicmn.SbiGetAuthDataMsg)
		//get the ueCtxt from ctxt
		var imsi types3gpp.Imsi
		supiStr := strings.TrimPrefix(authData.Supi, "imsi-")
		err := imsi.StoreImsiString(supiStr, types3gpp.CheckMncLen(supiStr))
		if err != nil {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "failed to get imsi")
			return types.ErrFailGetImsi
		}

		if authData.Data == nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "failed to get AuthData from UDM for imsi(%s)", imsi.String())
			return fmt.Errorf("failed to get AuthData from UDM for imsi(%s)", imsi.String())
		}

		heav := sbicmn.Trans_ModelsToN11_AuthDataFormat(authData.Data)
		ueContext, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
			return fmt.Errorf("failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
		}
		// get sn name
		var snName string
		//get SMF config mgr from ctxt
		snName = mmutils.GenerateSnName(configure.AmfConf.PlmnList.List[0])

		// get ue security context
		ueAvCtxt, err := manager.GetAusfAvContext(imsi.String())
		if err != nil {
			// no ue security context, create a new ue security context
			err, ueAvCtxt = manager.CreateAusfAvContext(imsi.String())
			if err != nil {
				return fmt.Errorf("failed to create ue security context, error(%s_", err)
			}
		}
		//derive ausf av
		err = derivevec.DeriveAusfAv(&heav, snName, ueAvCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, authData.Supi, "failed to derive ausf av, error(%s)", err)
			return fmt.Errorf("failed to derive ausf av, error(%s)", err)
		}
		// store the Auth data from ARPF
		seav := &sectypes.SeAvType{}
		seav.Autn = ueAvCtxt.Autn
		seav.HXresStar = ueAvCtxt.HXresStar
		seav.Rand = ueAvCtxt.Rand

		//store seav
		//ueContext.SecurityCtxt.AuthVector.Rand = seav.Rand
		ueContext.TempSecCtxt.AuthVector.Rand = seav.Rand
		//ueContext.SecurityCtxt.AuthVector.Autn = seav.Autn
		ueContext.TempSecCtxt.AuthVector.Autn = seav.Autn
		//ueContext.SecurityCtxt.AuthVector.HXresStar = seav.HXresStar
		ueContext.TempSecCtxt.AuthVector.HXresStar = seav.HXresStar
		//ueContext.Abba = [2]byte{0x00, 0x00}
		ueContext.TempSecCtxt.Abba = [2]byte{0x00, 0x00}

		//update ngKSI
		//ueContext.SecurityCtxt.NgKsi.Update()
		//reset nas counter 33.501 6.4.5
		//ueContext.SecurityCtxt.UplinkNasCount = 0
		//ueContext.SecurityCtxt.DownlinkNasCount = 0

		//reset derived keys
		//ueContext.KnasIntKey = nil
		ueContext.TempSecCtxt.KnasIntKey = nil
		//ueContext.KnasEncKey = nil
		ueContext.TempSecCtxt.KnasEncKey = nil
		//ueContext.Kamf = nil
		ueContext.TempSecCtxt.Kamf = nil
		//ueContext.Kgnb = nil
		ueContext.TempSecCtxt.Kgnb = nil

		ctx = context.WithValue(ctx, types.UeContextCK, ueContext)
		err = statemgr.TriggerFsm(ctx,
			statemgr.Register,
			statetype.StateRegisterWfUdmGetAuth,
			statetype.EventRegisterUdmGetAuth)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	case sbicmn.GetSmfSelDataMsgResponse:
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "received GetSmfSelectDataMsgResponse")
		smfSelData := message.MsgData.(*sbicmn.SbiGetSmfSelDataMsg)
		//get the ueCtxt from ctxt
		var imsi types3gpp.Imsi
		supiStr := strings.Split(smfSelData.Supi, "-")[1]
		err := imsi.StoreImsiString(supiStr, 2)
		if err != nil {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "failed to get imsi")
			return types.ErrFailGetImsi
		}

		ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
			return fmt.Errorf("failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
		}
		smfS := sbicmn.Trans_ModelsToN11_SmfSelDataFormat(smfSelData.Data)
		ueCtxt.SmfSelSubsData = &smfS
	case sbicmn.GetAmDataMsgResponse:
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "received GetAmDataMsgResponse")
		amData := message.MsgData.(*sbicmn.SbiGetAmDataMsg)
		//get the ueCtxt from ctxt
		var imsi types3gpp.Imsi
		supiStr := strings.Split(amData.Supi, "-")[1]
		err := imsi.StoreImsiString(supiStr, 2)
		if err != nil {
			rlogger.Trace(types.ModuleSmfN11, rlogger.ERROR, nil, "failed to get imsi")
			return types.ErrFailGetImsi
		}

		ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
			return fmt.Errorf("failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
		}
		amsd := sbicmn.Trans_ModelsToN11_AMSDDataFormat(amData.Data)
		ueCtxt.AccMobSubsData = &amsd

		ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)

		err = statemgr.TriggerFsm(ctx,
			statemgr.Register,
			statetype.StateRegisterWfUdmGetAmDataResp,
			statetype.EventUdmGetAmData)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
			return types.ErrFailTriggerFsm
		}
	case sbicmn.PduSessCreateSMContextResp:
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "received PostCreateSmContextResp")
		smResp := message.MsgData.(*sbicmn.SbiPostCreateSmContext)
		err := n11layer.HandleCreateSmCtxtResponse(ctx, smResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "failed to handle PostCreateSmContextResp")
		}
	case sbicmn.N1N2MessageTransferReq:
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "received N1N2MsgTransferReq")
		smData := message.MsgData.(*sbicmn.SbiHandlerMessage)
		err := n11layer.HandleN1N2MsgRequest(ctx, smData)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "failed to handle N1N2MsgTransferReq")
		}
	case sbicmn.PduSessUpdateSMContextResp:
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "received PostModifySmContextResp")
		smResp := message.MsgData.(*sbicmn.SbiPostModifySmContext)
		err := n11layer.HandleUpdateSmCtxtResponse(ctx, smResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "failed to handle PostModifySmContextResp")
		}
	case sbicmn.PduSessReleaseSMContextResp:
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "received PostReleaseSmContextResp")
		smResp := message.MsgData.(*sbicmn.SbiPostReleaseSmContext)
		err := n11layer.HandleReleaseSmCtxtResponse(ctx, smResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfSc, rlogger.ERROR, nil, "failed to handle PostReleaseSmContextResp")
		}
	case sbicmn.NudmFailMsg:
		rlogger.Trace(types.ModuleAmfSc, rlogger.INFO, nil, "received NudmFailMsg")
		sbiFailData := message.MsgData.(*sbicmn.SbiHandleFailMsg)
		var imsi types3gpp.Imsi
		imsistr := strings.TrimPrefix(sbiFailData.Supi, "imsi-")
		imsi.StoreImsiString(imsistr, types3gpp.CheckMncLen(imsistr))
		ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
		if err != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil, "failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
			return fmt.Errorf("failed to find the ue context with AmfUeNgApId %s,error %s",
				imsi, err)
		}
		rejectMsg := nasmsg.RegistrationRejectMsg{}
		rejectMsg.MMCause = nas.IllegalUE
		// encode nas message
		bytes, err := rejectMsg.Encode()
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt,
				"failed to encode registration accept message")
			return fmt.Errorf("failed to encode registration accept message")
		}

		// add the security header
		procNasMsg := mmsender.VerifyAndBuildSecProtectNasMsg(ueCtxt, bytes)

		err, _ = mmsender.SendDownLinkNasMsg(ctx, ueCtxt, procNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"fail to send downlink nas ngap msg")
			return types.ErrFailSendNgapMsg
		}

		//set state for ue
		err = ueCtxt.SetRmState(types.StateRmDeRegistered)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil,
				"fail to set RM state")
			return fmt.Errorf("fail to set RM state")
		}

		//release ue ctxt
		err = mmutils.ReleaseUECtxt(ueCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to release ue context")
			return fmt.Errorf("fail to release ue context")
		}

		//msg counter
		pm.PegCounter(statistics.RegistrationRejectCounter)
	}
	return nil
}
