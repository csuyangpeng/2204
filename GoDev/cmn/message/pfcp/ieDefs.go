package pfcp

import (
	"bytes"
	"encoding/binary"
)

// 3GPP TS 29.244 V15.5.0 (2019-03)
// N4 消息
// 8.1.2	Information Element Types
// Message IE
// Figure 8.1.1-2: 3GPP defined Information Element Format,
// the IE type shall be within the range of 0 to 32767
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = xxx (decimal)
3 to 4	Length = n
5 to (n+4)	IE specific data or content of a grouped IE*/

const (
	IE_Reserved                             = 0
	IE_Create_PDR                           = 1
	IE_PDI                                  = 2
	IE_Create_FAR                           = 3
	IE_Forwarding_Parameters                = 4
	IE_Duplicating_Parameters               = 5
	IE_Create_URR                           = 6
	IE_Create_QER                           = 7
	IE_Created_PDR                          = 8
	IE_Update_PDR                           = 9
	IE_Update_FAR                           = 10
	IE_Update_Forwarding_Parameters         = 11
	IE_Update_BAR                           = 12 // (PFCP Session Report Response)
	IE_Update_URR                           = 13
	IE_Update_QER                           = 14
	IE_Remove_PDR                           = 15
	IE_Remove_FAR                           = 16
	IE_Remove_URR                           = 17
	IE_Remove_QER                           = 18
	IE_Cause                                = 19
	IE_Source_Interface                     = 20
	IE_F_TEID                               = 21
	IE_Network_Instance                     = 22
	IE_SDF_Filter                           = 23
	IE_Application_ID                       = 24
	IE_Gate_Status                          = 25
	IE_MBR                                  = 26
	IE_GBR                                  = 27
	IE_QER_Correlation_ID                   = 28
	IE_Precedence                           = 29
	IE_Transport_Level_Marking              = 30
	IE_Volume_Threshold                     = 31
	IE_Time_Threshold                       = 32
	IE_Monitoring_Time                      = 33
	IE_Subsequent_Volume_Threshold          = 34
	IE_Subsequent_Time_Threshold            = 35
	IE_Inactivity_Detection_Time            = 36
	IE_Reporting_Triggers                   = 37
	IE_Redirect_Information                 = 38
	IE_Report_Type                          = 39
	IE_Offending_IE                         = 40
	IE_Forwarding_Policy                    = 41
	IE_Destination_Interface                = 42
	IE_UP_Function_Features                 = 43
	IE_Apply_Action                         = 44
	IE_Downlink_Data_Service_Information    = 45
	IE_Downlink_Data_Notification_Delay     = 46
	IE_DL_Buffering_Duration                = 47
	IE_DL_Buffering_Suggested_Packet_Count  = 48
	IE_PFCPSMReq_Flags                      = 49
	IE_PFCPSRRsp_Flags                      = 50
	IE_Load_Control_Information             = 51
	IE_Sequence_Number                      = 52
	IE_Metric                               = 53
	IE_Overload_Control_Information         = 54
	IE_Timer                                = 55
	IE_Packet_Detection_Rule_ID             = 56
	IE_F_SEID                               = 57
	IE_Application_IDs_PFDs                 = 58
	IE_PFD_context                          = 59
	IE_Node_ID                              = 60
	IE_PFD_contents                         = 61
	IE_Measurement_Method                   = 62
	IE_Usage_Report_Trigger                 = 63
	IE_Measurement_Period                   = 64
	IE_FQ_CSID                              = 65 //There is no IE in N4
	IE_Volume_Measurement                   = 66
	IE_Duration_Measurement                 = 67
	IE_Application_Detection_Information    = 68
	IE_Time_of_First_Packet                 = 69
	IE_Time_of_Last_Packet                  = 70
	IE_Quota_Holding_Time                   = 71
	IE_Dropped_DL_Traffic_Threshold         = 72
	IE_Volume_Quota                         = 73
	IE_Time_Quota                           = 74
	IE_Start_Time                           = 75
	IE_End_Time                             = 76
	IE_Query_URR                            = 77
	IE_Usage_Report_Modification_Response   = 78 //(in Session Modification Response)
	IE_Usage_Report_Deletion_Response       = 79 //(Session Deletion Response)
	IE_Usage_Report_Request                 = 80 //(Session Report Request)
	IE_URR_ID                               = 81
	IE_LinkedURR_ID                         = 82
	IE_Downlink_Data_Report                 = 83
	IE_Outer_Header_Creation                = 84
	IE_Create_BAR                           = 85
	IE_Update_BAR_Modification_Request      = 86 //(Session Modification Request)
	IE_Remove_BAR                           = 87
	IE_BAR_ID                               = 88
	IE_CP_Function_Features                 = 89
	IE_Usage_Information                    = 90
	IE_Application_Instance_ID              = 91
	IE_Flow_Information                     = 92
	IE_UE_IP_Address                        = 93
	IE_Packet_Rate                          = 94 //There is no IE in N4
	IE_Outer_Header_Removal                 = 95
	IE_Recovery_Time_Stamp                  = 96
	IE_DL_Flow_Level_Marking                = 97 //There is no IE in N4
	IE_Header_Enrichment                    = 98
	IE_Error_Indication_Report              = 99
	IE_Measurement_Information              = 100
	IE_Node_Report_Type                     = 101
	IE_User_Plane_Path_Failure_Report       = 102
	IE_Remote_GTPU_Peer                     = 103
	IE_UR_SEQN                              = 104
	IE_Update_Duplicating_Parameters        = 105
	IE_Activate_Predefined_Rules            = 106
	IE_Deactivate_Predefined_Rules          = 107
	IE_FAR_ID                               = 108
	IE_QER_ID                               = 109
	IE_OCI_Flags                            = 110
	IE_IE_PFCP_Association_Release_Request  = 111
	IE_Graceful_Release_Period              = 112
	IE_PDN_Type                             = 113
	IE_Failed_Rule_ID                       = 114
	IE_Time_Quota_Mechanism                 = 115 //There is no IE in N4
	IE_User_Plane_IP_Resource_Information   = 116
	IE_User_Plane_Inactivity_Timer          = 117
	IE_Aggregated_URRs                      = 118 //There is no IE in N4
	IE_Multiplier                           = 119 //There is no IE in N4
	IE_Aggregated_URR_ID                    = 120 //There is no IE in N4
	IE_Subsequent_Volume_Quota              = 121
	IE_Subsequent_Time_Quota                = 122
	IE_RQI                                  = 123
	IE_QFI                                  = 124
	IE_Query_URR_Reference                  = 125
	IE_Additional_Usage_Reports_Information = 126
	IE_Create_Traffic_Endpoint              = 127
	IE_Created_Traffic_Endpoint             = 128
	IE_Update_Traffic_Endpoint              = 129
	IE_Remove_Traffic_Endpoint              = 130
	IE_Traffic_Endpoint_ID                  = 131
	IE_Ethernet_Packet_Filter               = 132
	IE_MAC_address                          = 133
	IE_C_TAG                                = 134
	IE_S_TAG                                = 135
	IE_Ethertype                            = 136
	IE_Proxying                             = 137
	IE_Ethernet_Filter_ID                   = 138
	IE_Ethernet_Filter_Properties           = 139
	IE_Suggested_Buffering_Packets_Count    = 140
	IE_User_ID                              = 141
	IE_Ethernet_PDU_Session_Information     = 142
	IE_Ethernet_Traffic_Information         = 143
	IE_MAC_Addresses_Detected               = 144
	IE_MAC_Addresses_Removed                = 145
	IE_Ethernet_Inactivity_Timer            = 146
	IE_Additional_Monitoring_Time           = 147
	// 2019.3 update
	IE_Event_Quota                = 148
	IE_Event_Threshold            = 149
	IE_Subsequent_Event_Quota     = 150
	IE_Subsequent_Event_Threshold = 151
	IE_Trace_Information          = 152
	IE_Framed_Route               = 153
	IE_Framed_Routing             = 154
	IE_Framed_IPv6_Route          = 155
	IE_Event_Time_Stamp           = 156
	IE_Averaging_Window           = 157
	IE_Paging_Policy_Indicator    = 158
	IE_Spare                      = 159
)
const (
	Byte_MAX_Value   = 255
	Byte_2_MAX_Value = 65535
)

