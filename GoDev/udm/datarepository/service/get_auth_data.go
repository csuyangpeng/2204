/** Copyright(C),2020-2022
* Author: Jaytan
* Date: 11/23/20 10:15 PM
* Description:
 */
package service

import (
	"encoding/hex"
	"fmt"
	"lite5gc/amf/sc/mobmgnt/mmutils"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/openapi/models"
	"lite5gc/udm/arpf/derivevec"
	"lite5gc/udm/arpf/manager"
	"lite5gc/udm/datarepository/dao"
	"lite5gc/udm/msg"
	"net/http"
	"strings"
)

func GetAuthData(supi string) (*models.AuthenticationInfoResult, *models.ProblemDetails) {
	rlogger.FuncEntry(ModuleTag, nil)
	// 5gc\cmn\udmAgentLayer\dbmgr\getAuthData.go start
	authResult := models.AuthenticationInfoResult{}
	snssaiData, err := dao.GetSnssaiBySupi(supi)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "USER_NOT_FOUND",
			Detail: fmt.Sprintf("failed to get snssai with supi, key(%s), error(%s)", supi, err),
		}
		return nil, &problemDetails
	}
	//fetch data
	authData, err := dao.GetAuthDataBySupi(supi)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "USER_NOT_FOUND",
			Detail: fmt.Sprintf("failed to get supi id with supi, key(%s), error(%s)", supi, err),
		}
		return nil, &problemDetails
	} else {
		data := &msg.AuthData{}

		//store ki
		if authData.Ki != "" || strings.Contains(authData.Ki, " ") {
			ki, err := hex.DecodeString(authData.Ki)
			if err != nil || len(ki) != int(types.KSize) {
				rlogger.Trace(ModuleTag, rlogger.ERROR, supi, "invalid ki from db")
				problemDetails := models.ProblemDetails{
					Status: http.StatusInternalServerError,
					Cause:  "SYSTEM_FAILURE",
					Detail: fmt.Sprintf("invalid ki from db, error(%s),len(%d)", err, len(ki)),
				}
				return nil, &problemDetails
			}
			for i, v := range ki {
				data.Ki[i] = v
			}
		}

		//store opc
		if authData.Opc != "" || strings.Contains(authData.Opc, " ") {
			opc, err := hex.DecodeString(authData.Opc)
			if err != nil || len(opc) != int(types.OpcSize) {
				rlogger.Trace(ModuleTag, rlogger.DEBUG, supi, "invalid opc from db")
				problemDetails := models.ProblemDetails{
					Status: http.StatusInternalServerError,
					Cause:  "SYSTEM_FAILURE",
					Detail: fmt.Sprintf("invalid opc from db, error(%s),len(%d)", err, len(opc)),
				}
				return nil, &problemDetails
			}
			for i, v := range opc {
				data.Opc[i] = v
			}
			data.IsOpc = true
		}

		//store op
		if authData.Op != "" || strings.Contains(authData.Op, " ") {
			op, err := hex.DecodeString(authData.Op)
			if err != nil || len(op) != int(types.OpSize) {
				rlogger.Trace(ModuleTag, rlogger.DEBUG, supi, "invalid op from db")
				problemDetails := models.ProblemDetails{
					Status: http.StatusInternalServerError,
					Cause:  "SYSTEM_FAILURE",
					Detail: fmt.Sprintf("invalid op from db, error(%s),len(%d)", err, len(op)),
				}
				return nil, &problemDetails
			}
			for i, v := range op {
				data.Op[i] = v
			}
			data.IsOpc = false
		}

		//store amf

		if authData.Amf != "" || strings.Contains(authData.Amf, " ") {
			amf, err := hex.DecodeString(authData.Amf)
			if err != nil || len(amf) != int(types.AmfSize) {
				rlogger.Trace(ModuleTag, rlogger.DEBUG, supi, "invalid amf from db")
				problemDetails := models.ProblemDetails{
					Status: http.StatusInternalServerError,
					Cause:  "SYSTEM_FAILURE",
					Detail: fmt.Sprintf("invalid amf from db, error(%s),len(%d)", err, len(amf)),
				}
				return nil, &problemDetails
			}
			for i, v := range amf {
				data.Amf[i] = v
			}
		}
		//5gc\cmn\udmAgentLayer\dbmgr\getAuthData.go stop

		//todo Sn name should be send from amf
		// get sn name
		// snName := mmutils.GenerateSnName(configure.UdmConf.Service.Plmn)
		snName := mmutils.GenerateSnName(configure.UdmConf.PlmnList.List[0])

		var ueSecCtxt *types.UeSecContext
		ueSecCtxt = &types.UeSecContext{}
		// store the Auth data from DB
		ueSecCtxt.Key = data.Ki
		ueSecCtxt.IsOpc = data.IsOpc
		ueSecCtxt.Op = data.Op
		ueSecCtxt.Opc = data.Opc
		ueSecCtxt.Amf = data.Amf
		ueSecCtxt.Sqn = [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00} //sqn init
		ueSecCtxt.SnName = snName
		//generate HeAV
		manager.RetreiveSqn(ueSecCtxt, nil)

		heAv, err := derivevec.DeriveHeAv(ueSecCtxt)
		if err != nil {
			problemDetails := models.ProblemDetails{
				Status: http.StatusInternalServerError,
				Cause:  "SYSTEM_FAILURE",
				Detail: fmt.Sprintf("failed to derive HeAv, error(%s)", err),
			}
			return nil, &problemDetails
		}
		rlogger.Trace(ModuleTag, rlogger.INFO, nil, "DeriveHeAv Infos:%s", heAv)
		//var authResult models.AuthenticationInfoResult
		var auVec models.AuthenticationVector
		auVec.AvType = models.AvType__5_G_HE_AKA
		auVec.Rand = fmt.Sprintf("%x", heAv.Rand[:])
		auVec.Xres = fmt.Sprintf("%x", heAv.XRes[:])
		auVec.Autn = fmt.Sprintf("%x", heAv.Autn[:])
		auVec.CkPrime = fmt.Sprintf("%x", heAv.CK[:])
		auVec.IkPrime = fmt.Sprintf("%x", heAv.IK[:])
		auVec.XresStar = fmt.Sprintf("%x", heAv.XResStar[:])
		auVec.Kausf = fmt.Sprintf("%x", heAv.Kausf[:])

		rlogger.Trace(ModuleTag, rlogger.INFO, nil, "authentication vectors(%s)", auVec.String())
		authResult.AuthenticationVector = &auVec
		authResult.SupportedFeatures = snssaiData.SupFeatures
		authResult.Supi = supi
		authResult.AuthType = models.AuthType__5_G_AKA
	}

	return &authResult, nil
}
