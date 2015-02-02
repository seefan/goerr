//一个简单的生成错误的帮助类
package goerr

import (
	"fmt"
)

var (
	FormatString = "%s\n\n%s"
)

//按格式返回一个错误
func NewError(err error, format string, p ...interface{}) error {
	return fmt.Errorf(FormatString, err, fmt.Sprintf(format, p...))
}
