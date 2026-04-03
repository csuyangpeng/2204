package n11msg

type N11Msg struct {
	AmfSmCtxtId uint32
	SmfSmCtxtId uint32
	MsgType     N11MsgType
	MsgData     N11MsgData
}

func (p N11Msg) IpcMsgDataIf() {}

type N11MsgType uint8

const (
	CreateSmCtxtReq N11MsgType = iota
	CreateSmCtxtResp
	N1N2MsgReq
	N1N2MsgResp
	SmCtxtStatusNotify
	UpdateSmCtxtReq
	UpdateSmCtxtResp
	ReleaseSmCtxtReq
	ReleaseSmCtxtResp
)

type N11MsgData interface {
	N11MsgDataIf()
}
