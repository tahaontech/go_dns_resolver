package main

import (
	"fmt"
	"net"

	"github.com/tahaontech/go_dns_resolver/pkg/dns"
)

func main() {
	fmt.Printf("Starting DNS Server...\n")
	packetConnection, err := net.ListenPacket("udp", ":53")
	if err != nil {
		panic(err)
	}
	defer packetConnection.Close()

	for {
		buf := make([]byte, 512)
		bytesRead, addr, err := packetConnection.ReadFrom(buf)
		if err != nil {
			fmt.Printf("Read error from %s: %s\n", addr.String(), err)
			continue
		}

		go dns.HandlePacket(packetConnection, addr, buf[:bytesRead])
	}

}
