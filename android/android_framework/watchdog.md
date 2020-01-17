= watchdog
wangxinlong <wangxinlong@iauto.com>
v1.0, 2019-09-10
:toc: right

:numbered:

== Introduction
watchdog ：安卓系统设计的软件层面的watchdog，用于保护一些重要的系统服务，当出现故障时，通常会让安卓系统重启。由于这种机制的存在，就会存在一些system_server 进程被watchdog 杀掉而导致手机重启的问题。

主要作用：

1. 接收系统内部reboot请求，重启系统
2. 监护SystemServer进程，防止系统死锁

== Watchdog 类图

给出 watchdog 的类图以便于清楚了解该类的结构和基本功能。

image::images/watchdog/WatchdogClass.png[width=1200,height=514]

== Watchdog 内部架构和主要接口
根据上面的类图可以清楚地知道 WatchDog 内部主要部件和接口函数

* RebootRequestReceiver：负责接收系统内部发出的重启Intent消息，并进行系统重启。

* HandlerChecker: 用于检查线程的Looper是否处于长时间的工作状态。

* Monitor：每个被监护对象必须要实现的接口，由WatchDog在运行中调用，以实现监护功能。用于检查该对象是否出现死锁。


== Watchdog 工作机制时序图

watchdog 的工作机制时序图如下：

image::images/watchdog/watchdogseq.png[width=1300,height=1600]

== Watchdog 的创建

在SystemServer 进程创建过程中，会启动一系列 Service ，在启动 service 的第三个阶段StartOtherServices() 中会执行 
 [blue]#watchdog.getInstance()# 

watchdog.getIndtace() 是通过如下代码的单例模式的实现的，旨在只能创建一个watchdog对象
[source, java]
----
 public static Watchdog getInstance() {
        if (sWatchdog == null) {
            sWatchdog = new Watchdog();
        }

        return sWatchdog;
    }
----

在创建watchdog对象的同时，会执行该对象的私有构造方法，进行实例化
[source, java]
----
 private Watchdog() {
        super("watchdog");
        // 将前台线程加入到mHandlerCheckers队列
        mMonitorChecker = new HandlerChecker(FgThread.getHandler(),
                "foreground thread", DEFAULT_TIMEOUT);
        mHandlerCheckers.add(mMonitorChecker);
        // 将主线程加入到mHandlerCheckers对列
        mHandlerCheckers.add(new HandlerChecker(new Handler(Looper.getMainLooper()),
                "main thread", DEFAULT_TIMEOUT));
        // 将ui线程加入到mHandlerCheckers队列
        mHandlerCheckers.add(new HandlerChecker(UiThread.getHandler(),
                "ui thread", DEFAULT_TIMEOUT));
        // 将i/o 线程加入到mHandlerCheckers队列
        mHandlerCheckers.add(new HandlerChecker(IoThread.getHandler(),
                "i/o thread", DEFAULT_TIMEOUT));
        // 将display线程加入到mHandlerCheckers队列
        mHandlerCheckers.add(new HandlerChecker(DisplayThread.getHandler(),
                "display thread", DEFAULT_TIMEOUT));

        // 为Binder线程初始化monitor
        addMonitor(new BinderThreadMonitor());
        // 创建一个fdMonitor对象
        mOpenFdMonitor = OpenFdMonitor.create();

        // See the notes on DEFAULT_TIMEOUT.
        assert DB ||
                DEFAULT_TIMEOUT > ZygoteConnectionConstants.WRAPPED_PID_TIMEOUT_MILLIS;
    }

----

Watchdog 类继承自Thread，创建的线程名为watchdog，mHandlerCheckers队列将foreground线程、主线程、ui线程、i/o线程、display线程的HandlerChecker
对象放入到队列中了。

=== HandlerChecker 
HandlerChecker 类实现了Runnable接口，该线程有一个成员变量如下代码：
[source, java]
----
public final class HandlerChecker implements Runnable {
    private final Handler mHandler; //Handler对象
    private final String mName; //线程描述名
    private final long mWaitMax; //最长等待时间
    //记录着监控的服务
    private final ArrayList<Monitor> mMonitors = new ArrayList<Monitor>();
    private boolean mCompleted; //开始检查时先设置成false
    private Monitor mCurrentMonitor;
    private long mStartTime; //开始准备检查的时间点

    HandlerChecker(Handler handler, String name, long waitMaxMillis) {
        mHandler = handler;
        mName = name;
        mWaitMax = waitMaxMillis;
        mCompleted = true;
    }
}
----

=== addmonitor()

