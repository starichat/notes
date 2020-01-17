= PMS 研究
wangxinlong <wangxinlong@iauto.com>
:toc: left

:numbered:

== 概述
PackageManagerService 是安卓的核心服务之一，管理着安卓的package，实现扫描package路径并解析package，以及进一步的安装、卸载。
PackageManagerService 由 system_server 进程启动，主要有三个过程：

* PackageManagerService.main ：执行PackageManager的启动和初始化等过程

* PackageManagerService.updatePackagesIfNeeded：执行判断是否需要对package进行更新，以及优化等操作

* PackageManagerService.systemReady ：通知系统PackageManagerService进入就绪状态

先整体看下PMS流程图:

PackageManagerService 整个过程流程图如下：

image:images/pms/pms.png[width=1000]

相关类图如下：

image:images/pms/class.png[width=1000]


== PackageManagerService.main

该过程主要通过获取PackageManagerService对象，执行其构造函数，并实现以下功能：

该过程主要实现了以下功能：

1. UserManagerService、Settings、SysConfig等对象的创建以及初始化
2. 通过Setting读取配置文件、获取权限等信息并保存
3. 扫描系统目录，并通过ParallelPackageParser并发解析package信息并保存到Settings中，最后持久化到packages.xml中.
4. 通过消息机制，通知pms执行安装，清理残余文件等文件.

=== 初始化Settings,SystemConfig
创建Settings并执行其构造函数时创建了一些系统文件夹何一些包管理文件:

* packages.xml和package-backup.xml，用于描述系统所安装的Package信息，其中packages-backup.xml是package.xml的备份。
* packages-list用于描述系统中存在的所有非系统自带的apk信息及UID大于10000的apk。
* packages-stopped.xml 和 packages-stopped-backup.xml，强制停止应用信息的备份，其中packages-stopped-backup.xml是packages-stopped.xml的备份



在Settings中,mSharedUsers是一个map对象,通过来管理SharedUserSettings对象，Settings中的mOtherUserIds和mUserIds也类似原理，基于userId管理着SharedUserSettings对象。
同时ShareUserSettings持有一组PackageSetting,PakcageSetting 保存者package的信息.

SystemConfig 对象维护着系统配置相关的内容,从/system, /vendor, /odm,/oem, 等目录中读取xml文件,读取其中的权限配置信息,诸如feature、library、permission、assign-permission等节点下的信息,保存到SystemConfig中


=== 扫描 Package

扫描并解析package的详细时序图如下：

image:images/pms/scan.png[width=1000]

在 PackageManagerService 的构造函数中通过 scanDirTracedLI(),扫描如下路径的package

* /vendor/overlay
* /product/overlay
* /product_services/overlay
* /oem/overlay
* /odm/overlay
* /system/priv-app
* /system/framework
* /system/app
* /vendor/app
* /vendor/priv-app
* /odm/priv-app
* /odm/app
* /oem/app
* /product/priv-app
* /product/app
* /product_service/priv-app
* /product_services/app

接下来调用一个并发解析器,根据以下流程执行解析任务，依次解析轻量级Package信息（获取是否为核心apk等信息），再执行parseBaseapk对apk详细信息进行解析。解析获取Package的内容,并通过commitScanResultsLocked() 提交扫描结果给PMS,PMS通过commitPackageSettings将包信息提交到Setting里,解析流程图如下:

image:images/pms/parsepackage.jpg[align=left]

解析完成之后,依次执行以下步骤将package提交到settings中保存.

1. addForInitLI():将扫描包的应用包信息存储到PMS中,并对包的签名/版本等信息进行处理,移除不合格的包等操作

2. commitScanResultsLocked():提交包扫描并修改系统状态

3. commitPackageSettings():执行保存动作,并最终将package信息保存到Setting中.

这里可以通过adb shell可以查看解析到的package具体内容：

