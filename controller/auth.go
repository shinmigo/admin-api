package controller

import (
	"goshop/admin-api/filter"
)

var AuthFilter *filter.Auth

type Auth struct {
	Base
}

func (a *Auth) Initialise() {
	AuthFilter = filter.NewAuth(a.Context)
}

func (a *Auth) Login() {
	res, err := AuthFilter.Login()
	if err != nil {
		a.SetResponse(nil, err)
		return
	}

	a.SetResponse(res)
}

func (a *Auth) Logout() {
	if err := AuthFilter.Logout(); err != nil {
		a.SetResponse(nil, err)
		return
	}

	a.SetResponse()
}
