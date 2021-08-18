package rule

import (
	"errors"
	"fmt"
	"git.xiaojukeji.com/chenyeung/garg/check"
)

//处理 in 或者not in
type Express interface {
	Cal(actualVal interface{}) (bool, error)
}

type CollectionExpression struct {
	op    OperatorType
	elems []interface{}
}

func NewCollectionExpression() *CollectionExpression {
	return &CollectionExpression{}
}

func (c CollectionExpression) Cal(actualVal interface{}) (pass bool, err error) {
	switch c.op {
	case IN_OperatorType:
		pass = false
		for _, ele := range c.elems {
			if ele == actualVal {
				pass = true
				break
			}
		}
		if !pass {
			err = errors.New(fmt.Sprintf("expect %v in %v,but it's not contains", actualVal, c.elems))
		}
	case NI_OperatorType:
		pass = true
		for _, ele := range c.elems {
			if ele == actualVal {
				pass = false
				break
			}
		}
		if !pass {
			err = errors.New(fmt.Sprintf("expect %v not in %v,but it's in", actualVal, c.elems))
		}
	}
	return pass, err
}

type CalBucket struct {
	andBucket []Express
	orBucket  []Express
}

func NewCalBucket() *CalBucket {
	return &CalBucket{}
}

func (c CalBucket) Cal(actualVal interface{}) (pass bool, err error) {
	//result := garg.NewResult()
	//处理 &桶中express
	for _, exp := range c.andBucket {
		if pass, err := exp.Cal(actualVal); !pass {
			//result.Add(common.And_OPErrKey, err)
			return false, err
		}
	}
	//处理 ｜｜桶中的express
	for _, exp := range c.andBucket {
		if pass, err = exp.Cal(actualVal); pass {
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

func NewNormalExpression() *NormalExpression {
	return &NormalExpression{}
}

/*func (n NormalExpression) Cal(actual interface{}) (pass bool, err error) {
	atp := reflect.TypeOf(actual)
	av := reflect.ValueOf(actual)
	switch n.Op {
	case GE_OperatorType:
		switch atp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ev, err := strconv.ParseInt(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Int() >= ev
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			ev, err := strconv.ParseInt(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Int() >= ev
		case reflect.Float32, reflect.Float64:
			ev, err := strconv.ParseFloat(n.expected.(string), 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Float() >= ev
		case reflect.String:
			ev, ok := n.expected.(string)
			if !ok {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.String() >= ev
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + OperateType2String(n.Op))
		}
	case GT_OperatorType:
		switch atp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ev, err := strconv.ParseInt(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Int() > ev
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			ev, err := strconv.ParseInt(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Int() > ev
		case reflect.Float32, reflect.Float64:
			ev, err := strconv.ParseFloat(n.expected.(string), 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Float() > ev
		case reflect.String:
			ev, ok := n.expected.(string)
			if !ok {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.String() > ev
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + OperateType2String(n.Op))
		}
	case LT_OperatorType:
		switch atp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ev, err := strconv.ParseInt(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Int() < ev
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			ev, err := strconv.ParseInt(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Int() < ev
		case reflect.Float32, reflect.Float64:
			ev, err := strconv.ParseFloat(n.expected.(string), 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Float() < ev
		case reflect.String:
			ev, ok := n.expected.(string)
			if !ok {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.String() > ev
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + OperateType2String(n.Op))
		}
	case EQ_OperatorType:
		switch atp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ev, err := strconv.ParseInt(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Int() == ev
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			ev, err := strconv.ParseUint(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Uint() == ev
		case reflect.Float32, reflect.Float64:
			ev, err := strconv.ParseFloat(n.expected.(string), 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Float() == ev
		case reflect.String:
			ev, ok := n.expected.(string)
			if !ok {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.String() == ev
		case reflect.Bool:
			ev, err := strconv.ParseBool(n.expected.(string))
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Bool() == ev
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + OperateType2String(n.Op))
		}

	case LE_OperatorType:
		switch atp.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ev, err := strconv.ParseInt(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Int() <= ev
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			ev, err := strconv.ParseInt(n.expected.(string), 10, 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Int() <= ev
		case reflect.Float32, reflect.Float64:
			ev, err := strconv.ParseFloat(n.expected.(string), 64)
			if err != nil {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.Float() <= ev
		case reflect.String:
			ev, ok := n.expected.(string)
			if !ok {
				err = errors.New(fmt.Sprintf("类型转换错误, rule中的类型 与 实际字段值类型不一致"))
				break
			}
			pass = av.String() <= ev
		default:
			pass = false
			err = errors.New("不支持逻辑运算类型，op=" + OperateType2String(n.Op))
		}
	default:
		pass = false
		err = errors.New("暂不支持该操作符，或rule配置错误")
	}
	if !pass && err == nil {
		err = errors.New(fmt.Sprintf("expect %v %v, but act=%v, while the op is [%v]", OperateType2String(n.Op), n.expected, actual, OperateType2String(n.Op)))
	}
	return pass, err
}*/
func (n NormalExpression) Cal(actual interface{}) (pass bool, err error) {
	switch n.Op {
	case LE_OperatorType:
		return check.LE(actual, n.expected)
	case GE_OperatorType:
		return check.GE(actual, n.expected)
	case NE_OperatorType:
		return check.NE(actual, n.expected)
	case EQ_OperatorType:
		return check.E(actual, n.expected)
	case LT_OperatorType:
		return check.LT(actual, n.expected)
	case GT_OperatorType:
		return check.GT(actual, n.expected)
	case IN_OperatorType:
		return check.Contains(actual, n.expected)
	case NI_OperatorType:
		return check.NotContains(actual, n.expected)
	case NEED_OperatorType:
		return check.Required(actual)
	default:
		err = errors.New(OperateType2String(Illegal_OperatorType))
	}
	return false, err
}
