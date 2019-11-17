# docker 使用指南
> Docker 属于 Linux 容器的一种封装，提供简单易用的容器使用接口。它是目前最流行的 Linux 容器解决方案。

Docker 将应用程序与该程序的依赖，打包在一个文件里面。运行这个文件，就会生成一个虚拟容器。程序在这个虚拟容器里运行，就好像在真实的物理机上运行一样。有了 Docker，就不用担心环境问题。

总体来说，Docker 的接口相当简单，用户可以方便地创建和使用容器，把自己的应用放入容器。容器还可以进行版本管理、复制、分享、修改，就像管理普通的代码一样。 [参考阮一峰博客](https://www.ruanyifeng.com/blog/2018/02/docker-tutorial.html)

## 简介
docker 用途：
1. 提供一次性的环境
2. 提供弹性的云服务
3. 组建微服务架构，在本机即可模拟出微服务环境。

## 使用
docker 的安装比较简单，官方文档都给出了详细的步骤，[安装步骤](),这里就不做详细赘述了，直接默认安装好了docker环境，我使用的mac电脑，其他系统都是类似的。

### image 文件
docker 把应用程序及其依赖，打包在image文件中。通过这个文件，就可以生成docker容器。image文件可以看作是容器的静态模版，即容器是image文件的动态运行时。做系统级开发的可能对image比较了解，在刷系统镜像时候都是先编译生成一个系统image文件，将该文件刷入系统，然后通过boot引导读取该image，启动系统，系统的运行时即是image的动态运行时。

1. docker images 列出本机下载的镜像

2. docker search mysql 搜索镜像

3. docker pull mysql:lastest 拉取镜像

### docker 容器
上一节我们已经知道了容器，即image文件生成的容器实例，本身也是一个文件，探究计算机原理，所有的东西都是文件，甚至是二进制。容器一旦生成，就会同时存在于两个文件：image文件和容器文件。而且关闭容器文件并不会删除容器文件，只是容器停止运行了而已。
1. docker container ls 列出所有正在运行的容器

2. docker container ls -all 列出所有饿容器，包括终止运行的容器

3. docker container rm [containerId] 删除容器

4. docker run 运行容器
以下列出常见操作：
```
  -d, --detach=false         指定容器运行于前台还是后台，默认为false     
  -i, --interactive=false   打开STDIN，用于控制台交互    
  -t, --tty=false            分配tty设备，该可以支持终端登录，默认为false    
  -u, --user=""              指定容器的用户    
  -a, --attach=[]            登录容器（必须是以docker run -d启动的容器）  
  -w, --workdir=""           指定容器的工作目录   
  -c, --cpu-shares=0        设置容器CPU权重，在CPU共享场景使用    
  -e, --env=[]               指定环境变量，容器中可以使用该环境变量    
  -m, --memory=""            指定容器的内存上限    
  -P, --publish-all=false    指定容器暴露的端口    
  -p, --publish=[]           指定容器暴露的端口   
  -h, --hostname=""          指定容器的主机名    
  -v, --volume=[]            给容器挂载存储卷，挂载到容器的某个目录    
  --volumes-from=[]          给容器挂载其他容器上的卷，挂载到容器的某个目录  
```
[code] docker run -name mybusybox -it -v .... -d busybox:lastest

5. docker ps 查看容器

6. docker top containerName 查看容器内进程

7. docker stop containerName/containerId

8. docker restart

9. docker exec/attach 进入容器，exec 退出容器，不会导致容器停止

10. docker export dockerID > mydoker.tar 导出容器

## docker 网络模式
docker是一个容器，要在微服务上大展身手，自然和网络分不开，docker有四种网络迷失
* bridge 模式
// TODO
docker默认网络模式，此模式会为每个容器分配Network namespace，设置ip等，并和一个主机的docker容器连接到一个虚拟网桥上。当docker server启动时，会创建一个名为docker0的虚拟网桥，此主机上启动的docker容器会连接到这个虚拟网桥上，虚拟网桥的工作方式和物理交换机类似，这样主机上所有的容器就通过交换机连接在了一个二层网络中，接下来为容器分配ip，Docker会从RFC1918所定义的私有IP网段中，选择一个和宿主机不同的IP地址和子网分配给docker0，连接到docker0的容器就从这个子网中选择一个未占用的IP使用。如一般Docker会使用172.17.0.0/16这个网段，并将172.17.42.1/16分配给docker0网桥（在主机上使用ifconfig命令是可以看到docker0的，可以认为它是网桥的管理端口，在宿主机上作为一块虚拟网卡使用）。

```
docker run --name b1 -it --network bridge --rm busybox:latest // 创建一个bridge模式的docker容器
/ # ifconfig // 查看该容器的网络配置
eth0      Link encap:Ethernet  HWaddr 02:42:AC:11:00:02  
          inet addr:172.17.0.2  Bcast:172.17.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:63 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:8100 (7.9 KiB)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

/ # route -n // 查看其路由
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         172.17.0.1      0.0.0.0         UG    0      0        0 eth0
172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 eth0
/ # ping 192.168.164.168 //能否和宿主机ping通
PING 192.168.164.168 (192.168.164.168): 56 data bytes
64 bytes from 192.168.164.168: seq=0 ttl=63 time=0.481 ms
64 bytes from 192.168.164.168: seq=1 ttl=63 time=0.567 ms
64 bytes from 192.168.164.168: seq=2 ttl=63 time=0.547 ms
64 bytes from 192.168.164.168: seq=3 ttl=63 time=0.614 ms
^C
--- 192.168.164.168 ping statistics ---
4 packets transmitted, 4 packets received, 0% packet loss
round-trip min/avg/max = 0.481/0.552/0.614 ms
/ # 

```
再创建一个bridge的容器二，两者能够ping通
// TODO
注意，上面是使用的默认bridge，bridge还可以自定义bridge
可以使用如下命令进行自定义bridge
docker network create my_bridge

然后可以创建使用自定义bridge的容器，自定义bridge容器可以使用容器名进行dns解析。
docker run --name b1 -it --network my_bridge --rm busybox:latest
docker run --name b2 -it --network my_bridge --rm busybox:latest

b1 可以通过b2的名字找到b2的ip。但是使用默认网桥docker0就不行了。同时自定义网桥和默认网桥不互通，当然自定义网桥之间耶不互通，只有在同一网桥下的容器才互通，但是可以通过添加网桥，使不同网桥的容器可以互通。
// TODO
* host 模式
如果启动容器的时候使用host模式，那么这个容器将不会获得一个独立的Network Namespace，而是和宿主机共用一个Network Namespace。容器将不会虚拟出自己的网卡，配置自己的IP等，而是使用宿主机的IP和端口。使用host模式启动容器后可以发现，使用ip addr查看网络环境时，看到的都是宿主机上的信息。这种方式创建出来的容器，可以看到host上的所有网络设备。一般不推荐使用。
* container 模式
这个模式指定新创建的容器和已经存在的一个容器共享一个Network Namespace，而不是和宿主机共享。新创建的容器不会创建自己的网卡，配置自己的IP，而是和一个指定的容器共享IP、端口范围等。同样，两个容器除了网络方面，其他的如文件系统、进程列表等还是隔离的。两个容器的进程可以通过lo网卡设备通信。

```
// TODO
```
* none 模式
该模式将容器放置在自己的网络栈中，但是并不进行任何配置。实际上，该模式关闭了容器的网络功能，不分配任何网卡。

## dockerfile
以上，基本列出了docker的基本使用，但是我们发现使用docker使用以上规范时，比较麻烦，在一行命令行要输入好多指令，所以，以下我们学习dockfile创建容器，dockerfile的内容不多，主要内容如下：
```
// TODO
```

## docker-compose
既然docker涉及到的是容器，自然会和很多容器打交道，为每一个容器写一套dockerfile，分别运行，固然可行，但是对于有依赖的容器，就显得很复杂了。所有对多容器环境，选择使用docker-compose
### docker-compose.yml 语法
### 利用docker-compose 搭建一个集mysql、redis、go-web的简易多容器应用
