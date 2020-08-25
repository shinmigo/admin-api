package router

import (
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"
	"goshop/api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("product-spec", new(controller.ProductSpec), r, middleware.Cors(), middleware.VerifyToken())
		g.Get("/index")
		g.Post("/add")
		g.Post("/edit")
		g.Post("/delete")
	})
}
