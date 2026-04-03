package n4layer

import (
	"fmt"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/pfcpudp"
	"lite5gc/cmn/message/pfcp/utils"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types/configure"
	"lite5gc/upf/context/n4context"
	"lite5gc/upf/context/pdrcontext"
	"net"
)

// 作为client时，发起N4请求
// N4 Data Notification --> PFCP Session Report Request
func SessionReportRequest(n4Ctxt *n4context.N4SessionContext, dfCxt *pdrcontext.DataFlowContext) error {
	rlogger.FuncEntry(moduleTag, n4Ctxt)

	/*// get N4 context
	n4Ctxt, err := upfcontext.GetN4Context(upfcontext.N4SessionIDKey(dfCxt.SEID))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil,  "Failure to get N4 Context:%s", err)
		// 本地无上下文，返回成功
		return nil
	}*/

	// 2.	The UPF sends an N4 session Report request message to the SMF.
	//*pfcp.SessionReleaseRequest
	reqN4, err := CreateReportRequestMsg(n4Ctxt, dfCxt)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Ctxt, "Failed to Create N4 session release request message :%s", err)
		return err
	}
	// 3 发送请求消息，异步处理响应
	var resN4 pfcp.SessionReportResponse
	peerAddr := &net.UDPAddr{IP: n4Ctxt.SmfSEID.IPv4Addr,
		Port: configure.UpfConf.N4.Smf.Port}
	//upfIpN4Port := n4Ctxt.SmfSEID.IPv4Addr.String() + ":" + strconv.Itoa(configure.UpfConf.N4.Smf.Port) //":8806" //127.0.0.1
	err = SendMsgI(reqN4, &resN4, peerAddr)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, n4Ctxt, "Failed to send N4 message:%s", err)
		return err
	}
	rlogger.Trace(moduleTag, rlogger.INFO, n4Ctxt, "Send N4 report request message to :%s", peerAddr)
	// 4 响应成功,SMF响应失败消息
	/*if resN4.IE.Cause.CauseValue != pfcp.Cause_Request_accepted {
		return fmt.Errorf("Failed to Report N4 context,cause=%d", resN4.IE.Cause.CauseValue)
	}*/
	/*//1.	SMF receives the trigger to remove the N4 session context for the PDU Session.
	// delete N4 context
	// 在upf成功响应后释放本地 N4 context
	err = smfcontext.DeleteN4Context(smfcontext.N4SessionIDKey(seid))
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil,  "Failure to delete N4 Context:%s", err)
		return err
	}*/

	// 返回成功
	return nil
}

// IEsSessionReportRequest
func CreateReportRequestMsg(n4Ctxt *n4context.N4SessionContext,
	dfCxt *pdrcontext.DataFlowContext) (*pfcp.SessionReportRequest, error) {
	rlogger.FuncEntry(moduleTag, n4Ctxt)
	var reqN4 pfcp.SessionReportRequest
	// 获取UPF的SEID
	reqN4.PfcpHeader.Set()
	reqN4.PfcpHeader.MessageType = pfcp.PFCP_Session_Report_Request
	reqN4.PfcpHeader.SEID = n4Ctxt.SmfSEID.SEID
	reqN4.PfcpHeader.SequenceNumber = SequenceNumber
	SequenceNumber += 1

	// 2 请求消息构造
	reqN4.IE = pfcp.IEsSessionReportRequest{}
	// todo
	//reqN4.IE.ReportType.Set(pfcp.DLDR)
	reqN4.IE.ReportType.Type = pfcp.IE_Report_Type
	reqN4.IE.ReportType.DLDR = true

	//set optional IE
	reqN4.IE.IeFlags.Set(pfcp.IE_Downlink_Data_Report)
	reqN4.IE.DownlinkDataReport = &pfcp.IEDownlinkDataReport{}
	reqN4.IE.DownlinkDataReport.Set()
	reqN4.IE.DownlinkDataReport.PDRID.Set(dfCxt.RuleID)

	reqN4.IE.DownlinkDataReport.IeFlags.Set(pfcp.IE_Downlink_Data_Service_Information)
	reqN4.IE.DownlinkDataReport.DLDataServiceInfo = &pfcp.IEDLDataServiceInfo{}
	reqN4.IE.DownlinkDataReport.DLDataServiceInfo.Set()

	reqN4.IE.DownlinkDataReport.DLDataServiceInfo.QFI = dfCxt.DP.QFI
	//reqN4.IE.DownlinkDataReport.DLDataServiceInfo.Flag &= pfcp.QFII_Flag
	reqN4.IE.DownlinkDataReport.DLDataServiceInfo.QFII = true

	if utils.IsSetByte(dfCxt.DP.PPP, 7) { // todo 定义7
		reqN4.IE.DownlinkDataReport.DLDataServiceInfo.PPIValue = dfCxt.DP.PPI
		//reqN4.IE.DownlinkDataReport.DLDataServiceInfo.Flag &= pfcp.QFII_Flag
		reqN4.IE.DownlinkDataReport.DLDataServiceInfo.PPI = true
	}

	return &reqN4, nil
}

func SendMsgI(req *pfcp.SessionReportRequest, res interface{}, peerAddr *net.UDPAddr) error {
	rlogger.FuncEntry(moduleTag, nil)
	// pfcp encode
	request := pfcpv1.Message{}
	request.HeaderSet(req.PfcpHeader)
	request.BodySet(req)
	data, err := request.Marshal()
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Pfcp msg marshal err %s", err)
		return err
	}
	UdpConn := pfcpudp.PfcpServer.UdpConn
	_, err = UdpConn.WriteToUDP(data, peerAddr)
	if err != nil {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil, "Pfcp msg send err %s", err)
		return err
	}
	//fmt.Printf("SessionModifyResponse: %+v\n", resN4)

	return nil
}

func SessionReportResponseHandle(res *pfcp.SessionReportResponse) error {
	// 4 响应成功,SMF响应失败消息
	if res.IE.Cause.CauseValue != pfcp.Cause_Request_accepted {
		rlogger.Trace(moduleTag, rlogger.ERROR, nil,
			"Failed to Session Report Response, cause: %d", res.IE.Cause.CauseValue)
		return fmt.Errorf("Failed to Report N4 context,cause=%d", res.IE.Cause.CauseValue)
	}

	return nil
}
