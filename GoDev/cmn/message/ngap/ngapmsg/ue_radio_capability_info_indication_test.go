/** Copyright(C),2020-2022
* Author: zmj
* Date: 12/22/20 4:02 PM
* Description:
 */
package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	"testing"
)

func TestUeRadioCapInfoIndMsg_Encode(t *testing.T) {
	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewUeRadioCapInfoIndMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)

	msg.AmfUeNgapId = 100
	msg.RanUeNgapId = 200

	msg.UeRadioCapbility = "123456789"

	encodeMsg := msg.Encode()

	fmt.Println("encode msg: ", encodeMsg)
}

func TestUeRadioCapInfoIndMsg_Decode(t *testing.T) {
	encMsg := []byte{0, 44, 64, 29, 0, 0, 3, 0, 10, 0, 2, 0, 100, 0, 85, 0, 2, 0, 200, 0, 117, 64, 10, 9, 49, 50, 51, 52, 53, 54, 55, 56, 57}

	ossCtxt := codec.NewOssCtxt()
	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	cmnNgAp := codec.NewNgapCodec()
	cmnNgAp.SetEncBuffer(&encMsg[0], len(encMsg))

	cmnNgAp.DecodeHeader(ossCtxtPtr)

	fmt.Println("Get the header of incoming ngap Message: ", cmnNgAp.GetMsgType(), cmnNgAp.GetProcedureCode())

	msg := NewUeRadioCapInfoIndMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	decode := msg.Decode(encMsg)
	fmt.Println("result", decode)
	fmt.Printf("decode Msg: %s", msg)
}
