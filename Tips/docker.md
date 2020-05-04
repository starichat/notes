# docker 入门指北
> 基本盖链

## 虚拟机和docker区别

##　基本操作
### 安装
### busybox　示例docker各种操作
1.　容器声明周期管理

2.　容器操作

3.　容器ｒｏｏｆｔｓ命令

4.　镜像相关
## docker search 搜索镜像
｀｀｀
docker search busybox

NAME                      DESCRIPTION                                     STARS               OFFICIAL            AUTOMATED
busybox                   Busybox base image.                             1722                [OK]                
progrium/busybox                                                          71                                      [OK]

｀｀｀

## docker　pull　获取镜像
docker pull [name]:[version]
```
wangxinlong@wangxinlong:~$ docker pull busybox
Using default tag: latest
latest: Pulling from library/busybox
0f8c40e1270f: Pull complete 
Digest: sha256:1303dbf110c57f3edf68d9f5a16c082ec06c4cf7604831669faf2c712260b5a0
Status: Downloaded newer image for busybox:latest
docker.io/library/busybox:latest
```
## docker images　查看镜像相关信息

docker images
```
busybox             latest              020584afccce        2 weeks ago         1.22MB

```

## docker run 运行容器

## docker ps 查看容器

## 查看容器内进程

## 查看容器内进程

docker top containerName

## 守护态运行
docker run --name webserver -d -P nginx

参数解释
｀｀｀
runoob@runoob:~$ docker run -it nginx:latest /bin/bash  
root@b8573233d675:/#   



Usage: docker run [OPTIONS] IMAGE [COMMAND] [ARG...]    
02.  
03.  -d, --detach=false         指定容器运行于前台还是后台，默认为false     
04.  -i, --interactive=false   打开STDIN，用于控制台交互    
05.  -t, --tty=false            分配tty设备，该可以支持终端登录，默认为false    
06.  -u, --user=""              指定容器的用户    
07.  -a, --attach=[]            登录容器（必须是以docker run -d启动的容器）  
08.  -w, --workdir=""           指定容器的工作目录   
09.  -c, --cpu-shares=0        设置容器CPU权重，在CPU共享场景使用    
10.  -e, --env=[]               指定环境变量，容器中可以使用该环境变量    
11.  -m, --memory=""            指定容器的内存上限    
12.  -P, --publish-all=false    指定容器暴露的端口    
13.  -p, --publish=[]           指定容器暴露的端口   
14.  -h, --hostname=""          指定容器的主机名    
15.  -v, --volume=[]            给容器挂载存储卷，挂载到容器的某个目录    
16.  --volumes-from=[]          给容器挂载其他容器上的卷，挂载到容器的某个目录  
17.  --cap-add=[]               添加权限，权限清单详见：http://linux.die.net/man/7/capabilities    
18.  --cap-drop=[]              删除权限，权限清单详见：http://linux.die.net/man/7/capabilities    
19.  --cidfile=""               运行容器后，在指定文件中写入容器PID值，一种典型的监控系统用法    
20.  --cpuset=""                设置容器可以使用哪些CPU，此参数可以用来容器独占CPU    
21.  --device=[]                添加主机设备给容器，相当于设备直通    
22.  --dns=[]                   指定容器的dns服务器    
23.  --dns-search=[]            指定容器的dns搜索域名，写入到容器的/etc/resolv.conf文件    
24.  --entrypoint=""            覆盖image的入口点    
25.  --env-file=[]              指定环境变量文件，文件格式为每行一个环境变量    
26.  --expose=[]                指定容器暴露的端口，即修改镜像的暴露端口    
27.  --link=[]                  指定容器间的关联，使用其他容器的IP、env等信息    
28.  --lxc-conf=[]              指定容器的配置文件，只有在指定--exec-driver=lxc时使用    
29.  --name=""                  指定容器名字，后续可以通过名字进行容器管理，links特性需要使用名字    
30.  --net="bridge"             容器网络设置:  
31.                                bridge 使用docker daemon指定的网桥       
32.                                host    //容器使用主机的网络    
33.                                container:NAME_or_ID  >//使用其他容器的网路，共享IP和PORT等网络资源    
34.                                none 容器使用自己的网络（类似--net=bridge），但是不进行配置   
35.  --privileged=false         指定容器是否为特权容器，特权容器拥有所有的capabilities    
36.  --restart="no"             指定容器停止后的重启策略:  
37.                                no：容器退出时不重启    
38.                                on-failure：容器故障退出（返回值非零）时重启   
39.                                always：容器退出时总是重启    
40.  --rm=false                 指定容器停止后自动删除容器(不支持以docker run -d启动的容器)    
41.  --sig-proxy=true           设置由代理接受并处理信号，但是SIGCHLD、SIGSTOP和SIGKILL不能被代理    
｀｀｀

## 删除容器
docker rm [-f] containerName

## docker　网络模式
docker　有四种网络模式
host

bridge

Container

None

