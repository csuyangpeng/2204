package gnblayer

import (
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func (p *GnbLayer) handleNgResetMsg(msgBuf *types.MsgBuf) error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	ngResetMsg := ngapmsg.NewNgResetMsg()
	ngResetMsg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	err := ngResetMsg.Decode(msgBuf.Buffer)
	if err != nil {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to decode Ng Setup Request Message")
		return err
	}
	switch ngResetMsg.ResetTypeChoice.NgInterface {
	case types3gpp.ResetAll:
		p.SendGnbSctpShutdownMessages()
	}
	return nil
}
