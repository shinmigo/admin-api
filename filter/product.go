package filter

import (
	"encoding/json"
	"errors"
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
)

type Product struct {
	*gin.Context
}

func NewProduct(c *gin.Context) *Product {
	return &Product{Context: c}
}

func (m *Product) Index() (*productpb.ListProductRes, error) {
	id := m.Query("id")
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	name := m.DefaultQuery("name", "")
	status := m.Query("status")
	categoryId := m.Query("category_id")

	var idNum uint64
	var categoryIdNum uint64
	idLen := len(id)
	statusLen := len(status)
	categoryIdLen := len(categoryId)
	valid := validation.Validation{}
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if idLen > 0 {
		valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品数据不正确")
	}
	if len(name) > 0 {
		valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品名称格式错误")
	}
	if statusLen > 0 {
		valid.Match(status, regexp.MustCompile(`^1|2$`)).Message("商品状态格式错误！")
	}
	if categoryIdLen > 0 {
		valid.Match(categoryId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品分类格式不正确")
	}
	if valid.HasError() {
		return nil, valid.GetError()
	}
	if idLen > 0 {
		idNum, _ = strconv.ParseUint(id, 10, 64)
	}
	if categoryIdLen > 0 {
		categoryIdNum, _ = strconv.ParseUint(categoryId, 10, 64)
	}
	pageNum, _ := strconv.ParseUint(page, 10, 64)
	pageSizeNum, _ := strconv.ParseUint(pageSize, 10, 64)
	listProductReq := &productpb.ListProductReq{
		Page:       pageNum,
		PageSize:   pageSizeNum,
		Id:         idNum,
		Name:       name,
		CategoryId: categoryIdNum,
	}
	if statusLen > 0 {
		var statusNum productpb.ProductStatus
		if status == "1" {
			statusNum = productpb.ProductStatus_Disabled
		} else {
			statusNum = productpb.ProductStatus_Enabled
		}
		listProductReq.Status = statusNum
	}

	return service.NewProduct(m.Context).Index(listProductReq)
}

func (m *Product) Add() error {
	categoryId := m.PostForm("category_id")
	kindId := m.PostForm("kind_id")
	name := m.PostForm("name")
	shortDescription := m.PostForm("short_description")
	unit := m.PostForm("unit")
	images := m.PostForm("images")
	specType := m.PostForm("spec_type")
	spec := m.PostForm("spec")
	status := m.PostForm("status")
	tags := m.PostForm("tags")
	param := m.PostForm("param")
	description := m.PostForm("description")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	tagsList := make([]uint64, 0, 8)
	valid := validation.Validation{}
	valid.Required(categoryId).Message("请选择商品分类")
	valid.Match(categoryId, regexp.MustCompile(`^[0-9]+$`)).Message("商品分类格式错误")
	valid.Required(kindId).Message("请选择商品类型")
	valid.Match(kindId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品类型格式错误")
	valid.Required(name).Message("请填写商品名称")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品名称格式错误")
	valid.Required(shortDescription).Message("请填写商品简介")
	valid.Match(shortDescription, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品简介格式错误")
	valid.Required(unit).Message("请填写商品单位")
	valid.Match(unit, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品单位格式错误")
	valid.Required(images).Message("请上传商品轮播图")
	valid.Match(images, regexp.MustCompile(`^[a-zA-z0-9,]+$`)).Message("商品轮播图格式错误")
	valid.Required(specType).Message("请选择商品规格")
	valid.Match(specType, regexp.MustCompile(`^1|2$`)).Message("商品规格格式错误")
	valid.Required(spec).Message("请填写商品规格数据")
	valid.Required(status).Message("请选择商品上下架状态")
	valid.Match(status, regexp.MustCompile(`^1|2$`)).Message("商品上下架状态格式错误")
	valid.Required(param).Message("请填写商品详情参数")
	valid.Required(description).Message("请填写商品描述")
	//valid.Match(description, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品描述格式错误")
	if valid.HasError() {
		return valid.GetError()
	}
	if len(tags) > 0 {
		err := json.Unmarshal([]byte(tags), &tagsList)
		if err != nil {
			return errors.New("商品标签参数错误！")
		}
	}
	specParam := make([]*productpb.ProductSpec, 0, 8)
	err := json.Unmarshal([]byte(spec), &specParam)
	if err != nil {
		return errors.New("商品规格参数错误！")
	}
	paramList := make([]*productpb.ProductParam, 0, 8)
	err = json.Unmarshal([]byte(param), &paramList)
	if err != nil {
		return errors.New("商品参数格式错误！")
	}

	var specTypeReq productpb.ProductSpecType
	var statusReq productpb.ProductStatus
	categoryIdNum, _ := strconv.ParseUint(categoryId, 10, 64)
	kindIdNum, _ := strconv.ParseUint(kindId, 10, 64)
	imageList := strings.Split(images, ",")
	specTypeNum, _ := strconv.Atoi(specType)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	if specTypeNum == 1 {
		specTypeReq = 1
	} else {
		specTypeReq = 2
	}
	statusNum, _ := strconv.Atoi(status)
	if statusNum == 1 {
		statusReq = 1
	} else {
		statusReq = 2
	}
	if len(tagsList) == 0 {
		tagsList = nil
	}

	productParam := productpb.Product{
		CategoryId:       categoryIdNum,
		KindId:           kindIdNum,
		Name:             name,
		ShortDescription: shortDescription,
		Unit:             unit,
		Images:           imageList,
		SpecType:         specTypeReq,
		Status:           statusReq,
		Tags:             tagsList,
		Spec:             specParam,
		Param:            paramList,
		Description:      description,
		AdminId:          adminIdNum,
	}
	return service.NewProduct(m.Context).Add(&productParam)
}

func (m *Product) Edit() error {
	id := m.PostForm("id")
	categoryId := m.PostForm("category_id")
	kindId := m.PostForm("kind_id")
	name := m.PostForm("name")
	shortDescription := m.PostForm("short_description")
	unit := m.PostForm("unit")
	images := m.PostForm("images")
	specType := m.PostForm("spec_type")
	spec := m.PostForm("spec")
	status := m.PostForm("status")
	tags := m.PostForm("tags")
	param := m.PostForm("param")
	description := m.PostForm("description")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	tagsList := make([]uint64, 0, 8)
	valid := validation.Validation{}
	valid.Required(id).Message("请选择要修改的商品")
	valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品数据格式错误")
	valid.Required(categoryId).Message("请选择商品分类")
	valid.Match(categoryId, regexp.MustCompile(`^[0-9]+$`)).Message("商品分类格式错误")
	valid.Required(kindId).Message("请选择商品类型")
	valid.Match(kindId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品类型格式错误")
	valid.Required(name).Message("请填写商品名称")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品名称格式错误")
	valid.Required(shortDescription).Message("请填写商品简介")
	valid.Match(shortDescription, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品简介格式错误")
	valid.Required(unit).Message("请填写商品单位")
	valid.Match(unit, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品单位格式错误")
	valid.Required(images).Message("请上传商品轮播图")
	valid.Match(images, regexp.MustCompile(`^[a-zA-z0-9,]+$`)).Message("商品轮播图格式错误")
	valid.Required(specType).Message("请选择商品规格")
	valid.Match(specType, regexp.MustCompile(`^1|2$`)).Message("商品规格格式错误")
	valid.Required(spec).Message("请填写商品规格数据")
	valid.Required(status).Message("请选择商品上下架状态")
	valid.Match(status, regexp.MustCompile(`^1|2$`)).Message("商品上下架状态格式错误")
	valid.Required(param).Message("请填写商品详情参数")
	valid.Required(description).Message("请填写商品描述")
	//valid.Match(description, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品描述格式错误")
	if valid.HasError() {
		return valid.GetError()
	}
	if len(tags) > 0 {
		err := json.Unmarshal([]byte(tags), &tagsList)
		if err != nil {
			return errors.New("商品标签参数错误！")
		}
	}
	specParam := make([]*productpb.ProductSpec, 0, 8)
	err := json.Unmarshal([]byte(spec), &specParam)
	if err != nil {
		return errors.New("商品规格参数错误！")
	}
	paramList := make([]*productpb.ProductParam, 0, 8)
	err = json.Unmarshal([]byte(param), &paramList)
	if err != nil {
		return errors.New("商品参数格式错误！")
	}

	var specTypeReq productpb.ProductSpecType
	var statusReq productpb.ProductStatus
	idNum, _ := strconv.ParseUint(id, 10, 64)
	categoryIdNum, _ := strconv.ParseUint(categoryId, 10, 64)
	kindIdNum, _ := strconv.ParseUint(kindId, 10, 64)
	imageList := strings.Split(images, ",")
	specTypeNum, _ := strconv.Atoi(specType)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	if specTypeNum == 1 {
		specTypeReq = 1
	} else {
		specTypeReq = 2
	}
	statusNum, _ := strconv.Atoi(status)
	if statusNum == 1 {
		statusReq = 1
	} else {
		statusReq = 2
	}
	if len(tagsList) == 0 {
		tagsList = nil
	}

	productParam := productpb.Product{
		ProductId:        idNum,
		CategoryId:       categoryIdNum,
		KindId:           kindIdNum,
		Name:             name,
		ShortDescription: shortDescription,
		Unit:             unit,
		Images:           imageList,
		SpecType:         specTypeReq,
		Status:           statusReq,
		Tags:             tagsList,
		Spec:             specParam,
		Param:            paramList,
		Description:      description,
		AdminId:          adminIdNum,
	}
	return service.NewProduct(m.Context).Edit(&productParam)
}

func (m *Product) Delete() error {
	id := m.PostForm("id")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	valid := validation.Validation{}
	valid.Required(id).Message("请选择要删除的商品！")
	valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("要删除的商品格式错误")
	if valid.HasError() {
		return valid.GetError()
	}

	idNum, _ := strconv.ParseUint(id, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	delReq := &productpb.DelProductReq{
		ProductId: idNum,
		AdminId:   adminIdNum,
	}
	return service.NewProduct(m.Context).Delete(delReq)
}
