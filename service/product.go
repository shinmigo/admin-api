package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/api/pkg/grpc/gclient"
	"time"
)

type Product struct {
	*gin.Context
}

func NewProduct(c *gin.Context) *Product {
	return &Product{Context: c}
}

func (m *Product) Add(product *productpb.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductClient.AddProduct(ctx, product)
	cancel()

	return err
}
