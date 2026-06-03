package nasie

import (
	"fmt"
	"testing"
)

func TestPduSessionModReqMsg_Encode(t *testing.T) {
	var p BitRate
	p.Value = 1
	p.Uint = Pbps1
	fmt.Println(p.Tokbps())
}