通过 adb shell dumpsys package com.android.theme.icon_pack.circular.android 的　package　信息如下：
[source, txt]
----
Packages:
  Package [com.android.theme.icon_pack.circular.android] (85f0391):
    userId=10024
    pkg=Package{6b0e2f6 com.android.theme.icon_pack.circular.android}
    codePath=/system/product/overlay/IconPackCircularAndroid
    resourcePath=/system/product/overlay/IconPackCircularAndroid
    legacyNativeLibraryDir=/system/product/overlay/IconPackCircularAndroid/lib
    primaryCpuAbi=null
    secondaryCpuAbi=null
    versionCode=1 minSdk=29 targetSdk=29
    versionName=1.0
    splits=[base]
    apkSigningVersion=3
    applicationInfo=ApplicationInfo{ebbc6f7 com.android.theme.icon_pack.circular.android}
    flags=[ SYSTEM ALLOW_CLEAR_USER_DATA ALLOW_BACKUP ]
    privateFlags=[ PRIVATE_FLAG_ACTIVITIES_RESIZE_MODE_RESIZEABLE_VIA_SDK_VERSION ALLOW_AUDIO_PLAYBACK_CAPTURE PRODUCT ]
    dataDir=/data/user/0/com.android.theme.icon_pack.circular.android
    supportsScreens=[small, medium, large, xlarge, resizeable, anyDensity]
    timeStamp=2019-10-21 12:11:44
    firstInstallTime=2019-10-21 12:11:44
    lastUpdateTime=2019-10-21 12:11:44
    signatures=PackageSignatures{69d2964 version:3, signatures:[d7f1f224], past signatures:[]}
    installPermissionsFixed=false
    pkgFlags=[ SYSTEM ALLOW_CLEAR_USER_DATA ALLOW_BACKUP ]
    overlayTarget=android
    overlayCategory=android.theme.customization.icon_pack.android
    User 0: ceDataInode=22965 installed=true hidden=false suspended=false stopped=false notLaunched=false enabled=0 instant=false virtual=false
      runtime permissions:
    User 10: ceDataInode=37107 installed=true hidden=false suspended=false stopped=false notLaunched=false enabled=0 instant=false virtual=false
      runtime permissions:
----

扫描并解析完package后，同时保存package信息到settings中,然后基于消息机制,创建ThreadHandler线程，并以其Looper为参数创建PackageHandler对象，用于程序的安装和卸载.

== 优化阶段
上一阶段完成后,接着通过system_server调用PackageManagerService.updatePackagesIfNeeded()来对执行对系统进行优化。

优化过程时序图如下：

image:images/pms/update.png[width=1000]


以上过程，我们重点看看dexoptPath()函数，主要执行了以下功能：

1. 获取编译过滤器
2. 判断apk的dex文件是否需要优化
3. 判断是否需要在apk文件的同级目录下创建oat文件夹
4. 执行dexopt优化过程

=== 编译过滤器配置
系统根据编译原因进行组合获取编译过滤器的值，编译原因以及对应的编译过滤器属性内容如下：

[source,txt]
----
 "first-boot"：第一次开机 (--> pm.dexopt.first-boot)
 "boot"：启动  (--> pm.dexopt.boot)
 "install"：安装   (--> pm.dexopt.install)
 "bg-dexopt"：后台优化  (--> pm.dexopt.bg-dexopt)
 "ab-ota"：A/B分区ota升级 (--> pm.dexopt.ab-ota)
 "inactive"：不活跃代码  (--> pm.dexopt.inactive)
 "shared"：共享库代码  (--> pm.dexopt.shared)
----

根据编译的reason，通过getCompilerFilterForReason()得到pm.dexopt.REASON字符串，并由getAndCheckValidity去查询系统属性为pm.dexopt.REASON的值，比如，本次REASON，我们为first-boot,所以我们去获取pm.dexopt.first-boot 属性的值，该值在 runtime_libart.mk 指定，内容如下
[source, txt]
----
ifeq (eng,$(TARGET_BUILD_VARIANT))
    PRODUCT_SYSTEM_DEFAULT_PROPERTIES += \
        pm.dexopt.first-boot=extract \
        pm.dexopt.boot=extract
else
    PRODUCT_SYSTEM_DEFAULT_PROPERTIES += \
        pm.dexopt.first-boot=quicken \
        pm.dexopt.boot=verify
endif

# The install filter is speed-profile in order to enable the use of
# profiles from the dex metadata files. Note that if a profile is not provided
# or if it is empty speed-profile is equivalent to (quicken + empty app image).
PRODUCT_SYSTEM_DEFAULT_PROPERTIES += \
    pm.dexopt.install=speed-profile \
    pm.dexopt.bg-dexopt=speed-profile \
    pm.dexopt.ab-ota=speed-profile \
    pm.dexopt.inactive=verify \
    pm.dexopt.shared=speed
