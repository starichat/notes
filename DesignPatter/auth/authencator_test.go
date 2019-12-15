package auth

import (
	"auth/token"
	"strconv"
	"testing"
	"time"
)

func Test_auth(t *testing.T) {
	// 生成一个请求
	ts := time.Now().Unix()
	Ttoken := token.Genrate("http://www.starichat.com/api", "starichat", ts, "123")
	t.Log(Ttoken)
	t.Log(ts)
	ctoken := token.Init(Ttoken, ts)
	curl := "http://www.starichat.com/api?" + "appId=starichat&" + "token=" + ctoken.GetToken() + "&" + "timestamp=" + strconv.FormatInt(ts,10)
	//capi := apirequest.CreateFromFullUrl(curl)
	t.Log(curl)
	//outdateurl := "http://www.starichat.com/api?appId=starichat&token=f4be34f3ef4afb0333c2806d48314c73f1b7d339af99b6aa7479b95a5d63ab30&timestamp=1576314776"
	//fmt.Println(capi)
	auth(curl)
	//auth(outdateurl)
}
