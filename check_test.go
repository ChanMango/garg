package garg

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	student := NewStudent("yung", 29, false)
	Check(student)
}

type Student struct {
	ID       int64  `json:"id"`
	Name     string `json:"name,omitempty" arg:"required"`
	Age      int    `json:"age,omitempty" arg:">20"`
	AtScholl bool   `json:"at_scholl,omitempty" arg:"=ture"`
}

func NewStudent(name string, age int, atScholl bool) *Student {
	return &Student{Name: name, Age: age, AtScholl: atScholl}
}

func TestRegex(t *testing.T) {
	//ruleStr := "(>10 & <= 1 )or =5"
	//and_mh, _ := regexp.MatchString("and|&", ruleStr)
	//or_mh, _ := regexp.Compile("or|\\|")
	//
	//split := compile.Split(ruleStr, -1)
	//fmt.Println(split, len(split))
	//var str =int64(200)
	//bstr := "Uintptr"
	//fmt.Println(check.IsSameType(str, bstr))

}

func TestTmp(t *testing.T) {
	a := 10
	b := int64(20)
	at := reflect.TypeOf(a)
	println(at.Kind(), at.Name(), at.String())
	bt := reflect.TypeOf(b)
	bv := reflect.ValueOf(b)
	println(bt.Kind(), bt.Name(), bt.String(), bv.Type().String())

}
