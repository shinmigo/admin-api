package controller

import (
	"github.com/davecgh/go-spew/spew"
	"goshop/api/filter"
)

var MemberFilter *filter.Member

type Member struct {
	Base
}

func (m *Member) Initialise() {
	spew.Dump(111111)
	MemberFilter = filter.NewMember(m.Context)
}

// 会员列表
func (m *Member) Index() {
	str, err := MemberFilter.Index()
	if err != nil {
		m.SetResponse(str, err)
		return
	}
	
	m.SetResponse(str)
}

// 添加会员
func (m *Member) Add() {
	err := MemberFilter.Add()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}

// 会员编辑
func (m *Member) Edit() {
	err := MemberFilter.Edit()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}
