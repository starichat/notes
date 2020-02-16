package dao

import (
	"github.com/starichat/notes/DesignPatter/credit_system/config"
	"github.com/starichat/notes/DesignPatter/credit_system/model"
	"testing"
)

func Test_AddUserCredit(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	userCredit := &model.UserCredit{
		Id:       1,
		UserId:   1,
		CreditId: 1,
	}
	db.AddUserCredit(userCredit)
}

func Test_UpdateUserCredit(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	userCredit := &model.UserCredit{
		Id:       1,
		UserId:   1,
		CreditId: 2,
	}
	db.UpdateUserCredit(userCredit)
}

func Test_DeleteUserCredit(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	db.DeleteUserCredit(2)
}