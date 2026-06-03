package types3gpp

import (
	"fmt"
	"testing"
)

func TestSecurityCapability_String(t *testing.T) {
	sec := &SecurityCapability{}
	sec.SetNrEncAlgo(NEA0)
	sec.SetNrEncAlgo(NEA1)
	sec.SetNrEncAlgo(NEA2)
	sec.SetNrEncAlgo(NEA3)
	sec.SetNrIntPrctAlgo(NIA0)
	sec.SetNrIntPrctAlgo(NIA1)
	sec.SetNrIntPrctAlgo(NIA2)
	sec.SetNrIntPrctAlgo(NIA3)
	t.Logf("ue security:%s", sec)

	t.Logf("IntPrctAlgo: %v", sec.GetNrIntPrctAlgoU16())
}

func TestSecurityCapability_GetNrAlgoU16(t *testing.T) {
	exp := []byte{0xE0, 0x0}
	sec := &SecurityCapability{}
	sec.SetNrEncAlgo(NEA0)
	sec.SetNrEncAlgo(NEA1)
	sec.SetNrEncAlgo(NEA2)
	sec.SetNrEncAlgo(NEA3)
	sec.SetNrIntPrctAlgo(NIA0)
	sec.SetNrIntPrctAlgo(NIA1)
	sec.SetNrIntPrctAlgo(NIA2)
	sec.SetNrIntPrctAlgo(NIA3)
	val := sec.GetNrEncAlgoU16()
	if string(exp) != string(val) {
		t.Errorf("expect(%v),but retrurn(%v)", exp, val)
	}
	val1 := sec.GetNrIntPrctAlgoU16()
	if string(exp) != string(val1) {
		t.Errorf("expect(%v),but retrurn(%v)", exp, val1)
	}
}

func TestSecurityCapability_GetEutraAlgoU16(t *testing.T) {
	exp := []byte{0xF0, 0x0}
	sec := &SecurityCapability{}
	sec.SetEutraEncAlgo(EEA0)
	sec.SetEutraEncAlgo(EEA1)
	sec.SetEutraEncAlgo(EEA2)
	sec.SetEutraEncAlgo(EEA3)
	sec.SetEutraIntPrctAlgo(EIA0)
	sec.SetEutraIntPrctAlgo(EIA1)
	sec.SetEutraIntPrctAlgo(EIA2)
	sec.SetEutraIntPrctAlgo(EIA3)
	val := sec.GetEutraEncAlgoU16()
	if string(exp) != string(val) {
		t.Errorf("expect(%v),but retrurn(%v)", exp, val)
	}
	val1 := sec.GetEutraIntPrctAlgoU16()
	if string(exp) != string(val1) {
		t.Errorf("expect(%v),but retrurn(%v)", exp, val1)
	}
}

func TestSecurityCapability_SetEutraEncAlgoU16(t *testing.T) {
	input := []byte{0xF0, 0x0}
	sec := &SecurityCapability{}
	sec.SetNrEncAlgoU16(input)
	sec.SetNrIntPrctAlgoU16(input)
	sec.SetEutraEncAlgoU16(input)
	sec.SetEutraIntPrctAlgoU16(input)
	fmt.Println(sec)
}

func TestSecurityCapability_SetNrEncAlgoStr(t *testing.T) {
	sec := &SecurityCapability{}

	sec.SetNrIntAlgoStr("NIA0")
	fmt.Println(sec)
}

func TestSecurityCapability_SetNrIntAlgoStr(t *testing.T) {
	sec := &SecurityCapability{}

	sec.SetNrEncAlgoStr("NEA2")
	fmt.Println(sec)
}

func TestSecurityCapability_StoreNrIntAlgo(t *testing.T) {
	sec := &SecurityCapability{}
	sec.StoreNrIntAlgo("NIA0,NIA1,NIA2,NIA3")
	fmt.Println(sec)
}

func TestSecurityCapability_StoreNrEncAlgo(t *testing.T) {
	sec := &SecurityCapability{}
	sec.StoreNrEncAlgo("NEA0,NEA1,NEA2,NEA3")
	fmt.Println(sec)
}

func TestSecurityCapability_MatchNrEncAlgo(t *testing.T) {
	sec := &SecurityCapability{}
	sec.StoreNrEncAlgo("NEA0,NEA2")

	uesec := &SecurityCapability{}
	uesec.StoreNrEncAlgo("NEA0,NEA1,NEA2,NEA3")

	rt := sec.MatchNrEncAlgo(uesec.GetNrEncAlgo())
	fmt.Println(rt)
}

func TestSecurityCapability_MatchNrIntAlgo(t *testing.T) {
	sec := &SecurityCapability{}
	sec.StoreNrIntAlgo("NIA0,NIA2,")

	uesec := &SecurityCapability{}
	uesec.StoreNrIntAlgo("NIA0,NIA1,NIA2,NIA3")

	rt := sec.MatchNrIntAlgo(uesec.GetNrIntPrctAlgo())
	fmt.Println(rt)
}
