package service

import (
	"context"
	"fmt"
	"github.com/shinmigo/pb/memberpb"
	"goshop/api/pkg/grpc/gclient"
	"strconv"
	"time"
	
	"github.com/gin-gonic/gin"
)

type Member struct {
	*gin.Context
}

func NewMember(c *gin.Context) *Member {
	return &Member{Context: c}
}

// 会员列表
func (m *Member) Index(pNumber, pSize uint64) (*memberpb.ListRes, error) {
	req := &memberpb.ListReq{
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
		status, _ := strconv.ParseUint(m.Query("status"), 10, 32)
		req.Status = uint32(status)
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.Member.GetList(ctx, req)
	cancel()
	if err != nil {
		return nil, fmt.Errorf("获取会员列表失败， err：%v", err)
	}
	
	return resp, nil
}

// 添加会员
func (m *Member) Add() error {
	nickname := m.Query("nickname")
	mobile := m.Query("mobile")
	statusParam := m.Query("status")
	genderParam := m.DefaultQuery("gender","0")
	birthday := m.Query("birthday")
	memberLevelId := m.Query("member_level_id")
	password := m.Query("password")
	operator := m.Query("operator")
	status,_ := strconv.ParseUint(statusParam,10,32)
	gender,_ := strconv.ParseUint(genderParam,10,32)

	req := &memberpb.AddReq{
		Nickname: nickname,
		Mobile: mobile,
		Status: uint32(status),
		Gender: uint32(gender),
		Birthday: birthday,
		MemberLevelId: memberLevelId,
		Password: password,
		Operator: operator,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.Member.Add(ctx,req)
	cancel()
	if err!= nil {
		return fmt.Errorf("添加会员失败, err:%v", err)
	}

	if resp.State == 0 {
		return fmt.Errorf("添加失败")
	}

	return nil
}

// 会员编辑
func (m *Member) Edit() error {
	nickname := m.Query("nickname")
	mobile := m.Query("mobile")
	memberIdParam := m.Query("member_id")
	genderParam := m.DefaultQuery("gender","0")
	birthday := m.Query("birthday")
	memberLevelId := m.Query("member_level_id")
	operator := m.Query("operator")
	gender,_ := strconv.ParseUint(genderParam,10,32)
	memberId,_ := strconv.ParseUint(memberIdParam,10,32)

	req := &memberpb.EditReq{
		Nickname: nickname,
		Mobile: mobile,
		Gender: uint32(gender),
		Birthday: birthday,
		MemberLevelId: memberLevelId,
		Operator: operator,
		MemberId: memberId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := gclient.Member.Edit(ctx,req)
	cancel()
	if err!= nil {
		return fmt.Errorf("添加会员失败, err:%v",err)
	}

	if resp.State == 0 {
		return fmt.Errorf("添加失败")
	}

	return nil
}
