package httpfunc

import (
	"fmt"
	"net"
)

func IsPortAvailable(host string, port int) bool {
	addr := fmt.Sprintf("%s:%d", host, port)
	if _, err := net.Listen("tcp", addr); err != nil {
		return false
	}
	return true
}
