/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 12/2/20 8:37 PM
 * Description:
 */
package service

import (
	"fmt"
	"lite5gc/openapi/models"
	"lite5gc/udm/dbmgr"
	"net/http"
)

func GetSmfRegistration(gpsi string) (*models.SmfRegistration, *models.ProblemDetails) {

	db := dbmgr.DBGorm
	smfRegData := models.SmfRegistration{}
	err := db.Preloads("smf_snssai").First(&smfRegData).Error

	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusNoContent,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to get supi id with supi, key(%s), error(%s)", gpsi, err),
		}
		return nil, &problemDetails
	}

	return nil, nil
}
