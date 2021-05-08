package rpc

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/njun10/golibrary/logs"
)


// 标准返回结果数据结构封装。
func (s *HttpRpcInput) Request() {
	var res = ""
	// Add user info.
	client := g.Client().Use(ghttp.MiddlewareClientTracing)

	//
	client = client.Timeout(s.Timeout)

	// set cookie
	if len(s.Cookie)>0 {
		client = client.Cookie(s.Cookie)
	}

	// set header
	if len(s.Header)>0 {
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
				res = client.PostContent(s.Url, s.Data)
			}else{
				res = client.PostContent(s.Url)
			}
			break
		case "GET":
			if s.Data != nil {
				res = client.GetContent(s.Url, s.Data)
			}else{
				res = client.GetContent(s.Url)
			}
			break
		default:

	}
	logs.Write.Async().Info(s, res)
}


