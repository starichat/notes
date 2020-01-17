# 安卓系统启动之「AMS 到 home 的启动过程」

>  安卓系统启动阶段的最后一步就是通过 ActivityManagerService 启动 home。这里就以 ActivityManagerService 到 home 启动的过程为主线来进行讲解。

先整体看下 SystemServer 启动 AMS 的过程

![avatar](./images/systemserver.png)

1. SystemServer 进程启动,执行startBootStrapServices() 启动ActivityTaskManagerService, ActivityManagerService.
2. 执行 SystemServer.startOtherservice() 
3. 调用 ActivityManagerService.systemReady() ,处理 persistentAPP 的启动以及 home 的启动,开机广播的发送等

## AMS , ATMS启动过程分析
![avatar](./images/ATMS.png)

ATMS 启动过程如下:
1. 通过 SystemServiceManager.startService(),先通过反射<code>serviceClass = (Class<SystemService>)Class.forName(className)</code> 创建 ATMS 的内部类 Lifecycle 实例，Lifecycle 的构造方法里会创建 ATMS 的实例，ATMS 在初始化的时候会 new LocalService() 创建一个自己的 local service 实例。
2. 继续执行 SystemServiceManager 的 startService 里的 onStart 回调方法将 ATMS 加入到 ServiceManager 中,并将自己的 LocalService 变量注册到 LocalServices。



![avatar](./images/AMS.png)

AMS 的启动过程如下:
1. systemserver 调用 ActivityManagerService.Lifecycle 的静态方法 startService(mSystemServiceManager, atm) 
2. 借助 SytstemServiceManager 创建 ActivityManagerService.LifeCycle 实例,和 ATMS 一样，在内部类构造函数中会创建一个 ActivityManagerService 对象。
3. 最后通过 SytstemServiceManager 回调方法 onStart() 创建 AMS 的 LocalService 实例并注册到 LocalServices 中。

> 以上 AMS 、ATMS 启动完毕,接着会启动其他一系列系统服务,这里先不列出了。然后进入systemServer.startOtherService(), 在 SystemServer 的最后启动阶段就是通过进入 AMS 的     systemReady() 阶段，执行完 systemReady() 结束，整个系统就启动成功了。

下面进行本篇的重点，AMS 启动 home ，并发送广播标志系统启动完成。

## AMS.systemReady 阶段

![avatar](./images/AMS_SystemReady.png)

AMS 的 systemReady 阶段启动,其参数为一个 Runable 类型的 goingCallback, goingCallback 会执行启动一些系统service的启动，包括启动SystemUi。然后启动 home 应用，最后发送系统广播通知系统启动完成。


### goingCallback.run()

AMS 在调用一些系统服务的SystemReady或者SystemRuning方法，启动并标志这些服务启动完成，并执行一些清理进程和广播等行为之后，就会执行 goingRunable.run() 启动包括 SystemUi 在内的部分系统服务。

### 启动 persistent 进程

启动 persistent 进程,在 APK 的 AndroidManifest.xml 中定义了 persistent 标志位为 true 的 apk 会在该阶段被 AMS 拉起来。

### 启动 home 

所有的系统服务和必要的 apk 都已经启动了,接下来主要会执行以下工作启动 home。home 是开机展示的第一个 activity,下面将对启动 home 进行讲解。

启动流程：
1. 执行 ActivityManagerService.startHomeOnAllDisplays()
2. 执行 RootActivityContainer.startHomeOnAllDisplays() 使用如下一个 for 循环遍历启动每个屏幕的 home activity
```
boolean startHomeOnAllDisplays(int userId, String reason) {
        boolean homeStarted = false;
        for (int i = mActivityDisplays.size() - 1; i >= 0; i--) { // 遍历屏幕个数,启动各个 display 的 home
            final int displayId = mActivityDisplays.get(i).mDisplayId;
            homeStarted |= startHomeOnDisplay(userId, reason, displayId);
        }
        return homeStarted;
    }
```
3. 执行 startHomeOnDisplay(), 并调用 RootActivityContainer.startHomeOnAllDisplay() 
4. 执行 ActivityManagerService.getHomeIntent(), 获取启动 home activity 的 intent, AMS 会构造一个 android.intent.CATOGRY.home 的 intent。
5. 根据获得的 Intent, 基于 Intent 的饮食匹配调用 resolveHomeActivity() 去匹配合适的 activity。
![avatar](./images/resolveHomeActivity.png)
具体过程如下:
 * 5.1. 执行 PMS.resolveIntent()
 * 5.2. 执行 PMS.resolveIntentInternal()
 * 5.3. 执行 queryIntentActivitiesInternal() 查询具体的 activity。因为 PMS 在扫描 package 的时候就已经将 package 的activity的信息存到了 mActivities 中。所有这里就直接从存储的 mactivities 中查询。 
6. startHomeActivity()
7. 执行 ActivityStartController.obtainStarter(intent,reason),获取一个ActivityStater对象，并执行 ActivityStater 对象的 execute()方法
8. 执行 ActivityStarter.startActivity() 启动 home activity，之后就和普通的startActivity 一样了，通过 AMS 启动 home activity<code>关于startActivity具体如何启动进程以及展示出来的过程或者安卓的四大组件如何启动进程的过程这里先不列出了，后面会有专门的文章详细讲解</code>
9. 最后执行 postStartActivityProcessing() ，提交启动activity进程信息


