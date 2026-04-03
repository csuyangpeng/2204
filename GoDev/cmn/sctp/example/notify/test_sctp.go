package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"lite5gc/cmn/sctp"
)

const maxstreamid uint16 = 11

func HandleSctpNotifyEventOnServer(event []byte, conn sctp.SctpConnIf) error {
	header, err := sctp.ParseSctpNotifyHeader(event)
	if err != nil {
		return fmt.Errorf("failed to get notify event header, error(%s)", err)
	}

	log.Printf("receive sctp event %s", sctp.SCTPNotificationType(header.SnType))
	sctpConn, ok := conn.(*sctp.SCTPConn)
	if !ok {
		return fmt.Errorf("failed to get sctp conn, error(%v)", err)
	}
	stat, err := sctpConn.GetSctpStatus()
	fmt.Printf("stats(%s),err(%v)\n", stat, err)

	switch sctp.SCTPNotificationType(header.SnType) {
	case sctp.SCTP_ASSOC_CHANGE:
		sac, err := sctp.ParseSctpNotify_AssocChange(event)
		if err != nil {
			return fmt.Errorf("failed to get sctp association change error(%s)", err)
		}

		fmt.Println("sctp assocation change, status: ", sctp.SctpSacState(sac.SacState))

	case sctp.SCTP_PEER_ADDR_CHANGE:
		fmt.Println("start peer addr change")
		spac, err := sctp.ParseSctpNotify_PeerAddrChange(event)
		if err != nil {
			return fmt.Errorf("failed to get sctp peer addr change, error(%s)", err)
		}

		fmt.Println("sctp peer addr change, status: ", sctp.SctpSpcState(spac.SpcState))

		switch sctp.SctpSpcState(spac.SpcState) {
		case sctp.SCTP_ADDR_AVAILABLE:
		case sctp.SCTP_ADDR_UNREACHABLE:
			//sctpConn.Abort()
		case sctp.SCTP_ADDR_REMOVED:
		case sctp.SCTP_ADDR_ADDED:
		case sctp.SCTP_ADDR_MADE_PRIM:
		case sctp.SCTP_ADDR_CONFIRMED:
		default:
		}

	case sctp.SCTP_SEND_FAILED:
	case sctp.SCTP_REMOTE_ERROR:
	case sctp.SCTP_SHUTDOWN_EVENT:
		//conn.Close()
	case sctp.SCTP_PARTIAL_DELIVERY_EVENT:
	case sctp.SCTP_ADAPTATION_INDICATION:
	case sctp.SCTP_AUTHENTICATION_INDICATION:
	case sctp.SCTP_SENDER_DRY_EVENT:
	}

	return nil
}

func HandleSctpNotifyEventOnClient(event []byte, conn sctp.SctpConnIf) error {
	header, err := sctp.ParseSctpNotifyHeader(event)
	if err != nil {
		return fmt.Errorf("failed to get notify event header, error(%s)", err)
	}

	log.Printf("receive sctp event %s", sctp.SCTPNotificationType(header.SnType))

	switch sctp.SCTPNotificationType(header.SnType) {
	case sctp.SCTP_ASSOC_CHANGE:
	case sctp.SCTP_PEER_ADDR_CHANGE:
	case sctp.SCTP_SEND_FAILED:
	case sctp.SCTP_REMOTE_ERROR:
	case sctp.SCTP_SHUTDOWN_EVENT:
	case sctp.SCTP_PARTIAL_DELIVERY_EVENT:
	case sctp.SCTP_ADAPTATION_INDICATION:
	case sctp.SCTP_AUTHENTICATION_INDICATION:
	case sctp.SCTP_SENDER_DRY_EVENT:
	}

	return nil
}
func serveClient(conn net.Conn) error {
	for {
		buf := make([]byte, 254)
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		log.Printf("read: %d", n)
		n, err = conn.Write(buf[:n])
		if err != nil {
			return err
		}
		log.Printf("write: %d", n)
	}
}

func serveClient_v1(conn *sctp.SCTPConn) error {
	defer func() {
		log.Println("serveClient_v1 quit ", conn)
	}()

	conn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO | sctp.SCTP_EVENT_ASSOCIATION | sctp.SCTP_EVENT_ADDRESS | sctp.SCTP_EVENT_SHUTDOWN)
	//register event handler
	conn.SetSctpEventHandler(conn, HandleSctpNotifyEventOnServer)

	stat, err := conn.GetSctpStatus()
	fmt.Printf("stats(%s),err(%v)\n", stat, err)
	for {

		status, err := conn.GetSctpStatus()
		log.Printf("conn status(%s),error(%s)\n", status, err)

		buf := make([]byte, 254)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("serveClient_v1 conn.Read fail:", err)
			return err
		}
		log.Printf("read: %d: %v", n, buf[:n])
		n, err = conn.Write(buf[:n])
		if err != nil {
			log.Println("serveClient_v1 conn.Write fail:", err)
			return err
		}
		log.Printf("write: %d", n)
	}

}

