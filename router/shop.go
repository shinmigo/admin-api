package router

import (
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"
	"goshop/api/pkg/middleware"
	
	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		// 非登录
		g := routerhelper.NewGroupRouter("auth", new(controller.Auth), r, middleware.Cors())
		g.Post("/login")
		
		// 必须登录
		authLogin := routerhelper.NewGroupRouter("auth", new(controller.Auth), r, middleware.Cors(), middleware.VerifyToken())
		authLogin.Get("/logout")
	})
}
