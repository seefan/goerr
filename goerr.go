//一个简单的生成错误的帮助类
package goerr

import (
	"fmt"
)

var (
	FormatString = "%v\n\n%s"
)

//按格式返回一个错误
//同时携带原始的错误信息
func NewError(err error, format string, p ...interface{}) error {
	return fmt.Errorf(FormatString, err, fmt.Sprintf(format, p...))
}

//返回一个错误
func New(format string, p ...interface{}) error {
	return fmt.Errorf(format, p...)
}
