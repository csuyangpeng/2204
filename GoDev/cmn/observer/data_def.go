package observer

type MsgType string

type MsgParam interface{}
type MsgInfo struct {
	MsgType
	MsgParam
}

type MsgHandler func(param MsgParam) bool
type MsgHandlers []MsgHandler
type HandlerInfo map[MsgType]MsgHandlers
