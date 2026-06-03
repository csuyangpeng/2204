package procedure

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/smfcontext/gctxt"
	"lite5gc/smf/smfcontext/prcdctxt"
)

func PduSessEstbFinished(pCtxt *prcdctxt.PduSessionEstbPrcdCtxt) {
	rlogger.FuncEntry(types.ModuleSmfSM, pCtxt)
	if pCtxt == nil  {
		rlogger.Trace(types.ModuleSmfNas, rlogger.ERROR, nil, "input para is nil")
		return
	}
	//smf session context
	pduSessCtxt, err := gctxt.GetSessContext(gctxt.SeidKey(pCtxt.Seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, pCtxt, "failed to get pdu session context with seid(%d)", pCtxt.Seid)
		return
	}

	pduSessCtxt.SEID = pCtxt.Seid
	pduSessCtxt.AmfID = pCtxt.ServingNfId
	pduSessCtxt.DNN = pCtxt.DNN
	pduSessCtxt.SNSSAI = pCtxt.SNSSAI
	pduSessCtxt.UEIP = pCtxt.UEIP
	pduSessCtxt.IsAuthed = pCtxt.IsAuthed
	pduSessCtxt.IMSI = pCtxt.IMSI
	pduSessCtxt.DnnConfig = pCtxt.DnnConfig
	pduSessCtxt.SessionAmbr = pCtxt.SessionAmbr
	pduSessCtxt.SetPrcdCtxt(nil)
}
