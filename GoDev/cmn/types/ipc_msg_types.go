package types

type MsgData struct {
	MsgLen  int
	MsgData string
}

type IpcMsgData struct {
	Sender   string
	Receiver string
	Data     []byte
}
