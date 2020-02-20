package metrics

import (
	"testing"
	"time"
)

var m Metrics = Metrics{}

// 应用场景测试:登陆接口响应时间
func Test_login(t *testing.T) {
	startTime := time.Now().Unix()
	m.recordResponseTime("login",startTime)
	// todo login
	respTime := time.Now().Unix()-startTime
	m.RecordTimestamp("login",respTime)



}
