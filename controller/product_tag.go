package controller

import (
	"goshop/api/filter"
)

var productTagFilter *filter.ProdcutTag

type ProductTag struct {
	Base
}

func (m *ProductTag) Initialise() {
	productTagFilter = filter.NewProductTag(m.Context)
}

func (m *ProductTag) Index() {
	str, err := productTagFilter.Index()
	if err != nil {
		m.SetResponse(str, err)
		return
	}

	m.SetResponse(str)
}

func (m *User) Add() {
	str, err := userFilter.Test()
	if err != nil {
		//m.SetResponse(str, err, -110010) 也可以重写接受过来的错误码，自己定义
		m.SetResponse(str, err)
		return
	}

	m.SetResponse(str)
}
