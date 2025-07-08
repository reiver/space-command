package main

import (
	"net"

	"github.com/reiver/go-erorr"

	"github.com/reiver/space-command/srv/log"
)

const (
	errNilLocalAddr           = erorr.Error("nil local-addr")
	errNilMulticastUDPAddress = erorr.Error("nil multicast-udp-address")
)

func beaconserve(multicastUDPAddress *net.UDPAddr) <-chan error {
	log := logsrv.Prefix("beaconserve").Begin()
	defer log.End()

	ch := make(chan error)
	go _beaconserve(ch, multicastUDPAddress)
	log.Inform("beacon-daemon spawed 😈")
	return ch
}

func _beaconserve(ch chan error, multicastUDPAddress *net.UDPAddr) {
	log := logsrv.Prefix("_beaconserve").Begin()
	defer log.End()

	if nil == multicastUDPAddress {
		var err error = errNilMulticastUDPAddress

		ch <- err
		log.Error(err)
		return
	}

	udpConn, err := net.ListenMulticastUDP("udp", nil, multicastUDPAddress)
	if nil != err {
		ch <- err
		log.Errorf("ERROR: could not successfully listen to UDP address %v: %s", &multicastUDPAddress, err)
		return
	}
	defer udpConn.Close()
	log.Debug("Connected!")

	localAddr := udpConn.LocalAddr()
	if nil == localAddr {
		var err error = errNilLocalAddr

		ch <- err
		log.Errorf("ERROR: could not get UDP local-addr: %s", err)
		return
	}
	log.Debug("local-addr: ", localAddr)

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
				log.Errorf("ERROR: problem reading UDP package: %s", err)
				continue
			}
			log.Debugf("Read message of %d bytes long from %v", n, srcAddr)
			var p []byte = buffer[:n]

			log.Debugf("%q (%d)\n", p, len(p))
		}

	}
}
