package metrics

import "log"

type Storage interface {
	SaveRequestInfo(info RequestInfo)
	getRequestInfosByApi(apiName string) *[]RequestInfo
	getRequestInfos() map[string][]RequestInfo
}


type MetricsStorage struct {
}

type RedisStorage struct {
}

func (s *MetricsStorage) SaveRequestInfo(info RequestInfo) {
	log.Printf("i am storage and info's apiName is %s",info.ApiName)
}

func (s *RedisStorage) SaveRequestInfo(info RequestInfo) {
	log.Printf("i am Redis storage and info's apiName is %s",info.ApiName)
}

func (s *RedisStorage) getRequestInfosByApi(apiName string) *[]RequestInfo {
	return &[]RequestInfo{}
}

func (s *RedisStorage) getRequestInfos() map[string][]RequestInfo {
	return map[string][]RequestInfo{}
}