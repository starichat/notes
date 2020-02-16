package service

import (
	"github.com/starichat/notes/DesignPatter/credit_system/config"
	"github.com/starichat/notes/DesignPatter/credit_system/model"
	"testing"
)

func Test_EarnCredit(t *testing.T) {
	// 初始化配置
	c := config.NewDBConfig()
	// 初始化 db 连接
	s := New(c)
	req := &model.Req{
		UserId:    1,
		ChannelId: "213213",
		EventId:   "133131",
	}
	s.EarnCredit(req)

}