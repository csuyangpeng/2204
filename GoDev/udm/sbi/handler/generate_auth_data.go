package udmhdl

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/openapi/http_wrapper"
	"lite5gc/udm/datarepository/service"
	"net/http"
	"strings"
)

func HandleGenerateAuthData(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiAUT, nil)
	//msg counter
	//pm.PegCounter(statistics.AuthenticationInforRetrivslRequestsCounter)
	supiOrSuci := req.Params["supiOrSuci"]
	supiOrSuci = strings.TrimPrefix(supiOrSuci, "imsi-")

	authData, problemDetails := service.GetAuthData(supiOrSuci)

	if problemDetails != nil {
		//msg counter
		//pm.PegCounter(statistics.AuthenticationInforRetrivslFailureCounter)
		rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get Smf Selection data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		//msg counter
		//pm.PegCounter(statistics.AuthenticationInforRetrivslSuccessesCounter)
		return http_wrapper.NewResponse(http.StatusOK, nil, authData)
	}

}
