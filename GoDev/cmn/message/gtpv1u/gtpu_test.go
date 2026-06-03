package gtpv1u

import (
	"bytes"
	"fmt"
	"testing"
)

// 3GPP TS 29.281 V15.5.0 (2018-12)
//解码消息
var TestDataSet map[string][]byte

//var RecvMsgbuf N3MsgParser

//var TestEncodeData map[string]*GTPPDU
func init() {
	TestDataSet = map[string][]byte{}

	TestDataSet["T-PDU-ICMP"] = []byte{0x30, 0xff, 0x00, 0x3c, 0x04, 0x18, 0x01, 0x55,
		0x45, 0x00, 0x00, 0x3c, 0x10, 0xb0, 0x00, 0x00, 0x40, 0x01, 0xd5, 0x1f, 0x0a, 0x42, 0x48, 0x6f,
		0x40, 0x00, 0x02, 0x41, 0x08, 0x00, 0x4b, 0x8d, 0x00, 0x01, 0x01, 0xce, 0x61, 0x62, 0x63, 0x64,
		0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74,
		0x75, 0x76, 0x77, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69}
	TestDataSet["ErrorIndication"] = []byte{0x32, 0x1a, 0x00, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x10, 0x00, 0x00, 0x06, 0xf2, 0x85, 0x00, 0x04, 0x0a, 0x42, 0x48, 0x0b}

	TestDataSet["EchoRequest"] = []byte{0x32, 0x01, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	TestDataSet["EchoResponse"] = []byte{0x32, 0x02, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xf0, 0x00, 0x00, 0x0e, 0x00}

}

//RecvMsgbuf.Msgbuf = TestDataSet["T-PDU-ICMP"]

//TestEncodeData = map[string]*GTPPDU{}
//TestEncodeData["T-PDU-ICMP"] = &GtpMsgSendICMP

// 编码消息
// T-PDU-ICMP
var GtpMsgSendICMP = GTPPDU{
	Gtpv1uHeader: Gtpv1uHeader{
		Version:           1,
		PT:                1,
		Spare:             0,
		EFlag:             0,
		SFlag:             0,
		PNFlag:            0,
		MessageType:       255,
		Length:            60,
		TEID:              0x04180155,
		SequenceNumber:    0,
		NPDUNumber:        0,
		NextExtHeaderType: 0,
	},
	Gtpbody: []byte{
		0x45, 0x00, 0x00, 0x3c, 0x10, 0xb0, 0x00, 0x00, 0x40, 0x01, 0xd5, 0x1f, 0x0a, 0x42, 0x48, 0x6f,
		0x40, 0x00, 0x02, 0x41, 0x08, 0x00, 0x4b, 0x8d, 0x00, 0x01, 0x01, 0xce, 0x61, 0x62, 0x63, 0x64,
		0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f, 0x70, 0x71, 0x72, 0x73, 0x74,
		0x75, 0x76, 0x77, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69,
	},
}

// 编码消息
// ErrorIndication
var GtpMsgSendErrInd = ErrIndication{
	Gtpv1uHeader: Gtpv1uHeader{Version: 1,
		PT:                1,
		Spare:             0,
		EFlag:             0,
		SFlag:             1,
		PNFlag:            0,
		MessageType:       26,
		Length:            16,
		TEID:              0,
		SequenceNumber:    0,
		NPDUNumber:        0,
		NextExtHeaderType: 0,
	},
	IE: ErrIndicationIE{
		Teid: IETunEpIdI{Type: 0x10,
			TeidI: 0x000006f2,
		},
		PeerAddr: IEPeerAddr{Type: 0x85,
			Length: 0x0004,
			IpAddr: []byte{0x0a, 0x42, 0x48, 0x0b},
		},
	},
}

// 放到对应的模块，本包中存在循环import
/*func TestSendN3MsgHandle(t *testing.T) {
	fmt.Printf("target value: %#x\n", TestDataSet["T-PDU-ICMP"])
	// test dpe.SendN3MsgHandle
	data, err := dpe.SendN3MsgHandle(GtpMsgSendICMP.Gtpbody)
	if err != nil {
		t.Errorf("EncodeMsg fail %s", err)
	}
	if !bytes.Equal(TestDataSet["T-PDU-ICMP"], data) {
		t.Errorf("EncodeMsg fault.")
	}

}*/

//GtpMsgSend 编码处理测试
func TestEncodeMsgPDU(t *testing.T) {
	fmt.Println(GtpMsgSendICMP)
	data, err := GtpMsgSendICMP.EncodeMsg()
	if err != nil {
		t.Errorf("EncodeMsg fail")
	}
	fmt.Printf("target value: %#x\n", TestDataSet["T-PDU-ICMP"])
	fmt.Printf("Encode value: %#x\n", data)

	if !bytes.Equal(TestDataSet["T-PDU-ICMP"], data) {
		t.Errorf("EncodeMsg fault.")
	}
}

func TestEncodeMsgErrorInd(t *testing.T) {

	// TestEncodeData["ErrorIndication"] = &GtpMsgSendErrInd

	fmt.Println(GtpMsgSendErrInd)
	data, err := GtpMsgSendErrInd.EncodeMsg()
	if err != nil {
		t.Errorf("EncodeMsg fail")
	}
	fmt.Printf("target value: %#x\n", TestDataSet["ErrorIndication"])
	fmt.Printf("Encode value: %#x\n", data)

	if !bytes.Equal(TestDataSet["ErrorIndication"], data) {
		t.Errorf("EncodeMsg fault.")
	}

}

// EchoRequest
var GtpMsgSendEchoRequest = EchoRequest{
	Gtpv1uHeader: Gtpv1uHeader{Version: 1,
		PT:                1,
		Spare:             0,
		EFlag:             0,
		SFlag:             1,
		PNFlag:            0,
		MessageType:       1,
		Length:            4,
		TEID:              0,
		SequenceNumber:    240,
		NPDUNumber:        0,
		NextExtHeaderType: 0,
	},
	IE: EchoRequestIE{
		IEPriExt: IEPriExt{Type: 0x00, Length: 0x0000, ExtId: 0x0000, ExtValue: []byte{0x00}},
	},
}

// EchoResponse
var GtpMsgSendEchoResponse = EchoResponse{
	Gtpv1uHeader: Gtpv1uHeader{
		Version:           1,
		PT:                1,
		Spare:             0,
		EFlag:             0,
		SFlag:             1,
		PNFlag:            0,
		MessageType:       2,
		Length:            4,
		TEID:              0,
		SequenceNumber:    240,
		NPDUNumber:        0,
		NextExtHeaderType: 0,
	},
	IE: EchoResponseIE{
		Recovery: IERecovery{0x0e, 0x00},
	},
}

// EchoRequest
func TestEncodeMsgEchoRequest(t *testing.T) {

	// TestEncodeData["ErrorIndication"] = &GtpMsgSendEchoRequest

	fmt.Println(GtpMsgSendEchoRequest)
	data, err := GtpMsgSendEchoRequest.EncodeMsg()
	if err != nil {
		t.Errorf("EncodeMsg fail")
	}
	fmt.Printf("target value: %#x\n", TestDataSet["EchoRequest"])
	fmt.Printf("Encode value: %#x\n", data)

	if bytes.Equal(TestDataSet["EchoRequest"], data) {
		t.Errorf("EncodeMsg fault.")
	}

}

// EchoResponse
func TestEncodeMsgEchoResponse(t *testing.T) {

	// TestEncodeData["ErrorIndication"] = &GtpMsgSendEchoRequest

	fmt.Println(GtpMsgSendEchoResponse)
	data, err := GtpMsgSendEchoResponse.EncodeMsg()
	if err != nil {
		t.Errorf("EncodeMsg fail")
	}
	fmt.Printf("target value: %#x\n", TestDataSet["EchoResponse"])
	fmt.Printf("Encode value: %#x\n", data)

	if !bytes.Equal(TestDataSet["EchoResponse"], data) {
		t.Errorf("EncodeMsg fault.")
	}

}

// Decode 字节码转换为结构数据
func TestDecodeMsg(t *testing.T) {

	var icmptest N3MsgParser
	//T-PDU-ICMP
	//ErrorIndication
	//EchoRequest
	//EchoResponse
	fmt.Printf("Msg byte len:%d\n", len(TestDataSet["EchoResponse"]))
	fmt.Printf("Msg %#x\n", TestDataSet["EchoResponse"])
	fmt.Printf("Msg %8b\n", TestDataSet["EchoResponse"])

	icmptest.Msgbuf = TestDataSet["EchoResponse"]

	icmptest.DecodeMsg()
	// RecvMsgbuf.DecodeMsg()
	//fmt.Printf("%s\n", ICMP_DATA)
	//fmt.Printf("%x\n", ICMP_DATA)
	//fmt.Println( ICMP_DATA)
}

// T-PDU-ICMP Decode 字节码转换为结构数据
func TestDecodeMsgTPDUICMP(t *testing.T) {

	var N3Msgtest N3MsgParser
	//T-PDU-ICMP
	//ErrorIndication
	//EchoRequest
	//EchoResponse
	fmt.Printf("Msg byte len:%d\n", len(TestDataSet["T-PDU-ICMP"]))
	fmt.Printf("Msg %#x\n", TestDataSet["T-PDU-ICMP"])
	//fmt.Printf("Msg %8b\n", TestDataSet["T-PDU-ICMP"])

	N3Msgtest.Msgbuf = TestDataSet["T-PDU-ICMP"]
	fmt.Println(N3Msgtest.GtpuMsg)
	N3Msgtest.DecodeMsg()
	fmt.Printf("target header value: %v\n", GtpMsgSendICMP.Gtpv1uHeader)
	fmt.Printf("Decode header value: %v\n", N3Msgtest.GtpuMsg.Gtpv1uHeader)
	fmt.Printf("target body value: %#x\n", GtpMsgSendICMP.Gtpbody)
	fmt.Printf("Decode body value: %#x\n", N3Msgtest.GtpuMsg.Gtpbody)

	if N3Msgtest.GtpuMsg.Gtpv1uHeader != GtpMsgSendICMP.Gtpv1uHeader ||
		!bytes.Equal(N3Msgtest.GtpuMsg.Gtpbody, GtpMsgSendICMP.Gtpbody) {
		t.Errorf("DecodeMsg fault.")
	}
}

// ErrorIndication Decode 字节码转换为结构数据
func TestDecodeMsgErrorIndication(t *testing.T) {

	var N3Msgtest N3MsgParser
	//T-PDU-ICMP
	//ErrorIndication
	//EchoRequest
	//EchoResponse
	fmt.Printf("Msg byte len:%d\n", len(TestDataSet["ErrorIndication"]))
	fmt.Printf("Msg %#x\n", TestDataSet["ErrorIndication"])
	//fmt.Printf("Msg %8b\n", TestDataSet["ErrorIndication"])

	N3Msgtest.Msgbuf = TestDataSet["ErrorIndication"]
	//fmt.Println(N3Msgtest.ErrIndicationMsg)
	N3Msgtest.DecodeMsg()
	fmt.Printf("target header value: %v\n", GtpMsgSendErrInd)
	fmt.Printf("Decode header value: %v\n", *N3Msgtest.ErrIndicationMsg)
	//fmt.Printf("target body value: %#x\n", GtpMsgSendErrInd.Gtpbody)
	//fmt.Printf("Decode body value: %#x\n", N3Msgtest.GtpuMsg.Gtpbody)
	if N3Msgtest.ErrIndicationMsg.Gtpv1uHeader != GtpMsgSendErrInd.Gtpv1uHeader ||
		N3Msgtest.ErrIndicationMsg.ExtH != GtpMsgSendErrInd.ExtH ||
		N3Msgtest.ErrIndicationMsg.IE.Teid != GtpMsgSendErrInd.IE.Teid ||
		N3Msgtest.ErrIndicationMsg.IE.PeerAddr.Length != GtpMsgSendErrInd.IE.PeerAddr.Length ||
		!bytes.Equal((N3Msgtest.ErrIndicationMsg.IE.PeerAddr.IpAddr), (GtpMsgSendErrInd.IE.PeerAddr.IpAddr)) {
		t.Errorf("DecodeMsg fault.")
	}
}

// EchoRequest Decode 字节码转换为结构数据
func TestDecodeMsgEchoRequest(t *testing.T) {

	var N3Msgtest N3MsgParser
	//T-PDU-ICMP
	//ErrorIndication
	//EchoRequest
	//EchoResponse
	fmt.Printf("Msg byte len:%d\n", len(TestDataSet["EchoRequest"]))
	fmt.Printf("Msg %#x\n", TestDataSet["EchoRequest"])

	N3Msgtest.Msgbuf = TestDataSet["EchoRequest"]
	fmt.Println(N3Msgtest.GtpuMsg)
	N3Msgtest.DecodeMsg()
	fmt.Printf("target header value: %v\n", GtpMsgSendEchoRequest.Gtpv1uHeader)
	fmt.Printf("Decode header value: %v\n", N3Msgtest.EchoRequestMsg.Gtpv1uHeader)
	fmt.Printf("target body value: %v\n", GtpMsgSendEchoRequest.IE)
	fmt.Printf("Decode body value: %v\n", N3Msgtest.EchoRequestMsg.IE)
	// 异常用例，ExtValue 不等
	if N3Msgtest.EchoRequestMsg.Gtpv1uHeader != GtpMsgSendEchoRequest.Gtpv1uHeader ||
		bytes.Equal(N3Msgtest.EchoRequestMsg.IE.ExtValue, GtpMsgSendEchoRequest.IE.ExtValue) {
		t.Errorf("DecodeMsg fault.")
	}
}

// 字符串转换为字节码
func TestStrTobyteA(t *testing.T) {

	msgData := []byte{0x30, 0xff}
	msgDataB := []byte("xasc")

	fmt.Printf("%b\n", msgData)
	fmt.Printf("%x\n", msgData)
	fmt.Println(msgData)
	fmt.Println(msgDataB)

	var icmptest N3MsgParser
	icmptest.DecodeMsg()
	//fmt.Printf("%s\n", ICMP_DATA)
	//fmt.Printf("%x\n", ICMP_DATA)
	//fmt.Println( ICMP_DATA)
}

func TestErrdef(t *testing.T) {

	var aErr, bErr, cErr error
	aErr = ErrVersion
	bErr = ErrGTPMsgLen
	cErr = ErrOutSeqIE
	fmt.Println(aErr)
	fmt.Println(bErr)
	fmt.Println(cErr)
	var a uint64 = 1
	switch a {
	case Msg_Type_Echo_Request:
		fmt.Println("Msg_Type_Echo_Request")

	}

}

func TestUint8ToArray(t *testing.T) {
	//srcData := 0x30 //00110000
	srcData := 0x39 //00111001
	fmt.Printf("srcData %08b\n", srcData)
	dstData := ConvBinary(uint8(srcData))
	fmt.Printf("dstData type %T,%p\n", dstData, &dstData)
	fmt.Println(dstData)
	fmt.Println(len(dstData))
	for i := 7; i >= 0; i-- {
		fmt.Printf("%d: %d\n", i, dstData[i])
	}
	fmt.Println()
	//[0 0 0 0 1 1 0 0]
	if dstData[4] != 1 {
		t.Errorf("Wrong result of Data")
	}
	//srcData 00110000
	//dstBinary type [8]uint8,0xc00004c2f0
	//dstData type [8]uint8,0xc00004c2e8
	//[0 0 0 0 1 1 0 0]
}

func TestParse2toInt(t *testing.T) {

	intA, err := Parse2toInt("10000101")
	if err != nil || intA != 133 {
		t.Errorf("Wrong result of Data")
	}
	fmt.Println(intA)

}

func TestSlicefuzhi(t *testing.T) {
	s1 := make([]byte, 10)
	s2 := make([]byte, 10)
	testVar := []byte{1, 2, 3, 4}
	s3 := []byte{}
	var s4 []byte

	s1 = testVar
	s2 = testVar[:3]
	testVar[0] = 9
	fmt.Println(s1, len(s1))
	fmt.Println(s2, len(s2))
	fmt.Println(testVar)
	//[9 2 3 4] 4
	//[9 2 3] 3
	//[9 2 3 4]
	fmt.Println(&s3, s3, len(s3))
	fmt.Printf("s3 addr %p\n", s3)
	fmt.Println(&s4, s4, len(s4))
	fmt.Printf("s4 addr %p\n", s4)
	//&[] [] 0
	//s3 addr 0x97d3d0
	//&[] [] 0
	//s4 addr 0x0 // len()函数是nil指针安全的。
	var b []byte
	a := copy(b, "adb")
	fmt.Println(a)
	// 0

}

// go test -v
func testValidateUser(t *testing.T) {
	_, err := fmt.Println("test 1")
	if err != nil {
		t.Errorf("Wrong result of Println")
	}
}

func TestAll(t *testing.T) {
	t.Run("validate1", testValidateUser)
	t.SkipNow() //跳出以下测试
	t.Run("validate2", testValidateUser)
}

func TestMain(m *testing.M) {
	fmt.Println("Tests begins...") //第一步
	m.Run()                        //第二步,test模块可以运行
}

//go test -bench=.
func BenchmarkAll(b *testing.B) {
	var a int
	for i := 1; i < b.N; i++ {
		a = i
	}
	fmt.Println("Benchmark test finish", a, "b.N", b.N)
}
