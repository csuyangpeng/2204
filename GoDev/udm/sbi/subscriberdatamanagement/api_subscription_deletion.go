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

// Unsubscribe - unsubscribe from notifications
func Unsubscribe(c *gin.Context) {

	req := sbicmn.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["subscriptionId"] = c.Params.ByName("subscriptionId")

	HTTPResponse := http_wrapper.Response{}

	c.JSON(HTTPResponse.Status, HTTPResponse.Body)

}
