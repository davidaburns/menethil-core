package server

import "net"

type Server interface {
	Start()
	Stop()
	HandleConnection(con net.Conn)
}
