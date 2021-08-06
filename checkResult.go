package garg

import (
	"errors"
	"fmt"
	"strings"
)

//  返回参数检查错误字段的err信息
type Result map[string]error

//type ArgError struct {
//	error []error
//}
//
//func NewArgError(errors ...error) *ArgError {
//	var errs []error
//	errs = append(errs, errors...)
//	return &ArgError{error: errs}
//}

func NewResult() Result {
	r := make(Result, 0)
	return r
}

//为field 添加err记录
func (r Result) Add(msg string, err error) {
	r[msg] = err
}
func (r Result) CollectToError() error {
	sb := strings.Builder{}
	for k, v := range r {
		sb.WriteString(fmt.Sprintf("%v=>%v\n", k, v))
	}
	return errors.New(sb.String())
}
