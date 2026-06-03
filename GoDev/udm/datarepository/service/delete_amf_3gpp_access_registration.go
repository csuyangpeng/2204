/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 12/2/20 2:07 AM
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

func Delete3gppRegistration(supi string) *models.ProblemDetails {
	rlogger.FuncEntry(ModuleTag, nil)

	err := dao.Delete3gppreg(supi)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to  deregistration for 3GPP access , key(%s), error(%s)", supi, err),
		}
		return &problemDetails
	}

	return nil
}
