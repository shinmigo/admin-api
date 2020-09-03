package service

import (
	"context"
	"goshop/admin-api/pkg/grpc/gclient"
	"time"

	"github.com/shinmigo/pb/basepb"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/orderpb"
)

type Shipment struct {
	*gin.Context
}

func NewShipment(ctx *gin.Context) *Shipment {
	return &Shipment{Context: ctx}
}

func (s *Shipment) Add(req *orderpb.Shipment) (*basepb.AnyRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ShipmentClient.AddShipment(ctx, req)
	cancel()

	return resp, err
}
