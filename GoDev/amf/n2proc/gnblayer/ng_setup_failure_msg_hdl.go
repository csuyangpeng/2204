package gnblayer

import (
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func (p *GnbLayer) handleNgSetupFailureMsg(msgBuf *types.MsgBuf) error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	return nil
}

func (p *GnbLayer) sendNgSetupFailureMsg(ctype types3gpp.CauseType, cvalue types3gpp.CauseValue) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	ngSetupFailureMsg := ngapmsg.NgSetupFailureMsg{}
	ngSetupFailureMsg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())

	ngSetupFailureMsg.CauseType = uint16(ctype)
	ngSetupFailureMsg.CauseValue = uint16(cvalue)
	ngSetupFailureMsg.RelativeTimeToWait = uint32(types3gpp.Time5s)

	msgBuf := &types.MsgData{}
	msgBuf.MsgData = string(ngSetupFailureMsg.Encode())
	msgBuf.MsgLen = len(msgBuf.MsgData)

	p.sendch <- msgBuf
}
