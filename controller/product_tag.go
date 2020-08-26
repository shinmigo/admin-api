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

func (m *ProductTag) Add() {
	err := productTagFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductTag) Edit() {
	err := productTagFilter.Edit()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *ProductTag) Delete() {
	err := productTagFilter.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}
