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

// GetSmsData - retrieve a UE's SMS Subscription Data
func GetSmsData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
