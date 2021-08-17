package check

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

var default_type_compile, _ = regexp.Compile("[8|16|32|64]+|(ptr)?")

func TypeAndKind(v interface{}) (reflect.Type, reflect.Kind) {
	t := reflect.TypeOf(v)
	k := t.Kind()

	if k == reflect.Ptr {
		t = t.Elem()
		k = t.Kind()
	}
	return t, k
}
func TypeAndValue(v interface{}) (reflect.Type, reflect.Value) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		val = val.Elem()
	}
	return t, val
}

func toSetType(val string, atp reflect.Type) (interface{}, error) {
	var err error
	var ev interface{}
	switch atp.Kind() {
	case reflect.Int:
		x, errx := strconv.ParseInt(val, 10, 64)
		ev = int(x)
		err = errx
	case reflect.Int8:
		x, errx := strconv.ParseInt(val, 10, 8)
		err = errx
		ev = int8(x)
	case reflect.Int16:
		x, errx := strconv.ParseInt(val, 10, 16)
		err = errx
		ev = int16(x)
	case reflect.Int32:
		x, errx := strconv.ParseInt(val, 10, 32)
		err = errx
		ev = int32(x)
	case reflect.Int64:
		x, errx := strconv.ParseInt(val, 10, 64)
		err = errx
		ev = int64(x)
	case reflect.Uint:
		x, errx := strconv.ParseUint(val, 10, 64)
		err = errx
		ev = uint64(x)
	case reflect.Uint8:
		x, errx := strconv.ParseUint(val, 10, 8)
		err = errx
		ev = uint8(x)
	case reflect.Uint16:
		x, errx := strconv.ParseUint(val, 10, 16)
		err = errx
		ev = uint16(x)
	case reflect.Uint32:
		x, errx := strconv.ParseUint(val, 10, 32)
		err = errx
		ev = uint32(x)
	case reflect.Uint64:
		x, errx := strconv.ParseUint(val, 10, 64)
		err = errx
		ev = x
	case reflect.Uintptr:
		x, errx := strconv.ParseUint(val, 10, 64)
		err = errx
		f := float64(x)
		ev = &f
	case reflect.Float32:
		x, errx := strconv.ParseFloat(val, 32)
		err = errx
		ev = float32(x)
	case reflect.Float64:
		x, errx := strconv.ParseFloat(val, 64)
		err = errx
		ev = float64(x)
	case reflect.String:
		ev = val
	default:
		err = errors.New("暂不支持该类型校验，valueType=" + atp.Kind().String())
	}
	return ev, err
}

//通过pass去判断是否 校验通过 ==
func GE(obtain, expect interface{}) (bool, error) {
	atp, _ := TypeAndValue(obtain)
	isStr := IsString(expect)
	newExp := expect
	if isStr {
		setType, err := toSetType(expect.(string), atp)
		if err != nil {
			return false, err
		} else {
			newExp = setType
		}
	}
	pass, err := compareTwoValues(obtain, newExp, []CompareType{compareEqual, compareGreater})
	return pass, err
}
func LE(obtain, expect interface{}) (bool, error) {
	atp, _ := TypeAndValue(obtain)
	isStr := IsString(expect)
	newExp := expect
	if isStr {
		setType, err := toSetType(expect.(string), atp)
		if err != nil {
			return false, err
		} else {
			newExp = setType
		}

	}
	pass, err := compareTwoValues(obtain, newExp, []CompareType{compareEqual, compareLess})
	return pass, err
}
func GT(obtain, expect interface{}) (bool, error) {
	atp, _ := TypeAndValue(obtain)
	isStr := IsString(expect)
	newExp := expect
	if isStr {
		setType, err := toSetType(expect.(string), atp)
		if err != nil {
			return false, err
		} else {
			newExp = setType
		}

	}
	pass, err := compareTwoValues(obtain, newExp, []CompareType{compareGreater})
	return pass, err
}
func LT(obtain, expect interface{}) (bool, error) {
	atp, _ := TypeAndValue(obtain)
	isStr := IsString(expect)
	newExp := expect
	if isStr {
		setType, err := toSetType(expect.(string), atp)
		if err != nil {
			return false, err
		} else {
			newExp = setType
		}

	}
	pass, err := compareTwoValues(obtain, newExp, []CompareType{compareLess})
	return pass, err
}
func E(obtain, expect interface{}) (bool, error) {
	atp, _ := TypeAndValue(obtain)
	isStr := IsString(expect)
	newExp := expect
	if isStr {
		setType, err := toSetType(expect.(string), atp)
		if err != nil {
			return false, err
		} else {
			newExp = setType
		}

	}
	pass, err := compareTwoValues(obtain, newExp, []CompareType{compareEqual})
	return pass, err
}

//不想等
func NE(obtain, expect interface{}) (bool, error) {
	atp, _ := TypeAndValue(obtain)
	isStr := IsString(expect)
	newExp := expect
	if isStr {
		setType, err := toSetType(expect.(string), atp)
		if err != nil {
			return false, err
		} else {
			newExp = setType
		}

	}
	pass, err := compareTwoValues(obtain, newExp, []CompareType{compareEqual})
	return !pass, err
}

//是否是初始化值，或者默认零值
func Required(obtain interface{}) (bool, error) {
	var err error
	defer func() {
		x := recover()
		if x != nil {
			err = x.(error)
		}
	}()
	_, val := TypeAndValue(obtain)
	isZero := val.IsZero()
	return !isZero, err
}

