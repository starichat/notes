package metrics

import (
	"encoding/json"
	"fmt"
	"log"
)

type ConsoleReport struct {
	storage Storage
}

func NewReport(metricsStorage Storage) *MetricsCollector{
	return &MetricsCollector{storage:metricsStorage}
}

func (cr *ConsoleReport) startReport() {
	// 从数据库取出内容
	requestInfos := cr.storage.getRequestInfos()
	stats := map[string]RequestStat{}
	// 聚合统计
	for apiName, requestInfosPerApi := range requestInfos{
		requestStat := Aggregate(requestInfosPerApi, 100)
		stats[apiName]=requestStat
	}
	// 显示到控制台
	log.Println("->>>>>>>>>>>>>>>>>start<<<<<<<<<<<<<<<<<<-")
	statsJson, _ := json.Marshal(stats)
	fmt.Println(statsJson)
}

// 定时任务

