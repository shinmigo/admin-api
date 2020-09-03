package service

import (
	"context"
	"fmt"
	"goshop/admin-api/pkg/grpc/gclient"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
)

type ProductTag struct {
	*gin.Context
}

func NewProductTag(c *gin.Context) *ProductTag {
	return &ProductTag{Context: c}
}

func (m *ProductTag) Index(pNumber, pSize uint64) (*productpb.ListTagRes, error) {
	req := &productpb.ListTagReq{
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
	resp, err := gclient.ProductTag.GetTagList(ctx, req)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取标签列表失败")
	}

	return resp, nil
}

func (m *ProductTag) Add(req *productpb.Tag) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductTag.AddTag(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("添加标签失败")
	}

	if resp.State == 0 {
		return fmt.Errorf("添加失败")
	}

	return nil
}

func (m *ProductTag) Edit(req *productpb.Tag) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductTag.EditTag(ctx, req)
	cancel()

	return err
}

func (m *ProductTag) Delete(id uint64) error {
	req := &productpb.DelTagReq{
		TagId: id,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.ProductTag.DelTag(ctx, req)
	cancel()

	return err
}
