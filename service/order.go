package service

import (
	"context"
	"encoding/json"
	"goshop/admin-api/pkg/grpc/gclient"
	"time"

	"github.com/shinmigo/pb/memberpb"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/orderpb"
)

type Order struct {
	*gin.Context
}

type ListOrderRes struct {
	Total  uint64         `json:"total"`
	Orders []*orderDetail `json:"orders"`
}

type orderDetail struct {
	orderpb.OrderDetail
	Member *memberpb.MemberDetail `json:"member"`
}

func NewOrder(ctx *gin.Context) *Order {
	return &Order{Context: ctx}
}

func (o *Order) Index(req *orderpb.ListOrderReq) (*ListOrderRes, error) {
	var (
		orderList  *orderpb.ListOrderRes
		err        error
		orders     = make([]*orderDetail, 0, req.PageSize)
		jsonBytes  []byte
		memberIds  []uint64
		ctx        context.Context
		cancel     context.CancelFunc
		memberMaps = make(map[uint64]*memberpb.MemberDetail)
	)
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	if orderList, err = gclient.OrderClient.GetOrderList(ctx, req); err != nil {
		return nil, err
	}
	cancel()

	if orderList.Total == 0 {
		return &ListOrderRes{
			Total:  0,
			Orders: nil,
		}, nil
	}

	for _, ord := range orderList.Orders {
		memberIds = append(memberIds, ord.MemberId)
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	var memberList *memberpb.ListMemberRes
	if memberList, err = gclient.Member.GetMemberList(ctx, &memberpb.GetMemberReq{
		Page:     1,
		PageSize: uint64(len(memberIds)),
	}); err == nil {
		for _, memberDetail := range memberList.Members {
			memberMaps[memberDetail.MemberId] = memberDetail
		}
	}
	cancel()

	if jsonBytes, err = json.Marshal(orderList.Orders); err != nil {
		return nil, err
	}
	json.Unmarshal(jsonBytes, &orders)
	for _, ord := range orders {
		ord.Member = memberMaps[ord.MemberId]
	}

	return &ListOrderRes{
		Total:  orderList.Total,
		Orders: orders,
	}, err
}
