package main

import (
	"net/http"

	"github.com/reiver/go-erorr"

	"github.com/reiver/space-command/srv/http"
	"github.com/reiver/space-command/srv/log"
)

func httpserve(tcpaddr string) <-chan error {
	log := logsrv.Prefix("httpserve").Begin()
	defer log.End()

	log.Informf("serving HTTP on TCP address: %q", tcpaddr)

	ch := make(chan error)
	go _httpserve(ch, tcpaddr)
	log.Inform("http-daemon spawed 😈")
	return ch
}

func _httpserve(ch chan error, tcpaddr string) {
	log := logsrv.Prefix("_httpserve").Begin()
	defer log.End()

	err := http.ListenAndServe(tcpaddr, &httpsrv.Mux)
	if nil != err {
		err = erorr.Errorf("problem with serving HTTP on TCP address %q: %w", tcpaddr, err)
		log.Errorf("ERROR: %s", err)
		ch <- err
	}
}
