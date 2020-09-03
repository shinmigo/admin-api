package controller

import (
	"goshop/admin-api/filter"
)

var OrderFilter *filter.Order

type Order struct {
	Base
}

func (o *Order) Initialise() {
	OrderFilter = filter.NewOrder(o.Context)
}

func (o *Order) Index() {
	list, err := OrderFilter.Index()
	if err != nil {
		o.SetResponse(nil, err)
		return
	}

	o.SetResponse(list)
}
