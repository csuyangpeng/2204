package types3gpp

const (
	SizeofGtpTeid         = 4
	SizeofTransportLayAdd = 20
)

type GtpTeId [SizeofGtpTeid]byte

type AddGTPTunnel struct {
	transportLayerAddr [SizeofTransportLayAdd]uint8 //160 bit length
	GtpTeid            GtpTeId
}

func (p *AddGTPTunnel) SetTransportLayerAddr(transLayAddr []uint8) {
	copy(p.transportLayerAddr[:], transLayAddr)
}

func (p *AddGTPTunnel) GetTransportLayerAddr() []uint8 {
	transLayAddress := make([]uint8, 20)
	copy(transLayAddress, p.transportLayerAddr[:])
	return transLayAddress
}
