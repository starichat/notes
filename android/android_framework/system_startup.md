#  安卓系统启动
>  Android 系统底层是基于 Linux 操作系统。当用户按下开机之后，安卓设备要经过BootLoader、Linux Kernel 和安卓系统服务三个阶段才能完全启动。当Linux Kernel启动会创建Init进程，该进程是所有用户空间的开始，进而会启动ServiceManager，Zygote进程。Zygote进程继而创建SystemServer等进程。

![avatar](./images/ATMS.png)

## Boot ROM
用户按下开机键之后，系统执行一系列处理，触发执行 Boot ROM的代码，并将Bootloader加载到RAM中开始执行

## Boot loader
引导加载程序

## Kernel
Android 内核层，Android内核启动，会开始设置缓存，受保护的内存，调度和加载驱动程序，并开始准备启动init进程.并且内核进程都是由

## init 进程 
Android 中第一个被启动的进程是 init 进程，Android 中的其他进程都是由该进程fork出来的，该进程是 Android 系统的祖宗，进程号为 1。下面将根据系统启动顺序逐步介绍 Android 从 init 进程到SystemServer启动的过程。
init 进程代码路径为：<code>system/core/init/init.cpp </code>。当Krenel启动后就会调用init.cpp中的main方法，init进程main方法主要功能如下：

1. 赋予一些文件权限、初始化运行环境，创建磁盘并挂载文件系统，初始化内核log系统和用户log系统以及selinux初始化等操作 
2. 信号处理，处理子进程的终止功能，防止子进程形成僵尸进程占用cpu资源
3. 提供属性服务
4. 解析并运行Init.rc文件

这里重点解释一下后面三点
### 信号处理，处理子进程的终止功能，防止子进程形成僵尸进程占用cpu资源
代码路径： <code>/system/core/init/sigchld_handler.cpp</code>
看一下代码init 进程调用sigchld_handler_init() 来初始化信号处理
```
void sigchld_handler_init() {
    // 创建一个socket pair，一端用于写数据，一端用于读数据
    int s[2];
    if (socketpair(AF_UNIX, SOCK_STREAM | SOCK_NONBLOCK | SOCK_CLOEXEC, 0, s) == -1) {
        PLOG(FATAL) << "socketpair failed in sigchld_handler_init";
    }

    signal_write_fd = s[0];// 写
    signal_read_fd = s[1];// 读

    // 利用 act 接收捕获SIGNALD，如果捕获到子进程终止的信号则写入signal_write_fd
    struct sigaction act;
    memset(&act, 0, sizeof(act));
    act.sa_handler = SIGCHLD_handler;
    act.sa_flags = SA_NOCLDSTOP;
    sigaction(SIGCHLD, &act, 0);

    ReapAnyOutstandingChildren();// 进入waitpid并处理子进程是否退出

    register_epoll_handler(signal_read_fd, handle_signal);  // 当signal_read_fd 可读便会handle_signal函数
}
```
大体看一下信号处理的调用链：
1. signal_handler_init() 
2. socketpair() 
3. sigaction(SIGCHLD, &act, 0)
4. reap_any_outstanding_children()
5. register_epoll_handler(signal_read_fd, handle_signal)

> 采用信号处理方式可以有效防止僵尸进程的产生。在init.cpp的面函数中有signal_handler_init() 方法来初始化信号处理过程。该函数会创建一个socketpair的全双工通信管道用来读取和写入数据。socketpair()函数是Linux用于创建创建无名的相互连接的套接字可以在同一个进程中进行读写操作。每个进程在处理其他进程发送的signal信号时都需要先注册，当其子进程终止就会产生SIGCHLD信号，init进程捕捉到该信号写入signal_write_fd中，然后通过signal_read_fd可读，触发handle_signal 函数进一步调用ReapAnyOutstandingChildren()来调用ReapOneProcess函数的ServiceList::GetInstance().RemoveService(*service)来移除该进程。

### 提供属性服务
代码位置： <code>/system/core/init/propertice_service.cpp</code>

在启动属性服务之前，init 进程会先执行property_init()方法。
[source,cpp]
----
void property_init() {
    mkdir("/dev/__properties__", S_IRWXU | S_IXGRP | S_IXOTH);
    CreateSerializedPropertyInfo();
    if (__system_property_area_init()) {//创建属性共享内存空间
        LOG(FATAL) << "Failed to initialize property area";
    }
    if (!property_info_area.LoadDefaultPath()) {
        LOG(FATAL) << "Failed to load serialized property info file";
    }
}
----

