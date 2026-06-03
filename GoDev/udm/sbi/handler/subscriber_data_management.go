package udmhdl

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/openapi/http_wrapper"
	"lite5gc/openapi/models"
	"lite5gc/udm/datarepository/service"
	"net/http"
	"strings"
)

// const types.ModuleUdmSbiAUT = rlogger.PACKAGE_UDM_SBI_PRODUCER_types.ModuleUdmSbiAUT
//const types.ModuleUdmSbiAUT rlogger.types.ModuleUdmSbiAUT = types.ModuleUdmSbiAUT

func HandleGetAmData(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiAUT, nil)
	//msg counter
	// pm.PegCounter(statistics.AMDataRetrivslRequestsCounter)
	supi := req.Params["supi"]
	supi = strings.TrimPrefix(supi, "imsi-")

	// supiInt, err := strconv.Atoi(supi)
	//if err != nil {
	//	//msg counter
	//	// pm.PegCounter(statistics.AMDataRetrivslFailureCounter)
	//	rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get supi ,err(%s)", err)
	//}
	amsd, problemDetails := service.GetAccessAndMobilitySubscriptionData(supi)
	if problemDetails != nil {
		//msg counter
		//pm.PegCounter(statistics.AMDataRetrivslFailureCounter)
		rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get Access And Mobility Subscription data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		//msg counter
		//pm.PegCounter(statistics.AMDataRetrivslSuccessesCounter)

		return http_wrapper.NewResponse(http.StatusOK, nil, amsd)
	}

}

func HandleGetSmfSelectData(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiAUT, nil)
	//msg counter
	//pm.PegCounter(statistics.SMFSelectionDataRetrivslRequestsCounter)
	supi := req.Params["supi"]
	supi = strings.TrimPrefix(supi, "imsi-")
	//supiInt, err := strconv.Atoi(supi)
	//if err != nil {
	//	//msg counter
	//	//pm.PegCounter(statistics.SMFSelectionDataRetrivslFailureCounter)
	//	rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get supi ,err(%s)", err)
	//}
	smfSel, problemDetails := service.GetSmfSelectionSubscriptionData(supi)
	//if problemDetails != nil {
	//	//msg counter
	//	//pm.PegCounter(statistics.SMFSelectionDataRetrivslFailureCounter)
	//	rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get Smf Selection data ,err(%s)", err)
	//}

	if problemDetails != nil {
		//msg counter
		//pm.PegCounter(statistics.SMFSelectionDataRetrivslFailureCounter)
		rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get Smf Selection data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		//msg counter
		//pm.PegCounter(statistics.SMFSelectionDataRetrivslSuccessesCounter)
		return http_wrapper.NewResponse(http.StatusOK, nil, smfSel)
	}
}

func HandleGetSmData(req *sbicmn.Request) *http_wrapper.Response {

	rlogger.FuncEntry(types.ModuleUdmSbiAUT, nil)
	//msg counter
	//pm.PegCounter(statistics.SMDataRetrivslRequestsCounter)
	supi := req.Params["supi"]
	supi = strings.TrimPrefix(supi, "imsi-")
	//supiInt, err := strconv.Atoi(supi)
	//if err != nil {
	//	//msg counter
	//	pm.PegCounter(statistics.SMDataRetrivslFailureCounter)
	//	rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get supi ,err(%s)", err)
	//}
	sm, problemDetails := service.GetSmData(supi)

	//msg counter
	//pm.PegCounter(statistics.SMDataRetrivslSuccessesCounter)
	//udm_message.SendHttpResponseMessage(httpChannel, nil, http.StatusOK, sm)
	if problemDetails != nil {

		rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get Smf Selection data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {

		return http_wrapper.NewResponse(http.StatusOK, nil, sm)
	}
}

func HandleGetUeContextInSmfData(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiAUT, nil)
	//supiInt, err := strconv.Atoi(supi)
	supi := req.Params["supi"]
	supi = strings.TrimPrefix(supi, "imsi-")

	ucsd, problemDetails := service.GetUeContextInSmfData(supi)

	if problemDetails != nil {

		rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get UE Context in Smf data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {

		return http_wrapper.NewResponse(http.StatusOK, nil, ucsd)
	}
}

func HandleSubscribe(req *sbicmn.Request) *http_wrapper.Response {

	rlogger.FuncEntry(types.ModuleUdmSbiAUT, nil)
	supi := req.Params["supi"]
	supi = strings.TrimPrefix(supi, "imsi-")

	sdmSub := req.Body.(models.SdmSubscription)

	problemDetails := service.PostSdmSubscription(supi, sdmSub)

	if problemDetails != nil {

		rlogger.Trace(types.ModuleUdmSbiAUT, rlogger.ERROR, "fail to get UE Context in Smf data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {

		return http_wrapper.NewResponse(http.StatusOK, nil, sdmSub)
	}
	//udm_message.SendHttpResponseMessage(httpChannel, nil, http.StatusOK, sdmSub)
}
