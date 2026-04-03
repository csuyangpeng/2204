package gnblayer

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/juliangruber/go-intersect"
	"lite5gc/amf/tatable"
	"lite5gc/cmn/message/ngap/ngapmsg"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/configure"
	"lite5gc/cmn/types3gpp"
)

func (p *GnbLayer) handleNgSetupRequestMsg(msgBuf *types.MsgBuf) error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	ngSetupRequestMsg := ngapmsg.NewNgSetupRequestMsg()
	ngSetupRequestMsg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())
	err := ngSetupRequestMsg.Decode(msgBuf.Buffer)
	if err != nil {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
			"failed to decode Ng Setup Request Message")
		return err
	}

	// save the decoded ies
	p.gnbInfo.GnbId.Plmn = ngSetupRequestMsg.GRanNodeID.GNBID.Plmn
	p.gnbInfo.GnbId.GnbId = binary.LittleEndian.Uint32(ngSetupRequestMsg.GRanNodeID.GNBID.GNBID[:])
	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil,
		"gnb id(%d)", p.gnbInfo.GnbId.GnbId)
	p.gnbInfo.GnbId.IdType = types3gpp.GlobalGNB

	if ngSetupRequestMsg.RanNodeNamePrst {
		p.ranNodeName = ngSetupRequestMsg.RanNodeName
	}
	p.defPagingDrx = ngSetupRequestMsg.DefPagingDRX

	p.supTaList = ngSetupRequestMsg.SupportTAs

	// check gnb ip and gnb id
	ipExist, err := tatable.CheckGnbIpAddr(p.GetGnbIP())
	if (err == nil) && (ipExist == false) {
		exist, err := tatable.CheckGnbIdUniqueness(p.gnbInfo.GnbId.GnbId)
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil,
				"failed to check，error(%s)", err)
			return err
		}
		if exist {
			//trigger alarm
			// GnbId already exists, return gnbId duplicate
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"gnb id(%d) duplicate", p.gnbInfo.GnbId.GnbId)
			p.sendNgSetupFailureMsg(types3gpp.CT_RadioNetwork, types3gpp.Radiok_unknown_targetID)
			return errors.New("gnb id duplicate")

		}
	} else { // gnbIP already exists, return gnbIP duplicate
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
			"gnb ip (%s) already exist, err ", p.gnbInfo.GnbIP)
		p.sendNgSetupFailureMsg(types3gpp.CT_RadioNetwork, types3gpp.Radiok_unknown_targetID)
		return errors.New("gnb ip duplicate")
	}

	// check the support TA list for access control
	var checkResult bool
	for _, v := range ngSetupRequestMsg.SupportTAs {
		checkResult = CheckTai(v)
		if !checkResult {
			p.sendNgSetupFailureMsg(types3gpp.CT_Misc, types3gpp.Misc_unknown_PLMN)
			return nil
		}
	}

	//send ng setup response to gnb
	ngSetupResponseMsg := ngapmsg.NewNgSetupResponseMsg()
	ngSetupResponseMsg.SetOssCodecCtxt(p.ossCtxt.GetOssCtxtPtr_m())

	ngSetupResponseMsg.AmfName = configure.AmfConf.Service.AmfName
	ngSetupResponseMsg.RelativeAmfCapacity = uint16(configure.AmfConf.Service.AmfRelCap)

	for _, v := range configure.AmfConf.PlmnList.List {
		guami := &types3gpp.Guami{}
		guami.PlmnId = v
		guami.AmfId.SetAmfRegionID(configure.AmfConf.Service.AmfIdentifier.GetAmfRegionID())
		guami.AmfId.SetAmfPointer(configure.AmfConf.Service.AmfIdentifier.GetAmfPointer())
		guami.AmfId.SetAmfSetID(configure.AmfConf.Service.AmfIdentifier.GetAmfSetID())
		ngSetupResponseMsg.AddServedGuami(guami)
	}

	for _, v := range configure.AmfConf.PlmnList.List {
		bplmn := &types3gpp.BPlmn{}
		bplmn.Plmn = v

		for _, vv := range configure.AmfConf.Nssai {
			bplmn.AddSnssai(&vv)
		}

		ngSetupResponseMsg.AddPlmnSupport(bplmn)
	}

	respMsg := &types.MsgData{}
	respMsg.MsgData = string(ngSetupResponseMsg.Encode())
	respMsg.MsgLen = len(respMsg.MsgData)

	p.sendch <- respMsg
	//msg counter
	//pm.PegCounter(statistics.NGSetupResponseCounter)

	//creat the gnb info and save in ta table
	gnbInfo := &types3gpp.GnbInformation{}
	gnbInfo.GnbInstId = p.gnbInfo.GnbInstId
	gnbInfo.GnbId = p.gnbInfo.GnbId.GnbId
	gnbInfo.GnbPlmn = p.gnbInfo.GnbId.Plmn.String()
	gnbInfo.GnbIp = p.GetGnbIP()
	gnbInfo.RanNodeName = p.ranNodeName
	gnbInfo.DefPagingDrx = p.defPagingDrx
	for _, v := range p.supTaList {
		var tmp types3gpp.SprtTa
		tmp.Tac = v.Tac
		for _, v := range v.BPlmnList {
			var tv types3gpp.Bplmn
			tv.Plmn = v.Plmn.String()
			tv.SliceList = v.SliceSupportList
			tmp.BPlmns = append(tmp.BPlmns, tv)
		}
		gnbInfo.SprtTaList = append(gnbInfo.SprtTaList, tmp)
	}
	// save the tai into ta table
	for _, v := range p.supTaList {
		tai := types3gpp.TAI{
			Tac:  v.Tac,
			Plmn: v.BPlmnList[0].Plmn,
		}
		err := tatable.AddTai(tai, gnbInfo)
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"failed to add tai with gnb ip:(%s) for gnbInfo,error(%s)", p.gnbInfo.GnbIP, err)
		}
	}

	// send notify message to sc process
	p.SendTaTableUpdatedMessage()
	return nil
}

func CheckTai(tai types3gpp.SupportedTA) bool {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)
	//var result bool
	//check the bplmn list
	if len(tai.BPlmnList) == 0 {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
			"invalid support ta, no BPlmn")
		return false
	}

	//plmns
	ctai := types3gpp.TAI{}
	ctai.Tac = tai.Tac
	for _, v := range tai.BPlmnList {
		ctai.Plmn = v.Plmn

		amfHasTai := false
		for _, vv := range configure.AmfConf.TaiLists {
			if vv.HasTai(ctai) {
				amfHasTai = true
				break
			}
		}

		if amfHasTai {
			rst := intersect.Simple(configure.AmfConf.Nssai, v.SliceSupportList)
			val, _ := rst.([]interface{})
			if len(val) == 0 {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
					"gnb snssai is not supported on AMF.")

				var dbgStr string
				dbgStr = "config nssai ("
				for _, tv := range configure.AmfConf.Nssai {
					dbgStr += fmt.Sprintf("%s,", tv.String())
				}
				dbgStr += "), message nssai ("
				for _, tvv := range v.SliceSupportList {
					dbgStr += fmt.Sprintf("%s,", tvv.String())
				}
				dbgStr += ")"

				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
					"dump info: %s", dbgStr)

				return false
			}

		} else {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil,
				"gnb TAI(%s) is not supported on AMF.", ctai)
			return false
		}
	}

	return true
}
