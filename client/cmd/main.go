package main

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/huanmengerkong/example-kratos/pkg/sr" // 这个地方需要更新tag 才可以拉到最新版本
	v1 "helloworld/api/helloworld/v1"
	_ "image/jpeg"
	"strconv"
	v12 "user/protogo/adminuser/v1"
)

type Person struct {
	Name    string
	Age     string
	Address string
}

type ClientService struct {
	AdminUser  v12.AdminUserClient
	HelloWorld v1.GreeterClient
}

var ClientSvc ClientService

func mainsss() {
	ctx := context.Context(context.Background())
	// 实例化user 服务
	Init(ctx)
	s, err := ClientSvc.AdminUser.FrontedRegister(ctx, &v12.LoginRequest{
		Name:     "test",
		Password: "ps",
		Email:    "e@11.com",
	})
	fmt.Println(s, err)
	return
	// defer conn.Close()

}

func Init(ctx context.Context) {
	// 用户模块（管理员和用户）
	admin(ctx)
	helloworld(ctx)
}
func helloworld(ctx context.Context) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	ClientSvc.HelloWorld = v1.NewGreeterClient(conn)
	return
}

func admin(ctx context.Context) {
	hconsul := sr.NewAgent("localhost:8500")
	err := hconsul.Client(ctx)
	svc, err := hconsul.DiscorveryService(ctx, sr.RecoverQuest{ServiceName: "user"})
	if err != nil {
		return
	}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(svc.ServiceAddress+":"+strconv.Itoa(svc.ServicePort)),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	adminUser := v12.NewAdminUserClient(conn)
	ClientSvc = ClientService{AdminUser: adminUser}
	return
}
