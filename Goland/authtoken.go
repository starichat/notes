package auth

import (
	"crypto/sha256"
	"fmt"
	"io"
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

	return AuthToken{token:token,createTime:createTime,expiredTimeInterval:nil}
}

func InitByexpire(token string, createTime int64, expiredTimeInterval int64) AuthToken{
	return AuthToken{token:token,createTime:createTime,expiredTimeInterval:expiredTimeInterval}
}

func (a *AuthToken) Create(baseUrl string, createTime int64, params map[string]string) AuthToken{

	return AuthToken{}
}

func  Genrate(originalUrl string, appid string, timestamp int64, password string) string {
	var token strings.Builder
	token.WriteString(originalUrl)
	token.WriteString(appid)
	token.WriteString(password)
	token.WriteString(string(timestamp))
	h := sha256.New()
	io.WriteString(h, token.String())
	last := fmt.Sprintf("%x", h.Sum(nil))
	return last
}

func (a *AuthToken) isExpired() bool {
	return a.createTime+a.expiredTimeInterval < time.Now().Unix()

}
func (a *AuthToken) match(token string) bool {
	return a.token == token

}
