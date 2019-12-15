package apirequest

import (
	"log"
	"strconv"
	"strings"
)

// apirequest
type ApiRequest struct {
	baseUrl   string
	appId     string
	token     string
	timestamp int64
}

func InitApiRequest(baseUrl, token, appId string, timestamp int64) ApiRequest {
	return ApiRequest{baseUrl, appId, token, timestamp}
}

// url : http://xxxx.com/api?appId=xxx&token=xxx&timestamp=xxxxx
func CreateFromFullUrl(url string) ApiRequest {
	// 解析url
	urlinfos := strings.Split(url, "&")
	log.Println("len(url()): ", len(url))
	for _, u := range urlinfos {
		log.Println(u)
	}
	baseUrl := urlinfos[0][:strings.Index(urlinfos[0], "?")]
	// 截取第一个/之前的字符串
	appid := urlinfos[0][strings.Index(urlinfos[0], "=")+1:]
	// appid=
	token := urlinfos[1][strings.Index(urlinfos[1], "=")+1:]
	timestamp := urlinfos[2][strings.Index(urlinfos[2], "=")+1:]
	log.Println("ts: ", timestamp)
	ts, _ := strconv.ParseInt(timestamp, 10, 64)
	return ApiRequest{baseUrl, appid, token, ts}
}

func (a *ApiRequest) GetBaseUrl() string {
	return a.baseUrl
}
func (a *ApiRequest) GetToken() string {
	return a.token
}
func (a *ApiRequest) GetAppId() string {
	return a.appId
}

func (a *ApiRequest) GetTimestamp() int64 {
	return a.timestamp
}
