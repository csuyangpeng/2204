package sctp

import (
	"net"
	"sync"
	"syscall"
)

type SCTPListener struct {
	fd int
	m  sync.Mutex
}

func (ln *SCTPListener) Addr() net.Addr {
	laddr, err := sctpGetAddrs(ln.fd, 0, SCTP_GET_LOCAL_ADDRS)
	if err != nil {
		return nil
	}
	return laddr
}

// AcceptSCTP waits for and returns the next SCTP connection to the listener.
func (ln *SCTPListener) AcceptSCTP() (*SCTPConn, error) {
	fd, _, err := syscall.Accept4(ln.fd, 0)
	return NewSCTPConn(fd), err
}

// Accept waits for and returns the next connection connection to the listener.
func (ln *SCTPListener) Accept() (net.Conn, error) {
	return ln.AcceptSCTP()
}

func (ln *SCTPListener) Close() error {
	syscall.Shutdown(ln.fd, syscall.SHUT_RDWR)
	return syscall.Close(ln.fd)
}

func (ln *SCTPListener) SetHeatbeatInterval(hbinv uint32, pathMaxRxt uint32) error {
	return setHbInterval(ln.fd, hbinv, pathMaxRxt)
}

func (ln *SCTPListener) GetHeatbeatInterval() (uint32, error) {
	return getHbInterval(ln.fd)
}

func (ln *SCTPListener) SetInitMsg(info InitMsg) error {
	return setInitOpts(ln.fd, info)
}

func (ln *SCTPListener) GetInitMsg() (*InitMsg, error) {
	return getInitOpts(ln.fd)
}

func (ln *SCTPListener) SetNoDelay(flag bool) error {
	return setNoDelay(ln.fd, flag)
}
