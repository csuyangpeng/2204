package nasie

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

type Add5GSecInfo struct {
	Hdp   bool //Horizontal derivation parameter
	Rinmr bool //Retransmission of initial NAS message request
}

func (p *Add5GSecInfo) Reset() {
	p.Hdp = false
	p.Rinmr = false
}

func (p Add5GSecInfo) String() string {
	return fmt.Sprintf("HDP(%v),RINMR(%v)", p.Hdp, p.Rinmr)
}

// TLV 24.501 9.11.3.12
func (p *Add5GSecInfo) Encode() []byte {
	rlogger.FuncEntry(types.ModCmn, nil)
	encBuf := make([]byte, 3)

	//T
	typeOctec := byte(IeiAdd5GSecInfo)
	encBuf = append(encBuf, typeOctec)
	//L
	encBuf = append(encBuf, byte(1))
	//V
	var val byte
	if p.Hdp == true {
		val = val | 0x01
	}
	if p.Rinmr == true {
		val = val | 0x02
	}
	encBuf = append(encBuf, val)

	return encBuf
}
