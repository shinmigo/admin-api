package controller

import (
	"goshop/admin-api/filter"
)

var userFilter *filter.User

type User struct {
	Base
}

func (m *User) Initialise() {
	userFilter = filter.NewUser(m.Context)
}

func (m *User) GetListQuery() {
	str, err := userFilter.GetListQuery()
	if err != nil {
		//m.SetResponse(str, err, -110010) 也可以重写接受过来的错误码，自己定义
		m.SetResponse(str, err)
		return
	}

	m.SetResponse(str)
}

func (m *User) Test() {
	str, err := userFilter.Test()
	if err != nil {
		//m.SetResponse(str, err, -110010) 也可以重写接受过来的错误码，自己定义
		m.SetResponse(str, err)
		return
	}

	m.SetResponse(str)
}
