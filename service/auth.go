package service

import (
	"context"
	"time"
	
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/shoppb"
	"goshop/api/pkg/db"
	"goshop/api/pkg/grpc/gclient"
	"goshop/api/pkg/utils"
)

type UserLogin struct {
	UserId uint64 `json:"user_id"`
	Name   string `json:"name"`
	Token  string `json:"token"`
}

type Auth struct {
	*gin.Context
}

func NewAuth(c *gin.Context) *Auth {
	return &Auth{Context: c}
}

func (a *Auth) Login() (*UserLogin, error) {
	req := &shoppb.LoginReq{
		Username: a.PostForm("username"),
		Password: a.PostForm("password"),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ShopUser.Login(ctx, req)
	cancel()
	
	if err != nil {
		return nil, err
	}
	
	token, err := utils.GenerateToken(resp.UserId, resp.Username)
	if err != nil {
		return nil, err
	}
	
	// token保存在redis中
	redisKey := utils.UserTokenKey(resp.UserId)
	if err := db.Redis.Set(redisKey, token, time.Duration(utils.DEFAULT_EXPIRE_SECONDS)*time.Second).Err(); err != nil {
		return nil, err
	}
	
	return &UserLogin{
		UserId: resp.UserId,
		Name:   resp.Name,
		Token:  token,
	}, nil
}

func (a *Auth) Logout() error {
	userId := a.Request.Header.Get("goshop_user_id")
	spew.Dump(userId)
	return nil
}
