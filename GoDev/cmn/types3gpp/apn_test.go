package types3gpp

import (
	"bytes"
	"fmt"
	"testing"
)

func TestApn_StoreWithString(t *testing.T) {
	//expStr1 := "cmnet"
	//expStr2 := "cmnet"
	//expStr3 := "cmnet.shenzhen"
	//expStr4 := "cmnet.shenzhen.mcc460.mnc001.gprs"

	apn := &Apn{}
	apn.StoreWithString("cmnet342")
	fmt.Println(apn.String())

	//apn2 := &Apn{}
	//apn2.StoreWithString("cmnet")
	//outStr = fmt.Sprintf("%s", apn2)
	//if expStr2 != outStr {
	//	t.Errorf("exp(%s),but return(%s)", expStr2, outStr)
	//}
	//
	//apn3 := &Apn{}
	//apn3.StoreWithString("cmnet.shenzhen")
	//outStr = fmt.Sprintf("%s", apn3)
	//if expStr3 != outStr {
	//	t.Errorf("exp(%s),but return(%s)", expStr3, outStr)
	//}
	//
	//apn4 := &Apn{}
	//apn4.StoreWithString("cmnet.shenzhen.mcc460.mnc001.gprs")
	//outStr = fmt.Sprintf("%s", apn4)
	//if expStr4 != outStr {
	//	t.Errorf("exp(%s),but return(%s)", expStr4, outStr)
	//}
}

func TestApn_Encode(t *testing.T) {
	exp1 := []byte{0x05, 0x63, 0x6d, 0x6e, 0x65, 0x74, 0x03, 0x63, 0x6f, 0x6d, 0x06, 0x6d, 0x63, 0x63, 0x34, 0x36, 0x30, 0x05, 0x6d, 0x6e, 0x63, 0x30, 0x30, 0x04, 0x67, 0x70, 0x72, 0x73}
	apn := &Apn{}
	apn.StoreWithString("cmnet.com.mcc460.mnc00.gprs")
	outval := apn.Encode()
	if string(exp1) != string(outval) {
		t.Errorf("\nexp %v\nget %v", exp1, outval)
	}

	exp2 := []byte{0x05, 0x63, 0x6d, 0x6e, 0x65, 0x74}
	apn2 := &Apn{}
	apn2.StoreWithString("cmnet")
	outval = apn2.Encode()
	if string(exp2) != string(outval) {
		t.Errorf("\nexp %v\nget %v", exp2, outval)
	}
}

func TestApn_Decode(t *testing.T) {
	//encApn := []byte{0x05, 0x63, 0x6d, 0x6e, 0x65, 0x74}
	//exp := "cmnet"
	//apn := &Apn{}
	//msg := bytes.NewReader(encApn)
	//apn.Decode(msg)
	//if apn.String() != exp {
	//	t.Errorf("\nexp %v\nget %v", exp, apn.String())
	//}	//encApn := []byte{0x05, 0x63, 0x6d, 0x6e, 0x65, 0x74}
	//exp := "cmnet"
	//apn := &Apn{}
	//msg := bytes.NewReader(encApn)
	//apn.Decode(msg)
	//if apn.String() != exp {
	//	t.Errorf("\nexp %v\nget %v", exp, apn.String())
	//}

	//encApn = []byte{0x05, 0x63, 0x6d, 0x6e, 0x65, 0x74, 0x03, 0x63, 0x6f, 0x6d, 0x06, 0x6d, 0x63, 0x63, 0x34, 0x36, 0x30, 0x05, 0x6d, 0x6e, 0x63, 0x30, 0x30, 0x04, 0x67, 0x70, 0x72, 0x73}
	encApn2 := []byte{0x1d, 0x05, 0x63, 0x6d, 0x6e, 0x65, 0x74, 0x03, 0x63, 0x6f, 0x6d, 0x06, 0x6d, 0x63, 0x63, 0x34, 0x36, 0x30, 0x06, 0x6d, 0x6e, 0x63, 0x30, 0x30, 0x31, 0x04, 0x67, 0x70, 0x72, 0x73}
	exp2 := "cmnet.com.mcc460.mnc001.gprs"
	apn2 := &Apn{}
	msg2 := bytes.NewReader(encApn2)
	apn2.Decode(msg2)
	if apn2.String() != exp2 {
		t.Errorf("\nexp %v\nget %v", exp2, apn2.String())
	}
}
