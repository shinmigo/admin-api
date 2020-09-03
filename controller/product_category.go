package controller

import "goshop/admin-api/filter"

var productCategoryFilter *filter.ProductCategory

type ProductCategory struct {
	Base
}

func (m *ProductCategory) Initialise() {
	productCategoryFilter = filter.NewProductCategory(m.Context)
}

func (m *ProductCategory) Index() {
	list, err := productCategoryFilter.Index()
	if err != nil {
		m.SetResponse(list, err)
		return
	}

	m.SetResponse(list)
}

func (m *ProductCategory) Add() {
	err := productCategoryFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductCategory) Edit() {
	err := productCategoryFilter.Edit()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductCategory) EditStatus() {
	err := productCategoryFilter.EditStatus()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductCategory) Delete() {
	err := productCategoryFilter.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}