type IETypeLength struct {
	Type   uint16 `json:",omitempty"`
	Length uint16 `json:",omitempty"`
}

func (tl IETypeLength) EncodeTl() ([]byte, error) {
	// encode Tl
	encBuf := bytes.NewBuffer(nil)
	err := binary.Write(encBuf, binary.BigEndian, tl.Type)
	if err != nil {
		return nil, err
	}
	err = binary.Write(encBuf, binary.BigEndian, tl.Length)
	if err != nil {
		return nil, err
	}

	return encBuf.Bytes(), nil
}

func (tl IETypeLength) EncodeTlV(v []byte) ([]byte, error) {
	// encode Tl
	encBuf := bytes.NewBuffer(nil)
	err := binary.Write(encBuf, binary.BigEndian, tl.Type)
	if err != nil {
		return nil, err
	}
	err = binary.Write(encBuf, binary.BigEndian, tl.Length)
	if err != nil {
		return nil, err
	}
	_, err = encBuf.Write(v)
	if err != nil {
		return nil, err
	}
	return encBuf.Bytes(), nil
}

type IESet interface {
	Set(v uint8) error
}

type IEGet interface {
	Get() (v uint8, e error)
}

type IESetGet interface {
	IESet
	IEGet
}

// Octets 中的位对应的值
const (
	IEFlag_Bit_1 = 1
	IEFlag_Bit_2 = 2
	IEFlag_Bit_3 = 4
	IEFlag_Bit_4 = 8
	IEFlag_Bit_5 = 16
	IEFlag_Bit_6 = 32
	IEFlag_Bit_7 = 64
	IEFlag_Bit_8 = 128
)

// Figure 8.1.1-3: Vendor-Specific Information Element Format
// the IE type value shall be within the range of 32768 to 65535.
/*Bits
Octets	8	7	6	5	4	3	2	1
1 to 2	Type = xxx (decimal)
3 to 4	Length = n
5 to 6	Enterprise ID
7 to (n+4)	IE specific data or content of a grouped IE
*/
