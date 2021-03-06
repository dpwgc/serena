# Serena

## 基于Go + Gossip的KapokMQ集群服务发现与注册中心

![Go](https://img.shields.io/static/v1?label=LICENSE&message=Apache-2.0&color=orange)
![Go](https://img.shields.io/static/v1?label=Go&message=v1.17&color=blue)
[![github](https://img.shields.io/static/v1?label=Github&message=kapokmq&color=blue)](https://github.com/dpwgc/kapokmq)
[![star](https://gitee.com/dpwgc/kapokmq/badge/star.svg?theme=dark)](https://gitee.com/dpwgc/kapokmq/stargazers)
[![fork](https://gitee.com/dpwgc/kapokmq/badge/fork.svg?theme=dark)](https://gitee.com/dpwgc/kapokmq/members)

#### KapokMQ与Serena应用整合包下载

* https://github.com/dpwgc/kapokmq-server `github`

* https://gitee.com/dpwgc/kapokmq-server `gitee`

#### KapokMQ消息队列源码

* https://github.com/dpwgc/kapokmq `github`

* https://gitee.com/dpwgc/kapokmq `gitee`

***

### 实现功能

* 使用Gossip发现并同步消息队列服务节点信息。

* 消息队列生产者客户端可通过注册中心获取该集群所有消息队列服务节点列表，并与所有消息队列建立WebSocket连接，该生产者将随机选取一个消息队列节点投递消息。

***

### 开放接口

#### 获取所有消息队列节点信息

* 集群生产者通过此接口获取消息队列节点列表，进行负载均衡分配。

> http://127.0.0.1:8031/Registry/GetNodes

#### 请求方式
> POST

#### Content-Type
> form-data

#### 请求Header参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| secretKey     | test |  必填 | 安全访问密钥 |

#### 成功响应

```json
[
	{
		"Name": "mq:123.123.123.111:8011",
		"Addr": "123.123.123.111",
		"Port": "8011"
	},
	{
		"Name": "mq:123.123.123.222:8011",
		"Addr": "123.123.123.222",
		"Port": "8011"
	}
]
```

***

### 打包方法

##### 安装并配置go环境

##### 应用打包：

* 使用build.bat一键打包

```
Windows系统下直接运行./build.bat文件
自动执行打包命令
生成Linux二进制文件与Windows exe文件
```

* 打包成windows exe

```
  GoLand终端cd到项目根目录，执行go build命令，生成exe文件
```

* 打包成linux二进制文件

```
  cmd终端cd到项目根目录，依次执行下列命令：
  SET CGO_ENABLED=0
  SET GOOS=linux
  SET GOARCH=amd64
  go build
  生成二进制执行文件
```

***

### 部署方法

* 在服务器上部署

```
在Windows上部署

/serena                   # 文件根目录
    serena.exe            # 打包后的exe文件
    application.yaml      # 配置文件
    /log                  # 日志目录
```

```
在Linux上部署

/serena                   # 文件根目录
    serena                # 打包后的二进制文件(程序后台执行:setsid ./serena)
    application.yaml      # 配置文件
    /log                  # 日志目录
```

***

### 主要模块

##### 创建Gossip集群 `server/server.go`

* 使用 github.com/hashicorp/memberlist 创建一个Gossip集群的初始节点（注册中心节点），所有消息队列节点都要与注册中心节点相连。

* 对外给出一个GetNodes（http post接口），集群生产者将通过此接口获取消息队列节点列表。

***

### 配置文件

* ./application.yaml


