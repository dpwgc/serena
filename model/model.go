package model

// Node 消息队列节点结构体
type Node struct {
	Name string //消息队列节点名称
	Addr string //消息队列节点地址
	Port string //消息队列节点http服务端口
}
