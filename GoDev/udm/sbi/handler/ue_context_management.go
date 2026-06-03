package udmhdl

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/openapi/http_wrapper"
	"lite5gc/openapi/models"
	"lite5gc/udm/datarepository/service"
	"net/http"
	"strconv"
	"strings"
)

func HandleRegistrationAmf3gppAccess(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiUCX, nil)
	supi := req.Params["ueId"]
	supi = strings.TrimPrefix(supi, "imsi-")
	amf3GppAccessRegistration := req.Body.(models.Amf3GppAccessRegistration)
	problemDetails := service.PutAmf3GppAccessRegistration(supi, amf3GppAccessRegistration)

	if problemDetails != nil {

		rlogger.Trace(types.ModuleUdmSbiUCX, rlogger.ERROR, "fail to get UE Context in Amf data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {

		return http_wrapper.NewResponse(http.StatusCreated, nil, nil)
	}
}

func HandleDregistrationAmf3gppAccess(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiUCX, nil)
	supi := req.Params["ueId"]
	supi = strings.TrimPrefix(supi, "imsi-")
	amf3GppAccessRegistration := req.Body.(models.Amf3GppAccessRegistration)
	problemDetails := service.PutAmf3GppAccessRegistration(supi, amf3GppAccessRegistration)

	if problemDetails != nil {

		rlogger.Trace(types.ModuleUdmSbiUCX, rlogger.ERROR, "fail to get UE Context in Smf data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {

		return http_wrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
}

func HandleRegistrationSmf(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiUCX, nil)
	supi := req.Params["ueId"]
	supi = strings.TrimPrefix(supi, "imsi-")

	pduSessionId, err := strconv.ParseInt(req.Params["pduSessionId"], 10, 64)

	if err != nil {
		rlogger.Trace(types.ModuleUdmSbiUCX, rlogger.ERROR, "fail to get UE Context in Smf data ,err(%s)", err)
	}

	smfReg := req.Body.(models.SmfRegistration)
	problemDetails := service.PutSmfRegistration(supi, int32(pduSessionId), &smfReg)

	if problemDetails != nil {

		rlogger.Trace(types.ModuleUdmSbiUCX, rlogger.ERROR, "fail to get UE Context in Smf data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {

		return http_wrapper.NewResponse(http.StatusCreated, nil, nil)
	}
}

func HandleDregistrationSmf(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiUCX, nil)
	supi := req.Params["ueId"]
	supi = strings.TrimPrefix(supi, "imsi-")

	pduSessionId, err := strconv.ParseInt(req.Params["pduSessionId"], 10, 64)

	if err != nil {
		rlogger.Trace(types.ModuleUdmSbiUCX, rlogger.ERROR, "fail to get UE Context in Smf data ,err(%s)", err)
	}

	problemDetails := service.DeleteSmfRegistration(supi, int32(pduSessionId))

	if problemDetails != nil {
		rlogger.Trace(types.ModuleUdmSbiUCX, rlogger.ERROR, "fail to get UE Context in Smf data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {

		return http_wrapper.NewResponse(http.StatusNoContent, nil, nil)
	}
}

func HandleRetrievalSmf(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiUCX, nil)
	supi := req.Params["ueId"]
	supi = strings.TrimPrefix(supi, "imsi-")

	smfReg := req.Body.(models.SmfRegistration)
	problemDetails := service.PutSmfRegistration(supi, 2, &smfReg)

	if problemDetails != nil {

		rlogger.Trace(types.ModuleUdmSbiUCX, rlogger.ERROR, "fail to get UE Context in Smf data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {

		return http_wrapper.NewResponse(http.StatusCreated, nil, nil)
	}
}

func HandleRetrievalAmf3gppAccessReg(req *sbicmn.Request) *http_wrapper.Response {
	rlogger.FuncEntry(types.ModuleUdmSbiUCX, nil)
	gpsi := req.Params["ueId"]
	gpsi = strings.TrimPrefix(gpsi, "imsi-")

	registration, problemDetails := service.GetAmf3GppAccessRegistration(gpsi)

	if problemDetails != nil {

		rlogger.Trace(types.ModuleUdmSbiUCX, rlogger.ERROR, "fail to get UE Context in Amf3gpp data ,err(%s)", problemDetails)
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {

		return http_wrapper.NewResponse(http.StatusOK, nil, registration)
	}
}
