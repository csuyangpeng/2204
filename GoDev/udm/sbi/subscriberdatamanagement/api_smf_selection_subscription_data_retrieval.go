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
	"lite5gc/udm/sbi/handler"
)

// GetSmfSelectData - retrieve a UE's SMF Selection Subscription Data
func GetSmfSelectData(c *gin.Context) {

	req := sbicmn.NewRequest(c.Request, nil)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Query.Set("plmn-id", c.Query("plmn-id"))

	//handlerMsg := udm_message.NewHandlerMessage(udm_message.EventGetSmfSelectData, req)
	//udm_handler.SendMessage(handlerMsg)
	//
	//rsp := <-handlerMsg.ResponseChan

	HTTPResponse := udmhdl.HandleGetSmfSelectData(req)

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	return
}
