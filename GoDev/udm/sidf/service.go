package sidf

import "lite5gc/udm/sidf/ecies"

//-	The SIDF shall be a service offered by UDM.
func SUCIDeconcealing(suci []byte) (supi []byte, err error) {
	// 解析SCUI ：SUPI Type + Home Network Identifier + Routing Indicator +
	// Protection Scheme Identifier + Home Network Public Key Identifier + Scheme Output
	// 得到 Scheme Output

	// 当Protection Scheme Identifier是ProfileA = 0x1时，
	// Scheme Output：ECC ephemeral public key(32 bytes) + ciphertext value(5 bytes) + MAC tag value(8 bytes)
	schemeID := ecies.ProfileA
	switch schemeID {
	case ecies.ProfileA:
		// todo 解析SUCI
		var UEephPubKey, ciphertext, MacTag []byte
		var hnPublicKeyID uint8
		supi, err := SuciDecryptA(UEephPubKey, ciphertext, MacTag, hnPublicKeyID)
		if err != nil {
			return nil, err
		}
		return supi, nil
	default:
		return nil, ErrInvalidScheme
	}

}
