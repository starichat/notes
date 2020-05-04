# MySql 优化

1. sql 执行计划
2. 查询优化会干扰优化


ID：编号
select_type 查询类型
table 表
type类型
possible_keys 预测用到的索引
key 实际用到的索引
key_len 实际使用的索引的长度
ref 表之间的饮用
rows 通过索引查询的数据量
extra


## mysql 多大数据量会dangji

## 千万级数据测试
不加索引和加索引查询数据差别
1. explain select * from teacher where tid = 103213;
15:36:38	explain select * from teacher where tid = 103213	1 row(s) returned	0.00020 sec / 0.0000088 sec

15:37:29	select * from teacher where tid = 11222 LIMIT 0, 1000	1 row(s) returned	0.248 sec / 0.0000060 sec

2. alter table teacher add index indexName(tid);
15:39:33	alter table teacher add index indexName(tid)	0 row(s) affected Records: 0  Duplicates: 0  Warnings: 0	0.758 sec
select * from teacher where id = 212321
15:40:00	select * from teacher where tid = 11222 LIMIT 0, 1000	1 row(s) returned	0.00033 sec / 0.0000081 sec

'1', 'SIMPLE', 'teacher', NULL, 'ref', 'indexName', 'indexName', '5', 'const', '1', '100.00', NULL


'1', 'SIMPLE', 'COURSE', NULL, 'ALL', NULL, NULL, NULL, NULL, '996133', '100.00', 'Using filesort'

## 索引失效的几个方面
1. 查询条件不要有计算
2. like 查询要独立出现
3. 多索引使用顺序不能颠倒

