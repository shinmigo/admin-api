package service

import (
	"context"
	"encoding/json"
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

func (m *ProductParam) Add() error {
	contents := m.PostForm("contents")
	typeNumber, _ := strconv.ParseInt(m.PostForm("type"), 10, 32)
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)
	contentsList := make([]string, 0, 32)
	if err := json.Unmarshal([]byte(contents), &contentsList); err != nil {
		return fmt.Errorf("参数值解析失败, err: %v", err)
	}

	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	req := &productpb.Param{
		Name:     m.PostForm("name"),
		Type:     productpb.ParamType(typeNumber),
		Contents: contentsList,
		AdminId:  adminIdNum,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductParam.AddParam(ctx, req)
	cancel()

	if err != nil {
		return fmt.Errorf("添加失败, err: %v", err)
	}

	if resp.State == 0 {
		return fmt.Errorf("添加失败")
	}

	return nil
}

func (m *ProductParam) Edit() error {
	paramId := m.PostForm("param_id")
	typeStr := m.PostForm("type")
	contents := m.PostForm("contents")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	paramIdNumber, _ := strconv.ParseUint(paramId, 10, 64)
	typeNumber, _ := strconv.ParseInt(typeStr, 10, 32)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	contentsList := make([]string, 0, 32)
	if err := json.Unmarshal([]byte(contents), &contentsList); err != nil {
		return fmt.Errorf("参数值解析失败, err: %v", err)
	}

	req := &productpb.Param{
		ParamId:  paramIdNumber,
		Name:     m.PostForm("name"),
		Type:     productpb.ParamType(typeNumber),
		Contents: contentsList,
		AdminId:  adminIdNum,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductParam.EditParam(ctx, req)
	cancel()

	if err != nil {
		return fmt.Errorf("编辑失败, err: %v", err)
	}

	if resp.State == 0 {
		return fmt.Errorf("编辑失败")
	}

	return nil
}

func (m *ProductParam) Del() error {
	paramId := m.PostForm("param_id")
	paramIdNumber, _ := strconv.ParseUint(paramId, 10, 64)

	req := &productpb.DelParamReq{
		ParamId: paramIdNumber,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductParam.DelParam(ctx, req)
	cancel()

	if err != nil {
		return fmt.Errorf("删除失败, err: %v", err)
	}

	if resp.State == 0 {
		return fmt.Errorf("删除失败")
	}

	return nil
}

func (m *ProductParam) Detail(paramId uint64) (*productpb.ParamDetail, error) {
	req := &productpb.ListParamReq{
		Id:       paramId,
		Page:     1,
		PageSize: 1,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductParam.GetParamList(ctx, req)
	cancel()

	if err != nil {
		return nil, fmt.Errorf("获取商品参数失败")
	}

	if len(resp.Params) > 0 {
		return resp.Params[0], nil
	}

	return nil, fmt.Errorf("获取不到商品参数")
}
