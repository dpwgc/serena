package server

import (
	"fmt"
	"github.com/hashicorp/memberlist"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var list *memberlist.Memberlist

func InitRegistry() {

	//获取该节点的ip地址及端口号
	addr := viper.GetString("server.addr")
	if addr == "" {
		addr = "0.0.0.0"
	}
	port := viper.GetInt("registry.port")

	conf := memberlist.DefaultLANConfig()
	conf.Name = addr + ":" + strconv.Itoa(port)
	conf.BindPort = port
	conf.AdvertisePort = port

	var err error

	//创建一个注册中心节点
	list, err = memberlist.Create(conf)
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}

	//将注册中心节点加入到集群（创建一个集群）
	_, err = list.Join([]string{addr + ":" + strconv.Itoa(port)})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}

	discovery()
}

func discovery() {

	discoveryCycle := viper.GetInt("registry.discoveryCycle")

	for {
		// 获取当前集群的节点
		for _, member := range list.Members() {

			fmt.Printf("Member: %s %s %s\n", member.Name, member.Addr, strconv.Itoa(list.NumMembers()))
		}
		time.Sleep(time.Second * time.Duration(discoveryCycle))
	}
}
