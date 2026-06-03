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

// PutUpuAck - Nudm_Sdm Info for UPU service operation
func PutUpuAck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
