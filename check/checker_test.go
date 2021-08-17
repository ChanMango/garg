package check

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestGE(t *testing.T) {
	//a := 10
	//b := 10
	list := []int{}
	//pass, err := GE(&a, &b)
	//fmt.Println(pass, err)
	required, err := Required(list)
	fmt.Println(required, err)
}

func TestTmp(t *testing.T) {
	ast := require.New(t)
	ast.NotContains(10, "2010")
	of := reflect.TypeOf(int32(10))
	fmt.Println(of.String(), of.Kind(), of.Name())
}

func TestTT(t *testing.T) {
}

func TestSetValByType(t *testing.T) {
	byType, err := SetValByType("20", reflect.TypeOf(120.1))
	fmt.Println(byType, reflect.ValueOf(byType).Interface(), err)
}

func TestCreateContainer(t *testing.T) {
	a := "1"
	list := []string{"1", "2", "3", "4"}
	//list := []int64{1,2,3,4}
	container, err := CreateContainer(list, reflect.TypeOf(&a))
	fmt.Println(container, err)
	for _, x := range container {
		fmt.Printf("%p %v \n", x, x)
	}
}

type xp struct {
	name *string
}
