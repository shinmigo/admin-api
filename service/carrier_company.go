package service

import (
	"context"
	"goshop/admin-api/pkg/grpc/gclient"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/shoppb"
)

type CarrierCompany struct {
	*gin.Context
}

func NewCarrierCompany(c *gin.Context) *CarrierCompany {
	return &CarrierCompany{Context: c}
}

func (m *CarrierCompany) Index(param *shoppb.ListCarrierReq) (*shoppb.ListCarrierRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	list, err := gclient.ShopCarrier.GetCarrierList(ctx, param)
	cancel()

	return list, err
}

func (m *CarrierCompany) Add(carrier *shoppb.Carrier) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ShopCarrier.AddCarrier(ctx, carrier)
	cancel()

	return err
}

func (m *CarrierCompany) Edit(carrier *shoppb.Carrier) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ShopCarrier.EditCarrier(ctx, carrier)
	cancel()

	return err
}

func (m *CarrierCompany) Delete(carrier *shoppb.DelCarrierReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ShopCarrier.DelCarrier(ctx, carrier)
	cancel()

	return err
}

func (m *CarrierCompany) EditStatus(statusParam *shoppb.EditCarrierStatusReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ShopCarrier.EditCarrierStatus(ctx, statusParam)
	cancel()

	return err
}
