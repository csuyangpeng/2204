package ngapmsg

import (
	"fmt"
	codec "lite5gc/cmn/message/ngap/ngapcodec"
	T3 "lite5gc/cmn/types3gpp"
	"testing"
)

func TestInitialContextSetupReqMsg_Encode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()

	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewInitialContextSetupReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	msg.AmfUeNgapId = 5
	msg.RanUeNgapId = 6
	msg.OldAmfName = "shenzhen.amf"
	msg.OptFlags.Set(ICSR_AmfName)
	msg.UeAmbr.Uplink = 5000000
	msg.UeAmbr.Downlink = 6000000
	msg.OptFlags.Set(ICSR_Ueambr)
	msg.Guami.PlmnId.SetString("460001")
	msg.Guami.AmfId.SetAmfRegionID(1)
	msg.Guami.AmfId.SetAmfPointer(2)
	msg.Guami.AmfId.SetAmfSetID([]byte{1, 1})

	item := &T3.PduSessResSetupReqItem{}
	item.PduSessionId = 5
	item.IsNasPduPrst = true
	item.NasPdu = "987654321"
	item.PduSessResSetupReqTrans = string([]byte{88, 32, 7, 161, 32, 16, 78, 32, 1, 240, 10, 180, 9, 242, 55, 37, 0, 0, 0, 8, 20, 24, 11, 204, 112, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 9, 192, 4, 127, 255, 255, 253, 82, 10, 225, 192, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 5, 192, 4, 127, 255, 255, 253, 113, 0, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18})
	item.Snssai = &T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddPduSessResSetupReqItem(item)

	snssai := &T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddAllowedNssai(snssai)

	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA0)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA1)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA2)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA3)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA0)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA1)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA2)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA3)

	msg.SecKey = T3.SecurityKey{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	msg.UeRadioCap = []byte{1, 2, 3, 4, 5, 67}
	msg.OptFlags.Set(ICSR_UeRadioCap)
	msg.IndexRFSelPriority = 1
	msg.OptFlags.Set(ICSR_IndexRfSecPriority)
	msg.MaskedImeiSv = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	msg.OptFlags.Set(ICSR_MaskedImeiSv)
	msg.NasPdu = []byte{1, 1, 1, 1, 1, 2, 2, 2, 22, 3, 3, 3, 3, 3, 3}
	msg.OptFlags.Set(ICSR_Nas)

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestInitialContextSetupReqMsg_Decode(t *testing.T) {

}

func TestInitialContextSetupReqMsg01_Encode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()

	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewInitialContextSetupReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	msg.AmfUeNgapId = 1
	msg.RanUeNgapId = 1610612736
	//msg.OldAmfName = "shenzhen.amf"
	msg.OptFlags.Set(ICSR_AmfName)
	msg.UeAmbr.Uplink = 1073741824
	msg.UeAmbr.Downlink = 1073741824
	msg.OptFlags.Set(ICSR_Ueambr)
	msg.Guami.PlmnId.SetString("460001")
	msg.Guami.AmfId.SetAmfRegionID(1)
	msg.Guami.AmfId.SetAmfPointer(2)
	msg.Guami.AmfId.SetAmfSetID([]byte{1, 1})

	item := &T3.PduSessResSetupReqItem{}
	item.PduSessionId = 5
	item.IsNasPduPrst = true
	item.NasPdu = "987654321"
	item.PduSessResSetupReqTrans = string([]byte{88, 32, 7, 161, 32, 16, 78, 32, 1, 240, 10, 180, 9, 242, 55, 37, 0, 0, 0, 8, 20, 24, 11, 204, 112, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 9, 192, 4, 127, 255, 255, 253, 82, 10, 225, 192, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 5, 192, 4, 127, 255, 255, 253, 113, 0, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18})
	item.Snssai = &T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddPduSessResSetupReqItem(item)

	snssai := &T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddAllowedNssai(snssai)

	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA0)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA1)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA2)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA3)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA0)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA1)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA2)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA3)

	msg.SecKey = T3.SecurityKey{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	msg.UeRadioCap = []byte{1, 2, 3, 4, 5, 67}
	msg.OptFlags.Set(ICSR_UeRadioCap)
	msg.IndexRFSelPriority = 1
	msg.OptFlags.Set(ICSR_IndexRfSecPriority)
	msg.MaskedImeiSv = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	msg.OptFlags.Set(ICSR_MaskedImeiSv)
	msg.NasPdu = []byte{126, 2, 2, 162, 51, 240, 1, 126, 0, 78, 80, 2, 32, 0, 38, 2, 0, 0}
	msg.OptFlags.Set(ICSR_Nas)

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}

