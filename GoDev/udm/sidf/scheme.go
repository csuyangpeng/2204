package sidf

import (
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/udm/sidf/ecies"
)

//the protection scheme used to generate the SUCI.

// 当Protection Scheme Identifier是ProfileA = 0x1时，
// input :Scheme Output：ECC ephemeral public key(32 bytes) + ciphertext value(5 bytes) + MAC tag value(8 bytes)
// return supi
func SuciDecryptA(UEephPubKey []byte, ciphertext []byte, MacTag []byte,
	HomeNwPubKeyId uint8) ([]byte, error) {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	if len(UEephPubKey) != 32 || len(MacTag) != 8 {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "invalid byte length,UEephPublicKey(%d),MacTag(%d)",
			len(UEephPubKey), len(MacTag))
		return nil, ErrInvalidlength
	}
	if len(ciphertext) > (ecies.MaxSchemeOutput - len(UEephPubKey) - len(MacTag)) {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "invalid byte length,ciphertext(%d)",
			len(ciphertext))
		return nil, ErrInvalidlength
	}
	// Filling in parameters
	decode := ecies.NewHnEciesA()
	err := SliceToArrayByte(&decode.UePubKey, UEephPubKey)
	if err != nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "Failed to get UE public key:%s", err)
		return nil, err
	}
	//fmt.Printf("UePubKey:%x\n", decode.UePubKey)
	rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil, "Ue public key:%#x", decode.UePubKey)

	decode.Ciphertext = ciphertext
	rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil, "Ciphertext:%#x", decode.Ciphertext)

	decode.MacTag = MacTag
	// 根据收到的公钥ID得到对应的私钥
	HnPrivateKey, err := GetHNwPrivateKey(HomeNwPubKeyId)
	if err != nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil,
			"Failed to get home network public key for home network public key Id:%s", err)
		return nil, err
	}

	err = SliceToArrayByte(&decode.HnPriKey, HnPrivateKey) //default:ecies.HnPrivteKey
	if err != nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "Failed to get home network private key:%s", err)
		return nil, err
	}
	//fmt.Printf("HnPriKey,%d:%x\n", HomeNwPubKeyId, decode.HnPriKey)
	rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil, "Home Network private key,%d:%#x", HomeNwPubKeyId, decode.HnPriKey)

	//resolve the SUPI from the SUCI
	// 1 key agreement
	err = decode.KeyAgreement()
	if err != nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "Failed to generate shared key:%s", err)
		return nil, err
	}

	// 2 key derivation
	err = decode.KeyDerivation()
	if err != nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "Key derivation failed:%s", err)
		return nil, err
	}

	// 3 symmetric decryption
	supi, err := decode.SymDecrytion()
	if err != nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "Symmetric decryption failed:%s", err)
		return nil, err
	}

	// 4 MAC function verify
	ok := decode.KeyMacVerify()
	if !ok {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "User message verification failed:%s", err)
		return nil, ecies.ErrMacTagVerify
	}
	rlogger.Trace(types.ModuleUdm, rlogger.INFO, nil, "Suci to de-conceal and verify success,supi:%#x", supi)
	// The variable supi is MSIN
	return supi, nil
}
