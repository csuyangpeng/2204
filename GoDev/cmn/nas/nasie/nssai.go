package nasie

import (
	"bytes"
	"fmt"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"reflect"
)

type Nssai struct {
	Snssais []SNssai
}

func (p *Nssai) Reset() {
	for i := 0; i < len(p.Snssais); i++ {
		p.Snssais[i].Reset()
	}
}

const (
	sizeofSD             = 3
	MaxNumofRequestNSSAI = 8
	MaxNumofAllowNSSAI   = 8
	MaxNumofConfNSSAI    = 16
)

func (p *Nssai) Encode() []byte {
	var buf []byte
	for i := 0; i < len((*p).Snssais); i++ {
		SnssaiValueOctet := (*p).Snssais[i].Encode()
		SNssaiLength := byte(len(SnssaiValueOctet))
		// L
		buf = append(buf, SNssaiLength)
		// V
		buf = append(buf, SnssaiValueOctet[:]...)
	}
	return buf
}

func (p *Nssai) Decode(msgBuf *bytes.Reader) error {
	rlogger.FuncEntry(types.ModCmn, nil)
	length, _ := msgBuf.ReadByte()
	for length > 0 {
		snssai := SNssai{}
		err := snssai.Decode(msgBuf)
		if err != nil {
			return fmt.Errorf("failed to decode snssai")
		}
		//rlogger.Trace(types.ModCmn, rlogger.ERROR, nil,  "Decode Request Nssai snssai ==",snssai)
		p.Snssais = append(p.Snssais, snssai)
		length -= snssai.GetLen()
	}

	return nil
}

func (p *Nssai) AddSNssai(snssai SNssai) {
	//TODO check the length
	p.Snssais = append(p.Snssais, snssai)
}

func (p *Nssai) HasSnssai(snssai SNssai) bool {
	for _, v := range p.Snssais {
		if v == snssai {
			return true
		}
	}
	return false
}

func HasIntersectionSet(n1 Nssai, n2 Nssai) []SNssai {
	// 一模一样
	if reflect.DeepEqual(n1, n2) {
		return n1.Snssais
	} else {
		// 并集
		var intSet []SNssai
		for i := 0; i < len(n1.Snssais); i++ {
			for j := 0; j < len(n2.Snssais); j++ {
				if reflect.DeepEqual(n1.Snssais[i], n2.Snssais[j]) {
					intSet = append(intSet, n1.Snssais[i])
				}
			}
		}
		return intSet
	}
}
