package naslayer

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/statistics"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"lite5gc/oam/pm"
)

func (p *NasMgr) HandleSecMmMessage(ctx context.Context, n2connData *gctxt.N2ConnCtxt,
	secHeader nas.SecHeaderType, msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	// MAC
	mac := make([]byte, 4)
	err := binary.Read(msgBuf, binary.BigEndian, &mac)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to read MAC")
		return fmt.Errorf("fail to read MAC")
	}

	// Uplink NAS Counter
	nasSeqNum, err := msgBuf.ReadByte()
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to read nas counter")
		return fmt.Errorf("fail to read nas counter")
	}

	// NasMessage
	msgLen := msgBuf.Len()
	nasMsg := make([]byte, msgLen)
	err = binary.Read(msgBuf, binary.BigEndian, &nasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to read sec nas message")
		return fmt.Errorf("fail to read sec nas message")
	}

	ueCtxt, ok := ctx.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "fail to get ue context")
		return types.ErrFailFindUeCtxt
	}

	var intPrctResult IntegrationStatusType

	if secHeader == nas.IntegrityPrtctCipher || secHeader == nas.IntegrityPrtctCipherNewSecCtxt {
		//nasMsg可能是密文，第三个字节不能直接用
		// MAC verification
		if ok {
			intPrctResult = ValidateProtectedMsg(ueCtxt, nasSeqNum, mac, nasMsg)
		}
	} else {
		//明文
		nasMsgType := nasMsg[2] // the third byte in Plain NAS message is message type.
		// find ue context
		if !ok {
			rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, nil, "failed to get ue context")
			// for initial ue message, no ue context founded.
			//return fmt.Errorf("failed to get UE Context")
			//TODO verifictaion for mac, pass currently
			//intPrctResult = IntegrityValidated
			if nasMsgType == byte(nas.ServiceRequest) {
				rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "coming a ServiceRequest msg")
				//5G-S-TMSI
				var mobileIdentity nasie.MobileIdentity
				stmsiByte := bytes.NewReader(nasMsg[4:13])
				rlogger.Trace(types.ModuleAmfNas, rlogger.INFO, nil, "stmsiByte", len(nasMsg[4:13]))

				err = mobileIdentity.Decode(stmsiByte)
				if err != nil {
					return nas.ErrDecodeNasIeFailed
				}

				stmsiKey := mobileIdentity.Stmsi5g.GetKey()
				ueCtxt, err = gctxt.GetUeContext(gctxt.StmsiKey(stmsiKey))
				if err != nil {
					rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil,
						"failed to find the ue context for 5g-stmsi :%s", stmsiKey)
					return fmt.Errorf("failed to find the ue context for Stmsi5g :%s", stmsiKey)
				}
			}
		}
		// do mac validation
		if ueCtxt != nil {
			switch nasMsgType {
			//24.501 4.4.4.3
			//4.4.4.3	Integrity checking of NAS signalling messages in the AMF
			//Except the messages listed below,
			//no NAS signalling messages shall be processed by the receiving 5GMM entity in the AMF or
			//forwarded to the 5GSM entity, unless the secure exchange of NAS messages has been established
			//for the NAS signalling connection:
			//a)	REGISTRATION REQUEST;
			//b)	IDENTITY RESPONSE (if requested identification parameter is SUCI);
			//c)	AUTHENTICATION RESPONSE;
			//d)	AUTHENTICATION FAILURE;
			//e)	SECURITY MODE REJECT;
			//f)	DEREGISTRATION REQUEST; and
			//g)	DEREGISTRATION ACCEPT;
			case byte(nas.RegistrationRequest), byte(nas.IdentifyResponse), byte(nas.AuthenticationResponse),
				byte(nas.AuthenticationFailure), byte(nas.SecurityModeReject), byte(nas.DeregistrationRequestUe),
				byte(nas.DeregistrationAcceptUe):
				intPrctResult = IntegrityValidated
			default:
				// MAC verification
				intPrctResult = ValidateProtectedMsg(ueCtxt, nasSeqNum, mac, nasMsg)
			}
		} else {
			intPrctResult = IntegrityValidated
		}
	}

	if intPrctResult == IntegrityFailed {
		// integrity check failed
		// TODO send failure to UE ?
		rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, nil, "integration validation failed.")
		return fmt.Errorf("integration validation failed. ")
	} else {
		// integrity check passed
		if secHeader == nas.IntegrityPrtctCipher || secHeader == nas.IntegrityPrtctCipherNewSecCtxt {
			plainNasMsg, err := DecipherMessage(ueCtxt, nasMsg)
			if err != nil {
				// decipher nas message failed, drop message
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "decipher nas message failed.")
				return fmt.Errorf("decipher nas message failed")
			}

			//handle the decipher message
			if p.ProcessAndHandleProtectedMessage(ctx, n2connData, plainNasMsg) != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil,
					"failed to process and handle protected message")
				return fmt.Errorf("failed to process and handle protected message")
			}
		} else if secHeader == nas.IntegrityPrtc || secHeader == nas.IntegrityPrtctNewSecCtxt {
			// no need do the decipher
			if p.ProcessAndHandleProtectedMessage(ctx, n2connData, nasMsg) != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil,
					"failed to process and handle protected message")
				return fmt.Errorf("failed to process and handle protected message")
			}
		}
	}
	return nil
}

