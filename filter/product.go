package filter

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/api/pkg/validation"
	"goshop/api/service"
	"regexp"
	"strconv"
	"strings"
)

type Product struct {
	*gin.Context
}

func NewProduct(c *gin.Context) *Product {
	return &Product{Context: c}
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

	tagsList := make([]uint64, 0, 8)
	valid := validation.Validation{}
	valid.Required(categoryId).Message("请选择商品分类")
	valid.Match(categoryId, regexp.MustCompile(`^[1-9][0-9]$`)).Message("商品分类格式错误")
	valid.Required(kindId).Message("请选择商品类型")
	valid.Match(kindId, regexp.MustCompile(`^[1-9][0-9]$`)).Message("商品类型格式错误")
	valid.Required(name).Message("请填写商品名称")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]{1,30}$`)).Message("商品名称格式错误")
	valid.Required(shortDescription).Message("请填写商品简介")
	valid.Match(shortDescription, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]{1,50}$`)).Message("商品简介格式错误")
	valid.Required(unit).Message("请填写商品单位")
	valid.Match(unit, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品单位格式错误")
	valid.Required(images).Message("请上传商品轮播图")
	valid.Match(images, regexp.MustCompile(`^(http://|https://)[a-zA-z0-9.,]+$`)).Message("商品轮播图格式错误")
	valid.Required(specType).Message("请选择商品规格")
	valid.Match(specType, regexp.MustCompile(`^1|2$`)).Message("商品规格格式错误")
	valid.Required(spec).Message("请填写商品规格数据")
	valid.Required(status).Message("请选择商品上下架状态")
	valid.Match(status, regexp.MustCompile(`^1|2$`)).Message("商品上下架状态格式错误")
	valid.Required(param).Message("请填写商品详情参数")
	valid.Required(description).Message("请填写商品描述")
	valid.Match(description, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("商品描述格式错误")
	if valid.HasError() {
		return valid.GetError()
	}
	if len(tags) > 0 {
		err := json.Unmarshal([]byte(tags), &tagsList)
		if err != nil {
			return err
		}
	}
	specParam := make([]*productpb.ProductSpec, 0, 8)
	err := json.Unmarshal([]byte(spec), &specParam)
	if err != nil {
		return err
	}
	paramList := make([]*productpb.ProductParam, 0, 8)
	err = json.Unmarshal([]byte(param), &paramList)
	if err != nil {
		return err
	}

	var specTypeReq productpb.ProductSpecType
	var statusReq productpb.ProductStatus
	categoryIdNum, _ := strconv.ParseUint(categoryId, 10, 64)
	kindIdNum, _ := strconv.ParseUint(kindId, 10, 64)
	imageList := strings.Split(images, ",")
	specTypeNum, _ := strconv.Atoi(specType)
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
	}
	return service.NewProduct(m.Context).Add(&productParam)
}
