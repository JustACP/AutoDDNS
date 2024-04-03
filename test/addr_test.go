package test

import (
	"fmt"
	"net"
	"testing"
)

func TestAddr(t *testing.T) {
	interfaces, _ := net.Interfaces()
	for _, v := range interfaces {

		fmt.Printf(v.HardwareAddr.String() + "\n")
		addrs, _ := v.Addrs()
		for _, addr := range addrs {

			fmt.Printf("\taddr: %s \n", addr.String())
			ip, _, _ := net.ParseCIDR(addr.String())
			fmt.Printf("\tv6: %s v4: %s \n", ip.To16(), ip.To4())
		}
	}

}
