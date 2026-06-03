/*
* Copyright(C),2020‐2022
* Author: zoujun
* Date: 12/9/20 3:02 PM
* Description:
 */
package adapter

import (
	logger "lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/oam/cm/yaml"
)

func LoadConfigUPF(confFile string) {
	//load upf config
	yaml.Load(confFile, &configure.CmUpfConf)

	//load version information
	configure.UpfConf.Version = configure.CmUpfConf.Version
	if len(configure.UpfConf.Version.Main) == 0 {
		configure.UpfConf.Version.Main = "0.0.0"
	}

	if len(configure.UpfConf.Version.Patch) == 0 {
		configure.UpfConf.Version.Patch = "999"
	}

	//load logger information
	configure.UpfConf.Logger = configure.CmUpfConf.Logger
	if len(configure.UpfConf.Logger.Level) == 0 {
		configure.UpfConf.Logger.Level = "debug"
	}

	if len(configure.UpfConf.Logger.Path) == 0 {
		configure.UpfConf.Logger.Path = "log/upf.log"
	}

	// load N3 configuration
	configure.UpfConf.N3 = configure.CmUpfConf.N3
	if len(configure.UpfConf.N3.Ipv4) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.ERROR, nil, "Failed to load N3 ip from config file, set to default.")
		configure.UpfConf.N3.Ipv4 = "0.0.0.0"
	}

	if len(configure.UpfConf.N3.Mask) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.INFO, nil, "Failed to load N3 ip mask from config file, set to default.")
		configure.UpfConf.N3.Mask = "255.255.255.255"
	}

	// load N4 configuration
	configure.UpfConf.N4 = configure.CmUpfConf.N4
	if len(configure.UpfConf.N4.Local.Ipv4) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.ERROR, nil, "Failed to load N4 local ip from config file, set to default.")
		configure.UpfConf.N4.Local.Ipv4 = "127.0.0.1"
	}

	if configure.UpfConf.N4.Local.Port == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.ERROR, nil, "Failed to load N4 local port from config file, set to default.")
		configure.UpfConf.N4.Local.Port = 8805
	}

	// N4 smf configuration
	if len(configure.UpfConf.N4.Smf.Ipv4) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.ERROR, nil, "Failed to load N4 SMF ip from config file, set to default.")
		configure.UpfConf.N4.Smf.Ipv4 = "0.0.0.0"
	}

	if configure.UpfConf.N4.Smf.Port == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.ERROR, nil, "Failed to load N4 SMF port from config file, set to default.")
		configure.UpfConf.N4.Smf.Port = 8805
	}

	// dnn n6 config
	configure.UpfConf.N6 = configure.CmUpfConf.N6
	if len(configure.UpfConf.N6.Ipv4) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.ERROR, nil, "Failed to load N6 ip from config file, set to default.")
		configure.UpfConf.N6.Ipv4 = "0.0.0.0"
	}

	if len(configure.UpfConf.N6.Mask) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.INFO, nil, "Failed to load N6 ip mask from config file, set to default.")
		configure.UpfConf.N6.Mask = "255.255.255.255"
	}

	// nff config
	configure.UpfConf.Nff = configure.CmUpfConf.Nff
	if len(configure.UpfConf.Nff.DpdkArgs) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.INFO, nil, "Failed to load nff dpdk args from config file, set to default.")
		configure.UpfConf.Nff.DpdkArgs = "--log-level=7"
	}

	if len(configure.UpfConf.Nff.StatsServerAddress) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.INFO, nil,
			"Failed to load nff stats server address from config file, set to default.")
		configure.UpfConf.Nff.StatsServerAddress = "0.0.0.0"
	}
	if (configure.UpfConf.Nff.StatsServerPort) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.INFO, nil,
			"Failed to load nff stats server port from config file, set to default.")
		configure.UpfConf.Nff.StatsServerPort = 8080
	}

	// dnn_name_gw_ip_map
	// init config
	configure.UpfConf.DnnInfo = configure.CmUpfConf.DnnInfo
	configure.UpfConf.DnnNameGwIpMap = make(map[string]string)
	if len(configure.UpfConf.DnnInfo) == 0 {
		logger.Trace(types.ModuleUpfAdapter, logger.ERROR, nil, "Failed to load N6 dnn gateway ip from config file, set to default.")
		dnn := configure.CmDNNInformation{Dnn: "cmnet", DnnIp: "0.0.0.0"}
		configure.UpfConf.DnnInfo = append(configure.UpfConf.DnnInfo, dnn)
	}
	loadDnnListToMap(configure.UpfConf.DnnInfo, configure.UpfConf.DnnNameGwIpMap)

	//PM
	configure.UpfConf.Pm = configure.CmUpfConf.Pm

	return
}

func loadDnnListToMap(dnnList []configure.CmDNNInformation, ipMap map[string]string) {
	// cmnet,172.16.1.200
	for _, v := range dnnList {
		ipMap[v.Dnn] = v.DnnIp
	}
}
