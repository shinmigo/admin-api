package filter

import (
	"goshop/api/pkg/validation"
	"goshop/api/service"
	"regexp"
	"strconv"

	"github.com/shinmigo/pb/shoppb"

	"github.com/gin-gonic/gin"
)

type CarrierCompany struct {
	*gin.Context
}

func NewCarrierCompany(c *gin.Context) *CarrierCompany {
	return &CarrierCompany{Context: c}
}

func (m *CarrierCompany) Index() (*shoppb.ListCarrierRes, error) {
	id := m.Query("id")
	status := m.Query("status")
	code := m.Query("company_code")
	name := m.DefaultQuery("company_name", "")
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")

	var idNum uint64
	idLen := len(id)
	statusLen := len(status)
	valid := validation.Validation{}
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if idLen > 0 {
		valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("商品规格数据不正确")
	}
	if len(code) > 0 {
		valid.Match(code, regexp.MustCompile(`^[A-Za-z]{1,20}$`)).Message("物流编码格式不正确")
	}
	if len(name) > 0 {
		valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("物流名称格式错误")
	}
	if statusLen > 0 {
		valid.Match(status, regexp.MustCompile(`^1|2$`)).Message("状态格式错误！")
	}
	if valid.HasError() {
		return nil, valid.GetError()
	}

	if idLen > 0 {
		idNum, _ = strconv.ParseUint(id, 10, 64)
	}
	pageNum, _ := strconv.ParseUint(page, 10, 64)
	pageSizeNum, _ := strconv.ParseUint(pageSize, 10, 64)
	req := &shoppb.ListCarrierReq{
		Page:     pageNum,
		PageSize: pageSizeNum,
		Id:       idNum,
		Name:     name,
		Code:     code,
		Status:   2,
	}
	if statusLen > 0 {
		var statusNum shoppb.CarrierStatus
		if status == "1" {
			statusNum = shoppb.CarrierStatus_Enabled
		} else {
			statusNum = shoppb.CarrierStatus_Disabled
		}
		req.Status = statusNum
	}

	return service.NewCarrierCompany(m.Context).Index(req)
}

func (m *CarrierCompany) Add() error {
	sort := m.PostForm("sort")
	code := m.PostForm("company_code")
	name := m.PostForm("company_name")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	valid := validation.Validation{}
	valid.Required(name).Message("请填写物流公司名称！")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("物流公司名称格式错误")
	valid.Required(code).Message("请填写物流公司编码！")
	valid.Match(code, regexp.MustCompile(`^[A-Za-z]{1,20}$`)).Message("物流编码格式不正确")
	valid.Required(sort).Message("请填写物流公司排序！")
	valid.Match(sort, regexp.MustCompile(`^[0-9]*$`)).Message("物流公司排序格式错误！")
	if valid.HasError() {
		return valid.GetError()
	}

	sortNum, _ := strconv.ParseUint(sort, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	req := &shoppb.Carrier{
		Name:    name,
		Code:    code,
		Sort:    uint32(sortNum),
		Status:  2,
		AdminId: adminIdNum,
	}
	return service.NewCarrierCompany(m.Context).Add(req)
}

func (m *CarrierCompany) Edit() error {
	id := m.PostForm("id")
	sort := m.PostForm("sort")
	code := m.PostForm("company_code")
	name := m.PostForm("company_name")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	valid := validation.Validation{}
	valid.Required(id).Message("请选择要编辑的数据！")
	valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("物流公司数据格式错误")
	valid.Required(name).Message("请填写物流公司名称！")
	valid.Match(name, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("物流公司名称格式错误")
	valid.Required(code).Message("请填写物流公司编码！")
	valid.Match(code, regexp.MustCompile(`^[A-Za-z]{1,20}$`)).Message("物流编码格式不正确")
	valid.Required(sort).Message("请填写物流公司排序！")
	valid.Match(sort, regexp.MustCompile(`^[0-9]*$`)).Message("物流公司排序格式错误！")
	if valid.HasError() {
		return valid.GetError()
	}

	idNum, _ := strconv.ParseUint(id, 10, 64)
	sortNum, _ := strconv.ParseUint(sort, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	req := &shoppb.Carrier{
		CarrierId: idNum,
		Name:      name,
		Code:      code,
		Sort:      uint32(sortNum),
		Status:    2,
		AdminId:   adminIdNum,
	}
	return service.NewCarrierCompany(m.Context).Edit(req)
}

func (m *CarrierCompany) Delete() error {
	id := m.PostForm("id")

	valid := validation.Validation{}
	valid.Required(id).Message("请选择要删除的数据！")
	valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("物流公司数据格式错误")
	if valid.HasError() {
		return valid.GetError()
	}

	idNum, _ := strconv.ParseUint(id, 10, 64)
	req := &shoppb.DelCarrierReq{
		CarrierId: idNum,
	}
	return service.NewCarrierCompany(m.Context).Delete(req)
}
