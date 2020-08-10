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
