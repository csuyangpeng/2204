package mmutils

import (
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func GetPsiIdx(psi byte, numPsi uint8, idxSlice *[]byte) {
	var i uint8
	for i = 0; i < 8; i++ {
		var mask uint8 = 1
		mask = 1 << i
		if (psi & mask) != 0 {
			*idxSlice = append(*idxSlice, i+numPsi*8)
		}
	}
}

//求并集
func Union(slice1, slice2 []byte) []byte {
	m := make(map[byte]int)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

func GenerateSnName(plmn types3gpp.PlmnID) string {
	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	mncStr := plmn.GetMncString()
	mccStr := plmn.GetMccString()

	if len(mncStr) == 2 {
		mncStr = "mnc0" + mncStr
	} else {
		mncStr = "mnc" + mncStr
	}

	sn := fmt.Sprintf("5G:%s.mcc%s.3gppnetwork.org", mncStr, mccStr)
	rlogger.Trace(types.ModuleAmfMM, rlogger.DEBUG, nil,
		"generate sn_name(%s)", sn)
	return sn
	//return "5G:" + mncStr + ".mcc" + mccStr + ".3gppnetwork.org"
}
