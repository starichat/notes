package credit

import (
	"github.com/starichat/notes/DesignPatter/credit_sys/config"
	"github.com/starichat/notes/DesignPatter/credit_sys/dao"
	"log"
	"testing"
	"time"
)

func Test_AddCredit(t *testing.T) {
	// 初始化配置
	config.Init()
	// 初始化 db 连接
	db := dao.InitDB()
	defer db.Close()
	c:=&Credit{
		ChannelId:   "4sff",
		EventId:     "43334",
		Credit:      "2skeaj",
		CreatedTime:  time.Now(),
		ExpiredTime: time.Now(),
	}
	c.AddCredit()
}

func Test_UpdateCredit(t *testing.T) {
	// 初始化配置
	config.Init()
	// 初始化 db 连接
	db := dao.InitDB()
	defer db.Close()
	c:=&Credit{
		Id:   1,
		ChannelId:   "24241",
		EventId:     "4422",
		Credit:      "aa",
		CreatedTime:  time.Now(),
		ExpiredTime: time.Now(),
	}
	UpdateCredit(c)
}

func Test_FindTotalCredit(t *testing.T) {
	// 初始化配置
	config.Init()
	// 初始化 db 连接
	db := dao.InitDB()
	defer db.Close()
	credits ,err := FindCreditById(1)
	if err != nil {
		log.Println(err)
	}
	log.Println(credits)
}

func Test_FindLimitCredit(t *testing.T) {
	// 初始化配置
	config.Init()
	// 初始化 db 连接
	db := dao.InitDB()
	defer db.Close()
	credits ,err := FindLimitCredit(1,2)
	if err != nil {
		log.Println(err)
	}
	log.Println(credits)
}

func Test_DeleteCredit(t *testing.T) {
	// 初始化配置
	config.Init()
	// 初始化 db 连接
	db := dao.InitDB()
	defer db.Close()
	DeleteCredit(32)
}

