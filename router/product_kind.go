package router

import (
	"github.com/gin-gonic/gin"
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"
	"goshop/api/pkg/middleware"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("product-kind", new(controller.ProductKind), r,  middleware.VerifyToken())
		g.Get("/index")
		g.Post("/add")
		g.Post("/delete")
		g.Post("/edit")
		g.Post("/bind-param")
		g.Post("/bind-spec")
	})
}
