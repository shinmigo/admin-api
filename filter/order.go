package filter

import (
	"errors"
	"goshop/admin-api/pkg/utils"
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
	"regexp"
	"strconv"

	"github.com/shinmigo/pb/orderpb"

	"github.com/gin-gonic/gin"
)

type Order struct {
	validation validation.Validation
	*gin.Context
}

func NewOrder(ctx *gin.Context) *Order {
	return &Order{Context: ctx}
}

// 订单列表
func (o *Order) Index() (*service.ListOrderRes, error) {
	var (
		id                 = o.Query("id")
		page               = o.DefaultQuery("page", "1")
		pageSize           = o.DefaultQuery("page_size", "10")
		startCreatedAt     = o.Query("start_created_at")
		endCreatedAt       = o.Query("end_created_at")
		orderStatus        = o.Query("order_status")
		startTime, endTime utils.JSONTime
		err                error
		order              service.Order
	)

	valid := validation.Validation{}
	if len(id) > 0 {
		valid.Match(id, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("订单号不正确")
	}

	valid.Match(page, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("页码不正确")
	valid.Match(pageSize, regexp.MustCompile(`^[0-9]{1,3}$`)).Message("每页记录数不正确")

	//验证开始时间，结束时间
	if len(startCreatedAt) > 0 {
		err = startTime.UnmarshalJSON([]byte(startCreatedAt))
		if err != nil {
			return nil, errors.New("开始时间格式不正确, 正确格式如：yyyy-MM-dd HH:mm:ss")
		}
	}
	if len(endCreatedAt) > 0 {
		err = endTime.UnmarshalJSON([]byte(endCreatedAt))
		if err != nil {
			return nil, errors.New("结束时间格式不正确, 正确格式如：yyyy-MM-dd HH:mm:ss")
		}
		if len(startCreatedAt) > 0 && startTime.After(endTime.Time) {
			return nil, errors.New("开始时间不能大于结束时间")
		}
	}

	orderStatusNum, _ := strconv.ParseInt(orderStatus, 10, 32)
	if _, ok := orderpb.OrderStatus_name[int32(orderStatusNum)]; !ok {
		return nil, errors.New("订单状态不存在")
	}

	if valid.HasError() {
		return nil, valid.GetError()
	}

	pageNum, _ := strconv.ParseUint(page, 10, 64)
	pageSizeNum, _ := strconv.ParseUint(pageSize, 10, 64)
	idNum, _ := strconv.ParseUint(id, 10, 64)
	return order.Index(&orderpb.ListOrderReq{
		Page:           pageNum,
		PageSize:       pageSizeNum,
		OrderId:        idNum,
		StartCreatedAt: startCreatedAt,
		EndCreatedAt:   endCreatedAt,
		OrderStatus:    orderpb.OrderStatus(orderStatusNum),
	})
}