1. 属性初始化会先执行__system_property_area_init()，用于创建跨进程的共享内存，用于存储增加和修改的属性内容。
2. 接着系统会将默认的prop文件加载到共享内存中，根据函数实现主要有以下文件会先加载到共享内存中去：

 * /system/build.prop；
 * /vendor/build.prop；
 * /factory/factory.prop；
 * /data/local.prop；
 * /data/property路径下的persist属性

3. 启动属性服务，系统启动属性服务会先创建一个名为 property_service 的 socket ，用来监听来自进程写系统属性的请求，同时注册在该socket上，注册一个epoll事件的回调函数handle_property_set_fd,用来接收socket绑定请求，并接收发送来的数据，调用property_set 来设置属性到共享内存中去。

4.  获取属性

安卓系统有两个接口访问系统属性：

 * 一是Java层通过调用framework层的 SystemProperties.java 中提供的接口来访问来访问安卓系统属性，该接口的具体实现是在native层实现的，位置在<code>/frameworks/base/core/jni/android_os_SystemProperties.cpp</code>,通过SystemProperties_get（）来调用<code>system/core/base/properties.cpp</code> 里的GetProperty()从共享内存中取数据。

 * 二是C++层 <code>/bionic/libc/bionic/system_property_api.cpp</code> 里的__system_property_get函数通过最终是调用system_properties.Get(name, value)来获取共享内存里的系统属性值。

### 解析并运行Init.rc文件
> rc 文件是一种脚本文件，由action，commands，service，options 4种类型的声明构成

1. Action 动作:响应某个事件的过程，通过触发器，以on开头的语句来决定是否执行相应的service
2. Commands 命令: 
> 常用命令如下：

* class_start <service_class_name>： 启动属于同一个class的所有服务；
* start <service_name>： 启动指定的服务，若已启动则跳过；
* stop <service_name>： 停止正在运行的服务
* setprop <name> <value>：设置属性值
* mkdir <path>：创建指定目录
* symlink <target> <sym_link>： 创建连接到<target>的<sym_link>符号链接；
* write <path> <string>： 向文件path中写入字符串；
* exec： fork并执行，会阻塞init进程直到程序完毕；
* exprot <name> <name>：设定环境变量；
* loglevel <level>：设置log级别

3. services 服务
由init进程启动的service，运行在init的子进程中。，每一个service的启动都是由fork的方式生成的。
例如：
[source,txt]
----
on post-fs
 
    load_system_props
    
    start servicemanager
----
即表示启动servicemanager服务

4. Options 选项

--------------

示例：
当解析rc文件解析到以下语句时
[source,txt]
----
service zygote /system/bin/app_process64 -Xzygote /system/bin --zygote --start-system-server --socket-name=zygote
    class main
    priority -20
    user root
    group root readproc reserved_disk
    socket zygote stream 660 root system
    onrestart write /sys/android_power/request_state wake
    onrestart write /sys/power/state on
    onrestart restart audioserver
    onrestart restart cameraserver
    onrestart restart media
    onrestart restart netd
    onrestart restart wificond
    writepid /dev/cpuset/foreground/tasks
----
以上rc文件可以解析出来的语句为：
[source,txt]
----
ServiceName:zygote
Path: /system/bin/app_process64
Arguments:-Xzygote /system/bin --zygote --start-system-server
----
意思是启动Zygote，因为是64位系统，该进程对应的进程路径为app_process64,对应的main函数的参数有-Xzygote /system/bin --zygote --start-system-server,以下就进入Zygote的启动过程了。
类似的

## Zygote 启动
先奉上zygote启动的时序图
![avatar](./images/ATMS.png)

init 进程通过解析 rc 进程来启动 app_main 进程。
下面就直接根据下面调用链进行启动：

1. App_main.main()
2. AndroidRuntime.start()
3. startVm()
4. startReg()
5. ZygoteInit.main()
6. registerZygotesocket()
7. preload()
8. startSystemServer()
9. runSelectLoop()

### app_main
这一段的代码比较简单明了，就不贴代码了，最后
根据main函数的上述代码以及rc文件解析的内容的传入参数中可以得出zygote = true 和 startSystemServer = true。
然后根据以上结果继续执行main函数，在以下代码段中

