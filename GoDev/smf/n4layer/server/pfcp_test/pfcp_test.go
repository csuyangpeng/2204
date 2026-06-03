package test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/smf/n4layer/server/pfcptst"
	"net"
	"testing"
	"time"
)

//3GPP TS 23.502 V15.4.0 (2018-12)

//4.4.3	N4 Node Level Procedures
//4.4.3.1	N4 Association Setup Procedure
//Message Type: PFCP Association Setup Request (5)
// smf side
// send request:smf --> upf
func TestPfcpMarshal5(t *testing.T) {
	//PFCP Association Setup Request (5)
	ExpectedReqdata, _ := hex.DecodeString("2005001500000100003c0005000aca610200600004e07227a1")
	//请求编码
	req := pfcpv1.Message{}
	request := &pfcp.PFCPAssociationSetupRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Association_Setup_Request,
			Length:         0, // todo 编码后填充
			SequenceNumber: 1},
	}
	request.IE = &pfcp.IEPFCPAssociationSetupRequest{}
	request.IE.NodeID = &pfcp.IENodeID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Node_ID,
			Length: 5,
		},
		NodeIDType:  0,
		NodeIDvalue: []byte{10, 202, 97, 2}, // smf ip
	}

	request.IE.RecoveryTimeStamp = &pfcp.IERecoveryTimeStamp{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Recovery_Time_Stamp,
			Length: 4,
		},
		// 1556588833 = nt64(0xe07227a1(wireshark) - 0x83aa7e80(1900到1970的秒差))
		RecoveryTimeStamp: time.Unix(1556588833, 0),
	}
	fmt.Println("time", request.IE.RecoveryTimeStamp.RecoveryTimeStamp.UTC().String())
	req.Header.Version = request.PfcpHeader.Version
	req.Header.MessageType = pfcp.PFCP_Association_Setup_Request
	req.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	//todo 编码后填充
	req.Header.Length = 0
	req.Body = request
	fmt.Println(req.String())

	// todo pfcp encode
	data, err := req.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal err %s", err)
	}
	fmt.Printf("%x\n", ExpectedReqdata)
	fmt.Printf("%x\n", data)
	if !bytes.Equal(data, ExpectedReqdata) {
		t.Errorf("Pfcp msg marshal failed")
	}
	fmt.Println("------response------")
	// 响应解码
	resdata, _ := hex.DecodeString("2006001a00000100003c0005000aca5e02001300010100600004e07227a1")
	fmt.Printf("resdata:%x\n", resdata)
	msg := pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 6
	//Length: 26
	//SequenceNumber: 1
	if msg.Header.MessageType != 6 || msg.Header.Length != 26 ||
		msg.Header.SequenceNumber != 1 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.PFCPAssociationSetupResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	//10.202.97.2
	//2019-04-30 01:47:13 +0000 UTC
	fmt.Println(net.IP(v.IE.NodeID.NodeIDvalue))
	fmt.Println((v.IE.Cause.CauseValue))
	fmt.Println(v.IE.RecoveryTimeStamp.RecoveryTimeStamp.UTC())

	err = pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}
	// upf ip
	if net.IP(v.IE.NodeID.NodeIDvalue).String() != "10.202.94.2" ||
		v.IE.RecoveryTimeStamp.RecoveryTimeStamp.UTC().String() != "2019-04-30 01:47:13 +0000 UTC" {
		t.Errorf("Pfcp header unmarshal failed")
	}
}

//4.4.3.2	N4 Association Update Procedure
// smf side
// send request:smf --> upf
//Message Type: PFCP Association Update Request (7)
func TestPfcpMarshal7(t *testing.T) {
	//PFCP Association Update Request (7)
	ExpectedReqdata, _ := hex.DecodeString("2007001200000300003c0005000aca5e02006f000101")
	//请求编码
	req := pfcpv1.Message{}
	request := &pfcp.PFCPAssociationUpdateRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Association_Update_Request,
			Length:         0, // todo 编码后填充
			SequenceNumber: 3},
	}
	request.IE = &pfcp.IEPFCPAssociationUpdateRequest{}
	request.IE.NodeID = &pfcp.IENodeID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Node_ID,
			Length: 5,
		},
		NodeIDType:  0,
		NodeIDvalue: []byte{10, 202, 94, 2}, // upf ip
	}

	request.IE.PFCPAssociationReleaseRequest = &pfcp.IEPFCPAssociationReleaseRequest{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_IE_PFCP_Association_Release_Request,
			Length: 1,
		},
		SARR: true,
	}

	request.IE.IeFlags.Set(pfcp.IE_IE_PFCP_Association_Release_Request)

	req.Header.Version = request.PfcpHeader.Version
	req.Header.MessageType = request.PfcpHeader.MessageType
	req.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	//todo 编码后填充
	req.Header.Length = 0
	req.Body = request
	fmt.Println(req.String())

	// todo pfcp encode
	data, err := req.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal err %s", err)
	}
	fmt.Printf("%x\n", ExpectedReqdata)
	fmt.Printf("%x\n", data)
	if !bytes.Equal(data, ExpectedReqdata) {
		t.Errorf("Pfcp msg marshal failed")
	}
	fmt.Println("------response------")
	// 响应解码
	resdata, _ := hex.DecodeString("2008001200000300003c0005000aca61020013000101")
	fmt.Printf("resdata:%x\n", resdata)
	msg := pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 8
	//Length: 18
	//SequenceNumber: 3
	if msg.Header.MessageType != 8 || msg.Header.Length != 18 ||
		msg.Header.SequenceNumber != 3 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.PFCPAssociationUpdateResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	// 10.202.97.2 //smf ip
	fmt.Println(net.IP(v.IE.NodeID.NodeIDvalue))
	fmt.Println((v.IE.Cause.CauseValue))
	//fmt.Println(v.IE.RecoveryTimeStamp.RecoveryTimeStamp.UTC())

	err = pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}
	// upf ip
	if net.IP(v.IE.NodeID.NodeIDvalue).String() != "10.202.97.2" ||
		v.IE.Cause.CauseValue != 1 {
		t.Errorf("Pfcp header unmarshal failed")
	}
}

