package metrics

import (
	"log"
	"math"
	"sort"
)

// 聚合统计
func Aggregate(requestInfos []RequestInfo, durationInMillis int64) (stat RequestStat){
	stat = RequestStat{
		MaxResponseTime:  math.SmallestNonzeroFloat64,
		MinResponseTime:  math.MaxFloat64,
		AvgResponseTime:  0,
		P99ResponseTime:  0,
		P999ResponseTime: 0,
		Count:            0,
		Tps:              0,
	}
	var sumTime float64 = 0
	
	for _,v := range requestInfos {
		log.Printf("before %f", v.RespTime)
		stat.Count++
		if v.RespTime > stat.MaxResponseTime {
			stat.MaxResponseTime = v.RespTime
		}
		if v.RespTime < stat.MinResponseTime {
			stat.MinResponseTime = v.RespTime
		}
		sumTime += v.RespTime
	}
		if (stat.Count!=0){
			stat.AvgResponseTime = sumTime / float64(stat.Count)
		}
		stat.Tps = stat.Count / durationInMillis * 1000

		// 排序 requestInfos
		sort.Sort(RequestWrapper{requestInfos, func (p, q *RequestInfo) bool {
			return q.RespTime > p.RespTime    // Age 递减排序
		}})
		for _,vv := range requestInfos {
			log.Printf("sort %f" ,vv.RespTime)
		}
		idx999 := int(float64(stat.Count) * 0.999)
		idx99 := int(float64(stat.Count) * 0.99)
		if (stat.Count != 0) {
			stat.P99ResponseTime = requestInfos[idx99].RespTime
			stat.P999ResponseTime = requestInfos[idx999].RespTime
		}

	return stat

}
