package gctxt

import (
	"lite5gc/cmn/types3gpp"
)

type ConnStateType uint8

const (
	NgDisconnected        ConnStateType = 0
	NgDisConnectInProgess ConnStateType = 1
	NgConnected           ConnStateType = 2
)

type N2ConnCtxt struct {
	GnbConnID             uint32
	GnbInfo               types3gpp.GnbInfo
	NeedInitCtxtSetupPrcd bool
	AmfUeNgapID           AmfUeNgApId
	//unused
	ConnState ConnStateType
	Tai       types3gpp.TAI
}