// smf side
// receive request:smf <-- upf
//Message Type: PFCP Association Update Request (7)
func TestPfcpUnmarshal7(t *testing.T) {
	//PFCP Association Update Request (7)
	data, _ := hex.DecodeString("2007001200000300003c0005000aca5e02006f000101")
	fmt.Printf("%x\n", data)
	msg := pfcpv1.Message{}
	msg.Unmarshal(data)
	//
	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 7
	//Length: 18
	//SequenceNumber: 3
	if msg.Header.MessageType != 7 || msg.Header.Length != 18 ||
		msg.Header.SequenceNumber != 3 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	req, ok := msg.Body.(*pfcp.PFCPAssociationUpdateRequest)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	NodeIDStr := net.IP(req.IE.NodeID.NodeIDvalue).String()
	Sarr := req.IE.PFCPAssociationReleaseRequest.SARR
	fmt.Println(NodeIDStr)
	fmt.Println(Sarr)
	if NodeIDStr != "10.202.94.2" || Sarr != true {
		t.Errorf("Pfcp body unmarshal failed")
	}

	fmt.Println("------response------")

	res := &pfcpv1.Message{}
	// send PFCP Association Update Response(8)
	pfcptst.Dispatch(msg, res)

	data, err := res.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal failed")
	}
	ExpectedData, _ := hex.DecodeString("2008001200000300003c0005000aca61020013000101")

	fmt.Printf("ExpectedData:%x\n", ExpectedData)
	fmt.Printf("result      :%x\n", data)

	if !bytes.Equal(ExpectedData, data) {
		t.Errorf("Pfcp marshal failed")
	}
}

//4.4.3.3	N4 Association Release Procedure
// smf side
// send request:smf --> upf
//Message Type: PFCP Association Release Request (9)
func TestPfcpMarshal9(t *testing.T) {
	//PFCP Association Release Request (9)
	ExpectedReqdata, _ := hex.DecodeString("2009000d00000f00003c0005000aca6102")
	//请求编码
	req := pfcpv1.Message{}
	request := &pfcp.PFCPAssociationReleaseRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Association_Release_Request,
			Length:         0, // 编码后填充
			SequenceNumber: 15},
	}
	request.IE = &pfcp.IEGroupedPFCPAssociationReleaseRequest{}
	request.IE.NodeID = &pfcp.IENodeID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Node_ID,
			Length: 5,
		},
		NodeIDType:  0,
		NodeIDvalue: []byte{10, 202, 97, 2}, // upf ip
	}

	req.Header.Version = request.PfcpHeader.Version
	req.Header.MessageType = request.PfcpHeader.MessageType
	req.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	// 编码后填充
	req.Header.Length = 0
	req.Body = request
	fmt.Println(req.String())

	// pfcp encode
	data, err := req.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal err %s", err)
	}
	fmt.Printf("%x\n", ExpectedReqdata)
	fmt.Printf("%x\n", data)
	if !bytes.Equal(data, ExpectedReqdata) {
		t.Errorf("Pfcp msg marshal failed")
	}
	fmt.Println("------response------")
	//PFCP Association Release Response (10)
	// 响应解码
	resdata, _ := hex.DecodeString("200a001200000f00003c0005000aca5e020013000101")
	fmt.Printf("resdata:%x\n", resdata)
	msg := pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 10
	//Length: 18
	//SequenceNumber: 15
	if msg.Header.MessageType != 10 || msg.Header.Length != 18 ||
		msg.Header.SequenceNumber != 15 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.PFCPAssociationReleaseResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	// 10.202.94.2 //smf ip
	fmt.Println(net.IP(v.IE.NodeID.NodeIDvalue))
	fmt.Println((v.IE.Cause.CauseValue))

	// 响应处理
	err = pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}
	// upf ip
	if net.IP(v.IE.NodeID.NodeIDvalue).String() != "10.202.94.2" ||
		v.IE.Cause.CauseValue != 1 {
		t.Errorf("Pfcp header unmarshal failed")
	}
}

// smf side
// receive request:smf <-- upf
//Message Type: PFCP Association Release Request (9)
func TestPfcpUnmarshal9(t *testing.T) {
	//PFCP Association Release Request (9)
	data, _ := hex.DecodeString("2009000d00000f00003c0005000aca6102")
	fmt.Printf("%x\n", data)
	msg := pfcpv1.Message{}
	msg.Unmarshal(data)
	//
	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 9
	//Length: 13
	//SequenceNumber: 15
	if msg.Header.MessageType != 9 || msg.Header.Length != 13 ||
		msg.Header.SequenceNumber != 15 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	req, ok := msg.Body.(*pfcp.PFCPAssociationReleaseRequest)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	NodeIDStr := net.IP(req.IE.NodeID.NodeIDvalue).String()
	fmt.Println(NodeIDStr)
	if NodeIDStr != "10.202.97.2" {
		t.Errorf("Pfcp body unmarshal failed")
	}

	fmt.Println("------response------")

	res := &pfcpv1.Message{}
	// send PFCP Association Release Response(10)
	pfcptst.Dispatch(msg, res)

	data, err := res.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal failed")
	}
	ExpectedData, _ := hex.DecodeString("200a001200000f00003c0005000aca5e020013000101")

	fmt.Printf("ExpectedData:%x\n", ExpectedData)
	fmt.Printf("result      :%x\n", data)

	if !bytes.Equal(ExpectedData, data) {
		t.Errorf("Pfcp marshal failed")
	}
}

