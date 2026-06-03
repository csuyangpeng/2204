/*
 * Copyright(C),2020-2022
 * Author: Jaytan
 * Date: 11/23/20 9:36 PM
 * Description:
 */
package addsbirouters

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sbicmn"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/udm/sbi/subscriberdatamanagement"
	"lite5gc/udm/sbi/ueauthentication"
	"lite5gc/udm/sbi/uecontextmanagement"
)

type SbiUdm struct {
	sbi sbicmn.SbiInterface
}

func (p *SbiUdm) Initialize() {
	// rlogger.FuncEntry(moduleTag, nil)
	// p.sbi.Initialize(configure.UdmConf.SBI.Addr.Ipv4, configure.UdmConf.SBI.Addr.Port)

	p.sbi.Initialize(configure.UdmConf.Sbi.Udm.Addr.Ip, configure.UdmConf.Sbi.Udm.Addr.Port)

	p.sbi.SetSbiLogger()
	p.AddService()
}

func Start(appContext *types.AppContext) {
	sbiIf := &SbiUdm{}
	sbiIf.Initialize()
	go sbiIf.Run(appContext)
}

func (p *SbiUdm) Run(appContext *types.AppContext) error {
	rlogger.FuncEntry(types.ModuleUdmRouter, nil)
	go p.sbi.Start()
	appContext.Wg.Add(1)
	defer func() {
		appContext.Wg.Done()
	}()
	return nil
}

func (p *SbiUdm) AddService() {
	rlogger.FuncEntry(types.ModuleUdmRouter, nil)
	// add sbi service
	subscriberdatamanagement.AddService(p.sbi.GetRouter())
	UEAuthentication.AddService(p.sbi.GetRouter())
	UEContextManagement.AddService(p.sbi.GetRouter())
}
