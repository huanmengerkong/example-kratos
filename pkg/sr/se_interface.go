package sr

import "context"

type RecoverQuest struct {
	ServiceName string `json:"service_name"`
	IP          string `json:"ip"`
	Port        int    `json:"port"`
}

type SRInter interface {
	RegisterService(ctx context.Context, req RecoverQuest) error
	GetServiceList(ctx context.Context) error
	Client(ctx context.Context) error
	Monitoring() error
}
