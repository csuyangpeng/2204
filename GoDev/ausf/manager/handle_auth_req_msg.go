package manager

import (
	"context"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types3gpp"
	"lite5gc/udm/arpf/derivevec"
	arpfmgr "lite5gc/udm/arpf/manager"
)

type AusfMgr struct {
}

//func (p *AusfMgr) HandleUeAuthenticationReqMsg(ctxt context.Context,
func HandleUeAuthenticationReqMsg(ctxt context.Context,
	suci *types3gpp.Suci,
	snName string,
	reRsyncInfo *types.AuthReRsyncData) (error, *types3gpp.Supi, *types.SeAvType) {
	rlogger.FuncEntry(types.ModuleAusf, suci)

	//check suci
	imsi, err := suci.GetImsi()
	if err == nil {
		//get imsi from suci
		imsiStr := imsi.String()

		// get HeAv from udm-arpf
		// get ue security context
		ueAvCtxt, err := GetAusfAvContext(imsiStr)
		if err != nil {
			// no ue security context, create a new ue security context
			err, ueAvCtxt = CreateAusfAvContext(imsiStr)
			if err != nil {
				return fmt.Errorf("failed to create ue security context, error(%s_", err), nil, nil
			}
		}

		// get auth data from database
		// udm sbi start
		err, uesupi, heav := arpfmgr.HandleUeAuthenticatonGetReqMsg(ctxt, suci, snName, reRsyncInfo)

		if err != nil {
			rlogger.Trace(types.ModuleAusf, rlogger.ERROR, suci, "failed to get heav from arpf")
			return fmt.Errorf("failed to get heav from arpf"), nil, nil
		}

		//derive ausf av
		err = derivevec.DeriveAusfAv(heav, snName, ueAvCtxt)
		if err != nil {
			rlogger.Trace(types.ModuleAusf, rlogger.ERROR, suci, "failed to derive ausf av, error(%s)", err)
			return fmt.Errorf("failed to derive ausf av, error(%s)", err), nil, nil
		}

		// store the Auth data from ARPF
		seav := &types.SeAvType{}
		seav.Autn = ueAvCtxt.Autn
		seav.HXresStar = ueAvCtxt.HXresStar
		seav.Rand = ueAvCtxt.Rand
		return nil, uesupi, seav

	} else {
		// suci

	}

	return fmt.Errorf("failed to handle ue auth req msg in AUSF"), nil, nil
}
