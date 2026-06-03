package udmdata

import (
	"lite5gc/cmn/nas/nasie"
)

//29503 Table 6.1.6.2.8-1
type SessMgntSubscripitonData struct {
	SingleNssai nasie.SNssai

	DnnConfigs map[string]*nasie.DNNConfiguration

	//internalGroupIds	array(GroupId)	O	1..N
	// List of internal group identifier; see 3GPP TS 23.501 [2] subclause 5.9.7
	internalGroupIds []string

	//sharedDnnConfigurationsId	SharedDataId	O	1..0
	// Identifier of shared data.
	sharedDnnConfigurationsId string
}

//func (p *SessMgntSubscripitonData) String() string {
//	var rtString string
//
//	rtString = fmt.Sprintf("SessMgntSubsriptionData{\n Snssai(%s)\n DnnConfigs{ \n", p.SingleNssai)
//
//	for k, v := range p.DnnConfigs {
//		rtString += fmt.Sprintf("  Dnn(%s)\n  DnnConfig(%s) \n", k, v)
//	}
//
//	rtString += " }\n"
//
//	rtString += "}"
//
//	return rtString
//}
func FindSessMgntSubsDataWithSnssai(snssaiStr string, smsdList []SessMgntSubscripitonData) *SessMgntSubscripitonData {
	for _, v := range smsdList {
		if v.SingleNssai.String() == snssaiStr {
			return &v
		}
	}

	return nil
}
