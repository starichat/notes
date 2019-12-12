package auth

import "fmt"

// 获取持久化对象
var myStorage MyStorage



// 根据url执行认证
func auth(url string) {
	apiRequest := CreateFromFullUrl(url)
	authCommon(&apiRequest)
}

func authCommon(a *ApiRequest){
	id := a.appId
	token := a.token
	ts := a.timestamp
	url := a.baseUrl

	clientToken := Init(token,ts)

	if clientToken.isExpired(){
		fmt.Println("Token is out of date")
	}
	password := myStorage.getStorage(id)

	sToken := Genrate(url,id,ts,password)

	// 验证服务端token和客户端token
	if(clientToken.match(sToken)){
		fmt.Println("pass auth")
	}

}

