package pkg

import "sync/atomic"

//GetIncreaseID 并发环境下生成一个增长的id,按需设置局部变量或者全局变量
func GetIncreaseID(ID *uint64) uint64 {
	var n, v uint64
	for {
		v = atomic.LoadUint64(ID)
		n = v + 1
		if atomic.CompareAndSwapUint64(ID, v, n) {
			break
		}
	}
	return n
}
