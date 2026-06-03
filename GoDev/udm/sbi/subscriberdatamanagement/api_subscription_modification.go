/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 11/23/20 9:36 PM
 * Description:
 */

package subscriberdatamanagement

import (
	"lite5gc/cmn/sbicmn"
	"lite5gc/openapi/models"
	udmhdl "lite5gc/udm/sbi/handler"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Modify - modify the subscription
func Modify(c *gin.Context) {

	var sdmSubsModificationReq models.SdmSubsModification
	err := c.ShouldBindJSON(&sdmSubsModificationReq)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		//todo logger trace
		//rlogger.Handlelog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := sbicmn.NewRequest(c.Request, sdmSubsModificationReq)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	//handleMsg := udm_message.NewHandlerMessage(udm_message.EventSubscribe, req)
	//udm_handler.SendMessage(handleMsg)
	//
	//rsp := <-handleMsg.ResponseChan
	//
	HTTPResponse := udmhdl.HandleSubscribe(req)

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}

// ModifyForSharedData - modify the subscription
func ModifyForSharedData(c *gin.Context) {

	var sharedDataSubscriptions models.SdmSubsModification
	err := c.ShouldBindJSON(&sharedDataSubscriptions)
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		//todo logger trace
		//rlogger.SdmLog.Errorln(problemDetail)
		c.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := sbicmn.NewRequest(c.Request, sharedDataSubscriptions)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	//handleMsg := udm_message.NewHandlerMessage(udm_message.EventSubscribe, req)
	//udm_handler.SendMessage(handleMsg)
	//
	//rsp := <-handleMsg.ResponseChan

	HTTPResponse := udmhdl.HandleSubscribe(req)

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
