package sr

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"time"
)

type Hconsul struct {
	clientHost string
	client     *api.Client
}

func NewAgent(client string) *Hconsul {
	return &Hconsul{clientHost: client}
}

func (h *Hconsul) RegisterService(ctx context.Context, req RecoverQuest) error {

	// 注册一个新的服务
	registration := new(api.AgentServiceRegistration)
	registration.ID = req.ServiceName
	registration.Name = req.ServiceName
	registration.Address = req.IP
	registration.Port = req.Port
	registration.Tags = []string{"tag1"}
	registration.Check = &api.AgentServiceCheck{
		H2PingUseTLS: false,
		H2PING:       req.IP,
		// TTL: "10s",
		Interval: "5s",
	}
	/*check := &api.AgentServiceCheck{
		HTTP:                           data.Grpc + "/health",
		Interval:                       "10s",
		Timeout:                        "1s",
		DeregisterCriticalServiceAfter: "1m",
	}
	registration.Check = check*/
	err := h.client.Agent().ServiceRegister(registration)
	fmt.Println(111)
	return err
}

func (h *Hconsul) Client(ctx context.Context) error {
	configs := api.DefaultConfig()
	configs.Address = h.clientHost
	client, err := api.NewClient(configs)
	if err != nil {
		panic(fmt.Sprintf("consul err :%v", err))
	}
	h.client = client
	return err
}
func (h *Hconsul) DiscorveryService(ctx context.Context, req RecoverQuest) (svc *api.CatalogService, err error) {
	calot, meta, err := h.client.Catalog().Service(req.ServiceName, req.Tag, &api.QueryOptions{})
	if err != nil {
		return
	}
	if meta.LastIndex == 0 {
		return nil, errors.New("没有服务了")
	}
	for _, service := range calot {
		if service.Checks.AggregatedStatus() == api.HealthPassing {
			svc = service
			break
		}
	}
	return
}
func (h *Hconsul) GetServiceList(ctx context.Context) error {
	services, _, err := h.client.Catalog().Service("my-service", "", nil)
	if err != nil {
		return err
	}
	for _, service := range services {
		fmt.Println(service.ServiceAddress)
		fmt.Println(service.ServicePort)
		fmt.Println(service.ServiceTags)
	}
	return err
}

func (h *Hconsul) Monitoring() error {
	// 监视服务列表的变化
	params := make(map[string]interface{})
	params["type"] = "service"
	params["service"] = "my-service"
	params["passingonly"] = true

	q := api.QueryOptions{
		WaitIndex: 0,
		WaitTime:  time.Minute,
	}

	for {
		services, meta, err := h.client.Catalog().Service("my-service", "", &q)
		if err != nil {
			return err
		}
		q.WaitIndex = meta.LastIndex

		for _, service := range services {
			fmt.Println(service.ServiceAddress)
			fmt.Println(service.ServicePort)
			fmt.Println(service.ServiceTags)
		}
	}
}
