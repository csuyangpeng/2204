package naslayer

import (
	"bytes"
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/cmn/message/nasmsg"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
)

type NasMgr struct {
	ScInstId uint32
	// msg carrier
	registrationRequest    nasmsg.RegistrationRequestMsg
	registrationComplete   nasmsg.RegistrationCompleteMsg
	upLinkNasTransport     nasmsg.UplinkNasTransportMsg
	deRegistrationRequest  nasmsg.DeRegistrationRequestMsg
	serviceRequest         nasmsg.ServiceRequestMsg
	authenticationResponse nasmsg.AuthenticationResponseMsg
	authenticationFailure  nasmsg.AuthenticationFailureMsg
	securityModeComplete   nasmsg.SecurityModeCompleteMsg
	securityModeReject     nasmsg.SecurityModeRejectMsg
	identityResponse       nasmsg.IdentityResponseMsg
}

func (p NasMgr) String() string {
	strBuf := fmt.Sprintf("NasMgr Info:\n")
	strBuf += fmt.Sprintln("ScInstId:", p.ScInstId)
	return strBuf
}

func NewLayerMgr(scInst uint32) *NasMgr {
	return &NasMgr{ScInstId: scInst}
}

// HandleIncomingNasMsg process nas initial ue message
func (p *NasMgr) HandleIncomingNasMsg(ctx context.Context, n2connData *gctxt.N2ConnCtxt, nasData []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "received nas Msg:%v", nasData)
	nasMsg := bytes.NewReader(nasData)

	// 1. decode extract the EPD
	epd, err := nas.GetEpd(nasMsg)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to decode nas epd")
		return fmt.Errorf("failed to decode nas epd")
	}

	switch epd {
	case nas.Epd5gsMobMgntMsg:
		//2. decode Security header type and Spare half octet
		secHeaderType, err := nas.GetSecHeaderType(nasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to decode nas security header type")
			return fmt.Errorf("failed to decode nas security header type")
		}
		// save the security header type in incoming nas message
		ctx = context.WithValue(ctx, types.SecHeaderTypeInNasMsgCK, secHeaderType)

		switch secHeaderType {
		case nas.PlainNasMsg:
			err = p.HandleMmMessage(ctx, n2connData, nasMsg)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to handle mm message")
				return fmt.Errorf("failed to handle mm message")
			}
		case nas.IntegrityPrtc, nas.IntegrityPrtctNewSecCtxt,
			nas.IntegrityPrtctCipherNewSecCtxt, nas.IntegrityPrtctCipher:
			err = p.HandleSecMmMessage(ctx, n2connData, secHeaderType, nasMsg)
			if err != nil {
				rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "failed to handle sec mm message")
				return fmt.Errorf("failed to handle sec mm message")
			}
		default:
			rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "unsupported header (%x)", secHeaderType)
		}
	default:
		rlogger.Trace(types.ModuleAmfNas, rlogger.DEBUG, nil, "unsupported epd (%x)", epd)
		return fmt.Errorf("unsupported epd (%x)", epd)
	}
	return nil
}
