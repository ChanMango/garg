package garg

import (
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

type DefaultResolver struct {
	StructName string
	Value  reflect.Value
}

//只能使用此方法进行初始化

func NewDefaultResolver(target interface{}) *DefaultResolver {
	tp := reflect.TypeOf(target)
	tv := reflect.ValueOf(target)
	if tp.Kind() != reflect.Struct {
		//之前已经确认是是struct或者struct pointer类型, 所以这里一定是pointer类型
		tp = reflect.TypeOf(target).Elem()
		tv = reflect.ValueOf(target).Elem()
	}
	//是否已经存在结构信息，不存在才解析
	stuctName := tp.Name()
	_, ok := defaultStructDescribe[stuctName]
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
		defaultStructDescribe[tp.Name()] = describe
	}
	return &DefaultResolver{StructName: stuctName, Value: tv}
}

func (rvr DefaultResolver) parseStructure(tp reflect.Type) {

}

func (rvr DefaultResolver) verify() (pass bool, rt Result) {
	rt = NewResult()
	//一定存在
	describe := defaultStructDescribe[rvr.StructName]
	for _, field := range describe.needVerifyTagField {
		tags := describe.fieldTagMap[field].Get(common.VERIFY_LABEL)
		express, err := rule.Parse(tags)
		if err != nil {
			rt.Add(field, err)
			pass = false
			continue
		}
		pass, err = express.Cal(rvr.Value.FieldByName(field).Interface())
		if err != nil {
			rt.Add(field, err)
		}
	}
	return pass, rt
}