func (p *NasMgr) ProcessAndHandleProtectedMessage(ctx context.Context,
	n2connData *gctxt.N2ConnCtxt, plainNasMsg []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "Nas Message: %x", plainNasMsg)

	nasMsg := bytes.NewReader(plainNasMsg)

	// decode extract the EPD
	epd, err := nas.GetEpd(nasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to decode nas epd")
		return fmt.Errorf("failed to decode nas epd")
	}

	switch epd {
	case nas.Epd5gsMobMgntMsg:
		secHeaderType, err := nas.GetSecHeaderType(nasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to decode nas security header type")
			return fmt.Errorf("failed to decode nas security header type")
		}
		switch secHeaderType {
		case nas.PlainNasMsg:
			err = p.HandleMmMessage(ctx, n2connData, nasMsg)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, nil, "fail to hand mm msg")
				return fmt.Errorf("fail to hand mm msg")
			}
		case nas.IntegrityPrtc:
		case nas.IntegrityPrtctNewSecCtxt:
		case nas.IntegrityPrtctCipherNewSecCtxt:
		case nas.IntegrityPrtctCipher:
		default:
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "unsupported header (%x)", secHeaderType)
		}

	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "unsupported epd (%x)", epd)
	}

	return nil
}

func (p *NasMgr) HandleMmMessage(ctx context.Context, n2connData *gctxt.N2ConnCtxt,
	plainNasMsg *bytes.Reader) error {

	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	//3.Decode Message Type identity
	msgType, err := plainNasMsg.ReadByte()
	if err != nil {
		return nas.ErrInvalidMmMsgType
	}

	switch nas.MmMsgType(msgType) {
	case nas.RegistrationRequest:
		//msg counter
		//pm.PegCounter(statistics.RegistrationRequestCounter)
		cause := p.HandleRegistrationRequestMsg(ctx, n2connData, plainNasMsg)
		if cause != nas.SuccessAccept {
			err := RegisterFailHandler(ctx, p.ScInstId, n2connData, cause)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to send register reject")
				return fmt.Errorf("failed to send register reject")
			}
		}
	case nas.RegistrationComplete:
		//msg counter
		pm.PegCounter(statistics.RegistrationCompleteCounter)
		err := p.HandleRegistrationCompleteMsg(ctx, plainNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"registration complete message: %s", err)
			return err
		}
	case nas.ULNasTransport:
		//msg counter
		pm.PegCounter(statistics.ULNasTransportCounter)
		err := p.HandleULNasTransportMsg(ctx, n2connData, plainNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"upLink nas transport message: %s", err)
			return err
		}
	case nas.DeregistrationRequestUe:
		//msg counter
		pm.PegCounter(statistics.DeregistrationRequestUeCounter)
		err := p.HandleUeDeRegistReqMsg(ctx, n2connData, plainNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"ue deRegistration request message: %s", err)
			return err
		}
	case nas.DeregistrationAcceptUeT:
		err := p.HandleUeDeRegistAcceptMsg(ctx, n2connData)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"ue deRegistration request message: %s", err)
			return err
		}
	case nas.ServiceRequest:
		//msg counter
		pm.PegCounter(statistics.ServiceRequestCounter)
		//pm.IncCounter(statistics.ServiceRequestCounter, 1)
		err := p.HandleServiceRequestMsg(ctx, n2connData, plainNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"service message: %s", err)
			return err
		}
	case nas.IdentifyResponse:
		//msg counter
		pm.PegCounter(statistics.IdentifyResponseCounter)
		err := p.HandleIdentityResponse(ctx, n2connData, plainNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"identity response message: %s", err)
			return err
		}
	case nas.AuthenticationResponse:
		// msg counter
		pm.PegCounter(statistics.AuthenticationResponseCounter)
		err := p.HandleAuthRespMsg(ctx, plainNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"Authentication response message: %s", err)
			return err
		}
	case nas.SecurityModeComplete:
		//msg counter
		pm.PegCounter(statistics.SecurityModeCompleteCounter)
		err := p.HandleSecModCmpMsg(ctx, n2connData, plainNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"Authentication reject message: %s", err)
			return err
		}
	case nas.AuthenticationFailure:
		//msg counter
		pm.PegCounter(statistics.AuthenticationFailureCounter)
		err := p.HandleAuthFailure(ctx, n2connData, plainNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"Authentication Failure message: %s", err)
			return err
		}
	case nas.AuthenticationReject:
		//msg counter
		pm.PegCounter(statistics.AuthenticationRejectCounter)
		err := p.HandleAuthReject(ctx)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"Security Mode Complete message: %s", err)
			return err
		}
	case nas.SecurityModeReject:
		//msg counter
		pm.PegCounter(statistics.SecurityModeRejectCounter)
		err := p.HandleSecurityModeReject(ctx, n2connData, plainNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle "+
				"Authentication reject message: %s", err)
			return err
		}
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "unsupported mm message currently, "+
			"msgType(%d)", msgType)
		return nas.ErrInvalidMmMsgType
	}

	return nil
}
