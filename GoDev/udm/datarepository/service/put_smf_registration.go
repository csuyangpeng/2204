/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 12/2/20 1:59 AM
 * Description:
 */
package service

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/openapi/models"
	"lite5gc/udm/datarepository/dao"
	"lite5gc/udm/datarepository/model"
	"lite5gc/udm/dbmgr"
	"net/http"
)

func PutSmfRegistration(supi string, pdussId int32, smfReg *models.SmfRegistration) *models.ProblemDetails {
	//Step 1
	rlogger.FuncEntry(ModuleTag, nil)
	//todo

	//Step 2
	tx := dbmgr.DBGorm.Begin()

	//Step 3
	smfreg := &model.SmfRegistration{
		Supi:                        supi,
		SmfInstanceId:               smfReg.SmfInstanceId,
		SupportedFeatures:           smfReg.SupportedFeatures,
		PduSessionId:                pdussId,
		PcscfRestorationCallbackUri: smfReg.PcscfRestorationCallbackUri,
		PlmnId:                      smfReg.PlmnId.Mcc + smfReg.PlmnId.Mnc,
		PgwFqdn:                     "",
	}
	regId, err := dao.InsertSmfReg(tx, smfreg)

	if err != nil {
		tx.Callback()
		return &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to insert into smf_registration table :%v", err),
		}
	}

	ss := model.SmfSnssai{
		SmfRegId: regId,
		Sst:      smfReg.SingleNssai.Sst,
		Sd:       smfReg.SingleNssai.Sd,
	}
	err = dao.InsertSmfSnssai(tx, &ss)

	if err != nil {
		tx.Callback()
		return &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to insert into smf_snssai table :%v", err),
		}
	}
	tx.Commit()

	return nil
}