// Message Type: PFCP_Version_Not_Supported_Response (11)
// smf side
// receive response test,send response test
func TestPfcpUnmarshal11(t *testing.T) {
	//PFCP Version Not Supported Response (11)
	data, _ := hex.DecodeString("200b000400000700")
	fmt.Printf("%x\n", data)
	msg := pfcpv1.Message{}
	msg.Unmarshal(data)
	//
	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 11
	//Length: 4
	//SequenceNumber: 7
	if msg.Header.MessageType != 11 || msg.Header.Length != 4 ||
		msg.Header.SequenceNumber != 7 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	_, ok := msg.Body.(*pfcp.PFCPVersionNotSupportedResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}

	fmt.Println("------response------")

	res := &pfcpv1.Message{}
	// 发送 version not support response 11
	//pfcptst.Dispatch(msg, res)
	res.Header.Version = 1
	res.Header.MessageType = pfcp.PFCP_Version_Not_Supported_Response
	res.Header.Length = 0
	res.Header.SequenceNumber = 7

	request := &pfcp.PFCPVersionNotSupportedResponse{}
	res.Body = request
	data, err := res.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal failed")
	}
	ExpectedData, _ := hex.DecodeString("200b000400000700")

	fmt.Printf("ExpectedData:%x\n", ExpectedData)
	fmt.Printf("result      :%x\n", data)

	if !bytes.Equal(ExpectedData, data) {
		t.Errorf("Pfcp marshal failed")
	}
}

//4.4.3.4	N4 Report Procedure
// receive request:smf <-- upf
//Message Type: PFCP Node Report Request (12)
func TestPfcpUnmarshal12(t *testing.T) {
	//PFCP Node Report Request (12)
	data, _ := hex.DecodeString("200c001f00000200003c0005000aca5e0200650001010066000900670005020aca6402")
	fmt.Printf("%x\n", data)
	msg := pfcpv1.Message{}
	msg.Unmarshal(data)
	//
	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 12
	//Length:
	//SequenceNumber: 2
	if msg.Header.MessageType != 12 || msg.Header.Length != 31 ||
		msg.Header.SequenceNumber != 2 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	req, ok := msg.Body.(*pfcp.PFCPNodeReportRequest)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	NodeIDStr := net.IP(req.IE.NodeID.NodeIDvalue).String()
	NodeRType := req.IE.NodeReportType.UPFR
	NodeUPR := req.IE.UserPlanePathFailureReport.RemoteGTPUPeer.IPv4address.String()

	fmt.Println(NodeIDStr)
	fmt.Println(NodeRType)
	fmt.Println(NodeUPR)
	if NodeIDStr != "10.202.94.2" ||
		NodeRType != true ||
		NodeUPR != "10.202.100.2" {
		t.Errorf("Pfcp body unmarshal failed")
	}

	fmt.Println("------response------")
	// send PFCP Node Report Response(13)
	res := &pfcpv1.Message{}
	pfcptst.Dispatch(msg, res)

	data, err := res.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal failed")
	}
	ExpectedData, _ := hex.DecodeString("200d001200000200003c0005000aca61020013000101")

	fmt.Printf("ExpectedData:%x\n", ExpectedData)
	fmt.Printf("result      :%x\n", data)

	if !bytes.Equal(ExpectedData, data) {
		t.Errorf("Pfcp marshal failed")
	}
}

//4.4.3.5	N4 PFD management Procedure
// smf side
// send request:smf --> upf
//Message Type: PFCP PFD Management Request(3)
func TestPfcpMarshal3(t *testing.T) {
	//PFCP Association Release Request (9)
	ExpectedReqdata, _ := hex.DecodeString("2003002100000200003a00190018000461707031003b000d003d000901000005666c6f7731")
	//请求编码
	req := pfcpv1.Message{}
	request := &pfcp.PFCPPFDManagementRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_PFD_Management_Request,
			Length:         0, // 编码后填充
			SequenceNumber: 2},
	}
	request.IE = &pfcp.IEPFCPPFDManagementRequest{}
	request.IE.ApplicationIDsPFDs = &pfcp.IEApplicationIDsPFDs{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Application_IDs_PFDs,
			Length: 0,
		},
	}
	request.IE.ApplicationIDsPFDs.ApplicationID = &pfcp.IEApplicationID{} // ipv4
	request.IE.ApplicationIDsPFDs.ApplicationID.ApplicationIdentifier = []byte("app1")

	request.IE.ApplicationIDsPFDs.IeFlags.Set(pfcp.IE_PFD_context)
	request.IE.ApplicationIDsPFDs.PFD = &pfcp.IEPFDContext{}
	request.IE.ApplicationIDsPFDs.PFD.PFDContexts = &pfcp.IEPFDContents{}
	request.IE.ApplicationIDsPFDs.PFD.PFDContexts.FD = true
	request.IE.ApplicationIDsPFDs.PFD.PFDContexts.LengthofFlowDes = 5
	request.IE.ApplicationIDsPFDs.PFD.PFDContexts.FlowDescription = "flow1"

	req.Header.Version = request.PfcpHeader.Version
	req.Header.MessageType = request.PfcpHeader.MessageType
	req.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	// 编码后填充
	req.Header.Length = 0
	req.Body = request

	// pfcp encode
	data, err := req.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal err %s", err)
	}
	fmt.Println(req.String())
	fmt.Printf("%x\n", ExpectedReqdata)
	fmt.Printf("%x\n", data)
	if !bytes.Equal(data, ExpectedReqdata) {
		t.Errorf("Pfcp msg marshal failed")
	}
	fmt.Println("------response------")
	//PFCP Association Release Response (10)
	// 响应解码
	resdata, _ := hex.DecodeString("20040009000002000013000101")
	fmt.Printf("resdata:%x\n", resdata)
	msg := pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 4
	//Length: 9
	//SequenceNumber: 2
	if msg.Header.MessageType != 4 || msg.Header.Length != 9 ||
		msg.Header.SequenceNumber != 2 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.PFCPPFDManagementResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	fmt.Println((v.IE.Cause.CauseValue))

	// 响应处理
	err = pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}

	if v.IE.Cause.CauseValue != 1 {
		t.Errorf("Pfcp header unmarshal failed")
	}
}