[source, java]
----
 public void addMonitor(Monitor monitor) {
        synchronized (this) {
            if (isAlive()) {
                throw new RuntimeException("Monitors can't be added once the Watchdog is running");
        }
            mMonitorChecker.addMonitor(monitor);
    }
}
----

[source, java]
----
 // HandlerChecker 对象
 public void addMonitor(Monitor monitor) {
            mMonitors.add(monitor);
        }

----

即在Watchdog初始化的时候会执行 [blue]#addMonitor(new BinderThreadMonitor())#  函数，会将BinderThreadMonitor 添加到HandlerChecker的mMonitors队列中去。

addmonitor() 是将Binder线程添加到watchdog的监控中并通过线程的handler(mMonitorChecker)来检查是否工作正常.

==== BinderThreadMonitor 
再来看BinderThreadMonitor对象，BinderThreadMonitor 实现了Watchdog.Monitor 接口。

所有实现该接口的service都会被watchdog监控。

[source, java]
----
private static final class BinderThreadMonitor implements Watchdog.Monitor {
    public void monitor() {
        Binder.blockUntilThreadAvailable();
    }
}
----
Binder.blockUntilThreadAvailable()的具体实现是在native层实现的，native代码如下：
 [blue]#framework/native/libs/IPCThreadState.cpp#
[source, cpp]
----
void IPCThreadState::blockUntilThreadAvailable()
{
    pthread_mutex_lock(&mProcess->mThreadCountLock);
    while (mProcess->mExecutingThreadsCount >= mProcess->mMaxThreads) {
        ALOGW("Waiting for thread to be free. mExecutingThreadsCount=%lu mMaxThreads=%lu\n",
                static_cast<unsigned long>(mProcess->mExecutingThreadsCount),
                static_cast<unsigned long>(mProcess->mMaxThreads));
        // 等待正在执行的binder线程
        pthread_cond_wait(&mProcess->mThreadCountDecrement, &mProcess->mThreadCountLock);
    }
    pthread_mutex_unlock(&mProcess->mThreadCountLock);
}
----
该函数的作用是等待空闲的线程。

=== 监控对象
watchdog 主要监控两类对象

1. watchdog 会监控以上线程的Looper是否出现超时则进入watchdog处理机制
+
[options="header"]
|===
|线程名|对应HandlerChecker
|前台线程|new HandlerChecker(FgThread.getHandler(), "foreground thread", DEFAULT_TIMEOUT)
|主线程|new HandlerChecker(new Handler(Looper.getMainLooper()),"main thread", DEFAULT_TIMEOUT)
|ui 线程|new HandlerChecker(UiThread.getHandler(), "ui thread", DEFAULT_TIMEOUT)
|i/o 线程|new HandlerChecker(IoThread.getHandler(),"i/o thread", DEFAULT_TIMEOUT)
|display 线程|new HandlerChecker(DisplayThread.getHandler(), "display thread", DEFAULT_TIMEOUT)
|===



2. watchdog也会监控所有实现watchdog.monitor接口的线程有没有被阻塞，如果阻塞也会触发watchdog处理机制
+

通过调查可以知道实现该接口的类有：
+
* ActivityManagerService
* WindowManagerService
* InputManagerService
* PowerManagerService
* NetworkManagementService
* MountService
* NativeDaemonConnector
* BinderThreadMonitor
* MediaProjectionManagerService
* MediaRouterService
* MediaSessionService
* BinderThreadMonitor

== watchdog 的 init
执行完以上 watchdog,getInstance() 的watchdog的对象的创建之后，开始执行watchdog的初始化：

init主要作用就是执行watchdog的接受系统内部重启的功能。

[blue]#watchdog.getInstance().init()# 
[source, java]
----
public void init(Context context, ActivityManagerService activity) {
        mResolver = context.getContentResolver();
        mActivity = activity;
        context.registerReceiver(new RebootRequestReceiver(),
                new IntentFilter(Intent.ACTION_REBOOT),
                android.Manifest.permission.REBOOT, null);
    }
----
以上代码执行watchdog的初始化，主要注册了一个广播接收器。
[source, java]
----
final class RebootRequestReceiver extends BroadcastReceiver {
        @Override
        public void onReceive(Context c, Intent intent) {
            if (intent.getIntExtra("nowait", 0) != 0) {
                // 接收到广播之后在这里执行watchdog.rebootSystem()
                rebootSystem("Received ACTION_REBOOT broadcast");
                return;
            }
            Slog.w(TAG, "Unsupported ACTION_REBOOT broadcast: " + intent);
        }
    }
----

[source, java]
----
void rebootSystem(String reason) {
        Slog.i(TAG, "Rebooting system because: " + reason);
        IPowerManager pms = (IPowerManager)ServiceManager.getService(Context.POWER_SERVICE);
        try {
            // 通过powerManager执行reboot操作
            pms.reboot(false, reason, false);
        } catch (RemoteException ex) {
        }
    }
