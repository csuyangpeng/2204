package nasie

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

type QoSRules struct {
	QoSRules []QoSRule
}

func (p QoSRules) String() string {
	var rt string
	for i, v := range p.QoSRules {
		rt = rt + fmt.Sprintf("\nQosRule{%d : %s}", i+1, v)
	}
	rt = rt + fmt.Sprintf("\n")
	return rt
}

const (
	MinNumOfPF  = 0
	MaxNumOfPF  = 15
	MinNumOfQFI = 0
	MaxNumOfQFI = 63
)

type QoSRule struct {
	QoSRuleID             byte
	RuleOprCode           RuleOperationCode
	DefaultDQR            bool
	NumberOfPacketFilters byte // 0~15
	PacketFilterLists     PacketFilterLists
	QoSRulePrecedence     byte
	QoSFlowIdentifier     byte // 0~63
	Segregation           bool
}

func (p QoSRule) String() string {
	return fmt.Sprintf(
		"QosRuleId(%v),"+
			"RuleOprCode(%v),"+
			"DefaultDQR(%v),"+
			"NumberOfPacketFilters(%v),"+
			"QosRulePrecedence(%v),"+
			"QFI(%v),"+
			"Segregation(%v),"+
			"PacketFilterLists(%v),",
		p.QoSRuleID,
		p.RuleOprCode,
		p.DefaultDQR,
		p.NumberOfPacketFilters,
		p.QoSRulePrecedence,
		p.QoSFlowIdentifier,
		p.Segregation,
		p.PacketFilterLists)
}

//Rule operation code (bits 8 to 6 of octet 7)
//Bits
//8 7 6
//0 0 0	Reserved
//0 0 1	Create new QoS rule
//0 1 0	Delete existing QoS rule
//0 1 1	Modify existing QoS rule and add packet filters
//1 0 0	Modify existing QoS rule and replace all packet filters
//1 0 1	Modify existing QoS rule and delete packet filters
//1 1 0	Modify existing QoS rule without modifying packet filters
//1 1 1	Reserved
type RuleOperationCode byte

const (
	Reserved1                                          RuleOperationCode = 0
	CreateNewQoSRule                                   RuleOperationCode = 1
	DeleteExistingQoSRule                              RuleOperationCode = 2
	ModifyExistingQoSRuleAndAddPacketFilters           RuleOperationCode = 3
	ModifyExistingQoSRuleAndReplaceAllPacketFilters    RuleOperationCode = 4
	ModifyExistingQoSRuleAndDeletePacketFilters        RuleOperationCode = 5
	ModifyExistingQoSRuleWithoutModifyingPacketFilters RuleOperationCode = 6
	Reserved2                                          RuleOperationCode = 7
)

func (p *RuleOperationCode) StoreWithString(val string) error {
	switch val {
	case "create new qos rule":
		*p = CreateNewQoSRule
	case "delete existing qos rule":
		*p = DeleteExistingQoSRule
	case "modify existing qos rule and add packet filters":
		*p = ModifyExistingQoSRuleAndAddPacketFilters
	case "modify existing qos rule and replace all packet filters":
		*p = ModifyExistingQoSRuleAndReplaceAllPacketFilters
	case "modify existing qos rule and delete packet filters":
		*p = ModifyExistingQoSRuleAndDeletePacketFilters
	case "modify existing qos rule without modifying packet filters":
		*p = ModifyExistingQoSRuleWithoutModifyingPacketFilters
	default:
		return fmt.Errorf("invalid rule opr code(%s)", val)
	}
	return nil
}

