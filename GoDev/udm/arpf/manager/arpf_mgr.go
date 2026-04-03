package manager

import (
	"context"
	"encoding/binary"
	"fmt"
	"lite5gc/cmn/message/udmdata"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/udm/arpf/derivevec"
)

type ArpfMgr struct {
}

// handle ue authentication get request message from ausf
func HandleUeAuthenticatonGetReqMsg(ctxt context.Context,
	suci *types3gpp.Suci,
	snName string,
	reRsyncInfo *types.AuthReRsyncData) (error, *types3gpp.Supi, *types.HeAvType) {

	rlogger.FuncEntry(types.ModuleUdm, suci)

	var err error

	//check suci
	imsi, err := suci.GetImsi()
	if err != nil {
		return fmt.Errorf("failed to get imsi"), nil, nil
	}

	imsiStr := imsi.String()
	var ueSecCtxt *types.UeSecContext

	supi := &types3gpp.Supi{}
	err = supi.SetType(types3gpp.IMSIType)
	if err != nil {
		return fmt.Errorf("wrong supi type"), nil, nil
	}
	supi.SetImsi(imsi)

	// get ue security context
	ueSecCtxt, err = GetUeSecContext(imsiStr)
	rlogger.Trace(types.ModuleUdm, rlogger.INFO, supi, "err:", err)
	if err != nil {
		// no ue security context, create a new ue security context
		err, ueSecCtxt = CreateUeSecContext(imsiStr)
		if err != nil {
			return fmt.Errorf("failed to create ue security context, error(%s_", err), nil, nil
		}

		// get auth data from database
		//udmAgent, ok := ctxt.Value(types.UdmAgentCK).(*udmAgentLayer.LayerMgr)
		//if !ok {
		//	rlogger.Trace(types.ModuleUdm, rlogger.ERROR, supi, "failed to get udm agent")
		//	return fmt.Errorf("failed to get UdmAgent"), nil, nil
		//}

		//authData, err := udmAgent.Egress.GetAuthData(supi)
		//if err != nil {
		//	rlogger.Trace(types.ModuleUdm, rlogger.ERROR, supi, "failed to get AuthData from db")
		//	return fmt.Errorf("failed to get AuthData from DB"), nil, nil
		//}
		authData := udmdata.AuthData{}

		// store the Auth data from DB
		ueSecCtxt.Key = authData.Ki
		ueSecCtxt.IsOpc = authData.IsOpc
		ueSecCtxt.Op = authData.Op
		ueSecCtxt.Opc = authData.Opc
		ueSecCtxt.Amf = authData.Amf
		ueSecCtxt.Sqn = [6]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00} //sqn init

		rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, supi, "SN Name (%s)", snName)
		ueSecCtxt.SnName = snName
	}

	//generate HeAV
	RetreiveSqn(ueSecCtxt, reRsyncInfo)

	heAv, err := derivevec.DeriveHeAv(ueSecCtxt)
	if err != nil {
		return fmt.Errorf("failed to derive HeAv, error(%s)", err), nil, nil
	}

	return nil, supi, heAv
}

func RetreiveSqn(ueSecCtxt *types.UeSecContext, reRsyncInfo *types.AuthReRsyncData) {
	rlogger.FuncEntry(types.ModuleUdm, ueSecCtxt)

	if ueSecCtxt == nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, nil, "Critical ERROR, Invalid UeSecCtxt")
		return
	}

	// normal procedure, just increase sqn number
	if reRsyncInfo == nil {
		ueSecCtxt.Sqn = IncreaseSQN(ueSecCtxt.Sqn)
		return
	}

	// need resync sqn
	ak2, err := derivevec.ComputeAK2(ueSecCtxt, reRsyncInfo.Rand)
	if err != nil {
		rlogger.Trace(types.ModuleUdm, rlogger.ERROR, ueSecCtxt, "failed to derive AK2, error(%s)", err)
		return
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, ueSecCtxt, "With rand, derive ak2 (%x)", ak2)

	var sqnms_xor_ak2 [6]byte
	for i := 0; i < 6; i++ {
		sqnms_xor_ak2[i] = reRsyncInfo.Auts[i]
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, ueSecCtxt, "Auts, sqnms_xor_ak (%x)", sqnms_xor_ak2)

	// get SQNms
	var SQNms [6]byte
	for i := 0; i < 6; i++ {
		SQNms[i] = sqnms_xor_ak2[i] ^ ak2[i]
	}
	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, ueSecCtxt, "Auts, Sqnms (%x)", SQNms)

	if SQNinRange(SQNms, ueSecCtxt.Sqn, 5, 268435456, 32) == false {
		ueSecCtxt.Sqn = SQNms
	}

	ueSecCtxt.Sqn = IncreaseSQN(ueSecCtxt.Sqn)
}

