package garg

import (
	"crypto/md5"
	"errors"
	"fmt"
	"git.xiaojukeji.com/chenyeung/garg/check"
	"git.xiaojukeji.com/chenyeung/garg/common"
	"git.xiaojukeji.com/chenyeung/garg/rule"
	"reflect"
)

var defaultStructDescribe = make(map[string]*StructDescribe)

type StructDescribe struct {
	fieldTypeMap       map[string]reflect.Type
	fieldNameMap       map[string]string
	fieldTagMap        map[string]reflect.StructTag
	needVerifyTagField []string
}

//只能使用此方法进行初始化
func NewStructDescribe() *StructDescribe {
	var fieldTypeMap = make(map[string]reflect.Type, 0)
	var fieldNameMap = make(map[string]string, 0)
	var fieldTagMap = make(map[string]reflect.StructTag, 0)
	var needVerifyTag = make([]string, 0)
	return &StructDescribe{fieldTypeMap: fieldTypeMap, fieldNameMap: fieldNameMap, fieldTagMap: fieldTagMap, needVerifyTagField: needVerifyTag}
}

type defaultResolver struct {
	StructName string
	TargetV    reflect.Value
	ErrResult  Result
}

//只能使用此方法进行初始化

func genStrutNameKey(val interface{}) string {
	tp := reflect.TypeOf(val)
	name := tp.Name()
	if name == "" {
		name = "anonym_struct"
		md5Val := fmt.Sprintf("%x", md5.Sum([]byte(tp.String()))) //为什么要用md5  因为 匿名struct的名字特别长且有特殊符号导致map取不到对应的key
		return name + md5Val
	}
	return name
}
func NewDefaultResolver(val interface{}) *defaultResolver {
	structName := genStrutNameKey(val)
	tp := reflect.TypeOf(val).Elem()
	tp, tv := check.TypeAndValue(val)
	//是否已经存在结构信息，不存在才解析
	_, ok := defaultStructDescribe[structName]
	if !ok {
		describe := NewStructDescribe()
		for i := 0; i < tp.NumField(); i++ {
			fieldName := tp.Field(i).Name
			describe.fieldTypeMap[fieldName] = tp.Field(i).Type
			describe.fieldNameMap[fieldName] = tp.Field(i).Name
			describe.fieldTagMap[fieldName] = tp.Field(i).Tag
			tagV := tp.Field(i).Tag.Get(common.VERIFY_LABEL)
			if tagV != "" {
				//存在需要处里arg tag
				describe.needVerifyTagField = append(describe.needVerifyTagField, fieldName)
			}
		}
		defaultStructDescribe[structName] = describe
	}
	return &defaultResolver{StructName: structName, TargetV: tv, ErrResult: NewResult()}
}

func (resover defaultResolver) parseStructure(tp reflect.Type) {

}

func (resover defaultResolver) verify() (bool, Result) {
	pass := true
	//一定存在
	describe, ok := defaultStructDescribe[resover.StructName]
	if !ok {
		resover.ErrResult.Add(resover.StructName, "internal", errors.New("can not get have describe for sturct"))
		return false, resover.ErrResult
	}
	for _, field := range describe.needVerifyTagField {
		tags := describe.fieldTagMap[field].Get(common.VERIFY_LABEL)
		express, err := rule.NewParser(resover.TargetV, tags).Parse()
		if err != nil {
			resover.ErrResult.Add(resover.StructName, field, err)
			pass = false
			continue
		}
		_, err = express.Cal(resover.TargetV.FieldByName(field).Interface())
		if err != nil {
			pass = false
			resover.ErrResult.Add(resover.StructName, field, err)
		}
	}
	return pass, resover.ErrResult
}
