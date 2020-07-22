package main

import (
	"goshop/api/command/user"
	"goshop/api/pkg/grpc/gclient"
)

func initService() {
	go gclient.DialGrpcService()
	go user.Hello()
}
