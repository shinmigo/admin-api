package filter

import (
	"encoding/json"
	"errors"
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
)

type ProductCategory struct {
	*gin.Context
}

func NewProductCategory(c *gin.Context) *ProductCategory {
	return &ProductCategory{Context: c}
}

func (m *ProductCategory) Index() (*productpb.ListCategoryRes, error) {
	id := m.Query("id")
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	name := m.DefaultQuery("category_name", "")
	status := m.Query("status")

	var idNum uint64
	idLen := len(id)
	statusLen := len(status)
	valid := validation.Validation{}
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if idLen > 0 {
		valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品分类数据不正确")
	}
	if len(name) > 0 {
		valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品分类名称格式错误")
	}
	if statusLen > 0 {
		valid.Match(status, regexp.MustCompile(`^0|1$`)).Message("商品分类状态格式错误！")
	}
	if valid.HasError() {
		return nil, valid.GetError()
	}

	if idLen > 0 {
		idNum, _ = strconv.ParseUint(id, 10, 64)
	}
	pageNum, _ := strconv.ParseInt(page, 10, 64)
	pageSizeNum, _ := strconv.ParseInt(pageSize, 10, 64)
	listCategoryReq := &productpb.ListCategoryReq{
		Page:     pageNum,
		PageSize: pageSizeNum,
		Name:     name,
		Id:       idNum,
	}
	if statusLen > 0 {
		var statusNum productpb.CategoryStatus
		if status == "0" {
			statusNum = productpb.CategoryStatus_InActive
		} else {
			statusNum = productpb.CategoryStatus_Active
		}
		listCategoryReq.StatusPresent = &productpb.ListCategoryReq_Status{Status: statusNum}
	}

	return service.NewProductCategory(m.Context).Index(listCategoryReq)
}

func (m *ProductCategory) Add() error {
	parentId := m.PostForm("parent_id")
	name := m.PostForm("name")
	sort := m.PostForm("sort")
	status := m.PostForm("status")
	icon := m.PostForm("icon")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	var parentIdNum uint64
	parentIdLen := len(parentId)
	valid := validation.Validation{}
	if parentIdLen > 0 {
		valid.Match(parentId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("上级分类数据格式错误！")
	}
	valid.Required(name).Message("请填写商品分类名称！")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品分类名称格式错误")
	valid.Match(sort, regexp.MustCompile(`^[0-9]*$`)).Message("商品分类排序格式错误！")
	valid.Match(status, regexp.MustCompile(`^0|1$`)).Message("商品分类状态格式错误！")
	valid.Required(icon).Message("请上传商品分类图标！")
	valid.Match(icon, regexp.MustCompile(`^(http://|https://)[a-zA-z0-9.]+$`)).Message("商品分类图标数据错误")
	if valid.HasError() {
		return valid.GetError()
	}

	var statusNum productpb.CategoryStatus
	if parentIdLen > 0 {
		parentIdNum, _ = strconv.ParseUint(parentId, 10, 64)
	}
	sortNum, _ := strconv.ParseUint(sort, 10, 64)
	if status == "1" {
		statusNum = productpb.CategoryStatus_Active
	} else {
		statusNum = productpb.CategoryStatus_InActive
	}
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	reqProductCategoryParam := &productpb.Category{
		ParentId: parentIdNum,
		Name:     name,
		Icon:     icon,
		Status:   statusNum,
		Sort:     sortNum,
		AdminId:  adminIdNum,
	}
	return service.NewProductCategory(m.Context).Add(reqProductCategoryParam)
}

func (m *ProductCategory) Edit() error {
	categoryId := m.PostForm("id")
	parentId := m.PostForm("parent_id")
	name := m.PostForm("name")
	sort := m.PostForm("sort")
	status := m.PostForm("status")
	icon := m.PostForm("icon")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	var parentIdNum uint64
	parentIdLen := len(parentId)
	valid := validation.Validation{}
	valid.Required(categoryId).Message("请提交要编辑的商品分类！")
	valid.Match(categoryId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品分类数据格式错误！")
	if parentIdLen > 0 {
		valid.Match(parentId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("上级分类数据格式错误！")
	}
	valid.Required(name).Message("请填写商品分类名称！")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品分类名称格式错误")
	valid.Match(sort, regexp.MustCompile(`^[0-9]*$`)).Message("商品分类排序格式错误！")
	valid.Match(status, regexp.MustCompile(`^0|1$`)).Message("商品分类状态格式错误！")
	valid.Required(icon).Message("请上传商品分类图标！")
	valid.Match(icon, regexp.MustCompile(`^(http://|https://)[a-zA-z0-9.]+$`)).Message("商品分类图标数据错误")
	if valid.HasError() {
		return valid.GetError()
	}

	var statusNum productpb.CategoryStatus
	if parentIdLen > 0 {
		parentIdNum, _ = strconv.ParseUint(parentId, 10, 64)
	}
	sortNum, _ := strconv.ParseUint(sort, 10, 64)
	categoryIdNum, _ := strconv.ParseUint(categoryId, 10, 64)
	if status == "1" {
		statusNum = productpb.CategoryStatus_Active
	} else {
		statusNum = productpb.CategoryStatus_InActive
	}
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	reqProductCategoryParam := &productpb.Category{
		CategoryId: categoryIdNum,
		ParentId:   parentIdNum,
		Name:       name,
		Icon:       icon,
		Status:     statusNum,
		Sort:       sortNum,
		AdminId:    adminIdNum,
	}
	return service.NewProductCategory(m.Context).Edit(reqProductCategoryParam)
}

func (m *ProductCategory) EditStatus() error {
	categoryId := m.PostForm("id")
	status := m.PostForm("status")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	valid := validation.Validation{}
	valid.Required(categoryId).Message("请提交要编辑的商品分类！")
	valid.Match(status, regexp.MustCompile(`^0|1$`)).Message("商品分类状态格式错误！")
	if valid.HasError() {
		return valid.GetError()
	}

	idParam := make([]uint64, 0, 32)
	err := json.Unmarshal([]byte(categoryId), &idParam)
	if err != nil {
		return errors.New("要删除的商品分类数据格式错误！")
	}
	var statusNum productpb.CategoryStatus
	if status == "1" {
		statusNum = productpb.CategoryStatus_Active
	} else {
		statusNum = productpb.CategoryStatus_InActive
	}
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	param := &productpb.EditCategoryStatusReq{
		CategoryId: idParam,
		Status:     statusNum,
		AdminId:    adminIdNum,
	}
	return service.NewProductCategory(m.Context).EditStatus(param)
}

func (m *ProductCategory) Delete() error {
	id := m.PostForm("id")

	valid := validation.Validation{}
	valid.Required(id).Message("请选择要删除的分类！")
	if valid.HasError() {
		return valid.GetError()
	}

	idParam := make([]uint64, 0, 32)
	err := json.Unmarshal([]byte(id), &idParam)
	if err != nil {
		return errors.New("要删除的商品分类数据格式错误！")
	}

	return service.NewProductCategory(m.Context).Delete(idParam)
}
