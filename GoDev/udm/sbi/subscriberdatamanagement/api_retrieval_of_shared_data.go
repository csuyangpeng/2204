/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 11/23/20 9:36 PM
 * Description:
 */

package subscriberdatamanagement

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/openapi/http_wrapper"
)

// GetSharedData - retrieve shared data
func GetSharedData(c *gin.Context) {
	rlogger.FuncEntry(ModuleTag, nil)
	req := sbicmn.NewRequest(c.Request, nil)
	req.Query["sharedDataIds"] = c.QueryArray("shared-data-ids")

	//handleMsg := udm_message.NewHandlerMessage(udm_message.EventGetSharedData, req)
	//udm_handler.SendMessage(handleMsg)
	//
	//rsp := <-handleMsg.ResponseChan

	//ToDo
	HTTPResponse := http_wrapper.Response{}

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
	fmt.Println(HTTPResponse)
}
