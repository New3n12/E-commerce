package controlador

import (
	"net"
)

func GetWifiIPv4() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "error"
	}

	for _, iface := range interfaces {
		if iface.Name == "Wi-Fi" {
			addrs, err := iface.Addrs()
			if err != nil {
				return "error"
			}

			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						return ipNet.IP.String()
					}
				}
			}
		}
	}

	return "error"
}
