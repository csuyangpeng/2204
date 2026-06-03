package pfcptst

import (
	"fmt"
	"net"
	"time"

	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
)

// 用于测试
func Dispatch(msg pfcpv1.Message, res *pfcpv1.Message) error {
	switch msg.Header.MessageType {
	case pfcp.PFCP_Heartbeat_Request:
		request, ok := msg.Body.(*pfcp.HeartbeatRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		request.PfcpHeader.Version = msg.Header.Version
		request.PfcpHeader.MessageType = msg.Header.MessageType
		request.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber

		response := &pfcp.HeartbeatResponse{}
		if ok {
			err := pfcpv1.HandlePfcpHeartbeatRequest(*request, response)
			if err != nil {
				return err
			}
			pfcpHeader := pfcp.PfcpHeader{}
			pfcpHeader.Version = response.PfcpHeader.Version
			pfcpHeader.MessageType = response.PfcpHeader.MessageType
			pfcpHeader.Length = response.PfcpHeader.Length
			pfcpHeader.SequenceNumber = response.PfcpHeader.SequenceNumber

			res.Header = pfcpHeader
			res.Body = response
			return nil
		}
	case pfcp.PFCP_Heartbeat_Response:
		msg, ok := msg.Body.(*pfcp.HeartbeatResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		err := pfcpv1.HandlePfcpHeartbeatResponse(msg)
		if err != nil {
			return err
		}

		return nil

	case pfcp.PFCP_PFD_Management_Request:
		request, ok := msg.Body.(*pfcp.PFCPPFDManagementRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		request.PfcpHeader.Version = msg.Header.Version
		request.PfcpHeader.MessageType = msg.Header.MessageType
		request.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber
		err := pfcpv1.HandlePfcpPFDManagementRequest(request)
		if err != nil {
			return err
		}

		response := &pfcp.PFCPPFDManagementResponse{}
		response.IE = &pfcp.IEPFCPPFDManagementResponse{
			Cause: &pfcp.IECause{
				CauseValue: pfcp.Cause_Request_accepted,
			},
		}

		pfcpHeader := pfcp.PfcpHeader{}
		pfcpHeader.Version = pfcp.Version
		pfcpHeader.MessageType = pfcp.PFCP_PFD_Management_Response
		pfcpHeader.Length = 0
		pfcpHeader.SequenceNumber = request.PfcpHeader.SequenceNumber

		res.Header = pfcpHeader
		res.Body = response

		return nil
	case pfcp.PFCP_PFD_Management_Response:
		rmsg, ok := msg.Body.(*pfcp.PFCPPFDManagementResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		pfcpv1.HandlePfcpPFDManagementResponse(rmsg)

		return nil
	case pfcp.PFCP_Association_Setup_Request:
		rmsg, ok := msg.Body.(*pfcp.PFCPAssociationSetupRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		err := pfcpv1.HandleAssociationSetupRequest(rmsg)
		if err != nil {
			return err
		}
		response := &pfcp.PFCPAssociationSetupResponse{
			PfcpHeader: pfcp.PfcpHeaderforNode{
				Version:        pfcp.Version,
				MessageType:    pfcp.PFCP_Association_Setup_Response,
				Length:         0, // todo 编码后填充
				SequenceNumber: msg.Header.SequenceNumber},
		}
		response.IE = &pfcp.IEPFCPAssociationSetupResponse{}
		response.IE.NodeID = &pfcp.IENodeID{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Node_ID,
				Length: 5,
			},
			NodeIDType: 0,
			//NodeIDvalue: net.ParseIP(n4layer.UpfN4Layer.UpfIp).To4(), //[]byte{10, 202, 94, 1},
		}
		response.IE.Cause = &pfcp.IECause{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Cause,
				Length: 1,
			},
			CauseValue: 1,
		}

		response.IE.RecoveryTimeStamp = &pfcp.IERecoveryTimeStamp{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Recovery_Time_Stamp,
				Length: 4,
			},
			RecoveryTimeStamp: time.Unix(time.Now().Unix(), 0), //time.Unix(1556588833, 0),
		}
		res.Header = msg.Header
		res.Header.MessageType = pfcp.PFCP_Association_Setup_Response
		//todo 编码后填充
		res.Header.Length = 0
		res.Body = response
	case pfcp.PFCP_Association_Setup_Response:
		rmsg, ok := msg.Body.(*pfcp.PFCPAssociationSetupResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		pfcpv1.HandleAssociationSetupResponse(rmsg)

		return nil
	case pfcp.PFCP_Association_Update_Request:
		rmsg, ok := msg.Body.(*pfcp.PFCPAssociationUpdateRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		err := pfcpv1.HandleAssociationUpdateRequest(rmsg)
		if err != nil {
			return err
		}
		response := &pfcp.PFCPAssociationUpdateResponse{
			PfcpHeader: pfcp.PfcpHeaderforNode{
				Version:        pfcp.Version,
				MessageType:    pfcp.PFCP_Association_Update_Response,
				Length:         0, // 编码后填充
				SequenceNumber: msg.Header.SequenceNumber},
		}
		response.IE = &pfcp.IEPFCPAssociationUpdateResponse{}
		response.IE.NodeID = &pfcp.IENodeID{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Node_ID,
				Length: 5,
			},
			NodeIDType:  0,
			NodeIDvalue: net.ParseIP("10.202.97.2").To4(),
		}
		response.IE.Cause = &pfcp.IECause{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Cause,
				Length: 1,
			},
			CauseValue: 1,
		}

		res.Header = msg.Header
		res.Header.MessageType = response.PfcpHeader.MessageType
		// 编码后填充
		res.Header.Length = 0
		res.Body = response

		return nil
	case pfcp.PFCP_Association_Update_Response:
		rmsg, ok := msg.Body.(*pfcp.PFCPAssociationUpdateResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		pfcpv1.HandleAssociationUpdateResponse(rmsg)

		return nil
	case pfcp.PFCP_Association_Release_Request:
		rmsg, ok := msg.Body.(*pfcp.PFCPAssociationReleaseRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		err := pfcpv1.HandleAssociationReleaseRequest(rmsg)
		if err != nil {
			return err
		}
		response := &pfcp.PFCPAssociationReleaseResponse{
			PfcpHeader: pfcp.PfcpHeaderforNode{
				Version:        pfcp.Version,
				MessageType:    pfcp.PFCP_Association_Release_Response,
				Length:         0, // todo 编码后填充
				SequenceNumber: msg.Header.SequenceNumber},
		}
		response.IE = &pfcp.IEPFCPAssociationReleaseResponse{}
		response.IE.NodeID = &pfcp.IENodeID{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Node_ID,
				Length: 5,
			},
			NodeIDType:  0,
			NodeIDvalue: net.ParseIP("10.202.94.2").To4(),
		}
		response.IE.Cause = &pfcp.IECause{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Cause,
				Length: 1,
			},
			CauseValue: 1,
		}

		res.Header = msg.Header
		res.Header.MessageType = response.PfcpHeader.MessageType
		// 编码后填充
		res.Header.Length = 0
		res.Body = response
		return nil
	case pfcp.PFCP_Association_Release_Response:
		msg, ok := msg.Body.(*pfcp.PFCPAssociationReleaseResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		err := pfcpv1.HandleAssociationReleaseResponse(msg)
		if err != nil {
			return err
		}
		return nil
	case pfcp.PFCP_Node_Report_Request:
		request, ok := msg.Body.(*pfcp.PFCPNodeReportRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		err := pfcpv1.HandleNodeReportRequest(request)
		if err != nil {
			return err
		}

		response := &pfcp.PFCPNodeReportResponse{
			PfcpHeader: pfcp.PfcpHeaderforNode{
				Version:        pfcp.Version,
				MessageType:    pfcp.PFCP_Node_Report_Response,
				Length:         0, // 编码后填充
				SequenceNumber: msg.Header.SequenceNumber},
		}
		response.IE = &pfcp.IEPFCPNodeReportResponse{}
		response.IE.NodeID = &pfcp.IENodeID{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Node_ID,
				Length: 5,
			},
			NodeIDType:  0,
			NodeIDvalue: net.ParseIP("10.202.97.2").To4(),
		}
		response.IE.Cause = &pfcp.IECause{
			IETypeLength: pfcp.IETypeLength{
				Type:   pfcp.IE_Cause,
				Length: 1,
			},
			CauseValue: 1,
		}

		res.Header = msg.Header
		res.Header.MessageType = response.PfcpHeader.MessageType
		// 编码后填充
		res.Header.Length = 0
		res.Body = response

		return nil
	case pfcp.PFCP_Node_Report_Response:
		msg, ok := msg.Body.(*pfcp.PFCPNodeReportResponse)
		if !ok {
			return fmt.Errorf("type error")
		}
		err := pfcpv1.HandleNodeReportResponse(msg)
		if err != nil {
			return err
		}
		return nil
		//case pfcp.PFCP_SESSION_SET_DELETION_REQUEST:
		//	pfcp_handler.HandlePfcpSessionSetDeletionRequest(msg)
		//case pfcp.PFCP_SESSION_SET_DELETION_RESPONSE:
		//	pfcp_handler.HandlePfcpSessionSetDeletionResponse(msg)
	case pfcp.PFCP_Session_Establishment_Request:
		/*var n4 n4layer.N4Msg
		request, ok := msg.Body.(*pfcp.SessionEstablishmentRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		//解码消息头填充处理消息头
		request.PfcpHeader.Version = msg.Header.Version
		request.PfcpHeader.MPFlag = msg.Header.MPFlag
		request.PfcpHeader.SFlag = msg.Header.SFlag

		request.PfcpHeader.MessageType = msg.Header.MessageType
		request.PfcpHeader.Length = msg.Header.Length
		request.PfcpHeader.SEID = msg.Header.SEID
		request.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber
		request.PfcpHeader.MessagePriority = msg.Header.MessagePriority

		response := &pfcp.SessionEstablishmentResponse{}
		if ok {
			err := n4.SessionEstablishmentRequest(*request, response)
			if err != nil {
				return err
			}
			pfcpHeader := pfcp.PfcpHeader{}
			pfcpHeader.Version = response.PfcpHeader.Version
			pfcpHeader.MPFlag = response.PfcpHeader.MPFlag
			pfcpHeader.SFlag = response.PfcpHeader.SFlag

			pfcpHeader.MessageType = response.PfcpHeader.MessageType
			pfcpHeader.Length = response.PfcpHeader.Length
			pfcpHeader.SEID = response.PfcpHeader.SEID
			pfcpHeader.SequenceNumber = response.PfcpHeader.SequenceNumber
			pfcpHeader.MessagePriority = response.PfcpHeader.MessagePriority

			res.Header = pfcpHeader
			res.Body = response
		}

		return nil*/
	case pfcp.PFCP_Session_Establishment_Response:
		_, ok := msg.Body.(*pfcp.SessionEstablishmentResponse)
		if !ok {
			return fmt.Errorf("type error")
		}

		return nil
	case pfcp.PFCP_Session_Modification_Request:
		//var n4 n4layer.N4Msg
		request, ok := msg.Body.(*pfcp.SessionModifyRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		//解码消息头填充处理消息头
		request.PfcpHeader.Version = msg.Header.Version
		request.PfcpHeader.MPFlag = msg.Header.MPFlag
		request.PfcpHeader.SFlag = msg.Header.SFlag

		request.PfcpHeader.MessageType = msg.Header.MessageType
		request.PfcpHeader.Length = msg.Header.Length
		request.PfcpHeader.SEID = msg.Header.SEID
		request.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber
		request.PfcpHeader.MessagePriority = msg.Header.MessagePriority

		response := &pfcp.SessionModifyResponse{}

		//err := n4.SessionModifyRequest(*request, response)
		//if err != nil {
		//	return err
		//}
		pfcpHeader := pfcp.PfcpHeader{}
		pfcpHeader.Version = response.PfcpHeader.Version
		pfcpHeader.MPFlag = response.PfcpHeader.MPFlag
		pfcpHeader.SFlag = response.PfcpHeader.SFlag

		pfcpHeader.MessageType = response.PfcpHeader.MessageType
		pfcpHeader.Length = response.PfcpHeader.Length
		pfcpHeader.SEID = response.PfcpHeader.SEID
		pfcpHeader.SequenceNumber = response.PfcpHeader.SequenceNumber
		pfcpHeader.MessagePriority = response.PfcpHeader.MessagePriority

		res.Header = pfcpHeader
		res.Body = response

		return nil
	case pfcp.PFCP_Session_Modification_Response:
		_, ok := msg.Body.(*pfcp.SessionModifyResponse)
		if !ok {
			return fmt.Errorf("type error")
		}

		return nil
	case pfcp.PFCP_Session_Deletion_Request:
		//var n4 n4layer.N4Msg
		request, ok := msg.Body.(*pfcp.SessionReleaseRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		//解码消息头填充处理消息头
		request.PfcpHeader.Version = msg.Header.Version
		request.PfcpHeader.MPFlag = msg.Header.MPFlag
		request.PfcpHeader.SFlag = msg.Header.SFlag

		request.PfcpHeader.MessageType = msg.Header.MessageType
		request.PfcpHeader.Length = msg.Header.Length
		request.PfcpHeader.SEID = msg.Header.SEID
		request.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber
		request.PfcpHeader.MessagePriority = msg.Header.MessagePriority

		response := &pfcp.SessionReleaseResponse{}

		//err := n4.SessionReleaseRequest(*request, response)
		//if err != nil {
		//	return err
		//}
		pfcpHeader := pfcp.PfcpHeader{}
		pfcpHeader.Version = response.PfcpHeader.Version
		pfcpHeader.MPFlag = response.PfcpHeader.MPFlag
		pfcpHeader.SFlag = response.PfcpHeader.SFlag

		pfcpHeader.MessageType = response.PfcpHeader.MessageType
		pfcpHeader.Length = response.PfcpHeader.Length
		pfcpHeader.SEID = response.PfcpHeader.SEID
		pfcpHeader.SequenceNumber = response.PfcpHeader.SequenceNumber
		pfcpHeader.MessagePriority = response.PfcpHeader.MessagePriority

		res.Header = pfcpHeader
		res.Body = response

		return nil
	case pfcp.PFCP_Session_Deletion_Response:
		_, ok := msg.Body.(*pfcp.SessionReleaseResponse)
		if !ok {
			return fmt.Errorf("type error")
		}

		return nil
	case pfcp.PFCP_Session_Report_Request:
		request, ok := msg.Body.(*pfcp.SessionReportRequest)
		if !ok {
			return fmt.Errorf("type error")
		}
		//解码消息头填充处理消息头
		request.PfcpHeader.Version = msg.Header.Version
		request.PfcpHeader.MPFlag = msg.Header.MPFlag
		request.PfcpHeader.SFlag = msg.Header.SFlag

		request.PfcpHeader.MessageType = msg.Header.MessageType
		request.PfcpHeader.Length = msg.Header.Length
		request.PfcpHeader.SEID = msg.Header.SEID
		request.PfcpHeader.SequenceNumber = msg.Header.SequenceNumber
		request.PfcpHeader.MessagePriority = msg.Header.MessagePriority
		// todo msg handle

		// send response
		response := &pfcp.SessionReportResponse{}
		response.IE = pfcp.IEsSessionReportResponse{
			Cause: pfcp.IECause{
				CauseValue: pfcp.Cause_Request_accepted,
			},
		}

		pfcpHeader := pfcp.PfcpHeader{}
		pfcpHeader.Version = pfcp.Version
		pfcpHeader.SFlag = pfcp.Flag
		pfcpHeader.MessageType = pfcp.PFCP_Session_Report_Response
		pfcpHeader.Length = 0
		pfcpHeader.SEID = 0x0000000000004e20
		pfcpHeader.SequenceNumber = request.PfcpHeader.SequenceNumber

		res.Header = pfcpHeader
		res.Body = response

		return nil
	case pfcp.PFCP_Session_Report_Response:
		return nil

	default:
		return fmt.Errorf("Unknown PFCP message type: %d", msg.Header.MessageType)

	}
	return nil

}
