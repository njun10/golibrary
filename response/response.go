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
	var response context.Response
	con := context.ContextSer.Get(r.Context())
	gconv.Struct(con.Response, &response)

	r.Response.WriteJson(context.Response{
		Code:    code,
		Message: message,
		Data:    responseData,
		LogId:   gconv.String(r.GetCtxVar(logs.Id)),
		StChange: response.StChange,
		StChannel1: response.StChannel1,
		StChannel2: response.StChannel2,
		StChannel3: response.StChannel3,
		StType: response.StType,
		StId: response.StId,
		PlatId: response.PlatId,
		Version: response.Version,
		Channel: response.Channel,
		PlatForm: response.PlatForm,
	})
	context.ContextSer.SetErrMsg(r.Context(), code, message)
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	logs.Noti(r)
	r.Exit()
}

func WxJson(r *ghttp.Request, res interface{} , data ...interface{}) {
	r.Response.WriteJson(res)
}
