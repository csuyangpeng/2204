package types3gpp

import (
	"fmt"
	"testing"
)

func TestTAC_SetTac(t *testing.T) {
	var tac TAC

	tac.SetTac(1)
	rtStr := tac.String()

	expStr := "TAC(1)"

	if expStr != rtStr {
		t.Errorf("exp: %s, but return :%s", expStr, rtStr)
	}
}

func TestTAI_String(t *testing.T) {
	var tai TAI
	tai.Plmn.SetString("46000")
	tai.Tac.SetTac(512)

	rtStr := fmt.Sprintf("%s", &tai)
	expStr := "tai[46000-tac(512)]"

	if expStr != rtStr {
		t.Errorf("exp: %s, but return :%s", expStr, rtStr)
	}
}
