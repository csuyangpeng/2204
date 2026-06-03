package types3gpp

import "testing"

func Test_SetAmfSetIdU16_1(t *testing.T) {
	amfIdentifier := AmfIdentifier{}
	amfIdentifier.SetAmfRegionID(10)
	amfIdentifier.SetAmfPointer(5)
	amfIdentifier.SetAmfSetIdU16(0x1234)

	regionId := amfIdentifier.GetAmfRegionID()
	pointer := amfIdentifier.GetAmfPointer()
	setId := amfIdentifier.GetAmfSetIdU16()
	if regionId != 10 {
		t.Errorf("expect %v, but get %v", 10, regionId)
	}
	if pointer != 5 {
		t.Errorf("expect %v, but get %v", 5, pointer)
	}
	if setId != 0x1234 {
		t.Errorf("expect %v, but get %v", 0x1234, setId)
	}

	expect := AmfIdentifier{
		RegionId: 10,
		Pointer:  5,
		SetId:    [2]uint8{0x12, 0x34},
	}

	if amfIdentifier != expect {
		t.Errorf("expect %v, but get %v", expect, amfIdentifier)
	}

	//t.Logf("amfIdentifier: %v", amfIdentifier)
}
