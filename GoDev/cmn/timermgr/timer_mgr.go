// Package timermgr manage timers in application,
// include period timer and after timer.
package timermgr

import (
	"container/heap"
	"context"
	"fmt"
	"lite5gc/cmn/rlogger"
	T "lite5gc/cmn/types"
	"sync"
	"time"
)

// TimerMgr manages all the timed task.
type TimerMgr struct {
	timeOutChan chan *OnTimeOut
	timers      timerHeapType
	ticker      *time.Ticker
	wg          *sync.WaitGroup
	addChan     chan *timerType // add timer in loop
	cancelChan  chan int64      // cancel timer in loop
	sizeChan    chan int        // get size in loop
	ctx         context.Context
	cancel      context.CancelFunc
	timerIds    *T.AtomicInt64 // timer identity

	// for debug timer manager
	AfterTimerAddCounter       *T.AtomicInt64
	AfterTimerTimeroutCounter  *T.AtomicInt64
	AfterTimerCancelCounter    *T.AtomicInt64
	PeriodTimerAddCounter      *T.AtomicInt64
	PeriodTimerCancelCounter   *T.AtomicInt64
	PeriodTimerTimeroutCounter *T.AtomicInt64
	TimerPushHeapCounter       *T.AtomicInt64
	TimerPopHeapCounter        *T.AtomicInt64
}

func (p TimerMgr) String() (strbuf string) {
	strbuf += fmt.Sprintf("TimerManager Info:\n")
	strbuf += fmt.Sprintln("Timers:\n", p.timers.Len())
	strbuf += fmt.Sprintln("ticker:\n", *p.ticker)
	return strbuf
}

// NewTimerMgr return a TimerMagr*,
// will create a Timer Manager, initialized and ready for using
func NewTimerMgr(ctx context.Context, numGo int, tickPeriodMicosec int) *TimerMgr {
	timerMgr := &TimerMgr{
		timeOutChan: make(chan *OnTimeOut, T.ChanBufSize1024),
		timers:      make(timerHeapType, 0),
		ticker:      time.NewTicker(time.Duration(tickPeriodMicosec) * time.Microsecond),
		wg:          &sync.WaitGroup{},
		addChan:     make(chan *timerType, T.ChanBufSize1024),
		cancelChan:  make(chan int64, T.ChanBufSize1024),
		sizeChan:    make(chan int),
		timerIds:    T.NewAtomicInt64(0),
	}
	timerMgr.ctx, timerMgr.cancel = context.WithCancel(ctx)
	heap.Init(&timerMgr.timers)

	timerMgr.timerIds = T.NewAtomicInt64(0)

	timerMgr.AfterTimerAddCounter = T.NewAtomicInt64(0)
	timerMgr.AfterTimerTimeroutCounter = T.NewAtomicInt64(0)
	timerMgr.AfterTimerCancelCounter = T.NewAtomicInt64(0)
	timerMgr.PeriodTimerAddCounter = T.NewAtomicInt64(0)
	timerMgr.PeriodTimerCancelCounter = T.NewAtomicInt64(0)
	timerMgr.PeriodTimerTimeroutCounter = T.NewAtomicInt64(0)
	timerMgr.TimerPushHeapCounter = T.NewAtomicInt64(0)
	timerMgr.TimerPopHeapCounter = T.NewAtomicInt64(0)

	go func() {
		timerMgr.wg.Add(1)
		defer timerMgr.wg.Done()
		timerMgr.start()
	}()

	for i := 0; i < numGo; i++ {
		go func() {
			timerMgr.wg.Add(1)
			defer timerMgr.wg.Done()
			timerMgr.handleTimeroutEvent()
		}()
	}
	return timerMgr
}

// TimeOutChannel returns the timeout channel.
func (tm *TimerMgr) TimeOutChannel() chan *OnTimeOut {
	return tm.timeOutChan
}

// AddAfterTimer return timer id, which
// create a after timer with duration and callback,
// insert the timer into the manager to schedule and eanble the timer
func (tm *TimerMgr) AddAfterTimer(sec time.Duration, to *OnTimeOut) int64 {
	return tm.addTimer(time.Now().Add(sec*time.Second), 0, to)
}

// AddPeriodTimer return timer id, which
// create a after timer with duration and callback,
// insert the timer into the manager to schedule and eanble the timer
func (tm *TimerMgr) AddPeriodTimer(sec time.Duration, to *OnTimeOut) int64 {
	return tm.addTimer(time.Now(), sec*time.Second, to)
}

