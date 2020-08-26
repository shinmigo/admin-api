package service

import (
	"context"
	"fmt"
	"strconv"
	"time"
	
	"goshop/api/pkg/grpc/gclient"
	
	"github.com/shinmigo/pb/basepb"
	
	"github.com/shinmigo/pb/memberpb"
	
	"github.com/gin-gonic/gin"
)

type Member struct {
	*gin.Context
}

func NewMember(c *gin.Context) *Member {
	return &Member{Context: c}
}

// 会员列表
func (m *Member) Index(pNumber, pSize uint64) (*memberpb.ListMemberRes, error) {
	req := &memberpb.GetMemberReq{
		Page:     pNumber,
		PageSize: pSize,
	}
	
	if len(m.Query("mobile")) > 0 {
		req.Mobile = m.Query("mobile")
	}
	
	if len(m.Query("member_id")) > 0 {
		id, _ := strconv.ParseUint(m.Query("member_id"), 10, 64)
		req.MemberId = id
	}
	
	if len(m.Query("status")) > 0 {
		status, _ := strconv.ParseInt(m.Query("status"), 10, 32)
		req.Status = int32(status)
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.Member.GetMemberList(ctx, req)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取会员列表失败， err：%v", err)
	}
	
	return resp, nil
}

// 添加会员
func (m *Member) Add() error {
	nickname := m.PostForm("nickname")
	mobile := m.PostForm("mobile")
	statusParam := m.PostForm("status")
	genderParam := m.PostForm("gender")
	birthday := m.PostForm("birthday")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)
	memberLevelIdParam := m.PostForm("member_level_id")
	status, _ := strconv.ParseInt(statusParam, 10, 32)
	gender, _ := strconv.ParseInt(genderParam, 10, 32)
	memberLevelId, _ := strconv.ParseUint(memberLevelIdParam, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	
	req := &memberpb.Member{
		Nickname:      nickname,
		Mobile:        mobile,
		Status:        memberpb.MemberStatus(status),
		Gender:        memberpb.MemberGender(gender),
		Birthday:      birthday,
		MemberLevelId: memberLevelId,
		AdminId:       adminIdNum,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.Member.AddMember(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("添加会员失败, err:%v", err)
	}
	
	if resp.State == 0 {
		return fmt.Errorf("添加失败")
	}
	
	return nil
}

// 会员编辑
func (m *Member) Edit() error {
	nickname := m.PostForm("nickname")
	mobile := m.PostForm("mobile")
	memberIdParam := m.PostForm("member_id")
	genderParam := m.PostForm("gender")
	birthday := m.PostForm("birthday")
	adminId, _ := m.Get("goshop_user_id")
	adminIdString, _ := adminId.(string)
	memberLevelIdParam := m.PostForm("member_level_id")
	gender, _ := strconv.ParseInt(genderParam, 10, 32)
	memberId, _ := strconv.ParseUint(memberIdParam, 10, 32)
	memberLevelId, _ := strconv.ParseUint(memberLevelIdParam, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	
	req := &memberpb.Member{
		Nickname:      nickname,
		Mobile:        mobile,
		Gender:        memberpb.MemberGender(gender),
		Birthday:      birthday,
		MemberLevelId: memberLevelId,
		MemberId:      memberId,
		AdminId:       adminIdNum,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.Member.EditMember(ctx, req)
	cancel()
	if err != nil {
		return fmt.Errorf("更新会员失败, err:%v", err)
	}
	
	if resp.State == 0 {
		return fmt.Errorf("更新失败")
	}
	
	return nil
}

// 会员详情
func (m *Member) Info(memberId uint64) (*memberpb.MemberDetail, error) {
	req := &basepb.GetOneReq{
		Id: memberId,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.Member.GetMemberDetail(ctx, req)
	cancel()
	
	if err != nil {
		return nil, fmt.Errorf("获取会员失败, err: %v", err)
	}
	
	return resp, nil
}

// 更新会员状态
func (m *Member) EditStatus(memberId, adminId uint64, status int32) error {
	req := &basepb.EditStatusReq{
		Id:      memberId,
		Status:  status,
		AdminId: adminId,
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.Member.EditMemberStatus(ctx, req)
	cancel()
	
	if err != nil {
		return fmt.Errorf("更新会员状态失败, err:%v", err)
	}
	
	if resp.State == 0 {
		return fmt.Errorf("更新失败")
	}
	
	return nil
}
