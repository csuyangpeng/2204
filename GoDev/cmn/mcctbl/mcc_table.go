package mcctbl

import (
	"fmt"
	"lite5gc/cmn/utils"
)

const (
	sizeofMcc         = 3
	twoMncLen   uint8 = 2
	threeMncLen uint8 = 3
)

//MccTable, store the mcc - mnc length mapping table
type MccTable map[string]uint8

func (p *MccTable) AddItem(mcc string, mncLen uint8) error {
	err := IsValidMccString(mcc)
	if err != nil {
		return fmt.Errorf("invalid mcc string, error:%s", err)
	}
	if (mncLen != twoMncLen) && (mncLen != threeMncLen) {
		return fmt.Errorf("invalid mnc length, shoulb be"+
			" %d or %d, but receive %d", twoMncLen, threeMncLen, mncLen)
	}

	if _, ok := (*p)[mcc]; !ok {
		(*p)[mcc] = mncLen
	} else {
		return fmt.Errorf("already exist in mcc table")
	}

	return nil
}

func (p *MccTable) GetMncLen(mcc string) (uint8, error) {
	err := IsValidMccString(mcc)
	if err != nil {
		return 0, fmt.Errorf("invalid mcc string, error:%s", err)
	}

	if v, ok := (*p)[mcc]; ok {
		return v, nil
	} else {
		return 0, fmt.Errorf("failed to find mcc(%s) in table", mcc)
	}
}

func (p *MccTable) String() string {
	var tblStr string
	for key, val := range *p {
		tblStr += fmt.Sprintf("%s : %d ", key, val)
	}
	return "mcc table info - { " + tblStr + "}"
}

func IsValidMccString(mcc string) error {
	if len(mcc) != sizeofMcc {
		return fmt.Errorf("invalid mcc, length "+
			"should be %d, but get %d", sizeofMcc, len(mcc))
	}

	if !utils.IsDigitString(mcc) {
		return fmt.Errorf("invalid mcc, not digit string")
	}

	return nil
}
