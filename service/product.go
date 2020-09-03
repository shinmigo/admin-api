package service

import (
	"context"
	"goshop/admin-api/pkg/grpc/gclient"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
)

type Product struct {
	*gin.Context
}

func NewProduct(c *gin.Context) *Product {
	return &Product{Context: c}
}

func (m *Product) Index(listReq *productpb.ListProductReq) (*productpb.ListProductRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	list, err := gclient.ProductClient.GetProductList(ctx, listReq)
	cancel()

	return list, err
}

func (m *Product) Add(product *productpb.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductClient.AddProduct(ctx, product)
	cancel()

	return err
}

func (m *Product) Edit(product *productpb.Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductClient.EditProduct(ctx, product)
	cancel()

	return err
}

func (m *Product) Delete(product *productpb.DelProductReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductClient.DelProduct(ctx, product)
	cancel()

	return err
}
