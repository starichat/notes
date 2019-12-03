# docker 集群实践
> docker swarm 管理集群

## what is swarm
一个集群由多个docker主机组成，这些docker主机以集群模式运行，充当管理者和工作节点。

给定docker主机可以是管理接节点，也可以是工作节点。

## 使用
> 下面我们将学习到 docker-swarm 如何仅仅通过几条简单的指令就实现了docker集群的创建，部署应用，发布分布式服务。主要又1. 加入节点/管理/发现节点。
### 准备三台主机
docker swarm init --advertise-addr <ip>
初始化docker节点

docker swarm join ....


## 远程api
假设我们在192.168.1.123这台主机上开启了docker服务，监听了2375端口，那么我们就可以在同一网段的其他主机上（比如192.168.1.233）通过docker -H tcp://192.168.1.123:2345 <command>的方式调用到该主机上的docker服务。