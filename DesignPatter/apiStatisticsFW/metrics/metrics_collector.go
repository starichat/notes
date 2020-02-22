package metrics

import "log"

type MetricsCollector struct {
	storage Storage
}

// DI
func Init(metricsStorage Storage) *MetricsCollector{
	return &MetricsCollector{storage:metricsStorage}
}


func (mc *MetricsCollector) ReocrdRequest(info RequestInfo) {

	// 数据预处理
	if &info == nil || info.ApiName==""{
		log.Printf("ERROR:request Info is null")
		return
	}
	log.Printf("ERROR: Info  %p",&mc.storage)
	log.Println(mc.storage)
	mc.storage.SaveRequestInfo(info)





}