package check

import (
	"errors"
	"git.xiaojukeji.com/chenyeung/garg/common"
	"reflect"
	"regexp"
)

var default_type_compile, _ = regexp.Compile("[8|16|32|64]+|(ptr)?")

func GT(expect, obtain interface{}) {

}
func LT(expect, obtain interface{}) {

}
func GE(expect, obtain interface{}) {

}
func LE(expect, obtain interface{}) {

}
func Required(field string, expect interface{}, result map[string]string) {

}
func IsStruct(target interface{}) (bool, error) {
	tp := reflect.TypeOf(target)
	if tp.Kind() == reflect.Struct || tp.Elem().Kind() == reflect.Struct {
		return true, nil
	}
	return false, common.NotStructTypeError
}
func IsSameType(sTp, tTp reflect.Type) (bool, error) {
	tpa := sTp.Kind().String()
	tpb := tTp.Kind().String()
	if default_type_compile.ReplaceAllString(tpa, "") != default_type_compile.ReplaceAllString(tpb, "") {
		return false, errors.New("one's type =" + tpa + ", the another =" + tpb)
	}
	return true, nil
}