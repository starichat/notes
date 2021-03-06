# 接口调用统计框架

> 需求：开发一个小框架，能够获取接口调用的各种统计信息，比如，响应时间的最大值/最小值/平均值/百分位值/接口调用次数/频率等。
并且支持将统计结果输出到各种终端，以便于查看。

> 目标：
>1. 掌握线框图的需求分析手段
>2. 掌握测试驱动设计原则和最小原型设计原则
>3. 掌握框架开发的设计思路

## 进度与版本迭代
* 最小版本迭代设计 [tag](df8bd7407240f24eadf0f2945c8cde65bba7cb2f) 100% 
* 面向对象设计设计
    1. 类图，时序图，block图设计
    2. 
   
## 需求分析
性能计数器作为一个跟业务无关的功能，我们完全可以把它开发成一个独立的框架或者类库，集成到很多业务系统中。

从需求中拆解出以下功能：
* 接口统计信息：包括接口响应时间的统计信息，以及接口调用次数的统计信息
* 统计信息的类型：max/min/avg/percentile/count/tps
* 统计信息显示格式：Json/Html。。。。
* 统计信息显示终端：Console。。。。

##  线框图分析
命令行数据：
```json
{
  "api1": {
    "max": 239,
    "min": 10,
    "avg": 124,
    "p99": 265,
    "p999": 309,
    "count": 6789,
    "tps": 101
  },
  "api2": {
    //...  
  } 
}
```
* 统计触发方式
* 统计时间区间
* 统计时间间隔

>非功能性设计
>* 易用性：跟业务代码是否松耦合，提供的接口是否够灵活
>* 性能：框架本身对cpu，对内存消耗很低，并且要低延时
>* 扩展性：在不修改或者尽量少修改代码的情况下添加新功能。使用者在不修改框架源码的情况下，为框架扩展新的功能。
>* 容错性：对外抛出的所有运行时/非运行时异常都进行捕获处理
>* 通用性：。。。
>
## 功能设计
这里结合TDD和Prototype的思想，聚焦于一个简单的应用场景，基于此设计实现一个简单的原型。

例如：这个性能计数器的小框架就可以聚焦于于一个小功能，比如统计用户注册、登陆这两个接口的响应时间的max、min、avg，并将该结果以json格式
输出到命令行中。

> 要输出接口的响应时间的最大值、平均值和接口调用次数，我们首先要采集每次接口请求的响应时间，并且存储起来，然后按照某个时间间隔做聚合统计，最后才是将结果输出。在原型系统的代码实现中，我们可以把所有代码都塞到一个类中，暂时不用考虑任何代码质量、线程安全、性能、扩展性等等问题，怎么简单怎么来就行。

最小原型设计：
响应时间、 响应时间戳分别用来记录[接口]请求的响应时间和访问时间,并通过一个定时调度任务定时执行。
### 综合分析：
根据以上最小原型的分析，我们能够得到以下系统设计方案：

1 数据采集：获取响应时间、访问时间。数据采集过程要高度容错，不能影响到接口本身的可用性。
2 存储：将数据存储到内存/数据库/日志/文件中。为了尽量地减少对接口性能（比如响应时间）的影响，采集和存储的过程异步完成。
3 聚合统计：max，min。。。。。将原始数据聚合为统计数据。
4 显示：json。。。。负责将统计数据以某种格式显示到终端。

## 详细设计
通过以上设计，大概列出了以下方案：
1. 划分职责
1. MetricsCollector 采集数据
2. MetricsStorage 存储
3. Aggregator 聚合统计
4. Report 显示报告

### 类之间的接口关系定义
// TODO

## restful api 接口定义

### request 
HTTP/1.1 200 OK
Content-Type: application/json
{
    
}

### response 

## 设计原则
依赖注入：不同组件通过一个统一的interface获取其他组件中的对象和状态。