## bridge 模式,访问外部网络通过ｄｏｃｋｅｒ０转发
｀｀｀
wangxinlong@wangxinlong:~$ docker run --name b3 -it --rm busybox:latest 
/ # ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:11:00:03  
          inet addr:172.17.0.3  Bcast:172.17.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:24 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:2931 (2.8 KiB)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

/ # netstat -nutl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       
/ # echo "hello world b1" > /tmp/index.html
/ # httpd -h /tmp/ 
/ # netstat -nutl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       
tcp        0      0 :::80                   :::*                    LISTEN      
/ # ls
bin   dev   etc   home  proc  root  sys   tmp   usr   var
/ # route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         172.17.0.1      0.0.0.0         UG    0      0        0 eth0
172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 eth0
/ # ping 172.17.0.2
PING 172.17.0.2 (172.17.0.2): 56 data bytes
64 bytes from 172.17.0.2: seq=0 ttl=64 time=0.057 ms
64 bytes from 172.17.0.2: seq=1 ttl=64 time=0.042 ms
64 bytes from 172.17.0.2: seq=2 ttl=64 time=0.042 ms
64 bytes from 172.17.0.2: seq=3 ttl=64 time=0.038 ms
^C
--- 172.17.0.2 ping statistics ---
4 packets transmitted, 4 packets received, 0% packet loss
round-trip min/avg/max = 0.038/0.044/0.057 ms
/ # traceroute 172.17.0.2
traceroute to 172.17.0.2 (172.17.0.2), 30 hops max, 46 byte packets
 1  172.17.0.2 (172.17.0.2)  0.004 ms  0.021 ms  0.017 ms
/ # traceroute 192.168.164.168
traceroute to 192.168.164.168 (192.168.164.168), 30 hops max, 46 byte packets
 1  172.17.0.1 (172.17.0.1)  0.017 ms  0.017 ms  0.011 ms
 2  192.168.164.168 (192.168.164.168)  0.432 ms  0.554 ms  0.464 ms
/ # 


｀｀｀
｀｀｀
wangxinlong@wangxinlong:~$ ｀｀｀

host:
```
wangxinlong@wangxinlong:~$ docker run --name b2 -it --network host --rm busybox:latest
/ # ifconfig
br-08d9e5f6c764 Link encap:Ethernet  HWaddr 02:42:23:3D:29:3A  
          inet addr:172.20.0.1  Bcast:172.20.255.255  Mask:255.255.0.0
          inet6 addr: fe80::42:23ff:fe3d:293a/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1 errors:0 dropped:0 overruns:0 frame:0
          TX packets:170 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:28 (28.0 B)  TX bytes:24144 (23.5 KiB)

br-5ef4b885534c Link encap:Ethernet  HWaddr 02:42:22:85:B8:BD  
          inet addr:172.19.0.1  Bcast:172.19.255.255  Mask:255.255.0.0
          UP BROADCAST MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

br-72ad7e46c6a5 Link encap:Ethernet  HWaddr 02:42:FB:29:0E:6C  
          inet addr:172.18.0.1  Bcast:172.18.255.255  Mask:255.255.0.0
          UP BROADCAST MULTICAST  MTU:1500  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

docker0   Link encap:Ethernet  HWaddr 02:42:E3:3B:1E:AF  
          inet addr:172.17.0.1  Bcast:172.17.255.255  Mask:255.255.0.0
          inet6 addr: fe80::42:e3ff:fe3b:1eaf/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:6 errors:0 dropped:0 overruns:0 frame:0
          TX packets:46 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:392 (392.0 B)  TX bytes:5762 (5.6 KiB)

eno1      Link encap:Ethernet  HWaddr 84:A9:3E:5F:A7:B6  
          inet addr:192.168.164.201  Bcast:192.168.165.255  Mask:255.255.254.0
          inet6 addr: fe80::2b2e:3e23:a160:43dc/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:285974 errors:0 dropped:0 overruns:0 frame:0
          TX packets:106499 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:31934283 (30.4 MiB)  TX bytes:8601348 (8.2 MiB)
          Interrupt:16 Memory:e0000000-e0020000 

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:18084 errors:0 dropped:0 overruns:0 frame:0
          TX packets:18084 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:1737546 (1.6 MiB)  TX bytes:1737546 (1.6 MiB)

veth9779e22 Link encap:Ethernet  HWaddr 2A:80:54:B6:2A:CF  
          inet6 addr: fe80::2880:54ff:feb6:2acf/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:6 errors:0 dropped:0 overruns:0 frame:0
          TX packets:77 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:476 (476.0 B)  TX bytes:9284 (9.0 KiB)

veth9d287c9 Link encap:Ethernet  HWaddr B2:FF:64:3D:82:6C  
          inet6 addr: fe80::b0ff:64ff:fe3d:826c/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1291 errors:0 dropped:0 overruns:0 frame:0
          TX packets:1473 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:72476 (70.7 KiB)  TX bytes:96169 (93.9 KiB)

vethbf1600d Link encap:Ethernet  HWaddr 7A:53:61:B3:75:F2  
          inet6 addr: fe80::7853:61ff:feb3:75f2/64 Scope:Link
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:1292 errors:0 dropped:0 overruns:0 frame:0
          TX packets:1511 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:72689 (70.9 KiB)  TX bytes:101476 (99.0 KiB)

/ # route -n
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
0.0.0.0         192.168.164.1   0.0.0.0         UG    100    0        0 eno1
169.254.0.0     0.0.0.0         255.255.0.0     U     1000   0        0 eno1
172.17.0.0      0.0.0.0         255.255.0.0     U     0      0        0 docker0
172.18.0.0      0.0.0.0         255.255.0.0     U     0      0        0 br-72ad7e46c6a5
172.19.0.0      0.0.0.0         255.255.0.0     U     0      0        0 br-5ef4b885534c
172.20.0.0      0.0.0.0         255.255.0.0     U     0      0        0 br-08d9e5f6c764
172.26.10.96    192.168.164.1   255.255.255.255 UGH   100    0        0 eno1
192.168.164.0   0.0.0.0         255.255.254.0   U     100    0        0 eno1
/ # ping 192.168.164.168
PING 192.168.164.168 (192.168.164.168): 56 data bytes
64 bytes from 192.168.164.168: seq=0 ttl=64 time=0.488 ms
64 bytes from 192.168.164.168: seq=1 ttl=64 time=0.651 ms
64 bytes from 192.168.164.168: seq=2 ttl=64 time=0.635 ms
^C
--- 192.168.164.168 ping statistics ---
3 packets transmitted, 3 packets received, 0% packet loss
round-trip min/avg/max = 0.488/0.591/0.651 ms
/ # 

```

