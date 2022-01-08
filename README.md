# Serena

## 基于Go + Gossip的KapokMQ集群服务发现与注册中心

### KapokMQ与Serena应用整合包下载

* https://github.com/dpwgc/kapokmq-server

* https://gitee.com/dpwgc/kapokmq-server

### KapokMQ消息队列源码

* https://github.com/dpwgc/kapokmq

* https://gitee.com/dpwgc/kapokmq

`Golang` `Gossip` `Registry`

***

### 配置说明

* config/application.yaml

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

* 消息队列生产者客户端可通过注册中心获取该集群所有消息队列服务节点列表，并与所有消息队列建立WebSocket连接，该生产者将随机选取一个消息队列节点投递消息。

#### 获取所有消息队列节点信息 - 接口URL
> http://127.0.0.1:8031/Registry/GetNode

#### 请求方式
> POST

#### Content-Type
> form-data

#### 请求Header参数

| 参数        | 示例值   | 是否必填   |  参数描述  |
| :--------   | :-----  | :-----  | :----  |
| secretKey     | test |  必填 | 安全访问密钥 |

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
    serena                # 打包后的二进制文件(程序后台执行:setsid ./serena)
    /config               # 配置目录
        application.yaml  # 配置文件
    /log                  # 日志目录
```

