package types3gpp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

// const definitions
const (
	SizeofTac = 3
	MaxTAINum = 16
)

type TypeOfPlmnList byte

//TAIList : type of list
const (
	OnePlmnMultiTAC   byte = 0x00
	OnePlmnOneTac     byte = 0x01
	MultiPlmnMultiTAC byte = 0x02
)

// TAIList : a list of TAI, <= 16
type TAIList struct {
	Tais []TAI
	// whether contain one or more PLMN, true- multiple, false - single
	MultiplePLMN bool
}

func (p TAIList) String() string {
	var s string
	var taiStr string
	if p.MultiplePLMN == true {
		s += ", Multiple PLMN "
		if len(p.Tais) == 1 {
			taiStr += ", PLMN(" + p.Tais[0].Plmn.String() + ") " + p.Tais[0].Tac.String()
		} else {
			for k, _ := range p.Tais {
				taiStr += p.Tais[k].Plmn.String() + p.Tais[k].Tac.String()
			}
		}
	} else {
		s += ", Single PLMN "
		s += ", PLMN(" + p.Tais[0].Plmn.String() + ")"
		if len(p.Tais) == 1 {
			taiStr += " " + p.Tais[0].Tac.String()
		} else {
			for k, _ := range p.Tais {
				taiStr += " " + p.Tais[k].Tac.String()
			}
		}
	}

	s += taiStr
	return s
}
func (p *TAIList) GetTais() []TAI {
	return p.Tais
}

func (p *TAIList) GetTaisPointers() []*TAI {
	var tList []*TAI
	for k, _ := range p.Tais {
		tList = append(tList, &p.Tais[k])
	}
	return tList
}

func (p *TAIList) GetMultiplePLMN() bool {
	return p.MultiplePLMN
}

func (p *TAIList) SetMultiplePLMN(b bool) {
	p.MultiplePLMN = b
}

func (p *TAIList) GetLen() int {
	return len(p.Tais)
}

func (p *TAIList) AddTai(tai TAI) {
	p.Tais = append(p.Tais, tai)

	if p.MultiplePLMN == false {
		if len(p.Tais) > 0 {
			if p.Tais[0].Plmn != tai.Plmn {
				p.MultiplePLMN = true
			}
		}
	}
}

func (p *TAIList) HasTai(tai TAI) bool {
	for _, v := range p.Tais {
		if v.ComPare(&tai) {
			return true
		}
	}
	return false
}

// If you have any questions about this document, please refer to TS24.501 Figure 9.11.3.9.1
func (p *TAIList) Encode() ([]byte, error) {
	var allTaiOctet []byte
	lenTais := len(p.Tais)

	if lenTais > MaxTAINum {
		return allTaiOctet, fmt.Errorf("wrong Length of TAIList")
	}

	// TAIList may contain one or more PLMN
	// if there is only one PLMN
	if p.MultiplePLMN == false {
		// if this PLMN contains only one Tac

		if lenTais == 1 {
			//refer...TODO
			// octet 1
			typeOfList := OnePlmnOneTac << 5
			octet1 := typeOfList | byte(len((*p).Tais))
			allTaiOctet = append(allTaiOctet, octet1)

			// octet 2 ~ 4
			/*
				if     : allTaiOctet = append(allTaiOctet, (*p).Tais[0].Plmn.GetValue(BigEndian)[:]...)
				error  : "slice of unaddressable Value"
				reason : 因为当(*p).Tais[0].Plmn.GetValue(BigEndian)还没有赋值给任何变量的时候，编译器不知道它的地址，所以编译[:]时会报错
				优先级
				method : 先将(*p).Tais[0].Plmn.GetValue(BigEndian)赋值给某个变量, 然后再对这个变量切片
			*/
			t := (*p).Tais[0].Plmn.GetValue(BigEndian) // only for temporary conversion
			allTaiOctet = append(allTaiOctet, t[:]...)

			// octet 5 ~ 7
			rlogger.Trace(types.ModuleCmn3gtp, rlogger.DEBUG, nil, "(*p).Tais[0].Tac", (*p).Tais[0].Tac)
			allTaiOctet = append(allTaiOctet, (*p).Tais[0].Tac[:]...)

		} else { // if this PLMN contains multi Tac

			// octet 1
			typeOfList := OnePlmnMultiTAC << 5
			octet1 := typeOfList | byte(len((*p).Tais)-1)
			allTaiOctet = append(allTaiOctet, octet1)
			// octet 2 ~ 4
			t := (*p).Tais[0]                // only for temporary conversion
			tt := t.Plmn.GetValue(BigEndian) // only for temporary conversion
			allTaiOctet = append(allTaiOctet, tt[:]...)
			// octet 5 ~ 7 to octet 3k+2* ~ 3k+4*
			// go through each Tac in TAI of this TAIList
			for i := 0; i < len((*p).Tais); i++ {
				t := ((*p).Tais)[i]
				allTaiOctet = append(allTaiOctet, t.Tac[:]...)
			}
		}

	} else {
		// if there is multi PLMN
		// octet 1
		typeOfList := MultiPlmnMultiTAC << 5
		octet1 := typeOfList | byte(len((*p).Tais)-1)
		allTaiOctet = append(allTaiOctet, octet1)

		// go through each Tac in TAI of this TAIList
		for i := 0; i < len((*p).Tais); i++ {
			t := (*p).Tais[i]

			// octet 2 ~ 4 …… 6k-4* ~ 6k-2*
			tt := t.Plmn.GetValue(BigEndian) // only for temporary conversion
			allTaiOctet = append(allTaiOctet, tt[:]...)

			// octet 5 ~ 7 to octet 6k-1* ~  6k+1*
			allTaiOctet = append(allTaiOctet, t.Tac[:]...)
		}
	}

	return allTaiOctet, nil
}

func (t *TAI) Decode(msgBuf *bytes.Reader) error {
	plmnBytes := make([]byte, SizeofPlmnID)

	// octet 2~4
	binary.Read(msgBuf, binary.BigEndian, &plmnBytes)
	err := t.Plmn.SetValue(plmnBytes, BigEndian)
	if err != nil {
		return fmt.Errorf("failed to set decode plmnt")
	}
	// octet 5~6
	binary.Read(msgBuf, binary.BigEndian, &t.Tac)

	return nil
}
