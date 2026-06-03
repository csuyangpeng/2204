package configure

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

//UDM configuration
type UdmConfig struct {
	UdmInstanceId string
	Version       VerConfig
	Logger        LoggerConfig
	Sbi           SBI
	PlmnList      types3gpp.PlmnList
}

func (p *UdmConfig) ConversionUdmConf(cmData *CmUdmConfig) error {
	// udmIntsanceId
	p.UdmInstanceId = cmData.UdmInstanceId
	// version
	p.Version = cmData.Version
	// logger config
	p.Logger = cmData.Logger
	// sbi config
	p.Sbi = cmData.Sbi

	// plmn list
	for _, v := range cmData.PlmnList {
		var plmn types3gpp.PlmnID
		err := plmn.SetString(v)
		if err != nil {
			rlogger.Trace(types.ModuleCmnConf, rlogger.ERROR, nil,
				"failed to set plmn string(%s)", v)
			continue
		}

		p.PlmnList.AddPlmn(plmn)
	}
	return nil
}
