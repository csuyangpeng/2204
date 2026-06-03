package ngapmsg

import (
	"C"
	"encoding/binary"
	"fmt"
	"github.com/willf/bitset"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	T "lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"net"
	"unsafe"
)

//PduSessResMdfyRespTransfer struct definition
type PduSessResMdfyRespTransfer struct {
	DlGtpTunnel                T.GtpTunnel
	UlGtpTunnel                T.GtpTunnel
	AddQosFlowPerTNLInfo       []T.AddQosFlowPerTnlInfo
	QosFlowAddOrModifyRespList []byte
	QosFlowFailedMdfyList      []*T.QosFlowDesc
	OptFlags                   bitset.BitSet

	ctxt codec.NgapOssCtxt
}

const (
	PSRMRespTDlGtpTunnel = iota
	PSRMRespTUlGtpTunnel
	PSRMRespTAddQosFlowTNLInfo
	PSRMRespTQosFlowAddOrModifyResp
	PSRMRespTQosFlowFailAddOrModify
)

func (p *PduSessResMdfyRespTransfer) String() string {
	outStr := fmt.Sprintf("PduSessResMdfyRespTransfer{"+
		"[%v DlTransportLayerInfo(%s)],[%v UlTransportLayerInfo(%s)], AddQosFlowTNLInfoPrst[%v](",
		p.OptFlags.Test(PSRMRespTDlGtpTunnel), p.DlGtpTunnel, p.OptFlags.Test(PSRMRespTUlGtpTunnel), p.UlGtpTunnel,
		p.OptFlags.Test(PSRMRespTAddQosFlowTNLInfo))

	var AddQosStr string
	outStr += "AddQosPerTnlInfoList{"
	for i, v := range p.AddQosFlowPerTNLInfo {
		AddQosStr += fmt.Sprintf("\n    { #%d %+v}", i+1, v.AssocQosFlowList)
	}

	//outStr += "AssociatedQosFlowList{"
	//for _, v := range p.QosFlowIndlist {
	//	outStr += fmt.Sprintf("%+v ", v)
	//}

	outStr = outStr + AddQosStr + "QosFlowAddOrModifyRespList{"
	outStr += fmt.Sprintf("[%v]", p.OptFlags.Test(PSRMRespTQosFlowAddOrModifyResp))
	for _, v := range p.QosFlowAddOrModifyRespList {
		outStr += fmt.Sprintf("%+v ", v)
	}

	outStr += ",QosFlowFailedToAddOrModifyList{"
	outStr += fmt.Sprintf("[%v]", p.OptFlags.Test(PSRMRespTQosFlowFailAddOrModify))
	for _, v := range p.QosFlowFailedMdfyList {
		outStr += fmt.Sprintf("%+v,", v)
	}
	outStr += "}"
	return outStr
}

//AddQosFlowPerTNLInformationList add a QosFlowPerTNLInformation
func (p *PduSessResMdfyRespTransfer) AddQosFlowPerTNLInfoList(item *T.AddQosFlowPerTnlInfo) {
	p.AddQosFlowPerTNLInfo = append(p.AddQosFlowPerTNLInfo, *item)
}

//AddQosFlowAddOrMdfyRespList add a QosFlowAddOrMdfyResp into PduSessResMdfyRespTransfer
func (p *PduSessResMdfyRespTransfer) AddQosFlowAddOrModifyRespList(item byte) {
	p.QosFlowAddOrModifyRespList = append(p.QosFlowAddOrModifyRespList, item)
}

//AddQosFlowFailedMdfyList add a QosFlowFailedMdfy into PduSessResMdfyRespTransfer
func (p *PduSessResMdfyRespTransfer) AddQosFlowFailedMdfyList(item *T.QosFlowDesc) {
	p.QosFlowFailedMdfyList = append(p.QosFlowFailedMdfyList, item)
}