[source,cpp]
----
if (zygote) {
        runtime.start("com.android.internal.os.ZygoteInit", args, zygote);
    } else if (className) {
        runtime.start("com.android.internal.os.RuntimeInit", args, zygote);
    } else {
        fprintf(stderr, "Error: no class name or --zygote supplied.\n");
        app_usage();
        LOG_ALWAYS_FATAL("app_process: no class name or --zygote supplied.");
    }
----

可以发现zygote进程会直接进入zygote分支，进而启动runtime.start()函数。执行runtime.start（）会调用AndroidRuntime的start函数。

### AndroidRuntime
这是安卓的核心，是安卓的虚拟机，具体细节在后面ART部分会涉及，这里暂时只简单了解一下ART的功能：
AndroidRuntime的start函数主要执行以下功能：

1. 初始化JNI环境
在启动虚拟机之前，需要调用jni_invocation.Init(NULL)来初始化当前运行环境，获取JNI_CreateJavaVM环境等。
2. 启动虚拟机。虚拟机创建成功，执行回调函数，执行回调函数通知调用者
调用startVm()来启动虚拟机，主要执行了：
 * 初始化虚拟机参数
 * 调用了JNI_CreatejavaVM启动安卓虚拟机，并通过回调函数onVmCreated回调函数通知调用者。
3. 注册native函数
调用startReg(env)注册本地函数，为native本地函数提供加载到虚拟机的入口。
4. 通过反射执行目标函数主函数。
通过调用 env->CallStaticVoidMethod(startClass, startMeth, strArray)找到目标函数ZygoteInit的main函数。

### ZygoteInit

下面分析ZygoteInit的main函数，代码位置[blue]#/frameworks/base/core/java/com/android/internal/os/ZygoteInit.java#

ZygoteInit主要完成了以下功能：

1. 注册socket
通过调用zygoteServer.registerServerSocketFromEnv(socketName)注册一个socket，调用LocalsocketImpl.create()创建一个socket fd，然后将socket和fd注册到环境变量中，以便于其他进程通过getenv就能汇过去到fd。接下来只需要监听AndroidRuntime是否有fork新进程了。
2. 通过preload预加载各类资源。
3. 因为init.rc里的参数有start-system-server,所以这里可以通过调用forkSystemServer(abiList, socketName, zygoteServer)来fork  Zygote的第一个进程System Server.通过forkSystemServer时会获得一个pid=0的新生进程，进一步会调用handleSystemServerProcess函数来启动ystem server。

handleSystemServerProcess是启动SystemServer的核心
主要执行了以下工作：

执行zygoteinit（）进行初始化 -->调用commonInit（）初始化一些基础组件，剩下的初始化工作大概分两步，先在native层里实现，接着在Java层完成。  --> native层的工作在AppRuntime(AndroidRuntime的子类）的onZygoteInit()完成，启动了一个Thread, 这个Thread是SystemServer的主线程, 负责接收来至其他进程的Binder调用请求。--> 然后通过RuntimeInit.applicationInit(targetSdkVersion, argv, classLoader)来调用findStaticMain函数。findStaticMain利用反射获取ZygoteInit的main方法入口，--> 最后通过返回一个new MethodAndArgsCaller(m, argv)来被ZygoteInit.java 捕获并调用caller.run()进而调用SystemServer的main方法。

最后Zygote执行zygoteServer.runSelectLoop(abiList)方法进入一个死循环，当ActivityManagerServie需要请求创建新的app时才会唤醒zygote。

## SystemServer
SystemServer 的启动过程比较简单，主要通过main()函数调用run()方法，调用Looper.prepareMainLooper()创建一个主线程Looper,加载一些android——services的lib资源，处理启动System Services的相关环境配置等。
准备工作做好了，接下来就准备启动services
[source,java]
----
 try {
            traceBeginAndSlog("StartServices");
            startBootstrapServices();
            startCoreServices();
	    Log.d("wxl","otherservices start");
            startOtherServices();
            SystemServerInitThreadPool.shutdown();
----
由上可知systemserver启动过程如下：

1. 调用Looper.prepareMainLooper()创建一个主线程Looper
2. 加载一些android——services的lib资源，处理启动System Services的相关环境配置等
3. 启动startBootstrapServices()系统引导相关服务
4. 启动startCoreServices()系统核心服务
5. 启动startOtherServices()其他服务

至此，整个系统就已经完全启动了SystemServer将进入Looper.loop()的长循环中。

[注] system
server启动的服务过多，所以就不全部列出了

