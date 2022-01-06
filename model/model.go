package model

import "net"

type Node struct {
	Name string
	Addr net.IP
	Port uint16
}
