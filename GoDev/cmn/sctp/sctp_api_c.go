package sctp

import "C"
import (
	"fmt"
	"unsafe"
)

//Code for set heatbeat

//#include <sys/socket.h>
//#include <netinet/sctp.h>
//#include <string.h>
//
//int cSetHeatbeat(int sock, int hbintv, int pathmaxrxt) {
//	struct sctp_paddrparams heartbeat;
//  memset(&heartbeat,  0, sizeof(struct sctp_paddrparams));
//	heartbeat.spp_flags = SPP_HB_ENABLE;
//	heartbeat.spp_hbinterval = hbintv;
//	heartbeat.spp_pathmaxrxt = pathmaxrxt;
//	if(setsockopt(sock, SOL_SCTP, SCTP_PEER_ADDR_PARAMS , &heartbeat, sizeof(heartbeat)) != 0){
//		return -1;
//	}
//  return 0;
//}
//
//int cAbortConnection(int fd)
//{
//  struct linger l_lger;
//  int ret;
//
//  if (fd <= 0) {
//    return -1;
//  }
//
//  memset(&l_lger, 0, sizeof(l_lger));
//  l_lger.l_onoff = 1;
//  l_lger.l_linger = 0;
//  ret = setsockopt(fd, SOL_SOCKET, SO_LINGER, &l_lger, sizeof(l_lger));
//  if (ret < 0) {
//   return -2 ;
//  }
//  close(fd);
//  return 0;
//}
import "C"

func setHbInterval(fd int, hbInv uint32, pathMaxRxt uint32) error {
	//param := &SctpPaddrParams{}
	//param.SppHbInterval = hbInv
	//param.SppFlags = SPP_HB_ENABLE
	//
	//paramlen := unsafe.Sizeof(*param)
	//_, _, err := setsockopt(fd, SCTP_PEER_ADDR_PARAMS, uintptr(unsafe.Pointer(param)), uintptr(paramlen))
	//
	//fmt.Printf("Set SctpPaddrParams, error(%v) len(%d), value %v \n", err, paramlen, param)

	if C.cSetHeatbeat(C.int(fd), C.int(hbInv), C.int(pathMaxRxt)) != 0 {
		return fmt.Errorf("failed to set heatbeat")
	}
	return nil
}

func getHbInterval(fd int) (uint32, error) {
	param := &SctpPaddrParams{}
	paramlen := unsafe.Sizeof(*param)

	_, _, err := getsockopt(fd, SCTP_PEER_ADDR_PARAMS, uintptr(unsafe.Pointer(param)), uintptr(unsafe.Pointer(&paramlen)))
	if err != nil {
		return 0, err
	}
	//fmt.Printf("Get SctpPaddrParams: error(%v),len(%d),value %v\n", err, paramlen, param)
	return param.SppHbInterval, err
}

func abortSctpConnection(fd int) error {
	rt := C.cAbortConnection(C.int(fd))
	if rt != 0 {
		return fmt.Errorf("failed to abort sctp connections.")
	}
	return nil
}
