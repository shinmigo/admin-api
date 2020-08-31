package engine

import (
	"goshop/api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func NewGinDefault() *gin.Engine {
	r := gin.Default()
	//全局中间件
	r.Use(middleware.Cors())

	r.Static("/static", "./static")

	R = r

	return r
}
