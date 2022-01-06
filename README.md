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
  port: 8030

registry:
  # 安全密钥
  secretKey: test
  # Gossip服务运行端口号
  gossipPort: 8020
```

***

### 实现功能

* 发现消息队列服务节点

* 获取所有消息队列服务节点列表