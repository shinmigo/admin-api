package filter

import (
	"goshop/api/pkg/validation"
	"goshop/api/service"
	"regexp"
	"strconv"

	"github.com/shinmigo/pb/memberpb"

	"github.com/gin-gonic/gin"
)

type Member struct {
	validation validation.Validation
	*gin.Context
}

func NewMember(c *gin.Context) *Member {
	return &Member{Context: c, validation: validation.Validation{}}
}

// 会员列表
func (m *Member) Index() (*memberpb.ListRes, error) {
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")
	m.validation.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	m.validation.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}

	pNumber, _ := strconv.ParseUint(page, 10, 32)
	pSize, _ := strconv.ParseUint(pageSize, 10, 32)
	list, err := service.NewMember(m.Context).Index(pNumber, pSize)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 添加会员
func (m *Member) Add() error {
	nickname := m.Query("nickname")
	mobile := m.Query("mobile")
	status := m.Query("status")
	gender := m.Query("gender")
	birthday := m.Query("birthday")
	memberLevelId := m.Query("member_level_id")
	password := m.Query("password")
	operator := m.Query("operator")
	m.validation.Required(nickname).Message("昵称不能为空！")
	m.validation.Mobile(mobile).Message("手机号格式不正确！")
	m.validation.Required(status).Message("状态不能为空！")
	m.validation.Required(gender).Message("性别不能为空！")
	m.validation.Required(birthday).Message("生日不能为空！")
	m.validation.Required(memberLevelId).Message("会员等级不能为空！")
	m.validation.Required(password).Message("密码不能为空！")
	m.validation.Required(operator).Message("操作人不能为空！")

	if m.validation.HasError() {
		return m.validation.GetError()
	}

	if err := service.NewMember(m.Context).Add(); err != nil {
		return err
	}

	return nil
}

// 会员编辑
func (m *Member) Edit() error {
	nickname := m.Query("nickname")
	mobile := m.Query("mobile")
	memberId := m.Query("member_id")
	gender := m.DefaultQuery("gender", "0")
	birthday := m.Query("birthday")
	memberLevelId := m.Query("member_level_id")
	operator := m.Query("operator")

	m.validation.Required(nickname).Message("昵称不能为空！")
	m.validation.Mobile(mobile).Message("手机号格式不正确！")
	m.validation.Required(gender).Message("性别不能为空！")
	m.validation.Required(birthday).Message("生日不能为空！")
	m.validation.Required(memberLevelId).Message("会员等级不能为空！")
	m.validation.Required(operator).Message("操作人不能为空！")
	m.validation.Required(memberId).Message("member_id不能为空！")

	if m.validation.HasError() {
		return m.validation.GetError()
	}

	if err := service.NewMember(m.Context).Edit(); err != nil {
		return err
	}

	return nil
}

// 会员详情
func (m *Member) Info() (*memberpb.Member, error) {
	memberId := m.Query("member_id")
	m.validation.Required(memberId).Message("MemberId不能为空！")

	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}

	req, err := service.NewMember(m.Context).Info()
	if err != nil {
		return nil, err
	}

	return req, nil
}

// 更新会员状态
func (m *Member) EditStatus() error {
	status := m.Query("status")
	memberId := m.Query("member_id")
	operator := m.Query("operator")

	m.validation.Required(operator).Message("操作人不能为空！")
	m.validation.Required(memberId).Message("member_id不能为空！")
	m.validation.Required(status).Message("状态不能为空！")

	if m.validation.HasError() {
		return m.validation.GetError()
	}

	if err := service.NewMember(m.Context).EditStatus(); err != nil {
		return err
	}

	return nil
}
