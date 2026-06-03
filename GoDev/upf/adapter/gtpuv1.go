package adapter

import (
	"fmt"
	"lite5gc/cmn/message/gtpv1u"
	"lite5gc/upf/context/pdrcontext"
)

func SendN3MsgHandleExt(PayloadMsgbuf []byte, Msgcxt *pdrcontext.DataFlowContext) ([]byte, error) {
	Msgcxt.Rw.RLock()
	if Msgcxt == nil {
		Msgcxt.Rw.RUnlock()
		//rlogger.Trace(moduleTag, types.ERROR, nil, "Input parameter check failed !")
		return nil, fmt.Errorf("Input parameter check failed !")
	}
	// Msgbuf 转换为GTPPDU，N6侧消息为Msg_Type_G_PDU，直接加GTPv1U协议头发送到N3侧
	// 需要获取RAN地址与通道信息
	N3EncodeMsg := &gtpv1u.GPDUSessionContDL{}

	N3EncodeMsg.Gtpbody = PayloadMsgbuf
	//填充GTPv1U协议头部
	N3EncodeMsg.Version = gtpv1u.Protocol_version
	N3EncodeMsg.PT = gtpv1u.Protocol_Type
	N3EncodeMsg.MessageType = gtpv1u.Msg_Type_G_PDU
	N3EncodeMsg.TEID = uint32(Msgcxt.GnbTEID) //0x80000000 //0x04180155

	// Msg_Type_G_PDU PDU_Type_DL_PDU_Session_Information
	N3EncodeMsg.EFlag = gtpv1u.Protocol_Present
	N3EncodeMsg.NextExtHeaderType = gtpv1u.ExtHT_PDU_SESSION_CONTAINER

	//N3EncodeMsg.PDUSessionContainer.DLPDUSession = Msgcxt.DP
	N3EncodeMsg.PDUSessionContainer.Length = 2
	N3EncodeMsg.PDUSessionContainer.PDUType = gtpv1u.PDU_Type_DL_PDU_Session_Information
	N3EncodeMsg.PDUSessionContainer.QFI = Msgcxt.DP.QFI //9 //10
	N3EncodeMsg.PDUSessionContainer.PPP = Msgcxt.DP.PPP
	N3EncodeMsg.PDUSessionContainer.PPI = Msgcxt.DP.PPI
	N3EncodeMsg.PDUSessionContainer.NextExtHeaderType = 0
	/*Octets 1  Extension Header Length
	2 – m		Extension Header Content
	m+1		     Next Extension Header Type
	*/
	Msgcxt.Rw.RUnlock()
	N3EncodeMsg.Length = (uint16(len(PayloadMsgbuf)) +
		uint16(gtpv1u.GTPV1_U_HEADER_OPTIONAL_FIELDS_LEN) +
		uint16(N3EncodeMsg.PDUSessionContainer.Length*4)) // 可选头域9-12+ 扩展头+ payload length
	// test 删除IP与UDP头
	//N3EncodeMsg.Gtpbody = N3EncodeMsg.Gtpbody[28:]
	Msgbuf, err := N3EncodeMsg.EncodeMsg()
	if err != nil {
		//fmt.Println(err)
		//rlogger.Trace(moduleTag, types.ERROR, Msgcxt, "Failed to N6 message Eecode!")
		// N3DecodeMsg failed,discard message
		return Msgbuf, err
	}
	//fmt.Println(len(Msgbuf))
	//fmt.Printf("Encode value: %#x\n", Msgbuf)
	//rlogger.Trace(moduleTag, types.DEBUG, Msgcxt, "Encode value: %#x\n", Msgbuf)
	return Msgbuf, nil
}
