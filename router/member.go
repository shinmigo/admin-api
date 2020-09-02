package router

import (
	"github.com/gin-gonic/gin"
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"
	"goshop/api/pkg/middleware"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("member", new(controller.Member), r, middleware.VerifyToken())
		g.Get("/index")
		g.Post("/add")
		g.Post("/edit")
		g.Get("/info")
		g.Post("/edit-status")
	})
}
