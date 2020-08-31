package routerhelper

import (
	"github.com/gin-gonic/gin"
	"goshop/api/pkg/middleware"
)

type RouterFun func(r *gin.Engine)

var rList = make([]RouterFun, 0, 8)

var R *gin.Engine

func Use(p ...RouterFun) {
	rList = append(rList, p...)
}

func EntryRouterTree(e *gin.Engine) {
	for k := range rList {
		rList[k](e)
	}
}

func InitRouter() {
	//r := utils.NewGinDefault()
	r:=gin.Default()
	r.Use(middleware.Cors())
	R = r
	
	r.Static("/static", "./static")
	EntryRouterTree(r)
}
