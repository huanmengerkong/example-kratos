package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

// 这个玩意用在服务端的，但是用了grpc 我感觉其实不需要它，直接用grpc 的就可以了
func mainx() {
	router := gin.Default()
	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))
	// /da
	router.GET("say-hello", func(c *gin.Context) {
		name := c.Param("name")
		if name == "error" {
			// 返回kratos error
			kgin.Error(c, errors.Unauthorized("auth_error", "no authentication"))
		} else {
			c.JSON(200, map[string]string{"welcome": name})
		}
	})
	httpSrv := http.NewServer(http.Address(":8000"))
	httpSrv.HandlePrefix("/", router)

	app := kratos.New(
		kratos.Name("gin"),
		kratos.Server(
			httpSrv,
		),
	)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
