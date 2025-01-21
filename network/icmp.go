package network

import (
	"time"

	"CyberGuardian/logger"
	probing "github.com/prometheus-community/pro-bing"
)

func ICMP(szPingCount int, szPingTimeout int, host string) bool {
	// Ping initialization
	ping, err := probing.NewPinger(host)
	if err != nil {
		return false
	}
	ping.SetPrivileged(true)
	if szPingCount == 0 {
		logger.ERROR("Ping count is 0 !!!!!!")
	}
	ping.Count = szPingCount // Set send ping count
	ping.Timeout = time.Duration(szPingTimeout) * time.Second

	// Run ping
	err = ping.Run()
	if err != nil {
		return false
	}

	// Get ping statistics
	stats := ping.Statistics()
	if (stats.PacketsSent - stats.PacketsRecv) == stats.PacketsSent {
		return false
	} else {
		return true
	}
}
