package logs

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
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
