package types3gpp

//
//import (
//	"fmt"
//)
//
//const (
//	sizeofSD             = 3
//	MaxNumofRequestNSSAI = 8
//	MaxNumofAllowNSSAI   = 8
//	MaxNumofConfNSSAI    = 16
//)
//
//type StdSST int
//
//const (
//	_     StdSST = iota
//	EMBB
//	URLCC
//	MIOT
//)
//
//// Snssai :identifies a Network Slice.
//type Snssai struct {
//	sst      uint8
//	isSdPrst bool
//	sd       [sizeofSD]byte
//}
//
//func (p *Snssai) SetSST(sst uint8) {
//	p.sst = sst
//}
//func (p *Snssai) GetSST() uint8 {
//	return p.sst
//}
//
//func (p *Snssai) SetSD(sd [sizeofSD]byte) error {
//	if len(sd) != sizeofSD {
//		return fmt.Errorf("invalid length of SD, should "+
//			"be %d, but %d", sizeofSD, len(sd))
//	}
//	for i, v := range sd {
//		p.sd[i] = v
//	}
//	p.isSdPrst = true
//	return nil
//}
//
//func (p *Snssai) GetSD() [sizeofSD]byte {
//	return p.sd
//}
//
//func (p *Snssai) IsSdPrst() bool {
//	return p.isSdPrst
//}
//
//func (p *Snssai) String() string {
//	return fmt.Sprintf("SST(%d),SD(%s) ", p.sst, string(p.sd[:]))
//}
//
//// NSSAI
//type NssaiType int
//
//const (
//	Subscribed NssaiType = iota
//	Requested
//	Allowed
//	Rejected
//	Configed
//)
//
//// NSSAI :Network Slice Selection Assistance Information
//type Nssai struct {
//	nssaiType NssaiType
//	snssais   []Snssai
//}
//
//func (p *Nssai) SetType(t NssaiType) {
//	p.nssaiType = t
//}
//func (p *Nssai) GetType() NssaiType {
//	return p.nssaiType
//}
//func (p *Nssai) AddSNssai(snssai Snssai) {
//	//TODO check the length
//	p.snssais = append(p.snssais, snssai)
//}
//func (p *Nssai) HasSnssai(snssai Snssai) bool {
//	for _, v := range p.snssais {
//		if v == snssai {
//			return true
//		}
//	}
//	return false
//}
//func (p *Nssai) GetSnssais() []Snssai {
//	return p.snssais
//}
//func (p *Nssai) String() string {
//	var outStr string
//	outStr = fmt.Sprintf("NSSAI: type(%s), S-NSSAI: { ", p.GetTypeString())
//	for _, v := range p.snssais {
//		outStr += v.String()
//	}
//	outStr += " }"
//
//	return outStr
//}
//func (p *Nssai) GetTypeString() string {
//	var typeStr string
//	switch p.nssaiType {
//	case Subscribed:
//		typeStr = "Subscribed NSSAI"
//	case Requested:
//		typeStr = "Requested NSSAI"
//	case Allowed:
//		typeStr = "Allowed NSSAI"
//	case Rejected:
//		typeStr = "Rejected NSSAI"
//	case Configed:
//		typeStr = "Configed NSSAI"
//	}
//	return typeStr
//}
