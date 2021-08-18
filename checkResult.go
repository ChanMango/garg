package garg

import (
	"bytes"
	"encoding/json"
	"errors"
)

//  返回参数检查错误字段的err信息
type Result map[string]map[string]string

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
func (r Result) Add(structureName, msg string, err error) {
	_, ok := r[structureName]
	if !ok {
		r[structureName] = map[string]string{}
	}
	r[structureName][msg] = err.Error()
}
func (r Result) CollectToError() error {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(r)
	return errors.New(buffer.String())
}
func (r Result) AddAll(others ...Result) {
	for i := range others {
		other := others[i]
		for s := range other {
			r[s] = other[s]
		}
	}
}
