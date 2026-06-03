package nasie

type NetSliceIndicationType struct {
	NSSCI bool
	DCNI  bool
}

func (p *NetSliceIndicationType) Reset() {
	p.NSSCI = false
	p.DCNI = false
}

func (p NetSliceIndicationType) String() string {
	var s string
	if p.NSSCI == true {
		s += "NSSCI(True), "
	} else {
		s += "NSSCI(False), "
	}
	if p.DCNI == true {
		s += "DCNI(True) "
	} else {
		s += "DCNI(False) "
	}
	return s
}
