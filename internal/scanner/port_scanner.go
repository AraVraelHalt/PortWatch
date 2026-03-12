package scanner

import (
	"fmt"
	"sync"
	"time"
	"net"
	"port-scanner/internal/services"
)

func ScanPort(ip string, port int) bool {
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 400*time.Millisecond)
	
  if err != nil {
		return false
	}
  
	conn.Close()
  
	return true
}

func ScanHost(ip string, startPort, endPort int) []int {
	var openPorts []int
	var wg sync.WaitGroup

	results := make(chan int, endPort-startPort+1) 
	semaphore := make(chan struct{}, 100)         // concurrency limit

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(p int) {
			defer wg.Done()
			if ScanPort(ip, p) {
				results <- p
			}
			<-semaphore
		}(port)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for port := range results {
		service := services.Services[port]
		
    if service != "" {
			fmt.Printf("%s → Open port %d (%s)\n", ip, port, service)
		} else {
			fmt.Printf("%s → Open port %d\n", ip, port)
		}
  
		openPorts = append(openPorts, port)
	}

	return openPorts
}
