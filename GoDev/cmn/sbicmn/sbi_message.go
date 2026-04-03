package sbicmn

import (
	"lite5gc/cmn/nas"
	"lite5gc/openapi/models"
)

type MsgType int

const (
	//Nudm
	GetAmDataMsgRequest MsgType = iota
	GetAmDataMsgResponse

	GetSmfSelDataMsgRequest
	GetSmfSelDataMsgResponse

	GetAuthDataMsgRequest
	GetAuthDataMsgResponse

	PostAmf3gppAccessRegistration
	PostSdmSubscription

	GetSmDataMsgRequest
	GetSmDataMsgResponse

	GetUeCtxtInSmfDataMsgRequest
	GetUeCtxtInSmfDataMsgResponse

	// message between amf and smf
	PduSessCreateSMContextReq
	PduSessCreateSMContextResp

	PduSessUpdateSMContextReq
	PduSessUpdateSMContextResp

	PduSessReleaseSMContextReq
	PduSessReleaseSMContextResp

	N1N2MessageTransferReq
	N1N2MessageTransferResp

	// sbi消息处理失败场景
	NudmFailMsg
)

type SbiMessage struct {
	MsgType  MsgType
	MsgData  SbiMsgData
	ScInstId uint32
}

func (p *SbiMessage) IpcMsgDataIf() {}

type SbiMsgData interface {
	dumpMsg()
}

// SBI handler message
type SbiHandlerMessage struct {
	HTTPRequest  *Request
	ResponseChan chan *SbiHandlerResponseMessage
}

type SbiHandlerResponseMessage struct {
	HTTPResponse *Response
}

func NewSbiHandlerResponseMessage() *SbiHandlerResponseMessage {
	resp := &Response{}
	return &SbiHandlerResponseMessage{HTTPResponse: resp}
}
func (p SbiHandlerMessage) dumpMsg() {}

func NewSbiHandlerMsg(event MsgType, httpRequest *Request) (msg *SbiMessage) {
	msg = &SbiMessage{}
	msg.MsgType = event

	sbiHandlerMsg := &SbiHandlerMessage{}
	sbiHandlerMsg.ResponseChan = make(chan *SbiHandlerResponseMessage)
	sbiHandlerMsg.HTTPRequest = httpRequest

	msg.MsgData = sbiHandlerMsg

	msg.ScInstId = 0

	return msg
}

type ScSbiMsg interface {
	dumpMsg()
}

// release sm context
type SbiPostReleaseSmContext struct {
	SmContextRef string
	ReqData      *models.SmContextReleaseData
}

func (p *SbiPostReleaseSmContext) dumpMsg() {}

// modify sm context
type SbiPostModifySmContext struct {
	SmContextRef string
	ReqData      *models.UpdateSmContextRequest
	RespData     *models.UpdateSmContextResponse
}

func (p *SbiPostModifySmContext) dumpMsg() {}

// create sm context
type SbiPostCreateSmContext struct {
	Supi     string
	ReqData  *models.SmContextCreateData
	RespData *models.SmContextCreatedData
}

func (p *SbiPostCreateSmContext) dumpMsg() {}

// Amf3GppAccessRegistration data
type SbiPostAmf3gppAccessRegistration struct {
	Supi string
	Data *models.Amf3GppAccessRegistration
}

func (p *SbiPostAmf3gppAccessRegistration) dumpMsg() {}

// SdmSubscription data
type SbiPostSdmSubscription struct {
	Supi string
	Data *models.SdmSubscription
}

func (p *SbiPostSdmSubscription) dumpMsg() {}

// am data
type SbiGetAmDataMsg struct {
	Supi string
	Data *models.AccessAndMobilitySubscriptionData
}

func (p *SbiGetAmDataMsg) dumpMsg() {}

// authentication data
type SbiHandleFailMsg struct {
	Supi string
	Cause string
}

//AuthData
func (p *SbiHandleFailMsg) dumpMsg() {}

// smf select data
type SbiGetSmfSelDataMsg struct {
	Supi string
	Data *models.SmfSelectionSubscriptionData
}

func (p *SbiGetSmfSelDataMsg) dumpMsg() {}

// authentication data
type SbiGetAuthDataMsg struct {
	Supi string
	Data *models.AuthenticationInfoResult
}

//AuthData
func (p *SbiGetAuthDataMsg) dumpMsg() {}

type SbiGetSharedDataMsg struct {
	SharedDataIds []string
	SharedDatas   []models.SharedData
}

func (p SbiGetSharedDataMsg) dumpMsg() {}

// message: get sm data from udm
type SbiGetSmDataMsg struct {
	Supi string
	Psi  nas.PduSessID
	Data *[]models.SessionManagementSubscriptionData
}

func (p *SbiGetSmDataMsg) dumpMsg() {}

// message: get ue context in smf data from udm
type SbiGetUeCtxtInSmfDataMsg struct {
	Supi string
	Data *models.UeContextInSmfData
}

func (p *SbiGetUeCtxtInSmfDataMsg) dumpMsg() {}

// message: post n1n2 msg transfer req message
type SbiPostN1N2MsgTransferMsg struct {
	Supi     string
	Psi      nas.PduSessID
	ReqData  *models.N1N2MessageTransferRequest
	RespData *models.N1N2MessageTransferRspData
}

func (p *SbiPostN1N2MsgTransferMsg) dumpMsg() {}
