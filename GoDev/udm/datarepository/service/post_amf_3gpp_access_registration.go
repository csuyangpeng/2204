/** Copyright(C),2020-2022
* Author: Jaytan
* Date: 11/23/20 11:01 PM
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
	"strings"
)

func PutAmf3GppAccessRegistration(supi string, amf3GppAccessRegistration models.Amf3GppAccessRegistration) *models.ProblemDetails {
	rlogger.FuncEntry(ModuleTag, nil)
	tx := dbmgr.DBGorm.Begin()

	supiId, err := dao.GetSupiIdBySupi(tx, supi)

	if err != nil {
		tx.Callback()
		return &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "USER_NOT_FOUND",
			Detail: fmt.Sprintf("failed to insert into supi table :%v", err),
		}
	}

	amf3gppAccessReg := model.Amf3gppAccessRegistration{
		SupiId:                      supiId,
		AmfInstanceId:               amf3GppAccessRegistration.AmfInstanceId,
		DeregCallbackUri:            amf3GppAccessRegistration.DeregCallbackUri,
		GamAmfId:                    amf3GppAccessRegistration.Guami.AmfId,
		GamPlmnId:                   amf3GppAccessRegistration.Guami.PlmnId.Mcc + amf3GppAccessRegistration.Guami.PlmnId.Mnc,
		RatType:                     utils.CheckEnumIsNull(amf3GppAccessRegistration.RatType),
		SupportedFeatures:           amf3GppAccessRegistration.SupportedFeatures,
		PurgeFlag:                   utils.CheckBoolToStringIsTure(amf3GppAccessRegistration.PurgeFlag),
		Pei:                         amf3GppAccessRegistration.Pei,
		ImsVoPs:                     utils.CheckEnumIsNull(amf3GppAccessRegistration.ImsVoPs),
		AmfServiceNameDereg:         utils.CheckEnumIsNull(amf3GppAccessRegistration.AmfServiceNameDereg),
		PcscfRestorationCallbackUri: amf3GppAccessRegistration.PcscfRestorationCallbackUri,
		AmfServiceNamePcscfRest:     utils.CheckEnumIsNull(amf3GppAccessRegistration.AmfServiceNamePcscfRest),
		InitialRegistrationInd:      utils.CheckBoolToStringIsTure(amf3GppAccessRegistration.InitialRegistrationInd),
		DrFlag:                      utils.CheckBoolToStringIsTure(amf3GppAccessRegistration.DrFlag),
	}
	amfRegId, err := dao.InsertAmf3gppRegInfo(tx, &amf3gppAccessReg)
	if err != nil {
		tx.Callback()
		return &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to insert into amf3gppReg table :%v", err),
		}
	}

	//Step 4
	if (amf3GppAccessRegistration.BackupAmfInfo != nil) && len(amf3GppAccessRegistration.BackupAmfInfo) != 0 {
		sqlStr := "INSERT INTO guami (amf_gpp_reg_id,amf_back_up_id,amf_id,plmn_id) VALUES "
		const rowSQL = "(?,?,?,?)"
		for _, elem := range amf3GppAccessRegistration.BackupAmfInfo {
			backUpAmf := &model.AmfBackupInfo{
				AmfGppRegId: amfRegId,
				BackUpAmf:   elem.BackupAmf,
			}
			backUpAmfId, err := dao.InsertAmfBackInfo(tx, backUpAmf)
			if err != nil {
				tx.Rollback()
				return &models.ProblemDetails{
					Status: http.StatusForbidden,
					Cause:  "SYSTEM_FAILURE",
					Detail: fmt.Sprintf("failed to insert into AmfBackUp table :%v", err),
				}
			}
			if elem.GuamiList != nil && len(elem.GuamiList) != 0 {
				for _, elem := range elem.GuamiList {
					inserts := []string{}
					inserts = append(inserts, rowSQL)
					vals := []interface{}{}
					vals = append(vals, amfRegId, backUpAmfId, elem.AmfId, elem.PlmnId.Mcc+elem.PlmnId.Mnc)
					sqlExec := sqlStr + strings.Join(inserts, ",")
					err := tx.Exec(sqlExec, vals...).Error
					if err != nil {
						tx.Rollback()
						return &models.ProblemDetails{
							Status: http.StatusForbidden,
							Cause:  "SYSTEM_FAILURE",
							Detail: fmt.Sprintf("failed to insert into guami table :%v", err),
						}
					}
				}
			}
		}
	}
	tx.Commit()
	return nil
}
