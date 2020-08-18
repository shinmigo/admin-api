package router

import (
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"
	"goshop/api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("member", new(controller.Member), r, middleware.Cors())
		g.Get("/index")
		g.Get("/add")
		g.Get("/edit")
		g.Get("/info")
		g.Get("/edit-status")
	})
}
