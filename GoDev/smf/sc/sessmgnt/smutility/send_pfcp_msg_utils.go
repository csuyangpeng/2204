package smutility

import (
	"context"
	"lite5gc/cmn/message/pfcp"
	"lite5gc/cmn/message/pfcp/v1"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/router"
	"lite5gc/cmn/routeragent"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/utils"
	"lite5gc/smf/n4layer/typedef"
	"lite5gc/smf/smfcontext/gctxt"
	"net"
)

func SendMsg2PfcpNode(ctxt context.Context,
	pduSessCtxt *gctxt.PduSessContext, msgData *gctxt.ScN4MsgData, id pfcp.PFCPMSG) error {

	rlogger.FuncEntry(types.ModuleSmfSM, nil)
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "PFCP MSG:", id)

	if pduSessCtxt == nil || msgData == nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "input para is nil")
		return types.ErrInputParaNil
	}
	// set common data
	msgData.SEID = pduSessCtxt.SEID
	//msgData.DNN = pduSessCtxt.DNN
	msgData.SNSSAI = pduSessCtxt.SNSSAI
	msgData.UEIP = pduSessCtxt.UEIP

	//1 todo:the selected UPF
	//cmnet-1_1:172.20.0.50
	//dnnSnssaiUpfIpMap := fmt.Sprintf("%s-%s:%d", configure.SmfConf.UpfSelection[0].DnnName,
	//	configure.SmfConf.UpfSelection[0].Snssai, configure.SmfConf.UpfSelection[0].UpfIp)
	// trans snssai to string format
	temp := make([]byte, 4)
	copy(temp[1:4], pduSessCtxt.SNSSAI.Sd[:])
	//snssaiStr := fmt.Sprintf("%d_%d", pduSessCtxt.SNSSAI.Sst, binary.BigEndian.Uint32(temp))
	//key := pduSessCtxt.DNN.String() + "-" + snssaiStr
	upfIp := configure.SmfConf.UpfSelection[0].UpfIp
	//upfIp, ok := dnnSnssaiUpfIpMap[key]
	//rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "dnnSnssaiUpfIpMap,key", dnnSnssaiUpfIpMap, key, upfIp)
	//if !ok {
	//	rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, nil, "fail to get upf ip,configure infor:(%s),input key:(%s)", dnnSnssaiUpfIpMap, key)
	//	return fmt.Errorf("fail to get upf ip,configure infor:(%s),input key:(%s)", dnnSnssaiUpfIpMap, key)
	//}
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "upf ip", upfIp)

	//2 send n4 session establish request
	// CN Tunnel Info allocated by SMF, so send n4 session establishment request to UPF
	// sync RPC, finished here.
	msg := pfcpv1.SmfToNode{}
	msg.ID = id
	msg.Cxt = msgData

	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "pduSessCtxt.PFCPParameters: %+v", msgData.PFCPParameters)

	serviceMsg := pfcpv1.ServiceMsg{}
	addr := net.UDPAddr{}
	addr.IP = net.ParseIP(upfIp)
	addr.Port = pfcp.Port
	serviceMsg.RemoteAdd = &addr
	serviceMsg.Msg = msg

	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "serviceMsg.RemoteAdd：", serviceMsg.RemoteAdd)

	var scId uint32
	_, err := typedef.GetNode(upfIp)
	if err != nil {
		scId = 0
	} else {
		scId = utils.IpToUint32(addr.IP)
	}
	rlogger.Trace(types.ModuleSmfSM, rlogger.INFO, nil, "scId:", scId)

	err = routeragent.SendIpcMessage(ctxt, router.PfcpNodeGR, scId, serviceMsg)
	if err != nil {
		rlogger.Trace(types.ModuleSmfSM, rlogger.ERROR, ctxt, "failed to send ipv msg")
		return err
	}
	return nil
}
