package sidf

import (
	"fmt"
)

// Subscription Identifier De-concealing Function
// 3GPP TS 33.501 V15.5.0 (2019-06)

//The SIDF is responsible for de-concealment of the SUCI and shall fulfil the following requirements:
//-	The SIDF shall be a service offered by UDM.
//-	The SIDF shall resolve the SUPI from the SUCI based on the protection scheme used to generate the SUCI.

var (
	ErrInvalidScheme   = fmt.Errorf("sidf: invalid protection Scheme identifier")
	ErrInvalidParams   = fmt.Errorf("sidf: invalid parameters")
	ErrInvalidlength   = fmt.Errorf("sidf: invalid byte length")
	ErrInvalidPubKeyId = fmt.Errorf("sidf: invalid home network public key identifier")
)

//SliceToArray 32 byte 到32 byte
func SliceToArrayByte(dst *[32]byte, src []byte) error {

	if len(src) != len(dst) {
		return ErrInvalidlength
	}

	for i := range dst {
		dst[i] = src[i]
	}
	return nil
}
