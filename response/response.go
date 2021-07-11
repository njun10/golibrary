package response

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"github.com/njun10/golibrary/context"
	"github.com/njun10/golibrary/logs"
)

// 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = new(struct{})
	}
	r.Response.WriteJson(context.Response{
		Code:    code,
		Message: message,
		Data:    responseData,
		LogId:   gconv.String(r.GetCtxVar(logs.Id)),
	})
	context.ContextSer.SetErrMsg(r.Context(), code, message)
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	logs.Error(r)
	r.Exit()
}

func WxJson(r *ghttp.Request, res interface{} , data ...interface{}) {
	r.Response.WriteJson(res)
}
