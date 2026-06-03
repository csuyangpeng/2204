package create_test

import (
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"net"
	"testing"
)

func TestNewFSEID(t *testing.T) {

	for i := 0; i < 100; i++ {
		seid, err := GetSEID()
		ip := net.ParseIP("127.0.0.1")
		NewFSEID(seid, pfcp.IEFSEID_IPv4_address, ip)
		//fmt.Printf("Fully Qualified SEID:%v", fSEID)
		if err != nil {
			t.Errorf("Get new F-SEID failed.")
		}
	}

}

func TestCreatPDR(t *testing.T) {
	//Abs := make([]*pfcp.IECreatePDR,0)
	var Abs []*pfcp.IECreatePDR
	fmt.Println(Abs, len(Abs))
	pdr := pfcp.IECreatePDR{}
	Abs = append(Abs, &pdr)
	fmt.Println(Abs, len(Abs))
}
