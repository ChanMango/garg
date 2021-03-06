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

var defaultStructDescribe = make(map[string]*structDescribe)

type structDescribe struct {
	fieldTypeMap        map[string]reflect.Type
	fieldNameMap        map[string]string
	fieldTagMap         map[string]reflect.StructTag
	fieldRuleExpressMap map[string]rule.Express
	needVerifyTagField  []string
}

//只能使用此方法进行初始化
func NewStructDescribe() *structDescribe {
	var fieldTypeMap = make(map[string]reflect.Type, 0)
	var fieldNameMap = make(map[string]string, 0)
	var expressMap = make(map[string]rule.Express, 0)
	var fieldTagMap = make(map[string]reflect.StructTag, 0)
	var needVerifyTag = make([]string, 0)
	return &structDescribe{fieldTypeMap: fieldTypeMap, fieldNameMap: fieldNameMap, fieldRuleExpressMap: expressMap, fieldTagMap: fieldTagMap, needVerifyTagField: needVerifyTag}
}

type defaultResolver struct {
	StructName string
	TargetV    reflect.Value
	ErrResult  Result
}
type mapResolver struct {
	TargetV   reflect.Value
	ErrResult Result
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
func NewDefaultResolver(val interface{}) (*defaultResolver, Result) {
	structName := genStrutNameKey(val)
	tp, tv := check.TypeAndValue(val)
	//是否已经存在结构信息，不存在才解析
	_, ok := defaultStructDescribe[structName]
	if !ok {
		describe := NewStructDescribe()
		r := NewResult()
		for i := 0; i < tp.NumField(); i++ {
			fieldName := tp.Field(i).Name
			describe.fieldTypeMap[fieldName] = tp.Field(i).Type
			describe.fieldNameMap[fieldName] = tp.Field(i).Name
			describe.fieldTagMap[fieldName] = tp.Field(i).Tag
			tagV := tp.Field(i).Tag.Get(common.VERIFY_LABEL)
			if tagV != "" {
				//存在需要处里arg tag
				describe.needVerifyTagField = append(describe.needVerifyTagField, fieldName)
				express, err := rule.NewParser(tagV).Parse(tp.Field(i).Type)
				if err != nil {
					r.Add(structName, fieldName, err)
					continue
				}
				describe.fieldRuleExpressMap[fieldName] = express
			}
		}
		if r.CollectToError() != nil {
			return nil, r
		}
		defaultStructDescribe[structName] = describe
	}
	return &defaultResolver{StructName: structName, TargetV: tv, ErrResult: NewResult()}, nil
}

func (resolver defaultResolver) verifyByTagRule() (bool, Result) {
	pass := true
	//一定存在
	describe, ok := defaultStructDescribe[resolver.StructName]
	if !ok {
		resolver.ErrResult.Add(resolver.StructName, "internal", errors.New("can not get have describe for sturct"))
		return false, resolver.ErrResult
	}
	for _, field := range describe.needVerifyTagField {
		tags := describe.fieldTagMap[field].Get(common.VERIFY_LABEL)
		fieldType := describe.fieldTypeMap[field]
		express, err := rule.NewParser(tags).Parse(fieldType)
		if err != nil {
			resolver.ErrResult.Add(resolver.StructName, field, err)
			pass = false
			continue
		}
		_, err = express.Cal(resolver.TargetV.FieldByName(field).Interface())
		if err != nil {
			pass = false
			resolver.ErrResult.Add(resolver.StructName, field, err)
		}
	}
	return pass, resolver.ErrResult
}
func (resolver defaultResolver) verifyByMapRule(rulMap map[string]string) (bool, error) {
	pass := true
	//一定存在
	describe, ok := defaultStructDescribe[resolver.StructName]
	if !ok {
		resolver.ErrResult.Add(resolver.StructName, "internal", errors.New("can not get have describe for sturct"))
		return false, resolver.ErrResult.CollectToError()
	}
	for field, rl := range rulMap {
		fieldType, ok := describe.fieldTypeMap[field]
		if !ok {
			resolver.ErrResult.Add(resolver.StructName, field, errors.New("no such field "))
			pass = false
			continue
		}
		express, err := rule.NewParser(rl).Parse(fieldType)
		if err != nil {
			resolver.ErrResult.Add(resolver.StructName, field, err)
			pass = false
			continue
		}
		_, err = express.Cal(resolver.TargetV.FieldByName(field).Interface())
		if err != nil {
			pass = false
			resolver.ErrResult.Add(resolver.StructName, field, err)
		}
	}

	return pass, resolver.ErrResult.CollectToError()
}
