package ngapmsg

import "C"
import (
	"fmt"
	"github.com/willf/bitset"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	T "lite5gc/cmn/types3gpp"
	"lite5gc/cmn/utils"
	"unsafe"
)

//Pdu Session Resource Release Response Transfer struct definition
type PduSessResRelRespTransfer struct {
	//optional
	SecRATUsageInfo T.SecRATUsageInformation

	OptFlags bitset.BitSet

	ctxt codec.NgapOssCtxt
}

const (
	PSRRRT_SecRATUsageInformation = iota
)

func (p *PduSessResRelRespTransfer) String() string {
	outStr := fmt.Sprintf("SecRatUsageInfo(%v)"+"PduSessUsageReport(RatType:%s)",
		p.OptFlags.Test(PSRRRT_SecRATUsageInformation),
		p.SecRATUsageInfo.PduSessUsageReport.RatType.String())
	for _, v := range p.SecRATUsageInfo.PduSessUsageReport.VolumeTimeReportList {
		outStr += fmt.Sprintf("%+v", v)
	}
	outStr += fmt.Sprintf("QosFlowUsageTimeReport:")
	for _, v := range p.SecRATUsageInfo.QosFlowUsageReportList {
		outStr += fmt.Sprintf("%+v", v)
	}

	outStr += "}"
	return outStr
}

//NewPduSessResRelRespTransfer
func NewPduSessResRelRespTransfer() *PduSessResRelRespTransfer {
	return &PduSessResRelRespTransfer{}
}

// SetOssCodecCtxt set OSS codec context
func (p *PduSessResRelRespTransfer) SetOssCodecCtxt(ctxt codec.NgapOssCtxt) {
	p.ctxt = ctxt
}

// AddPduSessVolTimeReport add PduSessVolTimeReport into PduSessVolTimeReportList
func (p *PduSessResRelRespTransfer) AddPduSessVolTimeReport(timeReport *T.VolumeTimeReport) {
	p.SecRATUsageInfo.PduSessUsageReport.VolumeTimeReportList = append(p.SecRATUsageInfo.PduSessUsageReport.VolumeTimeReportList, *timeReport)
}

// AddQosFlowUsageReport add QosFlowUsageReport into QosFlowUsageReportList
func (p *PduSessResRelRespTransfer) AddQosFlowUsageReport(qosFlowReport *T.QosFlowUsageReport) {
	p.SecRATUsageInfo.QosFlowUsageReportList = append(p.SecRATUsageInfo.QosFlowUsageReportList, *qosFlowReport)
}

