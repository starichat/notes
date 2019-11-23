# docker 集群实践
> docker swarm 管理集群

## what is swarm
一个集群由多个docker主机组成，这些docker主机以集群模式运行，充当管理者和工作节点。

给定docker主机可以是管理接节点，也可以是工作节点。

## 使用
### 准备三台主机
docker swarm init --advertise-addr <ip>
初始化docker节点

docker swarm join ....
