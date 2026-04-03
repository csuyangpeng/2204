package naslayer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

func (p *NasMgr) HandleUeDeRegistAcceptMsg(
	ctx context.Context,
	n2connData *gctxt.N2ConnCtxt) error {
	td := []interface{}{ctx, n2connData}
	rlogger.FuncEntry(types.ModuleAmfNas, td)

	//get ue context with amf ngap id
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "there is no ue context")
		return types.ErrFailFindUeCtxt
	}
	//release smf and upf resource
	imsi := ueCtxt.GetImsi()
	for i, v := range ueCtxt.GetPDUSessionCtxts() {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "release pdu session in ue(%s), psi(%d)",
			imsi.String(), ueCtxt.GetPDUSessionCtxts()[i].Psi)
		releaseSmCtxtReq := n11msg.ReleaseSMContextRequestData{}
		releaseSmCtxtReq.NgApCause.Type = types3gpp.CT_RadioNetwork
		releaseSmCtxtReq.NgApCause.Value = types3gpp.Radiok_release_due_to_5gc_generated_reason
		mmsender.SendReleaseSMCtxtSBIMsg(ctx, v.Psi, ueCtxt.GetImsi(), releaseSmCtxtReq)
	}
	//cancel T3522
	err := gctxt.CancelPrcdTimer(ueCtxt, ctx, gctxt.T3522)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to cancel T3522")
		return fmt.Errorf("failed to cancel T3522")
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "cancel t3522 timer for deRegister request msg")

	//  get the scnglayer
	sender, ok := ctx.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return types.ErrFailFindNgapLayer
	}
	// send the message in scng layer
	err = sender.SendUEContextReleaseCommand(ueCtxt, n2connData.GnbInfo.GnbInstId, types3gpp.Nas_deregister)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, ctx, "failed to send UEContext Release Command message,err :", err)
		return err
	}

	uedeRegistPrcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.DeRegistrationPrcdCtxt)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get DeRegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
		return types.ErrFailGetProcedureCtxt
	}
	err = uedeRegistPrcdCtxt.SetNextState(statetype.StateDeRegisterWfUeCtxtRelCmp)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, td, "fail to set state")
		return fmt.Errorf("fail to set state")
	}

	return nil
}
