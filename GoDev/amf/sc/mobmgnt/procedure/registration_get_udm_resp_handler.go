/*
* Copyright(C),2020-2022
* Author:  xiaoyun
* Date:    11/16/20
* Description:
	收到UDM的回复消息后，处理信息，发送注册接受消息或ICSR消息
*/
package procedure

import (
	"context"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/mobmgnt/mmutils"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types/convsnssai"
	statetype "lite5gc/cmn/types/statetypes"
)

func HandleRegistrationUdmGetAmDataMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)
	cause := processUdmRespData(ctxt)
	if cause != nas.SuccessAccept {
		err := mmsender.SendRegisterRejectMsg(ctxt, cause)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ctxt, "failed to generate register reject msg")
			return
		}
	}
}

func processUdmRespData(ctxt context.Context) nas.Mm5gCause {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	//get ue context
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ctxt, "no ue context.")
		return nas.SystemFailure
	}

	//get register procedure context
	prcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return nas.SystemFailure
	}

	//compare the plmn in UE's imsi and config file
	imsi := ueCtxt.GetImsi()
	imsiPlmnStr := imsi.GetPlmnStr()
	confPlmnStr := configure.AmfConf.PlmnList.List[0].String()
	if imsiPlmnStr != confPlmnStr {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, imsi,
			"Plmn Not Allowed, from IMSI(%s), from Config(%s)", imsiPlmnStr, confPlmnStr)
		return nas.PlmnNotAllowed
	}

	// reject NSSAI
	//requestNssai := prcdCtxt.RequestNssai
	//configNssai := configure.AmfConf.Nssai
	//_, tag := nasie.HasUnionPart(requestNssai, configNssai)
	//switch tag {
	//case nasie.Total_Identical:
	//case nasie.Different_And_No_Union, nasie.Different_But_Has_Union:
	//	ueCtxt.RejectNssai = requestNssai
	//}

	// allow NSSAI
	var dftSubsNssai nasie.Nssai
	dftSubsNssai.AddSNssai(ueCtxt.AccMobSubsData.Nssai.DefSnssai)
	if len(prcdCtxt.RequestNssai.Snssais) == 0 {
		if len(dftSubsNssai.Snssais) != 0 {
			// use the nssai in subscription data
			ueCtxt.AllowedNssai = ueCtxt.AllowedNssai[0:0]
			for _, v := range dftSubsNssai.Snssais {
				tmp := convsnssai.Convert3gtpSnssai(v)
				ueCtxt.AllowedNssai = append(ueCtxt.AllowedNssai, tmp)
			}
		} else {
			// us the nssai in configuration data
			ueCtxt.AllowedNssai = configure.AmfConf.Nssai
		}
	} else {
		// nssai exist in request msg
		SnssaiSet := nasie.HasIntersectionSet(dftSubsNssai, prcdCtxt.RequestNssai)
		if len(SnssaiSet) > 0 {
			ueCtxt.AllowedNssai = ueCtxt.AllowedNssai[0:0]
			for _, v := range SnssaiSet {
				tmp := convsnssai.Convert3gtpSnssai(v)
				ueCtxt.AllowedNssai = append(ueCtxt.AllowedNssai, tmp)
			}
		} else {
			ueCtxt.AllowedNssai = configure.AmfConf.Nssai
		}
	}

	// TA list
	// todo get tai from ta table and get the the ta list
	ueCtxt.TaiList = configure.AmfConf.TaiLists[0]

	// ServiceAreaList
	at := nasie.ServiceAreaList{}
	at.SerAreaListType = nasie.AllowArea
	at.AreaList = configure.AmfConf.TaiLists[0]
	ueCtxt.ServiceAreaList = at

	// timer
	ueCtxt.T3502 = configure.GetAmfNasT3502Timer()
	ueCtxt.T3512 = configure.GetAmfNasT3512Timer()

	// guti
	if ueCtxt.Guti5g != nil {
		// recycling old guti resource
		cause := mmutils.RecyclingGUTIResources(ueCtxt)
		if cause != nas.SuccessAccept {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "recycling old guti resource, error(%s)", cause)
			return cause
		}
	}
	// allocate new guti resource
	cause := mmutils.AllocatingGUTIResources(ueCtxt)
	if cause != nas.SuccessAccept {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "allocate new guti resource, error(%s)", cause)
		return cause
	}

	prcdCtxt.NeedInitCtxtSetupPrcd = true
	if prcdCtxt.NeedInitCtxtSetupPrcd {
		n2ctxt, err := gctxt.GetN2ConnContext(gctxt.AmfUeNgApId(ueCtxt.GetAmfUeNgapId()))
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to get n2 context, id(%d), err(%s)",
				ueCtxt.GetAmfUeNgapId(), err)
			return nas.SystemFailure
		}
		//send initial context setup request msg
		sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
		if !ok {
			rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get ngap layer .")
			return nas.SystemFailure
		}
		err = sender.SendInitialContextSetupRequest(ueCtxt, n2ctxt.GnbInfo.GnbInstId, nil, nil)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to send ICSR msg, err:", err)
			return nas.SystemFailure
		}
		//reset the n2ctxt flag
		n2ctxt.NeedInitCtxtSetupPrcd = false
		//set the next status
		err = prcdCtxt.SetNextState(statetype.StateRegisterWfInitCtxtSetupResp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to set state")
			return nas.SystemFailure
		}
	} else {
		//send register accept msg
		err := mmsender.SendRegistrationAccept(ctxt, ueCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to send register accept msg, err:", err)
			return nas.SystemFailure
		}
		//set the next state
		err = prcdCtxt.SetNextState(statetype.StateRegisterWfRegisterCmp)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, imsi, "failed to set state")
			return nas.SystemFailure
		}
	}
	return nas.SuccessAccept
}
