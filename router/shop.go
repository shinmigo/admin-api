package router

import (
	"goshop/admin-api/controller"
	"goshop/admin-api/pkg/core/routerhelper"
	"goshop/admin-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		// 非登录
		g := routerhelper.NewGroupRouter("auth", new(controller.Auth), r)
		g.Post("/login")

		// 必须登录
		authLogin := routerhelper.NewGroupRouter("auth", new(controller.Auth), r, middleware.VerifyToken())
		authLogin.Get("/logout")
	})
}
