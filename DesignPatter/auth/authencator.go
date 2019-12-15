package auth

import (
	"auth/apirequest"
	"auth/token"
	"log"
)


// 根据url执行认证
func auth(url string) {
	apiRequest := apirequest.CreateFromFullUrl(url)
	log.Println(apiRequest)
	authCommon(&apiRequest)
}

func authCommon(a *apirequest.ApiRequest){
	id := a.GetAppId()
	ts := a.GetTimestamp()
	mtoken := a.GetToken()
	url := a.GetBaseUrl()

	clientToken := token.Init(mtoken,ts)
	log.Println(clientToken)

	if clientToken.IsExpired(){
		log.Println("token is out of date")
	}
	password := "123"

	sToken := token.Genrate(url,id,ts,password)
	log.Println(sToken)


	// 验证服务端token和客户端token
	if !clientToken.Match(sToken) {
		log.Println("authencator is failed")
	}
	log.Println("authencator is passed")


}
