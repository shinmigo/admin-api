package router

import (
	"goshop/admin-api/controller"
	"goshop/admin-api/pkg/core/routerhelper"
	"goshop/admin-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("user", new(controller.User), r, middleware.Test(), middleware.VerifyToken())
		g.Get("/get-list-query", "GetListQuery")
		g.Get("/test")
	})
}