//Message Type: PFCP Heartbeat Request (1)
// smf side
// send request：smf --> upf
func TestPfcpMarshal1(t *testing.T) {
	//PFCP Heartbeat Request (1)
	ExpectedReqdata, _ := hex.DecodeString("2001000c0000070000600004e07227a1")
	//请求编码
	req := pfcpv1.Message{}
	request := &pfcp.HeartbeatRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Heartbeat_Request,
			Length:         0, // 编码后填充
			SequenceNumber: 7},
	}
	//add IE
	//2019-04-30 01:47:13 +0000 UTC
	request.IE = &pfcp.IERecoveryTimeStamp{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Recovery_Time_Stamp,
			Length: 4},
		RecoveryTimeStamp: time.Unix(1556588833, 0),
	}

	req.Header.Version = request.PfcpHeader.Version
	req.Header.MessageType = request.PfcpHeader.MessageType
	req.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	//编码后填充
	req.Header.Length = 0
	req.Body = request
	fmt.Println(req.String())
	fmt.Println(request.IE.RecoveryTimeStamp.UTC().String())
	// todo pfcp encode
	data, err := req.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal err %s", err)
	}
	fmt.Printf("Expected:%x\n", ExpectedReqdata)
	fmt.Printf("Result  :%x\n", data)
	if !bytes.Equal(data, ExpectedReqdata) {
		t.Errorf("Pfcp msg marshal failed")
	}
	fmt.Println("------response------")
	// 响应解码
	//Message Type: PFCP Heartbeat Response (2)
	resdata, _ := hex.DecodeString("2002000c0000070000600004e07227a1")
	fmt.Printf("resdata:%x\n", resdata)
	msg := pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 2
	//Length: 12
	//SequenceNumber: 7
	if msg.Header.MessageType != 2 || msg.Header.Length != 12 ||
		msg.Header.SequenceNumber != 7 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.HeartbeatResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	//2019-04-30 01:47:13 +0000 UTC
	fmt.Println(v.IE.RecoveryTimeStamp.UTC())

	err = pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}

	if v.PfcpHeader.Length != 12 ||
		v.PfcpHeader.MessageType != 2 ||
		v.PfcpHeader.SequenceNumber != 7 ||
		v.IE.RecoveryTimeStamp.UTC().String() != "2019-04-30 01:47:13 +0000 UTC" {
		t.Errorf("Pfcp header unmarshal failed")
	}
}

//Message Type: PFCP Heartbeat Request (1)
// smf side
// receive request：smf <-- upf
func TestPfcpUnmarshal1(t *testing.T) {
	//PFCP Heartbeat Request (1)
	data, _ := hex.DecodeString("2001000c0000010000600004e07227a1")
	fmt.Printf("%x\n", data)
	msg := pfcpv1.Message{}
	msg.Unmarshal(data)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 1
	//Length: 12
	//SequenceNumber: 1
	if msg.Header.MessageType != 1 || msg.Header.Length != 12 ||
		msg.Header.SequenceNumber != 1 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.HeartbeatRequest)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}

	//2019-04-30 01:47:13 +0000 UTC
	fmt.Println(v.IE.RecoveryTimeStamp.UTC())

	fmt.Println("------response------")

	res := &pfcpv1.Message{}

	pfcptst.Dispatch(msg, res)
	// todo pfcp encode
	data, err := res.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal failed")
	}
	ExpectedData, _ := hex.DecodeString("2002000c0000010000600004e07227a1")

	fmt.Printf("ExpectedData:%x\n", ExpectedData)
	fmt.Printf("result      :%x\n", data)

	if !bytes.Equal(ExpectedData, data) {
		t.Errorf("Pfcp marshal failed")
	}
}

//4.4.1	N4 session management procedures

