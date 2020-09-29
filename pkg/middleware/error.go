package middleware

import (
	"fmt"
	"goshop/admin-api/pkg/utils"
	"runtime"

	"github.com/gin-gonic/gin"
)

func NotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Abort()
		c.JSON(200, utils.ResponseList{
			RunTime: 0,
			Code:    -404,
			Message: "接口不存在",
			Data:    struct{}{},
		})

	}
}

func ServerError() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recInfo := recover(); recInfo != nil {
				var str string
				str = fmt.Sprint(recInfo) + "; "
				for i := 1; ; i++ {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					str = str + fmt.Sprintf("%s:%d; ", file, line)
				}

				c.Abort()
				c.JSON(200, utils.ResponseList{
					RunTime: 0,
					Code:    -500,
					Message: str,
					Data:    struct{}{},
				})
			}

		}()

		c.Next()
	}
}
