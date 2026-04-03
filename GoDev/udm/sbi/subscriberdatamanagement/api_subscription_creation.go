/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 11/23/20 9:36 PM
 * Description:
 */

package subscriberdatamanagement

import (
	"github.com/gin-gonic/gin"
	"lite5gc/cmn/sbicmn"
	"lite5gc/openapi/models"
	"lite5gc/udm/sbi/handler"
	"net/http"
)

// Subscribe - subscribe to notifications
func Subscribe(c *gin.Context) {

	var sdmSubscriptionReq models.SdmSubscription

	err := c.ShouldBindJSON(&sdmSubscriptionReq)
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

	req := sbicmn.NewRequest(c.Request, sdmSubscriptionReq)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	//handleMsg := udm_message.NewHandlerMessage(udm_message.EventSubscribe, req)
	//udm_handler.SendMessage(handleMsg)
	//
	//rsp := <-handleMsg.ResponseChan

	HTTPResponse := udmhdl.HandleSubscribe(req)

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