func main() {
	var server = flag.Bool("server", false, "")
	var lip = flag.String("lip", "0.0.0.0", "")
	var rip = flag.String("rip", "0.0.0.0", "")
	var rport = flag.Int("rport", 0, "")
	var lport = flag.Int("lport", 0, "")

	flag.Parse()

	lips := []net.IPAddr{}
	for _, i := range strings.Split(*lip, ",") {
		if a, err := net.ResolveIPAddr("ip", i); err == nil {
			log.Printf("Resolved lip address '%s' to %s", i, a)
			lips = append(lips, *a)
		} else {
			log.Printf("Error resolving address '%s': %v", i, err)
		}
	}

	rips := []net.IPAddr{}
	for _, i := range strings.Split(*rip, ",") {
		if a, err := net.ResolveIPAddr("ip", i); err == nil {
			log.Printf("Resolved rip address '%s' to %s", i, a)
			rips = append(rips, *a)
		} else {
			log.Printf("Error resolving rip '%s': %v", i, err)
		}
	}

	laddr := &sctp.SCTPAddr{
		IPAddrs: lips,
		Port:    *lport,
	}
	raddr := &sctp.SCTPAddr{
		IPAddrs: rips,
		Port:    *rport,
	}
	log.Printf("raw addr: %+v\n", laddr.ToRawSockAddrBuf())

	if *server {
		//ln, err := sctp.ListenSCTP("sctp", laddr)
		ln, err := sctp.ListenSCTPExt("sctp", laddr, sctp.InitMsg{NumOstreams: maxstreamid, MaxInstreams: maxstreamid})
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("Listen on %s", ln.Addr())

		initmsg := sctp.InitMsg{NumOstreams: 5, MaxInstreams: 5, MaxAttempts: 6, MaxInitTimeout: 5}
		err = ln.SetInitMsg(initmsg)
		if err != nil {
			log.Println("failed to set init msg, error ", err)
		}

		initmsgPtr, err := ln.GetInitMsg()
		if err != nil {
			log.Println("failed to get init msg, error ", err)
		}
		log.Printf("init msg (%s)", initmsgPtr)

		//inv, err := ln.GetHeatbeatInterval()
		//log.Printf("heatbeat interval (%d), error(%v)\n", inv, err)

		err = ln.SetHeatbeatInterval(2000, 2)
		if err != nil {
			log.Fatalf("failed to set heatbeat interval, error(%s)", err)
		}

		inv, err := ln.GetHeatbeatInterval()
		log.Printf("heatbeat interval (%d), error(%v)\n", inv, err)

		for {
			conn, err := ln.Accept()

			//conn, err := ln.AcceptWithEvent(HandleSctpNotifyEventOnServer)
			if err != nil {
				log.Fatalf("failed to accept: %v", err)
			}
			log.Printf("Accepted Connection from RemoteAddr: %s", conn.RemoteAddr())
			// wconn := sctp.NewSCTPSndRcvInfoWrappedConn(conn.(*sctp.SCTPConn))
			// go serveClient(wconn)
			go serveClient_v1(conn.(*sctp.SCTPConn))
		}

	} else {
		conn, err := sctp.DialSCTP("sctp", laddr, raddr)
		if err != nil {
			log.Fatalf("failed to dial: %v", err)
		}
		log.Printf("Dail LocalAddr: %s; RemoteAddr: %s", conn.LocalAddr(), conn.RemoteAddr())
		streamid := uint16(0)
		for {
			if streamid > maxstreamid {
				streamid = 0
			} else {
				//streamid ++
			}

			info := &sctp.SndRcvInfo{
				Stream: uint16(streamid),
				PPID:   uint32(0),
			}

			conn.SubscribeEvents(sctp.SCTP_EVENT_ALL)
			conn.SetSctpEventHandler(conn, HandleSctpNotifyEventOnClient)
			n, err := conn.SCTPWrite([]byte("hello"), info)
			if err != nil {
				log.Fatalf("failed to write: %v", err)
			}
			log.Printf("write: %d", n)
			buf := make([]byte, 254)
			n, info, err = conn.SCTPRead(buf)
			if err != nil {
				log.Fatalf("failed to read: %v", err)
			}
			log.Printf("read: %d, msg: %s", n, string(buf[:n]))
			log.Printf("read: info: %+v", info)
			time.Sleep(time.Second)
		}
	}
}
