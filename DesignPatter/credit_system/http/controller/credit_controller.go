package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/starichat/notes/DesignPatter/credit_system/config"
	"github.com/starichat/notes/DesignPatter/credit_system/model"
	"github.com/starichat/notes/DesignPatter/credit_system/pkg"
	"github.com/starichat/notes/DesignPatter/credit_system/service"
	"net/http"
	"time"
)

const(
	EXPIREMENT  = 10*24*60*60*1000000000 // 10 天
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}



func Attendance(c *gin.Context) {
	// 初始化配置
	cf := config.NewDBConfig()
	// 初始化 db 连接
	s := service.New(cf)
	var request *model.Attendance
	if c.BindJSON(&request) != nil {
		if request.UserId == 0  {
			c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}
	credit := &model.CreditReq{
		UserId:      request.UserId,
		ChannelId:   pkg.GenerateChannleId(request.UserId, 0, 0),
		EventId:     pkg.GenerateEventId(request.UserId, 0),
		Credit:      5,
		ExpiredTime: time.Now().Add(EXPIREMENT),
	}
	result := s.EarnCredit(credit)
	c.JSON(http.StatusOK, gin.H{
		"msg": result,
	})



}

func consume(c *gin.Context){

}