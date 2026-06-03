package nassecurity

import (
	"encoding/binary"
	"fmt"
	"lite5gc/amf/security/secalgos"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func GenerateIntegrityMac(intAlgo types3gpp.NRIntPrctAlgo,
	knasInt []byte,
	nasCoutner uint32,
	bearerId byte,
	direction byte,
	plainNasMsg []byte) (error, [4]byte) {

	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	counter := make([]byte, 4)
	binary.BigEndian.PutUint32(counter, nasCoutner)

	bitlen := len(plainNasMsg) * 8

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Generate Int Mac, Input knas_int_key(%x),counter(%x),bearerId(%x),direction(%x),bitlen(%d),plainNasMsg(%x)",
		knasInt, counter, bearerId, direction, bitlen, plainNasMsg)

	var mac []byte
	var err error
	switch intAlgo {
	case types3gpp.NIA0, types3gpp.NIA1:
		mac = make([]byte, 4)
	case types3gpp.NIA2:
		mac, err = secalgos.Nia2CMAC(knasInt, counter, bearerId, direction, plainNasMsg, bitlen)
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "MAC (%x)", mac)
	case types3gpp.NIA3:
		mac, err = secalgos.Nia3ZMAC(knasInt, counter, bearerId, direction, plainNasMsg, bitlen)
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "NIA3 MAC (%x)", mac)
	default:
		mac = make([]byte, 4)
	}

	// return mac
	rtMac := [4]byte{}
	if len(mac) != 4 {
		//invalid generate mac
		if intAlgo != types3gpp.NIA0 {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "unsupported int protect algo(%s)", intAlgo)
			err = fmt.Errorf("unsupported int protect algo(%s)", intAlgo)
			return err, rtMac
		}
	} else {
		if err != nil {
			return fmt.Errorf("failed to GenerateIntegrity mac, error(%s)", err), rtMac
		} else {
			// return mac
			for i := 0; i < len(rtMac); i++ {
				rtMac[i] = mac[i]
			}
		}
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "return mac(%x)", rtMac)
	return nil, rtMac
}

func VerifyIntegrity(intAlgo types3gpp.NRIntPrctAlgo,
	knasInt []byte,
	nasCoutner uint32,
	bearerId byte,
	direction byte,
	plainNasMsg []byte,
	mac []byte) bool {

	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	if knasInt == nil {
		// no knasint key found for the ue context, directly pass the int mac check
		return true
	}

	counter := make([]byte, 4)
	binary.BigEndian.PutUint32(counter, nasCoutner)

	bitlen := len(plainNasMsg) * 8

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Generate Int Mac, Input knas_int_key(%x),counter(%x),bearerId(%x),direction(%x),bitlen(%d),plainNasMsg(%x)",
		knasInt, counter, bearerId, direction, bitlen, plainNasMsg)

	var expectMac []byte
	var err error
	switch intAlgo {
	case types3gpp.NIA0:
		return true
	case types3gpp.NIA1:
		expectMac = make([]byte, 4)
	case types3gpp.NIA2:
		expectMac, err = secalgos.Nia2CMAC(knasInt, counter, bearerId, direction, plainNasMsg, bitlen)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Nia2CMAC failed error(%s)", err)
		}
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Expect MAC (%x), Input MAC(%x)", expectMac, mac)
	case types3gpp.NIA3:
		expectMac, err = secalgos.Nia3ZMAC(knasInt, counter, bearerId, direction, plainNasMsg, bitlen)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "NIA3 MAC failed error(%s)", err)
		}
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Expect NIA3 MAC (%x), Input MAC(%x)", expectMac, mac)
	default:
		expectMac = make([]byte, 4)
	}

	// check the mac and expectMac
	if len(mac) == len(expectMac) {
		for i := 0; i < len(mac); i++ {
			if mac[i] != expectMac[i] {
				rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, nil, "Input MAC(%x) mismatch with ExpectMAC(%x)", mac, expectMac)
				return false
			}
		}
	} else {
		return false
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "pass nas message integrity verification")
	return true
}
