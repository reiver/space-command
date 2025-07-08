package env

import (
	"os"
)

var TcpPort string = tcpPort()

// tcpPort returns the TCP-port that the HTTP server should use.
//
// It defaults to TCP-port 8080.
//
// But that can be overridden by a value set in the "PORT" environment variable.
func tcpPort() string {
	value := os.Getenv("PORT")
	if "" == value {
		value = "80"
	}

	return value
}
