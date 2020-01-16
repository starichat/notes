# 积分兑换系统
> 通过开发一个积分兑换系统完美展示软件项目开发整个过程，从中学习项目开发如何运用设计模式的。

 > 需求：积分是一种常见的营销手段，很多产品都会通过它来促进消费、增加用户粘性，比如淘宝积分、信用卡积分、商场消费积分等等，我们需要设计一个积分兑换系统来处理用户的积分。

 ## 需求分析
 最好的需求分析就是通过设计产品的线框图，用例图来进行梳理。接下来，我们结合用例图来进行分析该需求：
 img：用例图

 * 用户获取积分时，告知积分的有效期
 * 用户使用积分时，优先使用快过期的积分
 * 用户在查询积分明细的时候，显示积分有效期和状态
 * 用户查询总可用积分的时候，会排除掉过期的积分

 根据上述用例分析，可以作出如下设计：
 1. 积分获取和兑换规则
 2. 积分消费和兑换规则
 3. 积分及其明细查询

 ## 系统设计
系统设计不同于功能设计，更聚焦于架构的层面，主要是对各个模块的设计，当然可以理解为更大层面的功能设计。功能设计一般是针对某个模块进行设计，实现代码的功能完善，高内聚等。系统设计就是实现模块间的组织、高内聚等。下面就来对该系统进行设计：

|模块|功能|完成度|
|:---:|:---:|:---:|
|积分获取和兑换模块|积分赚取渠道及兑换规则||
|积分消费模块|消费渠道及兑换规则的管理和维护||

## 功能设计
业务开发无外乎有以下三个方面：
1. 接口设计
2. 数据库设计
3. 业务模型设计

前期的接口和数据库设计非常重要，因为这部分迁移或者改动的成本非常大，所以对于数据库的设计需要非常谨慎。

### 接口和数据库设计
整个系统，涉及到持久层（需要存储）的数据无非就是积分数据和用户数据，下面就对此来进行数据库和接口设计：

#### 积分明细表：

|字段名|含义|
|:-:|:-:|
|id|积分明细ID|
|channel_id|积分赚取或消费渠道ID|
|event_id|相关事件id，比如订单id，评论id，优惠卷兑换交易id等|
|credit|积分|
|create_time|积分赚取或消费时间
|expired_time|积分过期时间


> 接口设计：接口设计需符合单一职责原则，粒度越小通用性就越好。但是，粒度不能太小，因为粒度太小就需要设计很多小接口，如果这些接口需要通过网络请求，特别是一些微服务的对外http公网请求接口，会严重影响性能。另外，粒度太小也会破坏事务的原子性。因而，在接口设计方面通常都需要坚固性能和设计原则。

本系统接口如下：

|接口|含义|参数|返回值|
|:---:|:---:|:---:|:---:|
|earnPoint|赚取积分|userId,channelId,eventId,credit,expiredTime|积分明细ID|
|consumePoint|消费积分|userId,channelId,eventId,credit,expiredTime|积分明细ID|
|getPoints|查询积分|userId|总可用积分|
|getPointsDetail|查询总积分明细|userId+分页参数|userId,channelId,eventId,credit,expiredTime|
|getDetailByEarn|查询赚取积分明细|userId+分页参数|userId,channelId,eventId,credit,expiredTime|
|getDetailByConsume|查询消费积分明细|userId+分页参数|userId,channelId,eventId,credit,expiredTime|

2. 业务模型设计

```sql
CREATE TABLE `credits`(
    `id` INT PRIMARY KEY AUTO_INCREMENT NOT NULL COMMENT "订单明细id",
    `channel_id` VARCHAR(32) DEFAULT NULL COMMENT "积分赚取或消费渠道ID",
    `event_id` VARCHAR(32) DEFAULT NULL COMMENT "相关事件id，比如订单id，评论id，优惠卷兑换交易id等", 
    `credit` INT DEFAULT NULL COMMENT "积分",
    `created_time` DATETIME DEFAULT NULL COMMENT "消费时间或者赚取时间",
    `expired_time` DATETIME DEFAULT NULL COMMENT "积分过期时间"
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

