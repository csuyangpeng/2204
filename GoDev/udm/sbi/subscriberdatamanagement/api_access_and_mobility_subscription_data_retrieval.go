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
	udmhdl "lite5gc/udm/sbi/handler"
)

// GetAmData - retrieve a UE's Access and Mobility Subscription Data
func GetAmData(c *gin.Context) {
	rlogger.FuncEntry(ModuleTag, nil)
	req := sbicmn.NewRequest(c.Request, nil)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Query.Set("plmn-id", c.Query("plmn-id"))
	fmt.Println(c.FullPath())

	//handlerMsg := udm_message.NewHandlerMessage(udm_message.EventGetAmData, req)

	//data, _ := udm_producer.HandleGetAmData(req)

	//response := sbicmn.NewResponse(http.StatusOK, nil, data)

	HTTPResponse := udmhdl.HandleGetAmData(req)

	//c.JSON(response.Status, response.Body)
	c.JSON(HTTPResponse.Status, HTTPResponse.Body)
}
