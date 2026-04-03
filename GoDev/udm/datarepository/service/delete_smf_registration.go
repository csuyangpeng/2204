/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 12/2/20 2:08 AM
 * Description:
 */
package service

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/openapi/models"
	"lite5gc/udm/datarepository/dao"
	"net/http"
)

func DeleteSmfRegistration(supi string, pdussId int32) *models.ProblemDetails {

	rlogger.FuncEntry(ModuleTag, nil)

	err := dao.DeleteSmfReg(supi, pdussId)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to  deregistration for smf reg , key(%s),key(%d), error(%s)", supi, pdussId, err),
		}
		return &problemDetails
	}

	return nil

}
