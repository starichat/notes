= PersistentDataBlockService
wangxinlong <wangxinlong@iauto.com>
v1.0, 2019-08-30
:toc: right

:numbered:

== introduction
PersistentDataBlockService ：用于读取和写入ro.frp.pst指定的持久化分区即 [blue]#/dev/block/bootdevice/by-name/frp# 的服务。该分区对应着Account Data Blocks存放key的数据。

如果不是通过设置ui进行出厂重置的初始化，该分区的数据将会保留。当设备通过设置进行出厂重置时，将擦除该分区的数据。


== 类图
该服务的类图如下

image::images/persistentdatablock/PersistentDataBlockServiceclass.png[width=1200,height=514]

== 启动过程

该服务主要启动的调用过程如下：

1. SystemServer.startOtherServices()
2. mSystemServiceManager.startService(PersistentDataBlockService.class)
3. mServices.add(service)
4. SystemServerInitThreadPool.get().submit()
5. SystemService.publishBinderService()
6. ServiceManager.addService()

== 启动过程解析
安卓的第一个进程init进程在解析init.rc 文件时启动了servicemanager，然后在系统启动过程中，通过Zygote 进程fork出了SystemServer进程，在启动各级系统service时，通过SystemServer.startOtherServices()调用mSystemServiceManager.startService(PersistentDataBlockService.class)执行PersistentDataBlockService的启动。
通过 mSystemServiceManager.startService(PersistentDataBlockService.class)主要实现了以下内容：

1. 通过SystemServiceManager 创建PersistentDataBlockService 对象并执行该类的构造函数
2. 将该对象添加到mSystemServiceManager的成员变量mServices。
3. 执行该对象的onStart() 方法。


回到 [blue]#frameworks/base/services/core/java/com/android/server/PersistentDataBlockService.java# 中
在其中的onStart 方法中

[source, java]
----
 @Override
    public void onStart() {
        // Do init on a separate thread, will join in PHASE_ACTIVITY_MANAGER_READY
        SystemServerInitThreadPool.get().submit(() -> { 
            mAllowedUid = getAllowedUid(UserHandle.USER_SYSTEM);
            enforceChecksumValidity();
            formatIfOemUnlockEnabled();
            publishBinderService(Context.PERSISTENT_DATA_BLOCK_SERVICE, mService);
            mInitDoneSignal.countDown();
        }, TAG + ".onStart");
    }

----

SystemServiceInitThreadPool 是初始化system server的线程池，这里利用线程池开辟一个独立的线程
对PersistentDataBlockService 进程初始化。并调用 publishBinderService 方法，这里是通过调用继承的父类SystemService.java的方法，在SystemService中可以发现，publishBinderService最后还是通过
ServiceManager.addService()来添加注册该服务。至此，PersistentDataBlockService 已经启动完成了。

== PersistentDataBlockService 主要功能描述

PersistentDataBlockService 作为服务端通过AIDL接口定义远程通信的方法，原理是基于framework binder的架构，通过Binder机制在客户端和服务端两个不同的进程间进行通信。

PersistentDataBlockService内部定义了一个 IBinder 内部类，继承
自IPersistentDataBlockService.Stub()，
并实现了 IPersistentDataBlockService.aidl 接口定义的方法。

该 service 主要实现了对 ro.frp.pst 指定的存储分区的读写以及擦除数据的功能。当没有通过终端的设置进行恢复出厂设置不会擦除该分区数据。

* write: 首先通过enforceUid(Binder.getCallingUid()) 确认调用者是否是允许的uid，防止越权操作。获取ro.frp.pst指定的文件路径的文件流对象，然后通过一个ByteBuffer 分配空间，想该buffer中写入数据。将ByteBUffer中的数据写入文件流对象中。
* read： 和write一样，首先确认调用者是否允许的uid，然后通过getTotalDataSizeLocked获取数据大小，获取ro.frp.pst指定的文件路径的文件流对象并从中读取数据。
* wipe: 该功能是整个service的核心，擦除擦除存储分区的数据。

wipe 主要流程：

1. com_android_server_PersistentDataBlockService_wipe()
2. wipe_block_device(fd) 
3. get_block_device_size(fd)
4. ioctl(fd,BLKSECDISCSRD,&range)

