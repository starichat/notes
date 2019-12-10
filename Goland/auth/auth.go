package main

type AuthToken struct {
	token               string
	createTime          int64
	expiredTimeInterval int64
}

func (a *AuthToken) AuthToken(token string, createTime int64) {

}

func (a *AuthToken) AuthTokenByexpire(token string, createTime int64, expiredTimeInterval int64) {

}

func (a *AuthToken) create(baseUrl string, createTime int64) {

}

func (a *AuthToken) getToken() {

}
func (a *AuthToken) isExpired() {

}
func (a *AuthToken) match() {

}

// apirequest
type ApiRequest struct {
	baseUrl   string
	token     string
	appId     string
	timestamp int64
}

func initApiRequest(baseUrl, token, appId string, timestamp int64) {

}

func createFromFullUrl(url string) ApiRequest {

}

func getBaseUrl() {

}
func getToken() {

}
func getAppId() {

}

func getTimestamp() {

}

// credentialstorage
type Credentialstorage interface {
	getpasswordByAppId(appid string)
}

// api auth
type ApiAuthencator interface {
	auth(url string)
	authbyApi(api *ApiRequest)
}

func auth(){

	ApiRequest apiRquest = ApiRequest.buildFromUrl(url) // 从请求url中解析出api数据


}

func authByApi(){
	// 从api中获取数据
	// 判断token是否过期
	// 从认证服务器中获取密码
	// 生成认证token
	// 校验token是否匹配
	
}



func main() {

}
