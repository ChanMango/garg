package garg

import (
	"fmt"
	"git.xiaojukeji.com/chenyeung/garg/check"
	"reflect"
	"testing"
	"time"
)

func Test(t *testing.T) {
	student := NewStudent("yung", 29, true)
	//newStu := NewStudent("kang", 10, true)
	m := make(map[string]string)
	m["name"] = "textName"
	m["age"] = "30"
	CompareAndUpdate(m, student)
	fmt.Println("old student", student)
	//check, result := CheckByTag(student)
	//fmt.Println(check, result)
}

type Student struct {
	ID       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty" arg:">20 & <= 29"`
	AtScholl bool   `json:"at_scholl,omitempty" arg:"=true"`
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

func TestChangeSliceValue(t *testing.T) {
	var list []Student
	list = append(list, *NewStudent("yeung", 20, false))

	toChange(list)
	for i := range list {
		fmt.Println(list[i])
	}
}
func toChange(tmp []Student) {
	for i := range tmp {
		tmp[i].Age = 100
		tmp[i].Name = "test"
	}
}

func TestIsPointer(t *testing.T) {
	x := NewStudent("yeung", 20, false)
	err := check.IsPtrVal(x)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCheckRule(t *testing.T) {
	var data = struct {
		Id          uint64    `json:"id" db:"id"`
		TaskId      uint64    `json:"task_id" db:"task_id"`
		CityId      uint64    `json:"city_id" db:"city_id"`
		Phone       string    `json:"phone" db:"phone"`
		CountryCode string    `json:"country_code" db:"country_code"`
		QueryWord   string    `json:"query_word" db:"query_word"`
		Lat         float64   `json:"lat" db:"lat" arg:">=-90 and <=90"`
		Lng         float64   `json:"lng" db:"lng" arg:">=-180 and <=180"`
		CreateTime  time.Time `json:"create_time" db:"create_time"`
		Comment     string    `json:"comment" db:"comment"`
	}{}

	data.Lat = 113.4
	data.Lng = 32.33444
	data.Phone = "12344345"
	data.CountryCode = "MX"
	data.QueryWord = "kfc"
	start := time.Now().Nanosecond()
	pass, result := CheckByTag(&data)
	end := time.Now().Nanosecond()
	fmt.Println(pass, result)
	fmt.Printf("耗时=%vms \n", float32(end-start)/float32(time.Millisecond))
}
