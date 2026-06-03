package gctxt

import (
	"context"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/timermgr"
	"lite5gc/cmn/types"
	"time"
)

type TimerType uint8

// 24.501 Table 10.2.2
const (
	T3550 TimerType = iota
	T3560
	T3570
	T3513
	T3522
	T3555
	T3565
	T3590
	T3591
	T3592
	T3593
)

func (p TimerType) String() string {
	var rt string
	switch p {
	case T3550:
		rt = "T3550"
	case T3560:
		rt = "T3560"
	case T3570:
		rt = "T3570"
	case T3513:
		rt = "T3513"
	case T3522:
		rt = "T3522"
	case T3555:
		rt = "T3555"
	case T3565:
		rt = "T3565"
	case T3590:
		rt = "T3590"
	case T3591:
		rt = "T3591"
	case T3592:
		rt = "T3592"
	case T3593:
		rt = "T3593"
	}
	return rt
}

const (
	MaxRestransNumT3550 uint8 = 4 //3gpp 24.501 5.5.1.2.8 c)
	MaxRestransNumT3560 uint8 = 4 //3gpp 24.501 6.3.3.5 a)
	MaxRestransNumT3570 uint8 = 4 //3gpp 24.501 5.4.3.6 b)
	MaxRestransNumT3522 uint8 = 4 //3gpp 24.501 5.5.2.3.5 a)
	MaxRestransNumT3513 uint8 = 2 //network determined
)

type TimerCtxt struct {
	TimerId       int64
	MaxResentNum  uint8
	Duration      int64
	Coutner       uint8
	ResendNasMsg  []byte
	TimeoutCBFunc *timermgr.OnTimeOut
}

func NewTimerCtxt(timerType TimerType,
	cbFunc timermgr.TimerFunc,
	para context.Context) (*TimerCtxt, error) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	tc := &TimerCtxt{}
	tc.TimeoutCBFunc = timermgr.NewOnTimeOut(cbFunc, para)
	tc.Coutner = 0
	switch timerType {
	case T3550:
		//tc.Duration = int64(configure.AmfConf.NAS.T3550sec)
		tc.Duration = 1
		tc.MaxResentNum = MaxRestransNumT3550 //24.501 5.5.1.2.8 c)
	case T3560:
		//tc.Duration = int64(configure.AmfConf.NAS.T3560sec)
		tc.Duration = 1
		tc.MaxResentNum = MaxRestransNumT3560 //24.501 5.4.1.2.4.4 a)
	case T3570:
		//tc.Duration = int64(configure.AmfConf.NAS.T3570sec)
		tc.Duration = 1
		tc.MaxResentNum = MaxRestransNumT3570 //24.501 5.4.3.6 b)
	case T3513:
		//tc.Duration = int64(configure.AmfConf.NAS.T3513Sec)
		tc.Duration = 1
		tc.MaxResentNum = MaxRestransNumT3513 //network determined
	case T3522:
		//tc.Duration = int64(configure.AmfConf.NAS.T3522sec)
		tc.Duration = 1
		tc.MaxResentNum = MaxRestransNumT3522 //24.501 5.5.2.3.5 a)
	case T3555:
	case T3565:
	case T3590:
	case T3591:
	case T3592:
	case T3593:
	default:
	}
	return tc, nil
}

type PrcdTimerCtxt map[TimerType]*TimerCtxt

func StartPrcdTimer(ueCtxt *UeContext,
	tm *timermgr.TimerMgr,
	timerType TimerType,
	tc *TimerCtxt,
	nasMsg []byte) error {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)
	//input parameter
	//check the timer for the ue
	_, ok := ueCtxt.PrcdTimerCtxt[timerType]
	if ok {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, "find the timer(%s),the ueCtxt.PrcdTimerCtxt is %s",
			timerType, ueCtxt.PrcdTimerCtxt)
		//cancel the timer
		//CancelPrcdTimer(ueCtxt, tm, timerType)
	}
	tc.ResendNasMsg = nasMsg
	duration := time.Duration(tc.Duration)
	tc.TimerId = tm.AddAfterTimer(duration, tc.TimeoutCBFunc)
	ueCtxt.PrcdTimerCtxt[timerType] = tc
	return nil
}

func CancelPrcdTimer(ueCtxt *UeContext,
	ctxt context.Context,
	timerType TimerType) error {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)

	tm, ok := ctxt.Value(types.ScTimerMgrCK).(*timermgr.TimerMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, ueCtxt.GetImsiPtr(), "failed to get timer manager")
		return types.ErrFailFindTimerMgr
	}

	tc, ok := ueCtxt.PrcdTimerCtxt[timerType]
	if !ok {
		rlogger.Trace(types.ModuleAmfCtxt, rlogger.DEBUG, ueCtxt, "failed to find timer(%s),the ueCtxt.PrcdTimerCtxt is %s",
			timerType, ueCtxt.PrcdTimerCtxt)
		return fmt.Errorf("failed to find timer(%v),the ueCtxt.PrcdTimerCtxt is %v",
			timerType, ueCtxt.PrcdTimerCtxt)
	}
	tm.CancelTimer(tc.TimerId)
	delete(ueCtxt.PrcdTimerCtxt, timerType)
	return nil
}

func RestartPrcdTimer(ctxt context.Context, tc *TimerCtxt) {
	rlogger.FuncEntry(types.ModuleAmfCtxt, nil)

	tm, ok := ctxt.Value(types.ScTimerMgrCK).(*timermgr.TimerMgr)
	if !ok {
		rlogger.Trace(types.ModuleAmfMM, rlogger.ERROR, nil, "failed to get timer manager")
		return
	}

	tm.CancelTimer(tc.TimerId)
	duration := time.Duration(tc.Duration)
	tc.TimerId = tm.AddAfterTimer(duration, tc.TimeoutCBFunc)
}
