package configure

import (
	"lite5gc/cmn/types3gpp"
	"lite5gc/openapi/models"
)

var (
	SysConf SystemConfig
	AmfConf AmfConfig
	SmfConf SmfConfig
	UpfConf UpfConfig
	UdmConf UdmConfig

	// CM global
	CmSysConf CmSysConfig
	CmAmfConf CmAmfConfig
	CmSmfConf CmSmfConfig
	CmUpfConf CmUpfConfig
	CmUdmConf CmUdmConfig
)

func GetTypesGuami() types3gpp.Guami {
	guami := types3gpp.Guami{}
	guami.PlmnId = AmfConf.PlmnList.List[0]
	guami.AmfId = AmfConf.Service.AmfIdentifier
	return guami
}

func GetModelsGuami() models.Guami {
	guamiAmfStr := AmfConf.Service.AmfIdentifier.String()
	guami := models.Guami{PlmnId: &models.PlmnId{}}
	guami.PlmnId.Mcc = AmfConf.PlmnList.List[0].GetMccString()
	guami.PlmnId.Mnc = AmfConf.PlmnList.List[0].GetMncString()
	guami.AmfId = guamiAmfStr
	return guami
}

func GetGuamiStr() string {
	guami := GetModelsGuami()
	guamiAmfStr := AmfConf.Service.AmfIdentifier.String()
	guamiStr := guami.PlmnId.Mcc + guami.PlmnId.Mnc + guamiAmfStr
	return guamiStr
}
