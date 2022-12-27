package util

import (
	"net"
)

func GetIPAddress() string {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		return addr.String()
	}
	return addrs[0].String()
}
