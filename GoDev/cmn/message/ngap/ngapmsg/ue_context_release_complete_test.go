package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"lite5gc/cmn/types3gpp"
	"testing"
)

func TestUeContextReleaseCmpMsg_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewUeContextReleaseCmpMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.AmfUeNgapId = 100
	msg.RanUeNgapId = 200
	msg.IsUliPrst = true
	msg.UserLocInfo.IsTimeStampPrst = true
	msg.UserLocInfo.TimeStmp = types3gpp.TimeStampType{1, 2, 3, 4}
	msg.UserLocInfo.Tai.Tac.SetTac(512)
	msg.UserLocInfo.Tai.Plmn.SetString("46000")
	msg.UserLocInfo.NrCgi.PlmnID.SetString("46001")
	msg.UserLocInfo.NrCgi.NrCellId = types3gpp.NRCellId{1, 2, 3, 4, 5}

	msg.IsRpiPrst = true
	cell := types3gpp.RecommendCell4Page{}
	cell.IsTSPrst = true
	cell.TimeStayed = 200
	cell.Cgi.PlmnID.SetString("46012")
	cell.Cgi.NrCellId = types3gpp.NRCellId{2, 2, 3, 4, 5}
	msg.RcmdPagingInfo.CellInfo = append(msg.RcmdPagingInfo.CellInfo, cell)

	node := types3gpp.RecommendRanNode4Page{}
	node.PageTarget = types3gpp.PT_RanNode
	node.RanNodeId.GNBID.Plmn.SetString("46012")
	node.RanNodeId.GNBID.GNBID = [types3gpp.SizeofGNBID]byte{1, 2, 3, 4}
	node.RanNodeId.GNBID.GNBLength = 32
	msg.RcmdPagingInfo.RanNodeInfo = append(msg.RcmdPagingInfo.RanNodeInfo, node)

	msg.PduSessIdList = append(msg.PduSessIdList, 5)
	msg.PduSessIdList = append(msg.PduSessIdList, 6)
	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestUeContextReleaseCmpMsg_Decode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	encMsg := []byte{32, 41, 0, 73, 0, 0, 5, 0, 60, 0, 5, 1, 0, 6, 0, 5, 0, 32, 64, 22, 0, 64, 100, 240, 33, 2, 2, 3, 4, 0, 0, 200, 0, 0, 100, 240, 33, 80, 1, 2, 3, 4, 0, 121, 64, 19, 80, 100, 240, 16, 1, 2, 3, 4, 0, 100, 240, 0, 0, 2, 0, 1, 2, 3, 4, 0, 10, 64, 2, 0, 100, 0, 85, 64, 2, 0, 200}

	cmnNgAp := codec.NewNgapCodec()
	cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))

	cmnNgAp.DecodeHeader(ossCtxtPtr)

	fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	msg := NewUeContextReleaseCmpMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Printf("decode Msg: %s", msg)
}
