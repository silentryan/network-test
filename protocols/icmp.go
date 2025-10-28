package protocols

// This package will provide:
// 1. ping accessible
// 2. ping average latency
// 3. traceroute: route path checking

import (
	"log"

	"github.com/prometheus-community/pro-bing"
)

// A host could be a IP or url without scheme

func Ping(host string) {
	pinger, err := probing.NewPinger(host)
	if err != nil {
		log.Fatal(err)
	}
	pinger.Run()
}
