package n2sctp

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"lite5gc/cmn/idmgr"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/sctp"
	"lite5gc/cmn/types"
	"sync"
)

// Start create three goroutie to process all the message both from sctp and ipc
func (p *N2SctpConn) Start(ctxt context.Context) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	fmt.Printf("ngap (%d) is running...\n", p.instId)

	loopers := []func(*N2SctpConn, *sync.WaitGroup){readLoop, writeLoop}
	for _, loop := range loopers {
		go loop(p, p.wg)
	}

	go func() {
		p.wg.Add(1)
		defer p.wg.Done()
		p.gnbMgr.LoopProcScMsg(ctxt)
	}()
}

// Close terminate the sctp connection
func (p *N2SctpConn) Close() {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	p.closeOnce.Do(func() {
		//close all the related go routine
		p.cancel()

		//disconnect,release n2aplayer
		p.gnbMgr.CleanUp()

		//close the conn
		err := p.conn.Close()
		if err != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, p,
				"failed to close sctp conn. error(%v)", err)
		}

		//return the instance id
		idmgr.GetInst().ReturnID(string(types.NGAP), p.GetInstId())

		// cleanup gnb information
		//err = ngap.DeleteGnbInfo(ngap.GnbIPKey(p.gnbIpAddr))
		//if err != nil {
		//	rlogger.Trace(MODULE_ID, types.ERROR, p, "failed to release GnbInfo by index(GnbIdKey:%d) "+
		//		",error(%s)", p.GetInstID(), err)
		//}
		//rlogger.Trace(MODULE_ID, types.WARN, nil, "SCTP Connection Close: ngap (%d) exit", p.GetInstID())

		fmt.Printf("SCTP Connection Close: ngap (%d) exit \n", p.GetInstId())
		SctpSer.DelSctpConn(p.GetInstId())
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "delete ngap id (%d)", p.GetInstId())
		//trigger alarm
		//oamagent.SetAlarm(webTypes.N2SctpConnDisconneted, fmt.Sprintf("gnb(%s) disconnected.", p.gnbIpAddr))

	})
}

// Abort the sctp connection
func (p *N2SctpConn) Abort() {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)
	p.closeOnce.Do(func() {
		//close all the related go routine
		p.cancel()

		//disconnect,release n2aplayer
		//p.ngapLayer.CleanUp()

		//close the conn
		p.conn.Abort()

		//return the instance id
		idmgr.GetInst().ReturnID(string(types.NGAP), p.GetInstId())

		//err := ngap.DeleteGnbInfo(ngap.GnbIPKey(p.gnbIpAddr))
		//if err != nil {
		//	rlogger.Trace(MODULE_ID, types.ERROR, nil, "failed to release GnbInfo by index(GnbIdKey:%d) "+
		//		",error(%s)", p.GetInstID(), err)
		//}

		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, p, "SCTP Connection Abort: ngap (%d) exit", p.GetInstId())
		fmt.Printf("SCTP Connection Abort: ngap (%d) exit \n", p.GetInstId())

		////trigger alarm
		//oamagent.SetAlarm(webTypes.N2SctpConnDisconneted, configure.GetAmfWebNo(),
		//	fmt.Sprintf("gnb(%s) disconnected.", p.gnbIpAddr), p.gnbIpAddr)
		//agent.AlarMapIdAlarmFlag[webTypes.N2SctpConnDisconneted] = true

		fmt.Printf("SCTP Connection Close: ngap (%d) exit \n", p.GetInstId())
		SctpSer.DelSctpConn(p.GetInstId())
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "delete ngap id (%d)", p.GetInstId())
	})
}

// readLoop read msg data from the sctp socket,
// send into recv channel
func readLoop(n2conn *N2SctpConn, wg *sync.WaitGroup) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	conn := n2conn.conn
	cDone := n2conn.ctx.Done()
	sDone := n2conn.belong.ctx.Done()
	//cRecvch := n2conn.recvch

	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, "SCTP Connection readloop routine start")
	wg.Add(1)
	defer func() {
		if p := recover(); p != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "panics: %v", p)
		}
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "SCTP Connection readloop routine exit")

		n2conn.Close()
		wg.Done()

	}()
	go func() {
		select {
		case <-cDone: // connection closed
			n2conn.Close()
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "SCTP Connection readLoop monitor receiving cancel signal from conn")
			return
		case <-sDone: // server closed
			n2conn.Close()
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "SCTP Connection readLoop monitor receiving cancel signal from server")
			return
		}
	}()

	for {
		select {
		default:
			msgBuf := &types.MsgBuf{}
			msgBuf.Buffer = make([]byte, types.BufSize8192)
			n, err := conn.Read(msgBuf.Buffer)
			if err != nil {
				if err == io.EOF {
					rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "sctp connection closed. err : %s", err)
					return //sctp will be close in defer
				}
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "Read Error. err : %s", err)
				return //TODO handle different error
			}
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil, "read msg len %d, msg ", n)
			msgBuf.MsgLen = n
			n2conn.gnbMgr.HandleSctpMsg(msgBuf)
		}
	}
}

