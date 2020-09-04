package controller

import (
	"goshop/admin-api/filter"
)

var carrierCompany *filter.CarrierCompany

type CarrierCompany struct {
	Base
}

func (m *CarrierCompany) Initialise() {
	carrierCompany = filter.NewCarrierCompany(m.Context)
}

func (m *CarrierCompany) Index() {
	list, err := carrierCompany.Index()
	if err != nil {
		m.SetResponse(list, err)
		return
	}

	m.SetResponse(list)
}

func (m *CarrierCompany) Add() {
	err := carrierCompany.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *CarrierCompany) Edit() {
	err := carrierCompany.Edit()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *CarrierCompany) Delete() {
	err := carrierCompany.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *CarrierCompany) EditStatus() {
	err := carrierCompany.EditStatus()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}
