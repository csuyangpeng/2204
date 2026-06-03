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

// GetTraceData - retrieve a UE's Trace Configuration Data
func GetTraceData(c *gin.Context) {

	req := sbicmn.NewRequest(c.Request, nil)
	req.Params["supi"] = c.Params.ByName("supi")
	req.Query.Set("plmn-id", c.Query("plmn-id"))

	HTTPResponse := http_wrapper.Response{}

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
