package metrics

import (
	"encoding/json"
	"fmt"
)

// data struct for recordTime
var recordTimes map[string]int64

var recordTimestamps map[string]int64

// data struct to show
type Stat struct {
	Max int
	Min int
	Avg int
	P99 int
	P999 int
	Count int
	tps int
}

type Metrics struct {

}

// 开始接口
func (m *Metrics) RecordTimestamp(api string, time int64){
	recordTimestamps[api] = time

}

// 统计时间
func (m *Metrics) recordResponseTime(api string, time int64) {
	recordTimes[api] = time
}

// 定时刷新数据
func (m *Metrics) startRepeatedReport() {
	// 实现一个定时任务，通过参数传递多久执行一次
	// TODO ： 实现逻辑代码
	// 1. 遍历两个map，将数据提取出来显示到json中去
	var results []map[string]Stat
	stat := Stat{
		Max:   0,
		Min:   0,
		Avg:   0,
		P99:   0,
		P999:  0,
		Count: 0,
		tps:   0,
	}
	m2 := map[string]Stat{}
	m2["api"] = stat
	results = append(results,m2)
	data, _ := json.Marshal(results)
	fmt.Println(string(data))
}

// 获取某接口的最大响应时间
func (m *Metrics) Max(api string) int64 {
	return 0
}