//4.4.1.2	N4 Session Establishment procedure
// smf side
// send request:smf --> upf
//Message Type: PFCP Session Establishment Request (50)
func TestPfcpMarshal50(t *testing.T) {
	//PFCP Session Establishment Request (50)
	ExpectedReqdata, _ := hex.DecodeString("2132013d000000000000000000000200003c0005000aca6102" +
		"0039000d0200000000000027100aca610200010078003800020001001d0004000000ff0002005" +
		"100140001000015000901000000010aca5e0100170036010000327065726d6974206f75742031" +
		"362066726f6d20616e79203530303020746f2061737369676e656420313032352d36353533350" +
		"07c000101005f000100006c000400000001006d00040000000100010035003800020002001d00" +
		"04000000ff000200130014000101005d00050601010101007c000101006c000400000002006d0" +
		"0040000000100030016006c000400000001002c00010200040005002a00010100030012006c00" +
		"0400000002002c00010c005800010100070020006d0004000000010019000100001a000a00000" +
		"000010000000001007c0001010055000500580001010071000101")
	//请求编码
	msg := pfcpv1.Message{}
	request := &pfcp.SessionEstablishmentRequest{
		PfcpHeader: pfcp.PfcpHeaderforSession{
			Version:        pfcp.Version,
			SFlag:          pfcp.Flag,
			MessageType:    pfcp.PFCP_Session_Establishment_Request,
			Length:         0, // 编码后填充
			SequenceNumber: 2},
	}
	request.IE = pfcp.IEsSessionEstablishmentRequest{}
	request.IE.NodeID = pfcp.IENodeID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Node_ID,
			Length: 5, // 编码后填充
		},
		NodeIDType:  pfcp.NodeIDType_IPv4_address,
		NodeIDvalue: []byte{10, 202, 97, 2}, // smf ip
	}
	// FSEID
	request.IE.CPFSEID = pfcp.IEFSEID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_F_SEID,
			Length: 0,
		},
		V4Flag:   pfcp.Flag,
		SEID:     0x0000000000002710,
		IPv4Addr: []byte{10, 202, 97, 2}, // smf ip
	}
	// PDR
	request.IE.CreatePDRs = make([]*pfcp.IECreatePDR, 2)
	{ // pdr 1
		pdr := &pfcp.IECreatePDR{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Create_PDR,
				Length: 0,
			},
		}
		pdr.PDRID.Type = pfcp.IE_Packet_Detection_Rule_ID
		pdr.PDRID.RuleID = 1
		pdr.Precedence.Type = pfcp.IE_Precedence
		pdr.Precedence.PrecedenceValue = 255

		pdr.PDI.Type = pfcp.IE_PDI
		{
			pdr.PDI.SourceInterface.Type = pfcp.IE_Source_Interface
			pdr.PDI.SourceInterface.InterfaceValue = pfcp.Access

			// optional IE
			pdr.PDI.IeFlags.Set(pfcp.IE_F_TEID)
			pdr.PDI.LocalFTEID = &pfcp.IEFTEID{}
			pdr.PDI.LocalFTEID.Type = pfcp.IE_F_TEID
			pdr.PDI.LocalFTEID.V4Flag = pfcp.Flag
			pdr.PDI.LocalFTEID.TEID = 0x00000001
			pdr.PDI.LocalFTEID.IPv4Addr = []byte{10, 202, 94, 1}

			pdr.PDI.IeFlags.Set(pfcp.IE_SDF_Filter)
			pdr.PDI.SDFFilters = make([]*pfcp.IESDFFilter, 1)
			pdr.PDI.SDFFilters[0] = &pfcp.IESDFFilter{}
			pdr.PDI.SDFFilters[0].Type = pfcp.IE_SDF_Filter
			pdr.PDI.SDFFilters[0].FDFlag = true
			pdr.PDI.SDFFilters[0].FlowDescription = []byte("permit out 16 from any 5000 to assigned 1025-65535")
			pdr.PDI.SDFFilters[0].LengthofFlowDescription = uint16(len(pdr.PDI.SDFFilters[0].FlowDescription))

			pdr.PDI.IeFlags.Set(pfcp.IE_QFI)
			pdr.PDI.QFIs = make([]*pfcp.IEQFI, 1)
			pdr.PDI.QFIs[0] = &pfcp.IEQFI{}
			pdr.PDI.QFIs[0].Type = pfcp.IE_QFI
			pdr.PDI.QFIs[0].Value = 0x01
		}
		// optional IE
		pdr.IeFlags.Set(pfcp.IE_Outer_Header_Removal)
		pdr.OuterHeaderRemoval = &pfcp.IEOuterHeaderRemoval{}
		pdr.OuterHeaderRemoval.Type = pfcp.IE_Outer_Header_Removal
		pdr.OuterHeaderRemoval.Description = uint8(pfcp.IEOuterHRemoval_GTPU_UDP_IPv4)

		pdr.IeFlags.Set(pfcp.IE_FAR_ID)
		pdr.FARID = &pfcp.IEFARID{}
		pdr.FARID.Type = pfcp.IE_FAR_ID
		pdr.FARID.Value = 1

		pdr.IeFlags.Set(pfcp.IE_QER_ID)
		pdr.QERIDs = make([]*pfcp.IEQERID, 1)
		pdr.QERIDs[0] = &pfcp.IEQERID{}
		pdr.QERIDs[0].Type = pfcp.IE_QER_ID
		pdr.QERIDs[0].Value = 1

		request.IE.CreatePDRs[0] = pdr

		// pdr 2
		pdr = &pfcp.IECreatePDR{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Create_PDR,
				Length: 0,
			},
		}
		pdr.PDRID.Type = pfcp.IE_Packet_Detection_Rule_ID
		pdr.PDRID.RuleID = 2
		pdr.Precedence.Type = pfcp.IE_Precedence
		pdr.Precedence.PrecedenceValue = 255

		pdr.PDI.Type = pfcp.IE_PDI
		pdr.PDI.SourceInterface.Type = pfcp.IE_Source_Interface
		pdr.PDI.SourceInterface.InterfaceValue = pfcp.Core

		// optional IE
		pdr.PDI.IeFlags.Set(pfcp.IE_UE_IP_Address)
		pdr.PDI.UEIPaddress = &pfcp.IEUEIPaddress{}
		pdr.PDI.UEIPaddress.Type = pfcp.IE_UE_IP_Address
		pdr.PDI.UEIPaddress.SD = pfcp.Flag
		pdr.PDI.UEIPaddress.V4Flag = pfcp.Flag
		pdr.PDI.UEIPaddress.IPv4Addr = []byte{1, 1, 1, 1}

		pdr.PDI.IeFlags.Set(pfcp.IE_QFI)
		pdr.PDI.QFIs = make([]*pfcp.IEQFI, 1)
		pdr.PDI.QFIs[0] = &pfcp.IEQFI{}
		pdr.PDI.QFIs[0].Type = pfcp.IE_QFI
		pdr.PDI.QFIs[0].Value = 0x01

		// optional IE
		pdr.IeFlags.Set(pfcp.IE_FAR_ID)
		pdr.FARID = &pfcp.IEFARID{}
		pdr.FARID.Type = pfcp.IE_FAR_ID
		pdr.FARID.Value = 2

		pdr.IeFlags.Set(pfcp.IE_QER_ID)
		pdr.QERIDs = make([]*pfcp.IEQERID, 1)
		pdr.QERIDs[0] = &pfcp.IEQERID{}
		pdr.QERIDs[0].Type = pfcp.IE_QER_ID
		pdr.QERIDs[0].Value = 1

		request.IE.CreatePDRs[1] = pdr
	}
	// FAR
	request.IE.CreateFARs = make([]*pfcp.IECreateFAR, 2)
	{ // far 1
		far := &pfcp.IECreateFAR{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Create_FAR,
				Length: 0,
			},
		}
		far.FARID.Type = pfcp.IE_FAR_ID
		far.FARID.Value = 1

		far.ApplyAction.Type = pfcp.IE_Apply_Action
		far.ApplyAction.Flag = pfcp.FORW
		//optional IE
		far.IeFlags.Set(pfcp.IE_Forwarding_Parameters)
		far.ForwardingParameters = &pfcp.IEForwardingParameters{}
		far.ForwardingParameters.Type = pfcp.IE_Forwarding_Parameters
		far.ForwardingParameters.DstInterface.Type = pfcp.IE_Destination_Interface
		far.ForwardingParameters.DstInterface.Value = pfcp.IEDestinationInterface_Core

		request.IE.CreateFARs[0] = far

		// far 2
		far = &pfcp.IECreateFAR{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Create_FAR,
				Length: 0,
			},
		}
		far.FARID.Type = pfcp.IE_FAR_ID
		far.FARID.Value = 2

		far.ApplyAction.Type = pfcp.IE_Apply_Action
		far.ApplyAction.Flag = pfcp.NOCP + pfcp.BUFF
		//optional IE
		far.IeFlags.Set(pfcp.IE_BAR_ID)
		far.BARID = &pfcp.IEBARID{}
		far.BARID.Type = pfcp.IE_BAR_ID
		far.BARID.Value = 1

		request.IE.CreateFARs[1] = far
	}
	//BAR
	request.IE.IeFlags.Set(pfcp.IE_Create_BAR)
	request.IE.CreateBAR = &pfcp.IECreateBAR{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Create_BAR,
			Length: 0,
		},
	}
	request.IE.CreateBAR.BARID.Type = pfcp.IE_BAR_ID
	request.IE.CreateBAR.BARID.Value = 1

	//QER
	request.IE.IeFlags.Set(pfcp.IE_Create_QER)
	request.IE.CreateQERs = make([]*pfcp.IECreateQER, 1)
	request.IE.CreateQERs[0] = &pfcp.IECreateQER{}
	request.IE.CreateQERs[0].Type = pfcp.IE_Create_QER
	request.IE.CreateQERs[0].QERID.Type = pfcp.IE_QER_ID
	request.IE.CreateQERs[0].QERID.Value = 1
	request.IE.CreateQERs[0].GateStatus.Type = pfcp.IE_Gate_Status
	request.IE.CreateQERs[0].GateStatus.ULGate = pfcp.IEGateStatus_OPEN
	request.IE.CreateQERs[0].GateStatus.DLGate = pfcp.IEGateStatus_OPEN
	//optional IE
	request.IE.CreateQERs[0].IeFlags.Set(pfcp.IE_MBR)
	request.IE.CreateQERs[0].MaximumBitrate.Type = pfcp.IE_MBR
	request.IE.CreateQERs[0].MaximumBitrate.ULMBR = 1
	request.IE.CreateQERs[0].MaximumBitrate.DLMBR = 1

	request.IE.CreateQERs[0].IeFlags.Set(pfcp.IE_QFI)
	request.IE.CreateQERs[0].QoSflowidentifier.Type = pfcp.IE_QFI
	request.IE.CreateQERs[0].QoSflowidentifier.Value = 1

	// PDNType
	request.IE.IeFlags.Set(pfcp.IE_PDN_Type)
	request.IE.PDNType = &pfcp.IEPDNType{}
	request.IE.PDNType.Type = pfcp.IE_PDN_Type
	request.IE.PDNType.PDNType = pfcp.PDNType_IPv4

	// Encoding message filling
	msg.Header.Version = request.PfcpHeader.Version
	msg.Header.SFlag = request.PfcpHeader.SFlag
	msg.Header.MessageType = request.PfcpHeader.MessageType
	msg.Header.SEID = request.PfcpHeader.SEID
	msg.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	// Fill after coding
	msg.Header.Length = 0
	msg.Body = request
	fmt.Println(msg.String())

	// pfcp encode
	data, err := msg.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal err %s", err)
	}
	fmt.Printf("%x\n", ExpectedReqdata)
	fmt.Printf("%x\n", data)
	if !bytes.Equal(data, ExpectedReqdata) {
		t.Errorf("Pfcp msg marshal failed")
	}
	fmt.Println("------response------")
	// 响应解码
	resdata, _ := hex.DecodeString("2133002b000000000000271000000200003c000500" +
		"0aca5e0200130001010039000d020000000000004e200aca5e02")
	fmt.Printf("resdata:%x\n", resdata)
	msg = pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 1
	//Message Type: 51
	//Length: 43
	//SEID: 0x0000000000002710
	//SequenceNumber: 1
	if msg.Header.MessageType != 51 || msg.Header.Length != 43 ||
		msg.Header.SEID != 0x0000000000002710 ||
		msg.Header.SequenceNumber != 2 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.SessionEstablishmentResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.SFlag = msg.Header.SFlag
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SEID = msg.Header.SEID
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	//10.202.94.2
	fmt.Println(net.IP(v.IE.NodeID.NodeIDvalue))
	fmt.Println((v.IE.Cause.CauseValue))
	fmt.Println(v.IE.UPFSEID.SEID)
	fmt.Println(v.IE.UPFSEID.IPv4Addr)

	err = pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}
	// upf ip
	if net.IP(v.IE.NodeID.NodeIDvalue).String() != "10.202.94.2" ||
		v.IE.Cause.CauseValue != 1 ||
		v.IE.UPFSEID.SEID != 0x0000000000004e20 ||
		!v.IE.UPFSEID.IPv4Addr.Equal(net.ParseIP("10.202.94.2")) {
		t.Errorf("Pfcp header unmarshal failed")
	}
}

