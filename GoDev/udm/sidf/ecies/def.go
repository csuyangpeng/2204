package ecies

import "fmt"

// 3GPP TS 33.501 V15.5.0 (2019-06)
// C.3	Elliptic Curve Integrated Encryption Scheme (ECIES)
var (
	ErrImport                     = fmt.Errorf("ecies: failed to import key")
	ErrInvalidCurve               = fmt.Errorf("ecies: invalid elliptic curve")
	ErrInvalidParams              = fmt.Errorf("ecies: invalid ECIES parameters")
	ErrInvalidPublicKey           = fmt.Errorf("ecies: invalid public key")
	ErrSharedKeyIsPointAtInfinity = fmt.Errorf("ecies: shared key is point at infinity")
	ErrSharedKeyTooBig            = fmt.Errorf("ecies: shared key params are too big")
	ErrPublicKeyVerify            = fmt.Errorf("ecies: public key verification error")

	ErrDerivationFailed = fmt.Errorf("ecies: failed to generate derivation key")

	ErrGeneratePublicKeyFailed     = fmt.Errorf("ecies: failed to generate public key")
	ErrGeneratePrivatelicKeyFailed = fmt.Errorf("ecies: failed to generate private key")

	ErrDecryptionBlock = fmt.Errorf("ecies: failed to create new cipher")
	ErrMacTagVerify    = fmt.Errorf("ecies: MAC Tag verification error")
)

//C.1	Introduction
const (
	NullScheme uint8 = 0x0
	ProfileA   uint8 = 0x1
	ProfileB   uint8 = 0x2
)

const MaxSchemeOutput int = 3000

//The UE shall not send, and the network may reject SUCIs larger than the maximum size of scheme-output.
func CheckSchemeOutputSize(suciSo []byte) bool {

	if len(suciSo) > MaxSchemeOutput {
		return false
	}
	return true
}

// C.3.2	Processing on UE side
// The final output shall be the concatenation of the ECC ephemeral public key,
// the ciphertext value, the MAC tag value, and any other parameters, if applicable.

//3GPP TS 23.003 V15.6.0 (2018-12)
// Figure 2.2B-3: Scheme Output for Elliptic Curve Integrated Encryption Scheme Profile A
type UeSchemeOutputA struct {
	// The ECC ephemeral public key is formatted as 64 hexadecimal digits, which allows to encode 256 bits.
	EphPubKey []byte
	// The ciphertext value is formatted as a variable length of hexadecimal digits.(5个16进制数字)
	Ciphertext []byte
	// The MAC tag value is formatted as 16 hexadecimal digits, which allows to encode 64 bits.
	MacTag []byte
}

// Figure 2.2B-4: Scheme Output for Elliptic Curve Integrated Encryption Scheme Profile B
type UeSchemeOutputB struct {
	// The ECC ephemeral public key is formatted as 66 hexadecimal digits, which allows to encode 264 bits.
	EphPubKey []byte
	// The ciphertext value is formatted as a variable length of hexadecimal digits.
	Ciphertext []byte
	// The MAC tag value is formatted as 16 hexadecimal digits, which allows to encode 64 bits.
	MacTag []byte
}

//C.3.3	Processing on home network side
// Profile A
// input :the ECC ephemeral public key(256bits) + the ciphertext value(10数字) + the MAC tag value（64bits） + other parameters(0bit)
//
