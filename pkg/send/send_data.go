package send

import (
	"context"
)

type SendInt interface {
	Send(c context.Context, req SendMessageRequest) error
}

type SendMessageRequest struct {
	To     []string `json:"to"`
	CC     []string `json:"cc"`
	Mobile string   `json:"mobile"`
}
type Send struct {
	SendInt
}

func (e Send) SendMessage(c context.Context, req SendMessageRequest) error {
	return e.SendInt.Send(c, req)
}
