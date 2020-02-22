package metrics

import "testing"

func TestMetricsCollector_ReocrdRequest(t *testing.T) {
	s:= &RedisStorage{}
	mc := Init(s)
	mc.ReocrdRequest(RequestInfo{
		ApiName:   "login",
		RespTime:  12,
		Timestamp: 123,
	})
}