func (p *PduSessResRelRespTransfer) Encode() []byte {
	transfer := codec.NewPduSessResRelRespTransferCodec()
	defer codec.DeletePduSessResRelRespTransferCodec(transfer)

	if p.OptFlags.Test(PSRRRT_SecRATUsageInformation) {
		secondaryRATUsageInformation := codec.NewSecRatUsageInformation()
		defer codec.DeleteSecRatUsageInformation(secondaryRATUsageInformation)

		if p.SecRATUsageInfo.IsPduSessUsageReportPrst {
			secondaryRATUsageInformation.SetPduSessUsageReportPrst(true)

			pduSessUsageRprt := codec.NewPduSessUsageReport()
			defer codec.DeletePduSessUsageReport(pduSessUsageRprt)

			pduSessUsageRprt.SetRatType(byte(p.SecRATUsageInfo.PduSessUsageReport.RatType))

			timeRprtList := codec.NewVolumeTimeReportVector()
			defer codec.DeleteVolumeTimeReportVector(timeRprtList)

			for _, v := range p.SecRATUsageInfo.PduSessUsageReport.VolumeTimeReportList {
				timeRprt1 := codec.NewVolumeTimeReport()
				defer codec.DeleteVolumeTimeReport(timeRprt1)

				timeRprt1.SetStartTimeStamp(&(v.StartTimeStamp.GetByteSlice()[0]))
				timeRprt1.SetEndTimeStamp(&(v.EndTimeStamp.GetByteSlice()[0]))
				timeRprt1.SetUsageCountUL(v.UsageCountUplink)
				timeRprt1.SetUsageCountDL(v.UsageCountDownlink)

				timeRprtList.Add(timeRprt1)
			}
			pduSessUsageRprt.SetVolumeTimeReportList(timeRprtList)

			secondaryRATUsageInformation.SetPduSessUsageReport(pduSessUsageRprt)
		}

		if p.SecRATUsageInfo.IsQosFlowUsageReportListPrst {
			secondaryRATUsageInformation.SetQosFlowUsageReportListPrst(true)

			qosFlowUsageRprtList := codec.NewQosFlowUsageReportVector()
			defer codec.DeleteQosFlowUsageReportVector(qosFlowUsageRprtList)

			for _, v := range p.SecRATUsageInfo.QosFlowUsageReportList {
				qosFlowUsage1 := codec.NewQosFlowUsageReport()
				defer codec.DeleteQosFlowUsageReport(qosFlowUsage1)

				qosFlowUsage1.SetRatType(byte(v.RatType))
				qosFlowUsage1.SetQosFlowId(v.QosFlowId)

				qosFlowTimeRprtList := codec.NewVolumeTimeReportVector()
				defer codec.DeleteVolumeTimeReportVector(qosFlowTimeRprtList)

				for _, vv := range v.QosFlowsTimeReportList {
					qosFlowTimeRprt1 := codec.NewVolumeTimeReport()
					defer codec.DeleteVolumeTimeReport(qosFlowTimeRprt1)

					qosFlowTimeRprt1.SetStartTimeStamp(&(vv.StartTimeStamp.GetByteSlice()[0]))
					qosFlowTimeRprt1.SetEndTimeStamp(&(vv.EndTimeStamp.GetByteSlice()[0]))
					qosFlowTimeRprt1.SetUsageCountUL(vv.UsageCountUplink)
					qosFlowTimeRprt1.SetUsageCountDL(vv.UsageCountDownlink)

					qosFlowTimeRprtList.Add(qosFlowTimeRprt1)
				}
				qosFlowUsage1.SetQosFlowsTimeReportList(qosFlowTimeRprtList)
				qosFlowUsageRprtList.Add(qosFlowUsage1)
			}
			secondaryRATUsageInformation.SetQosFlowUsageReportList(qosFlowUsageRprtList)
		}
		transfer.SetSecRatUsageInfo(secondaryRATUsageInformation)
	}
	msgBuffer := transfer.Encode(p.ctxt)
	defer codec.DeleteMsgBuffer(msgBuffer)
	bufLen := msgBuffer.GetLength()
	bufValue := msgBuffer.GetValue()
	fmt.Println("the code lenth:", bufLen)
	encodeBuffer := C.GoBytes(unsafe.Pointer(bufValue), C.int(bufLen))
	return encodeBuffer
}

