package service

import (
	"context"
	"fmt"
	"goshop/api/pkg/grpc/gclient"
	"strconv"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
)

type ProductParam struct {
	*gin.Context
}

func NewProductParam(c *gin.Context) *ProductParam {
	return &ProductParam{Context: c}
}

func (m *ProductParam) Index(pNumber, pSize uint64) (*productpb.ListParamRes, error) {
	req := &productpb.ListParamReq{
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
	resp, err := gclient.ProductParam.GetParamList(ctx, req)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取商品参数列表失败")
	}
	
	return resp, nil
}
