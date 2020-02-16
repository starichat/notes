package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/starichat/notes/DesignPatter/credit_system/http/controller"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	//IndexApi为一个Handler
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"msg":"hello world",
		})
	})
	v1 := router.Group("api")
	{
		v1.POST("/attendance",controller.Attendance)
		//v1.POST("credit", controller.GetCredits)
		//v1.GET("get",controller.GetCredits)
	}

	return router
}