## container 模式
```
wangxinlong@wangxinlong:~$ docker run --name b3 -it --rm busybox:latest 
/ # ifconfig
eth0      Link encap:Ethernet  HWaddr 02:42:AC:11:00:03  
          inet addr:172.17.0.3  Bcast:172.17.255.255  Mask:255.255.0.0
          UP BROADCAST RUNNING MULTICAST  MTU:1500  Metric:1
          RX packets:24 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:0 
          RX bytes:2931 (2.8 KiB)  TX bytes:0 (0.0 B)

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:0 errors:0 dropped:0 overruns:0 frame:0
          TX packets:0 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:0 (0.0 B)  TX bytes:0 (0.0 B)

/ # netstat -nutl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       
/ # echo "hello world b1" > /tmp/index.html
/ # httpd -h /tmp/ 
/ # netstat -nutl
Active Internet connections (only servers)
Proto Recv-Q Send-Q Local Address           Foreign Address         State       
tcp        0      0 :::80                   :::*                    LISTEN      
/ # ls
bin   dev   etc   home  proc  root  sys   tmp   usr   var
/ # 

```

## none　模式

## docker-compose
docker compose 是docker官方编排项目．

使用　docker-compose up

docker-compose.yml 语法：
```
version: '3' # 定义版本，不指定默认为版本 1，新版本功能更多

services: # 容器，就像 docker run
   db: # 名称，它也是 network 中 DNS 名称
     image: mysql:5.7 # 镜像，如果像自定义镜像可以不指定这个参数，而用 build
     volumes: # 定义数据卷，类似 -v
       - db_data:/var/lib/mysql
       - .:/aaa # 挂载当前目录到容器中的 /aaa 无需使用绝对路径
     restart: always # 类似 --restart
     # 'no' 默认，不自动重启，以为 no 是 yaml 关键字所以加引号
     # always 总是自动重启
     # on-failure 当失败时自动重启，也就是 exit code 不为 0 时
     # unless-stopped 除非手动停止，否者一直重启
     environment: # 定义环境变量，类似 -e
       MYSQL_ROOT_PASSWORD: somewordpress
       MYSQL_DATABASE: wordpress
       MYSQL_USER: wordpress
       MYSQL_PASSWORD: wordpress
   wordpress: # 第二个容器
     labels:
       com.example.description: "This label will appear on all containers for the web service"
     # 为容器添加 Docker 元数据（metadata）信息。例如可以为容器添加辅助说明信息。
     depends_on: # 帮助 compose 理解容器之间的关系
     # db 将会在 wordpress 之前被启动
     # 关闭时 wordpress 将会在 db 之前关闭
     # 我们指定只启动 wordpress，db 也会跟着启动
       - db
     image: wordpress:latest
     ports: # 端口，类似 -p
       - "8000:80"
     restart: always
     environment:
       WORDPRESS_DB_HOST: db:3306
       WORDPRESS_DB_USER: wordpress
       WORDPRESS_DB_PASSWORD: wordpress

volumes: # 可选，需要创建的数据卷，类似 docker volume create
  db_data:

networks: # 可选，需要创建的网络，类似 docker network create
```

## 用户自定义bridge　网络
通过docker network create命令创建自定义bridge网络，下面的命令创建了一个名为my_bri的网络

不用通ｂｒｉｄｇｅ之间是无法互相通信的