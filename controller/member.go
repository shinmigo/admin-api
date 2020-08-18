package controller

import (
	"goshop/api/filter"

	"github.com/davecgh/go-spew/spew"
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
	result, err := MemberFilter.Index()
	if err != nil {
		m.SetResponse(result, err)
		return
	}

	m.SetResponse(result)
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

// 会员详情
func (m *Member) Info() {
	result, err := MemberFilter.Info()
	if err != nil {
		m.SetResponse(result, err)
		return
	}

	m.SetResponse(result)
}

// 更新会员状态
func (m *Member) EditStatus() {
	err := MemberFilter.EditStatus()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse()
}
