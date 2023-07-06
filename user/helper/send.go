package helper

import "github.com/huanmengerkong/example-kratos/pkg/send"

type Email struct {
}

func (e *Email) Send(c context.Context, req send.SendMessageRequest) error {
	return nil
}
