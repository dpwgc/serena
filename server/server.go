package server

import (
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

	//获取该节点的ip地址及Gossip服务端口号
	addr := viper.GetString("server.addr")
	//addr为空默认设为0.0.0.0
	if addr == "" {
		addr = "0.0.0.0"
	}
	gossipPort := viper.GetInt("registry.gossipPort")

	conf := memberlist.DefaultLANConfig()
	conf.Name = "[r]-" + addr + ":" + strconv.Itoa(gossipPort)
	conf.BindPort = gossipPort
	conf.AdvertisePort = gossipPort

	var err error

	//创建一个注册中心节点
	list, err = memberlist.Create(conf)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	//将注册中心节点加入到集群（创建一个集群）
	_, err = list.Join([]string{addr + ":" + strconv.Itoa(gossipPort)})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
}

// GetNode 获取除了注册中心之外的集群所有节点
func GetNode(c *gin.Context) {

	var nodeList []model.Node

	// 获取当前集群的节点（除去注册中心）
	for _, member := range list.Members() {
		//如果该节点是注册中心，跳过
		if strings.Split(member.Name, "-")[0] == "[r]" {
			continue
		}
		node := model.Node{
			Name: member.Name,
			Addr: member.Addr,
			Port: member.Port,
		}
		nodeList = append(nodeList, node)
	}

	c.JSON(0, gin.H{
		"code": 0,
		"data": nodeList,
	})
}
