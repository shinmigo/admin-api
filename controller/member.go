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

func (m *Member) Index() {
	str, err := MemberFilter.Index()
	if err != nil {
		m.SetResponse(str, err)
		return
	}
	
	m.SetResponse(str)
}
