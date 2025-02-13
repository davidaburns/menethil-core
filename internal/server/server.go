package server

import "net"

type Server interface {
	Start() error
	Stop()
	HandleConnection(con net.Conn)
}