// writeLoop write the data from cSendch, and send out to sctp socket
func writeLoop(n2conn *N2SctpConn, wg *sync.WaitGroup) {
	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)

	conn := n2conn.conn
	cDone := n2conn.ctx.Done()
	sDone := n2conn.belong.ctx.Done()
	cSendch := n2conn.sendch

	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, " SCTP Connection writeLoop routine start")
	wg.Add(1)
	defer func() {
		if p := recover(); p != nil {
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "panics: %v", p)
		}
		n2conn.Close()
		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "SCTP Connection writeLoop routine exit")
		wg.Done()
	}()

	for {
		select {
		case <-cDone: // connection closed
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "SCTP Connection writeLoop monitor receiving cancel signal from conn")
			return
		case <-sDone: // server closed
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "SCTP Connection writeLoop monitor receiving cancel signal from server")
			return
		case msgBuf := <-cSendch: //send a message
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil,
				"send the message, len(%d), body: \n%v ",
				msgBuf.MsgLen, hex.Dump([]byte(msgBuf.MsgData)))
			if msgBuf.MsgLen <= 0 {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "invalid message with len %d", msgBuf.MsgLen)
				continue
			}
			//n, err := conn.Write(msgBuf.Buffer[:msgBuf.MsgLen])
			info := &sctp.SndRcvInfo{
				Stream: uint16(1),          //send with stream 1
				PPID:   uint32(0x3c000000), //ngap ppid, 60
			}
			n, err := conn.SCTPWrite([]byte(msgBuf.MsgData)[:msgBuf.MsgLen], info)
			if err != nil {
				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "failed to send message, error: ", err)
				continue
			}
			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil, "Send (%d) Message Success.", n)
		}
	}
}

// handleLoop handel the message data from cRecvch, from Sctp Socket or SC goroutie
//func handleLoop(n2conn *N2SctpConn, wg *sync.WaitGroup) {
//	rlogger.FuncEntry(types.ModuleAmfN2Proc, nil)
//
//	cDone := n2conn.ctx.Done()
//	sDone := n2conn.belong.ctx.Done()
//	cRecvch := n2conn.recvch
//
//	rlogger.Trace(types.ModuleAmfN2Proc, rlogger.INFO, nil, "SCTP Connection handleLoop routine start")
//	wg.Add(1)
//	defer func() {
//		if p := recover(); p != nil {
//			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "panics: %v", p)
//		}
//		n2conn.Close()
//		rlogger.Trace(types.ModuleAmfN2Proc, rlogger.ERROR, nil, "SCTP Connection handleLoop routine exit")
//		wg.Done()
//	}()
//
//	n2conn.ngapLayer.SetGnbIP(n2conn.gnbIpAddr)
//	for {
//		select {
//		case <-cDone: // connection closed
//			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "SCTP Connection handleLoop monitor receiving cancel signal from conn")
//			return
//
//		case <-sDone: // server closed
//			rlogger.Trace(types.ModuleAmfN2Proc, rlogger.WARN, nil, "SCTP Connection handleLoop monitor receiving cancel signal from server")
//			return
//
//		case msgBuf := <-cRecvch: //received messages from ran node, handle by ngap layer
//			//rlogger.Trace(MODULE_ID, types.DEBUG, nil,  "receive a message, len(%d), body(%s).", msgBuf.MsgLen, string(msgBuf.Buffer))
//			err := sc.ngapLayer.HandleSctpMsg(msgBuf)
//			if err != nil {
//				rlogger.Trace(types.ModuleAmfN2Proc, rlogger.DEBUG, nil, "handle sctp message failed. err:", err)
//			}
//			//default:
//			//	//receive messages from sc goroutie and handley by ngap layer
//			//	for _, v := range recvScChans {
//			//		select {
//			//		case scMsg := <-v:
//			//			err := sc.ngapLayer.HandleScMsg(scMsg)
//			//			if err != nil {
//			//				rlogger.Trace(MODULE_ID, types.DEBUG, nil,  "handle sc message failed. err:", err)
//			//			}
//			//		default:
//			//			break //break select
//			//		}
//			//	}
//			//	time.Sleep(time.Microsecond * 10)
//		}
//	}
//}
