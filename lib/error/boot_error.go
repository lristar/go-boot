package bootError

import "fmt"

type IError interface {
	Error() string
	Code() int
}

// BootError 请求类型错误
type BootError struct {
	msg  string
	code int
}

func (r BootError) Error() string {
	return r.msg
}

// New 新建error
func New(msg string, code int) BootError {
	return BootError{msg: msg, code: code}
}

// Errorf 输出msg并新建error
func Errorf(format string, code int, a ...interface{}) BootError {
	return BootError{msg: fmt.Sprintf(format, a...), code: code}
}

// Code 返回code
func (r BootError) Code() int {
	return r.code
}
