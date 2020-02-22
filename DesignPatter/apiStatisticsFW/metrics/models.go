package metrics


type RequestInfo struct {
	ApiName string
	RespTime float64
	Timestamp int64
}

type RequestStat struct {
	MaxResponseTime float64
	MinResponseTime float64
	AvgResponseTime float64
	P99ResponseTime float64
	P999ResponseTime float64
	Count int64
	Tps int64
}

// 包装类
type RequestWrapper struct {
	requests []RequestInfo
	by func(p, q * RequestInfo) bool
}
// 重写 Len() 方法
func (pw RequestWrapper) Len() int {
	return len(pw.requests)
}

// 重写 Swap() 方法
func (pw RequestWrapper) Swap(i, j int){
	pw.requests[i], pw.requests[j] = pw.requests[j], pw.requests[i]
}

// 重写 Less() 方法
func (pw RequestWrapper) Less(i, j int) bool {
	return pw.by(&pw.requests[i], &pw.requests[j])
}