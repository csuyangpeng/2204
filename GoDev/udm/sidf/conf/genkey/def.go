package genkey

// TS23.003
// Figure 2.2B-3:
// Scheme Output for Elliptic Curve Integrated Encryption Scheme Profile A
// key pair
type HomeNetworkPKI struct {
	// Home network public key identifier (PKI)
	// Home network PKI value (1-254)
	HomeNetworkPKI uint8
	HomeNetworkKeyPair
}

// key pair
type HomeNetworkKeyPair struct {
	// The ECC ephemeral public key is formatted as 64 hexadecimal digits,
	// which allows to encode 256 bits.
	// Home network public key:Hexadecimal value
	HnPublicKey []byte
	//Home Network Private Key:Hexadecimal value
	HnPrivteKey []byte
}

type HomeNetworkPKIStr struct {
	// Home network public key identifier (PKI)
	// Home network PKI value (1-254)
	HomeNetworkPKI uint8 `json:"Home network public key identifier"`
	HomeNetworkKeyPairStr
}
type HomeNetworkKeyPairStr struct {
	// The ECC ephemeral public key is formatted as 64 hexadecimal digits,
	// which allows to encode 256 bits.
	// Home network public key:Hexadecimal value
	HnPublicKey string `json:"Home network public key"`
	//Home Network Private Key:Hexadecimal value
	HnPrivteKey string `json:"Home network private key"`
}

var PKISet []HomeNetworkPKI
