/*
* Copyright(C),2020-2022
* Author:  xiaoyun
* Date:    11/16/20
* Description:
	无需鉴权或鉴权完成后，继续处理注册消息，处理至给UDM发出SBI消息
*/
package procedure

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/openapi/models"
)

func HandleRegistrationMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)
	cause := HandleInitialRegistration(ctxt)
	if cause != nas.SuccessAccept {
		//if NAS msg has error, send reject msg
		err := mmsender.SendRegisterRejectMsg(ctxt, cause)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to generate register reject msg")
			return
		}
	}
}

func HandleInitialRegistration(ctxt context.Context) nas.Mm5gCause {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ctxt, "no ue context.")
		return nas.SystemFailure
	}

	// get imsi
	imsi := ueCtxt.GetImsi()

	// get supi
	supi := &types3gpp.Supi{}
	supi.SetImsi(&imsi)

	// get the procedure context
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return nas.SystemFailure
	}

	// DO all the process logic here...
	// 8.AUSFSelection()

	// 9.Authentication()

	// Interaction with UDM
	// 14a Nudm_UECM_Registration
	// 14b Nudm_SDM_Get
	// 14C Nudm_SDM_Subscribe

	scId, ok := ctxt.Value(types.ScIdCK).(uint32)
	if !ok {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "no sc id found")
		return nas.SystemFailure
	}
	rlogger.Trace(types.ModuleAmfMM, rlogger.INFO, ueCtxt.GetImsiPtr(), "sc id: %d", scId)

	// guami
	guami := configure.GetModelsGuami()

	// url
	amfSbi := configure.AmfConf.Sbi.Amf
	RegUrl := fmt.Sprintf("%s://%s:%d/namf-comm/v1/nudm-uecm/imsi-%s;guami=%s",
		amfSbi.Scheme, amfSbi.Addr.Ip, amfSbi.Addr.Port, imsi.String(), configure.GetGuamiStr())

	//14a Nudm_UECM_Registration
	amf3gppAR := models.Amf3GppAccessRegistration{}
	amf3gppAR.AmfInstanceId = configure.AmfConf.Service.AmfInstanceId
	amf3gppAR.DeregCallbackUri = RegUrl
	amf3gppAR.Guami = &guami
	amf3gppAR.RatType = models.RatType_NR
	amf3gppAR.InitialRegistrationInd = true
	//sbi models
	amfRegist := &sbicmn.SbiPostAmf3gppAccessRegistration{}
	amfRegist.Supi = imsi.AddIMSIPrefix()
	amfRegist.Data = &amf3gppAR
	//sc to sbi msg
	amfRegMessage := sbicmn.SbiMessage{}
	amfRegMessage.MsgData = amfRegist
	amfRegMessage.MsgType = sbicmn.PostAmf3gppAccessRegistration
	amfRegMessage.ScInstId = scId
	err := routeragent.SendIpcMessage(ctxt, router.SbiGR, 1, &amfRegMessage)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to send sbi msg to udm. error(%s)", err)
		return nas.SystemFailure
	}
	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, imsi, "send SBI Nudm_UECM_Registration")

	//14b Nudm_SDM_Get
	//sbi models
	getam := &sbicmn.SbiGetAmDataMsg{}
	getam.Supi = imsi.AddIMSIPrefix()
	//sc to sbi msg
	amMessage := sbicmn.SbiMessage{}
	amMessage.MsgData = getam
	amMessage.MsgType = sbicmn.GetAmDataMsgRequest
	amMessage.ScInstId = scId
	err = routeragent.SendIpcMessage(ctxt, router.SbiGR, 1, &amMessage)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to send sbi msg to udm. error(%s)", err)
		return nas.SystemFailure
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, imsi, "send SBI Nudm_SDM_GetAm")

	/*//smf select data
	//sbi models
	getsmfSel := &sbicmn.SbiGetSmfSelDataMsg{}
	getsmfSel.Supi = imsi.AddIMSIPrefix()
	//sc to sbi msg
	smfSelmessage := sbicmn.SbiMessage{}
	smfSelmessage.MsgData = getsmfSel
	smfSelmessage.MsgType = sbicmn.GetSmfSelDataMsgRequest
	smfSelmessage.ScInstId = scId
	err = routeragent.SendIpcMessage(ctxt, router.SbiGR, 1, &smfSelmessage)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to send sbi msg to udm. error(%s)", err)
		return nas.SystemFailure
	}

	// 14C Nudm_SDM_Subscribe
	// url
	monitorUrls := make([]string, 2)
	monitorUrls[0] = fmt.Sprintf("%s://%s:%d/nudm-sdm/v1/imsi-%s/am-data",
		amfSbi.Scheme, amfSbi.Addr.Ip, amfSbi.Addr.Port, imsi.String())
	monitorUrls[1] = fmt.Sprintf("%s://%s:%d/nudm-sdm/v1/imsi-%s/smf-select-data",
		amfSbi.Scheme, amfSbi.Addr.Ip, amfSbi.Addr.Port, imsi.String())
	subsUrl := fmt.Sprintf("%s://%s:%d/namf-comm/v1/nudm-sdm/imsi-%s;guami=%s",
		amfSbi.Scheme, amfSbi.Addr.Ip, amfSbi.Addr.Port, imsi.String(), configure.GetGuamiStr())

	sdmSubsData := models.SdmSubscription{}
	sdmSubsData.NfInstanceId = configure.AmfConf.Service.AmfInstanceId
	sdmSubsData.ImplicitUnsubscribe = true
	sdmSubsData.CallbackReference = subsUrl
	sdmSubsData.MonitoredResourceUris = monitorUrls
	//sbi models
	var sdmSub = &sbicmn.SbiPostSdmSubscription{}
	sdmSub.Supi = imsi.AddIMSIPrefix()
	sdmSub.Data = &sdmSubsData
	//sc to sbi msg
	sdmSubsMessage := sbicmn.SbiMessage{}
	sdmSubsMessage.MsgData = sdmSub
	sdmSubsMessage.MsgType = sbicmn.PostSdmSubscription
	sdmSubsMessage.ScInstId = scId
	err = routeragent.SendIpcMessage(ctxt, router.SbiGR, 1, &sdmSubsMessage)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to send sbi msg to udm. error(%s)", err)
		return nas.SystemFailure
	}*/

	// 15-16.PCFStart()
	// 17.PDUSessionUpdate()

	err = prcdCtxt.SetNextState(statetype.StateRegisterWfUdmGetAmDataResp)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to set state")
		return nas.SystemFailure
	}

	return nas.SuccessAccept
}
