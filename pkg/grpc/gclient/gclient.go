package gclient

import (
	"fmt"
	"goshop/api/pkg/grpc/etcd3"
	"goshop/api/pkg/utils"
	"log"
	"strings"

	"github.com/shinmigo/pb/memberpb"

	"github.com/shinmigo/pb/productpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

var (
	ProductTag            productpb.TagServiceClient
	ProductParam          productpb.ParamServiceClient
	ProductKind           productpb.KindServiceClient
	Member                memberpb.MemberServiceClient
	ProductCategoryClient productpb.CategoryServiceClient
)

func DialGrpcService() {
	pms()
	crm()
}

func crm() {
	r := etcd3.NewResolver(utils.C.Etcd.Host)
	resolver.Register(r)
	conn, err := grpc.Dial(r.Scheme()+"://author/"+utils.C.Grpc.Name["crm"], grpc.WithBalancerName("round_robin"), grpc.WithInsecure())
	if err != nil {
		log.Panicf("grpc没有连接上%s, err: %v \n", utils.C.Grpc.Name["crm"], err)
	}
	fmt.Printf("连接成功：%s, host分别为: %s \n", utils.C.Grpc.Name["crm"], strings.Join(utils.C.Etcd.Host, ","))
	Member = memberpb.NewMemberServiceClient(conn)
}

func pms() {
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
	ProductCategoryClient = productpb.NewCategoryServiceClient(conn)
	ProductKind = productpb.NewKindServiceClient(conn)
}