----

接收到ACTION_REBOOT广播后，通过 powerManager 执行reboot操作。

== Watchdog 的启动
经过上面的创建和初始化，下面进入到[blue]#watchdog.getInstance().start()# 执行watchdog线程的run()方法，进行watchdog的检测和处理过程中。

1. 首先基于for循环遍历出所有HandlerCheckers，并执行所有HandlerCheckers的[orange]#scheduleCheckLocked()# 的方法，为每个HandlerChecker记录当前的mStartTime
+
[source, java]
----
public void scheduleCheckLocked() {
            if (mMonitors.size() == 0 && mHandler.getLooper().getQueue().isPolling()) {
              
                mCompleted = true; // 当目标looper正在轮询状态则返回
                return;
            }

            if (!mCompleted) {
                // 当有一个check正在处理中，则无需重复发送
                return;
            }

            mCompleted = false;
            mCurrentMonitor = null;
            // 记录当下时间
            mStartTime = SystemClock.uptimeMillis();
            //发送消息，插入消息对列最开头
            mHandler.postAtFrontOfQueue(this);
    }
        public void run() {
            final int size = mMonitors.size();
            for (int i = 0 ; i < size ; i++) {
                synchronized (Watchdog.this) {
                    mCurrentMonitor = mMonitors.get(i);
                }
                // 回调相应service的monitor方法
                mCurrentMonitor.monitor();
            }

            synchronized (Watchdog.this) {
                mCompleted = true;
                mCurrentMonitor = null;
            }
    }
----
+
该方法主要功能：向watchdog的监控的looper池的最头部执行该HandlerChecker.run()方法，在该方法中调用monitor()，
执行完成后设置mCompleted=true,当handler消息池当前的消息，导致没有机会执行monitor方法，则会触发watchdog。
+
[red]#这里简单介绍下安卓的消息机制:Runnable和Message可以被压入Looper维护的MessageQueue的消息对列中，然后Looper循环去执通过Handler执行任务，便会不停地从MessageQueue中取出Runnable或者Message。# 
+
即以上代码的 [orange]#mHandler.postAtFrontOfQueue(this);# 会调用handler的sendMessageAtFrontOfQueue(getPostMessage(r)) 来取出当前线程执行。就会调用下面的run方法，从而回调监控的service的monitor方法
+
其中postAtFrontOfQueue(this),该方法输入参数为Runnable对象，根据消息机制，最终会回调HandlerChecker中的run方法，该方法会循环
遍历所有Watchdog.Monitor 接口，具体服务实现该接口。


2. 执行等待30s

3. 调用evaluateCheckerCompletionLocked() 来评估Checker的状态


主要调用关系如下：
evaluateCheckerCompletionLocked() 
Math.max(state, hc.getCompletionStateLocked())
[source, java]
----
private int evaluateCheckerCompletionLocked() {
    int state = COMPLETED;
    for (int i=0; i<mHandlerCheckers.size(); i++) {
        HandlerChecker hc = mHandlerCheckers.get(i);
        state = Math.max(state, hc.getCompletionStateLocked());
    }
    return state;
}
----

[source, java]
----
public int getCompletionStateLocked() {
    if (mCompleted) {
        return COMPLETED;
    } else {
        long latency = SystemClock.uptimeMillis() - mStartTime;
        // mWaitMax默认是60s
        if (latency < mWaitMax/2) {
            return WAITING;
        } else if (latency < mWaitMax) {
            return WAITED_HALF;
        }
    }
    return OVERDUE;
}
----

依据 getCompletionStateLocked() 函数，可以计算出mHandlerCheckers队列中等待的状态值最大的state属于下面哪个值，以便进行进一步处理

* COMPLETED = 0：等待完成；
* WAITING = 1：等待时间小于DEFAULT_TIMEOUT的一半，即30s；
* WAITED_HALF = 2：等待时间处于30s~60s之间；
* OVERDUE = 3：等待时间大于或等于60s。


以上三步就是执行了watchdog的检测机制

