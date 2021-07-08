// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package context

const (
	// 上下文变量存储键名
	ContextKey = "ContextKey"
)

// 请求上下文结构
type ContextM struct {
	Error  int           `json:"error"`
	Msg    string        `json:"msg"`
	Agent  *ContextAgent `json:"agent"` // 上下文用户信息
	Params interface{}   `json:"params"`
	Data   map[string]interface{}   `json:"data"`
	Res    interface{}   `json:"res"`
	Time   int64         `json:"time"` // 请求耗时
}

// 请求上下文中的用户信息
type ContextAgent struct {
	AgentId int64
	PlatId  int
	Uid     int64
	Lat float64
	Lng float64
	Channel string
	Version string
}

// 数据返回通用JSON数据结构
type Response struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
	LogId   string      `json:"log_id"`
}