//4.4.1.3	N4 Session Modification procedure
// smf side
// send request:smf --> upf
//Message Type: PFCP Session Modification Request (52)
func TestPfcpMarshal52(t *testing.T) {
	//PFCP Session Modification Request (52)
	ExpectedReqdata, _ := hex.DecodeString("213400340000000000004e2000000300000a0024006c00040000" +
		"0002000b0013002a0001000054000a0100000000010aca5e0b002c000102")
	//请求编码
	msg := pfcpv1.Message{}
	request := &pfcp.SessionModifyRequest{
		PfcpHeader: pfcp.PfcpHeaderforSession{
			Version:        pfcp.Version,
			SFlag:          pfcp.Flag,
			MessageType:    pfcp.PFCP_Session_Modification_Request,
			Length:         0, // 编码后填充
			SEID:           0x0000000000004e20,
			SequenceNumber: 3},
	}
	request.IE = pfcp.IEsSessionModifyRequest{}

	// FAR
	request.IE.IeFlags.Set(pfcp.IE_Update_FAR)
	request.IE.UpdateFARs = make([]*pfcp.IEUpdateFAR, 1)
	{ // far 1
		far := &pfcp.IEUpdateFAR{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Update_FAR,
				Length: 0,
			},
		}
		far.FARID.Type = pfcp.IE_FAR_ID
		far.FARID.Value = 2
		//optional IE
		far.IeFlags.Set(pfcp.IE_Apply_Action)
		far.ApplyAction.Type = pfcp.IE_Apply_Action
		far.ApplyAction.Flag = pfcp.FORW

		far.IeFlags.Set(pfcp.IE_Update_Forwarding_Parameters)
		far.UpdateForwardingPara = &pfcp.IEUpdateForwardingParameters{}
		//far.UpdateForwardingPara.Type = pfcp.IE_Update_Forwarding_Parameters

		far.UpdateForwardingPara.IeFlags.Set(pfcp.IE_Destination_Interface)
		far.UpdateForwardingPara.DstInterface = &pfcp.IEDestinationInterface{}
		far.UpdateForwardingPara.DstInterface.Value = pfcp.IEDestinationInterface_Access

		far.UpdateForwardingPara.IeFlags.Set(pfcp.IE_Outer_Header_Creation)
		far.UpdateForwardingPara.OuterHeaderCreation = &pfcp.IEOuterHeaderCreation{}
		far.UpdateForwardingPara.OuterHeaderCreation.Description = pfcp.IEOuterHCreation_GTPU_UDP_IPv4
		far.UpdateForwardingPara.OuterHeaderCreation.TEID = 0x00000001
		far.UpdateForwardingPara.OuterHeaderCreation.IPv4Addr = []byte{10, 202, 94, 11}

		request.IE.UpdateFARs[0] = far
	}

	// Encoding message filling
	msg.Header.Version = request.PfcpHeader.Version
	msg.Header.SFlag = request.PfcpHeader.SFlag
	msg.Header.MessageType = request.PfcpHeader.MessageType
	msg.Header.SEID = request.PfcpHeader.SEID
	msg.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	// Fill after coding
	msg.Header.Length = 0
	msg.Body = request
	fmt.Println(msg.String())

	// pfcp encode
	data, err := msg.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal err %s", err)
	}
	fmt.Printf("Expected:%x\n", ExpectedReqdata)
	fmt.Printf("Result  :%x\n", data)
	if !bytes.Equal(data, ExpectedReqdata) {
		t.Errorf("Pfcp msg marshal failed")
	}
	fmt.Println("------response------")
	// 响应解码
	resdata, _ := hex.DecodeString("213500110000000000002710000003000013000101")
	fmt.Printf("ResData:%x\n", resdata)
	msg = pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 1
	//Message Type: 53
	//Length: 17
	//SEID: 0x0000000000002710
	//SequenceNumber: 3
	if msg.Header.MessageType != 53 || msg.Header.Length != 17 ||
		msg.Header.SEID != 0x0000000000002710 ||
		msg.Header.SequenceNumber != 3 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.SessionModifyResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.SFlag = msg.Header.SFlag
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SEID = msg.Header.SEID
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	fmt.Println((v.IE.Cause.CauseValue))

	err = pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}
	// upf ip
	if v.IE.Cause.CauseValue != 1 {
		t.Errorf("Pfcp header unmarshal failed")
	}
}

