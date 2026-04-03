package types

// MsgBuf define the messages transfer during gorouties
type MsgBuf struct {
	MsgLen int
	Buffer []byte
}

func (p *MsgBuf) DumpMsg()      {}
func (p *MsgBuf) IpcMsgDataIf() {}

// IpcMsg define the interface for all IPC Messages
type IpcMsg interface {
	DumpMsg()
}

// Chan key definition
type ChanKey string

var (
	ScNotifyChan ChanKey = "sc notify channel"
)

type StringValid struct {
	V     string
	Valid bool
}
type Uint32Valid struct {
	V     uint32
	Valid bool
}
type Uint64Valid struct {
	V     uint64
	Valid bool
}
type Int64Valid struct {
	V     int64
	Valid bool
}
