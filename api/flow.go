package api

import "github.com/gogf/gf/net/ghttp"

type Baseflow interface {
	Context(*ghttp.Request)
	DealParams(*ghttp.Request)
	Prepare(*ghttp.Request)
	Action(*ghttp.Request)
	Response(*ghttp.Request)
	Logger(*ghttp.Request)
}

func BfIndex(bf Baseflow, r *ghttp.Request) {
	bf.Context(r)
	bf.DealParams(r)
	bf.Prepare(r)
	bf.Action(r)
	bf.Response(r)
	bf.Logger(r)
	r.Exit()
}