//4.4.1.4	N4 Session Release procedure
// smf side
// send request:smf --> upf
//Message Type: PFCP Session Release Request (54)
func TestPfcpMarshal54(t *testing.T) {
	//PFCP Session Release Request (54)
	ExpectedReqdata, _ := hex.DecodeString("2136000c0000000000004e2000000400")
	//请求编码
	msg := pfcpv1.Message{}
	request := &pfcp.SessionReleaseRequest{
		PfcpHeader: pfcp.PfcpHeaderforSession{
			Version:        pfcp.Version,
			SFlag:          pfcp.Flag,
			MessageType:    pfcp.PFCP_Session_Deletion_Request,
			Length:         0, // 编码后填充
			SEID:           0x0000000000004e20,
			SequenceNumber: 4},
	}
	request.IE = pfcp.IEsSessionDelRequest{}

	// Encoding message filling
	msg.Header.Version = request.PfcpHeader.Version
	msg.Header.SFlag = request.PfcpHeader.SFlag
	msg.Header.MessageType = request.PfcpHeader.MessageType
	msg.Header.SEID = request.PfcpHeader.SEID
	msg.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	// Fill after coding
	msg.Header.Length = 0
	msg.Body = request
	fmt.Println(msg.String())

	// pfcp encode
	data, err := msg.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal err %s", err)
	}
	fmt.Printf("Expected:%x\n", ExpectedReqdata)
	fmt.Printf("Result  :%x\n", data)
	if !bytes.Equal(data, ExpectedReqdata) {
		t.Errorf("Pfcp msg marshal failed")
	}
	fmt.Println("------response------")
	// 响应解码
	resdata, _ := hex.DecodeString("213700110000000000002710000004000013000101")
	fmt.Printf("ResData:%x\n", resdata)
	msg = pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 1
	//Message Type: 55
	//Length: 17
	//SEID: 0x0000000000002710
	//SequenceNumber: 4
	if msg.Header.MessageType != 55 || msg.Header.Length != 17 ||
		msg.Header.SEID != 0x0000000000002710 ||
		msg.Header.SequenceNumber != 4 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.SessionReleaseResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.SFlag = msg.Header.SFlag
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SEID = msg.Header.SEID
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	fmt.Println((v.IE.Cause.CauseValue))

	err = pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}
	// upf ip
	if v.IE.Cause.CauseValue != 1 {
		t.Errorf("Pfcp header unmarshal failed")
	}
}

