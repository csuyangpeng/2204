package nasmsg

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"testing"
)

func TestPduSessionModCmdMsg_Encode(t *testing.T) {
	msg := PDUSessionModifyCommandMsg{}

	msg.MsgHeader.PrcdTransactionID = 1
	msg.MsgHeader.PduSessionID = 5
	msg.MsgHeader.MessageType = nas.PduSessionModRequest

	msg.SMCause = nas.SuccessNoReason
	msg.IeFlags.Set(Ieid_PduSessionModCmd_SMCause)

	msg.AlwaysOnPduSessReq = true
	msg.IeFlags.Set(Ieid_PduSessionModCmd_AlwaysOnPduSessReq)

	msg.RQTimer.Uint = nasie.TwoSeconds
	msg.RQTimer.TimerValue = 2
	msg.IeFlags.Set(Ieid_PduSessionModCmd_RQTimer)

	msg.Ambr.UpLink.Value = 1
	msg.Ambr.UpLink.Uint = nasie.Kbps1
	msg.Ambr.DownLink.Value = 1
	msg.Ambr.DownLink.Uint = nasie.Gbps1
	msg.IeFlags.Set(Ieid_PduSessionModCmd_Ambr)

	qosRule := nasie.QoSRule{}
	qosRule.QoSFlowIdentifier = 1
	qosRule.Segregation = false
	qosRule.QoSRulePrecedence = 255
	qosRule.NumberOfPacketFilters = 1
	qosRule.RuleOprCode = 1
	qosRule.DefaultDQR = true
	qosRule.QoSRuleID = 1

	pfl := nasie.PacketFilterList{}
	pfl.PktFilterIdentifier = 1
	pfl.PktFilterDirection = 2
	pfl.PacketFilterContents.PacketFilterContentID[0] = nasie.MatchAlltype

	qosRule.PacketFilterLists.PFList = append(qosRule.PacketFilterLists.PFList, pfl)
	msg.RequestQosRules.QoSRules = append(msg.RequestQosRules.QoSRules, qosRule)
	msg.IeFlags.Set(Ieid_PduSessionModReq_RequestQosRules)

	qosRuleDesc := nasie.QoSFlowDescription{}
	qosRuleDesc.QFI = 1
	qosRuleDesc.OperationCode = nasie.CreateNewQoSFlowDescription
	qosRuleDesc.E = true
	qosRuleDesc.NumberOfParameters = 1
	list := nasie.ParametersIE{}
	list.ParameterID = nasie.FiveQI
	list.QI5Content = nasie.QI9
	qosRuleDesc.ParameterList.ParmsList = append(qosRuleDesc.ParameterList.ParmsList, list)
	msg.RequestQosFlowDesc.Descr = append(msg.RequestQosFlowDesc.Descr, qosRuleDesc)
	msg.IeFlags.Set(Ieid_PduSessionModReq_RequestQosFlowDesc)

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
	fmt.Printf("Encode Msg: %x\n", buf[3:])
	bufMsg := bytes.NewReader(buf[3:])
	msg2 := PDUSessionModifyCommandMsg{}
	err = msg2.Decode(bufMsg)
	//fmt.Println(err)
	fmt.Println(msg2)
}
