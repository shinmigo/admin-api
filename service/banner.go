package service

import (
	"context"
	"goshop/admin-api/pkg/grpc/gclient"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/shoppb"
)

type BannerAd struct {
	*gin.Context
}

func NewBannerAd(c *gin.Context) *BannerAd {
	return &BannerAd{Context: c}
}

func (m *BannerAd) Index(param *shoppb.ListBannerAdReq) (*shoppb.ListBannerAdRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	list, err := gclient.BannerAdClient.GetBannerAdList(ctx, param)
	cancel()

	return list, err
}

func (m *BannerAd) Add(param *shoppb.BannerAd) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.BannerAdClient.AddBannerAd(ctx, param)
	cancel()

	return err
}

func (m *BannerAd) Edit(param *shoppb.BannerAd) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.BannerAdClient.EditBannerAd(ctx, param)
	cancel()

	return err
}

func (m *BannerAd) EditStatus(param *shoppb.EditBannerAdStatusReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.BannerAdClient.EditBannerAdStatus(ctx, param)
	cancel()

	return err
}

func (m *BannerAd) Delete(param *shoppb.DelBannerAdReq) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	_, err := gclient.BannerAdClient.DelBannerAd(ctx, param)
	cancel()

	return err
}
