package cm

import (
	"fmt"
	"lite5gc/cmn/types/configure"
)

//func ExampleLoadCmnConfig() {
//	LoadCmnConfig()
//	fmt.Printf("%+v\n", configure.CmCmnConf)
//	// Output:
//	// {PlmnList:[] TaiList:[] Nssai:[{Sst:0 Sd:[0 0 0]}] DnnInfo:{DnnIp: DnnNameIpRangeString: DnnSnssaiUpfIpString:} Sbi:{AmfSBI:{Addr:{Ipv4: Port:0} Scheme:} UdmSBI:{Addr:{Ipv4: Port:0} Scheme:} SmfSBI:{Addr:{Ipv4: Port:0} Scheme:}} Logger:{AmfLogger:{LogLevel: LogPath: LogFile:false LogConsole:false} SmfLogger:{LogLevel: LogPath: LogFile:false LogConsole:false} UdmLogger:{LogLevel: LogPath: LogFile:false LogConsole:false}} Version:{AmfV:{MainVer: PatchVer:} SmfV:{MainVer: PatchVer:} UdmV:{MainVer: PatchVer:}}}
//}

func ExampleLoadSysConfig() {
	LoadSysConfig()
	fmt.Printf("%+v\n", configure.SysConf)
	// Output:
	// {N2SctpParam:{Initmsg:{NumOstreams:0 MaxInstreams:0 MaxAttempts:0 MaxInitTimeout:0} SctpHeatbeat:{HeatbeatInterval:0 PathMaxRXT:0}} SC:{NumOfAmfScInst:0 NumOfSmfScInst:0}}
}

func ExampleLoadAmfConfig() {
	LoadAmfConfig()
	fmt.Printf("%+v\n", configure.CmAmfConf)
	// Output:
	// {N2:{Ipv4:11.11.11.11 Port:0} Service:AmfService Info:
	// AmfInstanceId:  amfinstanceid
	// AmfName:  amfname
	// AmfIdentifer:  000000
	// AmfRelCap:  0
	//  NAS:{SecEnabled:false EncAlgo: IntAlgo: T3512min:0 T3513Sec:0 T3502min:0 T3550sec:0 T3560sec:0 T3570sec:0 T3522sec:0 T3555sec:0 T3565sec:0 TICSsec:2}}
}

func ExampleLoadSmfConfig() {
	LoadSmfConfig()
	fmt.Printf("%+v\n", configure.CmSmfConf)
	// Output:
	// {Service:{SmfInstanceId: SSCMode:Unknown SessType:unknown} Rules:[{QoSRuleID:0 RuleOprCode:0 DefaultDQR:false NumberOfPacketFilters:0 QoSRulePrecedence:0 FlowDescr:{QFI:0 OperationCode:0 E:false NumberOfParameters:0 ParameterList:[{ParameterID:0 QI5Content:0 GFBRUplinkContent:{Value:0 Uint:0} GFBRDownlinkContent:{Value:0 Uint:0} MFBRUplinkContent:{Value:0 Uint:0} MFBRDownlinkContent:{Value:0 Uint:0} AveragingWindowContent:0}]} Segregation:false}] N4Conf:{SMFIP:<nil> SMFPort:0} PFCPParameters:PDRs0:{QFI:0 Precedence:0 OuterHeaderRemoval:0 FarI:{Action:0 DstInterface:0}},BAR:{%!s(uint8=0) %!s(uint8=0)}, Timer:{Timer3592:0 Timer3593:0 Timer3590:0 Timer3591:0}}
}

func ExampleLoadUpfConfig() {
	LoadUpfConfig()
	fmt.Printf("%+v\n", configure.UpfConf)
	// Output:
	// {N3:{Ipv4: Port:0} N4:{Local:{Ipv4: Port:0} Smf:{Ipv4: Port:0}} N6:}
}

func ExampleLoadUdmConfig() {
	LoadUdmConfig()
	fmt.Printf("%+v\n", configure.UdmConf)
	// Output:
	// {UdmInstanceId: Version:{Main:0.1 Patch:0} Logger:{Level:debug Path:. Control:false} Sbi:{Amf:{Addr:{Ip: Port:0} Scheme:} Udm:{Addr:{Ip:10.18.1.56 Port:29503} Scheme:http} Smf:{Addr:{Ip: Port:0} Scheme:}}}

}
