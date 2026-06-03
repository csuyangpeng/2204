package types3gpp

import (
	"fmt"
)

const (
	MinSizeOfPlmnList = 1
	MaxSizeOfPlmnList = 15
)

type PlmnList struct {
	List []PlmnID
}

func (p *PlmnList) AddPlmn(plmn PlmnID) {
	p.List = append(p.List, plmn)
}
func (p *PlmnList) GetPlmn(index uint) *PlmnID {
	return &p.List[index]
}
func (p *PlmnList) Encode() ([]byte, error) {
	var allPlmnOctet []byte
	if len(p.List) > MaxSizeOfPlmnList {
		return allPlmnOctet, fmt.Errorf("wrong length of PlmnList")
	}
	for i := 0; i < len(p.List); i++ {
		plmnValue := p.List[i].GetValue(BigEndian)
		allPlmnOctet = append(allPlmnOctet, plmnValue[:]...)
	}
	return allPlmnOctet, nil
}
