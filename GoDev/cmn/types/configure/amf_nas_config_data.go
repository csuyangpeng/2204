package configure

import (
	"fmt"
	"lite5gc/cmn/nas/nasie"
	"lite5gc/cmn/types3gpp"
)

// nas configuration
type NasConfig struct {
	SecCap     types3gpp.SecurityCapability
	T3512min   int
	T3513Sec   int
	T3502min   int
	T3550sec   int //Retransmission of REGISTRATION ACCEPT message
	T3560sec   int //Retransmission of AUTHENTICATION REQUEST message or SECURITY MODE COMMAND message
	T3570sec   int //Retransmission of IDENTITY REQUEST message
	T3522sec   int //Retransmission of DEREGISTRATION REQUEST message
	T3555sec   int //Retransmission of CONFIGURATION UPDATE COMMAND message
	T3565sec   int //Retransmission of NOTIFICATION message
	TICSsec    int //Initial Context Setup message 2s
}

func (p NasConfig) String() (strbuf string) {
	strbuf += fmt.Sprintf("NasConfig Info:\n")
	strbuf += fmt.Sprintln("T3512min: ", p.T3512min)
	strbuf += fmt.Sprintln("T3502min: ", p.T3502min)
	strbuf += fmt.Sprintln("T3550sec: ", p.T3550sec)
	strbuf += fmt.Sprintln("T3560sec: ", p.T3560sec)
	strbuf += fmt.Sprintln("T3570sec: ", p.T3570sec)
	strbuf += fmt.Sprintln("T3522sec: ", p.T3522sec)
	strbuf += fmt.Sprintln("T3555sec: ", p.T3555sec)
	strbuf += fmt.Sprintln("T3565sec: ", p.T3565sec)
	return strbuf
}

func GetAmfNasT3512Timer() nasie.GprsTimer3 {
	AmfConf.ConfRWlock.RLock()
	defer AmfConf.ConfRWlock.RUnlock()

	timer := nasie.GprsTimer3{}
	timer.Uint = nasie.OneMin
	timer.TimerValue = uint8(AmfConf.NAS.T3512min)

	return timer
}

func GetAmfNasT3502Timer() nasie.GprsTimer2 {
	AmfConf.ConfRWlock.RLock()
	defer AmfConf.ConfRWlock.RUnlock()

	timer := nasie.GprsTimer2{}
	timer.Uint = nasie.OneMinute
	timer.TimerValue = uint8(AmfConf.NAS.T3502min)

	return timer
}