wipe 函数的核心最后就是通过 ioctl(fd,BLKSECDISCSRD,&range) 来操作的，ioctl 函数是一个系统调用，作用于一个文件描述符。是设备驱动程序中对设备的 I/O 通道进行管理的函数。

其中，BLKSECDISCARD：安全擦除操作。所以即对ro.frp.pst指定的文件路径丢弃range范围内的数据来进行数据擦除。

== PersistentDataBlockManagerInternal

PersistentDataBlockService 还定义了一个 PersistentDataBlockManagerInternal 的内部类，用于实现PersistentDataBlockManagerInternal 接口
主要作用是对外开放以下接口：

* setFrpCredentialHandle(byte[] handle)
基于 FileChannel 将handle的内容写入 ro.frp.pst 指定的文件路径中.

* getFrpCredentialHandle()：
获取取frp认证处理比设置要简单，和上文类似，这里从ro.frp.pst中指定的分区块读取字节内容。

* forceOemUnlockEnabled(boolean enabled)：
强制打开OEM锁，强制设置系统属性OEM_UNLOCK_PROP=0

== PersistentDataBlockManager

系统将 PersistentDataBlockService 的操作封装system-api中，客户端通过system-api 来对 PersistentDataBlockManager 来操作对用户数据清空等操作  具体通过调用 PersistentDataBlockManager 来调用以下api，同时也需要必要权限才能操作PersistentDataBlockManagerd的相关api，如下表所示：

|===
|api | 特殊权限
|write(byte[] data) |无
|read()  | 无
|getDataBlockSize()  |android.Manifest.permission.ACCESS_PDB_STATE
|getMaximumDataBlockSize()  | 无
|wipe()|android.Manifest.permission.OEM_UNLOCK_STATE
|setOemUnlockEnabled(boolean enabled)|android.Manifest.permission.OEM_UNLOCK_STATE
|getOemUnlockEnabled() | anyOf = {
            android.Manifest.permission.READ_OEM_UNLOCK_STATE,
            android.Manifest.permission.OEM_UNLOCK_STATE
    }
|getFlashLockState() | anyOf = {
            android.Manifest.permission.READ_OEM_UNLOCK_STATE,
            android.Manifest.permission.OEM_UNLOCK_STATE
    }

|===

== frp 机制功能
根据以上研究，可以发现 PersistentDataBlockService 的主要作用就是为了实现frp机制，即出厂设置保护。

主要有以下功能：

* 有设备的 google 账户，就可以通过 ADM 远程锁屏
* frp 开启的情况下，只有在设置apk下进行出厂设置才能清楚ro.frp.pst 指定的文件路径的数据，即设备不存在用户的任何信息了，不需要密码就可以操作了。
* frp 开启时，bootloader下或者ADM下做factory reset不会清除ro.frp.pst 指定的文件路径的数据，即属于未信任服务，进行一些操作仍然需要密码。


== 通过设置apk恢复出厂设置流程图
image::images/persistentdatablock/PersistentDataBlockService.png[width=1800,height=1200]

== 流程解析

1. 在终端设备设置里触发出厂设置按钮，会调/dev/block/bootdevice/by-name/frp用 
 [blue]#/home/wangxinlong/Android_9/packages/apps/Settings/src/com/android/settings/MasterClearConfirm.java#  里的onclick方法
2. 进而会执行new AsyncTask<Void, Void, Void>().execute() 依次调用 PersistentDataBlockManager.wipe() 方法
3. PersistentDataBlockManager 基于Binder机制，通过 AIDL 接口调用 PersistentDataBlockService 里wipe方法
4. wipe 的具体实现是通过native函数在jni层实现的，上文有提到，最终通过调用ioctl函数对ro.frp.pst指定的文件路径的数据进行擦除，当然前提是得有相应的权限。
5. 最后会发送广播通知MasterClearReceiver 去调用RecoverySystem.rebootWipeUserData() 去擦除用户数据

[red]#进行最后一步擦除用户数据之前，需要获取到 UserManager 信息，才能擦除设备的用户信息。#
















