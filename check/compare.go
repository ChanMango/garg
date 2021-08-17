package check

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type CompareType int

const (
	compareLess CompareType = iota - 1
	compareEqual
	compareGreater
	compareError CompareType = 100
)

var (
	intType   = reflect.TypeOf(int(1))
	int8Type  = reflect.TypeOf(int8(1))
	int16Type = reflect.TypeOf(int16(1))
	int32Type = reflect.TypeOf(int32(1))
	int64Type = reflect.TypeOf(int64(1))

	uintType   = reflect.TypeOf(uint(1))
	uint8Type  = reflect.TypeOf(uint8(1))
	uint16Type = reflect.TypeOf(uint16(1))
	uint32Type = reflect.TypeOf(uint32(1))
	uint64Type = reflect.TypeOf(uint64(1))

	float32Type = reflect.TypeOf(float32(1))
	float64Type = reflect.TypeOf(float64(1))
	boolType    = reflect.TypeOf(false)

	stringType = reflect.TypeOf("")
)

//比较两个同类型数据的成员值
func CompareStruct(sa, sb interface{}) bool {
	var diffFileNum = 0
	ta, va := TypeAndValue(sa)
	_, vb := TypeAndValue(sb)
	if IsStruct(sa) || IsStruct(sb) {
		return false
	}
	isSameType, _ := IsSameType(sa, sb)
	if !isSameType {
		return false
	}

	for i := 0; i < va.NumField(); i++ {
		fa := va.Field(i)
		curFieldName := ta.Field(i).Name
		fb := vb.FieldByName(curFieldName)
		switch fa.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			va := fa.Int()
			vb := fb.Int()
			compare(va, vb)
		case reflect.String:
			va := fa.String()
			vb := fb.String()
			compare(va, vb)
		case reflect.Float64:
			va := fa.Float()
			vb := fb.Float()
			compare(va, vb)
		case reflect.Bool:
			va := fa.Bool()
			vb := fb.Bool()
			compare(va, vb)
		case reflect.Uint64, reflect.Uint8, reflect.Uint32, reflect.Uint:
			va := fa.Uint()
			vb := fb.Uint()
			compare(va, vb)
		default:
			fmt.Println("UnSupport Kind=", fa.Kind().String())
		}

	}
	if diffFileNum == 0 {
		fmt.Println("无Diff字段")
	}

	return false
}
func compare(obj1, obj2 interface{}) (CompareType, bool) {

	obj1Value := reflect.ValueOf(obj1)
	obj2Value := reflect.ValueOf(obj2)

	// throughout this switch we try and avoid calling .Convert() if possible,
	// as this has a pretty big performance impact
	switch obj1Value.Kind() {
	case reflect.Int:
		{
			intobj1, ok := obj1.(int)
			if !ok {
				intobj1 = obj1Value.Convert(intType).Interface().(int)
			}
			intobj2, ok := obj2.(int)
			if !ok {
				intobj2 = obj2Value.Convert(intType).Interface().(int)
			}
			if intobj1 > intobj2 {
				return compareGreater, true
			}
			if intobj1 == intobj2 {
				return compareEqual, true
			}
			if intobj1 < intobj2 {
				return compareLess, true
			}
		}
	case reflect.Int8:
		{
			int8obj1, ok := obj1.(int8)
			if !ok {
				int8obj1 = obj1Value.Convert(int8Type).Interface().(int8)
			}
			int8obj2, ok := obj2.(int8)
			if !ok {
				int8obj2 = obj2Value.Convert(int8Type).Interface().(int8)
			}
			if int8obj1 > int8obj2 {
				return compareGreater, true
			}
			if int8obj1 == int8obj2 {
				return compareEqual, true
			}
			if int8obj1 < int8obj2 {
				return compareLess, true
			}
		}
	case reflect.Int16:
		{
			int16obj1, ok := obj1.(int16)
			if !ok {
				int16obj1 = obj1Value.Convert(int16Type).Interface().(int16)
			}
			int16obj2, ok := obj2.(int16)
			if !ok {
				int16obj2 = obj2Value.Convert(int16Type).Interface().(int16)
			}
			if int16obj1 > int16obj2 {
				return compareGreater, true
			}
			if int16obj1 == int16obj2 {
				return compareEqual, true
			}
			if int16obj1 < int16obj2 {
				return compareLess, true
			}
		}
	case reflect.Int32:
		{
			int32obj1, ok := obj1.(int32)
			if !ok {
				int32obj1 = obj1Value.Convert(int32Type).Interface().(int32)
			}
			int32obj2, ok := obj2.(int32)
			if !ok {
				int32obj2 = obj2Value.Convert(int32Type).Interface().(int32)
			}
			if int32obj1 > int32obj2 {
				return compareGreater, true
			}
			if int32obj1 == int32obj2 {
				return compareEqual, true
			}
			if int32obj1 < int32obj2 {
				return compareLess, true
			}
		}
	case reflect.Int64:
		{
			int64obj1, ok := obj1.(int64)
			if !ok {
				int64obj1 = obj1Value.Convert(int64Type).Interface().(int64)
			}
			int64obj2, ok := obj2.(int64)
			if !ok {
				int64obj2 = obj2Value.Convert(int64Type).Interface().(int64)
			}
			if int64obj1 > int64obj2 {
				return compareGreater, true
			}
			if int64obj1 == int64obj2 {
				return compareEqual, true
			}
			if int64obj1 < int64obj2 {
				return compareLess, true
			}
		}
	case reflect.Uint:
		{
			uintobj1, ok := obj1.(uint)
			if !ok {
				uintobj1 = obj1Value.Convert(uintType).Interface().(uint)
			}
			uintobj2, ok := obj2.(uint)
			if !ok {
				uintobj2 = obj2Value.Convert(uintType).Interface().(uint)
			}
			if uintobj1 > uintobj2 {
				return compareGreater, true
			}
			if uintobj1 == uintobj2 {
				return compareEqual, true
			}
			if uintobj1 < uintobj2 {
				return compareLess, true
			}
		}
	case reflect.Uint8:
		{
			uint8obj1, ok := obj1.(uint8)
			if !ok {
				uint8obj1 = obj1Value.Convert(uint8Type).Interface().(uint8)
			}
			uint8obj2, ok := obj2.(uint8)
			if !ok {
				uint8obj2 = obj2Value.Convert(uint8Type).Interface().(uint8)
			}
			if uint8obj1 > uint8obj2 {
				return compareGreater, true
			}
			if uint8obj1 == uint8obj2 {
				return compareEqual, true
			}
			if uint8obj1 < uint8obj2 {
				return compareLess, true
			}
		}
	case reflect.Uint16:
		{
			uint16obj1, ok := obj1.(uint16)
			if !ok {
				uint16obj1 = obj1Value.Convert(uint16Type).Interface().(uint16)
			}
			uint16obj2, ok := obj2.(uint16)
			if !ok {
				uint16obj2 = obj2Value.Convert(uint16Type).Interface().(uint16)
			}
			if uint16obj1 > uint16obj2 {
				return compareGreater, true
			}
			if uint16obj1 == uint16obj2 {
				return compareEqual, true
			}
			if uint16obj1 < uint16obj2 {
				return compareLess, true
			}
		}
	case reflect.Uint32:
		{
			uint32obj1, ok := obj1.(uint32)
			if !ok {
				uint32obj1 = obj1Value.Convert(uint32Type).Interface().(uint32)
			}
			uint32obj2, ok := obj2.(uint32)
			if !ok {
				uint32obj2 = obj2Value.Convert(uint32Type).Interface().(uint32)
			}
			if uint32obj1 > uint32obj2 {
				return compareGreater, true
			}
			if uint32obj1 == uint32obj2 {
				return compareEqual, true
			}
			if uint32obj1 < uint32obj2 {
				return compareLess, true
			}
		}
	case reflect.Uint64:
		{
			uint64obj1, ok := obj1.(uint64)
			if !ok {
				uint64obj1 = obj1Value.Convert(uint64Type).Interface().(uint64)
			}
			uint64obj2, ok := obj2.(uint64)
			if !ok {
				uint64obj2 = obj2Value.Convert(uint64Type).Interface().(uint64)
			}
			if uint64obj1 > uint64obj2 {
				return compareGreater, true
			}
			if uint64obj1 == uint64obj2 {
				return compareEqual, true
			}
			if uint64obj1 < uint64obj2 {
				return compareLess, true
			}
		}
	case reflect.Float32:
		{
			float32obj1, ok := obj1.(float32)
			if !ok {
				float32obj1 = obj1Value.Convert(float32Type).Interface().(float32)
			}
			float32obj2, ok := obj2.(float32)
			if !ok {
				float32obj2 = obj2Value.Convert(float32Type).Interface().(float32)
			}
			if float32obj1 > float32obj2 {
				return compareGreater, true
			}
			if float32obj1 == float32obj2 {
				return compareEqual, true
			}
			if float32obj1 < float32obj2 {
				return compareLess, true
			}
		}
	case reflect.Float64:
		{
			float64obj1, ok := obj1.(float64)
			if !ok {
				float64obj1 = obj1Value.Convert(float64Type).Interface().(float64)
			}
			float64obj2, ok := obj2.(float64)
			if !ok {
				float64obj2 = obj2Value.Convert(float64Type).Interface().(float64)
			}
			if float64obj1 > float64obj2 {
				return compareGreater, true
			}
			if float64obj1 == float64obj2 {
				return compareEqual, true
			}
			if float64obj1 < float64obj2 {
				return compareLess, true
			}
		}
	case reflect.String:
		{
			stringobj1, ok := obj1.(string)
			if !ok {
				stringobj1 = obj1Value.Convert(stringType).Interface().(string)
			}
			stringobj2, ok := obj2.(string)
			if !ok {
				stringobj2 = obj2Value.Convert(stringType).Interface().(string)
			}
			if stringobj1 > stringobj2 {
				return compareGreater, true
			}
			if stringobj1 == stringobj2 {
				return compareEqual, true
			}
			if stringobj1 < stringobj2 {
				return compareLess, true
			}
		}
	case reflect.Bool:
		boolobj1, ok := obj1.(bool)
		if !ok {
			boolobj1 = obj1Value.Convert(boolType).Interface().(bool)
		}
		boolobj2, ok := obj2.(bool)
		if !ok {
			boolobj2 = obj2Value.Convert(boolType).Interface().(bool)
		}
		if boolobj1 == boolobj2 {
			return compareEqual, true
		}
	case reflect.Ptr:
		if obj1 == nil || obj2 == nil {
			//有一个值是空指针，不能比
			return compareError, false
		} else {
			return compare(obj1Value.Elem().Interface(), obj2Value.Elem().Interface())
		}
	}
	return compareEqual, false
}

