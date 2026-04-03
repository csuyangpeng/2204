/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 12/1/20 4:04 AM
 * Description:
 */
package utils

import (
	"lite5gc/openapi/models"
)

func CheckBoolToStringIsTure(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func CheckIntToBoolIsTure(b int) bool {
	if b == 0 {
		return false
	} else {
		return true
	}
}

// TODO check all enum type and set init value
func CheckEnumIsNull(em interface{}) string {
	switch em.(type) {
	case models.ImsVoPs:
		if em == nil || em.(models.ImsVoPs) == "" {
			return "HOMOGENEOUS_SUPPORT"
		}
		return string(em.(models.ImsVoPs))
	case models.RatType:
		if em == nil || em.(models.RatType) == "" {
			return "NR"
		}
		return string(em.(models.RatType))
	case models.ServiceName:
		if em == nil || em.(models.ServiceName) == "" {
			return ""
		}
		return string(em.(models.ServiceName))
	default:
		return ""
	}
}
