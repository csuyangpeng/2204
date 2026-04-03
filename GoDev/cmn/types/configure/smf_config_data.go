package configure

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"net"
)

//SMF configuration
type SmfConfig struct {
	Version      VerConfig
	Logger       LoggerConfig
	Service      SmfService
	DnnInfo      []CmDnnInfo
	UpfSelection []CmUpfSelection
	N4Conf       SmfN4Config
	Bar          CmBAR
	Rules        nasie.QoSRules
	FlowDescr    nasie.QoSFlowsDesc
	Sbi          SBI
	Timer        NasTimer
}

func (p SmfConfig) String() string {
	return fmt.Sprintf(
		"Version(%s),\nLogger(%v),\nService(%v),\nDnnInfo(%s),\n"+
			"UpfSelection(%s),\nN4Conf(%v),\nBar(%v),\nRules(%v),\n"+
			"FlowDescr(%v),\nSbi(%v),\nTimer(%d)\n",
		p.Version, p.Logger, p.Service, p.DnnInfo, p.UpfSelection,
		p.N4Conf, p.Bar, p.Rules, p.FlowDescr, p.Sbi, p.Timer)
}

type SmfService struct {
	InstanceId string
	SSCMode    nas.SSCMode
	SessType   types3gpp.PduSessType
}

type SmfN4Config struct {
	SMFIP   net.IP
	SMFPort int
}

type NasTimer struct {
	Timer3592 int `yaml:"t3590 sec"`
	Timer3593 int `yaml:"t3591 sec"`
	Timer3590 int `yaml:"t3592 sec"`
	Timer3591 int `yaml:"t3593 sec"`
}

