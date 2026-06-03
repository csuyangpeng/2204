package sbiInit

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/smf/sbismf/pdusession"
)

var SbiSmf sbicmn.SbiInterface

func Initialize() {
	rlogger.FuncEntry(types.ModuleSmfSbi, nil)
	//load config ip and port
	SbiSmf.Initialize(configure.SmfConf.Sbi.Smf.Addr.Ip, configure.SmfConf.Sbi.Smf.Addr.Port)
	SbiSmf.SetSbiLogger()
	//add service
	pdusession.AddService(SbiSmf.GetRouter())
	//start listening...
	go SbiSmf.Start()
}
