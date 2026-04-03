package sbicmn

import (
	"fmt"
	"lite5gc/cmn/nas"
	"testing"
)

func TestGetSmfKeys(t *testing.T) {

	imsi, psi, err := nas.GetSmfKeys("imsi-460000123456001-5")
	fmt.Println(imsi.String(), psi, err)
}
