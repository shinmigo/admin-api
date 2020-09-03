package controller

import "goshop/admin-api/filter"

var ProductFilter *filter.Product

type Product struct {
	Base
}

func (m *Product) Initialise() {
	ProductFilter = filter.NewProduct(m.Context)
}

func (m *Product) Index() {
	list, err := ProductFilter.Index()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse(list)
}

func (m *Product) Add() {
	err := ProductFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *Product) Edit() {
	err := ProductFilter.Edit()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *Product) Delete() {
	err := ProductFilter.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}
