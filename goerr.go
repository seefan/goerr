//一个简单的生成错误的帮助类
package goerr

import (
	"bytes"
	"fmt"
	"strconv"
)

// new errorContext with error and string
//
// return error
func Errorf(err error, format string, p ...interface{}) error {
	return &errorContext{text: fmt.Sprintf(format, p...), err: err, code: -1}
}

// new errorContext with error
//
//  return  error
func Error(err error) *errorContext {
	return &errorContext{err: err, code: -1}
}

// new errorContext with string
func String(format string, p ...interface{}) *errorContext {
	return &errorContext{text: fmt.Sprintf(format, p...), code: -1}
}

//new error
type errorContext struct {
	code int
	text string
	file string
	line int
	err  error
}

func (e *errorContext) Format(format string, p ...interface{}) *errorContext {
	e.text = fmt.Sprintf(format, p...)
	return e
}
func (e *errorContext) Code(code int) *errorContext {
	e.code = code
	return e
}
func (e *errorContext) Line(line int) *errorContext {
	e.line = line
	return e
}
func (e *errorContext) File(file string) *errorContext {
	e.file = file
	return e
}

//实现error接口
func (e *errorContext) Error() string {
	var buffer bytes.Buffer
	if e.file != "" {
		buffer.WriteString("File: ")
		buffer.WriteString(e.file)
		buffer.WriteString("\t")
	}
	if e.line != -1 {
		bs := strconv.AppendInt(nil, int64(e.line), 10)
		buffer.WriteString("Line: ")
		buffer.Write(bs)
		buffer.WriteString("\t")
	}
	if e.code != 0 {
		bs := strconv.AppendInt(nil, int64(e.code), 10)
		buffer.WriteString("Error Code: ")
		buffer.Write(bs)
		buffer.WriteString("\t")
	}
	if e.text != "" {
		buffer.WriteString("Text: ")
		buffer.WriteString(e.text)
		buffer.WriteString("\t")
	}
	if e.err != nil {
		buffer.WriteString("\nTrace: ")
		buffer.WriteString(e.err.Error())
	}
	return buffer.String()
}
