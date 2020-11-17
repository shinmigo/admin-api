package controller

import (
	"goshop/admin-api/filter"
)

var BannerAdFilter *filter.BannerAd

type BannerAd struct {
	Base
}

func (m *BannerAd) Initialise() {
	BannerAdFilter = filter.NewBannerAd(m.Context)
}

func (m *BannerAd) Index() {
	list, err := BannerAdFilter.Index()
	if err != nil {
		m.SetResponse(list, err)
		return
	}

	m.SetResponse(list)
}

func (m *BannerAd) Add() {
	err := BannerAdFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *BannerAd) Edit() {
	err := BannerAdFilter.Edit()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *BannerAd) EditStatus() {
	err := BannerAdFilter.EditStatus()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

func (m *BannerAd) Delete() {
	err := BannerAdFilter.Delete()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}
