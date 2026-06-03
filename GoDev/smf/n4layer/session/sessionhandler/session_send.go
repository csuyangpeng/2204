package sessionhandler

import (
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

func sendSessionMsg(msg interface{}, header pfcp.PfcpHeaderforSession, n *pfcpv1.Node) error {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	//switch request := msg.(type) {
	//case pfcp.SessionEstablishmentRequest:
	/*if n.State != pfcpv1.NodeActive {
		return fmt.Errorf(strconv.Itoa(int(n.State)))
	}*/
	encode := &pfcpv1.Message{}
	// Encoding message filling
	encode.HeaderSet(header)
	encode.BodySet(msg)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "send Session Msg:%s", encode.String())
	//fmt.Println(msg.String())
	// pfcp encode
	data, err := encode.Marshal()
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
			"Pfcp msg marshal err %s,send msg:%s", err, encode.String())
		return fmt.Errorf("Pfcp msg marshal err %s,send msg:%s", err, encode.String())
	}
	err = n.SendUdpMsg(data)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to send session msg,error:%s,msg:%s", err, encode.String())
		return err
	}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
		"success to send Session Msg:<%s>--><%s>: %#x\n", n.Server.UdpConn.LocalAddr(), n.PeerAddr, data)
	return nil
	/*case pfcp.SessionModifyRequest:
		//	节点发送请求消息
		if n.State != pfcpv1.NodeActive {
			return fmt.Errorf(strconv.Itoa(int(n.State)))
		}
		msg := &pfcpv1.Message{}
		// Encoding message filling
		msg.HeaderSet(request.PfcpHeader)
		msg.BodySet(request)
		//fmt.Println(msg.String())
		// pfcp encode
		data, err := msg.Marshal()
		if err != nil {
			return fmt.Errorf("Pfcp msg marshal err %s", err)
		}
		err = n.SendUdpMsg(data)
		if err != nil {
			return err
		}
		return nil
	case pfcp.SessionReleaseRequest:
		//	节点发送请求消息
		if n.State != pfcpv1.NodeActive {
			return fmt.Errorf(strconv.Itoa(int(n.State)))
		}
		msg := &pfcpv1.Message{}
		// Encoding message filling
		msg.HeaderSet(request.PfcpHeader)
		msg.BodySet(request)
		//fmt.Println(msg.String())
		// pfcp encode
		data, err := msg.Marshal()
		if err != nil {
			return fmt.Errorf("Pfcp msg marshal err %s", err)
		}
		err = n.SendUdpMsg(data)
		if err != nil {
			return err
		}
		return nil
	case pfcp.SessionReportResponse:
		// 转发会话响应消息
		if n.State != pfcpv1.NodeActive {
			return fmt.Errorf(strconv.Itoa(int(n.State)))
		}
		msg := &pfcpv1.Message{}
		// Encoding message filling
		msg.HeaderSet(request.PfcpHeader)
		msg.BodySet(request)
		//fmt.Println(msg.String())
		// pfcp encode
		data, err := msg.Marshal()
		if err != nil {
			return fmt.Errorf("Pfcp msg marshal err %s", err)
		}
		err = n.SendUdpMsg(data)
		if err != nil {
			return err
		}
	default:
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "service message type error")
		return nil
	}*/
	//return nil
}
