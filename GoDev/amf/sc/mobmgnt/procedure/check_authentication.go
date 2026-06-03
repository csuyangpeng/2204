/*
* Copyright(C),2020-2022
* Author:  xiaoyun
* Date:    11/13/20 3:29 AM
* Description:
	判断是否需要进行鉴权流程
*/
package procedure

import (
	"lite5gc/amf/context/gctxt"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
)

func CheckAuthentication(secTxt gctxt.SecurityCtxt,
	ksi nasie.NasKSI,
	ueSecCap types3gpp.SecurityCapability) bool {

	rlogger.FuncEntry(types.ModuleAmfMM, nil)

	authNeeded := false

	// 2. if ksi in request msg is invalid, true
	if ksi.IsValid() == false {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Invalid KSI(%s)", ksi)
		authNeeded = true
	}

	// 3. if ksi in request msg is different with ksi stored in AMF, true
	if secTxt.NgKsi.Ksi != ksi.Ksi {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil,
			"ngKSI mismatched.(%s) - (%s)", secTxt.NgKsi, ksi)
		authNeeded = true
	}

	// 4. if algs in UE Net don't match the algs stored in ue context,trigger authentication
	if !secTxt.UeSecCapablity.IsMatched(ueSecCap) {
		authNeeded = true
	}

	// 5. Check the force authentication flag
	if secTxt.ForceAuthNeed {
		authNeeded = true
	}

	return authNeeded
}
