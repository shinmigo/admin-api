package router

import (
	"goshop/api/controller"
	"goshop/api/pkg/core/routerhelper"

	"github.com/gin-gonic/gin"
)

func init() {
	routerhelper.Use(func(r *gin.Engine) {
		g := routerhelper.NewGroupRouter("user2", new(controller.User), r)
		g.Get("/get-list-query2", "GetListQuery")
	})
}
