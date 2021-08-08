package garg

import (
	"reflect"
	"strings"
)

var defaultStructTagMap = make(map[string]map[string]string)

var TagType_JSON = "json"

type StructTool struct {
	Name string
}

func NewStructTool(val interface{}) *StructTool {
	//of := reflect.ValueOf(val)
	tp := reflect.TypeOf(val)
	if tp.Kind() == reflect.Ptr {
		tp = tp.Elem()
	}
	m, ok := defaultStructTagMap[tp.String()]
	if !ok {
		//初次
		m = make(map[string]string)
	}
	//记录jie构
	for i := 0; i < tp.NumField(); i++ {
		tag := strings.Split(tp.Field(i).Tag.Get("json"), ",")[0]
		m[tag] = tp.Field(i).Name
	}
	defaultStructTagMap[tp.String()] = m
	return &StructTool{Name: tp.String()}
}
func (st *StructTool) Tag2FieldName(tag string) string {
	m, ok := defaultStructTagMap[st.Name]
	if !ok {
		return ""
	}
	return m[tag]
}
