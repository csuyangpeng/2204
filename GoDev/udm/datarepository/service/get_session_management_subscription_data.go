/** Copyright(C),2020-2022
* Author: Jaytan
* Date: 11/23/20 10:56 PM
* Description:
 */
package service

import (
	"fmt"
	"lite5gc/openapi/models"
	"lite5gc/udm/datarepository/dao"
	"net/http"
	"strings"
)

//( *models.AccessAndMobilitySubscriptionData, *models.ProblemDetails)

func GetSmData(supi string) (*[]models.SessionManagementSubscriptionData, *models.ProblemDetails) {
	//fetch data
	//single nssai
	var sm models.SessionManagementSubscriptionData

	snssaiData, err := dao.GetSnssaiBySupi(supi)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to get snssai with supi, key(%s), error(%s)", supi, err),
		}
		return nil, &problemDetails
	}
	var sn models.Snssai
	sn.Sd = snssaiData.Sd
	sn.Sst = int32(snssaiData.Sst)
	sm.SingleNssai = &sn
	sm.DnnConfigurations = make(map[string]models.DnnConfiguration)

	dnnCfg, err := dao.GetDnnConfigByDnnName(supi)
	if err != nil {
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: fmt.Sprintf("failed to get snssai with supi, key(%s), error(%s)", supi, err),
		}
		return nil, &problemDetails
	}
	//dnn configuration
	var dnnC models.DnnConfiguration
	//default pdu session types
	var pst models.PduSessionTypes
	switch dnnCfg.DefPduSessType {
	case "IPV4":
		pst.DefaultSessionType = models.PduSessionType_IPV4
	case "IPV6":
		pst.DefaultSessionType = models.PduSessionType_IPV6
	case "IPV4V6":
		pst.DefaultSessionType = models.PduSessionType_IPV4_V6
	case "UNSTRUCTURED":
		pst.DefaultSessionType = models.PduSessionType_UNSTRUCTURED
	case "ETHERNET":
		pst.DefaultSessionType = models.PduSessionType_ETHERNET
	default:
		pst.DefaultSessionType = models.PduSessionType_IPV4
	}
	//allow pdu session types
	apsts := strings.Split(dnnCfg.AllowedPduSessTypes, ",")
	ts := make([]models.PduSessionType, len(apsts))
	for _, v := range apsts {
		switch v {
		case "IPV4":
			ts = append(ts, models.PduSessionType_IPV4)
		case "IPV6":
			ts = append(ts, models.PduSessionType_IPV6)
		case "IPV4V6":
			ts = append(ts, models.PduSessionType_IPV4_V6)
		case "UNSTRUCTURED":
			ts = append(ts, models.PduSessionType_UNSTRUCTURED)
		case "ETHERNET":
			ts = append(ts, models.PduSessionType_ETHERNET)
		default:
			ts = append(ts, models.PduSessionType_IPV4)
		}
	}
	ts = ts[1:] // delete the first one value(null)
	pst.AllowedSessionTypes = ts
	dnnC.PduSessionTypes = &pst
	//default ssc modes
	var sms models.SscModes
	switch dnnCfg.DefSscMode {
	case "SSC_MODE_1":
		sms.DefaultSscMode = models.SscMode__1
	case "SSC_MODE_2":
		sms.DefaultSscMode = models.SscMode__2
	case "SSC_MODE_3":
		sms.DefaultSscMode = models.SscMode__3
	default:
		sms.DefaultSscMode = models.SscMode__1
	}
	//allow ssc modes
	asms := strings.Split(dnnCfg.AllowedSscMode, ",")
	asm := make([]models.SscMode, len(asms))
	for _, v := range apsts {
		switch v {
		case "SSC_MODE_1":
			asm = append(asm, models.SscMode__1)
		case "SSC_MODE_2":
			asm = append(asm, models.SscMode__2)
		case "SSC_MODE_3":
			asm = append(asm, models.SscMode__3)
		default:
			asm = append(asm, models.SscMode__1)
		}
	}
	asm = asm[1:] // delete the first one value(null)
	sms.AllowedSscModes = asm
	dnnC.SscModes = &sms
	//IwkEpsInd
	switch dnnCfg.IwkEpsInd {
	case "TRUE":
		dnnC.IwkEpsInd = true
	case "FALSE":
		dnnC.IwkEpsInd = false
	default:
		dnnC.IwkEpsInd = false
	}
	//Var5gQosProfile
	var subdq models.SubscribedDefaultQos
	subdq.Var5qi = int32(dnnCfg.FiveQI)
	subdq.PriorityLevel = int32(dnnCfg.PriorityLevel)

	var arp models.Arp
	temp := int32(dnnCfg.ArpPriorityLevel)
	arp.PriorityLevel = temp
	switch dnnCfg.PreemptCap {
	case "NOT_PREEMPT":
		arp.PreemptCap = models.PreemptionCapability_NOT_PREEMPT
	case "MAY_PREEMPT":
		arp.PreemptCap = models.PreemptionCapability_MAY_PREEMPT
	default:
		arp.PreemptCap = models.PreemptionCapability_NOT_PREEMPT
	}
	switch dnnCfg.PreemptVuln {
	case "NOT_PREEMPTABLE":
		arp.PreemptVuln = models.PreemptionVulnerability_NOT_PREEMPTABLE
	case "PREEMPTABLE":
		arp.PreemptVuln = models.PreemptionVulnerability_PREEMPTABLE
	default:
		arp.PreemptVuln = models.PreemptionVulnerability_PREEMPTABLE
	}
	subdq.Arp = &arp
	dnnC.Var5gQosProfile = &subdq
	//SessionAmbr
	var ambr models.Ambr
	ambr.Downlink = dnnCfg.SessAmbrDownlink
	ambr.Uplink = dnnCfg.SessAmbrUplink
	dnnC.SessionAmbr = &ambr
	//Var3gppChargingCharacteristics
	dnnC.Var3gppChargingCharacteristics = dnnCfg.ChargingChart
	//StaticIpAddress
	var ip models.IpAddress
	ip.Ipv4Addr = dnnCfg.StaticIpv4Addr
	ip.Ipv6Addr = dnnCfg.StaticIpv6Addr
	dnnC.StaticIpAddress = make([]models.IpAddress, 1)
	dnnC.StaticIpAddress[0] = ip
	//UpSecurity
	var ups models.UpSecurity
	switch dnnCfg.UpSecurityConfid {
	case "REQUIRED":
		ups.UpConfid = models.UpConfidentiality_REQUIRED
	case "PREFERRED":
		ups.UpConfid = models.UpConfidentiality_PREFERRED
	case "NOT_NEEDED":
		ups.UpConfid = models.UpConfidentiality_NOT_NEEDED
	default:
		ups.UpConfid = models.UpConfidentiality_NOT_NEEDED
	}
	switch dnnCfg.UpSecurityIntegr {
	case "REQUIRED":
		ups.UpIntegr = models.UpIntegrity_REQUIRED
	case "PREFERRED":
		ups.UpIntegr = models.UpIntegrity_PREFERRED
	case "NOT_NEEDED":
		ups.UpIntegr = models.UpIntegrity_NOT_NEEDED
	default:
		ups.UpIntegr = models.UpIntegrity_NOT_NEEDED
	}
	dnnC.UpSecurity = &ups
	sm.DnnConfigurations[dnnCfg.Dnn] = dnnC
	//todo internalGroupIds
	//todo sharedDnnConfigurationsId
	smsd := make([]models.SessionManagementSubscriptionData, 1)
	//smsd = make([]models.SessionManagementSubscriptionData, 1)
	smsd[0] = sm
	return &smsd, nil
}
