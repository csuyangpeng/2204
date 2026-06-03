package utils

import (
	"fmt"
	"testing"
)

func TestUpdateRestartCounter(t *testing.T) {

	UpdateRestartCounter("amf_rst.dat")

	count := GetAmfRestartCounter()
	fmt.Println(count)

}
