/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 11/30/20 2:46 AM
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
	"lite5gc/udm/utils"
	"net/http"
)

func GetAmf3GppAccessRegistration(gpsi string) (*models.Amf3GppAccessRegistration, *models.ProblemDetails) {

	rlogger.FuncEntry(ModuleTag, nil)

	db := dbmgr.DBGorm
	amf3gppReg := &models.Amf3GppAccessRegistration{}

	amf3gppdata, err := dao.Get3gppAmfRegData(db, gpsi)
	fmt.Println("--------------", amf3gppdata)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusNoContent,
			Cause:  "USER_NOT_FOUND",
			Detail: fmt.Sprintf("failed to get supi id with supi, key(%s), error(%s)", gpsi, err),
		}
		return nil, &problemDetails
	}
	backupamfs := []model.AmfBackupInfo{}
	err = db.Where("amf_gpp_reg_id = ?", amf3gppdata.Id).Preload("Guami").Find(&backupamfs).Error

	// amf3gppReg
	amf3gppReg.AmfInstanceId = amf3gppdata.AmfInstanceId
	amf3gppReg.PcscfRestorationCallbackUri = amf3gppdata.PcscfRestorationCallbackUri
	amf3gppReg.SupportedFeatures = amf3gppdata.SupportedFeatures
	amf3gppReg.InitialRegistrationInd = utils.CheckIntToBoolIsTure(amf3gppdata.InitialRegistrationInd)
	amf3gppReg.DrFlag = utils.CheckIntToBoolIsTure(amf3gppdata.DrFlag)
	amf3gppReg.AmfServiceNamePcscfRest = models.ServiceName(amf3gppdata.AmfServiceNamePcscfRest)
	amf3gppReg.AmfServiceNameDereg = models.ServiceName(amf3gppdata.AmfServiceNameDereg)
	amf3gppReg.ImsVoPs = models.ImsVoPs(amf3gppdata.ImsVoPs)
	amf3gppReg.PurgeFlag = utils.CheckIntToBoolIsTure(amf3gppdata.PurgeFlag)
	amf3gppReg.RatType = models.RatType(amf3gppdata.RatType)
	amf3gppReg.Pei = amf3gppdata.Pei

	// gumai
	gumai := &models.Guami{}

	// plmnId
	plmnId := &models.PlmnId{
		Mcc: amf3gppdata.GamPlmnId[0:3],
		Mnc: amf3gppdata.GamPlmnId[3:],
	}
	gumai.PlmnId = plmnId
	gumai.AmfId = amf3gppdata.GamAmfId
	amf3gppReg.Guami = gumai

	if backupamfs != nil && len(backupamfs) != 0 {

		for _, value := range backupamfs {
			bk := models.BackupAmfInfo{
				BackupAmf: value.BackUpAmf,
			}
			if len(value.Guami) != 0 {
				for _, v := range value.Guami {
					gm := models.Guami{
						PlmnId: &models.PlmnId{
							Mcc: v.PlmnId[0:3],
							Mnc: v.PlmnId[3:],
						},
						AmfId: v.AmfId,
					}
					bk.GuamiList = append(bk.GuamiList, gm)
				}
			}
			amf3gppReg.BackupAmfInfo = append(amf3gppReg.BackupAmfInfo, bk)
		}
	}
	return amf3gppReg, nil
}
