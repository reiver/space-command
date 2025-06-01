package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("space-command ⚡")

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


	udpConn, err := net.ListenMulticastUDP("udp", nil,  &multicastUDPAddress)
	if nil != err {
		fmt.Printf("ERROR: could not successfully listen to UDP address %v: %s", &multicastUDPAddress, err)
		return
	}
	defer udpConn.Close()
	fmt.Println("Connected!")

	localAddr := udpConn.LocalAddr()
	if nil == localAddr {
		fmt.Printf("ERROR: could not get UDP local-address: %s\n", err)
		return
	}

	{
		// The max size of a UDP packet is 65,535 (= 2¹⁶ - 1).
		//
		// But, want the buffer to be a multiple of the system's cache line length
		// (as this often makes the operations faster) so rounding it up to 65536 (= 2^16).
		//
		// Although, could probably make this much smaller the "safe" UDP packet length is 508 bytes,
		var buffer [65536]byte

		for {
			n, srcAddr, err := udpConn.ReadFromUDP(buffer[:])
			if nil != err {
				fmt.Printf("ERROR: problem reading UDP package: %s\n", err)
				continue
			}
			fmt.Printf("Read message of %d bytes long from %v\n", n, srcAddr)
			var p []byte = buffer[:n]

			fmt.Printf("%q (%d)\n", p, len(p))
		}

	}
}
