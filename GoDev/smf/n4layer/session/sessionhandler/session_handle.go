package sessionhandler

import (
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/n4layer/session/create"
	"lite5gc/smf/smfcontext/gctxt"
)

func SmfSessionHandle(n *pfcpv1.Node, msg pfcpv1.ServiceMsg) (*pfcpv1.ServiceMsg, error) {
	rlogger.FuncEntry(types.ModuleSmfN4, nil)
	// 创建N4会话
	/*scMsg, ok := msg.Msg.(pfcpv1.SmfToNode)
	if !ok {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "service message type error")
		return nil, fmt.Errorf("service message type error")
	}*/
	// 适配消息
	//var message interface{}
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil,
		"node:%s,Session message processing(%v)", n.NodeID, msg.Msg.ID)
	if msg.Msg.Cxt == nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to SessionMsg:SmfToNode cxt is nil")
		return nil, nil
	}
	switch msg.Msg.ID {
	case pfcp.PFCP_Session_Establishment_Request:
		if msg.Msg.Cxt == nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
				"msg.Msg.Cxt is nil (%+v)", msg.Msg.Cxt)
			return nil, fmt.Errorf("failed to ScN4Msg is nil")
		} else {
			// 1 构造消息
			reqN4, err := create.SessionEstablishmentRequest(msg.Msg.Cxt, n)
			if err != nil {
				res := &pfcpv1.ServiceMsg{}
				res.Msg.ID = pfcp.PFCP_Session_Establishment_Response
				res.Msg.Cxt = msg.Msg.Cxt
				res.Msg.Cxt.Cause = gctxt.Cause_System_failure
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to create SessionMsg:%s", err)
				return res, err
			}

			// 3 发送成功，保存上下文, TS 23.502 4.4.1.2	N4 Session Establishment procedure
			n4Cxt := &gctxt.N4SessionContext{}
			//n4Cxt.PduCxt = msg.Msg.Cxt //for paging procedure
			create.StoreSmfN4Context(reqN4, n4Cxt)
			// 返回SEID给UE context
			//rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, msg.Msg.Cxt.GetPrcdCtxt())
			//pCtxt, ok := msg.Msg.Cxt.GetPrcdCtxt().(*prcdctxt.PduSessionEstbPrcdCtxt)
			//if ok {
			//	pCtxt.Seid = n4Cxt.SEID
			//} else {
			//	rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,
			//		"prcdctxt is not PduSessionEstbPrcdCtxt,but:",
			//		gctxt.PrcdCtxt2String(msg.Msg.Cxt.GetPrcdCtxt()))
			//	return nil, fmt.Errorf("procedure context is not PduSessionEstbPrcdCtxt")
			//}

			//TODO should set here???
			msg.Msg.Cxt.SEID = n4Cxt.SEID

			// 2 发送消息
			err = sendSessionMsg(reqN4, reqN4.PfcpHeader, n)
			if err != nil {
				rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to send SessionMsg:%s", err)
			}
			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "success to send SessionMsg:%d", reqN4.PfcpHeader.MessageType)
		}

	case pfcp.PFCP_Session_Modification_Request:
		// 1 构造消息
		reqN4, err := create.SessionModificationRequest(msg.Msg.Cxt, n)
		if err != nil {
			res := &pfcpv1.ServiceMsg{}
			res.Msg.ID = pfcp.PFCP_Session_Modification_Response
			res.Msg.Cxt = msg.Msg.Cxt
			res.Msg.Cxt.Cause = gctxt.Cause_System_failure
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to create SessionMsg:%s", err)
			return res, err
		}
		// 2 发送消息
		err = sendSessionMsg(reqN4, reqN4.PfcpHeader, n)
		if err != nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to send SessionMsg:%s", err)
		}
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "success to send SessionMsg:%d", reqN4.PfcpHeader.MessageType)
		// 3 发送成功，保存上下文, TS 23.502 4.4.1.2	N4 Session Establishment procedure
		create.UpdateSmfN4ContextV1(reqN4, msg.Msg.Cxt.SEID)

	case pfcp.PFCP_Session_Deletion_Request:
		// 1 构造消息
		reqN4, err := create.SessionReleaseRequest(msg.Msg.Cxt.SEID, n)
		if err != nil {
			res := &pfcpv1.ServiceMsg{}
			res.Msg.ID = pfcp.PFCP_Session_Deletion_Request
			res.Msg.Cxt = msg.Msg.Cxt
			res.Msg.Cxt.Cause = gctxt.Cause_System_failure
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to create SessionMsg:%s", err)
			return res, err
		}
		// 2 发送消息
		err = sendSessionMsg(reqN4, reqN4.PfcpHeader, n)
		if err != nil {
			rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "failed to send SessionMsg:%s", err)
		}
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, nil, "success to send SessionMsg:%d", reqN4.PfcpHeader.MessageType)
	case pfcp.PFCP_Session_Report_Response:
		// 1 构造消息
	default:
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil, "service message type error")
		return nil, nil
	}
	return nil, nil
}
