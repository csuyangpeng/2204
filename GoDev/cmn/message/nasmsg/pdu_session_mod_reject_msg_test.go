package nasmsg

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func TestPduSessionModRejectMsg_Encode(t *testing.T) {
	//msg := PduSessionModifyRejectMsg{}
	//msg.MsgHeader.MessageType = nas.PduSessEstabishReject
	//msg.MsgHeader.PrcdTransactionID = 1
	//msg.MsgHeader.PduSessionID = 5
	//
	//msg.SMCause = nas.NotSupportedSSCMode
	//
	//msg.BackOffTimer.Uint = nasie.TenHour
	//msg.BackOffTimer.TimerValue = 1
	//msg.IeFlags.Set(Ieid_PduSessionModReject_BackOffTimer)
	//
	//msg.SMCongestionReAttempInd = true
	//msg.IeFlags.Set(Ieid_PduSessionModReject_SMCongestionReAttempInd)
	//
	//buf, err := msg.Encode()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Printf("Encode Msg: %x\n", buf)
	//fmt.Printf("Encode Msg: %x\n", buf[3:])
	//
	//bufMsg := bytes.NewReader(buf[3:])
	//msg2 := PduSessionModifyRejectMsg{}
	//err = msg2.Decode(bufMsg)
	////fmt.Println(err)
	//fmt.Println(msg2)
	tac := make([]byte, 4)
	binary.LittleEndian.PutUint32(tac, 2)
	fmt.Println(tac)
	fmt.Println(binary.LittleEndian.Uint32(tac))
}
