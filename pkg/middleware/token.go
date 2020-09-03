package middleware

import (
	"strconv"

	"goshop/admin-api/pkg/db"
	"goshop/admin-api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if len(token) == 0 {
			c.Abort()
			c.JSON(401, utils.ResponseList{
				RunTime: 0,
				Code:    -110,
				Message: "请登录",
				Data: struct {
				}{},
			})
			return
		}

		u, err := utils.ValidateToken(token)
		if err != nil || u == nil {
			c.Abort()
			c.JSON(401, utils.ResponseList{
				RunTime: 0,
				Code:    -110,
				Message: "验证失败，重新登录。",
				Data: struct {
				}{},
			})
			return
		}

		userTokenKey := utils.UserTokenKey(u.UserId)
		uToken := db.Redis.Get(userTokenKey).Val()
		if len(uToken) == 0 || uToken != token {
			c.Abort()
			c.JSON(401, utils.ResponseList{
				RunTime: 0,
				Code:    -110,
				Message: "token丢失，重新登录。",
				Data: struct {
				}{},
			})
			return
		}

		idStr := strconv.FormatUint(u.UserId, 10)
		c.Set("goshop_user_id", idStr)
	}
}
