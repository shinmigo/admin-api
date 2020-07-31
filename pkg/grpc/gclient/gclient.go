package gclient

import (
	"fmt"
	"goshop/api/pkg/grpc/etcd3"
	"goshop/api/pkg/utils"
	"log"
	"strings"
	"time"

	"github.com/shinmigo/pb/productpb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var Hello productpb.HelloServiceClient

func DialGrpcService() {
	r := etcd3.NewResolver(utils.C.Etcd.Host)
	resolver.Register(r)

	//这里后面会有多个grpc服务，
	conn, err := grpc.Dial(r.Scheme()+"://author/"+utils.C.Grpc.Name["pms"], grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Panicf("grpc没有连接上%s, err: %v \n", utils.C.Grpc.Name["pms"], err)
	}
	fmt.Printf("连接成功：%s, host分别为: %s \n", utils.C.Grpc.Name["pms"], strings.Join(utils.C.Etcd.Host, ","))
	Hello = productpb.NewHelloServiceClient(conn)

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		resp, err := Hello.Echo(ctx, &productpb.Payload{Data: "hello"}, grpc.FailFast(true))
		cancel()
		if err != nil {
			fmt.Println(err, "err---")
		} else {
			fmt.Println(resp)
		}

		<-time.After(time.Second)
	}
}
