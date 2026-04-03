package types3gpp

import (
	"fmt"
)

type SelNasSecAlgo struct {
	nrEncryptAlgo NREncryptAlgo
	nrIntPrctAlgo NRIntPrctAlgo
}

func (p SelNasSecAlgo) String() string {
	return fmt.Sprintf("NR Enc Algo (%d), NR Int Algo (%d)", p.nrEncryptAlgo, p.nrIntPrctAlgo)
}

func (p *SelNasSecAlgo) Reset() {
	p.nrIntPrctAlgo = 0
	p.nrIntPrctAlgo = 0
}
func (p *SelNasSecAlgo) SetNrEncAlgo(ind NREncryptAlgo) {
	p.nrEncryptAlgo = ind
}

func (p *SelNasSecAlgo) GetNrEncAlgo() NREncryptAlgo {
	return p.nrEncryptAlgo
}

func (p *SelNasSecAlgo) SetNrIntPrctAlgo(ind NRIntPrctAlgo) {
	p.nrIntPrctAlgo = ind
}

func (p *SelNasSecAlgo) GetNrIntPrctAlgo() NRIntPrctAlgo {
	return p.nrIntPrctAlgo
}

//24.501 9.11.3.34  encode as V
func (p *SelNasSecAlgo) Encode() byte {

	return byte(p.nrEncryptAlgo<<4) | byte(p.nrIntPrctAlgo)
}

//24.501 9.11.3.34
func (p *SelNasSecAlgo) Decode(val byte) error {

	p.nrEncryptAlgo = NREncryptAlgo(val >> 4)
	p.nrIntPrctAlgo = NRIntPrctAlgo(val & 0x0F)

	return nil
}
