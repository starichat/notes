# Go 语言性能优化工具

## GoRerport 代码质量评估报告

## dingo-hunter 在Go程序中找出deadlocks的静态分析起
## flen 获取函数长度信息
## safesql 防止sql注入

# 实战
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



id
药品名称
生产厂家
生成日期
保质期
用途
规格
进货价
销售价
员工编号
客户编号

CREATE TABLE [medication_info] (
[id]	INT	NOT NULL IDENTITY(1,1) PRIMARY KEY，
[name]	VARCHAR(20)	NOT NULL ,
[type]	VARCHAR(20)	NOT NULL,
[manufacturer]	VARCHAR(20)	NOT NULL,
[manu_date]	DATETIME	NOT NULL,
[deadline]	DATETIME	NOT NULL,
[purchase_price]	float	NOT NULL,
[sale_price]	float	NOT NULL
)

CREATE TABLE [staff] (
[id]	INT	NOT NULL IDENTITY(1,1) PRIMARY KEY,
[name]	VARCHAR(20)	NOT NULL,
[age]	VARCHAR(20)	NOT NULL,
[education]	VARCHAR(20)	NOT NULL,
[position]	VARCHAR(20)	NOT NULL  
)

CREATE TABLE [purchase] (
[id]	INT	NOT NULL IDENTITY(1,1) PRIMARY KEY,
[purchase_time]	DATETIME	NOT NULL,
[number]	INT	NOT NULL,
[medication_id]	INT	NOT NULL,
[staff_id]	INT	NOT NULL
)

CREATE TABLE [sale] (
[id]	INT	NOT NULL IDENTITY(1,1) PRIMARY KEY,
[created_time]	DATETIME	NOT NULL,
[number]	INT	NOT NULL,
[medication_id]	INT	NOT NULL,
[staff_id]	INT	NOT NULL,
[member_id]	INT	NOT NULL 
)

CREATE TABLE [member] (
[id]	INT	NOT NULL IDENTITY(1,1) PRIMARY KEY,
[name]	VARCHAR(20)	NOT NULL,
[age]	INT	NOT NULL,
[mobile]	VARCHAR(11)	NOT NULL
)
CREATE TABLE [supplier] (
[id]]	INT	NOT NULL IDENTITY(1,1) PRIMARY KEY,
[supplier_name]	VARCHAR(20)	NOT NULL
)

存储过程：
create procedure INSERT_YAOPIN
@StuNo    nvarchar(64)='001'        --设置默认值
as
begin
    INSERT YAOPIN VALUE(S_StuNo=@StuNo,S_StuNo=@StuNo,S_StuNo=@StuNoS_StuNo=@StuNo)
end





员工编号
姓名
年龄
学历
职务


订单编号
进货时间
数量
药品编号
员工编号
会员编号

编号
进货时间
数量
药品编号
员工编号
会员编号

会员编号
姓名
年龄
联系电话

# 存储过程
CREATE VIEW [out_of_date_view]
AS
SELECT * FROM [medication_info] 
WHERE DATEDIFF(day, [deadline], getDate()) <= 7

# 
Create Procedure Add_Member
(@Name Nvarchar(20)='',
@AGE NINT=0,
@MOBILE NVACCHAR(20)=''
)
as
begin
    INSERT [dbo].[memeber] (name,age,mobile)
    VALUES(@NAME,@AGE,@MOBILE)
end
GO

Create Procedure Add_Medication_Member
(@Name Nvarchar(20)='',
@AGE NINT=0,
@EDUCATION NVACCHAR(20)='',
@POSITION NVACCHAR(20)='',
)
as
begin
    INSERT [dbo].[tb_Demo_MultiRowsInsert] ([name],[age],[education],[position])
    VALUES(@NAME,@AGE,@EDUCATION,@POSITION)
end
GO


Create Procedure Add_Supplier
(@Name Nvarchar(20)='',
)
as
begin
    INSERT [dbo].[tb_Demo_MultiRowsInsert] (supplier_name)
    VALUES(@NAME)
end
GO

Create Procedure Add_Medication
(@Name Nvarchar(20)='',
@TYPE Nvarchar(20)='',
@SUPPIER NVACCHAR(20)='',
@MANUDATE ,
@DEADLINE NVACCHAR(20)='',
@PURCHASEPRICE FLOAT,
@SALEPRICE FLOAT,
)
as
begin
    INSERT [dbo].[tb_Demo_MultiRowsInsert] ([name],[type],[manufacturer],[manu_date],[deadline],[purchase_price],[sale_price])
    VALUES(@NAME,@TYPE,@SUPPIER,@MANUDATE,@DEADLINE,@PURCHASEPRICE,@SALEPRICE)
end
GO

Create Procedure Add_Stock
(@STOCKTIME Nvarchar(20)='',
@NUMBER NVACCHAR(20)='',
@MEDICATIONID ,
@STAFFID NVACCHAR(20)='',
)
as
INSERT [dbo].[tb_Demo_MultiRowsInsert] ([purchase_time],[number],[medication_id],[staff_id])
VALUES(@STOCKTIME,@NUMBER,@MEDICATIONID,@STAFFID)
GO

Create Procedure Add_Order
(@ORDERTIME Nvarchar(20)='',
@NUMBER NVACCHAR(20)='',
@MEDICATIONID ,
@STAFFID NVACCHAR(20)='',
@MEMBERID NVACCHAR(20)='',
)
as
begin
    INSERT [dbo].[order] ([created_time],[number],[medication_id],[staff_id],[member_id])
    VALUES(@NAME,@NUMBER,@MEDICATIONID,@STAFFID,@MEMBERID)
end
GO


CREATE PROCEDURE Sales.uspGetEmployeeSalesYTD  
AS    
   SET NOCOUNT ON;  
   SELECT LastName, SalesYTD  
   FROM Sales.SalesPerson AS sp  
   JOIN HumanResources.vEmployee AS e ON e.BusinessEntityID = sp.BusinessEntityID   
RETURN  


SELECT * from SHANGPINGBIAO WHERE TIME < STARTTIME AND TIME > ENDTIMME

DELIMITER $$
USE `test_ drugstore`$$
DROP TRIGGER /*!50032 IF EXISTS */ `insert_出售`$$
CREATE
    /*!50017 DEFINER = 'root'@'localhost' */
    TRIGGER `insert_出售` before INSERT ON `order` 
    FOR EACH ROW BEGIN
call update_date 
if num < kuocun
roollback
END;
$$
DELIMITER ;
### 
select * from stock where created_time betweenn [startTime] and [endTime]

SELECT medication.name, medication.deadline FROM YAOPIN WHERE ID = @ID 

SELECT medication.id, medication.name, sum((m.sale_price-m.purchase_price)*s.number) from
medication, sale where medication.id = sale.medication_id where s.time betweenn [startTime] and [endTime]

SELECT medication.id, medication.name,s.* from
medication, sale where medication.id = sale.medication_id where s.time betweenn [startTime] and [endTime]


select m.*,s.number from medication, stock where stock.medication_id = medicaiton.id 