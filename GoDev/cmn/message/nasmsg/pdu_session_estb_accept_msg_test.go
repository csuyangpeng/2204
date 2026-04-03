package nasmsg

import (
	"fmt"
	"lite5gc/cmn/nas/nasie"
	"testing"
)

func TestPduSessionEstbAcceptMsg_Encode(t *testing.T) {
	msg := PduSessionEstbAcceptMsg{}
	msg.SessionType.StoreWithString("IPV4")
	msg.SessionAMBR.StoreWithString("256 kbps", "256 kbps")
	msg.SNSSAI.Sst = 1
	msg.DNN.StoreWithString("cmnet")
	qosRule := nasie.QoSRule{}
	qosRule.QoSFlowIdentifier = 1
	qosRule.Segregation = false
	qosRule.QoSRulePrecedence = 1
	qosRule.NumberOfPacketFilters = 1
	qosRule.RuleOprCode = 1
	qosRule.DefaultDQR = true
	qosRule.QoSRuleID = 1

	pfl := nasie.PacketFilterList{}
	pfl.PktFilterIdentifier = 1
	pfl.PktFilterDirection = 1
	pfl.PacketFilterContents.PacketFilterContentID[0] = 1
	qosRule.PacketFilterLists.PFList = append(qosRule.PacketFilterLists.PFList, pfl)

	msg.AuthorizedQoSRules.QoSRules = append(msg.AuthorizedQoSRules.QoSRules, qosRule)

	buf, err := msg.Encode()
	if err != nil {
		fmt.Println(err)
	}
	//out := fmt.Sprintf("%x", buf)
	//expect := "7e0056000200002147114711471147114711471147114711201011495219c670b9b9d8b9676b45dad5fc"
	//
	fmt.Printf("Encode Msg: %x\n", buf)
	//fmt.Printf("Expect msg: %s\n", expect)
	//
	//if out != expect {
	//	t.Errorf("failed to encode.")
	//}
}
