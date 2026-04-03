package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/mobmgnt/mmutils"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
)

func (p *NasMgr) HandleAuthFailure(ctx context.Context, n2connData *gctxt.N2ConnCtxt,
	plainNasMsg *bytes.Reader) error {
	td := []interface{}{ctx, n2connData}
	rlogger.FuncEntry(types.ModuleAmfNas, td)

	p.authenticationFailure.Reset()

	err := p.authenticationFailure.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to decode "+
			"authentication failure nas message")
		return fmt.Errorf("failed to decode authentication failure message")
	}

	//get ue context with amf ngap id
	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "type assersion failed for ue ctext")
		return types.ErrFailFindUeCtxt
	}

	// stop timer T3560
	err = gctxt.CancelPrcdTimer(ueCtxt, ctx, gctxt.T3560)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to cancel T3560")
		return fmt.Errorf("failed to cancel T3560")
	}

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "authentication failure, cause(%x) in hex", p.authenticationFailure.Cause)

	syncFailureProceedingFlag := ueCtxt.SecurityCtxt.SyncFailureProceeding

	// reset the ueCtxt.SecurityCtxt.SyncFailureProceeding if needed.
	if ueCtxt.SecurityCtxt.SyncFailureProceeding {
		ueCtxt.SecurityCtxt.SyncFailureProceeding = false
	}

	// 33.501 6.1.3.3
	switch p.authenticationFailure.Cause {
	case nas.SynchFailure:
		if syncFailureProceedingFlag {
			// 24.501 5.4.1.3.7 f)
			// NOTE 4:	Upon receipt of two consecutive AUTHENTICATION FAILURE messages from the UE with 5GMM
			// cause #21 "synch failure", the network may terminate the 5G AKA based primary authentication
			// and key agreement procedure by sending an AUTHENTICATION REJECT message.
			// send authentication reject to UE
			rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, td,
				"consecutive authentication failure msg received, send auth reject to ue")
			err = mmsender.SendAuthenticationRejectMsg(ctx)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, td, "fail to send auth reject msg")
				return types.ErrFailSendNasMsg
			}
			return nil
		}
		//retreive the AV with AUTS
		// item f
		reRsyncInfo := &types.AuthReRsyncData{}
		reRsyncInfo.Rand = ueCtxt.SecurityCtxt.TempSecCtxt.AuthVector.Rand
		reRsyncInfo.Auts = p.authenticationFailure.Auts
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "ResyncInfo, rand(%x), auts(%x)", reRsyncInfo.Rand, reRsyncInfo.Auts)

		suci := &types3gpp.Suci{}
		suci.SetFromImsi(ueCtxt.GetSupi().GetImsi())
		err := mmutils.RetrieveAuthVector(ctx, ueCtxt, suci, reRsyncInfo)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to retreive auth vector. error(%s)", err)
			return fmt.Errorf("failed to retreive auth vector. error(%s)", err)
		}
		//send authentication request message to ue
		err = mmsender.SendAuthenticationRequestMsg(ctx)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, td, "fail to send auth request msg")
			return types.ErrFailSendNasMsg
		}
		ueCtxt.SecurityCtxt.SyncFailureProceeding = true
	case nas.NgksiInUsed:
		// 24.501 5.4.1.3.7 item e
		// the network performs necessary actions to select a new ngKSI and
		// send the same 5G authentication challenge to the UE.
		// the network may also re-initiate the 5G AKA based primary authentication and
		// key agreement procedure (see subclause 5.4.1.3.2).
		suci := &types3gpp.Suci{}
		suci.SetFromImsi(ueCtxt.GetSupi().GetImsi())
		err := mmutils.RetrieveAuthVector(ctx, ueCtxt, suci, nil)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to retreive auth vector. error(%s)", err)
			return fmt.Errorf("failed to retreive auth vector. error(%s)", err)
		}
		ueCtxt.SecurityCtxt.NgKsi.Update()
		//send authentication request message to ue
		err = mmsender.SendAuthenticationRequestMsg(ctx)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, td, "fail to send auth request msg")
			return types.ErrFailSendNasMsg
		}
	case nas.Non5gAuthUnaccept:
		// 24.501 5.4.1.3.7, item d.
		// Upon the first receipt of an AUTHENTICATION FAILURE message from the UE with 5GMM cause
		// #26 "non-5G authentication unacceptable" and #20 "MAC failure",

		// the network may initiate the identification procedure described in subclause 5.4.3.
		//If the mapping of 5G-GUTI to SUPI in the network was incorrect,
		// the network should respond by sending a new AUTHENTICATION REQUEST message to the UE.

		// If the mapping of 5G-GUTI to SUPI in the network was correct,
		// the network should terminate the 5G AKA based primary authentication and key agreement procedure
		// by sending an AUTHENTICATION REJECT message (see subclause 5.4.1.3.5) and 5.4.3.1.7 c)
		fallthrough
	case nas.MacFailure:
		// item c
		err = mmsender.SendIdentityRequest(ctx, ueCtxt, nasie.Suci)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to send identity response msg")
		}
		ueRegistPrcdCtxt, ok := ueCtxt.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
		if !ok {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to get RegistrationPrcdCtxt, and :", ueCtxt.GetProcCtxt())
			return types.ErrFailGetProcedureCtxt
		} else {
			err = ueRegistPrcdCtxt.SetNextState(statetype.StateRegisterAuthStart) //与guti注册时返回的identity resp区分开来
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, td, "fail to set state")
				return fmt.Errorf("fail to set state")
			}
		}
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "authentication failure, cause(%x) in hex", p.authenticationFailure.Cause)
	}

	return nil
}
