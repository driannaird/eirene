package helper

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (r *Response) Error() string {
	return r.Message
}

func Success(msg string, data interface{}) error {
	return &Response{
		Code:    200,
		Message: msg,
		Data:    data,
	}
}

func InternalServerError(msg string) error {
	return &Response{
		Code:    500,
		Message: msg,
	}
}

func BadRequest(msg string) error {
	return &Response{
		Code:    400,
		Message: msg,
	}
}

func Created(msg string) error {
	return &Response{
		Code:    201,
		Message: msg,
	}
}

func NotFound(msg string) error {
	return &Response{
		Code:    404,
		Message: msg,
	}
}

func Unauthorize(msg string) error {
	return &Response{
		Code:    401,
		Message: msg,
	}
}

func Forbidden(msg string) error {
	return &Response{
		Code:    403,
		Message: msg,
	}
}
