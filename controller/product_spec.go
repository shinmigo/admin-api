package controller

import "goshop/admin-api/filter"

var productSpecFilter *filter.ProductSpec

type ProductSpec struct {
	Base
}

func (m *ProductSpec) Initialise() {
	productSpecFilter = filter.NewProductSpec(m.Context)
}

func (m *ProductSpec) Index() {
	list, err := productSpecFilter.Index()
	if err != nil {
		m.SetResponse(list, err)
		return
	}

	m.SetResponse(list)
}

func (m *ProductSpec) Add() {
	err := productSpecFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductSpec) Edit() {
	err := productSpecFilter.Edit()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductSpec) Delete() {
	err := productSpecFilter.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

// 获取可绑定规格列表， 无分页返回全部
func (m *ProductSpec) BindableSpecs() {
	str, err := productSpecFilter.BindableSpecs()
	if err != nil {
		m.SetResponse(str, err)
		return
	}
	
	m.SetResponse(str)
}
