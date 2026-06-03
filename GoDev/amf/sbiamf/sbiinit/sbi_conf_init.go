package sbiinit

import (
	"lite5gc/amf/sbiamf/communication"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
)
var SbiAmf sbicmn.SbiInterface

func Initialize() {
	rlogger.FuncEntry(types.ModuleAmfSbi, nil)
	//load config ip and port
	SbiAmf.Initialize(configure.AmfConf.Sbi.Amf.Addr.Ip,configure.AmfConf.Sbi.Amf.Addr.Port)
	SbiAmf.SetSbiLogger()
	//add service
	Communication.AddService(SbiAmf.GetRouter())
	//start listening...
	go SbiAmf.Start()
}
