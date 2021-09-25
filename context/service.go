package context

import (
	"context"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"time"
)

// 上下文管理服务
var ContextSer = new(contextService)

type contextService struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextService) Init(r *ghttp.Request) {
	c := new(ContextM)
	c.Res = new(struct{})
	c.Agent = new(ContextAgent)
	c.Params = new(struct{})
	c.Data = make(map[string]interface{})
	c.Data["start"] = time.Now()
	r.ParseQuery(&c.Agent)
	c.Agent.RemoteIp = r.GetClientIp()
	r.SetCtxVar(ContextKey, c)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *contextService) Get(ctx context.Context) *ContextM {
	value := ctx.Value(ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*ContextM); ok {
		return localCtx
	}
	return nil
}

// 将错误信息标识 记录到上下文
func (s *contextService) SetErrMsg(ctx context.Context, error int, msg string) {
	s.Get(ctx).Error = error
	s.Get(ctx).Msg = msg
}

// 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *contextService) SetAgent(ctx context.Context, agent *ContextAgent) {
	s.Get(ctx).Agent = agent
}

// 将接口需要的参数加入上下文中
func (s *contextService) SetParams(ctx context.Context, params interface{}) {
	s.Get(ctx).Params = params
}

// 将返回的data加入上下文中
func (s *contextService) SetResData(ctx context.Context, data interface{}) {
	s.Get(ctx).Res = data
}

// 向返回的data加入上下文中 添加数据
func (s *contextService) AddResData(ctx context.Context, key string, add interface{}) {
	addmap := gconv.Map(s.Get(ctx).Res)
	addmap[key] = add
	s.Get(ctx).Res = addmap
}

// 向context的data加入上下文中 添加数据
func (s *contextService) AddContextData(ctx context.Context, key string, add interface{}) {
	addmap := gconv.Map(s.Get(ctx).Data)
	addmap[key] = add
	s.Get(ctx).Data = addmap
}
// SetExtra 设置额外信息用于埋点
func (s *contextService) SetExtra(ctx context.Context, extra interface{}) {
	s.Get(ctx).Extra = extra
}

// AddExtraData 向Extra中添加信息
func (s *contextService) AddExtraData(ctx context.Context, key string, value interface{})  {
	extra := gconv.Map(s.Get(ctx).Extra)
	extra[key] = value
	s.Get(ctx).Extra = extra
}