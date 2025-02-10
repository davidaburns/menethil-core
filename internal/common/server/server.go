package server

import "net"

type Server interface {
	Start()
	Stop()
	Update()
	HandleConnection(con net.Conn)
}
