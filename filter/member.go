package filter

import (
	"github.com/shinmigo/pb/memberpb"
	"goshop/api/pkg/validation"
	"goshop/api/service"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Member struct {
	*gin.Context
}

func NewMember(c *gin.Context) *Member {
	return &Member{Context: c}
}

func (m *Member) Index() (*memberpb.ListRes, error) {
	valid := validation.Validation{}
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if valid.HasError() {
		return nil, valid.GetError()
	}


	pNumber, _ := strconv.ParseUint(page, 10, 32)
	pSize, _ := strconv.ParseUint(pageSize, 10, 32)
	list, err := service.NewMember(m.Context).Index(pNumber, pSize)
	if err != nil {
		return nil, err
	}
	
	return list, nil
}
