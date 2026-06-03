package create

import (
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/utils"
	"lite5gc/smf/smfcontext/gctxt"
)

// TS 23.502 4.4.1.4	N4 Session Release procedure
/*
1.	SMF receives the trigger to remove the N4 session context for the PDU Session.
2.	The SMF sends an N4 session release request message to the UPF.
3.	The UPF identifies the N4 session context to be removed by the N4 Session ID and removes the whole session context.
The UPF responds with an N4 session release response message containing any information that the UPF has to provide to the SMF.
4.	The SMF interacts with the network entity which triggered this procedure (e.g. AMF or PCF).
*/
func SessionReleaseRequest(seid uint64, n *pfcpv1.Node) (*pfcp.SessionReleaseRequest, error) {
	rlogger.FuncEntry(types.ModuleSmfN4,nil)
	/*if smfUeCxt == nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,  "Input parameter check failed")
		return errors.New("smfcontext is nil") // todo err 返回一个error对象中可以增加错误码，对应具体的错误
	}
	// get SEID
	pduSessContext := smfUeCxt.GetPduSessCtxt(psi)
	if pduSessContext == nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, nil,  "Failed to get pdu session Context:%s")
		return errors.New("pduSessContext is nil")
	}

	seid := pduSessContext.SEID*/
	// get N4 context
	n4Ctxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n4Ctxt, "Failure to get N4 Context:%s", err)
		// 本地无上下文，返回成功
		return nil, err
	}

	// 2.	The SMF sends an N4 session release request message to the UPF.
	//*pfcp.SessionReleaseRequest
	reqN4, err := CreateN4ReleaseRequestMsgV1(n4Ctxt, n)
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, n4Ctxt, "Failed to Create N4 session release request message :%s", err)
		return nil, err
	}

	// 返回成功
	return reqN4, nil
}

// IEsSessionDelRequest
func CreateN4ReleaseRequestMsgV1(n4Ctxt *gctxt.N4SessionContext, n *pfcpv1.Node) (
	*pfcp.SessionReleaseRequest, error) {

	var reqN4 pfcp.SessionReleaseRequest
	// 获取UPF的SEID
	reqN4.PfcpHeader.Set()
	reqN4.PfcpHeader.MessageType = pfcp.PFCP_Session_Deletion_Request
	reqN4.PfcpHeader.SEID = n4Ctxt.UpfSEID.SEID
	utils.Uint32AtomicAdd(&n.SequenceNumber, 1)
	reqN4.PfcpHeader.SequenceNumber = n.SequenceNumber

	// 2 请求消息构造
	reqN4.IE = pfcp.IEsSessionDelRequest{}

	return &reqN4, nil
}
