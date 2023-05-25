package cmd

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "helloworld/api/helloworld/v1"
)

func main() {
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
	// defer conn.Close()
	cc := v1.NewGreeterClient(conn)
	r, err := cc.SayNihHao(context.Background(), &v1.NiHaoRequest{Name: "test"})
	u := v1.NewUserClient(conn)
	fmt.Println(r)
	s, err := u.GetUser(context.Background(), &v1.GetUserRequest{Id: 1})
	fmt.Println(s, err)
}
