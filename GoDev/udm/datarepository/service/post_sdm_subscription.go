/** Copyright(C),2020-2022
* Author: Jaytan
* Date: 11/23/20 11:02 PM
* Description:
 */
package service

import (
	logger "lite5gc/cmn/rlogger"
	"lite5gc/openapi/models"
)

func PostSdmSubscription(supi string, sdmSubs models.SdmSubscription) *models.ProblemDetails {
	logger.FuncEntry(ModuleTag, nil)
	//todo
	//fmt.Println("sdmSubs",sdmSubs)
	return nil
}
