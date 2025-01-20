package network

import (
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func ICMP(pingCount int, host string) bool {
	// Ping initialization
	ping, err := probing.NewPinger(host)
	if err != nil {
		return false
	}
	ping.SetPrivileged(true)
	ping.Count = pingCount // Set send ping count
	ping.Timeout = 3 * time.Second

	// Run ping
	err = ping.Run()
	if err != nil {
		return false
	}

	// Get ping statistics
	stats := ping.Statistics()
	return stats.PacketsRecv > 0 // Return true if packets received
}
