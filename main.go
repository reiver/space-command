package main

import (
	"net"

	"github.com/reiver/space-command/cfg"
	"github.com/reiver/space-command/srv/log"
)

func main() {
	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Debug("space-command ⚡")

	var beacondaemon <-chan error
	{
		// 239.83.80.67 (0xEF535043)
		var multicastIPAddress net.IP = net.IPv4(239, 'S', 'P', 'C')
		log.Debug("(beacon) multicast ip-address: ", multicastIPAddress)

		// 21328 (0x5350)
		var udpPort uint16 = (uint16('S') << 8) | uint16('P')
		log.Debugf("(beacon) UDP port: %v (0x%X)", udpPort, udpPort)

		var multicastUDPAddress = net.UDPAddr{
			IP: multicastIPAddress,
			Port: int(udpPort),
		}
		log.Debug("(beacon) UDP address: ", &multicastUDPAddress)

		beacondaemon = beaconserve(&multicastUDPAddress)
	}

	var httpdaemon <-chan error
	{
		var httptcpaddr string = cfg.WebServerTCPAddress()
		log.Debug("(http) TCP address: ", httptcpaddr)

		httpdaemon = httpserve(httptcpaddr)
	}

	{
		var err error
		select {
		case err = <-beacondaemon:
			log.Error("beacon-daemon lost: ", err)
		case err = <-httpdaemon:
			log.Errorf("http-daemon lost: %s", err)
		}
		panic(err)
	}
}
