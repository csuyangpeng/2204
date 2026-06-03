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

// GetSmData - retrieve a UE's Session Management Subscription Data
func GetSmData(c *gin.Context) {

	req := sbicmn.NewRequest(c.Request, nil)
	req.Params["supi"] = c.Param("supi")
	req.Query.Set("plmn-id", c.Query("plmn-id"))

	//handleMsg := udm_message.NewHandlerMessage(udm_message.EventGetSmData, req)
	//udm_handler.SendMessage(handleMsg)
	//
	//rsp := <-handleMsg.ResponseChan
	HTTPResponse := udmhdl.HandleGetSmData(req)

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
