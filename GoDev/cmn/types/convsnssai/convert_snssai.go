/** Copyright(C),2020-2022
* Author: zmj
* Date: 12/8/20 9:55 AM
* Description:
 */
package convsnssai

import (
	"encoding/binary"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/types3gpp"
)

func Convert3gtpSnssai(s nasie.SNssai) types3gpp.Snssai {

	out := types3gpp.Snssai{}

	out.Sst = s.Sst
	if s.Ind == nasie.SstSd {
		out.Sd = binary.LittleEndian.Uint32(s.Sd[:])
		out.SdPrst = true
	}

	return out
}

func ConvertSnssai(s types3gpp.Snssai) nasie.SNssai {
	out := nasie.SNssai{}

	out.Sst = s.Sst

	if s.SdPrst == true {
		out.Ind = nasie.SstSd
		tmp := make([]uint8, 4)
		binary.BigEndian.PutUint32(tmp, s.Sd)

		for i := 0; i < nasie.SizeofSd; i++ {
			out.Sd[i] = tmp[i]
		}
	} else {
		out.Ind = nasie.SstOnly
	}
	return out
}