func (p *SmfConfig) ConvertSmfConfig(cmData *CmSmfConfig) error {
	rlogger.FuncEntry(types.ModuleCmn3gtp, nil)
	// check the input parameter
	if cmData == nil {
		return types.ErrInputParaNil
	}

	// version
	p.Version = cmData.Version

	// logger
	p.Logger = cmData.Logger

	// SBI
	p.Sbi = cmData.Sbi

	// nas
	p.Timer = cmData.Timer

	p.DnnInfo = cmData.Dnn
	p.UpfSelection = cmData.UpfSel
	p.Bar = cmData.BAR

	// n4
	if cmData.N4.IP == "" || cmData.N4.Port == 0 {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "n4 is empty")
		return types.ErrInputParaNil
	} else {
		p.N4Conf.SMFIP = net.ParseIP(cmData.N4.IP)
		p.N4Conf.SMFPort = cmData.N4.Port
	}

	// smf service
	u := uuid.NewV4()
	p.Service.InstanceId = u.String()
	if cmData.Service.SessionType == "" {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
			"config item is empty, set to default(ipv4)")
		p.Service.SessType = types3gpp.Ipv4
	} else {
		err := p.Service.SessType.StoreWithString(cmData.Service.SessionType)
		if err != nil {
			rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
				"Invalid config item, set to default(ipv4)")
			p.Service.SessType = types3gpp.Ipv4
		}
	}
	if cmData.Service.SSCMode == "" {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
			"config item is empty, set to default(ssc mode 1)")
		p.Service.SSCMode = nas.SSCMode1
	} else {
		err := p.Service.SSCMode.StoreWithString(cmData.Service.SSCMode)
		if err != nil {
			rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
				"Invalid config item, set to default(ssc mode 1)")
			p.Service.SSCMode = nas.SSCMode1
		}
	}

	if len(cmData.Rules) != 0 {
		for i := 0; i < len(cmData.Rules); i++ {
			cmRule := cmData.Rules[i]
			rule := nasie.QoSRule{}
			rule.QoSRuleID = cmRule.RuleID
			err := rule.RuleOprCode.StoreWithString(cmRule.OperationCode)
			if err != nil {
				rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
					"Invalid config item, set to default(create new qos rule)")
				rule.RuleOprCode = nasie.CreateNewQoSRule
			}
			rule.DefaultDQR = cmRule.IsDefaultDQR
			rule.QoSRulePrecedence = cmRule.Precedence
			rule.QoSFlowIdentifier = cmRule.QFI
			rule.Segregation = cmRule.Segregation
			rule.NumberOfPacketFilters = byte(len(cmRule.PacketFilterLists))
			if len(cmRule.PacketFilterLists) != 0 {
				for i := 0; i < len(cmRule.PacketFilterLists); i++ {
					var filter nasie.PacketFilterList
					filter.PktFilterIdentifier = cmRule.PacketFilterLists[i].Id
					err = filter.PktFilterDirection.StoreWithString(cmRule.PacketFilterLists[i].Direction)
					if err != nil {
						rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
							"Invalid config item, set to default(Bidirectional)")
						filter.PktFilterDirection = nasie.Bidirectional
					}
					if cmRule.PacketFilterLists[i].Descriptions[0] == "permit out ip from any to assigned" {
						filter.PacketFilterContents.PacketFilterContentID = append(
							filter.PacketFilterContents.PacketFilterContentID, nasie.MatchAlltype)
					} else {
						// 解析 todo
					}
					// append
					rule.PacketFilterLists.PFList = append(rule.PacketFilterLists.PFList, filter)
				}
			} else {
				rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "packet filter lists config is empty")
			}
			// append
			p.Rules.QoSRules = append(p.Rules.QoSRules, rule)
		}
	} else {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "rule config is empty")
	}

	if len(cmData.FlowDesc) != 0 {
		for i := 0; i < len(cmData.FlowDesc); i++ {
			cmFlowDesc := cmData.FlowDesc[i]
			flowDesc := nasie.QoSFlowDescription{}
			flowDesc.QFI = cmFlowDesc.QFI
			err := flowDesc.OperationCode.StoreWithString(cmFlowDesc.OperationCode)
			if err != nil {
				rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
					"Invalid config item, set to default(CreateNewQoSFlowDescription)")
				flowDesc.OperationCode = nasie.CreateNewQoSFlowDescription
			}
			flowDesc.E = cmFlowDesc.E
			flowDesc.NumberOfParameters = byte(len(cmFlowDesc.ParameterList))
			if len(cmFlowDesc.ParameterList) != 0 {
				for i := 0; i < len(cmFlowDesc.ParameterList); i++ {
					var p nasie.ParametersIE
					err = p.ParameterID.StoreWithString(cmFlowDesc.ParameterList[i].Id)
					if err != nil {
						rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
							"Invalid config item, set to default(FiveQI 9)")
						p.ParameterID = nasie.FiveQI
						p.QI5Content = 9
					}
					switch p.ParameterID {
					case nasie.FiveQI:
						err = p.QI5Content.StoreWithInt(int(cmFlowDesc.ParameterList[i].FiveQI))
						if err != nil {
							rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
								"Invalid config item, set to default(FiveQI 9)")
							p.QI5Content = 9
						}
					case nasie.MFBRUplink:
						err = p.MFBRUplinkContent.StoreWithString(cmFlowDesc.ParameterList[i].MFBRUplink)
						if err != nil {
							rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
								"Invalid config item, set to default(1 Gbps)")
							p.MFBRUplinkContent.Value = 1
							p.MFBRUplinkContent.Uint = nasie.Gbps1
						}
					case nasie.MFBRDownlink:
						err = p.MFBRDownlinkContent.StoreWithString(cmFlowDesc.ParameterList[i].MFBRDownlink)
						if err != nil {
							rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
								"Invalid config item, set to default(1 Gbps)")
							p.MFBRDownlinkContent.Value = 1
							p.MFBRDownlinkContent.Uint = nasie.Gbps1
						}
					case nasie.GFBRUplink:
						err =	p.GFBRUplinkContent.StoreWithString(cmFlowDesc.ParameterList[i].GFBRUplink)
						if err != nil {
							rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
								"Invalid config item, set to default(1 Gbps)")
							p.GFBRUplinkContent.Value = 1
							p.GFBRUplinkContent.Uint = nasie.Gbps1
						}
					case nasie.GFBRDownlink:
						err =p.GFBRDownlinkContent.StoreWithString(cmFlowDesc.ParameterList[i].GFBRDownlink)
						if err != nil {
							rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
								"Invalid config item, set to default(1 Gbps)")
							p.GFBRDownlinkContent.Value = 1
							p.GFBRDownlinkContent.Uint = nasie.Gbps1
						}
					}
					flowDesc.ParameterList.ParmsList = append(flowDesc.ParameterList.ParmsList,p)
				}
			} else {
				rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil,
					"flow desc param list config is empty")
			}
			// append
			p.FlowDescr.Descr = append(p.FlowDescr.Descr, flowDesc)
		}
	} else {
		rlogger.Trace(types.ModuleCmn3gtp, rlogger.ERROR, nil, "flow desc config is empty")
	}

	return nil
}
