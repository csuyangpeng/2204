package manager

import (
	"fmt"
	"testing"
)

func TestHandleUeAuthenticatonGetReqMsg(t *testing.T) {

	//var imsi types3gpp.Imsi
	//imsi.StoreImsiString("460000234560001", 2)
	//
	//suci := &types3gpp.Suci{}
	//suci.SetFromImsi(&imsi)
	//
	//dbmgr.DbInit()
	//
	//udmAgent := udmAgentLayer.NewUdmAgentLayer()
	//
	//ctxt := context.WithValue(context.Background(), types.UdmAgentCK, udmAgent)
	//
	//err, supi, heav := HandleUeAuthenticatonGetReqMsg(ctxt, suci, "5gc:460001", nil)
	//if err != nil {
	//	fmt.Println("failed to handle ue auth get req msg", err)
	//	return
	//}
	//
	//fmt.Println("supi (", supi, ") heav(", heav, ")")

}

func TestIncreaseSQN(t *testing.T) {
	inputSqn := [6]byte{0x00, 0x00, 0x00, 0x00, 0xf6, 0x1f}
	sqn1 := IncreaseSQN(inputSqn)
	sqn2 := IncreaseSQN(sqn1)
	fmt.Printf("Input Sqn %x, Output Sqn1 %x  Sqn2 %x", inputSqn, sqn1, sqn2)
}

func TestSQNinRange(t *testing.T) {
	sqnms := [6]byte{0x00, 0x00, 0x00, 0x00, 0xf6, 0x20}
	sqnhs := [6]byte{0x00, 0x00, 0x00, 0x00, 0xf6, 0x43}

	v := SQNinRange(sqnms, sqnhs, 5, 268435456, 32)
	fmt.Println("result : ", v)
}
