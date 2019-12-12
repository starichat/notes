package auth

// apirequest
type ApiRequest struct {
	baseUrl   string
	appId     string
	token     string
	timestamp int64
}

func InitApi(baseUrl, token, appId string, timestamp int64) ApiRequest{
	return ApiRequest{baseUrl,appId,token,timestamp}
}

func CreateFromFullUrl(url string) ApiRequest {
	// 解析url
	// 截取第一个/之前的字符串
	// appid=
	// token
	// timestamp
	return ApiRequest{"a","aa","aaa",1}
}

func getBaseUrl(a *ApiRequest) string {
	return a.baseUrl
}
func getToken(a *ApiRequest) string {
	return a.token
}
func getAppId(a *ApiRequest) string {
	return a.appId
}

func getTimestamp(a *ApiRequest) int64{
	return a.timestamp
}