func (p *PduSessResRelRespTransfer) Decode(msgbuf []byte) error {
	pduSessResRelRespTransferCodec := codec.NewPduSessResRelRespTransferCodec()
	defer codec.DeletePduSessResRelRespTransferCodec(pduSessResRelRespTransferCodec)

	msgBuffer := codec.NewMsgBuffer()
	defer codec.DeleteMsgBuffer(msgBuffer)

	msgBuffer.SetLength(int64(len(msgbuf)))
	msgBuffer.SetValue(&msgbuf[0])

	if pduSessResRelRespTransferCodec.Decode(p.ctxt, msgBuffer) == true {
		//Secondary RAT Usage Information
		if pduSessResRelRespTransferCodec.IsSecRatUsageInfoPrst() {
			p.OptFlags.Set(PSRRRT_SecRATUsageInformation)
			secRATUsageInfo := pduSessResRelRespTransferCodec.GetSecRatUsageInfo()
			if secRATUsageInfo.GetPduSessUsageReportPrst() {
				p.SecRATUsageInfo.IsPduSessUsageReportPrst = true
				p.SecRATUsageInfo.PduSessUsageReport.RatType = T.RATType(secRATUsageInfo.GetPduSessUsageReport().GetRatType())
				pduSessUsageRprtList := secRATUsageInfo.GetPduSessUsageReport().GetVolumeTimeReportList()
				for i := 0; i < int(pduSessUsageRprtList.Size()); i++ {
					pduSessUsageRprt := pduSessUsageRprtList.Get(i)

					cPduSessUsageRprt := T.VolumeTimeReport{}

					startTimeStamp := utils.Conv2ByteSlice(pduSessUsageRprt.GetStartTimeStamp(), T.SizeofTAC)
					for i, v := range startTimeStamp {
						cPduSessUsageRprt.StartTimeStamp[i] = v
					}
					endTimeStamp := utils.Conv2ByteSlice(pduSessUsageRprt.GetEndTimeStamp(), T.SizeofTAC)
					for i, v := range endTimeStamp {
						cPduSessUsageRprt.EndTimeStamp[i] = v
					}
					cPduSessUsageRprt.UsageCountUplink = pduSessUsageRprt.GetUsageCountUL()
					cPduSessUsageRprt.UsageCountDownlink = pduSessUsageRprt.GetUsageCountDL()

					p.AddPduSessVolTimeReport(&cPduSessUsageRprt)
				}
			}

			if secRATUsageInfo.GetQosFlowUsageReportListPrst() {
				p.SecRATUsageInfo.IsQosFlowUsageReportListPrst = true
				qosFlowUsageRprtList := secRATUsageInfo.GetQosFlowUsageReportList()
				for i := 0; i < int(qosFlowUsageRprtList.Size()); i++ {
					qosFlowUsageRprt := qosFlowUsageRprtList.Get(i)
					cQosFlowUsageRprt := T.QosFlowUsageReport{}

					cQosFlowUsageRprt.RatType = T.RATType(qosFlowUsageRprt.GetRatType())
					cQosFlowUsageRprt.QosFlowId = qosFlowUsageRprt.GetQosFlowId()
					qosFlowTimeRprtList := qosFlowUsageRprt.GetQosFlowsTimeReportList()
					fmt.Println("the di er ge xun huan de ci shu:", int(qosFlowTimeRprtList.Size()))
					for j := 0; j < int(qosFlowTimeRprtList.Size()); j++ {
						qosFlowTimeRprt := qosFlowTimeRprtList.Get(i)
						cQosFlowTimeRprt := T.VolumeTimeReport{}

						startTimeStamp := utils.Conv2ByteSlice(qosFlowTimeRprt.GetStartTimeStamp(), T.SizeofTAC)
						for i, v := range startTimeStamp {
							cQosFlowTimeRprt.StartTimeStamp[i] = v
						}
						endTimeStamp := utils.Conv2ByteSlice(qosFlowTimeRprt.GetEndTimeStamp(), T.SizeofTAC)
						for i, v := range endTimeStamp {
							cQosFlowTimeRprt.EndTimeStamp[i] = v
						}
						cQosFlowTimeRprt.UsageCountUplink = qosFlowTimeRprt.GetUsageCountUL()
						cQosFlowTimeRprt.UsageCountDownlink = qosFlowTimeRprt.GetUsageCountDL()
						fmt.Println("fu zhi qian(uplink):", qosFlowTimeRprt.GetUsageCountUL())
						fmt.Println("fu zhi hou(downlink):", cQosFlowTimeRprt.UsageCountUplink)

						cQosFlowUsageRprt.AddQosFlowTimeReport(&cQosFlowTimeRprt)
					}
					p.AddQosFlowUsageReport(&cQosFlowUsageRprt)
				}
			}
		}
	} else {
		return fmt.Errorf("Failed to decode msg Bufer")
	}

	return nil
}
