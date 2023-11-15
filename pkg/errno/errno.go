package errno

import "fmt"

type Errno struct {
	Code int
	Msg  string
	Desc string
}

func (e *Errno) Error() string {
	return e.Msg
}

// SetMessage 设置 Errno 类型错误中的 Msg 字段.
func (e *Errno) SetMessage(format string, args ...interface{}) *Errno {
	e.Msg = fmt.Sprintf(format, args...)
	return e
}

// SetDescription 设置 Errno 类型错误中的 Desc 字段.
func (e *Errno) SetDescription(format string, args ...interface{}) *Errno {
	e.Desc = fmt.Sprintf(format, args...)
	return e
}

// Decode 尝试从 err 中解析出业务错误码和错误信息.
func Decode(err error) (code int, msg string, desc string) {
	if err == nil {
		return OK.Code, OK.Msg, OK.Desc
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.Code, typed.Msg, typed.Desc
	default:
		// 默认返回未知错误码和错误信息. 该错误代表服务端出错
		return InternalServerError.Code, InternalServerError.Msg, err.Error()
	}
}
