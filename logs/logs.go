package logs

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"github.com/njun10/golibrary/context"
	"time"
)

var Write = g.Log()
var Id = "logid"

type Logger struct {
	*glog.Logger
}

//生产唯一的logid
func GenLogid() string {
	return NewObjectId().Hex()
}
func Noti(r *ghttp.Request) {
	Write.Async().Notice(r.GetCtxVar(Id), getLogInfo(r))
}

func Error(r *ghttp.Request) {
	Write.Async().Error(r.GetCtxVar(Id), getLogInfo(r))
}


func Info(r *ghttp.Request, data interface{}) {
	Write.Async().Info(r.GetCtxVar(Id), data)
}

func getLogInfo(r *ghttp.Request) interface{} {
	c := context.ContextSer.Get(r.Context())
	c.Time = time.Since(gconv.Time(c.Data["start"])).Milliseconds()
	return c
}
