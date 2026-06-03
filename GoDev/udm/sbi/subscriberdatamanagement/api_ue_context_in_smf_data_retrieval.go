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

// GetUeContextInSmfData - retrieve a UE's UE Context In SMF Data
func GetUeContextInSmfData(c *gin.Context) {

	req := sbicmn.NewRequest(c.Request, nil)
	req.Params["supi"] = c.Params.ByName("supi")

	//handleMsg := udm_message.NewHandlerMessage(udm_message.EventGetUeContextInSmfData, req)
	//udm_handler.SendMessage(handleMsg)
	//
	//rsp := <-handleMsg.ResponseChan

	HTTPResponse := udmhdl.HandleGetUeContextInSmfData(req)

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