func TestInitialContextSetupReqMsg3_Encode(t *testing.T) {

	ossCtxt := codec.NewOssCtxt()

	ossCtxtPtr := ossCtxt.GetOssCtxtPtr_m()
	defer codec.DeleteOssCtxt(ossCtxt)

	msg := NewInitialContextSetupReqMsg()
	msg.SetOssCodecCtxt(ossCtxtPtr)
	msg.AmfUeNgapId = 5
	msg.RanUeNgapId = 6
	msg.OldAmfName = "shenzhen.amf"
	msg.OptFlags.Set(ICSR_AmfName)
	msg.UeAmbr.Uplink = 5000000
	msg.UeAmbr.Downlink = 6000000
	msg.OptFlags.Set(ICSR_Ueambr)
	msg.Guami.PlmnId.SetString("460001")
	msg.Guami.AmfId.SetAmfRegionID(1)
	msg.Guami.AmfId.SetAmfPointer(2)
	msg.Guami.AmfId.SetAmfSetID([]byte{1, 1})

	item := &T3.PduSessResSetupReqItem{}
	item.PduSessionId = 5
	item.IsNasPduPrst = true
	item.NasPdu = "987654321"
	item.PduSessResSetupReqTrans = string([]byte{88, 32, 7, 161, 32, 16, 78, 32, 1, 240, 10, 180, 9, 242, 55, 37, 0, 0, 0, 8, 20, 24, 11, 204, 112, 2, 19, 136, 11, 1, 10, 0, 8, 0, 1, 244, 0, 3, 232, 9, 192, 4, 127, 255, 255, 253, 82, 10, 225, 192, 1, 128, 2, 0, 200, 0, 1, 244, 0, 3, 232, 5, 192, 4, 127, 255, 255, 253, 113, 0, 1, 173, 176, 32, 1, 134, 160, 32, 3, 91, 96, 32, 3, 13, 64, 0, 2, 88, 0, 1, 244, 18})
	item.Snssai = &T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddPduSessResSetupReqItem(item)

	snssai := &T3.Snssai{
		Sst:    0x11,
		Sd:     0x112233,
		SdPrst: true,
	}
	msg.AddAllowedNssai(snssai)

	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA0)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA1)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA2)
	msg.UeSecCap.SetNrIntPrctAlgo(T3.NIA3)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA0)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA1)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA2)
	msg.UeSecCap.SetNrEncAlgo(T3.NEA3)

	msg.SecKey = T3.SecurityKey{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	msg.UeRadioCap = []byte{1, 2, 3, 4, 5, 67}
	msg.OptFlags.Set(ICSR_UeRadioCap)
	msg.IndexRFSelPriority = 1
	msg.OptFlags.Set(ICSR_IndexRfSecPriority)
	msg.MaskedImeiSv = [8]byte{1, 2, 3, 4, 5, 6, 7, 8}
	msg.OptFlags.Set(ICSR_MaskedImeiSv)
	//msg.NasPdu = []byte{126,2,2,162,51,240,1}
	msg.NasPdu = []byte{126, 2, 2, 162, 51, 240, 1, 126, 0, 78, 80, 2, 32, 0, 38, 2}

	//msg.NasPdu = []byte{126, 0, 78, 80, 2, 32, 0 ,38 ,2 ,0, 0}
	msg.OptFlags.Set(ICSR_Nas)

	encodeMsg := msg.Encode()
	fmt.Println("Encode Msg: ", encodeMsg)
}
