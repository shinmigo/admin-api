package controller

import (
	"goshop/api/filter"
)

var productParamFilter *filter.ProdcutParam

type ProductParam struct {
	Base
}

func (m *ProductParam) Initialise() {
	productParamFilter = filter.NewProductParam(m.Context)
}

func (m *ProductParam) Index() {
	str, err := productParamFilter.Index()
	if err != nil {
		m.SetResponse(str, err)
		return
	}
	
	m.SetResponse(str)
}
