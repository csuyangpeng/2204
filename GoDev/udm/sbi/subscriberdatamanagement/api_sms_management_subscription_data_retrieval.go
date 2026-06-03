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

// GetSmsMngData - retrieve a UE's SMS Management Subscription Data
func GetSmsMngData(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
