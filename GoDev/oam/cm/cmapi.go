package cm

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/oam/cm/yaml"
)

//var sys_conf_file = "data/cm_sys_conf.yaml"
//var amf_conf_file = "data/cm_amf_conf.yaml"
//var smf_conf_file = "data/cm_smf_conf.yaml"
//var upf_conf_file = "data/cm_upf_conf.yaml"
//var udm_conf_file = "data/cm_udm_conf.yaml"

// LoadSysConfig: load sys config from /path/to/sys.yaml to cmglobal
func LoadSysConfig(confFile string) {
	rlogger.FuncEntry(types.ModuleOamCm, nil)
	yaml.Load(confFile, &configure.CmSysConf)
	configure.SysConf.ConvertSystemConfig(&configure.CmSysConf)
}

// LoadAmfConfig: load amf config from /path/to/amf.yaml to cmglobal
func LoadAmfConfig(confFile string) error {
	rlogger.FuncEntry(types.ModuleOamCm, nil)
	yaml.Load(confFile, &configure.CmAmfConf)
	return configure.AmfConf.ConvertAmfConfig(&configure.CmAmfConf)
}

// LoadSmfConfig: load smf config from /path/to/smf.yaml to cmglobal
func LoadSmfConfig(confFile string) error{
	rlogger.FuncEntry(types.ModuleOamCm, nil)
	yaml.Load(confFile, &configure.CmSmfConf)
	return configure.SmfConf.ConvertSmfConfig(&configure.CmSmfConf)
}

// LoadUpfConfig: load upf config from /path/to/upf.yaml to cmglobal
func LoadUpfConfig(confFile string) {
	rlogger.FuncEntry(types.ModuleOamCm, nil)
	yaml.Load(confFile, &configure.UpfConf)
}

// LoadUdmConfig: load udm config from /path/to/udm.yaml to cmglobal
func LoadUdmConfig(confFile string) {
	rlogger.FuncEntry(types.ModuleOamCm, nil)
	yaml.Load(confFile, &configure.CmUdmConf)
	configure.UdmConf.ConversionUdmConf(&configure.CmUdmConf)
}
