// use to generator template
package yaml

import (
	"fmt"
	"lite5gc/cmn/types/configure"
)

type Study struct {
	CourseName string `yaml:"CourseName"`
	Score      int    `yaml:"Score"`
}

type Student struct {
	Name      string  `yaml:"name"`
	Address   string  `yaml:"addr"`
	ScoreList []Study `yaml:"ScoreList"`
}

func ExampleDump() {
	//slice := unsafeheader.Slice{
	//	Data: nil,
	//	Len:  0,
	//	Cap:  0,
	//}
	//fmt.Println(slice)
	src := "template/in/student.yaml"

	var newStu Student
	Dump(newStu, src)
	fmt.Printf("%+v\n", newStu)

	// Output: {Name:George Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:数学 Score:22}]}
}

func ExampleLoad() {
	src := "template/in/student.yaml"
	des := "template/out/student.yaml"

	var newStu Student
	Load(src, &newStu)

	newStu.Name = "lucy"
	study := Study{}
	newStu.ScoreList = append(newStu.ScoreList, study)
	newStu.ScoreList = append(newStu.ScoreList, study)
	newStu.ScoreList[1].CourseName = "英语"
	//newStu.ScoreList[1].s = 23
	Dump(newStu, des)

	var newStuExpect Student
	Load(des, &newStuExpect)
	fmt.Printf("%+v\n", newStuExpect)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

//var (
//	CmnConf CmnConfigData
//	SysConf SystemConfig
//	AmfConf AmfConfig
//	SmfConf SmfConfig
//	UpfConf UpfConfig
//	UdmConf UdmConfig
//)
func ExampleDump_CmCmnConfig() {
	//des := "template/in/cm_cmn_conf.yaml"
	//var config configure.CmCmnConfig
	//Dump(config, des)
	//fmt.Printf("%+v\n", config)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

func ExampleLoad_CmCmnConfig() {
	//src := "template/in/cm_cmn_conf.yaml"
	//des := "template/out/cm_cmn_conf.yaml"
	//
	//var config configure.CmCmnConfig
	//Load(src, &config)
	//config.PlmnList = append(config.PlmnList, "")
	//config.TaiList = append(config.TaiList, "")
	//var nssai configure.CmSNssai
	//config.Nssai = append(config.Nssai, nssai)
	//
	//fmt.Printf("%+v\n", config)
	//Dump(config, des)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

func ExampleDump_SystemConfig() {
	des := "template/in/cm_sys_conf.yaml"
	var conf configure.SystemConfig
	Dump(conf, des)
	fmt.Printf("%+v\n", conf)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

func ExampleLoad_SystemConfig() {
	src := "template/in/cm_sys_conf.yaml"
	des := "template/out/cm_sys_conf.yaml"

	var conf configure.SystemConfig
	Load(src, &conf)

	// todo revise
	Dump(conf, des)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

func ExampleDump_CmAmfConf() {
	des := "template/in/cm_amf_conf.yaml"
	var conf configure.CmAmfConfig
	Dump(conf, des)
	fmt.Printf("%+v\n", conf)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

func ExampleLoad_CmAmfConf() {
	src := "template/in/cm_amf_conf.yaml"
	des := "template/out/cm_amf_conf.yaml"

	var conf configure.CmAmfConfig
	Load(src, &conf)
	fmt.Printf("%+v\n", conf)

	// todo revise
	Dump(conf, des)
	// Output:
	// {N2:{Ipv4: Port:0} Service:AmfService Info:
	//AmfInstanceId:
	//AmfName:
	//AmfIdentifer:  000840
	//AmfRelCap:  0
	//  NAS:{SecEnabled:false SecCap:sddedsdf T3512min:0 T3513Sec:0 T3502min:0 T3550sec:0 T3560sec:0 T3570sec:0 T3522sec:0 T3555sec:22 T3565sec:0 TICSsec:0}}
}

func ExampleDump_CmSmfConf() {
	des := "template/cm_smf_conf.yaml"
	var conf configure.CmSmfConfig
	Dump(conf, des)
	fmt.Printf("%+v\n", conf)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

func ExampleLoad_CmSmfConf() {
	src := "template/in/cm_smf_conf.yaml"
	des := "template/out/cm_smf_conf.yaml"

	var conf configure.CmSmfConfig
	Load(src, &conf)

	var rule configure.CmQoSRule
	conf.Rules = append(conf.Rules, rule)
	var filter configure.CmPacketFilterList
	conf.Rules[0].PacketFilterLists = append(conf.Rules[0].PacketFilterLists, filter)
	var desc string
	conf.Rules[0].PacketFilterLists[0].Descriptions = append(conf.Rules[0].PacketFilterLists[0].Descriptions, desc)

	var ruleD configure.CmQoSFlowDesc
	conf.FlowDesc = append(conf.FlowDesc, ruleD)
	var paramIe configure.CmParameters
	conf.FlowDesc[0].ParameterList = append(conf.FlowDesc[0].ParameterList, paramIe)

	var dnn configure.CmDnnInfo
	conf.Dnn = append(conf.Dnn, dnn)

	var u configure.CmUpfSelection
	conf.UpfSel = append(conf.UpfSel, u)

	Dump(conf, des)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

func ExampleDump_UpfConfig() {
	des := "template/in/cm_upf_conf.yaml"
	var conf configure.CmUpfConfig
	Dump(conf, des)
	fmt.Printf("%+v\n", conf)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

func ExampleLoad_UpfConfig() {
	src := "template/in/cm_upf_conf.yaml"
	des := "template/out/cm_upf_conf.yaml"

	var conf configure.CmUpfConfig
	Load(src, &conf)
	fmt.Printf("%+v\n", conf)

	var dnn configure.CmDNNInformation
	conf.DnnInfo = append(conf.DnnInfo, dnn)


	// todo revise
	Dump(conf, des)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}
func ExampleDump_UdmConfig() {
	des := "template/in/cm_udm_conf.yaml"
	var conf configure.UdmConfig
	Dump(conf, des)
	fmt.Printf("%+v\n", conf)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}

func ExampleLoad_UdmConfig() {
	src := "template/in/cm_udm_conf.yaml"
	des := "template/out/cm_udm_conf.yaml"

	var conf configure.UdmConfig
	Load(src, &conf)

	// todo revise
	Dump(conf, des)
	// Output: {Name:lucy Address:北京 ScoreList:[{CourseName:语文 Score:21} {CourseName:英语 Score:23}]}
}
