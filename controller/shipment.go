package controller

import (
	"goshop/admin-api/filter"
)

var ShipmentFilter *filter.Shipment

type Shipment struct {
	Base
}

func (s *Shipment) Initialise() {
	ShipmentFilter = filter.NewShipment(s.Context)
}

func (s *Shipment) Add() {
	resp, err := ShipmentFilter.Add()
	if err != nil {
		s.SetResponse(nil, err)
		return
	}

	s.SetResponse(resp)
}
