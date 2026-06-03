package create

import (
	"errors"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/smf/smfcontext/gctxt"
	"strconv"
)

func SessionReportRequest(req pfcp.SessionReportRequest,
	res *pfcp.SessionReportResponse) error {

	rlogger.FuncEntry(types.ModuleSmfN4, &req.PfcpHeader)

	if req.PfcpHeader.MessageType != pfcp.PFCP_Session_Report_Request {
		return errors.New("Session Report Request message type error: " +
			strconv.Itoa(int(req.PfcpHeader.MessageType)))
	}

	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &req.PfcpHeader, "Session Report Request:")
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &req.PfcpHeader, "  .MessageType    %v", req.PfcpHeader.MessageType)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &req.PfcpHeader, "  .SEID           %v", req.PfcpHeader.SEID)
	rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &req.PfcpHeader, "  .SequenceNumber %v", req.PfcpHeader.SequenceNumber)

	res.IE.Cause.Set()
	// N4 context
	n4Cxt, err := gctxt.GetN4Context(gctxt.N4SessionIDKey(req.PfcpHeader.SEID))
	if err != nil {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, &req.PfcpHeader, "Failed to get N4 Context:%s", err)
		res.IE.Cause.CauseValue = pfcp.Cause_Session_context_not_found
		res.PfcpHeader.Version = 1
		res.PfcpHeader.MessageType = pfcp.PFCP_Session_Report_Response
		res.PfcpHeader.SEID = 0
		res.PfcpHeader.SequenceNumber = req.PfcpHeader.SequenceNumber
		return nil
	}

	// N4 Session Report Response
	res.PfcpHeader.Version = 1
	res.PfcpHeader.SFlag = pfcp.Flag
	res.PfcpHeader.MessageType = pfcp.PFCP_Session_Report_Response
	res.PfcpHeader.SEID = n4Cxt.UpfSEID.SEID
	res.PfcpHeader.SequenceNumber = req.PfcpHeader.SequenceNumber
	/*if !req.IE.IeFlags.Test(pfcp.IE_Report_Type) {
		rlogger.Trace(types.ModuleSmfN4, rlogger.ERROR, &req.PfcpHeader, "Report type error:%s", req.IE.ReportType.TypeValue)
		res.IE.Cause.CauseValue = pfcp.Cause_Mandatory_IE_incorrect
		return nil
	}*/

	if req.IE.ReportType.DLDR {
		// PDR ID m
		if req.IE.DownlinkDataReport == nil ||
			req.IE.DownlinkDataReport.PDRID.Type != pfcp.IE_Packet_Detection_Rule_ID {
			res.IE.Cause.CauseValue = pfcp.Cause_Mandatory_IE_missing
			res.IE.OffendingIE.Set()
			res.IE.OffendingIE.TypeOffendingIE = pfcp.IE_Packet_Detection_Rule_ID
			return nil
		}
		pdrId := req.IE.DownlinkDataReport.PDRID
		if req.IE.DownlinkDataReport.DLDataServiceInfo != nil {
			serviceInfo := req.IE.DownlinkDataReport.DLDataServiceInfo

			rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &req.PfcpHeader, "Down link Data Report:", pdrId, serviceInfo)
		}
		rlogger.Trace(types.ModuleSmfN4, rlogger.INFO, &req.PfcpHeader, "Down link Data Report:", pdrId)
		res.IE.Cause.CauseValue = pfcp.Cause_Request_accepted

		// Update BAR
		//Table 7.5.9.2-1: Update BAR IE in PFCP Session Report Response
		//Octet 1 and 2		Update BAR IE Type = 12 (decimal)
		//Octets 3 and 4		Length = n
		//Information elements	P                  IE Type
		//
		//BAR ID	M                                  BAR ID
		//Downlink Data Notification Delay	C      Downlink Data Notification Delay
		//DL Buffering Duration	C                  DL Buffering Duration
		//DL Buffering Suggested Packet Count	O      DL Buffering Suggested Packet Count
		//Suggested Buffering Packets Count	C      Suggested Buffering Packets Count

		res.IE.UpdateBAR.BARID.Set(n4Cxt.BAR.BARID.Value)
		//res.IE.UpdateBAR.DLDataNotificationDelay.Set(configure.SmfConf.N4Conf.BAR.DLDataNotificationDelay.Value)
		//res.IE.UpdateBAR.DLBufferingDuration.Set()
		//res.IE.UpdateBAR.DLBufferingSuggestedPacketCount.Set(uint16(configure.SmfConf.N4Conf.BAR.SugBuffPacketsCount.CountValue))

		// PFCPSRRsp-Flags
		res.IE.PFCPSRRspFlags.Set(false)

		//n4Cxt.SEID;n4Cxt.PduSessionId;serviceInfo.QFI;serviceInfo.PPIValue;n4Cxt.PDRs[pdrId].PDI.LocalFTEID
		return nil

	}
	res.IE.Cause.CauseValue = pfcp.Cause_Request_rejected
	return nil
}
