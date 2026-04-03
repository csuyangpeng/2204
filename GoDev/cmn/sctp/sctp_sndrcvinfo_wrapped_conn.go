package sctp

import (
	"net"
	"syscall"
	"time"
	"unsafe"
)

type SCTPSndRcvInfoWrappedConn struct {
	conn *SCTPConn
}

func NewSCTPSndRcvInfoWrappedConn(conn *SCTPConn) *SCTPSndRcvInfoWrappedConn {
	conn.SubscribeEvents(SCTP_EVENT_DATA_IO)
	return &SCTPSndRcvInfoWrappedConn{conn}
}

func (c *SCTPSndRcvInfoWrappedConn) Write(b []byte) (int, error) {
	if len(b) < int(sndRcvInfoSize) {
		return 0, syscall.EINVAL
	}
	info := (*SndRcvInfo)(unsafe.Pointer(&b[0]))
	n, err := c.conn.SCTPWrite(b[sndRcvInfoSize:], info)
	return n + int(sndRcvInfoSize), err
}

func (c *SCTPSndRcvInfoWrappedConn) Read(b []byte) (int, error) {
	if len(b) < int(sndRcvInfoSize) {
		return 0, syscall.EINVAL
	}
	n, info, err := c.conn.SCTPRead(b[sndRcvInfoSize:])
	if err != nil {
		return n, err
	}
	copy(b, toBuf(info))
	return n + int(sndRcvInfoSize), err
}

func (c *SCTPSndRcvInfoWrappedConn) Close() error {
	return c.conn.Close()
}

func (c *SCTPSndRcvInfoWrappedConn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

func (c *SCTPSndRcvInfoWrappedConn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

func (c *SCTPSndRcvInfoWrappedConn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

func (c *SCTPSndRcvInfoWrappedConn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

func (c *SCTPSndRcvInfoWrappedConn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

func (c *SCTPSndRcvInfoWrappedConn) SetWriteBuffer(bytes int) error {
	return c.conn.SetWriteBuffer(bytes)
}

func (c *SCTPSndRcvInfoWrappedConn) GetWriteBuffer() (int, error) {
	return c.conn.GetWriteBuffer()
}

func (c *SCTPSndRcvInfoWrappedConn) SetReadBuffer(bytes int) error {
	return c.conn.SetReadBuffer(bytes)
}

func (c *SCTPSndRcvInfoWrappedConn) GetReadBuffer() (int, error) {
	return c.conn.GetReadBuffer()
}
