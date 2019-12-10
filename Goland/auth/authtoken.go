package auth

import (
	"crypto/sha256"
	"fmt"
	"io"
	"strings"
)

type AuthToken struct {
	token               string
	createTime          int64
	expiredTimeInterval int64
}

func (a *AuthToken) new(token string, createTime int64) AuthToken{

	return AuthToken{"ad",1,1}
}

func (a *AuthToken) newByexpire(token string, createTime int64, expiredTimeInterval int64) {

}

func (a *AuthToken) create(baseUrl string, createTime int64) {

}

func (a *AuthToken) Genrate(originalUrl string, appid string, timestamp int64, password string) string {
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
	return a.createTime+a.expiredTimeInterval < 10000

}
func (a *AuthToken) match(au *AuthToken) bool {
	return a.token == au.token

}
