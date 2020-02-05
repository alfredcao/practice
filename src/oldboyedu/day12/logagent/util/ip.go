package util

import (
	"fmt"
	"net"
)

var LocalIps []string

func init() {
	err := GetLocalIp()
	if err != nil {
		fmt.Println("get local ip failed, err :", err)
	} else {
		fmt.Println("get local ip success, value :", LocalIps)
	}

}

func GetLocalIp() (err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			LocalIps = append(LocalIps, ipnet.IP.String())
		}
	}

	return
}