----

上面提到的编译过滤器,在安卓art中定义,内容如下,箭头指向为对应java层的表达式内容:
[source, txt]
----
enum Filter {
    kAssumeVerified,      // 跳过验证，但无论如何都将所有类别标记为已验证。--> assume-verified
    kExtract,             // 将验证延迟到运行时，不要编译任何东西。 --> extract
    kVerify,              // 仅验证类。   --> verify
    kQuicken,             // 验证，加快和编译JNI存根。  --> space
    kSpaceProfile,        // 根据配置文件最大程度地节省空间。  --> space-profile
    kSpace,               // 最大限度地节省空间。   --> space-profile
    kSpeedProfile,        // 根据配置文件最大化运行时性能。  --> speed-profile
    kSpeed,               // 最大化运行时性能。  --> speed
    kEverythingProfile,   // 编译所有可以基于配置文件编译的内容。--> everything-profile
    kEverything,          // 编译所有可以编译的东西。  --> everything
  };
----

===  判断apk的dex文件是否进行优化
 
判断是否对apk进行优化的判断逻辑是在native层实现的，代码位置 [blue]#art/runtime/native/dalvik_system_DexFile.cc# 代码过多，此处就先不贴了。

主要有两个关键的地方进行判断：

1. 如果是启动类路径的代码，始终视为最新的，直接返回kNoDexOptNeeded。
+
[source, cpp]
----
if (oat_file_assistant.IsInBootClassPath()) {
    return OatFileAssistant::kNoDexOptNeeded;
  }
----
因为对于启动类路径的代码进行了预编译，启动类路径的默认的预编译过滤器选项为speed过滤器。

2. 如果非启动类路径代码，会进行一系列的文件检查，最后根据文件状态信息会得到以下返回值。
[source, cpp]
----
enum DexOptNeeded {
 kNoDexOptNeeded = 0, //已经编译，不需要执行　dexopt　来更新该apk　
 kDex2OatFromScratch = 1, //有dex文件，但还没编过
 kDex2OatForBootImage = 2,//oat文件不能匹配boot image（系统升级 boot image会变化）
 kDex2OatForFilter = 3,//oat文件不能匹配compiler filter(更改了compiler　filter)
}
----

除了不需优化的kNoDexOptNeeded选项外，其他的都需要再进行优化编译。

以下对具体进行检查的逻辑进行解析，大致过程如下图所示，这里重点讲解几个关键的校验方法：

image:images/pms/dexoptneeded.jpg[align=left]

* 过滤器不匹配校验:
+

 获取oat文件，并取得该文件的编译过滤器,如果用于生成文件的编译过滤器大于等于给定的编译过滤器，则返回true,表示不需要进行再次编译了。
 如果dexopt的目的是降低编译器过滤器，那么profile_changed应该为true，以表示最近更改的配置文件并返回false。返回faslse，
 则就会进入检查dex文件.

* oat 文件校验:
+

 1. 获取oat文件并判断oat文件头中的concurrent-copying字段信息,是否和kUseReadBarrier(该变量作用是生成屏障,表示是否触发并发复制gc)相等,即判断系统的并发复制gc和oat文件是否一致,是则继续校验vdex文件,否则返回false,表示该文件在该系统中不能打开
 
 2. 校验vdex文件,具体校验在下面具体讲解
 
 3. 校验该oat文件的编译过滤器是否需要验证,如果需要则继续检验Bootclasspath的checksum,否则直接校验manshipdex文件.
 
 4. 校验BootClassPath的checksum,在下面会详细讲解如何校验,如果校验成功则继续校验
  ,有该字段则继续进行vdex校验,过滤器认证(检验该文件的编译过滤器是否需要认证)，BootClasspath校验、dex文件校验等，没有则表示该oat文件不可用，则直接校验dex文件。
  
 5. 校验dex文件

