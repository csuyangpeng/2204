/** Copyright(C),2020-2022
* Author: zmj
* Date: 12/11/20 10:33 AM
* Description:
 */
package sbicmn

import (
	"github.com/gin-gonic/gin"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"time"
)

func LogGin(c *gin.Context) {
	startTime := time.Now()

	c.Next()

	endTime := time.Now()

	latencyTime := endTime.Sub(startTime)

	reqMethod := c.Request.Method

	reqUri := c.Request.RequestURI

	statusCode := c.Writer.Status()

	clientIP := c.ClientIP()

	rlogger.Trace(types.ModCmn, rlogger.INFO, nil, "| %3d | %13v | %15s | %s | %s |",
		statusCode,
		latencyTime,
		clientIP,
		reqMethod,
		reqUri)
}