//4.4.2.2	N4 Session Level Reporting Procedure
// smf side
// receive request：smf <-- upf
//Message Type: PFCP Session Report Request (56)
func TestPfcpUnmarshal56(t *testing.T) {
	//PFCP Session Report Request (56)
	data, _ := hex.DecodeString("2138001b000000000000271000000200002700010100530006003800020002")
	fmt.Printf("%x\n", data)
	msg := pfcpv1.Message{}
	msg.Unmarshal(data)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 1
	//Message Type: 56
	//Length: 27
	//SEID: 0x0000000000002710
	//SequenceNumber: 2
	if msg.Header.MessageType != 56 || msg.Header.Length != 27 ||
		msg.Header.SEID != 0x0000000000002710 ||
		msg.Header.SequenceNumber != 2 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.SessionReportRequest)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}

	fmt.Println("DLDR:", v.IE.ReportType.DLDR)
	fmt.Println("DownlinkDataReport,RuleID:", v.IE.DownlinkDataReport.PDRID.RuleID)
	if v.IE.ReportType.DLDR != true ||
		v.IE.DownlinkDataReport.PDRID.RuleID != 2 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	fmt.Println("------response------")

	res := &pfcpv1.Message{}

	pfcptst.Dispatch(msg, res)
	// pfcp encode
	data, err := res.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal failed")
	}
	ExpectedData, _ := hex.DecodeString("213900110000000000004e20000002000013000101")

	fmt.Printf("ExpectedData:%x\n", ExpectedData)
	fmt.Printf("result      :%x\n", data)

	if !bytes.Equal(ExpectedData, data) {
		t.Errorf("Pfcp marshal failed")
	}
}

func TestPfcpMarshal5tmp(t *testing.T) {
	//PFCP Association Setup Request (5)
	ExpectedReqdata, _ := hex.DecodeString("2005001500000100003c000500ac14001000600004e2634527")
	//请求编码
	req := pfcpv1.Message{}
	request := &pfcp.PFCPAssociationSetupRequest{
		PfcpHeader: pfcp.PfcpHeaderforNode{
			Version:        pfcp.Version,
			MessageType:    pfcp.PFCP_Association_Setup_Request,
			Length:         0, // todo 编码后填充
			SequenceNumber: 1},
	}
	request.IE = &pfcp.IEPFCPAssociationSetupRequest{}
	request.IE.NodeID = &pfcp.IENodeID{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Node_ID,
			Length: 5,
		},
		NodeIDType:  0,
		NodeIDvalue: []byte{172, 20, 0, 16}, // smf ip
	}

	request.IE.RecoveryTimeStamp = &pfcp.IERecoveryTimeStamp{
		IETypeLength: pfcp.IETypeLength{
			Type:   pfcp.IE_Recovery_Time_Stamp,
			Length: 4,
		},
		// 1556588833 = nt64(0xe07227a1(wireshark) - 0x83aa7e80(1900到1970的秒差))
		RecoveryTimeStamp: time.Unix(1556588833, 0),
	}
	fmt.Println("time", request.IE.RecoveryTimeStamp.RecoveryTimeStamp.UTC().String())
	req.Header.Version = request.PfcpHeader.Version
	req.Header.MessageType = pfcp.PFCP_Association_Setup_Request
	req.Header.SequenceNumber = request.PfcpHeader.SequenceNumber
	//todo 编码后填充
	req.Header.Length = 0
	req.Body = request
	fmt.Println(req.String())

	// todo pfcp encode
	data, err := req.Marshal()
	if err != nil {
		t.Errorf("Pfcp msg marshal err %s", err)
	}
	fmt.Printf("%x\n", ExpectedReqdata)
	fmt.Printf("%x\n", data)
	if !bytes.Equal(data, ExpectedReqdata) {
		t.Errorf("Pfcp msg marshal failed")
	}
	fmt.Println("------response------")
	// 响应解码
	resdata, _ := hex.DecodeString("2006003c00000100003c000500ac100002001300010100600004e0b577cc002b000200010074000b2980ac100002046e346370007400092980ac100003026e33")
	fmt.Printf("resdata:%x\n", resdata)
	msg := pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 6
	//Length: 26
	//SequenceNumber: 1
	if msg.Header.MessageType != 6 ||
		msg.Header.SequenceNumber != 1 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.PFCPAssociationSetupResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	//10.202.97.2
	//2019-04-30 01:47:13 +0000 UTC
	fmt.Println(net.IP(v.IE.NodeID.NodeIDvalue))
	fmt.Println((v.IE.Cause.CauseValue))
	fmt.Println(v.IE.RecoveryTimeStamp.RecoveryTimeStamp.UTC())

	err = pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}
	// upf ip
	if net.IP(v.IE.NodeID.NodeIDvalue).String() != "172.20.0.16" {
		t.Errorf("Pfcp header unmarshal failed")
	}
}

func TestPfcpMarshal6(t *testing.T) {
	//PFCP Association Setup Request (5)
	fmt.Println("------response------")
	// 响应解码
	resdata, _ := hex.DecodeString("2006003600000100003c000500ac110005001300010100600004e27b2035002b00060001000000000074000e6902ac11000503646674026e3300")
	fmt.Printf("resdata:%x\n", resdata)
	msg := pfcpv1.Message{}
	msg.Unmarshal(resdata)

	fmt.Println(msg.String())
	//Version: 1
	//MP: 0
	//S: 0
	//Message Type: 6
	//Length: 26
	//SequenceNumber: 1
	if msg.Header.MessageType != 6 || msg.Header.Length != 54 ||
		msg.Header.SequenceNumber != 1 {
		t.Errorf("Pfcp header unmarshal failed")
	}

	v, ok := msg.Body.(*pfcp.PFCPAssociationSetupResponse)
	if !ok {
		t.Errorf("Pfcp body unmarshal failed")
	}
	// 消息头填充
	v.PfcpHeader.Version = msg.Header.Version
	v.PfcpHeader.MessageType = msg.Header.MessageType
	v.PfcpHeader.Length = msg.Header.Length
	v.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

	//10.202.97.2
	//2019-04-30 01:47:13 +0000 UTC
	fmt.Println(net.IP(v.IE.NodeID.NodeIDvalue))
	fmt.Println((v.IE.Cause.CauseValue))
	fmt.Println(v.IE.RecoveryTimeStamp.RecoveryTimeStamp.UTC())

	err := pfcptst.Dispatch(msg, nil)
	if err != nil {
		t.Errorf("Pfcp msg handler err %s", err)
	}
	// upf ip
	if net.IP(v.IE.NodeID.NodeIDvalue).String() != "172.17.0.5" {
		t.Errorf("Pfcp header unmarshal failed")
	}
}
