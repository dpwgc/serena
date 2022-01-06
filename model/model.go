package model

import "net"

// Node 消息队列节点结构体
type Node struct {
	Name string
	Addr net.IP
	Port string
}
