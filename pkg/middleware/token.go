package middleware

import (
	"goshop/api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("GoshopToken")
		//查询缓存或rpc
		val := "hello world"
		if token != val {
			res := utils.ResponseList{
				RunTime: 0,
				Code:    -110,
				Message: "你没有权限登录，请联系管理员。",
				Data: struct {
				}{},
			}

			c.Abort()
			c.JSON(200, res)
		}
	}
}
