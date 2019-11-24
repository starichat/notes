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
docker 把应用程序及其依赖，打包在 image 文件中。通过这个文件，就可以生成 docker 容器。image 文件可以看作是容器的静态模版，即容器是 image 文件的动态运行时。做系统级开发的可能对 image 比较了解，在刷系统镜像时候都是先编译生成一个系统 image 文件，将该文件刷入系统，然后通过 boot 引导读取该 image，启动系统，系统的运行时即是 image 的动态运行时。

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
docker是一个容器，要在微服务上大展身手，自然和网络分不开，docker有四种网络模式:

https://user-gold-cdn.xitu.io/2019/1/7/16828bdd24359cfe?imageslim

* bridge 模式
docker默认网络模式，此模式会为每个容器分配Network namespace，设置ip等，并和一个主机的docker容器连接到一个虚拟网桥上。当docker server启动时，会创建一个名为docker0的虚拟网桥，此主机上启动的docker容器会连接到这个虚拟网桥上，虚拟网桥的工作方式和物理交换机类似，这样主机上所有的容器就通过交换机连接在了一个二层网络中，接下来为容器分配ip，Docker会从RFC1918所定义的私有IP网段中，选择一个和宿主机不同的IP地址和子网分配给docker0，连接到docker0的容器就从这个子网中选择一个未占用的IP使用。如一般Docker会使用172.17.0.0/16这个网段，并将172.17.42.1/16分配给docker0网桥（在主机上使用ifconfig命令是可以看到docker0的，可以认为它是网桥的管理端口，在宿主机上作为一块虚拟网卡使用）。

https://user-gold-cdn.xitu.io/2019/1/7/16828bdd2287ee1c?imageView2/0/w/1280/h/960/format/webp/ignore-error/1

```
docker run --name b1 -it --network bridge --rm busybox:latest
/ # ifcinfig
sh: ifcinfig: not found
/ # ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:11:00:03  
          inet addr:172.17.0.3  Bcast:172.17.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:31 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:4485 (4.3 KiB)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

/ # 


```
再创建一个bridge的容器二，两者能够ping通
```
docker run --name b2 -it --network bridge --rm busybox:latest
Unable to find image 'busybox:latest' locally
latest: Pulling from library/busybox
0f8c40e1270f: Pull complete 
Digest: sha256:1303dbf110c57f3edf68d9f5a16c082ec06c4cf7604831669faf2c712260b5a0
Status: Downloaded newer image for busybox:latest
/ # ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:11:00:02  
          inet addr:172.17.0.2  Bcast:172.17.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:43 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:6053 (5.9 KiB)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

/ # 

```

注意，上面是使用的默认bridge，bridge还可以自定义bridge
可以使用如下命令进行自定义bridge
docker network create my_bridge
然后通过docker network ls 就可以看到所有的docker network.
```
 docker network ls
NETWORK ID          NAME                     DRIVER              SCOPE
14c72228792d        bridge                   bridge              local
7bb73a988fcf        docker_gwbridge          bridge              local
72ad7e46c6a5        go_simpleweibo_default   bridge              local
08d9e5f6c764        hello_default            bridge              local
e2ff4b881875        host                     host                local
we7vwdicha1b        ingress                  overlay             swarm
34767993fe9f        my_bri                   bridge              local
950a42fa6823        my_bridge                bridge              local
ccd4c6d929e2        none                     null                local
5ef4b885534c        starichat_default        bridge              local

```

然后就可以创建使用自定义 bridge 的容器，自定义 bridge 容器可以使用容器名进行 dns 解析。
docker run --name b3 -it --network my_bridge --rm busybox:latest
docker run --name b4 -it --network my_bridge --rm busybox:latest

b3 可以通过 b4 的名字找到 b4 的ip。但是使用默认网桥docker0就不行了。同时自定义网桥和默认网桥不互通，当然自定义网桥之间耶不互通，只有在同一网桥下的容器才互通，但是可以通过添加网桥，使不同网桥的容器可以互通。

```
docker run --name b3 -it --network my_bridge --rm busybox:latest
/ # ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:17:00:02  
          inet addr:172.23.0.2  Bcast:172.23.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:74 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:10836 (10.5 KiB)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

```

```
/ # ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:17:00:03  
          inet addr:172.23.0.3  Bcast:172.23.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:44 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:6327 (6.1 KiB)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

/ # 

```
在 b3 的容器内直接 ping b4 的容器名可以直接 ping 通,但是默认的 bridge 模式的 docker0 网桥就不行,因为默认的网桥不会使用dns解析.

不通网桥之间,网络是不互通的,这时候 

docker network connect my_bridge b1 将 b1 加入到 my_bridge 网桥上,这样,b1 就可以和 b3,b4 ping通了

我们看下 b1 的网卡
```
# ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:11:00:03  
          inet addr:172.17.0.3  Bcast:172.17.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:147 errors:0 dropped:0 overruns:0 frame:0
          TX packets:22 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:22665 (22.1 KiB)  TX bytes:2044 (1.9 KiB)

eth1      Link encap:Ethernet  HWaddr 02:42:AC:17:00:04  
          inet addr:172.23.0.4  Bcast:172.23.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:52 errors:0 dropped:0 overruns:0 frame:0
          TX packets:9 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:7203 (7.0 KiB)  TX bytes:770 (770.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

```
我们发现 b1 有两个网卡,eth1 显然是在my_bridge上新增了一个节点,当和b3通信的时候,就使用这个节点进行通信,当和b2通信,使用eth0网桥.


