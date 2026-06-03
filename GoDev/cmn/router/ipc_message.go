package router

import "fmt"

// instance types
type InstType uint8

const (
	GnbGR InstType = iota + 1
	ScGR
	SbiGR
	SmfScGR
	CmdCliGR
	SmfCmdCliGR
	PfcpNodeGR
)

func (p InstType) String() string {
	switch p {
	case GnbGR:
		return "GNB"
	case ScGR:
		return "SC"
	case SbiGR:
		return "SBI"
	case SmfScGR:
		return "SmfSC"
	case CmdCliGR:
		return "CmdCli"
	case SmfCmdCliGR:
		return "SmfCmdCli"
	case PfcpNodeGR:
		return "PfcpNodeGR"
	default:
		return "INVLD"
	}
}

// route address for routing message between goroutines
type RouteAddr struct {
	Type InstType
	Id   uint32
}

func (p *RouteAddr) String() string {
	return fmt.Sprintf("%s.%d", p.Type, p.Id)
}

const (
	UnknownId   = 0xFFFFFFFD
	BroadcastId = 0xFFFFFFFE
	InvalidId   = 0xFFFFFFFF
)

// message types
type MsgType uint8

const (
	DP MsgType = iota
	CP
)

func (p MsgType) String() string {
	switch p {
	case DP:
		return "DataMsg"
	case CP:
		return "CtrlMsg"
	default:
		return "Invalid"
	}
}

// ipc message
type IpcMsg struct {
	MsgT MsgType
	//DestAddr RouteAddr
	//SrcAddr  RouteAddr
	MsgD MsgContent
}

//func (p *IpcMsg) String() string {
//	return fmt.Sprintf("IpcMessage{Type:%s, DestAddr:%s, SrcAddr:%s}",
//		p.MsgT, p.DestAddr, p.SrcAddr)
//}

// message data interface
type MsgContent interface {
	MsgContentIf()
}

// control message
type OpType uint8

const (
	Register OpType = iota
	Deregister
)

type ControlMsg struct {
	Op         OpType
	SrcAddr    RouteAddr
	PubChannel DataChannel
}

func (p ControlMsg) MsgContentIf() {}

// data message
type DataMsg struct {
	SrcAddr  RouteAddr
	DestAddr RouteAddr
	MsgData  IpcMsgData
}

func (p DataMsg) MsgContentIf() {}

// date message playload
type IpcMsgData interface {
	IpcMsgDataIf()
}
