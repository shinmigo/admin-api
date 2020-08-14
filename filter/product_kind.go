package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/api/pkg/validation"
	"goshop/api/service"
	"regexp"
	"strconv"
)

type ProductKind struct {
	*gin.Context
}

func NewProductKind(c *gin.Context) *ProductKind  {
	return &ProductKind{Context: c}
}

func (m *ProductKind) Index() (*productpb.ListKindRes, error)  {
	valid := validation.Validation{}
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if valid.HasError() {
		return nil, valid.GetError()
	}

	pNumber, _ := strconv.ParseUint(page, 10, 16)
	pSize, _ := strconv.ParseUint(pageSize, 10, 16)

	list, err := service.NewProductKind(m.Context).Index(pNumber, pSize)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (m *ProductKind) Add() error  {
	valid := validation.Validation{}
	name := m.PostForm("name")

	valid.Required(name).Message("名称不能为空")
	if valid.HasError() {
		return valid.GetError()
	}

	if err := service.NewProductKind(m.Context).Add(); err != nil {
		return err
	}

	return nil
}
