package filter

import (
	"goshop/api/pkg/validation"
	"goshop/api/service"
	"regexp"
	"strconv"
	
	"github.com/shinmigo/pb/productpb"
	
	"github.com/gin-gonic/gin"
)

type ProdcutParam struct {
	*gin.Context
}

func NewProductParam(c *gin.Context) *ProdcutParam {
	return &ProdcutParam{Context: c}
}

func (m *ProdcutParam) Index() (*productpb.ListParamRes, error) {
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
	list, err := service.NewProductParam(m.Context).Index(pNumber, pSize)
	if err != nil {
		return nil, err
	}
	
	return list, nil
}