// Positive asserts that the specified element is positive
//
//    assert.Positive( 1)
//    assert.Positive( 1.23)
//func Positive(e interface{}, msgAndArgs ...interface{}) bool {
//	zero := reflect.Zero(reflect.TypeOf(e))
//	return compareTwoValues(e, zero.Interface(), []CompareType{compareGreater})
//}

// Negative asserts that the specified element is negative
//
//    assert.Negative( -1)
//    assert.Negative( -1.23)
//func Negative(e interface{}, msgAndArgs ...interface{}) bool {
//	zero := reflect.Zero(reflect.TypeOf(e))
//	return compareTwoValues(e,t,zero.Interface(), []CompareType{compareLess})
//}

func compareTwoValues(e1 interface{}, e2 interface{}, allowedComparesResults []CompareType) (bool, error) {
	e1Type, _ := TypeAndKind(e1)
	e2Type, _ := TypeAndKind(e2)
	if e1Type != e2Type {
		return false, errors.New("Elements should be the same type")
	}

	compareResul, isComparable := compare(e1, e2)
	if !isComparable {
		return false, errors.New(fmt.Sprintf("Can not compare type %s", reflect.TypeOf(e1)))
	}

	if !containsValue(allowedComparesResults, compareResul) {
		return false, errors.New("")
	}

	return true, nil
}

func containsValue(values []CompareType, value CompareType) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func includeElement(element interface{}, list interface{}) (found bool, err error) {

	listValue := reflect.ValueOf(list)
	listKind := reflect.TypeOf(list).Kind()
	defer func() {
		if e := recover(); e != nil {
			found = false
			err = e.(error)
		}
	}()

	if listKind == reflect.String {
		elementValue := reflect.ValueOf(element)
		contains := strings.Contains(listValue.String(), elementValue.String())
		return contains, err
	}

	if listKind == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if ObjectsAreEqual(mapKeys[i].Interface(), element) {
				return true, err
			}
		}
		return false, err
	}

	for i := 0; i < listValue.Len(); i++ {
		if ObjectsAreEqual(listValue.Index(i).Interface(), element) {
			return true, err
		}
	}
	return false, err

}

// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}
