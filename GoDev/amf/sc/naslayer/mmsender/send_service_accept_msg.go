package mmsender

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/nassecurity"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/oam/pm"
)

func SendServiceAccept(ctxt context.Context) error {
	rlogger.FuncEntry(types.ModuleAmfNas,ctxt)

	//get the ueCtxt from ctxt
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ctxt,
			"no ue context")
		return types.ErrFailFindUeCtxt
	}

	// get procedure context
	pCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.ServiceRequestPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ServiceRequestPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return types.ErrFailGetProcedureCtxt
	}

	serviceAccept := nasmsg.ServiceAcceptMsg{}
	serviceAccept.PDUSessionStatus = ueCtxt.GetUePsiStatus()
	serviceAccept.IeFlags.Set(nasmsg.IeidServiceacptPdusessionstatus)
	serviceAcceptNasBytes, err := serviceAccept.Encode()
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ctxt, "fail to encode serviceAccept")
		return fmt.Errorf("fail to encode serviceAccept")
	}

	// add the security header
	procNasMsg := VerifyAndBuildSecProtectNasMsg(ueCtxt, serviceAcceptNasBytes)
	if procNasMsg == nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ctxt, "protectNasMsg is nil!!")
	}

	//msg counter
	pm.PegCounter(statistics.ServiceAcceptCounter)

	//  get the scnglayer
	sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return types.ErrFailFindNgapLayer
	}

	var setupReqList []*types3gpp.PduSessResSetupReqItem
	for k, v := range pCtxt.N2SmInfo {
		// get pdu session
		pduSessCtxt := ueCtxt.GetPduSessCtxt(nas.PduSessID(k))

		//construct snssai
		snssai := &types3gpp.Snssai{}
		switch pduSessCtxt.SNssai.Ind {
		case nasie.SstOnly, nasie.SstMapSst:
			snssai.Sst = pduSessCtxt.SNssai.Sst
		case nasie.SstSd, nasie.SstSdMapSst, nasie.SstSdMapSstMapSd:
			snssai.Sst = pduSessCtxt.SNssai.Sst
			snssai.Sd = types3gpp.ConvertSdToU32(pduSessCtxt.SNssai.Sd[:],types.BigEndian)
		}
		item := &types3gpp.PduSessResSetupReqItem{}
		item.PduSessResSetupReqTrans = string(v)
		item.Snssai = snssai
		item.PduSessionId = uint8(pduSessCtxt.Psi)
		setupReqList = append(setupReqList, item)
	}
	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, ctxt, "setupReqList:", setupReqList)

	// send the message in scng layer
	err = sender.SendInitialContextSetupRequest(ueCtxt, pCtxt.GnbInfo.GnbInstId, procNasMsg, setupReqList)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ctxt, "fail to SendInitialContextSetupRequest")
		return fmt.Errorf("fail to SendInitialContextSetupRequest")
	}

	nassecurity.UpdateDownlinkNasCounter(ueCtxt)

	return nil
}
