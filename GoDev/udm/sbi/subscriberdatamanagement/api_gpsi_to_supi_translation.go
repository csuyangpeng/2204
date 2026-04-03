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

// GetIdTranslationResult - retrieve a UE's SUPI
func GetIdTranslationResult(c *gin.Context) {

	req := sbicmn.NewRequest(c.Request, nil)
	req.Params["gpsi"] = c.Params.ByName("gpsi")
	req.Query.Set("SupportedFeatures", c.Query("supported-features"))

	//handleMsg := udm_message.NewHandlerMessage(udm_message.EventGetIdTranslationResult, req)
	//udm_handler.SendMessage(handleMsg)
	//
	//rsp := <-handleMsg.ResponseChan
	//Todo
	HTTPResponse := http_wrapper.Response{}

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
