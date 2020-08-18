package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/api/pkg/grpc/gclient"
	"time"
)

type ProductSpec struct {
	*gin.Context
}

func NewProductSpec(c *gin.Context) *ProductSpec {
	return &ProductSpec{Context: c}
}

func (m *ProductSpec) Index(param *productpb.ListSpecReq) (*productpb.ListSpecRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	list, err := gclient.ProductSpecClient.GetSpecList(ctx, param)
	cancel()

	return list, err
}

func (m *ProductSpec) Add(spec *productpb.Spec) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductSpecClient.AddSpec(ctx, spec)
	cancel()

	return err
}

func (m *ProductSpec) Edit(spec *productpb.Spec) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductSpecClient.EditSpec(ctx, spec)
	cancel()

	return err
}

func (m *ProductSpec) Delete(idParam uint64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	ProductSpecParam := &productpb.DelSpecReq{
		SpecId: idParam,
	}
	_, err := gclient.ProductSpecClient.DelSpec(ctx, ProductSpecParam)
	cancel()

	return err
}
