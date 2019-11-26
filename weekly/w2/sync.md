# go 语言sync包解析

> golang 的并发机制广受热爱，但是当多个goroutine同时进行处理任务，就会遇到抢占资源的情况，某一个goroutine等待另一个goroutine处理完某一个步骤后才能继续。这时候就需要一个锁/共享内存等机制的出现，来帮助goroutine进行协同合作，这个工具就是sync。

## 锁

## 临时对象池

## Once

## WaitGroup 和 Cond

