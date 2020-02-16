package dao

import (
	"github.com/starichat/notes/DesignPatter/credit_system/config"
	"github.com/starichat/notes/DesignPatter/credit_system/model"
	"log"
	"testing"
	"time"
)
const(
  EXPIR  = 60*60*1000000000 // one hour
)

func Test_AddCredit(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	credit := &model.CreditInfo{
		ChannelId:   "4s222ff1",
		EventId:     "432334",
		Credit:      10011,
		CreatedTime:  time.Now(),
		ExpiredTime: time.Now().Add(EXPIR),
	}
	log.Println(EXPIR)
	db.AddCredit(credit)
}

func Test_UpdateCredit(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	credit := &model.CreditInfo{
		Id:   3,
		ChannelId:   "2424221",
		EventId:     "442222",
		Credit:     122222,
		CreatedTime:  time.Now(),
		ExpiredTime: time.Now(),
	}
	db.UpdateCredit(credit)
}

func Test_FindCreditById(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	credit ,err := db.FindCreditById(3)
	if err != nil {
		log.Println(err)
	}
	log.Println(credit)
}

func Test_FindLimitCredit(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	credits ,err := db.FindLimitCredit(2,0)
	if err != nil {
		log.Println(err)
	}
	log.Println(credits)
}

func Test_DeleteCredit(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	db.DeleteCredit(1)
}

