package garg

import (
	"errors"
	"fmt"
	"git.xiaojukeji.com/chenyeung/garg/check"
	"reflect"
	"strconv"
)

type CheckerFunc = func(val interface{}) (bool, error)

type CheckIfc interface {
	Check(cf CheckerFunc) (bool, Result)
}

//根据stuct tag上定义的规则，校验字段内容
func CheckByTag(vals ...interface{}) (bool, error) {
	r := NewResult()
	for i := range vals {
		//执行参数检查
		isStruct := check.IsStruct(vals[i])
		if !isStruct {
			r.Add("nil", "arg_type_err", errors.New(fmt.Sprintf("type of arg index %v shou be struct", i)))
			continue
		}
		pass, result := NewDefaultResolver(vals[i]).verify()
		if !pass {
			r.AddAll(result)
		}
	}
	if len(r) > 0 {
		return false, r.CollectToError()
	}
	return true, nil

}

//根据map里边制定字段对应的rule， 进行stuct 字段的参数校验
func CheckByMap(val interface{}, ruleMap map[string]interface{}) (bool, Result) {

	return false, nil

}

//使用自定义校验规则，校验val内容
func CustomChecker(value interface{}, ckFun CheckerFunc) (bool, error) {
	return ckFun(value)
}

//比较内容并进行更改
func CompareAndUpdate(patchMap map[string]string, toUpdate interface{}) error {
	vpp := reflect.ValueOf(patchMap)
	vpo := reflect.ValueOf(toUpdate)
	if vpp.Kind() == reflect.Ptr {
		vpp = vpp.Elem()
	}
	if vpp.Kind() != reflect.Map {
		return errors.New("数据更新patch 仅支持map 或mapPtr")
	} else {

	}
	if vpo.Kind() != reflect.Ptr {
		return errors.New("仅支持引用型数据类型")
	}
	//空指针判断
	if vpo.IsNil() {
		return errors.New("<nil pointer>")
	}
	if !vpo.IsValid() {
		return errors.New("<zero targetVP>")
	}
	////数据更新处理
	for field, patchVal := range patchMap {
		fieldName := NewStructTool(toUpdate).Tag2FieldName(field)
		valueOf := vpo.Elem().FieldByName(fieldName)
		if !valueOf.IsValid() {
			fmt.Println(field, " | not have this fied")
			continue
		}
		if valueOf.Kind() != reflect.Ptr {
			//饮用类型才可以 修改值
			valueOf = valueOf.Addr()
		}
		err := setValue(valueOf.Interface(), patchVal)
		if err != nil {
			//有错误 更新失败
			return err
		}
	}
	return nil
}
func setValue(target interface{}, val interface{}) error {
	of := reflect.ValueOf(target)
	if of.Kind() != reflect.Ptr {
		errors.New("仅支持值类型数据,否则 更新数据时会报错")
	}
	switch of.Elem().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ov, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			errors.New(fmt.Sprintf("类型不匹配,targetType=%v ,updateVal=%v", of.Kind().String(), val))
		} else {
			of.Elem().SetInt(ov)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ov, err := strconv.ParseUint(val.(string), 10, 64)
		if err != nil {
			errors.New(fmt.Sprintf("类型不匹配,targetType=%v ,updateVal=%v", of.Kind().String(), val))
		} else {
			of.Elem().SetUint(ov)
		}
	case reflect.Float32, reflect.Float64:
		ov, err := strconv.ParseFloat(val.(string), 64)
		if err != nil {
			errors.New(fmt.Sprintf("类型不匹配,targetType=%v ,updateVal=%v", of.Kind().String(), val))
		} else {
			of.Elem().SetFloat(ov)
		}
	case reflect.String:
		ov, ok := val.(string)
		if !ok {
			errors.New(fmt.Sprintf("类型不匹配,targetType=%v ,updateVal=%v", of.Kind().String(), val))
		} else {
			of.Elem().SetString(ov)
		}
	case reflect.Bool:
		ov, err := strconv.ParseBool(val.(string))
		if err != nil {
			errors.New(fmt.Sprintf("类型不匹配,targetType=%v ,updateVal=%v", of.Kind().String(), val))
		} else {
			of.Elem().SetBool(ov)
		}
		//todo 对这些改值进行判断
	/*case reflect.Ptr:
		//todo
		CompareAndUpdate(vpp.Interface(), vpo.Interface())
	case reflect.Array, reflect.Slice:
		//todo 数据复制有些问题
		reflect.AppendSlice(vpo, vpp)
	case reflect.Map:
		if vpo.IsNil() {
			value := reflect.MakeMap(vpo.Type())
			for _, key := range vpo.MapKeys() {
				value.SetMapIndex(key, vpo.MapIndex(key))
			}
		}*/
	//不为空的情况不知道呢
	default:
		errors.New("暂不支持处理该类型" + of.Kind().String())
	}
	return nil
}
