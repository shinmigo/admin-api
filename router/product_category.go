package router

import (
	"github.com/gin-gonic/gin"
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"
	"goshop/api/pkg/middleware"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("product-category", new(controller.ProductCategory), r, middleware.Cors())
		g.Get("/index")
		g.Post("/add")
		g.Post("/edit")
		g.Post("/edit-status", "EditStatus")
		g.Post("/delete")
	})
}
