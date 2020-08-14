package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"golang.org/x/net/context"
	"goshop/api/pkg/grpc/gclient"
	"strconv"
	"time"
)

type ProductKind struct {
	*gin.Context
}

func NewProductKind(c *gin.Context)  *ProductKind {
	return &ProductKind{Context: c}
}

func (m *ProductKind) Index(pNumber, pSize uint64) (*productpb.ListKindRes, error) {
	req := &productpb.ListKindReq{
		Page:     pNumber,
		PageSize: pSize,
	}

	if len(m.Query("name")) > 0 {
		req.Name = m.Query("name")
	}

	if len(m.Query("id")) > 0 {
		id, _ := strconv.ParseUint(m.Query("id"), 10, 64)
		req.Id = id
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductKind.GetKindList(ctx, req)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取商品类型失败")
	}

	return resp, nil
}

func (m *ProductKind) Add() error {
	storeId := m.DefaultPostForm("store_id", "0")
	sId, _ := strconv.ParseUint(storeId, 10, 64)

	req := &productpb.Kind{
		StoreId:              sId,
		Name:                 m.PostForm("name"),
		AdminId:              0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductKind.AddKind(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("添加商品类型失败")
	}

	if resp.State == 0 {
		return fmt.Errorf("添加失败")
	}

	return nil
}