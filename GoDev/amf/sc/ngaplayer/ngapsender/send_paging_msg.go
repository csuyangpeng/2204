package ngapsender

import (
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
)

func (p *NgapSender) SendPaging(ueCtxt *gctxt.UeContext) error {
	rlogger.FuncEntry(types.ModuleAmfNgap, nil)

	if ueCtxt.GetProcCtxt() != nil {
		rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, nil,
			"ueCtxt.GetProcCtxt() != nil, stop send paging msg",
			ueCtxt.GetProcCtxt())
		return fmt.Errorf("stop send paging msg")
	}
	msg := ngapmsg.NewPagingMsg()
	msg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	msg.UePagingId = ueCtxt.Guti5g.GetStmsi()
	rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, nil, "ueCtxt.Guti5g", ueCtxt.Guti5g)
	rlogger.Trace(types.ModuleAmfNgap, rlogger.INFO, nil, "ueCtxt.Guti5g.GetStmsi", ueCtxt.Guti5g.GetStmsi())
	//TODO msg.TaiListForPaging = configure.AmfConf.TaiLists.GetTaisPointers()

	msgBuf := types.MsgBuf{}
	msgBuf.Buffer = msg.Encode() //encode ngap message
	msgBuf.MsgLen = len(msgBuf.Buffer)

	//msg counter
	pm.PegCounter(statistics.PagingCounter)

	//get all gnb info
	//gnbCtxts, err := ngap.ValuesOfGnbInfoTbl(ngap.GnbIPKeyType)
	//if err != nil {
	//	rlogger.Trace(types.ModAmfScNgap, rlogger.ERROR, ueCtxt, "Failed to get gnb info:%s", err)
	//	return err
	//}

	//遍历所有基站，如果基站所在的TAI在核心网配置的TAI List中，就给该基站发送paging消息
	//for _, v := range gnbCtxts {
	//	for _, vv := range v.SupTaList {
	//		tai := t3.TAI{}
	//		tai.Plmn = vv.BPlmnList[0].Plmn // only one plmn supported now
	//		tai.Tac = vv.Tac
	//		if configure.AmfConf.Service.TaiList.HasTai(tai) {
	//			p.SendNgapMsg(v.GnbInfo.GnbInstId, &msgBuf)
	//		}
	//	}
	//}

	// todo
	// 实际适用情况：
	// 查询UE所在的TAI，找到该TAI所属TAI LIST，往该list下的所有基站广播paging消息
	return nil
}
