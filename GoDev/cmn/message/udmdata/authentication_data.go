package udmdata

import (
	"fmt"
	"lite5gc/cmn/types"
)

type AuthData struct {
	Ki    [types.KSize]byte
	IsOpc bool
	Opc   [types.OpcSize]byte
	Op    [types.OpSize]byte
	Amf   [types.AmfSize]byte
}

func (p AuthData) String() string {
	return fmt.Sprintf("KI(%x),IsOpc(%v),Opc(%x),Op(%x),Amf(%x)", p.Ki, p.IsOpc, p.Opc, p.Op, p.Amf)
}
