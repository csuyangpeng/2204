package ngapmsg

import "C"
import (
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

//PduSessResSetupRespTransfer struct definition
type PduSessResSetupRespTransfer struct {
	GtpTunnel      T.GtpTunnel
	QosFlowIndlist []*T.AssocQosFlowItem

	AddQosFlowPerTNLInfo   []T.AddQosFlowPerTnlInfo
	SecResult              T.SecurityResult
	QosFlowFailedSetupList []*T.QosFlowDesc
	OptFlags               bitset.BitSet

	ctxt codec.NgapOssCtxt
}

const (
	PSRSRespT_SecRst = iota
	PSRSRespT_QosFlowFailList
	PSRSRespT_AddQosFlowTNLInfo
)

func (p *PduSessResSetupRespTransfer) String() string {
	outStr := fmt.Sprintf("PduSessResSetupRespTransfer{"+
		"TransportLayerInfo(%s), QosFlowIndList(", p.GtpTunnel)

	for _, v := range p.QosFlowIndlist {
		outStr += fmt.Sprintf("%s ", v)
	}

	var AddQosStr string
	outStr += "AddQosPerTnlInfoList{"
	for i, v := range p.AddQosFlowPerTNLInfo {
		AddQosStr += fmt.Sprintf("\n    { #%d %+v}", i+1, v.AssocQosFlowList)
	}

	outStr += fmt.Sprintf(", [%v SecResult(%s)]",
		p.OptFlags.Test(PSRSRespT_SecRst), p.SecResult)

	outStr += ",QosFlowFailedSetup{"
	for _, v := range p.QosFlowFailedSetupList {
		outStr += fmt.Sprintf("%s,", v)
	}
	outStr += "}"
	return outStr
}

//AddQosFlowFailedSetupList add a QosFlowDesc into PduSessResSetupRespTransfer
func (p *PduSessResSetupRespTransfer) AddQosFlowIndList(item *T.AssocQosFlowItem) {
	p.QosFlowIndlist = append(p.QosFlowIndlist, item)
}

//AddQosFlowPerTNLInformationList add a QosFlowPerTNLInformation
func (p *PduSessResSetupRespTransfer) AddQosFlowPerTNLInfoList(item *T.AddQosFlowPerTnlInfo) {
	p.AddQosFlowPerTNLInfo = append(p.AddQosFlowPerTNLInfo, *item)
}

//AddQosFlowFailedSetupList add a QosFlowDesc into PduSessResSetupRespTransfer
func (p *PduSessResSetupRespTransfer) AddQosFlowFailedSetupList(item *T.QosFlowDesc) {
	p.QosFlowFailedSetupList = append(p.QosFlowFailedSetupList, item)
}

// NewPduSessResSetupReqTransfer create a new Message
func NewPduSessResSetupRespTransfer() *PduSessResSetupRespTransfer {

	return &PduSessResSetupRespTransfer{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResSetupRespTransfer) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

func (p *PduSessResSetupRespTransfer) Encode() []byte {
	transfer := codec.NewPduSesResSetupRespTransferCodec()
	defer codec.DeletePduSesResSetupRespTransferCodec(transfer)

	ipaddr := make([]byte, 4)
	binary.BigEndian.PutUint32(ipaddr, p.GtpTunnel.GetIpv4Long())
	teid := make([]byte, 4)
	binary.BigEndian.PutUint32(teid, p.GtpTunnel.GetTeid())

	gtpTunnel := codec.NewGtpTunnel()
	defer codec.DeleteGtpTunnel(gtpTunnel)
	gtpTunnel.SetIpType(0) //hardcode to ipv4, in ngap spec, ipv4:0, ipv6:1
	gtpTunnel.SetTransportLayerAddr(&ipaddr[0])
	gtpTunnel.SetGtpTeid(&teid[0])

	transfer.SetGtpTunnel(gtpTunnel)
	for _, v := range p.QosFlowIndlist {
		item := codec.NewAssQosFlowItem()
		defer codec.DeleteAssQosFlowItem(item)
		item.SetQosFlowInd(uint(v.QosFlowIdentifier))
		if v.IsMappingIndPrst {
			item.SetQosFlowMapIndPrst(true)
			item.SetQosFlowMapInd(byte(v.MappingIndication))
		}
		transfer.AddQosFlowIndList(item)
	}

	//AddDlQosFlowPerTNLInformationList
	if p.OptFlags.Test(PSRSRespT_AddQosFlowTNLInfo) {
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

	// security result
	if p.OptFlags.Test(PSRSRespT_SecRst) {
		transfer.SetSecResult(byte(p.SecResult.IntPrctRst), byte(p.SecResult.ConfdPrctRst))
	}

	// QosFlowFailedSetupList
	if p.OptFlags.Test(PSRSRespT_QosFlowFailList) {
		for _, v := range p.QosFlowFailedSetupList {
			qosflow := codec.NewQosFlowCodecItem()
			defer codec.DeleteQosFlowCodecItem(qosflow)
			qosflow.SetQosFlowInd(uint(v.QowFlowInd))
			qosflow.SetCauseType(byte(v.Cause.Type))
			qosflow.SetCauseValue(byte(v.Cause.Value))
			transfer.AddQosFlowFailedSetupList(qosflow)
		}
	}

	msgBuffer := transfer.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResSetupRespTransfer) Decode(msgbuf []byte) error {

	transfer := codec.NewPduSesResSetupRespTransferCodec()
	defer codec.DeletePduSesResSetupRespTransferCodec(transfer)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if transfer.Decode(p.ctxt, msgBuffer) == true {

		// pdu gtp tunnel info
		if transfer.GetGtpTunnel().GetIpType() == 0 { //only ipv4 supported, todo support ipv6
			ipaddr := utils.Conv2ByteSlice(transfer.GetGtpTunnel().GetTransportLayerAddr(), 4)
			p.GtpTunnel.SetIpAddr(net.IPv4(ipaddr[0], ipaddr[1], ipaddr[2], ipaddr[3]))
		} else {
			rlogger.Trace(types.ModCmn, rlogger.DEBUG, nil, "ipv6 unsupported")
		}
		teid := utils.Conv2ByteSlice(transfer.GetGtpTunnel().GetGtpTeid(), 4)
		p.GtpTunnel.SetTeid(binary.BigEndian.Uint32(teid))

		qfiList := transfer.GetQosFlowIndList()
		for i := 0; i < int(qfiList.Size()); i++ {
			pItem := &T.AssocQosFlowItem{}
			pItem.QosFlowIdentifier = int(qfiList.Get(i).GetQosFlowInd())
			pItem.IsMappingIndPrst = qfiList.Get(i).GetQosFlowMapIndPrst()
			if pItem.IsMappingIndPrst {
				pItem.MappingIndication = T.QosFlowMapInd(qfiList.Get(i).GetQosFlowMapInd())
			}
			p.QosFlowIndlist = append(p.QosFlowIndlist, pItem)
		}

		//add dl qos flow per tnl information list
		if transfer.IsQosFlowTNLInfoListPrst() {
			p.OptFlags.Set(PSRSRespT_AddQosFlowTNLInfo)

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

		if transfer.IsSecResultPrst() {
			p.OptFlags.Set(PSRSRespT_SecRst)
			p.SecResult.IntPrctRst = T.PerformRst(transfer.GetIndPrctResult())
			p.SecResult.ConfdPrctRst = T.PerformRst(transfer.GetConfdPrctResult())
		}

		if transfer.IsFailedSetupListPrst() {
			p.OptFlags.Set(PSRSRespT_QosFlowFailList)
			qosflowList := transfer.GetQosFlowFailedSetupList()
			for i := 0; i < int(qosflowList.Size()); i++ {
				qosflow := qosflowList.Get(i)

				pQosFlow := &T.QosFlowDesc{}
				pQosFlow.QowFlowInd = uint8(qosflow.GetQosFlowInd())
				pQosFlow.Cause.Type = T.CauseType(qosflow.GetCauseType())
				pQosFlow.Cause.Value = T.CauseValue(qosflow.GetCauseValue())

				p.QosFlowFailedSetupList = append(p.QosFlowFailedSetupList, pQosFlow)
			}
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
