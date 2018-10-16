balance:负载均衡库
=================================
balance.go: blance interface
instance.go: instance实体
mgr.go: 负载均衡管理功能,注册和发现
random.go: 随机balancer
roundrobin.go: 轮训balancer



main:调用方
=================================
hash.go:用户定义balancer
main.go: