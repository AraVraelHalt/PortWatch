package scanner

import (
	"fmt"
	"net"
	"port-scanner/internal/utils"
)

// Scan all hosts in a CIDR network
func ScanNetwork(cidr string, startPort, endPort int) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	
  if err != nil {
		fmt.Println("Invalid CIDR:", err)
		return
	}

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); utils.Inc(ip) {
		target := ip.String()
		fmt.Println("Scanning", target)
		ScanHost(target, startPort, endPort)
	}
}
