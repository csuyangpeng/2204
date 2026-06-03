package types3gpp

import (
	"fmt"
	"strconv"
)

type Tac int

func (p *Tac) StoreString(tac string) error {
	if len(tac) > 3 {
		return fmt.Errorf("invalid tac string(%s)", tac)
	}
	ret, err := strconv.ParseInt("0x"+tac, 0, 64)
	if err != nil {
		return fmt.Errorf("failed to store tac, error(%s)", err)
	}
	*p = Tac(ret)
	return nil
}

type Area struct {
	Tacs []Tac
	//AreaCodes []string // operator specific
	AreaCodes string // operator specific
}
