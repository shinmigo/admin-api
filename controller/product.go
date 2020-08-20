package controller

import "goshop/api/filter"

var ProductFilter *filter.Product

type Product struct {
	Base
}

func (m *Product) Initialise() {
	ProductFilter = filter.NewProduct(m.Context)
}

func (m *Product) Add() {
	err := ProductFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}
