package service

import (
	"context"
	"fmt"
	"goshop/admin-api/pkg/grpc/gclient"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
)

type ProductCategory struct {
	*gin.Context
}

func NewProductCategory(c *gin.Context) *ProductCategory {
	return &ProductCategory{Context: c}
}

func (m *ProductCategory) Index(param *productpb.ListCategoryReq) (*productpb.ListCategoryRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	list, err := gclient.ProductCategoryClient.GetCategoryList(ctx, param)
	cancel()

	return list, err
}

func (m *ProductCategory) Add(category *productpb.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductCategoryClient.AddCategory(ctx, category)
	cancel()

	return err
}

func (m *ProductCategory) Edit(category *productpb.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductCategoryClient.EditCategory(ctx, category)
	cancel()

	return err
}

func (m *ProductCategory) EditStatus(statusParam *productpb.EditCategoryStatusReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductCategoryClient.EditCategoryStatus(ctx, statusParam)
	cancel()

	return err
}

func (m *ProductCategory) Delete(idParam []uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	productCategoryParam := &productpb.DelCategoryReq{
		CategoryId: idParam,
	}
	_, err := gclient.ProductCategoryClient.DelCategory(ctx, productCategoryParam)
	cancel()
	fmt.Println(err.Error())

	return err
}
