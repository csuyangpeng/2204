package types3gpp

import (
	"fmt"
	"lite5gc/cmn/utils"
	"net"
)

type GtpTunnel struct {
	ipAddr net.IP
	teid   uint32
}

func (p *GtpTunnel) SetIpAddr(addr net.IP) {
	p.ipAddr = make(net.IP, net.IPv6len)
	copy(p.ipAddr, addr)
}

func (p *GtpTunnel) GetIpAddr() net.IP {
	return p.ipAddr
}

func (p *GtpTunnel) GetIpv4Long() uint32 {
	return utils.Ip2long(p.ipAddr)
}

func (p *GtpTunnel) SetTeid(teid uint32) {
	p.teid = teid
}

func (p *GtpTunnel) GetTeid() uint32 {
	return p.teid
}

func (p GtpTunnel) String() string {
	return fmt.Sprintf("GtpTunnel(%s:%d)", p.ipAddr.String(), p.teid)
}
