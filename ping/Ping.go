package ping

import (
	"fmt"

	"github.com/go-ping/ping"
)

// from given IP list(file) to get a IP String Array
func ReadIPs(file string) []string {

	var ips []string
	return ips
}

// from given IP
func Ping(ips []string) {
	pinger, err := ping.NewPinger(ips[0])
	if err != nil {
		panic(err)
	}

	// count = 10
	pinger.Count = 10

	// this is for Linux
	pinger.SetPrivileged(true)

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}

	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
	err = pinger.Run()
	if err != nil {
		panic(err)
	}
}

// Using for calculate average ping rtt and get the network performance
func AverageRTT() {

}
