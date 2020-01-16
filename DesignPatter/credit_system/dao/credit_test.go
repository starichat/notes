package dao

import (
	"github.com/starichat/notes/DesignPatter/credit_sys/config"
	"github.com/starichat/notes/DesignPatter/credit_sys/model"
	"log"
	"testing"
	"time"
)

func Test_AddCredit(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	credit := &model.CreditInfo{
		ChannelId:   "4sff",
		EventId:     "43334",
		Credit:      100,
		CreatedTime:  time.Now(),
		ExpiredTime: time.Now(),
	}
	db.AddCredit(credit)
}
//
//func Test_UpdateCredit(t *testing.T) {
//	// 初始化配置
//	config.Init()
//	// 初始化 db 连接
//	db := InitDB()
//	defer db.Close()
//	c:=&Credit{
//		Id:   1,
//		ChannelId:   "24241",
//		EventId:     "4422",
//		Credit:     123,
//		CreatedTime:  time.Now(),
//		ExpiredTime: time.Now(),
//	}
//	UpdateCredit(c)
//}
//
func Test_FindCreditById(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	db := New(c)
	defer db.Close()
	credit ,err := db.FindCreditById(1)
	if err != nil {
		log.Println(err)
	}
	log.Println(credit)
}

//func Test_FindLimitCredit(t *testing.T) {
//	// 初始化配置
//	config.Init()
//	// 初始化 db 连接
//	db := InitDB()
//	defer db.Close()
//	credits ,err := FindLimitCredit(1,2)
//	if err != nil {
//		log.Println(err)
//	}
//	log.Println(credits)
//}
//
//func Test_DeleteCredit(t *testing.T) {
//	// 初始化配置
//	config.Init()
//	// 初始化 db 连接
//	db := InitDB()
//	defer db.Close()
//	DeleteCredit(32)
//}

