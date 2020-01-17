# Android 启动优化
> android 系统性能优化考虑的方面无非就是两个方面，时间和空间。当然放大或者缩小到计算机的任何地方，都可以从这两个方向着手。结合android 系统启动的实际情景，下面从系统开机时间（时间方面），内存、cpu（空间方面）进行优化。

开机时间：从按下 power 到 launcher 界面展示出来这个过程所花费的总时间。
开机内存：系统完全启动之后，系统内存占用率。
开机cpu：系统完全启动之后，cpu占用率。

## 开机启动时间优化
> 分析工具：bootchart，systrace,logcat,eventlog,dmesg等

使用方法：
adb shell touch /data/bootchart/enabled 之后再重启机器，执行
```
```
就可以绘制出系统启动的bootchart图

[systrace]()


> eventlog 是记录系统启动重要节点的log机制。

打印系统启动关键节点log如下所示：
```
```
各个对应阶段如下表：
```
```
### zygote 启动时间优化
由上图可知，从 AndroidRuntime.start() 到 systemServer 总共耗时，其中比较耗时比重最大的阶段是
Zygote 预加载 class 花费的时间。 预加载类主要作用是预加载/system/framework/config/preload-classes 中配置的class

针对此阶段，我们可以推迟某些 class 的预加载时机，可以将与系统启动无关的 class 推迟到 ui 启动之后。
即预加载资源可以分为两个阶段：
1. early_preloadclasses()
 1. 大部分 android.*
 2. 大部分 java library 诸如：
    * java.io.*
	* java.lang.*
	* java.math.*
	* java.net.*
	* java.nio.*
	* java.security.*
	* java.text.*
	* java.util.*
	* javax.net.ssl.*
	* javax.security.*
	* sun.security
 3. libcore.*
2. late_preloadclasses()
  其他系统启动用不到的class 就可以全部放到ui起来之后再进行加载

### systemserver 启动优化
在systemserver 结束位置## 创建 SystemServiceManager
systemserver 根据创建好的系统上下文将打log，获取systemserver总启动时间，
```
```
systemserver 阶段有三个阶段
startBootStrapServices()
startCoreServices()
startOtherServices()
前两个阶段的 service 涉及到 android 的核心服务，最好不要变动，针对Otherservices 中的非系统启动必要服务也可以采取和zygote预加载资源一样的措施，将 一些service推迟到ui起来之后再启动。
比如：
1. Vibrator service
2. Consumer IR service
3. WiFi NAN and RTT services
4. Network Stats service
5. EthernetService

然后再来结合 eventlog 查看一下system启动比较耗时的服务：
pms

系统第一次启动，pms 扫描阶段主要耗时有：
1. 打开androidmanifest 文件
2. 读取相关属性..
3. 检查签名
4. 保存 package list
第二次启动：
1. 检查 package 是否被更新,是则重新扫描该包
2. 没有更新的,直接从package list中读取

同时，在 pms 中，还将扫描时间超过100ms以上的log输出出来了。可以针对该项，优化100ms以上的log
pms dexopt
部分apk会在系统启动的时候会进行优化，针对优化，我们有两种解决方式，不优化，或者进行预编译优化。在编译期间，我们就将 apk 进行优化。

## Bootanimation

## 内存和cpu
1.在开机用adb指令查看CPU的运行频率及online的CPU
adb shell cat /sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_cur_freq
cat /sys/devices/system/cpu/cpu1/online


## 裁剪
裁剪一些不必要的 apk 或者 service 能够提升系统启动速率和内存占用。





