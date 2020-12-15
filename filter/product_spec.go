package filter

import (
	"encoding/json"
	"errors"
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
)

type ProductSpec struct {
	*gin.Context
}

func NewProductSpec(c *gin.Context) *ProductSpec {
	return &ProductSpec{Context: c}
}

func (m *ProductSpec) Index() (*productpb.ListSpecRes, error) {
	id := m.Query("id")
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	name := m.DefaultQuery("spec_name", "")

	var idNum uint64
	idLen := len(id)
	valid := validation.Validation{}
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if idLen > 0 {
		valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品规格数据不正确")
	}
	if len(name) > 0 {
		valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品规格名称格式错误")
	}
	if valid.HasError() {
		return nil, valid.GetError()
	}

	if idLen > 0 {
		idNum, _ = strconv.ParseUint(id, 10, 64)
	}
	pageNum, _ := strconv.ParseUint(page, 10, 64)
	pageSizeNum, _ := strconv.ParseUint(pageSize, 10, 64)
	listSpecReq := &productpb.ListSpecReq{
		Page:     pageNum,
		PageSize: pageSizeNum,
		Id:       idNum,
		Name:     name,
	}

	return service.NewProductSpec(m.Context).Index(listSpecReq)
}

func (m *ProductSpec) Add() error {
	name := m.PostForm("name")
	sort := m.PostForm("sort")
	values := m.PostForm("values")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	valid := validation.Validation{}
	valid.Required(name).Message("请填写商品规格名称！")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品规格名称格式错误")
	valid.Match(sort, regexp.MustCompile(`^[0-9]*$`)).Message("商品规格排序格式错误！")
	valid.Required(values).Message("请提交商品规格值！")
	//valid.Match(values, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9,]+$`)).Message("商品规格值数据格式错误！")
	if valid.HasError() {
		return valid.GetError()
	}

	sortNum, _ := strconv.ParseUint(sort, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	valuesList := strings.Split(values, ",")
	reqParam := &productpb.Spec{
		Name:     name,
		Sort:     sortNum,
		Contents: valuesList,
		AdminId:  adminIdNum,
	}

	return service.NewProductSpec(m.Context).Add(reqParam)
}

func (m *ProductSpec) Edit() error {
	specId := m.PostForm("id")
	name := m.PostForm("name")
	sort := m.PostForm("sort")
	values := m.PostForm("values")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	valid := validation.Validation{}
	valid.Required(specId).Message("请提交要编辑的商品规格！")
	valid.Match(specId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品规格数据格式错误！")
	valid.Required(name).Message("请填写商品规格名称！")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品规格名称格式错误")
	valid.Match(sort, regexp.MustCompile(`^[0-9]*$`)).Message("商品规格排序格式错误！")
	valid.Required(values).Message("请提交商品规格值！")
	if valid.HasError() {
		return valid.GetError()
	}

	sortNum, _ := strconv.ParseUint(sort, 10, 64)
	
	valuesList := make([]*productpb.SpecValue, 0, 32)
	if err := json.Unmarshal([]byte(values), &valuesList); err != nil {
		return fmt.Errorf("参数值解析失败, err: %v", err)
	}
	specIdNum, _ := strconv.ParseUint(specId, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	reqParam := &productpb.EditSpecReq{
		SpecId:   specIdNum,
		Name:     name,
		Sort:     sortNum,
		Contents: valuesList,
		AdminId:  adminIdNum,
	}

	return service.NewProductSpec(m.Context).Edit(reqParam)
}

func (m *ProductSpec) Delete() error {
	id := m.PostForm("id")

	valid := validation.Validation{}
	valid.Required(id).Message("请选择要删除的商品规格！")
	if valid.HasError() {
		return valid.GetError()
	}

	idParam := make([]uint64, 0, 32)
	err := json.Unmarshal([]byte(id), &idParam)
	if err != nil {
		return errors.New("要删除的商品规格数据格式错误！")
	}
	return service.NewProductSpec(m.Context).Delete(idParam)
}

func (m *ProductSpec) BindableSpecs() (*productpb.BindSpecAllRes, error) {
	name := m.Query("name")
	list, err := service.NewProductSpec(m.Context).BindableSpecs(name)
	if err != nil {
		return nil, err
	}

	return list, nil
}
