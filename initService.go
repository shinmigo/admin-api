package main

import (
	"goshop/admin-api/command/user"
	"goshop/admin-api/pkg/grpc/gclient"
)

func initService() {
	go gclient.DialGrpcService()
	go user.Hello()
}
