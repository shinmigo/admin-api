package service

import (
	"context"
	"goshop/admin-api/pkg/grpc/gclient"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/orderpb"
)

type Order struct {
	*gin.Context
}

func NewOrder(ctx *gin.Context) *Order {
	return &Order{Context: ctx}
}

func (o *Order) Index(req *orderpb.ListOrderReq) (*orderpb.ListOrderRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	list, err := gclient.OrderClient.GetOrderList(ctx, req)
	cancel()

	return list, err
}
