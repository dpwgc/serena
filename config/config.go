package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Conf struct {
	Server   Server   `yaml:"server"`
	Registry Registry `yaml:"registry"`
}
type Server struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`
}
type Registry struct {
	SecretKey  string `yaml:"secretKey"`
	GossipPort int    `yaml:"gossipPort"`
}

var Get Conf

// InitConfig 初始化项目配置
func InitConfig() {
	yamlFile, err := ioutil.ReadFile("application.yaml")
	if err != nil {
		panic(err)
	} // 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &Get)
	if err != nil {
		panic(err)
	}
}
