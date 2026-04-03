package gnblayer

import (
	"context"
	"fmt"
	"lite5gc/amf/tatable"
	"lite5gc/cmn/redisclt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"math/rand"
	"time"

	codec "lite5gc/cmn/message/ngap/ngapcodec"
)

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

// GnbLayer define for manage gnb layer ngap message
type GnbLayer struct {
	ossCtxt codec.OssCtxt

	// send sctp msg channel
	sendch chan<- *types.MsgData

	gnbInfo      types3gpp.GnbInfo
	ranNodeName  string
	defPagingDrx types3gpp.PagingDRX
	supTaList    []types3gpp.SupportedTA
}

func (p *GnbLayer) SetGnbIP(i string) {
	p.gnbInfo.GnbIP = i
}

func (p GnbLayer) GetGnbIP() string {
	return p.gnbInfo.GnbIP
}

// Init will initialize the NgAp layer manager
func (p *GnbLayer) Init(instID uint32,
	ip string,
	sendCh chan<- *types.MsgData,
//recvCh <-chan *types.MsgBuf,
	ctxt context.Context) error {

	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	p.gnbInfo.GnbInstId = instID
	p.gnbInfo.GnbIP = ip

	p.sendch = sendCh

	p.ossCtxt = codec.NewOssCtxt()

	// save the gnb id in redis
	// publish gnb id
	//_, err := n2sctp.RedisClt.Publish("amf_gnb", fmt.Sprintf("gnb_%d", instID))
	//if err != nil {
	//	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, "failed to publish gnb id(%d)", instID)
	//}

	// remove the redis list for amf_gnb_<id>
	redisclt.Agent.Del(fmt.Sprintf("%s%s", types.GnbProc, fmt.Sprintf("%d", instID)))

	return nil
}

// CleanUp layer manager cleanup the resources
func (p *GnbLayer) CleanUp() {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	//notify all the sc, this ngap layer is destroyed
	p.notifyDestroy()

	//delete oss ctxt
	codec.DeleteOssCtxt(p.ossCtxt)

	// TODO remove in redis for ta ran table info
	tatable.RemoveGnbInfo(p.GetGnbIP())
}

func (p *GnbLayer) LoopProcScMsg(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	key := fmt.Sprintf("%s%d", types.GnbProc, p.gnbInfo.GnbInstId)
	//scan the gnb_<id> msg queue from redis and dispatcher the messages
	for {
		data := &types.IpcMsgData{}

		err := redisclt.Agent.BLPopObject(key, 0, data)
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
				"receive msg from amf sc, but error happend(%s)", err)
		} else {
			p.handleAmfScMsg(data)
		}
	}
}

func (p *GnbLayer) notifyDestroy() {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)
	// notify the sctp shutdown, trigger sc to clean all ue release information.

	err := p.SendGnbSctpShutdownMessages()
	if err != nil {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to send gnb sctp shutdown msg to sc,error(%s)", err)
	}

}

func (p *GnbLayer) handleAmfScMsg(data *types.IpcMsgData) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	if nil == data {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
			"receive invalid message from amf sc module")
		return
	}

	msgData := &types.MsgData{}
	msgData.MsgLen = len(data.Data)
	msgData.MsgData = string(data.Data)
	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
		"receive msg from sc (%s) to gnb(%s),msg len(%d)",
		data.Sender, data.Receiver, msgData.MsgLen)

	p.sendch <- msgData
}
