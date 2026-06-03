package configure

import (
	"fmt"
	"net"
)

var (
	LocalIps []string
)

func init() {
	addr, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("get local ip failed. err: %s", err))
	}
	for _, address := range addr {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// fmt.Println(ipnet.IP.String())
				LocalIps = append(LocalIps, ipnet.IP.String())
			}
		}
	}
}
