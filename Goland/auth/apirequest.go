package auth

// apirequest
type ApiRequest struct {
	BaseUrl   string
	Token     string
	AppId     string
	Timestamp int64
}

func initApiRequest(baseUrl, token, appId string, timestamp int64) {

}

func CreateFromFullUrl(url string) ApiRequest {
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
