package rpc

import (
	"context"
	"github.com/gogf/gf/net/gtrace"
	"time"
)

// http 请求的输入参数
type HttpRpcInput struct {
	Url           string `json:"url"`
	Data          interface{} `json:"data"`
	Header        map[string]string `json:"header"`
	Cookie        map[string]string `json:"cookie"`
	Ctx           context.Context `json:"ctx"`
	Timeout       time.Duration `json:"timeout"`
	Retry         int `json:"retry"`
	RetryInterval time.Duration `json:"retryinterval"`
	Method        string `json:"method"`
	AuthName      string `json:"authname"`
	AuthPass      string `json:"authpass"`
}

// http 请求的输入参数
type HttpRpcLog struct {
	Input    *HttpRpcInput `json:"input"`
	Res    string `json:"res"`
	Time int64 `json:"time"`
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
