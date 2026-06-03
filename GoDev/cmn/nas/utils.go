package nas

import (
	"fmt"
	"lite5gc/cmn/types3gpp"
	"strconv"
	"strings"
)

func GetSmfKeys(smCtxtRef string) (types3gpp.Imsi, PduSessID, error) {

	imsi := types3gpp.Imsi{}
	var psi PduSessID

	infos := strings.Split(smCtxtRef, "-")
	if len(infos) != 3 {
		return imsi, psi, fmt.Errorf("invalid smCtxtRef(%s)", smCtxtRef)
	}

	imsiStr := infos[1]
	psiStr := infos[2]

	_ = imsi.StoreImsiString(imsiStr, types3gpp.CheckMncLen(imsiStr))
	psiInt, err := strconv.Atoi(psiStr)
	if err != nil {
		return imsi, psi, fmt.Errorf("invalid psi(%s) in smCtxtRef(%s)", psiStr, smCtxtRef)
	}

	psi = PduSessID(psiInt)

	return imsi, psi, nil
}
