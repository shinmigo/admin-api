package service

import (
	"fmt"
	"goshop/admin-api/pkg/grpc/gclient"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"golang.org/x/net/context"
)

type ProductKind struct {
	*gin.Context
}

func NewProductKind(c *gin.Context) *ProductKind {
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
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)
	sId, _ := strconv.ParseUint(storeId, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)

	req := &productpb.Kind{
		StoreId: sId,
		Name:    m.PostForm("name"),
		AdminId: adminIdNum,
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

func (m *ProductKind) Delete() error {
	kindId := m.PostForm("kind_id")
	kingIdNumber, _ := strconv.ParseUint(kindId, 10, 64)

	req := &productpb.DelKindReq{
		KindId: kingIdNumber,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductKind.DelKind(ctx, req)
	cancel()

	if err != nil {
		return fmt.Errorf("删除失败, err: %v", err)
	}

	if resp.State == 0 {
		return fmt.Errorf("删除失败")
	}

	return nil
}

func (m *ProductKind) Edit() error {
	name := m.PostForm("name")
	kindId := m.PostForm("kind_id")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	kindIdNumber, _ := strconv.ParseUint(kindId, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)

	req := &productpb.Kind{
		KindId:  kindIdNumber,
		Name:    name,
		AdminId: adminIdNum,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductKind.EditKind(ctx, req)
	cancel()

	if err != nil {
		return fmt.Errorf("编辑失败, err: %v", err)
	}

	if resp.State == 0 {
		return fmt.Errorf("编辑失败")
	}

	return nil
}

func (m *ProductKind) BindParam() error {
	kindId := m.PostForm("kind_id")
	paramIds := m.PostForm("param_ids")

	kindIdNumber, _ := strconv.ParseUint(kindId, 10, 64)

	paramIdArr := strings.Split(paramIds, ",")

	var paramIdNums = []uint64{}
	for _, i := range paramIdArr {
		j, _ := strconv.ParseUint(i, 10, 64)
		paramIdNums = append(paramIdNums, j)
	}

	req := &productpb.BindParamReq{
		KindId:   kindIdNumber,
		ParamIds: paramIdNums,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductKind.BindParam(ctx, req)
	cancel()

	if err != nil {
		return fmt.Errorf("绑定参数失败, err: %v", err)
	}

	if resp.State == 0 {
		return fmt.Errorf("绑定参数失败")
	}

	return nil
}

func (m *ProductKind) BindSpec() error {
	kindId := m.PostForm("kind_id")
	paramIds := m.PostForm("spec_ids")

	kindIdNumber, _ := strconv.ParseUint(kindId, 10, 64)

	specIdArr := strings.Split(paramIds, ",")

	var specIdNums = []uint64{}
	for _, i := range specIdArr {
		j, _ := strconv.ParseUint(i, 10, 64)
		specIdNums = append(specIdNums, j)
	}

	req := &productpb.BindSpecReq{
		KindId:  kindIdNumber,
		SpecIds: specIdNums,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.ProductKind.BindSpec(ctx, req)
	cancel()

	if err != nil {
		return fmt.Errorf("绑定规格失败, err: %v", err)
	}

	if resp.State == 0 {
		return fmt.Errorf("绑定规格失败")
	}

	return nil
}
