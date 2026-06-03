package utils

import "testing"

func TestIsDigitString_1(t *testing.T) {
	str1 := "123456789"
	rc := IsDigitString(str1)
	exp := true
	if exp != rc {
		t.Errorf("exp: %v, but return: %v", exp, rc)
	}
}
func TestIsDigitString_2(t *testing.T) {
	str1 := "123456a79"
	rc := IsDigitString(str1)
	exp := false
	if exp != rc {
		t.Errorf("exp: %v, but return: %v", exp, rc)
	}
}

func TestIsDigitString_3(t *testing.T) {
	str1 := []byte{1, 2, 3, 4, 5, 6}
	rc := IsDigitString(string(str1))
	exp := false
	if exp != rc {
		t.Errorf("exp: %v, but return: %v", exp, rc)
	}
}

func TestStoreSupportedFeatures(t *testing.T) {
	str := "0x80AB"
	rc, _ := StoreSupportedFeatures(str)
	exp := uint64(32939)
	if exp != rc {
		t.Errorf("exp:%d, but return: %d", exp, rc)
	}

	str = "0x80AB0011"
	rc, _ = StoreSupportedFeatures(str)
	exp = uint64(2158690321)
	if exp != rc {
		t.Errorf("exp:%d, but return: %d", exp, rc)
	}

	str = "0X1234"
	rc, _ = StoreSupportedFeatures(str)
	exp = uint64(4660)
	if exp != rc {
		t.Errorf("exp:%d, but return: %d", exp, rc)
	}
}
