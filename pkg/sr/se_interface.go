package sr

import (
	"context"
	"github.com/hashicorp/consul/api"
)

type RecoverQuest struct {
	ServiceName string `json:"service_name"`
	IP          string `json:"ip"`
	Port        int    `json:"port"`
	Tag         string `json:"tag"`
}

type SRInter interface {
	RegisterService(ctx context.Context, req RecoverQuest) error
	GetServiceList(ctx context.Context) error
	Client(ctx context.Context) error
	Monitoring() error
	DiscorveryService(ctx context.Context, req RecoverQuest) (svc *api.CatalogService, err error)
}
