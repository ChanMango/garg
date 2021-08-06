package rule

import (
	"errors"
	"fmt"
	"git.xiaojukeji.com/chenyeung/garg/check"
	"reflect"
)

//处理 in 或者not in
type Express interface {
	Cal() (bool, error)
}

type CollectionExpression struct {
	op    OperatorType
	elems []interface{}
}

func (c CollectionExpression) Cal(target interface{}) (pass bool) {
	switch c.op {
	case IN_OperatorType:
		pass = false
		for _, ele := range c.elems {
			if ele == target {
				pass = true
				break
			}
		}
	case NI_OperatorType:
		pass = true
		for _, ele := range c.elems {
			if ele == target {
				pass = false
				break
			}
		}
	}
	return pass
}

type CalBucket struct {
	andBucket []Express
	orBucket  []Express
}

func (c CalBucket) Cal() (pass bool, err error) {
	//result := garg.NewResult()
	//处理 &桶中express
	for _, exp := range c.andBucket {
		if pass, err := exp.Cal(); !pass {
			//result.Add(common.And_OPErrKey, err)
			return false, err
		}
	}
	//处理 ｜｜桶中的express
	for _, exp := range c.andBucket {
		if pass, err = exp.Cal(); pass {
			return true, nil
		}
	}
	//todo 处理错误回传
	return false, err

}

type NormalExpression struct {
	Op       OperatorType
	expected interface{}
}

func (n NormalExpression) Cal(actual interface{}) (pass bool, err error) {
	etp := reflect.TypeOf(n.expected)
	atp := reflect.TypeOf(actual)
	ev := reflect.ValueOf(n.expected)
	av := reflect.ValueOf(actual)
	sameType, err := check.IsSameType(etp, atp)
	if !sameType {
		return false, err
	}
	switch n.Op {
	case GE_OperatorType:
		switch etp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			pass = av.Int() >= ev.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			pass = av.Int() >= ev.Int()
		case reflect.Float32, reflect.Float64:
			pass = av.Float() >= ev.Float()
		case reflect.String:
			pass = av.String() >= ev.String()
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + n.Op.String())
		}
	case NE_OperatorType:
		switch etp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			pass = ev.Int() != av.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			pass = ev.Uint() != av.Uint()
		case reflect.Float32, reflect.Float64:
			pass = ev.Float() != av.Float()
		case reflect.Complex64, reflect.Complex128:
			pass = ev.Complex() != av.Complex()
		case reflect.String:
			pass = ev.String() != av.String()
		case reflect.Bool:
			pass = ev.Bool() != av.Bool()
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + n.Op.String())
		}
	case GT_OperatorType:
		switch etp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			pass = av.Int() > ev.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			pass = av.Int() > ev.Int()
		case reflect.Float32, reflect.Float64:
			pass = av.Float() > ev.Float()
		case reflect.String:
			pass = av.String() > ev.String()
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + n.Op.String())
		}
	case LT_OperatorType:
		switch etp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			pass = av.Int() < ev.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			pass = av.Int() < ev.Int()
		case reflect.Float32, reflect.Float64:
			pass = av.Float() < ev.Float()
		case reflect.String:
			pass = av.String() < ev.String()
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + n.Op.String())
		}
	case EQ_OperatorType:
		switch etp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			pass = ev.Int() == av.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			pass = ev.Uint() == av.Uint()
		case reflect.Float32, reflect.Float64:
			pass = ev.Float() == av.Float()
		case reflect.Complex64, reflect.Complex128:
			pass = ev.Complex() == av.Complex()
		case reflect.String:
			pass = ev.String() == av.String()
		case reflect.Bool:
			pass = ev.Bool() == av.Bool()
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + n.Op.String())
		}

	case LE_OperatorType:
		switch etp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			pass = av.Int() <= ev.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			pass = av.Int() <= ev.Int()
		case reflect.Float32, reflect.Float64:
			pass = av.Float() <= ev.Float()
		case reflect.String:
			pass = av.String() <= ev.String()
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + n.Op.String())
		}
	default:
		pass = false
		err = errors.New("不支持逻辑运算类型，op=" + n.Op.String())
	}
	if !pass && err == nil {
		err = errors.New(fmt.Sprintf("expect= %v bug act= %v,while the op=%v", n.expected, actual, n.Op.String()))
	}
	return pass, err
}
