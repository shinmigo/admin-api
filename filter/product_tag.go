package filter

import (
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
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
	name := m.PostForm("name")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	valid := validation.Validation{}
	valid.Required(name).Message("名称不能为空")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品标签名称格式错误")
	if valid.HasError() {
		return valid.GetError()
	}

	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	req := &productpb.Tag{
		Name:    name,
		AdminId: adminIdNum,
	}
	if err := service.NewProductTag(m.Context).Add(req); err != nil {
		return err
	}

	return nil
}

func (m *ProdcutTag) Edit() error {
	id := m.PostForm("id")
	name := m.PostForm("name")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	valid := validation.Validation{}
	valid.Required(id).Message("标签不能为空")
	valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("要编辑的标签格式错误")
	valid.Required(name).Message("名称不能为空")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品标签名称格式错误")
	if valid.HasError() {
		return valid.GetError()
	}

	idNum, _ := strconv.ParseUint(id, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	req := &productpb.Tag{
		TagId:   idNum,
		Name:    name,
		AdminId: adminIdNum,
	}
	return service.NewProductTag(m.Context).Edit(req)
}

func (m *ProdcutTag) Delete() error {
	id := m.PostForm("id")

	valid := validation.Validation{}
	valid.Required(id).Message("标签不能为空")
	valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("要删除的标签格式错误")
	if valid.HasError() {
		return valid.GetError()
	}

	idNum, _ := strconv.ParseUint(id, 10, 64)
	return service.NewProductTag(m.Context).Delete(idNum)
}
