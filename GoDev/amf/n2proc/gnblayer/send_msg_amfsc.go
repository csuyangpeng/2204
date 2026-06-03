package gnblayer

import (
	"fmt"
	"lite5gc/cmn/redisclt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// SendScMsg send message to sc goroutine
func (p *GnbLayer) SendMsg2AmfSC(scInst uint32, msg *types3gpp.Gnb2AmfScMsg) error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	var scid uint32 = scInst
	if scInst == types3gpp.InvalidInstId {
		// get all sc instance id randomly
		ids, err := GetAllScInstId()
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
				"failed to get sc instance id(%s)", err)
			return err
		}
		scid, err = GetRandValue(ids)
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
				"failed to get random sc instance id(%s)", err)
			return err
		}
	}

	key := fmt.Sprintf("%s%d", types.AmfProc, scid)
	err := redisclt.Agent.LPush(key, *msg)
	if err != nil {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to send msg to sc(%d)", scid)
		return fmt.Errorf("failed to send msg to sc(%d)", scid)
	}

	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil, "send message to %s, %s", key, msg)
	return nil
}

func (p *GnbLayer) Broadcast2AmfSc(msg *types3gpp.Gnb2AmfScMsg) error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	// get all sc instance id
	ids, err := GetAllScInstId()
	if err != nil {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to get sc ids.(%s)", err)
		return nil
	}
	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
		"ids %v", ids)

	for _, v := range ids {
		err := p.SendMsg2AmfSC(v, msg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to send msg to sc (%s)", err)
		}
	}

	return nil
}
func (p *GnbLayer) SendGnbSctpShutdownMessages() error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	//construct Ngap 2 sc message
	msg := &types3gpp.Gnb2AmfScMsg{}
	msg.MsgType = types3gpp.SctpShutdown
	msg.PrcdCode = types3gpp.MaxProcedureCode
	msg.GnbInfo = p.gnbInfo
	msg.NgapMsg = nil

	return p.Broadcast2AmfSc(msg)
}

func (p *GnbLayer) SendTaTableUpdatedMessage() error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	//construct Ngap 2 sc message
	msg := &types3gpp.Gnb2AmfScMsg{}
	msg.MsgType = types3gpp.TaTblUpdate
	msg.PrcdCode = types3gpp.MaxProcedureCode
	msg.GnbInfo = p.gnbInfo
	msg.NgapMsg = nil

	return p.Broadcast2AmfSc(msg)
}

func GetAllScInstId() ([]uint32, error) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	var ids []uint32
	scids, err := redisclt.Agent.SetMembers(types.AmfScInstId)
	if err != nil {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to get sc inst ids")
		return nil, fmt.Errorf("failed to get sc inst ids")
	}

	for _, v := range scids {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
			"scid type(%s), value(%s)", reflect.TypeOf(v), v)
		vstr := fmt.Sprintf("%s", v)
		//vstr, ok := v.(string)
		//if !ok {
		//	return nil, fmt.Errorf("invalid sc inst id")
		//}
		id, err := strconv.Atoi(vstr)
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"invalid sc inst id(%s)", vstr)
			continue
		}

		ids = append(ids, uint32(id))
	}

	return ids, nil
}

// random select a element from the slice
func GetRandValue(ids []uint32) (uint32, error) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	if len(ids) == 0 {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil, "empty slice for ids")
		return 0, fmt.Errorf("empty slice for ids")
	}

	return ids[rand.Intn(len(ids))], nil
}
