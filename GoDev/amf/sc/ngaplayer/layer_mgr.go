package ngaplayer

import (
	"context"
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/redisclt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

type LayerMgr struct {
	scId    uint32
	ossCtxt codec.OssCtxt
	ctxt    context.Context
}

func (p LayerMgr) GetSCID() uint32 {
	return p.scId
}

func (p LayerMgr) String() (strbuf string) {
	strbuf += fmt.Sprintf("ngap layer manager Info:\n")
	strbuf += fmt.Sprintln("ossCtxt: ", p.ossCtxt)
	return strbuf
}

func NewLayerMgr(scId uint32) *LayerMgr {
	return &LayerMgr{scId: scId}
}

func (p *LayerMgr) Init() {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)
	p.ossCtxt = codec.NewOssCtxt()
}

func (p *LayerMgr) GetOssCtxt() codec.OssCtxt {
	return p.ossCtxt
}

func (p *LayerMgr) Cleanup() {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)
	codec.DeleteOssCtxt(p.ossCtxt)
}

func (p *LayerMgr) LoopProcN2Msg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	//save the base context
	p.ctxt = ctxt

	key := fmt.Sprintf("%s%d", types.AmfProc, p.scId)
	rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, nil,
		"start scan redis list(%s)", key)

	//scan the amf_sc_<id> msg queue from redis and dispatcher the messages
	for {
		data := &types3gpp.Gnb2AmfScMsg{}
		err := redisclt.Agent.BLPopObject(key, 0, data)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, nil, "no msg pop, error(%s)", err)
		} else {
			rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, nil, "ngap message comming.")
			p.HandleNgapMsg(data)
		}
	}
}

func (p *LayerMgr) HandleNgapMsg(data *types3gpp.Gnb2AmfScMsg) {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	if data == nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "receive invalid message from amf gnb module")
		return
	}

	//new a context for message handler
	type key = struct{}
	ctxtNew := context.WithValue(p.ctxt, key{}, "start handle ngap message")

	err := p.HandlerIncomingMsg(ctxtNew, data)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil,
			"failed to HandlerIncomingMsg, err(%s) ", err)
	}

	return
}
