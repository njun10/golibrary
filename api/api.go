package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/njun10/golibrary/context"
	"github.com/njun10/golibrary/logs"
	"github.com/njun10/golibrary/response"
)

var BaseApi = new(Base)
var ContextInfo = new(context.ContextM)
var ReqInfo = new(ghttp.Request)

// api 基础
type Base struct {
}

//初始化函数
func init() {

}

// @summary 入口函数
// @tags    基础类服务
// @produce json
// @param   entity  body
// @router
// @success 200 {object} response.JsonResponse "执行结果"
func (b *Base) Index(r *ghttp.Request) {
	BfIndex(b, r)
}

// 初始化通用上下文信息
func (b *Base) Context(r *ghttp.Request) {
	var logId = r.Header.Get("LogId")
	if len(logId) > 0 {
		r.SetCtxVar(logs.Id, logId)
	} else {
		r.SetCtxVar(logs.Id, logs.GenLogid())
	}
	context.ContextSer.Init(r)
	agent := new(context.ContextAgent)
	if err := r.ParseForm(&agent); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
}

// 接口参数处理
func (b *Base) DealParams(r *ghttp.Request) {
}

// 准备数据
func (b *Base) Prepare(r *ghttp.Request) {
}

// 主业务逻辑
func (b *Base) Action(r *ghttp.Request) {
}

// 处理返回参数
func (b *Base) Response(r *ghttp.Request) {
	response.Json(r, 0, "success", context.ContextSer.Get(r.Context()).Res)
}

// 处理日志
func (b *Base) Logger(r *ghttp.Request) {
	logs.Noti(r)
}