// NewPduSessResMdfyRespTransfer create a new Message
func NewPduSessResMdfyRespTransfer() *PduSessResMdfyRespTransfer {
	return &PduSessResMdfyRespTransfer{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResMdfyRespTransfer) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *PduSessResMdfyRespTransfer) Encode() []byte {
	transfer := codec.NewPduSessResMdfyRespTransferCodec()
	defer codec.DeletePduSessResMdfyRespTransferCodec(transfer)

	if p.OptFlags.Test(PSRMRespTDlGtpTunnel) {
		dlIpaddr := make([]byte, 4)
		binary.BigEndian.PutUint32(dlIpaddr, p.DlGtpTunnel.GetIpv4Long())
		dlTeid := make([]byte, 4)
		binary.BigEndian.PutUint32(dlTeid, p.DlGtpTunnel.GetTeid())

		dlGtpTunnel := codec.NewGtpTunnel()
		defer codec.DeleteGtpTunnel(dlGtpTunnel)
		dlGtpTunnel.SetIpType(0) //hardcode to ipv4, in ngap spec, ipv4:0, ipv6:1
		dlGtpTunnel.SetTransportLayerAddr(&dlIpaddr[0])
		dlGtpTunnel.SetGtpTeid(&dlTeid[0])

		transfer.SetDlGtpTunnel(dlGtpTunnel)
	}

	if p.OptFlags.Test(PSRMRespTUlGtpTunnel) {
		ulIpaddr := make([]byte, 4)
		binary.BigEndian.PutUint32(ulIpaddr, p.UlGtpTunnel.GetIpv4Long())
		ulTeid := make([]byte, 4)
		binary.BigEndian.PutUint32(ulTeid, p.UlGtpTunnel.GetTeid())

		ulGtpTunnel := codec.NewGtpTunnel()
		defer codec.DeleteGtpTunnel(ulGtpTunnel)
		ulGtpTunnel.SetIpType(0) //hardcode to ipv4, in ngap spec, ipv4:0, ipv6:1
		ulGtpTunnel.SetTransportLayerAddr(&ulIpaddr[0])
		ulGtpTunnel.SetGtpTeid(&ulTeid[0])

		transfer.SetUlGtpTunnel(ulGtpTunnel)
	}

	//AddQosFlowPerTNLInformationList
	if p.OptFlags.Test(PSRMRespTAddQosFlowTNLInfo) {
		for _, qosFlowTNLInfo := range p.AddQosFlowPerTNLInfo {
			addQosFlowInfo := codec.NewAddQosFlowPerTNLInfo()
			defer codec.DeleteAddQosFlowPerTNLInfo(addQosFlowInfo)

			addIpaddr := make([]byte, 4)
			binary.BigEndian.PutUint32(addIpaddr, qosFlowTNLInfo.AddUpTransLayerInfo.GetIpv4Long())
			addTeid := make([]byte, 4)
			binary.BigEndian.PutUint32(addTeid, qosFlowTNLInfo.AddUpTransLayerInfo.GetTeid())

			addGtpTunnel := codec.NewGtpTunnel()
			defer codec.DeleteGtpTunnel(addGtpTunnel)
			addGtpTunnel.SetIpType(0) //hardcode to ipv4, in ngap spec, ipv4:0, ipv6:1
			addGtpTunnel.SetTransportLayerAddr(&addIpaddr[0])
			addGtpTunnel.SetGtpTeid(&addTeid[0])

			addQosFlowInfo.SetUpTransportLayerInfo(addGtpTunnel)

			qosFlowInfoVector := codec.NewAssQosFlowItemVector()
			defer codec.DeleteAssQosFlowItemVector(qosFlowInfoVector)

			for _, v := range qosFlowTNLInfo.AssocQosFlowList {
				item := codec.NewAssQosFlowItem()
				defer codec.DeleteAssQosFlowItem(item)
				item.SetQosFlowInd(uint(v.QosFlowIdentifier))
				if v.IsMappingIndPrst {
					item.SetQosFlowMapIndPrst(true)
					item.SetQosFlowMapInd(byte(v.MappingIndication))
				}
				qosFlowInfoVector.Add(item)
			}
			addQosFlowInfo.SetAssQosFlowList(qosFlowInfoVector)
			transfer.AddQosFlowTNLInfoList(addQosFlowInfo)
		}
	}

	//QosFlowAddOrModifyResponse
	if p.OptFlags.Test(PSRMRespTQosFlowAddOrModifyResp) {
		for _, v := range p.QosFlowAddOrModifyRespList {
			transfer.AddQosFlowAddOrMdfyRespList(uint(v))
		}
	}

	// QosFlowFailedToModifyList
	if p.OptFlags.Test(PSRMRespTQosFlowFailAddOrModify) {
		for _, v := range p.QosFlowFailedMdfyList {
			qosflow := codec.NewQosFlowCodecItem()
			defer codec.DeleteQosFlowCodecItem(qosflow)
			qosflow.SetQosFlowInd(uint(v.QowFlowInd))
			qosflow.SetCauseType(byte(v.Cause.Type))
			qosflow.SetCauseValue(byte(v.Cause.Value))
			transfer.AddQosFlowFailedAddOrMdfyList(qosflow)
		}
	}

	msgBuffer := transfer.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResMdfyRespTransfer) Decode(msgbuf []byte) error {

	transfer := codec.NewPduSessResMdfyRespTransferCodec()
	defer codec.DeletePduSessResMdfyRespTransferCodec(transfer)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if transfer.Decode(p.ctxt, msgBuffer) == true {

		//dl gtp tunnel
		if transfer.IsDlGtpTunnelPrst() {
			p.OptFlags.Set(PSRMRespTDlGtpTunnel)
			if transfer.GetDlGtpTunnel().GetIpType() == 0 { //only ipv4 supported, todo support ipv6
				ipaddr := utils.Conv2ByteSlice(transfer.GetDlGtpTunnel().GetTransportLayerAddr(), 4)
				p.DlGtpTunnel.SetIpAddr(net.IPv4(ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]))
			} else {
				rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ipv6 unsupported")
			}
			teid := utils.Conv2ByteSlice(transfer.GetDlGtpTunnel().GetGtpTeid(), 4)
			p.DlGtpTunnel.SetTeid(binary.BigEndian.Uint32(teid))
		}

		//ul gtp tunnel
		if transfer.IsUlGtpTunnelPrst() {
			p.OptFlags.Set(PSRMRespTUlGtpTunnel)
			if transfer.GetUlGtpTunnel().GetIpType() == 0 { //only ipv4 supported, todo support ipv6
				ipaddr := utils.Conv2ByteSlice(transfer.GetUlGtpTunnel().GetTransportLayerAddr(), 4)
				p.UlGtpTunnel.SetIpAddr(net.IPv4(ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]))
			} else {
				rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ipv6 unsupported")
			}
			teid := utils.Conv2ByteSlice(transfer.GetUlGtpTunnel().GetGtpTeid(), 4)
			p.UlGtpTunnel.SetTeid(binary.BigEndian.Uint32(teid))
		}

		//add qos flow per tnl information list
		if transfer.IsQosFlowTNLInfoListPrst() {
			p.OptFlags.Set(PSRMRespTAddQosFlowTNLInfo)

			addQosFlowTnlInfoList := transfer.GetQosFlowTNLInfoList()
			for i := 0; i < int(addQosFlowTnlInfoList.Size()); i++ {
				addQosFlowTnlInfo := addQosFlowTnlInfoList.Get(i)
				CPaddQosFlowTnlInfo := T.AddQosFlowPerTnlInfo{}

				if addQosFlowTnlInfo.GetUpTransportLayerInfo().GetIpType() == 0 { //only ipv4 supported, todo support ipv6
					ipaddr := utils.Conv2ByteSlice(addQosFlowTnlInfo.GetUpTransportLayerInfo().GetTransportLayerAddr(), 4)

					CPaddQosFlowTnlInfo.AddUpTransLayerInfo.SetIpAddr(net.IPv4(ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]))
				} else {
					rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ipv6 unsupported")
				}

				teid := utils.Conv2ByteSlice(addQosFlowTnlInfo.GetUpTransportLayerInfo().GetGtpTeid(), 4)
				CPaddQosFlowTnlInfo.AddUpTransLayerInfo.SetTeid(binary.BigEndian.Uint32(teid))

				assQosFlowList := addQosFlowTnlInfo.GetAssQosFlowList()
				for j := 0; j < int(assQosFlowList.Size()); j++ {
					aItem := &T.AssocQosFlowItem{}
					aItem.QosFlowIdentifier = int(assQosFlowList.Get(i).GetQosFlowInd())
					aItem.IsMappingIndPrst = assQosFlowList.Get(i).GetQosFlowMapIndPrst()
					if aItem.IsMappingIndPrst {
						aItem.MappingIndication = T.QosFlowMapInd(assQosFlowList.Get(i).GetQosFlowMapInd())
					}
					CPaddQosFlowTnlInfo.AddAssocQosFlowItem(aItem)
				}

				p.AddQosFlowPerTNLInfo = append(p.AddQosFlowPerTNLInfo, CPaddQosFlowTnlInfo)
			}
		}

		if transfer.IsQosFlowAddOrMdfyRespListPrst() {
			p.OptFlags.Set(PSRMRespTQosFlowAddOrModifyResp)
			qosFlowIdVec := transfer.GetQosFlowAddOrMdfyRespList()
			for i := 0; i < int(qosFlowIdVec.Size()); i++ {
				psi := qosFlowIdVec.Get(i)
				p.QosFlowAddOrModifyRespList = append(p.QosFlowAddOrModifyRespList, byte(psi))
			}
		}

		if transfer.IsFailedAddOrMdfyListPrst() {
			p.OptFlags.Set(PSRMRespTQosFlowFailAddOrModify)
			qosflowList := transfer.GetQosFlowFailedAddOrMdfyList()
			for i := 0; i < int(qosflowList.Size()); i++ {
				qosflow := qosflowList.Get(i)

				pQosFlow := &T.QosFlowDesc{}
				pQosFlow.QowFlowInd = uint8(qosflow.GetQosFlowInd())
				pQosFlow.Cause.Type = T.CauseType(qosflow.GetCauseType())
				pQosFlow.Cause.Value = T.CauseValue(qosflow.GetCauseValue())

				p.QosFlowFailedMdfyList = append(p.QosFlowFailedMdfyList, pQosFlow)
			}
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
