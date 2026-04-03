package nas

import (
	"bytes"
	"context"
)

type CmnNasInfo struct {
	N2ConnDataId    uint32
	ConnectMsg      bool
	NasCounter      uint8
	IntegrityStatus bool
}

//NasMessage
type NasMessage interface {
	ExecuteMessage(ctx context.Context, msgBuf *bytes.Reader)
}

func (p *CmnNasInfo) Reset() {
	p.N2ConnDataId = 0
	p.ConnectMsg = false
	p.NasCounter = 0
	p.IntegrityStatus = false
}
