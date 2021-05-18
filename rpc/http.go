package rpc

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
	"github.com/njun10/golibrary/logs"
)

// 标准返回结果数据结构封装。
func (s *HttpRpcInput) Request(r *ghttp.Request) string {
	var log HttpRpcLog
	log.Input = s
	log.Res = "empty"
	start := time.Now()
	// Add user info.
	client := g.Client().Use(ghttp.MiddlewareClientTracing)

	//
	client = client.Timeout(s.Timeout)

	// set cookie
	if len(s.Cookie) > 0 {
		client = client.Cookie(s.Cookie)
	}

	// set header
	if len(s.Header) > 0 {
		client = client.Header(s.Header)
	}

	// set retry
	if s.Retry > 0 {
		client = client.Retry(s.Retry, s.RetryInterval)
	}

	// auth
	if "" != s.AuthName {
		client = client.SetBasicAuth(s.AuthName, s.AuthPass)
	}
	client = client.Ctx(s.Ctx)
	// send
	switch s.Method {
	case "POST":
		if s.Data != nil {
			log.Res = client.PostContent(s.Url, s.Data)
		} else {
			log.Res = client.PostContent(s.Url)
		}
		break
	case "GET":
		if s.Data != nil {
			log.Res = client.GetContent(s.Url, s.Data)
		} else {
			log.Res = client.GetContent(s.Url)
		}
		break
	case "PUT":
		if s.Data != nil {
			log.Res = client.PutContent(s.Url, s.Data)
		} else {
			log.Res = client.PutContent(s.Url)
		}
		break
	default:

	}
	log.Time = time.Since(start).Milliseconds()
	logs.Info(r, log)
	return log.Res
}
