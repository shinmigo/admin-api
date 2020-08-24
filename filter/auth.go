package filter

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"goshop/api/pkg/validation"
	"goshop/api/service"
)

type Auth struct {
	validation validation.Validation
	*gin.Context
}

func NewAuth(c *gin.Context) *Auth {
	return &Auth{Context: c, validation: validation.Validation{}}
}

func (a *Auth) Login() (interface{}, error) {
	username := a.PostForm("username")
	password := a.PostForm("password")
	
	a.validation.Required(username).Message("用户名不能为空！")
	a.validation.Required(password).Message("密码不能为空！")
	
	if a.validation.HasError() {
		return nil, a.validation.GetError()
	}
	
	res, err := service.NewAuth(a.Context).Login()
	if err != nil {
		return nil, err
	}
	
	return res, nil
}

func (a *Auth) Logout() error {
	spew.Dump(123)
	if err := service.NewAuth(a.Context).Logout(); err != nil {
		return err
	}
	return nil
}
