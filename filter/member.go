package filter

import (
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
	"regexp"
	"strconv"
	"strings"

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
func (m *Member) Index() (*memberpb.ListMemberRes, error) {
	mobile := m.Query("mobile")
	status := m.Query("status")
	nickname := m.Query("nickname")
	memberId := m.Query("member_id")
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "10")

	var statusNum int
	var memberIdNum uint64
	statusLen := len(status)
	memberIdLen := len(memberId)
	valid := validation.Validation{}
	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的编号 不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页面的数量 不正确")
	if len(nickname) > 0 {
		valid.Match(nickname, regexp.MustCompile(`^[\p{Han}a-zA-Z0-9]+$`)).Message("会员名称格式错误")
	}
	if len(mobile) > 0 {
		valid.Match(mobile, regexp.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$")).Message("会员手机号格式错误")
	}
	if statusLen > 0 {
		valid.Match(status, regexp.MustCompile(`^1|2$`)).Message("会员状态错误")
	}
	if memberIdLen > 0 {
		valid.Match(memberId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("会员编号格式错误")
	}
	if valid.HasError() {
		return nil, valid.GetError()
	}
	if statusLen > 0 {
		statusNum, _ = strconv.Atoi(status)
	}
	if memberIdLen > 0 {
		memberIdNum, _ = strconv.ParseUint(memberId, 10, 64)
	}
	pNumber, _ := strconv.ParseUint(page, 10, 32)
	pSize, _ := strconv.ParseUint(pageSize, 10, 32)
	req := &memberpb.GetMemberReq{
		Nickname: nickname,
		MemberId: memberIdNum,
		Status:   int32(statusNum),
		Mobile:   mobile,
		Page:     pNumber,
		PageSize: pSize,
	}
	list, err := service.NewMember(m.Context).Index(req)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 添加会员
func (m *Member) Add() error {
	nickname := m.PostForm("nickname")
	mobile := m.PostForm("mobile")
	status := m.PostForm("status")
	gender := m.PostForm("gender")
	birthday := m.PostForm("birthday")
	memberLevelId := m.PostForm("member_level_id")
	password := m.PostForm("password")
	operator := m.PostForm("operator")

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
	nickname := m.PostForm("nickname")
	mobile := m.PostForm("mobile")
	memberId := m.PostForm("member_id")
	gender := m.PostForm("gender")
	birthday := m.PostForm("birthday")
	memberLevelId := m.PostForm("member_level_id")
	operator := m.PostForm("operator")

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
func (m *Member) Info() (*memberpb.MemberDetail, error) {
	memberIdParam := m.Query("member_id")
	m.validation.Required(memberIdParam).Message("MemberId不能为空！")

	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}

	memberId, _ := strconv.ParseUint(memberIdParam, 10, 64)
	req, err := service.NewMember(m.Context).Info(memberId)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// 更新会员状态
func (m *Member) EditStatus() error {
	statusParam := m.PostForm("status")
	memberIdParam := m.PostForm("member_id")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)

	valid := validation.Validation{}
	valid.Required(memberIdParam).Message("会员编号不能为空！")
	valid.Match(memberIdParam, regexp.MustCompile("^[1-9][0-9,]*$")).Message("会员编号格式错误")
	valid.Required(statusParam).Message("状态不能为空！")
	valid.Match(statusParam, regexp.MustCompile(`^1|2$`)).Message("会员状态错误")
	if valid.HasError() {
		return valid.GetError()
	}

	status, _ := strconv.ParseInt(statusParam, 10, 32)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	memberIdList := strings.Split(memberIdParam, ",")
	memberIdLen := len(memberIdList)
	memberIdNumList := make([]uint64, 0, memberIdLen)
	for i := range memberIdList {
		id, _ := strconv.ParseUint(memberIdList[i], 10, 64)
		memberIdNumList = append(memberIdNumList, id)
	}
	if err := service.NewMember(m.Context).EditStatus(memberIdNumList, adminIdNum, int32(status)); err != nil {
		return err
	}

	return nil
}
