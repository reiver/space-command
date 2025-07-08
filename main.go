package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("space-command ⚡")

	var beacondaemon <-chan error
	{
		// 239.83.80.67 (0xEF535043)
		var multicastIPAddress net.IP = net.IPv4(239, 'S', 'P', 'C')
		fmt.Printf("multicast ip-address: %v\n", multicastIPAddress)

		// 21328 (0x5350)
		var udpPort uint16 = (uint16('S') << 8) | uint16('P')
		fmt.Printf("UDP port: %v (0x%X)\n", udpPort, udpPort)

		var multicastUDPAddress = net.UDPAddr{
			IP: multicastIPAddress,
			Port: int(udpPort),
		}
		fmt.Printf("UDP address: %v\n", &multicastUDPAddress)

		beacondaemon = beaconserve(&multicastUDPAddress)
	}

	{
		var err error
		select {
		case err = <-beacondaemon:
			fmt.Printf("beacon-daemon lost: %s\n", err)
//		case err = <-httpdaemon:
//			log.Errorf("http-daemon lost: %s", err)
		}
		panic(err)
	}
}
