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
	"lite5gc/openapi/http_wrapper"
)

// GetNssai - retrieve a UE's subscribed NSSAI
func GetNssai(c *gin.Context) {

	req := sbicmn.NewRequest(c.Request, nil)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Query.Set("plmn-id", c.Query("plmn-id"))

	//handleMsg := udm_message.NewHandlerMessage(udm_message.EventGetNssai, req)
	//udm_handler.SendMessage(handleMsg)
	//
	//rsp := <-handleMsg.ResponseChan
	// ToDo
	HTTPResponse := http_wrapper.Response{}

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
