package auth

import "fmt"

// 获取持久化对象
var myStorage MyStorage



// 根据url执行认证
func auth(url string) {
	apiRequest := CreateFromFullUrl(url)
	authCommon(apiRequest)
}

func authCommon(a ApiRequest){
	id := a.AppId
	token := a.Token
	ts := a.Timestamp
	url := a.BaseUrl

	//clientToken := AuthToken.new(token,ts,ts)

	//if clientToken.isExpired(){
	//	fmt.Println("Token is expired")
	//}
	b := new(AuthToken)
	password := myStorage.getStorage(id)

	sToken := b.Genrate(url,id,ts,password)

	// 验证服务端token和客户端token

}

