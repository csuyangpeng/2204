package n2sctp

import (
	"context"
	"sync"

	"fmt"
	"lite5gc/amf/n2proc/gnblayer"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sctp"
	"lite5gc/cmn/types"
)

// N2SctpConn define the n2 sctp connection
type N2SctpConn struct {
	instId    uint32 //instance id
	gnbIpAddr string //gNodeB Ip Address
	conn      *sctp.SCTPConn
	closeOnce sync.Once

	belong *N2Listener
	wg     *sync.WaitGroup

	sendch chan *types.MsgData
	//recvch chan *types.MsgBuf

	gnbMgr gnblayer.GnbLayer
	ctx    context.Context
	cancel context.CancelFunc
}

func (p *N2SctpConn) SctpConnDump() {}

// SetGnbIpAddr set gnb ip address
func (p *N2SctpConn) SetGnbIpAddr(addr string) {
	p.gnbIpAddr = addr
	p.gnbMgr.SetGnbIP(addr)
}

// GetGnbIpAddr return gnb ip address
func (p *N2SctpConn) GetGnbIpAddr() string {
	return p.gnbIpAddr
}

// SetInstId set gnb sctp instance id
func (p *N2SctpConn) SetInstId(id uint32) {
	//p.rwMu.RLock()
	//defer p.rwMu.RUnlock()

	p.instId = id
}

// GetInstId get gnb conn instance id
func (p *N2SctpConn) GetInstId() uint32 {
	//p.rwMu.RLock()
	//defer p.rwMu.RUnlock()

	return p.instId
}

// NewN2SctpConn return a sctp connection struct pointer
func NewN2SctpConn(id uint32, listener *N2Listener, sctpConn *sctp.SCTPConn, appWg *sync.WaitGroup) *N2SctpConn {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	n2SctpConn := &N2SctpConn{
		instId: id,
		belong: listener,
		conn:   sctpConn,
		wg:     appWg,
		//sendch: make(chan *types.MsgData, listener.bufferSize), //send sctp message channel for uplayer
		sendch: make(chan *types.MsgData), //send sctp message channel for uplayer
		//recvch: make(chan *types.MsgBuf, listener.bufferSize), //receive sctp message channel for uplayer to handle
	}

	// subscriber sctp event
	n2SctpConn.conn.SubscribeEvents(
		sctp.SCTP_EVENT_DATA_IO | sctp.SCTP_EVENT_ASSOCIATION | sctp.SCTP_EVENT_ADDRESS | sctp.SCTP_EVENT_SHUTDOWN)

	//register event handler
	n2SctpConn.conn.SetSctpEventHandler(n2SctpConn, HandleSctpNotifyEvent)

	ctxt := context.WithValue(listener.ctx, types.NgSctpSerCK, listener)
	n2SctpConn.ctx, n2SctpConn.cancel = context.WithCancel(ctxt)

	n2SctpConn.gnbMgr.Init(id, n2SctpConn.gnbIpAddr, n2SctpConn.sendch, n2SctpConn.ctx)

	return n2SctpConn
}

func HandleSctpNotifyEvent(event []byte, conn sctp.SctpConnIf) error {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, conn)

	header, err := sctp.ParseSctpNotifyHeader(event)
	if err != nil {
		return fmt.Errorf("failed to get notify event header, error(%s)", err)
	}

	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, conn, "receive sctp event %s", sctp.SCTPNotificationType(header.SnType))

	sctpConn, ok := conn.(*N2SctpConn)
	if !ok {
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, conn, "invalid convert connection")
	}

	switch sctp.SCTPNotificationType(header.SnType) {
	case sctp.SCTP_ASSOC_CHANGE:
		sac, err := sctp.ParseSctpNotify_AssocChange(event)
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, conn, "failed to get sctp association change error(%s)", err)
			return fmt.Errorf("failed to get sctp association change error(%s)", err)
		}
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, conn, "sctp assocation change, status: ", sctp.SctpSacState(sac.SacState))

		//TODO implement different actions here
		switch sctp.SctpSacState(sac.SacState) {
		case sctp.SCTP_COMM_UP:
		case sctp.SCTP_COMM_LOST:
			sctpConn.conn.Close()
		case sctp.SCTP_RESTART:
		case sctp.SCTP_SHUTDOWN_COMP:
		case sctp.SCTP_CANT_STR_ASSOC:
		default:
		}
		// call back handler here
	case sctp.SCTP_PEER_ADDR_CHANGE:
		spac, err := sctp.ParseSctpNotify_PeerAddrChange(event)
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, conn, "failed to get sctp peer addr change, error(%s)", err)
			return fmt.Errorf("failed to get sctp peer addr change, error(%s)", err)
		}
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, conn, "sctp peer addr change, status: ", sctp.SctpSpcState(spac.SpcState))

		switch sctp.SctpSpcState(spac.SpcState) {
		case sctp.SCTP_ADDR_AVAILABLE:
		case sctp.SCTP_ADDR_UNREACHABLE:
			//TODO only one peer arress, abort is ok, if multi-address applied, should not abort the sctp connection directly
			sctpConn.conn.Abort()
		case sctp.SCTP_ADDR_REMOVED:
		case sctp.SCTP_ADDR_ADDED:
		case sctp.SCTP_ADDR_MADE_PRIM:
		case sctp.SCTP_ADDR_CONFIRMED:
		default:
		}

	case sctp.SCTP_SEND_FAILED:
	case sctp.SCTP_REMOTE_ERROR:
	case sctp.SCTP_SHUTDOWN_EVENT:
		sctpConn.conn.Close()
	case sctp.SCTP_PARTIAL_DELIVERY_EVENT:
	case sctp.SCTP_ADAPTATION_INDICATION:
	case sctp.SCTP_AUTHENTICATION_INDICATION:
	case sctp.SCTP_SENDER_DRY_EVENT:
	}

	return nil
}
