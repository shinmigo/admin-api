package router

import (
	"github.com/gin-gonic/gin"
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"
	"goshop/api/pkg/middleware"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("product-spec", new(controller.ProductSpec), r, middleware.Cors())
		g.Get("/index")
		g.Post("/add")
		g.Post("/edit")
		g.Post("/delete")
	})
}
