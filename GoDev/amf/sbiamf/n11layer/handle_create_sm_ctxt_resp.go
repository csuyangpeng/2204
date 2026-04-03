package n11layer

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"strings"
)

func HandleCreateSmCtxtResponse(ctx context.Context, createtSmResp *sbicmn.SbiPostCreateSmContext) error {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)

	//get the ueCtxt from ctxt
	var imsi types3gpp.Imsi
	imsistr := strings.TrimPrefix(createtSmResp.Supi, "imsi-")
	imsi.StoreImsiString(imsistr, types3gpp.CheckMncLen(imsistr))

	ueCtxt, err := gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
	if err != nil {
		rlogger.Trace(types.ModuleAmfSbi, rlogger.ERROR, nil, "failed to find the ue context with AmfUeNgApId %s,error %s",
			imsi, err)
		return fmt.Errorf("failed to find the ue context with AmfUeNgApId %s,error %s",
			imsi, err)
	}

	n11CreatedData := sbicmn.Trans_ModelsToN11_SmContextCreatedDataFormat(createtSmResp.RespData)

	if n11CreatedData.PduSessionId != 0 {
		pduSessCtxt := ueCtxt.GetPduSessCtxt(n11CreatedData.PduSessionId)
		if pduSessCtxt != nil {
			rlogger.Trace(types.ModuleAmfSbi, rlogger.DEBUG, nil, "pduSessCtxt.psi", pduSessCtxt.Psi) //todo
		}
	}

	return nil
}
