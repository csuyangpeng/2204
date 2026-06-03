package nasmsg

import (
	"fmt"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func TestSecurityModeCommandMsg_Encode(t *testing.T) {
	msg := SecurityModeCommandMsg{}

	msg.SelectNasSecAlg.SetNrIntPrctAlgo(types3gpp.NIA0)
	msg.SelectNasSecAlg.SetNrEncAlgo(types3gpp.NEA0)

	msg.NgKSI.Tsc = false
	msg.NgKSI.Ksi = 0

	msg.UeSecCap.SetNrEncAlgo(types3gpp.NEA0)
	msg.UeSecCap.SetNrIntPrctAlgo(types3gpp.NIA0)

	msg.ImeiSvReq = true
	msg.IeFlags.Set(IeidSecmodcmdImeireq)

	encBuf, err := msg.Encode()
	if err != nil {
		t.Errorf("failed to encode security mode command msg")
	}
	fmt.Printf("Encode Msg: %x\n", encBuf)
	//7e005d0000028080e1

}
