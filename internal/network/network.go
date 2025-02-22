package network

import (
	"fmt"
	"net"
	"time"
)

func IsResovable(address string, port int) bool {
	timeout := time.Second * 2
	addr := fmt.Sprintf("%s:%d", address, port)

	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
