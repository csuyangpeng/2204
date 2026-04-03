package nassecurity

import (
	"encoding/binary"
	"fmt"
	"lite5gc/amf/security/secalgos"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func CipherNasMsg(encAlgo types3gpp.NREncryptAlgo, encKey []byte, nasCount uint32, bearId byte, direction byte, plainNasMsg []byte) ([]byte, error) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	counter := make([]byte, 4)
	binary.BigEndian.PutUint32(counter, nasCount)

	bitlen := len(plainNasMsg) * 8

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "cipher nas msg, enc_algo(%d),knas_enc_key(%x),counter(%x),bearerId(%x),direction(%x),bitlen(%d),plainNasMsg(%x)",
		encAlgo, encKey, counter, bearId, direction, bitlen, plainNasMsg)

	switch encAlgo {
	case types3gpp.NEA0:
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "NEA0 enc algos.")
		return plainNasMsg, nil
	case types3gpp.NEA1:
	case types3gpp.NEA2:
		cipherNasMsg, err := secalgos.Nea2Encrypt(encKey, counter, bearId, direction, plainNasMsg, bitlen)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to encode ciper nas message. error(%s)", err)
			return nil, err
		}
		return cipherNasMsg, nil
	case types3gpp.NEA3:
		cipherNasMsg, err := secalgos.Nea3Encrypt(encKey, counter, bearId, direction, plainNasMsg, bitlen)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to encode ciper nas message for nea3. error(%s)", err)
			return nil, err
		}
		return cipherNasMsg, nil
	default:
	}

	return nil, fmt.Errorf("failed to ciper nas message")

}

func DecipherNasMsg(encAlgo types3gpp.NREncryptAlgo, encKey []byte, nasCount uint32, bearId byte, direction byte, cipherMsg []byte) ([]byte, error) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	counter := make([]byte, 4)
	binary.BigEndian.PutUint32(counter, nasCount)

	bitlen := len(cipherMsg) * 8

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "decipher nas msg, enc_algo(%d),knas_enc_key(%x),counter(%x),bearerId(%x),direction(%x),bitlen(%d),cipherNasMsg(%x)",
		encAlgo, encKey, counter, bearId, direction, bitlen, cipherMsg)

	switch encAlgo {
	case types3gpp.NEA0:
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "using NEA0")
		return cipherMsg, nil
	case types3gpp.NEA1:
	case types3gpp.NEA2:
		plainNasMsg, err := secalgos.Nea2Decrypt(encKey, counter, bearId, direction, cipherMsg, bitlen)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to decode ciper nas message. error(%s)", err)
			return nil, err
		}
		rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "---NEA2 decode ciper nas message：%x", plainNasMsg)
		return plainNasMsg, nil
	case types3gpp.NEA3:
		plainNasMsg, err := secalgos.Nea3Decrypt(encKey, counter, bearId, direction, cipherMsg, bitlen)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to decode ciper nas message for nea3. error(%s)", err)
			return nil, err
		}
		rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "---NEA3 decode ciper nas message：%x", plainNasMsg)
		return plainNasMsg, nil
	default:

	}

	return nil, fmt.Errorf("failed to deciper nas message")
}