//encode QoSRule to nas octet stream
func (p *QoSRule) Encode() ([]byte, error) {
	rlogger.FuncEntry(types.ModCmn, nil)

	var encBuf []byte

	// not encode the TV yet, octet 4 and 5-6 pass

	//octet 7
	octet7 := byte(p.RuleOprCode) << 5
	octet7 |= utils.BoolToByte(p.DefaultDQR) << 4
	octet7 |= byte(p.NumberOfPacketFilters)

	encBuf = append(encBuf, octet7)

	//Packet filter list
	switch p.RuleOprCode {
	case ModifyExistingQoSRuleAndDeletePacketFilters:
		for i := 0; i < len(p.PacketFilterLists.PFList); i++ {
			encBuf = append(encBuf, byte(p.PacketFilterLists.PFList[i].PktFilterIdentifier))
		}
	case CreateNewQoSRule, ModifyExistingQoSRuleAndAddPacketFilters, ModifyExistingQoSRuleAndReplaceAllPacketFilters:
		for i := 0; i < len(p.PacketFilterLists.PFList); i++ {
			encBuf = append(encBuf, p.PacketFilterLists.PFList[i].Encode()...)
		}
	}
	// QoSRulePrecedence
	encBuf = append(encBuf, p.QoSRulePrecedence)

	// QoSFlowIdentifier && Segregation
	// 24.501 Figure 9.11.4.13.2
	//bit8(Spare) |	bit7(Segregation) | bit6-1	(QoS flow identifier)
	QosAndSegByte := utils.BoolToByte(p.Segregation) << 6
	QosAndSegByte |= (p.QoSFlowIdentifier & 0x3F)
	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "QFI_Seg(%x)", QosAndSegByte)

	encBuf = append(encBuf, QosAndSegByte)

	return encBuf, nil
}

func (p *QoSRule) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	//octet 7
	octet7, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "fail to read byte")
		return fmt.Errorf("fail to read byte")
	}
	//fmt.Println("octet7",octet7 )

	p.NumberOfPacketFilters, _ = utils.GetBitsValue(octet7, 1, 4)
	p.DefaultDQR, _ = utils.GetBitValue(octet7, 5)
	code, _ := utils.GetBitsValue(octet7, 6, 8)
	p.RuleOprCode = RuleOperationCode(code >> 5)

	//fmt.Println("p.NumberOfPacketFilters ",p.NumberOfPacketFilters )
	//fmt.Println("p.DefaultDQR ",p.DefaultDQR )
	//fmt.Println("p.RuleOprCode ",p.RuleOprCode )
	//octet 8*-m*
	//Packet filter list
	p.PacketFilterLists.PFList = []PacketFilterList{}
	switch p.RuleOprCode {
	case ModifyExistingQoSRuleAndDeletePacketFilters:
		for i := 0; i < int(p.NumberOfPacketFilters); i++ {
			flist := PacketFilterList{}
			id, err := msgBuf.ReadByte()
			if err != nil {
				rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "fail to read byte")
				return fmt.Errorf("fail to read byte")
			}
			flist.PktFilterIdentifier, _ = utils.GetBitsValue(id, 1, 4)
			p.PacketFilterLists.PFList = append(p.PacketFilterLists.PFList, flist)
		}
	case CreateNewQoSRule, ModifyExistingQoSRuleAndAddPacketFilters, ModifyExistingQoSRuleAndReplaceAllPacketFilters:
		for i := 0; i < int(p.NumberOfPacketFilters); i++ {
			flist := PacketFilterList{}
			flist.Decode(msgBuf)
			p.PacketFilterLists.PFList = append(p.PacketFilterLists.PFList, flist)
		}
	case DeleteExistingQoSRule:
		rlogger.Trace(types.ModCmn, rlogger.INFO, nil, "no more bytes")
		return nil
	}
	//fmt.Println("p.PacketFilterLists.PFList ",p.PacketFilterLists.PFList)

	//octec m+1
	p.QoSRulePrecedence, _ = msgBuf.ReadByte()
	//fmt.Println("p.QoSRulePrecedence ",p.QoSRulePrecedence )

	//octec m+2
	octetm2, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModCmn, rlogger.ERROR, nil, "fail to read byte")
		return fmt.Errorf("fail to read byte")
	}
	//fmt.Println("octetm2",octetm2)

	p.QoSFlowIdentifier, _ = utils.GetBitsValue(octetm2, 1, 6)
	p.Segregation, _ = utils.GetBitValue(octetm2, 7)
	//fmt.Println("p.QoSFlowIdentifier ",p.QoSFlowIdentifier )
	//fmt.Println("p.Segregation ",p.Segregation )

	return nil
}
