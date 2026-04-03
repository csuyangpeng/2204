package nasie

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

const (
	MaxnumofTAIinList      = 16
	MinLengthofSerAreaList = 6
	MaxLengthofSerAreaList = 114

	OnePlmnMultiTAC    byte = 0x00
	OnePlmnOneTac      byte = 0x01
	MultiPlmnMultiTAC  byte = 0x02
	OnePlmnAllAllowTAC byte = 0x03
)

type TypeString byte

const (
	AllowArea    TypeString = 0
	NonAllowArea TypeString = 1
)

func (p TypeString) String() string {
	if p == 0 {
		return "AllowArea"
	} else {
		return "NonAllowArea"
	}
}

type ServiceAreaList struct {
	SerAreaListType TypeString
	AreaList        types3gpp.TAIList
}

func (p ServiceAreaList) String() string {
	var s string
	s += p.SerAreaListType.String()
	s += p.AreaList.String()
	return s
}
func (p *ServiceAreaList) Encode() ([]byte, error) {

	var allTaiOctet []byte

	lenTais := p.AreaList.GetLen()
	if lenTais > MaxnumofTAIinList {
		return allTaiOctet, fmt.Errorf("wrong length of TAIList")
	}

	// TAIList may contain one or more PLMN
	// if there is only one PLMN
	if p.AreaList.GetMultiplePLMN() == false {
		// if this PLMN contains only one Tac
		switch lenTais {
		case 0:
			//plmn without TAC, type of list = "11":
			//Allowed type shall be coded as "0" and number of elements shall be ignored, and
			//If allowed type is coded as "1", it shall be interpreted as "0".
			// octet 1
			typeOfList := byte(p.SerAreaListType) << 7
			typeOfList |= OnePlmnAllAllowTAC << 5
			var numOfTAC byte
			if p.SerAreaListType == AllowArea {
				numOfTAC = 0
			} else {
				numOfTAC = 1
			}
			octet1 := typeOfList | numOfTAC
			allTaiOctet = append(allTaiOctet, octet1)
			// octet 2 ~ 4
			/*
				if     : allTaiOctet = append(allTaiOctet, (*p).tais[0].Plmn.GetValue(BigEndian)[:]...)
				error  : "slice of unaddressable value"
				reason : 因为当(*p).tais[0].Plmn.GetValue(BigEndian)还没有赋值给任何变量的时候，编译器不知道它的地址，所以编译[:]时会报错
				优先级
				method : 先将(*p).tais[0].Plmn.GetValue(BigEndian)赋值给某个变量, 然后再对这个变量切片
			*/
			t := (*p).AreaList.GetTais()[0].Plmn.GetValue(types3gpp.BigEndian) // only for temporary conversion
			allTaiOctet = append(allTaiOctet, t[:]...)
		case 1:
			//refer...TODO
			// octet 1
			typeOfList := byte(p.SerAreaListType) << 7
			typeOfList |= OnePlmnOneTac << 5
			octet1 := typeOfList | byte((*p).AreaList.GetLen())
			allTaiOctet = append(allTaiOctet, octet1)

			// octet 2 ~ 4
			t := (*p).AreaList.GetTais()[0].Plmn.GetValue(types3gpp.BigEndian) // only for temporary conversion
			allTaiOctet = append(allTaiOctet, t[:]...)

			// octet 5 ~ 7
			allTaiOctet = append(allTaiOctet, (*p).AreaList.GetTais()[0].Tac[:]...)
		default:
			// if this PLMN contains multi Tac
			// octet 1
			typeOfList := byte(p.SerAreaListType) << 7
			typeOfList |= OnePlmnMultiTAC << 5
			octet1 := typeOfList | byte((*p).AreaList.GetLen()-1)
			allTaiOctet = append(allTaiOctet, octet1)

			// octet 2 ~ 4
			// plmn
			t := (*p).AreaList.GetTais()[0]            // only for temporary conversion
			tt := t.Plmn.GetValue(types3gpp.BigEndian) // only for temporary conversion
			allTaiOctet = append(allTaiOctet, tt[:]...)

			// octet 5 ~ 7 to octet 3k+2* ~ 3k+4*
			// go through each Tac in TAI of this TAIList
			for i := 0; i < (*p).AreaList.GetLen(); i++ {
				t := ((*p).AreaList.GetTais())[i]
				allTaiOctet = append(allTaiOctet, t.Tac[:]...)
			}
		}

	} else {
		// if there is multi PLMN
		// todo
		rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "multi PLMN not support yet")
	}
	return allTaiOctet, nil
}
