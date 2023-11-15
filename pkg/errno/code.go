package errno

var (
	// OK 代表请求成功.
	OK = &Errno{Code: 0, Msg: "ok", Desc: ""}

	// InternalServerError 表示所有未知的服务器端错误.
	InternalServerError = &Errno{Code: 500, Msg: "InternalError", Desc: "Internal usercenter error."}

	// ErrPageNotFound 表示路由不匹配错误.
	ErrPageNotFound = &Errno{Code: 404, Msg: "Page not found.", Desc: ""}

	ErrParameterInvalid = &Errno{
		Code: 40000,
		Msg:  "参数校验错误",
		Desc: "",
	}

	ErrEntityNull = &Errno{
		Code: 40400,
		Msg:  "实体对象不存在",
		Desc: "",
	}

	ErrDBFailed = &Errno{
		Code: 50001,
		Msg:  "数据库操作异常",
		Desc: "",
	}

	ErrUnauthorization = &Errno{
		Code: 40100,
		Msg:  "用户认证失败",
		Desc: "",
	}

	ErrForbidden = &Errno{
		Code: 40300,
		Msg:  "无权限",
		Desc: "",
	}

	ErrEntityExists = &Errno{
		Code: 40002,
		Msg:  "实体对象已存在",
		Desc: "",
	}
)