//func StoreSqn(sqn [6]byte) uint64 {
//	rlogger.FuncEntry()
//
//	sqnArr := make([]byte, 8)
//	sqnArr[0] = 0
//	sqnArr[1] = 0
//	for i := 0; i < 6; i++ {
//		sqnArr[i+2] = sqn[i]
//	}
//
//	return binary.BigEndian.Uint64(sqnArr)
//}

//func IncreaseSqnUint64(sqn uint64) uint64 {
//	rlogger.FuncEntry()
//
//	sqnArr := make([]byte, 8)
//	binary.BigEndian.PutUint64(sqnArr, sqn)
//	var sqnInput [6]byte
//	for i := 0; i < 6; i++ {
//		sqnInput[i] = sqnArr[i+2]
//	}
//
//	sqnOutput := IncreaseSQN(sqnInput)
//
//	sqnArr[0] = 0
//	sqnArr[1] = 0
//	for i := 0; i < 6; i++ {
//		sqnArr[i+2] = sqnOutput[i]
//	}
//
//	return binary.BigEndian.Uint64(sqnArr)
//}

// 33.102 C.3.2 not time-based.
func IncreaseSQN(sqn [6]byte) [6]byte {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	//SQN = SEQ | IND (5 bit)
	ind := sqn[5] & 0x1F
	ind++
	if ind&0x1F == 0x00 {
		ind = 0x00
	}

	var seq [6]byte
	seq = sqn

	seq[5] = seq[5] & 0xE0 // 1110 0000
	seq[5] = seq[5] + 0x20 // seq++
	if seq[5] == 0x00 {
		seq[4]++
		if seq[4] == 0x00 {
			seq[3]++
			if seq[3] == 0x00 {
				seq[2]++
				if seq[2] == 0x00 {
					seq[1]++
					if seq[1] == 0x00 {
						seq[0]++
						if seq[0] == 0x00 {
							seq[5] = 0x20
						}
					}
				}
			}
		}
	}

	seq[5] = seq[5] | byte(ind)

	rlogger.Trace(types.ModuleUdm, rlogger.DEBUG, nil, "input SQN: %x, output SQN %x", sqn, seq)
	return seq
}

func SQNinRange(sqnMs [6]byte, sqnHe [6]byte, ind_len uint, delta int64, L int64) bool {
	rlogger.FuncEntry(types.ModuleUdm, nil)

	sqnMsBytes := make([]byte, 8)
	for i := 0; i < 6; i++ {
		sqnMsBytes[i+2] = sqnMs[i]
	}

	sqnHeBytes := make([]byte, 8)
	for i := 0; i < 6; i++ {
		sqnHeBytes[i+2] = sqnHe[i]
	}

	sqnMsU64 := binary.BigEndian.Uint64(sqnMsBytes)
	sqnHeU64 := binary.BigEndian.Uint64(sqnHeBytes)

	seqMS := sqnMsU64 >> ind_len
	seqHE := sqnHeU64 >> ind_len

	if int64(seqHE)-int64(seqMS) > delta {
		return false
	}

	if int64(seqMS)-int64(seqHE) > L {
		return false
	}

	if int64(seqHE) <= int64(seqMS) {
		return false
	}

	return true
}
