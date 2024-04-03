package util

import (
	"github.com/JustACP/AutoDDNS/logging"
	"net"
)

var LocalAddress map[string][]net.IP = make(map[string][]net.IP)

func GetLocalAddr() {
	interfaces, err := net.Interfaces()
	if err != nil {
		logging.Error("get network interface error %v", err)
	}

	for _, currInterface := range interfaces {
		currMAC := currInterface.HardwareAddr.String()
		addrs, err := currInterface.Addrs()
		if err != nil {
			logging.Info("get network interface mac: %s error %v", currMAC, err)
		}

		LocalAddress[currMAC] = make([]net.IP, len(addrs))
		for i, addr := range addrs {
			ip, _, _ := net.ParseCIDR(addr.String())
			LocalAddress[currMAC][i] = ip
		}
	}
}
