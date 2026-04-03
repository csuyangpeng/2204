package types3gpp

import (
	"fmt"
)

// SupiType : type of SUPI
type SupiType uint8

// enum for SUPIType, 2-7 reserved.
const (
	IMSIType    SupiType = 0
	NAIType     SupiType = 1
	InvalidType SupiType = 2
)

// NAI: TS23.003 28.7.2 form : "<IMSI>@nai.5gc.mnc<MNC>.mcc<MCC>.3gppnetwork.org"
// NAI store and retrieve from IMSI
// Supi  : Subscription Permanent Identifier
type Supi struct {
	supiType SupiType
	imsi     Imsi
}

func CreateATestSUPI() Supi {
	imsi := CreateATestIMSI()
	return Supi{supiType: IMSIType, imsi: imsi}
}

func (p *Supi) String() string {
	if p.supiType == IMSIType {
		return p.imsi.String()
	} else {
		return p.imsi.GetNAIString()
	}
}

func (p *Supi) SetType(t SupiType) error {
	if t != IMSIType && t != NAIType {
		return fmt.Errorf("invalid supi type(0-imsi, 1-nai), %d", t)
	}
	p.supiType = t
	return nil
}

func (p *Supi) GetType() SupiType {
	return p.supiType
}

func (p *Supi) SetImsi(imsi *Imsi) {
	p.supiType = IMSIType
	p.imsi = *imsi
}

func (p *Supi) GetImsi() *Imsi {
	return &(p.imsi)
}

func (p *Supi) SetNAI(nai string, mncDigitsLen uint8) error {

	err := p.imsi.StoreWithNAI(nai, mncDigitsLen)
	if err != nil {
		return fmt.Errorf("failed to store nai. error: %s", err)
	}
	p.supiType = NAIType
	return nil
}

func (p *Supi) GetNAI() string {
	return p.imsi.GetNAIString()
}

func (p *Supi) Store(supiStr string, supiType SupiType) {
	p.SetType(supiType)

	if supiType == IMSIType {
		var mnclen uint8 = 2
		if len(supiStr) > 15 {
			mnclen = 3
		}
		imsi := &Imsi{}
		imsi.StoreImsiString(supiStr, mnclen)
		p.SetImsi(imsi)
	}

	return
}
