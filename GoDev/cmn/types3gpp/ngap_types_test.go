/** Copyright(C),2020-2022
* Author: zmj
* Date: 12/7/20 5:45 PM
* Description:
 */
package types3gpp

import (
	"encoding/binary"
	"lite5gc/cmn/types"
	"testing"
)

func TestConvertSdToU32_LittleEndian(t *testing.T) {
	var tmp [SizeofSD]byte
	tmp[0] = 1
	tmp[1] = 1
	tmp[2] = 1

	out := ConvertSdToU32(tmp[:], types.LittleEndian)

	var expOut uint32 = 0x10101

	if out != expOut {
		t.Errorf("exp: %d, but return :%d", expOut, out)
	}
}

func TestConvertSdToU32_BigEndian(t *testing.T) {
	var tmp [SizeofSD]byte
	tmp[0] = 1
	tmp[1] = 1
	tmp[2] = 1

	out := ConvertSdToU32(tmp[:], types.BigEndian)

	var expTmp [4]byte
	expTmp[0]=1
	expTmp[1]=1
	expTmp[2]=1
	expTmp[3]=0
	expOut := binary.BigEndian.Uint32(expTmp[:])

	if out != expOut {
		t.Errorf("exp: %d, but return :%d", expOut, out)
	}
}

//func TestConvertU32ToSd(t *testing.T) {
//	ConvertU32ToSd()
//}
