package controller

import (
	"goshop/admin-api/filter"
)

var productKindFilter *filter.ProductKind

type ProductKind struct {
	Base
}

func (m *ProductKind) Initialise() {
	productKindFilter = filter.NewProductKind(m.Context)
}

func (m *ProductKind) Index() {
	str, err := productKindFilter.Index()
	if err != nil {
		m.SetResponse(str, err)
		return
	}

	m.SetResponse(str)
}

func (m *ProductKind) Add() {
	err := productKindFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductKind) Delete() {
	err := productKindFilter.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductKind) Edit() {
	err := productKindFilter.Edit()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductKind) BindParam() {
	err := productKindFilter.BindParam()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductKind) BindSpec() {
	err := productKindFilter.BindSpec()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}
