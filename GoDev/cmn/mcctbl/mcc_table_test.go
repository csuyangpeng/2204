package mcctbl

import (
	"fmt"
	"testing"
)

func TestMccTable_AddItem(t *testing.T) {
	mccTable := &MccTable{}

	err := mccTable.AddItem("460", 2)
	if err != nil {
		t.Errorf("failed to add item, error: %s", err)
		return
	}

	expStr := "mcc table info - { 460 : 2 }"
	rtStr := fmt.Sprint(mccTable)

	if expStr != rtStr {
		t.Errorf("exp: %s, but return: %s", expStr, rtStr)
	}
}

func TestMccTable_GetMncLen(t *testing.T) {
	mccTable := &MccTable{}

	_ = mccTable.AddItem("460", 2)
	_ = mccTable.AddItem("461", 3)
	_ = mccTable.AddItem("462", 3)
	_ = mccTable.AddItem("463", 2)

	var exp uint8 = 2
	rt, err := mccTable.GetMncLen("463")
	if err != nil {
		t.Errorf("failed to get mnc len. error: %s", err)
		return
	}
	if exp != rt {
		t.Errorf("exp: %d, but return: %d", exp, rt)
	}
}