=== Watchdog 处理过程
完成了watchdog的监测机制后，下面就进一步执行watchdog的处理机制
[source, java]
----
 public void run() {
        boolean waitedHalf = false;
        while (true) {
            final List<HandlerChecker> blockedCheckers;
            final String subject;
            final boolean allowRestart;
            int debuggerWasConnected = 0;
            synchronized (this) {
              ....
                    // 如果超时才会进入该分支，即获取被阻塞的checkers
                    blockedCheckers = getBlockedCheckersLocked();
                    // 获取描述信息
                    subject = describeCheckersLocked(blockedCheckers);
                } else {
                    blockedCheckers = Collections.emptyList();
                    subject = "Open FD high water mark reached";
                }
                // mAllowRestart = true
                allowRestart = mAllowRestart;
            }
            EventLog.writeEvent(EventLogTags.WATCHDOG, subject);
            ArrayList<Integer> pids = new ArrayList<>();
            pids.add(Process.myPid());
            if (mPhonePid > 0) pids.add(mPhonePid);
        
            final File stack = ActivityManagerService.dumpStackTraces(
                    !waitedHalf, pids, null, null, getInterestingNativePids());
            // 阻塞 2s
            SystemClock.sleep(2000);

            // 触发kernel来dump所有阻塞线程
            doSysRq('w');
            doSysRq('l');

            // 新建dropbox线程并输出dropbox信息
            Thread dropboxThread = new Thread("watchdogWriteToDropbox") {
                    public void run() {
                        mActivity.addErrorToDropBox(
                                "watchdog", null, "system_server", null, null,
                                subject, null, stack, null);
                    }
                };
            dropboxThread.start();
            try {
                dropboxThread.join(2000);  // wait up to 2 seconds for it to return.
            } catch (InterruptedException ignored) {}

            IActivityController controller;
            synchronized (this) {
                controller = mController;
            }
            if (controller != null) {
                // 将阻塞信息报告给 IActivityController
                Slog.i(TAG, "Reporting stuck state to activity controller");
                try {
                    Binder.setDumpDisabled("Service dumps disabled due to hung system process.");
                    // 1 = keep waiting, -1 = kill system
                    int res = controller.systemNotResponding(subject);
                    if (res >= 0) {
                        Slog.i(TAG, "Activity controller requested to coninue to wait");
                        waitedHalf = false;
                        continue;
                    }
                } catch (RemoteException e) {
                }
            }
          // 只有在debugger没有attach时，裁杀死进程
            if (Debug.isDebuggerConnected()) {
                debuggerWasConnected = 2;
            }
            if (debuggerWasConnected >= 2) {
                Slog.w(TAG, "Debugger connected: Watchdog is *not* killing the system process");
            } else if (debuggerWasConnected > 0) {
                Slog.w(TAG, "Debugger was connected: Watchdog is *not* killing the system process");
            } else if (!allowRestart) {
                Slog.w(TAG, "Restart not allowed: Watchdog is *not* killing the system process");
            } else {
                Slog.w(TAG, "*** WATCHDOG KILLING SYSTEM PROCESS: " + subject);
                //打印堆栈信息
                WatchdogDiagnostics.diagnoseCheckers(blockedCheckers);
                Slog.w(TAG, "*** GOODBYE!");
                // 杀死system_server进程
                Process.killProcess(Process.myPid());
                System.exit(10);
            }

            waitedHalf = false;
        }
    }

----

watchdog检测到异常的信息收集工作：

* ActivityManagerService.dumpStackTraces()
* doSysRq('w');
* doSysRq('l');
* dropboxThread.run().mActivity.addErrorToDropBox()

[source, java]
----
public static File dumpStackTraces(boolean clearTraces, ArrayList<Integer> firstPids, ProcessCpuTracker processCpuTracker, SparseArray<Boolean> lastPids, String[] nativeProcs) {
    //默认为 data/anr/traces.txt
    String tracesPath = SystemProperties.get("dalvik.vm.stack-trace-file", null);
    if (tracesPath == null || tracesPath.length() == 0) {
        return null;
    }

    File tracesFile = new File(tracesPath);
    try {
        //当clearTraces，则删除已存在的traces文件
        if (clearTraces && tracesFile.exists()) tracesFile.delete();
        //创建traces文件
        tracesFile.createNewFile();
        // -rw-rw-rw-
        FileUtils.setPermissions(tracesFile.getPath(), 0666, -1, -1);
    } catch (IOException e) {
        return null;
    }
    //输出trace内容
    dumpStackTraces(tracesPath, firstPids, processCpuTracker, lastPids, nativeProcs);
    return tracesFile;
}
----
输出system_server和mediaserver,/sdcard,surfaceflinger这3个native进程的traces信息

=== dropBox
触发dropbox输出文件到/data/system/dropbox ，内容是traces和相应的blocked信息。

=== systemserver 进程重启

收集完以上信息后调用WatchdogDiagnostics.diagnoseCheckers(blockedCheckers);打印堆栈信息，然后Process.killProcess(Process.myPid());

systemserver 死了之后，会执行信号处理机制，向init进程发送子进程退出信号，从而zygote也会杀死自己。init进程会重新启动zygote->systemserver 以达到重新启动系统的目的。
