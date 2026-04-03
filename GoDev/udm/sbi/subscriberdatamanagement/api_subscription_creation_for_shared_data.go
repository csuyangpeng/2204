/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 11/23/20 9:36 PM
 * Description:
 */

package subscriberdatamanagement

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/openapi/http_wrapper"
	"lite5gc/openapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SubscribeToSharedData - subscribe to notifications for shared data
func SubscribeToSharedData(c *gin.Context) {
	rlogger.FuncEntry(ModuleTag, nil)
	var sharedDataSubsReq models.SdmSubscription

	err := c.ShouldBindJSON(&sharedDataSubsReq)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		//todo logger
		//rlogger.SdmLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := sbicmn.NewRequest(c.Request, sharedDataSubsReq)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	HTTPResponse := http_wrapper.Response{}

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
