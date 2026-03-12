package main

import (
	"flag"
	"fmt"
	"port-scanner/internal/baseline"
	"port-scanner/internal/scanner"
)

func main() {
  // Flags
	ip := flag.String("ip", "", "Target IP address")
	networkCIDR := flag.String("network", "", "CIDR network to scan (example 192.168.1.0/24)")
	startPort := flag.Int("start", 1, "Start port")
	endPort := flag.Int("end", 1024, "End port")
	flag.Parse()

  // Baseline
	baselineFile := "data/baseline.json"
	baselineMap := baseline.LoadBaseline(baselineFile)

	if *networkCIDR != "" {
		scanner.ScanNetwork(*networkCIDR, *startPort, *endPort)
		return
	}

	if *ip == "" {
		fmt.Println("You must provide --ip or --network")
		return
	}

	openPorts := scanner.ScanHost(*ip, *startPort, *endPort)

  // Display open ports
	for _, port := range openPorts {
		if !baselineMap[port] {
			fmt.Println("⚠️ ALERT: New open port detected →", port)
		}
	}
  
  // Save data
	baseline.SaveBaseline(baselineFile, openPorts)
}
