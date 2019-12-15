package token

import (
"crypto/sha256"
"fmt"
"io"
	"log"
	"strconv"
	"strings"
"time"
)

type AuthToken struct {
	token               string
	createTime          int64
	expiredTimeInterval int64
}

// init 一个token
func Init(token string, createTime int64) AuthToken{

	var expiredTimeInterval int64 = 1 * 60 * 10
	return AuthToken{token:token,createTime:createTime,expiredTimeInterval:expiredTimeInterval}
}

func InitByexpire(token string, createTime int64, expiredTimeInterval int64) AuthToken{
	return AuthToken{token:token,createTime:createTime,expiredTimeInterval:expiredTimeInterval}
}

func  Create(originalUrl string, appid string, timestamp int64, password string) AuthToken{

	token := Genrate(originalUrl,appid,timestamp,password)
	return Init(token,time.Now().Unix())
}

// 供外部调用生成token的算法
// 加密算法
// TODO:
func  Genrate(originalUrl string, appid string, timestamp int64, password string) string {
	var token strings.Builder
	token.WriteString(originalUrl)
	token.WriteString(appid)
	token.WriteString(password)

	token.WriteString(strconv.FormatInt(timestamp,10))
	h := sha256.New()
	io.WriteString(h, token.String())
	last := fmt.Sprintf("%x", h.Sum(nil))
	return last
}


func (a *AuthToken) GetToken() string {
	return a.token
}
func (a *AuthToken) IsExpired() bool {
	log.Println("createTime: ",a.createTime)
	return a.createTime+a.expiredTimeInterval < time.Now().Unix()

}
func (a *AuthToken) Match(token string) bool {
	log.Println(a.token)
	return a.token == token

}
