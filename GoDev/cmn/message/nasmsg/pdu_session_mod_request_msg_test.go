package nasmsg

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"net"
	"testing"
)

func TestPduSessionModReqMsg_Encode(t *testing.T) {
	msg := PduSessionModifyRequestMsg{}

	msg.MsgHeader.PrcdTransactionID = 1
	msg.MsgHeader.PduSessionID = 5
	msg.MsgHeader.MessageType = nas.PduSessionModRequest

	msg.SmCapability.MPTCP = true
	msg.SmCapability.RqoS = true
	msg.SmCapability.ATSLL = true
	msg.SmCapability.MH6PDU = true
	msg.SmCapability.Spare = []byte{2, 4}
	msg.SmCapability.EPTS1 = true
	msg.IeFlags.Set(Ieid_PduSessionModReq_SmCapability)

	msg.SMCause = nas.SuccessNoReason
	msg.IeFlags.Set(Ieid_PduSessionModReq_SMCause)

	msg.IntMaxDataRate.MaxRateUpLink = nas.FullDataRate
	msg.IntMaxDataRate.MaxRateDownLink = nas.KBPS_64
	msg.IeFlags.Set(Ieid_PduSessionModReq_IntMaxDataRate)

	msg.AlwaysOnPduSessReq = true
	msg.IeFlags.Set(Ieid_PduSessionModReq_AlwaysOnPduSessReq)

	msg.MaxNumOfSupPckFilter = 100
	msg.IeFlags.Set(Ieid_PduSessionModReq_MaxNumOfSupPckFilter)

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
	pfl.PacketFilterContents.PacketFilterContentID = make([]nasie.PacketFilterContentIdentifier, 2)
	pfl.PacketFilterContents.PacketFilterContentID[0] = nasie.IPv4RemoteAddressType
	_, ipnet, _ := net.ParseCIDR("192.0.2.1/24")
	pfl.PacketFilterContents.IPV4RemoteAddr = *ipnet

	pfl.PacketFilterContents.PacketFilterContentID[1] = nasie.ProtocolIdentifierNextHeaderType
	pfl.PacketFilterContents.V4ProtocolIdOrV6NextHeader = 1
	fmt.Println(pfl.PacketFilterContents)

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
	//fmt.Printf("Encode Msg: %x\n", buf)
	//fmt.Printf("Expect msg: %s\n", expect)
	//
	//if out != expect {
	//	t.Errorf("failed to encode.")
	//}
	fmt.Printf("Encode Msg: %x\n", buf[3:])
	testBuf := []byte{122, 0, 4, 1, 0, 1, 64, 121, 0, 3, 1, 64, 0}
	//fmt.Printf("Encode Msg: %x\n", buf[3:])
	//bufMsg := bytes.NewReader(buf[3:])
	bufMsg := bytes.NewReader(testBuf)
	msg2 := PduSessionModifyRequestMsg{}
	err = msg2.Decode(bufMsg)
	//fmt.Println(err)
	fmt.Println(msg2.RequestQosRules)
	fmt.Println(msg2.RequestQosFlowDesc)
}
