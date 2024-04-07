package entity

// 定义result结构体
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// 定义错误码
const (
	// 成功
	CODE_SUCCESS int = 1
	// 失败
	CODE_ERROR int = -1
	// 自定义
)

// 设置错误码 使用结构体的函数参数叫方法，他紧跟着func 函数名称在方法之后
func (res *Result) SetCode(code int) *Result {
	res.Code = code
	return res
}

// 设置错误信息
func (res *Result) SetMessage(msg string) *Result {
	res.Message = msg
	return res
}

// 设置返回数据
func (res *Result) SetData(data interface{}) *Result {
	res.Data = data
	return res
}
