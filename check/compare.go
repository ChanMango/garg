package check

import (
	"fmt"
	"reflect"
)

//比较两个同类型数据的成员值
func CompareStruct(DA, DB interface{}) {
	var diffFileNum = 0
	tpA := reflect.TypeOf(DA)
	tpB := reflect.TypeOf(DB)
	valA := reflect.ValueOf(DA)
	valB := reflect.ValueOf(DB)
	if tpA.Kind() == reflect.Ptr {
		if tpA.Elem().Kind() == reflect.Struct {
			tpA = tpA.Elem()
			valA = reflect.ValueOf(DA).Elem()
		} else {
			fmt.Println("仅支持struct或struct指针类型")
			return
		}
	}
	if tpB.Kind() == reflect.Ptr {
		if tpB.Elem().Kind() == reflect.Struct {
			tpB = tpB.Elem()
			valB = reflect.ValueOf(DB).Elem()
		} else {
			fmt.Println("仅支持struct或struct指针类型")
			return
		}
	}
	if tpA.Kind() != tpB.Kind() {
		fmt.Println("Diff参数A与B的类型不一致，不能进行diff")
		return
	}

	for i := 0; i < valA.NumField(); i++ {
		fa := valA.Field(i)
		curFieldName := tpA.Field(i).Name
		fb := valB.FieldByName(curFieldName)
		switch fa.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			va := fa.Int()
			vb := fb.Int()
			compare(curFieldName, va, vb, &diffFileNum)
		case reflect.String:
			va := fa.String()
			vb := fb.String()
			compare(curFieldName, va, vb, &diffFileNum)
		case reflect.Float64:
			va := fa.Float()
			vb := fb.Float()
			compare(curFieldName, va, vb, &diffFileNum)
		case reflect.Bool:
			va := fa.Bool()
			vb := fb.Bool()
			compare(curFieldName, va, vb, &diffFileNum)
		case reflect.Uint64, reflect.Uint8, reflect.Uint32, reflect.Uint:
			va := fa.Uint()
			vb := fb.Uint()
			compare(curFieldName, va, vb, &diffFileNum)
		default:
			fmt.Println("UnSupport Kind=", fa.Kind().String())
		}

	}
	if diffFileNum == 0 {
		fmt.Println("无Diff字段")
	}

}
func compare(fieldName string, va, vb interface{}, counter *int) {
	if va != vb {
		print(fmt.Sprintf("[%s]存在Diff ---> [(%v) != (%v)]\n", fieldName, va, vb))
		*counter += 1
	} else {
		//fmt.Println(fieldName+" Is OK")
	}
}
