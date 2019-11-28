# jit 收集热点代码并后台编译成机器码

app 在安装的时候会调用 create_app_data 在 misc/* 创建对应的 profile 文件

## profile 数据收集

app 启动会开启 profiler 线程收集数据

流程：
AT.main()
AT.performLaunchActivity()
LoadAPk.getClassLoader

setupJitProfileSupport()
VMruntime.registerAppInfo
start
run()

getmethodsvisitor
info.addaddmethofsandClasses

当前进程中所有已经load的class，如果这些class是apk中的class，则会被添加profiler中

对于要记录的method需要满足以下条件：
1. 函数调用需要有以下threshold
2. 


[参考链接]
https://blog.csdn.net/hl09083253cy/article/details/78418809