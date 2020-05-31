# Go 语言性能优化

## 概述
Go 语言以性能著称，为了让我们的GO程序更好地利用资源，这里我们来看看Goland性能优化。
性能优化通常指的是：
1. 性能指标
cpu 内存 IO操作
* 系统吞吐量：每秒可以处理的请求数
* 响应时间：从客户端发出请求，到收到回包的总耗时

2. 定位瓶颈
## benchmark
go test -v gate_test.go -run=none -bench=. -benchtime=3s -cpuprofile cpu.prof -memprofile mem.prof
run 单次测试，一般用于代码逻辑验证
bench=. 执行所有的benchmark，也可以通过用例函数名来指定部分测试用例
benchtime 指定测试执行时长
cpuprofile 输出cpu的pprof信息文件
memprofile 输出所有heap的pprof文件
blockprofile 阻塞分析，记录goroutine阻塞等待同步的位置
mutexprofile 互斥锁分析，报告互斥锁的竞争情况

benchmark 测试用例常用函数
b.ReportAllocs 输出单次循环使用的内存数量和对象allocs信息
b.RunParallel 使用协程并发测试
b.SetBytes(n int64) 设置单次循环使用的内存数量

## pprof
* 生成方式
runtime/pprof: 手动调用如runtime.StartCPUProfile或者runtime.StopCPUProfile等 API 来生成和写入采样文件，灵活性高。主要用于本地测试。
net/http/pprof: 通过 http 服务获取 Profile 采样文件，简单易用，适用于对应用程序的整体监控。通过 runtime/pprof 实现。主要用于服务器端测试。
go test: 通过 go test -bench . -cpuprofile cpuprofile.out生成采样文件，主要用于本地基准测试。可用于重点测试某些函数。
* 查看方式
go tool pprof [options][binary]...
 - --text 纯文本
 - -- Web 
 - --sv
 - --list funcname 筛选出正则匹配funcname的函数信息
 - http=":port" 本地浏览器打开

 go tool pprof -base profile1 profile2
  - 对比查看2个profile，一般用于代码修改前后对比，定位差异点
通过命令行方式查看 profile 时，可以在命令行对话中，使用下列命令，查看相关信息
flat: 采样时，该函数正在运行的次数*采样频率(10ms)，即得到估算的函数运行”采样时间”。这里不包括函数等待子函数返回。

flat%: flat / 总采样时间值

sum%: 前面所有行的 flat% 的累加值，如第三行 sum% = 71.24% = 27.56% + 50.58%

cum: 采样时，该函数出现在调用堆栈的采样时间，包括函数等待子函数返回。因此 flat <= cum
3.
cum%: cum / 总采样时间值

topN [-cum] 查看前 N 个数据：

