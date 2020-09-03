package filter

import (
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
	"regexp"
	"strconv"

	"github.com/shinmigo/pb/productpb"

	"github.com/gin-gonic/gin"
)

type ProdcutParam struct {
	validation validation.Validation
	*gin.Context
}

func NewProductParam(c *gin.Context) *ProdcutParam {
	return &ProdcutParam{Context: c}
}

func (m *ProdcutParam) Index() (*productpb.ListParamRes, error) {
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	m.validation.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	m.validation.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}

	pNumber, _ := strconv.ParseUint(page, 10, 16)
	pSize, _ := strconv.ParseUint(pageSize, 10, 16)
	list, err := service.NewProductParam(m.Context).Index(pNumber, pSize)
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (m *ProdcutParam) Add() error {
	name := m.PostForm("name")
	typeStr := m.PostForm("type")
	contents := m.PostForm("contents")

	m.validation.Required(name).Message("参数名称不能为空！")
	m.validation.Match(typeStr, regexp.MustCompile(`^[0-3]{1}$`)).Message("参数类型不正确")
	m.validation.Required(contents).Message("参数值不能为空！")

	if m.validation.HasError() {
		return m.validation.GetError()
	}

	if err := service.NewProductParam(m.Context).Add(); err != nil {
		return err
	}

	return nil
}

func (m *ProdcutParam) Edit() error {
	paramId := m.PostForm("param_id")
	name := m.PostForm("name")
	typeStr := m.PostForm("type")
	contents := m.PostForm("contents")

	m.validation.Required(paramId).Message("paramId不能为空！")
	m.validation.Required(name).Message("参数名称不能为空！")
	m.validation.Match(typeStr, regexp.MustCompile(`^[0-3]{1}$`)).Message("参数类型不正确")
	m.validation.Required(contents).Message("参数值不能为空！")

	if m.validation.HasError() {
		return m.validation.GetError()
	}

	if err := service.NewProductParam(m.Context).Edit(); err != nil {
		return err
	}
	return nil
}

func (m *ProdcutParam) Del() error {
	paramId := m.PostForm("param_id")

	m.validation.Required(paramId).Message("paramId不能为空！")
	if m.validation.HasError() {
		return m.validation.GetError()
	}

	if err := service.NewProductParam(m.Context).Del(); err != nil {
		return err
	}
	return nil
}

func (m *ProdcutParam) Detail() (*productpb.ParamDetail, error) {
	paramId := m.Query("param_id")
	m.validation.Required(paramId).Message("paramId不能为空！")

	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}

	paramIdNumber, _ := strconv.ParseUint(paramId, 10, 64)
	res, err := service.NewProductParam(m.Context).Detail(paramIdNumber)

	if err != nil {
		return nil, err
	}

	return res, nil
}
