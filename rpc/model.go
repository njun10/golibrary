package rpc

import (
	"context"
	"github.com/gogf/gf/net/gtrace"
	"time"
)

// http 请求的输入参数
type HttpRpcInput struct {
	Url           string
	Data          interface{}
	Header        map[string]string
	Cookie        map[string]string
	Ctx           context.Context
	Timeout       time.Duration
	Retry         int
	RetryInterval time.Duration
	Method        string
	AuthName      string
	AuthPass      string
}

func Init() *HttpRpcInput {
	ctx, span := gtrace.NewSpan(context.Background(), "HttpRpcRequests")
	defer span.End()
	return &HttpRpcInput{
		Url:           "",
		Data:          nil,
		Header:        make(map[string]string),
		Cookie:        make(map[string]string),
		Timeout:       100000000,
		Retry:         3,
		RetryInterval: 1000000,
		Method:        "POST",
		AuthPass:      "",
		AuthName:      "",
		Ctx:           ctx,
	}
}
