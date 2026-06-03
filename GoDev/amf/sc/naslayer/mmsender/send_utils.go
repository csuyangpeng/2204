package mmsender

import (
	"context"
	"fmt"
	"lite5gc/amf/context/gctxt"
	"lite5gc/amf/sc/ngaplayer/ngapsender"
	"lite5gc/amf/sc/utils"
	"lite5gc/cmn/nas"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
	"reflect"
)

func VerifyAndBuildSecProtectNasMsg(ueCtxt *gctxt.UeContext, inNasMsg []byte) []byte {
	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)

	var secPrctNasMsg []byte
	var err error

	if !ueCtxt.IsSecCtxtEstb ||
		ueCtxt.SecHdrTypeIncomingMsg == nas.PlainNasMsg ||
		ueCtxt.SecHdrTypeIncomingMsg == nas.IntegrityPrtc {
		return inNasMsg
	} else {
		// check and build the security protected nas message
		// set security protected header
		secPrctNasMsg, err = utils.EncodeSecPrctNasMsg(ueCtxt, nas.IntegrityPrtctCipher, inNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "failed to encode sec header for nas messages. error(%s)", err)
			secPrctNasMsg = inNasMsg
		}
	}

	return secPrctNasMsg
}

func StartNasTimer(ctx context.Context, ueCtxt *gctxt.UeContext, timerType gctxt.TimerType, nasMsg []byte) error {
	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)

	if ueCtxt == nil || nasMsg == nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "invalid input parameters.")
		return fmt.Errorf("startNasTimer: invalid input parameters")
	}

	// start timer for register accept message
	ctx = context.WithValue(ctx, types.UeContextCK, ueCtxt)
	ctx = context.WithValue(ctx, types.TimerTypeCK, timerType)

	tc, err := gctxt.NewTimerCtxt(timerType, NasTimeoutCallbackFunc, ctx)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "failed to create timer context, error(%s)", err)
		return fmt.Errorf("failed to create timer context, error(%s)", err)
	}

	timerMgr, ok := ctx.Value(types.ScTimerMgrCK).(*timermgr.TimerMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "failed to get timer manager")
		return types.ErrFailFindTimerMgr
	}

	var nasMsgPayload []byte
	nasMsgPayload = append(nasMsgPayload, nasMsg[:]...)
	err = gctxt.StartPrcdTimer(ueCtxt, timerMgr, timerType, tc, nasMsgPayload)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "failed to start procedure timer, error(%s)", err)
		return fmt.Errorf("failed to start procedure timer, error(%s)", err)
	}

	return nil
}

func NasTimeoutCallbackFunc(params interface{}) {
	args := reflect.ValueOf(params) //interface to value, which is a slice

	lens := args.Len()
	if lens != 1 {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to handle call back with params.")
		return
	}

	ctxt := (args.Index(0).Interface()).(context.Context)
	ueCtxt, ok := ctxt.Value(types.UeContextCK).(*gctxt.UeContext)
	if !ok {
		return
	}

	rlogger.FuncEntry(types.ModuleAmfNas, ueCtxt)

	timerType := ctxt.Value(types.TimerTypeCK).(gctxt.TimerType)
	tc, ok := ueCtxt.PrcdTimerCtxt[timerType]
	if !ok {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "failed to find timer context for %s", timerType)
		return
	}

	if tc.Coutner >= tc.MaxResentNum {
		err := gctxt.CancelPrcdTimer(ueCtxt, ctxt, timerType)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
				"failed to cancel timer:", timerType)
			return
		}
		// 24.504 5.4.1.3.7 b)	Expiry of timer T3560.
		// The network shall, on the first expiry of the timer T3560,
		// retransmit the AUTHENTICATION REQUEST message and shall reset and start timer T3560.
		// This retransmission is repeated four times, i.e. on the fifth expiry of timer T3560,
		// the network shall abort the 5G AKA based primary authentication and key agreement procedure
		// and any ongoing 5GMM specific procedure and release the N1 NAS signalling connection.
		ueCtxt.SetProcCtxt(nil)
		// TODO an release

		return
	}

	// send the message in scng layer
	sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return
	}

	//get gnb instance id for sending ngap msg
	gnbInstId, err := gctxt.GetGnbInstIdByAmfUeNgapId(ueCtxt.GetAmfUeNgapId())
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil, "failed to get gnb instance id")
		return
	}

	if tc.ResendNasMsg != nil {
		err = sender.ReSendDownlinkNasTransport(gnbInstId, tc.ResendNasMsg)
		if err != nil {
			rlogger.Trace(types.ModuleAmfNas, rlogger.WARN, []interface{}{ueCtxt}, "failed to send registration accetp message")
			return
		}
	} else {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt, "internal error, resend nas msg is nil for timerout event.")
	}

	tc.Coutner++

	gctxt.RestartPrcdTimer(ctxt, tc)

	return
}

func SendDownLinkNasMsg(ctxt context.Context, ueCtxt *gctxt.UeContext, encBuf []byte) (error, []byte) {
	rlogger.FuncEntry(types.ModuleAmfNas, nil)

	var encNgapMsg []byte
	//get gnb instance id for sending ngap msg
	gnbInstId, err := gctxt.GetGnbInstIdByAmfUeNgapId(ueCtxt.GetAmfUeNgapId())
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"fail to get GnbInstId by AmfUeNgapId(%d)", ueCtxt.GetAmfUeNgapId())
		return types.ErrFailFindGnbInstId, encNgapMsg
	}

	//  get the scnglayer
	sender, ok := ctxt.Value(types.NgapSenderCK).(*ngapsender.NgapSender)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(),
			"failed to get ngap layer .")
		return types.ErrFailFindNgapLayer, encNgapMsg
	}

	// send ngap message
	err, encNgapMsg = sender.SendDownlinkNasTransport(ueCtxt, gnbInstId, encBuf)
	if err != nil {
		rlogger.Trace(types.ModuleAmfNas, rlogger.ERROR, nil,
			"failed to send Downlink Nas Transport message")
		return types.ErrFailSendNgapMsg, encNgapMsg
	}
	return nil, encNgapMsg
}
