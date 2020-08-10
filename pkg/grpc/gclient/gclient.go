package gclient

import (
	"fmt"
	"github.com/shinmigo/pb/memberpb"
	"goshop/api/pkg/grpc/etcd3"
	"goshop/api/pkg/utils"
	"log"
	"strings"

	"github.com/shinmigo/pb/productpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var ProductTag productpb.TagServiceClient
var ProductParam productpb.ParamServiceClient
var Member memberpb.MemberServiceClient

func DialGrpcService() {
	r := etcd3.NewResolver(utils.C.Etcd.Host)
	resolver.Register(r)

	//这里后面会有多个grpc服务，
	conn, err := grpc.Dial(r.Scheme()+"://author/"+utils.C.Grpc.Name["pms"], grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Panicf("grpc没有连接上%s, err: %v \n", utils.C.Grpc.Name["pms"], err)
	}
	fmt.Printf("连接成功：%s, host分别为: %s \n", utils.C.Grpc.Name["pms"], strings.Join(utils.C.Etcd.Host, ","))
	ProductTag = productpb.NewTagServiceClient(conn)
	ProductParam = productpb.NewParamServiceClient(conn)

	connCrm, err := grpc.Dial(r.Scheme()+"://author/"+utils.C.Grpc.Name["crm"], grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Panicf("grpc没有连接上%s, err: %v \n", utils.C.Grpc.Name["crm"], err)
	}
	fmt.Printf("连接成功：%s, host分别为: %s \n", utils.C.Grpc.Name["crm"], strings.Join(utils.C.Etcd.Host, ","))
	Member = memberpb.NewMemberServiceClient(connCrm)

	//for {
	//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//	resp, err := Hello.Echo(ctx, &productpb.Payload{Data: "hello"}, grpc.FailFast(true))
	//	cancel()
	//	if err != nil {
	//		fmt.Println(err, "err---")
	//	} else {
	//		fmt.Println(resp)
	//	}
	//
	//	<-time.After(time.Second)
	//}
}
