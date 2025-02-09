package client

import "net"

type Client struct {
	Connection net.Conn
}

func FromConnection(conn net.Conn) *Client {
	return &Client{Connection: conn}
}
