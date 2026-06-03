package pfcp

const ( // Presence class
	Mandatory           = 1
	Conditional         = 2
	MonditionalOptional = 3
	Optional            = 4
)

// 消息定义信息
type MsgDefInfo struct {
	Ies []IeDefInfo
}

// Grouped IE 定义信息
type GroupedIe struct {
	IeId int
	Ies  []IeDefInfo
}

// IE定义信息
type IeDefInfo struct {
	IeId     int
	Presence uint8
	Grouped  bool
}

// 消息定义信息
var (
	PfcpMsgConst = make([]*MsgDefInfo, IE_Spare)
)

func init() {

	//PFCP_Session_Establishment_Request
	MsgDef := MsgDefInfo{}
	PfcpMsgConst[PFCP_Session_Establishment_Request] = &MsgDef

	Ies := make([]IeDefInfo, 0)
	Ie := IeDefInfo{IE_Create_PDR, Mandatory, true}
	Ies = append(Ies, Ie)

	MsgDef.Ies = Ies

}