func (tm *TimerMgr) addTimer(when time.Time, interv time.Duration, to *OnTimeOut) int64 {
	if to == nil {
		return int64(-1)
	}
	//create a new timer
	timer := &timerType{
		id:         tm.timerIds.GetAndIncrement(),
		expiration: when,
		interval:   interv,
		timeout:    to,
	}

	tm.addChan <- timer

	if interv == 0 {
		tm.AfterTimerAddCounter.PegCounter()
	} else {
		tm.PeriodTimerAddCounter.PegCounter()
	}

	return timer.id
}

// Size returns the number of current running timer
func (tm *TimerMgr) Size() int {
	return <-tm.sizeChan
}

// CancelTimer cancels a timed task with specified timer ID.
func (tm *TimerMgr) CancelTimer(timerID int64) {
	tm.cancelChan <- timerID
}

// Stop stops the timer manager, stop the timer goroutie
func (tm *TimerMgr) Stop() {
	tm.cancel()
	tm.wg.Wait()
}

func (tm *TimerMgr) getExpired() []*timerType {
	expired := make([]*timerType, 0)
	for tm.timers.Len() > 0 {
		timer := heap.Pop(&tm.timers).(*timerType)
		tm.TimerPopHeapCounter.PegCounter()

		elapsed := time.Since(timer.expiration).Seconds()

		if elapsed > 1.0 {
			// fmt.Printf("timer(%d) elapsed %f\n", timer.id, elapsed)
		}

		if elapsed > 0.0 {
			expired = append(expired, timer)
			continue
		} else {
			heap.Push(&tm.timers, timer)
			tm.TimerPushHeapCounter.PegCounter()
			break
		}
	}
	return expired
}

func (tm *TimerMgr) update(timers []*timerType) {
	if timers != nil {
		for _, t := range timers {
			if t.isRepeat() { // repeatable timer task
				tm.PeriodTimerTimeroutCounter.PegCounter()

				t.expiration = t.expiration.Add(t.interval)
				// if task time out for at least 10 seconds, the expiration time needs
				// to be updated in case this task executes every time timer wakes up.
				if time.Since(t.expiration).Seconds() >= 10.0 {
					t.expiration = time.Now() //should invoke ASSP
				}

				heap.Push(&tm.timers, t)
				tm.TimerPushHeapCounter.PegCounter()
			} else {
				tm.AfterTimerTimeroutCounter.PegCounter()
			}
		}
	}
}

func (tm *TimerMgr) start() {
	rlogger.Trace(T.ModCmn, rlogger.INFO, nil, "TimerMgr %p work routine start", tm)
	defer rlogger.Trace(T.ModCmn, rlogger.ERROR, nil, "TimerMgr %p work routine exit", tm)

	for {
		select {
		case timerID := <-tm.cancelChan:
			index := tm.timers.getIndexByID(timerID)
			if index >= 0 {
				heap.Remove(&tm.timers, index)
			}

		case tm.sizeChan <- tm.timers.Len():

		case <-tm.ctx.Done():
			tm.ticker.Stop()
			return

		case timer := <-tm.addChan:
			heap.Push(&tm.timers, timer)
			tm.TimerPushHeapCounter.PegCounter()

		case <-tm.ticker.C:
			timers := tm.getExpired()
			for _, t := range timers {
				tm.TimeOutChannel() <- t.timeout
			}
			tm.update(timers)
		}
	}
}

func (tm *TimerMgr) handleTimeroutEvent() {
	rlogger.Trace(T.ModCmn, rlogger.INFO, nil, "TimerMgr %p handleTimeroutEvent Loop routine start", tm)
	defer rlogger.Trace(T.ModCmn, rlogger.ERROR, nil, "TimerMgr %p handleTimeroutEvent Loop routine exit", tm)

	for {
		select {
		case timeout := <-tm.TimeOutChannel():
			timeout.Callback(timeout.Params)
		case <-tm.ctx.Done():
			return
		}
	}
}

func (tm *TimerMgr) Destroy() {
	rlogger.Trace(T.ModCmn, rlogger.INFO, nil, "TimerMgr to Destroy")
	tm.cancel()
	tm.wg.Wait()
	rlogger.Trace(T.ModCmn, rlogger.INFO, nil, "TimerMgr destroyed")
}

// ResetPmCounter reset all timer counter to zero
func (tm *TimerMgr) ResetPmCounter() {
	tm.AfterTimerAddCounter.Set(0)
	tm.AfterTimerTimeroutCounter.Set(0)
	tm.AfterTimerCancelCounter.Set(0)
	tm.PeriodTimerAddCounter.Set(0)
	tm.PeriodTimerCancelCounter.Set(0)
	tm.PeriodTimerTimeroutCounter.Set(0)

	tm.TimerPushHeapCounter.Set(0)
	tm.TimerPopHeapCounter.Set(0)
}
