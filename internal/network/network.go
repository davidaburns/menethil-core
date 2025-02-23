package network

import (
	"errors"
	"fmt"
	"net"
)

type IP struct {
	net.IP
}

func (ip *IP) Scan(src any) error {
	switch v := src.(type) {
	case []byte:
		parsed := net.ParseIP(string(v))
		if parsed == nil {
			return errors.New("invalid IP address")
		}

		ip.IP = parsed
	case string:
		parsed := net.ParseIP(string(v))
		if parsed == nil {
			return errors.New("invalid IP address")
		}

		ip.IP = parsed
	default:
		return fmt.Errorf("cannot scan type %T into a valid IP address", src)
	}

	return nil
}

func (ip *IP) IsResovable() bool {
	ips, err := net.LookupIP(ip.String())
	if err != nil || len(ips) == 0 {
		return false
	}

	return true
}
