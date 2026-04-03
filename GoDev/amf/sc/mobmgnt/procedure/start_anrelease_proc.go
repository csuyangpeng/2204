package procedure

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/rlogger"

	"context"
	"fmt"
	"lite5gc/cmn/message/n11msg"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func StartAnRelPrcd(ctxt context.Context, ueContext *gctxt.UeContext, cause types3gpp.CauseValue) error {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	relCause := types3gpp.Nas_normal_release

	//release smf and upf resource
	if len(ueContext.GetPDUSessionCtxts()) != 0 {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueContext.GetImsiPtr(),
			"release pdu session in smf first")
		prcdCtxt := prcdctxt.NewAnReleasePrcdCtxtNetTrigger() //anReleaseCtxt.status = st.StateAnRelWfUpSmCtxtAck
		ueContext.SetProcCtxt(prcdCtxt)
		//release smf and upf resource
		imsi := ueContext.GetImsi()
		for i := range ueContext.GetPDUSessionCtxts() {
			rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueContext.GetImsiPtr(),
				"release pdu session in ue(%s), psi(%d)",imsi.String(), ueContext.GetPDUSessionCtxts()[i].Psi)
			smCtxtUpdateData := n11msg.UpdateSMContextRequestData{}
			smCtxtUpdateData.UpCnxState = n11msg.DEACTIVATED
			smCtxtUpdateData.IeFlags.Set(n11msg.Ieid_upCnxState)
			mmsender.SendUpdateSMCtxtSBIMsg(ctxt, i, imsi, smCtxtUpdateData)
		}
	} else {
		rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueContext.GetImsiPtr(),
			"no pdu session , send rel cmd msg directly")
		prcdCtxt := prcdctxt.NewAnReleasePrcdCtxtAfterReg() //anReleaseCtxt.status = st.StateAnRelWfRelCmp
		ueContext.SetProcCtxt(prcdCtxt)

		//send ue context release command msg to ran
		err := mmsender.SendUeContextReleaseCmd(ctxt, relCause)
		if err != nil {
			rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, ueContext.GetImsiPtr(),
				"fail to send ue context release command msg to RAN,err:", err)
			return fmt.Errorf("fail to send ue context release command msg to ran, err:", err)
		}
	}
	return nil
}
