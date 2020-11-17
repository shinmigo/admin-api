package controller

import (
	"bytes"
	"context"
	"fmt"
	"goshop/admin-api/pkg/grpc/gclient"
	"io"
	"io/ioutil"

	"github.com/shinmigo/pb/shoppb"
)

type Image struct {
	Base
}

func (m *Image) Initialise() {

}

func (m *Image) GetImage() {
	name := m.DefaultQuery("name", "")
	if len(name) == 0 {
		m.SetResponse(nil, fmt.Errorf("文件名不合法"))
		return
	}

	req := shoppb.GetImageReq{
		ImageId: m.DefaultQuery("name", ""),
	}

	res, err := gclient.ImageClient.GetImage(context.Background(), &req)
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	_, _ = io.Copy(m.Writer, bytes.NewBuffer(res.Content))
}

func (m *Image) Upload() {
	f, head, err := m.Request.FormFile("my_file")
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	b, _ := ioutil.ReadAll(f)
	req := shoppb.UploadReq{
		Content: b,
		Name:    head.Filename,
	}
	res, err := gclient.ImageClient.Upload(context.Background(), &req)
	if err != nil {
		m.SetResponse(nil, err)
		return
	}

	m.SetResponse(res.ImageId)
}
