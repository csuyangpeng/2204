package mmsender

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/convsnssai"
	"lite5gc/oam/pm"
)

func SendRegistrationAccept(ctxt context.Context, ueCtxt *gctxt.UeContext) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	registerAcceptMsg := nasmsg.RegistrationAcceptMsg{}
	registerAcceptMsg.IeiMark = make(map[nasie.Iei]bool)

	registerAcceptMsg.RegResult = nasmsg.Access3gpp
	registerAcceptMsg.AllowSmsOverNAS = ueCtxt.AllowSmsOverNAS

	registerAcceptMsg.Guti5g.Guti5g = ueCtxt.Guti5g
	registerAcceptMsg.Guti5g.IdType = nasie.Guti5g
	registerAcceptMsg.IeiMark[nasie.IeiGuti5G] = true

	registerAcceptMsg.EquivalentPlmns = ueCtxt.EquivalentPlmns
	registerAcceptMsg.IeiMark[nasie.IeiEquivalentPLMNs] = true

	for _, v := range ueCtxt.AllowedNssai {
		tmp := convsnssai.ConvertSnssai(v)
		registerAcceptMsg.AllowedNssai.AddSNssai(tmp)
	}

	registerAcceptMsg.IeiMark[nasie.IeiAllowedNSSAI] = true

	registerAcceptMsg.ServiceAreaList = ueCtxt.ServiceAreaList
	registerAcceptMsg.IeiMark[nasie.IeiServiceAreaList] = true

	registerAcceptMsg.TaiList = ueCtxt.TaiList
	registerAcceptMsg.IeiMark[nasie.IeiTAIList] = true

	registerAcceptMsg.T3502 = ueCtxt.T3502
	registerAcceptMsg.IeiMark[nasie.IeiT3502Value] = true

	// encode nas message
	bytes, err := registerAcceptMsg.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to encode "+
			"registration accept message")
		return err
	}

	// add the security header
	procNasMsg := VerifyAndBuildSecProtectNasMsg(ueCtxt, bytes)

	err, encNgapMsg := SendDownLinkNasMsg(ctxt, ueCtxt, procNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to send downlink nas ngap msg")
		return types.ErrFailSendNgapMsg
	}

	// update downlink nas counter
	nassecurity.UpdateDownlinkNasCounter(ueCtxt)

	//msg counter
	pm.PegCounter(statistics.RegistrationAcceptCounter)

	// start t3550 timer for register accept message
	err = StartNasTimer(ctxt, ueCtxt, gctxt.T3550, encNgapMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to start T3550,error:", err)
		return fmt.Errorf("failed to start t3550 nas timer, error(%s)", err)
	}

	return nil
}