* host 模式
如果启动容器的时候使用host模式，那么这个容器将不会获得一个独立的Network Namespace，而是和宿主机共用一个Network Namespace。容器将不会虚拟出自己的网卡，配置自己的IP等，而是使用宿主机的IP和端口。使用host模式启动容器后可以发现，使用ip addr查看网络环境时，看到的都是宿主机上的信息。这种方式创建出来的容器，可以看到host上的所有网络设备。一般不推荐使用。
* container 模式
这个模式指定新创建的容器和已经存在的一个容器共享一个Network Namespace，而不是和宿主机共享。新创建的容器不会创建自己的网卡，配置自己的IP，而是和一个指定的容器共享IP、端口范围等。同样，两个容器除了网络方面，其他的如文件系统、进程列表等还是隔离的。两个容器的进程可以通过lo网卡设备通信。

```
 docker run --name b5 -it --net=container:b3 --rm busybox:latest
/ # ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:17:00:02  
          inet addr:172.23.0.2  Bcast:172.23.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:430 errors:0 dropped:0 overruns:0 frame:0
          TX packets:18 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:79851 (77.9 KiB)  TX bytes:1398 (1.3 KiB)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:56 errors:0 dropped:0 overruns:0 frame:0
          TX packets:56 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:3742 (3.6 KiB)  TX bytes:3742 (3.6 KiB)

/ # 

```
* none 模式
该模式将容器放置在自己的网络栈中，但是并不进行任何配置。实际上，该模式关闭了容器的网络功能，不分配任何网卡。

## dockerfile
以上，基本列出了docker的基本使用，但是我们发现使用docker使用以上规范时，比较麻烦，在一行命令行要输入好多指令，所以，以下我们学习dockfile创建容器，dockerfile的内容不多，主要内容如下：
```
FROM：指定基础镜像，必须为第一个命令
MAINTAINER: 维护者信息
RUN：构建镜像时执行的命令
ADD：将本地文件添加到容器中，tar类型文件会自动解压(网络压缩资源不会被解压)，可以访问网络资源，类似wget
COPY：功能类似ADD，但是是不会自动解压文件，也不能访问网络资源
CMD：构建容器后调用，也就是在容器启动时才进行调用。
ENTRYPOINT：配置容器，使其可执行化。配合CMD可省去"application"，只使用参数。
LABEL：用于为镜像添加元数据.
ENV：设置环境变量
EXPOSE：指定于外界交互的端口
VOLUME：用于指定持久化目录
WORKDIR：工作目录，类似于cd命令
USER:指定运行容器时的用户名或 UID，后续的 RUN 也会使用指定用户。使用USER指定用户时，可以使用用户名、UID或GID，或是两者的组合。当服务不需要管理员权限时，可以通过该命令指定运行用户。并且可以在之前创建所需要的用户
ARG：用于指定传递给构建运行时的变量
ONBUILD：用于设置镜像触发器
```

### example:

## docker-compose
既然docker涉及到的是容器，自然会和很多容器打交道，为每一个容器写一套dockerfile，分别运行，固然可行，但是对于有依赖的容器，就显得很复杂了。所有对多容器环境，选择使用docker-compose
### docker-compose.yml 语法
```
image: 指定为镜像名称或镜像 ID。如果镜像在本地不存在，Compose 将会尝试拉去这个镜像。
build: 指定 Dockerfile 所在文件夹的路径。 Compose 将会利用它自动构建这个镜像，然后使用这个镜像。
command: 覆盖容器启动后默认执行的命令。
links: 链接到其它服务中的容器。使用服务名称（同时作为别名）或服务名称：服务别名 （[SERVICE:ALIAS](service:ALIAS)） 格式都可以。
external_links: 链接到 docker-compose.yml 外部的容器，甚至 并非 Compose 管理的容器。参数格式跟 links 类似。
ports: 暴露端口信息。使用宿主：容器 （HOST:CONTAINER）格式或者仅仅指定容器的端口（宿主将会随机选择端口）都可以。
expose: 暴露端口，但不映射到宿主机，只被连接的服务访问。仅可以指定内部端口为参数
volumes: 卷挂载路径设置。可以设置宿主机路径 （HOST:CONTAINER） 或加上访问模式 （HOST:CONTAINER:ro）。
volumes_from: 从另一个服务或容器挂载它的所有卷。
environment: 设置环境变量。你可以使用数组或字典两种格式。只给定名称的变量会自动获取它在 Compose 主机上的值，可以用来防止泄露不必要的数据。
env_file: 从文件中获取环境变量，可以为单独的文件路径或列表。如果通过 docker-compose -f FILE 指定了模板文件，则 env_file 中路径会基于模板文件路径。如果有变量名称与 environment 指令冲突，则以后者为准。
extends: 基于已有的服务进行扩展。例如我们已经有了一个 webapp 服务，模板文件为 common.yml。
# common.yml
webapp:
build: ./webapp
environment:
 - DEBUG=false
 - SEND_EMAILS=false
 编写一个新的 development.yml 文件，使用 common.yml 中的 webapp 服务进行扩展。
 # development.yml
web:
extends:
file: common.yml
service: webapp
ports:
 - "8000:8000"
links:
 - db
environment:
- DEBUG=true
db:
image: postgres

后者会自动继承 common.yml 中的 webapp 服务及相关环节变量。

net: 设置网络模式。使用和 docker client 的 --net 参数一样的值。

pid: 跟主机系统共享进程命名空间。打开该选项的容器可以相互通过进程 ID 来访问和操作。

dns: 配置 DNS 服务器。可以是一个值，也可以是一个列表。

cap_add, cap_drop: 添加或放弃容器的 Linux 能力（Capabiliity）。

dns_search: 配置 DNS 搜索域。可以是一个值，也可以是一个列表。
```

#### 一个简单的example

### 利用docker-compose 搭建一个集mysql、redis、go-web的简易多容器应用
