package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/memberlist"
	"github.com/spf13/viper"
	"serena/model"
	"strconv"
	"strings"
)

//节点列表
var list *memberlist.Memberlist

// InitRegistry 初始化注册中心
func InitRegistry() {

	//获取注册中心的服务ip地址
	addr := viper.GetString("server.addr")
	//addr为空默认设为0.0.0.0
	if addr == "" {
		addr = "0.0.0.0"
	}
	//获取注册中心的Gin Http服务端口号
	port := viper.GetString("server.port")

	//获取注册中心的Gossip服务端口号
	gossipPort := viper.GetInt("registry.gossipPort")

	conf := memberlist.DefaultLANConfig()
	conf.Name = fmt.Sprintf("%s%s%s%s", "r-", addr, ":", port) //前缀r-表明这是注册中心，前缀mq-表明这是消息队列节点
	//注册中心的Gossip服务端口号
	conf.BindPort = gossipPort
	conf.AdvertisePort = gossipPort

	var err error

	//创建一个注册中心节点
	list, err = memberlist.Create(conf)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	//由注册中心来创建一个集群
	_, err = list.Join([]string{addr + ":" + strconv.Itoa(gossipPort)})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
}

// GetNodes 获取除了注册中心之外的集群所有节点
func GetNodes(c *gin.Context) {

	var nodes []model.Node

	// 获取当前集群的消息队列节点信息（除去注册中心）
	for _, member := range list.Members() {
		//如果该节点是注册中心，跳过
		if strings.Split(member.Name, "-")[0] == "r" {
			continue
		}
		node := model.Node{
			Name: member.Name,
			Addr: member.Addr,
			Port: strings.Split(member.Name, ":")[1],
		}
		nodes = append(nodes, node)
	}

	data, _ := json.Marshal(nodes)
	c.String(0, string(data))
}
