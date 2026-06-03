/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 11/23/20 9:36 PM
 * Description:
 */
package subscriberdatamanagement

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUeContextInSmsfData - retrieve a UE's UE Context In SMSF Data
func GetUeContextInSmsfData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
