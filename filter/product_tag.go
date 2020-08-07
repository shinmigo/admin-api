package filter

import (
	"goshop/api/pkg/validation"
	"goshop/api/service"
	"regexp"
	"strconv"

	"github.com/shinmigo/pb/productpb"

	"github.com/gin-gonic/gin"
)

type ProdcutTag struct {
	*gin.Context
}

func NewProductTag(c *gin.Context) *ProdcutTag {
	return &ProdcutTag{Context: c}
}

func (m *ProdcutTag) Index() (*productpb.ListTagRes, error) {
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
	list, err := service.NewProductTag(m.Context).Index(pNumber, pSize)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (m *ProdcutTag) Add() error {
	valid := validation.Validation{}
	name := m.Query("name")
	valid.Required(name).Message("名称不能为空")
	if valid.HasError() {
		return valid.GetError()
	}

	if err := service.NewProductTag(m.Context).Add(); err != nil {
		return err
	}

	return nil
}
