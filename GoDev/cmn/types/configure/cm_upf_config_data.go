/*
* Copyright(C),2020-2022
* Author:  xiaoyun
* Date:    12/10/20 5:35 AM
* Description:
 */
package configure

type CmUpfConfig struct {
	Version VerConfig
	Logger  LoggerConfig
	N3      IpUpConfig
	N4      N4Config
	N6      IpUpConfig
	Nff     NffConfig
	DnnInfo []CmDNNInformation `yaml:"dnn info"`
	Pm      Pm
}

type CmDNNInformation struct {
	Dnn                  string
	DnnIp                string `yaml:"ip"`
	DnnNameIpRangeString string
	DnnSnssaiUpfIpString string
}

type Pm struct {
	Startmodulecount  bool `yaml:"start module count"`
	Startsessioncount bool `yaml:"start session count"`
}
