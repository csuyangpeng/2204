/** Copyright(C),2020-2022
* Author: zmj
* Date: 11/24/20 3:40 PM
* Description:
 */
package ngapsender

import (
	"fmt"
	"lite5gc/cmn/redisclt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"math/rand"
	"time"
)

func (p *NgapSender) SendNgapMsg(gnbid uint32, msgbuf types.MsgBuf) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	msgdata := types.IpcMsgData{}

	msgdata.Receiver = fmt.Sprintf("%s%d", types.GnbProc, gnbid)
	msgdata.Sender = fmt.Sprintf("%s%d", types.AmfProc, p.scId)
	//msgdata.Data.MsgLen = msgbuf.MsgLen
	//msgdata.Data.MsgData = string(msgbuf.Buffer)
	msgdata.Data = msgbuf.Buffer

	rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, nil,
		"msg data to gnb(%s - >%s, msg data(%x)", msgdata.Sender, msgdata.Receiver, msgdata.Data)

	if gnbid == types3gpp.InvalidInstId {
		//TODO select a randon sc instanceid
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "invalid gnb instance id(%d)", gnbid)
		return fmt.Errorf("invalid gnb instance id(%d)", gnbid)
	}

	err := redisclt.Agent.LPush(fmt.Sprintf("%s%d", types.GnbProc, gnbid), msgdata)
	if err != nil {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		rlogger.Trace(types.ModuleAmfNgap, rlogger.ERROR, nil, "failed to send msg to gnb(%d)", gnbid)
		return fmt.Errorf("failed to send msg to gnb(%d)", gnbid)
	}
	rlogger.Trace(types.ModuleAmfNgap, rlogger.DEBUG, nil,
		"sending msg to gnb success")

	return nil
}
