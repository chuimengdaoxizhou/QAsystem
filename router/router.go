package router

import (
	"github.com/gin-gonic/gin"
	"qasystem/R_L/controllers"
	"qasystem/R_L/middlewares"
)

func SetupRouter() *gin.Engine {
	// 初始化一个引擎
	r := gin.New()
	// 注册全局中间件
	r.Use(middlewares.AuthMiddleWare())
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Loginuser)
		auth.POST("/register", controllers.RegisterUser)
	}

	return r
}
