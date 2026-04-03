/** Copyright(C),2020-2022
* Author: Jaytan
* Date: 11/23/20 10:58 PM
* Description:
 */
package service

import (
	"fmt"
	"lite5gc/openapi/models"
	"lite5gc/udm/datarepository/dao"
	"net/http"
	"strconv"
)

func GetSmfSelectionSubscriptionData(supi string) (*models.SmfSelectionSubscriptionData, *models.ProblemDetails) {

	smfSel := models.SmfSelectionSubscriptionData{}
	//fetch data
	snssaiData, err := dao.GetSnssaiBySupi(supi)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to get snssai with supi, key(%s), error(%s)", supi, err),
		}
		return nil, &problemDetails
	}

	dnns, err := dao.GetDnnInfoBySnssaiID(snssaiData.Id)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to get dnn with snssai id, key(%s), error(%s)", snssaiData.Id, err),
		}
		return nil, &problemDetails
	}
	var snssaiI models.SnssaiInfo
	snssaiI.DnnInfos = make([]models.DnnInfo, 0)

	for _, v := range dnns {
		var d models.DnnInfo
		d.Dnn = v.Dnn
		d.LboRoamingAllowed = v.LboRoamingAllowed
		d.IwkEpsInd = v.IwkEpsInd
		d.DefaultDnnIndicator = v.DefaultDnnInd
		snssaiI.DnnInfos = append(snssaiI.DnnInfos, d)
	}

	var snssaikey string
	switch snssaiData.Sd {
	case "":
		snssaikey = strconv.Itoa(snssaiData.Sst)
	case "10000": //65536
		snssaikey = strconv.Itoa(snssaiData.Sst) + "-1"
	default:
		snssaikey = strconv.Itoa(snssaiData.Sst) + "-" + snssaiData.Sd
	}
	smfSel.SupportedFeatures = snssaiData.SupFeatures
	smfSel.SubscribedSnssaiInfos = make(map[string]models.SnssaiInfo)
	smfSel.SubscribedSnssaiInfos[snssaikey] = snssaiI
	//todo sharedSnssaiInfosId
	return &smfSel, nil
}
