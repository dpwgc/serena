# Serena

## 基于Go + Gossip的KapokMQ集群服务发现与注册中心

### KapokMQ消息队列：

* https://github.com/dpwgc/kapokmq

* https://gitee.com/dpwgc/kapokmq

`Golang` `Gossip` `Registry`

***

### 配置说明

```yaml
server:
  # ip地址/域名
  addr: 0.0.0.0
  # Gin服务端口号
  port: 8031

registry:
  # 安全密钥
  secretKey: test
  # Gossip服务运行端口号
  gossipPort: 8041
```

***

### 实现功能

* 使用Gossip发现并同步消息队列服务节点信息。

* 消息队列生产者客户端可通过注册中心获取该集群所有消息队列服务节点列表，并与所有消息队列建立WebSocket连接，进行负载均衡分配。

***

### 打包方式

* 填写application.yaml内的配置。

* 运行项目：

```
（1）GoLand直接运行main.go(调试)
```

```
（2）打包成exe运行(windows部署)

  GoLand终端cd到项目根目录，执行go build命令，生成exe文件
```

```
（3）打包成二进制文件运行(linux部署)

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
    /config               # 配置目录
        application.yaml  # 配置文件
    /log                  # 日志目录
```

```
在Linux上部署

/serena                   # 文件根目录
    serena                # 打包后的二进制文件(程序后台执行:setsid ./KapokMQ)
    /config               # 配置目录
        application.yaml  # 配置文件
    /log                  # 日志目录
```

