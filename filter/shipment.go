package filter

import (
	"goshop/admin-api/pkg/validation"
	"goshop/admin-api/service"
	"regexp"
	"strconv"

	"github.com/shinmigo/pb/orderpb"

	"github.com/shinmigo/pb/basepb"

	"github.com/gin-gonic/gin"
)

type Shipment struct {
	validation validation.Validation
	*gin.Context
}

func NewShipment(ctx *gin.Context) *Shipment {
	return &Shipment{Context: ctx}
}

// 订单列表
func (s *Shipment) Add() (*basepb.AnyRes, error) {
	var (
		orderId        = s.PostForm("order_id")
		carrierId      = s.PostForm("carrier_id")
		trackingNumber = s.PostForm("tracking_number")
		adminId, _     = s.Get("goshop_user_id")
		adminIdString  = adminId.(string)
	)

	valid := validation.Validation{}
	valid.Required(orderId).Message("订单号不为空")
	valid.Required(carrierId).Message("所选物流不为空")
	valid.Required(trackingNumber).Message("运单号不为空")

	if len(orderId) > 0 {
		valid.Match(orderId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("订单号不正确")
	}

	if len(carrierId) > 0 {
		valid.Match(carrierId, regexp.MustCompile(`^[1-9][0-9]*$`)).Message("物流不正确")
	}

	if valid.HasError() {
		return nil, valid.GetError()
	}

	orderIdNum, _ := strconv.ParseUint(orderId, 10, 64)
	carrierIdNum, _ := strconv.ParseUint(carrierId, 10, 64)
	adminIdNum, _ := strconv.ParseUint(adminIdString, 10, 64)
	return service.NewShipment(s.Context).Add(&orderpb.Shipment{
		OrderId:        orderIdNum,
		CarrierId:      carrierIdNum,
		TrackingNumber: trackingNumber,
		AdminId:        adminIdNum,
	})
}
