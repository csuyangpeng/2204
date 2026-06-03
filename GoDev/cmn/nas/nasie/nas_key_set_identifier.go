package nasie

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
)

// 24.501 9.11.3.32
//Type of security context flag (TSC) (octet 1)
//
//Bit
//4
//0			native security context (for KSIAMF)
//1			mapped security context (for KSIASME)
//
//TSC does not apply for NAS key set identifier value "111".
//
//NAS key set identifier (octet 1)
//
//Bits
//3	2	1
//0	0	0
//through	possible values for the NAS key set identifier
//1	1	0
//
//1	1	1	no key is available (UE to network);
//reserved (network to UE)

const (
	MinAllowableKsi_c uint8 = 0x00
	MaxAllowableKsi_c uint8 = 0x06
	ReservedKsi_v     uint8 = 0x07
)
const (
	NativeSecCtxt bool = false
	MappedSecCtxt bool = true
)

type NasKSI struct {
	Ksi uint8
	Tsc bool
}

func (p NasKSI) String() string {
	tsc := "Native"
	if p.Tsc == MappedSecCtxt {
		tsc = "Mapped"
	}
	return fmt.Sprintf("TSC(%s),KSI(%d)", tsc, p.Ksi)
}

func (p *NasKSI) Reset() {
	p.Ksi = 0
	p.Tsc = NativeSecCtxt
}

func (p *NasKSI) Decode(msgBuf *bytes.Reader) error {
	nasKsiSetIdOctet, err := msgBuf.ReadByte()
	if err != nil {
		return fmt.Errorf("failed to read Nas ksi set identifier")
	}
	ksi, _ := utils.GetBitsValue(nasKsiSetIdOctet, 1, 3)
	p.Ksi = uint8(ksi)
	p.Tsc, _ = utils.GetBitValue(nasKsiSetIdOctet, 4)

	return nil
}

func (p *NasKSI) Encode() byte {
	ngKsi := p.Ksi
	if p.Tsc == MappedSecCtxt {
		ngKsi |= 1 << 3 //bit4
	}

	rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ngKSI(%x)", ngKsi)
	return ngKsi
}

func (p *NasKSI) IsValid() bool {
	if p.Ksi >= MinAllowableKsi_c && p.Ksi <= MaxAllowableKsi_c {
		return true
	}
	return false
}

func (p *NasKSI) Update() {
	rlogger.FuncEntry(types.ModCmn, nil)

	val := p.Ksi + 1
	if val > MaxAllowableKsi_c {
		p.Ksi = MinAllowableKsi_c
	} else {
		p.Ksi = val
	}
}
