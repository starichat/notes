package metrics

import (
	"log"
	"testing"
)

func TestAggregate(t *testing.T) {
	r := RequestInfo{
		ApiName:   "login",
		RespTime:  2,
		Timestamp: 123,
	}
	r1 := RequestInfo{
		ApiName:   "login1",
		RespTime:  1,
		Timestamp: 123,
	}
	r2 := RequestInfo{
		ApiName:   "login1",
		RespTime:  3,
		Timestamp: 123,
	}
	rs := []RequestInfo{}
	rs = append(rs, r)
	rs = append(rs, r1)
	rs = append(rs,r2)
	log.Println(Aggregate(rs,100))
}
