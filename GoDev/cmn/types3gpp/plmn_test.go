package types3gpp

import (
	"fmt"
	"testing"
)

func TestSetValue_1(t *testing.T) {
	plmn := PlmnID{}

	val := []byte{0x46, 0x0f, 0x01}

	plmn.SetValue(val, LittleEndian)

	outStr := fmt.Sprintf("%s", plmn)

	expectStr := "46001"

	if outStr != expectStr {
		t.Errorf("expect %s, but get %s", expectStr, outStr)
	}
}

func TestSetValue_2(t *testing.T) {
	plmn := PlmnID{}

	val := []byte{0x64, 0xf0, 0x10}

	plmn.SetValue(val, BigEndian)

	outStr := plmn.String()

	expectStr := "46001"

	if outStr != expectStr {
		t.Errorf("expect %s, but get %s", expectStr, outStr)
	}
}

func TestGetValue_1(t *testing.T) {
	plmn := PlmnID{}
	err := plmn.SetString("46001")
	if err != nil {
		t.Errorf("failed to set string.")
	}

	expect := [3]byte{0x46, 0x0f, 0x01}

	val := plmn.GetValue(LittleEndian)

	if val != expect {
		t.Errorf("expect %v, but get %v", expect, val)
	}
}

func TestGetValue_2(t *testing.T) {
	plmn := PlmnID{}
	err := plmn.SetString("460123")
	if err != nil {
		t.Errorf("failed to set string.")
	}

	expect := [3]byte{0x46, 0x03, 0x12}

	val := plmn.GetValue(LittleEndian)

	if val != expect {
		t.Errorf("expect %v, but get %v", expect, val)
	}
}

func TestGetValue_3(t *testing.T) {
	plmn := PlmnID{}
	err := plmn.SetString("460123")
	if err != nil {
		t.Errorf("failed to set string.")
	}
	expect := [3]byte{0x64, 0x30, 0x21}

	val := plmn.GetValue(BigEndian)
	if val != expect {
		t.Errorf("expect %v, but get %v", expect, val)
	}
}

func TestGetValue_4(t *testing.T) {
	plmn := PlmnID{}
	err := plmn.SetString("00101")
	if err != nil {
		t.Errorf("failed to set string.")
	}

	fmt.Println("plmn: ", plmn.String())
	expect := [3]byte{0x00, 0x1f, 0x01}

	val := plmn.GetValue(LittleEndian)

	if val != expect {
		t.Errorf("expect %v, but get %v", expect, val)
	}
}
