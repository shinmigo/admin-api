package router

import (
	"github.com/gin-gonic/gin"
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"
	"goshop/api/pkg/middleware"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("product", new(controller.Product), r, middleware.Cors())
		g.Post("/add")
	})
}
