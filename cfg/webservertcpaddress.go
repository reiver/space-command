package cfg

import (
	"fmt"

	"github.com/reiver/space-command/env"
)

func WebServerTCPAddress() string {
	return fmt.Sprintf(":%s", env.TcpPort)
}