func IsStruct(target interface{}) bool {
	tp := reflect.TypeOf(target)
	if tp.Kind() == reflect.Struct || tp.Elem().Kind() == reflect.Struct {
		return true
	}
	return false
}
func Contains(obtain interface{}, coll interface{}) (bool, error) {

	found, err := includeElement(obtain, coll)
	if !found && err == nil {
		err = errors.New(fmt.Sprintf("%v not contains %v", coll, obtain))
	}
	return found, err
}
func NotContains(obtain interface{}, coll interface{}) (bool, error) {
	found, err := includeElement(obtain, coll)
	if found && err == nil {
		err = errors.New(fmt.Sprintf("%v should not contains %v", coll, obtain))
	}
	return !found, err
}

//有错误，就代表不是穿的指针引用的值
func IsPtrVal(target interface{}) bool {
	tp := reflect.TypeOf(target)
	if tp.Kind() != reflect.Ptr {
		return false
	}
	return true
}
func IsSameType(expected, actual interface{}) (bool, error) {
	if actual == nil {
		return false, errors.New("actual is nil")
	}
	if expected == nil {
		return false, errors.New("expected is nil")
	}
	et, _ := TypeAndKind(expected)
	at, _ := TypeAndKind(actual)

	if et.Kind() != et.Kind() {
		return false, errors.New(fmt.Sprintf(" expected %v(%v) but actual %v(%v)", et.String(), expected, at.String(), actual))
	}
	if et != at {
		return false, errors.New(fmt.Sprintf(" expected %v(%v) but actual %v(%v)", et.String(), expected, at.String(), actual))
	}

	//if ek != reflect.Struct && ek != reflect.Map && ek != reflect.Slice && ek != reflect.Array && ek != reflect.String {
	//	return false,errors.New("")
	//}

	return true, nil
}
func IsString(target interface{}) bool {
	tp := reflect.TypeOf(target)
	if tp.Kind() == reflect.Ptr {
		if tp.Elem().Kind() == reflect.String {
			return true
		}
	} else if tp.Kind() == reflect.String {
		return true
	}
	return false

}

//暂时仅支持切片类型
func CreateContainer(eles []string, of reflect.Type) ([]interface{}, error) {
	tof := of
	var isPtr = of.Kind() == reflect.Ptr
	if isPtr {
		tof = of.Elem()
	}
	res := reflect.MakeSlice(reflect.TypeOf([]interface{}{}), 0, len(eles))
	for i := range eles {
		var toSet reflect.Value
		switch tof.Kind() {
		case reflect.Int8:
			v, err := strconv.ParseInt(eles[i], 10, 8)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetInt(v)
		case reflect.Int16:
			v, err := strconv.ParseInt(eles[i], 10, 16)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetInt(v)
		case reflect.Int32:
			v, err := strconv.ParseInt(eles[i], 10, 32)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetInt(v)
		case reflect.Int64, reflect.Int:
			v, err := strconv.ParseInt(eles[i], 10, 64)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetInt(v)
		case reflect.Uint8:
			v, err := strconv.ParseUint(eles[i], 10, 8)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetUint(v)
		case reflect.Uint16:
			v, err := strconv.ParseUint(eles[i], 10, 16)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetUint(v)
		case reflect.Uint32:
			v, err := strconv.ParseUint(eles[i], 10, 32)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetUint(v)
		case reflect.Uint64, reflect.Uint:
			v, err := strconv.ParseUint(eles[i], 10, 64)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetUint(v)
		case reflect.Float64:
			v, err := strconv.ParseFloat(eles[i], 64)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetFloat(v)
		case reflect.Float32:
			v, err := strconv.ParseFloat(eles[i], 32)
			if err != nil {
				return nil, err
			}
			toSet = reflect.New(tof)
			toSet.Elem().SetFloat(v)
		case reflect.String:
			toSet = reflect.New(tof)
			toSet.Elem().SetString(eles[i])
		default:
			//不支持的类型，跳过
			continue
		}

		if isPtr {
			res = reflect.Append(res, toSet)
		} else {
			res = reflect.Append(res, toSet.Elem())
		}
	}
	return res.Interface().([]interface{}), nil
}
func SetValByType(ele string, of reflect.Type) (interface{}, error) {
	eleOf := of
	if of.Kind() == reflect.Ptr {
		eleOf = of.Elem()
	}
	newV := reflect.New(eleOf)
	switch eleOf.Kind() {
	case reflect.Int8:
		v, err := strconv.ParseInt(ele, 10, 8)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetInt(v)
	case reflect.Int16:
		v, err := strconv.ParseInt(ele, 10, 16)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetInt(v)
	case reflect.Int32:
		v, err := strconv.ParseInt(ele, 10, 32)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetInt(v)
	case reflect.Int64, reflect.Int:
		v, err := strconv.ParseInt(ele, 10, 64)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetInt(v)
	case reflect.Uint8:
		v, err := strconv.ParseUint(ele, 10, 8)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetUint(v)
	case reflect.Uint16:
		v, err := strconv.ParseUint(ele, 10, 16)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetUint(v)
	case reflect.Uint32:
		v, err := strconv.ParseUint(ele, 10, 32)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetUint(v)
	case reflect.Uint64, reflect.Uint:
		v, err := strconv.ParseUint(ele, 10, 64)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetUint(v)
	case reflect.Float64:
		v, err := strconv.ParseFloat(ele, 64)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetFloat(v)
	case reflect.Float32:
		v, err := strconv.ParseFloat(ele, 32)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetFloat(v)
	case reflect.String:
		newV.Elem().SetString(ele)
	case reflect.Bool:
		v, err := strconv.ParseBool(ele)
		if err != nil {
			return nil, err
		}
		newV.Elem().SetBool(v)
	default:
		return nil, errors.New("SetValByType: unsupport type " + of.String())
	}
	//普通类型
	if of.Kind() != reflect.Ptr {
		return newV.Elem().Interface(), nil
	}
	//指针
	return newV.Interface(), nil
}
