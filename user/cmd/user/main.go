package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	log2 "github.com/huanmengerkong/example-kratos/log"
	"os"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse()
	/*logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)*/
	logger := log2.NewLogrusLogger()
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	// 注册服务
	// agent := sr.NewAgent("localhost:8500")
	recover(RecoverQuest{
		ServiceName: bc.Server.ServiceName,
		IP:          bc.Server.Grpc.Addr,
		Port:        int(bc.Server.Grpc.Port),
	})
	fmt.Println(bc.Data)
	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

type RecoverQuest struct {
	ServiceName string `json:"service_name"`
	IP          string `json:"ip"`
	Port        int    `json:"port"`
}

func recover(data RecoverQuest) {
	// https://github.com/hashicorp/consul/tree/master/api。
	// 这个例子中，我们首先创建了一个Consul客户端，然后使用该客户端注册服务。接下来，我们使用该客户端获取服务列表，并使用watch机制监视服务列表的变化。当服务列表发生变化时，我们会收到通知并更新服务列表。以下是示例代码：

	// 创建一个新的Consul客户端
	configs := api.DefaultConfig()
	configs.Address = "localhost:8500"
	client, err := api.NewClient(configs)
	if err != nil {
		log.Fatal(err)
	}

	// 注册一个新的服务
	registration := new(api.AgentServiceRegistration)
	registration.ID = data.ServiceName
	registration.Name = data.ServiceName
	registration.Address = data.IP
	registration.Port = data.Port
	registration.Tags = []string{"tag1"}
	registration.Check = &api.AgentServiceCheck{
		TTL: "15",
	}
	/*check := &api.AgentServiceCheck{
		HTTP:                           data.Grpc + "/health",
		Interval:                       "10s",
		Timeout:                        "1s",
		DeregisterCriticalServiceAfter: "1m",
	}
	registration.Check = check*/

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal(err)
	}

	// 获取服务列表

}