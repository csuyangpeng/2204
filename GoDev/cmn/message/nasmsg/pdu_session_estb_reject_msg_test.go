package nasmsg

import (
	"fmt"
	"lite5gc/cmn/nas"
	"testing"
)

func TestPduSessionEstbRejectMsg_Encode(t *testing.T) {
	msg := PduSessionEstbRejectMsg{}
	msg.MsgHeader.MessageType = nas.PdqSessEstablishReject

	msg.SMCause = nas.NotSupportedSSCMode

	msg.AllowedSSCMode.Ssc1 = true
	msg.AllowedSSCMode.Ssc2 = false
	msg.AllowedSSCMode.Ssc3 = false
	msg.IeFlags.Set(Ieid_PduSessionEstbReject_AllowedSSCMode)

	msg.Print()

	buf, err := msg.Encode()
	if err != nil {
		fmt.Println(err)
	}
	//out := fmt.Sprintf("%x", buf)
	//expect := "7e0056000200002147114711471147114711471147114711201011495219c670b9b9d8b9676b45dad5fc"
	//
	fmt.Printf("Encode Msg: %x\n", buf)
	//fmt.Printf("Expect msg: %s\n", expect)
	//
	//if out != expect {
	//	t.Errorf("failed to encode.")
	//}
}
