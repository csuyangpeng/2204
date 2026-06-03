package timermgr

import "time"

// timerType define a Timer data struct
// 'expiration' is the time when timer time out, if 'interval' > 0
// the timer will time out periodically, 'timeout' contains the callback
// to be called when times out
type timerType struct {
	id         int64
	expiration time.Time
	interval   time.Duration
	timeout    *OnTimeOut
	index      int // for container/heap
}

func (t *timerType) isRepeat() bool {
	return int64(t.interval) > 0
}

// TimerFunc define callback function type after timerout
type TimerFunc func(interface{})

// OnTimeOut define a timerout task.
type OnTimeOut struct {
	Callback TimerFunc
	Params   interface{}
}

// NewOnTimeOut returns OnTimeOut.
func NewOnTimeOut(cb TimerFunc, p ...interface{}) *OnTimeOut {
	return &OnTimeOut{
		Callback: cb,
		Params:   p,
	}
}
