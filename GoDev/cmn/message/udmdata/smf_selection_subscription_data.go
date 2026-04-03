package udmdata

import (
	"encoding/hex"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/types3gpp"
	"strconv"
)

type SmfSelSubscribeData struct {
	SupFeatures    uint64
	SnssaiInfoList map[string]SnssaiInfo
}

func (p *SmfSelSubscribeData) GetDefDnn(snssai *nasie.SNssai) (apn *types3gpp.Apn, err error) {
	var snssaikey string
	switch hex.EncodeToString(snssai.Sd[:]) {
	case "":
		snssaikey = strconv.Itoa(int(snssai.Sst))
	case "10000":
		snssaikey = strconv.Itoa(int(snssai.Sst)) + "-1"
	default:
		snssaikey = strconv.Itoa(int(snssai.Sst)) + "-" + hex.EncodeToString(snssai.Sd[:])
	}
	apn = &p.SnssaiInfoList[snssaikey].DnnInfos[0].Dnn
	return apn, nil
}
