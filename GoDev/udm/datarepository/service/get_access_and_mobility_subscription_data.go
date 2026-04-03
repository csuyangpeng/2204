/** Copyright(C),2020-2022
* Author: Jaytan
* Date: 11/23/20 9:46 PM
* Description:
 */
package service

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/openapi/models"
	"lite5gc/udm/datarepository/dao"
	"net/http"
)

const ModuleTag rlogger.ModuleTag = types.ModuleUdmSbiDAO

func GetAccessAndMobilitySubscriptionData(supi string) (*models.AccessAndMobilitySubscriptionData, *models.ProblemDetails) {
	rlogger.FuncEntry(ModuleTag, nil)

	//fetch data
	amsd := &models.AccessAndMobilitySubscriptionData{}
	amsdData, err := dao.GetAmsdDataBySupi(supi)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "USER_NOT_FOUND",
			Detail: fmt.Sprintf("failed to get supi id with supi, key(%s), error(%s)", supi, err),
		}
		return amsd, &problemDetails
	}

	//conversion data format
	//store subscribedUeAmbr
	amsd.SubscribedUeAmbr = &models.AmbrRm{}
	if amsdData.UeAmbrDl != "" {
		amsd.SubscribedUeAmbr.Downlink = amsdData.UeAmbrDl
	}
	if amsdData.UeAmbrUl != "" {
		amsd.SubscribedUeAmbr.Uplink = amsdData.UeAmbrUl
	}

	//store rfsp index id
	if amsdData.RfspIndexId != 0 {
		temp := int32(amsdData.RfspIndexId)
		amsd.RfspIndex = temp
	}

	//store subs registor timer
	if amsdData.SubsRegTimer != 0 {
		temp := int32(amsdData.SubsRegTimer)
		amsd.SubsRegTimer = temp
	}

	//store active time
	if amsdData.ActiveTime != 0 {
		temp := int32(amsdData.ActiveTime)
		amsd.ActiveTime = temp
	}

	//store mpsPriority
	switch amsdData.MpsPriority {
	case "True":
		amsd.MpsPriority = true
	case "False":
		amsd.MpsPriority = false
	default:
		amsd.MpsPriority = false
	}

	//store mscPriority
	switch amsdData.McsPriority {
	case "True":
		amsd.McsPriority = true
	case "False":
		amsd.McsPriority = false
	default:
		amsd.McsPriority = false
	}

	//store micoAllowed
	switch amsdData.MicoAllowed {
	case "True":
		amsd.MicoAllowed = true
	case "False":
		amsd.MicoAllowed = false
	default:
		amsd.MicoAllowed = false
	}

	//store dlpacketcount
	if amsdData.DlPacketCount >= -1 {
		amsd.DlPacketCount = int32(amsdData.DlPacketCount)
	}

	//get support features
	if amsdData.SupFeatures != "" {
		amsd.SupportedFeatures = amsdData.SupFeatures
	}

	//store nssai
	amsd.Nssai = &models.Nssai{}
	amsd.Nssai.DefaultSingleNssais = make([]models.Snssai, 1)
	amsd.Nssai.DefaultSingleNssais[0].Sst = int32(amsdData.Sst)
	amsd.Nssai.DefaultSingleNssais[0].Sd = amsdData.Sd

	//todo store Gpsis
	//todo store InternalGroupIds
	//todo store RatRestrictions
	//todo store ForbiddenAreas
	//todo store ServiceAreaRestriction
	//todo store CoreNetworkTypeRestrictions
	//todo store UeUsageType
	//todo store SorInfo
	//todo store SharedAmDataIds
	//todo store OdbPacketServices

	return amsd, nil
}