* vdex文件校验:
+

  1. 根据Get... 获取该vdex对应的dex文件,并从dex文件中取得checksum(checksum是对除了magic和checksum以外的内容通过crc32算法得到的) 并存到checksums这样一个数组中.
  2. 根据获取到checksums的长度取得dex文件数量,并和对应的vdex文件中存取的checksums的dex文件个数进行比较,如果一样则继续比较checksums的值,否则返回false,校验失败.
  3. 比对vdex文件的checksums和对dex文件求得的checksums的值,一致则返回true,否则返回false.

* 是否有原始dex文件校验:
+

 判断获取的dex文件的magic是否和dex匹配，然后对于单个dex文件，校验单个头文件的checksum。
 对于zip文件，这是classes.dex和每个额外的multidex条目classes2.dex、classes3.dex等单个文件的CRC32算法的checksum并存入数组 
 checksums中。如果可以找到校验和，则返回true，否则返回false。
  
* BootClassPath校验:
+

 1. 从oat文件中获取bootclasspath字段的值,存储着系统bootclasspath包含的文件路径.如果获取不到则直接返回false,否则继续校验bootclasspath路径的文件的checksum,判断文件有没有被更改过.
 2. oat文件的bootclasspath-checksums存储着类似这样的内容i;11/a43b5245, 其中11表示bootclasspath的文件个数,根据该文件个数,对从系统中获取的真实的bootclasspath文件个数进行比较,不合格则返回false,否则返回继续逐个校验bcp单个文件的文件名和checksum,获取到每个文件的checksum追加到待比较的checksums中,并最后和从oat文件中的checksums进行比较.相等则返回true，否则返回true。
 
通过adb shelloatdump --header-only --oat-file=system@product@app@ModuleMetadata@ModuleMetadata.apk@classes.dex,可以看到其中的的关于校验的内容，concurrent-copying,CHECKSUM、LOCATION等

[source, txt]
----
......
MAGIC:
oat
170

LOCATION:
system@product@app@ModuleMetadata@ModuleMetadata.apk@classes.dex

CHECKSUM:
0x220f411b

INSTRUCTION SET:
X86_64
......
KEY VALUE STORE:
bootclasspath = /apex/com.android.runtime/javalib/core-oj.jar:/apex/com.android.runtime/javalib/core-libart.jar:/apex/com.android.runtime/javalib/okhttp.jar:/apex/com.android.runtime/javalib/bouncycastle.jar:/apex/com.android.runtime/javalib/apache-xml.jar:/system/framework/framework.jar:/system/framework/ext.jar:/system/framework/telephony-common.jar:/system/framework/voip-common.jar:/system/framework/ims-common.jar:/system/framework/android.test.base.jar
bootclasspath-checksums = i;11/a43b5245
classpath = PCL[]{PCL[/system/framework/android.hidl.manager-V1.0-java.jar*1684516937]{PCL[/system/framework/android.hidl.base-V1.0-java.jar*604948433]}#PCL[/system/framework/android.hidl.base-V1.0-java.jar*604948433]}
compilation-reason = first-boot
compiler-filter = quicken
concurrent-copying = true
debuggable = false
......
----  

===  是否创建 oat　文件夹
 
安卓在执行优化前中实现了以下逻辑，如果获取的package的applicationInfo.isSystem()为true且applicationInfo.isUpdateSystemApp也为false时，才会在　apk/jar　同级目录下创建oat文件夹,以便于之后生成odex文件直接使用。

===  执行dexopt优化过程

通过java层的installer.dexopt　借助socket和native层的installd通信，并判断oat文件夹是否存在？
若存在，则直接基于以下代码在oat文件夹下生成如下odex文件
[source, cpp]
----
std::string res_ = oat_dir_ + '/' + instruction_set + '/'
            + apk_path_.substr(start + 1, end - start - 1) + ".odex";
----
如果oat文件夹不存在，则创建如下oat文件
[source, cpp]
----
std::string res_ = android_data_dir + DALVIK_CACHE + '/' + instruction_set_ + src_
            + DALVIK_CACHE_POSTFIX;
----
其中 DALVIK_CACHE_POSTFIX = "@classes.dex"

最后生成通过 [blue]#open_vdex_files_for_dex2oat# 生成　vdex　文件,生成的代码目录在　data/dalvik-cache/<isa>/ (<isa> 指的是机器指令集比如本次模拟器使用的是x86_64)中创建vdex文件



