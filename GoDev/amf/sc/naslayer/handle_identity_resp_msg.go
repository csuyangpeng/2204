package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/context/prcdctxt"
	"lite5gc/amf/sc/naslayer/mmsender"
	"lite5gc/amf/sc/statemgr"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/cmn/types/statetypes"
	"lite5gc/cmn/types3gpp"
	"lite5gc/udm/sidf"
)

func (p *NasMgr) HandleIdentityResponse(ctx context.Context, n2connData *gctxt.N2ConnCtxt,
	plainNasMsg *bytes.Reader) error {
	td := []interface{}{ctx, n2connData}
	rlogger.FuncEntry(types.ModuleAmfNas, td)

	p.identityResponse.Reset()
	//decode the message
	err := p.identityResponse.Decode(plainNasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to decode "+
			"identity response message， %s", err)
		return fmt.Errorf("failed to decode "+
			"identity response message， %s", err)
	}

	ueContext, err := gctxt.GetUeContext(gctxt.AmfUeNgApId(n2connData.AmfUeNgapID))
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, td, "failed to find ue context with AmfUeNgapId(%d)", n2connData.AmfUeNgapID)
		return fmt.Errorf("failed to find ue context with AmfUeNgapId(%d)", n2connData.AmfUeNgapID)
	}

	//check amf ue ngap id valid
	if gctxt.AmfUeNgApId(ueContext.GetAmfUeNgapId()) != n2connData.AmfUeNgapID {
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, td, "AmfUENgapId mismatch between incoming message(%d) and UeContext(%d)",
			n2connData.AmfUeNgapID, ueContext.GetAmfUeNgapId())
		return fmt.Errorf("AmfUENgapId mismatch between incoming message and UeContext")
	}
	switch p.identityResponse.MobileIdentity.IdType {
	case nasie.Suci:
		imsiInUe := ueContext.GetImsi()
		if imsiInUe.GetLength() == 0 {
			switch p.identityResponse.MobileIdentity.Suci.GetProtectSchemeId() {
			case types3gpp.NullScheme:
				imsi, err := p.identityResponse.MobileIdentity.Suci.GetImsi()
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to get imsi from suci ", err)
					return fmt.Errorf("failed to get imsi from suci")
				}

				rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, td, "IMSI in SUCI: ", imsi)
				oldueContext := ueContext
				ueContext, err = gctxt.GetUeContext(gctxt.ImsiKey(imsi.GetValue()))
				if err != nil { // 找不到
					rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, td, "don't find ue context by imsi")
					ueContext = oldueContext
					ueContext.SetImsi(imsi)
					// add index(imsi) for amf UeContext,之前的ue context只有AmfUeNgApId RanUeNgApId
					err = gctxt.AddIndexUeContext(gctxt.ImsiKey(imsi.GetValue()), ueContext)
					if err != nil {
						rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "failed to add index(imsi:%s) "+
							"for ue context,error(%s)", imsi, err)
						return fmt.Errorf("failed to add index(imsi:%s) for ue context,error(%s)", imsi, err)
					}
				} else { // 找到了
					rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, td, "find ue context by imsi")
					//更新ue context的AmfUeNgApId RanUeNgApId
					err = gctxt.UpdateUeContext(gctxt.AmfUeNgApId(oldueContext.GetAmfUeNgapId()), ueContext)
					ueContext.SetAmfUeNgapID(oldueContext.GetAmfUeNgapId())
					ueContext.SetRanUeNgapId(oldueContext.GetRanUeNgapId())
					ueContext.SetProcCtxt(oldueContext.GetProcCtxt())
					ueContext.SecurityCtxt.UeSecCapablity = oldueContext.SecurityCtxt.UeSecCapablity
					ueContext.SecurityCtxt.IsSecCtxtEstb = false
				}

			case types3gpp.ProfileA:
				// SUCI Ciphertext de-concealing
				output := p.identityResponse.MobileIdentity.Suci.GetSchemeOutputA()
				HomeNwPubKeyId := p.identityResponse.MobileIdentity.Suci.GetHomeNwPubKeyId()
				MsinHex, err := sidf.SuciDecryptA(output.EphPublicKey,
					output.Ciphertext, output.MacTag, HomeNwPubKeyId)
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to de-conceal SUCI ", err)
					return fmt.Errorf("failed to de-conceal SUCI ", err)

				}
				p.identityResponse.MobileIdentity.Suci.SetMsinHex2SchemeOutput(MsinHex)
				imsi, err := p.identityResponse.MobileIdentity.Suci.GetImsi()
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to get imsi from suci ", err)
					return fmt.Errorf("failed to get imsi from suci")
				}

				rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, td, "IMSI in SUCI after ProfileA: ", imsi)

				ueContext.SetImsi(imsi)
			}
		} else {
			//mac failure 后的id request ,此时ue ctxt里是有imsi的
			rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, td, "there is already has a imsi in ueContext:%s", imsiInUe)
		}
	case nasie.Guti5g:
	case nasie.STmsi5g:
	case nasie.Imei:
	case nasie.ImeiSvi:
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, td, "invalid mobile id type:%s ", p.identityResponse.MobileIdentity.IdType)
		return fmt.Errorf("invalid mobile id type")
	}

	//cancel t3570
	err = gctxt.CancelPrcdTimer(ueContext, ctx, gctxt.T3570)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueContext.GetImsiPtr(),
			"failed to cancel T3570")
		return fmt.Errorf("failed to cancel T3570")
	}
	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, td, "cancel t3570 timer for identity request msg")

	ctx = context.WithValue(ctx, types.UeContextCK, ueContext)
	if ueContext.GetProcCtxt() != nil {
		switch ueContext.GetProcCtxt().(type) {
		case *prcdctxt.RegistrationPrcdCtxt:
			prcdCtxt, ok := ueContext.GetProcCtxt().(*prcdctxt.RegistrationPrcdCtxt)
			if !ok {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueContext.GetImsiPtr(),
					"failed to get RegistrationPrcdCtxt, and :", ueContext.GetProcCtxt())
				return types.ErrFailGetProcedureCtxt
			} else {
				imsiFromMsg, err := p.identityResponse.MobileIdentity.Suci.GetImsi()
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to get imsi from suci ", err)
					return fmt.Errorf("failed to get imsi from suci")
				}

				switch prcdCtxt.GetCurrentState() {
				case statetype.StateRegisterAuthStart:
					//auth failure，原因是MAC failure，core发起identity，identity resp获取suci后：
					if ueContext.GetImsi() != *imsiFromMsg {
						// 若ue ctxt中的imsi与返回的imsi不相等，重新发起鉴权
						err := statemgr.TriggerFsm(ctx,
							statemgr.Register,
							statetype.StateRegisterAuthStart,
							statetype.EventAuthRequest)
						if err != nil {
							rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
							return types.ErrFailTriggerFsm
						}
					} else {
						// 若ue ctxt中的imsi与返回的imsi相等，发出鉴权拒绝
						err = mmsender.SendAuthenticationRejectMsg(ctx)
						if err != nil {
							rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to send auth reject msg")
							return types.ErrFailSendNasMsg
						}
						ueContext.SetProcCtxt(nil)
						return nil
					}
				case statetype.StateRegisterWfIdentityResp:
					//guti 发起register request,找不到ueContext，core发起identity，identity resp获取suci后，发起鉴权
					err := statemgr.TriggerFsm(ctx,
						statemgr.Register,
						statetype.StateRegisterAuthStart,
						statetype.EventAuthRequest)
					if err != nil {
						rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to trigger fsm")
						return types.ErrFailTriggerFsm
					}
				}
			}
		case *prcdctxt.ServiceRequestPrcdCtxt:
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "there is a ServiceRequest procedure")
		case *prcdctxt.DeRegistrationPrcdCtxt:
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "there is a DeRegistration procedure")
		case *prcdctxt.AnReleasePrcdCtxt:
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "there is a AnRelease procedure")
		default:

			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "unsupported procedure context, error(%s), prcd(%s)",
				err, ueContext.GetProcCtxt())
			return fmt.Errorf("unsupported procedure context error(%s)", err)
		}
	} else {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, td, "there is no under going procedure")
		return fmt.Errorf("there is no under going procedure")
	}
	return nil
}
