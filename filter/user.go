package filter

import (
	"goshop/admin-api/pkg/utils"
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
	"regexp"

	"github.com/gin-gonic/gin"
)

type User struct {
	*gin.Context
}

func NewUser(c *gin.Context) *User {
	return &User{Context: c}
}

//test1
func (m *User) GetListQuery() (string, error) {
	valid := validation.Validation{}
	valid.Required(m.Query("user")).Message("用户名不能为空")
	valid.Required(m.Query("password")).Message("密码不能为空")
	valid.Match(m.Query("password"), regexp.MustCompile(`^[0-9]{2,20}$`)).Message("密码格式 不正确")
	if valid.HasError() {
		return "", valid.GetError()
	}

	list := service.NewUser(m.Context).GetListQuery("hello")

	return list, nil
}

//test2
func (m *User) Test() (string, *utils.ErrorCode) {
	valid := validation.Validation{}
	valid.Required(m.Query("user")).Message("用户名不能为空")
	valid.Required(m.Query("password")).Message("密码不能为空")
	valid.Match(m.Query("password"), regexp.MustCompile(`^[0-9]{2,20}$`)).Message("密码格式 不正确")
	if valid.HasError() {
		//注意，这里返回的错误，是带错误码的
		return "", utils.NewErrorCode(-110021, valid.GetErroString())
	}

	list := service.NewUser(m.Context).GetListQuery(m.Query("user"))

	return list, nil
}
