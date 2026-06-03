package configure

import (
	"net"
)

type CmSmfConfig struct {
	Version  VerConfig
	Logger   LoggerConfig
	Service  CmSmfService
	Dnn      []CmDnnInfo      `yaml:"dnn info"`
	UpfSel   []CmUpfSelection `yaml:"upf selection"`
	N4       CmSmfN4Config
	BAR      CmBAR
	Rules    []CmQoSRule `yaml:"qos rules"`
	FlowDesc []CmQoSFlowDesc `yaml:"flow descriptions"`
	Sbi      SBI
	Timer    NasTimer
}
type CmDnnInfo struct {
	Name    string
	Ip      string
	IpRange string `yaml:"ip range"`
}
type CmUpfSelection struct {
	DnnName string `yaml:"dnn name"`
	Snssai  string
	UpfIp   string `yaml:"upf ip"`
}

type CmBAR struct {
	DLDataNotificationDelay uint8 `yaml:"downlink data notification delay(ms)"`
	SugBuffPacketsCount     uint8 `yaml:"suggestion buffer packets count"`
}

type CmSmfService struct {
	SSCMode     string `yaml:"ssc mode"`
	SessionType string `yaml:"pdu session type"`
}

type CmQoSRule struct {
	RuleID            byte `yaml:"rule id"`
	OperationCode     string `yaml:"operation code"`
	IsDefaultDQR      bool `yaml:"is default rule"`
	Precedence        byte
	QFI               byte // 0~63
	Segregation       bool
	PacketFilterLists []CmPacketFilterList `yaml:"packet filter list"`
}

type CmPacketFilterList struct {
	Id           byte  `yaml:"packet filter id"` // 0~15
	Direction    string
	Descriptions []string
}

type CmQoSFlowDesc struct {
	QFI            byte
	OperationCode  string `yaml:"operation code"`
	E              bool
	ParameterList  []CmParameters `yaml:"parameter list"`
}

type CmParameters struct {
	Id string `yaml:"parameters id"`
	FiveQI byte
	//GBR
	GFBRUplink             string `yaml:"gfbr uplink"`
	GFBRDownlink           string `yaml:"gfbr downlink"`
	MFBRUplink             string `yaml:"mfbr uplink"`
	MFBRDownlink           string `yaml:"mfbr downlink"`
	AveragingWindowContent byte   `yaml:"averaging window(ms)"` //23.501 Table 5.7.4-1 when 5QI=9 AveragingWindowContent=N/A
}

type CmSmfN4Config struct {
	IP   string
	Port int
}

type CmUdmConfig struct {
	UdmInstanceId string
	Version       VerConfig
	Logger        LoggerConfig
	Sbi           SBI
	PlmnList      []string `yaml:"plmn list"`
}

type CmRedisConfig struct {
	ServerIP net.IP
	Port     int
}
