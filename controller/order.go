package controller

import (
	"goshop/admin-api/filter"
)

//var OrderFilter *filter.Order

type Order struct {
	Base
	OrderFilter *filter.Order
}

func (o *Order) Initialise() {
	o.OrderFilter = filter.NewOrder(o.Context)
}

func (o *Order) Index() {
	list, err := o.OrderFilter.Index()
	if err != nil {
		o.SetResponse(nil, err)
		return
	}

	o.SetResponse(list)
}

func (o *Order) Status() {
	list, err := o.OrderFilter.Status()
	if err != nil {
		o.SetResponse(nil, err)
		return
	}

	o.SetResponse(list)
}
